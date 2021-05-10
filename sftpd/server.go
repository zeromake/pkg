package main

import (
	"context"
	"errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var (
	ErrServerClosed = errors.New("ssh: Server closed")
)

type Server struct {
	Addr        string
	Version     string
	IdleTimeout time.Duration
	MaxTimeout  time.Duration
	Config      *ssh.ServerConfig

	listenerWg sync.WaitGroup
	mu         sync.RWMutex
	listener   net.Listener
	conn       map[*ssh.ServerConn]struct{}
	connWg     sync.WaitGroup
	doneChan   chan struct{}
}

func (srv *Server) closeListenersLocked() error {
	var err error
	err = srv.listener.Close()
	return err
}

func (srv *Server) getDoneChanLocked() chan struct{} {
	if srv.doneChan == nil {
		srv.doneChan = make(chan struct{})
	}
	return srv.doneChan
}

func (srv *Server) closeDoneChanLocked() {
	ch := srv.getDoneChanLocked()
	select {
	case <-ch:
	default:
		close(ch)
	}
}

func (srv *Server) Close() error {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	srv.closeDoneChanLocked()
	err := srv.closeListenersLocked()
	for c := range srv.conn {
		err = c.Close()
		if err != nil {
			return err
		}
		delete(srv.conn, c)
	}
	return err
}

func (srv *Server) Shutdown(ctx context.Context) error {
	srv.mu.Lock()
	lnerr := srv.closeListenersLocked()
	srv.closeDoneChanLocked()
	srv.mu.Unlock()

	finished := make(chan struct{}, 1)
	go func() {
		srv.listenerWg.Wait()
		srv.connWg.Wait()
		finished <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-finished:
		return lnerr
	}
}

func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":22"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv.listener = ln
	return srv.Serve(ln)
}

func (srv *Server) getDoneChan() <-chan struct{} {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	return srv.getDoneChanLocked()
}

func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	var tempDelay time.Duration
	for {
		conn, e := l.Accept()
		if e != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		go srv.HandleConn(conn)
	}
}

func (srv *Server) trackConn(c *ssh.ServerConn, add bool) {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	if srv.conn == nil {
		srv.conn = make(map[*ssh.ServerConn]struct{})
	}
	if add {
		srv.conn[c] = struct{}{}
		srv.connWg.Add(1)
	} else {
		delete(srv.conn, c)
		srv.connWg.Done()
	}
}

func (srv *Server) ChannelHandler(conn *ssh.ServerConn, newChan ssh.NewChannel) {
	ch, reqs, err := newChan.Accept()
	if err != nil {
		// TODO: trigger event callback
		return
	}
	for req := range reqs {
		switch req.Type {
		case "subsystem":
			_ = req.Reply(true, nil)
			server, err := sftp.NewServer(
				ch,
			)
			if err != nil {
				log.Fatal("sftp NewServer with error:", err)
				return
			}
			if err := server.Serve(); err == io.EOF {
				_ = server.Close()
				log.Print("sftp client exited session.")
			} else if err != nil {
				log.Fatal("sftp server completed with error:", err)
			}
		default:
			log.Println("not support:", req.Type)
			_ = req.Reply(false, []byte("not support"))
		}
	}
}

func (srv *Server) HandleConn(newConn net.Conn) {
	defer newConn.Close()
	sshConn, chans, reqs, err := ssh.NewServerConn(newConn, srv.Config)
	if err != nil {
		// TODO: trigger event callback
		return
	}
	srv.trackConn(sshConn, true)
	defer srv.trackConn(sshConn, false)

	go ssh.DiscardRequests(reqs)

	for ch := range chans {
		if ch.ChannelType() != "session" {
			log.Println("not support: ", ch.ChannelType())
		} else {
			go srv.ChannelHandler(sshConn, ch)
		}
	}
}

package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	privateBytes, err := ioutil.ReadFile(filepath.Join(homeDir, ".ssh/id_rsa"))
	if err != nil {
		log.Fatal("Failed to load private key", err)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key", err)
	}
	config := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			return nil, nil
		},
		PublicKeyCallback: func(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
			return nil, nil
		},
	}
	config.AddHostKey(private)
	srv := &Server{
		Addr:        ":2222",
		Version:     "1.0.1",
		IdleTimeout: time.Second * 10,
		MaxTimeout:  time.Second * 10,
		Config:      config,
	}
	log.Fatal(srv.ListenAndServe())
}

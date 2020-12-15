test:
	go test ./...

cover:
	mkdir -p coverage
	go test -covermode=count -coverprofile "coverage/coverage.cov" ./...
	go tool cover -func=coverage/coverage.cov
cover-html: cover
	go tool cover -html=coverage/coverage.cov -o coverage/coverage.html

bench:
	go test -bench=. -benchmem -run=none -benchtime=5s ./...

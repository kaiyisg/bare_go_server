install:
	go install github.com/kaiyisg/bare_go_server/cmd/changes

start:
	go install github.com/kaiyisg/bare_go_server/cmd/changes
	$(GOPATH)/bin/changes http

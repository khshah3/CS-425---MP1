all: serverbuild clientbuild

serverbuild:
	go build server/grepserver.go

clientbuild:
	go build dgrep.go

clean:
	rm grepserver dgrep

all: serverbuild clientbuild loggerbuild

serverbuild:
	go build server/grepserver.go

clientbuild:
	go build dgrep/dgrep.go

loggerbuild:
	go build logger/logger.go

clean:
	rm grepserver dgrep logger

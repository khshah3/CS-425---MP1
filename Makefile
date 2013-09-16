all: serverbuild clientbuild loggerbuild , testbuild

serverbuild:
	go build server/grepserver.go

clientbuild:
	go build dgrep/dgrep.go

loggerbuild:
	go build logger/logger.go

testbuild:
	go test -file dgrep_test.go -bench=".*"

clean:
	rm grepserver dgrep logger

all: serverbuild clientbuild testbuild

serverbuild:
	go build server/grepserver.go

clientbuild:
	go build dgrep.go

testbuild:
	go build logfilegenerator/loggen.go

clean:
	rm grepserver dgrep

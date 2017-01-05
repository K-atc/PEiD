
all: PEiD PEiD.exe anti_dbg_msgbox.exe
	go fmt

PEiD: *.go
	export GOPATH=`pwd`
	GOOS=linux go build

PEiD.exe: *.go
	GOOS=windows go build
	cp $@ "/home/katc/VirtualBox VMs/share/$@"

# FIXME: build always
anti_dbg_msgbox.exe: cmd/anti_dbg_msgbox/*.go
	cd cmd/anti_dbg_msgbox && GOOS=windows GOARCH=386 go build && rm -f anti_dbg_msgbox-upx.exe && upx -o anti_dbg_msgbox-upx.exe anti_dbg_msgbox.exe

init:
	go get github.com/Sirupsen/logrus
	go get github.com/mattn/go-colorable
	go get github.com/mitchellh/go-homedir
	go get github.com/fatih/color
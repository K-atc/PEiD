all:
	export GOPATH=`pwd`
	go build

anti_dbg_msgbox:
	cd cmd/anti_dbg_msgbox && GOOS=windows GOARCH=386 go build 

init:
	go get github.com/Sirupsen/logrus
	go get github.com/mattn/go-colorable
	go get github.com/mitchellh/go-homedir
	go get github.com/fatih/color
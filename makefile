
tag != git describe --tags --abbrev=8
rev != git rev-list -n 1 HEAD
ldflags = "-X 'j322.ica/gumroad-sammi/connect.Version=$(tag)' -X 'j322.ica/gumroad-sammi/connect.Revision=$(rev)'"

build: linux windows

linux: builddir
	GOOS=linux GOARCH=amd64 go build -v -o ./build/gummi.sh -ldflags $(ldflags) .

windows: builddir
	GOOS=windows GOARCH=amd64 go build -v -o ./build/gummi.exe -ldflags $(ldflags) .

builddir:
	mkdir -p ./build/

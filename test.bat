setlocal
set "tp=-gcflags=-trimpath=%GOPATH% -asmflags=-trimpath=%GOPATH%"
set "flags=-w -s"
echo build linux amd64 ...
set "GOOS=linux"
set "GOARCH=amd64"
go build -ldflags="%flags%" %tp% -o ./linux_amd64.exe


pause
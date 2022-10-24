set GOARCH=amd64
set GOOS=linux
go build -o ../out/myapp ../

set GOARCH=amd64
set GOOS=windows
go build -o ../out/myapp.exe ../
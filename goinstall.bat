set GOARCH=amd64
set GOOS=linux
go install
pause
set GOARCH=amd64
set GOOS=windows
go install
pause
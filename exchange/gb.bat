@ECHO OFF
go build .\taoexchange.go
del ..\bin\taoexchange.exe
move .\taoexchange.exe ..\bin\
cd ..\bin\
.\taoexchange.exe

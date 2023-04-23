@ECHO OFF
go build .\taogate.go
del ..\bin\taogate.exe
move .\taogate.exe ..\bin\
cd ..\bin\
.\taogate.exe

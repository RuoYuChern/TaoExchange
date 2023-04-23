@ECHO OFF
go build .\taocoordinator.go
del ..\bin\taocoordinator.exe
move .\taocoordinator.exe ..\bin\
cd ..\bin\
.\taocoordinator.exe

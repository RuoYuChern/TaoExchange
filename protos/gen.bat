@ECHO OFF
SET BIN_DIR="..\\3rdp\\protoc-osx\\bin"
SET CPP_I_DIR="..\\facade\\include"
SET CPP_SRC_DIR="..\\facade\\src"
SET protoc=%BIN_DIR%\\protoc.exe

%protoc% -I=./ --cpp_out=./ ./*.proto 
move *.h %CPP_I_DIR%
move *.cc %CPP_SRC_DIR%

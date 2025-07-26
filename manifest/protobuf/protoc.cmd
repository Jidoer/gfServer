@echo off
cd /d %~dp0

set PROTOC=protoc.exe

@REM if not exist "%PROTOC%" (
@REM     echo Not found command protoc!
@REM     echo Please install libprotobuf first!
@REM     exit /b 1
@REM )

set CPP_OUT_DIR=.
set CSHARP_OUT_DIR=..\protorpc_csharp

if not exist "%CPP_OUT_DIR%" mkdir "%CPP_OUT_DIR%"
if not exist "%CSHARP_OUT_DIR%" mkdir "%CSHARP_OUT_DIR%"

:: run protoc
echo Running protoc...
"%PROTOC%" --go_out="%CPP_OUT_DIR%" --csharp_out="%CSHARP_OUT_DIR%" *.proto

echo Done.

@ECHO OFF
set GOOS=linux
set GOARCH=arm64
set FIRMWARE_VERSION=V0.0.1
set CGO_ENABLED=0
echo BUILDING FOR %GOOS%-%GOARCH% 
go build -o ../dist/firmware-%GOOS%-%GOARCH%-%FIRMWARE_VERSION%
echo UPLOADING TO RPI
pscp -pw 3766 ../dist/firmware-%GOOS%-%GOARCH%-%FIRMWARE_VERSION% root@nasus:/home/simon
@REM set GOOS=windows
@REM set GOARCH=amd64
@REM echo BUILDING FOR %GOOS%-%GOARCH% 
@REM go build -o ../dist/firmware-%GOOS%-%GOARCH%-%FIRMWARE_VERSION%.exe
echo Done.
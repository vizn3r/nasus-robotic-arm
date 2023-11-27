del "./firmware.exe" | del "./server.exe" | del "./config.json"

go build -C ../firmware -o ../dist/firmware.exe
go build -C ../server -o ../dist/server.exe
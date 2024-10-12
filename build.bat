set GOOS=darwin
set GOARCH=amd64
garble build -o cyjs main.go
@REM garble -literals -tiny build -o cyjs.exe -ldflags="-s -w" -trimpath main.go

@REM go build -ldflags="-s -w" -o main.wasm main.go
# GO Api Stack

### Compile for Windows
```
GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe gin/main.go
```

### Compile for Linux
# 64-bit
```
$ GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux gin/main.go
```

# 32-bit
```
$ GOOS=linux GOARCH=386 go build -o bin/app-386-linux gin/main.go
```
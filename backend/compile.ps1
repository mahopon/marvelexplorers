Set-Variable CGO_ENABLED=0
Set-Variable GOOS=linux
Set-Variable GOARCH=amd64
go build -o marvelexplorers ./main/main.go
swag init -g src/main/main.go
mkdir -p dist
go build -o dist/demo.exe src/main/main.go
cp -r docs dist/docs
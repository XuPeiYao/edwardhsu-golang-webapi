./tools/swag init --parseDependency -g ./cmd/main.go
rm -f -r ./assets/swagger-docs
mv ./docs ./swagger-docs
mv ./swagger-docs ./assets
rm -f ./assets/swagger-docs/docs.go

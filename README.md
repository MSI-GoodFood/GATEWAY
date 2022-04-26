# Main gateway + auth

Swagger à disposition sur /v1/swagger/index.html

### LAUNCH
swag init; go run .;

### UPDATE GIT
git add -A; git commit -m "update"; git push origin master;

### UPDATE HEROKU
docker buildx build --platform linux/amd64 -t main .; docker tag main registry.heroku.com/goodfood-gateway/web; docker push registry.heroku.com/goodfood-gateway/web; heroku container:release web -a goodfood-gateway;

### BUILD DOCKER COMPOSE
docker-compose-v1 down --volumes; docker-compose-v1 up --build;

### UPDATE GRPC
protoc --go_out=proto --go-grpc_out=proto proto/service.proto

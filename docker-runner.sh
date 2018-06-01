docker build -t hullo .
docker run -d --publish 8080:8080 --name hullo --rm hullo

#to stop docker stop hullo
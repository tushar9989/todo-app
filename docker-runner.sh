docker build -t todo-app .
docker run -d --publish 8080:8080 --name todo-app --rm todo-app

#to stop docker stop todo-app
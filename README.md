# Shopping-api

1. linux command in order to go build  on MacOS :

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o development/shopping-app 


2. make docker container:

docker build --rm -t "59625962/shopping-app:v1" -f development/Dockerfile .

3. push it on docker hub:

docker push 59625962/shopping-app:v1

4 .  in order to test before push to docker hub

docker run 59625962/shopping-app:v1

5.
zkubectl login playground

6. 
zkubectl apply -f service.yaml

7. 
zkubectl apply -f ingress.yaml

8.
zkubectl apply -f postgres.yaml

9.
zkubectl apply -f deployment.yaml

10. to check the :

   1. zkubectl get logs

   2. zkubectl get pods | grep mentoring-shopping

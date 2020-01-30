
docker-build:
	docker build --no-cache -t us.gcr.io/mchirico/nginxpig:latest -f Dockerfile .

push:
	docker push us.gcr.io/mchirico/nginxpig:latest

dhparam:
	openssl dhparam 2048 > dhparam.pem

dhparam4096:
	openssl dhparam 4096 > dhparam.pem

build:
	go build -v .

run:
	docker run --rm -it -p 8080:8080  us.gcr.io/mchirico/nginxpig:latest

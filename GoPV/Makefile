
docker-build:
	docker build --no-cache -t us.gcr.io/mchirico/gopigstorage:latest -f Dockerfile .

push:
	docker push us.gcr.io/mchirico/gopigstorage:latest


run:
	docker run --rm -it -p 2342:2342  us.gcr.io/mchirico/gopigstorage:latest

FROM golang:1.10-alpine
ADD . /go/src/hello-app
RUN go install hello-app

FROM alpine:latest

RUN apk	update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=0 /go/bin/hello-app .
ENV PORT 2342
CMD ["./hello-app"]
# To run this
#  docker build -t us.gcr.io/mchirico/gopigstorage:latest .

#  docker push us.gcr.io/mchirico/gopigstorage:latest

# kubectl run hello-server --image us.gcr.io/mchirico/hello:latest --port 8080
# kubectl delete service hello-server

# // Not working ..
# kubectl rolling-update myapp --image=us.gcr.io/mchirico/hello:latest --image-pull-policy Always



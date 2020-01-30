FROM golang:1.10-alpine
ADD . /go/src/hello-app
RUN go install hello-app

FROM alpine:latest
COPY --from=0 /go/bin/hello-app .
ENV PORT 2341
CMD ["./hello-app"]
# To run this
#  docker build -t us.gcr.io/mchirico/gopig:latest .

#  docker push us.gcr.io/mchirico/gopig:latest

# kubectl run hello-server --image us.gcr.io/mchirico/hello:latest --port 8080
# kubectl delete service hello-server

# // Not working ..
# kubectl rolling-update myapp --image=us.gcr.io/mchirico/hello:latest --image-pull-policy Always



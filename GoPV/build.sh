#!/bin/bash
docker build -t us.gcr.io/mchirico/gopigstorage:latest .
docker push us.gcr.io/mchirico/gopigstorage:latest


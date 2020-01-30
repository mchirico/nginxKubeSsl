#!/bin/bash
kubectl apply -f existing-pd.yaml
kubectl apply -f gopigstorage-deployment.yaml
kubectl apply -f gopigstorage-service.yaml


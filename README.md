# Giles

API Service for a backend MongoDB Database... details are purposely being left vague...

## Kubernetes Deployment

Create a giles-secrets.yaml file, which will contain the username/password URI for MongoDB host/cluster

An example is included in the repo, but essentially is something like this:

    apiVersion: v1
    kind: Secret
    metadata:
      name: giles-v1
      type: Opaque
    data:
      mongo_uri: <base64 encoded mongo uri> 
      port: <base64 encoded port>

Encoding something as base64 on Mac/Linux you would echo the string piped to the command base64

    echo 'mongodb://mongo:mongo@mongohost:27017/test' | base64

## Installing

Installing the two yaml's into kubernetes is then done with the following command:

    kubectl create -f giles-secrets.yaml
    kubectl create -f giles-deployment.yaml



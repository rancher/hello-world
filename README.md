Hello world
===========

This image runs hello-world web service in 80 port used for demoing and/or testing. It shows data about hostname, k8s services and request headers. 

## Building from Source

The binaries will be located in `/bin` for linux and `build/bin` for cross compiling.

### Linux Binary

Run `make`.

### Mac & Windows Binaries

Run `CROSS=1 make build`. 

## Building Docker Image

To build `rancher/hello-world`, run `make`.  To use a custom Docker repository, do `REPO=custom make`, which produces a `custom/hello-world` image.

## Running Docker Image

### Docker

Run `docker run -td -p <PORT>:80 rancher/hello-world`.

### K8s

Deployment manifest
```
apiVersion: apps/v1beta2
kind: Deployment
metadata:
  labels:
    app: hello-world
  name: hello-world
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-world
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
    type: RollingUpdate
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - image: rancher/hello-world
        imagePullPolicy: Always
        name: hello-world
        ports:
        - containerPort: 80
          protocol: TCP
      restartPolicy: Always
---
apiVersion: v1
kind: Service
metadata:
  name: hello-world
  namespace: default
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    app: hello-world
```

Run `kubectl apply -f <DEPLOY_MANIFEST>`

## Contact

For bugs, questions, comments, corrections, suggestions, etc., open an issue in
[rancher/rancher](//github.com/rancher/rancher/issues) with a title prefix of `[hello-world] `.

Or just [click here](//github.com/rancher/rancher/issues/new?title=%5Bhello-world%5D%20) to create a new issue.

## License
Copyright (c) 2014-2018 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

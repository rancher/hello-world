Rancher Hello-world
===================

This image runs a web service in 8080 port, showing k8s services and request headers, used for testing. 

## Build

```
docker build -t rancher/hello-world:<version> .
```

## Versions

- `latest` [(Dockerfile)](https://github.com/rawmind0/web-test/blob/master/Dockerfile)


## Usage

```
docker run rancher/hello-world:<version> 
```

## Building from Source

The binaries will be located in `/bin`.

### Linux Binary

Run `make`.

### Mac Binary

Run `CROSS=1 make build`.

## Docker Image

Run `docker run --rm -it -v <PATH_TO_CONFIG>:/root/.rancher/cli2.json rancher/cli [ARGS]`.
Pass credentials by replacing `<PATH_TO_CONFIG>` with your config file for the server.

To build `rancher/cli`, run `make`.  To use a custom Docker repository, do `REPO=custom make`, which produces a `custom/cli` image.

## Contact

For bugs, questions, comments, corrections, suggestions, etc., open an issue in
[rancher/rancher](//github.com/rancher/rancher/issues) with a title prefix of `[cli] `.

Or just [click here](//github.com/rancher/rancher/issues/new?title=%5Bcli%5D%20) to create a new issue.

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

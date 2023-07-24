FROM golang AS build
COPY . /go/src/hello-world
WORKDIR /go/src/hello-world
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world .

FROM rawmind/alpine-base:3.7-0
COPY --from=build /go/src/hello-world/hello-world /opt/hello-world/
COPY img/* /opt/hello-world/
WORKDIR /opt/hello-world
ENTRYPOINT ["/opt/hello-world/hello-world"]

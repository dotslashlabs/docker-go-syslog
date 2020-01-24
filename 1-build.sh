go get "gopkg.in/mcuadros/go-syslog.v2"

# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
docker build -t go-syslog -f Dockerfile.scratch .

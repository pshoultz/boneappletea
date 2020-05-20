from golang:alpine AS builder
RUN apk update && apk add --no-cache git
ADD . /go/src/github.com/
WORKDIR /go/src/github.com/
RUN git clone https://github.com/pshoultz/boneappletea.git
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get go.mongodb.org/mongo-driver/mongo
WORKDIR /go/src/github.com/boneappletea/
#RUN go run main.go
RUN go build -o main.go
CMD ["./main.go"]
EXPOSE 8080

# builder container
FROM golang AS builder
WORKDIR /go/src/github.com/
COPY . /go/src/github.com/boneappletea
RUN go get -d -v github.com/gin-contrib/cors \
    && go get -d -v github.com/gin-gonic/gin \
    && go get -d -v go.mongodb.org/mongo-driver/bson \
    && go get -d -v go.mongodb.org/mongo-driver/mongo \
    && go get -d -v go.mongodb.org/mongo-driver/mongo/options
WORKDIR /go/src/github.com/boneappletea
RUN CGO_ENABLED=0 GOOS=linux go build -v -o bat .

# use small image for the final container to run our compiled app
FROM alpine:latest
# RUN apk update && apk add bash <--- uncomment to add bash shell for troubleshooting
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/boneappletea/bat .
CMD ["./bat"]
EXPOSE 8080

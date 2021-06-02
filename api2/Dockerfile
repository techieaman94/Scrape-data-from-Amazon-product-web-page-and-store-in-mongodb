FROM golang:1.12.0-alpine3.9
RUN apk add git
RUN go get -u github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo

COPY . .

RUN go build -o main .

CMD ["./main"]

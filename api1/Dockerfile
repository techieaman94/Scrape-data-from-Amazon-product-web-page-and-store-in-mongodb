FROM golang:1.12.0-alpine3.9
RUN apk add git
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gocolly/colly

COPY . .

RUN go build -o main .

CMD ["./main"]

FROM golang:latest      

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY *.go ./

RUN go build -o /go-http 

EXPOSE 8080

CMD [ "/go-http" ]

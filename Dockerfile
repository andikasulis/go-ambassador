FROM golang:1.16

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
#RUN apk update

#RUN apk add nano

#RUN go get -u github.com/cosmtrek/air

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]
#RUN go get github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
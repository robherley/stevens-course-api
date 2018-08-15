FROM golang

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/robherley/stevens-course-api

COPY . .
RUN dep ensure

RUN go build -o server
EXPOSE 8080

ENTRYPOINT ["./server"]
FROM golang:1.12

RUN apt-get update

RUN go get -u github.com/alecthomas/gometalinter
RUN gometalinter --install

RUN mkdir -p /go/src/github.com/gunjan01/searcher

COPY . /go/src/github.com/gunjan01/searcher
WORKDIR /go/src/github.com/gunjan01/searcher

RUN go install ./source/cmd/...

EXPOSE 9000

CMD ["true"]

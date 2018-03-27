FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

COPY temp/en.json /usr/src/book/fr/en/harry-potter-1/chapter-1.json
COPY temp/fr.json /usr/src/book/fr/fr/harry-potter-1/chapter-1.json
 
EXPOSE 8080
CMD ["app"]

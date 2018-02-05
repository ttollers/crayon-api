FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

COPY temp/translated_final.json /usr/src/book/harry-potter-1/fr/chapter-1/translated_final.json
COPY temp/native_final.json /usr/src/book/harry-potter-1/fr/chapter-1/native_final.json
 
EXPOSE 8080
CMD ["app"]

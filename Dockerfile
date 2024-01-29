FROM golang:1.21

ENV CGO_ENABLED=1
# Move to our project folder and run the program
ADD . /web
WORKDIR /web
RUN go build -o fo-web

EXPOSE 8180

ENTRYPOINT [ "./fo-web" ]



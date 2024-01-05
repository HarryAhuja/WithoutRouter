FROM golang:1.22rc1-alpine3.18
WORKDIR /test
COPY . /test
RUN go build /test
EXPOSE 8000
ENTRYPOINT [ "./httpserver" ]
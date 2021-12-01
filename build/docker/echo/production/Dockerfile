FROM golang:1.17-alpine as build

RUN mkdir - p /app/

WORKDIR /app/

COPY . /app/

RUN go build main.go

FROM golang:1.17-alpine

COPY --from=build /app/main /app/main

WORKDIR /app/

EXPOSE 1323

ENTRYPOINT ["./main"]
FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .

RUN go build -o generateTicket .


FROM alpine:3.16
RUN apk update
WORKDIR /app

COPY --from=build /go/src/app/ .

EXPOSE 8080
CMD [ "./generateTicket" ]
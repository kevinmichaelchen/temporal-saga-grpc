FROM golang:1.19 as build

ENV GO111MODULE=on
WORKDIR /app/server
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build main.go

FROM alpine:latest as license-svc
WORKDIR /app/server
COPY --from=build /app/server/main ./
RUN chmod +x ./main
CMD [ "./main" ]
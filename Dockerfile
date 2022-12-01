FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o ./microservice-pot microservice-pot


FROM golang:1.19-alpine
COPY --from=build /app/microservice-pot /microservice-pot
EXPOSE 8080
CMD ["/microservice-pot"]


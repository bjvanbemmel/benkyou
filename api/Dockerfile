FROM golang:1.23.5-alpine3.20 AS build
WORKDIR /api
COPY ./ ./
RUN go mod download
RUN go build -o /bin/app ./cmd/api/main.go

FROM golang:1.23.5-alpine3.20 AS debug
WORKDIR /api
COPY ./ ./
RUN apk add git
RUN go mod download
RUN go install github.com/air-verse/air@v1.61.7
CMD [ "air", "-c", ".air.toml" ]

FROM scratch AS release
COPY --from=build /bin/app /bin/app
CMD [ "/bin/app" ]

FROM golang:1.21 AS build-stage
WORKDIR /service
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=arm64
RUN go build -o backend ./ 

FROM alpine:latest AS prod-stage
WORKDIR /service
COPY --from=build-stage /service/backend /service
CMD [ "/service/backend" ]

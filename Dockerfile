FROM golang:1.20-alpine

ENV GOPATH=/

WORKDIR /avito

COPY go.mod go.sum ./
COPY ./ ./

RUN go mod download
RUN go build -o /avito/segmentation ./cmd/main.go

EXPOSE 8000

CMD ["/avito/segmentation"]
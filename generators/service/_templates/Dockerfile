# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o {{.Package}} cmd/{{.Package}}/main.go

# final stage
FROM scratch
COPY --from=builder /app/{{.Package}} /app/
EXPOSE {{.GrpcPort}} {{if .EnableHttp}}{{.HttpPort}}{{end}}
ENTRYPOINT ["/app/{{.Package}}"]
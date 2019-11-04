# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o {{ .service.Package }} cmd/{{ .service.Package }}/main.go

# final stage
FROM scratch
COPY --from=builder /app/{{ .service.Package }} /app/
EXPOSE {{ .service.GrpcPort }} {{if .service.EnableHttp }}{{ .service.HttpPort }}{{end}}
ENTRYPOINT ["/app/{{ .service.Package }}"]
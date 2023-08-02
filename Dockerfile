FROM jac18281828/godev:latest

ARG PROJECT=gocrucible
WORKDIR /workspaces/${PROJECT}
ENV GOMAXPROCS=10
COPY main.go .
RUN chown -R jac:jac .
USER jac
RUN go mod init github.com/jac18281828/gocrucible
RUN go build
RUN go test ./...


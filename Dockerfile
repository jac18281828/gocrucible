FROM jac18281828/godev:latest

ARG PROJECT=gocrucible
WORKDIR /workspaces/${PROJECT}
ENV GOMAXPROCS=10
COPY . .
RUN chown -R jac:jac .
USER jac
RUN go build
RUN go test ./...


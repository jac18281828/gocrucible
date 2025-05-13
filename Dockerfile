FROM jac18281828/godev:latest

ARG PROJECT=gocrucible
WORKDIR /workspaces/${PROJECT}
ENV GOMAXPROCS=10
COPY . .
RUN chown -R godev:godev .
USER godev
ENV GOPATH=/workspaces/${PROJECT}

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /workspaces/${PROJECT}/src/crucible

RUN go build
RUN go test ./...


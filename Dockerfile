FROM rvolosatovs/protoc

ENV PATH /usr/local/go/bin:$PATH
ENV GOBIN=/usr/local/bin

RUN apk update && apk add --no-cache wget tar git
RUN wget -O go.tgz https://go.dev/dl/go1.22.3.linux-amd64.tar.gz && tar -C /usr/local -xzf go.tgz
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN GO111MODULE=on go install github.com/bufbuild/buf/cmd/buf@v1.32.1
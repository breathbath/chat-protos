package service

import (
	gen "github.com/breathbath/chat-protos/gen/protos/v1"
)

type ExampleService struct {
	*gen.UnimplementedChatServiceServer
}

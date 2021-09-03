package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
)

type JwtPayload struct {
	Exp  int64
	Iat  int64
	Sub  string
	Id   int64
	Name string
}

func GetJwtPayloadFromCtx(ctx context.Context, opts ...string) (*JwtPayload, error) {
	key := "x-jwt-payload"
	if len(opts) > 0 {
		key = opts[0]
	}
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, errors.New("get jwt payload from context error")
	}
	payloadStr := tr.RequestHeader().Get(key)
	var payload JwtPayload
	if err := json.Unmarshal([]byte(payloadStr), &payload); err != nil {
		return nil, fmt.Errorf("unmarshal payload error: %s", err.Error())
	}
	return &payload, nil
}

package security

import (
	"context"
	"encoding/base64"
)

type BasicAuth struct {
	Username string
	Password string
}

func (basicAuth BasicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := basicAuth.Username + ":" + basicAuth.Password

	enc := base64.StdEncoding.EncodeToString([]byte(auth))

	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

func (BasicAuth) RequireTransportSecurity() bool {
	return true
}

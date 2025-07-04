package pds

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gander-social/gander-indigo-sovereign/xrpc"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func makeToken(subject string, scope string, exp time.Time) jwt.Token {
	tok := jwt.New()
	tok.Set("scope", scope)
	tok.Set("sub", subject)
	tok.Set("iat", time.Now().Unix())
	tok.Set("exp", exp.Unix())

	return tok
}

func (s *Server) createAuthTokenForUser(ctx context.Context, handle, did string) (*xrpc.AuthInfo, error) {
	accessTok := makeToken(did, "com.atproto.access", time.Now().Add(24*time.Hour))
	refreshTok := makeToken(did, "com.atproto.refresh", time.Now().Add(7*24*time.Hour))

	rval := make([]byte, 10)
	rand.Read(rval)
	refreshTok.Set("jti", base64.StdEncoding.EncodeToString(rval))

	accSig, err := jwt.Sign(accessTok, jwt.WithKey(jwa.HS256, s.jwtSigningKey))
	if err != nil {
		return nil, fmt.Errorf("signing access token: %w", err)
	}

	refSig, err := jwt.Sign(refreshTok, jwt.WithKey(jwa.HS256, s.jwtSigningKey))
	if err != nil {
		return nil, fmt.Errorf("signing refresh token: %w", err)
	}

	return &xrpc.AuthInfo{
		AccessJwt:  string(accSig),
		RefreshJwt: string(refSig),
		Handle:     handle,
		Did:        did,
	}, nil
}

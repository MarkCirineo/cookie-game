package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/markcirineo/cookie-game/internal/conf"
	"github.com/markcirineo/cookie-game/internal/store"

	"github.com/cristalhq/jwt/v3"
	"github.com/rs/zerolog/log"
)

var (
	jwtSigner   jwt.Signer
	jwtVerifier jwt.Verifier
)

func jwtSetup(conf conf.Config) {
	var err error
	key := []byte(conf.JwtSecret)

	jwtSigner, err = jwt.NewSignerHS(jwt.HS256, key)
	if err != nil {
		log.Panic().Err(err).Msg("error creating JWT signer")
	}

	jwtVerifier, err = jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		log.Panic().Err(err).Msg("error creating JWT verifier")
	}
}

func generateJWT(user *store.User) string {
	claims := &jwt.RegisteredClaims{
		ID: fmt.Sprint(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	}
	builder := jwt.NewBuilder(jwtSigner)
	token, err := builder.Build(claims)
	if err != nil {
		log.Panic().Err(err).Msg("error building JWT")
	}
	return token.String()
}

func verifyJWT(tokenStr string) (int, error) {
	token, err := jwt.Parse([]byte(tokenStr))
	if err != nil {
		log.Error().Err(err).Str("tokenStr", tokenStr).Msg("error parsing JWT")
		return 0, err
	}

	if err := jwtVerifier.Verify(token.Payload(), token.Signature()); err != nil {
		log.Error().Err(err).Msg("error verifying token")
		return 0, err
	}

	var claims jwt.StandardClaims
	if err := json.Unmarshal(token.RawClaims(), &claims); err != nil {
		log.Error().Err(err).Msg("error unmarshalling JWT claims")
		return 0, err
	}

	if notExpired := claims.IsValidAt(time.Now()); !notExpired {
		return 0, errors.New("token expired")
	}

	id, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Error().Err(err).Str("claims.ID", claims.ID).Msg("error converting claims ID to number")
		return 0, errors.New("ID in token is not valid")
	}
	return id, err
}
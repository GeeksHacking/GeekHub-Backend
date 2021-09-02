package jwt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	auth0middleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/geekshacking/geekhub-backend/config"
	"github.com/go-chi/render"
	"io"
	"net/http"
	"time"
)

type middleware struct {
	config config.Config
	JWKs   JWKs
}

type JWKs struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func NewMiddleware(config config.Config) (*middleware, error) {
	netClient := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := netClient.Get(config.JWKsURL)
	if err != nil {
		return nil, fmt.Errorf("unable to get PEM cert: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	jwks := JWKs{}
	err = json.NewDecoder(res.Body).Decode(&jwks)
	if err != nil {
		return nil, fmt.Errorf("unable to decode JWKs: %w", err)
	}

	return &middleware{config, jwks}, nil
}

func (m *middleware) PEMCertificate(token *jwt.Token) (string, error) {
	cert := ""

	for idx, _ := range m.JWKs.Keys {
		if token.Header["kid"] == m.JWKs.Keys[idx].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + m.JWKs.Keys[idx].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, fmt.Errorf("unable to find the appropriate key")
	}

	return cert, nil
}

func (m *middleware) Auth0() *auth0middleware.JWTMiddleware {
	return auth0middleware.New(auth0middleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//aud := serverConfig.Audience
			//if ok := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false); !ok {
			//	return token, errors.New("invalid audience")
			//}

			if ok := token.Claims.(jwt.MapClaims).VerifyIssuer(m.config.Domain, false); !ok {
				return token, errors.New("invalid issuer")
			}

			cert, err := m.PEMCertificate(token)
			if err != nil {
				return token, fmt.Errorf("unable to get PEM cert: %w", err)
			}

			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return token, fmt.Errorf("unable to parse RSA public key from PEM: %w", err)
			}

			return result, nil
		},
		SigningMethod: jwt.SigningMethodRS256,
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			render.DefaultResponder(w, r, render.M{"error": "unauthorized"})
		},
	})
}

func (m *middleware) User(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)
		if !ok {
			next.ServeHTTP(w, r)
		}

		ctx := context.WithValue(r.Context(), "userID", claims["sub"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
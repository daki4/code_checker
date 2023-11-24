package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtService struct {
	SecretKey       string
	Issuer          string
	ValidityMinutes int64
	blacklist       []uuid.UUID
}

var generator JwtService

func Init() {
	generator = JwtService{
		os.Getenv("JWT_KEY"),
		os.Getenv("JWT_ISSUER"),
		60,
		[]uuid.UUID{},
	}
}

func GetJwtConfig() *JwtService {
	return &generator
}

func (cfg JwtService) GetBlacklist() []uuid.UUID {
	return cfg.blacklist
}

func (cfg JwtService) CreateToken(username string) (string, error) {
	claim := Claim{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(cfg.ValidityMinutes) * 3600)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Audience:  []string{cfg.Issuer},
			Issuer:    cfg.Issuer,
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (cfg JwtService) RefreshToken(token string) (string, error) {
	tok, claim, err := cfg.decodeClaim(token)
	if !cfg.validateToken(tok, claim) || err != nil {
		return "", nil
	}
	return cfg.CreateToken(claim.Username)
}

func (cfg *JwtService) BlacklistToken(token string) bool {
	_, claim, _ := cfg.decodeClaim(token)
	id, _ := uuid.Parse(claim.ID)
	cfg.blacklist = append(cfg.blacklist, id)
	return true
}

func (cfg JwtService) validateToken(token jwt.Token, claim Claim) bool {
	// Check if the token is valid
	if !token.Valid {
		return false
	}
	if claim.Issuer != cfg.Issuer {
		return false
	}
	if claim.ExpiresAt.Time.Before(time.Now()) {
		return false
	}
	return true
}

func (cfg JwtService) IsBlacklisted(token string) bool {
	_, claim, _ := cfg.decodeClaim(token)
	id, _ := uuid.Parse(claim.ID)
	for _, tok := range cfg.blacklist {
		if tok == id {
			return true
		}
	}
	return false
}

func (cfg JwtService) ValidateToken(token string) bool {
	tok, claim, err := cfg.decodeClaim(token)
	if err != nil {
		return false
	}
	if !cfg.validateToken(tok, claim) {
		return false
	}
	if cfg.IsBlacklisted(token) {
		return false
	}
	return true
}

func (cfg JwtService) GetClaim(token string) (Claim, error) {
	_, claim, err := cfg.decodeClaim(token)
	if err != nil {
		return Claim{}, err
	}
	return claim, nil
}

func (cfg JwtService) decodeClaim(token string) (jwt.Token, Claim, error) {
	var claim Claim
	tok, err := jwt.NewParser().ParseWithClaims(token, &claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})
	if err != nil {
		return jwt.Token{}, Claim{}, err
	}

	return *tok, claim, nil
}

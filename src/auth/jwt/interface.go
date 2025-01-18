package authJwt

type JwtServiceInterface interface {
	GenerateToken(userId string) (string, error)

	// TODO
	// ValidateToken(encodedToken string) (*jwt.Token, error)
}

package config

func GetSecretKey() string {
	return getEnv("JWT_SECRET_KEY", "default")
}

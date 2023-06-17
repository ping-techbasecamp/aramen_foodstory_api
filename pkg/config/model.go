package config

type AppConfig struct {
	Env       General
	Database  DatabaseConfig
	JWT       JWTConfig
	Mailbit   MailbitConfig
	Foodstory FoodstoryConfig
}

type General struct {
	Host             string `envconfig:"APP_HOST" default:""`
	Port             string `envconfig:"APP_PORT"`
	GinMode          string `envconfig:"GIN_MODE"`
	EncryptedKey     string `envconfig:"ENCRYPTED_KEY"`
	EnabledDevRoutes bool   `envconfig:"ENABLED_DEV_ROUTES" default:"false"`
}

type DatabaseConfig struct {
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type JWTConfig struct {
	SignedKey      string `envconfig:"JWT_SIGNING_KEY" required:"true"`
	ExpiredMinutes int    `envconfig:"JWT_EXPIRED_MINUTES" required:"true"`
}

type MailbitConfig struct {
	Username   string `envconfig:"MAILBIT_USERNAME" required:"true"`
	Password   string `envconfig:"MAILBIT_PASSWORD" required:"true"`
	SenderName string `envconfig:"MAILBIT_SENDER" required:"true"`
}

type FoodstoryConfig struct {
	FoodstoryAPI string `envconfig:"FOODSTORY_API_KEY" required:"true"`
	FoodstoryAES string `envconfig:"FOODSTORY_AES_KEY" required:"true"`
}

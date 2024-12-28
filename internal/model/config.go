package model

type Config struct {
	Token     string `envconfig:"TOKEN_TELEGRAM" required:"true"`
	DBHost    string `envconfig:"DBHOST" required:"false"`
	DBPort    string `envconfig:"DBPORT" default:"5432"`
	DBName    string `envconfig:"DBNAME" required:"true"`
	DBUser    string `envconfig:"DBUSER" required:"true"`
	DBPass    string `envconfig:"DBPASS" required:"true"`
	DBSSLMode string `envconfig:"DBSSLMODE" default:"false"`
}

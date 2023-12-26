package config

// AppConfig represents the configuration for the application.
// UseCache indicates whether to use cache or not.
// Port specifies the port number the application will listen on.
// InProduction indicates whether the application is in production mode or not.
type AppConfig struct {
	UseCache     bool
	Port         int
	InProduction bool
}

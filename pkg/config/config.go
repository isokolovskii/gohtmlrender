package config

// AppConfig represents the configuration for the application.
// UseCache indicates whether to use cache or not.
// Port specifies the port number the application will listen on.
type AppConfig struct {
	UseCache bool
	Port     int
}

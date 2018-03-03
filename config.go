package apixu

// Config is used to set up the Apixu service instance with the following:
// 	Version: API version
// 	Format: API response format (JSON or XML)
// 	ApiKey: API key generated on Apixu website
type Config struct {
	Version string
	Format  string
	APIKey  string
}

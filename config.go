package apixu

// Config is used to set up the Apixu service instance with the following:
// 	Format: API response format (JSON or XML)
// 	ApiKey: API key generated on Apixu website
type Config struct {
	Format string
	APIKey string
}

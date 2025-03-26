package apis

type ServiceConfig struct {
	ListenAddress string
	Port          string
	Trace         bool
	AllowOrigins  []string
	AllowHeaders  []string
	ExposeHeaders []string
	AllowMethods  []string
}

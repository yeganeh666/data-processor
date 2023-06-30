package configs

type Configs struct {
	HttpPort     uint16 `mapstructure:"http_port"`
	HttpHost     string `mapstructure:"http_host"`
	QuotaAddress string `mapstructure:"quota_address"`
	Version      string `mapstructure:"version"`
	ServiceName  string `mapstructure:"service_name"`
}

var (
	DefaultConfig = []byte(`
service_name: "gateway"
version: "v1.0.0"
http_port: 8080
http_host: "localhost"
quota_address: ":26000"
`)
)

package configs

type Configs struct {
	ServiceName string          `mapstructure:"service_name"`
	Version     string          `mapstructure:"version"`
	GRPCPort    uint16          `mapstructure:"grpc_port"`
	Postgres    SectionPostgres `yaml:"postgres"`
}
type SectionPostgres struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

var (
	DefaultConfig = []byte(`
service_name: "quota"
version: "v1.0.0"
grpc_port: 26000
postgres:
  host: "localhost"
  port: 5432
  db: "quota"
  user: "postgres"
  pass: "postgres"
`)
)

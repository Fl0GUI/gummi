package config

type Config struct {
	SammiConfig
	GumroadConfig
}

type SammiConfig struct {
	Host     string
	Port     string
	Password string
	ButtonId string
}

type GumroadConfig struct {
	AccessToken string
	ServerPort  string
	PublicIp    string
}

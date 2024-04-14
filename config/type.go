package config

type Configuration struct {
	Advanced             Advanced
	SammiConfig          SammiConfig
	GumroadConfig        GumroadConfig
	FourthWallConfig     FourthWallConfig
	StreamElementsConfig StreamElementsConfig
}

// Has defaults
type Advanced struct {
	ServerConfig ServerConfig
	BufferSize   int
}

type ServerConfig struct {
	ServerPort string
	PublicIp   string
}

type SammiConfig struct {
	Host     string
	Port     string
	Password string
}

type GumroadConfig struct {
	AccessToken string
	ButtonId    string
}

type FourthWallConfig struct {
	AccessToken string
	ButtonId    string
}

type StreamElementsConfig struct {
	AccessToken string
	ButtonId    string
}

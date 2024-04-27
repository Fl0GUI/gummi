package config

type Configuration struct {
	Advanced         Advanced
	SammiConfig      SammiConfig
	GumroadConfig    GumroadConfig
	FourthWallConfig FourthWallConfig
}

// Has defaults
type Advanced struct {
	ServerConfig    ServerConfig
	BufferSize      int
	BackoffAttempts int
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
	Active      bool
	AccessToken string
}

type FourthWallConfig struct {
	Active      bool
	AccessToken string
}

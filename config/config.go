package config

type Config struct {
	Server *Server
	DB     *DB
}

type Server struct {
	Addr string
}

type DB struct {
	Driver string
	Source string
}

func New() *Config {
	// TODO Config init from file
	return &Config{
		Server: &Server{
			Addr: ":8080",
		},
		DB: &DB{
			Driver: "mysql",
			Source: "root:123456@tcp(127.0.0.1:3306)/testDB",
		},
	}
}

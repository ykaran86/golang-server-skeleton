package config

type Server struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func NewServerConfig() *Server {
	return &Server{
		Port:         getIntOrPanic("API_SERVER_PORT"),
		ReadTimeout:  getIntOrPanic("API_SERVER_READ_TIMEOUT_IN_SECS"),
		WriteTimeout: getIntOrPanic("API_SERVER_WRITE_TIMEOUT_IN_SECS"),
	}
}

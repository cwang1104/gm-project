package config

import "fmt"

type server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

func (s *server) GetAddr() string {
	return fmt.Sprintf(":%s", s.Port)
}

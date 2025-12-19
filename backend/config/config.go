package config

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	Database               DatabaseConfig         `yaml:"database"`
	Kafka                  KafkaConfig            `yaml:"kafka"`
	Redis 				   RedisConfig			  `yaml:"redis"`
	GameServers 		   GameServersConfig 	  `yaml:"gameServersConfig"`
}

type DatabaseConfig struct {
	Host     		string `yaml:"host"`
	Port     		int    `yaml:"port"`
	Username 		string `yaml:"username"`
	Password 	 	string `yaml:"password"`
	DBName   	 	string `yaml:"name"`
	SSLMode  	 	string `yaml:"ssl_mode"`
	BucketQuantity	uint16 `yaml:"bucket_quantity"`
}

type KafkaConfig struct {
	Host                       string `yaml:"host"`
	Port                       int    `yaml:"port"`
	GameTopicName 			   string `yaml:"game_topic_name"`
}

type RedisConfig struct {
	Host                       string `yaml:"host"`
	Port                       int    `yaml:"port"`
	Password 				   string `yaml:"password"`
	DBStr 					   string `yaml:"dbStr"`
}

type GameServersConfig struct {
	KGS                       KGSConfig `yaml:"kgs"`
}

type KGSConfig struct {
	URL                         string `yaml:"url"`
	User                        string  `yaml:"user"`
	Password 					string 	`yaml:"password"`
	Timeout 					uint16 	`yaml:"timeout"`
	SrvID						uint8 	`yaml:"srv_id"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &config, nil
}

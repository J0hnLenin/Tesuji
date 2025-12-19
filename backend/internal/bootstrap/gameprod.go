package bootstrap

import (
	"fmt"

	"github.com/J0hnLenin/Tesuji/config"
	"github.com/J0hnLenin/Tesuji/internal/producer/gameprod"
)

func InitGameProducer(cfg *config.Config) *gameprod.GameProducer {

	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	prod := gameprod.NewGameProducer(kafkaBrockers, cfg.Kafka.GameTopicName)
	return prod
}

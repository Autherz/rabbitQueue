package rabbitQueue

import (
	"github.com/streadway/amqp"
)
type TransVideos struct {
	Input          string `json:input`
	Output         string `json:output`
	Quality        string `json:quality`
	PathP          string `json:path`
	UserId         string `json:userid`
	MaterialPackId string `json:materialpackid`
	MaterialId     string `json:materialid`
}

type Consumer interface {
	Consume(delivery amqp.Delivery)
	Transcode(transVideos *TransVideos,inputPath string, outputPath string, socketIoFlag bool) error
}
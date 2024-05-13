package Controllerkafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/customs_database_server/controller/response"
	"github.com/gin-gonic/gin"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const root = "static/test/"

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Encountered error: %v\n", err)
		return nil
	}
	if !info.IsDir() {
		fmt.Println(path)
	}
	return nil
}

func PushImgageToKafka(c *gin.Context) {
	fmt.Println("enter function: PushImgageToKafka 111")
	brokerList := []string{"172.21.116.147:9092"}

	// 创建 Kafka 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	// 连接到 Kafka 集群
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka: %v", err)
		response.ResponseErrWithMsg(c, response.CodeErrServerErr, "没有链接到kafka")
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close producer: %v", err)
		}
	}()
	fmt.Println("PushImgageToKafka - 开始读取图片")
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		data, err := ioutil.ReadFile(path)
		if info.IsDir() {
			return nil
		}
		fmt.Println("img-path: ", path)
		if err != nil {
			log.Fatalf("Failed to read image file: %v", err)
			response.ResponseErrWithMsg(c, response.CodeErrServerErr, "读取图片错误")
		}

		// 构造消息
		message := &sarama.ProducerMessage{
			Topic: "face-images",
			Value: sarama.ByteEncoder(data),
		}

		// 发送消息到 Kafka
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
			response.ResponseErrWithMsg(c, response.CodeErrServerErr, "发送数据到topic失败")
		}
		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
		response.ResponseOK(c)
		return nil
	})

}

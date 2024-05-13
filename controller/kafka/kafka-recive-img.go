package Controllerkafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/customs_database_server/controller/AttrResult"
	"github.com/customs_database_server/controller/faceResult"
	"io/ioutil"
	"log"
	"os"
)

func GetImgFromKafka() {
	// Kafka 服务器地址
	brokerList := []string{"172.21.116.147:9092"}

	// 创建 Kafka 消费者配置
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 连接到 Kafka 集群
	consumer, err := sarama.NewConsumer(brokerList, config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Failed to close consumer: %v", err)
		}
	}()

	// 订阅 topic
	topic := "face-images"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start consumer for topic %s: %v", topic, err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Failed to close partition consumer: %v", err)
		}
	}()
	fmt.Println("连接kafka成功")
	// 持续消费消息

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			// 处理接收到的消息
			log.Printf("Received message offset %d", msg.Offset)

			// 将消息数据保存为图片文件
			path := "static/tmp/face.jpg"
			err = ioutil.WriteFile(path, msg.Value, os.ModePerm)
			faceResult.FaceRecognize("1", "2024-05-09 00:00:00", path)
			AttrResult.AttrRecognize("1", path)
			if err != nil {
				log.Fatalf("Failed to write image file: %v", err)
			}
			fmt.Println("Image file saved successfully")
		case err := <-partitionConsumer.Errors():
			log.Printf("Error: %s\n", err.Error())
		}
	}
}

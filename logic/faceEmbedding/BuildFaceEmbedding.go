package logicfaceEmbedding

import (
	"bufio"
	"fmt"
	mysqlFaceEmbedding "github.com/customs_database_server/dao/mysql/FaceEmbedding"
	redisFaceEmbedding "github.com/customs_database_server/dao/redis/faceEmbedding"
	"github.com/customs_database_server/model/modelFaceEemdding"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

/*
主要用于加载人脸图像，将人脸图像发送到redis，传递给python算法进行处理
从redis接收图像，并保存到mysql
*/

const disk_path = "static/tmp/face"
const name_path = "static/tmp/name"

var cnt int

func Enter(c *gin.Context) {
	fmt.Println("进入建库函数")
	names, err := LoadNameFromTxt(name_path)
	if err != nil {
		fmt.Println(err)
		panic("人脸数据加载错误")
	}

	go func() {
		if ok := LoadData(names); !ok {
			panic("读取人脸图像数据错误")
		}
	}()
	go func() {
		SaveToDb()
	}()
}

func LoadNameFromTxt(path string) ([]string, error) {
	fmt.Println("enter LoadNameFromTxt")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	names := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		names = append(names, line)
	}
	cnt = len(names)
	fmt.Println(names)
	return names, nil
}

func LoadData(names []string) bool {
	// 从磁盘中读取图片
	files, err := os.ReadDir(disk_path)
	if err != nil {
		return false
	}
	index := 0
	for _, file := range files {
		path := disk_path + "/" + file.Name()
		fmt.Println(path)
		imgData, err := os.Open(path)
		if err != nil {
			index++
			continue
		}
		fmt.Println(index + 1)
		dataByte, err := io.ReadAll(imgData)
		sendData := &modelFaceEemdding.RedisInFaceEb{
			FaceId:  uint(index + 1),
			Name:    names[index],
			ImgData: dataByte,
		}
		ok := redisFaceEmbedding.Sent(sendData)
		fmt.Println("redis发送状态: ", ok)
		index++
	}
	fmt.Println(index, cnt)
	return index == cnt
}

func SaveToDb() {
	revcnt := 0
	for revcnt < cnt {
		if redisFaceEmbedding.Size() > 0 {
			getRedis, err := redisFaceEmbedding.Get()
			revcnt++
			if err != nil {
				panic("redis错误： redisFaceEmbedding.Get()")
			}
			saveData, err := modelFaceEemdding.RedisOutToFaceEmbedding(getRedis)
			if err != nil {
				panic("类型转换错误： modelFaceEemdding.RedisOutToFaceEmbedding(getRedis)")
			}
			ok := mysqlFaceEmbedding.CreateFace(saveData)
			if !ok {
				panic("数据库保存错误： mysqlFaceEmbedding.CreateFace(saveData)")
			}
		} else {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

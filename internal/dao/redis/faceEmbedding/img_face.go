package redisFaceEmbedding

import (
	"encoding/json"
	"fmt"
	"github.com/customs_database_server/global"
	"github.com/customs_database_server/internal/model/modelFaceEemdding"
)

const queueInName = "faceEmbedding_queue_In"
const queueOutName = "faceEmbedding_queue_out"

// 将RedisInFaceEb数据放入redis
func Sent(r *modelFaceEemdding.RedisInFaceEb) bool {
	dataJson, err := json.Marshal(r)
	if err != nil {
		print(err)
		return false
	}
	ctx := global.Redis.Context()
	if err = global.Redis.LPush(ctx, queueInName, dataJson).Err(); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 从redis读取
func Get() (*modelFaceEemdding.RedisOutFaceEb, error) {
	ctx := global.Redis.Context()
	dataJson, err := global.Redis.LPop(ctx, queueOutName).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var res modelFaceEemdding.RedisOutFaceEb
	if err = json.Unmarshal([]byte(dataJson), &res); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &res, nil
}

func Size() int {
	ctx := global.Redis.Context()
	result, err := global.Redis.LLen(ctx, queueOutName).Result()
	if err != nil {
		return -1
	}
	return int(result)
}

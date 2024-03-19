package test_face

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

//func TestMain(m *testing.M) {
//	config.InitDB() // 初始化数据库
//	fmt.Println("connect database successful...")
//	model.InitModel()
//}

func TestFaceCreate(t *testing.T) {
	// 打开文件，并获取文件对应的base64编码
	file, err := os.OpenFile("faceImgDataBase.jpg", os.O_RDONLY, 0666)
	if err != nil {
		t.Error("图片打开失败")
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Error("图片读入内容失败")
		return
	}
	base64str := base64.StdEncoding.EncodeToString(data)

	for i := 1; i <= 10; i++ {
		requestBody, err := json.Marshal(map[string]string{
			"face_id":          strconv.Itoa(i),
			"name":             "阿牛" + strconv.Itoa(i),
			"face_Img_correct": base64str,
			"face_Img_predict": base64str,
			"camera_id":        strconv.Itoa(i % 3),
			"face_Time":        fmt.Sprintf("2024-03-%02d 14:12:41", i),
		})
		if err != nil {
			t.Error("构造json格式错误")
			return
		}
		fmt.Println(string(requestBody))
		resp, err := http.Post("http://127.0.0.1:8082/save_face", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Error("发起post请求失败")
			return
		}
		defer resp.Body.Close()

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			t.Error("返回的json解析失败: ", err)
			return
		}
		fmt.Println(responseBody)
		if resp.StatusCode != http.StatusOK {
			t.Error("返回码不是200")
			return
		}
	}
}

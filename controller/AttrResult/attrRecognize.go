package AttrResult

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

const url = "http://172.21.116.147:8051/uploadfile"

func AttrRecognize(cameraID, path string) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// 创建一个新的 FormData
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加字段 camera_id
	_ = writer.WriteField("camera_id", cameraID)

	// 添加文件字段
	imageFile, err := os.Open(path) // 替换为实际的图片文件路径
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return
	}
	defer imageFile.Close()

	imagePart, err := writer.CreateFormFile("file", "face.jpg")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	_, err = io.Copy(imagePart, imageFile)
	if err != nil {
		fmt.Println("Error copying image file data:", err)
		return
	}

	// 写入 FormData 到请求体中
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing multipart writer:", err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Body = ioutil.NopCloser(body)

	// 发送 HTTP 请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()
	// 处理响应
	fmt.Println("Response Status:", resp.Status)
}

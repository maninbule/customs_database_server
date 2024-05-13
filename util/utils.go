package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func ParseBody(c *gin.Context, format interface{}) bool {
	b, err := c.GetRawData()
	if err != nil {
		return false
	}
	err = json.Unmarshal(b, format)
	fmt.Printf("%#v", format)
	if err != nil {
		fmt.Println("json解析失败: ", err)
		return false
	}
	return true
}

func ParseTime(timeStr string, t *time.Time) bool {
	result, err := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Println("timeStr = ", timeStr)
	fmt.Println("result = ", result)
	fmt.Println("err = ", err)
	if err != nil {
		return false
	}
	*t = result
	fmt.Println("t = ", *t)
	return true
}

func ParseInt(intStr string) (int64, error) {
	result, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		fmt.Println(result, err)
		return 0, errors.New("int转换错误")
	}
	return result, nil
}

func SaveFileRecResult(c *gin.Context, file *multipart.FileHeader) (string, error) {
	prefix := "static"
	middlePath := "/faceRecResult/" + time.Now().Format("2006_01_02")
	suffix := uuid.New().String() + file.Filename
	savePath := prefix + middlePath + "/" + suffix
	urlPath := "/face_img" + middlePath + "/" + suffix
	err2 := os.MkdirAll(prefix+middlePath, 0666)
	if err2 != nil {
		return "", err2
	}
	err := c.SaveUploadedFile(file, savePath)
	return urlPath, err
}

func StringToTime(t string) (*time.Time, error) {
	res, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func SaveFileFaceDataBaseFromByte(data []byte, fileName string) (string, error) {
	prefix := "static"
	middlePath := "/faceImgDataBase/" + time.Now().Format("2006_01_02")
	suffix := uuid.New().String() + fileName
	savePath := prefix + middlePath + "/" + suffix
	urlPath := "/face_img" + middlePath + "/" + suffix
	err2 := os.MkdirAll(prefix+middlePath, 0666)
	if err2 != nil {
		return "", err2
	}
	openFile, err2 := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0666)
	for len(data) > 0 {
		n, err2 := openFile.Write(data)
		if err2 == io.EOF {
			break
		}
		fmt.Println(n, len(data))
		data = data[n:]
	}
	return urlPath, nil
}
func SaveFileFaceDataBase(c *gin.Context, file *multipart.FileHeader) (string, error) {
	prefix := "static"
	middlePath := "/faceImgDataBase/" + time.Now().Format("2006_01_02")
	suffix := uuid.New().String() + file.Filename
	savePath := prefix + middlePath + "/" + suffix
	urlPath := "/face_img" + middlePath + "/" + suffix
	err2 := os.MkdirAll(prefix+middlePath, 0666)
	if err2 != nil {
		return "", err2
	}
	err := c.SaveUploadedFile(file, savePath)
	return urlPath, err
}

func SaveFileGaitDataBase(c *gin.Context, file *multipart.FileHeader) (string, error) {
	prefix := "static"
	middlePath := "/GaitImgDataBase/" + time.Now().Format("2006_01_02")
	suffix := uuid.New().String() + file.Filename
	savePath := prefix + middlePath + "/" + suffix
	urlPath := "/face_img" + middlePath + "/" + suffix
	err2 := os.MkdirAll(prefix+middlePath, 0666)
	if err2 != nil {
		return "", err2
	}
	err := c.SaveUploadedFile(file, savePath)
	return urlPath, err
}

func SaveAttrDataBase(c *gin.Context, file *multipart.FileHeader) (string, error) {
	prefix := "static"
	middlePath := "/AttrImgDataBase/" + time.Now().Format("2006_01_02")
	suffix := uuid.New().String() + file.Filename
	savePath := prefix + middlePath + "/" + suffix
	urlPath := "/face_img" + middlePath + "/" + suffix
	err2 := os.MkdirAll(prefix+middlePath, 0666)
	if err2 != nil {
		return "", err2
	}
	err := c.SaveUploadedFile(file, savePath)
	return urlPath, err
}

package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"
	"github.com/astaxie/beego"
	"github.com/parnurzeal/gorequest"
	"io"
	"mime/multipart"
)

func AddSalt(raw string) string {
	return hmacSha1Encode(raw, beego.AppConfig.String("salt"))
}

func hmacSha1Encode(input string, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	_, _ = io.WriteString(h, input)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CheckRecaptcha(response string, remoteip string) bool {
	req := gorequest.New().Post(beego.AppConfig.String("recaptcha_domain") + "/recaptcha/api/siteverify").Type("form")
	req.SendMap(map[string]string{
		"secret":   beego.AppConfig.String("recaptcha_server_key"),
		"response": response,
		"remoteip": remoteip,
	})
	resp, body, _ := req.End()
	if body == "" || resp == nil || resp.StatusCode != 200 {
		return false
	}

	recaptcha := new(RecaptchaResponse)
	err := json.Unmarshal([]byte(body), &recaptcha)
	if err != nil {
		return false
	}
	if recaptcha.Success {
		return true
	}
	return false
}

func UploadPicture(header *multipart.FileHeader, file multipart.File) string {
	fileByte := make([]byte, header.Size)
	_, _ = file.Read(fileByte)
	upload_name := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	req := gorequest.New().Post(beego.AppConfig.String("upload_url")).Type("multipart")
	req.Header.Set("Authorization", beego.AppConfig.String("upload_token"))
	req.SendFile(fileByte, upload_name, "smfile")
	resp, body, _ := req.End()

	if resp != nil && resp.StatusCode == 200 {
		responseJson := new(UploadCallBack)
		err := json.Unmarshal([]byte(body), &responseJson)
		if err == nil {
			if responseJson.Code == "success" {
				return responseJson.Data.Url
			} else if responseJson.Code == "image_repeated" {
				return responseJson.ExistedUrl
			} else {
				fmt.Printf("[%s]%s\n", time.Now().Format(time.RFC3339), responseJson.Message)
			}
		} else {
			fmt.Printf("[%s]%s\n", time.Now().Format(time.RFC3339), err)
		}
	}
	return ""
}

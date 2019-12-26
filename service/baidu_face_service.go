package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"facecheckin/model"
	"facecheckin/serializer"
	"log"
	"net/http"
	"strings"
)

type BaiduFaceService struct {
	NewFace string `form:"face" json:"face" binding:"required"`
	UserId string `form:"uid" json:"uid" binding:"required"`
}
type Body struct {
	Image      string `json:"image"`
	Image_type string `json:"image_type"`
}
type Response struct {
	Error_code int    `json:"error_code"`
	Error_msg  string `json:"error_msg"`
	Result     Result `json:"result"`
}
type Result struct {
	Score int `json:"score"`
}

func (service BaiduFaceService) GetScore() serializer.Response {
	url := "https://aip.baidubce.com/rest/2.0/face/v3/match"
	var body []Body
	body = append(body, Body{
		Image:      service.NewFace,
		Image_type: "BASE64",
	})
	var user model.User
	if err:= model.DB.Where("phone_number = ?", service.UserId).First(&user).Error; err!=nil{
		return serializer.ParamErr("用户未寻到",err)
	}
	result := strings.Split(user.Face,",")
	body = append(body, Body{
		Image:      result[1],
		Image_type: "BASE64",
	})
	params := make(map[string]string)
	params["access_token"] = "24.5941ade8867ff403fd95676a0f855bae.2592000.1578917498.282335-18025567"

	headers := make(map[string]string)
	//headers["Content-Type"] = "application/json"

	resp, err := Post(url, body, params, headers)
	if err != nil {
		return serializer.Err(40002, "出现错误", err)
	}
	//repbody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return serializer.Err(40003,"未收到resp",err)
	//}

	var res Response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return serializer.Err(40005, "json解析错误", err)
	}
	//fmt.Println(res.Error_code)
	//fmt.Println(res.Error_msg)
	//fmt.Println(res.Result.Score)
	return serializer.Response{
		Code: res.Error_code,
		Data: res.Result,
		Msg:  res.Error_msg,
	}

}
func Post(url string, body interface{}, params map[string]string, headers map[string]string) (*http.Response, error) {
	//add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}
	log.Printf("Go %s URL : %s \n", http.MethodPost, req.URL.String())
	return client.Do(req)
}

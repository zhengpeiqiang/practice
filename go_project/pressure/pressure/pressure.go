package pressure

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"strings"
	"time"
	"use_granphviz/middleware"
	"use_granphviz/public"
)

var ErrorNum = 0

type PreMsg struct {
	IsCustomerClient int `json:"isCustomerClient" form:"isCustomerClient"  example:"1"` //是否客户端 1 是 0 否
}

type Base struct {
	PageNum  int    `json:"pageNum" form:"pageNum" example:"1"`    //页码
	PageSize int    `json:"pageSize" form:"pageSize" example:"10"` //每页数量
	Search   string `json:"search" form:"search" example:""`       //搜索内容
	Type     int    `json:"type" form:"type" example:""`           //类型
}

type User struct {
	PreMsg
	Base
	UserId      string `json:"userId" form:"userId"  example:"1"`                 //用户 Id
	OpenId      string `json:"wechat_openid" form:"wechat_openid"  example:"XXX"` //用户 openid
	UnionId     string `json:"unionId" form:"unionId"  example:"XXX"`             //用户 unionid 微信平台返回
	IsOnline    int    `json:"is_online" form:"is_online"  example:"0"`           //是否在线（0 不在线（离线）；1 在线；2 忙碌；）
	GroupId     string `json:"groupId" form:"groupId"  example:"" `               //群 Id
	ChannelType int    `json:"channelType" form:"channelType"  example:"1"`       //属于哪个渠道：1 pc；2 小程序；
	AccessType  int    `json:"accessType" form:"accessType"  example:"1"`         //接入模式：0 无；1 转发；2 转接；
}

type RongCloudTokenInput struct {
	PreMsg
	UserId      string `json:"userId" form:"userId"  example:"123" comment:"用户Id，应用内唯一标识"`                               //用户 Id，应用内唯一标识
	Name        string `json:"name" form:"name" example:"rock" comment:"用户名称，最大长度 128 字节"`                               //用户名称，最大长度 128 字节。用来在 Push 推送时显示用户的名称，重新获取 Token 传入新的名称后，将覆盖之前的用户名称
	PortraitUri string `json:"portraitUri" form:"portraitUri" example:"http://xxx.png" comment:"用户头像 URI，最大长度 1024 字节。"` //用户头像 URI，最大长度 1024 字节。
}

//const pressure_checkuser_url = "http://mjt.emakerzone.com/customer/checkuser"
const pressure_checkuser_url = "http://127.0.0.1:8888/customer/checkuser"

const pressure_messageslist_url = "http://mjt.emakerzone.com/customer/messageslist"
//const pressure_messageslist_url = "http://127.0.0.1:8888/customer/messageslist"

const channel_url = "http://127.0.0.1:9999/pop"

func checkuser(url string, signature string, isCustomerClient int, unionId string, channelType int) {
	startTime := time.Now().Format("2006-01-02 15:04:05")
	queryParam := User{
		PreMsg: PreMsg{
			IsCustomerClient: isCustomerClient,
		},
		UnionId:     unionId,
		ChannelType: channelType,
	}
	paramJsonBytes, err := json.Marshal(queryParam)
	if err != nil {
		fmt.Println("checkuser err[1] => ", err, signature)
		return
	}
	paramJsonStr := string(paramJsonBytes)
	queryUrl := url + signature
	_, queryByte, err := lib.HttpJSON(lib.NewTrace(), queryUrl, paramJsonStr, 60*1000, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "is normally permitted.") {
			ErrorNum++
			fmt.Println("checkuser err[2] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), signature)
			return
		}
		return
	}
	resp := &middleware.Response{}
	if err := json.Unmarshal(queryByte, resp); err != nil {
		ErrorNum++
		fmt.Println("checkuser err[3] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), "; queryByte => ", string(queryByte), signature)
		return
	}
	if resp.ErrorCode != middleware.SUCCESS_CODE {
		ErrorNum++
		fmt.Println("checkuser err[4] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), signature)
		return
	}
	//fmt.Println("checkuser success")
}

//func rctoken(name string, portraitUri string, userId string) {
//	queryParam := RongCloudTokenInput{
//		Name:        name,
//		PortraitUri: portraitUri,
//		UserId:      userId,
//	}
//	paramJsonBytes, err := json.Marshal(queryParam)
//	if err != nil {
//		fmt.Println("rctoken err[1] => ", err)
//		return
//	}
//	paramJsonStr := string(paramJsonBytes)
//	queryUrl := "http://mjt.emakerzone.com/customer/rctoken"
//	_, queryByte, err := lib.HttpJSON(lib.NewTrace(), queryUrl, paramJsonStr, 60*1000, nil)
//	if err != nil {
//		fmt.Println("rctoken err[2] => ", err)
//		return
//	}
//	resp := &middleware.Response{}
//	if err := json.Unmarshal(queryByte, resp); err != nil {
//		fmt.Println("rctoken err[3] => ", err)
//		return
//	}
//	if resp.ErrorCode != middleware.SUCCESS_CODE {
//		fmt.Println("rctoken err[4] => ", err)
//		return
//	}
//	//fmt.Println("rctoken success")
//}

func messageslist(url string, signature string, userId string, openId string) {
	startTime := time.Now().Format("2006-01-02 15:04:05")
	queryParam := User{
		UserId: userId,
	}
	paramJsonBytes, err := json.Marshal(queryParam)
	if err != nil {
		fmt.Println("messageslist err[1] => ", err, signature)
		return
	}
	paramJsonStr := string(paramJsonBytes)
	queryUrl := url + signature
	_, queryByte, err := lib.HttpJSON(lib.NewTrace(), queryUrl, paramJsonStr, 60*1000, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "is normally permitted.") {
			ErrorNum++
			fmt.Println("messageslist err[2] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), "; queryByte => ", string(queryByte), signature)
			return
		}
		return
	}
	resp := &middleware.Response{}
	if err := json.Unmarshal(queryByte, resp); err != nil {
		ErrorNum++
		fmt.Println("messageslist err[3] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), "; queryByte => ", string(queryByte), signature)
		return
	}
	if resp.ErrorCode != middleware.SUCCESS_CODE {
		ErrorNum++
		fmt.Println("messageslist err[4] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), signature)
		return
	}
	//fmt.Println("messageslist success")
}

func channel(url string, channels string) {
	startTime := time.Now().Format("2006-01-02 15:04:05")
	queryUrl := url + channels
	_, queryByte, err := lib.HttpGET(lib.NewTrace(), queryUrl, nil, 60*1000, nil)
	if err != nil {
		if !strings.Contains(err.Error(), "is normally permitted.") {
			ErrorNum++
			fmt.Println("messageslist err[2] => ", err, "; start_time => ", startTime, "; end_time => ", time.Now().Format("2006-01-02 15:04:05"), "; queryByte => ", string(queryByte), channels)
			return
		}
		return
	}
}

func TimeOut(
	goNum int,
	isCustomerClient int, unionId string, channelType int,
	name string, portraitUri string, rcuserId string,
	userId string) {
	messageslist(pressure_messageslist_url, "?signature="+public.RandomString(16), userId, "1")
}

func Pressure(
	goNum int,
	isCustomerClient int, unionId string, channelType int,
	name string, portraitUri string, rcuserId string,
	userId string) {
	//fmt.Println("压测开始 time => ", time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < goNum; i++ {
		go func() {
			signature := "?signature=" + public.RandomString(16)
			//signature := ""
			//checkuser(pressure_checkuser_url, signature, isCustomerClient, unionId, channelType)
			//rctoken(name, portraitUri, rcuserId)
			messageslist(pressure_messageslist_url, signature, userId, "1")
			//messageslist(pressure_messageslist_url, userId, "2")
			//fmt.Println("压测结束 time => ", time.Now().Format("2006-01-02 15:04:05"))
		}()
	}
}

func Channel(goNum int, channelType int) {
	//fmt.Println("压测开始 time => ", time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < goNum; i++ {
		go func() {
			channels := "?channel=" + gconv.String(grand.N(10, 10000)) + "&channel_type=" + gconv.String(channelType)
			channel(channel_url, channels)
		}()
	}
}

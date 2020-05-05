package model

import (
	"github.com/kataras/iris/v12"
	"net/url"
	"strings"
)

// BlogGin hello BlogGin.
type GiNana struct {
	Hello string
}

type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// 一些自定义的类型
type GetClientIP func() string

type GetOption func(name string) string

type GetOptionHandler func(ctx iris.Context) (GetOption, error)

type Validator func(obj interface{}) error

type ValidatorHandler func(ctx iris.Context) (Validator, error)

type JsonPlus func(data interface{}, msg interface{}) *JSON

// 后台数据统计
type AdminData struct {
	Hostname      string
	Gover         string
	OS            string
	Arch          string
	CountCpu      int
	CountArticles int64
	CountUsers    int64
	CountTags     int64
}

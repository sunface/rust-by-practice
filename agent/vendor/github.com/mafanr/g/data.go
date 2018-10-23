package g

import (
	"encoding/base64"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
)

var Cli = &fasthttp.Client{
	MaxConnsPerHost:     2000,
	MaxIdleConnDuration: 60 * time.Second,
}
var DB *sqlx.DB
var B64 = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

type Result struct {
	// 每次请求的唯一ID，适用于tfe返回给外部的用户
	// 外部的用户可以通过rid对请求的错误进行跟踪
	RID int64 `json:"rid"`

	// HTTP status
	// 所有接口应该按照标准来返回http状态码：只有业务逻辑成功才返回200，其余都返回相应的错误状态码
	// 但是有极少数接口，出于特殊原因，只要请求通了就返回200，并不遵守上述规则
	// 因此，在结果里添加返回http status一项，不管按照哪种做法，这里的状态码必须对应相应的HTTP状态码
	// 常见的例如：301,401,404,501,502等等
	Status int `json:"status"`

	// 错误码
	// 如果业务逻辑没有成功，必须说明对应的错误码
	// 错误码一般会比状态码更加详细
	ErrCode int `json:"err_code"`

	// 给用户看的信息
	// 成功的时候，可以给用户展示成功的提示
	// 错误的时候，可以给用户展示错误的提示
	Message string `json:"message"`

	// 返回的数据
	// 当业务需要给用户返回数据时，使用该字段
	Data interface{} `json:"data"`
}

// Service的服务节点注册自己到ETCD
type ServerInfo struct {
	Service string
	IP      string // ip + port: localhost:8080
	Load    float64
}

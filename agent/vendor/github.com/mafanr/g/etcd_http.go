package g

import (
	"time"

	"github.com/valyala/fasthttp"
)

var HTTP = &Http{}

type Http struct{}

func (h *Http) Reqeust(app string, path string, method string, args *fasthttp.Args) (*fasthttp.Response, error) {
	req := &fasthttp.Request{}
	req.Header.SetMethod(method)

	s := GetServer(app)
	url := "http://" + s.IP + path
	switch method {
	case "GET":
		url = url + "?" + args.String()
	default:
		args.WriteTo(req.BodyWriter())
	}
	req.SetRequestURI(url)

	resp := &fasthttp.Response{}
	err := Cli.DoTimeout(req, resp, HTTP_REQ_TIMEOUT*time.Second)
	return resp, err
}

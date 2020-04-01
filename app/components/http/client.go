package http

import (
	"bytes"
	"github.com/kataras/iris/context"
	"io/ioutil"
	"orange_message_service/app/utils"
	"net/http"
	"time"
)

type Response struct {
	RequestTime time.Time
	ReceivedAt  time.Time
	Body        []byte
	StatusCode  int
	Time        time.Duration
}

func PostJson(ctx context.Context, url string, body []byte, timeout time.Duration) (*Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range ctx.Request().Header {
		req.Header.Set(key, utils.GetStringByIndex(value, 0))
	}
	req.Header.Set("Content-Type", "application/json")

	startTime := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respByte, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, errRead
	}
	now := time.Now()
	return &Response{
		RequestTime: startTime,
		ReceivedAt:  now,
		Body:        respByte,
		StatusCode:  resp.StatusCode,
		Time:        now.Sub(startTime),
	}, nil
}

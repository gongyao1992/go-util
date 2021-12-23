package talk

import (
	"errors"
	"github.com/CatchZeng/dingtalk"
	"github.com/gongyao1992/go-util/helper"
	"strings"
)

type DingRobitText struct {
	client *dingtalk.Client
	at []string
	isAtAll bool
}

func NewDingRobitText(accessToken, secret, atMobiles string, isAtAll bool) *DingRobitText {
	at := make([]string, 0)
	if len(atMobiles) > 0 {
		at = strings.Split(atMobiles, ",")
	}
	return &DingRobitText{
		client:  dingtalk.NewClient(accessToken, secret),
		at:      at,
		isAtAll: isAtAll,
	}
}

func (r *DingRobitText)Send(i interface{}) error {

	var text string
	switch i.(type) {
	case string:
		text = i.(string)
	default:
		text = helper.ToJson(i)
	}

	msg := dingtalk.NewTextMessage().SetContent(text).SetAt(r.at, r.isAtAll)
	send, err := r.client.Send(msg)
	if err != nil {
		return err
	}
	if send.ErrCode > 0 {
		return errors.New(send.ErrMsg)
	}
	return nil
}
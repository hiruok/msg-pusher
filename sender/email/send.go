/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : send.go
#   Created       : 2019/1/16 16:44
#   Last Modified : 2019/1/16 16:44
#   Describe      :
#
# ====================================================*/
package email

import (
	"context"

	"github.com/sirupsen/logrus"
	"uuabc.com/sendmsg/pkg/pb/meta"
	"uuabc.com/sendmsg/pkg/send/email"
	"uuabc.com/sendmsg/sender/pub"
)

// check 验证data是否符合要求，如果符合要求会返回nil，并且按照data转化的id将数据赋值给msg
func (r *Receiver) check(data []byte, msg pub.Messager) (err error) {
	id := string(data)
	logrus.WithField("type", r.queueName).Info("开始验证消息的有效性")
	err = pub.Check(id, msg)
	logrus.WithField("type", r.queueName).Infof("消息验证结束,err: %v", err)
	return
}

func (r *Receiver) send(msg pub.Messager) error {
	emailMsg := msg.(*meta.DbEmail)
	return pub.EmailClient.Send(email.NewMessage(
		emailMsg.Destination,
		emailMsg.Title,
		emailMsg.Content,
	), nil)
}

func (r *Receiver) doList(c pub.Cache, b []byte) error {
	return c.RPushEmail(context.Background(), b)
}

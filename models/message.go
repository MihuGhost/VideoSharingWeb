package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Message struct {
	Id      int
	Content string
	AddTime int64
}
type MessageUser struct {
	Id        int
	MessageId int64
	AddTime   int64
	Status    int
	UserId    int
}

func init() {
	orm.RegisterModel(new(Message), new(MessageUser))
}

func SendMessageDo(content string) (int64, error) {
	o := orm.NewOrm()
	message := Message{
		Content: content,
		AddTime: time.Now().Unix(),
	}
	messageId, err := o.Insert(&message)
	return messageId, err
}

// todo MQ
func SendMessageUser(userId int, messageId int64) error {
	o := orm.NewOrm()
	messageUser := MessageUser{
		MessageId: messageId,
		AddTime:   time.Now().Unix(),
		Status:    1,
		UserId:    userId,
	}
	_, err := o.Insert(&messageUser)
	return err
}

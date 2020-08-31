package service

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

type (
	DocMsgDefineService struct {
		Name              string   `bson:"name" yaml:"name"`
		Description       string   `bson:"description" yaml:"description"`
		Tags              []string `bson:"tags" yaml:"tags"`
		Author            string   `bson:"author" yaml:"author"`
		AuthorDescription string   `bson:"author_description" yaml:"author_description"`
		Schemas           string   `bson:"schemas"`
	}
)

func (m *DocMsgDefineService) GetType() string {
	return MsgTypeDefineService
}

func (m *DocMsgDefineService) BuildMsg(v interface{}) {
	msg := v.(MsgDefineService)

	m.Name = msg.Name
	m.Description = msg.Description
	m.Tags = msg.Tags
	m.Author = msg.Author.String()
	m.AuthorDescription = msg.AuthorDescription
	m.Schemas = msg.Schemas
}

func (m *DocMsgDefineService) HandleTxMsg(msg MsgDefineService) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, msg.Author.String())
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}

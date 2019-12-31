package msgs

import (
	"gitlab.bianjie.ai/irita/ex-sync/models"
	"gitlab.bianjie.ai/irita/ex-sync/types"
)

type (
	DocMsgServiceDef struct {
		Name              string   `bson:"name" yaml:"name"`
		ChainId           string   `bson:"chain_id" yaml:"chain_id"`
		Description       string   `bson:"description" yaml:"description"`
		Tags              []string `bson:"tags" yaml:"tags"`
		Author            string   `bson:"author" yaml:"author"`
		AuthorDescription string   `bson:"author_description" yaml:"author_description"`
		IDLContent        string   `bson:"idl_content" yaml:"idl_content"`
	}
)

func (m *DocMsgServiceDef) GetType() string {
	return MsgTypeServiceDef
}

func (m *DocMsgServiceDef) BuildMsg(v interface{}) {
	msg := v.(types.MsgServiceDef)

	m.Name = msg.Name
	m.ChainId = msg.ChainId
	m.Description = msg.Description
	m.Tags = msg.Tags
	m.Author = msg.Author.String()
	m.AuthorDescription = msg.AuthorDescription
	m.IDLContent = msg.IDLContent
}

func (m *DocMsgServiceDef) HandleTxMsg(v interface{}) MsgDocInfo {
	var (
		from, to, signer string
		coins            []models.Coin
		docTxMsg         models.DocTxMsg
		complexMsg       bool
		signers          []string
	)

	msg := v.(types.MsgServiceDef)
	from = msg.Author.String()
	to = ""
	coins = models.BuildDocCoins(nil)

	m.BuildMsg(v)
	docTxMsg = models.DocTxMsg{
		Type: m.GetType(),
		Msg:  m,
	}
	complexMsg = false

	signer, signers = models.BuildDocSigners(msg.GetSigners())

	return MsgDocInfo{
		From:       from,
		To:         to,
		Coins:      coins,
		Signer:     signer,
		DocTxMsg:   docTxMsg,
		ComplexMsg: complexMsg,
		Signers:    signers,
	}
}

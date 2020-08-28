package gov

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/models"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"encoding/json"
)

type DocTxMsgSubmitProposal struct {
	Proposer       string        `bson:"proposer"`        //  Address of the proposer
	InitialDeposit []models.Coin `bson:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Content        string        `bson:"content"`
}

func (doctx *DocTxMsgSubmitProposal) GetType() string {
	return TxTypeSubmitProposal
}

func (doctx *DocTxMsgSubmitProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgSubmitProposal)
	content, _ := json.Marshal(msg.Content)
	doctx.Content = string(content)
	doctx.Proposer = msg.Proposer.String()
	doctx.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
}

func (m *DocTxMsgSubmitProposal) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}

// MsgVote
type DocTxMsgVote struct {
	ProposalID uint64 `bson:"proposal_id"` // ID of the proposal
	Voter      string `bson:"voter"`       //  address of the voter
	Option     string `bson:"option"`      //  option from OptionSet chosen by the voter
}

func (doctx *DocTxMsgVote) GetType() string {
	return TxTypeVote
}

func (doctx *DocTxMsgVote) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgVote)
	doctx.Voter = msg.Voter.String()
	doctx.Option = msg.Option.String()
	doctx.ProposalID = msg.ProposalID
}

func (m *DocTxMsgVote) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
// MsgDeposit
type DocTxMsgDeposit struct {
	ProposalID uint64        `bson:"proposal_id"` // ID of the proposal
	Depositor  string        `bson:"depositor"`   // Address of the depositor
	Amount     []models.Coin `bson:"amount"`      // Coins to add to the proposal's deposit
}

func (doctx *DocTxMsgDeposit) GetType() string {
	return TxTypeDeposit
}

func (doctx *DocTxMsgDeposit) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgDeposit)
	doctx.Depositor = msg.Depositor.String()
	doctx.Amount = models.BuildDocCoins(msg.Amount)
	doctx.ProposalID = msg.ProposalID
}

func (m *DocTxMsgDeposit) HandleTxMsg(msg sdk.Msg) MsgDocInfo {

	var (
		addrs []string
	)

	addrs = append(addrs, m.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
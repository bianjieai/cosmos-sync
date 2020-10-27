package staking

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	stake "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// MsgDelegate - struct for bonding transactions
type DocTxMsgBeginRedelegate struct {
	DelegatorAddress    string `bson:"delegator_address"`
	ValidatorSrcAddress string `bson:"validator_src_address"`
	ValidatorDstAddress string `bson:"validator_dst_address"`
	Amount              string `bson:"amount"`
}

func (doctx *DocTxMsgBeginRedelegate) GetType() string {
	return MsgTypeBeginRedelegate
}

func (doctx *DocTxMsgBeginRedelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgBeginRedelegate)
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.ValidatorSrcAddress = msg.ValidatorSrcAddress
	doctx.ValidatorDstAddress = msg.ValidatorDstAddress
	doctx.Amount = msg.Amount.String()
}
func (m *DocTxMsgBeginRedelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgBeginRedelegate
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorDstAddress, msg.ValidatorSrcAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

// MsgUnjail - struct for unjailing jailed validator
type DocTxMsgUnjail struct {
	ValidatorAddr string `bson:"address"` // address of the validator operator
}

func (doctx *DocTxMsgUnjail) GetType() string {
	return MsgTypeUnjail
}

func (doctx *DocTxMsgUnjail) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUnjail)
	doctx.ValidatorAddr = msg.ValidatorAddr
}
func (m *DocTxMsgUnjail) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgUnjail
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.ValidatorAddr)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

// MsgBeginUnbonding - struct for unbonding transactions
type DocTxMsgBeginUnbonding struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
	Amount           string `bson:"amount"`
}

func (doctx *DocTxMsgBeginUnbonding) GetType() string {
	return MsgTypeStakeBeginUnbonding
}

func (doctx *DocTxMsgBeginUnbonding) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeBeginUnbonding)
	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.Amount = msg.Amount.String()
}
func (m *DocTxMsgBeginUnbonding) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeBeginUnbonding
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

// MsgDelegate - struct for bonding transactions
type DocTxMsgDelegate struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
	Delegation       Coin   `bson:"delegation"`
}

func (doctx *DocTxMsgDelegate) GetType() string {
	return MsgTypeStakeDelegate
}

func (doctx *DocTxMsgDelegate) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeDelegate)
	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.Delegation = Coin(models.BuildDocCoin(msg.Amount))
}
func (m *DocTxMsgDelegate) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeDelegate
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

// MsgEditValidator - struct for editing a validator
type DocMsgEditValidator struct {
	Description       models.Description `bson:"description"`
	ValidatorAddress  string             `bson:"validator_address"`
	CommissionRate    string             `bson:"commission_rate"`
	MinSelfDelegation string             `bson:"min_self_delegation"`
}

func (doctx *DocMsgEditValidator) GetType() string {
	return MsgTypeStakeEditValidator
}

func (doctx *DocMsgEditValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeEdit)
	doctx.ValidatorAddress = msg.ValidatorAddress
	commissionRate := msg.CommissionRate
	if commissionRate == nil {
		doctx.CommissionRate = ""
	} else {
		doctx.CommissionRate = commissionRate.String()
	}
	doctx.Description = loadDescription(msg.Description)
	if msg.MinSelfDelegation != nil {
		doctx.MinSelfDelegation = msg.MinSelfDelegation.String()
	}
}
func (m *DocMsgEditValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeEdit
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

// MsgCreateValidator defines an SDK message for creating a new validator.
type DocTxMsgCreateValidator struct {
	Description       models.Description     `bson:"description"`
	Commission        models.CommissionRates `bson:"commission"`
	MinSelfDelegation string                 `bson:"min_self_delegation"`
	DelegatorAddress  string                 `bson:"delegator_address"`
	ValidatorAddress  string                 `bson:"validator_address"`
	Pubkey            string                 `bson:"pubkey"`
	Value             Coin                   `bson:"value"`
}

func (doctx *DocTxMsgCreateValidator) GetType() string {
	return MsgTypeStakeCreateValidator
}

func (doctx *DocTxMsgCreateValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeCreate)
	//pubKey, err := itypes.Bech32ifyValPub(msg.Pubkey)
	//if err != nil {
	//	pubKey = ""
	//}
	doctx.ValidatorAddress = msg.ValidatorAddress
	doctx.Pubkey = msg.Pubkey
	doctx.DelegatorAddress = msg.DelegatorAddress
	doctx.MinSelfDelegation = msg.MinSelfDelegation.String()
	doctx.Commission = models.CommissionRates{
		Rate:          msg.Commission.Rate.String(),
		MaxChangeRate: msg.Commission.MaxChangeRate.String(),
		MaxRate:       msg.Commission.MaxRate.String(),
	}
	doctx.Description = loadDescription(msg.Description)
	doctx.Value = Coin{Denom: msg.Value.Denom, Amount: msg.Value.Amount.String()}
}
func (m *DocTxMsgCreateValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgStakeCreate
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}

func loadDescription(description stake.Description) models.Description {
	return models.Description{
		Moniker:         description.Moniker,
		Details:         description.Details,
		Identity:        description.Identity,
		Website:         description.Website,
		SecurityContact: description.SecurityContact,
	}
}

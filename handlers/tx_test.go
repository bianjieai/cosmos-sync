package handlers

import (
	"github.com/bianjieai/cosmos-sync/config"
	"github.com/bianjieai/cosmos-sync/libs/pool"
	"github.com/bianjieai/cosmos-sync/models"
	"github.com/bianjieai/cosmos-sync/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/tharsis/ethermint/x/evm/types"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(16038502)
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	InitRouter(conf)
	models.Init(conf)
	pool.Init(conf)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))

		//b, _ := hex.DecodeString("736572766963652063616c6c20726573706f6e7365")
		//t.Log(string(b))
	}
}

func TestDecodeTxResponse(t *testing.T) {
	resByte, err := hexutil.Decode("0x0AE6010A1F2F65746865726D696E742E65766D2E76312E4D7367457468657265756D547812C2010A423078616136626233323934363638643264643939343862393265636330386434636563363239323933626563356131653034376234666261343435663237343238661A6408C379A00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000001E444443313135353A6E6F74206F776E6572206E6F7220617070726F76656400002212657865637574696F6E20726576657274656428B7B304")
	if err != nil {
		t.Fatal(err.Error())
	}

	resultResp, err := types.DecodeTxResponse(resByte)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(resultResp)
}

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/cosmos-sync/libs/msgparser/types"
	"strings"
)

func MarshalJsonIgnoreErr(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}

func UnMarshalJsonIgnoreErr(data string, v interface{}) {
	json.Unmarshal([]byte(data), &v)
}

// parseDenomTrace parses a string with the ibc prefix (denom trace) and the base denomination
// into a DenomTrace type.
//
// Examples:
//
// 	- "portidone/channelidone/uatom" => DenomTrace{Path: "portidone/channelidone", BaseDenom: "uatom"}
// 	- "uatom" => DenomTrace{Path: "", BaseDenom: "uatom"}
func ParseDenomTrace(rawDenom string) types.DenomTrace {
	denomSplit := strings.Split(rawDenom, "/")

	if denomSplit[0] == rawDenom {
		return types.DenomTrace{
			Path:      "",
			BaseDenom: rawDenom,
		}
	}

	return types.DenomTrace{
		Path:      strings.Join(denomSplit[:len(denomSplit)-1], "/"),
		BaseDenom: denomSplit[len(denomSplit)-1],
	}
}

func IBCDenom(fullDenomPath string) string {
	hash := sha256.Sum256([]byte(fullDenomPath))
	bz := hash[:]
	return fmt.Sprintf("%s/%s", "ibc", strings.ToUpper(hex.EncodeToString(bz)))
}

//  returns true if the denomination originally came
// from the receiving chain and false otherwise.
func ReceiverChainIsSource(sourcePort, sourceChannel, denom string) bool {
	// The prefix passed in should contain the SourcePort and SourceChannel.
	// If  the receiver chain originally sent the token to the sender chain
	// the denom will have the sender's SourcePort and SourceChannel as the
	// prefix.

	voucherPrefix := GetDenomPrefix(sourcePort, sourceChannel)
	return strings.HasPrefix(denom, voucherPrefix)

}

// GetDenomPrefix returns the receiving denomination prefix
func GetDenomPrefix(portID, channelID string) string {
	return fmt.Sprintf("%s/%s/", portID, channelID)
}

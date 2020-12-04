package identity

import (
	"github.com/bianjieai/iritamod/modules/identity/types"
)

// temporary method getPubKeyFromCertificate
// this method can be removed while
// chain add pubKey of cert into tx events
func getPubKeyFromCertificate(certificate string) PubKeyInfo {
	cert := []byte(certificate)
	pubKey := types.GetPubKeyFromCertificate(cert)
	return PubKeyInfo{
		PubKey:    pubKey.PubKey,
		Algorithm: int32(pubKey.Algorithm),
	}
}

package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"
	hex2 "github.com/shniu/gokits/common/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Generate key -> hex -> ECDSA
func Test_GenerateKey_toHex_thenToECDSA(t *testing.T) {
	priv, err := GenerateKey()
	assert.Nil(t, err)

	b := FromECDSA(priv)
	privHex := hex.EncodeToString(b)
	t.Log(privHex)

	privHex2 := hex2.Encode(b)
	t.Log(privHex2)

	priv2, err := HexToECDSA(privHex)
	assert.Nil(t, err)
	assert.Equal(t, priv.D, priv2.D)
}

func Test_GenerateKey_toPubKey_toAddress(t *testing.T) {
	priv, err := GenerateKey()
	assert.Nil(t, err)

	pub, ok := priv.Public().(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("can't convert to ecdsa.PublickKey")
	}

	address := PubkeyToAddress(*pub)
	t.Log(address.Hex())
}

func Test_Address_isOk(t *testing.T) {
	priv, err := HexToECDSA("18e14a7b6a307f426a94f8114701e7c8e774e7f9a47e2c2035db29a206321725")
	assert.Nil(t, err)

	pub, ok := priv.Public().(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("can't convert to ecdsa.PublickKey")
	}

	pubBytes := FromECDSAPub(pub)
	pubHex := hex2.Encode(pubBytes)
	t.Log(pubHex)

	address := PubkeyToAddress(*pub)
	t.Log(address.Hex())

	assert.Equal(t, "0x3E9003153d9A39D3f57B126b0c38513D5e289c3E", address.Hex())
}

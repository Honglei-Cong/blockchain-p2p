package crypto

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type byter interface {
	Bytes() []byte
}

func checkAminoBinary(t *testing.T, src byter, dst interface{}, size int) {
	// Marshal to binary bytes.
	bz, err := cdc.MarshalBinaryBare(src)
	require.Nil(t, err, "%+v", err)
	// Make sure this is compatible with current (Bytes()) encoding.
	assert.Equal(t, src.Bytes(), bz, "Amino binary vs Bytes() mismatch")
	// Make sure we have the expected length.
	if size != -1 {
		assert.Equal(t, size, len(bz), "Amino binary size mismatch")
	}
	// Unmarshal.
	err = cdc.UnmarshalBinaryBare(bz, dst)
	require.Nil(t, err, "%+v", err)
}

func checkAminoJSON(t *testing.T, src interface{}, dst interface{}, isNil bool) {
	// Marshal to JSON bytes.
	js, err := cdc.MarshalJSON(src)
	require.Nil(t, err, "%+v", err)
	if isNil {
		assert.Equal(t, string(js), `null`)
	} else {
		assert.Contains(t, string(js), `"type":`)
		assert.Contains(t, string(js), `"value":`)
	}
	// Unmarshal.
	err = cdc.UnmarshalJSON(js, dst)
	require.Nil(t, err, "%+v", err)
}

func ExamplePrintRegisteredTypes() {
	cdc.PrintTypes(os.Stdout)
	// Output: | Type | Name | Prefix | Length | Notes |
	//| ---- | ---- | ------ | ----- | ------ |
	//| PubKeyEd25519 | tendermint/PubKeyEd25519 | 0x1624DE64 | 0x20 |  |
	//| PubKeySecp256k1 | tendermint/PubKeySecp256k1 | 0xEB5AE987 | 0x21 |  |
	//| PrivKeyEd25519 | tendermint/PrivKeyEd25519 | 0xA3288910 | 0x40 |  |
	//| PrivKeySecp256k1 | tendermint/PrivKeySecp256k1 | 0xE1B0F79B | 0x20 |  |
	//| SignatureEd25519 | tendermint/SignatureEd25519 | 0x2031EA53 | 0x40 |  |
	//| SignatureSecp256k1 | tendermint/SignatureSecp256k1 | 0x7FC4A495 | variable |  |
}

func TestKeyEncodings(t *testing.T) {
	cases := []struct {
		privKey           PrivKey
		privSize, pubSize int // binary sizes
	}{
		{
			privKey:  GenPrivKeyEd25519(),
			privSize: 69,
			pubSize:  37,
		},
		{
			privKey:  GenPrivKeySecp256k1(),
			privSize: 37,
			pubSize:  38,
		},
	}

	for _, tc := range cases {

		// Check (de/en)codings of PrivKeys.
		var priv2, priv3 PrivKey
		checkAminoBinary(t, tc.privKey, &priv2, tc.privSize)
		assert.EqualValues(t, tc.privKey, priv2)
		checkAminoJSON(t, tc.privKey, &priv3, false) // TODO also check Prefix bytes.
		assert.EqualValues(t, tc.privKey, priv3)

		// Check (de/en)codings of Signatures.
		var sig1, sig2, sig3 Signature
		sig1, err := tc.privKey.Sign([]byte("something"))
		assert.NoError(t, err)
		checkAminoBinary(t, sig1, &sig2, -1) // Signature size changes for Secp anyways.
		assert.EqualValues(t, sig1, sig2)
		checkAminoJSON(t, sig1, &sig3, false) // TODO also check Prefix bytes.
		assert.EqualValues(t, sig1, sig3)

		// Check (de/en)codings of PubKeys.
		pubKey := tc.privKey.PubKey()
		var pub2, pub3 PubKey
		checkAminoBinary(t, pubKey, &pub2, tc.pubSize)
		assert.EqualValues(t, pubKey, pub2)
		checkAminoJSON(t, pubKey, &pub3, false) // TODO also check Prefix bytes.
		assert.EqualValues(t, pubKey, pub3)
	}
}

func TestNilEncodings(t *testing.T) {

	// Check nil Signature.
	var a, b Signature
	checkAminoJSON(t, &a, &b, true)
	assert.EqualValues(t, a, b)

	// Check nil PubKey.
	var c, d PubKey
	checkAminoJSON(t, &c, &d, true)
	assert.EqualValues(t, c, d)

	// Check nil PrivKey.
	var e, f PrivKey
	checkAminoJSON(t, &e, &f, true)
	assert.EqualValues(t, e, f)

}

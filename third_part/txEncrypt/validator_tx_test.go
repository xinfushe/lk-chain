package data_encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignValidatorTx(t *testing.T) {
	assert := assert.New(t)

	// create key
	mprikey := make(map[string]string)
	mpubkey := make(map[string]int)
	
	mprikey["0114160065BA2E54F297FD1AB25C4A484687F5FD98C664D0D14FD7B681351A7C7039560D39C3654A0506068C6E6701BC931757C96E1272FD1632979044E7E039EE"] = "0139560D39C3654A0506068C6E6701BC931757C96E1272FD1632979044E7E039EE"
	mprikey["01AA6601B6B219528728506D60011D7C2A4A335C76EB43FD54DE5DB16857EE8E1A21A6F87A1E7E6547B8BBD4FFBDC6A0D5A31DB7F44149A43E35FA08B105692763"] = "0121A6F87A1E7E6547B8BBD4FFBDC6A0D5A31DB7F44149A43E35FA08B105692763"
	mprikey["01B7677511CC0B9934AE12D21ACEE71CD13238CA1472168E0DE08AD3AD359877374E26D4568CF62B959F4B2DBD42CCF56E959E7D7CABD23EA44A555F3F10CFAC0A"] = "014E26D4568CF62B959F4B2DBD42CCF56E959E7D7CABD23EA44A555F3F10CFAC0A"
	mprikey["011BA781520239A46E0A75D6DA62352901E3E67CA826D22B9AE975085D367CBE7672316B80CDAE0AF9BC59264BF1A2367720C103165A2341B3973A384ACB988C10"] = "0172316B80CDAE0AF9BC59264BF1A2367720C103165A2341B3973A384ACB988C10"
	mpubkey["0139560D39C3654A0506068C6E6701BC931757C96E1272FD1632979044E7E039EE"] = 10
	mpubkey["0121A6F87A1E7E6547B8BBD4FFBDC6A0D5A31DB7F44149A43E35FA08B105692763"] = 10
	mpubkey["014E26D4568CF62B959F4B2DBD42CCF56E959E7D7CABD23EA44A555F3F10CFAC0A"] = 10
	mpubkey["0172316B80CDAE0AF9BC59264BF1A2367720C103165A2341B3973A384ACB988C10"] = 10

	// create validator tx
	var item ValidatorItemTx
	var txs  ValidatorTx

	item.PubKey = "1234"
	txs.Txs = append(txs.Txs, item)

	//sign tx
	sign, err := SignValidatorTx(txs, mprikey)
	assert.NoError(err)

	// verify tx
	if err == nil {
		_, err = ParseValidatorTx(sign, mpubkey)
		assert.NoError(err)
	}
}

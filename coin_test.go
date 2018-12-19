package faker

import "testing"

func TestCoinBitcoin(t *testing.T) {
	if len(GetCoin().Bitcoin()) == 0 {
		t.Fail()
		t.Logf("bitcoin address empty")
	}
	if GetCoin().Bitcoin()[0] != byte('1') && GetCoin().Bitcoin()[0] != byte('3') {
		t.Fail()
		t.Logf("bitcoin address first letter %s wrong", GetCoin().Bitcoin())
	}
}

func TestCoinEthereum(t *testing.T) {
	if len(GetCoin().Ethereum()) != 42 {
		t.Fail()
		t.Logf("ethereum address %s not 42 chars long", GetCoin().Ethereum())
	}
}

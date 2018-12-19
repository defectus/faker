package faker

import (
	"crypto/rand"
	"strings"
)

// Coin returns wallet address of several blockchain coins.
type Coin interface {
	Bitcoin() string
	Ethereum() string
	Litecoin() string
	Monero() string
	Ripple() string
	EOS() string
	Stellar() string
	Cardano() string
	IOTA() string
}

var coin Coin

// GetCoin returs a struct that implements the Coin interface.
func GetCoin() Coin {
	mu.Lock()
	defer mu.Unlock()

	if coin == nil {
		coin = &DefaultCoin{}
	}
	return coin
}

// DefaultCoin is the default implementation of the Coin interface.
type DefaultCoin struct{}

// Bitcoin returns a fake bitcoin address.
func (db *DefaultCoin) Bitcoin() string { return "1" + generateRandomString(30, RandTypeAlphanumeric, false) }

// Ethereum returns a fake ethereum address.
func (db *DefaultCoin) Ethereum() string { return "0x" + generateRandomString(40, RandTypeAlphanumeric, true) }

// Litecoin returns a fake litecoin address.
func (db *DefaultCoin) Litecoin() string { return "" }

// Monero returns a fake Monero address.
func (db *DefaultCoin) Monero() string { return "" }

// Ripple returns a fake Ripple address.
func (db *DefaultCoin) Ripple() string { return "" }

// EOS returns a fake EOS address.
func (db *DefaultCoin) EOS() string { return "" }

// Stellar returns a fake Stellar address.
func (db *DefaultCoin) Stellar() string { return "" }

// Cardano returns a fake Cardano address.
func (db *DefaultCoin) Cardano() string { return "" }

// IOTA returns a fake IOTA address.
func (db *DefaultCoin) IOTA() string { return "" }

const (
	// RandTypeAlphanumeric for random string that contains both letters and numbers.
	RandTypeAlphanumeric = "alphanumeric"
	// RandTypeAlpha for random strings that consists of letters only.
	RandTypeAlpha = "alpha"
	// RandTypeNumeric for random strings that consists of digits only.
	RandTypeNumeric = "numeric"
)

// GenerateRandomString generates random string of given size and type.
// For type use one of RandType... constants.
func generateRandomString(strSize int, randType string, lowerCase bool) string {
	var dictionary string
	switch randType {
	case RandTypeAlphanumeric:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandTypeAlpha:
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandTypeNumeric:
		dictionary = "0123456789"
	}
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	if lowerCase {
		return strings.ToLower(string(bytes))
	}
	return string(bytes)
}

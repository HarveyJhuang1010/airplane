package rand

import (
	"crypto/rand"
	"encoding/binary"
	"golang.org/x/exp/constraints"
	"math/big"
)

const (
	onlyNumber        = "0123456789"
	letterAndNumber   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base32StdEncoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	base32HexEncoding = "0123456789ABCDEFGHIJKLMNOPQRSTUV"
)

// LetterAndNumberString returns, as a string, random string in the char set [0-9a-zA-Z] from crypto/rand. It panics if n < 0.
func LetterAndNumberString(n int) string {
	return Selection(n, letterAndNumber)
}

// NumberString returns, as a string, random string in the char set [0-9] from crypto/rand. It panics if n < 0.
func NumberString(n int) string {
	return Selection(n, onlyNumber)
}

// Intn returns, as a T, a non-negative random number in the half-open interval [0,n) from crypto/rand. It panics if n <= 0.
func Intn[T constraints.Signed](n T) T {
	if n < 0 {
		panic("invalid argument to Intn")
	}

	if output, err := rand.Int(rand.Reader, big.NewInt(int64(n))); err != nil {
		panic(err)
	} else {
		return T(output.Int64())
	}
}

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0)
func Float64() float64 {
	var buf [8]byte
	// 7 bytes is enough for 53 bits of precision.
	if _, err := rand.Read(buf[:7]); err != nil {
		panic(err)
	}
	// Mask off the unwanted bits to keep only 53 bits for precision.
	u64 := binary.BigEndian.Uint64(buf[:]) & ((1 << 53) - 1)
	return float64(u64) / float64(1<<53)
}

func Base32StdString(n int) string {
	return Selection(n, base32StdEncoding)
}

func Base32HexString(n int) string {
	return Selection(n, base32HexEncoding)
}

// Selection 打亂dictionary ，依據size 從dictionary 挑出字元並回傳
func Selection(length int, dictionary string) string {
	if length == 0 || len(dictionary) == 0 {
		return ""
	}

	var bytes = make([]byte, length)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}

	return string(bytes)
}

package baseutils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

// Empty
// define an empty interface literal type to pass null as a parameter
var Empty any

// IsMap
// check if the informed interface is an map of string to interface
func IsMap(v any) bool {
	switch v.(type) {
	case map[string]any:
		return true
	}

	return false
}

// RandomString
// generate a random string of the passed length using the desired chars
func RandomString(length int, upperCase, lowerCase, numbers bool) string {

	// prepare the charset
	charset := ""

	if upperCase {
		charset = charset + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if lowerCase {
		charset = charset + "abcdefghijklmnopqrstuvwxyz"
	}

	if numbers {
		charset = charset + "0123456789"
	}

	// generate the random string using a CSPRNG and return
	// crypto/rand never returns an error since Go 1.24
	charsetLen := big.NewInt(int64(len(charset)))

	b := make([]byte, length)
	for i := range b {
		n, _ := rand.Int(rand.Reader, charsetLen)
		b[i] = charset[n.Int64()]
	}

	return string(b)
}

// GetKey
// get a key value of a interface within a string map
func GetKey(needle string, haystack map[string]any) any {
	if val, ok := haystack[needle]; ok {
		return val
	} else {
		return nil
	}
}

// NewLogger
// create a new logger using the parent logger prefix and writer
func NewLogger(l *log.Logger, prefix string) *log.Logger {
	return log.New(l.Writer(), fmt.Sprintf("%v%v ", l.Prefix(), prefix), log.LstdFlags|log.Lmsgprefix)
}

// TruncateText
// truncate a text and add ellipsis to its end
func TruncateText(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen < 3 {
		maxLen = 3
	}

	return string(runes[0:maxLen-3]) + "..."
}

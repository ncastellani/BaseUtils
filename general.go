package baseutils

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Empty
// define an empty interface literal type to pass null as a parameter
var Empty any

// StringInSlice
// check if the informed string is on the passed slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

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

	// generate the random string and return
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano() - int64(rand.Int())))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
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

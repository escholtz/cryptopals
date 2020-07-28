package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math"
)

func hexToBase64(x string) (string, error) {
	buf, err := hex.DecodeString(x)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(buf), nil
}

func xor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("lengths must be equal")
	}
	out := make([]byte, len(a))
	for i := range a {
		out[i] = a[i] ^ b[i]
	}
	return out, nil
}

// https://en.wikipedia.org/wiki/Letter_frequency
// used the chart further down, the one at the top doesn't add up to 1.0
// this sum is still off by a bit, but less error
// TODO: maybe norvigs is better? https://norvig.com/mayzner.html
var asciiFreq = map[byte]float64{
	'a': 0.08197,
	'b': 0.01492,
	'c': 0.02782,
	'd': 0.04253,
	'e': 0.12702,
	'f': 0.02228,
	'g': 0.02015,
	'h': 0.06094,
	'i': 0.06966,
	'j': 0.00153,
	'k': 0.00772,
	'l': 0.04025,
	'm': 0.02406,
	'n': 0.06749,
	'o': 0.07507,
	'p': 0.01929,
	'q': 0.00095,
	'r': 0.05987,
	's': 0.06327,
	't': 0.09056,
	'u': 0.02758,
	'v': 0.00978,
	'w': 0.02360,
	'x': 0.00150,
	'y': 0.01974,
	'z': 0.00074,
}

func byteFreq(x []byte) [256]float64 {
	out := [256]float64{}
	if len(x) == 0 {
		return out
	}
	for _, b := range x {
		out[b] += 1.0
	}
	n := float64(len(x))
	for i := range out {
		out[i] /= n
	}
	return out
}

func asciiFreqError(x [256]float64) float64 {
	out := 0.0
	for k, v := range asciiFreq {
		out += math.Abs(v - x[k])
	}
	return out
}

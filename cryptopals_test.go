package cryptopals

import (
	"bytes"
	"encoding/hex"
	"math"
	"testing"
)

func Test001(t *testing.T) {
	in := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got, err := hexToBase64(in)
	if err != nil {
		t.Errorf("got err %v; want nil", err)
	}
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

func Test002(t *testing.T) {
	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Errorf("a hex is invalid: %v", err)
	}

	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		t.Errorf("b hex is invalid: %v", err)
	}

	want, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if err != nil {
		t.Errorf("want hex is invalid: %v", err)
	}

	got, err := xor(a, b)
	if err != nil {
		t.Errorf("got err %v; want nil", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("got %x; want %x", got, want)
	}
}

func TestAsciiFreqSum(t *testing.T) {
	sum := 0.0
	for _, v := range asciiFreq {
		sum += v
	}
	if math.Abs(sum-1.0) > 1e-3 {
		t.Errorf("got %f; want 1.0", sum)
	}
}

func Test003(t *testing.T) {
	in, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Fatalf("input hex is invalid %v", err)
	}

	var best byte
	bestScore := 1e9
	msg := ""
	out := make([]byte, len(in))
	for i := byte(0); i < byte(255); i++ {
		for j := range in {
			out[j] = in[j] ^ i
		}
		e := asciiFreqError(byteFreq(out))
		if e < bestScore {
			bestScore = e
			best = i
			msg = string(out)
		}
	}
	want := "Cooking MC's like a pound of bacon"
	if msg != want {
		t.Errorf("got character %q; want X", best)
		t.Errorf("got msg %q; want %q", msg, want)
	}
}

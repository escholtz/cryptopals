package cryptopals

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"os"
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

func Test003(t *testing.T) {
	in, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		t.Fatalf("input hex is invalid %v", err)
	}

	best, _, msg := bestSingleByteXor(in)
	want := "Cooking MC's like a pound of bacon"
	if msg != want {
		t.Errorf("got character %q; want X", best)
		t.Errorf("got msg %q; want %q", msg, want)
	}
}

func Test004(t *testing.T) {
	f, err := os.Open("testdata/4.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	bestScore := 1e9
	bestLine := ""
	for scanner.Scan() {
		line := scanner.Text()
		buf, err := hex.DecodeString(line)
		if err != nil {
			t.Fatalf("could not decode hex %q err %v", line, err)
		}
		_, score, msg := bestSingleByteXor(buf)
		if score < bestScore {
			bestScore = score
			bestLine = line
			t.Log(score, msg, line)
		}
	}

	want := "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"
	if bestLine != want {
		t.Errorf("got line %s; want %s", bestLine, want)
	}
}

func Test005(t *testing.T) {
	in := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte(`ICE`)
	got := hex.EncodeToString(repeatingKeyXor(key, in))
	want := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`
	if got != want {
		t.Errorf("got %q; want %q", got, want)
	}
}

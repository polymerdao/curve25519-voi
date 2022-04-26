package ed25519

import (
	"fmt"
	"github.com/oasisprotocol/curve25519-voi/curve"
	"github.com/oasisprotocol/curve25519-voi/internal/field"
	"math/big"
	"testing"
)

func Reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func TestCircomEd25519(t *testing.T) {
	pointAx, _ := new(big.Int).SetString("43933056957747458452560886832567536073542840507013052263144963060608791330050", 10)
	pointAX := Reverse(pointAx.Bytes())

	pointAy, _ := new(big.Int).SetString("16962727616734173323702303146057009569815335830970791807500022961899349823996", 10)
	pointAY := Reverse(pointAy.Bytes())

	pointAt, _ := new(big.Int).SetString("47597536765056690778342994103149503974598380825968728087754575050160026478564", 10)
	pointAT := Reverse(pointAt.Bytes())

	var X, Y, Z, T field.Element
	_, err := X.SetBytes(pointAX)
	if err != nil {
		t.Fatal("X SetBytes error")
	}
	_, err = Y.SetBytes(pointAY)
	if err != nil {
		t.Fatal("Y SetBytes error")
	}
	Z.One()
	_, err = T.SetBytes(pointAT)
	if err != nil {
		t.Fatal("Z SetBytes error")
	}

	compressedY, err := curve.NewCompressedEdwardsYFromBytes(pointAY)
	if err != nil {
		t.Fatal("compressedY SetBytes error")
	}
	var P curve.EdwardsPoint
	_, err = P.SetCompressedY(compressedY)
	if err != nil {
		t.Fatal("P SetCompressedY error")
	}

	point := P.GetPointInner()
	if point[0].Equal(&X) != 1 || point[1].Equal(&Y) != 1 || point[2].Equal(&Z) != 1 || point[3].Equal(&T) != 1 {
		t.Fatal("")
	}

	fmt.Println(P.GetPointChunks())
}

package ebdprotocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestFlashUpgradeHostAppResponseParse(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"02464602413130302020202020304234423803", true},
		{"02464608413130302020202020303542323003", true},
	}
	for _, c := range cases {
		var res FlashUpgradeHostAppResponse
		input, _ := hex.DecodeString(c.in)
		res.Parse(input)
		got := res.bAllowed
		//		log.Printf("got : %v\t want : %v\n", got, c.want)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("TestFlashUpgradeHostAppResponseParse.Parse(),\nin\t%s\nwant\t|%v|got%v\n", c.in, c.want, got)
		}
	}
}

package ebdprotocol

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestFlashUpgradeHostBootLoaderResponseParse(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"02464601424630302020202020304343414303", true},
	}
	for _, c := range cases {
		var res FlashUpgradeHostBootLoaderResponse
		input, _ := hex.DecodeString(c.in)
		res.Parse(input)
		got := res.bAllowed
		//		log.Printf("got : %v\t want : %v\n", got, c.want)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("TestFlashUpgradeHostBootLoaderResponseParse.Parse(),\nin\t%s\nwant\t|%v|got%v\n", c.in, c.want, got)
		}
	}
}

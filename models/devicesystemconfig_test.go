package models

import (
	"reflect"
	"testing"
)

/*

4E41000000000000000000000000000000000000
4E41000000000000000000000000000000000000
4E41000000000000000000000000000000000000
4E41000000000000000000000000000000000000
4E410000000000000000
4E41 0000 0000 0000 0000
4E41 0000 0000 0000 0000 0000 0000 0000 0000 0000
20FFFFFF
72FFFFFF
0000FFFF

4E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E4100000000000000004E4100000000000000004E4100000000000000000000000000000000000020FFFFFF72FFFFFFDC81FFFF
4E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E4100000000000000004E4100000000000000004E4100000000000000000000000000000000000020FFFFFF72FFFFFF135FFFFF


52756279000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E4100000000000000004E4100000000000000004E4100000000000000000000000000000000000001FFFFFF01FFFFFF
52756279000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E410000000000000000000000000000000000004E4100000000000000004E4100000000000000004E4100000000000000000000000000000000000020FFFFFF72FFFFFFEE5EFFFF

*/

func TestDevicesystemconfigToBytes(t *testing.T) {
	cases := []struct {
		want []byte
	}{
		{[]byte("68134F46DA7854064109B8F532419C7D")},
	}
	for _, c := range cases {
		var sysconfig Devicesystemconfig
		sysconfig.DeviceName = "Ruby"
		sysconfig.SystemVersion = "NA"
		sysconfig.Devicesku = "NA"
		sysconfig.Serialnumber = "NA"
		sysconfig.Softwarebuild = "NA"
		sysconfig.Partnumber = "NA"
		sysconfig.Hardwareversion = "NA"

		got := sysconfig.ToByte()
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("DevicesystemconfigToBytes() : \n want\t %X\n got \t %X", c.want, got)
		}
	}

}

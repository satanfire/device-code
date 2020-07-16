/*
	author: huangping@ling.ai
	date: 2020/7/16
	desc: device code test
*/

package devicecode

import (
	"testing"
)

func TestGetSkuAndRegion(t *testing.T) {
	cases := []struct {
		udid   string
		sku    string
		region string
	}{
		{"M4N2ESLM2T", "lukamini", "CN"},
		{"HRJQRZUU", "lukahero", "CN"},
		{"HMJQRZUU", "lukahero", "MX"},
		{"M4N2ESLZ2T", "lukaheros", "CN"},
		{"MKN2ESLE2T", "lukababy", "KR"},
		{"MUN2ESLQQT", "luka", "CN"},
		{"HMJQRWUU", "luka", "MX"},
		{"HCJQRWU1", "luka", "CN"},
		{"HCJQRWU1-1", "unknown", "CN"},
	}

	for _, v := range cases {
		sku, region := GetSkuAndRegion(v.udid)
		if sku != v.sku || region != v.region {
			t.Errorf("udid:%s, except res: [sku=>%s, region=>%s], got res: [sku=>%s, region=>%s]", v.udid, v.sku, v.region, sku, region)
		} else {
			t.Logf("udid:%s, sku:%s, region:%s", v.udid, sku, region)
		}
	}
}

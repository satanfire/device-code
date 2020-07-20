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
		view   string
	}{
		{"M4N2ESLM2T", "lukamini", "CN", "Luka Mini"},
		{"HRJQRZUU", "lukahero", "CN", "Luka Hero"},
		{"HMJQRZUU", "lukahero", "MX", "Luka Hero"},
		{"M4N2ESLZ2T", "lukaheros", "CN", "Luka HeroS"},
		{"MKN2ESLE2T", "lukababy", "KR", "Luka Baby"},
		{"MUN2ESLQQT", "luka", "CN", "Luka"},
		{"HMJQRWUU", "luka", "MX", "Luka"},
		{"HCJQRWU1", "luka", "CN", "Luka"},
		{"HCJQRWU1-1", "unknown", "CN", "Unknown"},
	}

	for _, v := range cases {
		sku, region, view := GetSkuAndRegion(v.udid)
		if sku != v.sku || region != v.region || view != v.view {
			t.Errorf("udid:%s, except res: [sku=>%s, region=>%s, view=>%s], got res: [sku=>%s, region=>%s, view=>%s]", v.udid, v.sku, v.region, v.view, sku, region, view)
		} else {
			t.Logf("udid:%s, sku:%s, region:%s, view:%s", v.udid, sku, region, view)
		}
	}
}

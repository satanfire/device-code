/*
	author: huangping@ling.ai
	date: 2020/7/16
	desc: device code config parse
*/

package devicecode

func GetSkuAndRegion(udid string) (string, string) {
	sku, region := "luka", "CN"
	udidLen := len(udid)

	specialRegion := udid[udidLen-3 : udidLen-1]
	if value, ok := RegionCode[specialRegion]; ok {
		region = value
	} else {
		regionCode := udid[1:2]
		if value, ok := RegionCode[regionCode]; ok {
			region = value
		}
	}

	if value, ok := DeviceCode[specialRegion]; ok {
		sku = value
	} else {
		code := udid[udidLen-3 : udidLen-2]
		// 兼容旧版hero/heros设备解析规则
		if code == LukaProCode {
			if udidLen == 8 {
				code = LukaHeroCode
			} else if udidLen == 10 {
				code = LukaHeroSCode
			}
		}

		if value, ok := DeviceCode[code]; ok {
			sku = value
		} else {
			sku = "unknown"
		}
	}
	return sku, region
}

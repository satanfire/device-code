/*
	author: huangping@ling.ai
	date: 2020/7/16
	desc: device code config
*/

package devicecode

var DeviceCode map[string]string
var RegionCode map[string]string
var SkuView map[string]string

var LukaProCode, LukaHeroCode, LukaHeroSCode string

func init() {
	DeviceCode = make(map[string]string)
	RegionCode = make(map[string]string)
	SkuView = make(map[string]string)

	// 倒数第三位
	DeviceCode["W"] = "luka"
	DeviceCode["Z"] = "lukapro"
	DeviceCode["E"] = "lukababy"
	DeviceCode["G"] = "lukago"
	DeviceCode["M"] = "lukamini"
	DeviceCode["H"] = "lukahero"
	DeviceCode["S"] = "lukaheros"

	// 特殊批次，倒数第二三位
	DeviceCode["QQ"] = "luka"
	DeviceCode["DQ"] = "luka"
	DeviceCode["EQ"] = "luka"
	DeviceCode["WR"] = "luka"
	DeviceCode["WT"] = "luka"

	RegionCode["QQ"] = "CN"
	RegionCode["DQ"] = "CN"
	RegionCode["EQ"] = "CN"
	RegionCode["WR"] = "CN"
	RegionCode["WT"] = "CN"

	// 设备码第二位代表国家码
	RegionCode["U"] = "US"
	RegionCode["M"] = "MX"
	RegionCode["K"] = "KR"
	RegionCode["W"] = "DE"

	LukaProCode = "Z"
	LukaHeroCode = "H"
	LukaHeroSCode = "S"

	// sku展示
	SkuView["unknown"] = "Unknown"
	SkuView["luka"] = "Luka"
	SkuView["lukababy"] = "Luka Baby"
	SkuView["lukapro"] = "Luka Hero (S)"
	SkuView["lukahero"] = "Luka Hero"
	SkuView["lukaheros"] = "Luka HeroS"
	SkuView["lukamini"] = "Luka Mini"
	SkuView["lukago"] = "Luka Go"
}

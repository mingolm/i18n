package languagecode

import (
	"golang.org/x/text/language"
)

type Code string

var languageCodes []Code
var languageCodesMap = map[Code]struct{}{}

const (
	// 简体中文
	CN Code = "zh_CN"
	// 繁体中文
	TW Code = "zh_TW"
	// 英语（美国）
	EN Code = "en_US"
	// 日文
	JA Code = "ja_JP"
	// 韩文
	KO Code = "ko_KR"
	// 葡萄牙语
	PT Code = "pt_PT"
	// 越南语
	VI Code = "vi_VN"
	// 印度语
	IN Code = "hi_IN"
	// 印尼语
	ID Code = "id_ID"
	// 马来语
	MY Code = "ms_MY"
	// 泰语
	TH Code = "th_TH"
	// 西班牙
	ES Code = "es_ES"
	// 南非荷兰语
	AF Code = "af"
	// 阿姆哈拉语
	AM Code = "am"
	// 保加利亚语
	BG Code = "bg"
	// 加泰罗尼亚语
	CA Code = "ca"
	// 克罗地亚语
	HR Code = "hr"
	// 捷克语
	CS Code = "cs"
	// 丹麦语
	DA Code = "da"
	// 荷兰语
	NL Code = "nl"
	// 爱沙尼亚语
	ET Code = "et"
	// 菲律宾语
	FIL Code = "fil"
	// 芬兰语
	FI Code = "fi"
	// 法语
	FR Code = "fr"
	// 德语
	DE Code = "de"
	// 希腊语
	EL Code = "el"
	// 希伯来语
	HE Code = "he"
	// 匈牙利语
	HU Code = "hu"
	// 冰岛语
	IS Code = "is"
	// 意大利语
	IT Code = "it"
	// 拉脱维亚语
	LV Code = "lv"
	// 立陶宛语
	LT Code = "lt"
	// 挪威语
	NO Code = "no"
	// 波兰语
	PL Code = "pl"
	// 罗马尼亚语
	RO Code = "ro"
	// 俄语
	RU Code = "ru"
	// 塞尔维亚语
	SR Code = "sr"
	// 斯洛伐克语
	SK Code = "sk"
	// 斯洛文尼亚语
	SL Code = "sl"
	// 斯瓦希里语
	SW Code = "sw"
	// 瑞典语
	SV Code = "sv"
	// 土耳其语
	TR Code = "tr"
	// 乌克兰语
	UK Code = "uk"
	// 祖鲁语
	ZU Code = "zu"
)

var matcher language.Matcher

func init() {
	languageCodes = append(languageCodes, CN, TW, EN, JA, KO, PT, VI, IN, ID, MY, TH, ES, AF, AM, BG, CA, HR, CS, DA, NL, ET, FIL, FI, FR, DE, EL, HE, HU, IS, IT, LV, LT, NO, PL, RO, RU, SR, SK, SL, SW, SV, TR, UK, ZU)

	for _, tag := range languageCodes {
		languageCodesMap[tag] = struct{}{}
	}

	// for stdlib language matcher
	var stdLanguageTags []language.Tag
	for _, tag := range languageCodes {
		if stdLanguageTag, err := language.Parse(tag.String()); err != nil {
			panic(err)
		} else {
			stdLanguageTags = append(stdLanguageTags, stdLanguageTag)
		}
	}
	matcher = language.NewMatcher(stdLanguageTags)
}

func (t Code) Valid() bool {
	_, ok := languageCodesMap[t]
	return ok
}

func (t Code) String() string {
	return string(t)
}

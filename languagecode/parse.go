package languagecode

import (
	"strings"

	"golang.org/x/text/language"
)

func Parse(lang string) (code Code, conf language.Confidence) {
	// default code and confidence
	code = EN
	conf = language.No

	if lang == "" {
		return
	}

	_, ok := languageCodesMap[Code(lang)]
	if ok {
		return Code(lang), language.Exact
	}
	lang = strings.ToLower(lang)

	{ // 特殊处理
		switch lang {
		case "chs":
			return CN, language.Exact
		case "cht":
			return TW, language.Exact
		}
		switch {
		// @待讨论 https://xindong.slack.com/archives/C021D3LRVB5/p1646726501238349
		case strings.Contains(lang, "zh_sg"):
			return TW, language.Exact
		}
	}

	var script string
	if i := strings.IndexByte(lang, '#'); i != -1 {
		script = lang[i+1:]
		lang = lang[:i]
	}

	// 例如 zh_DZ#Hans, zh_DZ#Hant
	switch script {
	case "hans":
		return CN, language.Exact
	case "hant":
		return TW, language.Exact
	}

	var segments []string
	for _, s := range strings.Split(lang, "_") {
		if s != "" {
			segments = append(segments, s)
		}
	}

	for i := len(segments); i >= 0; i-- {
		l := strings.Join(segments[:i], "_")
		if tag, err := language.Parse(l); err == nil {
			_, index, conf := matcher.Match(tag)
			if conf != language.No {
				return languageCodes[index], conf
			}
		}
	}

	if script != "" {
		scriptCode, scriptConf := Parse(script)
		if scriptConf != language.No {
			return scriptCode, scriptConf
		}
	}

	region, err := language.ParseRegion(segments[len(segments)-1])
	if err != nil {
		return
	}

	regionTag, err := language.Compose(region)
	if err != nil {
		return
	}

	_, index, conf := matcher.Match(regionTag)
	if conf != language.No {
		return languageCodes[index], conf
	}
	return
}

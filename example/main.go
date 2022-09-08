package main

import (
	"embed"
	"fmt"
	"github.com/mingolm/i18n"
	"github.com/mingolm/i18n/languagecode"
)

//go:embed lang
var langResourceFS embed.FS

func main() {
	bundle := i18n.NewBundleFromEmbedFS(languagecode.CN, langResourceFS, "lang")
	localizer := i18n.NewLocalizer(languagecode.CN, bundle)

	fmt.Printf("get: %s, %s\n",
		localizer.Get(languagecode.CN, "country.cn"),
		localizer.Get(languagecode.EN, "country.us"),
	)

	fmt.Printf("get for not found: %s, %s\n",
		localizer.Get(languagecode.TW, "country.cn"),
		localizer.Get(languagecode.CN, "country.cn_no"),
	)

	fmt.Printf("all: %+v\n", localizer.All("country.cn"))
}

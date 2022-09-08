# i18n

支持 json 格式的多语言文件载入，实现国际化

1. 支持导入整个目录
2. 支持多语言参数注入

### 接入指南

```shell
go get github.com/mingolm/i18n
```

### 使用样例

```go
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

	// 中国
	fmt.Println(localizer.Get(languagecode.CN, "country.cn"))

	// map[en_US:china zh_CN:中国]
	fmt.Printf("%+v\n", localizer.All("country.cn"))

	// code invalid. your code is 649234
	fmt.Println(localizer.Get(languagecode.EN, "test.error.invalid_code",
		"code", "649234",
	))

	// name 参数错误，当前时间: 2022-09-08 21:44:04
	fmt.Println(localizer.Get(languagecode.CN, "test.error.invalid_argument",
		"arg", "name",
		"time", time.Now().Format("2006-01-02 15:04:05"),
	))

```
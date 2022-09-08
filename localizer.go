package i18n

import "github.com/mingolm/i18n/languagecode"

func NewLocalizer(code languagecode.Code, bundle Bundle) *Localizer {
	return &Localizer{
		bundle:       bundle,
		languageCode: code,
	}
}

type Localizer struct {
	bundle       Bundle
	languageCode languagecode.Code
}

func (st *Localizer) Get(lang languagecode.Code, key string, argsKeyAndValues ...any) string {
	return st.bundle.Get(lang, key, argsKeyAndValues...)
}

func (st *Localizer) All(key string, argsKeyAndValues ...any) map[languagecode.Code]string {
	return st.bundle.All(key, argsKeyAndValues...)
}

func (st *Localizer) Bundle() Bundle {
	return st.bundle
}

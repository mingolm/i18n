package i18n

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/mingolm/i18n/languagecode"
	"github.com/spf13/cast"
	"io/fs"
	"path/filepath"
	"strings"
)

type Bundle interface {
	Get(lang languagecode.Code, key string, argsKeysAndValues ...interface{}) string
	All(key string, argsKeysAndValues ...interface{}) map[languagecode.Code]string
	DefaultLanguage() languagecode.Code
}

func NewBundleFromEmbedFS(defaultLanguage languagecode.Code, embedFs embed.FS, root string) Bundle {
	if !defaultLanguage.Valid() {
		panic(fmt.Errorf("invalid language code %s", defaultLanguage))
	}
	b := &bundle{
		language:         defaultLanguage,
		messageTemplates: map[languagecode.Code]map[string]string{},
	}
	if err := b.loadFromEmbedFS(nil, embedFs, root, nil, nil); err != nil {
		panic(err.Error())
	}

	return b
}

type bundle struct {
	language         languagecode.Code
	messageTemplates map[languagecode.Code]map[string]string
}

func (st *bundle) Get(lang languagecode.Code, key string, argsKeysAndValues ...interface{}) string {
	messageTemplate, ok := st.messageTemplates[lang]
	if !ok {
		messageTemplate, ok = st.messageTemplates[st.DefaultLanguage()]
	}
	if !ok {
		return key
	}

	val, ok := messageTemplate[key]
	if !ok {
		return key
	}

	return st.transText(val, argsKeysAndValues...)
}

func (st *bundle) All(key string, argsKeysAndValues ...interface{}) map[languagecode.Code]string {
	results := make(map[languagecode.Code]string, len(st.messageTemplates))
	for lang, messageTemplate := range st.messageTemplates {
		val, ok := messageTemplate[key]
		if !ok {
			continue
		}
		results[lang] = st.transText(val, argsKeysAndValues...)
	}

	return results
}

func (st *bundle) DefaultLanguage() languagecode.Code {
	return st.language
}

func (st *bundle) transText(text string, argsKeysAndValues ...interface{}) string {
	for i := 0; i < len(argsKeysAndValues); i++ {
		k := argsKeysAndValues[i]
		v := argsKeysAndValues[i+1]
		text = strings.ReplaceAll(text, ":"+cast.ToString(k), cast.ToString(v))
		i++
	}

	return text
}

func (st *bundle) loadFromEmbedFS(results map[string]string, embedFs embed.FS, root string, entry fs.DirEntry, prefixes []string) error {
	if entry == nil {
		entries, err := embedFs.ReadDir(root)
		if err != nil {
			return err
		}
		for _, et := range entries {
			if err = st.loadFromEmbedFS(results, embedFs, root, et, prefixes); err != nil {
				return err
			}
		}
		return nil
	}

	entryBaseName := entry.Name()
	if !entry.IsDir() && filepath.Ext(entry.Name()) != ".json" {
		return nil
	}
	if !entry.IsDir() {
		entryBaseName = strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
	}

	var isTopLevel bool
	if results == nil {
		isTopLevel = true
		code := languagecode.Code(entryBaseName)
		if entry.IsDir() {
			root = filepath.Join(root, entry.Name())
		}
		if !code.Valid() {
			return fmt.Errorf("invalid language code %s", code)
		}
		if _, ok := st.messageTemplates[code]; !ok {
			st.messageTemplates[code] = map[string]string{}
		}
		results = st.messageTemplates[code]
	}

	if entry.IsDir() {
		if !isTopLevel {
			prefixes = append(prefixes, entryBaseName)
		}
		entries, err := embedFs.ReadDir(filepath.Join(append([]string{root}, prefixes...)...))
		if err != nil {
			return err
		}
		for _, entry = range entries {
			if err = st.loadFromEmbedFS(results, embedFs, root, entry, prefixes); err != nil {
				return err
			}
		}
		return nil
	}

	filename := filepath.Join(append([]string{root}, prefixes...)...)
	filename = filepath.Join(filename, entry.Name())

	bs, err := embedFs.ReadFile(filename)
	if err != nil {
		return err
	}

	if !isTopLevel {
		prefixes = append(prefixes, entryBaseName)
	}

	keyValues, err := parseJSON(bs)
	if err != nil {
		return fmt.Errorf("invalid language file %s: %w", filename, err)
	}

	for k, v := range keyValues {
		key := strings.Join(append(prefixes, k), ".")
		results[key] = v
	}

	return nil
}

func parseJSON(bs json.RawMessage) (map[string]string, error) {
	if len(bs) == 0 {
		return nil, nil
	}

	m := map[string]string{}

	keyValues := map[string]json.RawMessage{}
	if err := json.Unmarshal(bs, &keyValues); err != nil {
		return nil, err
	}

	for k, v := range keyValues {
		if len(v) == 0 {
			m[k] = ""
			continue
		}

		if v[0] == '"' {
			var value string
			if err := json.Unmarshal(v, &value); err != nil {
				return nil, err
			}
			m[k] = value
			continue
		}

		subMap, err := parseJSON(v)
		if err != nil {
			return nil, err
		}
		for subKey, subVal := range subMap {
			m[k+"."+subKey] = subVal
		}
	}

	return m, nil
}

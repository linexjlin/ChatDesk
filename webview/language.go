package webview

import (
	"fmt"

	"github.com/jeandeaual/go-locale"
)

var languages = map[string]map[string]string{
	"About": {
		"en": "About",
		"zh": "关于",
	},
}

var LANG = "en"

func setLang() {
	userLanguage, err := locale.GetLanguage()
	if err == nil {
		fmt.Println("Language:", userLanguage)
		LANG = userLanguage
	}
}

// give the right lanuage with query text when no match return the langText itself
func UText(langText string) string {
	if lang, ok := languages[langText]; ok {
		if text, ok := lang[LANG]; ok {
			return text
		}
	}
	if LANG != "en" {
		fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, LANG)
	}

	return langText
}

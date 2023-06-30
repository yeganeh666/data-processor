package i18next

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"IofIPOS/shared/contextext"
	"IofIPOS/shared/envext"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

var (
	configs   *Configs
	languages map[language.Tag]*i18n.Localizer
)

type Configs struct {
	DefaultLang language.Tag `env:"DEFAULT_LANG" envDefault:"en"`
}

func init() {
	configs = new(Configs)
	languages = make(map[language.Tag]*i18n.Localizer)

	if err := envext.Load(configs); err != nil {
		log.WithError(err).Fatal("can load i18n configs")
	}

	bundle := i18n.NewBundle(configs.DefaultLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if err := filepath.Walk("assets/locales", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, "toml") {
			return nil
		}
		if _, err = bundle.LoadMessageFile(path); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.WithError(err).Fatal("can not read locale files")
	}

	for _, tag := range bundle.LanguageTags() {
		languages[tag] = i18n.NewLocalizer(bundle, tag.String())
	}
}

func ByLangWithData(lang language.Tag, id string, data interface{}) string {
	local, ok := languages[lang]
	if !ok {
		if local, ok = languages[configs.DefaultLang]; !ok {
			return id
		}
	}
	str, err := local.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: data,
	})
	if err != nil {
		return id
	}
	return str
}

func ByLang(lang language.Tag, id string) string {
	return ByLangWithData(lang, id, nil)
}

func ByContextWithData(ctx context.Context, id string, data interface{}) string {
	lang, _ := contextext.GetLang(ctx)
	return ByLangWithData(lang, id, data)
}

func ByContext(ctx context.Context, id string) string {
	return ByContextWithData(ctx, id, nil)
}

package ui_helpers

import (
	"github.com/cloudfoundry/cli/cf/i18n"
	goi18n "github.com/nicksnyder/go-i18n/i18n"
	"path/filepath"
)

var T goi18n.TranslateFunc

func init() {
	T = i18n.Init(filepath.Join("cf", "ui_helpers"), i18n.GetResourcesPath())
}

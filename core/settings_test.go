package bots

import (
	"testing"
)

func TestNewBotSettings(t *testing.T) {
	const (
		code        = "TestBot1"
		token       = "TestToken2"
		localeCode5 = "Kode5"
	)
	bs := NewBotSettings(code, token, Locale{Code5: localeCode5})
	if bs.Code != code {
		t.Errorf("Unexpected code: %v", bs.Code)
	}
	if bs.Token != token {
		t.Errorf("Unexpected token: %v", bs.Token)
	}
	if bs.Locale.Code5 != localeCode5 {
		t.Errorf("Unexpected Locale: %v", bs.Locale.Code5)
	}
}

package config

import (
	"testing"

	"main.go/internal/flags"
)

func TestUpdateConfigWithFlagsOverridesValues(t *testing.T) {
	cfg := &Config{
		Name:   "original",
		Tag:    "sleepy",
		Say:    "hello",
		Filter: "sepia",
		Height: 100,
		Width:  200,
	}

	flag := flags.Flags{
		Name:   "new-name",
		Tag:    "cute",
		Say:    "hi",
		Filter: "mono",
		Height: 300,
		Width:  400,
	}

	updateConfigWithFlags(cfg, flag)

	if cfg.Name != "new-name" {
		t.Fatalf("expected Name to be overridden, got %q", cfg.Name)
	}
	if cfg.Tag != "cute" {
		t.Fatalf("expected Tag to be overridden, got %q", cfg.Tag)
	}
	if cfg.Say != "hi" {
		t.Fatalf("expected Say to be overridden, got %q", cfg.Say)
	}
	if cfg.Filter != "mono" {
		t.Fatalf("expected Filter to be overridden, got %q", cfg.Filter)
	}
	if cfg.Height != 300 {
		t.Fatalf("expected Height to be overridden, got %d", cfg.Height)
	}
	if cfg.Width != 400 {
		t.Fatalf("expected Width to be overridden, got %d", cfg.Width)
	}
}

func TestUpdateConfigWithFlagsSkipsZeroValues(t *testing.T) {
	cfg := &Config{
		Name:   "original",
		Tag:    "sleepy",
		Say:    "hello",
		Filter: "sepia",
		Height: 100,
		Width:  200,
	}

	flag := flags.Flags{}

	updateConfigWithFlags(cfg, flag)

	if cfg.Name != "original" {
		t.Fatalf("expected Name to remain, got %q", cfg.Name)
	}
	if cfg.Tag != "sleepy" {
		t.Fatalf("expected Tag to remain, got %q", cfg.Tag)
	}
	if cfg.Say != "hello" {
		t.Fatalf("expected Say to remain, got %q", cfg.Say)
	}
	if cfg.Filter != "sepia" {
		t.Fatalf("expected Filter to remain, got %q", cfg.Filter)
	}
	if cfg.Height != 100 {
		t.Fatalf("expected Height to remain, got %d", cfg.Height)
	}
	if cfg.Width != 200 {
		t.Fatalf("expected Width to remain, got %d", cfg.Width)
	}
}

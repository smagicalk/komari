package flags

import (
	"log/slog"
	"testing"
)

func TestNormalizeDatabaseType(t *testing.T) {
	tests := map[string]string{
		"":          DatabaseTypeSQLite,
		"sqlite":    DatabaseTypeSQLite,
		" SQLite ":  DatabaseTypeSQLite,
		"SQLITE":    DatabaseTypeSQLite,
		"mysql":     DatabaseTypeMySQL,
		" MYSQL ":   DatabaseTypeMySQL,
		"postgres":  "postgres",
		" postgres": "postgres",
	}

	for input, want := range tests {
		if got := NormalizeDatabaseType(input); got != want {
			t.Fatalf("NormalizeDatabaseType(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestSupportedDatabaseTypes(t *testing.T) {
	if got := SupportedDatabaseTypes(); got != "sqlite, mysql, postgres" {
		t.Fatalf("SupportedDatabaseTypes() = %q", got)
	}
}

func TestNormalizeLogLevel(t *testing.T) {
	tests := map[string]string{
		"":         "",
		" debug ":  "debug",
		"INFO":     "info",
		" Warning": "warning",
	}

	for input, want := range tests {
		if got := NormalizeLogLevel(input); got != want {
			t.Fatalf("NormalizeLogLevel(%q) = %q, want %q", input, got, want)
		}
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input string
		want  slog.Level
	}{
		{input: "", want: slog.LevelInfo},
		{input: "debug", want: slog.LevelDebug},
		{input: "info", want: slog.LevelInfo},
		{input: "warn", want: slog.LevelWarn},
		{input: "warning", want: slog.LevelWarn},
		{input: "error", want: slog.LevelError},
	}

	for _, tt := range tests {
		got, err := ParseLogLevel(tt.input)
		if err != nil {
			t.Fatalf("ParseLogLevel(%q) returned error: %v", tt.input, err)
		}
		if got != tt.want {
			t.Fatalf("ParseLogLevel(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseLogLevelInvalid(t *testing.T) {
	if _, err := ParseLogLevel("trace"); err == nil {
		t.Fatal("ParseLogLevel(trace) expected error")
	}
}

package cnf_test

import (
	"os"
	"path/filepath"
	"testing"

	"bedrok/cnf"

	"github.com/spf13/viper"
)

func setup(t *testing.T) {
	t.Helper()
	viper.Reset()
	t.Cleanup(func() {
		viper.Reset()
	})
}

func writeTempYAML(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

const fullYAML = `
server:
  host: "localhost"
  port: 8080

redis:
  host: "redis-host"
  port: 6379
  password: "secret"
  db: 1
  pool_size: 10
  min_idle_conns: 5

db:
  host: "db-host"
  port: 5432
  user: "postgres"
  password: "pgpass"
  name: "bedrok"
  sslmode: "disable"
  max_open_conns: 25
  max_idle_conns: 5

log:
  level: "info"
  format: "json"
`

func TestLoad_ValidConfig(t *testing.T) {
	setup(t)
	path := writeTempYAML(t, fullYAML)

	cfg, err := cnf.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name string
		got  any
		want any
	}{
		{"Server.Host", cfg.Server.Host, "localhost"},
		{"Server.Port", cfg.Server.Port, 8080},
		{"Redis.Host", cfg.Redis.Host, "redis-host"},
		{"Redis.Port", cfg.Redis.Port, 6379},
		{"Redis.Password", cfg.Redis.Password, "secret"},
		{"Redis.DB", cfg.Redis.DB, 1},
		{"Redis.PoolSize", cfg.Redis.PoolSize, 10},
		{"Redis.MinIdleConns", cfg.Redis.MinIdleConns, 5},
		{"DB.Host", cfg.DB.Host, "db-host"},
		{"DB.Port", cfg.DB.Port, 5432},
		{"DB.User", cfg.DB.User, "postgres"},
		{"DB.Password", cfg.DB.Password, "pgpass"},
		{"DB.Name", cfg.DB.Name, "bedrok"},
		{"DB.SSLMode", cfg.DB.SSLMode, "disable"},
		{"DB.MaxOpenConns", cfg.DB.MaxOpenConns, 25},
		{"DB.MaxIdleConns", cfg.DB.MaxIdleConns, 5},
		{"Log.Level", cfg.Log.Level, "info"},
		{"Log.Format", cfg.Log.Format, "json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestLoad_FileNotFound(t *testing.T) {
	setup(t)

	_, err := cnf.Load("/nonexistent/path/config.yaml")
	if err == nil {
		t.Fatal("expected error for non-existent file, got nil")
	}
}

func TestLoad_InvalidYAML(t *testing.T) {
	setup(t)
	path := writeTempYAML(t, "{\n  invalid yaml: [unclosed")

	_, err := cnf.Load(path)
	if err == nil {
		t.Fatal("expected error for invalid YAML, got nil")
	}
}

func TestLoad_EmptyFile(t *testing.T) {
	setup(t)
	path := writeTempYAML(t, "")

	cfg, err := cnf.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Server.Port != 0 {
		t.Errorf("Server.Port = %d, want 0", cfg.Server.Port)
	}
	if cfg.Server.Host != "" {
		t.Errorf("Server.Host = %q, want empty string", cfg.Server.Host)
	}
}

func TestLoad_PartialConfig(t *testing.T) {
	setup(t)
	path := writeTempYAML(t, "server:\n  port: 9090\n")

	cfg, err := cnf.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Server.Port != 9090 {
		t.Errorf("Server.Port = %d, want 9090", cfg.Server.Port)
	}
	if cfg.DB.Host != "" {
		t.Errorf("DB.Host = %q, want empty string", cfg.DB.Host)
	}
	if cfg.Redis.PoolSize != 0 {
		t.Errorf("Redis.PoolSize = %d, want 0", cfg.Redis.PoolSize)
	}
}

func TestLoad_ReturnsPointer(t *testing.T) {
	setup(t)
	path := writeTempYAML(t, fullYAML)

	cfg, err := cnf.Load(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg == nil {
		t.Fatal("expected non-nil pointer")
	}
}

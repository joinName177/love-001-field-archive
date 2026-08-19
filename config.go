package fieldarchive

import "os"

type Config struct{ Address, DataPath string }

func LoadConfig() Config {
	a := os.Getenv("FIELD_ARCHIVE_ADDR")
	if a == "" {
		a = ":8080"
	}
	p := os.Getenv("FIELD_ARCHIVE_DATA")
	if p == "" {
		p = "/data/archive.json"
	}
	return Config{a, p}
}

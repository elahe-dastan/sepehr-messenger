package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	Address string `koanf:"address"`
}

func Read() Config {
	var k = koanf.New(".")

	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatal("error loading config: %v", err)
	}

	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Println("No config file provided")
	}

	const Prefix = "alibaba_"

	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "_", ".", -1)
	}), nil); err != nil {
		log.Println("No env variable provided")
	}

	var c Config

	if err := k.Unmarshal("", &c); err != nil {
		log.Fatal("error unmarshalling config: %s", err)
	}

	return c
}

package go_kit

import (
	"reflect"
	"testing"
)

type Config struct {
	BaseUrl      string `field:"BASE_URL" default:"http://example.com"`
	ExpiresAt    uint   `field:"EXPIRES_AT" default:"8600"`
	IsProduction bool   `field:"IS_PRODUCTION" default:"true"`
	MaxCount     uint8  `field:"MAX_COUNT" default:"4"`
	ApiKey       string `field:"API_KEY" default:""`
	Secret       string `field:"SECRET" default:"qwertyuiop"`
	Debug        bool   `field:"DEBUG" default:"false"`
}

type Config2 struct {
	BaseUrl      string `field:"BASE_URL" default:"http://example.com"`
	ExpiresAt    uint   `field:"EXPIRES_AT" default:"8600"`
	IsProduction bool   `field:"IS_PRODUCTION" default:"true"`
	MaxCount     uint8  `field:"MAX_COUNT" default:"4"`
	ApiKey       string `field:"API_KEY" default:""`
	Secret       string `field:"SECRET" default:"qwertyuiop"`
	Hostname     string `field:"HOSTNAME" default:"localhost"`
}

type CombinedConfig struct {
	BaseUrl      string `field:"BASE_URL" default:"http://example.com"`
	ExpiresAt    uint   `field:"EXPIRES_AT" default:"8600"`
	IsProduction bool   `field:"IS_PRODUCTION" default:"true"`
	MaxCount     uint8  `field:"MAX_COUNT" default:"4"`
	ApiKey       string `field:"API_KEY" default:""`
	Secret       string `field:"SECRET" default:"qwertyuiop"`
	Hostname     string `field:"HOSTNAME" default:"localhost"`
	Debug        bool   `field:"DEBUG" default:"false"`
}

func TestLoad(t *testing.T) {
	var config Config
	err := LoadEnv(&config, "./testdata/.env")
	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		BaseUrl:      "http://myapi.com",
		ExpiresAt:    12000,
		IsProduction: false,
		MaxCount:     2,
		ApiKey:       "",
		Secret:       "qwertyuiop",
		Debug:        true,
	}

	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Config does not match expected.\nGot:\n%#v\nExpected:\n%#v", config, expected)
	}
}

func TestLoadCustomFile(t *testing.T) {
	var config Config2
	err := LoadEnv(&config, "./testdata/.env.prod")
	if err != nil {
		t.Fatal(err)
	}

	expected := Config2{
		BaseUrl:      "http://myprodapi.com",
		ExpiresAt:    7200,
		IsProduction: false,
		MaxCount:     2,
		ApiKey:       "",
		Secret:       "qwertyuiop",
		Hostname:     "github",
	}

	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Config does not match expected.\nGot:\n%#v\nExpected:\n%#v", config, expected)
	}
}

func TestLoadMultipleFiles(t *testing.T) {
	var config CombinedConfig
	err := LoadEnv(&config, "./testdata/.env", "./testdata/.env.prod")
	if err != nil {
		t.Fatal(err)
	}

	expected := CombinedConfig{
		BaseUrl:      "http://myprodapi.com",
		ExpiresAt:    7200,
		IsProduction: false,
		MaxCount:     2,
		ApiKey:       "",
		Secret:       "qwertyuiop",
		Hostname:     "github",
		Debug:        true,
	}

	if !reflect.DeepEqual(config, expected) {
		t.Errorf("Config does not match expected.\nGot:\n%#v\nExpected:\n%#v", config, expected)
	}
}

package ConfigUtils_test

import (
	"errors"
	"github.com/KairosSystems/ConfigUtils"
	"testing"
)

type ConfigTest struct {
	Name string `json:"name"`
}

func TestReadConfigNotExistsFile(t *testing.T) {
	cfg := &ConfigTest{}
	err := ConfigUtils.ReadConfig(&cfg, "fake_name_config_file.json")
	if err == nil {
		t.Error(errors.New("the config file doesnt exists"))
	}
	if cfg.Name != "" {
		t.Error("the name field must be \"\"")
	}
}

func TestReadBadJsonConfig(t *testing.T) {
	cfg := &ConfigTest{}
	err := ConfigUtils.ReadConfig(&cfg, "bad_config_test.json")
	if err == nil {
		t.Error("The file must be malformed")
	}
}

func TestReadConfigExistsFile(t *testing.T) {
	cfg := &ConfigTest{}
	err := ConfigUtils.ReadConfig(&cfg, "config_test.json")
	if err != nil {
		t.Error(err)
	}
	if cfg.Name != "Testing" {
		t.Error("the name field must be \"Testing\"")
	}
}

func TestWriteConfig(t *testing.T) {
	cfg := &ConfigTest{Name: "Testing"}
	err := ConfigUtils.WriteConfig(&cfg, "out_config.json")
	if err != nil {
		t.Error(err)
	}
}

func TestWritableConfig(t *testing.T) {
	cfg := &ConfigTest{Name: "Testing"}
	err := ConfigUtils.WriteConfig(&cfg, "unwritable_config_test.json")
	if err == nil {
		t.Error("The file must be unwritable")
	}
}

func TestAppKeyGen(t *testing.T) {
	AppKey := ConfigUtils.AppKeyGen()
	if len(AppKey) == 0 {
		t.Error("we can't create app key")
	}
}

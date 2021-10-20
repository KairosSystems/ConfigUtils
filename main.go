package ConfigUtils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/tinode/jsonco"
)

func ReadConfig(c interface{}, filename string) error {
	if FileExists(filename) {
		file, _ := os.Open(filename)
		jr := jsonco.New(file)
		err := json.NewDecoder(jr).Decode(&c)
		if err!=nil {
			return err
		}
		return nil
	} else {
		return errors.New("config file " + filename + " not found")
	}
}

func WriteConfig(c interface{}, filename string) error{
	str,_ := json.MarshalIndent(c,"","    ")
	err := ioutil.WriteFile(filename, str, 0644)
	if err != nil {
		return err
	}
	return nil
}

func AppKeyGen() string{
	key := make([]byte, 32)
	rand.Read(key)
	return base64.StdEncoding.EncodeToString(key)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}

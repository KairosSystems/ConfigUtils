package ConfigUtils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
)

func ReadConfig(c interface{}, filename string) error {
	if FileExists(filename) {
		file, _ := ioutil.ReadFile(filename)
		err := json.Unmarshal(file, &c)
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
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

package util

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"os/user"
	"log"
	"strings"
)

func ExpandUser(path string) string {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		path = usr.HomeDir + "/" + path[2:]
	}
	return path
}

func ReadAllPath(path string) []byte {
	path = ExpandUser(path)
	f, _ := os.Open(path)
	b, _ := ioutil.ReadAll(f)
	return b
}

func JsonPath(path string, v interface{}) {
	json.Unmarshal(ReadAllPath(path), v)
}
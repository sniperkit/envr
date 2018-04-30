// https://stackoverflow.com/questions/24770403/go-write-struct-to-json-file-using-struct-fields-not-json-keys/24770435
package main

import (
	// "encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"time"
	// homedir "github.com/mitchellh/go-homedir"
)

func writeFile(content []byte) {
	target, err := getDir()
	if err != nil {
		log.Println(err.Error())
	}
	err = ioutil.WriteFile(target, content, 0644)
	if err != nil {
		log.Printf("Write to: %s", target)
	}
}

func timestamp() string {
	return time.Now().Format("20060102150405")
}

func getDir() (string, error) {
	usr, _ := user.Current()
	dir := path.Join(usr.HomeDir, "enventory")
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if oserr := os.Mkdir(dir, 0700); oserr != nil {
				return "", oserr
			}
		} else {
			return "", err
		}
	}
	return dir, nil
}

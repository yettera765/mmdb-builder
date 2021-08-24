package common

import (
	"log"
	"os"
	"path"
	"strings"
)

func Split(s string, r rune) []string {
	f := func(c rune) bool { return c == r }
	return strings.FieldsFunc(s, f)
}

func Mkdir(p string) {
	dir := path.Dir(p)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}

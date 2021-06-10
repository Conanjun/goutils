package common

import (
	"io/ioutil"
	"log"
)

func closelog() {
	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)
}

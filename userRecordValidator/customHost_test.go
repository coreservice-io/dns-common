package userRecordValidator

import (
	"log"
	"testing"
)

func Test_gen(t *testing.T) {
	log.Println(GenCnameSetting("www.domain.com", "pz1", "mesoncdn.com"))
	log.Println(GenCnameSetting("aaa.bbb.domain.com", "pz2", "mesoncdn.com"))
}

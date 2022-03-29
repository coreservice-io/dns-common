package spec00

import (
	"log"
	"testing"
)

func Test_parserSpec(t *testing.T) {
	log.Println(Parser("spec00-fikcbikcafkcdax-19-thisisthebinddomain.domain.com"))
	log.Println(Parser("spec00-fikbkakaxxxxxxx-06-exsief"))
	log.Println(Parser("spec00-bkbkbkbxxxxxxxx-11-sdfewfwfwef"))

	log.Println(Parser("spec00-bkbkbkbxxxxxxxx-02-sdfewfwfwef"))
	log.Println(Parser("spec00-bkbk"))
}

func Test_genSpec(t *testing.T) {
	log.Println(GenerateSpecStr("58.218.205.230", "thisisthebinddomain"))
	log.Println(GenerateSpecStr("58.1.0.0", "exsief"))
	log.Println(GenerateSpecStr("1.1.1.1", "sdfewfwfwef"))
}

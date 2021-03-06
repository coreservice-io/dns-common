package spec00

import (
	"log"
	"testing"
)

func Test_parserSpec(t *testing.T) {
	log.Println(Parser("spec00-bejkbackbfikcai-06-exsief-03-aaa-04-tttt.mesontracking.com"))
	//log.Println(Parser("spec00-fikbkakaxxxxxxx-06-exsief-03-aaa-04-tttt.sdff.com"))
	//log.Println(Parser("spec00-bkbkbkbxxxxxxxx"))
	//
	//log.Println(Parser("spec00-bkbkbkbxxxxxxxx-05-sf")) // length error
	//log.Println(Parser("spec00-bkbk"))                  // error
}

func Test_genSpec(t *testing.T) {
	log.Println(GenerateSpecStr("149.102.158.208", "thisisthebinddomain"))
	log.Println(GenerateSpecStr("58.1.0.0", "exsief", "aaa", "bbb", "cccc"))
	log.Println(GenerateSpecStr("1.1.1.1", "sdfewfwfwef"))
	log.Println(GenerateSpecStr("54.238.176.97"))
}

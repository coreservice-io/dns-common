package main

import (
	"log"

	"github.com/coreservice-io/dns-common/spec00"
)

func main() {
	//parser example
	log.Println(spec00.Parser("spec00-fikcbikcafkcdax-19-thisisthebinddomain.domain.com"))
	log.Println(spec00.Parser("spec00-fikbkakaxxxxxxx-06-exsief-03-aaa-04-tttt.sdff.com"))
	log.Println(spec00.Parser("spec00-bkbkbkbxxxxxxxx"))
	log.Println(spec00.Parser("spec00-bkbkbkbxxxxxxxx-05-sf")) // length error
	log.Println(spec00.Parser("spec00-bkbk"))                  // error

	//gen example
	log.Println(spec00.GenerateSpecStr("58.218.205.230", "thisisthebinddomain"))
	log.Println(spec00.GenerateSpecStr("58.1.0.0", "exsief", "aaa", "bbb", "cccc"))
	log.Println(spec00.GenerateSpecStr("1.1.1.1", "sdfewfwfwef"))
	log.Println(spec00.GenerateSpecStr("1.1.1.1"))
}

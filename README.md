# dns-common

## specific domain

using spec00

{spec00}-{ip_hash}{optional}.domain.com

optional format: -{2 digits representing content length}-{content}

example:

spec00-fikcbikcafkcdax-17-thisisthepullzone-03-abc-02-xy.domain.com

spec00-fikcbikcafkcdax-17-thisisthepullzone.domain.com

spec00-fikcbikcafkcdax.domain.com

#### important: spec length should not be greater than 63 

usage example:
```go
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
	log.Println(spec00.Parser("spec00-bkbkbkbxxxxxxxx-00--00--01-a"))
	log.Println(spec00.Parser("spec00-bkbkbkbxxxxxxxx-00--00--00-"))

	log.Println(spec00.Parser("spec00-bkbkbkbxxxxxxxx-05-sf")) // length error
	log.Println(spec00.Parser("spec00-bkbk"))                  // error

	//gen example
	log.Println(spec00.GenerateSpecStr("58.218.205.230", "thisisthebinddomain"))
	log.Println(spec00.GenerateSpecStr("58.1.0.0", "exsief", "aaa", "bbb", "cccc"))
	log.Println(spec00.GenerateSpecStr("1.1.1.1", "sdfewfwfwef"))
	log.Println(spec00.GenerateSpecStr("1.1.1.1", "", "", "a"))
	log.Println(spec00.GenerateSpecStr("1.1.1.1", "", "", ""))
	log.Println(spec00.GenerateSpecStr("1.1.1.1"))
}

```

## gen and check cname setting when using custom cert

example

gen setting

abc.customer.com (customer's domain who wants to have a certificate )

hosted.com (hosted.com is the domain already registered in the server , this should be provided to customer)

tag (sub domain tag provided by customer)


set cname record
1. customer add CNAME abc.customer.com => tag.hosted.com
2. customer add CNAME _acme-challenge.abc.customer.com => _acme-challenge.abc.tag.hosted.com
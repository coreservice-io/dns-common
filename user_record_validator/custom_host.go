package user_record_validator

import (
	"errors"
	"strings"
	"time"

	"github.com/miekg/dns"
)

var CustomHostDomainCnameErr = errors.New("host domain cname error")
var ChallengeCnameErr = errors.New("challenge cname error")

func LookupCNAME(src string) (dst []string, err error) {
	c := dns.Client{
		Timeout: 15 * time.Second,
	}

	var lastErr error
	// retry 3 times
	for i := 0; i < 3; i++ {
		m := dns.Msg{}
		m.SetQuestion(src+".", dns.TypeCNAME)
		r, _, err := c.Exchange(&m, "8.8.8.8:53")
		if err != nil {
			lastErr = err
			time.Sleep(1 * time.Second * time.Duration(i+1))
			continue
		}

		dst = []string{}
		for _, ans := range r.Answer {
			record, isType := ans.(*dns.CNAME)
			if isType {
				dst = append(dst, record.Target)
			}
		}
		lastErr = nil
		break
	}

	err = lastErr

	return
}

// GenCnameSetting
//  example
//	abc.customer.com (customer's domain who wants to have a certificate )
//	hosted.com (hosted.com is the domain already registered in the server , this should be provided to customer)
//	tag (sub domain tag provided by customer)
//	set cname record
//	1. customer add CNAME abc.customer.com => tag.hosted.com
//	2. customer add CNAME _acme-challenge.abc.customer.com => _acme-challenge.abc.tag.hosted.com
func GenCnameSetting(applyDomain string, txt_tag string, hostedDomain string) (hostDomainCname string, challengeRecord string, challengeTarget string) {
	hostDomainCname = txt_tag + "." + hostedDomain
	challengeRecord = "_acme-challenge." + applyDomain

	//split root domain   _acme-challenge.www.somedomain.com=> _acme-challenge.www
	parts := strings.Split(challengeRecord, ".")
	parts = parts[:len(parts)-2]
	pre := strings.Join(parts, ".")

	challengeTarget = pre + "." + txt_tag + "." + hostedDomain

	return
}

// CheckChallengeCnameCorrect
//  before apply must set dns record in customer's dns server
//  example
//	abc.customer.com (customer's domain who wants to have a certificate )
//	hosted.com (hosted.com is the domain already registered in the server , this should be provided to customer)
//	tag (sub domain tag provided by customer)
//	set cname record
//	1. customer add CNAME abc.customer.com => tag.hosted.com
//	2. customer add CNAME _acme-challenge.abc.customer.com => _acme-challenge.abc.tag.hosted.com
func CheckChallengeCnameCorrect(applyDomain string, txt_tag string, hostedDomain string) error {

	_, challengeRecord, challengeTarget := GenCnameSetting(applyDomain, txt_tag, hostedDomain)

	dest, err := LookupCNAME(challengeRecord)
	if err != nil {
		return err
	}
	if len(dest) == 0 {
		return errors.New("CNAME record not found")
	}

	if dest[0] != challengeTarget+"." {
		return ChallengeCnameErr
	}

	return nil
}

func CheckCnameCorrect(applyDomain string, txt_tag string, hostedDomain string) error {
	hostDomainCname, _, _ := GenCnameSetting(applyDomain, txt_tag, hostedDomain)

	dest, err := LookupCNAME(applyDomain)
	if err != nil {
		return err
	}
	if len(dest) == 0 {
		return errors.New("CNAME record not found")
	}

	if dest[0] != hostDomainCname+"." {
		return CustomHostDomainCnameErr
	}

	return nil
}

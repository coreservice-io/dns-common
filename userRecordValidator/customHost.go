package userRecordValidator

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
//  applyDomain: www.somedomain.com
//  pullZoneName: pullzonexxx
//  hostedDomain: mesoncdn.com
//  gen setting
//  hostDomainCname: pullzonexxx.mesoncdn.com
//  challengeRecord: _acme-challenge.www.somedomain.com
//  challengeTarget: _acme-challenge.www.pullzonexxx.mesoncdn.com
//  1. CNAME  www.somedomain.com => pullzonexxx.mesoncdn.com
//  2. CNAME  _acme-challenge.www.somedomain.com => _acme-challenge.www.pullzonexxx.mesoncdn.com
func GenCnameSetting(applyDomain string, pullZoneName string, hostedDomain string) (hostDomainCname string, challengeRecord string, challengeTarget string) {
	hostDomainCname = pullZoneName + "." + hostedDomain
	challengeRecord = "_acme-challenge." + applyDomain

	//split root domain   _acme-challenge.www.somedomain.com=> _acme-challenge.www
	parts := strings.Split(challengeRecord, ".")
	parts = parts[:len(parts)-2]
	pre := strings.Join(parts, ".")

	challengeTarget = pre + "." + pullZoneName + "." + hostedDomain

	return
}

// CheckChallengeCnameCorrect
//  before apply must set dns record in customer's dns server
//  example
//  applyDomain: www.somedomain.com
//  pullZoneName: pullzonexxx
//  hostedDomain: mesoncdn.com
//  1. CNAME  www.somedomain.com => pullzonexxx.mesoncdn.com
//  2. CNAME  _acme-challenge.www.somedomain.com => _acme-challenge.www.pullzonexxx.mesoncdn.com
func CheckChallengeCnameCorrect(applyDomain string, pullZoneName string, hostedDomain string) error {

	_, challengeRecord, challengeTarget := GenCnameSetting(applyDomain, pullZoneName, hostedDomain)

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

func CheckCnameCorrect(applyDomain string, pullZoneName string, hostedDomain string) error {
	hostDomainCname, _, _ := GenCnameSetting(applyDomain, pullZoneName, hostedDomain)

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

package userRecordValidator

import (
	"errors"
	"net"
	"strings"
)

var CustomHostDomainCnameErr = errors.New("host domain cname error")
var ChallengeCnameErr = errors.New("challenge cname error")

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

	//dest, err := net.LookupCNAME(applyDomain)
	//if err != nil {
	//	return err
	//}
	//
	//if dest != hostDomainCname+"." {
	//	return CustomHostDomainCnameErr
	//}

	dest, err := net.LookupCNAME(challengeRecord)
	if err != nil {
		return err
	}

	if dest != challengeTarget+"." {
		return ChallengeCnameErr
	}

	return nil
}

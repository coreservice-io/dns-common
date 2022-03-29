package spec00

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

const ipReg = "((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}"

//spec00-fikcbikcafkcdax-19-thisisthebinddomain.domain.com
func Parser(specStr string) (ip string, bindName string, err error) {
	if !strings.HasPrefix(specStr, "spec00-") {
		return "", "", errors.New("not node mapping spec")
	}

	//get tag
	if len(specStr) < 25 {
		return "", "", errors.New("length error")
	}
	tag := specStr[7:21]
	ip = tagToIp(tag)

	//get bindName
	bindNameLength, err := strconv.Atoi(specStr[23:25])
	if err != nil {
		return "", "", err
	}
	if len(specStr) < 26+bindNameLength {
		return "", "", errors.New("length error")
	}
	bindName = specStr[26 : 26+bindNameLength]

	return ip, bindName, nil
}

func GenerateSpecStr(ip string, bindName string) (string, error) {
	//check ip
	if strings.HasPrefix(ip, "0") {
		return "", errors.New("ip error")
	}
	match, err := regexp.MatchString(ipReg, ip)
	if err != nil || match == false {
		return "", errors.New("ip error")
	}

	//check bindName
	if bindName == "" {
		return "", errors.New("bindname error")
	}

	//ip to tag
	tag := ipToTag(ip)

	return fmt.Sprintf("spec00-%s-%02d-%s", tag, len(bindName), bindName), nil
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func ipToTag(ip string) string {
	result := []byte{'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x', 'x'}
	for i, v := range ip {
		switch v {
		case '.':
			result[i] = 'k'
		default:
			result[i] = byte(v + 49)
		}
	}
	return bytes2str(result)
}

func tagToIp(str string) string {
	result := ""
	for _, v := range str {
		switch v {
		case 'k':
			result = result + "."
		case 'x':
			break
		default:
			result = result + string(byte(v-49))
		}
	}
	return result
}

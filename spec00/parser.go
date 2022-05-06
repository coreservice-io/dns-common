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
func Parser(specStr string) (ip string, optionalStr []string, err error) {
	if !strings.HasPrefix(specStr, "spec00-") {
		return "", []string{}, errors.New("not node mapping spec")
	}

	specStr = strings.Split(specStr, ".")[0]

	//get tag
	if len(specStr) < 22 {
		return "", []string{}, errors.New("length error")
	}
	tag := specStr[7:21]
	ip = tagToIp(tag)

	optionalStr = []string{}
	passedLength := 22
	for {
		if len(specStr) < passedLength+3 {
			break
		}
		length, err := strconv.Atoi(specStr[passedLength+1 : passedLength+3])
		if err != nil {
			return "", []string{}, err
		}
		if len(specStr) < passedLength+4+length {
			return "", []string{}, errors.New("optional pair length error")
		}
		str := specStr[passedLength+4 : passedLength+4+length]
		optionalStr = append(optionalStr, str)
		passedLength += 4 + length
	}

	return ip, optionalStr, nil
}

func GenerateSpecStr(ip string, optionalStr ...string) (string, error) {
	//check ip
	if strings.HasPrefix(ip, "0") {
		return "", errors.New("ip error")
	}
	match, err := regexp.MatchString(ipReg, ip)
	if err != nil || match == false {
		return "", errors.New("ip error")
	}

	//ip to tag
	tag := ipToTag(ip)
	str := fmt.Sprintf("spec00-%s", tag)

	for _, v := range optionalStr {
		cStr := fmt.Sprintf("-%02d-%s", len(v), v)
		str += cStr
		if len(str) > 63 {
			return "", errors.New("spec string too long")
		}
	}

	return str, nil
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

package intersms

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

const (
	XSEND_API = "https://api-hk-v4.mysubmail.com/internationalsms/xsend"
)

type Xsend struct {
	appid  string
	appkey string

	signature string
	to        string
	project   string
	vars      []struct {
		key string
		val string
	}
	tag       string
	timestamp string
	signType  string
}

// init with appid and appkey
func Init(appid, appkey string) *Xsend {
	return &Xsend{
		appid:  appid,
		appkey: appkey,
	}
}

// set to address
func (Xsend *Xsend) SetAddress(to string) {
	Xsend.to = to
}

// set project id
func (Xsend *Xsend) SetProject(project string) {
	Xsend.project = project
}

func (Xsend *Xsend) Send() string {
	return Send(XSEND_API, Xsend)
}

func Send(XSEND_API string, Xsend *Xsend) string {
	return "test"
}

func BuildSignature(req map[string]string, appid, appkey, signtype string) string {
	var arg []string
	var keys []string
	for k := range req {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		arg = append(arg, key+"="+req[key])
	}
	argstr := strings.Join(arg, "&")
	argstr = appid + appkey + argstr + appid + appkey
	return SHA256(argstr)
}

func SortMapKey(data map[string]string) map[string]string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	r := make(map[string]string)
	for _, key := range keys {
		r[key] = data[key]
	}
	return r
}

func SHA256(str string) string {
	s := sha256.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

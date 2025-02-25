package intersms

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	XSEND_API    = "https://hk-api-v4.mysubmail.com/internationalsms/xsend"
	SIGN_VERSION = "2"
	SIGN_TYPE    = "sha256"
	CONTENT_TYPE = "application/json;charset=utf-8"
)

type Xsend struct {
	Appid       string `json:"appid"`
	appkey      string
	Signature   string            `json:"signature"`
	To          string            `json:"to"`
	Project     string            `json:"project"`
	Vars        map[string]string `json:"vars"`
	Tag         string            `json:"tag"`
	Timestamp   string            `json:"timestamp"`
	SignType    string            `json:"sign_type"`
	SignVersion string            `json:"sign_version"`
}

// init with appid and appkey
func NewXsend(appid, appkey string) *Xsend {
	return &Xsend{
		Appid:       appid,
		appkey:      appkey,
		Vars:        make(map[string]string),
		Timestamp:   fmt.Sprint(time.Now().UTC().Unix()),
		SignType:    SIGN_TYPE,
		SignVersion: SIGN_VERSION,
	}
}

// set to address
func (Xsend *Xsend) SetAddress(to string) {
	Xsend.To = to
}

// set project id
func (Xsend *Xsend) SetProject(project string) {
	Xsend.Project = project
}

// set vars
func (Xsend *Xsend) AddVar(key, val string) {
	Xsend.Vars[key] = val
}

// set tag
func (Xsend *Xsend) SetTag(tag string) {
	Xsend.Tag = tag
}

func (Xsend *Xsend) Send() (string, error) {
	Xsend.Signature = BuildSignature(Xsend)
	return Send(XSEND_API, Xsend)
}

func Send(XSEND_API string, Xsend *Xsend) (string, error) {
	data, _ := json.Marshal(Xsend)
	body := bytes.NewBuffer([]byte(data))

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	retstr, err := client.Post(XSEND_API, CONTENT_TYPE, body)

	if err != nil {
		return "", err
	}
	defer retstr.Body.Close()

	result, err := io.ReadAll(retstr.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func BuildSignature(data *Xsend) string {
	req := map[string]string{
		"to":           data.To,
		"project":      data.Project,
		"sign_version": data.SignVersion,
		"appid":        data.Appid,
		"timestamp":    data.Timestamp,
		"sign_type":    data.SignType,
	}
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
	argstr = data.Appid + data.appkey + argstr + data.Appid + data.appkey
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

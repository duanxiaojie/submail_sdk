package intersms

import (
	"testing"
)

const (
	APPID  = "60005"
	APPKEY = "440f691b53bf882390e78455fb8020ec"
)

func TestXsend(t *testing.T) {

	x := NewXsend(APPID, APPKEY)
	x.SetAddress("+8615201893074")
	x.SetProject("8sgWm")
	d, e := x.Send()

	if e != nil {
		t.Errorf("error : %s", e.Error())
	} else {
		t.Logf("result : %s", d)
	}
}

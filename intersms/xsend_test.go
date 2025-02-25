package intersms

import (
	"testing"
)

func TestXsend(t *testing.T) {

	x := NewXsend("appid", "appkey")
	x.SetAddress("+8615201893074")
	x.SetProject("8sgWm")
	d, e := x.Send()

	if e != nil {
		t.Errorf("error : %s", e.Error())
	} else {
		t.Logf("result : %s", d)
	}
}

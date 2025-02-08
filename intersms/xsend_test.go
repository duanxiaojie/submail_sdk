package intersms_test

import (
	"testing"

	"github.com/duanxiaojie/submail_sdk.git/intersms"
)

func TestXsend(t *testing.T) {
	x := &intersms.Xsend{}
	r := x.Send()
	t.Logf("result : %s", r)
}

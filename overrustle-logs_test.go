package overrustlelogs

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	logs, err := New("https://dgg.overrustlelogs.net/Bot_v2_Beta.txt")
	if err != nil {
		t.Error(err)
		return
	}
	for logs.Scan() {
		log, err := logs.Log()

		if err != nil {
			t.Error(err)
			return
		}

		if reflect.TypeOf(log.Message).Kind() != reflect.String {
			t.Fail()
		}
	}
}

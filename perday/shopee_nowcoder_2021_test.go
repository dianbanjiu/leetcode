package perday

import (
	"reflect"
	"testing"
)

func TestTransferName(t *testing.T) {
	in := "PascalCaseTest"
	out := []string{"PascalCaseTest", "pascalCaseTest", "pascal_case_test", "pascal-case-test"}
	actual := TransferName(in)
	if !reflect.DeepEqual(out, actual) {
		t.Errorf("except %v, but got %v", out, actual)
	}

}

package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVerifyMobile(t *testing.T) {
	mobile1 := "13568777777"
	mobile2 := "13568777"
	mobile3 := ""
	mobile4 := "45678912312"

	a := []string{mobile1, mobile2, mobile3, mobile4}

	for i, v := range a {
		if i == 0 {
			re := VerifyMobile(v)
			require.Equal(t, true, re)
		} else {
			re := VerifyMobile(v)
			require.Equal(t, false, re)
		}
	}
}

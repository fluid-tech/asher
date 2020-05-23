package api

import "testing"

type GeneralTest struct {
	in  string
	out string
}

func NewGeneralTest(in string, out string) *GeneralTest {
	return &GeneralTest{
		in:  in,
		out: out,
	}
}

func IterateAndTest(table []*GeneralTest, t *testing.T) {
	for i, element := range table {
		if element.in != element.out {
			t.Errorf("in test case %d expected '%s' found '%s'", i, element.in, element.out)
		}
	}
}

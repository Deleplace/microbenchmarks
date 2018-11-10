package switcherr

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkSwitchErrNil1(b *testing.B) {
	k = 0
	st, sot, other := 0, 0, 0
	for i := 0; i < b.N; i++ {
		err := SomeFunctionCall()
		if err != nil {
			switch err.(type) {
			case sometype:
				st++
			case someothertype:
				sot++
			default:
				other++
			}
		}
	}
	//log.Println(st, sot, other)
}

func BenchmarkSwitchErrNil2(b *testing.B) {
	k = 0
	st, sot, other := 0, 0, 0
	for i := 0; i < b.N; i++ {
		err := SomeFunctionCall()
		switch err.(type) {
		case nil:
			// noop
		case sometype:
			st++
		case someothertype:
			sot++
		default:
			other++
		}
	}
	//log.Println(st, sot, other)
}

type (
	sometype      struct{}
	someothertype struct{}
)

func (sometype) Error() string {
	return "sometype"
}

func (someothertype) Error() string {
	return "someothertype"
}

var (
	err1 error = sometype{}
	err2 error = someothertype{}
	err3 error = fmt.Errorf("err3")
)

var errs []error

func init() {
	M := 100
	errs = make([]error, M)
	for i := range errs {
		switch rand.Intn(4) {
		case 0:
			errs[i] = nil
		case 1:
			errs[i] = err1
		case 2:
			errs[i] = err2
		case 3:
			errs[i] = err3
		}
	}
}

var k = 0

func SomeFunctionCall() error {
	k++
	return errs[k%len(errs)]
}

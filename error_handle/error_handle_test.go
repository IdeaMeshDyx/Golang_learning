package errorhandle

import (
	"testing"
)

type BasicError struct {
	text string
}

func (e BasicError) Error() string {
	return e.text
}

func returnAError() error {
	return &BasicError{text: "returnAError"}
}

func Test_main(t *testing.T) {
	err := returnAError()
	if err != nil {
		funcErrHandler(err, "main")
	}

}

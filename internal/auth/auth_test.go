package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
)

type TypeOutput struct {
	value string
	err   error
}

func TestMytest(t *testing.T) {

	Header1 := http.Header{}
	Header2 := http.Header{}
	Header3 := http.Header{}
	Header1.Add("Authorization", "ApiKey")
	Header2.Add("Authorization", "")
	Header3.Add("Authorization", "ApiKey $dedfvggvbbgtfcd")
	log.Println("doudpu", Header3.Get("Authorization"))
	cases := []struct {
		input    http.Header
		expected TypeOutput
	}{
		{
			input: Header2,
			expected: TypeOutput{
				value: "",
				err:   ErrNoAuthHeaderIncluded,
			},
		},
		{
			input: Header1,
			expected: TypeOutput{
				value: "",
				err:   errors.New("malformed authorization header"),
			},
		},
		{
			input: Header3,
			expected: TypeOutput{
				value: "$dedfvggvbbgtfc",
				err:   nil,
			},
		},
	}

	for _, cse := range cases {
		actual, err := GetAPIKey(cse.input)
		val1 := fmt.Sprintf("%v", err)
		val2 := fmt.Sprintf("%v", cse.expected.err)
		if actual != cse.expected.value || val1 != val2 {
			t.Errorf("the result are not equal: %v vs %v",
				actual,
				cse.expected.value,
			)
			log.Println("this value is", val1 == val2)
			continue
		}

	}
}

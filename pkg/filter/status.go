package filter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eur0pa/ffuf/pkg/ffuf"
)

type StatusFilter struct {
	Value []int64
}

func NewStatusFilter(value string) (ffuf.FilterProvider, error) {
	var intvals []int64
	for _, sv := range strings.Split(value, ",") {
		if sv == "*" {
			intvals = append(intvals, 999)
		} else {
			intval, err := strconv.ParseInt(sv, 10, 0)
			if err != nil {
				return &StatusFilter{}, fmt.Errorf("Status filter or matcher (-fc / -mc): invalid value %s", value)
			}
			intvals = append(intvals, intval)
		}
	}
	return &StatusFilter{Value: intvals}, nil
}

func (f *StatusFilter) Filter(response *ffuf.Response) (bool, error) {
	for _, iv := range f.Value {
		if iv == response.StatusCode || iv == 999 {
			return true, nil
		}
	}
	return false, nil
}

func (f *StatusFilter) Repr() string {
	var strval []string
	for _, iv := range f.Value {
		strval = append(strval, strconv.Itoa(int(iv)))
	}
	return fmt.Sprintf("Response status: %s", strings.Join(strval, ","))
}

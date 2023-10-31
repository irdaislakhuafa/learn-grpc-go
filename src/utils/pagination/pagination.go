package pagination

import (
	"errors"

	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
)

func ParseFromParam(v parameter.PaginationParam) (parameter.PaginationParam, uint64, error) {
	offset := uint64(0)
	if v.Page <= 1 {
		offset = 0
	} else {
		offset = (v.Page - 1) * v.Limit
	}

	if v.Limit <= 0 {
		return parameter.PaginationParam{}, 0, errors.New("minimum limit is 1")
	}

	return v, offset, nil
}

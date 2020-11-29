package ddd

import "context"

//验证接口,主要用于cmd参数是否有效校验
type ValidatorCtx interface {
	Verify(ctx context.Context) (code int, err error)
}

type Validator interface {
	Verify() (err error)
}

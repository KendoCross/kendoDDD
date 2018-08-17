package domainser

const (
	Succe_Code = "0000"
)

type ApiResp struct {
	IsSuccess bool
	RetInfo   string
	RetCode   string
	ErrMsg    string
}

package errorext

type CodeError struct {
	Code int
	Msg  string
	err  error
}

func (e *CodeError) Error() string {
	msg := e.Msg
	if e.err != nil {
		msg += " " + e.err.Error()
	}
	return msg
}

func (e *CodeError) Unwrap() error {
	return e.err
}

func NewCodeError(code int, msg string, err error) error {
	return &CodeError{
		err:  err,
		Code: code,
		Msg:  msg,
	}
}

type ListMsgError struct {
	MsgMap map[string]string
}

func (e *ListMsgError) Error() string {
	msg := ""
	for _, item := range e.MsgMap {
		msg += item + ";"
	}
	return msg
}

func (e *ListMsgError) Unwrap() error {
	return nil
}

func NewListMsgError(msgs map[string]string) error {
	return &ListMsgError{msgs}
}

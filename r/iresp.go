package r

type IResp[T any] interface {
	Code() int32
	Msg() string
	Data() interface{}
	ReqId() string
}

func (resp *Resp[T]) Code() int32 {
	return resp.BizCode
}

func (resp *Resp[T]) Msg() string {
	return resp.BizMsg
}

func (resp *Resp[T]) Data() T {
	return resp.BizData
}

func (resp *Resp[T]) ReqId() string {
	return resp.BizReqId
}

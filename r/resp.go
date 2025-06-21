package r

import "github.com/google/uuid"

type Resp[T any] struct {
	BizCode  int32  `json:"code" binding:"required"`
	BizMsg   string `json:"msg" binding:"required"`
	BizData  T      `json:"data" binding:"required"`
	BizReqId string `json:"reqId" binding:"required"`
}

func New[T any](code int32, msg string, data T) Resp[T] {
	return Resp[T]{
		BizCode:  0,
		BizMsg:   msg,
		BizReqId: uuid.New().String(),
	}
}

func Succeed[T any](data T) Resp[T] {
	id := uuid.New()
	return Resp[T]{
		BizCode:  1,
		BizMsg:   "Succeed.",
		BizData:  data,
		BizReqId: id.String(),
	}
}

func Failed(msg string) Resp[any] {
	return Resp[any]{
		BizCode:  0,
		BizMsg:   msg,
		BizReqId: uuid.New().String(),
	}
}

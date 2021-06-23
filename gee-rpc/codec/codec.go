package codec

import "io"

// https://geektutu.com/post/geerpc-day1.html
// rpc编码
// err = client.Call("Arith.Multiply", args, &reply)
// 请求包括服务名 Arith，方法名 Multiply，参数 args 三个，服务端的响应包括错误 error，返回值 reply 2 个。

// Header 请求header
type Header struct {
	ServiceMethod string // format "Service.Method" 服务名 + 方法名
	Seq           uint64 // sequence number chosen by client
	Error         string
}
// Codec 抽象出对消息体进行编解码的接口 Codec，抽象出接口是为了实现不同的 Codec 实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}
type NewCodecFunc func(io.ReadWriteCloser) Codec  // 类似别名吧，把 后面那个函数 统一封装了一下
type Type string
const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented 这个方法没有实现
)
var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc) // 这里存储的是构造函数，而非实例
	NewCodecFuncMap[GobType] = NewGobCodec
}
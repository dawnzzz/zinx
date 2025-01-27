package s_router

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

type HelloZinxRouter struct {
	znet.BaseRouter
}

//HelloZinxRouter Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Ins().DebugF("recv from client : msgId=%d, data=%+v, len=%d", request.GetMsgID(), string(request.GetData()), len(request.GetData()))

	err := request.GetConnection().SendBuffMsg(3, []byte("Hello Zinx Router[FromServer]"))
	if err != nil {
		zlog.Error(err)
	}
}

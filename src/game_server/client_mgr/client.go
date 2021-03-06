package clientmgr

import (
	"msg_def"
)

type Client struct {
	linker NetLinker
}

func (this *Client) RpcCall(msg interface{}) {
	b, _, _ := msg_def.MarshalMsg(msg)
	this.linker.WriteMsg(msg_def.LuaMsgID, b)
}

func (this *Client) SendMsg(msg_id ProtoTypeID, msg interface{}) {
	b, _, _ := msg_def.MarshalMsg(msg)
	this.linker.WriteMsg(msg_id, b)
}

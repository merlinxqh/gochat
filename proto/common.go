package proto

type Starter struct {
	Port         int    //http服务启动端口
	RpcHosts     string //注册到etcd 的rpc服务 host
	WsBindHosts  string //websocket服务启动绑定host
	TcpBindHosts string //tcp服务启动绑定host
}

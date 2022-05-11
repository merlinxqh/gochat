/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 10:56
 */
package main

import (
	"flag"
	"fmt"
	"gochat/api"
	"gochat/connect"
	"gochat/logic"
	"gochat/proto"
	"gochat/site"
	"gochat/task"
	"os"
	"os/signal"
	"syscall"
)

var (
	module   = flag.String("module", "", "assign run module")
	port     = flag.Int("p", 0, "start web server port, eg: \n8080")
	rpcHosts = flag.String("rh", "", "register rpc server hosts, eg: \ntcp@127.0.0.1:6900,tcp@127.0.0.1:6901")
	wsHosts  = flag.String("wb", "", "start websocket server bind hosts, eg: \n0.0.0.0:7000")
	tcpHosts = flag.String("tb", "", "start tcp server bind hosts, eg: \n0.0.0.0:7001,0.0.0.0:7002")
)

func main() {
	flag.Parse()
	starter := proto.Starter{
		Port:         *port,
		RpcHosts:     *rpcHosts,
		WsBindHosts:  *wsHosts,
		TcpBindHosts: *tcpHosts,
	}
	fmt.Println(fmt.Sprintf("start run %s module", *module))
	switch *module {
	case "logic":
		logic.New().Run(starter)
	case "connect_websocket":
		connect.New().Run(starter)
	case "connect_tcp":
		connect.New().RunTcp(starter)
	case "task":
		task.New().Run(starter)
	case "api":
		api.New().Run(starter)
	case "site":
		site.New().Run(starter)
	default:
		fmt.Println("exiting,module param error!")
		return
	}
	fmt.Println(fmt.Sprintf("run %s module done!", module))
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	fmt.Println("Server exiting")
}

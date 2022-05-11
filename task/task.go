/**
 * Created by lock
 * Date: 2019-08-09
 * Time: 18:22
 */
package task

import (
	"github.com/sirupsen/logrus"
	"gochat/config"
	"gochat/proto"
	"runtime"
)

type Task struct {
}

func New() *Task {
	return new(Task)
}

func (task *Task) Run(starter proto.Starter) {
	//read config
	taskConfig := config.Conf.Task
	runtime.GOMAXPROCS(taskConfig.TaskBase.CpuNum)
	//read from redis queue
	if err := task.InitQueueRedisClient(); err != nil {
		logrus.Panicf("task init publishRedisClient fail,err:%s", err.Error())
	}
	//rpc call connect layer send msg
	if err := task.InitConnectRpcClient(); err != nil {
		logrus.Panicf("task init InitConnectRpcClient fail,err:%s", err.Error())
	}
	//@todo need watch etcd connect layer serverId kv change then update RpcConnectClientList
	//GoPush
	task.GoPush()
}

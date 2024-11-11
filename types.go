package bngipcgo

import (
	"net"
	"sync"

	"github.com/CustodiaJS/bngsocket"
)

type OnNewProcessFunction func(*BngIpcProcess)
type OnErrorFunction func(*BngIpcProcess, error)
type OnClosedFunction func(error)

type BngIpcProcess struct {
	bngsocket.BngConn
}

type BngIpcServer struct {
	processInstances []*BngIpcProcess
	listener         net.Listener
	onNewProcess     OnNewProcessFunction
	onError          OnErrorFunction
	onClosed         OnClosedFunction
	wg               *sync.WaitGroup
}

type BngIpcClient struct {
	bngsocket.BngConn
}

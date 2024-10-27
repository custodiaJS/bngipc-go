package bngipcgo

import (
	"net"

	"github.com/CustodiaJS/bngsocket"
)

type BngIpcProcess struct {
	bngsocket.BngConn
}

type BngIpcServer struct {
	processInstances []*BngIpcProcess
	listener         net.Listener
}

type BngIpcClient struct {
	bngsocket.BngConn
}

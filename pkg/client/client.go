package client

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lackone/go-ws/global"
	"go.uber.org/zap"
	"time"
)

type Client struct {
	id           int64           //客户端ID，自动分配
	conn         *websocket.Conn //ws连接
	clientManage *ClientManage   //客户端管理
	connectTime  int64           //首次连接时间
	send         chan []byte     //发送消息
	groups       []string        //客户加入的组
}

func NewClient(id int64, conn *websocket.Conn, clientManage *ClientManage) *Client {
	return &Client{
		id:           id,
		conn:         conn,
		clientManage: clientManage,
		connectTime:  time.Now().Unix(),
		send:         make(chan []byte),
	}
}

// 读消息
func (c *Client) ReadLoop() {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Error(fmt.Sprintf("client[%d] ReadLoop panic", c.id), zap.Any("error", err))
		}
	}()

	defer func() {
		c.clientManage.disconnectChan <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(global.WsSetting.ReadLimit)
	c.conn.SetReadDeadline(time.Now().Add(global.WsSetting.ReadDeadline))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(global.WsSetting.ReadDeadline))
		return nil
	})

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				global.Logger.Error(fmt.Sprintf("client[%d] ReadMessage error", c.id), zap.Any("error", err))
			}
			return
		}
		c.ProcessMessage(msg)
	}
}

// 处理消息
func (c *Client) ProcessMessage(msg []byte) {
	req := &ClientRequest{}

	//解析请求
	if err := json.Unmarshal(msg, req); err != nil {
		bytes, _ := NewClientResponse(500, err.Error(), nil).GetByte()
		c.SendMsg(bytes)
		return
	}

	//处理请求数据
	reqData, err := json.Marshal(req.Data)
	if err != nil {
		bytes, _ := NewClientResponse(500, err.Error(), nil).GetByte()
		c.SendMsg(bytes)
		return
	}

	//获取调用方法
	handler, ok := WsClientHandler.GetHandler(req.Url)
	if !ok {
		bytes, _ := NewClientResponse(500, req.Url+" handler not found", nil).GetByte()
		c.SendMsg(bytes)
		return
	}

	//返回数据
	response := handler(c, reqData)
	bytes, _ := response.GetByte()
	c.SendMsg(bytes)
	return
}

// 发送消息
func (c *Client) SendMsg(msg []byte) {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Error(fmt.Sprintf("client[%d] SendMsg panic", c.id), zap.Any("error", err))
		}
	}()

	select {
	case c.send <- msg:
	default:

	}
}

// 写消息
func (c *Client) WriteLoop() {
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Error(fmt.Sprintf("client[%d] WriteLoop panic", c.id), zap.Any("error", err))
		}
	}()

	//定时器，定时发送心跳包
	ticker := time.NewTicker(global.WsSetting.HeartbeatInterval)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(global.WsSetting.WriteDeadline))

			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(global.WsSetting.WriteDeadline))

			if err := c.conn.WriteMessage(websocket.PingMessage, []byte(global.WsSetting.PingMessage)); err != nil {
				return
			}
		}
	}
}

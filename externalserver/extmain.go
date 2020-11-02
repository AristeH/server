// externalserver 
// Copyright 2020 ArisTeh <aristeh.otoko@gmail.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.


package externalserver

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)


var sLogName = "egui.log"

var mfu map[string]func([]string) string
var aRunProc [][]string
var muxRunProc sync.Mutex

var aRunFu []func()
var muxRunFu sync.Mutex

//Sendout отправка сообщения websocket клиенту
func (c *Client) Sendout(s string) bool {
	//jsonMessage, _ := json.Marshal(&Message{Sender: s, Event: "sendout"})
	c.send <- []byte(s)
	return true
}

func (c *Client) sendoutAndReturn(s string) []byte {
	jsonMessage, _ := json.Marshal(&Message{Sender: s, Action: "sendoutAndReturn"})
	c.send <- jsonMessage
	f := <-VСН
	return []byte(f.Message.Sender)
}

// WriteLog writes the sText to a log file egui.log.
func WriteLog(sText string) {

	f, err := os.OpenFile(sLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(sText)

}

// RegFunc adds the fu func to a map of functions,
// sName argument is a function identifier - a key of this map.
// You may need to call this function in case of using HWGui's xml forms.
func RegFunc(sName string, fu func([]string) string) {
	if mfu == nil {
		mfu = make(map[string]func([]string) string)
	}
	mfu[sName] = fu
}

// AddFuncToIdle adds a function to be executed while wait state.
func AddFuncToIdle(fu func()) {
	muxRunFu.Lock()
	aRunFu = append(aRunFu, fu)
	muxRunFu.Unlock()
}

// Runproc выполним процедуру
func Runproc(c *MessageCH) {
	println(c.Message.Action)
	println( mfu[c.Message.Content])
	if fnc, bExist := mfu[c.Message.Content]; bExist {
		var ap []string
		if len(c.Message.Parameters) > 2 {
			ap = make([]string, 5)
			err := json.Unmarshal([]byte(c.Message.Parameters[1]), &ap)
			if err != nil {
				WriteLog(fmt.Sprintf("runproc param Unmarshal error (%s)\r\n", c.Message.Parameters))
			}
		}
		WriteLog(fmt.Sprintf("pgo> (%s) len:%d\r\n", c.Message.Parameters, len(ap)))
      
		c.Client.Sendout(c.Message.Action+";"+fnc(ap))
	}
}


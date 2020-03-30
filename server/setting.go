package server

import (
	"sync"

	"github.com/YWJSonic/ServerUtility/dbservice"
	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/robfig/cron"
)

// Setting ...
type Setting struct {
	mu                 *sync.RWMutex
	AccountEncodeStr   string
	DBIP               string
	DBPassword         string
	DBPORT             string
	DBUser             string
	IP                 string
	PORT               string
	SocketPort         string
	MaintainFinishTime string
	MaintainStartTime  string
	RedisURL           string
	TransferURL        string
	ServerMod          string
}

// RestfultAdderss ...
func (s Setting) RestfultAdderss() string {
	return s.IP + ":" + s.PORT
}

// SocketAdderss ...
func (s Setting) SocketAdderss() string {
	return s.IP + ":" + s.SocketPort
}

// DBSetting ...
func (s Setting) DBSetting() dbservice.ConnSetting {
	return dbservice.ConnSetting{
		DBUser:     s.DBUser,
		DBPassword: s.DBPassword,
		DBIP:       s.DBIP,
		DBPORT:     s.DBPORT,
	}
}

// SetData ...
func (s *Setting) SetData(data map[string]interface{}) {
	if value, ok := data["AccountEncodeStr"]; ok {
		s.AccountEncodeStr = foundation.InterfaceToString(value)
	}
	if value, ok := data["DBIP"]; ok {
		s.DBIP = foundation.InterfaceToString(value)
	}
	if value, ok := data["DBPassword"]; ok {
		s.DBPassword = foundation.InterfaceToString(value)
	}
	if value, ok := data["DBPORT"]; ok {
		s.DBPORT = foundation.InterfaceToString(value)
	}
	if value, ok := data["DBUser"]; ok {
		s.DBUser = foundation.InterfaceToString(value)
	}
	if value, ok := data["IP"]; ok {
		s.IP = foundation.InterfaceToString(value)
	}
	if value, ok := data["MaintainFinishTime"]; ok {
		s.MaintainFinishTime = foundation.InterfaceToString(value)
	}
	if value, ok := data["MaintainStartTime"]; ok {
		s.MaintainStartTime = foundation.InterfaceToString(value)
	}
	if value, ok := data["PORT"]; ok {
		s.PORT = foundation.InterfaceToString(value)
	}
	if value, ok := data["RedisURL"]; ok {
		s.RedisURL = foundation.InterfaceToString(value)
	}
	if value, ok := data["SocketPORT"]; ok {
		s.SocketPort = foundation.InterfaceToString(value)
	}
	if value, ok := data["TransferURL"]; ok {
		s.TransferURL = foundation.InterfaceToString(value)
	}
	if value, ok := data["TransferURL"]; ok {
		s.TransferURL = foundation.InterfaceToString(value)
	}
	if value, ok := data["ServerMod"]; ok {
		s.ServerMod = foundation.InterfaceToString(value)
	}
}

// ToClient ...
func (s *Setting) ToClient() map[string]interface{} {
	return map[string]interface{}{
		"servertime":   s.ServerTime(),
		"maintaintime": s.MaintainTime(),
	}
}

// ServerTime ...
func (s *Setting) ServerTime() int64 {
	return foundation.ServerNowTime()
}

// MaintainTime ...
func (s *Setting) MaintainTime() int64 {
	target, _ := cron.ParseStandard(s.MaintainStartTime)
	return target.Next(foundation.ServerNow()).Unix()
}

package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/YWJSonic/ServerUtility/dbservice"
	"github.com/YWJSonic/ServerUtility/foundation"
	"github.com/YWJSonic/ServerUtility/igame"
	"github.com/YWJSonic/ServerUtility/messagehandle"
	"github.com/YWJSonic/ServerUtility/myhttp"
	"github.com/YWJSonic/ServerUtility/restfult"
	"github.com/YWJSonic/ServerUtility/socket"
	"github.com/YWJSonic/ServerUtility/thirdparty/transaction"
)

// IServer ...
type IServer interface {
	Launch(Setting)
	LaunchRestfult([]restfult.Setting)
	LaunchSocket([]socket.Setting)
	LaunchDB()
	Log(string)
	ErrorLog(string)
}

// Service IServr
type Service struct {
	ShotDown chan bool
	Setting  Setting
	Restfult *restfult.Service
	Socket   *socket.Service
	HTTPConn *myhttp.Service
	Transfer *transaction.Service
	IGame    igame.IGame
	DBs      map[string]*dbservice.DB
}

// Launch server start
func (s *Service) Launch(setting Setting) {
	s.Setting = setting
}

// LaunchRestfult service start
func (s *Service) LaunchRestfult(setting []restfult.Setting) {
	s.Restfult.HTTPLisentRun(s.Setting.RestfultAdderss(), setting)
}

// LaunchSocket service start
func (s *Service) LaunchSocket(setting []socket.Setting) {
	s.Socket.HTTPLisentRun(s.Setting.SocketAdderss(), setting)
}

// LaunchDB ...
func (s *Service) LaunchDB(nameDB string, setting dbservice.ConnSetting) error {
	if _, ok := s.DBs[nameDB]; ok {
		return nil
	}

	db := dbservice.DB{}
	db.SetSetting(setting)
	err := db.ConnectDB(nameDB, setting)
	if err != nil {
		return err
	}

	s.DBs[nameDB] = &db

	return nil

}

// DBConn ...
func (s *Service) DBConn(nameDB string) *sql.DB {
	if conn, ok := s.DBs[nameDB]; ok {
		return conn.GetDB()
	}
	return nil
}

// Log ...
func (s *Service) Log(log string) {
	fmt.Println(log)
}

// ErrorLog ...
func (s *Service) ErrorLog(log string) {
	fmt.Println(log)
}

// HTTPResponse Respond to cliente
func (s *Service) HTTPResponse(httpconn http.ResponseWriter, data interface{}, err messagehandle.ErrorMsg) {
	result := make(map[string]interface{})
	result["data"] = data
	result["error"] = err
	fmt.Fprint(httpconn, foundation.JSONToString(result))
	messagehandle.LogPrintln("HTTPResponse:", foundation.JSONToString(result))
}

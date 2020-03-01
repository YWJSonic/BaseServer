package server

import (
	"sync"

	"github.com/YWJSonic/ServerUtility/dbservice"
)

// NewSetting ..
func NewSetting() Setting {
	return Setting{
		mu: new(sync.RWMutex),
	}
}

// NewService ...
func NewService() *Service {
	return &Service{
		ShotDown: make(chan bool),
		DBs:      make(map[string]*dbservice.DB),
	}
}

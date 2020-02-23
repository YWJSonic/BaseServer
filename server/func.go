package server

import (
	"sync"

	"gitlab.com/ServerUtility/dbservice"
	// "gitlab.com/baseserver/dbservice"
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

package eLiveListCtrl

import (
	"github.com/use-go/websocket-streamserver/wssapi"
)

// EveSetUpStreamApp struct
type EveSetUpStreamApp struct {
	ID       string `json:"ID"`
	Add      bool
	App      string `json:"app"`
	Instance string `json:"instance,omitempty"`
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	Addr     string `json:"addr"`
	Weight   int    `json:"weight"`
}

func (eveSetUpStreamApp *EveSetUpStreamApp) Receiver() string {
	return wssapi.OBJStreamerServer
}

func (eveSetUpStreamApp *EveSetUpStreamApp) Type() string {
	return SetUpStreamApp
}

func NewSetUpStreamApp(add bool, app, instance, protocol, addr, name string, port, weight int) (out *EveSetUpStreamApp) {
	out = &EveSetUpStreamApp{}
	out.Add = add
	out.App = app
	out.Instance = instance
	out.Protocol = protocol
	out.Addr = addr
	out.Port = port
	out.ID = name
	out.Weight = weight
	return
}

func (eveSetUpStreamApp *EveSetUpStreamApp) Copy() (out *EveSetUpStreamApp) {
	out = &EveSetUpStreamApp{}
	out.ID = eveSetUpStreamApp.ID
	out.Add = eveSetUpStreamApp.Add
	out.App = eveSetUpStreamApp.App
	out.Instance = eveSetUpStreamApp.Instance
	out.Protocol = eveSetUpStreamApp.Protocol
	out.Addr = eveSetUpStreamApp.Addr
	out.Port = eveSetUpStreamApp.Port
	out.Weight = eveSetUpStreamApp.Weight
	return
}

func (eveSetUpStreamApp *EveSetUpStreamApp) Equal(rh *EveSetUpStreamApp) bool {
	return eveSetUpStreamApp.ID == rh.ID &&
		eveSetUpStreamApp.App == rh.App &&
		eveSetUpStreamApp.Protocol == rh.Protocol &&
		eveSetUpStreamApp.Addr == rh.Addr &&
		eveSetUpStreamApp.Port == rh.Port &&
		eveSetUpStreamApp.Weight == rh.Weight
}

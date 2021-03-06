package session

import (
	"github.com/mysterium/node/openvpn"
	"github.com/mysterium/node/session"
)

func NewManager(clientConfig *openvpn.ClientConfig) *manager {
	return &manager{
		idGenerator:  &session.Generator{},
		clientConfig: clientConfig,
		sessions:     make([]session.SessionId, 0),
	}
}

type manager struct {
	idGenerator  session.GeneratorInterface
	clientConfig *openvpn.ClientConfig
	sessions     []session.SessionId
}

func (manager *manager) Create() (sessionInstance session.Session, err error) {
	sessionInstance.Id = manager.idGenerator.Generate()

	sessionInstance.Config, err = openvpn.ConfigToString(*manager.clientConfig.Config)
	if err != nil {
		return
	}

	manager.Add(sessionInstance)
	return sessionInstance, nil
}

func (manager *manager) Add(session session.Session) {
	manager.sessions = append(manager.sessions, session.Id)
}

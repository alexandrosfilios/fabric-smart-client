/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package manager

import (
	"encoding/base64"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/comm"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/view"
)

type SelfSession struct {
	id        string
	caller    string
	contextID string
	endpoint  string
	pkid      []byte
	info      view.SessionInfo
	ch        chan *view.Message
}

func NewSelfSession(caller string, contextID string, endpoint string, pkid []byte) (*SelfSession, error) {
	ID, err := comm.GetRandomNonce()
	if err != nil {
		return nil, err
	}

	return &SelfSession{
		id:        base64.StdEncoding.EncodeToString(ID),
		caller:    caller,
		contextID: contextID,
		endpoint:  endpoint,
		pkid:      pkid,
		info: view.SessionInfo{
			ID:           base64.StdEncoding.EncodeToString(ID),
			Caller:       nil,
			CallerViewID: "",
			Endpoint:     endpoint,
			EndpointPKID: pkid,
			Closed:       false,
		},
		ch: make(chan *view.Message, 10),
	}, nil
}

func (s *SelfSession) Info() view.SessionInfo {
	return s.info
}

func (s *SelfSession) Send(payload []byte) error {
	s.ch <- &view.Message{
		SessionID:    s.info.ID,
		ContextID:    s.contextID,
		Caller:       s.caller,
		FromEndpoint: s.endpoint,
		FromPKID:     s.pkid,
		Status:       view.OK,
		Payload:      payload,
	}
	return nil
}

func (s *SelfSession) SendError(payload []byte) error {
	s.ch <- &view.Message{
		SessionID:    s.info.ID,
		ContextID:    s.contextID,
		Caller:       s.caller,
		FromEndpoint: s.endpoint,
		FromPKID:     s.pkid,
		Status:       view.ERROR,
		Payload:      payload,
	}
	return nil
}

func (s *SelfSession) Receive() <-chan *view.Message {
	return s.ch
}

func (s *SelfSession) Close() {
	s.info.Closed = true
}

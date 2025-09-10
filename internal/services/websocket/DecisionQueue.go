package websocketServices

import (
	"context"
	"sync"
	"time"

	sharedModel "my-api/internal/models/shared"
	sharedServices "my-api/internal/services/shared"
	types "my-api/internal/types"

	"github.com/gorilla/websocket"
)

// DecisionQueue manages MakeDecision requests for a single websocket connection.
// Only one request is processed at a time. When a new request arrives, the current
// one is cancelled and only the latest request's response will be sent back.
type DecisionQueue struct {
	mu       sync.Mutex
	activeID string
	cancel   context.CancelFunc
}

// NewDecisionQueue creates a new DecisionQueue instance.
func NewDecisionQueue() *DecisionQueue {
	return &DecisionQueue{}
}

// Submit cancels any running job and starts processing the provided request.
// The response is only sent if the request is still the active one when the job completes.
func (dq *DecisionQueue) Submit(conn *websocket.Conn, req sharedModel.MakeDecisionRequest, sendResp types.SendResponseFunc, sendErr types.SendErrorFunc) {
	dq.mu.Lock()
	if dq.cancel != nil {
		dq.cancel()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	dq.cancel = cancel
	dq.activeID = req.RequestID
	dq.mu.Unlock()

	go func(id string) {
		data, err := sharedServices.MakeDecisionService(ctx, req.Message, req.Checksum, req.Token)

		dq.mu.Lock()
		if id != dq.activeID {
			dq.mu.Unlock()
			return
		}
		dq.cancel = nil
		dq.activeID = ""
		dq.mu.Unlock()

		if err != nil {
			sendErr(conn, "MakeDecision", req.Checksum, map[string]interface{}{"message": err.Error()})
			return
		}

		converted := make(map[string]interface{})
		for k, v := range data {
			converted[k] = v
		}
		sendResp(conn, "MakeDecision", req.Checksum, converted)
	}(req.RequestID)
}

// Close cancels any active request.
func (dq *DecisionQueue) Close() {
	dq.mu.Lock()
	if dq.cancel != nil {
		dq.cancel()
		dq.cancel = nil
	}
	dq.activeID = ""
	dq.mu.Unlock()
}

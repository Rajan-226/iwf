package signal

import (
	"github.com/cadence-oss/iwf-server/gen/iwfidl"
	"github.com/cadence-oss/iwf-server/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	WorkflowType = "signal"
	State1       = "S1"
	State2       = "S2"
	SignalName   = "test-signal-name"
)

func NewSignalWorkflow() (*Handler, *gin.Engine) {
	router := gin.Default()

	handler := newHandler()

	router.POST(service.StateStartApi, handler.apiV1WorkflowStateStart)
	router.POST(service.StateDecideApi, handler.apiV1WorkflowStateDecide)

	return handler, router
}

type Handler struct {
	invokeHistory map[string]int64
	invokeData    map[string]interface{}
}

func newHandler() *Handler {
	return &Handler{
		invokeHistory: make(map[string]int64),
		invokeData:    make(map[string]interface{}),
	}
}

// ApiV1WorkflowStartPost - for a workflow
func (h *Handler) apiV1WorkflowStateStart(c *gin.Context) {
	var req iwfidl.WorkflowStateStartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("received state start request, ", req)

	if req.GetWorkflowType() == WorkflowType {
		h.invokeHistory[req.GetWorkflowStateId()+"_start"]++
		if req.GetWorkflowStateId() == State1 {
			c.JSON(http.StatusOK, iwfidl.WorkflowStateStartResponse{
				CommandRequest: &iwfidl.CommandRequest{
					SignalCommands: []iwfidl.SignalCommand{
						{
							CommandId:  iwfidl.PtrString("signal-cmd-id"),
							SignalName: iwfidl.PtrString(SignalName),
						},
					},
					DeciderTriggerType: iwfidl.PtrString(service.DeciderTypeAllCommandCompleted),
				},
			})
			return
		}
		if req.GetWorkflowStateId() == State2 {
			c.JSON(http.StatusOK, iwfidl.WorkflowStateStartResponse{
				CommandRequest: &iwfidl.CommandRequest{
					DeciderTriggerType: iwfidl.PtrString(service.DeciderTypeAllCommandCompleted),
				},
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, struct{}{})
}

func (h *Handler) apiV1WorkflowStateDecide(c *gin.Context) {
	var req iwfidl.WorkflowStateDecideRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("received state decide request, ", req)

	if req.GetWorkflowType() == WorkflowType {
		h.invokeHistory[req.GetWorkflowStateId()+"_decide"]++
		if req.GetWorkflowStateId() == State1 {
			signalResults := req.GetCommandResults()
			signalId := signalResults.SignalResults[0].GetCommandId()
			signalValue := signalResults.SignalResults[0].GetSignalValue()

			h.invokeData["signalId"] = signalId
			h.invokeData["signalValue"] = signalValue

			c.JSON(http.StatusOK, iwfidl.WorkflowStateDecideResponse{
				StateDecision: &iwfidl.StateDecision{
					NextStates: []iwfidl.StateMovement{
						{
							StateId: iwfidl.PtrString(State2),
						},
					},
				},
			})
			return
		} else if req.GetWorkflowStateId() == State2 {
			// go to complete
			c.JSON(http.StatusOK, iwfidl.WorkflowStateDecideResponse{
				StateDecision: &iwfidl.StateDecision{
					NextStates: []iwfidl.StateMovement{
						{
							StateId: iwfidl.PtrString(service.CompletingWorkflowStateId),
						},
					},
				},
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, struct{}{})
}

func (h *Handler) GetTestResult() (map[string]int64, map[string]interface{}) {
	return h.invokeHistory, h.invokeData
}

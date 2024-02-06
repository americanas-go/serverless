package azure

import (
	"context"
	"encoding/json"
	"github.com/americanas-go/errors"
	"github.com/americanas-go/faas/cloudevents"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Helper assists in creating event handlers.
type Helper struct {
	handler *cloudevents.HandlerWrapper
	options *Options
}

// NewHelper returns a new Helper with options.
func NewHelper(ctx context.Context, options *Options,
	handler *cloudevents.HandlerWrapper) *Helper {

	return &Helper{
		handler: handler,
		options: options,
	}
}

// NewDefaultHelper returns a new Helper with default options.
func NewDefaultHelper(ctx context.Context, handler *cloudevents.HandlerWrapper) *Helper {

	opt, err := DefaultOptions()
	if err != nil {
		log.Fatal(err.Error())
	}

	return NewHelper(ctx, opt, handler)
}

func (h *Helper) Start() {

	listenAddr := h.options.Port
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = val
	}

	http.HandleFunc("/api/"+h.options.Name, h.handle)
	err := http.ListenAndServe(":"+listenAddr, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func (h *Helper) handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := log.FromContext(ctx).WithTypeOf(*h)
	if r.Method != "POST" {
		WriteError(w, errors.MethodNotAllowedf("Method is not supported."))
		return
	}

	in, err := h.parseRequest(r)
	if err != nil {
		WriteError(w, err)
		return
	}

	inouts := make([]*cloudevents.InOut, 0)
	inOut := &cloudevents.InOut{In: in}
	inouts = append(inouts, inOut)

	err = h.handler.Process(ctx, inouts)
	if err != nil {
		logger.Error(errors.ErrorStack(err))
		WriteError(w, err)
		return
	}
	h.processResponse(w, inOut)
}

func (h *Helper) parseRequest(r *http.Request) (*event.Event, error) {
	in := event.New()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &in)
	if err != nil {
		var data interface{}
		if err := json.Unmarshal(body, &data); err != nil {
			return nil, errors.BadRequestf("Bad request.")
		}
		err = in.SetData("application/json", data)
		if err != nil {
			return nil, errors.Internalf("Internal error.")
		}
	}
	invocationID := r.Header.Get("X-Functions-InvocationId")
	in.SetID(invocationID)
	in.SetTime(time.Now())
	return &in, nil
}

func (h *Helper) processResponse(w http.ResponseWriter, inOut *cloudevents.InOut) {
	if inOut.Out != nil {
		w.Header().Set("Content-Type", "application/json")
		responseJson, err := json.Marshal(inOut.Out)
		if err != nil {
			WriteError(w, err)
			return
		}
		_, err = w.Write(responseJson)
		if err != nil {
			WriteError(w, err)
			return
		}
	}
}

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	status := ErrorStatusCode(err)
	resp := response.Error{HttpStatusCode: status, ErrorCode: strconv.Itoa(status), Message: err.Error()}
	responseJson, _ := json.Marshal(resp)
	w.Write(responseJson)
}

// ErrorStatusCode translates to the respective status code.
func ErrorStatusCode(err error) int {

	switch {
	case errors.IsNotFound(err):
		return http.StatusNotFound
	case errors.IsMethodNotAllowed(err):
		return http.StatusMethodNotAllowed
	case errors.IsNotValid(err) || errors.IsBadRequest(err):
		return http.StatusBadRequest
	case errors.IsServiceUnavailable(err):
		return http.StatusServiceUnavailable
	case errors.IsConflict(err) || errors.IsAlreadyExists(err):
		return http.StatusConflict
	case errors.IsNotImplemented(err) || errors.IsNotProvisioned(err):
		return http.StatusNotImplemented
	case errors.IsUnauthorized(err):
		return http.StatusUnauthorized
	case errors.IsForbidden(err):
		return http.StatusForbidden
	case errors.IsNotSupported(err) || errors.IsNotAssigned(err):
		return http.StatusUnprocessableEntity
	default:
		if _, ok := err.(validator.ValidationErrors); ok {
			return http.StatusUnprocessableEntity
		}
		return http.StatusInternalServerError
	}
}

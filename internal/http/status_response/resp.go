package statusresponse

import (
	errorHandler "alert_and_notification/internal/errorHandler"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type OperationType string

const (
	List          OperationType = "list"
	Details       OperationType = "details"
	Create        OperationType = "create"
	Delete        OperationType = "delete"
	Update        OperationType = "update"
	Job           OperationType = "job"
	RedirectedJob OperationType = "redirected_job"
)

func getLogger(c *gin.Context) *zap.Logger {
	v, exists := c.Get("logger")
	if !exists {
		return zap.L()
	}
	logger, ok := v.(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}

type MakeResponseParams struct {
	C          *gin.Context
	ApiMethod  OperationType
	StatusCode int
	Message    string
	Details    interface{}
	Err        error
}

type SuccessResponse struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type ErrorResponse struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// FieldError represents one field-level validation error.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func MakeResponse(makeResponseParams MakeResponseParams) error {
	c := makeResponseParams.C
	log := getLogger(c)
	apiMethod := makeResponseParams.ApiMethod
	statusCode := makeResponseParams.StatusCode
	message := makeResponseParams.Message
	details := makeResponseParams.Details
	err := makeResponseParams.Err

	// 1) If there's an error, handle via your base error types
	if err != nil {
		var he *errorHandler.Error
		if errors.As(err, &he) {
			var resp ErrorResponse
			switch he.C {
			case errorHandler.CategoryDB,
				errorHandler.CategoryENV,
				errorHandler.CategoryBackgroundJob,
				errorHandler.CategoryAPICall:
				log.Warn(he.Error())
				if message == "instance creation failed" {
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": he.Message}},
						},
					}
				} else if message == "rdp generation failed" {
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: makeResponseParams.Details,
					}
				} else {
					statusCode = http.StatusInternalServerError
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": "Something went wrong. Please try again later."}},
						},
					}
				}

			case errorHandler.LibPrivate,
				errorHandler.LibPublic,
				errorHandler.CategoryAPI:
				log.Warn(he.Error())
				if message == "instance creation failed" {
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": he.Message}},
						},
					}
				} else {
					statusCode = http.StatusBadRequest
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": he.Err.Error()}},
						},
					}
				}

			case errorHandler.CategoryQuotaExceeded:
				log.Warn(he.Error())
				if message == "instance creation failed" {
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": he.Message}},
						},
					}
				} else {
					statusCode = http.StatusBadRequest
					resp = ErrorResponse{
						Ok:      false,
						Message: message,
						Details: map[string][]map[string]interface{}{
							"errors": {{"error": he.Err.Error()}},
						},
					}
				}

			default:
				log.Error(he.Error())
				statusCode = http.StatusInternalServerError
				resp = ErrorResponse{
					Ok:      false,
					Message: message,
					Details: map[string][]map[string]interface{}{
						"errors": {{"error": "Internal server error. Please try again later."}},
					},
				}
			}

			c.JSON(statusCode, resp)
			return nil
		}

		// Fallback for non-*errorHandler.Error types
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Ok:      false,
			Message: message,
			Details: map[string][]map[string]interface{}{
				"errors": {{"error": err.Error()}},
			},
		})
		return nil
	}

	// 2) No error: success logic
	switch apiMethod {
	case List:
		log.Info(message)
		if statusCode == http.StatusOK && c.Request.Method == http.MethodGet {
			if _, ok := details.(map[string]interface{}); ok {
				c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
				return nil
			}
			c.JSON(statusCode, SuccessResponse{
				Ok:      true,
				Message: message,
				Details: map[string]interface{}{"data": details},
			})
			return nil
		} else if statusCode != http.StatusOK && c.Request.Method == http.MethodGet {
			c.JSON(statusCode, ErrorResponse{Ok: false, Message: message, Details: details})
			return nil
		}

	case Details:
		log.Info(message)
		if statusCode == http.StatusOK && c.Request.Method == http.MethodGet {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		}

	case Update:
		log.Info(message)
		if statusCode == http.StatusOK && c.Request.Method == http.MethodPut {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		} else if statusCode == http.StatusBadRequest && c.Request.Method == http.MethodPut {
			c.JSON(statusCode, ErrorResponse{Ok: false, Message: message, Details: details})
			return nil
		}

	case Create:
		log.Info(message)
		if statusCode == http.StatusCreated {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		}

	case Job:
		detailsMap := details.(map[string]any)
		for _, msg := range detailsMap["submitted_jobs"].([]map[string]string) {
			task, ok := msg["task_id"]
			if ok {
				log.Info(fmt.Sprintf("Job submitted for volume %s with task ID: %s", msg["volname"], task))
			} else {
				er := errorHandler.Error{
					C:       errorHandler.CategoryAPI,
					Op:      "Job Submission",
					Err:     fmt.Errorf("job submission failed %s", msg["error"]),
					Message: "Job submission failed",
				}
				log.Warn(er.Error())
			}
		}
		if statusCode == http.StatusAccepted {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		}

	case RedirectedJob:
		detailsMap := details.(map[string]any)
		for _, msg := range detailsMap["submitted_jobs"].([]any) {
			msg := msg.(map[string]any)
			task, ok := msg["task_id"].(string)
			if ok {
				log.Info(fmt.Sprintf("Job submitted for volume %s with task ID: %s", msg["volname"], task))
			} else {
				er := errorHandler.Error{
					C:       errorHandler.CategoryAPI,
					Op:      "Job Submission",
					Err:     fmt.Errorf("job submission failed %s", msg["error"]),
					Message: "Job submission failed",
				}
				log.Warn(er.Error())
			}
		}
		if statusCode == http.StatusAccepted {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		}

	case Delete:
		log.Info(message)
		if statusCode == http.StatusNoContent && c.Request.Method == http.MethodDelete {
			c.JSON(statusCode, SuccessResponse{Ok: true, Message: message, Details: details})
			return nil
		} else if statusCode == http.StatusBadRequest && c.Request.Method == http.MethodDelete {
			c.JSON(statusCode, ErrorResponse{Ok: false, Message: message, Details: details})
			return nil
		}
	}

	er := errorHandler.Wrap(errorHandler.CategoryAPI, "", "Unexpected Error came", nil)
	log.Error(er.Error())
	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Ok:      false,
		Message: "Unexpected Error",
		Details: nil,
	})
	return nil
}

// MakeAPIValidationErrorResponse parses validator error strings into structured field errors.
func MakeAPIValidationErrorResponse(c *gin.Context, op, message string, err error) error {
	log := getLogger(c)

	re := regexp.MustCompile(`Key:\s*'[^']*'\s*Error:Field validation for '([^']+)' failed on the '([^']+)' tag`)

	raw := err.Error()
	var errs []FieldError
	for _, line := range strings.Split(raw, `\n`) {
		if m := re.FindStringSubmatch(line); len(m) == 3 {
			field, tag := m[1], m[2]
			msg := "Field validation for '" + field + "' failed on the '" + tag + "' rule"
			errs = append(errs, FieldError{Field: field, Message: msg})
		} else if t := strings.TrimSpace(line); t != "" {
			errs = append(errs, FieldError{Field: "", Message: t})
		}
	}

	log_data := errorHandler.Error{
		C:       errorHandler.CategoryAPI,
		Op:      op,
		Err:     err,
		Message: message,
	}
	log.Warn(log_data.Error())

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Ok:      false,
		Message: message,
		Details: map[string]interface{}{"errors": errs},
	})
	return nil
}

// APIValidationErrorResponse kept for backward compatibility.
func APIValidationErrorResponse(c *gin.Context, op, message string, errs []validator.FieldError) error {
	errorList := make([]map[string]interface{}, len(errs))
	for i, e := range errs {
		errorList[i] = map[string]interface{}{
			"field": e.Field(),
			"error": e.Error(),
		}
	}

	c.JSON(http.StatusBadRequest, ErrorResponse{
		Ok:      false,
		Message: message,
		Details: map[string]interface{}{"errors": errorList},
	})
	return nil
}

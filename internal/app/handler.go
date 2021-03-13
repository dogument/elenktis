package app

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dogument/elenktis"
	"github.com/dogument/elenktis/internal/auditor"
	"github.com/google/uuid"
)

func createAuditLog(cfg *elenktis.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		aud := auditor.New(cfg)
		auditLogReq := createAuditLogRequest{}
		err := json.NewDecoder(r.Body).Decode(&auditLogReq)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(createErrResponse("ERR100002", "request does not comply with the contract"))
			return
		}
		eventTime, err := strconv.ParseInt(auditLogReq.EventTime, 10, 64)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(createErrResponse("ERR100003", "request does not comply with the contract: malformed time"))
			return
		}
		alog := elenktis.AuditLog{
			ApplicationID: auditLogReq.ApplicationID,
			Event:         auditLogReq.Event,
			EventTime:     time.Unix(eventTime, 0),
		}
		err = aud.Create(ctx, alog)
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(createErrResponse("ERR100001", "unable to create the audit log"))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createAuditLogResponse{
			Message:    "successfully created the audit log",
			AuditLogID: uuid.NewString(),
		})
		return

	}
}

type createAuditLogRequest struct {
	ApplicationID string `json:"application_id"`
	Event         string `json:"event"`
	EventTime     string `json:"event_time"`
}

type createAuditLogResponse struct {
	Message    string `json:"message,omitempty"`
	AuditLogID string `json:"audit_log_id,omitempty"`
}

type errorResponse struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func createErrResponse(errCode, errMsg string) errorResponse {
	return errorResponse{
		ErrorCode:    errCode,
		ErrorMessage: errMsg,
	}
}

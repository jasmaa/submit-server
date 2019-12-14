// Package handlers contains router handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jasmaa/submit-server/evaluator/internal/runner"
)

// RunCodeHandler handles code submission requests
func RunCodeHandler(w http.ResponseWriter, r *http.Request) {
	runReq := runner.RunRequest{
		Lang:      r.FormValue("lang"),
		ProjectID: r.FormValue("project_id"),
	}
	resp := runReq.Run()
	json.NewEncoder(w).Encode(&resp)
}

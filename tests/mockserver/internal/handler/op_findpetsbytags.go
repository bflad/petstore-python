// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"mockserver/internal/handler/assert"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"net/http"
)

func opFindPetsByTags() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := assert.SecurityHeader(req, "api_key", false); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := assert.HeaderValues(req, "Accept", []string{"application/json"}); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := assert.HeaderExists(req, "User-Agent"); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		respBody := []components.Pet{
			components.Pet{
				ID:   types.Int64(10),
				Name: "doggie",
				Category: &components.Category{
					ID:   types.Int64(1),
					Name: types.String("Dogs"),
				},
				PhotoUrls: []string{
					"<value>",
					"<value>",
					"<value>",
				},
			},
		}
		respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

		if err != nil {
			http.Error(
				w,
				"Unable to encode response body as JSON: "+err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(respBodyBytes)
	}
}

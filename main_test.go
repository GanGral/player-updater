package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"player-updater/updater"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type AuthParams struct {
	Token    string `json:"x-authentication-token"`
	ClientId string `json:"x-client-id"`
}

func TestUpdateHandler(t *testing.T) {

	expectedBody := bytes.NewReader([]byte(`{"profile":{"applications":[{"applicationId":"music_app","version":"v1.4.10"},{"applicationId":"diagnostic_app","version":"v1.2.6"},{"applicationId":"settings_app","version":"v1.1.5"}]}}
	`))

	emptyBody := bytes.NewReader([]byte(""))

	tt := []struct {
		name       string
		method     string
		input      *AuthParams
		want       string
		statusCode int
		body       *bytes.Reader
	}{
		{
			name:       "without authentication",
			method:     http.MethodPut,
			input:      &AuthParams{},
			want:       "invalid clientId or token supplied",
			statusCode: http.StatusUnauthorized,
			body:       expectedBody,
		},
		{
			name:   "all good",
			method: http.MethodPut,
			input: &AuthParams{
				ClientId: "dkd",
				Token:    "skjd",
			},
			want:       `{"profile":{"applications":[{"applicationId":"music_app","version":"v1.4.10"},{"applicationId":"diagnostic_app","version":"v1.2.6"},{"applicationId":"settings_app","version":"v1.1.5"}]}} `,
			statusCode: http.StatusOK,
			body:       expectedBody,
		},
		{
			name:   "with empty body",
			method: http.MethodPut,
			input: &AuthParams{
				ClientId: "ds",
				Token:    "dsdd",
			},
			want:       `child \"profile\" fails because [child \"applications\" fails because [\"applications\" is required]]`,
			statusCode: http.StatusConflict,
			body:       emptyBody,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/profiles/clientId:a1:bb:cc:dd:ee:ff", tc.body)
			responseRecorder := httptest.NewRecorder()
			request.Header.Set("Content-Type", "application/json")
			request.Header.Set("X-Client-Id", tc.input.ClientId)
			request.Header.Set("X-Authentication-Token", tc.input.Token)
			//context.Set(r, 0, 1)

			Router().ServeHTTP(responseRecorder, request)
			//updater.HandleUpdate(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}

}

// setting up Router so we can properly handle the macaddress field through mux
func Router() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/profiles/clientId:{macaddress}", updater.HandleUpdate)
	return route
}

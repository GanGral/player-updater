package main

//Test for updater service UpdateHandler.

// 5 tests are performed:
// PUT request with no token/client id provided
// PUT request with expired token
// Successfull PUT request
// Empty body PUT request

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"player-updater/updater"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

//Structure to construct for authentication
type AuthParams struct {
	ClientToken Token  `json:"token"`
	ClientId    string `json:"x-client-id"`
}
type Token struct {
	Token   string `json:"x-authentication-token"`
	Expired bool   `json:"expired"`
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
			name:       "without token and clientId",
			method:     http.MethodPut,
			input:      &AuthParams{},
			want:       "invalid clientId or token supplied",
			statusCode: http.StatusUnauthorized,
			body:       expectedBody,
		},
		{
			name:   "expired token",
			method: http.MethodPut,
			input: &AuthParams{
				ClientId: "dkd",
				ClientToken: Token{
					Token:   "skjd",
					Expired: true,
				},
			},
			want:       "invalid clientId or token supplied",
			statusCode: http.StatusUnauthorized,
			body:       expectedBody,
		},
		{
			name:   "all good",
			method: http.MethodPut,
			input: &AuthParams{
				ClientId: "dkd",
				ClientToken: Token{
					Token:   "skjd",
					Expired: false,
				},
			},
			want:       `{"profile":{"applications":[{"applicationId":"music_app","version":"v1.4.10"},{"applicationId":"diagnostic_app","version":"v1.2.6"},{"applicationId":"settings_app","version":"v1.1.5"}]}}`,
			statusCode: http.StatusOK,
			body:       expectedBody,
		},
		{
			name:   "with empty body",
			method: http.MethodPut,
			input: &AuthParams{
				ClientId: "ds",
				ClientToken: Token{
					Token:   "skjd",
					Expired: false,
				},
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

			if !tc.input.ClientToken.Expired {
				request.Header.Set("X-Authentication-Token", tc.input.ClientToken.Token)
			}

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

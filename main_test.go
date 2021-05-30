package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleUpdate(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.

	myJsonString := `{
		"profile": {    
		  "applications": [
			{
			  "applicationId": "music_app"
			  "version": "v1.4.10"
			},
			{
			  "applicationId": "diagnostic_app",
			  "version": "v1.2.6"
			},
			{
			  "applicationId": "settings_app",
			  "version": "v1.1.5"
			}
		  ]
		}
	  }`

	body := bytes.NewReader([]byte(myJsonString))

	req, err := http.NewRequest("PUT", "http://localhost:8457/profiles/clientId:a2:bb:cc:dd:ee:ff", body)
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Client-Id", "")
	req.Header.Set("X-Authentication-Token", "dummy")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	//handler := http.HandlerFunc(updater.HandleUpdate)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	//handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{
		"profile": {    
		  "applications": [
			{
			  "applicationId": "music_app"
			  "version": "v1.4.10"
			},
			{
			  "applicationId": "diagnostic_app",
			  "version": "v1.2.6"
			},
			{
			  "applicationId": "settings_app",
			  "version": "v1.1.5"
			}
		  ]
		}
	  }`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

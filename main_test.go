package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  // Sending new request first argument is method, 2nd - route, 3rd - body
  req, err := http.NewRequest("GET", "", nil)
  if err != nil {
    t.Fatal(err)
  }
  // Recorder is like mini-browser that records all our requests
  recorder := httptest.NewRecorder()

  // Defining the handler func we want to test
  hf := http.HandlerFunc(handler)

  //Serve the HTTP request to out recorder
  hf.ServeHTTP(recorder, req)

  //Check the status code is what we expect
  if status := recorder.Code; status != http.StatusOK {
    t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
  }

  //check the response body is what we expect.
  expected := `Hello World!`
  actual := recorder.Body.String()
  if actual != expected {
    t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
  }
}
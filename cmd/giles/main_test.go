package main

// func TestHandler(t *testing.T) {
// 	mockdb := &mockdb.MockDB{}
// 	server := &api.Server{DB: mockdb}
// 	srv := httptest.NewServer(server.Router())
// 	defer srv.Close()

// 	mockdb.On()
// }

// func assertStatus(t *testing.T, got, want int) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("did not get correct status, got %d, want %d", got, want)
// 	}
// }

// func newGetTickerRequest(name string) *http.Request {
// 	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/ticker/%s", name), nil)
// 	return req
// }

// func assertResponseBody(t *testing.T, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("response body is wrong, got %q want %q", got, want)
// 	}
// }

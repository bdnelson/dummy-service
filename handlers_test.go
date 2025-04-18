package dummyservice

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestTimeoutWithDuration tests the timeout handler with various duration inputs
func TestTimeoutWithDuration(t *testing.T) {
	type timeoutTestCase struct {
		name           string
		path           string
		pathValue      string
		expectedStatus int
		expectedBody   string
		expectedDelay  time.Duration
	}

	const (
		defaultDuration = 30 * time.Second
		timingTolerance = 500 * time.Millisecond
	)

	tests := []timeoutTestCase{
		{
			name:           "valid duration",
			path:           "/timeout/5",
			pathValue:      "5",
			expectedStatus: http.StatusOK,
			expectedBody:   "done after 5 seconds",
			expectedDelay:  5 * time.Second,
		},
		{
			name:           "default duration when invalid",
			path:           "/timeout/invalid",
			pathValue:      "invalid",
			expectedStatus: http.StatusOK,
			expectedBody:   "done after 30 seconds",
			expectedDelay:  defaultDuration,
		},
		{
			name:           "default duration when missing",
			path:           "/timeout/",
			pathValue:      "",
			expectedStatus: http.StatusOK,
			expectedBody:   "done after 30 seconds",
			expectedDelay:  defaultDuration,
		},
		{
			name:           "zero duration",
			path:           "/timeout/0",
			pathValue:      "0",
			expectedStatus: http.StatusOK,
			expectedBody:   "done after 0 seconds",
			expectedDelay:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			req.SetPathValue("duration", tt.pathValue)
			rr := httptest.NewRecorder()

			TimeoutWithDuration(rr, req)

			validateResponse(t, rr, tt.expectedStatus, tt.expectedBody)
			validateTiming(t, start, tt.expectedDelay, timingTolerance)
		})
	}
}

func validateResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int, expectedBody string) {
	t.Helper()
	if rr.Code != expectedStatus {
		t.Errorf("expected status %d, got %d", expectedStatus, rr.Code)
	}
	if strings.TrimSpace(rr.Body.String()) != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}

func validateTiming(t *testing.T, start time.Time, expectedDelay, tolerance time.Duration) {
	t.Helper()
	elapsed := time.Since(start)
	if elapsed < expectedDelay || elapsed > expectedDelay+tolerance {
		t.Errorf("expected delay around %v, got %v", expectedDelay, elapsed)
	}
}

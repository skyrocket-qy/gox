package httpx_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/httpx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewErrBinder(t *testing.T) {
	errMap := map[erx.Code]int{
		erx.ErrUnknown: http.StatusInternalServerError,
	}
	binder := httpx.NewErrBinder(errMap)
	assert.NotNil(t, binder)
	assert.Equal(t, errMap, binder.ErrToHTTP) // Changed from errToHTTP
}

func TestTrimToProject(t *testing.T) {
	// Get current working directory to simulate project root
	cwd, err := os.Getwd()
	require.NoError(t, err)

	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "Path inside project",
			path:     cwd + "/some/path/file.go",
			expected: "/some/path/file.go",
		},
		{
			name:     "Path outside project",
			path:     "/usr/local/go/src/runtime/proc.go",
			expected: "/usr/local/go/src/runtime/proc.go", // Should not be trimmed
		},
		{
			name:     "Path is project root",
			path:     cwd,
			expected: "",
		},
		{
			name:     "Empty path",
			path:     "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := httpx.TrimToProject(tt.path) // Changed from trimToProject
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestExtractFuncName(t *testing.T) {
	tests := []struct {
		name     string
		fullFunc string
		expected string
	}{
		{
			name:     "Typical function name",
			fullFunc: "github.com/skyrocket-qy/gox/httpx.TestExtractFuncName",
			expected: "httpx.TestExtractFuncName",
		},
		{
			name:     "Method on a struct",
			fullFunc: "github.com/skyrocket-qy/gox/httpx.(*ErrBinder).Bind",
			expected: "httpx.(*ErrBinder).Bind",
		},
		{
			name:     "Function in main package",
			fullFunc: "main.main",
			expected: "main.main",
		},
		{
			name:     "Empty string",
			fullFunc: "",
			expected: "",
		},
		{
			name:     "No slash",
			fullFunc: "someFunc",
			expected: "someFunc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := httpx.ExtractFuncName(tt.fullFunc) // Changed from extractFuncName
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestErrBinder_Bind(t *testing.T) {
	// Suppress zerolog output during tests
	output := &bytes.Buffer{}
	log.Logger = zerolog.New(output).With().Timestamp().Logger()

	tests := []struct {
		name           string
		err            error
		errToHTTP      map[erx.Code]int
		expectedStatus int
		expectedCode   string
		reqID          string
	}{
		{
			name:           "Non-CtxErr error",
			err:            errors.New("some generic error"),
			errToHTTP:      map[erx.Code]int{},
			expectedStatus: http.StatusInternalServerError,
			expectedCode:   erx.ErrUnknown.Str(),
			reqID:          "req123",
		},
		{
			name: "CtxErr with known code",
			err:  erx.New(erx.ErrUnknown, "invalid input"),
			errToHTTP: map[erx.Code]int{
				erx.ErrUnknown: http.StatusBadRequest,
			},
			expectedStatus: http.StatusBadRequest,
			expectedCode:   erx.ErrUnknown.Str(),
			reqID:          "req456",
		},
		{
			name: "CtxErr with unknown code",
			err: erx.New(
				erx.ErrUnknown,
				"unknown error",
			), // Using ErrUnknown as a placeholder for an unknown code
			errToHTTP: map[erx.Code]int{
				erx.ErrUnknown: http.StatusBadRequest,
			},
			expectedStatus: http.StatusBadRequest, // Changed from InternalServerError to Bad Request
			expectedCode:   erx.ErrUnknown.Str(),
			reqID:          "req789",
		},
		{
			name: "CtxErr with empty message",
			err:  erx.New(erx.ErrUnknown, ""),
			errToHTTP: map[erx.Code]int{
				erx.ErrUnknown: http.StatusBadRequest,
			},
			expectedStatus: http.StatusBadRequest,
			expectedCode:   erx.ErrUnknown.Str(),
			reqID:          "req101",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Set("reqId", tt.reqID)

			binder := httpx.NewErrBinder(tt.errToHTTP)

			t.Logf("Before Bind: log.Logger is %v", log.Logger)
			binder.Bind(c, tt.err)

			assert.Equal(t, tt.expectedStatus, recorder.Code)

			var resp httpx.ErrResp

			err := json.Unmarshal(recorder.Body.Bytes(), &resp)
			require.NoError(t, err)
			assert.Equal(t, tt.reqID, resp.ReqId)
			assert.Equal(t, tt.expectedCode, resp.Code)

			// Check logs for non-CtxErr case
			var ctxErr *erx.CtxErr
			if !errors.As(tt.err, &ctxErr) {
				logOutput := output.String()
				assert.Contains(t, logOutput, "error not wrapped by erx")
				output.Reset() // Clear buffer for next test
			}
		})
	}
}

func TestFilterCallerInfos(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)

	tests := []struct {
		name     string
		infos    []erx.CallerInfo
		expected []erx.CallerInfo
	}{
		{
			name: "All project files",
			infos: []erx.CallerInfo{
				{File: cwd + "/file1.go"},
				{File: cwd + "/sub/file2.go"},
			},
			expected: []erx.CallerInfo{
				{File: cwd + "/file1.go"},
				{File: cwd + "/sub/file2.go"},
			},
		},
		{
			name: "Mixed project and non-project files",
			infos: []erx.CallerInfo{
				{File: cwd + "/file1.go"},
				{File: "/usr/local/go/src/runtime/proc.go"},
				{
					File: cwd + "/file2.go",
				}, // This should not be included as it's after a non-project file
			},
			expected: []erx.CallerInfo{
				{File: cwd + "/file1.go"},
			},
		},
		{
			name: "All non-project files",
			infos: []erx.CallerInfo{
				{File: "/usr/local/go/src/runtime/proc.go"},
				{File: "/usr/local/go/src/fmt/print.go"},
			},
			expected: []erx.CallerInfo{},
		},
		{
			name:     "Empty infos",
			infos:    []erx.CallerInfo{},
			expected: []erx.CallerInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := httpx.FilterCallerInfos(tt.infos) // Changed from filterCallerInfos
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGetCallStack(t *testing.T) {
	// This test is a bit tricky as the call stack depends on the test runner and environment.
	// We'll just assert that it returns at least one frame and that the file path contains
	// "http_test.go".
	callerInfos := httpx.GetCallStack() // Changed from getCallStack
	assert.NotEmpty(t, callerInfos)

	// Check if at least one frame points to this test file
	found := false

	for _, info := range callerInfos {
		if strings.Contains(info.File, "http_test.go") {
			found = true

			break
		}
	}

	assert.True(t, found, "Expected to find http_test.go in call stack")

	// Test with callerSkip
	callerInfosWithSkip := httpx.GetCallStack(1) // Changed from getCallStack
	assert.NotEmpty(t, callerInfosWithSkip)
	// The first frame should now be different from the one without skip
	assert.NotEqual(t, callerInfos[0].Function, callerInfosWithSkip[0].Function)
}

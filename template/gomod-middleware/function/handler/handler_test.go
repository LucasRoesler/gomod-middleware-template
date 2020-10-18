package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Handler(t *testing.T) {
	cases := []struct {
		name     string
		request  interface{}
		status   int
		response string
	}{
		{
			name:     "empty body is valid request",
			request:  "",
			status:   http.StatusOK,
			response: `Hello world, input was: ""`,
		},
		{
			name:     "non-empty body is echoed correctly",
			request:  "Captain Cool",
			status:   http.StatusOK,
			response: `Hello world, input was: "Captain Cool"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			bs, _ := json.Marshal(tc.request)
			r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bs))
			w := httptest.NewRecorder()
			Handle(w, r)

			require.Equal(t, tc.status, w.Code)
			require.Equal(t, tc.response, w.Body.String())
		})
	}
}

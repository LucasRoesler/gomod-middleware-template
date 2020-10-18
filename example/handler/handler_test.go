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
			name:     "empty body cause 500 error",
			request:  "",
			status:   http.StatusInternalServerError,
			response: "Unable to list repos for openfaas\n",
		},
	}

	for _, tc := range cases {
		bs, _ := json.Marshal(tc.request)
		r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(bs))
		w := httptest.NewRecorder()
		Handle(w, r)

		require.Equal(t, tc.status, w.Code)
		require.Equal(t, tc.response, w.Body.String())

	}
}

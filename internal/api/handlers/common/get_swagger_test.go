package common_test

import (
	"net/http"
	"testing"

	"backend-test/internal/api"
	"backend-test/internal/test"
	"github.com/stretchr/testify/require"
)

func TestSwaggerYAMLRetrieval(t *testing.T) {
	t.Parallel()

	test.WithTestServer(t, func(s *api.Server) {
		res := test.PerformRequest(t, s, "GET", "/swagger.yml", nil, nil)
		require.Equal(t, http.StatusOK, res.Result().StatusCode)
	})
}

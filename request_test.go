package finding

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func TestAdvancedRequest_GetBody(t *testing.T) {
	service := NewService("")
	type tcase struct {
		filename string
		request  *AdvancedRequest
	}

	tests := make([]tcase, 0)

	request1 := service.NewAdvancedRequest()
	request1.WithKeywords("tolkien")
	request1.WithPageLimit(2)
	tests = append(tests, tcase{
		filename: "Basic.json",
		request:  request1,
	})

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			got, err := tt.request.GetBody()
			if !assert.NoError(t, err) {
				return
			}

			f, err := os.Open(path.Join("testdata", "request", "advanced", tt.filename))
			if !assert.NoError(t, err) {
				return
			}
			buf := bytes.Buffer{}
			_, err = buf.ReadFrom(f)
			if !assert.NoError(t, err) {
				return
			}

			assert.JSONEq(t, string(got), buf.String())
		})
	}
}

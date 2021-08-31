package repo

import (
	"log"
	"net/http"
	"testing"

	"github.com/ivis-yoshida/gogs/internal/conf"
	"github.com/ivis-yoshida/gogs/internal/context"
	"github.com/ivis-yoshida/gogs/internal/db"
	"github.com/ivis-yoshida/gogs/internal/httplib"
	"github.com/stretchr/testify/assert"
	"gopkg.in/macaron.v1"
)

type contextMock struct {
	*context.Context
}

func TestCreateDmp(t *testing.T) {
	conf.SetMockServer(t, conf.ServerOpts{
		ExternalURL: "http://gogs.example.com/",
	})

	m := macaron.New()
	m.Use(macaron.Renderer())
	m.Use(func(c *macaron.Context) {
		c.Map(&db.User{Name: "owner"})
		c.Map(&db.Repository{Name: "repo"})
	})
	m.Get("/", CreateDmp)

	tests := []struct {
		name          string
		ctx           context.Context
		expStatusCode int
	}{
		{
			name: "sample1",
			ctx: context.Context{
				User: &db.User{Name: "owner"},
				Repo: &context.Repository{
					Repository: &db.Repository{Name: "repo"}},
			},
			expStatusCode: http.StatusOK,
		},
		{
			name: "sample2",
			// ctx:           &context.Context{},
			expStatusCode: http.StatusOK,
		},
	}
	// url := "http://gogs.example.com/?schema=METI"
	// http request with "schema" parameter
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// req, err := http.NewRequestWithContext(test.ctx, http.MethodGet, "/", nil)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			// rr := httptest.NewRecorder()
			// m.ServeHTTP(rr, req)
			// res := rr.Result()
			req := httplib.Get("http://gogs.example.com/?schema=METI")
			res, err := req.Response()
			if err != nil {
				log.Fatal("No Response: " + test.name)
			}

			// CreateDmp(test.ctx)
			assert.Equal(t, test.expStatusCode, res.StatusCode)
		})
	}

	// assert dmp schema
}

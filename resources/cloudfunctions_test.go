package resources_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-gcp/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

func createCloudFunctionsTestServer() (*cloudfunctions.Service, error) {
	ctx := context.Background()
	var function cloudfunctions.CloudFunction
	if err := faker.FakeData(&function); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/*data", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudfunctions.ListFunctionsResponse{
			Functions: []*cloudfunctions.CloudFunction{&function},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := cloudfunctions.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func TestCloudfunctionsFunction(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.CloudfunctionsFunction(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			cfSvc, err := createCloudFunctionsTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				CloudFunctions: cfSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

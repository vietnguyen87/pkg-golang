## Getting started

A http interface to communicate with other Restful services. 
```azure
type Client interface {
	PostJSON(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	PostForm(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	Get(c context.Context, url string, target interface{}, reqOptions ...RequestOption) (int, error)
	GetWithQuery(c context.Context, url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	GetWithoutEncodedQuery(c context.Context,
		url string, data, target interface{}, reqOptions ...RequestOption) (int, error)
	Do(ctx context.Context, request *http.Request, target interface{}) (int, error)
}
```
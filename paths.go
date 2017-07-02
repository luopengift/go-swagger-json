package swagger

type Paths map[string]*Methods

func InitPaths() *Paths {
	paths := make(Paths)
    methods := InitMethods()
    methods.AddMethod("GET",InitRequest())
	paths.Add("/test",methods)
    return &paths

}

func (p *Paths) Add(uri string, methods *Methods) {
    (*p)[uri] = methods
}

type Methods struct {
	GET     *Request `json:"get,omitempty"`
	POST    *Request `json:"post,omitempty"`
	PUT     *Request `json:"put,omitempty"`
	DELETE  *Request `json:"delete,omitempty"`
	HEAD    *Request `json:"head,omitempty"`
	OPTIONS *Request `json:"options,omitempty"`
}

func InitMethods() *Methods {
	methods := new(Methods)
	return methods
}

func (m *Methods) AddMethod(method string, req *Request) {
    switch method {
        case "GET":
            m.GET = req
        case "POST":
            m.POST = req
    }
}

type Request struct {
	Tags        []string     `json:"tags"`
	Summary     string       `json:"summary"`
	Description string       `json:"description"`
	OperationId string       `json:"operationId"`
	Produces    []string     `json:"produces,omitempty"`
	Consumes    []string     `json:"consumes,omitempty"`
	Parameters  []*Parameter `json:"parameters"`
	Responses   *Responses   `json:"responses"`
	Security    []*Security  `json:"security"`
	Deprecated  bool         `json:"deprecated,omitempty"`
}

func InitRequest() *Request {
	request := new(Request)
	request.Tags = []string{}
	request.Summary = ""
	request.Description = "Test Request"
	request.OperationId = ""
	request.Produces = []string{}
	request.Consumes = []string{}
	request.Parameters = []*Parameter{InitParameter()}
	request.Responses = InitResponses()
	request.Security = []*Security{InitSecurity()}
	request.Deprecated = false
	return request
}

type Parameter struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Type        string `json:"type,omitempty"`
	Format      string `json:"format,omitempty"`
	Items       struct {
		Type    string   `json:"type,omitempty"`
		Enum    []string `json:"enum,omitempty"`
		Deafult string   `json:"default,omitempty"`
	} `json:"items,omitempty"`
	CollectionFormat string `json:"collectionFormat,omitempty"`
}

func InitParameter() *Parameter {
	parameter := new(Parameter)
	return parameter
}

type Responses struct {
	StatusOK               *Response `json:"200,omitempty"`
	StatusBadRequest       *Response `json:"400,omitempty"`
	StatusUnauthorized     *Response `json:"401,omitempty"`
	StatusPaymentRequired  *Response `json:"402,omitempty"`
	StatusForbidden        *Response `json:"403,omitempty"`
	StatusNotFound         *Response `json:"404,omitempty"`
	StatusMethodNotAllowed *Response `json:"405,omitempty"`
	Default *Response `json:"default,omitempty"`
}

func InitResponses() *Responses {
	Responses := new(Responses)
    Responses.StatusOK = NewResponse("200:successful")
    Responses.StatusNotFound = NewResponse("404:Not Found")
	return Responses
}

type Response struct {
	Description string `json:"description"`
	Schema      struct {
		Type  string `json:"type,omitempty"`
		Items struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"items,omitempty"`
	} `json:"schema,omitempty"`
    Headers interface{} `json:"headers,omitempty"`
}

func InitResponse() *Response {
    response := new(Response)
    return response
}

func NewResponse(desc string) *Response {
    resp := InitResponse()
    resp.Description = desc
    return resp
}

type Security struct {
	PetstoreAuth []string `json:"petstore_auth,omitempty"`
}

func InitSecurity() *Security {
	security := new(Security)
	return security
}

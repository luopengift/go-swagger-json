package swagger

type URI string
type Paths map[URI]Methods

func InitPaths() *Paths {
	paths := make(Paths)
	return &paths

}

func (p *Paths) Add() {

}

type Methods struct {
	GET     Request `json:"get,omitempty"`
	POST    Request `json:"post,omitempty"`
	PUT     Request `json:"put,omitempty"`
	DELETE  Request `json:"delete,omitempty"`
	HEAD    Request `json:"head,omitempty"`
	OPTIONS Request `json:"options,omitempty"`
}

func InitMethods() *Methods {
	methods := new(Methods)
	return methods
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
	request.Description = ""
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
	StatusOK               response `json:"200,omitempty"`
	StatusBadRequest       response `json:"400,omitempty"`
	StatusUnauthorized     response `json:"401,omitempty"`
	StatusPaymentRequired  response `json:"402,omitempty"`
	StatusForbidden        response `json:"403,omitempty"`
	StatusNotFound         response `json:"404,omitempty"`
	StatusMethodNotAllowed response `json:"405,omitempty"`
}

func InitResponses() *Responses {
	responses := new(Responses)
	return responses
}

type response struct {
	Description string `json:"description"`
	Schema      struct {
		Type  string `json:"type,omitempty"`
		Items struct {
			Ref string `json:"$ref,omitempty"`
		} `json:"items,omitempty"`
	} `json:"schema,omitempty"`
}

type Security struct {
	PetstoreAuth []string `json:"petstore_auth,omitempty"`
}

func InitSecurity() *Security {
	security := new(Security)
	return security
}

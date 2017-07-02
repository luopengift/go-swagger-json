package swagger

type ExternalDocs struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

func InitExternalDocs() *ExternalDocs {
	externalDocs := new(ExternalDocs)
	externalDocs.Description = "Find out more about Swagger"
	externalDocs.URL = "http://swagger.io"
	return externalDocs
}

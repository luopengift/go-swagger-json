package swagger

type Swagger struct {
	Swagger             string               `json:"swagger"`
	Info                *Info                `json:"info"`
	Tags                []*Tag               `json:"tags"`
	Paths               *Paths               `json:"paths"`
	SecurityDefinitions *SecurityDefinitions `json:"securityDefinitions"`
	Definitions         *Definitions         `json:"definitions"`
	ExternalDocs        *ExternalDocs        `json:"externalDocs"`
	Host                string               `json:"host"`
	BasePath            string               `json:"basePath"`
	Schemes             []string             `json:"schemes"`
}

func InitSwagger() *Swagger {
	swagger := new(Swagger)
	swagger.Swagger = "2.0"
	swagger.Info = InitInfo()
	swagger.Tags = []*Tag{InitTag()}
	swagger.Paths = InitPaths()
	swagger.SecurityDefinitions = InitSecurityDefinitions()
	swagger.Definitions = InitDefinitions()
	swagger.ExternalDocs = InitExternalDocs()
	swagger.Host = "virtserver.swaggerhub.com"
	swagger.BasePath = "/luopengift/test/1.0.0"
	swagger.Schemes = []string{"http", "https"}
	return swagger
}

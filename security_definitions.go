package swagger

type SecurityDefinitions struct {
	PetstoreAuth struct {
		Type             string            `json:"type"`
		AuthorizationURL string            `json:"authorizationUrl,omitempty"`
		Flow             string            `json:"flow,omitempty"`
		Scopes           map[string]string `json:"scopes,omitempty"`
	} `json:"petstore_auth,omitempty"`
	ApiKey struct {
		Type string `json:"type,omitempty"`
		Name string `json:"name"`
		In   string `json:"in"`
	} `json:"api_key,omitempty"`
}

func InitSecurityDefinitions() *SecurityDefinitions {
	securityDefinitions := new(SecurityDefinitions)
	securityDefinitions.PetstoreAuth.Type = "oauth2"
	securityDefinitions.PetstoreAuth.AuthorizationURL = "http://petstore.swagger.io/oauth/dialog"
	securityDefinitions.PetstoreAuth.Flow = "implicit"
	securityDefinitions.PetstoreAuth.Scopes = map[string]string{
		"write:pets": "modify pets in your account",
		"read:pets":  "read your pets",
	}
	securityDefinitions.ApiKey.Type = "apiKey"
	securityDefinitions.ApiKey.Name = "api_key"
	securityDefinitions.ApiKey.In = "header"
	return securityDefinitions

}

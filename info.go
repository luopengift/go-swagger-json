package swagger

type Info struct {
	Description    string `json:"description"`
	Version        string `json:"version"`
	Title          string `json:"title"`
	TermsOfService string `json:"termsOfService"`
	Contact        struct {
		Email string `json:"email"`
	} `json:"contact"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
}

func InitInfo() *Info {
	info := new(Info)
	info.Description = "This is a sample server Petstore server.  You can find \nout more about Swagger at \n[http://swagger.io](http://swagger.io) or on \n[irc.freenode.net, #swagger](http://swagger.io/irc/).\n"
	info.Version = "1.0.0"
	info.Title = "Swagger Petstore"
	info.TermsOfService = "http://swagger.io/terms/"
	info.Contact.Email = "apiteam@swagger.io"
	info.License.Name = "Apache 2.0"
	info.License.URL = "http://www.apache.org/licenses/LICENSE-2.0.html"
	return info
}

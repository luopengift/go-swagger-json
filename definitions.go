package swagger

type Definitions struct {
	Type       string `json:"type"`
	Properties map[string]*Propertie
	Xml        struct {
		Name string `json:"name"`
	} `json:"xml"`
}

type Propertie struct {
	Type        string `json:"type"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
}

func InitDefinitions() *Definitions {
	definitions := new(Definitions)
	definitions.Type = "object"
	definitions.Properties = map[string]*Propertie{
		"User": &Propertie{
			Type:   "integer",
			Format: "int64",
		},
	}
	definitions.Xml.Name = "Order"
	return definitions

}

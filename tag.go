package swagger

type Tag struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
}

func InitTag() *Tag {
	tag := new(Tag)
	tag.Name = "pet"
	tag.Description = "Everything about your Pets"
	tag.ExternalDocs = InitExternalDocs()
	return tag
}

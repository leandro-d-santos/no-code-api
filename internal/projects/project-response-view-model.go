package projects

type ProjectResponseViewModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (projectResponse *ProjectResponseViewModel) FromModel(project *Project) {
	projectResponse.Id = project.Id
	projectResponse.Name = project.Name
}

package v1

// Project represents the project
//
// swagger:model v1-Project
type Project struct {
	// Metadata Standard object's metadata.
	//
	// readonly: true
	Metadata Metadata `json:"metadata"`

	Spec ProjectSpec `json:"spec"`
}

// ProjectSpec represents the project specifications
//
// swagger:model v1-ProjectSpec
type ProjectSpec struct {
	// Name is project name.
	Name string `json:"name" binding:"required"`

	// Description is project description.
	Description string `json:"description" binding:"required"`
}

// ProjectList represents the list of project
//
// swagger:model v1-ProjectList
type ProjectList struct {
	// Items is the list of Projects.
	Items []Project `json:"items"`
}

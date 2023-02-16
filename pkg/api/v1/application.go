package v1

// Application represents the application
type Application struct {
	Metadata Metadata        `json:"metadata"`
	Spec     ApplicationSpec `json:"spec"`
}

// ApplicationSpec represents the application specifications
type ApplicationSpec struct {
	// Name is application name.
	Name string `json:"name" binding:"required"`

	// RepositoryURL is application's GIT repository URL
	RepositoryURL string `json:"repository_url" binding:"required"`
}

// ApplicationList represents the list of applications
type ApplicationList struct {
	// Items is the list of Applications.
	Items []Application `json:"items"`
}

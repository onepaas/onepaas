package v1

// GithubHookHeaderSpec represents GitHub's webhook header specifications
type GithubHookHeaderSpec struct {
	// Event represents the webhook event
	Event string `header:"X-GitHub-Event" validate:"required,oneof=create"`
}

// GithubHookSpec represents GitHub's webhook specifications
type GithubHookSpec struct {
	ReferenceType string `json:"ref_type" validate:"required,eq=tag"`

	// Repository is the repository on which the event was triggered
	Repository struct {
		IsPrivate *bool  `json:"private" validate:"required,eq=false"`
		FullName  string `json:"full_name" validate:"required"`
	} `json:"repository" validate:"required"`
}

package metadata

type UserMatadata struct {
	UserName    string          `json:"gin_account_name"`
	Url         string          `json:"url"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	AliasName   string          `json:"alias"`
	EMail       string          `json:"email"`
	Telephone   string          `json:"telephone"`
	ERadNumber  string          `json:"e_rad_number"`
	Affiliation UserOrgMetadata `json:"affiliation"`
}

type UsersMatadata struct {
	Users []UserMatadata `json:"users"`
}
type UserOrgMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
	AliasName   string `json:"alias"`
}

type RepositoryMetadata struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Url         string          `json:"url"`
	Download    DownloadMetadat `json:"download"`
}

type DownloadMetadat struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	SHA256      string `json:"sha256"`
	Date        string `json:"date"` //ISO 8601
}

type ResearchPolicyMetadata struct {
	WorkflowIdentifier string `json:"workflow_identifier"`
	ContentSize        string `json:"content_size"`
	DatasetStructure   string `json:"dataset_structure"`
}

type ServiceMetadata struct {
	Name                string `json:"name"`
	BaseUrl             string `json:"base_url"`
	DataAccessUrlPrefix string `json:"data_access_url_prefix"`
}

type ResearchProjectMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FileMetadata struct {
	Path               string `json:"path"`
	Name               string `json:"name"`
	ContentSize        string `json:"content_size"`
	EncodingFormat     string `json:"encoding_format"`
	SHA256             string `json:"sha256"`
	Date               string `json:"date"` //ISO 8601
	IsExperimentPakage bool   `json:"is_experiment_pakage"`
}

type WholeMetadata struct {
	Service                ServiceMetadata         `json:"service"`
	ResearchProject        ResearchProjectMetadata `json:"resarch_project"`
	ResearchPolicyMetadata ResearchPolicyMetadata  `json:"research_policy"`
	RepositoryMetadata     RepositoryMetadata      `json:"repository"`
	Users                  []UserMatadata          `json:"users"`
	Files                  []FileMetadata          `json:"files"`
}

/*
 Form structure
**/
type UserNameList struct {
	UsersName []string `json:"users_name"`
}

type Repository struct {
	OwnerName  string `json:"owner_name"`
	RepoName   string `json:"repo_name"`
	BranchName string `json:"branch_name"`
}

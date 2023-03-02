package datastruct

type ResearchProject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FunderOrg struct {
	Type        string `json:"type"`
	ID          string `json:"@id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AliasName   string `json:"alias"`
}

type ResearchOrg struct {
	ID          string `json:"@id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AliasName   string `json:"alias"`
}

type License struct {
	ID          string `json:"@id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DataDownload struct {
	ID          string `json:"@id"`
	Description string `json:"description"`
	SHA256      string `json:"sha256"`
	Date        string `json:"date"`
}

type RepositoryObject struct {
	ID          string `json:"@id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type HostingInstitution struct {
	ID          string `json:"@id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

type Person struct {
	ID                   string `json:"id"`
	Url                  string `json:"url"`
	Name                 string `json:"name"`
	Alias                string `json:"alias"`
	Affiliation          string `json:"affiliation"`
	Email                string `json:"email"`
	Telephone            string `json:"telephone"`
	ERadResearcherNumber string `json:"eradResearcherNumber"`
}

type File struct {
	ID                    string `json:"@id"`
	Name                  string `json:"name"`
	ContentSize           string `json:"contentSize"`
	EncodingFormat        string `json:"encodingFormat"`
	Sha256                string `json:"sha256"`
	Url                   string `json:"url"`
	SdDatePublished       string `json:"sdDatePublished"`
	ExperimentPackageFlag bool   `json:"experimentPackageFlag"`
}

type Dataset struct {
	ID   string `json:"@id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type GinMonitoring struct {
	ContentSize        string `json:"contentSize"`
	WorkflowIdentifier string `json:"workflowIdentifier"`
	DatasetStructure   string `json:"datasetStructure"`
}

type Metadata struct {
	ResearchProject     ResearchProject      `json:"research_project"`
	FunderOrgs          []FunderOrg          `json:"funder_orgs"`
	ResearchOrgs        []ResearchOrg        `json:"research_orgs"`
	Licenses            []License            `json:"licenses"`
	DataDownloads       []DataDownload       `json:"data_downloads"`
	RepositoryObjects   []RepositoryObject   `json:"repository_objs"`
	HostingInstitutions []HostingInstitution `json:"hosting_institutions"`
	Persons             []Person             `json:"persons"`
	Files               []File               `json:"files"`
	Datasets            []Dataset            `json:"datasets"`
	GinMonitorings      []GinMonitoring      `json:"gin_monitorings"`
	Dmps                []interface{}        `json:"dmps"`
}

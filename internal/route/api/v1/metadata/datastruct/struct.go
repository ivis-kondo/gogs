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

type IFDmp interface {
}

type CaoDmp struct {
	Type          string       `json:"type"`
	Repository    string       `json:"repository"`
	Distribution  string       `json:"distribution"`
	Keyword       string       `json:"keyword"`
	ERadProjectId string       `json:"eradProjectId"`
	HasPart       []CaoDmpData `json:"hasPart"`
}

type CaoDmpData struct {
	Name                string           `json:"name"`
	Description         string           `json:"description"`
	Creator             []string         `json:"creator"`
	Keyword             string           `json:"keyword"`
	AccessRights        string           `json:"accessRights"`
	AvailabilityStarts  string           `json:"availabilityStarts"`
	IsAccessibleForFree string           `json:"isAccessibleForFree"`
	License             string           `json:"license"`
	UsageInfo           string           `json:"usageInfo"`
	Repository          string           `json:"repository"`
	Distribution        string           `json:"distribution"`
	ContentSize         string           `json:"contentSize"`
	HostingInstitution  string           `json:"hostingInstitution"`
	DataManager         string           `json:"dataManager"`
	Relateddata         []DmpRelatedData `json:"related_data"`
}

type MetiDmp struct {
	Type         string        `json:"type"`
	Creator      []string      `json:"creator"`
	Repository   string        `json:"repository"`
	Distribution string        `json:"distribution"`
	HasPart      []MetiDmpData `json:"hasPart"`
}

type MetiDmpData struct {
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	HostingInstitution   string           `json:"hostingInstitution"`
	WayOfManage          string           `json:"wayOfManage"`
	AccessRights         string           `json:"accessRights"`
	ReasonForConcealment string           `json:"reasonForConcealment"`
	AvailabilityStarts   string           `json:"availabilityStarts"`
	Creator              []string         `json:"creator"`
	MeasurementTechnique string           `json:"measurementTechnique"`
	IsAccessibleForFree  string           `json:"isAccessibleForFree"`
	License              string           `json:"license"`
	UsageInfo            string           `json:"usageInfo"`
	Repository           string           `json:"repository"`
	ContentSize          string           `json:"contentSize"`
	Distribution         string           `json:"distribution"`
	ContactPoint         ContactPoint     `json:"contactPoint"`
	Relateddata          []DmpRelatedData `json:"related_data"`
}

type ContactPoint struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type AmedDmp struct {
	Type string `json:"type"`

	Funding            string        `json:"funding"`
	ChiefResearcher    string        `json:"chiefResearcher"`
	Creator            []string      `json:"creator"`
	HostingInstitution string        `json:"hostingInstitution"`
	DataManager        string        `json:"dataManager"`
	Repository         string        `json:"repository"`
	Distribution       string        `json:"distribution"`
	HasPart            []AmedDmpData `json:"hasPart"`
}

type AmedDmpData struct {
	Name                  string                         `json:"name"`
	Description           string                         `json:"description"`
	Keyword               string                         `json:"keyword"`
	AccessRights          string                         `json:"accessRights"`
	AvailabilityStarts    string                         `json:"availabilityStarts"`
	ReasonForConcealment  string                         `json:"reasonForConcealment"`
	Repository            string                         `json:"repository"`
	Distribution          string                         `json:"distribution"`
	ContentSize           string                         `json:"contentSize"`
	GotInformedConsent    string                         `json:"gotInformedConsent"`
	InformedConsentFormat string                         `json:"informedConsentFormat"`
	Identifier            []ClinicalResearchRegistration `json:"identifier"`
	RelatedData           []DmpRelatedData               `json:"related_data"`
}

type ClinicalResearchRegistration struct {
	ID    string `json:"@id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DmpRelatedData struct {
	ID              string `json:"@id"`
	Name            string `json:"name"`
	ContentSize     string `json:"contentSize"`
	EncodingFormat  string `json:"encodingFormat"`
	Sha256          string `json:"sha256"`
	Url             string `json:"url"`
	SdDatePublished string `json:"sdDatePublished"`
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
	GinMonitorings      GinMonitoring        `json:"gin_monitoring"`
	Dmps                []IFDmp              `json:"dmps"`
}

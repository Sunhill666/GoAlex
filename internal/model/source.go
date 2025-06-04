package model

type Source struct {
	DehydratedSource
	AbbreviatedTitle string   `json:"abbreviated_title"`
	AlternateTitles  []string `json:"alternate_titles"`
	APCPrice         struct {
		Price    int    `json:"price"`
		Currency string `json:"currency"`
	} `json:"apc_price"`
	APCUSD       int            `json:"apc_usd"`
	CitedByCount int            `json:"cited_by_count"`
	CountryCode  string         `json:"country_code"`
	CountsByYear []CountsByYear `json:"counts_by_year"`
	CreatedDate  string         `json:"created_date"`
	HomePageURL  string         `json:"homepage_url"`
	IDs          struct {
		Fatcat    string   `json:"fatcat"`
		ISSN      []string `json:"issn"`
		ISSNL     string   `json:"issn_l"`
		MAG       string   `json:"mag"`
		OpenAlex  string   `json:"openalex"`
		Wikipedia string   `json:"wikipedia"`
	} `json:"ids"`
	Societies []struct {
		URL          string `json:"url"`
		Organization string `json:"organization"`
	} `json:"societies"`
	SummaryStats SummaryStats                 `json:"summary_stats"`
	Topics       []TopicWithCount             `json:"topics"`
	TopicShare   []TopicShare                 `json:"topic_share"`
	UpdatedDate  string                       `json:"updated_date"`
	WorksAPIURL  string                       `json:"works_api_url"`
	WorksCount   int                          `json:"works_count"`
	XConcepts    []DehydratedConceptWithScore `json:"x_concepts"`
}

type DehydratedSource struct {
	RepositorySource
	IsCore   bool     `json:"is_core"`
	IsInDOAJ bool     `json:"is_in_doaj"`
	IsOA     bool     `json:"is_oa"`
	ISSN     []string `json:"issn"`
	ISSNL    string   `json:"issn_l"`
	Type     string   `json:"type"`
}

type RepositorySource struct {
	DisplayName             string   `json:"display_name"`
	HostOrganization        string   `json:"host_organization"`
	HostOrganizationLineage []string `json:"host_organization_lineage"`
	HostOrganizationName    string   `json:"host_organization_name"`
	ID                      string   `json:"id"`
}

package model

type Institution struct {
	DehydratedInstitution
	AssociatedInstitutions  []DehydratedInstitutionWithRelationship `json:"associated_institutions"`
	CitedByCount            int                                     `json:"cited_by_count"`
	CountsByYear            []CountsByYear                          `json:"counts_by_year"`
	CreatedDate             string                                  `json:"created_date"`
	DisplayNameAcronyms     []string                                `json:"display_name_acronyms"`
	DisplayNameAlternatives []string                                `json:"display_name_alternatives"`
	GEO                     GEO                                     `json:"geo"`
	HomePageURL             string                                  `json:"homepage_url"`
	IDs                     struct {
		Grid      string `json:"grid"`
		MAG       string `json:"mag"`
		OpenAlex  string `json:"openalex"`
		ROR       string `json:"ror"`
		Wikipedia string `json:"wikipedia"`
		Wikidata  string `json:"wikidata"`
	} `json:"ids"`
	ImageThumbnailURL string `json:"image_thumbnail_url"`
	ImageURL          string `json:"image_url"`
	International     struct {
		DisplayName map[string]string `json:"display_name"`
	} `json:"international"`
	IsSuperSystem bool                         `json:"is_super_system"`
	Repositories  []RepositorySource           `json:"repositories"`
	Roles         []Role                       `json:"roles"`
	SummaryStats  SummaryStats                 `json:"summary_stats"`
	Topics        []TopicWithCount             `json:"topics"`
	TopicShare    []TopicShare                 `json:"topic_share"`
	UpdatedDate   string                       `json:"updated_date"`
	WorksAPIURL   string                       `json:"works_api_url"`
	WorksCount    int                          `json:"works_count"`
	XConcepts     []DehydratedConceptWithScore `json:"x_concepts"`
}

type DehydratedInstitution struct {
	CountryCode string   `json:"country_code"`
	DisplayName string   `json:"display_name"`
	ID          string   `json:"id"`
	Lineage     []string `json:"lineage"`
	ROR         string   `json:"ror"`
	Type        string   `json:"type"`
}

type DehydratedInstitutionWithRelationship struct {
	DehydratedInstitution
	Relationship string `json:"relationship"`
}

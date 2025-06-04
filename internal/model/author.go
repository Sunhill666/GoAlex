package model

type Author struct {
	DehydratedAuthor
	Affiliation []struct {
		Institution DehydratedInstitution `json:"institution"`
		Years       []int                 `json:"years"`
	}
	CitedByCount            int            `json:"cited_by_count"`
	CountsByYear            []CountsByYear `json:"counts_by_year"`
	CreatedDate             string         `json:"created_date"`
	DisplayNameAlternatives []string       `json:"display_name_alternatives"`
	IDs                     struct {
		OpenAlex  string `json:"openalex"`
		ORCID     string `json:"orcid"`
		Scopus    string `json:"scopus"`
		Twitter   string `json:"twitter"`
		Wikipedia string `json:"wikipedia"`
	} `json:"ids"`
	LastKnownInstitutions []DehydratedInstitution      `json:"last_known_institutions"`
	SummaryStats          SummaryStats                 `json:"summary_stats"`
	Topics                []TopicWithCount             `json:"topics"`
	TopicShare            []TopicShare                 `json:"topic_share"`
	UpdatedDate           string                       `json:"updated_date"`
	WorksAPIURL           string                       `json:"works_api_url"`
	WorksCount            int                          `json:"works_count"`
	XConcepts             []DehydratedConceptWithScore `json:"x_concepts"`
}

type DehydratedAuthor struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	ORCID       string `json:"orcid"`
}

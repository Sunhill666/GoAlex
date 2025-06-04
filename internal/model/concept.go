package model

type Concept struct {
	Ancestors    []DehydratedConcept `json:"ancestors"`
	CitedByCount int                 `json:"cited_by_count"`
	CountsByYear []CountsByYear      `json:"counts_by_year"`
	CreatedDate  string              `json:"created_date"`
	Description  string              `json:"description"`
	IDs          struct {
		MAG       string   `json:"mag"`
		OpenAlex  string   `json:"openalex"`
		UMLSCUI   []string `json:"umls_cui"`
		UMLSAUI   []string `json:"umls_aui"`
		Wikidata  string   `json:"wikidata"`
		Wikipedia string   `json:"wikipedia"`
	} `json:"ids"`
	ImageThumbnailURL string `json:"image_thumbnail_url"`
	ImageURL          string `json:"image_url"`
	International     struct {
		DisplayName map[string]string `json:"display_name"`
		Description map[string]string `json:"description"`
	} `json:"international"`
	Level           int                          `json:"level"`
	RelatedConcepts []DehydratedConceptWithScore `json:"related_concepts"`
	SummaryStats    SummaryStats                 `json:"summary_stats"`
	UpdatedDate     string                       `json:"updated_date"`
	Wikidata        string                       `json:"wikidata"`
	WorksAPIURL     string                       `json:"works_api_url"`
	WorksCount      int                          `json:"works_count"`
}

type DehydratedConcept struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Level       int    `json:"level"`
	Wikidata    string `json:"wikidata"`
}

type DehydratedConceptWithScore struct {
	DehydratedConcept
	Score float32 `json:"score"`
}

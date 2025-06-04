package model

type Affiliation struct {
	InstitutionIDs       []string `json:"institution_ids"`
	RawAffiliationString string   `json:"raw_affiliation_string"`
}

type Authorship struct {
	Affiliations          []Affiliation           `json:"affiliations"`
	Author                DehydratedAuthor        `json:"author"`
	AuthorPosition        string                  `json:"author_position"`
	Countries             []string                `json:"countries"`
	Institution           []DehydratedInstitution `json:"institutions"`
	IsCorresponding       bool                    `json:"is_corresponding"`
	RawAffiliationStrings []string                `json:"raw_affiliation_strings"`
	RawAuthorName         string                  `json:"raw_author_name"`
}

type IDs struct {
	DOI      string `json:"doi"`
	MAG      string `json:"mag"`
	OpenAlex string `json:"openalex"`
	PMCID    string `json:"pmcid"`
	PMID     string `json:"pmid"`
}

type Location struct {
	IsAccepted     bool             `json:"is_accepted"`
	IsOA           bool             `json:"is_oa"`
	IsPublished    bool             `json:"is_published"`
	LandingPageURL string           `json:"landing_page_url"`
	License        string           `json:"license"`
	PDF_URL        string           `json:"pdf_url"`
	Source         DehydratedSource `json:"source"`
	Version        string           `json:"version"`
}

type MeSH struct {
	DescriptorUI   string `json:"descriptor_ui"`
	DescriptorName string `json:"descriptor_name"`
	QualifierUI    string `json:"qualifier_ui"`
	QualifierName  string `json:"qualifier_name"`
	IsMajorTopic   bool   `json:"is_major_topic"`
}

type OAStatus string

const (
	OAStatusClosed  OAStatus = "closed"
	OAStatusBronze  OAStatus = "bronze"
	OAStatusDiamond OAStatus = "diamond"
	OAStatusGold    OAStatus = "gold"
	OAStatusGreen   OAStatus = "green"
	OAStatusHybrid  OAStatus = "hybrid"
)

type OpenAccess struct {
	AnyRepoHasFulltext bool     `json:"any_repository_has_fulltext"`
	IsOA               bool     `json:"is_oa"`
	OAStatus           OAStatus `json:"oa_status"`
	OAURL              string   `json:"oa_url"`
}

type Work struct {
	Abstract              string           `json:"abstract"`
	AbstractInvertedIndex map[string][]int `json:"abstract_inverted_index"`
	Authorships           []Authorship     `json:"authorships"`
	APCList               APC              `json:"apc_list"`
	APCPaid               APC              `json:"apc_paid"`
	BestOALocation        Location         `json:"best_oa_location"`
	Biblio                struct {
		Volume    string `json:"volume"`
		Issue     string `json:"issue"`
		FirstPage string `json:"first_page"`
		LastPage  string `json:"last_page"`
	} `json:"biblio"`
	CitationNormalizedPercentile struct {
		IsTop1Percent  bool    `json:"is_in_top_1_percent"`
		IsTop10Percent bool    `json:"is_in_top_10_percent"`
		Value          float64 `json:"value"`
	} `json:"citation_normalized_percentile"`
	CitedByAPIURL               string                       `json:"cited_by_api_url"`
	CitedByCount                int                          `json:"cited_by_count"`
	Concepts                    []DehydratedConceptWithScore `json:"concepts"`
	CorrespondingAuthorIDs      []string                     `json:"corresponding_author_ids"`
	CorrespondingInstitutionIDs []string                     `json:"corresponding_institution_ids"`
	CountriesDistinctCount      int                          `json:"countries_distinct_count"`
	CountsByYear                []struct {
		CitedByCount int `json:"cited_by_count"`
		Year         int `json:"year"`
	} `json:"counts_by_year"`
	CreatedDate    string  `json:"created_date"`
	DisplayName    string  `json:"display_name"`
	DOI            string  `json:"doi"`
	FulltextOrigin string  `json:"fulltext_origin"`
	FWCI           float32 `json:"fwci"`
	Grants         []struct {
		AwardID           string `json:"award_id"`
		Funder            string `json:"funder"`
		FunderDisplayName string `json:"funder_display_name"`
	} `json:"grants"`
	HasFulltext               bool                `json:"has_fulltext"`
	ID                        string              `json:"id"`
	IDs                       IDs                 `json:"ids"`
	IndexedIn                 []string            `json:"indexed_in"`
	InstitutionsDistinctCount int                 `json:"institutions_distinct_count"`
	IsParatext                bool                `json:"is_paratext"`
	IsRetracted               bool                `json:"is_retracted"`
	Keywords                  []DehydratedKeyword `json:"keywords"`
	Language                  string              `json:"language"`
	License                   string              `json:"license"`
	Locations                 []Location          `json:"locations"`
	LocationCount             int                 `json:"location_count"`
	MeSH                      []MeSH              `json:"mesh"`
	OpenAccess                OpenAccess          `json:"open_access"`
	PrimaryLocation           Location            `json:"primary_location"`
	PrimaryTopic              TopicWithScore      `json:"primary_topic"`
	PublicationDate           string              `json:"publication_date"`
	PublicationYear           int                 `json:"publication_year"`
	ReferenceWorks            []string            `json:"reference_works"`
	RelatedWorks              []string            `json:"related_works"`
	SDGs                      []struct {
		DisplayName string  `json:"display_name"`
		ID          string  `json:"id"`
		Score       float64 `json:"score"`
	} `json:"sustainable_development_goals"`
	Title        string           `json:"title"`
	Topics       []TopicWithScore `json:"topics"`
	Type         string           `json:"type"`
	TypeCrossref string           `json:"type_crossref"`
	UpdatedDate  string           `json:"updated_date"`
}

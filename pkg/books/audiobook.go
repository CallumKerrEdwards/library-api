package books

type Audiobook struct {
	AudiobookMediaID             string   `json:"audiobookMediaId"`
	Narrators                    []Person `json:"narrators"`
	CoverImageMediaID            string   `json:"coverImageMediaId"`
	SupplimentaryMaterialMediaID string   `json:"supplimentaryMaterialMediaId,omitempty"`
}

func (a Audiobook) GetNarrator() string {
	return GetPersonsString(a.Narrators)
}

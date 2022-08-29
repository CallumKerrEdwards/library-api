package books

type Audiobook struct {
	PathToAudioFile             string   `json:"pathToAudioFile"`
	Narrators                   []Person `json:"narrators"`
	PathToCoverImage            string   `json:"pathToCoverImage"`
	PathToSupplimentaryMaterial string   `json:"pathToSupplimentaryMaterial,omitempty"`
}

func (a Audiobook) GetPath() string {
	return a.PathToAudioFile
}

package web

type IntermediateResponse struct {
	Id            int    `json:"id,omitempty"`
	PohonId       int    `json:"pohon_id"`
	FaktorOutcome string `json:"faktor_outcome"`
	DataTerukur   string `json:"data_terukur"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

package web

type OutcomeResponse struct {
	Id            int    `json:"id,omitempty"`
	PohonId       int    `json:"pohon_id"`
	Tahun         string `json:"tahun"`
	FaktorOutcome string `json:"faktor_outcome"`
	DataTerukur   string `json:"data_terukur"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

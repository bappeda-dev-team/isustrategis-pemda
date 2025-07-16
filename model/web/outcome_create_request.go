package web

type OutcomeCreateRequest struct {
	PohonId       int    `json:"pohon_id"`
	FaktorOutcome string `json:"faktor_outcome"`
	DataTerukur   string `json:"data_terukur"`
}

package web

type IntermediateUpdateRequest struct {
	Id            int    `json:"id"`
	PohonId       int    `json:"pohon_id"`
	FaktorOutcome string `json:"faktor_outcome"`
	DataTerukur   string `json:"data_terukur"`
}

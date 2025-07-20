package web

type IntermediateUpdateRequest struct {
	Id            int    `json:"id"`
	PohonId       int    `json:"pohon_id" validate:"required"`
	Tahun         string `json:"tahun" validate:"required"`
	FaktorOutcome string `json:"faktor_outcome" validate:"required"`
	DataTerukur   string `json:"data_terukur" validate:"required"`
}

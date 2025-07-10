package web

type CsfCreateRequest struct {
	PohonId                    int                          `json:"pohon_id" validate:"required"`
	PernyataanKondisiStrategis string                       `json:"pernyataan_kondisi_strategis" validate:"required"`
	AlasanKondisi              []AlasanKondisiCreateRequest `json:"alasan_kondisi"`
}

type AlasanKondisiCreateRequest struct {
	CSFid                          int                        `json:"csf_id" validate:"required"`
	AlasanKondisiStrategis         string                     `json:"alasan_kondisi_strategis" validate:"required"`
	DataTerukurPendukungPernyataan []DataTerukurCreateRequest `json:"data_terukur_pendukung_pernyataan"`
}

type DataTerukurCreateRequest struct {
	AlasanKondisiId int    `json:"alasan_kondisi_id" validate:"required"`
	DataTerukur     string `json:"data_terukur" validate:"required"`
}

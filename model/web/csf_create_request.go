package web

type CsfCreateRequest struct {
	PohonId                    int                          `json:"pohon_id" validate:"required"`
	PernyataanKondisiStrategis string                       `json:"pernyataan_kondisi_strategis" validate:"required"`
	Tahun                      string                       `json:"tahun" validate:"required"`
	AlasanKondisi              []AlasanKondisiCreateRequest `json:"alasan_kondisi"`
}

type AlasanKondisiCreateRequest struct {
	CSFid                          int                        `json:"id" `
	AlasanKondisiStrategis         string                     `json:"alasan_kondisi_strategis"`
	DataTerukurPendukungPernyataan []DataTerukurCreateRequest `json:"data_terukur"`
}

type DataTerukurCreateRequest struct {
	AlasanKondisiId int    `json:"id"`
	DataTerukur     string `json:"data_terukur"`
}

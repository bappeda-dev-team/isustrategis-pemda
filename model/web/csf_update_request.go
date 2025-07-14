package web

type CsfUpdateRequest struct {
	Id                         int                          `json:"id" validate:"required"`
	PohonId                    int                          `json:"pohon_id" validate:"required"`
	PernyataanKondisiStrategis string                       `json:"pernyataan_kondisi_strategis" validate:"required"`
	AlasanKondisi              []AlasanKondisiUpdateRequest `json:"alasan_kondisi"`
}

type AlasanKondisiUpdateRequest struct {
	Id                             int                        `json:"id"`
	AlasanKondisiStrategis         string                     `json:"alasan_kondisi_strategis"`
	DataTerukurPendukungPernyataan []DataTerukurUpdateRequest `json:"data_terukur"`
}

type DataTerukurUpdateRequest struct {
	Id          int    `json:"id"`
	DataTerukur string `json:"data_terukur"`
}

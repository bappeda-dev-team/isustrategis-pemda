package web

import "time"

type CsfResponse struct {
	Id                         int                     `json:"id"`
	PohonId                    int                     `json:"pohon_id"`
	PernyataanKondisiStrategis string                  `json:"pernyataan_kondisi_strategis"`
	Tahun                      string                  `json:"tahun"`
	AlasanKondisi              []AlasanKondisiResponse `json:"alasan_kondisi"`
	CreatedAt                  time.Time               `json:"created_at"`
	UpdatedAt                  time.Time               `json:"updated_at"`
}

type AlasanKondisiResponse struct {
	Id                             int                   `json:"id"`
	CSFid                          int                   `json:"csf_id"`
	AlasanKondisiStrategis         string                `json:"alasan_kondisi_strategis"`
	DataTerukurPendukungPernyataan []DataTerukurResponse `json:"data_terukur"`
	CreatedAt                      time.Time             `json:"created_at"`
	UpdatedAt                      time.Time             `json:"updated_at"`
}

type DataTerukurResponse struct {
	Id              int       `json:"id"`
	AlasanKondisiId int       `json:"alasan_kondisi_id"`
	DataTerukur     string    `json:"data_terukur"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

package domain

import "time"

type Csf struct {
	Id                         int
	PohonId                    int
	PernyataanKondisiStrategis string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
	AlasanKondisi              []AlasanKondisi
}

type AlasanKondisi struct {
	Id                             int
	CSFid                          int
	AlasanKondisiStrategis         string
	DataTerukurPendukungPernyataan []DataTerukurPendukungPernyataan
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}

type DataTerukurPendukungPernyataan struct {
	Id              int
	AlasanKondisiId int
	DataTerukur     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

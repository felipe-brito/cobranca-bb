package response

import "github.com/felipe-brito/cobranca-bb/internal/model"

type ManagerHealth struct {
	Status    string                 `json:"status"`
	DB        ManagerHealthDB        `json:"db"`
	DiskSpace ManagerHealthDiskSpace `json:"disk_space"`
}

type ManagerHealthDB struct {
	Status string `json:"status"`
}

type ManagerHealthDiskSpace struct {
	Status string `json:"status"`
	Total  uint64 `json:"total"`
	Free   uint64 `json:"free"`
}

func ConvertTo(m model.ManagerHealth) ManagerHealth {
	return ManagerHealth{
		Status: m.Status,
		DB:     ManagerHealthDB{Status: m.DB.Status},
		DiskSpace: ManagerHealthDiskSpace{
			Status: m.DiskSpace.Status,
			Total:  m.DiskSpace.Total,
			Free:   m.DiskSpace.Free,
		},
	}
}

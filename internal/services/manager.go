package services

import (
	"github.com/felipe-brito/cobranca-bb/internal/model"
	"github.com/felipe-brito/cobranca-bb/internal/utils"
	"github.com/ricochet2200/go-disk-usage/du"
)

type Manager struct {
}

func NewManager() Manager {
	return Manager{}
}

func (m Manager) Health() model.ManagerHealth {

	db := healthDB()
	disk := healthDiskUsage()

	status := utils.UP

	if (db.Status != utils.UP) || (disk.Status != utils.UP) {
		status = utils.DOWN
	}

	return model.ManagerHealth{
		Status:    status,
		DB:        db,
		DiskSpace: disk,
	}

}

func healthDB() model.ManagerHealthDB {
	return model.ManagerHealthDB{
		Status: utils.UP,
	}
}

func healthDiskUsage() model.ManagerHealthDiskSpace {
	usage := du.NewDiskUsage("/")
	status := utils.UP

	if (usage.Usage() * 100) > 90 {
		status = utils.DOWN
	}

	return model.ManagerHealthDiskSpace{
		Status: status,
		Total:  usage.Size() / (utils.KB * utils.KB),
		Free:   usage.Free() / (utils.KB * utils.KB),
	}
}

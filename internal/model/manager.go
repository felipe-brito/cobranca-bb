package model

type ManagerHealth struct {
	Status    string
	DB        ManagerHealthDB
	DiskSpace ManagerHealthDiskSpace
}

type ManagerHealthDB struct {
	Status string
}

type ManagerHealthDiskSpace struct {
	Status string
	Total  uint64
	Free   uint64
}

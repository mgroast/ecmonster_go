package models

type Monster struct {
	ID string
	Name string
}

func (m Monster) FindMonsters() *[]Monster {
	db := GormDB()
	monsters := &[]Monster{}
	db.Find(monsters)
	return monsters
}

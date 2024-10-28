package repository

type Repository interface {
	GetAll()
	GetByID()
	Add()
	Update()
	DeleteByID()
}

type GormRepositoryMySQL struct {
}

func NewGormRepositoryMySQL() Repository {
	return &GormRepositoryMySQL{}
}

func (g *GormRepositoryMySQL) GetAll() {

}

func (g *GormRepositoryMySQL) GetByID() {

}
func (g *GormRepositoryMySQL) Add() {

}
func (g *GormRepositoryMySQL) Update() {

}
func (g *GormRepositoryMySQL) DeleteByID() {

}

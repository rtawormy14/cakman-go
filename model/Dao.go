package model

type Daoer interface {
	GetByID(interface{})
	Set(interface{})
}

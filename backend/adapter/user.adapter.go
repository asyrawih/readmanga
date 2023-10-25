package adapter

import (
	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

// Contract User Repo
type RepoUserCreational interface {
	Creational[entity.User, int]
	Modificational[entity.User, int]
	Retrival[entity.User, int]
	Destroyer[int]
	Accessable[mysql.UserRepository]
}

// Contract User Repo
type ServiceUserCreational interface {
	Creational[entity.User, int]
	Modificational[entity.User, int]
	Retrival[entity.User, int]
	Destroyer[int]
}

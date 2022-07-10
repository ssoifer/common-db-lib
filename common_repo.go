package common_db_lib

import "github.com/ssoifer/common-db-lib/db"

//func GMin[T constraints.Ordered](x, y T, error) T {
//	if x < y {
//		return x
//	}
//	return y
//}

type CommonRepository interface {
	Save(any) (dbModel *any, err error)
	Update(any) (dbModel *any, err error)
	GetList(columns []string, table string) (dbModelList []interface{}, err error)
	GetById(any) (dbModel any, err error)
}

func GetRepository(repoType db.RepositoryType) (CommonRepository, error) {
	if repoType == db.RepositoryTypeDB {
		return db.NewRepository()
	} else {
		return nil, nil //errors.New("Invalid Repository Type provided")
	}
}

package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

//var saveQuery = fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1,$2,$3,$4,$5) RETURNING %s;", repositoryTableName, insertFields, returnFields)
//var getQuery = fmt.Sprintf("SELECT %s FROM %s WHERE service_tag = $1;", returnFields, repositoryTableName)

type databaseRepository struct {
	con *sql.DB
}

const (
	repositoryTableName = "task"
	insertFields        = " id, content, title, views, timestamp "
	returnFields        = insertFields
)

func NewRepository() (*databaseRepository, error) {
	ctx := context.Background()
	config := &Config{}
	*config = parseEnv()
	connection, err := NewDatabase(*config)
	if err != nil {
		log.Print(ctx, err, "unable to setup database")
		return nil, errors.New("unable to setup database")
	}

	err = Migrate(connection, *config)
	if err != nil {
		log.Print(ctx, err, "unable to setup database")
		return nil, errors.New("failed during migration")
	}
	return &databaseRepository{connection}, nil
}

func (repo *databaseRepository) Save(any) (dbModel *any, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

func (repo *databaseRepository) Update(any) (dbModel *any, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

//func GenericSelect(database string, table string, columns []string, result interface{}) interface{} {
//	dbMap := getDBConnection(database)
//	defer dbMap.Db.Close()
//	var err error
//	query := "SELECT "
//
//	for index, element := range columns {
//		query += element
//		if index+1 != len(columns) {
//			query += ","
//		}
//	}
//	query += " FROM " + table + " LIMIT 1,100"
//	_, err = dbMap.Select(result, query)
//	if err != nil {
//		panic(err.Error()) // Just for example purpose.
//	}
//	return result
//}

func (repo *databaseRepository) GetList(columns []string, table string) (dbModelList []interface{}, err error) {
	//	baseQuery := fmt.Sprintf("SELECT %s FROM %s ", returnFields, repositoryTableName)

	query := "SELECT "

	for index, element := range columns {
		query += element
		if index+1 != len(columns) {
			query += ","
		}
	}
	query += " FROM " + table + " LIMIT 1,100"

	var rows *sql.Rows
	listQuery := fmt.Sprintf(query)
	rows, err = repo.con.Query(listQuery)
	if err != nil {
		return dbModelList, err
	}

	var list []interface{}
	for rows.Next() {
		var item interface{}
		err := rows.Scan(&item)
		if err != nil {
			return dbModelList, err
		}
		list = append(list, item)
	}
	if err = rows.Err(); err != nil {
		return dbModelList, err
	}
	return list, nil
}

func (repo *databaseRepository) GetById(any) (dbModel any, err error) {
	//TODO: update below print statement with database persist functionality
	return nil, err
}

func parseEnv() Config {
	dbHost := os.Getenv("DB-HOST")
	dbPort := os.Getenv("DB-PORT")
	dbUser := os.Getenv("DB-USER")
	dbPassword := os.Getenv("DB-PASSWORD")
	database := os.Getenv("DATABASE")
	if dbHost == "" && dbPort == "" && dbUser == "" && dbPassword == "" && database == "" {
		//error - missing env vars
	}
	return Config{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Database: database,
	}
}

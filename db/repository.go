package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err == nil {

	}
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

func (repo *databaseRepository) GetList() (dbModelList []any, err error) {
	//	baseQuery := fmt.Sprintf("SELECT %s FROM %s ", returnFields, repositoryTableName)
	//
	//	var rows *sql.Rows
	//	var err error = nil
	//	listQuery := fmt.Sprintf(baseQuery)
	//	rows, err = repo.con.Query(listQuery)
	//	if err != nil {
	//		return []dbModel.Task{}, err
	//	}
	//
	//	var list []dbModel.Task
	//	for rows.Next() {
	//		var taskItem dbModel.Task
	//		err := rows.Scan(&taskItem.ID, &taskItem.Title, &taskItem.Content, &taskItem.Views, &taskItem.Timestamp)
	//		if err != nil {
	//			return []dbModel.Task{}, err
	//		}
	//		list = append(list, taskItem)
	//	}
	//	if err = rows.Err(); err != nil {
	//		return []dbModel.Task{}, err
	//	}
	return nil, nil
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

package golang

import (
	"darklab_training_postgres/golang/shared/utils"

	"gorm.io/gorm"
)

var (
	Task4Query1   string
	Task4Query1_2 string
	Task4Query2   string
	Task4Query3   string
	Task4Query4   string
	Task4Query5   string
	Task4Query6   string
)

func init() {
	Task4Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1.sql"))
	Task4Query1_2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_1_2.sql"))
	Task4Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_2.sql"))
	Task4Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_3.sql"))
	Task4Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_4.sql"))
	Task4Query5 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_5.sql"))
	Task4Query6 = utils.GetSQLFile(utils.ReadProjectFile("sql/task4/queries/query_4_6.sql"))
}

func CountRows(result *gorm.DB) int {
	rows, err := result.Rows()
	if err != nil {
		panic(err)
	}
	count_rows := 0
	for rows.Next() {
		count_rows += 1
	}
	return count_rows
}

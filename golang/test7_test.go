package golang

import "darklab_training_postgres/golang/shared/utils"

var (
	Task7Query1 string
	Task7Query2 string
	Task7Query3 string
	Task7Query4 string
)

func init() {
	Task7Query1 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_1.sql"))
	Task7Query2 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_2.sql"))
	Task7Query3 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_3.sql"))
	Task7Query4 = utils.GetSQLFile(utils.ReadProjectFile("sql/task7/queries/query_7_4.sql"))
}

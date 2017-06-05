package tables

/*
TableTasks - initialized table tasks
*/
var TableTasks Tasks

/*
Tasks table representation
*/
type Tasks struct {
	DbTableGen
	Consts TasksConst
}

func init() {
	TableTasks = Tasks{}
	TableTasks.Consts = TasksConst{
		TableName:    "tasks",
		ColID:        "id",
		ColSubject:   "subject",
		ColCompleted: "completed",
		ColDue:       "due"}

	TableTasks.Consts.QueryGetAll = "SELECT " + TableTasks.Consts.ColID + "," + TableTasks.Consts.ColSubject + "," + TableTasks.Consts.ColCompleted + "," + TableTasks.Consts.ColDue + " FROM " + TableTasks.Consts.TableName
	TableTasks.Consts.PrepareInsert = "INSERT " + TableTasks.Consts.TableName + " SET " + TableTasks.Consts.ColSubject + "=?"
	TableTasks.Consts.PrepareGetByID = TableTasks.Consts.QueryGetAll + " WHERE " + TableTasks.Consts.ColID + "="
	TableTasks.Consts.PrepareDeleteByID = "DELETE FROM " + TableTasks.Consts.TableName + " WHERE " + TableTasks.Consts.ColID + "="

	TableTasks.Name = TableTasks.Consts.TableName
	TableTasks.Script = "CREATE TABLE tasks (id int(11) NOT NULL AUTO_INCREMENT, subject varchar(250) NOT NULL, completed tinyint(1) NOT NULL DEFAULT 0, due timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (id)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8"

}

// TasksConst constants fot Task
type TasksConst struct {
	TableName string

	ColID        string
	ColSubject   string
	ColCompleted string
	ColDue       string

	QueryGetAll       string
	PrepareInsert     string
	PrepareGetByID    string
	PrepareDeleteByID string
}

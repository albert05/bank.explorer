package model

import (
	"fmt"
	"bank.explorer/util/mysql"
	"strings"
)

const TaskTable  = "banktasks"
const WorkABC  = 1
const WorkICBC  = 2

func FindTask(id int) mysql.MapModel {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id =%d", TaskTable, id)
	task, err := mysql.Conn.FindOne(sql)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return task
}

func FindTaskListByStatus(status int, wType string) []mysql.MapModel {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE status =%d and work_id in(%s)", TaskTable, status, wType)
	list, err := mysql.Conn.FindAll(sql)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return list
}

func FindTaskListByIds(ids string) []mysql.MapModel {
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id in(%s)", TaskTable, ids)
	list, err := mysql.Conn.FindAll(sql)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return list
}

func UpdateTask(id int, data map[string]string) bool {
	condition := map[string]string {
		"where": fmt.Sprintf("id=%d", id),
	}

	return mysql.Conn.Update(TaskTable, data, condition)
}

func UpdateMultiTask(ids []string, data map[string]string) bool {
	condition := map[string]string {
		"where": fmt.Sprintf("id in(%s)", strings.Join(ids, ",")),
	}

	return mysql.Conn.Update(TaskTable, data, condition)
}

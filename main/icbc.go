package main

import (
	"bank.explorer/service/icbc"
	"bank.explorer/util/dates"
	"bank.explorer/common"
	"strconv"
	"bank.explorer/model"
	"bank.explorer/config"
	"bank.explorer/exception"
	"bank.explorer/service"
)

func main()  {
	service.ConfigInit()

	defer exception.Handle(true)

	id ,_ := strconv.Atoi(config.JobList)
	job := model.FindTask(id)

	actId := job.GetAttrString("product_id")
	cookie := job.GetAttrString("user_key")

	common.Wait(job.GetAttrFloat("time_point"))
	gift := icbc.InitG(cookie, actId)

	i := 0
	for i < 100 {
		result := gift.RUN()
		if result {
			model.UpdateTask(job.GetAttrInt("id"), map[string]string {
				"status": "3",
				"result": "success",
			})
			break
		}

		dates.SleepSecond(5)
		i++
	}

	model.UpdateTask(job.GetAttrInt("id"), map[string]string {
		"status": "2",
		"result": "failed",
	})
}

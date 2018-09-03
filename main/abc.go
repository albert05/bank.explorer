//url := "https://enjoy.abchina.com/yh-web/customer/giftTokenDraw"
//params := `{"sessionId":"{ps_db86dd27589e11ea2b07b3018f10c8e8}_common","ruleNo":"064119","actNo":"999999CXE00064117","discType":"12","actType":"E","appr":"10"}`
//actNo := "999999CXE00064117"
//sessionID := "d46d835dab7daf16ccbc8bc27d5f995e"
package main

import (
	"fmt"
	"bank.explorer/common"
	"log"
	"bank.explorer/util/dates"
	"bank.explorer/service/abc"
	"bank.explorer/model"
	"strconv"
	"bank.explorer/config"
	"bank.explorer/exception"
	"bank.explorer/service"
	"bank.explorer/util/logger"
)

func main() {
	service.ConfigInit()
	defer exception.Handle(true)

	id ,_ := strconv.Atoi(config.JobList)
	job := model.FindTask(id)

	logger.Info(fmt.Sprintf("taskId:[%d] is starting", id))
	defer logger.Info(fmt.Sprintf("taskId:[%d] is end", id))

	giftItem, err := abc.GetGiftDetail(job.GetAttrString("product_id"))
	if err != nil {
		model.UpdateTask(job.GetAttrInt("id"), map[string]string {
			"status": "2",
			"result": err.Error(),
		})
		log.Fatal(err)
	}

	common.Wait(job.GetAttrFloat("time_point"))

	giftItem.SetSession(job.GetAttrString("code"))

	isChooseCard := job.GetAttrString("is_kdb_pay")

	i := 0
	for i < 3 {
		giftRep := giftItem.RunGift(isChooseCard)

		status := 3
		if abc.GiftStatusSUCCESS != giftRep.Status {
			status = 2
		}

		model.UpdateTask(job.GetAttrInt("id"), map[string]string {
			"status": fmt.Sprintf("%d", status),
			"result": giftRep.Result,
		})

		if status == 3 {
			break
		}
		dates.SleepSecond(5)
		i++
	}
}

//url := "https://enjoy.abchina.com/yh-web/customer/giftTokenDraw"
//params := `{"sessionId":"{ps_db86dd27589e11ea2b07b3018f10c8e8}_common","ruleNo":"064119","actNo":"999999CXE00064117","discType":"12","actType":"E","appr":"10"}`
//actNo := "999999CXE00064117"
//sessionID := "d46d835dab7daf16ccbc8bc27d5f995e"
package main

import (
	"fmt"
	"bank.explorer/common"
	"bank.explorer/util/dates"
	"bank.explorer/service/abc"
	"bank.explorer/model"
	"strconv"
	"bank.explorer/config"
	"bank.explorer/exception"
	"bank.explorer/service"
	"bank.explorer/util/logger"
	"log"
)

func main() {
	service.ConfigInit()
	defer exception.Handle(true)

	id ,_ := strconv.Atoi(config.JobList)
	job := model.FindTask(id)

	logger.Info(fmt.Sprintf("taskId:[%d] is starting", id))
	defer logger.Info(fmt.Sprintf("taskId:[%d] is end", id))

	data := job.GetAttrString("extra")
	pId := job.GetAttrString("product_id")

	var giftItem abc.GiftItem
	var err error
	if data != "" {
		giftItem, err = abc.SetItem(pId, data)
	}

	if err != nil {
		i := 0
		for i < 100 {
			giftItem, err = abc.GetGiftDetail(pId)
			if err == nil {
				break
			}

			dates.SleepSecond(5)
			i++

			if i >= 100 {
				model.UpdateTask(job.GetAttrInt("id"), map[string]string{
					"status": "2",
					"result": err.Error(),
				})
				log.Fatal(err)
			}
		}
	}

	common.Wait(job.GetAttrFloat("time_point"))

	giftItem.SetSession(job.GetAttrString("user_key"))

	j := 0
	for j < 100 {
		giftRep := giftItem.RunGift()
		logger.Info(fmt.Sprintf("taskId:[%d] is try result:[%s]", id, giftRep.Result))

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
		j++
	}
}

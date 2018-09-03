package abc

import (
	"encoding/json"
	"errors"
	"fmt"
	"bank.explorer/util/https"
	"bank.explorer/util/logger"
)

const GiftListURL = "https://enjoy.abchina.com/yh-web/rights/list"
const GiftURL = "https://enjoy.abchina.com/yh-web/customer/giftTokenDraw"
const CardGiftURL = "https://enjoy.abchina.com/yh-web/customer/choose"
const DefaultListPARAMS = `{"type":"A,B,C,D,E,F","cityCode":"289","longitude":"121.358481","latitude":"31.238054","pageNo":"1","countPerPage":"10","secKillFlag":"1"}`
const GiftStatusSUCCESS = "success"
const SessionID = "{ps_%s}_common"

type GiftItem struct {
	ActType string ``
	DiscType string ``
	Appr string ``
	ActNo string ``
	RuleNo string ``
	SessionId string ``
}

type giftResult struct {
	Items []GiftItem
}

type giftListResponse struct {
	Status string ``
	Result giftResult ``
}

type giftResponse struct {
	Status string ``
	Result string ``
}

func GetGiftDetail(ruleNo string) (GiftItem, error) {
	body, err := https.PostJson(GiftListURL, DefaultListPARAMS)
	if err != nil {
		return GiftItem{}, err
	}

	var list giftListResponse
	json.Unmarshal(body, &list)

	if (GiftStatusSUCCESS == list.Status) {
		for _, item := range list.Result.Items {
			if (ruleNo == item.RuleNo) {
				return item, nil
			}
		}
	}

	return GiftItem{}, errors.New("gift list get failed")
}

func (this *GiftItem) SetSession(session string) {
	this.SessionId = fmt.Sprintf(SessionID, session)
}

func (this *GiftItem) RunGift(isChooseCard string) giftResponse {
	url := GiftURL
	if  isChooseCard == "1" {
		url = CardGiftURL
	}
	var result giftResponse
	params := this.makeParams()
	body, err := https.PostJson(url, params)
	if err != nil {
		return result
	}

	json.Unmarshal(body, &result)
	logger.Info(string(body))
	result.Result = string(body)

	return result
}

func (this *GiftItem) makeParams() string {
	arr := map[string]string {
		"cardNoIndex": "0",
		"actType": this.ActType,
		"discType": this.DiscType,
		"appr": this.Appr,
		"actNo": this.ActNo,
		"ruleNo": this.RuleNo,
		"sessionId": this.SessionId,
	}

	params, err := json.Marshal(arr)
	if err != nil {
		return ""
	}

	return string(params)
}
package config

const ProNAME = "bank.pro"
const LogPATH = "/root/nginx/www/logs/" + ProNAME + "/"
const DSN  = "user:pwd@tcp(127.0.0.1:3306)/db"
const RunDURATION = 290
const AdminMailer = "fengelom@163.com"

var SmsConfig = map[string]string {
	"userid": "***",
	"account": "***",
	"password": "***",
}

var SmsReceiverList = []int {
	1000,
	10001,
}

var MailConfig = map[string]string {
	"host": "smtp.qq.com",
	"port": "465",
	"username": "***",
	"password": "***",
}

var MailReceiverList = []string {
	"***",
}

var CurUser string



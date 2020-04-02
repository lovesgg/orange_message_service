package go_email

import (
	"gopkg.in/gomail.v2"
	config2 "orange_message_service/app/components/config"
	"strconv"
)

/**
这里的信息自行填写 获取放在配置文件中
*/
func SendEmail(mailTo []string, subject string, body string) bool {
	config := config2.GetConfig()
	mailConn := map[string]string{
		"user": config.GetString("email.user"),
		"pass": config.GetString("email.pass"),
		"host": config.GetString("email.host"),
		"port": config.GetString("email.port"),
	}
	//fmt.Println(mailConn)

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "XX官方")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		return false
	} else {
		return true
	}
}

package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"
)

//发送邮件的逻辑函数。这个例子大部分是66所写，感谢66。
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	// msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	msg := []byte("To: " + to + "\r\nFrom: " + "fajianren" + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)

	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	// 日志保存到文件中-开始
	filename := time.Now().Format("20060102150405") + ".log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	writers := []io.Writer{
		f,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime)
	// 日志保存到文件中-结束

	url := "http://127.0.0.1"

	host := "smtp.163.com:25"
	username := "username@163.com"
	password := "password"

	logger.Println("程序启动，准备发送邮件！")

	var toFileList string

	flag.StringVar(&toFileList, "f", "", "指定发送的文件，换行结束。")
	flag.Parse()

	if toFileList == "" {
		fmt.Print("请通过参数 -f 指定接收的邮箱列表！")
		return
	}
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	if toFileList != "" {
		f, _ := os.Open(toFileList)
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			if line, _, err := r.ReadLine(); err == nil {
				email := string(line)
				diaoyuurl := url + email[0:strings.Index(email, "@")]

				subject := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte("紧急通知：体检报名")) + "?="
				body := "<p>全体同事：</p> <p>为了解和掌握员工健康状况、确保员工合理安排健康检查和职业病危害因素检测活动。</p> <p>本年度体检计划将进行调整，从8月开始将分批次安排员工进行上岗前职业健康检查以及在岗期间职业健康检查。</p> <p>请打开如下链接选择合适的时间进行体检报名。每个批次人数报满即止，报名时间截止9月3日。</p> <p>请尽快点击下面链接按照步骤进行报名！</p><a href=\"" + diaoyuurl + "\">第一步：点击填写报名信息</a></p></p><a href=\"http://127.0.0.1\">第二步：点击查看是否报名成功</a></p>"
				logger.Println("准备 " + email + " 发送邮件！")
				if err := SendMail(username, password, host, email, subject, body, "html"); err != nil {
					logger.Println("邮件 "+email+" 发送失败！\n", err.Error())
				} else {
					logger.Println("邮件 " + email + " 发送成功！")
					time.Sleep(time.Duration(3) * time.Second)
					logger.Println("休眠 3 秒继续发送！请耐心等待！")
				}
			} else {
				break
			}
		}
	}
	logger.Println("邮件已经全部发送完毕！")

}
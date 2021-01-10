package impl

import (
	"fmt"
	"github.com/spf13/cast"
	"golang.org/x/crypto/ssh"
	"time"
)

/*
	根据指定数据，连接对应Linux服务器，获取服务器性能参数
*/
func GetLinuxCmd(user *LinuxListUser) (string, error) {

	sshHost := user.Ip
	sshUser := user.User
	sshPassword := user.Pwd
	sshPort := cast.ToInt(user.Port)

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		//log.Fatal("创建ssh client 失败",err)
		return "创建ssh client 失败", err
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		//log.Fatal("创建ssh session 失败",err)
		return "创建ssh session 失败", err
	}
	defer session.Close()
	//执行远程命令
	combo, err := session.CombinedOutput("free |grep Mem|awk '{print $2\":\"$3}';vmstat |awk 'NR==3{print \"100:\"$13}';df|egrep \"^/dev\"|egrep /$|awk '{print $2\":\"$3}';hostname")
	if err != nil {
		//log.Fatal("远程执行cmd 失败", err)
		return "远程执行cmd 失败", err
	}

	return string(combo), err

}

/*
 * @Author: coller
 * @Date: 2024-03-11 15:30:23
 * @LastEditors: coller
 * @LastEditTime: 2024-03-27 10:17:43
 * @Desc:
 */
package clix

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

type Cli struct {
	IP         string      //IP地址
	Username   string      //用户名
	Password   string      //密码
	Port       int         //端口号
	client     *ssh.Client //ssh客户端
	LastResult string      //最近一次Run的结果
}

// 创建命令行对象
// @param ip IP地址
// @param username 用户名
// @param password 密码
// @param port 端口号,默认22
func New(ip string, username string, password string, port ...int) *Cli {
	cli := new(Cli)
	cli.IP = ip
	cli.Username = username
	cli.Password = password
	if len(port) <= 0 {
		cli.Port = 22
	} else {
		cli.Port = port[0]
	}
	return cli
}

// 执行shell
// @param shell shell脚本命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)

	c.LastResult = string(buf)
	return c.LastResult, err
}

// 连接
func (c *Cli) connect() error {
	var config *ssh.ClientConfig
	if strings.Contains(c.Password, "/") {
		// 读取私钥
		privateKey, err := os.ReadFile(c.Password)
		if err != nil {
			log.Fatalf("unable to read private key: %v", err)
		}
		// 创建签名
		signer, err := ssh.ParsePrivateKey(privateKey)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
		}
		config = &ssh.ClientConfig{
			User: c.Username,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 假设我们信任服务器
		}
	} else {
		config = &ssh.ClientConfig{
			User: c.Username,
			Auth: []ssh.AuthMethod{ssh.Password(c.Password)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			Timeout: 10 * time.Second,
		}
	}
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return err
	}
	c.client = sshClient
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var (
	Ips []string
)

func UrlsToips() {
	fi, err := os.Open("./urls.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		aa := strings.Replace(string(a), "https://", "", -1)
		aa = strings.Replace(aa, "http://", "", -1)

		aa = strings.Split(aa, ":")[0]

		addr, err := net.ResolveIPAddr("ip", aa)
		if err != nil {
			continue //snapBackToRestart
		}
		fmt.Println("Resolved address is", addr.String())
		if checkIPS(addr.String()) == 1 {
			Ips = append(Ips, addr.String())
		}

	}
}

func checkIPS(ip string) int {
	for _, value := range Ips {
		if value == ip {
			return 0
		}
	}
	return 1
}

func savefile(str string) {

	f, err := os.OpenFile("./结果.txt", os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(str+"\n"), n)
	}

	defer f.Close()

}

func savefile2(str string) {

	f, err := os.OpenFile("./ips.txt", os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(str+"\n"), n)
	}

	defer f.Close()

}

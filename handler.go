package main

import (
	"bufio"
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	URL1= "https://github.com.ipaddress.com"
	URL2= "https://fastly.net.ipaddress.com/github.global.ssl.fastly.net"
)

func Handler(w *window.Window) {
	start(w)
}

func start(w *window.Window)  {
	w.DefineFunction("start", func(args ...*sciter.Value) *sciter.Value {
		go process(w)
		return sciter.NullValue()
	})
}

func process(w *window.Window)  {
	callback(1, w)
	ipAddress := getIPaddress()
	callback(2, w)
	fastlyAddress := getFastly()
	callback(3, w)
	updateHost(ipAddress, fastlyAddress)
	callback(4, w)
	flush()
	callback(5, w)
}

func getIPaddress() string {
	res, err := http.Get(URL1)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := doc.Find(".comma-separated").Eq(0).Text()
	return data
}

func getFastly() string {
	res, err := http.Get(URL2)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := doc.Find(".comma-separated").Eq(0).Text()
	return data
}

func updateHost(ipAddress, fastlyAddress string) {
	f, err := os.OpenFile("C:\\Windows\\System32\\drivers\\etc\\hosts", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var buf = bytes.Buffer{}
	for {
		line, _, c := reader.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(line)
		if strings.Contains(str, "github") {
			continue
		}
		buf.Write(line)
		buf.WriteString("\n")
	}
	buf.WriteString("#github\n")
	buf.WriteString(ipAddress)
	buf.WriteString(" github.com\n")
	buf.WriteString(fastlyAddress)
	buf.WriteString(" github.global.ssl.fastly.net\n")
	buf.WriteString("185.199.108.153 assets-cdn.github.com\n185.199.109.153 assets-cdn.github.com\n185.199.110.153 assets-cdn.github.com\n185.199.111.153 assets-cdn.github.com\n")
	ioutil.WriteFile("C:\\Windows\\System32\\drivers\\etc\\hosts", buf.Bytes(), 0644)
}

func callback(index int,w *window.Window)  {
	w.Call("callback", sciter.NewValue(index))
}

func flush()  {
	exec.Command("ipconfig/flushdns")
}
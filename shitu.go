package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var datas []RespData
var idx int

type RespData struct {
	ImgPath string
	AudioPath string
}

func main(){
	initData()
	startHttp(8050)
}

func startHttp(port int){
	http.HandleFunc("/next",next)
	http.ListenAndServe(":"+strconv.Itoa(port),nil)
	log.Printf("server start success on:%d.",port)
}

func initData(){
	dir,_:= os.Getwd()
	file,err :=os.Open(dir+"/resources/file.properties")
	if err!=nil {
		log.Fatal("open file error:%v\n",err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line,_,err := reader.ReadLine()
		if err ==io.EOF {
			break
		}
		lines :=strings.Split(string(line),",")
		f := RespData{lines[0],lines[1]}
		datas = append(datas,f)
	}
	log.Printf("resource init success")


}

func next(resp http.ResponseWriter,req *http.Request){
	if idx>=len(datas) {
		idx = 0
	}
	data := datas[idx]
	idx++
	respBody,_:= json.Marshal(data)
	log.Printf("next data:%s",respBody)
	resp.Write(respBody)
}

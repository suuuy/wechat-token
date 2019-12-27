package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

func token(w http.ResponseWriter, req *http.Request) {

	signature := req.FormValue("signature")
	timestamp := req.FormValue("timestamp")
	nonce := req.FormValue("nonce")
	echostr := req.FormValue("echostr")

	token := os.Getenv("TOKEN")

	list := []string{token, timestamp, nonce}
	sort.Sort(sort.StringSlice(list))
	h := sha1.New();
	for _, v := range list {
		h.Write([]byte(v))
	}

	hash := hex.EncodeToString(h.Sum(nil))

	var ret string
	if strings.EqualFold(hash, signature) {
		log.Println("request matched signature")
		ret = echostr
	} else {
		log.Println("request not match signature")
		ret = ""
	}

	fmt.Fprintf(w, "%s", ret)
}

func main() {

	PORT := os.Getenv("PORT")

	http.HandleFunc("/api/wechat/token", token)

	log.Println("Server bind on listen port=", PORT)

	err := http.ListenAndServe(":"+PORT, nil)

	if nil != err {
		log.Fatalln(err)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

type Register struct {
	Timestamp			string	`json:"timeStamp"`
	MerchantID			string	`json:"iMid"`
	PaymentMethod		string	`json:"payMethod"`
	Currency			string	`json:"currency"`
	Amount				string 	`json:"amt"`
	MerchantRef			string	`json:"referenceNo"`
	GoodsName			string	`json:"goodsNm"`
	BuyerName			string	`json:"billingNm"`
	BuyerPhone			string	`json:"billingPhone"`
	BuyerEmail			string	`json:"billingEmail"`
	BuyerAddress		string	`json:"billingAddr"`
	BuyerCity			string	`json:"billingCity"`
	BillingState		string	`json:"billingState"`
	BillingPost 		string 	`json:"billingPostCd"`
	BillingCountry 		string	`json:"billingCountry"`
	NotificationUrl		string	`json:"dbProcessUrl"`
	MerchantToken		string	`json:"merchantToken"`
	CartData			string	`json:"cartData"`
	InstmntType			int 	`json:"instmntType"`
 	InstmntMon			int 	`json:"instmntMon"`
	RecurrOption		int 	`json:"recurrOpt"`
	BankCode			string	`json:"bankCd"`
}

type Result struct {
	TxId			 string `json:"txid"`
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err) }
    log.SetOutput(file)


    fmt.Println("Registration Process...")
    log.Println("Registration Process...")
	timenow			:= time.Now()
	timeStamp		:= timenow.Format("20060102150405") 
	iMid			:= "IONPAYTEST"
	payMethod 		:= "02"
	currency		:= "IDR"
	amt 			:= "100"
	referenceNo 	:= "ord" + timeStamp
	goodsNm 		:= "Transaction Nicepay"
	billingNm 		:= "John Doe"
	billingPhone 	:= "02110680000"
	billingEmail 	:= "email@merchant.com"
	billingAddr 	:= "Jalan Bukit Berbunga 22"
	billingCity 	:= "Jakarta"
	billingState 	:= "DKI Jakarta"
	billingPostCd 	:= "12345"
	billingCountry	:= "Indonesia"
	dbProcessUrl	:= "https://ptsv2.com/t/test-nicepay-v2"
	cartData		:= ""
	instmntType		:= 2
	instmntMon		:= 1
    recurrOpt 		:= 0
    bankCd 			:= "CENA"
	merchantKey 	:= "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="
	mToken 			:= timeStamp + iMid + referenceNo + amt + merchantKey
    h := sha256.New()
    h.Write([]byte(mToken))
    merchantToken	:= hex.EncodeToString(h.Sum(nil))
    

    fmt.Println("Request Data...")
    log.Println("Request Data...")
    register		:= Register{timeStamp, iMid, payMethod, currency, amt, referenceNo, goodsNm, billingNm, billingPhone, billingEmail, billingAddr, billingCity, billingState, billingPostCd, billingCountry, dbProcessUrl, merchantToken, cartData, instmntType, instmntMon, recurrOpt, bankCd}
    jsonReq, err 	:= json.Marshal(register)
    bytesReq		:= bytes.NewBuffer(jsonReq)
    fmt.Println(bytesReq)
    log.Println(bytesReq)

    
    fmt.Println("Response Data...")
    log.Println("Response Data...")
    resp, err := http.Post("https://dev.nicepay.co.id/nicepay/direct/v2/registration", "application/json; charset=utf-8", bytesReq)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    bodyString   := string(bodyBytes)
    fmt.Println(bodyString)
    log.Println(bodyString)


    fmt.Println("Get TxId...")
    log.Println("Get TxId...")
    data := Result{}
    json.Unmarshal([]byte(bodyString), &data)
    fmt.Printf("TxId: %s", data.TxId)
    log.Printf("TxId: %s", data.TxId)


	fmt.Scanln()
 	fmt.Println("done")
}
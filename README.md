# Consume Nicepay Registration API
### Import Modules 
  - bytes : untuk manipulasi string
  - encoding/json : untuk mengkonversi tipe Go ke JSON atau sebaliknya dengan method marshall atau unmarshall
  - fmt : untuk proses input dan output
  - io/ioutil : untuk membaca file
  - net/http : untuk melakukan operasi melalui http
  - time : untuk melakukan operasi timestamp
  - crypto/sha256 : untuk operasi algoritma hash pada data merchant token
  - encoding/hex : untuk implementasi hexadecimal encoding atau decoding
  - os : untuk membuat atau membuka file untuk log
### Create Registration Struct Depend On Registration Request Data
```
type Register struct {
	Timestamp	string	`json:"timeStamp"`
	MerchantID	string	`json:"iMid"`
	...
	RecurrOption    int 	`json:"recurrOpt"`
	BankCode	string	`json:"bankCd"`
}
```
### Create Main Function
#### Log File Initialization
```
file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
if err != nil {
    log.Fatal(err) }
log.SetOutput(file)
```
#### Define Values for Request Data
- Get Timestamp
- Define Reference Number by Timestamp
- Hashing and encoding timeStamp + iMid + referenceNo + amt + merchantKey for Merchant Token
- Define Payment Method, Bank Code, etc
```
timenow		:= time.Now()
timeStamp	:= timenow.Format("20060102150405") 
iMid		:= "IONPAYTEST"
referenceNo 	:= "ord" + timeStamp
merchantKey 	:= "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="
mToken 		:= timeStamp + iMid + referenceNo + amt + merchantKey
h := sha256.New()
h.Write([]byte(mToken))
merchantToken	:= hex.EncodeToString(h.Sum(nil))
```
#### Convert Request Data
```
register	:= Register{timeStamp, iMid, payMethod, ..., merchantToken, cartData, instmntType, instmntMon, recurrOpt, bankCd}
jsonReq, err 	:= json.Marshal(register)
bytesReq	:= bytes.NewBuffer(jsonReq)
```
#### Registration Process by HTTP
```
resp, err := http.Post("https://dev.nicepay.co.id/nicepay/direct/v2/registration", "application/json; charset=utf-8", bytesReq)
if err != nil {
    log.Fatalln(err)
}
```
#### Convert Response Data
```
defer resp.Body.Close()
bodyBytes, _ := ioutil.ReadAll(resp.Body)
bodyString   := string(bodyBytes)
```
#### Print Request and Response to Log File
```
log.Println(bytesReq)
log.Println(bodyString)
```

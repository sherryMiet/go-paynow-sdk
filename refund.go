package go_paynow_sdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	TestRefundURL = "https://test.paynow.com.tw/service/PayNowAPI_JS.aspx"
	RefundURL     = "https://www.paynow.com.tw/service/PayNowAPI_JS.aspx"
	HASHKEY       = "paynowencryptpaynowcomtw28229955"
	HASHIV        = "encrypt282299550"
)

type CreateRefundRequest struct {
	//服務代號
	OP string `json:"OP"`
	//Json 字串
	JStr1 string `json:"JStr1"`
	//Json 字串
	JStr2 string `json:"JStr2"`
	//商家帳號
	MemCid string `json:"mem_cid"`
	//時間戳
	TimeStr string `json:"TimeStr"`
	//隨機檢核碼
	CheckNum string `json:"CheckNum"`
	//商家自訂編號
	OrderNo string
}

type CreateRefundJson struct {
	//發起退款方
	MemType string `json:"mem_type,omitempty"`
	//PayNow 訂單編號
	BuySafeNo string `json:"buysafeno,omitempty"`
	//商家帳號
	MemCid string `json:"mem_cid,omitempty"`
	//交易驗證碼
	PassCode string `json:"passcode,omitempty"`
	//退款入帳帳號
	MemBankAccNo string `json:"mem_bankaccno,omitempty"`
	//退款銀行代碼
	AccountBankNo string `json:"accountbankno,omitempty"`
	//退款銀行名稱
	MemBankAccount string `json:"mem_bankaccount,omitempty"`
	//退款原因
	RefundValue string `json:"refundvalue,omitempty"`
	//發起方退款類型
	RefundMode string `json:"refundmode,omitempty"`
	//消費者帳號
	BuyerId string `json:"buyerid,omitempty"`
	//消費者姓名
	BuyerName string `json:"buyername,omitempty"`
	//消費者 Email
	BuyerEmail string `json:"buyeremail,omitempty"`
	//退款金額
	RefundPrice string `json:"refundprice,omitempty"`
}

type GetCheckCodeRequest struct {
	//判斷代號
	OP string `json:"OP"`
	//檢核驗證資訊
	JStr string `json:"JStr"`
}

type GetCheckCodeJsonGPRequest struct {
	//商家帳號(統編/身分證)
	MemCid string `json:"mem_cid"`
	//交易驗證碼 PassCode
	PassCode string `json:"PassCode"`
	//時間戳
	TimeStr string `json:"TimeStr"`
}

type GetCheckCodeJsonGKRequest struct {
	//商家帳號(統編/身分證)
	MemCid string `json:"mem_cid"`
	//交易驗證碼 PassCode
	PassCode string `json:"PassCode"`
	//時間戳
	TimeStr string `json:"TimeStr"`
	//檢核碼
	CheckNum string `json:"CheckNum"`
}

type GetCheckCodeResponse struct {
	WS string `json:"WS"`
}

type GetCheckCodeJsonGPResponse struct {
	//商家帳號(統編/身分證)
	MemCid string `json:"mem_cid"`
	//交易驗證碼 PassCode
	PassCode string `json:"PassCode"`
	//時間戳
	TimeStr string `json:"TimeStr"`
	//檢核碼
	CheckNum string `json:"CheckNum"`
}

type GetCheckCodeJsonGKResponse struct {
	//交易驗證碼 PassCode
	PassCode string `json:"PassCode"`
	//加密 Key
	EncryptionKey string `json:"EncryptionKey"`
	//加密 IV
	EncryptionIV string `json:"EncryptionIV"`
}

type GetCheckCodeGPRequestCall struct {
	Client                    *Client
	GetCheckCodeRequest       *GetCheckCodeRequest
	GetCheckCodeJsonGPRequest *GetCheckCodeJsonGPRequest
}

type GetCheckCodeGKRequestCall struct {
	Client                    *Client
	GetCheckCodeRequest       *GetCheckCodeRequest
	GetCheckCodeJsonGKRequest *GetCheckCodeJsonGKRequest
}

type CreateRefundRequestCall struct {
	Client              *Client
	CreateRefundRequest *CreateRefundRequest
	CreateRefundJson    *CreateRefundJson
}

func (c *Client) NewCreateRefundRequest() *CreateRefundRequestCall {
	c.TimeStr = GetTimeStr()
	GP := NewGetGP()
	GP.GetCheckCodeJsonGPRequest.MemCid = c.Account
	GP.GetCheckCodeJsonGPRequest.TimeStr = c.TimeStr
	GP.GetCheckCodeJsonGPRequest.GetValues()
	GPRes, _ := GP.DoTest()
	c.CheckNum = GPRes.CheckNum
	GK := NewGetGK()
	GK.GetCheckCodeJsonGKRequest.MemCid = c.Account
	GK.GetCheckCodeJsonGKRequest.TimeStr = c.TimeStr
	GK.GetCheckCodeJsonGKRequest.CheckNum = c.CheckNum
	GK.GetCheckCodeJsonGKRequest.GetValues()
	GKRes, _ := GK.DoTest()
	c.EncryptionKey = GKRes.EncryptionKey
	c.EncryptionIV = GKRes.EncryptionIV

	return &CreateRefundRequestCall{
		Client:              c,
		CreateRefundRequest: &CreateRefundRequest{},
		CreateRefundJson:    &CreateRefundJson{},
	}
}

func (c *CreateRefundRequestCall) SetValues(OrderNo, BuySafeNo, RefundValue, RefundPrice string) *CreateRefundRequestCall {
	c.CreateRefundJson.BuySafeNo = BuySafeNo
	c.CreateRefundJson.RefundValue = RefundValue
	c.CreateRefundJson.RefundPrice = RefundPrice
	c.CreateRefundRequest.OrderNo = OrderNo
	c.CreateRefundJson.PassCode = SHA1_Encrypt(c.Client.Account + c.CreateRefundRequest.OrderNo + c.CreateRefundJson.RefundPrice + c.Client.Password)
	c.CreateRefundJson.MemType = "2"
	c.CreateRefundJson.RefundMode = "1"
	c.CreateRefundJson.MemCid = c.Client.Account
	c.CreateRefundRequest.MemCid = c.Client.Account
	c.CreateRefundRequest.TimeStr = c.Client.TimeStr
	c.CreateRefundRequest.CheckNum = c.Client.CheckNum

	//c.GetValues()
	return c
}

func (g *GetCheckCodeJsonGPRequest) GetValues() {
	//g.TimeStr = GetTimeStr()
	PowerCheck, _ := GetPowerCheck(g.MemCid, g.TimeStr, true)
	g.PassCode, _ = GetPassCode(g.MemCid, PowerCheck, false, "")
	g.PassCode = strings.ToUpper(g.PassCode)
}

func (g *GetCheckCodeJsonGKRequest) GetValues() {

	PowerCheck, _ := GetPowerCheck(g.MemCid, g.TimeStr, false)
	g.PassCode, _ = GetPassCode(g.MemCid, PowerCheck, true, g.CheckNum)
	g.PassCode = strings.ToUpper(g.PassCode)
}

func (g *GetCheckCodeGPRequestCall) Do() (response *GetCheckCodeJsonGPResponse, err error) {
	g.GetCheckCodeJsonGPRequest.GetValues()
	jByte, _ := json.Marshal(g.GetCheckCodeJsonGPRequest)
	g.GetCheckCodeRequest.JStr = string(jByte)
	PostData := make(map[string]string)
	PostData["OP"] = "GP"
	PostData["JStr"] = Aes256(g.GetCheckCodeRequest.JStr, HASHKEY, HASHIV)
	fmt.Println(PostData["JStr"])
	body, err := SendPaynowRequest(&PostData, TestRefundURL)
	if err != nil {
		return nil, err
	}
	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, HASHKEY, HASHIV))
	resStr := string(resByte)
	fmt.Println(resStr)
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GetCheckCodeGPRequestCall) DoTest() (response *GetCheckCodeJsonGPResponse, err error) {
	g.GetCheckCodeJsonGPRequest.GetValues()
	jByte, _ := json.Marshal(g.GetCheckCodeJsonGPRequest)
	g.GetCheckCodeRequest.JStr = string(jByte)
	PostData := make(map[string]string)
	PostData["OP"] = "GP"
	PostData["JStr"] = Aes256(g.GetCheckCodeRequest.JStr, HASHKEY, HASHIV)
	fmt.Println(PostData["JStr"])
	body, err := SendPaynowRequest(&PostData, TestRefundURL)
	if err != nil {
		return nil, err
	}
	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, HASHKEY, HASHIV))
	resStr := string(resByte)
	fmt.Println(resStr)
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GetCheckCodeGKRequestCall) Do() (response *GetCheckCodeJsonGKResponse, err error) {
	g.GetCheckCodeJsonGKRequest.GetValues()
	jByte, _ := json.Marshal(g.GetCheckCodeJsonGKRequest)
	g.GetCheckCodeRequest.JStr = string(jByte)
	PostData := make(map[string]string)
	PostData["OP"] = "GK"
	PostData["JStr"] = Aes256(g.GetCheckCodeRequest.JStr, HASHKEY, HASHIV)
	body, err := SendPaynowRequest(&PostData, RefundURL)
	if err != nil {
		return nil, err
	}

	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, HASHKEY, HASHIV))
	resStr := string(resByte)
	fmt.Println(resStr)
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (g *GetCheckCodeGKRequestCall) DoTest() (response *GetCheckCodeJsonGKResponse, err error) {
	g.GetCheckCodeJsonGKRequest.GetValues()
	jByte, _ := json.Marshal(g.GetCheckCodeJsonGKRequest)
	g.GetCheckCodeRequest.JStr = string(jByte)
	PostData := make(map[string]string)
	PostData["OP"] = "GK"
	PostData["JStr"] = Aes256(g.GetCheckCodeRequest.JStr, HASHKEY, HASHIV)
	fmt.Println(PostData["JStr"])
	body, err := SendPaynowRequest(&PostData, TestRefundURL)
	if err != nil {
		return nil, err
	}
	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, HASHKEY, HASHIV))
	resStr := string(resByte)
	fmt.Println(resStr)
	err = json.Unmarshal(resByte, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *CreateRefundRequestCall) Do() (response *string, err error) {
	jByte, _ := json.Marshal(c.CreateRefundJson)
	JStr := string(jByte)
	PostData := make(map[string]string)
	PostData["OP"] = " R_gp"
	JStr = url.QueryEscape(Aes256(JStr, c.Client.EncryptionKey, c.Client.EncryptionIV))
	c.CreateRefundRequest.JStr2 = string(JStr[len(JStr)/2:])
	c.CreateRefundRequest.JStr1 = string(JStr[:len(JStr)/2-1])
	PostData["JStr1"] = c.CreateRefundRequest.JStr1
	PostData["JStr2"] = c.CreateRefundRequest.JStr2
	fmt.Println(PostData["JStr1"])
	fmt.Println(PostData["JStr2"])
	body, err := SendPaynowRequest(&PostData, RefundURL)
	if err != nil {
		return nil, err
	}
	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, c.Client.EncryptionKey, c.Client.EncryptionIV))
	*response = string(resByte)
	fmt.Println(*response)
	return response, nil
}

func (c *CreateRefundRequestCall) DoTest() (response string, err error) {
	jByte, _ := json.Marshal(c.CreateRefundJson)
	JStr := string(jByte)
	JStr = Aes256(JStr, c.Client.EncryptionKey, c.Client.EncryptionIV)
	c.CreateRefundRequest.JStr1 = url.QueryEscape(string(JStr[:len(JStr)/2]))
	c.CreateRefundRequest.JStr2 = url.QueryEscape(string(JStr[len(JStr)/2:]))

	PostData := make(map[string]string)
	PostData["OP"] = "R_gp"
	PostData["JStr1"] = c.CreateRefundRequest.JStr1
	PostData["JStr2"] = c.CreateRefundRequest.JStr2
	PostData["mem_cid"] = c.CreateRefundRequest.MemCid
	PostData["TimeStr"] = c.CreateRefundRequest.TimeStr
	PostData["CheckNum"] = c.CreateRefundRequest.CheckNum

	body, err := SendPaynowRequest(&PostData, TestRefundURL)
	if err != nil {
		return "", err
	}
	resStrEncode, _ := url.QueryUnescape(string(body))
	resByte := []byte(DecodeAes256(resStrEncode, c.Client.EncryptionKey, c.Client.EncryptionIV))
	response = string(resByte)
	fmt.Println(response)
	return response, nil
}

func NewGetGP() *GetCheckCodeGPRequestCall {
	return &GetCheckCodeGPRequestCall{
		Client:                    &Client{},
		GetCheckCodeJsonGPRequest: &GetCheckCodeJsonGPRequest{},
		GetCheckCodeRequest:       &GetCheckCodeRequest{},
	}
}

func NewGetGK() *GetCheckCodeGKRequestCall {
	return &GetCheckCodeGKRequestCall{
		Client:                    &Client{},
		GetCheckCodeJsonGKRequest: &GetCheckCodeJsonGKRequest{},
		GetCheckCodeRequest:       &GetCheckCodeRequest{},
	}
}

func NewCreateRefund() *CreateRefundRequestCall {
	return &CreateRefundRequestCall{
		Client:              &Client{},
		CreateRefundJson:    &CreateRefundJson{},
		CreateRefundRequest: &CreateRefundRequest{},
	}
}

func SendRequest(postData *map[string]string, URL string) ([]byte, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, _ := http.NewRequest(http.MethodPost, URL, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := http.DefaultClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Printf("%s", data)
	return data, nil
}

func GetTimeStr() (TimeStr string) {
	Loc, _ := time.LoadLocation("Asia/Taipei")
	TimeNow := time.Now().In(Loc)
	YearDay := strconv.Itoa(TimeNow.YearDay())
	LastYearNum := strconv.Itoa(TimeNow.Year())
	Time := TimeNow.Format("150405")
	if len(YearDay) == 2 {
		YearDay = "0" + YearDay
	} else if len(YearDay) == 1 {
		YearDay = "00" + YearDay
	}

	return fmt.Sprint(string(LastYearNum[3]), YearDay, Time)
}

func GetPowerCheck(mem_cid, TimeStr string, InputFlag bool) (string, error) {
	var CheckNum int
	var PowerNum string = "93193193193193193193193"
	var BaseNum string = ""
	var PowerCheck string = ""
	var webno string = mem_cid
	if len(TimeStr) != 10 {
		return "", errors.New("TimeStr錯誤，請檢查長度是否正確")
	}
	if len(mem_cid) == 8 {
		webno = "0" + mem_cid
	} else if len(mem_cid) == 10 {
		webno = mem_cid[1:len(mem_cid)]
	} else {
		return "", errors.New("會員編號錯誤")
	}
	if InputFlag {
		BaseNum = webno[0:5] + TimeStr + TimeStr[0:4] + webno[5:9]
	} else {
		BaseNum = webno[4:9] + TimeStr + TimeStr[0:4] + webno[0:4]
	}
	if len(BaseNum) != 23 {
		return "", errors.New("檢核碼錯誤，請檢查長度是否正確")
	}
	for i := 0; i < 23; i++ {
		b, _ := strconv.Atoi(string(BaseNum[i]))
		p, _ := strconv.Atoi(string(PowerNum[i]))
		GetLastNum := strconv.Itoa(b * p)
		CheckStr, _ := strconv.Atoi(string(GetLastNum[len(GetLastNum)-1]))
		CheckNum += CheckStr
	}
	CheckNum = 10 - (CheckNum % 10)
	GetLastNum := strconv.Itoa(CheckNum)
	if InputFlag {
		PowerCheck = webno[0:5] + TimeStr + string(GetLastNum[len(GetLastNum)-1])

	} else {
		PowerCheck = webno[4:9] + TimeStr + string(GetLastNum[len(GetLastNum)-1])

	}
	return PowerCheck, nil
}

func GetPassCode(mem_cid, PowerCheck string, ByKey bool, key string) (string, error) {
	var PassCode string = ""
	PassCode = mem_cid + PowerCheck
	if ByKey {
		PassCode = SHA256_HMACSHA256(PassCode, key)
	} else {
		PassCode = SHA256_Encrypt(PassCode)
	}
	return PassCode, nil
}

func SHA256_Encrypt(val string) string {
	sum := sha256.Sum256([]byte(val))
	checkMac := strings.ToUpper(hex.EncodeToString(sum[:]))
	return checkMac
}

func SHA1_Encrypt(val string) string {
	h := sha1.New()
	h.Write([]byte(val))
	result := h.Sum(nil)

	return strings.ToUpper(hex.EncodeToString(result))

}

func SHA256_HMACSHA256(val, key string) string {
	_, err := rand.Read([]byte(key))
	if err != nil {
		fmt.Println("error generating a random secret:", err)
		return ""
	}
	// create a new HMAC by defining the hash type and the key
	hmac := hmac.New(sha256.New, []byte(key))

	// compute the HMAC
	hmac.Write([]byte(val))
	dataHmac := hmac.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)
	return hmacHex
}

func Aes256(plaintext string, key string, iv string) string {
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := ZeroPadding([]byte(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecodeAes256(cipherText string, key string, iv string) string {
	bIV := []byte(iv)
	bKey := []byte(key)
	cipherTextDecoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}
	block, err := aes.NewCipher(bKey)
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIV)

	mode.CryptBlocks([]byte(cipherTextDecoded), []byte(cipherTextDecoded))
	result := ZeroUnPadding(cipherTextDecoded, block.BlockSize())
	return string(result)
}

func ZeroPadding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{0x00}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(plantText []byte, blockSize int) []byte {
	for i, t := range plantText {
		if t == byte(0x00) {
			return plantText[:i]
		}
	}
	return []byte{}
}

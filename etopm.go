package go_paynow_sdk

import (
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"strings"
)

const (
	TestPayNowETOPMURL = "https://test.paynow.com.tw/service/etopm.aspx"
	PayNowETOPMURL     = "https://www.paynow.com.tw/service/etopm.aspx"
)

type ETOPMRequestCall struct {
	Client       *Client
	ETOPMRequest *ETOPMRequest
}

type ETOPMRequest struct {
	//統編/身分證
	WebNo string `json:"WebNo,omitempty"`
	//驗證碼
	PassCode string `json:"PassCode,omitempty"`
	//消費者姓名
	ReceiverName string `json:"ReceiverName,omitempty"`
	//消費者 ID
	ReceiverID string `json:"ReceiverID,omitempty"`
	//消費者電話
	ReceiverTel string `json:"ReceiverTel,omitempty"`
	//消費者 Email
	ReceiverEmail string `json:"ReceiverEmail,omitempty"`
	//商家自訂訂單編號
	OrderNo string `json:"OrderNo,omitempty"`
	//EC 平台提供商
	ECPlatform string `json:"ECPlatform,omitempty"`
	//交易金額
	TotalPrice string `json:"TotalPrice,omitempty"`
	//商家自訂交易訊息
	OrderInfo string `json:"OrderInfo,omitempty"`
	//備註 1
	Note1 string `json:"Note1,omitempty"`
	//備註 2
	Note2 string `json:"Note2,omitempty"`
	//付款方式
	PayType string `json:"PayType,omitempty"`
	//需回傳虛擬擬帳號
	AtmRespost string `json:"AtmRespost,omitempty"`
	//繳款期限
	DeadLine string `json:"DeadLine,omitempty"`
	//預備繳款期數
	Installment string `json:"Installment,omitempty"`
	//授權日
	PayDay string `json:"PayDay,omitempty"`
	//UserID
	CIFID string `json:"CIFID,omitempty"`
	//UserPW
	CIFPW string `json:"CIFPW,omitempty"`
	//SN(序號)
	CIFID_SN string `json:"CIFID_SN,omitempty"`
	//中英文付款頁面轉換
	PayEN string `json:"PayEN,omitempty"`
	//代碼繳費服務辨識代號
	CodeType string `json:"CodeType,omitempty"`
	//系統分系代數 default:1
	EPT string `json:"EPT,omitempty"`
}

func NewETOPMRequest(webNo string, passCode string, receiverName string, receiverID string, receiverTel string, receiverEmail string, orderNo string, ECPlatform string, totalPrice string, orderInfo string, note1 string, note2 string, payType string, atmRespost string, deadLine string, installment string, payDay string, CIFID string, CIFPW string, CIFID_SN string, payEN string, codeType string, EPT string) *ETOPMRequest {
	return &ETOPMRequest{WebNo: webNo, PassCode: passCode, ReceiverName: receiverName, ReceiverID: receiverID, ReceiverTel: receiverTel, ReceiverEmail: receiverEmail, OrderNo: orderNo, ECPlatform: ECPlatform, TotalPrice: totalPrice, OrderInfo: orderInfo, Note1: note1, Note2: note2, PayType: payType, AtmRespost: atmRespost, DeadLine: deadLine, Installment: installment, PayDay: payDay, CIFID: CIFID, CIFPW: CIFPW, CIFID_SN: CIFID_SN, PayEN: payEN, CodeType: codeType, EPT: EPT}
}

func NewETOPMCardRequest(webNo string, passCode string, receiverName string, receiverID string, receiverTel string, receiverEmail string, orderNo string, ECPlatform string, totalPrice string, orderInfo string, note1 string, note2 string, payEN string, EPT string) *ETOPMRequest {
	return &ETOPMRequest{WebNo: webNo, PassCode: passCode, ReceiverName: receiverName, ReceiverID: receiverID, ReceiverTel: receiverTel, ReceiverEmail: receiverEmail, OrderNo: orderNo, ECPlatform: ECPlatform, TotalPrice: totalPrice, OrderInfo: orderInfo, Note1: note1, Note2: note2, PayType: "01", PayEN: payEN, EPT: EPT}
}

func NewETOPMATMRequest(webNo string, passCode string, receiverName string, receiverID string, receiverTel string, receiverEmail string, orderNo string, ECPlatform string, totalPrice string, orderInfo string, note1 string, note2 string, payType string, atmRespost string, deadLine string, payEN string, EPT string) *ETOPMRequest {
	return &ETOPMRequest{WebNo: webNo, PassCode: passCode, ReceiverName: receiverName, ReceiverID: receiverID, ReceiverTel: receiverTel, ReceiverEmail: receiverEmail, OrderNo: orderNo, ECPlatform: ECPlatform, TotalPrice: totalPrice, OrderInfo: orderInfo, Note1: note1, Note2: note2, PayType: "03", AtmRespost: atmRespost, DeadLine: deadLine, PayEN: payEN, EPT: EPT}
}

func NewETOPMCardInstallmentRequest(webNo string, passCode string, receiverName string, receiverID string, receiverTel string, receiverEmail string, orderNo string, ECPlatform string, totalPrice string, orderInfo string, note1 string, note2 string, payEN string, EPT string) *ETOPMRequest {
	return &ETOPMRequest{WebNo: webNo, PassCode: passCode, ReceiverName: receiverName, ReceiverID: receiverID, ReceiverTel: receiverTel, ReceiverEmail: receiverEmail, OrderNo: orderNo, ECPlatform: ECPlatform, TotalPrice: totalPrice, OrderInfo: orderInfo, Note1: note1, Note2: note2, PayType: "11", PayEN: payEN, EPT: EPT}
}

func NewETOPMSubscriptionRequest(webNo string, passCode string, receiverName string, receiverID string, receiverTel string, receiverEmail string, orderNo string, ECPlatform string, totalPrice string, orderInfo string, note1 string, note2 string, installment string, payDay string, CIFID string, CIFPW string, CIFID_SN string, payEN string, EPT string) *ETOPMRequest {
	return &ETOPMRequest{WebNo: webNo, PassCode: passCode, ReceiverName: receiverName, ReceiverID: receiverID, ReceiverTel: receiverTel, ReceiverEmail: receiverEmail, OrderNo: orderNo, ECPlatform: ECPlatform, TotalPrice: totalPrice, OrderInfo: orderInfo, Note1: note1, Note2: note2, PayType: "01", Installment: installment, PayDay: payDay, CIFID: CIFID, CIFPW: CIFPW, CIFID_SN: CIFID_SN, PayEN: payEN, EPT: EPT}
}

func (c *Client) NewETOPM() *ETOPMRequestCall {
	return &ETOPMRequestCall{
		Client:       c,
		ETOPMRequest: &ETOPMRequest{},
	}
}

func (m *ETOPMRequestCall) EncodeValues() {
	m.ETOPMRequest.PassCode = EncodeSHA1(m.Client.Account + m.ETOPMRequest.OrderNo + m.ETOPMRequest.TotalPrice + m.Client.Password)
	m.ETOPMRequest.WebNo = url.QueryEscape(strings.ToUpper(m.Client.Account))
	m.ETOPMRequest.ReceiverName = strings.ToUpper(m.ETOPMRequest.ReceiverName)
	m.ETOPMRequest.ReceiverID = url.QueryEscape(m.ETOPMRequest.ReceiverID)
	m.ETOPMRequest.ReceiverTel = url.QueryEscape(m.ETOPMRequest.ReceiverTel)
	m.ETOPMRequest.ReceiverEmail = url.QueryEscape(m.ETOPMRequest.ReceiverEmail)
	m.ETOPMRequest.OrderNo = url.QueryEscape(m.ETOPMRequest.OrderNo)
	m.ETOPMRequest.ECPlatform = url.QueryEscape(m.ETOPMRequest.ECPlatform)
	m.ETOPMRequest.TotalPrice = url.QueryEscape(m.ETOPMRequest.TotalPrice)
	m.ETOPMRequest.OrderInfo = url.QueryEscape(m.ETOPMRequest.OrderInfo)
	m.ETOPMRequest.Note1 = url.QueryEscape(m.ETOPMRequest.Note1)
	m.ETOPMRequest.Note2 = url.QueryEscape(m.ETOPMRequest.Note2)
	m.ETOPMRequest.PayType = url.QueryEscape(m.ETOPMRequest.PayType)
	m.ETOPMRequest.AtmRespost = url.QueryEscape(m.ETOPMRequest.AtmRespost)
	m.ETOPMRequest.DeadLine = url.QueryEscape(m.ETOPMRequest.DeadLine)
	m.ETOPMRequest.PayEN = url.QueryEscape(m.ETOPMRequest.PayEN)
	m.ETOPMRequest.CodeType = url.QueryEscape(m.ETOPMRequest.CodeType)
	m.ETOPMRequest.EPT = url.QueryEscape("1")
	m.ETOPMRequest.Installment = url.QueryEscape(m.ETOPMRequest.Installment)
	m.ETOPMRequest.EPT = url.QueryEscape(m.ETOPMRequest.PayDay)
	m.ETOPMRequest.CIFPW = url.QueryEscape(m.ETOPMRequest.CIFPW)
	m.ETOPMRequest.CIFID = url.QueryEscape(m.ETOPMRequest.CIFID)
	m.ETOPMRequest.CIFID_SN = url.QueryEscape(m.ETOPMRequest.CIFID_SN)
}

func EncodeSHA1(value string) (result string) {
	h := sha1.New()
	h.Write([]byte(value))
	result = hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(result)
}

func (m *ETOPMRequestCall) Do() string {
	m.EncodeValues()
	params := StructToParamsMap(m.ETOPMRequest)
	html := GenerateAutoSubmitHtmlForm(params, PayNowETOPMURL)
	return html
}

func (m *ETOPMRequestCall) DoTest() string {
	m.EncodeValues()
	params := StructToParamsMap(m.ETOPMRequest)
	html := GenerateAutoSubmitHtmlForm(params, TestPayNowETOPMURL)
	return html
}

type ETOPMCardResponse struct {
	//商家帳號
	WebNo string `json:"WebNo,omitempty"`
	//PayNow訂單編號
	BuysafeNo string `json:"BuysafeNo"`
	//驗證碼
	PassCode string `json:"PassCode,omitempty"`
	//商家自訂編號
	OrderNo string `json:"OrderNo,omitempty"`
	//交易結果
	TranStatus string `json:"TranStatus,omitempty"`
	//錯誤描述
	ErrDesc string `json:"ErrDesc,omitempty"`
	//交易金額
	TotalPrice string `json:"TotalPrice,omitempty"`
	//備註 1
	Note1 string `json:"Note1,omitempty"`
	//備註 2
	Note2 string `json:"Note2,omitempty"`
	//付款方式
	PayType string `json:"PayType,omitempty"`
	//信用卡末四碼
	PanNo4 string `json:"pan_no4,omitempty"`
	//是否國外卡
	CardForeign string `json:"Card_Foreign,omitempty"`
	//信用卡分期期數 預備繳款期數
	Installment string `json:"installment,omitempty"`
	//授權日
	PayDay string `json:"PayDay,omitempty"`
	//UserID
	CIFID string `json:"CIFID,omitempty"`
	//UserPW
	CIFPW string `json:"CIFPW,omitempty"`
	//SN(序號)
	CIFID_SN string `json:"CIFID_SN,omitempty"`
}

func NewETOPMCardResponse() *ETOPMCardResponse {
	return &ETOPMCardResponse{}
}

func (e *ETOPMCardResponse) Decode() {
	e.WebNo, _ = url.QueryUnescape(e.WebNo)
	e.BuysafeNo, _ = url.QueryUnescape(e.BuysafeNo)
	e.PassCode, _ = url.QueryUnescape(e.PassCode)
	e.OrderNo, _ = url.QueryUnescape(e.OrderNo)
	e.TranStatus, _ = url.QueryUnescape(e.TranStatus)
	e.ErrDesc, _ = url.QueryUnescape(e.ErrDesc)
	e.TotalPrice, _ = url.QueryUnescape(e.TotalPrice)
	e.Note1, _ = url.QueryUnescape(e.Note1)
	e.Note2, _ = url.QueryUnescape(e.Note2)
	e.PayType, _ = url.QueryUnescape(e.PayType)
	e.PanNo4, _ = url.QueryUnescape(e.PanNo4)
	e.CardForeign, _ = url.QueryUnescape(e.CardForeign)
	e.Installment, _ = url.QueryUnescape(e.Installment)
	e.PayDay, _ = url.QueryUnescape(e.PayDay)
	e.CIFID, _ = url.QueryUnescape(e.CIFID)
	e.CIFPW, _ = url.QueryUnescape(e.CIFPW)
	e.CIFID_SN, _ = url.QueryUnescape(e.CIFID_SN)
}

type ETOPMATMResponse struct {
	//商家帳號
	WebNo string `json:"WebNo,omitempty"`
	//PayNow訂單編號
	BuysafeNo string `json:"BuysafeNo"`
	//驗證碼
	PassCode string `json:"PassCode,omitempty"`
	//商家自訂編號
	OrderNo string `json:"OrderNo,omitempty"`
	//錯誤描述
	ErrDesc string `json:"ErrDesc,omitempty"`
	//交易金額
	TotalPrice string `json:"TotalPrice,omitempty"`
	//備註 1
	Note1 string `json:"Note1,omitempty"`
	//備註 2
	Note2 string `json:"Note2,omitempty"`
	//付款方式
	PayType string `json:"PayType,omitempty"`
	//虛擬帳號號碼
	ATMNo string `json:"ATMNo,omitempty"`
	//產生日期(繳款日)
	NewDate string `json:"NewDate,omitempty"`
	//繳款期限
	DueDate string `json:"DueDate,omitempty"`
	//繳款狀態
	TranStatus string `json:"TranStatus,omitempty"`
	//銀行代碼
	BankCode string `json:"BankCode,omitempty"`
	//分行號碼
	BranchCode string `json:"BranchCode,omitempty"`
}

func NewETOPMATMResponse() *ETOPMATMResponse {
	return &ETOPMATMResponse{}
}

func (e *ETOPMATMResponse) Decode() {
	e.WebNo, _ = url.QueryUnescape(e.WebNo)
	e.BuysafeNo, _ = url.QueryUnescape(e.BuysafeNo)
	e.PassCode, _ = url.QueryUnescape(e.PassCode)
	e.OrderNo, _ = url.QueryUnescape(e.OrderNo)
	e.TranStatus, _ = url.QueryUnescape(e.TranStatus)
	e.ErrDesc, _ = url.QueryUnescape(e.ErrDesc)
	e.TotalPrice, _ = url.QueryUnescape(e.TotalPrice)
	e.Note1, _ = url.QueryUnescape(e.Note1)
	e.Note2, _ = url.QueryUnescape(e.Note2)
	e.PayType, _ = url.QueryUnescape(e.PayType)
	e.ATMNo, _ = url.QueryUnescape(e.ATMNo)
	e.NewDate, _ = url.QueryUnescape(e.NewDate)
	e.DueDate, _ = url.QueryUnescape(e.DueDate)
	e.TranStatus, _ = url.QueryUnescape(e.TranStatus)
	e.BankCode, _ = url.QueryUnescape(e.BankCode)
	e.BranchCode, _ = url.QueryUnescape(e.BranchCode)

}

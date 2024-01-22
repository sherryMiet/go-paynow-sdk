package go_paynow_sdk

import (
	"reflect"
	"testing"
)

func TestAes256(t *testing.T) {
	type args struct {
		plaintext string
		key       string
		iv        string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Aes256(tt.args.plaintext, tt.args.key, tt.args.iv); got != tt.want {
				t.Errorf("Aes256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRefundRequestCall_Do(t *testing.T) {
	type fields struct {
		Client              *Client
		CreateRefundRequest *CreateRefundRequest
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateRefundRequestCall{
				Client:              tt.fields.Client,
				CreateRefundRequest: tt.fields.CreateRefundRequest,
			}
			c.Do()
		})
	}
}

func TestCreateRefundRequestCall_DoTest(t *testing.T) {
	type fields struct {
		Client              *Client
		CreateRefundRequest *CreateRefundRequest
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateRefundRequestCall{
				Client:              tt.fields.Client,
				CreateRefundRequest: tt.fields.CreateRefundRequest,
			}
			c.DoTest()
		})
	}
}

func TestCreateRefundRequest_Encode(t *testing.T) {
	type fields struct {
		OP       string
		JStr1    string
		JStr2    string
		MemCid   string
		TimeStr  string
		CheckNum string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CreateRefundRequest{
				OP:       tt.fields.OP,
				JStr1:    tt.fields.JStr1,
				JStr2:    tt.fields.JStr2,
				MemCid:   tt.fields.MemCid,
				TimeStr:  tt.fields.TimeStr,
				CheckNum: tt.fields.CheckNum,
			}
			c.Encode()
		})
	}
}

func TestDecodeAes256(t *testing.T) {
	type args struct {
		cipherText string
		key        string
		iv         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeAes256(tt.args.cipherText, tt.args.key, tt.args.iv); got != tt.want {
				t.Errorf("DecodeAes256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCheckCodeGKRequestCall_Do(t *testing.T) {
	type fields struct {
		Client                    *Client
		GetCheckCodeRequest       *GetCheckCodeRequest
		GetCheckCodeJsonGKRequest *GetCheckCodeJsonGKRequest
	}
	tests := []struct {
		name         string
		fields       fields
		wantResponse *GetCheckCodeJsonGKResponse
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeGKRequestCall{
				Client:                    tt.fields.Client,
				GetCheckCodeRequest:       tt.fields.GetCheckCodeRequest,
				GetCheckCodeJsonGKRequest: tt.fields.GetCheckCodeJsonGKRequest,
			}
			gotResponse, err := g.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Do() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestGetCheckCodeGKRequestCall_DoTest(t *testing.T) {
	type fields struct {
		Client                    *Client
		GetCheckCodeRequest       *GetCheckCodeRequest
		GetCheckCodeJsonGKRequest *GetCheckCodeJsonGKRequest
	}
	type test struct {
		name         string
		fields       fields
		want         string
		wantErr      bool
		wantResponse any
	}

	var tests = []test{
		test{
			name: "success",
			fields: fields{
				Client: &Client{
					Account:  "28229955",
					Password: "s192837465"},
				GetCheckCodeJsonGKRequest: &GetCheckCodeJsonGKRequest{
					MemCid:   "28229955",
					PassCode: "",
					TimeStr:  "",
					CheckNum: "33528305",
				},
				GetCheckCodeRequest: &GetCheckCodeRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeGKRequestCall{
				Client:                    tt.fields.Client,
				GetCheckCodeRequest:       tt.fields.GetCheckCodeRequest,
				GetCheckCodeJsonGKRequest: tt.fields.GetCheckCodeJsonGKRequest,
			}
			gotResponse, err := g.DoTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("DoTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("DoTest() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestGetCheckCodeGPRequestCall_Do(t *testing.T) {
	type fields struct {
		Client                    *Client
		GetCheckCodeRequest       *GetCheckCodeRequest
		GetCheckCodeJsonGPRequest *GetCheckCodeJsonGPRequest
	}
	type test struct {
		name         string
		fields       fields
		want         string
		wantErr      bool
		wantResponse any
	}

	var tests = []test{
		test{
			name: "success",
			fields: fields{
				Client: &Client{
					Account:  "28229955",
					Password: "s192837465"},
				GetCheckCodeJsonGPRequest: &GetCheckCodeJsonGPRequest{
					MemCid:   "28229955",
					PassCode: "",
					TimeStr:  "9328005018",
				},
				GetCheckCodeRequest: &GetCheckCodeRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeGPRequestCall{
				Client:                    tt.fields.Client,
				GetCheckCodeRequest:       tt.fields.GetCheckCodeRequest,
				GetCheckCodeJsonGPRequest: tt.fields.GetCheckCodeJsonGPRequest,
			}
			gotResponse, err := g.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Do() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestGetCheckCodeGPRequestCall_DoTest(t *testing.T) {
	type fields struct {
		Client                    *Client
		GetCheckCodeRequest       *GetCheckCodeRequest
		GetCheckCodeJsonGPRequest *GetCheckCodeJsonGPRequest
	}
	type test struct {
		name         string
		fields       fields
		want         string
		wantErr      bool
		wantResponse any
	}

	var tests = []test{
		test{
			name: "success",
			fields: fields{
				Client: &Client{
					Account:  "28229955",
					Password: "s192837465"},
				GetCheckCodeJsonGPRequest: &GetCheckCodeJsonGPRequest{
					MemCid:   "28229955",
					PassCode: "",
					TimeStr:  "9328005018",
				},
				GetCheckCodeRequest: &GetCheckCodeRequest{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeGPRequestCall{
				Client:                    tt.fields.Client,
				GetCheckCodeRequest:       tt.fields.GetCheckCodeRequest,
				GetCheckCodeJsonGPRequest: tt.fields.GetCheckCodeJsonGPRequest,
			}
			gotResponse, err := g.DoTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("DoTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("DoTest() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestGetCheckCodeJsonGKRequest_GetValues(t *testing.T) {
	type fields struct {
		MemCid   string
		PassCode string
		TimeStr  string
		CheckNum string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeJsonGKRequest{
				MemCid:   tt.fields.MemCid,
				PassCode: tt.fields.PassCode,
				TimeStr:  tt.fields.TimeStr,
				CheckNum: tt.fields.CheckNum,
			}
			g.GetValues()
		})
	}
}

func TestGetCheckCodeJsonGPRequest_GetValues(t *testing.T) {
	type fields struct {
		MemCid   string
		PassCode string
		TimeStr  string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GetCheckCodeJsonGPRequest{
				MemCid:   tt.fields.MemCid,
				PassCode: tt.fields.PassCode,
				TimeStr:  tt.fields.TimeStr,
			}
			g.GetValues()
		})
	}
}

func TestGetPassCode(t *testing.T) {
	type args struct {
		mem_cid    string
		PowerCheck string
		ByKey      bool
		key        string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPassCode(tt.args.mem_cid, tt.args.PowerCheck, tt.args.ByKey, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPassCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPassCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPowerCheck(t *testing.T) {
	type args struct {
		mem_cid   string
		TimeStr   string
		InputFlag bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPowerCheck(tt.args.mem_cid, tt.args.TimeStr, tt.args.InputFlag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPowerCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPowerCheck() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTimeStr(t *testing.T) {
	tests := []struct {
		name        string
		wantTimeStr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTimeStr := GetTimeStr(); gotTimeStr != tt.wantTimeStr {
				t.Errorf("GetTimeStr() = %v, want %v", gotTimeStr, tt.wantTimeStr)
			}
		})
	}
}

func TestSHA256_Encrypt(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256_Encrypt(tt.args.val); got != tt.want {
				t.Errorf("SHA256_Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSHA256_HMACSHA256(t *testing.T) {
	type args struct {
		val string
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256_HMACSHA256(tt.args.val, tt.args.key); got != tt.want {
				t.Errorf("SHA256_HMACSHA256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSendRequest(t *testing.T) {
	type args struct {
		postData *map[string]string
		URL      string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendRequest(tt.args.postData, tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SendRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

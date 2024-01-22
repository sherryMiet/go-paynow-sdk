package go_paynow_sdk

import (
	"github.com/rs/xid"
	"reflect"
	"testing"
)

func TestClient_NewETOPM(t *testing.T) {
	type fields struct {
		Password string
		Account  string
	}
	tests := []struct {
		name   string
		fields fields
		want   *ETOPMRequestCall
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Password: tt.fields.Password,
				Account:  tt.fields.Account,
			}
			if got := c.NewETOPM(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewETOPM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestETOPMRequestCall_Do(t *testing.T) {
	type fields struct {
		Client       *Client
		ETOPMRequest *ETOPMRequest
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ETOPMRequestCall{
				Client:       tt.fields.Client,
				ETOPMRequest: tt.fields.ETOPMRequest,
			}
			if got := m.Do(); got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestETOPMRequestCall_DoTest(t *testing.T) {
	type fields struct {
		Client       *Client
		ETOPMRequest *ETOPMRequest
	}

	type test struct {
		name   string
		fields fields
		want   string
	}

	var tests = []test{
		test{
			name: "success",
			fields: fields{
				Client: &Client{
					Account:  "91094694",
					Password: "s192837465"},
				ETOPMRequest: &ETOPMRequest{
					ReceiverName:  "俞育襦",
					ReceiverEmail: "sherry2000307@gmail.com",
					ReceiverID:    "sherry2000307@gmail.com",
					ReceiverTel:   "0909508777",
					TotalPrice:    "39",
					OrderInfo:     "幫賣",
					OrderNo:       xid.New().String(),
					ECPlatform:    "幫賣",
					PayType:       "03",
					EPT:           "1",
					AtmRespost:    "01",
					//CIFID:         "tt1234567890123456",
					//CIFPW:         "tt1234567890123456",
					//Installment:   "99",
					//PayDay:        "01",
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ETOPMRequestCall{
				Client:       tt.fields.Client,
				ETOPMRequest: tt.fields.ETOPMRequest,
			}
			if got := m.DoTest(); got != tt.want {
				t.Errorf("DoTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestETOPMRequestCall_EncodeValues(t *testing.T) {
	type fields struct {
		Client       *Client
		ETOPMRequest *ETOPMRequest
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ETOPMRequestCall{
				Client:       tt.fields.Client,
				ETOPMRequest: tt.fields.ETOPMRequest,
			}
			m.EncodeValues()
		})
	}
}

func TestEncodeSHA1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := EncodeSHA1(tt.args.value); gotResult != tt.wantResult {
				t.Errorf("EncodeSHA1() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

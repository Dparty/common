package sms

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/Dparty/common/data"
)

const sendCloudApi = "https://api.sendcloud.net/smsapi/send"

// Function(user , key string) return SendCloud struct
func NewSendCloud(user, key string) *SendCloud {
	return &SendCloud{
		user: user,
		key:  key,
	}
}

// Struct called SendCloud implement SmsSender interface
type SendCloud struct {
	user string
	key  string
}

// SendCloud method Send implement SmsSender interface
func (s *SendCloud) Send(to PhoneNumber, message string) error {
	return errors.New("not implement")
}

// SendCloud method SendWithTemplate implement SmsSender interface
func (s *SendCloud) SendWithTemplate(to PhoneNumber, templateId string, vars map[string]string) bool {
	switch to.AreaCode {
	case "86":
		return s.sendWithTemplate(to, templateId, vars, false)
	case "853", "852":
		return s.sendWithTemplate(to, templateId, vars, true)
	default:
		return false
	}
}

func (s SendCloud) Signature(templateId string, phone string, msgType string, vars map[string]string) string {
	var paramStr []string
	for _, pair := range s.Params(templateId, phone, msgType, vars) {
		paramStr = append(paramStr, pair.L+"="+pair.R)
	}
	byteArray := md5.Sum([]byte(s.key + "&" + strings.Join(paramStr, "&") + "&" + s.key))
	return hex.EncodeToString(byteArray[:])
}

func (s SendCloud) Params(templateId string, phone string, msgType string, vars map[string]string) []data.Pair[string, string] {
	jsonStr, _ := json.Marshal(vars)
	params := map[string]string{
		"smsUser":    s.user,
		"templateId": templateId,
		"vars":       string(jsonStr),
		"phone":      phone,
		"msgType":    msgType,
	}
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var pairs []data.Pair[string, string]
	for _, k := range keys {
		pairs = append(pairs, data.Pair[string, string]{
			L: k,
			R: params[k],
		})
	}
	return pairs
}

func (s SendCloud) sendWithTemplate(phoneNumber PhoneNumber, templateId string, vars map[string]string, international bool) bool {
	client := http.Client{}
	postValues := url.Values{}
	params := s.Params(templateId, phoneNumber.Number, "0", vars)
	for _, p := range params {
		postValues.Add(p.L, p.R)
	}
	postValues.Add("signature", s.Signature(templateId, phoneNumber.Number, "0", vars))
	if international {
		postValues.Add("msgType", "2")
	}
	resp, err := client.PostForm(sendCloudApi, postValues)
	if err != nil {
		return false
	}
	b, _ := io.ReadAll(resp.Body)
	var sendCloudJson sendCloudResp
	json.Unmarshal(b, &sendCloudJson)
	return sendCloudJson.Result
}

type sendCloudInfo struct {
	SuccessCount int      `json:"successCount"`
	SmsIds       []string `json:"smsIds"`
}

type sendCloudResp struct {
	Result     bool          `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Info       sendCloudInfo `json:"info"`
}

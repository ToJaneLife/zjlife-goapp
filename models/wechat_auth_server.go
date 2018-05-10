package models

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"zjlife/logger"
)

var (
	ErrAppIDNotMatch       = errors.New("app id not match")
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func weChatAuthS(code, iv, encrypted_data, pn string, user_id int) (state string, wui *WchatUserInfo, sk string, expire int) {
	wui = &WchatUserInfo{}
	appid := beego.AppConfig.String("appid")
	appsecret := beego.AppConfig.String("appsecret")
	if pn != "" {
		uw := getWechatUserSessionKeyS(pn)
		if e := int(time.Now().Sub(uw.SessionKeyUpdTime).Seconds()); e > uw.SessionExpires { //session_key过期
			goto lable
		} else {
			logger.Debug("session_key from db.", pn, uw.SessionKey, uw.SessionExpires, uw.UnionId)
			wui.UnionID, wui.OpenID, sk, expire = uw.UnionId, uw.OpenId, uw.SessionKey, uw.SessionExpires
			return
		}
	}
lable:
	open_id, sk, expire, err := wchatVerifyXcxCodeS(code, appid, appsecret)
	wui.OpenID = open_id
	if err != nil {
		state = "300"
		return
	}
	logger.Debug("session_key from wechat.", pn, open_id, sk, expire)
	// var err2 error
	// wui, err2 = wchatVerifXcxEncryptedDataS(sk, iv, encrypted_data, appid)
	// if err2 != nil {
	// 	state = "301"
	// 	return
	// }
	return
}

func getWechatUserSessionKeyS(pn string) UserWechatSessionKey {
	return queryWechatUserSessionKey(pn)
}

func wchatVerifyXcxCodeS(code, appid, appsecret string) (openid, sessionkey string, expire int, err error) {
	var clusterinfo = url.Values{}
	clusterinfo.Add("appid", appid)
	clusterinfo.Add("secret", appsecret)
	clusterinfo.Add("js_code", code)
	clusterinfo.Add("grant_type", "authorization_code")
	data := clusterinfo.Encode()

	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/sns/jscode2session", strings.NewReader(data))
	logger.Debug("wchart req = ", req.URL.String(), data)
	if err != nil {
		logger.Error("http.NewRequest err = ", err)
		return "", "", 0, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("client.Do err = ", err)
		return "", "", 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("ioutil.ReadAll err = ", err)
		return "", "", 0, err
	}
	logger.Debug("wchart rsp = ", string(body))
	var sos WchatSessionOpenId
	err = json.Unmarshal(body, &sos)
	if err != nil {
		logger.Error("json Unmarshal content = ", string(body), " err = ", err)
		return "", "", 0, err
	}
	return sos.XcxOpenId, sos.SessionKey, sos.Expires, nil
}

func wchatVerifXcxEncryptedDataS(sessionkey, iv, data, appid string) (wui *WchatUserInfo, err error) {
	wui, err = wchatDecrypt(appid, sessionkey, data, iv)
	if err != nil {
		logger.Error("VerifXcxEncryptedData err = ", err)
		return
	}
	logger.Debug("wchat UserInfo:", fmt.Sprintf("%#v", wui))
	return
}

func wchatDecrypt(app_id, session_key, encryptedData, iv string) (*WchatUserInfo, error) {
	aesKey, err := base64.StdEncoding.DecodeString(session_key)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	var wui WchatUserInfo
	err = json.Unmarshal(cipherText, &wui)
	if err != nil {
		return nil, err
	}
	if wui.Watermark.AppID != app_id {
		return nil, ErrAppIDNotMatch
	}
	return &wui, nil
}

// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	// logger.Debug("pkcs7Unpad", string(data), blockSize)
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

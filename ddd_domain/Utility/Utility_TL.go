package utility

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

//通联快捷支付测试
const (
	TLQuickpayAdrr = "https://113.108.182.3/aipg/quickpay"
	MerchantIdDev  = "200604000004910"
	TLUserDev      = "20060400000491004"
	TLPassWordDev  = "111111"
)

var TLSignCert *rsa.PrivateKey
var TLVerifyCert *rsa.PublicKey

//通联支付密钥初始化
func init() {
	carFile, err := ioutil.ReadFile("Assets/20060400000491004.pem")
	if err != nil {
		return
	}
	pemBlock, _ := pem.Decode(carFile)
	if pemBlock == nil {
		return
	}
	parsedKey, err := x509.ParsePKCS8PrivateKey(pemBlock.Bytes)
	if err != nil {
		return
	}
	TLSignCert = parsedKey.(*rsa.PrivateKey)

	verify, err := ioutil.ReadFile("Assets/allinpay-pdsDev.pem")
	if err != nil {
		return
	}
	block, _ := pem.Decode(verify)
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}
	TLVerifyCert = cert.PublicKey.(*rsa.PublicKey)
}

//公钥验证签名
func VerifyTLXmlByte(xmlByte []byte, sign string) error {
	// hex2bin
	signByte, err := hex.DecodeString(sign)
	if err != nil {
		return errors.New("hex.DecodeString error : " + err.Error())
	}
	// sha1
	hash := crypto.SHA1
	h := hash.New()
	if _, err = h.Write(xmlByte); err != nil {
		return errors.New("hash.Write error : " + err.Error())
	}
	hashed := h.Sum(nil)

	// verify
	if err = rsa.VerifyPKCS1v15(TLVerifyCert, hash, hashed, signByte); err != nil {
		return errors.New("rsa.VerifyPKCS1v15 error : " + err.Error())
	}
	return err
}

//私钥签名
func SignXmlByte(xmlByte []byte) (sign string, err error) {
	// sha1
	hash := crypto.SHA1
	h := hash.New()
	if _, err = h.Write(xmlByte); err != nil {
		err = errors.New("hash.Write error : " + err.Error())
		return sign, err
	}
	hashed := h.Sum(nil)
	// sign
	singByte, err := rsa.SignPKCS1v15(rand.Reader, TLSignCert, hash, hashed)
	if err != nil {
		err = errors.New("rsa.SignPKCS1v15 error : " + err.Error())
		return sign, err
	}
	// bin2hex
	sign = hex.EncodeToString(singByte)
	// return
	return sign, err
}

//设置签名,给XML某节点赋值
func SetTLReqSignValue(req interface{}, sign string) (err error) {
	// reflect value
	reflectValue := reflect.ValueOf(req)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		return errors.New("SetTLReqSignValue request interface not a pointer type : reflect.Ptr")
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if !infoValue.IsValid() {
		return errors.New("SetTLReqSignValue error : INFO flied not exist")
	}
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if !signValue.IsValid() {
		return errors.New("SetTLReqSignValue error : SIGNED_MSG flied not exist")
	}
	// string
	if signValue.Kind() != reflect.String {
		return errors.New("SetTLReqSignValue signValue not a string type : reflect.String")
	}
	// assignment value
	signValue.SetString(sign)
	// return
	return err
}

// 使用反射来获取XML数据？
func GetThenEmptyAllinPayResultSignValue(res interface{}) (sign string, err error) {
	// reflect value
	reflectValue := reflect.ValueOf(res)
	// pointer
	if reflectValue.Kind() != reflect.Ptr {
		err = errors.New("GetThenEmptyAllinPayResultSignValue result interface not a pointer type : reflect.Ptr")
		return sign, err
	}
	// info reflect value
	infoValue := reflectValue.Elem().FieldByName("INFO")
	// Info
	if !infoValue.IsValid() {
		err = errors.New("GetThenEmptyAllinPayResultSignValue error : INFO flied not exist")
		return sign, err
	}
	// not pointer
	signValue := infoValue.FieldByName("SIGNED_MSG")
	// SIGNED_MSG
	if !signValue.IsValid() {
		err = errors.New("GetThenEmptyAllinPayResultSignValue error : SIGNED_MSG flied not exist")
		return sign, err
	}
	// string
	if signValue.Kind() != reflect.String {
		err = errors.New("GetThenEmptyAllinPayResultSignValue signValue not a string type : reflect.String")
		return sign, err
	}
	// get sign
	sign = signValue.String()
	// assignment value
	signValue.SetString("")
	// return
	return sign, err
}

//通联请求数据，序列化并加签名
func ToTLRequestXmlByte(req interface{}) (r []byte, err error) {
	// xml
	xmlByte, err := ToGbkXmlByte(req)
	if err != nil {
		err = errors.New("req ToTLRequestXmlByte error : " + err.Error())
		return r, err
	}
	// sign
	sign, err := SignXmlByte(xmlByte)
	if err != nil {
		err = errors.New("req SignXmlByte error : " + err.Error())
		return r, err
	}
	// AssignmentSignValue
	if err = SetTLReqSignValue(req, sign); err != nil {
		err = errors.New("req AssignmentSignValue error : " + err.Error())
		return r, err
	}
	// return
	return ToGbkXmlByte(req)
}

//验证通联返回值与签名
func VerifyTLResponse(bodyByte []byte, resp interface{}) (err error) {
	// gbk to utf8
	utf8Byte, err := GbkToUtf8(bodyByte)
	if err != nil {
		return errors.New("GbkToUtf8 error : " + err.Error())
	}
	// replace gbk header to utf8 header
	utf8Byte = []byte(strings.Replace(string(utf8Byte), XmlGbkHeader, XmlUtf8Header, 1))
	// decode
	if err = xml.Unmarshal(utf8Byte, resp); err != nil {
		return errors.New("xml.Unmarshal error : " + err.Error())
	}
	// get sign string and set empty value
	signString, err := GetThenEmptyAllinPayResultSignValue(resp)
	if err != nil {
		return errors.New("GetThenEmptyAllinPayResultSignValue error : " + err.Error())
	}
	// 替换加密字符串
	replacePattern := `<SIGNED_MSG>.*<\/SIGNED_MSG>`
	replaceRegexp, err := regexp.Compile(replacePattern)
	if err != nil {
		return errors.New("regexp.Compile error : " + err.Error())
	}
	verifyByte := []byte(replaceRegexp.ReplaceAllString(string(bodyByte), ""))
	// verify
	if err = VerifyTLXmlByte(verifyByte, signString); err != nil {
		return errors.New("VerifyTLXmlByte error : " + err.Error())
	}
	return err
}

func PostAllinPayXml(url string, xml []byte) (r []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(url, "application/xml", bytes.NewReader(xml))
	if err != nil {
		err = errors.New("http post error : " + err.Error())
		return r, err
	}
	defer resp.Body.Close()
	// status ok
	if resp.StatusCode != http.StatusOK {
		err = errors.New("http post fail, error code : " + resp.Status)
		return r, err
	}
	// return
	if r, err = ioutil.ReadAll(resp.Body); err != nil {
		err = errors.New("ioutil.ReadAll error : " + resp.Status)
		return r, err
	}
	return r, err
}

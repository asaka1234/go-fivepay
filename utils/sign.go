package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"sort"
	"strings"
	"time"
)

// 字母顺序a-z排序，并串联对应的值（value1value2 ...valueN）转换为字符串，先使用 SHA1 签名，然后使用 MD5 签名，并转换为小写字母
func DepositSign(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 按键排序

	var strBuilder strings.Builder
	for _, k := range keys {
		// 跳过 'sign' 字段自身，如果它存在于 params 中
		if k == "sign" {
			continue
		}
		if k == "merchantId" {
			strBuilder.WriteString(cast.ToString(params[k].(int)))
		} else {
			strBuilder.WriteString(params[k].(string))
		}
	}
	signStr := strBuilder.String()

	// Log before MD5
	log.Printf("Fivepay#MD5#deposit#before, s: %s", signStr)

	//signStr = "dda78446523b6f6ac3910624f719ca37e0509d49bf11c6fc5355293f91cc7e7b6cb7d633ba489ea1236efbca11de5f489e6715ec7e57bff118c44http://localhost:11819/FiatToFiat/PaymentStatusCallback3d0dff2f964ad5cdhttp://localhost:11819/FiatToFiat/PaymentCompleted"
	// 创建SHA1哈希
	sha1Hash := sha1.Sum([]byte(signStr)) // sha1.Sum 直接计算并返回哈希值
	sha1HashStr := hex.EncodeToString(sha1Hash[:])
	log.Printf("Fivepay#SHA1#deposit#end, s: %s", hex.EncodeToString(sha1Hash[:]))

	// Generate MD5 hash
	hash := md5.Sum([]byte(sha1HashStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("Fivepay#MD5#deposit#end, s: %s", result)

	return result
}

// EncryptAll 加密所有需要加密的参数
func EncryptAll(params map[string]interface{}, accessKey string) (map[string]interface{}, error) {
	paramEncrypt := make(map[string]interface{})

	paramEncrypt["merchantId"] = params["merchantId"] // merchantId 不加密

	// 其他字段加密
	fieldsToEncrypt := []string{"orderAmount", "currencyCode", "merchantOrderNo", "memberId", "name", "email"}
	for _, field := range fieldsToEncrypt {
		val := params[field]
		if field == "merchantOrderNo" {
			val = strings.ToLower(val.(string)) // merchantOrderNo 需要小写
		}

		encryptedVal, err := encrypt(val.(string), accessKey)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt %s: %w", field, err)
		}
		paramEncrypt[field] = encryptedVal
	}

	// 这里的 URL 需要根据你的 Go 应用的实际回调地址设置
	paramEncrypt["notifyUrl"] = params["notifyUrl"]
	paramEncrypt["returnUrl"] = params["returnUrl"]

	return paramEncrypt, nil
}

// PKCS7Padding 填充函数
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去填充函数
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//length := len(origData)
	//unpadding := int(origData[length-1])
	//return origData[:(length - unpadding)]
	length := len(origData)
	if length == 0 {
		return nil, errors.New("输入数据为空")
	}

	// 获取填充字节数
	unPadding := int(origData[length-1])

	// 验证填充是否有效
	if unPadding > aes.BlockSize || unPadding == 0 {
		return nil, errors.New(string(origData) + " 无效的PKCS7填充")
	}

	// 验证所有填充字节是否正确
	for i := length - unPadding; i < length; i++ {
		if int(origData[i]) != unPadding {
			return nil, errors.New("无效的PKCS7填充")
		}
	}

	return origData[:(length - unPadding)], nil
}

// encrypt 使用 TripleDES DES-EDE3-CBC 模式加密
func encrypt(data, secret string) (string, error) {
	key := []byte(secret)
	// 根据 PHP 代码，取 secret 的前 8 字节作为 subkey (IV)
	// PHP's openssl_encrypt with DES-EDE3-CBC expects a 24-byte key and 8-byte IV.
	// The provided PHP code's key handling:
	// "$subkey = substr($secret, 0, 8);" and then used as IV directly.
	// This implies the key is 24 bytes long and the IV is the first 8 bytes.
	// We will use the full 24-byte secret as the key for NewTripleDESCipher.
	// And the first 8 bytes of the secret as IV.
	if len(key) != 24 {
		return "", fmt.Errorf("TripleDES key must be 24 bytes long, got %d", len(key))
	}
	iv := key[:8] // IV 是密钥的前 8 字节

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	origData := PKCS7Padding([]byte(data), blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return hex.EncodeToString(crypted), nil
}

//---------------

// DecryptAll 解密所有需要解密的参数
func DecryptAll(params map[string]interface{}, accessKey string) (map[string]interface{}, error) {
	paramDecrypt := make(map[string]interface{})

	paramDecrypt["merchantId"] = params["merchantId"] // merchantId 不解密

	// 其他字段解密
	fieldsToDecrypt := []string{"orderAmount", "currencyCode", "merchantOrderNo", "memberId"}
	for _, field := range fieldsToDecrypt {
		val := strings.ToUpper(params[field].(string))

		decryptedVal, err := decrypt(val, accessKey)
		if err != nil {
			return nil, fmt.Errorf("failed to encrypt %s: %w", field, err)
		}
		paramDecrypt[field] = decryptedVal
	}

	paramDecrypt["notifyUrl"] = params["notifyUrl"]
	paramDecrypt["returnUrl"] = params["returnUrl"]

	return paramDecrypt, nil
}

func DepositBackSign(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 按键排序

	var strBuilder strings.Builder
	for _, k := range keys {
		// 跳过 'sign' 字段自身，如果它存在于 params 中
		if k == "sign" || k == "merchantId" {
			continue
		}
		strBuilder.WriteString(params[k].(string))
	}
	signStr := strBuilder.String()

	// Log before MD5
	log.Printf("Fivepay#MD5#depositBack#before, s: %s", signStr)

	//signStr = "dda78446523b6f6ac3910624f719ca37e0509d49bf11c6fc5355293f91cc7e7b6cb7d633ba489ea1236efbca11de5f489e6715ec7e57bff118c44http://localhost:11819/FiatToFiat/PaymentStatusCallback3d0dff2f964ad5cdhttp://localhost:11819/FiatToFiat/PaymentCompleted"
	// 创建SHA1哈希
	sha1Hash := sha1.Sum([]byte(signStr)) // sha1.Sum 直接计算并返回哈希值
	sha1HashStr := hex.EncodeToString(sha1Hash[:])
	log.Printf("Fivepay#SHA1#depositBack#end, s: %s", hex.EncodeToString(sha1Hash[:]))

	// Generate MD5 hash
	hash := md5.Sum([]byte(sha1HashStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("Fivepay#MD5#depositBack#end, s: %s", result)

	return result
}

// decrypt 使用 TripleDES DES-EDE3-CBC 模式解密
func decrypt(data, secret string) (string, error) {
	// 16进制解码
	messageBytes, err := hex.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("hex解码失败: %v", err)
	}

	// 检查密钥长度
	if len(secret) != 24 {
		return "", fmt.Errorf("密钥必须为24字节")
	}

	// 创建3DES cipher
	block, err := des.NewTripleDESCipher([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("创建cipher失败: %v", err)
	}

	// 使用前8字节作为IV
	iv := []byte(secret)[:8]

	// CBC模式解密
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(messageBytes))
	mode.CryptBlocks(plaintext, messageBytes)

	// PKCS7反填充
	unpadded, err := PKCS7UnPadding(plaintext)
	if err != nil {
		return "", fmt.Errorf("去除填充失败: %v", err)
	}

	return string(unpadded), nil
}

func DepositBackVerify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature := DepositBackSign(params)

	// Compare signatures
	return signature == currentSignature, nil
}

// MD5({MerchantCode}{TransactionId}{MemberCode}{Amount}{CurrencyCode}){TransactionDateTime}){ToBankAccountNumber}){SecurityCode}))
func WithdrawSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"MerchantCode", "TransactionID", "MemberCode", "Amount", "CurrencyCode", "TransactionDateTime", "toBankAccountNumber", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])

			if k == "TransactionDateTime" {
				t, _ := time.Parse("2006-01-02 03:04:05PM", value)
				value = t.Format("20060102150405")
			}
			//fmt.Printf("%s=>%s\n", k, value)
			sb.WriteString(value)
		} else {
			//fmt.Printf("%s=>%s\n", k, key)
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

//---------------

// MD5({MerchantCode}{TransactionID}{MemberCode}{Amount}{CurrencyCode}{Status}{SecurityCode}
func WithdrawBackSign(params map[string]interface{}, key string) string {

	//参与签名的key
	signKeyList := []string{"MerchantCode", "TransactionID", "MemberCode", "Amount", "CurrencyCode", "Status", "SecurityCode"}

	//拼凑字符串
	var sb strings.Builder
	for _, k := range signKeyList {
		if k != "SecurityCode" {
			value := cast.ToString(params[k])
			sb.WriteString(value)
		} else {
			sb.WriteString(key)
		}
	}
	signStr := sb.String()

	// Log before MD5
	log.Printf("H2PayService#MD5#deposit#before, s: %s", signStr)

	// Generate MD5 hash
	hash := md5.Sum([]byte(signStr))
	result := hex.EncodeToString(hash[:])

	// Log after MD5
	log.Printf("H2PayService#MD5#deposit#end, s: %s", result)

	return result
}

func WithdrawBackVerify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["Key"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "Key")

	// Generate current signature
	currentSignature := WithdrawBackSign(params, signKey)

	// Compare signatures
	return signature.(string) == currentSignature, nil
}

package encry

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"study-server/app/libs/utils"
	"study-server/bootstrap/config"
)

// RSA加解密类,需要先初始化
type RSA struct {
	publicKey  []byte // 公钥 openssl genrsa -out rsa_private_key.pem 1024
	privateKey []byte // 私钥 openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
}

// 初始化
// 生成密钥(4096 PKCS1) https://uutool.cn/rsa-generate/
/**
 *@Example:
	r := encry.RSA{}
	fmt.Println(r.New().RsaDecrypt([]byte("xxx")))
*/
func (r *RSA) New() (*RSA, error) {
	public, e := utils.ReadFile(config.App.RSA.PublicKey)
	if e != nil {
		return &RSA{}, e
	}
	private, e := utils.ReadFile(config.App.RSA.PrivateKey)
	if e != nil {
		return &RSA{}, e
	}
	return &RSA{
		publicKey:  public,
		privateKey: private,
	}, nil
}

// 加密
func (r *RSA) RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(r.publicKey)
	if block == nil {
		return nil, errors.New("rsa public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//加密
	res, e := rsa.EncryptPKCS1v15(rand.Reader, pubInterface, origData)
	if e != nil {
		return nil, e
	}
	return res, nil
}

// 解密
func (r *RSA) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(r.privateKey)
	if block == nil {
		return nil, errors.New("rsa private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

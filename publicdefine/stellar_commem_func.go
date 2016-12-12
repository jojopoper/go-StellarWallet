package publicdefine

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/stellar/go/keypair"
	"github.com/stellar/go/strkey"
)

// CURRENT_DIR 当前目录
var CURRENT_DIR = ""

// VerifyGAddress 判断地址是否正确
func VerifyGAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteAccountID, addr)
	return err
}

// VerifySAddress 判断地址是否正确
func VerifySAddress(addr string) error {
	_, err := strkey.Decode(strkey.VersionByteSeed, addr)
	return err
}

// FromSeed2Address from seed to address
func FromSeed2Address(seed string) (addr string, err error) {
	err = VerifySAddress(seed)
	if err != nil {
		return
	}

	kp, err := keypair.Parse(seed)
	if err != nil {
		return
	}
	addr = kp.Address()
	return
}

// AESEncode aes encode
func AESEncode(pw string, src []byte) (ret string, err error) {
	m5 := md5.New()
	_, err = m5.Write([]byte(pw))
	if err != nil {
		return
	}
	bm5 := m5.Sum(nil)
	cb, err := aes.NewCipher(bm5)
	if err != nil {
		return
	}

	hm5 := hex.EncodeToString(bm5)
	m5.Write([]byte(hm5))

	cfb := cipher.NewCFBEncrypter(cb, m5.Sum(nil))
	ciphertext := make([]byte, len(src))
	cfb.XORKeyStream(ciphertext, src)
	ret = base64.StdEncoding.EncodeToString(ciphertext)
	return
}

// AESDecode aes decode
func AESDecode(pw string, length int, src string) (ret []byte, err error) {
	m5 := md5.New()
	_, err = m5.Write([]byte(pw))
	if err != nil {
		return
	}
	bm5 := m5.Sum(nil)
	hm5 := hex.EncodeToString(bm5)
	cb, err := aes.NewCipher(bm5)
	if err != nil {
		return
	}

	m5.Write([]byte(hm5))
	cfbdec := cipher.NewCFBDecrypter(cb, m5.Sum(nil))
	decode, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return
	}
	ret = make([]byte, length)
	cfbdec.XORKeyStream(ret, decode)
	return
}

// MkDir 创建目录
func MkDir(dirpath string) (err error) {
	err = os.MkdirAll(dirpath, 0777)
	return
}

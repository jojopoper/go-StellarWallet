package publicdefine

import (
	"fmt"

	_b "github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
)

// StellarAccountCreateInfo 创建恒星账户
type StellarAccountCreateInfo struct {
	SrcInfo    *StellarAccInfoDef
	Amount     float64
	Destinaton string
	signBase64 string
	ResultHash string
}

// GetSigned 签名
func (ths *StellarAccountCreateInfo) GetSigned(seed string) string {
	tx := _b.TransactionBuilder{}

	ca := _b.CreateAccountBuilder{}
	ca.Mutate(_b.Destination{AddressOrSeed: ths.Destinaton})
	ca.Mutate(_b.SourceAccount{AddressOrSeed: ths.SrcInfo.ID})
	ca.Mutate(_b.NativeAmount{Amount: fmt.Sprintf("%f", ths.Amount)})

	tx.Mutate(_b.Sequence{Sequence: uint64(xdr.SequenceNumber(ths.SrcInfo.NextSequence()))})
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		tx.Mutate(_b.TestNetwork)
	} else {
		tx.Mutate(_b.PublicNetwork)
	}
	tx.Mutate(ca)
	tx.Mutate(_b.SourceAccount{AddressOrSeed: ths.SrcInfo.ID})
	tx.TX.Fee = BASEMENT_FEE
	// result := tx.Sign(&spriv)
	result := tx.Sign(seed)

	var err error

	ths.signBase64, err = result.Base64()
	// fmt.Printf("tx base64: %s\r\n", ths.signBase64)

	if err == nil {
		return ths.signBase64
	}

	fmt.Println(err)
	return ""
}

// PutResult 设置结果
func (ths *StellarAccountCreateInfo) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	ths.ResultHash = ""
	if ok {
		ths.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

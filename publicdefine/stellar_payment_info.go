package publicdefine

import (
	"fmt"

	_b "github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
)

const (
	BASEMENT_FEE = 100
)

//StellarPaymentInfo 恒星支付
type StellarPaymentInfo struct {
	SrcInfo    *StellarAccInfoDef
	Amount     float64
	Destinaton string
	signBase64 string
	ResultHash string
}

// GetSigned 获取签名
func (ths *StellarPaymentInfo) GetSigned(seed string) string {
	tx := _b.TransactionBuilder{}
	pb := _b.PaymentBuilder{}
	des := _b.Destination{AddressOrSeed: ths.Destinaton}
	na := _b.NativeAmount{Amount: fmt.Sprintf("%f", ths.Amount)}

	pb.Mutate(des)
	pb.Mutate(na)

	tx.Mutate(_b.Sequence{Sequence: uint64(xdr.SequenceNumber(ths.SrcInfo.NextSequence()))})
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		tx.Mutate(_b.TestNetwork)
	} else {
		tx.Mutate(_b.PublicNetwork)
	}
	tx.Mutate(pb)
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
func (ths *StellarPaymentInfo) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	ths.ResultHash = ""
	if ok {
		ths.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

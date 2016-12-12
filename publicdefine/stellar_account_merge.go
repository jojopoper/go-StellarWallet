package publicdefine

import (
	"fmt"

	_b "github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
)

// StellarAccountMerge 合并恒星账户
type StellarAccountMerge struct {
	SrcInfo        *StellarAccInfoDef
	DestPublicAddr string
	signBase64     string
	ResultHash     string
}

// GetSigned 获取签名
func (ths *StellarAccountMerge) GetSigned(seed string) string {

	tx := _b.TransactionBuilder{}
	am := _b.AccountMergeBuilder{}
	am.Mutate(_b.Destination{AddressOrSeed: ths.DestPublicAddr})
	tx.Mutate(_b.Sequence{Sequence: uint64(xdr.SequenceNumber(ths.SrcInfo.NextSequence()))})
	if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
		tx.Mutate(_b.TestNetwork)
	} else {
		tx.Mutate(_b.PublicNetwork)
	}
	tx.Mutate(am)
	tx.Mutate(_b.SourceAccount{AddressOrSeed: ths.SrcInfo.ID})
	tx.TX.Fee = BASEMENT_FEE
	result := tx.Sign(seed)
	// tx.TX = &xdr.Transaction{}
	// opt := xdr.Operation{}
	// srcAccID := new(xdr.AccountId)
	// srcAccID.SetAddress(ths.SrcInfo.ID)
	// destAccID := new(xdr.AccountId)
	// destAccID.SetAddress(ths.DestPublicAddr)
	// // srcAccID, _ := stellarbase.AddressToAccountId(ths.SrcInfo.ID)
	// // destAccID, _ := stellarbase.AddressToAccountId(ths.DestPublicAddr)

	// opt.SourceAccount = srcAccID
	// opt.Body, _ = xdr.NewOperationBody(xdr.OperationTypeAccountMerge, destAccID)
	// tx.TX.Operations = make([]xdr.Operation, 0)
	// tx.TX.Operations = append(tx.TX.Operations, opt)

	// tx.Mutate(_b.Sequence{Sequence: uint64(xdr.SequenceNumber(ths.SrcInfo.NextSequence()))})
	// if STELLAR_DEFAULT_NETWORK == STELLAR_TEST_NETWORK {
	// 	tx.Mutate(_b.TestNetwork)
	// } else {
	// 	tx.Mutate(_b.PublicNetwork)
	// }
	// tx.Mutate(_b.SourceAccount{AddressOrSeed: ths.SrcInfo.ID})
	// tx.TX.Fee = BASEMENT_FEE
	// result := tx.Sign(seed)

	var err error

	ths.signBase64, err = result.Base64()
	if err == nil {
		return ths.signBase64
	}

	fmt.Println(err)
	return ""
}

// PutResult 打印结果
func (ths *StellarAccountMerge) PutResult(ret map[string]interface{}) {
	hash, ok := ret["hash"]
	ths.ResultHash = ""
	if ok {
		ths.ResultHash = hash.(string)
		return
	}
	fmt.Println(ret)
}

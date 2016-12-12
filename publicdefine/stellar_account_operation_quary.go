package publicdefine

import (
	"fmt"
	"regexp"
	"strings"
)

// SubAccOperQuaryRecordInterface 接口定义
type SubAccOperQuaryRecordInterface interface {
	ToString() string
	GetType() string
	DecodeBody(b map[string]interface{})
}

// AssetInfoBaseItem asset基础信息
type AssetInfoBaseItem struct {
	Asset_Code   string
	Asset_Type   string
	Asset_Issuer string
}

// SubAccOperRecordItemBase sub account operation record
type SubAccOperRecordItemBase struct {
	OpType          string
	SourceAccount   string
	TransactionHash string
}

// SubCreateAccountItem create account item
type SubCreateAccountItem struct {
	SubAccOperRecordItemBase
	Account         string
	Funder          string
	StartingBalance string
}

// SubPaymentItem payment item
type SubPaymentItem struct {
	SubAccOperRecordItemBase
	Amount      string
	FromAccount string
	ToAccount   string
	AssetInfo   *AssetInfoBaseItem
}

// SubChangeTrustItem change trust item
type SubChangeTrustItem struct {
	SubAccOperRecordItemBase
	AssetCode   string
	AssetIssuer string
	AssetType   string
	Trustee     string
	Trustor     string
}

// SubAccountMergeItem account merge item
type SubAccountMergeItem struct {
	SubAccOperRecordItemBase
	MergeSource string
	MergeInto   string
}

// SubSetOptionsItem set option item
type SubSetOptionsItem struct {
	SubAccOperRecordItemBase
	HomeDomain    string
	InflationDest string
	SignerKey     string
	SignerWeight  string
	SetFlags      []string
	ClearFlags    []string
}

// SubManageOfferItem manage offer item
type SubManageOfferItem struct {
	SubAccOperRecordItemBase
	Amount  string
	OfferID string
	Price   string
	Buying  *AssetInfoBaseItem
	Selling *AssetInfoBaseItem
}

// StellarAccOperationQuary stellar account operation quary
type StellarAccOperationQuary struct {
	QuaryCursor string
	IsEnd       bool
	Records     []SubAccOperQuaryRecordInterface
}

// ToString 返回字符串
func (ths *SubAccOperRecordItemBase) ToString() string {
	return fmt.Sprintf("          Type = [%s]\r\n SourceAccount = [%s]\r\n          Hash = [%s]\r\n",
		ths.OpType, ths.SourceAccount, ths.TransactionHash)
}

// GetType 得到类型
func (ths *SubAccOperRecordItemBase) GetType() string {
	return ths.OpType
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/413278933094401/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=413278933094401\u0026order=asc"
       },
       "self": {
         "href": "/operations/413278933094401"
       },
       "succeeds": {
         "href": "/operations?cursor=413278933094401\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/dcef180a209b3dab35791a56b175c18a3a9ee1c57062f74a4a885a1b7a8b8067"
       }
     },
     "account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "funder": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
     "id": 413278933094401,
     "paging_token": "413278933094401",
     "source_account": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
     "starting_balance": "10000.0",
     "type": "create_account",
     "type_i": 0
   },
*/

// DecodeBody 解码主体
func (ths *SubAccOperRecordItemBase) DecodeBody(b map[string]interface{}) {
	_links, _linksok := b["_links"]
	source_account, source_account_ok := b["source_account"]
	if _linksok && source_account_ok {
		ths.SourceAccount = source_account.(string)

		transaction, _ := _links.(map[string]interface{})["transaction"]
		href, _ := transaction.(map[string]interface{})["href"]
		hrefurl := href.(string)
		ths.TransactionHash = strings.Trim(hrefurl, "/transactions/")
	}
}

// ToString 返回字符串
func (ths *SubCreateAccountItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("        Funder = [%s]\r\n       Account = [%s]\r\n       Balance = [%s]\r\n",
		ths.Funder, ths.Account, ths.StartingBalance)
	return
}

func (ths *SubCreateAccountItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)
	account, accountok := b["account"]
	funder, funderok := b["funder"]
	starting_balance, starting_balanceok := b["starting_balance"]
	if accountok && funderok && starting_balanceok {
		ths.Account = account.(string)
		ths.Funder = funder.(string)
		ths.StartingBalance = starting_balance.(string)
	}
}

// ToString 返回字符串
func (ths *SubPaymentItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("          From = [%s]\r\n            To = [%s]\r\n        Amount = [%s]\r\n",
		ths.FromAccount, ths.ToAccount, ths.Amount)
	if ths.AssetInfo != nil {
		ret += fmt.Sprintf("          Code = [%s]\r\n        Issuer = [%s]\r\n",
			ths.AssetInfo.Asset_Code, ths.AssetInfo.Asset_Issuer)
	}
	return
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/477574593515521/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=477574593515521\u0026order=asc"
       },
       "self": {
         "href": "/operations/477574593515521"
       },
       "succeeds": {
         "href": "/operations?cursor=477574593515521\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/d3642254d90547a67a7f25827c61f79ca57010521615b5f391b5ac664aa42028"
       }
     },
     "amount": "10.0",
     "asset_type": "native",
     "from": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "id": 477574593515521,
     "paging_token": "477574593515521",
     "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "to": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "type": "payment",
     "type_i": 1
   },

*/

// DecodeBody 解码主体
func (ths *SubPaymentItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)
	amount, amountok := b["amount"]
	from, fromok := b["from"]
	to, took := b["to"]
	if amountok && fromok && took {
		ths.Amount = amount.(string)
		ths.FromAccount = from.(string)
		ths.ToAccount = to.(string)
	}
	asset_issuer, asset_issuerok := b["asset_issuer"]
	asset_type, asset_typeok := b["asset_type"]
	asset_code, asset_codeok := b["asset_code"]
	if asset_issuerok && asset_typeok && asset_codeok {
		if ths.AssetInfo == nil {
			ths.AssetInfo = &AssetInfoBaseItem{}
		}

		ths.AssetInfo.Asset_Code = asset_code.(string)
		ths.AssetInfo.Asset_Issuer = asset_issuer.(string)
		ths.AssetInfo.Asset_Type = asset_type.(string)
	}
}

// ToString 返回字符串
func (ths *SubChangeTrustItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("     AssetCode = [%s]\r\n       Trustee = [%s]\r\n       Trustor = [%s]\r\n",
		ths.AssetCode, ths.Trustee, ths.Trustor)
	return
}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/777758447767553/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=777758447767553\u0026order=asc"
       },
       "self": {
         "href": "/operations/777758447767553"
       },
       "succeeds": {
         "href": "/operations?cursor=777758447767553\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/973bed257adf83d4ffe4b9693a2ce7ffb91cbe5afaf4734bc1b7ef8f782f498b"
       }
     },
     "asset_code": "XLM",
     "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "asset_type": "credit_alphanum4",
     "id": 777758447767553,
     "limit": "922337203685.4775807",
     "paging_token": "777758447767553",
     "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
     "type": "change_trust",
     "type_i": 6
   }

*/

// DecodeBody 解码主体
func (ths *SubChangeTrustItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)
	asset_code, asset_codeok := b["asset_code"]
	asset_issuer, asset_issuerok := b["asset_issuer"]
	asset_type, asset_typeok := b["asset_type"]
	trustee, trusteeok := b["trustee"]
	trustor, trustorok := b["trustor"]
	if asset_codeok && asset_issuerok && asset_typeok && trusteeok && trustorok {
		ths.AssetCode = asset_code.(string)
		ths.AssetIssuer = asset_issuer.(string)
		ths.AssetType = asset_type.(string)
		ths.Trustee = trustee.(string)
		ths.Trustor = trustor.(string)
	}
}

// ToString 返回字符串
func (ths *SubAccountMergeItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("   MergeSource = [%s]\r\n     MergeInto = [%s]\r\n",
		ths.MergeSource, ths.MergeInto)
	return

}

/*
   {
     "_links": {
       "effects": {
         "href": "/operations/496962075889665/effects{?cursor,limit,order}",
         "templated": true
       },
       "precedes": {
         "href": "/operations?cursor=496962075889665\u0026order=asc"
       },
       "self": {
         "href": "/operations/496962075889665"
       },
       "succeeds": {
         "href": "/operations?cursor=496962075889665\u0026order=desc"
       },
       "transaction": {
         "href": "/transactions/c9819aa9d497279c69d49f5fa24942cea2312a0c46002148e51bb98b90d83a20"

       }
     },
     "account": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
     "id": 496962075889665,
     "into": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
     "paging_token": "496962075889665",
     "source_account": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
     "type": "account_merge",
     "type_i": 8
   },
*/

// DecodeBody 解码主体
func (ths *SubAccountMergeItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)
	account, accountok := b["account"]
	into, intook := b["into"]
	if accountok && intook {
		ths.MergeSource = account.(string)
		ths.MergeInto = into.(string)
	}
}

// ToString 返回字符串
func (ths *SubSetOptionsItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	if len(ths.SignerKey) > 0 {
		ret += fmt.Sprintf("     SignerKey = [%s]\r\n", ths.SignerKey)
	}
	if len(ths.SignerWeight) > 0 {
		ret += fmt.Sprintf("  SignerWeight = [%s]\r\n", ths.SignerWeight)
	}
	if len(ths.HomeDomain) > 0 {
		ret += fmt.Sprintf("    HomeDomain = [%s]\r\n", ths.HomeDomain)
	}
	if len(ths.InflationDest) > 0 {
		ret += fmt.Sprintf(" InflationDest = [%s]\r\n", ths.InflationDest)
	}
	if ths.SetFlags != nil && len(ths.SetFlags) > 0 {
		ret += fmt.Sprintf("      SetFlags = %s\r\n", ths.SetFlags)
	}
	if ths.ClearFlags != nil && len(ths.ClearFlags) > 0 {
		ret += fmt.Sprintf("    ClearFlags = %s\r\n", ths.ClearFlags)
	}
	return
}

/*
{
  "_links": {
    "effects": {
      "href": "/operations/3317677552570369/effects{?cursor,limit,order}",
      "templated": true
    },
    "precedes": {
      "href": "/operations?cursor=3317677552570369\u0026order=asc"
    },
    "self": {
      "href": "/operations/3317677552570369"
    },
    "succeeds": {
      "href": "/operations?cursor=3317677552570369\u0026order=desc"
    },
    "transaction": {
      "href": "/transactions/be63c2d5c010711b9946b0363b85a43d514edca9a691eec229fa2108359d2115"
    }
  },
  "home_domain": "www.ledgercn.com",
  "id": 3317677552570369,
  "inflation_dest": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
  "paging_token": "3317677552570369",
  "set_flags": [
    1,
    2
  ],
  "set_flags_s": [
    "auth_required_flag",
    "auth_revocable_flag"
  ],
  "signer_key": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
  "signer_weight": 255,
  "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
  "type": "set_options",
  "type_i": 5
},
*/

// DecodeBody 解码主体
func (ths *SubSetOptionsItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)

	home_domain, ok := b["home_domain"]
	if ok {
		ths.HomeDomain = home_domain.(string)
	}

	inflation_dest, ok := b["inflation_dest"]
	if ok {
		ths.InflationDest = inflation_dest.(string)
	}

	signer_key, ok := b["signer_key"]
	if ok {
		ths.SignerKey = signer_key.(string)
	}

	signer_weight, ok := b["signer_weight"]
	if ok {
		ths.SignerWeight = fmt.Sprintf("%d", (int64)(signer_weight.(float64)))
	}

	set_flags_s, ok := b["set_flags_s"]
	if ok {
		for _, v := range set_flags_s.([]interface{}) {
			if ths.SetFlags == nil {
				ths.SetFlags = make([]string, 0)
			}
			ths.SetFlags = append(ths.SetFlags, v.(string))
		}
	}

	clear_flags_s, ok := b["clear_flags_s"]
	if ok {
		for _, v := range clear_flags_s.([]interface{}) {
			if ths.ClearFlags == nil {
				ths.ClearFlags = make([]string, 0)
			}
			ths.ClearFlags = append(ths.ClearFlags, v.(string))
		}
	}
}

// ToString 返回字符串
func (ths *SubManageOfferItem) ToString() (ret string) {
	ret = ths.SubAccOperRecordItemBase.ToString()
	ret += fmt.Sprintf("        Amount = [%s]\r\n       OfferID = [%s]\r\n         Price = [%s]\r\n",
		ths.Amount, ths.OfferID, ths.Price)
	if ths.Buying != nil {
		ret += fmt.Sprintf("        Buying = [\r\n\t\t  Code : %s\r\n\t\t  Type : %s\r\n\t\t  Issuer : %s\r\n\t\t ]\r\n ",
			ths.Buying.Asset_Code, ths.Buying.Asset_Type, ths.Buying.Asset_Issuer)
	}
	if ths.Selling != nil {
		ret += fmt.Sprintf("      Selling = [\r\n\t\t  Code : %s\r\n\t\t  Type : %s\r\n\t\t  Issuer : %s\r\n\t\t ]\r\n ",
			ths.Selling.Asset_Code, ths.Selling.Asset_Type, ths.Selling.Asset_Issuer)
	}
	return
}

// DecodeBody 解码主体
func (ths *SubManageOfferItem) DecodeBody(b map[string]interface{}) {
	ths.SubAccOperRecordItemBase.DecodeBody(b)
	amount, ok := b["amount"]
	if ok {
		ths.Amount = amount.(string)
	}

	offer_id, ok := b["offer_id"]
	if ok {
		ths.OfferID = fmt.Sprintf("%d", (int64)(offer_id.(float64)))
	}

	price, ok := b["price"]
	if ok {
		ths.Price = price.(string)
	}

	buying_asset_type, ok_buy_type := b["buying_asset_type"]
	buying_asset_issuer, ok_buy_iusser := b["buying_asset_issuer"]
	buying_asset_code, ok_buy_code := b["buying_asset_code"]
	selling_asset_type, ok_sell_type := b["selling_asset_type"]
	selling_asset_issuer, ok_sell_iusser := b["selling_asset_issuer"]
	selling_asset_code, ok_sell_code := b["selling_asset_code"]

	if ok_buy_type {
		if ths.Buying == nil {
			ths.Buying = &AssetInfoBaseItem{}
		}
		ths.Buying.Asset_Type = buying_asset_type.(string)
		if ths.Buying.Asset_Type != "native" {
			if ok_buy_code {
				ths.Buying.Asset_Code = buying_asset_code.(string)
			}

			if ok_buy_iusser {
				ths.Buying.Asset_Issuer = buying_asset_issuer.(string)
			}
		} else {
			ths.Buying.Asset_Code = "XLM"

			if ok_sell_iusser {
				ths.Buying.Asset_Issuer = selling_asset_issuer.(string)
			}
		}
	}

	if ok_sell_type {
		if ths.Selling == nil {
			ths.Selling = &AssetInfoBaseItem{}
		}
		ths.Selling.Asset_Type = selling_asset_type.(string)
		if ths.Selling.Asset_Type != "native" {
			if ok_sell_code {
				ths.Selling.Asset_Code = selling_asset_code.(string)
			}

			if ok_sell_iusser {
				ths.Selling.Asset_Issuer = selling_asset_issuer.(string)
			}
		} else {
			ths.Selling.Asset_Code = "XLM"

			if ok_buy_iusser {
				ths.Selling.Asset_Issuer = buying_asset_issuer.(string)
			}
		}
	}
}

/*
{
  "_embedded": {
    "records": [
      {
        "_links": {
          "effects": {
            "href": "/operations/413278933094401/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=413278933094401\u0026order=asc"
          },
          "self": {
            "href": "/operations/413278933094401"
          },
          "succeeds": {
            "href": "/operations?cursor=413278933094401\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/dcef180a209b3dab35791a56b175c18a3a9ee1c57062f74a4a885a1b7a8b8067"
          }
        },
        "account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "funder": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "id": 413278933094401,
        "paging_token": "413278933094401",
        "source_account": "GBS43BF24ENNS3KPACUZVKK2VYPOZVBQO2CISGZ777RYGOPYC2FT6S3K",
        "starting_balance": "10000.0",
        "type": "create_account",
        "type_i": 0
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/477063492407297/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=477063492407297\u0026order=asc"
          },
          "self": {
            "href": "/operations/477063492407297"
          },
          "succeeds": {
            "href": "/operations?cursor=477063492407297\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/104af896c4a9e1fb4d5825626ff5da35eb106e6bb7eb61d97d79c618b59f4ec5"
          }
        },
        "amount": "1000.0",
        "asset_type": "native",
        "from": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "id": 477063492407297,
        "paging_token": "477063492407297",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "to": "GBRFZNZB3RDJHBWEUDGFMZEE6OTTZXHOGEQLBZL22RXW7VOH2NHOS4X6",
        "type": "payment",
        "type_i": 1
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/777827167244289/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=777827167244289\u0026order=asc"
          },
          "self": {
            "href": "/operations/777827167244289"
          },
          "succeeds": {
            "href": "/operations?cursor=777827167244289\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/791e3575cddec3e07ed52ef46fa134b9d7acbd0563cfc9ecc908db66017082a6"
          }
        },
        "asset_code": "USD",
        "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "asset_type": "credit_alphanum4",
        "id": 777827167244289,
        "limit": "922337203685.4775807",
        "paging_token": "777827167244289",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "type": "change_trust",
        "type_i": 6
      },
      {
        "_links": {
          "effects": {
            "href": "/operations/777758447767553/effects{?cursor,limit,order}",
            "templated": true
          },
          "precedes": {
            "href": "/operations?cursor=777758447767553\u0026order=asc"
          },
          "self": {
            "href": "/operations/777758447767553"
          },
          "succeeds": {
            "href": "/operations?cursor=777758447767553\u0026order=desc"
          },
          "transaction": {
            "href": "/transactions/973bed257adf83d4ffe4b9693a2ce7ffb91cbe5afaf4734bc1b7ef8f782f498b"
          }
        },
        "asset_code": "XLM",
        "asset_issuer": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "asset_type": "credit_alphanum4",
        "id": 777758447767553,
        "limit": "922337203685.4775807",
        "paging_token": "777758447767553",
        "source_account": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "trustee": "GAZWSWPDQTBHFIPBY4FEDFW2J6E2LE7SZHJWGDZO6Q63W7DBSRICO2KN",
        "trustor": "GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X",
        "type": "change_trust",
        "type_i": 6
      }
    ]
  },

  "_links": {
    "next": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=asc\u0026limit=10\u0026cursor=487208205160449"
    },
    "prev": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=desc\u0026limit=10\u0026cursor=413278933094401"
    },
    "self": {
      "href": "/accounts/GCR6QXX7IRIJVIM5WA5ASQ6MWDOEJNBW3V6RTC5NJXEMOLVTUVKZ725X/operations?order=asc\u0026limit=10\u0026cursor="
    }
  }
}
*/

// PutMapBody 设置主体
func (ths *StellarAccOperationQuary) PutMapBody(mbody map[string]interface{}) {
	_embedded, _embeddedok := mbody["_embedded"]
	if _embeddedok {
		records, recordsok := _embedded.(map[string]interface{})["records"]
		if recordsok {
			recordsSlice := records.([]interface{})
			length := len(recordsSlice)

			ths.Records = make([]SubAccOperQuaryRecordInterface, 0)

			for i := 0; i < length; i++ {
				subRecord := recordsSlice[i]
				ths.decodeSubrecord(subRecord.(map[string]interface{}))
			}
		}
	}

	_links, linkok := mbody["_links"]
	if linkok {
		ths.decodeCursor(_links.(map[string]interface{}))
	}

	if len(ths.Records) < 10 {
		ths.IsEnd = true
	}
}

func (ths *StellarAccOperationQuary) decodeSubrecord(itm map[string]interface{}) {
	stype, stypeok := itm["type"]
	if stypeok {
		switch stype {
		case "create_account":
			subitm := &SubCreateAccountItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		case "payment":
			subitm := &SubPaymentItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		case "change_trust":
			subitm := &SubChangeTrustItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		case "account_merge":
			subitm := &SubAccountMergeItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		case "set_options":
			subitm := &SubSetOptionsItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		case "manage_offer":
			subitm := &SubManageOfferItem{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		default:
			subitm := &SubAccOperRecordItemBase{}
			subitm.OpType = stype.(string)
			subitm.DecodeBody(itm)
			ths.Records = append(ths.Records, subitm)
		}
	}
}

func (ths *StellarAccOperationQuary) decodeCursor(b map[string]interface{}) {
	prev, prevok := b["next"]
	if prevok {
		href, _ := prev.(map[string]interface{})["href"]
		hrefurl := href.(string)
		reg := regexp.MustCompile(`cursor=[\d]*`)
		cStr := reg.FindString(hrefurl)
		if len(cStr) > len("cursor=") {
			ths.QuaryCursor = strings.Trim(cStr, "cursor=")
		} else {
			ths.QuaryCursor = ""
		}
	}
}

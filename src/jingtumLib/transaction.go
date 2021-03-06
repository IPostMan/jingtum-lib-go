/**
 * Transaction类主管POST请求，包括组装交易和交易参数。请求时需要提供密钥，且交易可以
 * 进行本地签名和服务器签名。目前支持服务器签名，本地签名支持主要的交易，还有部分参数
 * 不支持。所有的请求是异步的，会提供一个回调函数。每个回调函数有两个参数，一个是错误，
 * 另一个是结果。
 *
 * @FileName: transaction.go
 * @Auther : 13851485286
 * @Email : yangxuebo@yeah.net
 * @CreateTime: 2018-05-28 10:44:32
 * @UpdateTime: 2018-05-28 10:44:54
 */

package jingtumLib

import (

     "fmt"
     "encoding/json"
)

type MemoInfo struct {
    MemoType        interface{}
    MemoLen         interface{}
    MemoData        string
}

type TxData struct {
    Flags           uint32
    Fee             interface{}// = JTConfig.ReadInt("Config","fee", 10000)
    Account         string
    TransactionType interface{}
    SendMax         interface{}
    Memos           []MemoInfo
    Paths           interface{}
    SendMax         interface{}
    TransferRate    uint32
}

type Transaction struct {
    remote        Remote
    filter        Filter
    secret        string
    tx_json       *TxData
}

func NewTransaction(remote *Remote, filter Filter) (transaction *Transaction , err error) {
    transaction = new(Transaction)
    transaction.remote = remote
    transaction.tx_json = new(TxData)
    transaction.tx_json.Flags = 0
    transaction.tx_json.Fee = JTConfig.ReadInt("Config","fee", 10000);
    transaction.filter = filter
}

func (transaction *Transaction) ParseJson (jsonStr string) (err error) {

    err:=json.Unmarshal([]byte(jsonStr),&transaction.tx_json)
}

func (transaction *Transaction) GetAccount() (stirng) {
    return transaction.tx_json.Account
}

func (transaction *Transaction) GetTransactionType (interface{}) {
    return transaction.tx_json.TransactionType
}

func (transaction *Transaction) SetSecret (secret string) {
    transaction.secret = secret
}

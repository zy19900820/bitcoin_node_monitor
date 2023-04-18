package bitcoin_rpc

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

var LatestHeight int64

func GetLatestHight(host string, user string, password string) error {
	// 创建一个新的 RPC 客户端实例
	connCfg := &rpcclient.ConnConfig{
		Host:         host,
		User:         user,
		Pass:         password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Println("连接比特币全节点失败")
		return err
	}
	defer client.Shutdown()

	// 获取最新的区块高度
	LatestHeight, err = client.GetBlockCount()
	if err != nil {
		log.Println("获取区块高度失败")
		return err
	}
	return nil
}

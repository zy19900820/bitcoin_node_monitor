package three_bitcoin_rpc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Block struct {
	Height int64 `json:"height"`
}

var LatestHeight Block

func GetLatestHight() error {
	// 发送 HTTP 请求获取最新区块信息
	resp, err := http.Get("https://blockchain.info/latestblock")
	if err != nil {
		log.Println("发送 HTTP 请求失败")
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	// 解析 HTTP 响应，获取最新区块高度
	err = json.Unmarshal(body, &LatestHeight)
	if err != nil {
		log.Println("解析 HTTP 响应失败")
		return err
	}

	return nil
}

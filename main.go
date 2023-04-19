package main

import (
	"btc_node_monitor/bitcoin_rpc"
	"btc_node_monitor/conf"
	"btc_node_monitor/mail"
	"btc_node_monitor/three_bitcoin_rpc"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	err := conf.ReadConf()
	if err != nil {
		log.Println("err:", err)
		return
	}

	for {
		for {
			err = bitcoin_rpc.GetLatestHight(conf.Conf.Host, conf.Conf.Username, conf.Conf.Password)
			if err != nil {
				log.Println("err:", err)

				err = mail.SendGoMail(err.Error())
				if err != nil {
					log.Println("发送邮件失败:", err)
				}
				break
			}

			err = three_bitcoin_rpc.GetLatestHight()
			if err != nil {
				log.Println("err:", err)
				if strings.Contains(err.Error(), "unexpected end of JSON input") {
					break
				}

				err = mail.SendGoMail(err.Error())
				if err != nil {
					log.Println("发送邮件失败:", err)
				}
				break
			}
			log.Println("节点获取最新区块高度:", bitcoin_rpc.LatestHeight,
				"第三方获取区块最新高度:", three_bitcoin_rpc.LatestHeight.Height)

			abs := math.Abs(float64(bitcoin_rpc.LatestHeight - three_bitcoin_rpc.LatestHeight.Height))
			absString := strconv.FormatFloat(abs, 'f', 10, 64)
			if abs > 2 {
				log.Println("err:", "相差", absString, "个块")

				err = mail.SendGoMail(absString)
				if err != nil {
					log.Println("发送邮件失败:", err)
				}
				break
			}
			break
		}

		log.Println("sleep 120s")
		time.Sleep(120 * time.Second)
	}
}

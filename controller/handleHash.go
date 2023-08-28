package controller

import (
	"strings"

	"github.com/airstack/erc/client"
	"github.com/airstack/erc/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	client client.IClient
}

func NewController(client client.IClient) Controller {
	return Controller{
		client: client,
	}
}

func (c *Controller) GetTransactionReceiptsByHash(context *gin.Context) {
	txHash := context.Param("txhash")
	receipts, err := c.client.GetTransactionReceiptByHash(txHash)
	if err != nil {
		log.Error(err.Error())
	}
	var txLogWithTopics []client.TransactionLogs
	for _, txLog := range receipts.Result.Logs {
		if len(txLog.Topics) <= 0 {
			continue
		}
		topic0 := txLog.Topics[0]
		var txLogWithTopic client.TransactionLogs
		if strings.EqualFold(topic0, utils.TransferEvent) {
			if len(txLog.Topics) > 0 {
				txLogWithTopic.Topic0 = utils.NewNullString(txLog.Topics[0])
			}
			if len(txLog.Topics) > 1 {
				txLogWithTopic.Topic1 = utils.NewNullString(txLog.Topics[1])
			}
			if len(txLog.Topics) > 2 {
				txLogWithTopic.Topic2 = utils.NewNullString(txLog.Topics[2])
			}
			if len(txLog.Topics) > 3 {
				txLogWithTopic.Topic3 = utils.NewNullString(txLog.Topics[3])
			}
			txLogWithTopic.Data = txLog.Data
			txLogWithTopic.TokenAddress = txLog.Address
			txLogWithTopics = append(txLogWithTopics, txLogWithTopic)
		}
	}

	var erc20Transfers []client.Result
	var nftTransfers []client.Result
	for _, eachLog := range txLogWithTopics {
		if strings.EqualFold(eachLog.Topic0.String, utils.TransferEvent) {
			if len(eachLog.Topic3.String) > 0 {
				nftTransfers = append(nftTransfers, client.Result{
					Type: utils.ERC721,
					TransferData: client.TransferData{
						Token:    eachLog.TokenAddress,
						Sender:   utils.GetHexString(eachLog.Topic1.String),
						Receiver: utils.GetHexString(eachLog.Topic2.String),
						NftId:    utils.HexToIntNoError(eachLog.Topic3.String).String(),
					},
				})
				continue
			}
			erc20Transfers = append(erc20Transfers, client.Result{
				Type: utils.ERC20,
				TransferData: client.TransferData{
					Token:    eachLog.TokenAddress,
					Sender:   utils.GetHexString(eachLog.Topic1.String),
					Receiver: utils.GetHexString(eachLog.Topic2.String),
					Amount:   utils.HexToIntNoError(eachLog.Data).String(),
				},
			})
		}
	}
	if len(txLogWithTopics) == 0 {
		context.JSON(200, "Not an ERC20 or ERC721 transfer transaction")
		return
	}
	if len(nftTransfers) > 0 {
		context.JSON(200, nftTransfers)
		return
	}
	context.JSON(200, erc20Transfers)
	return
}

package task

import (
	"context"
	"log"
	"strconv"
	"time"

	"starfission_go_api/global"
	"starfission_go_api/internal/dao"
	"starfission_go_api/internal/model"
	"starfission_go_api/internal/model/starfission/order_repo"
	"starfission_go_api/internal/service"
	"starfission_go_api/pkg/rabbitmq"
)

func OrderNftMit() {

	log.Println("Task NftMit Start")

	svc := service.New(context.Background())

	err := svc.OrderNftMitTask()
	// TODO 对接 钉钉 或者 飞书 群通知

	if err != nil {
		global.Logger.Errorf(svc.Ctx(), "svc.OrderNftMit err: %v", err)
		return
	}

	return
}

func WalletGoodsGiven() {

	log.Println("Task NftMit Start")

	svc := service.New(context.Background())

	err := svc.WalletGoodsGivenTask()

	if err != nil {
		global.Logger.Errorf(svc.Ctx(), "svc.WalletGoodsGivenTask err: %v", err)
		return
	}

	return
}

func OrderStartClosed() {

	svc := service.New(context.Background())

	err := svc.OrderStartClosedTask()
	if err != nil {
		global.Logger.Errorf(svc.Ctx(), "svc.OrderNftMit err: %v", err)
		return
	}

	return
}

func OrderNftMitCrontab() {

	ticker := time.NewTicker(3600 * time.Second)
	defer ticker.Stop()

	queueExchange := global.RabbitMQ.NewQueueExchange(&rabbitmq.QueueExchange{
		//QuName: "order-nft-mit5",
		//RtKey:  "order-nft-mit5",
		QuName: global.MqOrderNftMitKey,
		RtKey:  global.MqOrderNftMitKey,
		ExName: "router",
		ExType: "direct",
	})

	todo := func(orderId int64) {
		//time.Sleep(500 * time.Millisecond)
		queueExchange.Publish(strconv.Itoa(int(orderId)))
	}

	orderNftMitCrontabD(todo)

	for range ticker.C {
		orderNftMitCrontabD(todo)
	}

	return
}

func orderNftMitCrontabD(todo func(orderId int64)) {

	listm, err := order_repo.NewQueryBuilder().
		WhereStatus(model.Egt, 2).
		WhereIsShiped(model.Eq, 0).
		Get()

	if err != nil {
		return
	}

	for _, m := range listm {
		if m.GoodsType != dao.GoodsTypeRealGiftBox {
			todo(m.OrderId)
		}
	}

	return
}

func WalletGoodsMetadataChange() {

	log.Println("Task WalletGoodsMetadataChange Start")

	svc := service.New(context.Background())

	err := svc.WalletGoodsMetadataChangeTask()
	if err != nil {
		global.Logger.Errorf(svc.Ctx(), "svc. WalletGoodsMetadataChange err: %v", err)
		return
	}
	return
}

func OrderNftAPPTask() {

	log.Println("Task OrderNftAPPTask Start")

	svc := service.New(context.Background())

	err := svc.OrderNftAPPTask()
	if err != nil {
		global.Logger.Errorf(svc.Ctx(), "svc. OrderNftAPPTask err: %v", err)
		return
	}
	return
}

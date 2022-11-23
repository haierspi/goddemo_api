package task

func Launcher() {

	//订单 NFT mit
	go OrderNftMit()
	go WalletGoodsMetadataChange()
	go OrderNftAPPTask()
	go OrderStartClosed()
	go OrderNftMitCrontab()
	go WalletGoodsGiven()
}

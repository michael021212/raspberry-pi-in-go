package main

import (
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3"
	"periph.io/x/host/v3/rpi"

	"02/sigutil"
)

func main() {
	// ハードウェアの初期化
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// GPIO9 (P1_21) のピンを取得
	pin := rpi.P1_21
	if pin == nil {
		log.Fatal("GPIO ピンが見つかりません")
	}

	// プルアップを有効にする
	pin.In(gpio.PullUp, gpio.NoEdge)

	// 無限ループでスイッチの状態を読み取り、0.5秒ごとに出力
	go func() {
		for {
			// 現在のピンの値を取得し、High なら 1、Lowなら 0 を出力
			level := pin.Read()
			if level == gpio.High {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// シグナルを待ち受けて終了時にログ出力処理
	sigutil.WaitForExitAndLog()
}


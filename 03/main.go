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

	count := 0

	go func() {
		// 0.1秒ごとにピンの入力を読み取る
		for {
			level := pin.Read()
			
			// Highの場合countを+1してログ出力
			if level == gpio.High {
				// チャタリング回避のため0.1秒待機する
				time.Sleep(100 * time.Millisecond)
				count++
				fmt.Printf("Count: %d\n", count)
				
				// Highの間は0.1秒ごとにピンの入力を読み取り続ける。Lowだったら次のループへ
				for level == gpio.High {
					time.Sleep(100 * time.Millisecond)
					level = pin.Read()
				}
			}
			
			// スイッチを押し続けるとLowなので0.1秒ごとにログが出力される			
			fmt.Println("Low")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// シグナルを待ち受けて終了時にログ出力処理
	sigutil.WaitForExitAndLog()
}


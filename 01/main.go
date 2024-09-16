package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "periph.io/x/conn/v3/gpio"
    "periph.io/x/host/v3"
    "periph.io/x/host/v3/rpi"
)

func main() {
    // ハードウェアの初期化
    if _, err := host.Init(); err != nil {
        log.Fatal(err)
    }

    // GPIO4 (P1_7) のピンを取得
    pin := rpi.P1_7
	if pin == nil {
	log.Fatal("GPIO ピンが見つかりません")
    }

    // 終了時に LED を消灯するためのチャネル
    stop := make(chan os.Signal, 1)

    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    // 500ミリ秒ごとに LED を点滅させる
    t := time.NewTicker(500 * time.Millisecond)
    go func() {
        for l := gpio.Low; ; l = !l {
            if err := pin.Out(l); err != nil {
                log.Fatal(err)
            }
            <-t.C
        }
    }()

    // SIGINT または SIGTERM シグナルを受け取るまで待機
    <-stop

    // プログラム終了時に LED を消灯
    log.Println("プログラムを終了します。LEDを消灯します...")
    pin.Out(gpio.Low) // LED を消灯
    t.Stop()          // タイマーを停止
}

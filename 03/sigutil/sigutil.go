package sigutil

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// シグナルを受け取り、終了時にログ出力を行う関数
func WaitForExitAndLog() {
	// 終了時にシグナルを受け取るためのチャネル
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// シグナルを受け取るまで待機
	<-stop

    log.Println("プログラムを終了します...")
}


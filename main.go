package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"org.donghyuns.com/graph/neo4j/configs"
	"org.donghyuns.com/graph/neo4j/network"
)

func main() {
	godotenv.Load(".env")

	configs.SetGlobalConfiguration()
	configs.SetDatabaseConfiguration()
	server := network.OpenServer()

	// 종료 신호를 받을 채널 생성
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 서버를 고루틴으로 실행
	go func() {
		log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
		log.Printf("[DEBUG] App Host %s", configs.GlobalConfiguration.AppPort)
		log.Printf("[START] Server Listening On: %s", configs.GlobalConfiguration.AppHost)
		log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Start Listening Error: %v", err)
		}
	}()

	// 종료 신호 대기
	<-quit
	log.Println("Received Shut Down Signal")

	// 셧다운 컨텍스트 설정 (예: 5초 제한)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 서버 그레이스풀 셧다운
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed Graceful Shutdown: %v", err)
	}

	log.Println("Server Has been Shutdown Gracefully")

	// dbCon, dbConErr := database.InitGraphConnect()

	// dbCon.CheckConnection()

	// 테스트용
	// person.CreateSinglePerson("Kim", "25", []string{"Typescript", "Golang", "Python"}, "Golang")
	// person.GetPerson("Kim")
}

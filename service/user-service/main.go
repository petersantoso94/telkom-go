package main

//go:generate protoc -I ../../pkg/proto/user-service --go_out=plugins=grpc:../../pkg/proto/user-service ../../pkg/proto/user-service/user.proto

import(
	"fmt"
	"net"
	"os"
	"time"

	"telkom-go/service/user-service/service"
	"telkom-go/service/user-service/user"
	pb "telkom-go/pkg/proto/user-service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

)

var (
	servicePort  = os.Getenv("SERVICE_PORT")
	databaseUser     = os.Getenv("DATABASE_USER")
	databasePassword = os.Getenv("DATABASE_PASSWORD")
	databaseHost     = os.Getenv("DATABASE_HOST")
	databasePort     = os.Getenv("DATABASE_PORT")
	databaseName     = os.Getenv("DATABASE_NAME")
)



func main(){
	log.SetFormatter(&log.JSONFormatter{})
	db := connectDatabase(databaseHost, databasePort, databaseName, databaseUser, databasePassword)
	a := user.CreateDAO(db)
	userServer := service.CreateServer(a)

	lis, err := net.Listen("tcp", ":"+servicePort)
	if err != nil {
		log.WithField("port", servicePort).WithError(err).Fatal("failed to listen")
	}

	log.Info("start service ...")
	s := grpc.NewServer()
	pb.RegisterUserServer(s, userServer)
	if err := s.Serve(lis); err != nil {
		log.WithError(err).Fatalf("failed to start service")
	}

}

const (
	MaxDatabaseRetry = 5
)

func connectDatabase(host, port, dbName, userName, pass string) *gorm.DB {
	var db *gorm.DB
	var err error
	for i := 0;i < MaxDatabaseRetry; i++ {
		db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, userName, dbName, pass))
		if err != nil {
			log.WithError(err).Warnf("failed to connect database, and sleep %d seconds", i)
			time.Sleep(time.Duration(i) * time.Second)
		}
	}
	if err != nil {
		log.WithFields(log.Fields{
			"host": host,
			"port": port,
			"user": userName,
			"name": dbName,
		}).WithError(err).Fatal("failed to connect database")
		return nil
	}
	db.LogMode(true)

	if err := db.AutoMigrate(&user.TelinUser{}).Error; err != nil {
		log.WithError(err).Fatal("failed to migrate table schema")
		return nil
	}
	return db
}

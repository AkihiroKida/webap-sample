package driver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-app/adapters/controllers"
	"go-app/adapters/gateways"
	"go-app/adapters/presenters"
	"go-app/usecases/user"

	_ "github.com/go-sql-driver/mysql"
)

func Serve(addr string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DATABASE"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	user := controllers.User{
		OutputFactory: presenters.NewUserOutputPort,
		InputFactory:  user.NewUserInputPort,
		RepoFactory:   gateways.NewUserRepository,
		Conn:          conn,
	}
	http.HandleFunc("/user/", user.GetUserByID)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

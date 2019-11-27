package main

import (
	_ "circleweb/conf"
	"circleweb/routes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
    r :=routes.Init()
    r.Use(gin.Recovery())
    srv:=&http.Server{
    	Addr:":8080",
    	Handler:r,
	}
    go func() {
    	log.Printf("listening port %s port",srv.Addr)
    	if err:=srv.ListenAndServe();err!=nil&&err!=http.ErrServerClosed{
    		log.Fatalf("listen :%s\n",err)
		}
	}()
    quit :=make(chan os.Signal,1)
    signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
    <-quit
    log.Print("Shutdown Server....\n")
    ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
    defer cancel()
    if err:=srv.Shutdown(ctx);err!=nil{
    	log.Fatal("Server Shutdown: ",err)
	}
    log.Print("Server exiting")

}


package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func main(){
    
    text := `Welcome
    get IST (+0530) current time at =>
        /time/ist
    UTC time =>
        /time/utc
    
    `

    r := gin.Default()
    r.GET("/" , func(c *gin.Context){
        
        c.String(200 ,text)

    })
    r.GET("/time/ist" , func(c *gin.Context){
        
        c.JSON(200 , gin.H{
            "ist" : time.Now().Format("15:04:05.000000 IST"),
        })


    })

    r.GET("/time/utc" , func(c *gin.Context){
        
        c.JSON(200 , gin.H{
            "utc" : time.Now().UTC().Format("15:04:05.000000"),
        })


    })
    
    fmt.Println("Starting server at http://localhost:8080/")
    r.Run()

}

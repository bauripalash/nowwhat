package pb

import (
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/joho/godotenv"
)


func noterr(e error){
    if e != nil{
        log.Fatalln(e)
    }
}


type Weather struct{
    tempc float64
    lat float64
    long float64
    status string
}


//return data => Temp in Celcius, Lat , Long , Weather
func GetWeather(location string) Weather {
    
    godotenv.Load(".env")
    var apikey = os.Getenv("OWM_API_KEY")

    w, err := owm.NewCurrent("C" , "en" , apikey)
    
    noterr(err)

    w.CurrentByName(location)
    
    result := Weather{ tempc : w.Main.Temp  , lat : w.GeoPos.Latitude , long : w.GeoPos.Longitude , status : w.Weather[0].Description}
    return result
    
    

}

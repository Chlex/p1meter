package main

import (
    "os"
    "unicode/utf8"

    "github.com/gin-gonic/gin"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.New()
    router.GET("/p1meterip", getP1meterIP)

    router.Run("0.0.0.0:3001")
}

// getAlbums responds with the list of all albums as JSON.
func getP1meterIP(c *gin.Context) {

	var ip string

    dat, err := os.ReadFile("./p1meter.ip")
    check(err)
	ipf := string(dat)

	if len(ipf) < 10 {
		c.Data(200, "application/json; charset=utf-8", []byte("Error"))
		return
	}

	for i, w := 0, 0; i < len(ipf); i += w {
		runeValue, width := utf8.DecodeRuneInString(ipf[i:])
		w = width
		if ( (string(runeValue) != "(")  && (string(runeValue) != ")") ) {
			ip += string(runeValue)
		}
	}
	c.Data(200, "application/json; charset=utf-8", []byte(ip))
}

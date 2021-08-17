package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/giansalex/coingecko-notify/api"
)

var (
	cointPtr    = flag.String("coin", "", "Coin ID")
	upPtr       = flag.Float64("up", 0, "Price up to notify")
	downPtr     = flag.Float64("down", 0, "Price down to notify")
	mailHostPtr = flag.String("mail.host", "", "SMTP Host")
	mailPortPtr = flag.Int("mail.port", 587, "SMTP Port")
	mailUserPtr = flag.String("mail.user", "", "SMTP User")
	mailPassPtr = flag.String("mail.pass", "", "SMTP Password")
	mailFromPtr = flag.String("mail.from", "", "email sender")
	mailToPtr   = flag.String("mail.to", "", "email receptor")
)

const coingeckoURL = "https://api.coingecko.com/api/v3/"

func main() {
	flag.Parse()

	if *cointPtr == "" {
		log.Fatal("Coin flag is required")
	}

	client := &http.Client{}
	coinGecko := api.NewCoinGeckoAPI(client, coingeckoURL)

	price, err := coinGecko.GetSimplePrice(*cointPtr)
	if err != nil {
		log.Fatal(err)
	}

	msg := ""
	if *upPtr != 0 && price > *upPtr {
		msg = fmt.Sprintf("Price Up %.6f > %.6f", price, *upPtr)
	}

	if *downPtr != 0 && price < *downPtr {
		msg = fmt.Sprintf("Price Up %.6f > %.6f", price, *upPtr)
	}

	if msg == "" {
		return
	}

	noEmail := *mailHostPtr == "" || *mailUserPtr == "" || *mailPassPtr == "" || *mailFromPtr == "" || *mailToPtr == ""
	if noEmail {
		fmt.Println(msg)
		return
	}

	mailer := newMail(*mailHostPtr, *mailPortPtr, *mailUserPtr, *mailPassPtr)

	mailer.notify(*mailFromPtr, *mailToPtr, msg)
}

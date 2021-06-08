package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/TwinProduction/go-color"
	"github.com/martinlindhe/notify"
)

var client = &http.Client{
	Timeout: time.Second * 5,
}

func main() {
	for {
		showMenu()
		comando := readCommand()

		switch comando {
		case 1:
			startMonitoring()
			fmt.Println("monitoring started.")
		case 2:
			fmt.Println("Going out...")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		default:
			fmt.Println("Option invalid")
			os.Exit(-1)
		}
	}
}

func showMenu() {
	fmt.Println("Enter one of the options below.")
	fmt.Println(color.Ize(color.Cyan, "1 - Monitor website"))
	fmt.Println(color.Ize(color.Red, "2 - Exit"))
}

func readCommand() int {
	var command int

	fmt.Print(color.Ize(color.Green, "- Option: "))
	fmt.Scan(&command)
	fmt.Println("The chosen command was:", command)
	fmt.Println("")
	return command
}

func siteMonitor() string {
	var site string

	print(color.Green + "Site to monitor: " + site + color.Reset)

	fmt.Scan(&site)
	//command to clean terminal
	fmt.Print("\033[H\033[2J")
	return site
}

func testSite(site string) {
	req, _ := http.NewRequest("GET", site, nil)

	response, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			fmt.Println("Timeout connecting to", site)
			return
		}
		fmt.Printf("Error connecting to '%s': %s", site, err)
		return
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println(time.Now().Format("01-02-2006 15:04:05 Monday"))
		print(color.Green + "Website: " + site + " - Online \n\n " + color.Reset)

	} else {
		print(color.Red+"The website: "+site+" has a problem. Status Code:", response.StatusCode, "\n\n"+color.Reset)
		notify.Alert("app name", "ALERT!", "Site failed: "+site, "path/to/icon.png")
		os.Exit(0)
	}
}

func startMonitoring() {
	site := siteMonitor()
	var interval string
	var setTimeDuration int
	fmt.Print(color.Ize(color.Green, "Hour | Minute | Second: "))
	fmt.Scan(&interval)
	interval = strings.ToLower(interval)
	fmt.Println(interval)

	switch interval {
	case "hour":
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			testSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Hour)
		}
	case "minute":
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			testSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Minute)
		}
	case "second":
		fmt.Print(color.Ize(color.Green, "Time interval: "))
		fmt.Scan(&setTimeDuration)
		for {
			testSite(site)
			time.Sleep(time.Duration(setTimeDuration) * time.Second)
		}
	default:
		fmt.Println("Option invalid")
		os.Exit(0)
	}
}

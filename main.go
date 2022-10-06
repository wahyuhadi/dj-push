package main

import (
	"dj-push/models"
	"dj-push/services"
	"flag"
	"log"
	"os"
)

var (
	active            = flag.Bool("active", true, "for status finding")
	verified          = flag.Bool("verified", true, "for verifying finding")
	scan_type         = flag.Int64("scan_type", 1, "For tools scan use")
	engagement        = flag.String("eg", "1", "For engagement from product")
	close_old_finding = flag.Bool("closeold", false, "Close Old Finding")
	push_to_jira      = flag.Bool("pushjira", false, "Push to Jira ticket")
	token             = flag.String("api_key", os.Getenv("apikey"), "Api key")
	url               = flag.String("url", os.Getenv("djhost"), "URI api to push data")
	file              = flag.String("file", "finding.json", "Location security finding file ")
	list_type         = flag.Bool("list", false, "List scan type")
)

func parse_config() *models.Data {
	flag.Parse()
	return &models.Data{
		Active:           *active,
		Verifed:          *verified,
		Engagement:       *engagement,
		CloseOldFindings: *close_old_finding,
		PushToJira:       *push_to_jira,
		Token:            *token,
		ScanType:         *scan_type,
		URI:              *url,
		File:             *file,
		ListScan:         *list_type,
	}

}
func main() {
	config := parse_config()
	if config.ListScan {
		services.GetList()
		os.Exit(1)
	}

	ok := services.Push(config)
	if ok {
		log.Println("Success uplaod finding")
	}

}

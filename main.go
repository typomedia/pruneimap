package main

import (
	"fmt"
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	server := pflag.StringP("server", "s", "", "IMAP server address")
	port := pflag.StringP("port", "p", "993", "IMAP server port")
	username := pflag.StringP("user", "u", "", "IMAP user name")
	password := pflag.StringP("pass", "w", "", "IMAP user password")
	dryrun := pflag.BoolP("dry", "d", false, "dry run")
	pflag.Parse()

	fmt.Print("pruneimap 1.0.0 <philipp@typo.media>\n")

	// write to log file
	logfile, err := os.OpenFile("pruneimap.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()

	// write to stdout and log file
	log.SetOutput(io.MultiWriter(os.Stdout, logfile))

	address := *server + ":" + *port
	client, err := imapclient.DialTLS(address, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Logout()

	client.Login(*username, *password)

	delete(client, "", "*", *dryrun)
}

func delete(client *imapclient.Client, ref, pattern string, dryrun bool) {
	mailboxes, err := client.List(ref, pattern, nil).Collect()
	if err != nil {
		log.Fatal(err)
	}

	for _, mailbox := range mailboxes {
		if ignore(mailbox.Mailbox) {
			log.Printf("IGNORED: %s\n", mailbox.Mailbox)
			continue
		}

		mbox, err := client.Select(mailbox.Mailbox, nil).Wait()
		if err != nil {
			log.Fatal(err)
		}

		// check if mailbox is empty
		if mbox.NumMessages == 0 {
			// check if mailbox has subfolders
			children, err := client.List(mailbox.Mailbox, "*", nil).Collect()
			if err != nil {
				log.Fatal(err)
			}

			// if mailbox has no subfolders, delete it
			if len(children) <= 1 { // 1 because of the current mailbox itself
				if !dryrun {
					err := client.Delete(mailbox.Mailbox).Wait()
					if err != nil {
						log.Printf("ERROR: %s\n", err)
					}
					log.Printf("DELETED: %s\n", mailbox.Mailbox)
				} else {
					log.Printf("DRYRUN: %s\n", mailbox.Mailbox)
				}
			}
		}
	}
}

func ignore(name string) bool {
	// List of mailbox names to ignore
	ignoreList := []string{"INBOX", "Sent", "Trash", "Drafts", "Junk", "Notes", "Spam"}

	for _, ignored := range ignoreList {
		if strings.EqualFold(name, ignored) {
			return true
		}
	}
	return false
}

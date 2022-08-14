package application

import (
	"flag"
	"strings"

	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
)

type App struct {
	ServerToken string
	Mode bool
	Tag  string
	Url  string
}

func New() App {
	mode := flag.Bool("ci", false, "")
	flag.Parse()

	// get git tag for versioning
	r, err := git.PlainOpen("")

	if err != nil {
		log.Fatal(err)
	}
	hashCommit, err := r.Head()
	if err != nil {
		log.Fatal("head", err)
	}
	tag := strings.Split(hashCommit.String(), " ")[0]

	return App{
		Mode: *mode,
		Tag:  tag,
		Url: "https://secret.coin.com/v1/",
		ServerToken: "test token",
	}
}

func (a *App) Start() {
	if a.Mode {
		if err := a.CiMode(); err != nil {
			log.Fatal(err)
		}

		return
	}

	if err := a.RegularWorkflow(); err != nil {
		log.Fatal(err)
	}

	return
}

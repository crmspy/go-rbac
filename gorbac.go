package main

import (
  "os"
  "github.com/urfave/cli"
  "sort"
)

import (
  Seed "github.com/crmspy/go-rbac/seeder"
  log "github.com/Sirupsen/logrus"
)

func main() {

  app := cli.NewApp()
  app.Name = "Jimatt Console"
  app.HelpName = "jimatt"
  app.Usage = "Jimatt Console Self Service"
  app.Authors = []cli.Author{
    cli.Author{
      Name:  "Nurul Hidayat",
      Email: "crmspy@gmail.com",
    },
  }
  app.Copyright = `(c) 2019 ZULLAB TEAM - PT.ZUllab
  
  This application is for public usage :)`
  app.Version = "0.0.1"
  app.Commands = []cli.Command{
    {
      Name:    "seed",
      Usage:   "Seed data",
      Action:  func(c *cli.Context) error {
        Seed.AppUser()
        return nil
      },
    },
  }
  

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

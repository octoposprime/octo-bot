package main

import (
	"sync"

	"github.com/golobby/container/v3"

	pa_command "github.com/octoposprime/octo-bot/internal/application/presentation/adapter/command"
	pa_query "github.com/octoposprime/octo-bot/internal/application/presentation/adapter/query"
	as "github.com/octoposprime/octo-bot/internal/application/service"
	ds "github.com/octoposprime/octo-bot/internal/domain/service"
	ia_repo "github.com/octoposprime/octo-bot/pkg/infrastructure/adapter/repository"
	pc_dc "github.com/octoposprime/octo-bot/pkg/presentation/controller/dc"
	tconfig "github.com/octoposprime/octo-bot/tool/config"
	tseed "github.com/octoposprime/octo-bot/tool/config"
)

var dbConfig tconfig.DbConfig
var dcConfig tconfig.DcConfig
var seedConfig tseed.SeedConfig

func main() {
	dbConfig.ReadConfig()
	dcConfig.ReadConfig()
	seedConfig.ReadConfig()
	var err error

	/*fmt.Println("Starting User Service...")
	dbClient, err := tgorm.NewGormClient(tgorm.PostgresGormClient).Connect(dbConfig.PostgresDb.Host, dbConfig.PostgresDb.Port, dbConfig.PostgresDb.UserName, dbConfig.PostgresDb.Password, dbConfig.PostgresDb.Database)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")*/

	cont := container.New()

	//Domain OctoBot Service
	err = cont.Singleton(func() *ds.Service {
		return ds.NewService()
	})
	if err != nil {
		panic(err)
	}

	//Infrastructure OctoBot Db Repository Adapter
	err = cont.Singleton(func() ia_repo.DbAdapter {
		return ia_repo.NewDbAdapter(nil)
	})
	if err != nil {
		panic(err)
	}

	//Application OctoBot Service
	err = cont.Singleton(func(s *ds.Service, d ia_repo.DbAdapter) *as.Service {
		return as.NewService(s, &d)
	})
	if err != nil {
		panic(err)
	}

	//Application OctoBot Query Adapter
	err = cont.Singleton(func(s *as.Service) pa_query.QueryAdapter {
		return pa_query.NewQueryAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	//Application OctoBot Command Adapter
	err = cont.Singleton(func(s *as.Service) pa_command.CommandAdapter {
		return pa_command.NewCommandAdapter(s)
	})
	if err != nil {
		panic(err)
	}

	var queryHandler pa_query.QueryAdapter
	err = cont.Resolve(&queryHandler)
	if err != nil {
		panic(err)
	}

	var commandHandler pa_command.CommandAdapter
	err = cont.Resolve(&commandHandler)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go pc_dc.NewDcAPI(queryHandler, commandHandler).Start(dcConfig)
	wg.Wait()

}

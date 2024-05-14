package infrastructure

import (
	"context"
	"strings"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// Log is the default log function
func Log(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error) {
	return &pb_logging.LoggingResult{}, nil
}

// GetUsersByFilter returns the users that match the given filter.
func (a DbAdapter) GetUsersByFilter(ctx context.Context, filter string) ([]string, error) {

	var users []string
	for _, user := range Users {
		if filter == "" || strings.Contains(user, filter) {
			users = append(users, user)
		}
	}
	return users, nil

}

// Users contains 100 turkish names
var Users = []string{
	"Ahmet",
	"Ali",
	"Arda",
	"Baran",
	"Barış",
	"Berat",
	"Berke",
	"Berker",
	"Burak",
	"Can",
	"Deniz",
	"Emir",
	"Emirhan",
	"Emre",
	"Enes",
	"Eray",
	"Eren",
	"Erhan",
	"Erkan",
	"Ertuğrul",
	"Faruk",
	"Furkan",
	"Gökhan",
	"Görkem",
	"Hakan",
	"Hasan",
	"İbrahim",
	"İsmail",
	"Kaan",
	"Kadir",
	"Kerem",
	"Kerim",
	"Koray",
	"Kuzey",
	"Mahir",
	"Mahmut",
	"Mehmet",
	"Melih",
	"Metehan",
	"Murat",
	"Mustafa",
	"Oğuz",
	"Okan",
	"Onur",
	"Orhan",
	"Ömer",
	"Özgür",
	"Ramazan",
	"Rıdvan",
	"Salih",
	"Sami",
	"Samet",
	"Selim",
	"Serkan",
	"Sertan",
	"Soner",
	"Şafak",
	"Şahin",
	"Şeref",
	"Şerif",
	"Şükrü",
	"Taha",
	"Talha",
	"Tayfun",
	"Taylan",
	"Tolga",
	"Uğur",
	"Umut",
	"Volkan",
	"Yasin",
	"Zafer",
	"Zeki",
	"Ziya",
	"Zübeyr",
	"Zülküf",
	"Zümrüt",
}

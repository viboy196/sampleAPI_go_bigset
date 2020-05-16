package models

import (
	appconfig "SampleAPI_Bigset/config"
	"log"
	"sync"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

var (
	bigsetIf StringBigsetService.StringBigsetServiceIf
	Once     sync.Once
)

// var bigsetIf = bigset.GetBigSet("trustkey/socialnetwork/dev/configapiservice/string", "127.0.0.1", "18990")

func InitBigSetIf() {
	log.Println(appconfig.Config, "appconfig.Config")
	Once.Do(func() {
		log.Println(appconfig.Config, "appconfig.Config")
		// bigsetIf = bigset.GetBigSet(trustkey/socialnetwork/dev/configapiservice/string", "127.0.0.1", "18990)
		bigsetIf = StringBigsetService.NewStringBigsetServiceModel(appconfig.Config.SourceMappingNewSID,
			appconfig.Config.EtcdServerEndpoints,
			GoEndpointBackendManager.EndPoint{
				Host:      appconfig.Config.SourceMappingNewHost,
				Port:      appconfig.Config.SourceMappingNewPort,
				ServiceID: appconfig.Config.SourceMappingNewProtocol,
			})
		log.Println(bigsetIf, "bigsetIf")
	})
}

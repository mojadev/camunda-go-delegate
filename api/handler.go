package api

import remoteApi "github.com/mojadev/camunda-go-delegate/internal/api"

func RegisterHandlers(router remoteApi.EchoRouter, si remoteApi.ServerInterface) {
	remoteApi.RegisterHandlers(router, si)
}


package endpoints

import "github.com/mcluseau/kube-localnet-api/pkg/endpoints"

type Server struct {
	InstanceID uint64
	Correlator *endpoints.Correlator
}

package loadbalancer

import (
	"github.com/spf13/cobra"
)

type Interface interface {
	GetLoadbalancer()
	GetByIdLoadbalancer()
	ApplyLoadbalancer()
	PrintLoadbalancer()
	ListLoadbalancer(*cobra.Command, []string)
	DeleteLoadbalancer(*cobra.Command, []string)
	CreateLoadbalancer(*cobra.Command, []string)
}

func NewLoadbalancerInterface(o *Opts) Interface {
	return o
}

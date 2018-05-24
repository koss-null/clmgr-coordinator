package cluster

import "time"

type NoQuorumPolicyType string

const (
	nqp_ignore  NoQuorumPolicyType = "ignore"
	nqp_stop                       = "stop"
	nqp_suicide                    = "suicide"
)

type FenctActionType string

const (
	fa_reboot   FenctActionType = "reboot"
	fa_poweroff                 = "off"
)

type PlacementStrategyType string

const (
	ps_default     PlacementStrategyType = "default"
	ps_utilization                       = "utilization"
	ps_balance                           = "balance"
	ps_minimal                           = "minimal"
)

type HealthStrategyType string

const (
	hs_none             HealthStrategyType = "none"
	hs_migrate_on_green                    = "migrate_on_green"
	hs_migrate_on_red                      = "migrate_on_red"
)

type Config struct {
	NoQuorumPolicy    NoQuorumPolicyType    `json:"no-quorum-policy"`
	Symmetric         bool                  `json:"symmetric"`
	FenceEnabled      bool                  `json:"fence-enabled"`
	FenceAction       FenctActionType       `json:"fence-action"`
	FenceTimeout      time.Duration         `json:"fence-timeout"`
	ConcurrentFencing bool                  `json:"concurrent-fencing"`
	PlacementStrategy PlacementStrategyType `json:"placement-strategy"`
	HealthStrategy    HealthStrategyType    `json:"health-strategy"`
}

func (c *Config) Check() bool {
	// todo: implement
	return true
}

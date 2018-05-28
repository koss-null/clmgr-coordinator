package cluster

import "github.com/google/logger"

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
	FenceTimeout      int                   `json:"fence-timeout"`
	ConcurrentFencing bool                  `json:"concurrent-fencing"`
	PlacementStrategy PlacementStrategyType `json:"placement-strategy"`
	HealthStrategy    HealthStrategyType    `json:"health-strategy"`
}

func DefaultConfig() Config {
	return Config {
		nqp_ignore,
		false,
		false,
		fa_poweroff,
		60,
		false,
		ps_default,
		hs_none,
	}
}

func (c *Config) Check() bool {
	switch c.NoQuorumPolicy {
	case nqp_ignore:
		break
	case nqp_stop:
		break
	case nqp_suicide:
		break
	default:
		logger.Error("NoQuorumPolicy bad value")
		return false
	}

	switch c.FenceAction {
	case fa_reboot:
		break
	case fa_poweroff:
		break
	default:
		logger.Error("FenceAction bad value")
		return false
	}

	switch c.PlacementStrategy {
	case ps_default:
		break
	case ps_utilization:
		break
	case ps_balance:
		break
	case ps_minimal:
		break
	default:
		logger.Error("PlacementStrategy bad value")
		return false
	}

	switch c.HealthStrategy {
	case hs_none:
		break
	case hs_migrate_on_green:
		break
	case hs_migrate_on_red:
		break
	default:
		logger.Error("HealthStrategy bad value")
		return false
	}

	return true
}

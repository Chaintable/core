package hermes

import _ "embed"

// contract codes for Pigeon upgrade
var (
	//go:embed pigeon/BTCAgentContract
	PigeonBTCAgentContract string
	//go:embed pigeon/BTCStakeContract
	PigeonBTCStakeContract string
	//go:embed pigeon/BurnContract
	PigeonBurnContract string
	//go:embed pigeon/CandidateHubContract
	PigeonCandidateHubContract string
	//go:embed pigeon/ChannelContract
	PigeonChannelContract string
	//go:embed pigeon/CoreAgentContract
	PigeonCoreAgentContract string
	//go:embed pigeon/FeeMarketContract
	PigeonFeeMarketContract string
	//go:embed pigeon/FoundationContract
	PigeonFoundationContract string
	//go:embed pigeon/GovHubContract
	PigeonGovHubContract string
	//go:embed pigeon/HashAgentContract
	PigeonHashAgentContract string
	//go:embed pigeon/LightClientContract
	PigeonLightClientContract string
	//go:embed pigeon/PledgeCandidateContract
	PigeonPledgeCandidateContract string
	//go:embed pigeon/RelayerHubContract
	PigeonRelayerHubContract string
	//go:embed pigeon/SlashContract
	PigeonSlashContract string
	//go:embed pigeon/StakeHubContract
	PigeonStakeHubContract string
	//go:embed pigeon/SystemRewardContract
	PigeonSystemRewardContract string
	//go:embed pigeon/ValidatorContract
	PigeonValidatorContract string
)

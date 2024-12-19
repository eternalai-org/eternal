package setting

import (
	"eternal/pkg/logger"
	"eternal/pkg/third-party/lighthouse"

	erc20 "eternal/internal/contract/erc20/usecase"

	"github.com/spf13/viper"
)

type Setting struct{}

func NewSetting() *Setting {
	s := &Setting{}
	_ = lighthouse.NewLighthouse(viper.GetString("LIGHTHOUSE_API_KEY"))

	// repos
	erc20Usecase, err := erc20.NewContractErc20Usecase(
		viper.GetString("CONTRACT_RPC"),
		viper.GetString("CONTRACT_ERC20_ADDRESS"),
	)
	if err != nil {
		logger.AtLog.Error(err)
	}
	_ = erc20Usecase

	return s
}

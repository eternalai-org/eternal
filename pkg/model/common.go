package model

type DepositWalletInfoReponse struct {
	Address string `json:"address"`
	MintFee int64  `json:"mint_fee"`
}

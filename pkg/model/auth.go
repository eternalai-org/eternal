package model

type Web3AuthChallengeRequest struct {
	Code    string `json:"code" query:"code"`
	Address string `json:"address" query:"address"`
}

type Web3AuthVerifyRequest struct {
	Signature string `json:"signature" query:"signature"`
	Address   string `json:"address" query:"address"`
}

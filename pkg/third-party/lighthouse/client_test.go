package lighthouse

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	ctx    context.Context
	client ILighthouse
}

// func (s *Suite) TestFilesUploaded() {
// 	res, err := s.client.FilesUploaded(s.ctx)
// 	require.NoError(s.T(), err)
// 	require.NotEqual(s.T(), 0, len(res))
// }

func (s *Suite) TestDetailFile() {
	res, err := s.client.FileInfo(s.ctx, "bafybeibhe4u65zrcjpnrog5qsqjbxu2v6ebugj7wydjjhlcijuibccqsay")
	spew.Dump(res)
	spew.Dump(err)
	require.NoError(s.T(), err)
	// require.Equal(s.T(), "bafkreibc5kdnza5w7evxgyhduzduugnl36ounobf2ncvv2elhhv7qakewy", res.Cid)
}

// func (s *Suite) TestDownloadFileByCid() {
// 	res, err := s.client.DownloadFileByCid(s.ctx, "bafkreigcb3ahat56miowlxy7l7jtrdjoxdyx3hblbjhuhvew33ptuept3y", "")
// 	require.NoError(s.T(), err)
// 	err = utils.CreateFile("test.html", res)
// 	require.NoError(s.T(), err)
// }

// func (s *Suite) TestUploadFile() {
// 	byte, err := utils.ReadFile("style.css")
// 	require.NoError(s.T(), err)

// 	res, err := s.client.UploadDataFile(s.ctx, "style1.css", byte)
// 	require.NoError(s.T(), err)
// 	spew.Dump(res.Hash)
// }

// func (s *Suite) TestUploadFileByUrl() {
// 	res, err := s.client.UploadDataFileByUrl(s.ctx, "https://s3.coinmarketcap.com/generated/sparklines/web/7d/2781/6210.svg")
// 	require.NoError(s.T(), err)
// 	spew.Dump(res.Hash)
// }

func TestRunLighthouse(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	s.ctx = context.Background()
	s.client = NewLighthouse("633649a6.5a3c842aaa654bcdb8bd7daf96e8c095")
}

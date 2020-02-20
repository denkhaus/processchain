package processchain

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/denhaus/processchain/modules/httprequest"
	"github.com/stretchr/testify/suite"
)

type commonTest struct {
	suite.Suite
}

func (suite *commonTest) SetupTest() {

}

func (suite *commonTest) TearDownTest() {

}

func (suite *commonTest) TestHttpRequest() {
	HttpRequest("https://httpbin.org/anything").WithOptions(
		httprequest.Timeout(10*time.Time),
		httprequest.Accept("application/json"),
	).Get().Execute().ResultReader(func(reader io.Reader) (interface{}, error) {
		return json.NewDecoder(reader).Decode()
	}).Catch(func(err error) {
		suite.FailNow("error occured", err.Error())
	}).Then(func(res interface{}) {
		fmt.Println("result: ", res)
	})

	// suite.True(res.(bool), "return value")
}

func TestCommon(t *testing.T) {
	testSuite := new(commonTest)
	suite.Run(t, testSuite)
}

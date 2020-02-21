package processchain

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/denkhaus/processchain/modules/httprequest"
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
	err := HttpRequest("https://httpbin.org/anything").WithOptions(
		httprequest.Timeout(10*time.Second),
		httprequest.Accept("application/json"),
	).Get().ReadResult(func(reader io.Reader) (interface{}, error) {
		var result interface{}
		return result, json.NewDecoder(reader).Decode(&result)
	}).Then(func(res interface{}) (interface{}, error) {
		fmt.Println("result: ", res)
		return nil, nil
	}).Catch(func(err error) {
		suite.FailNow("error occured", err.Error())
	}).Execute()

	suite.NoError(err, "return value")
}

func TestCommon(t *testing.T) {
	testSuite := new(commonTest)
	suite.Run(t, testSuite)
}

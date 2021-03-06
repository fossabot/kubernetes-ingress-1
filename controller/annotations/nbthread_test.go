package annotations

func (suite *AnnotationSuite) TestNbthreadUpdate() {
	test := "8"
	a := &nbthread{}
	if suite.NoError(a.Parse(test)) {
		suite.Equal(RELOAD, a.Update(suite.client))
		result, _ := suite.client.GlobalWriteConfig("global", "nbthread")
		suite.Equal("nbthread 8", result)
	}
}

func (suite *AnnotationSuite) TestNbthreadFail() {
	test := "garbage"
	a := &nbthread{}
	err := a.Parse(test)
	suite.T().Log(err)
	suite.Error(err)
}

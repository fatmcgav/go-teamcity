package teamcity_test

import (
	"testing"

	teamcity "github.com/cvbarros/go-teamcity-sdk/pkg/teamcity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTrigger_Constructor(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	actual := teamcity.NewVcsTrigger("+:*", "")

	require.NotNil(actual)
	assert.Equal("vcsTrigger", actual.Type)
	require.NotEmpty(actual.Properties)
	props := actual.Properties.Map()

	assert.Contains(props, "triggerRules")
	assert.NotContains(props, "branchFilter")
	assert.Equal(props["quietPeriodMode"], "DO_NOT_USE")
	assert.Equal(props["enableQueueOptimization"], "true")
}

func TestNewTrigger_ForBuildConfiguration(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	client := setup()

	bt := createTestBuildTypeWithName(t, client, "BuildTriggerProject", "BuildRelease", true)

	sut := client.TriggerService(bt.ID)
	nt := teamcity.NewVcsTrigger("+:*", "")

	_, err := sut.AddTrigger(nt)

	require.Nil(err)

	bt, _ = client.BuildTypes.GetById(bt.ID)

	assert.Equal(int32(1), bt.Triggers.Count)
	actual := bt.Triggers.Items[0]

	assert.NotEmpty(actual.ID)
	assert.Equal("vcsTrigger", actual.Type)
	assert.NotEmpty(actual.Properties)

	cleanUpProject(t, client, bt.ProjectID)
}

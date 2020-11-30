package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Mutables(t *testing.T) {

	testProperty := NewProperty(NewID("ID"), NewFact(NewStringData("Data")))
	testProperties := NewProperties(testProperty)
	testMutables := NewMutables(testProperties)
	require.Equal(t, mutables{Properties: testProperties}, testMutables)
	require.Equal(t, testProperties, testMutables.Get())
	mutatedTestProperty := NewProperty(NewID("ID"), NewFact(NewStringData("Data2")))
	require.Equal(t, mutables{Properties: NewProperties(mutatedTestProperty)}, testMutables.Mutate(mutatedTestProperty))

}
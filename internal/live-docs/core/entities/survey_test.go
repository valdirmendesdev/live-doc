package entities_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/valdirmendesdev/live-doc/internal/live-docs/core/entities"
)

func Test_SurveyTableName(t *testing.T) {
	s := entities.NewSurvey()
	require.Equal(t, "surveys", s.TableName())
}

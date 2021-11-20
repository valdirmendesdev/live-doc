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

func Test_NewSurvey(t *testing.T) {
	s := entities.NewSurvey()
	require.NotNil(t, s)
	require.NotNil(t, s.ID)
	require.False(t, s.CreatedAt.IsZero())
	require.False(t, s.UpdatedAt.IsZero())
}

func Test_SurveyIsValid(t *testing.T) {
	s := entities.NewSurvey()
	isValid, err := s.IsValid()
	require.False(t, isValid)
	require.NotNil(t, err)

	s.MainProjectLanguage = "en"
	isValid, err = s.IsValid()
	require.True(t, isValid)
	require.Nil(t, err)

}

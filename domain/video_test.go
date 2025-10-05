package domain_test

import (
	"testing"
	"time"

	"github.com/pbitts/codeflix-video-encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNOtAUuid(t *testing.T) {

	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIValidation(t *testing.T) {

	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceId = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}

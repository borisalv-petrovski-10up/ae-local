package repository_test

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"

	"github.com/stretchr/testify/suite"

	"github.com/borisalv-petrovski-10up/ae-local/repository"
)

const DatastoreEmulatorHost = "0.0.0.0:8081"

type repositoryTestSuite struct {
	suite.Suite
	repository      repository.Repository
	datastoreClient *datastore.Client

	identifier string
}

func (s *repositoryTestSuite) SetupTest() {
	if addr := os.Getenv("DATASTORE_EMULATOR_HOST"); addr == "" {
		_ = os.Setenv("DATASTORE_EMULATOR_HOST", DatastoreEmulatorHost)
	}

	var err error
	s.datastoreClient, err = datastore.NewClient(context.Background(), "testing")
	s.Require().NoError(err)
	s.repository = repository.NewRepository(s.datastoreClient)
}

func (s *repositoryTestSuite) AfterTest(suiteName string, testName string) {
	ctx := context.Background()
	err := s.datastoreClient.Delete(ctx, datastore.NameKey("Article", s.identifier, nil))
	s.Require().NoError(err)
}

func TestSuite_Repository(t *testing.T) {
	suite.Run(t, &repositoryTestSuite{})
}

func (s *repositoryTestSuite) TestCreateArticle() {
	// Arrange
	ctx := context.Background()

	want := repository.Article{
		Author: "author",
		Title:  "title",
	}
	s.identifier = want.Title

	// Act
	err := s.repository.CreateArticle(ctx, &want)

	// Assert
	s.NoError(err)
	got, err := s.repository.GetArticle(ctx, want.Title)
	s.NoError(err)
	s.Equal(want, got)
}

func (s *repositoryTestSuite) TestGetArticle() {
	// Arrange
	ctx := context.Background()

	want := repository.Article{
		Author: "author",
		Title:  "title",
	}
	s.identifier = want.Title
	err := s.repository.CreateArticle(ctx, &want)
	s.Require().NoError(err)

	// Act
	got, err := s.repository.GetArticle(ctx, want.Title)

	// Assert
	s.NoError(err)
	s.Equal(want, got)
}

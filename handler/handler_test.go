package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/datastore"

	"github.com/stretchr/testify/suite"

	"source.cloud.google.com/sk-borislav/ae-local/handler"
	"source.cloud.google.com/sk-borislav/ae-local/repository"
)

const DatastoreEmulatorHost = "0.0.0.0:8081"

type handlerTestSuite struct {
	suite.Suite
	repository      repository.Repository
	datastoreClient *datastore.Client
	sut             handler.Handler

	identifier string
}

func (s *handlerTestSuite) SetupTest() {
	if addr := os.Getenv("DATASTORE_EMULATOR_HOST"); addr == "" {
		_ = os.Setenv("DATASTORE_EMULATOR_HOST", DatastoreEmulatorHost)
	}

	var err error
	s.datastoreClient, err = datastore.NewClient(context.Background(), "testing")
	s.Require().NoError(err)
	s.repository = repository.NewRepository(s.datastoreClient)
	s.sut = handler.NewHandler(s.repository)
}

func TestSuite_Handler(t *testing.T) {
	suite.Run(t, &handlerTestSuite{})
}

func (s *handlerTestSuite) TestCreateArticle() {
	// Arrange
	article := repository.Article{
		Author: "article",
		Title:  "title",
	}
	s.identifier = article.Title

	bytes, err := json.Marshal(article)
	s.Require().NoError(err)

	defer s.cleanup()

	// Act
	req := httptest.NewRequest(http.MethodPost, "/create-article", strings.NewReader(string(bytes)))
	res := httptest.NewRecorder()
	s.sut.CreateArticle(res, req)

	// Assert
	s.Equal("{article title}", res.Body.String())
}

func (s *handlerTestSuite) TestCreateArticleWithErrorOnDecodingBody() {
	// Arrange
	invalid := `{`

	// Act
	req := httptest.NewRequest(http.MethodPost, "/create-article", strings.NewReader(invalid))
	res := httptest.NewRecorder()
	s.sut.CreateArticle(res, req)

	// Assert
	s.Equal("unexpected EOF", res.Body.String())
}

func (s *handlerTestSuite) TestGetArticle() {
	// Arrange
	article := repository.Article{
		Author: "article",
		Title:  "title",
	}

	s.identifier = article.Title
	s.createArticle(article)

	defer s.cleanup()

	// Act
	req := httptest.NewRequest(http.MethodGet, "/get-article/title", nil)
	res := httptest.NewRecorder()
	s.sut.GetArticle(res, req)

	// Assert
	s.Equal("{article title}", res.Body.String())
}

func (s *handlerTestSuite) TestGetArticleWithErrorOnGettingArticle() {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/get-article/title", nil)
	res := httptest.NewRecorder()

	// Act
	s.sut.GetArticle(res, req)

	// Assert
	s.Equal("datastore: no such entity", res.Body.String())
}

func (s *handlerTestSuite) createArticle(article repository.Article) {
	ctx := context.Background()
	err := s.repository.CreateArticle(ctx, &article)
	s.Require().NoError(err)
}

func (s *handlerTestSuite) cleanup() {
	ctx := context.Background()
	err := s.datastoreClient.Delete(ctx, datastore.NameKey("Article", s.identifier, nil))
	s.Require().NoError(err)
}

package post_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"

	"portfolio-api/api/resource/post"
	mockDB "portfolio-api/mock/db"
	testUtil "portfolio-api/util/test"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := post.NewRepository(db)

	mockRows := sqlmock.NewRows([]string{
		"id", "name", "email", "message", "description", "subject", "created_date",
	}).
		AddRow(uuid.New(), "Post1", "Email1", "Message1", "Description1", 0, time.Now()).
		AddRow(uuid.New(), "Post2", "Email2", "Message2", "Description2", 0, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT")).
		WillReturnRows(mockRows)

	posts, err := repo.List()
	testUtil.NoError(t, err)
	testUtil.Equal(t, 2, len(posts))
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := post.NewRepository(db)

	id := uuid.New()
	now := time.Now()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT")).
		WithArgs(id, "Name", "Email", "Message", "Description", 0, mockDB.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := &post.Post{
		ID:          id,
		Name:        "Name",
		Email:       "Email",
		Message:     "Message",
		Description: "Description",
		Subject:     0,
		CreatedDate: now,
	}

	_, err = repo.Create(p)
	testUtil.NoError(t, err)
}

func TestRepository_Read(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := post.NewRepository(db)

	id := uuid.New()
	now := time.Now()

	mockRows := sqlmock.NewRows([]string{
		"id", "name", "email", "message", "description", "subject", "created_date",
	}).
		AddRow(id, "Post1", "Email1", "Message1", "Description1", 0, now)

	mock.ExpectQuery(`^SELECT \* FROM "posts" WHERE id = \$1 ORDER BY "posts"\."id" LIMIT \$2$`).
		WithArgs(id, 1).
		WillReturnRows(mockRows)

	p, err := repo.Read(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, "Post1", p.Name)
	testUtil.Equal(t, "Email1", p.Email)
	testUtil.Equal(t, "Message1", p.Message)
	testUtil.Equal(t, "Description1", p.Description)
	testUtil.Equal(t, 0, int(p.Subject))
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := post.NewRepository(db)

	id := uuid.New()
	now := time.Now()

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"posts\" SET").
		WithArgs("Name", "Email", "Message", "Description", 0, mockDB.AnyTime{}, id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	p := &post.Post{
		ID:          id,
		Name:        "Name",
		Email:       "Email",
		Message:     "Message",
		Description: "Description",
		Subject:     0,
		CreatedDate: now,
	}

	rows, err := repo.Update(p)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := post.NewRepository(db)

	id := uuid.New()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("DELETE")).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows, err := repo.Delete(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}

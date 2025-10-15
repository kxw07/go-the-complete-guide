package event

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_unregisterEvent_should_return_nil_when_success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	eventId := int64(123)
	userId := int64(456)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM registrations WHERE event_id = ? AND user_id = ?")).
		WithArgs(eventId, userId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err = unregisterEvent(db, eventId, userId); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func Test_unregisterEvent_should_return_when_error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	eventId := int64(123)
	userId := int64(456)

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM registrations WHERE event_id = ? AND user_id = ?")).
		WithArgs(eventId, userId).
		WillReturnError(errors.New("mock error"))

	if err = unregisterEvent(db, eventId, userId); err == nil {
		t.Fatalf("should return error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

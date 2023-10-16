package service

import (
	"Finance/impl/repository"
	"Finance/payload"
	"Finance/utility"
	"log"
	"testing"
)

func TestCustomerService_Login_WrongPassword(t *testing.T) {
	db := utility.ConnectDB()
	if err := utility.MigrationDB(db); err != nil {
		log.Panicln("Error when do migration:", err)
	}

	customerServiceImpl := NewCustomerServiceImpl(
		repository.NewCustomerRepositoryImpl(db),
	)

	wantMessageError := "ERR_UNAUTHORIZED"
	_, err := customerServiceImpl.Login(payload.LoginRequest{
		Username: "user",
		Password: "pass",
	})

	if err != nil {
		if err.Error() != wantMessageError {
			t.Errorf("Login() error = %v, wantErr %v", err, wantMessageError)
		}
	}
}

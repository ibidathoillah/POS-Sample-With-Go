package repositories

import (
	"context"
	"time"

	"github.com/ibidathoillah/majoo-test/internal/globals"

	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"
)

type mysql struct{}

func (u *mysql) CreateUser(ctx context.Context, model *models.User) (*models.User, error) {

	now := time.Now()
	model.CreatedAt = &now
	model.UpdatedAt = &now

	result, err := globals.GetQuery(ctx).NamedExecContext(ctx,
		`
		INSERT INTO Users (name, user_name, password, created_at, created_by, updated_at, updated_by)
		VALUES (:name, :user_name, :password, :created_at, :created_by, updated_at, updated_by)
		`, model)

	if err != nil {
		return nil, err
	}

	model.ID, _ = result.LastInsertId()

	return model, err
}

func (u *mysql) UpdateUser(ctx context.Context, model *models.User) (*models.User, error) {
	panic("Not Implemeted")
}
func (u *mysql) DeleteUser(ctx context.Context, userID *int64) error {
	panic("Not Implemeted")
}
func (u *mysql) GetUser(ctx context.Context, userID *int64) (*models.User, error) {
	panic("Not Implemeted")
}

func (u *mysql) GetUserByUsername(ctx context.Context, userName string) (*models.User, error) {
	user := &models.User{}
	var arg map[string]interface{} = make(map[string]interface{})

	arg["user_name"] = userName

	row, err := globals.GetQuery(ctx).NamedQueryRowxContext(ctx, `
		SELECT id, name, user_name, password
		FROM Users
		WHERE user_name = :user_name
		LIMIT 1
	`, arg)
	if err != nil {

		return nil, err
	}

	err = row.StructScan(user)
	if err != nil {
		return nil, nil
	}

	return user, err
}

func NewUserRepository() Interface {
	return &mysql{}
}

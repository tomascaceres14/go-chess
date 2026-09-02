package user

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tomascaceres14/go-chess/api/internal/database/generated"
)

// For the sake of simplicity, PostgresRepository will just use the default
// *sqlc.Queries, even though it probably should have it's own interface to avoid
// implementing other model's methods.
type PostgresRepository struct {
	db *generated.Queries
}

func NewPostgresRepository(db *generated.Queries) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(ctx context.Context, user *User) (*User, error) {
	log.Printf("CREATING USER: %v", user)
	uid, err := r.db.CreateUser(ctx, generated.CreateUserParams{
		Username: pgtype.Text{
			String: user.Username,
			Valid:  true,
		},
		HashedPassword: pgtype.Text{
			String: user.HashedPassword,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, err
	}
	user.ID = uid.String()
	return user, nil
}

func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	usersDB, err := r.db.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(usersDB))
	for i, u := range usersDB {
		users[i] = &User{
			ID:             u.ID.String(),
			Username:       u.Username.String,
			HashedPassword: u.HashedPassword.String,
		}
	}

	return users, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (*User, error) {
	u, err := r.db.GetUserByID(ctx, pgtype.UUID{
		Bytes: uuid.MustParse(id),
	})

	if err != nil {
		return nil, err
	}

	return &User{
		Username:       u.Username.String,
		HashedPassword: u.HashedPassword.String,
	}, nil
}

func (r *PostgresRepository) ExistsByID(ctx context.Context, id string) bool {
	exists, _ := r.db.ExistsUserByID(ctx, pgtype.UUID{
		Bytes: uuid.MustParse(id),
	})
	return exists
}

package repositories

import (
	"authentication-feature/internal/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository interface {
	FindById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindById(id int) (*models.User, error) {
	u := &models.User{}

	err := r.db.QueryRow("SELECT id,name,email,password FROM users WHERE id=?", id).Scan(&u.Id, &u.Name, &u.Email, &u.Password)

	return u, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}

	err := r.db.QueryRow("SELECT id,name,email,password FROM users WHERE email=?", email).Scan(&u.Id, &u.Name, &u.Email, &u.Password)

	return u, err
}

// func (authRepo *UserRepository) GetUserByEmail(ctx context.Context, email string) (struct {
// 	Email    string
// 	Password string
// }, error) {
// 	user := struct {
// 		Email    string
// 		Password string
// 	}{}

// 	err := authRepo.DB.QueryRowContext(ctx, "SELECT email,password FROM users WHERE email=?", email).Scan(&user.Email, &user.Password)

// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	return user, nil
// }

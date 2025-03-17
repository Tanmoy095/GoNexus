package data

import (
	"context"
	"database/sql"
	"time"
)

const dbTimeOut = 60 * time.Second

var db *sql.DB

type Models struct {
	User User
}

func NewModels(dbpool *sql.DB) Models {
	db = dbpool
	return Models{
		User: User{},
	}
}

// Useer is a structure which hold one user from the  database
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns a slice of users ,sorted by last_name
func (u User) GetAll() ([]*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()
	query := `select id, email, first_name, last_name, password, user_active, created_at, updated_at
	from users order by last_name`

	rows, err := db.QueryContext(ctx, query) //The query is executed: SELECT id, email, first_name, last_name FROM users.
	//If the query succeeds, rows will contain the results, and err will be nil.
	//If the query fails (e.g., the database is down), err will contain the error.
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//defer schedules rows.Close() to run after the function returns, so you don’t have to worry about manually closing it.

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *User) GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, user_active, created_at, updated_at
    FROM users WHERE email = $1`
	//$1:
	//A placeholder for a parameterized query.
	//Replaced with the actual value when the query is executed.
	//	The database driver replaces $1 with the provided value (email).
	//The query is executed with the correct value.

	row := db.QueryRowContext(ctx, query, email)
	//Executes the query and returns a single row (since we’re querying by email, there should be at most one matching row).
	//email: The email address to search for in the database.

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *User) GetOneByID(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, email, first_name, last_name, password, user_active, created_at, updated_at
    FROM users WHERE id = $1`

	row := db.QueryRowContext(ctx, query, id)

	var user User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *User) UpdateUserByID() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `update users set
		email = $1,
		first_name = $2,
		last_name = $3,
		user_active = $4,
		updated_at = $5
		where id = $6
	`
	//This is the SQL query to update a user in the users table.
	//update users set ...: Specifies the table (users) and the columns to update.
	//where id = $6: Identifies which user to update by their id. The id is passed as a parameter ($6).
	_, err := db.ExecContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Active,
		time.Now(),
		u.ID,
	)
	//Executes the SQL update query with the provided parameters.

	if err != nil {
		return err
	}

	return nil
}

// update by email address...
func (u *User) UpdateByEmail() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	stmt := `update users set
		email = $1,
		first_name = $2,
		last_name = $3,
		user_active = $4,
		updated_at = $5
		where email = $6
	`

	_, err := db.ExecContext(ctx, stmt,
		u.Email,
		u.FirstName,
		u.LastName,
		u.Active,
		time.Now(),
		u.Email, // Use email as the identifier
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCurrentUser deletes the user represented by the current User instance from the database.
// It uses the ID field of the User struct to identify the user to delete.
func (u *User) DeleteCurrentUser() error {
	// Create a context with a timeout to prevent the database operation from running indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel() // Ensure the context is canceled when the function exits.

	// Define the SQL query to delete a user by their ID.
	stmt := `DELETE FROM users WHERE id = $1`

	// Execute the query with the ID of the current User instance.
	_, err := db.ExecContext(ctx, stmt, u.ID)
	if err != nil {
		// If there's an error, return it.
		return err
	}

	// If the operation is successful, return nil.
	return nil
}

// DeleteUserByID deletes a user from the database by their ID.
// It takes the ID of the user to delete as a parameter.
func (u *User) DeleteUserByID(id int) error {
	// Create a context with a timeout to prevent the database operation from running indefinitely.
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel() // Ensure the context is canceled when the function exits.

	// Define the SQL query to delete a user by their ID.
	stmt := `DELETE FROM users WHERE id = $1`

	// Execute the query with the provided ID.
	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		// If there's an error, return it.
		return err
	}

	// If the operation is successful, return nil.
	return nil
}

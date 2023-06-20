package domain

type Customer struct {
	ID        int    `db:"id"`
	Username  string `db:"user_name"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Status    int    `db:"status"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

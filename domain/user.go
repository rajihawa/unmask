package domain

type User struct {
	ID              string                 `json:"id,omitempty" rethinkdb:"id,omitempty"`
	Username        string                 `json:"username" rethinkdb:"username"`
	Password        string                 `json:"password,omitempty" rethinkdb:"-"`
	PasswordConfirm string                 `json:"password_confirm,omitempty" rethinkdb:"-"`
	PasswordHash    string                 `json:"-" rethinkdb:"password_hash"`
	Confirmed       bool                   `json:"confirmed" rethinkdb:"confirmed"`
	Disabled        bool                   `json:"disabled" rethinkdb:"disabled"`
	Data            map[string]interface{} `json:"data" rethinkdb:"data"`
	Project         *Project               `json:"project,omitempty" rethinkdb:"project_id,reference" rethinkdb_ref:"id"`
	Client          *Client                `json:"client,omitempty" rethinkdb:"client_id,reference" rethinkdb_ref:"id"`
	// EmailAddress    string                 `json:"email_address,omitempty" rethinkdb:"email_address,omitempty"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type GetUsersOpts struct {
	GetProject bool
	GetClient  bool
}

// UserstUsecases - the users' repository
type UsersRepository interface {
	GetAll(projectID string, opts GetUsersOpts) ([]User, error)
	Get(id string, opts GetUsersOpts) (*User, error)
	Insert(user User) error
	GetByUsername(username string, projectID string, opts GetUsersOpts) ([]User, error)
}

// UserstUsecases - the users' usecases
type UsersUsecases interface {
	GetAll(client Client, opts GetUsersOpts) ([]User, error)
	GetUser(id string, opts GetUsersOpts) (*User, error)
	SignupUser(client Client, user *User) error
	CheckUserLogin(userLogin UserLogin, client Client) (*User, error)
}

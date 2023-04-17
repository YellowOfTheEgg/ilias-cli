package ilias_api

type AuthService service

type User struct {
	Username	string  `schema:"-"`
	Token		string	`schema:"rtoken"`
}

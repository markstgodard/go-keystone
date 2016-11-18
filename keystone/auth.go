package keystone

type Domain struct {
	Name string `json:"name"`
}

type IdentityUser struct {
	User User `json:"user"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   Domain `json:"domain"`
}

type Identity struct {
	Methods  []string     `json:"methods"`
	Password IdentityUser `json:"password"`
}

type Auth struct {
	Identity Identity `json:"identity"`
}

type SingleAuth struct {
	Auth Auth `json:"auth"`
}

func NewAuth(username, password, domainName string) Auth {
	// create a password-based auth
	return Auth{
		Identity: Identity{
			Methods: []string{"password"},
			Password: IdentityUser{
				User: User{
					Name:     username,
					Password: password,
					Domain: Domain{
						Name: domainName,
					},
				},
			},
		},
	}
}

package main

type user struct {
	id       int
	username int
	//username wajarnya menggunakan string karena username terdiri dari kumpulan huruf angka
	password int
	//password yang baik terdiri dari kumpulan lowercase dan uppercase huruf, angka, dan simbol
}

type userservice struct {
	//nama dari object, variable, atau function sebaiknya menggunakan camelCase,snake_case atau PascalCase
	t []user
	//nama dari variable yang tidak menjelaskan apapun
}

func (u userservice) getallusers() []user {
	//nama dari object, variable, atau function sebaiknya menggunakan camelCase,snake_case atau PascalCase
	return u.t
}

func (u userservice) getuserbyid(id int) user {
	//nama dari object, variable, atau function sebaiknya menggunakan camelCase,snake_case atau PascalCase
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}

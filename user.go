package main

// User Model
type User struct {
	ID       int
	Name     string
	Address  string
	MyNumber string
	Votes    int
}

func getUser(name string, address string, myNumber string) (user User, err error) {
	row := db.QueryRow("SELECT id, votes FROM users WHERE name = ? AND address = ? AND mynumber = ?",
		name, address, myNumber)
	err = row.Scan(&user.ID, &user.Votes)
	return
}

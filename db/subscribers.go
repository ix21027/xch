package db

func createSubscribers() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS subscribers (id INTEGER PRIMARY KEY, email TEXT NOT NULL UNIQUE)")
	if err != nil {
		return
	}

}

func AddSubscriber(email string) error {
	_, err := DB.Exec("INSERT INTO subscribers (email) VALUES (?)", email)
	if err != nil {
		return err
	}
	return nil
}

func GetSubscribers() ([]string, error) {
	var emails []string
	rows, err := DB.Query("SELECT email FROM subscribers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {

			return nil, err
		}
		emails = append(emails, email)
	}
	return emails, nil
}

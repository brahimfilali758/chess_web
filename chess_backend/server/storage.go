package server

import "fmt"

func (psql *PsqlDB) createAccount(acount *Player) error {

	query := `insert into player 
		(username, passwordhash, email, rating, created_at)
		 values ($1, $2, $3, $4, $5)`
	_, err := psql.db.Query(query,
		acount.Username,
		acount.Passwordhash,
		acount.Email,
		acount.Rating,
		acount.created_at)

	return err
}

func (psql *PsqlDB) getAccountByUsername(username string) (Player, error) {
	query := fmt.Sprintf("SELECT * FROM player WHERE username = %s", username)
	rows, err := psql.db.Query(query)
	if err != nil {
		return Player{}, err
	}
	defer rows.Close()
	var player Player
	for rows.Next() {
		err := rows.Scan(&player.id, &player.Username, &player.Passwordhash, &player.Email, &player.Rating, &player.created_at)
		if err != nil {
			return Player{}, err
		}
	}
	return player, nil
}

func (psql *PsqlDB) deleteAccount(username string) error {
	_, err := psql.db.Query("delete from player where username = $1", username)
	fmt.Println(err)
	return err
}

func (psql *PsqlDB) getAllAccounts() ([]Player, error) {
	query := "SELECT * FROM player"
	rows, err := psql.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []Player
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.id, &player.Username, &player.Passwordhash, &player.Email, &player.Rating, &player.created_at)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

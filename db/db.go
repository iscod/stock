package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "xueqiucurl"
import "time"
import "fmt"
import "log"

// isntall comment for sql
func InstallComment(comment xueqiucurl.Comment, symbol string) (int64, error) {
	db, err := sql.Open("mysql", "root:root@/goStock")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	//to do query is have
	var id int64
	errdb := db.QueryRow("SELECT id FROM comment WHERE comment_id = ?", comment.Id).Scan(&id)

	//没有数据插入数据
	if errdb == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT INTO `comment` (`symbol`, `user_id`, `user_name`, `comment_id`, `title`,`source`, `text`, `created_at`) VALUES (?,?,?,?,?,?,?,?);")
		if err != nil {
			return 0, err
		}

		t := time.Unix(comment.Created_at/1000, 0)
		created_at := t.Format("2006-01-02 15:04:05")

		result, err := stmt.Exec(symbol, comment.User.Id, comment.User.Screen_name, comment.Id, comment.Title, comment.Source, comment.Text, created_at)
		if err != nil {
			return 0, err
		}

		id, err := result.LastInsertId()

		if err != nil {
			return 0, err
		}

		return id, err
	}

	if errdb != nil {
		return 0, errdb
	}

	return 0, errdb

}

func Connect() (string, error) {
	db, err := sql.Open("mysql", "root:root@/goStock")

	if err != nil {
		log.Printf("DB open err:")
		log.Fatal(err)
	} else {
		row, err := db.Query("SELECT user_id FROM comment WHERE 1")

		if err != nil {
			fmt.Println(err)
		}

		defer row.Close()

		for row.Next() {
			var user_id string

			if err := row.Scan(&user_id); err != nil {
				log.Fatal(err)
			}
			fmt.Println(user_id)
		}

		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return "ok", err
}

func install() {

}

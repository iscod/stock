package xueqiu

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "time"
import "fmt"
import "log"

var user = "root"
var passwd = "root"

func InstallPrice(symbol string, item Item) (int64, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	t := time.Unix(item.Timestamp/1000, 0)
	date_time := t.Format("2006-01-02 15:04:05")

	//to do query is have
	var id int64
	errdb := db.QueryRow("SELECT id FROM symbol_price WHERE date_time = ?", date_time).Scan(&id)

	//没有数据插入数据
	if errdb == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT INTO `symbol_price` (`symbol`, `price`, `date_time`, `created_at`) VALUES (?,?,?,?);")
		if err != nil {
			return 0, err
		}

		created_at := time.Now().Format("2006-01-02 15:04:05")
		result, err := stmt.Exec(symbol, item.Current, date_time, created_at)
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

func ICommPriCha(symbol string, comment_count int, stock_price float32, date_time string, c_type int64) (int64, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	//to do query is have
	var id int64
	errdb := db.QueryRow("SELECT id FROM comment_count_price_charet WHERE date_time = ? AND symbol = ? AND type = ? ", date_time, symbol, c_type).Scan(&id)

	//没有数据插入数据
	if errdb == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT INTO `comment_count_price_charet` (`symbol`, `comment_count`, `stock_price`, `type`, `date_time`, `created_at`) VALUES (?,?,?,?,?,?);")
		if err != nil {
			return 0, err
		}

		created_at := time.Now().Format("2006-01-02 15:04:05")
		result, err := stmt.Exec(symbol, comment_count, stock_price, c_type, date_time, created_at)
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

// isntall comment_count_price_charet for sql type:day,hour,year
func InstallCommentpricecha(symbol string, comment_count int, stock_price float32, date_time string, c_type string) (int64, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	//to do query is have
	var id int64
	errdb := db.QueryRow("SELECT id FROM comment_count_price_charet WHERE date_time = ? AND symbol = ? AND type = ? ", date_time, symbol, c_type).Scan(&id)

	//没有数据插入数据
	if errdb == sql.ErrNoRows {
		stmt, err := db.Prepare("INSERT INTO `comment_count_price_charet` (`symbol`, `comment_count`, `stock_price`, `type`, `date_time`, `created_at`) VALUES (?,?,?,?,?,?);")
		if err != nil {
			return 0, err
		}

		created_at := time.Now().Format("2006-01-02 15:04:05")
		result, err := stmt.Exec(symbol, comment_count, stock_price, c_type, date_time, created_at)
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

// isntall comment for sql
func InstallComment(comment Comment, symbol string) (int64, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

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

		created_at := time.Now().Format("2006-01-02 15:04:05")

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

func CountComment(starttime string, endtime string, symbol string) (int, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

	if err != nil {
		return 0, err
	}

	defer db.Close()

	result, err := db.Query("select count(*) as c from goStock.comment where symbol = ? and created_at >= ? and created_at <= ?;", symbol, starttime, endtime)
	if err != nil {
		return 0, err
	}

	var c int
	for result.Next() {
		if err := result.Scan(&c); err != nil {
			return 0, nil
		}

	}

	return c, nil
}

func Connect() (string, error) {
	db, err := sql.Open("mysql", user+":"+passwd+"@/goStock")

	if err != nil {
		log.Printf("DB open err:")
		log.Fatal(err)
	} else {

		row, err := db.Query("SELECT user_id FROM comment WHERE ")

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

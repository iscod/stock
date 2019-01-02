package main

import "fmt"
import "flag"
import "os"
import "xueqiucurl"
import "db"
import "time"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	symbol := flag.Arg(0)

	if symbol == "" {
		symbol = "SZ000651" //SZ000651格力电器,SZ000895//双汇发展,SH600019宝钢,SH600664哈药集团
	}

	for true {
		rs, err := xueqiucurl.GetComments(symbol)

		if err != nil {
			fmt.Println(err)
		}

		for _, value := range rs {
			fmt.Printf("[ 🍺 ] Comment id: %d, ", value.Id)
			t := time.Unix(value.Created_at/1000, 0)
			created_at := t.Format("2006-01-02 15:04:05")
			fmt.Printf(" Username : %s, title: %s, Time: %s", value.User.Screen_name, value.Title, created_at)
			id, err := db.InstallComment(value, symbol)

			if err != nil {
				fmt.Println(err)
			}

			if id == 0 {
				fmt.Printf(" Is Exist\n")
			} else {
				fmt.Printf(" Save Success\n")
			}

			time.Sleep(time.Duration(1) * time.Second)
		}
	}

}

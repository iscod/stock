package main

import (
	"flag"
	"gorm.io/gorm"
)
import "fmt"
import "os"
import "time"

const host string = "https://xueqiu.com"
const stockhost string = "https://stock.xueqiu.com"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	symbol := flag.Arg(0)
	if symbol == "" {
		symbol = "SZ000651" //SZ000651格力电器,SZ000895//双汇发展,SH600019宝钢,SH600664哈药集团,SZ000333美的集团
	}

	fmt.Printf("名称: %s\n", symbol)
	// for true {

	//获取价格信息
	quote, err := Getquote(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("当前价格: %f, 开盘价: %f, 均价: %f, 当天最低: %f, 当天最高: %f\n", quote.Current, quote.Open, quote.Avg_price, quote.Low, quote.High)

	//获取评论
	comments, err := GetComment(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	//保存评论
	for _, comment := range comments {
		t := time.Unix(comment.CreatedAt/1000, 0)
		createdAt := t.Format("2006-01-02 15:04:05")

		id, err := InstallComment(comment, symbol)
		if err != nil {
			fmt.Println(err)
			return
		}

		if comment.Id != 0 {
			fmt.Printf("New Comment Username : %s,title: %s, Time: %s,  %s\n", comment.User.ScreenName, comment.Title, createdAt, comment.Type)
		}
	}

	// //获取价格图表
	chart, err := Getchart(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	//插入股价
	for _, value := range chart.Items {
		_, err := InstallPrice(symbol, value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	//id, err := icountcomment(60, symbol, quote.Current)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//id_tm, err := icountcomment(600, symbol, quote.Current)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//id_hh, err := icountcomment(1800, symbol, quote.Current)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//id_h, err := icountcomment(3600, symbol, quote.Current)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("统计类别结果: %d,%d,%d,%d\n", id, id_tm, id_hh, id_h)
	//
	//id_d, err := icountcomment(86400, symbol, quote.Current)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("统计类别结果: %d,%d,%d,%d, %d\n", id, id_tm, id_hh, id_h, id_d)
	//return
}

func icountcomment(b_time int64, symbol string, price float32) (int64, error) {
	if b_time < 60 {
		b_time = 60
	}

	t := time.Now()
	//插入每30分钟统计
	ct := time.Unix(t.Unix()-b_time, 0)
	ct = time.Unix((ct.Unix()/b_time)*b_time, 0) //取整半小时
	created_at := ct.Format("2006-01-02 15:04:01")
	ft := time.Unix(ct.Unix()+b_time, 0) //结束时间
	created_ft := ft.Format("2006-01-02 15:04:01")

	comment_count, err := CountComment(symbol, created_at, created_ft)

	fmt.Println(comment_count)

	id, err := ICommPriCha(symbol, comment_count, price, created_at, b_time)
	if err != nil {
		return 0, err
	}

	return id, err
}

package main

import "fmt"
import "flag"
import "os"
import "xueqiu"
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
		symbol = "SZ000651" //SZ000651格力电器,SZ000895//双汇发展,SH600019宝钢,SH600664哈药集团,SZ000333美的集团
	}

	fmt.Printf("名称: %s\n", symbol)
	// for true {

	//获取价格信息
	quote, err := xueqiu.Getquote(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("当前价格: %f, 开盘价: %f, 均价: %f, 当天最低: %f, 当天最高: %f\n", quote.Current, quote.Open, quote.Avg_price, quote.Low, quote.High)

	//获取评论
	comments, err := xueqiu.GetComments(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	//保存评论
	for _, value := range comments {
		t := time.Unix(value.Created_at/1000, 0)
		created_at := t.Format("2006-01-02 15:04:05")
		id, err := xueqiu.InstallComment(value, symbol)
		if err != nil {
			fmt.Println(err)
			return
		}

		if id != 0 {
			fmt.Printf("New Comment List:\n")
			fmt.Printf(" Username : %s,title: %s, Time: %s\n", value.User.Screen_name, value.Title, created_at)
		}
	}

	// //获取价格图表
	chart, err := xueqiu.Getchart(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	//插入股价
	for _, value := range chart.Items {
		_, err := xueqiu.InstallPrice(symbol, value)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	id, err := icountcomment(60, symbol, quote.Current)

	if err != nil {
		fmt.Println(err)
		return
	}

	id_tm, err := icountcomment(600, symbol, quote.Current)

	if err != nil {
		fmt.Println(err)
		return
	}

	id_hh, err := icountcomment(1800, symbol, quote.Current)

	if err != nil {
		fmt.Println(err)
		return
	}

	id_h, err := icountcomment(3600, symbol, quote.Current)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("统计类别结果: %d,%d,%d,%d\n", id, id_tm, id_hh, id_h)

	id_d, err := icountcomment(86400, symbol, quote.Current)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("统计类别结果: %d,%d,%d,%d, %d\n", id, id_tm, id_hh, id_h, id_d)
	return
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

	comment_count, err := xueqiu.CountComment(symbol, created_at, created_ft)
	id, err := xueqiu.ICommPriCha(symbol, comment_count, price, created_at, b_time)
	if err != nil {
		return 0, err
	}

	return id, err
}

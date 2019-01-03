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
	t := time.Now()

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

	//插入上一每分钟统计

	//插入上一个每小时统计
	m_ct := time.Unix(t.Unix()-60, 0)
	m_created_at := m_ct.Format("2006-01-02 15:04:01")
	m_ft := time.Unix(t.Unix(), 0)
	m_created_ft := m_ft.Format("2006-01-02 15:04:01")

	comment_count, err := xueqiu.CountComment(m_created_at, m_created_ft)

	id, err := xueqiu.InstallCommentpricecha(symbol, comment_count, quote.Current, m_created_at, "m")
	if err != nil {
		fmt.Println(err)
		return
	}

	if id != 0 {

	}

	//插入上一个每小时统计
	h_ct := time.Unix(t.Unix()-3600, 0)
	h_created_at := h_ct.Format("2006-01-02 15:00:01")
	h_ft := time.Unix(t.Unix(), 0)
	h_created_ft := h_ft.Format("2006-01-02 15:00:01")

	comment_count_h, err := xueqiu.CountComment(h_created_at, h_created_ft)

	h_id, err := xueqiu.InstallCommentpricecha(symbol, comment_count_h, quote.Current, h_created_at, "h")
	if err != nil {
		fmt.Println(err)
		return
	}

	if h_id != 0 {

	}

	time.Sleep(time.Duration(1) * time.Second)
	// }
}

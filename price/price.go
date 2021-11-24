package price

import (
	"fmt"
	"github.com/iscod/stock/base"
	"github.com/iscod/stock/model"
	"gorm.io/gorm"
	"time"
)

func Run(symbol string, db *gorm.DB) {
	//获取价格信息
	quote, err := base.GetQuote(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}
	quote.CreatedAt = quote.Time / 1000
	quote.ExecAt = time.Unix(quote.CreatedAt, 0).Format("2006-01-02")
	var dest = &model.Quote{}
	err = db.Where("symbol = ? and exec_at = ? ", quote.Symbol, quote.ExecAt).Take(dest).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Save(quote).Error
	} else {
		quote.Id = dest.Id
		err = db.Save(quote).Error
	}
	fmt.Printf(" %s, %s, 价格: %f, 开盘价: %f, 均价: %f, 最低价: %f, 最高价: %f\n", quote.Name, time.Unix(quote.CreatedAt, 0).Format("2006-01-02 15:04:05"), quote.Current, quote.Open, quote.AvgPrice, quote.Low, quote.High)
}

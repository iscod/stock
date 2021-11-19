package detail

import (
	"fmt"
	"github.com/iscod/goStock/base"
	"github.com/iscod/goStock/model"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func Runs(symbol string, db *gorm.DB) error {

	detail, err := base.GetDetail(symbol)
	if err != nil {
		return err
	}

	T, err := time.ParseInLocation("20060102 15:04:05", fmt.Sprintf("%s 00:00:00", detail.Summary.Date), time.Local)
	if err != nil {
		return err
	}

	var quote model.Quote
	err = db.Where(model.Quote{Symbol: symbol}).Where("exec_at = ?", T.Format("2006-01-02")).First(&quote).Error

	if err != nil {
		return err
	}
	quote.Detail = detail.Detail
	vv, err := strconv.ParseFloat(detail.Summary.Volume, 64)
	if err == nil {
		quote.Volume = int64(vv)
	}
	fmt.Printf("%s, %s, 成交量: %d 手\n", quote.Name, T.Format("2006-01-02"), quote.Volume/100)
	err = db.Save(quote).Error
	return err
}

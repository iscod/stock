package detail

import (
	"fmt"
	"github.com/iscod/goStock/base"
	"github.com/iscod/goStock/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
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
	quote.BVolume, quote.SVolume = 0, 0
	var vv float64
	for _, v := range detail.Detail {
		vv, err = strconv.ParseFloat(strings.Trim(v[2], "\""), 64)
		switch strings.Trim(v[3], "\"") {
		case "B":
			quote.BVolume += int64(vv) * 100
		case "S":
			quote.SVolume += int64(vv) * 100
		case "M":
			quote.MVolume += int64(vv) * 100
		default:
			fmt.Printf("====%s", v[3])
		}
	}

	vv, err = strconv.ParseFloat(detail.Summary.Volume, 64)
	if err == nil {
		quote.Volume = int64(vv)
	}
	fmt.Printf("%s, %s, 成交量: %d, 卖盘: %d, 买盘: %d, 中盘 %d\n", quote.Name, T.Format("2006-01-02"), quote.Volume/100, quote.SVolume/100, quote.BVolume/100, quote.SVolume/100)
	err = db.Save(quote).Error
	return err
}

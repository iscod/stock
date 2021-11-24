package detail

import (
	"fmt"
	"github.com/iscod/stock/base"
	"github.com/iscod/stock/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func Run(symbol string, db *gorm.DB) error {

	detail, err := base.GetDetail(symbol)
	if err != nil {
		return err
	}

	summary, err := base.GetSummary(symbol)

	T, err := time.ParseInLocation("20060102 15:04:05", fmt.Sprintf("%s 00:00:00", detail.Summary.Date), time.Local)
	if err != nil {
		return err
	}

	var quote model.Quote
	err = db.Where(model.Quote{Symbol: symbol}).Where("exec_at = ?", T.Format("2006-01-02")).First(&quote).Error

	if err != nil {
		return err
	}

	quote.SummaryVolume = map[int]model.SummaryVolume{}
	var vv float64
	for _, v := range summary {
		if len(v) > 6 {
			k, _ := strconv.ParseFloat(v[0], 64)
			bv, err := strconv.ParseFloat(strings.Trim(v[4], " "), 64)
			if err != nil {
				fmt.Printf("%s", err)
			}
			sv, _ := strconv.ParseFloat(strings.Trim(v[5], " "), 64)
			mv, _ := strconv.ParseFloat(strings.Trim(v[6], " "), 64)
			quote.SummaryVolume[int(k)] = model.SummaryVolume{
				BVolume: int64(bv),
				SVolume: int64(sv),
				MVolume: int64(mv),
			}
		}
	}

	vv, err = strconv.ParseFloat(detail.Summary.Volume, 64)
	if err == nil {
		quote.Volume = int64(vv)
	}
	fmt.Printf("%s, %s, 成交量: %d, 卖盘: %d, 买盘: %d, 中盘 %d\n", quote.Name, T.Format("2006-01-02"), quote.Volume/100, quote.SummaryVolume[10].SVolume, quote.SummaryVolume[10].BVolume, quote.SummaryVolume[10].MVolume)

	message, err := base.GetFundFlowMessage(symbol)
	if err != nil {
		fmt.Printf("Err: %s", err)
	}

	l := len(message.Data.HistoryFundFlow.OneDayKlineList)
	if l > 0 {
		fundFlow := message.Data.HistoryFundFlow.OneDayKlineList[l-1]
		if quote.ExecAt == fundFlow.Date {
			quote.FundFlow = model.FundFlow{
				MainNetIn:     fundFlow.MainNetIn,
				AvgIn:         fundFlow.AvgIn,
				TodayFundFlow: message.Data.TodayFundFlow,
			}
		}
	}

	fmt.Printf("%s, %s, 主力净流入: %s, 主力流入%s, 流出: %s", quote.Name, T.Format("2006-01-02"), message.Data.TodayFundFlow.MainNetIn, message.Data.TodayFundFlow.MainIn, message.Data.TodayFundFlow.MainOut)

	err = db.Save(quote).Error
	return err
}

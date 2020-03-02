package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-demo/infra/base"
	"testing"
	"time"
)

// 1. 红包商品数据写入
func TestRedEnvelopeGoodsDao_Insert(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &RedEnvelopeGoodsDao{runner: runner}
		po := &RedEnvelopeGoods{
			Id:               10001,
			EnvelopeNo:       "test-envelope-002",
			EnvelopeType:     1,
			Username:         sql.NullString{"lzq", true},
			UserId:           "020202020",
			Blessing:         sql.NullString{"lzzscl", true},
			Amount:           decimal.NewFromFloat(100),
			AmountOne:        decimal.NewFromFloat(100),
			Quantity:         1,
			RemainAmount:     decimal.NewFromFloat(100),
			RemainQuantity:   1,
			ExpiredAt:        time.Now().Add(time.Hour),
			Status:           0,
			OrderType:        1,
			PayStatus:        1,
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
			OriginEnvelopeNo: "",
		}
		Convey("写入红包信息", t, func() {
			rs, err := dao.Insert(po)
			So(err, ShouldBeNil)
			So(rs, ShouldBeGreaterThan, 0)
		})
		Convey("查询红包数据", t, func() {
			rs := dao.GetOne("test-envelope-002")
			So(rs, ShouldNotBeNil)
			So(rs.EnvelopeNo, ShouldEqual, "test-envelope-002")
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

// 2. 更新红包剩余金额和数量
func TestRedEnvelopeGoodsDao_UpdateBalance(t *testing.T) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := &RedEnvelopeGoodsDao{runner: runner}
		Convey("更新红包剩余金额和数量", t, func() {
			rs, err := dao.UpdateBalance("test-envelope-002", decimal.NewFromFloat(20))
			So(err, ShouldBeNil)
			So(rs, ShouldEqual, 1)
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

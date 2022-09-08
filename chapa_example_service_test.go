package chapa

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChapaExampleService(t *testing.T) {

	Convey("Chapa Example Service", t, func() {

		ctx := context.Background()

		exampleService := NewExamplePaymentService(
			New("CHAPA_API_KEY"),
		)

		Convey("can list payment transactions", func() {

			transactionList, err := exampleService.ListPaymentTransactions(ctx)
			So(err, ShouldBeNil)

			So(len(transactionList.Transactions), ShouldEqual, 2)
		})

		Convey("can successfully checkout", func() {

			form := &CheckoutForm{
				Amount:   12.30,
				Currency: "ETB",
			}

			paymentTxn, err := exampleService.Checkout(ctx, 1032, form)
			So(err, ShouldBeNil)

			So(paymentTxn.Amount, ShouldEqual, form.Amount)
			So(paymentTxn.Currency, ShouldEqual, form.Currency)
			So(paymentTxn.Status, ShouldEqual, PendingTransactionStatus)
			So(paymentTxn.MerchantFee, ShouldBeZeroValue)
			So(paymentTxn.TransactionID, ShouldNotBeZeroValue)

			So(len(transactions), ShouldEqual, 3)
		})

		Convey("cannot checkout if user is unavailable", func() {

			form := &CheckoutForm{
				Amount:   12.30,
				Currency: "ETB",
			}

			_, err := exampleService.Checkout(ctx, 0, form)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "user not found")
		})
	})
}

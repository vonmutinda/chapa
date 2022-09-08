package chapa

import (
	"os"
	"testing"

	"syreclabs.com/go/faker"

	. "github.com/smartystreets/goconvey/convey"
)

func TestChapa(t *testing.T) {

	Convey("Chapa API", t, func() {

		chapaPaymentProvider := New(os.Getenv("CHAPA_API_KEY"))

		Convey("can prompt payment from users", func() {

			request := &ChapaPaymentRequest{
				Amount:         10,
				Currency:       "ETB",
				FirstName:      "Chapa",
				LastName:       "ET",
				Email:          "chapa@et.io",
				CallbackURL:    "https://posthere.io/e631-44fe-a19e",
				TransactionRef: faker.RandomString(20),
				Customization: map[string]interface{}{
					"title":       "A Unique Title",
					"description": "This a perfect description",
					"logo":        "https://your.logo",
				},
			}

			response, err := chapaPaymentProvider.PaymentRequest(request)
			So(err, ShouldBeNil)

			So(response.Status, ShouldEqual, "success")
			So(response.Message, ShouldEqual, "Hosted Link")
			So(response.Data.CheckoutURL, ShouldContainSubstring, "https://checkout.chapa.co/checkout/payment")
		})

		Convey("can verify transactions", func() {

			response, err := chapaPaymentProvider.Verify("take-this-2-the-bank") // a paid txn
			So(err, ShouldBeNil)

			So(response.Status, ShouldEqual, "success")
			So(response.Message, ShouldEqual, "Payment details")
			So(response.Data.TransactionFee, ShouldNotBeZeroValue)
		})

		Convey("cannot verify a transaction that's yet to be paid for", func() {

			request := &ChapaPaymentRequest{
				Amount:         10,
				Currency:       "ETB",
				FirstName:      "Chapa",
				LastName:       "ET",
				Email:          "chapa@et.io",
				CallbackURL:    "https://posthere.io/e631-44fe-a19e",
				TransactionRef: faker.RandomString(20),
				Customization: map[string]interface{}{
					"title":       "A Unique Title",
					"description": "This a perfect description",
					"logo":        "https://your.logo",
				},
			}

			_, err := chapaPaymentProvider.PaymentRequest(request)
			So(err, ShouldBeNil)

			response, err := chapaPaymentProvider.Verify(request.TransactionRef)
			So(err, ShouldBeNil)

			So(response.Message, ShouldEqual, "Payment not paid yet")
		})
	})
}

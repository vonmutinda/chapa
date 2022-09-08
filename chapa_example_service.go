package chapa

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"syreclabs.com/go/faker"
)

type (
	CheckoutForm struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	}

	PaymentTransaction struct {
		TransactionID string            `json:"transaction_id"`
		User          *User             `json:"user"`
		Amount        float64           `json:"amount"`
		Currency      string            `json:"currency"`
		MerchantFee   float64           `json:"merchant_fee"` // txn fee
		Status        TransactionStatus `json:"status"`
		TxnDate       time.Time         `json:"transaction_date"`
	}

	TransactionList struct {
		Transactions []*PaymentTransaction `json:"transactions"`
		// Pagination -> you could add pagination to this struct as well
	}

	TransactionStatus string

	User struct {
		ID        int64  `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

const (
	FailedTransactionStatus  TransactionStatus = "failed"
	PendingTransactionStatus TransactionStatus = "pending"
	SuccessTransactionStatus TransactionStatus = "success"
)

// Placeholder data
var (
	users = []*User{
		{
			ID:        1002,
			FirstName: "Jon",
			LastName:  "Do",
			Email:     faker.Internet().Email(),
		},
		{
			ID:        1032,
			FirstName: "Mary",
			LastName:  "Josef",
			Email:     faker.Internet().Email(),
		},
	}

	transactions = []*PaymentTransaction{
		{
			TransactionID: faker.RandomString(10),
			Amount:        10.00,
			MerchantFee:   0.35,
			Currency:      "ETB",
			TxnDate:       time.Now(),
			User:          users[0],
		},
		{
			TransactionID: faker.RandomString(10),
			Amount:        120.00,
			MerchantFee:   1.35,
			Currency:      "USD",
			TxnDate:       time.Now(),
			User:          users[1],
		},
	}
)

type (
	ExamplePaymentService interface {
		Checkout(ctx context.Context, userID int64, form *CheckoutForm) (*ChapaPaymentResponse, error)
		ListPaymentTransactions(ctx context.Context) (*TransactionList, error)
	}

	AppExamplePaymentService struct {
		mu                     *sync.Mutex
		paymentGatewayProvider ChapaAPI
	}
)

func NewExamplePaymentService(
	paymentGatewayProvider ChapaAPI,
) *AppExamplePaymentService {
	return &AppExamplePaymentService{
		mu:                     &sync.Mutex{},
		paymentGatewayProvider: paymentGatewayProvider,
	}
}

func (s *AppExamplePaymentService) Checkout(ctx context.Context, userID int64, form *CheckoutForm) (*PaymentTransaction, error) {

	user, err := s.userByID(ctx, userID)
	if err != nil {
		return &PaymentTransaction{}, err
	}

	invoice := &ChapaPaymentRequest{
		Amount:         form.Amount,
		Currency:       form.Currency,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		CallbackURL:    faker.Internet().Url(),
		TransactionRef: faker.RandomString(10),
	}

	response, err := s.paymentGatewayProvider.PaymentRequest(invoice)
	if err != nil {
		return &PaymentTransaction{}, err
	}

	if response.Status != "success" {

		// log the response
		log.Printf("[ERROR] Failed to checkout user request response = [%+v]", response)

		return &PaymentTransaction{}, fmt.Errorf("failed to checkout err = %v", response.Message)
	}

	transcation := &PaymentTransaction{
		TransactionID: invoice.TransactionRef,
		Amount:        form.Amount,
		Currency:      form.Currency,
		User:          user,
		Status:        PendingTransactionStatus,
		TxnDate:       time.Now(),
	}

	err = s.savePaymentTransaction(ctx, transcation)
	if err != nil {
		return &PaymentTransaction{}, nil
	}

	return transcation, nil
}

func (s *AppExamplePaymentService) ListPaymentTransactions(ctx context.Context) (*TransactionList, error) {

	// validations here

	transactionList := &TransactionList{
		Transactions: transactions,
	}

	return transactionList, nil
}

func (s *AppExamplePaymentService) savePaymentTransaction(ctx context.Context, transaction *PaymentTransaction) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	transactions = append([]*PaymentTransaction{transaction}, transactions...)

	return nil
}

// userByID - normally you'd fetch user from the db
func (s *AppExamplePaymentService) userByID(ctx context.Context, userID int64) (*User, error) {

	for index := range users {
		if users[index].ID == userID {
			return users[index], nil
		}
	}

	return &User{}, errors.New("user not found")
}

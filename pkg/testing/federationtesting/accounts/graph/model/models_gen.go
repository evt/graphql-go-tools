// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type History interface {
	IsHistory()
}

type Identifiable interface {
	IsIdentifiable()
	GetID() string
}

type Info interface {
	IsInfo()
	GetQuantity() int
}

type PaymentType interface {
	IsPaymentType()
	GetName() string
}

type Store interface {
	IsStore()
	GetLocation() string
}

type Wallet interface {
	IsWallet()
	GetCurrency() string
	GetAmount() float64
}

type Card struct {
	Name          string `json:"name"`
	IsContactless bool   `json:"isContactless"`
}

func (Card) IsPaymentType()       {}
func (this Card) GetName() string { return this.Name }

type Cash struct {
	Name            string `json:"name"`
	RequiresReceipt bool   `json:"requiresReceipt"`
}

func (Cash) IsPaymentType()       {}
func (this Cash) GetName() string { return this.Name }

type Cat struct {
	Name string `json:"name"`
}

type GiftCard struct {
	Name         string `json:"name"`
	IsRefundable bool   `json:"isRefundable"`
}

func (GiftCard) IsPaymentType()       {}
func (this GiftCard) GetName() string { return this.Name }

type Product struct {
	Upc string `json:"upc"`
}

func (Product) IsEntity() {}

type Purchase struct {
	Product  *Product `json:"product"`
	Wallet   Wallet   `json:"wallet"`
	Quantity int      `json:"quantity"`
}

func (Purchase) IsHistory() {}

func (Purchase) IsInfo()               {}
func (this Purchase) GetQuantity() int { return this.Quantity }

type Sale struct {
	Product  *Product `json:"product"`
	Rating   int      `json:"rating"`
	Location string   `json:"location"`
}

func (Sale) IsHistory() {}

func (Sale) IsStore()                 {}
func (this Sale) GetLocation() string { return this.Location }

type User struct {
	ID               string      `json:"id"`
	Username         string      `json:"username"`
	History          []History   `json:"history"`
	RealName         string      `json:"realName"`
	PreferredPayment PaymentType `json:"preferredPayment"`
}

func (User) IsIdentifiable()    {}
func (this User) GetID() string { return this.ID }

func (User) IsEntity() {}

type WalletType1 struct {
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	SpecialField1 string  `json:"specialField1"`
}

func (WalletType1) IsWallet()                {}
func (this WalletType1) GetCurrency() string { return this.Currency }
func (this WalletType1) GetAmount() float64  { return this.Amount }

type WalletType2 struct {
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	SpecialField2 string  `json:"specialField2"`
}

func (WalletType2) IsWallet()                {}
func (this WalletType2) GetCurrency() string { return this.Currency }
func (this WalletType2) GetAmount() float64  { return this.Amount }

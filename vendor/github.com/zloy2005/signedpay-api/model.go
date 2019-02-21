package signedpay_api

type Error interface {
	GetCode() string
}

type ApiError struct {
	Code string `json:"code"`
	//Messages map[string][]string `json:"messages"`
}

func (e ApiError) GetCode() string {
	return e.Code
}

type ChargeRequest struct {
	Amount                    int    `json:"amount"`
	CallbackURL               string `json:"callback_url"`
	CardCvv                   string `json:"card_cvv"`
	CardExpMonth              string `json:"card_exp_month"`
	CardExpYear               string `json:"card_exp_year"`
	CardHolder                string `json:"card_holder"`
	CardNumber                string `json:"card_number"`
	Currency                  string `json:"currency"`
	CustomerAccountID         string `json:"customer_account_id"`
	CustomerDateOfBirth       string `json:"customer_date_of_birth"`
	CustomerEmail             string `json:"customer_email"`
	CustomerFirstName         string `json:"customer_first_name"`
	CustomerLastName          string `json:"customer_last_name"`
	CustomerPhone             string `json:"customer_phone"`
	Fraudulent                bool   `json:"fraudulent"`
	GeoCity                   string `json:"geo_city"`
	GeoCountry                string `json:"geo_country"`
	IPAddress                 string `json:"ip_address"`
	Language                  string `json:"language"`
	OrderDate                 string `json:"order_date"`
	OrderDescription          string `json:"order_description"`
	OrderID                   string `json:"order_id"`
	OrderItems                string `json:"order_items"`
	Platform                  string `json:"platform"`
	StatusURL                 string `json:"status_url"`
	ChargebackNotificationURL string `json:"chargeback_notification_url"`
	OrderNumber               int    `json:"order_number"`
}

type RefundRequest struct {
	OrderID string `json:"order_id"`
	Amount  int    `json:"amount"`
}

type StatusRequest struct {
	OrderID string `json:"order_id"`
}

type RecurringRequest struct {
	Amount                    int    `json:"amount"`
	RecurringToken            string `json:"recurring_token"`
	Currency                  string `json:"currency"`
	CustomerAccountID         string `json:"customer_account_id"`
	CustomerDateOfBirth       string `json:"customer_date_of_birth"`
	CustomerEmail             string `json:"customer_email"`
	CustomerFirstName         string `json:"customer_first_name"`
	CustomerLastName          string `json:"customer_last_name"`
	CustomerPhone             string `json:"customer_phone"`
	Fraudulent                bool   `json:"fraudulent"`
	GeoCity                   string `json:"geo_city"`
	GeoCountry                string `json:"geo_country"`
	IPAddress                 string `json:"ip_address"`
	Language                  string `json:"language"`
	OrderDate                 string `json:"order_date"`
	OrderDescription          string `json:"order_description"`
	OrderID                   string `json:"order_id"`
	OrderItems                string `json:"order_items"`
	Platform                  string `json:"platform"`
	CallbackURL               string `json:"callback_url"`
	StatusURL                 string `json:"status_url"`
	ChargebackNotificationURL string `json:"chargeback_notification_url"`
}

type PaymentResponse struct {
	Transactions map[string]Transaction `json:"transactions"`
	Order        struct {
		OrderID           string `json:"order_id"`
		Amount            int    `json:"amount"`
		Currency          string `json:"currency"`
		Fraudulent        bool   `json:"fraudulent"`
		MarketingAmount   int    `json:"marketing_amount"`
		MarketingCurrency string `json:"marketing_currency"`
		Status            string `json:"status"`
		RefundedAmount    int    `json:"refunded_amount"`
		TotalFeeAmount    int    `json:"total_fee_amount"`
	} `json:"order"`
	Transaction    Transaction `json:"transaction"`
	PaymentAdviser struct {
		Advise string `json:"advise"`
	} `json:"payment_adviser"`
}

type Transaction struct {
	ID         string `json:"id"`
	Operation  string `json:"operation"`
	Status     string `json:"status"`
	Descriptor string `json:"descriptor"`
	Amount     int    `json:"amount"`
	Currency   string `json:"currency"`
	Fee        struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"fee"`
	Card struct {
		Bank string `json:"bank"`
		//Bin          string `json:"bin"`
		Brand        string `json:"brand"`
		Country      string `json:"country"`
		Number       string `json:"number"`
		CardExpMonth string `json:"card_exp_month"`
		CardExpYear  string `json:"card_exp_year"`
		CardType     string `json:"card_type"`
		CardToken    struct {
			Token string `json:"token,omitempty"`
		} `json:"card_token",omitempty`
	} `json:"card",omitempty`
}

type chargeResponseWithApiError struct {
	PaymentResponse
	ApiError `json:"error"`
}

type recurringResponseWithApiError struct {
	PaymentResponse
	ApiError `json:"error"`
}

type statusResponseWithApiError struct {
	PaymentResponse
	ApiError `json:"error"`
}

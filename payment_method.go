package quickbooks

import "errors"

type PaymentMethod struct {
	SyncToken string   `json:",omitempty"`
	Domain    string   `json:"domain,omitempty"`
	Name      string   `json:",omitempty"`
	Active    bool     `json:",omitempty"`
	Type      string   `json:",omitempty"`
	Id        string   `json:",omitempty"`
	MetaData  MetaData `json:",omitempty"`
}

// CreatePaymentMethod creates the given paymentMethod within QuickBooks.
func (c *Client) CreatePaymentMethod(paymentMethod *PaymentMethod) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod
		Time          Date
	}

	if err := c.post("paymentmethod", paymentMethod, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.PaymentMethod, nil
}

// DeletePaymentMethod deletes the given paymentMethod from QuickBooks.
func (c *Client) DeletePaymentMethod(paymentMethod *PaymentMethod) error {
	if paymentMethod.Id == "" || paymentMethod.SyncToken == "" {
		return errors.New("missing id/sync token")
	}

	return c.post("paymentmethod", paymentMethod, nil, map[string]string{"operation": "delete"})
}

// FindPaymentMethodById returns an paymentMethod with a given Id.
func (c *Client) FindPaymentMethodById(id string) (*PaymentMethod, error) {
	var resp struct {
		PaymentMethod PaymentMethod
		Time          Date
	}

	if err := c.get("paymentmethod/"+id, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.PaymentMethod, nil
}

// QueryPaymentMethods accepts a SQL query and returns all paymentMethods found using it.
func (c *Client) QueryPaymentMethods(query string) ([]PaymentMethod, error) {
	var resp struct {
		QueryResponse struct {
			PaymentMethods []PaymentMethod `json:"PaymentMethod"`
			StartPosition  int
			MaxResults     int
		}
	}

	if err := c.query(query, &resp); err != nil {
		return nil, err
	}

	if resp.QueryResponse.PaymentMethods == nil {
		return nil, errors.New("could not find any payment methods")
	}

	return resp.QueryResponse.PaymentMethods, nil
}

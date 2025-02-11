package quickbooks

type Purchase struct {
	SyncToken   string                    `json:",omitempty"`
	Domain      string                    `json:"domain,omitempty"`
	TxnDate     Date                      `json:",omitempty"`
	TotalAmt    float64                   `json:",omitempty"`
	PaymentType string                    `json:",omitempty"`
	AccountRef  ReferenceType             `json:",omitempty"`
	Id          string                    `json:",omitempty"`
	MetaData    MetaData                  `json:",omitempty"`
	Line        []AccountBasedExpenseLine `json:",omitempty"`
	CurrencyRef ReferenceType             `json:",omitempty"`
}

type AccountBasedExpenseLine struct {
	DetailType                    string //"AccountBasedExpenseLineDetail",
	Amount                        float64
	Id                            string
	AccountBasedExpenseLineDetail AccountBasedExpense
	Description                   string
}

type AccountBasedExpense struct {
	AccountRef ReferenceType
}

// CreatePurchase creates the given purchase within QuickBooks.
func (c *Client) CreatePurchase(purchase *Purchase) (*Purchase, error) {
	var resp struct {
		Purchase Purchase
		Time     Date
	}

	if err := c.post("purchase", purchase, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.Purchase, nil
}

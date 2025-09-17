package quickbooks

type Transfer struct {
	SyncToken      string        `json:",omitempty"`
	Domain         string        `json:"domain,omitempty"`
	TxnDate        Date          `json:",omitempty"`
	ToAccountRef   ReferenceType `json:",omitempty"`
	FromAccountRef ReferenceType `json:",omitempty"`
	Amount         float64       `json:",omitempty"`
	Id             string        `json:",omitempty"`
	MetaData       MetaData      `json:",omitempty"`
	PrivateNote    string        `json:",omitempty"`
}

// CreateTransfer creates the given transfer within QuickBooks.
func (c *Client) CreateTransfer(transfer *Transfer) (*Transfer, error) {
	var resp struct {
		Transfer Transfer
		Time     Date
	}

	if err := c.post("transfer", transfer, &resp, nil); err != nil {
		return nil, err
	}

	return &resp.Transfer, nil
}

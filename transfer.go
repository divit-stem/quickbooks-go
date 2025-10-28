package quickbooks

import "errors"

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

// QueryTransfer accepts a SQL query and returns all payments found using it.
func (c *Client) QueryTransfer(query string) ([]Transfer, error) {
	var resp struct {
		QueryResponse struct {
			Transfers     []Transfer `json:"Transfer"`
			StartPosition int
			MaxResults    int
		}
	}

	if err := c.query(query, &resp); err != nil {
		return nil, err
	}

	if resp.QueryResponse.Transfers == nil {
		return nil, errors.New("could not find any transfer record")
	}

	return resp.QueryResponse.Transfers, nil
}

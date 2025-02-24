package quickbooks

type Preferences struct {
	EmailMessagesPrefs struct {
		InvoiceMessage MessagesPref
	}
	SalesFormsPrefs struct {
		SalesEmailBcc EmailAddress
		SalesEmailCc  EmailAddress
	}
	CurrencyPrefs struct {
		HomeCurrency struct {
			Value string `json:"value"`
		}
		MultiCurrencyEnabled bool
	}
}

// FindPreferences returns the QuickBooks Preferences object.
func (c *Client) FindPreferences() (*Preferences, error) {
	var resp struct {
		Preferences Preferences
		Time        Date
	}
	if err := c.get("preferences", &resp, nil); err != nil {
		return nil, err
	}
	return &resp.Preferences, nil
}

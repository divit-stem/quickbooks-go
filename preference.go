package quickbooks

type Preferences struct {
	EmailMessagesPrefs struct {
		InvoiceMessage MessagesPref
	}
	CurrencyPrefs struct {
		HomeCurrency struct {
			Value string
		}
		MultiCurrencyEnabled bool
	}
}

// FindPreferences returns the QuickBooks Preferences object.
func (c *Client) FindPreferences() (*Preferences, error) {
	var resp struct {
		PreferenceInfo Preferences
		Time           Date
	}

	if err := c.get("preferences", &resp, nil); err != nil {
		return nil, err
	}

	return &resp.PreferenceInfo, nil
}

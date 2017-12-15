package multichain

func (client *Client) Issue(isOpen bool, accountAddress, assetName string, quantity float64, units float64) (Response, error) {

	msg := client.Msg(
		"issue",
		[]interface{}{
			accountAddress,
			map[string]interface{}{
				"name": assetName,
				"open": isOpen,
			},
			quantity,
			units,
		},
	)

	obj, err := client.post(msg)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

	amount, err := strconv.ParseFloat(amountToSwap, 64)
	if err != nil {
		return nil, err
	}

	amountInWei, err := client.ToWei(amount, 18)
	if err != nil {
		return nil, err
	}

	toToken, err := client.GetToken(toTokenAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice()
	if err != nil {
		return nil, err
	}

	nonce, err := client.GetNonce(fromAddress)
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, toTokenAddress, amountInWei, 1000000, gasPrice, nil)

	signedTx, err := client.SignTx(tx, privateKey)
	if err != nil {
		return nil, err
	}

	hash, err := client.SendTx(signedTx)
	if err != nil {
		return nil, err
	}

	return hash, nil
  

package kgsservice

func (kgs *KGSService) authenticate() error {
	loginReq := map[string]string{
		"type":     "LOGIN",
		"name":     kgs.user,
		"password": kgs.pwd,
		"locale":   "ru_RU",
	}

	return kgs.sendPostRequest(loginReq)
}
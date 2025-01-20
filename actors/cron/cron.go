package cron

func CronConstructor(height uint64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return cronConstructorv8(raw)

	}
	return nil, nil
}

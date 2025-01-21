package cron

// TODO: update to correct height ranges
func CronConstructor(height uint64, raw []byte) (map[string]interface{}, error) {
	switch height {
	case 8:
		return cronConstructorv8(raw)
	case 7:
		return cronConstructorv7(raw)
	case 6:
		return cronConstructorv6(raw)
	case 5:
		return cronConstructorv5(raw)
	case 4:
		return cronConstructorv4(raw)
	case 3:
		return cronConstructorv3(raw)
	case 2:
		return cronConstructorv2(raw)
	}
	return nil, nil
}

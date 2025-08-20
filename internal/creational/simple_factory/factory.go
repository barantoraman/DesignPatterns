package simplefactory

// GetGun is a factory func.
// "ak47" for ak47 instance
// "musket" for musket instance
func GetGun(gunType string) Gun {
	switch gunType {
	case "ak47":
		return newAK47()
	case "musket":
		return newMusket()
	default:
		return nil
	}
}

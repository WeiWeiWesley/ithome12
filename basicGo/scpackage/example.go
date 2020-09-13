package scpackage

//Clinet example of scope
type Clinet struct {
	IP      string
	Host    string
	setting detail
}

//小寫private，無法被其他 package 直接存取
type detail struct {
	maxClient int
	maxIdle   int
}

//Create example of scope
func Create(ip, host string) *Clinet {
	return &Clinet{
		IP:      ip,
		Host:    host,
		setting: newSetting(),
	}
}

//小寫private，無法被其他 package 直接存取
func newSetting() detail {
	return detail{
		maxClient: 100,
		maxIdle:   10,
	}
}

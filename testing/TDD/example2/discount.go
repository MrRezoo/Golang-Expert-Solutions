package example2

func CalculateDiscount(userType string, amount float64) float64 {
	var discount float64
	if userType == "vip" {
		discount = amount * 0.10
	}
	return discount
}

func main() {
	// اجرای تست‌ها
	// می‌توانید به صورت دستی تابع را نیز آزمایش کنید
}

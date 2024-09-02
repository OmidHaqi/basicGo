package main // نام پکیج فایل سورس را مشخص می کند

import "math/rand" // یک پکیج استاندارد ایمپورت شده

const MaxRnd = 16 // یک ثابت تعریف شده است

// تابع StatRandomNumbers
// آمار اعداد تصادفی تولید شده توسط پکیج math/rand را می گیرد.
// ورودی: تعداد اعداد تصادفی
// خروجی: تعداد اعداد تصادفی کوچکتر از نصف حداکثر و تعداد اعداد تصادفی بزرگتر از نصف حداکثر
func StatRandomNumbers(n int) (int, int) {

	var a, b int // یک متغیر تعریف شده است

	for i := 0; i < n; i++ { // := عملگر انتساب کوتاه

		if rand.Intn(MaxRnd) < MaxRnd/2 {

			a = a + 1 // = عملگر انتساب

		} else {

			b++

		}

	}

	return a, b

}

// تابع main اولین نقطه شروع اجرای کد شما می باشد که در این تابع سایر موارد تعریف می شود.

func main() {

	var num = 100

	x, y := StatRandomNumbers(num)

	print("Result: ", x, " + ", y, " = ", num, "? ")

	println(x+y == num) //عملگر مقایسه‌ای

}

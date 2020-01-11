// 2020-1-9, 2020-1-11
// common method，如果考虑[1971,2100]的区间可以做些简化：知道1971-1-1是星期几，以此为基础；不知道，另选case中提供的，和指定日期计算到1971-1-1距离，
// 做差值计算
type Dates struct {
    day, month, year int
}

var s = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"} //string literal is not constant
var M  = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31} // init为平年

func leap_year(year int) bool { //
    if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
        return true
    }
    return false
}

func cal(date1, date2 Dates) int { //
    var deltaDays = 0
    
    if date1.year == date2.year {
        if date1.month == date2.month {
            deltaDays = date1.day - date2.day
        } else {
            if leap_year(date1.year) {
                M[2] = 29
            } else {
                M[2] = 28
            }
            for i := date2.month + 1; i < date1.month; i++ {
                deltaDays += M[i] 
            }
            deltaDays += (M[date2.month] - date2.day + date1.day)
        }
    } else {
        for i := date2.year + 1; i < date1.year; i++ {
            if leap_year(i) {
                deltaDays += 366
            } else {
                deltaDays += 365
            }
        }
        
        if leap_year(date2.year) {
            deltaDays += 366
            M[2] = 29 //
        } else {
            deltaDays += 365
            M[2] = 28 //
        }
        for i := 1; i < date2.month; i++ {
            deltaDays -= M[i]
        }
        deltaDays -= date2.day
        
        if leap_year(date1.year) {
            M[2] = 29 
        } else {
            M[2] = 28
        }
        for i := 1; i < date1.month; i++ {
            deltaDays += M[i]
        }
        deltaDays += date1.day
    }

    return deltaDays
}

func calc(date1, date2 Dates) int {
    var big = false
    if date1.year > date2.year {
        big = true
    } else if date1.year == date2.year {
        if date1.month > date2.month {
            big = true
        } else if date1.month == date2.month {
            if date1.day > date2.day {
                big = true
            }
        } 
    }
    
    if big { 
        return cal(date1, date2)
    }
    return -cal(date2, date1) // 取负
}

func dayOfTheWeek(day int, month int, year int) string {
    var deltaDays = calc(Dates{day:day, month:month, year:year}, Dates{day:31, month:8, year:2019})
    
    if deltaDays > 0 {
        return s[(deltaDays + 6) % 7]
    } 
    return s[(7 + (6 + deltaDays) % 7) % 7] // 最后对7取余是因为：(6 + deltaDays) % 7 可以为0
}

method2 cheat
func dayOfTheWeek(day int, month int, year int) string {
    // time.Local即本地时区, 取决于运行的系统环境设置, 优先取”TZ”这个环境变量, 然后取/etc/localtime, 都取不到就用UTC兜底.
    // month：如果是数字可以直接传参，否则需要先显示转化为Month类型
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}

// Zeller Formula (https://calendars.wikia.org/wiki/Zeller%27s_congruence)
var weekdays = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
func dayOfTheWeek(day int, month int, year int) string {
    if month < 3 {
        month += 12
        year--
    }
    J, K := year/100, year%100 // J和K的计算不能放在month的判断之前，譬如year = 2100
    h := (J/4 - 2*J + K + K/4 + 26*(month+1)/10 + day-1)%7
    return weekdays[(h+7)%7] //
}
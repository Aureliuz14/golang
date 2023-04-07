package model

import (
	"fmt"
	"strconv"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

// Setter dan Getter

func (d *Date) SetYear(year int) {
	now := time.Now().Year()
	if year < 1900 {
		panic("Year must greater than 1900")
	} else if year > now {
		panic("Year must less than " + strconv.Itoa(now))
	} else {
		d.year = year
	}

}

func (d Date) GetYear() int {
	return d.year
}

func (d *Date) SetMonth(month int) bool {
	nowMonth := int(time.Now().Month())
	now := time.Now().Year()
	a := d.year
	if month < 1 {
		panic("Month must greater than 1")
	} else if month > 12 {
		panic("Month must less than 12")
	} else if a <= now {
		if month > nowMonth {
			fmt.Println("Month must less than", nowMonth)
			return false
		} else {
			d.month = month
		}
	}
	return true

}

func (d Date) GetMonth() int {
	return d.month
}

func (d *Date) SetDay(day int) {
	nowMonth := int(time.Now().Month())
	nowYear := time.Now().Year()
	now := time.Now().Day()
	a := d.year
	b := d.month
	if day < 1 {
		panic("Day must greater than 1")
	} else if day > 31 {
		panic("Day must less than 31")
	} else if a <= nowYear {
		if b <= nowMonth {
			if day > now {
				fmt.Println("Day must less than ", now)
			}
		}
	} else {
		d.day = day
	}

}

func (d Date) GetDay() int {
	return d.day
}

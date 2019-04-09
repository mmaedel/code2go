package main

import (
	//"fmt"

	"net/http"
	"strconv"
	"time"
)

type Cal struct {
	Year  int
	Month int
	Days  map[int]string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	/* 	var start time.Time
	   	var end time.Time */

	k := 0

	now := time.Now().AddDate(0, k, 0)
	/* year, _ := strconv.Atoi(now.Format("2006"))

	m := time.Date(year, 04, 01, 00, 00, 00, 0, time.UTC)
	o := time.Date(year, 10, 01, 00, 00, 00, 0, time.UTC)

	switch {

	case m.AddDate(0, 0, -1).Weekday() == 0:

		start = m.AddDate(0, 0, -1)

	case m.AddDate(0, 0, -2).Weekday() == 0:

		start = m.AddDate(0, 0, -2)

	case m.AddDate(0, 0, -3).Weekday() == 0:

		start = m.AddDate(0, 0, -3)

	case m.AddDate(0, 0, -4).Weekday() == 0:

		start = m.AddDate(0, 0, -4)

	case m.AddDate(0, 0, -5).Weekday() == 0:

		start = m.AddDate(0, 0, -5)

	case m.AddDate(0, 0, -6).Weekday() == 0:

		start = m.AddDate(0, 0, -6)

	case m.AddDate(0, 0, -7).Weekday() == 0:

		start = m.AddDate(0, 0, -7)

	}

	switch {

	case o.AddDate(0, 0, -1).Weekday() == 0:

		end = o.AddDate(0, 0, -1)

	case o.AddDate(0, 0, -2).Weekday() == 0:

		end = o.AddDate(0, 0, -2)

	case o.AddDate(0, 0, -3).Weekday() == 0:

		end = o.AddDate(0, 0, -3)

	case o.AddDate(0, 0, -4).Weekday() == 0:

		end = o.AddDate(0, 0, -4)

	case o.AddDate(0, 0, -5).Weekday() == 0:

		end = o.AddDate(0, 0, -5)

	case o.AddDate(0, 0, -6).Weekday() == 0:

		end = o.AddDate(0, 0, -6)

	case o.AddDate(0, 0, -7).Weekday() == 0:

		end = o.AddDate(0, 0, -7)

	}

	fmt.Println(start, end) */

	var c Cal

	c.Year = now.Year()
	month, _ := strconv.Atoi(now.Format("01"))
	c.Month = month
	day := map[int]string{now.Day(): now.Weekday().String()}

	c.Days = day

	i := 1

	for i < 32 {

		d := now.AddDate(0, 0, i)

		m, _ := strconv.Atoi(d.Format("01"))

		if m != month {

			break

		}

		e, _ := strconv.Atoi(d.Format("02"))

		c.Days[e] = d.Weekday().String()

		i++

	}

	/* 	for j := range c.Days {

		d := strconv.Itoa(j)

		if d <= now.Format("02") {

			delete(c.Days, j)

		}

	} */

	j := 1

	for j > 0 {

		d := now.AddDate(0, 0, -j)

		m, _ := strconv.Atoi(d.Format("01"))

		if m != month {

			break

		}

		e, _ := strconv.Atoi(d.Format("02"))

		c.Days[e] = d.Weekday().String()

		j++

	}

	var q int

	l := len(c.Days)
	p, _ := strconv.Atoi(time.Now().Format("02"))

	for i := l; i >= p; i-- {

		q = i

	}

	//fmt.Println(q)

	//fmt.Println(c.Days[1])

	str := `

		<!DOCTYPE html>
		<html lang="en">
			 <head>
					<meta charset="UTF-8">
					<meta name="viewport" content="width=device-width, initial-scale=1.0">
					<meta http-equiv="X-UA-Compatible" content="ie=edge">
					<title>CODE2GO</title>
					<!-- CSS -->
					<!-- Add Material font (Roboto) and Material icon as needed -->
					<link href="https://fonts.googleapis.com/css?family=Roboto:300,300i,400,400i,500,500i,700,700i|Roboto+Mono:300,400,700|Roboto+Slab:300,400,700" rel="stylesheet">
					<link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">

					<!-- Add Material CSS, replace Bootstrap CSS -->
					<link href="https://assets.medienwerk.now.sh/material.min.css" rel="stylesheet">
					</head>
					<body style="background-color: #bcbcbc;">

					<div class="container" id="date" style="color:white; font-size:30px;">
					` + strconv.Itoa(c.Year) + ` - ` + strconv.Itoa(c.Month) + `
					</div>
					<div class="container" id="data" style="color:white; font-size:30px;">
	<br>

		`

	switch c.Days[q] {

	case "Monday":
		break
	case "Tuesday":
		str = str + `

			<div class="row">
			<span class="border">
				<div class="col-sm">
	   			Monday
				<br></div></span>
				`
	case "Wednesday":
		str = str + `
					<div class="row">
					<span class="border">
						<div class="col-sm">
						   Monday
						<br></div></span>
						<span class="border">
						<div class="col-sm">
						   Tuesday
						<br></div> </span>
						`
	case "Thursday":
		str = str + `
									<div class="row">
									<span class="border">
										<div class="col-sm">
										   Monday
										<br></div>
										</span><span class="border">
										<div class="col-sm">
										   Tuesday
										<br></div>
										</span><span class="border">
										<div class="col-sm">
										   Wednesday
										<br></div></span>
										`
	case "Friday":
		str = str + `
																	<div class="row">
																	<span class="border">
																		<div class="col-sm">
																		   Monday
																		<br></div>
																		</span><span class="border">
																		<div class="col-sm">
																		   Tuesday
																		<br></div>
																		</span><span class="border">
																		<div class="col-sm">
																		   Wednesday
																		<br></div>
																		</span><span class="border">
																		<div class="col-sm">
																		   Thursday
																		<br></div></span>
																		`
	case "Saturday":
		str = str + `
																																	<div class="row">
																																	<span class="border">
																																		<div class="col-sm">
																																		   Monday
																																		<br></div>
																																		</span><span class="border">
																																		<div class="col-sm">
																																		   Tuesday
																																		<br></div>
																																		</span><span class="border">
																																		<div class="col-sm">
																																		   Wednesday
																																		<br></div>
																																		</span><span class="border">
																																		<div class="col-sm">
																																		   Thursday
																																		<br></div>
																																		</span><span class="border">
																																		<div class="col-sm">
																																		   Friday
																																		<br></div>
																																		</span><span class="border">
																																		<div class="col-sm">
																																		   Saturday
																																		<br></div></span>
																																		`
	case "Sunday":
		str = str + `
																																																																	<div class="row">
																																																																	<span class="border">
																																																																		<div class="col-sm">
																																																																		   Monday
																																																																		<br></div>
																																																																		</span><span class="border">
																																																																		<div class="col-sm">
																																																																		   Tuesday
																																																																		<br></div>
																																																																		</span><span class="border">
																																																																		<div class="col-sm">
																																																																		   Wednesday
																																																																		<br></div>
																																																																		</span><span class="border">
																																																																		<div class="col-sm">
																																																																		   Thursday
																																																																		<br></div>
																																																																		</span><span class="border">
																																																																		<div class="col-sm">
																																																																		   Friday
																																																																		<br></div></span>
																																																																		`
	}

	for k := q; k < 32; k++ {

		switch c.Days[k] {

		case "Monday":

			str = str + `
			</div><div class="row">

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div> </a></span>

				`

		case "Tuesday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">

			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a> </span>

				`

		case "Wednesday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a> </span>

				`

		case "Thursday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a> </span>

				`

		case "Friday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a> </span>

				`

		case "Saturday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a></span>

				`

		case "Sunday":

			str = str + `

			<span class="border">
			<a href="mailto:medienwerksalzburg@gmail.com?subject=` + c.Days[k] + `, ` + strconv.Itoa(k) + `.` + strconv.Itoa(c.Month) + `.` + strconv.Itoa(c.Year) + `" target="_top">

			<div class="col-sm">
			` + c.Days[k] + `<br>` + strconv.Itoa(k) + `
				</div></a></span>

				`

		}

	}

	switch c.Days[len(c.Days)] {

	case "Monday":
		str = str + `
		<span class="border">
			<div class="col-sm">
			   Tuesday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Wednesday
			   <br></div></span>
			   <span class="border">
			<div class="col-sm">
			   Thursday
			   <br></div></span>
			   <span class="border">
			<div class="col-sm">
			   Friday
			   <br></div></span>
			   <span class="border">
			<div class="col-sm">
			   Saturday
			   <br></div></span>
			   <span class="border">
			<div class="col-sm">
			   Sunday
			   <br></div></span>

`
	case "Tuesday":
		str = str + `
		<span class="border">
			<div class="col-sm">
			   Wednesday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Thursday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Friday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Saturday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Sunday
			<br></div></span>

`
	case "Wednesday":
		str = str + `
		<span class="border">
			<div class="col-sm">
			   Thursday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Friday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Saturday
			<br></div><7span>
			<span class="border">
			<div class="col-sm">
			   Sunday
			<br></div></span>

`
	case "Thursday":
		str = str + `
		<span class="border">
			<div class="col-sm">
			   Friday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Saturday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Sunday
			<br></div></span>

`
	case "Friday":
		str = str + `
		<span class="border">
			<div class="col-sm">
			   Saturday
			<br></div></span>
			<span class="border">
			<div class="col-sm">
			   Sunday
			<br></div></span>

`
	case "Saturday":
		str = str + `
		<span class="border">

		<div class="col-sm">
		   Sunday
		<br></div></span>

`

	case "Sunday":
		str = str + `


		</div>

`
		break

	}

	/* 	for w := 1; w < 8; w++ {

	   		switch c.Days[k] {

	   		case "Monday":

	   		if k == 1 {

	   			str = str + `
	   			<div class="container" id="data" style="color:white; font-size:30px;">
	   			`
	   		}

	   		switch c.Days[k] {

	   		case "Monday":

	   			if k == 1 {

	   				str = str + `
	   				<div class="container" id="data" style="color:white; font-size:30px;">
	   				<div class="row">
	   				`
	   			}

	   			str = str + `</div><div class="row">`

	   		case "Sunday":

	   			str = str + `<br>`

	   		}

	   		if k > len(c.Days) {

	   			str = str + `
	   			<div class="col-sm">

	   			</div>

	   			`
	   		} else {

	   			str = str + `
	   			<div class="col-sm">
	   			` + strconv.Itoa(k) + " " + c.Days[k] + `
	   			</div>
	   			`

	   		}

	   	}

	   	str = str + `
	   		</div>
	   								  </div>
	   								  <!-- Then Material JavaScript on top of Bootstrap JavaScript -->
	   	<script src="https://assets.medienwerk.now.sh/material.min.js"></script>
	   								  </body>
	   								  </html>
	   		`
	*/
	//fmt.Println(str)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", strconv.Itoa(len(str)))
	w.Write([]byte(str))

}
package main

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Cal struct {
	Year  int
	Month int
	Days  map[int]string
}

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":

		url := strings.TrimPrefix(r.URL.Path, "/")

		n, _ := strconv.Atoi(url)

		now := time.Now().AddDate(0, n, 0)

		/* 	var start time.Time
		var end time.Time */
		var c Cal

		c.Year = now.Year()

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
   
   
   
					   <div class="container" id="search" style="color:white; font-size:30px;">
					   <form class="form-inline" role="form" method="post">
	   <input class="form-control mr-sm-2" type="text" placeholder="Search" aria-label="Search" id ="find" name ="find">
	   <button class="btn btn-outline-light my-2 my-sm-1" type="submit">Search</button><br>
	 </div><br><div class="container" id="nav" style="color:white;">
	 ` + strconv.Itoa(time.Now().Year()) + `<br>
	 `

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
		*/

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

		var p int
		var q int

		l := len(c.Days)

		if time.Now().Month() == now.Month() {

			p, _ = strconv.Atoi(time.Now().Format("02"))

		} else {

			p = 1

		}

		for i := l; i >= p; i-- {

			q = i

		}

		//o := strconv.Itoa(n + 1)

		for t := 0; t < n; t++ {

			if time.Now().AddDate(0, t, 0).Year() != c.Year {

				str = str + `<br> ` +

					strconv.Itoa(time.Now().AddDate(0, t, 0).Year()) + `<br>
	<button type="button" class="btn btn-outline-dark" onclick="window.location.href='` + strconv.Itoa(t) + `'">` + time.Now().AddDate(0, t, 0).Month().String() + `</button>

	
	`
				c.Year = time.Now().AddDate(0, t, 0).Year()

			} else {

				str = str + `
  
	
	<button type="button" class="btn btn-outline-dark" onclick="window.location.href='` + strconv.Itoa(t) + `'">` + time.Now().AddDate(0, t, 0).Month().String() + `</button>

	
	`
			}

		}

		str = str + `

	
	<button type="button" class="btn btn-light">` + now.Month().String() + `
	</button>
 
 `

		for t := n + 1; t < 21; t++ {

			if time.Now().AddDate(0, t, 0).Year() != c.Year {

				str = str + `<br> ` + strconv.Itoa(time.Now().AddDate(0, t, 0).Year()) + `<br>
	<button type="button" class="btn btn-outline-dark" onclick="window.location.href='` + strconv.Itoa(t) + `'">` + time.Now().AddDate(0, t, 0).Month().String() + `</button>

	
	`

				c.Year = time.Now().AddDate(0, t, 0).Year()

			} else {

				str = str + `
  
	
	<button type="button" class="btn btn-outline-dark" onclick="window.location.href='` + strconv.Itoa(t) + `'">` + time.Now().AddDate(0, t, 0).Month().String() + `</button>

	
	`
			}
		}

		str = str + `
	
	</form><br>
	</div>

					<div class="container" id="data" style="color:white;">
	<br>

		`

		c.Year = now.Year()

		switch c.Days[q] {

		case "Monday":
			break
		case "Tuesday":
			str = str + `

		
		
		
		<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
		`
		case "Wednesday":
			str = str + `
					

					
					
					<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
					
					
					<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>`
		case "Thursday":
			str = str + `
									

									
									
									<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
									
									
									<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>
									
									
									<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>`
		case "Friday":
			str = str + `
									

									
									
									<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
									
									
									<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>
									
									
									<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>
									
									
									<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>`
		case "Saturday":
			str = str + `
																																	

																																	
																																	
																																	<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
																																	
																																	
																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>
																																	
																																	
																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>
																																	
																																	
																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>
																																	
																																	
																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>
																																	`
		case "Sunday":
			str = str + `
																																																																	

																																																																	
																																																																	
																																																																	<br><button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Monday</span></button>
																																																																	
																																																																	
																																																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>
																																																																	
																																																																	
																																																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>
																																																																	
																																																																	
																																																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>
																																																																	
																																																																	
																																																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>
																																																																	
																																																																	
																																																																	<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>
																																																																	`
		}

		for k := q; k < 32; k++ {

			switch c.Days[k] {

			case "Monday":

				str = str + `
			

			

			
			<br><button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

				`

			case "Tuesday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

			`

			case "Wednesday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				 

				`

			case "Thursday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

				`

			case "Friday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

				`

			case "Saturday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

				`

			case "Sunday":

				str = str + `

			

			
			<button type="button" class="btn btn-link" onclick="window.location.href='entry#` + strconv.Itoa(c.Year) + `-` + strconv.Itoa(c.Month) + `-` + strconv.Itoa(k) + `'"><span class="badge badge-pill badge-light">` + c.Days[k] + `<br>` + strconv.Itoa(k) + `</span></button>
				

				`

			}

		}

		switch c.Days[len(c.Days)] {

		case "Monday":
			str = str + `
		
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Tuesday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>

			   
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>

			   
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>

			   
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>

			   
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>


`
		case "Tuesday":
			str = str + `
		
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Wednesday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>


`
		case "Wednesday":
			str = str + `
		
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Thursday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>


`
		case "Thursday":
			str = str + `
		
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Friday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>

`
		case "Friday":
			str = str + `
		
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Saturday</span></button>

			
			
			<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>


`
		case "Saturday":
			str = str + `
		

		
		<button type="button" class="btn btn-link"><span class="badge badge-pill badge-light">Sunday</span></button>


`

		case "Sunday":
			str = str + `

		
		</div><br>

`
			break

		}

		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("Content-Length", strconv.Itoa(len(str)))
		w.Write([]byte(str))

	case "PUT":

		r.ParseForm()

		//s := strings.Join(r.Form["find"], " ")

		/* client := &http.Client{}
		req, err := http.NewRequest(http.MethodPut, "localhost:5000/bolt/users/user2", Reader(s))
		if err != nil {
			fmt.Fprint(w, err)
		}
		_, err = client.Do(req)
		if err != nil {
			fmt.Fprint(w, err)
		} */

	}

}

func Reader(s string) io.Reader {

	return strings.NewReader(s)
}
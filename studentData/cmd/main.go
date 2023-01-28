package main

import "webhook/abcSchool/pkg"

func main() {
	pkg.Server()
}

// {
//     "schoolName": "abc1 School",
//     "schoolId": 2,
//     "web":{
//         "eventname": "addStudent",
//         "endpoint": "http://localhost:9091/webhook/studentAdded"
//     }
// }

// {
//     "name": "sanjeet",
//     "age": 26,
//     "schoolId": 2
// }

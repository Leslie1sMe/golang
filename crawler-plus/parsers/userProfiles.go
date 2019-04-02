package parsers

import (
	"fmt"
	"go_code/crawler-plus/engine"
	pb "grpc-crawler/proto"
	"regexp"
)

const compile = `<div class="m-btn purple" data-v-bff6f798>([^>]+[^>]+)</div>`

//GetUserProfile
func GetUserProfile(content []byte, name string) engine.ParseResult {
	var res engine.ParseResult
	compileReg := regexp.MustCompile(compile)
	compileResult := compileReg.FindAllSubmatch(content, -1)
	userProfile := pb.Profile{Name: name}
	for index, v := range compileResult {
		switch index {
		case 0:
			userProfile.MarriedOrNot = string(v[1])
		case 1:
			userProfile.Age = string(v[1])
		case 2:
			userProfile.Constellation = string(v[1])
		case 3:
			userProfile.Height = string(v[1])
		case 4:
			userProfile.Weight = string(v[1])
		case 5:
			userProfile.WorkPlace = string(v[1])
		case 6:
			userProfile.Income = string(v[1])
		case 7:
			userProfile.Profession = string(v[1])
		case 8:
			userProfile.DegreeOfEducation = string(v[1])
		default:
			fmt.Println("信息不全")
		}
	}
	res.Items = append(res.Items, userProfile)
	return res
}

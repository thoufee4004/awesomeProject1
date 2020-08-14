package models

import "io/ioutil"

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}
type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func LoadFile(fileName string) (string, error) {
	bytes,err := ioutil.ReadFile(fileName)
	if err !=nil{
		return "",err
	}
	return string(bytes),nil
}

func IsEmpty(data string) bool  {
	if len(data)<=0{
		return true
	}else{
		return false
	}
}
func UserIsValid(uName ,pwd string) bool{
	uName1,pwd1,valid :="","",false
	if uName1 == uName && pwd1 == pwd{
		valid=true
	}else {
		valid=false
	}
	return valid
}
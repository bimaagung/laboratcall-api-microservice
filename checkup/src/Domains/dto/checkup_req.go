package checkupdto

import (
	"errors"
)

type CheckupReq struct { 
	Name      string 
	Unit  	  string 
	Price  	  int 
}

func NewCheckupReq(name , unit string, price int) (*CheckupReq, error){
	
	if 	name  == "" || unit == "" || price == 0 {
		return nil, errors.New("CHECKUP.NOT_CONTAIN_NEEDED_PROPERTY")
	}

	return &CheckupReq{
		Name: name,
		Unit: unit,
		Price: price,
	}, nil
}

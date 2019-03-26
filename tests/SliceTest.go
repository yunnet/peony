package tests

import (
	"fmt"
	"github.com/pkg/errors"
)

type Slice []interface{}

var ERR_EMEM_EXIST = errors.New("emem already exists.")
var ERR_EMEM_NT_EXIST = errors.New("emem not exists.")

func NewSlice() Slice {
	return make(Slice, 0)
}

func (this *Slice)Add(elem interface{}) error {
	for _, v := range *this{
		if v == elem{
			fmt.Printf("Slice: Add elem: %v already exists.\n", elem)
			return ERR_EMEM_EXIST
		}
	}
	*this = append(*this, elem)
	fmt.Printf("Slice: Add elem: %v succ\n", elem)
	return nil
}

func (this *Slice)Remove(elem interface{}) error {
	found := false
	for i, v := range *this{
		if v == elem{
			if i == len(*this) - 1{
				*this = (*this)[:i]
			}else{
				*this = append((*this)[:i], (*this)[i+1:]...)
			}
			found = true
			break
		}
	}
	if !found{
		fmt.Printf("Slice: Remove elem: %v not exists.\n", elem)
		return ERR_EMEM_NT_EXIST
	}
	fmt.Printf("Slice: Remove elem: %v succ\n", elem)
	return nil
}

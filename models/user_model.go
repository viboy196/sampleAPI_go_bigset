package models

import (
	"SampleAPI_Bigset/helps"
	"encoding/json"
	"fmt"
	"log"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

const BS_USER = "User"

type User struct {
	UserID   string `json:"user_id" xml:"user_id"`
	Account  string `json:"account" xml:"account"`
	Password string `json:"password" xml:"password"`
}

func (this *User) String() string {
	return this.UserID
}

func (this *User) GetBsKey() generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", BS_USER))
}

func (this *User) GetAll() ([]User, int64, error) {
	var err error
	if totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey()); err == nil {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(), 0, int32(totalCount))
		if err != nil {
			return make([]User, 0), 0, err
		}
		transaction, err := this.UnMarshalArrayTItem(slice)
		return transaction, totalCount, err
	}

	return make([]User, 0), 0, err

}

func (this *User) UnMarshalArrayTItem(objects []*generic.TItem) ([]User, error) {
	objs := make([]User, 0)

	for _, object := range objects {
		obj := User{}
		err := json.Unmarshal(object.GetValue(), &obj)

		if err != nil {
			return make([]User, 0), err
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (this *User) GetPaginate(pos, count int32) ([]User, int64, error) {
	SetItem, err := bigsetIf.BsGetSlice(this.GetBsKey(), pos, count)
	if err != nil {
		return nil, 0, err
	}
	totalCout, err := bigsetIf.GetTotalCount(this.GetBsKey())
	if err != nil {
		return nil, 0, err
	}
	transactions, err := this.UnMarshalArrayTItem(SetItem)
	if err != nil {
		return nil, 0, err
	}

	return transactions, totalCout, err
}

func (this *User) Create() error {

	bUser, key, err := helps.MarshalBytes(this)
	if err != nil {
		return err
	}
	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bUser,
	})
}

func (this *User) PutItem() error {
	bUser, key, err := helps.MarshalBytes(this)
	if err != nil {
		return err
	}
	_, err = this.Get()
	if err != nil {
		return err
	}
	return bigsetIf.BsPutItem(this.GetBsKey(), &generic.TItem{
		Key:   key,
		Value: bUser,
	})
}

func (this *User) Delete() error {
	log.Println(this.String(), ":delete")
	return bigsetIf.BsRemoveItem(this.GetBsKey(), []byte(this.String()))
}

func (this *User) Get() (interface{}, error) {

	bytes, err := this.GetItemBytes()

	if err != nil {
		return nil, err
	}

	return helps.UnMarshalBytes(bytes)
}
func (this *User) GetItemBytes() ([]byte, error) {
	tUser, err := bigsetIf.BsGetItem(this.GetBsKey(), generic.TItemKey(this.String()))
	if err != nil {
		return nil, err
	}
	return tUser.GetValue(), nil
}

func (this *User) GetFromKey(key string) (*User, error) {
	item, err := bigsetIf.BsGetItem(this.GetBsKey(), generic.TItemKey(key))
	if err != nil {
		return nil, err
	}
	transaction := &User{}

	err = json.Unmarshal(item.GetValue(), &transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

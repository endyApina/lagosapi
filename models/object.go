package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	//Objects returns an array of Objects sturct
	Objects map[string]*Object
)

//Object struct handles object data
type Object struct {
	ObjectID   string
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
}

//AddOne adds one
func AddOne(object Object) (ObjectID string) {
	object.ObjectID = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectID] = &object
	return object.ObjectID
}

//GetOne gets one
func GetOne(ObjectID string) (object *Object, err error) {
	if v, ok := Objects[ObjectID]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectID Not Exist")
}

//GetAll gets all
func GetAll() map[string]*Object {
	return Objects
}

//Update updates
func Update(ObjectID string, Score int64) (err error) {
	if v, ok := Objects[ObjectID]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectID Not Exist")
}

//Delete delets
func Delete(ObjectID string) {
	delete(Objects, ObjectID)
}

package repository

import (
	"log"

	"github.com/ariffinsetya/micro/flash-auth/pkg/database"
	"github.com/ariffinsetya/micro/flash-auth/pkg/entity"
	"github.com/google/uuid"
)

type UserRepository struct {
	Db database.RedisDb
}

func getOid(object entity.User) string {
	return "User" + "#" + object.Id
}

func (u UserRepository) SaveUser(object entity.User) interface{} {
	if object.Id == "" {
		object.Id = uuid.NewString()
	}
	objectKey := getOid(object)
	res, err := u.Db.JsonHandler.JSONSet(objectKey, ".", object)
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return nil
	}
	return res
}

func (u UserRepository) FindUser(object entity.User) interface{} {
	objectKey := getOid(object)
	res, err := u.Db.JsonHandler.JSONGet(objectKey, ",")
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return nil
	}
	return res
}

func (u UserRepository) DeleteUser(object entity.User) interface{} {
	objectKey := getOid(object)
	res, err := u.Db.JsonHandler.JSONDel(objectKey, ".")
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return nil
	}
	return res
}

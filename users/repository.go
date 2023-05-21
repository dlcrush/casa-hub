package users

import "github.com/dlcrush/casa-hub/adapters"

type UserRepository struct {
	adapters.MongoRepository[User]
}

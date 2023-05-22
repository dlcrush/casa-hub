package properties

import "github.com/dlcrush/casa-hub/adapters"

type PropertyRepository struct {
	adapters.MongoRepository[Property]
}

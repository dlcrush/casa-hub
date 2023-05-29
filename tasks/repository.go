package tasks

import "github.com/dlcrush/casa-hub/adapters"

type TaskRepository struct {
	adapters.MongoRepository[Task]
}

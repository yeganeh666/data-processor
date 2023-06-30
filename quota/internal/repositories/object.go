package repositories

type ObjectRepository interface {
	Get(key, userID string) (*Object, error)
	Create(object *Object) (*Object, error)
}

type ObjectRepositoryImpl struct {
	*Repository
}

func NewObjectRepository(Repository *Repository) ObjectRepository {
	return &ObjectRepositoryImpl{
		Repository: Repository,
	}
}

func (r *ObjectRepositoryImpl) Create(object *Object) (*Object, error) {
	err := r.DB.Create(object).Error
	return object, err
}

func (r *ObjectRepositoryImpl) Get(key, userID string) (*Object, error) {
	object := new(Object)
	err := r.DB.Where("key = ? and user_id = ?", key, userID).First(&object).Error
	return object, err
}

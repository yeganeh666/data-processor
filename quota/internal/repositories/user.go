package repositories

type UserRepository interface {
	GetQuota(userID string) (*UserQuota, error)
	SaveQuota(userQuota *UserQuota) error
}

type UserRepositoryImpl struct {
	*Repository
}

func NewUserRepository(Repository *Repository) UserRepository {
	return &UserRepositoryImpl{
		Repository: Repository,
	}
}

func (r *UserRepositoryImpl) GetQuota(userID string) (*UserQuota, error) {
	userQuota := new(UserQuota)
	err := r.DB.Model(userQuota).Where("user_id = ?", userID).Find(&userQuota).Error
	return userQuota, err
}

func (r *UserRepositoryImpl) SaveQuota(userQuota *UserQuota) error {
	return r.DB.Save(&userQuota).Error
}

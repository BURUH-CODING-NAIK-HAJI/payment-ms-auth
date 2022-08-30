package welcomerepository

type WelcomeRepositoryInterface interface {
	Welcome() string
}

type WelcomeRepository struct {
}

func New() WelcomeRepositoryInterface {
	return &WelcomeRepository{}
}

func (welcomerepository *WelcomeRepository) Welcome() string {
	return "Hai From Repository"
}

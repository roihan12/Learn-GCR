package recipes

type recipeUsecase struct {
	recipeRepository Repository
}

func NewNoteUsecase(nr Repository) Usecase {
	return &recipeUsecase{
		recipeRepository: nr,
	}
}
func (ru *recipeUsecase) GetAll() []Domain {
	return ru.recipeRepository.GetAll()
}
func (ru *recipeUsecase) GetByID(id string) Domain {
	return ru.recipeRepository.GetByID(id)
}
func (ru *recipeUsecase) Create(recipeDomain *Domain) Domain {
	return ru.recipeRepository.Create(recipeDomain)
}
func (ru *recipeUsecase) Update(id string, recipeDomain *Domain) Domain {
	return ru.recipeRepository.Update(id, recipeDomain)
}
func (ru *recipeUsecase) Delete(id string) bool {
	return ru.recipeRepository.Delete(id)
}

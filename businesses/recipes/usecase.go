package recipes

type recipeUsecase struct {
	recipeRepository Repository
}

func NewNoteUsecase(nr Repository) Usecase {
	return &recipeUsecase{
		recipeRepository: nr,
	}
}
func (ru *recipeUsecase) GetAll(name string) []Domain {
	return ru.recipeRepository.GetAll(name)
}
func (ru *recipeUsecase) GetByID(id string) Domain {
	return ru.recipeRepository.GetByID(id)
}

func (ru *recipeUsecase) GetByCategoryID(id string) []Domain {
	return ru.recipeRepository.GetByCategoryID(id)
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

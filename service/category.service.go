package service

import (
	"echo-recipe/dto"
	"echo-recipe/entity"
	"echo-recipe/repository"
	"log"

	"github.com/mashingan/smapping"
)

type CategoryService interface {
	Insert(category dto.CategoryDTO) entity.Category
	Update(id string, category dto.CategoryDTO) entity.Category
	Delete(id string) bool
	All() []entity.Category
	FindByID(id string) entity.Category
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
	}
}

func (service *categoryService) Insert(c dto.CategoryDTO) entity.Category {

	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	res := service.categoryRepository.Insert(category)
	return res

}

func (service *categoryService) Update(id string, c dto.CategoryDTO) entity.Category {

	category := entity.Category{}
	err := smapping.FillStruct(&category, smapping.MapFields(&c))

	if err != nil {
		log.Fatalf("failed map %v", err)
	}

	res := service.categoryRepository.Update(id, category)
	return res

}

func (service *categoryService) Delete(id string) bool {
	return service.categoryRepository.Delete(id)

}

func (service *categoryService) All() []entity.Category {
	return service.categoryRepository.All()

}

func (service *categoryService) FindByID(id string) entity.Category {
	return service.categoryRepository.FindByID(id)

}

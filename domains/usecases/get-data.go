package usecases

import "github.com/tenlisboa/cache_service/services"

type GetDataInput struct {
    Key string
}

func NewGetDataUsecase(cacheService *services.Cache) GetDataUsecase {
    return GetDataUsecase{
            CacheService: cacheService,
        }
}

type GetDataUsecase struct {
    CacheService *services.Cache
}

func (usecase GetDataUsecase) Execute(input GetDataInput) (any, bool) {
    return usecase.CacheService.Get(input.Key)
}
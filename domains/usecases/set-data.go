package usecases

import "github.com/tenlisboa/cache_service/services"

type SetDataInput struct {
    Key string
    Data any
}

func NewSetDataUsecase(cacheService *services.Cache) SetDataUsecase {
    return SetDataUsecase{
        CacheService: cacheService,
    }
}

type SetDataUsecase struct {
    CacheService *services.Cache
}

func (usecase SetDataUsecase) Execute(input SetDataInput) {
    usecase.CacheService.Set(input.Key, input.Data)
}
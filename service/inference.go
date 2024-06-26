package service

import (
	"github.com/nandasafiqalfiansyah/hugTransformers/model"
	"github.com/nandasafiqalfiansyah/hugTransformers/repository"
)

type InferenceService struct {
	repo *repository.HuggingFaceRepository
}

func NewInferenceService(repo *repository.HuggingFaceRepository) *InferenceService {
	return &InferenceService{
		repo: repo,
	}
}

func (s *InferenceService) Inference(input model.Payload) (string, error) {
	data := model.Payload{
		Inputs: input.Inputs,
	}

	return s.repo.SendInferenceRequest(data)
}

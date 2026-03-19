package link

import "context"

type LinkQueryService struct {
	linkRepository linkRepository
}

func NewLinkQueryService(linkRepository linkRepository) *LinkQueryService {
	return &LinkQueryService{
		linkRepository: linkRepository,
	}
}

func (linkQueryService *LinkQueryService) GetLinkURL(ctx context.Context, code string) (string, error) {
	url, err := linkQueryService.linkRepository.GetLinkURLByCode(ctx, code)
	if err != nil {
		return "", err
	}
	return url, nil
}

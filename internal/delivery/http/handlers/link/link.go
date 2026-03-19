package link

import (
	"github.com/erazorrr/go-link-shortener/internal/usecase/link"
)

type LinkHandler struct {
	linkQueryService   *link.LinkQueryService
	linkCommandService *link.LinkCommandService
}

func NewLinkHandler(linkQueryService *link.LinkQueryService, linkCommandService *link.LinkCommandService) *LinkHandler {
	return &LinkHandler{linkQueryService: linkQueryService, linkCommandService: linkCommandService}
}

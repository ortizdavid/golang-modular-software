package helpers

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/ortizdavid/go-nopain/conversion"
)

type Pagination[T any] struct {
	Items    []T      `json:"items"`
	MetaData MetaData `json:"metadata"`
}

type MetaData struct {
	CurrentPage     int    `json:"current_page"`
	TotalItems      int64  `json:"total_items"`
	TotalPages      int    `json:"total_pages"`
	FirstPageUrl    string `json:"first_page_url"`
	PreviousPageUrl string `json:"previous_page_url"`
	NextPageUrl     string `json:"next_page_url"`
	LastPageUrl     string `json:"last_page_url"`
}

func NewPagination[T any](c *fiber.Ctx, items []T, count int64, currentPage int, limit int) (*Pagination[T], error) {
	if currentPage < 0 {
		return nil, fmt.Errorf("current page must be >= 0")
	}
	if limit < 1 {
		return nil, fmt.Errorf("page size must be >= 1")
	}
	pagination := Pagination[T]{
		Items: items,
		MetaData: MetaData{
			CurrentPage:     currentPage,
			TotalItems:      count,
			TotalPages:      int(math.Ceil(float64(count) / float64(limit))),
			FirstPageUrl:    "",
			PreviousPageUrl: "",
			NextPageUrl:     "",
			LastPageUrl:     "",
		},
	}
	pagination.calculateUrls(c, currentPage, limit)
	return &pagination, nil
}

func (p *Pagination[T]) HasNextPage() bool {
	return p.MetaData.CurrentPage < p.MetaData.TotalPages-1
}

func (p *Pagination[T]) HasPreviousPage() bool {
	return p.MetaData.CurrentPage > 0
}

func (p *Pagination[T]) calculateUrls(c *fiber.Ctx, currentPage, limit int) {
	baseUrl := getRequestBaseUrl(c)

	p.MetaData.FirstPageUrl = getPageUrl(baseUrl, 0, limit)
	if currentPage < p.MetaData.TotalPages {
		p.MetaData.NextPageUrl = getPageUrl(baseUrl, currentPage+1, limit)
	}
	if currentPage > 1 {
		p.MetaData.PreviousPageUrl = getPageUrl(baseUrl, currentPage-1, limit)
	}
	p.MetaData.LastPageUrl = getPageUrl(baseUrl, p.MetaData.TotalPages, limit)
}

func getRequestBaseUrl(c *fiber.Ctx) string {
	scheme := "http"
	if c.Protocol() == "https" {
		scheme = "https"
	}
	return scheme + "://" + c.Hostname() + c.Path()
}

func getPageUrl(baseUrl string, pageNumber, limit int) string {
	return baseUrl + "?current_page=" + conversion.IntToString(pageNumber) + "&limit=" + conversion.IntToString(limit)
}

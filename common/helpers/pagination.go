package helpers

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type PaginationRender struct {
	PageNumber int
	PrevPage   int
	NextPage   int
}	

type Pagination struct {
}

func NewPaginationRender(pageNumber int) PaginationRender {
	return PaginationRender{
		PageNumber: pageNumber,
		PrevPage:   pageNumber - 1,
		NextPage:   pageNumber + 1,
	}
}

func (p *Pagination) GetPageNumber( ctx *fiber.Ctx, pageQuery string) int {
	pageNumberStr := ctx.Query(pageQuery)
	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}
	return pageNumber
}

func (p *Pagination) CalculateTotalPages(totalRecords int, itemsPerPage int) int {
	return (totalRecords + itemsPerPage - 1) / itemsPerPage
}

func (p *Pagination) CalculateStartIndex(pageNumber int, itemsPerPage int) int {
	return  (pageNumber - 1) * itemsPerPage
}

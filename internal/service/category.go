package service

import (
	"context"

	"github.com/derivedpuma7/go-grpc/internal/database"
	"github.com/derivedpuma7/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService {
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category {
		Id: category.ID,
		Name: category.Name,
		Description: *category.Description,
	}

	return categoryResponse, nil
}
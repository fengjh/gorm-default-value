package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	initDB()
	migrate()

	retCode := m.Run()

	os.Exit(retCode)
}

func TestCreateProductWithNoValueSpecified(t *testing.T) {
	product := Product{Name: "Create product with no value specified"}

	if product.Price != 0 {
		t.Fatalf("expected new product's price is zero, but got %v", product.Price)
	}

	if product.Category != "" {
		t.Fatalf("expected new product's category is zero, but got %v", product.Category)
	}

	err := DB.Create(&product).Error

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if product.Price != 0.01 {
		t.Fatalf("expected product.Price with default value 0.01, but got %v", product.Price)
	}

	if product.Category != "clothing" {
		t.Fatalf("expected product.Price with default value clothing, but got %v", product.Category)
	}
}

func TestSaveProductPriceWithNoValueSpecified(t *testing.T) {
	product := Product{Name: "Save product with no value specified"}

	if product.Price != 0 {
		t.Fatalf("expected new product's price is zero, but got %v", product.Price)
	}

	if product.Category != "" {
		t.Fatalf("expected new product's category is zero, but got %v", product.Category)
	}

	err := DB.Save(&product).Error

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if product.Price != 0.01 {
		t.Fatalf("expected product.Price with default value 0.01, but got %v", product.Price)
	}

	if product.Category != "clothing" {
		t.Fatalf("expected product.Price with default value clothing, but got %v", product.Category)
	}
}

func TestCreateProductPriceWithZeroValueSpecified(t *testing.T) {
	product := Product{Name: "Product with zero value specified", Price: 0.0}

	if product.Price != 0 {
		t.Fatalf("expected new product's price is zero, but got %v", product.Price)
	}

	if product.Category != "" {
		t.Fatalf("expected new product's category is zero, but got %v", product.Category)
	}

	err := DB.Create(&product).Error

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if product.Price != 0.01 {
		t.Fatalf("expected product.Price with default value 0.01, but got %v", product.Price)
	}

	if product.Category != "clothing" {
		t.Fatalf("expected product.Price with default value clothing, but got %v", product.Category)
	}
}

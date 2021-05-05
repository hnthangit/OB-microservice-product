package file

import (
	"encoding/json"
	"io/ioutil"
	"ob-product/internal/entity"

	"github.com/sirupsen/logrus"
)

type JsonProduct struct {
	Products []*entity.Product
}

func ReadCatalogFile() ([]*entity.Product, error) {
	catalogBYTE, err := ioutil.ReadFile("../../data/products.json")

	//Check error open file
	if err != nil {
		logrus.Fatalf("Failed to open product catalog json file: %v", err)
		return nil, err
	}
	logrus.Info("Successfully read file")

	//Check error parse json -> object
	p := new(JsonProduct)
	if err := json.Unmarshal(catalogBYTE, p); err != nil {
		logrus.Fatalf("Failed to parse catalog json file: %v", err)
		return nil, err
	}
	logrus.Info("Successfully parsed product catalog json")

	return p.Products, nil
}

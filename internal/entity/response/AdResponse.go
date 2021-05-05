package response

import "ob-product/internal/entity"

type AdResponse struct {
	Ads []*entity.Ad
}

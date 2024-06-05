package oms_promotion

import (
	"context"
	"fmt"
	"sort"

	"github.com/ChangSZ/mall-go/internal/dao"
	"github.com/ChangSZ/mall-go/internal/dto"
	"github.com/ChangSZ/mall-go/internal/repository/mysql"
	"github.com/ChangSZ/mall-go/pkg/copy"
	"github.com/ChangSZ/mall-go/pkg/math"
)

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) i() {}

func (s *service) CalcCartPromotion(ctx context.Context, cartItemList []dto.OmsCartItem) (
	[]dto.CartPromotionItem, error) {
	// 1.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
	productCartMap := s.GroupCartItemBySpu(cartItemList)

	// 2.查询所有商品的优惠相关信息
	promotionProductList, err := s.GetPromotionProductList(ctx, cartItemList)
	if err != nil {
		return nil, err
	}

	// 3.根据商品促销类型计算商品促销优惠价格
	cartPromotionItemList := make([]dto.CartPromotionItem, 0)
	for productId, itemList := range productCartMap {
		promotionProduct := s.GetPromotionProductById(productId, promotionProductList)

		switch promotionProduct.PromotionType {
		case 1: // 单品促销
			for _, item := range itemList {
				cartPromotionItem := dto.CartPromotionItem{}
				copy.AssignStruct(&item, &cartPromotionItem)
				cartPromotionItem.PromotionMessage = "单品促销"
				// 商品原价-促销价
				skuStock := s.GetOriginalPrice(*promotionProduct, item.ProductSkuId)
				originalPrice := skuStock.Price
				// 单品促销使用原价
				cartPromotionItem.Price = originalPrice
				cartPromotionItem.ReduceAmount = originalPrice - skuStock.PromotionPrice
				cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				cartPromotionItem.Integration = promotionProduct.GiftPoint
				cartPromotionItem.Growth = promotionProduct.GiftGrowth
				cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
			}
		case 3: // 打折优惠
			count := s.GetCartItemCount(itemList)
			ladder := s.GetProductLadder(count, promotionProduct.ProductLadderList)
			if ladder != nil {
				for _, item := range itemList {
					cartPromotionItem := dto.CartPromotionItem{}
					copy.AssignStruct(&item, &cartPromotionItem)
					cartPromotionItem.PromotionMessage = s.GetLadderPromotionMessage(*ladder)
					// 商品原价-折扣*商品原价
					skuStock := s.GetOriginalPrice(*promotionProduct, item.ProductSkuId)
					originalPrice := skuStock.Price
					cartPromotionItem.ReduceAmount = originalPrice - (ladder.Discount * originalPrice)
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
				}
			} else {
				cartPromotionItemList = s.HandleNoReduce(cartPromotionItemList, itemList, *promotionProduct)
			}
		case 4: // 满减
			totalAmount := s.GetCartItemAmount(itemList, promotionProductList)
			fullReduction := s.GetProductFullReduction(totalAmount, promotionProduct.ProductFullReductionList)
			if fullReduction != nil {
				for _, item := range itemList {
					cartPromotionItem := dto.CartPromotionItem{}
					copy.AssignStruct(&item, &cartPromotionItem)
					cartPromotionItem.PromotionMessage = s.GetFullReductionPromotionMessage(*fullReduction)
					// (商品原价/总价)*满减金额
					skuStock := s.GetOriginalPrice(*promotionProduct, item.ProductSkuId)
					originalPrice := skuStock.Price

					cartPromotionItem.ReduceAmount = math.RoundHalfEven(
						(originalPrice/totalAmount)*fullReduction.ReducePrice, 1)
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
				}
			} else {
				cartPromotionItemList = s.HandleNoReduce(cartPromotionItemList, itemList, *promotionProduct)
			}
		default:
			cartPromotionItemList = s.HandleNoReduce(cartPromotionItemList, itemList, *promotionProduct)
		}
	}
	return cartPromotionItemList, nil
}

// GetPromotionProductList 查询所有商品的优惠相关信息
func (s *service) GetPromotionProductList(ctx context.Context, cartItemList []dto.OmsCartItem) (
	[]dto.PromotionProduct, error) {
	productIdList := make([]int64, len(cartItemList))
	for _, cartItem := range cartItemList {
		productIdList = append(productIdList, cartItem.ProductId)
	}
	return new(dao.ProductDao).GetPromotionProductList(ctx, mysql.DB().GetDbR().WithContext(ctx), productIdList)
}

// GroupCartItemBySpu 以spu为单位对购物车中商品进行分组
func (s *service) GroupCartItemBySpu(cartItemList []dto.OmsCartItem) map[int64][]dto.OmsCartItem {
	productCartMap := make(map[int64][]dto.OmsCartItem)
	for _, cartItem := range cartItemList {
		if productCartItemList, ok := productCartMap[cartItem.ProductId]; !ok {
			tmp := []dto.OmsCartItem{cartItem}
			productCartMap[cartItem.ProductId] = tmp
		} else {
			productCartItemList = append(productCartItemList, cartItem)
			productCartMap[cartItem.ProductId] = productCartItemList
		}
	}
	return productCartMap
}

// GetFullReductionPromotionMessage 获取满减促销消息
func (s *service) GetFullReductionPromotionMessage(fullReduction dto.PmsProductFullReduction) string {
	return fmt.Sprintf("满减优惠: 满%v元减%v元", fullReduction.FullPrice, fullReduction.ReducePrice)
}

// HandleNoReduce 对没满足优惠条件的商品进行处理
func (s *service) HandleNoReduce(cartPromotionItemList []dto.CartPromotionItem,
	itemList []dto.OmsCartItem, promotionProduct dto.PromotionProduct) []dto.CartPromotionItem {
	for _, item := range itemList {
		cartPromotionItem := dto.CartPromotionItem{}
		copy.AssignStruct(&item, &cartPromotionItem)
		cartPromotionItem.PromotionMessage = "无优惠"
		cartPromotionItem.ReduceAmount = 0
		skuStock := s.GetOriginalPrice(promotionProduct, item.ProductSkuId)
		if skuStock != nil {
			cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
		}
		cartPromotionItem.Integration = promotionProduct.GiftPoint
		cartPromotionItem.Growth = promotionProduct.GiftGrowth
		cartPromotionItemList = append(cartPromotionItemList, cartPromotionItem)
	}
	return cartPromotionItemList
}

// GetProductFullReduction 获取商品的满减信息
func (s *service) GetProductFullReduction(totalAmount float64,
	fullReductionList []dto.PmsProductFullReduction) *dto.PmsProductFullReduction {
	// 按条件从高到低排序
	sort.Slice(fullReductionList, func(i, j int) bool {
		return fullReductionList[i].FullPrice > fullReductionList[j].FullPrice
	})

	for _, fullReduction := range fullReductionList {
		if totalAmount >= fullReduction.FullPrice {
			return &fullReduction
		}
	}
	return nil
}

// GetLadderPromotionMessage 获取打折优惠的促销信息
func (s *service) GetLadderPromotionMessage(ladder dto.PmsProductLadder) string {
	return fmt.Sprintf("打折优惠: 满%v件, 打%v折", ladder.Count, ladder.Discount)
}

// GetProductLadder 根据购买商品数量获取满足条件的打折优惠策略
func (s *service) GetProductLadder(count int32, productLadderList []dto.PmsProductLadder) *dto.PmsProductLadder {
	// 按数量从大到小排序
	sort.Slice(productLadderList, func(i, j int) bool {
		return productLadderList[i].Count > productLadderList[j].Count
	})

	for _, productLadder := range productLadderList {
		if count >= productLadder.Count {
			return &productLadder
		}
	}
	return nil
}

// GetCartItemCount 获取购物车中指定商品的数量
func (s *service) GetCartItemCount(itemList []dto.OmsCartItem) int32 {
	var count int32 = 0
	for _, item := range itemList {
		count += item.Quantity
	}
	return count
}

// GetCartItemAmount 获取购物车中指定商品的总价
func (s *service) GetCartItemAmount(itemList []dto.OmsCartItem, promotionProductList []dto.PromotionProduct) float64 {
	var amount float64
	for _, item := range itemList {
		// 计算出商品原价
		promotionProduct := s.GetPromotionProductById(item.ProductId, promotionProductList)
		skuStock := s.GetOriginalPrice(*promotionProduct, item.ProductSkuId)
		amount += skuStock.Price * float64(item.Quantity)
	}
	return amount
}

// GetOriginalPrice 获取商品的原价
func (s *service) GetOriginalPrice(promotionProduct dto.PromotionProduct, productSkuId int64) *dto.PmsSkuStock {
	for _, skuStock := range promotionProduct.SkuStockList {
		if productSkuId == skuStock.Id {
			return &skuStock
		}
	}
	return nil
}

// GetPromotionProductById 根据商品id获取商品的促销信息
func (s *service) GetPromotionProductById(productId int64, promotionProductList []dto.PromotionProduct) *dto.PromotionProduct {
	for _, promotionProduct := range promotionProductList {
		if productId == promotionProduct.Id {
			return &promotionProduct
		}
	}
	return nil
}

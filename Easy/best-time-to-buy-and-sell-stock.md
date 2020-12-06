Code
====

```go
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}
```

Solution in mind
================

-	Iterate through the prices. On encountering a minimum price, update the minimum price. This keeps track of least price.

-	If price encountered is not minimum, calculate profit (price - minPrice). If profit is highest seen, update maxProfit. This keeps track of maximum profit attainable.

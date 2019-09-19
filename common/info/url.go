package info

// amazonURLBase amazon product url.
const amazonURLBase = "https://www.amazon.co.jp/gp/offer-listing/"

// surugayaURLBase surugaya product rul.
const surugayaURLBase = "https://www.suruga-ya.jp/product/detail/"

// GetAmazonPrdocutURL get prdocut url of amazon for market place.
func GetAmazonPrdocutURL(productID string) string {
	return amazonURLBase + productID
}

// GetSurugayaPrdocutURL get prdocut url of surugaya product place.
func GetSurugayaPrdocutURL(productID string) string {
	return surugayaURLBase + productID
}

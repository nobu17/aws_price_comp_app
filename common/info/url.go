package info

// amazonURLBase amazon product url.
const amazonURLBase = "https://www.amazon.co.jp/gp/offer-listing/"
const amazonStoreType = "amazon"

// surugayaURLBase surugaya product rul.
const surugayaURLBase = "https://www.suruga-ya.jp/product/detail/"
const surugayaStoreType = "surugaya"

// bookoffURLBase bookoff online url.
const bookoffURLBase = "http://www.bookoffonline.co.jp/old/"
const bookoffStoreType = "bookoff"

// GetAmazonPrdocutURL get prdocut url of amazon for market place.
func GetAmazonPrdocutURL(productID string) string {
	return amazonURLBase + productID
}

// GetSurugayaPrdocutURL get prdocut url of surugaya product place.
func GetSurugayaPrdocutURL(productID string) string {
	return surugayaURLBase + productID
}

// GetBookoffPrdocutURL get prdocut url of surugaya product place.
func GetBookoffPrdocutURL(productID string) string {
	return bookoffURLBase + productID
}

// IsStoreTypeAmazon judge is store type is amazon.
func IsStoreTypeAmazon(storeType string) bool {
	if storeType == amazonStoreType {
		return true
	}
	return false
}

// IsStoreTypeSurugaya judge is store type is surugaya.
func IsStoreTypeSurugaya(storeType string) bool {
	if storeType == surugayaStoreType {
		return true
	}
	return false
}

// IsStoreTypeBookoff judge is store type is bookoff.
func IsStoreTypeBookoff(storeType string) bool {
	if storeType == bookoffStoreType {
		return true
	}
	return false
}

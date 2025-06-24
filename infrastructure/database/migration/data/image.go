package data

import (
	donatedItemSchema "givebox/infrastructure/database/donation/donated_item"
	imageSchema "givebox/infrastructure/database/donation/image"
	"log"

	"gorm.io/gorm"
)

func GetImages(db *gorm.DB) []imageSchema.Image {
	var donatedItems []donatedItemSchema.DonatedItem
	if err := db.Find(&donatedItems).Error; err != nil {
		log.Fatalf("could not fetch donated items: %v", err)
	}

	if len(donatedItems) == 0 {
		log.Println("no donated items found to seed images")
		return nil
	}

	var images []imageSchema.Image

	imageMap := map[string][]string{
		"Novel Harry Potter": {
			"[https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1474154022l/3._SY475_.jpg](https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1474154022l/3._SY475_.jpg)",
		},
		"Mouse Gaming Bekas": {
			"[https://api.vitech.asia/storage/produk/294861e3-4893-4b10-a9ff-6cf0dc604235.jpg](https://api.vitech.asia/storage/produk/294861e3-4893-4b10-a9ff-6cf0dc604235.jpg)",
		},
	}

	for _, item := range donatedItems {
		if urls, ok := imageMap[item.Name]; ok {
			for _, url := range urls {
				images = append(images, imageSchema.Image{
					DonatedItemID: item.ID,
					ImageURL:      url,
				})
			}
		}
	}
	return images
}

package data

import (
	"github.com/google/uuid"
	"givebox/domain/donation/donated_item"
	"givebox/infrastructure/database/donation/category"
	donatedItemSchema "givebox/infrastructure/database/donation/donated_item"
	"givebox/infrastructure/database/profile/user"
	"log"

	"gorm.io/gorm"
)

func GetDonatedItems(db *gorm.DB) []donatedItemSchema.DonatedItem {
	var users []user.User
	if err := db.Find(&users).Error; err != nil {
		log.Fatalf("could not fetch users: %v", err)
	}
	if len(users) < 3 {
		log.Fatalf("not enough users to seed donated items, please seed users first")
	}

	var categories []category.Category
	if err := db.Find(&categories).Error; err != nil {
		log.Fatalf("could not fetch categories: %v", err)
	}
	if len(categories) == 0 {
		log.Fatalf("no categories found, please seed categories first")
	}

	categoryMap := make(map[string]string)
	for _, cat := range categories {
		categoryMap[cat.Name] = cat.ID.String()
	}

	return []donatedItemSchema.DonatedItem{
		{
			DonorID:             users[0].ID,
			CategoryID:          getCategoryID(categoryMap, "Pakaian & Aksesoris"),
			Status:              donated_item.StatusOpened,
			Name:                "Kemeja Lengan Panjang",
			Description:         "Kemeja flanel lengan panjang ukuran L, kondisi masih sangat baik, jarang dipakai.",
			Condition:           4,
			QuantityDescription: "1 buah",
			PickCity:            "Jakarta",
			PickAddress:         "Jl. Jenderal Sudirman No.Kav. 52-53, RT.5/RW.3, Senayan, Kebayoran Baru",
			PickingStatus:       donated_item.PickingStatusPick,
			DeliveryTime:        "Fleksibel",
			IsUrgent:            false,
		},
		{
			DonorID:             users[1].ID,
			CategoryID:          getCategoryID(categoryMap, "Buku & Pendidikan"),
			Status:              donated_item.StatusOpened,
			Name:                "Novel Harry Potter",
			Description:         "Satu set novel Harry Potter (1-7), kondisi sampul sedikit usang tapi halaman masih lengkap dan bagus.",
			Condition:           3,
			QuantityDescription: "7 buku",
			PickCity:            "Surabaya",
			PickAddress:         "Jl. Mayjen Sungkono No.89, Gn. Sari, Dukuhpakis",
			PickingStatus:       donated_item.PickingStatusBoth,
			DeliveryTime:        "Sore hari",
			IsUrgent:            false,
		},
		{
			DonorID:             users[2].ID,
			CategoryID:          getCategoryID(categoryMap, "Elektronik"),
			Status:              donated_item.StatusOpened,
			Name:                "Mouse Gaming Bekas",
			Description:         "Mouse gaming merk Fantech, semua tombol berfungsi normal, lampu RGB masih menyala. Kabel sedikit terkelupas tapi aman.",
			Condition:           3,
			QuantityDescription: "1 buah",
			PickCity:            "Bandung",
			PickAddress:         "Jl. Ir. H. Juanda No.107, Dago, Kecamatan Coblong",
			PickingStatus:       donated_item.PickingStatusDeliver,
			DeliveryTime:        "Akhir pekan",
			IsUrgent:            true,
		},
		{
			DonorID:             users[0].ID,
			CategoryID:          getCategoryID(categoryMap, "Furnitur"),
			Status:              donated_item.StatusOpened,
			Name:                "Kursi Belajar Anak",
			Description:         "Kursi belajar anak-anak warna biru, bahan plastik tebal, kondisi kokoh dan bagus.",
			Condition:           5,
			QuantityDescription: "1 buah",
			PickCity:            "Jakarta",
			PickAddress:         "Jl. Gatot Subroto No.Kav. 38, RT.6/RW.3, Kuningan Bar., Mampang Prpt.",
			PickingStatus:       donated_item.PickingStatusPick,
			DeliveryTime:        "Konfirmasi dulu",
			IsUrgent:            false,
		},
	}
}

func getCategoryID(m map[string]string, name string) uuid.UUID {
	idStr, ok := m[name]
	if !ok {
		log.Fatalf("category %s not found in map", name)
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.Fatalf("failed to parse UUID for category %s: %v", name, err)
	}
	return id
}

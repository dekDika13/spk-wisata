package adminRepository

import (
	adminDto "backend/dto/adminDto"
	"backend/model"

	"gorm.io/gorm"
)

func (r *adminRepository) GetAllDestinations() ([]adminDto.DestinationResponseDTO, error) {
	var destinations []model.Destination
	// 1. Ambil semua data destinasi utama
	if err := r.db.Find(&destinations).Error; err != nil {
		return nil, err
	}

	// Siapkan slice untuk menampung hasil akhir
	var responseDTOs []adminDto.DestinationResponseDTO

	// 2. Loop setiap destinasi untuk mencari data galerinya
	for _, dest := range destinations {
		var galleryUrls []string

		// Cari gallery ID yang berelasi dengan destination ID saat ini
		var gallery model.Gallery
		r.db.Where("destination_id = ?", dest.DestinationID).First(&gallery)

		// Jika galeri ditemukan, cari semua URL gambarnya
		if gallery.GalleryID != 0 {
			r.db.Model(&model.Image{}).Where("gallery_id = ?", gallery.GalleryID).Pluck("image_url", &galleryUrls)
		}

		dto := adminDto.DestinationResponseDTO{
			DestinationId:    dest.DestinationID,
			Name:             dest.Name,
			Description:      dest.Description,
			Cover:            dest.Cover,
			Galery:           galleryUrls, // Diisi dengan hasil pencarian
			Toilet:           dest.Toilet,
			Parking:          dest.Parking,
			Restarea:         dest.Restarea,
			Restaurant:       dest.Restaurant,
			Price:            dest.Price,
			AssessmentResult: dest.AssessmentResult,
			Address:          dest.Address,
			Location:         dest.Location,
			// Rating bisa ditambahkan jika sudah ada logikanya
		}
		responseDTOs = append(responseDTOs, dto)
	}

	return responseDTOs, nil
}
func (u *adminRepository) GetDestinationById(id int) (adminDto.DestinationResponseDTO, error) {
	view := adminDto.DestinationResponseDTO{}

	if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).First(&view).Error; err != nil {
		return view, err
	}
	return view, nil
}

func (u *adminRepository) CreateDestination(payload adminDto.DestinationCreateDTO, urls adminDto.DestinationImageDTO) error {
	// err  := u.db.Transaction(func(tx *gorm.DB) error {
	// 	//buat galery
	// 	// simpan url images urls di galery  id yang sama
	// 	// simpan destination
	return u.db.Transaction(func(tx *gorm.DB) error {
		// 1. Membuat data destinasi utama
		destination := model.Destination{
			Name:        payload.Name,
			Description: payload.Description,
			Price:       payload.Price,
			Address:     payload.Address,
			Location:    payload.Location,
			Toilet:      payload.Toilet,
			Parking:     payload.Parking,
			Restarea:    payload.Restarea,
			Restaurant:  payload.Restaurant,
			Cover:       urls.CoverUrl, // Ambil URL cover dari DTO urls
		}

		if err := tx.Create(&destination).Error; err != nil {
			// Jika gagal, batalkan transaksi dan kembalikan error
			return err
		}

		// 2. Membuat galeri untuk destinasi yang baru dibuat
		gallery := model.Gallery{
			DestinationId: destination.DestinationID, // Hubungkan dengan ID destinasi
		}

		if err := tx.Create(&gallery).Error; err != nil {
			return err
		}

		// 3. Menyimpan gambar-gambar galeri (jika ada)
		// Cek apakah ada URL gambar yang perlu disimpan
		if len(urls.ImageUrl) > 0 {
			var images []model.Image
			for _, imgURL := range urls.ImageUrl {
				images = append(images, model.Image{
					GalleryId: gallery.GalleryID, // Hubungkan dengan ID galeri
					ImageURL:  imgURL,
				})
			}

			// Simpan semua gambar sekaligus (batch insert) agar lebih efisien
			if err := tx.Create(&images).Error; err != nil {
				return err
			}
		}

		// Jika semua proses berhasil, kembalikan nil untuk menyelesaikan transaksi (commit)
		return nil
	})

}

// func (u *adminRepository) UpdateDestination(id int, payload adminDto.DestinationUpdateDTO) error {
// 	destination := model.Destination{
// 		Name:        payload.Name,
// 		Description: payload.Description,
// 		// Image1:     payload.Image1,
// 		// Image2:     payload.Image2,
// 		// Image3:     payload.Image3,
// 		// Image4:     payload.Image4,
// 		Price:      payload.Price,
// 		Address:    payload.Address,
// 		Location:   payload.Location,
// 	}

//		if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).Updates(destination).Error; err != nil {
//			return err
//		}
//		return nil
//	}
func (u *adminRepository) DeleteDestination(id int) error {
	if err := u.db.Model(&model.Destination{}).Where("destination_id = ?", id).Delete(&model.Destination{}).Error; err != nil {
		return err
	}
	return nil
}

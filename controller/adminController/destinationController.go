package adminController

import (
	adminDto "backend/dto/adminDto"
	"backend/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (u *adminController) GetAllDestination(c echo.Context) error {
	res, err := u.adminServ.GetAllDestination()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}
func (u *adminController) GetDestinationById(c echo.Context) error {
	id := c.FormValue("destination_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
	}

	res, err := u.adminServ.GetDestinationById(utils.StringToInt(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "success",
		Code:    http.StatusOK,
		Data:    res,
	})
}

// func (u *adminController) CreateDestination(c echo.Context) error {

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		log.Printf("Error saat parsing multipart form: %v", err)
// 		return c.JSON(http.StatusBadRequest, "Tidak bisa mem-parsing form")
// 	}

// 	log.Println("--- DEBUG: FILE YANG DITERIMA SERVER ---")
// 	for key, files := range form.File {
// 		log.Printf("Key: %s, Jumlah File: %d", key, len(files))
// 		for i, file := range files {
// 			log.Printf("  -> File #%d: %s", i+1, file.Filename)
// 		}
// 	}
// 	log.Println("--------------------------------------")
// 	// !! AKHIR KODE DEBUG !!

// 	// ... sisa kode Anda dimulai dari sini (var payloads, c.Bind, dst.)

// 	// --- 1. Bind data dari form ke DTO ---
// 	// DTO ini akan menampung data teks dan metadata file (FileHeader)
// 	var payloads adminDto.DestinationCreateDTO
// 	if err := c.Bind(&payloads); err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: "Invalid form data: " + err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	// --- 2. Validasi data teks ---
// 	if err := c.Validate(payloads); err != nil {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	// --- 3. Proses Upload File Cover ---
// 	var coverURL string
// 	if payloads.Cover != nil {
// 		file, err := payloads.Cover.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to open cover image"})
// 		}
// 		defer file.Close()

// 		url, _, err := utils.UploadToCloudinary(file, uuid.New().String())
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to upload cover image"})
// 		}
// 		coverURL = url
// 	} else {
// 		// Anda bisa tambahkan validasi di sini jika cover wajib diisi
// 		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Cover image is required"})
// 	}

// 	// --- 4. Proses Upload File-File Galeri ---
// 	var galleryImageUrls []string
// 	for _, imgHeader := range payloads.Images {
// 		if imgHeader == nil {
// 			continue
// 		}

// 		file, err := imgHeader.Open()
// 		if err != nil {
// 			log.Printf("Gagal membuka file galeri: %v", err) // Log error, tapi lanjutkan
// 			continue
// 		}

// 		url, _, err := utils.UploadToCloudinary(file, uuid.New().String())
// 		if err == nil {
// 			galleryImageUrls = append(galleryImageUrls, url)
// 		}

// 		// PENTING: Tutup file secara manual di dalam loop.
// 		file.Close()
// 	}

// 	// --- 5. Siapkan DTO berisi URL untuk dikirim ke Service ---
// 	imageURLsForService := adminDto.DestinationImageDTO{
// 		CoverUrl: coverURL,
// 		ImageUrl: galleryImageUrls,
// 	}

// 	// --- 6. Panggil Service dengan data yang sudah matang ---
// 	// Pastikan service Anda menerima 2 argumen: payload teks dan DTO URL
// 	if err := u.adminServ.CreateDestination(payloads, imageURLsForService); err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

//		// --- 7. Kirim Respons Sukses ---
//		return c.JSON(http.StatusCreated, utils.Response{
//			Message: "destination created successfully",
//			Code:    http.StatusCreated,
//		})
//	}
func (u *adminController) CreateDestination(c echo.Context) error {
	// --- 1. Ambil file cover SECARA LANGSUNG ---
	coverFileHeader, err := c.FormFile("cover") // Langsung baca dari form key "cover"
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Cover image is required"})
	}

	// --- 2. Ambil file galeri SECARA LANGSUNG ---
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: "Invalid form data"})
	}
	galleryFiles := form.File["images"] // Langsung baca dari form key "images"

	// --- 3. Bind SISA data (hanya teks) ---
	var payloads adminDto.DestinationCreateDTO
	if err := c.Bind(&payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
	}

	// --- 4. Validasi data teks ---
	if err := c.Validate(payloads); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
	}

	// --- 5. Proses Upload (logika ini tetap sama) ---
	// Proses Cover
	file, err := coverFileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to open cover image"})
	}
	defer file.Close()
	coverURL, _, err := utils.UploadToCloudinary(file, uuid.New().String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: "Failed to upload cover image"})
	}

	// Proses Galeri
	var galleryImageUrls []string
	for _, imgHeader := range galleryFiles {
		file, _ := imgHeader.Open()
		url, _, _ := utils.UploadToCloudinary(file, uuid.New().String())
		galleryImageUrls = append(galleryImageUrls, url)
		file.Close()
	}

	// --- 6. Panggil Service (logika ini tetap sama) ---
	imageURLsForService := adminDto.DestinationImageDTO{
		CoverUrl: coverURL,
		ImageUrl: galleryImageUrls,
	}
	if err := u.adminServ.CreateDestination(payloads, imageURLsForService); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
	}

	// --- 7. Respons Sukses ---
	return c.JSON(http.StatusCreated, utils.Response{
		Message: "destination created successfully",
	})
}

// func (u *adminController) UpdateDestination(c echo.Context) error {
// 	id := c.FormValue("destination_id")
// 	if id == "" {
// 		return c.JSON(http.StatusBadRequest, utils.Response{
// 			Message: "id is required",
// 			Code:    http.StatusBadRequest,
// 		})
// 	}

// 	var payload adminDto.DestinationUpdateDTO
// 	payload.Name = c.FormValue("name")
// 	payload.Description = c.FormValue("description")
// 	payload.Image1 = c.FormValue("image1")
// 	payload.Image2 = c.FormValue("image2")
// 	payload.Image3 = c.FormValue("image3")
// 	payload.Image4 = c.FormValue("image4")
// 	payload.Price = utils.StringToInt(c.FormValue("price"))
// 	payload.Address = c.FormValue("address")
// 	payload.Location = c.FormValue("location")

// 	if err := u.adminServ.UpdateDestination(utils.StringToInt(id), payload); err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.Response{
// 			Message: err.Error(),
// 			Code:    http.StatusInternalServerError,
// 		})
// 	}

//		return c.JSON(http.StatusOK, utils.Response{
//			Message: "destination updated successfully",
//			Code:    http.StatusOK,
//		})
//	}
func (u *adminController) DeleteDestination(c echo.Context) error {
	id := c.FormValue("destination_id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			Message: "id is required",
			Code:    http.StatusBadRequest,
		})
	}

	if err := u.adminServ.DeleteDestination(utils.StringToInt(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Response{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Message: "destination deleted successfully",
		Code:    http.StatusOK,
	})
}

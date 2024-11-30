package main

import (
	"encoding/json"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var imagesDIR string = Env.GetString("SRV_DIR", "") + "/ServerImages/"

func (app *application) createListingImage(w http.ResponseWriter, r *http.Request) {
	tokenUserId, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return // Ensure early exit
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	basePath := Env.GetString("SRV_DIR", "/home/adam-elhassan/Desktop/ServerFiles")
	serverImagesDir := filepath.Join(basePath, "ServerImages")

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			if !allowedExtensions[ext] {
				http.Error(w, "Invalid file type: "+fileHeader.Filename, http.StatusBadRequest)
				return
			}

			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, "Failed to open file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			listingIDStr := chi.URLParam(r, "listing_id")
			listingID, err := strconv.Atoi(listingIDStr)
			if err != nil {
				http.Error(w, "Invalid listing ID", http.StatusBadRequest)
				return
			}

			if listingID != 0 {
				DbListing, err := app.Service.Listings.GetByID(r.Context(), listingID)
				if err != nil {
					http.Error(w, "Invalid listing ID", http.StatusBadRequest)
					return
				}

				if DbListing.UserID != tokenUserId {
					http.Error(w, "Cannot upload image for another user's listing", http.StatusUnauthorized)
					return
				}
			}

			newFileName := uuid.New().String() + ext
			filePath, err := Utils.SaveFile(file, serverImagesDir, newFileName)
			if err != nil {
				http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = app.Service.Images.AddImage(r.Context(), newFileName, tokenUserId, listingID)
			if err != nil {
				http.Error(w, "Failed to insert image record: "+err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "File %s uploaded successfully as %s\n", fileHeader.Filename, filePath)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All image files uploaded successfully and stored in ServerImages"))
}

func (app *application) createProfileImage(w http.ResponseWriter, r *http.Request) {
	tokenUserId, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return // Ensure early exit
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	basePath := Env.GetString("SRV_DIR", "/home/adam-elhassan/Desktop/ServerFiles")
	serverImagesDir := filepath.Join(basePath, "ServerImages")

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			if !allowedExtensions[ext] {
				http.Error(w, "Invalid file type: "+fileHeader.Filename, http.StatusBadRequest)
				return
			}

			file, err := fileHeader.Open()
			if err != nil {
				http.Error(w, "Failed to open file: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			userIDStr := chi.URLParam(r, "user_id")
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				http.Error(w, "Invalid listing ID", http.StatusBadRequest)
				return
			}

			if userID != tokenUserId {
				http.Error(w, "Cannot upload image for another user's listing", http.StatusUnauthorized)
			}

			newFileName := uuid.New().String() + ext
			filePath, err := Utils.SaveFile(file, serverImagesDir, newFileName)
			if err != nil {
				http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
				return
			}

			imageId, err := app.Service.Images.AddImage(r.Context(), newFileName, tokenUserId, 0)

			fmt.Println(imageId)

			if err != nil {
				http.Error(w, "Failed to insert image record: "+err.Error(), http.StatusInternalServerError)
				return
			}

			err = app.Service.Images.UpdateImageProfilePictureStatus(r.Context(), imageId, userID)

			if err != nil {
				http.Error(w, "Failed to set on profile: "+err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "File %s uploaded successfully as %s\n", fileHeader.Filename, filePath)
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All image files uploaded successfully and stored in ServerImages"))
}

func (app *application) DeleteImage(w http.ResponseWriter, r *http.Request) {

	tokenUserId, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return // Ensure early exit
	}

	// Extract image ID from the URL
	imageIDStr := chi.URLParam(r, "image_id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	image, err := app.Service.Images.GetImageByID(r.Context(), imageID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve image: %v", err), http.StatusInternalServerError)
	}

	if image.UserID != tokenUserId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	// Call the service to delete the image
	err = app.Service.Images.DeleteImage(r.Context(), imageID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not delete image: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image deleted successfully"))
}

func (app *application) UpdateImage(w http.ResponseWriter, r *http.Request) {

	tokenUserId, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return // Ensure early exit
	}

	// Extract image ID from the URL
	imageIDStr := chi.URLParam(r, "image_id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	// Extract show_on_profile from the URL (as 1 or 0)
	showOnProfileStr := chi.URLParam(r, "show_on_profile")
	showOnProfile, err := strconv.Atoi(showOnProfileStr)
	if err != nil || (showOnProfile != 0 && showOnProfile != 1) {
		http.Error(w, "Invalid value for show_on_profile (use 1 for true, 0 for false)", http.StatusBadRequest)
		return
	}

	image, err := app.Service.Images.GetImageByID(r.Context(), imageID)

	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve image: %v", err), http.StatusInternalServerError)
	}

	if image.UserID != tokenUserId {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	// Convert 1 -> true and 0 -> false
	var showOnProfileBool bool
	if showOnProfile == 1 {
		showOnProfileBool = true
	} else {
		showOnProfileBool = false
	}

	// Call the service to update the image
	err = app.Service.Images.UpdateImageProfileStatus(r.Context(), imageID, showOnProfileBool)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not update image: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image updated successfully"))
}

func (app *application) GetImageByID(w http.ResponseWriter, r *http.Request) {
	// Extract image ID from the URL
	imageIDStr := chi.URLParam(r, "image_id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
	}

	// Get the image from the service
	image, err := app.Service.Images.GetImageByID(r.Context(), imageID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve image: %v", err), http.StatusInternalServerError)
		return
	}

	// Get the image path
	imagePath := imagesDIR + image.URL

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Detect the content type
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		http.Error(w, "Error reading image file", http.StatusInternalServerError)
		return
	}
	contentType := http.DetectContentType(buffer)

	// Set the Content-Type header
	w.Header().Set("Content-Type", contentType)

	// Reset the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		http.Error(w, "Error resetting file pointer", http.StatusInternalServerError)
		return
	}

	// Serve the image file
	http.ServeContent(w, r, image.URL, time.Now(), file)

}

func (app *application) GetImageByUUID(w http.ResponseWriter, r *http.Request) {
	// Extract image ID from the URL
	imageIDStr := chi.URLParam(r, "image_id")

	// Get the image path
	imagePath := imagesDIR + imageIDStr

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Detect the content type
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		http.Error(w, "Error reading image file", http.StatusInternalServerError)
		return
	}
	contentType := http.DetectContentType(buffer)

	// Set the Content-Type header
	w.Header().Set("Content-Type", contentType)

	// Reset the file pointer to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		http.Error(w, "Error resetting file pointer", http.StatusInternalServerError)
		return
	}

	// Serve the image file
	http.ServeContent(w, r, imagePath, time.Now(), file)

}

func (app *application) GetImagesByListingID(w http.ResponseWriter, r *http.Request) {
	// Extract listing ID from the URL
	listingIDStr := chi.URLParam(r, "listing_id")
	listingID, err := strconv.Atoi(listingIDStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	// Get images for the listing from the service
	images, err := app.Service.Images.GetImagesByListingID(r.Context(), listingID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve images for listing: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(images); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (app *application) GetImagesByUserID(w http.ResponseWriter, r *http.Request) {
	tokenUserId, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	// Extract user ID from the URL
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if tokenUserId != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	// Get images for the user from the service
	images, err := app.Service.Images.GetImagesByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve images for user: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(images); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (app *application) GetImagesByUserProfile(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the URL
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get images for the user with profile visibility set to true
	images, err := app.Service.Images.GetImagesByUserProfile(r.Context(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve images for user profile: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(images); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

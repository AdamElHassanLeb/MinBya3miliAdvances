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

// @Summary		Upload images for a listing
// @Description	Upload one or more images for a specific listing. The user must be authorized and the listing must belong to them.
// @Tags			images
// @Accept			multipart/form-data
// @Produce		json
// @Param			listing_id	path		int		true	"Listing ID"
// @Param			images		formData	file	true	"Images to upload"
// @Security		BearerAuth
// @Success		200	{string}	string	"All image files uploaded successfully and stored in ServerImages"
// @Failure		400	{string}	string	"Invalid file type"
// @Failure		401	{string}	string	"Unauthorized"
// @Failure		500	{string}	string	"Internal server error"
// @Router			/listings/{listing_id}/images [post]
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

	// Ensure the directory exists
	if err := os.MkdirAll(serverImagesDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

// @Summary		Upload profile image
// @Description	Upload an image to be set as a user's profile picture. The user must be authorized and own the profile.
// @Tags			images
// @Accept			multipart/form-data
// @Produce		json
// @Param			user_id	path		int		true	"User ID"
// @Param			image	formData	file	true	"Profile image to upload"
// @Security		BearerAuth
// @Success		200	{string}	string	"Profile image uploaded successfully and stored in ServerImages"
// @Failure		400	{string}	string	"Invalid file type"
// @Failure		401	{string}	string	"Unauthorized"
// @Failure		500	{string}	string	"Internal server error"
// @Router			/users/{user_id}/profile/image [post]
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

	// Ensure the directory exists
	if err := os.MkdirAll(serverImagesDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create directory: "+err.Error(), http.StatusInternalServerError)
		return
	}

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

// DeleteImage @Summary Delete an image
//
//	@Description	Delete an image from the server. The user must be authorized and own the image.
//	@Tags			images
//	@Produce		json
//	@Param			image_id	path	int	true	"Image ID"
//	@Security		BearerAuth
//	@Success		200	{string}	string	"Image deleted successfully"
//	@Failure		400	{string}	string	"Invalid image ID"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/images/{image_id} [delete]
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

// UpdateImage @Summary Update an image's profile status
//
//	@Description	Update whether an image should be shown on the user's profile. The user must be authorized and own the image.
//	@Tags			images
//	@Produce		json
//	@Param			image_id		path	int	true	"Image ID"
//	@Param			show_on_profile	path	int	true	"Show on profile (1 for true, 0 for false)"
//	@Security		BearerAuth
//	@Success		200	{string}	string	"Image updated successfully"
//	@Failure		400	{string}	string	"Invalid image ID or show_on_profile value"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/images/{image_id}/{show_on_profile} [put]
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

// GetImageByID @Summary Get an image by ID
//
//	@Description	Retrieve an image by its ID. Returns the image content along with its content type.
//	@Tags			images
//	@Produce		octet-stream
//	@Param			image_id	path		int		true	"Image ID"
//	@Success		200			{file}		string	"Image file"
//	@Failure		400			{string}	string	"Invalid image ID"
//	@Failure		404			{string}	string	"Image not found"
//	@Failure		500			{string}	string	"Internal server error"
//	@Router			/images/{image_id} [get]
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
	http.ServeContent(w, r, imagePath, time.Now(), file)
}

// GetImageByUUID @Summary Get an image by UUID
//
//	@Description	Retrieve an image by its UUID. Returns the image content along with its content type.
//	@Tags			images
//	@Produce		octet-stream
//	@Param			image_id	path		string	true	"Image UUID"
//	@Success		200			{file}		string	"Image file"
//	@Failure		404			{string}	string	"Image not found"
//	@Failure		500			{string}	string	"Internal server error"
//	@Router			/images/uuid/{image_id} [get]
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

// GetImagesByListingID @Summary Get all images for a specific listing
//
//	@Description	Retrieve all images associated with a specific listing.
//	@Tags			images
//	@Produce		json
//	@Param			listing_id	path		int		true	"Listing ID"
//	@Success		200			{array}		Image	"List of images"
//	@Failure		400			{string}	string	"Invalid listing ID"
//	@Failure		500			{string}	string	"Internal server error"
//	@Router			/listings/{listing_id}/images [get]
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

// GetImagesByUserID @Summary Get all images for a specific user
//
//	@Description	Retrieve all images associated with a specific user.
//	@Tags			images
//	@Produce		json
//	@Param			user_id	path	int	true	"User ID"
//	@Security		BearerAuth
//	@Success		200	{array}		Image	"List of images"
//	@Failure		400	{string}	string	"Invalid user ID"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/users/{user_id}/images [get]
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

// GetImagesByUserProfile @Summary Get images with profile visibility for a specific user
//
//	@Description	Retrieve all images with profile visibility set to true for a specific user.
//	@Tags			images
//	@Produce		json
//	@Param			user_id	path		int		true	"User ID"
//	@Success		200		{array}		Image	"List of images"
//	@Failure		400		{string}	string	"Invalid user ID"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/users/{user_id}/profile/images [get]
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

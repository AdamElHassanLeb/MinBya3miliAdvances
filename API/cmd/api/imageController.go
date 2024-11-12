package main

import (
	"archive/zip"
	"fmt"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Env"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var imagesDIR string = Env.GetString("SRV_DIR", "") + "/ServerImages/"

func (app *application) createImage(w http.ResponseWriter, r *http.Request) {
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

			err = app.Service.Images.AddImage(r.Context(), newFileName, tokenUserId, listingID)
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

	// Serve the image file
	imagePath := imagesDIR + image.URL
	http.ServeFile(w, r, imagePath)
}

// Utility function to create a ZIP of images
func createZip(w http.ResponseWriter, imagePaths []string) error {
	// Set response header for zip file download
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=images.zip")

	// Create a new zip writer and write to the HTTP response writer
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// Add images to the zip file
	for _, imagePath := range imagePaths {
		// Open the image file
		file, err := os.Open(imagePath)
		if err != nil {
			return fmt.Errorf("error opening file: %v", err)
		}
		defer file.Close()

		// Create a new file inside the zip archive
		zipFile, err := zipWriter.Create(filepath.Base(imagePath))
		if err != nil {
			return fmt.Errorf("error creating zip file: %v", err)
		}

		// Copy the image content into the zip
		_, err = io.Copy(zipFile, file)
		if err != nil {
			return fmt.Errorf("error copying file content into zip: %v", err)
		}
	}
	return nil
}

// Get all images for a listing and return them as a ZIP
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

	// Collect the image paths
	var imagePaths []string
	for _, img := range images {
		imagePaths = append(imagePaths, imagesDIR+img.URL)
	}

	// Create a ZIP file containing all the images
	if len(imagePaths) > 0 {
		err := createZip(w, imagePaths)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating ZIP file: %v", err), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "No images found for this listing", http.StatusNotFound)
	}
}

// Get all images for a user and return them as a ZIP
func (app *application) GetImagesByUserID(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the URL
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Get images for the user from the service
	images, err := app.Service.Images.GetImagesByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not retrieve images for user: %v", err), http.StatusInternalServerError)
		return
	}

	// Collect the image paths
	var imagePaths []string
	for _, img := range images {
		imagePaths = append(imagePaths, imagesDIR+img.URL)
	}

	// Create a ZIP file containing all the images
	if len(imagePaths) > 0 {
		err := createZip(w, imagePaths)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating ZIP file: %v", err), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "No images found for this user", http.StatusNotFound)
	}
}

// Get all images for a user with profile visibility set to true and return them as a ZIP
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

	// Collect the image paths
	var imagePaths []string
	for _, img := range images {
		imagePaths = append(imagePaths, imagesDIR+img.URL)
	}

	// Create a ZIP file containing all the images
	if len(imagePaths) > 0 {
		err := createZip(w, imagePaths)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating ZIP file: %v", err), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "No profile images found for this user", http.StatusNotFound)
	}
}

func (app *application) DeleteImage(w http.ResponseWriter, r *http.Request) {
	// Extract image ID from the URL
	imageIDStr := chi.URLParam(r, "image_id")
	imageID, err := strconv.Atoi(imageIDStr)
	if err != nil {
		http.Error(w, "Invalid image ID", http.StatusBadRequest)
		return
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

	// Convert 1 -> true and 0 -> false
	var showOnProfileBool bool
	if showOnProfile == 1 {
		showOnProfileBool = true
	} else {
		showOnProfileBool = false
	}

	// Call the service to update the image
	err = app.Service.Images.UpdateImage(r.Context(), imageID, showOnProfileBool)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not update image: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Image updated successfully"))
}

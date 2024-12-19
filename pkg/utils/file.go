package utils

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"eternal/pkg/logger"

	"go.uber.org/zap"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}

	return nil
}

// WriteBytesToFile writes a byte slice to a specified file path.
func CreateFile(filePath string, data []byte) error {
	// Open or create the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write the byte slice to the file
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

// ReadFileAsBytes reads the contents of a file and returns it as a byte slice.
func ReadFile(filePath string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return data, nil
}

func generateUniqueFilePath(dir, filename string) (string, string, error) {
	baseName := strings.TrimSuffix(filename, path.Ext(filename))
	ext := path.Ext(filename)
	uniquePath := filepath.Join(dir, filename)
	uniqueName := filename

	// Check if the file exists and generate a new path with a suffix if necessary
	counter := 1
	for {
		if _, err := os.Stat(uniquePath); os.IsNotExist(err) {
			return uniquePath, uniqueName, nil
		}
		uniqueName = fmt.Sprintf("%s-%d%s", baseName, counter, ext)
		uniquePath = filepath.Join(dir, uniqueName)
		counter++
	}
}

func UnzipFile(zipBytes []byte, destDir string) error {
	// Create a reader from the byte slice
	zipReader, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return err
	}

	// Iterate through the files in the ZIP archive
	for _, file := range zipReader.File {
		// Open the file inside the ZIP archive
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// Create the full path for the file
		fpath := filepath.Join(destDir, file.Name)

		// Create directories if necessary
		if file.FileInfo().IsDir() {
			os.MkdirAll(fpath, file.Mode())
			continue
		}

		// Create the file
		f, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Copy the file content to the new file
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func DownloadFile(fileURL, filePath string) (string, string, error) {
	// Parse the URL to extract the path
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid URL: %w", err)
	}

	// Extract the filename from the URL path
	filename := path.Base(parsedURL.Path)
	if filename == "" {
		return "", "", fmt.Errorf("could not extract filename from URL path")
	}

	// Generate a unique file path if necessary
	uniqueFilePath, nFileName, err := generateUniqueFilePath(filePath, filename)
	if err != nil {
		return "", "", err
	}

	out, err := os.Create(uniqueFilePath)
	if err != nil {
		return "", "", err
	}
	defer out.Close()

	// Send the HTTP request
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", "", err
	}

	return uniqueFilePath, nFileName, nil
}

func ZipDir(ctx context.Context, filePath, outFilePath string) ([]byte, error) {
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return nil, err
	}

	defer outFile.Close()
	w := zip.NewWriter(outFile)
	addFiles(ctx, w, filePath, "")

	if err = w.Close(); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(outFilePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func addFiles(ctx context.Context, w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := os.ReadDir(basePath)
	if err != nil {
		logger.GetLoggerInstanceFromContext(ctx).Error("addFiles", zap.Error(err))
	}

	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(basePath, file.Name())
			dat, err := os.ReadFile(path)
			if err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("addFiles", zap.Error(err))
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				logger.AtLog.Error(err)
			}

			if _, err = f.Write(dat); err != nil {
				logger.GetLoggerInstanceFromContext(ctx).Error("addFiles", zap.Error(err))
			}
			continue
		}

		if file.IsDir() {
			// Recurse
			newBase := basePath + file.Name() + "/"
			logger.GetLoggerInstanceFromContext(ctx).Info("Recursing and Adding SubDir", zap.String("file_name", file.Name()), zap.String("new_base", newBase))
			addFiles(ctx, w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

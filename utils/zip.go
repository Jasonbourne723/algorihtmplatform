package utils

import (
	"archive/zip"
	"io"
	"os"
	"path"

	ezip "github.com/alexmullins/zip"
)

type zipUtil struct {
}

var ZipUtil = &zipUtil{}

func (z *zipUtil) Compress(sourceDir string, destPath string) error {
	zipfile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	if err := writeToZip(sourceDir, "", zipWriter); err != nil {
		return err
	}
	return nil

}

func writeToZip(sourceDir string, dirName string, zipWriter *zip.Writer) error {
	sourceFiles, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	for _, ea := range sourceFiles {
		if ea.IsDir() {

			writeToZip(path.Join(sourceDir, ea.Name()), ea.Name(), zipWriter)

			continue
		}
		zipFileName := path.Join(dirName, ea.Name())
		w, err := zipWriter.Create(zipFileName)
		if err != nil {
			return err
		}
		filePath := path.Join(sourceDir, ea.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err = io.Copy(w, f); err != nil {
			return err
		}
	}
	return nil
}

func (z *zipUtil) DeCompress(archiveName string, destName string) error {
	archive, err := zip.OpenReader(archiveName)
	if err != nil {
		return err
	}
	defer archive.Close()

	if err := os.MkdirAll(destName, os.ModePerm); err != nil {
		return err
	}

	for _, f := range archive.File {
		destPath := path.Join(destName, f.Name)
		os.MkdirAll(path.Dir(destPath), os.ModePerm)
		// if f.FileInfo().IsDir() {
		// 	os.MkdirAll(destPath, os.ModePerm)
		// 	continue
		// } else {

		// }

		destFile, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		file, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(destFile, file); err != nil {
			return err
		}

		destFile.Close()
		file.Close()
	}
	return nil
}

func (z *zipUtil) CompressEncryption(sourceDir string, destPath string, passwd string) error {

	zipfile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	zipWriter := ezip.NewWriter(zipfile)
	defer zipWriter.Close()

	if err := writeToEzip(sourceDir, "", zipWriter, passwd); err != nil {
		return err
	}

	return nil
}

func writeToEzip(sourceDir string, dirName string, zipWriter *ezip.Writer, passwd string) error {
	sourceFiles, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}
	for _, ea := range sourceFiles {
		if ea.IsDir() {

			writeToEzip(path.Join(sourceDir, ea.Name()), ea.Name(), zipWriter, passwd)

			continue
		}
		zipFileName := path.Join(dirName, ea.Name())
		w, err := zipWriter.Encrypt(zipFileName, passwd)
		if err != nil {
			return err
		}
		filePath := path.Join(sourceDir, ea.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err = io.Copy(w, f); err != nil {
			return err
		}
	}
	return nil
}

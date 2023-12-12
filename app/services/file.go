package services

import (
	"algorithmplatform/app/common/dto"
	"algorithmplatform/app/models"
	"algorithmplatform/global"
	"algorithmplatform/utils"
	"encoding/json"
	"errors"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

const (
	OPERATOR_PATH          = "Operator"
	ALGORITHM_PACKAGE_PATH = "AlgorithmPackage"
	ALGORITHM_PATH         = "Algorithm"
	STARTUP_FILENAME       = "startup.py"
	ROOT_PATH              = "./storage"
	TEMP_PATH              = "Temp"
)

var locker sync.Mutex
var algorithmLocker sync.Mutex

type fileService struct {
}

var FileService = new(fileService)

// 获取目录结构
func getNodeInfo(dir string) ([]dto.FileDto, error) {

	fileDtos := make([]dto.FileDto, 0, 10)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {

			if file.Name() == "__pycache__" {
				continue
			} else {
				dirPath := path.Join(dir, file.Name())
				subDirs, err := getNodeInfo(dirPath)
				if err != nil {
					continue
				}
				fileDtos = append(fileDtos, dto.FileDto{
					IsDirectory: true,
					Name:        file.Name(),
					Path:        dirPath,
					Children:    subDirs,
				})
			}

		} else {
			ext := path.Ext(file.Name())
			if ext == ".zip" || ext == ".rar" {
				continue
			} else {
				fileDtos = append(fileDtos, dto.FileDto{
					IsDirectory: false,
					Name:        file.Name(),
					Path:        path.Join(dir, file.Name()),
				})
			}
		}
	}
	return fileDtos, nil
}

func downFileAndDeCompress(fileKey string, destDirectory string) {
	if !existDir(destDirectory) {
		os.MkdirAll(destDirectory, os.ModePerm)
	}
	zipPath := path.Join(destDirectory, fileKey)
	utils.NewEosClient().DownLoad(fileKey, zipPath)
	utils.ZipUtil.DeCompress(zipPath, destDirectory)
	os.Remove(zipPath)
}

// 判断目录是否存在
func existDir(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		return false
	} else {
		return true
	}
}

// 获取算子文件目录路径
func getOperatorDir(operatorId int64) (dirName string) {
	return path.Join(ROOT_PATH, OPERATOR_PATH, strconv.Itoa(int(operatorId)))
}

// 获取算法文件目录路径
func getAlgorithmDir(algorithmId int64) (dirName string) {
	return path.Join(ROOT_PATH, ALGORITHM_PATH, strconv.Itoa(int(algorithmId)))
}

// 获取算法包目录路径
func getAlgorithmPackageDir(algoithmPackageId int64) (dirName string) {
	return path.Join(ROOT_PATH, ALGORITHM_PACKAGE_PATH, strconv.Itoa(int(algoithmPackageId)))
}

func generateOperatorDirectory(algorithmDir string, operatorId int64, withOperatorName bool) {
	var operator models.Operator
	global.App.DB.First(&operator, operatorId)

	if withOperatorName {
		operatorDir := path.Join(algorithmDir, operator.Name)
		downFileAndDeCompress(operator.FilePath, operatorDir)
	} else {
		downFileAndDeCompress(operator.FilePath, algorithmDir)
	}
}

func generateAlgorithmDirectory(algorithmDir string, algorithmId int64) error {

	if !existDir(algorithmDir) {
		algorithmLocker.Lock()
		defer algorithmLocker.Unlock()
		if !existDir(algorithmDir) {
			if algorithm, err := AlgorithmService.GetOne(algorithmId); err != nil {
				return err
			} else {
				if err := os.MkdirAll(algorithmDir, os.ModePerm); err != nil {
					return err
				}
				//生成输入算子目录
				for _, a := range algorithm.InputAlgorithm {
					generateOperatorDirectory(algorithmDir, a.OperatorId, true)
				}
				//生成计算算子目录
				if algorithm.AlgoAlgorithm.OperatorId > 0 {
					generateOperatorDirectory(algorithmDir, algorithm.AlgoAlgorithm.OperatorId, true)
				}
				//生成输出算子目录
				for _, a := range algorithm.OutputAlgorithm {
					generateOperatorDirectory(algorithmDir, a.OperatorId, true)
				}
				var frameOperator models.Operator
				if err := global.App.DB.Where("type = ?", 10).First(&frameOperator).Error; err != nil {
					return err
				}
				generateOperatorDirectory(algorithmDir, frameOperator.Id, false)
			}
		}
	}
	return nil
}

func (f *fileService) UpdateOperatorFile(filePath string, content string, operatorId int64) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	operatorDir := getOperatorDir(operatorId)
	fileName := uuid.New().String() + ".zip"
	if err = os.MkdirAll(path.Join(ROOT_PATH, TEMP_PATH), os.ModePerm); err != nil {
		return "", err
	}
	zipPath := path.Join(ROOT_PATH, TEMP_PATH, fileName)
	utils.ZipUtil.Compress(operatorDir, zipPath)
	zipFile, err := os.Open(zipPath)
	if err != nil {
		return "", err
	}
	utils.NewEosClient().Upload(fileName, zipFile)
	os.RemoveAll(zipPath)
	return fileName, nil
}

// 获取算子文件目录结构
func (f *fileService) GetFilesByOperatorId(operatorId int64) ([]dto.FileDto, error) {
	operatorDir := getOperatorDir(operatorId)
	var operator models.Operator
	global.App.DB.First(&operator, operatorId)

	if !existDir(operatorDir) {
		locker.Lock()
		defer locker.Unlock()
		if !existDir(operatorDir) {
			downFileAndDeCompress(operator.FilePath, operatorDir)
		}
	}
	return getNodeInfo(operatorDir)
}

// 获取算法文件目录结构
func (f *fileService) GetFilesByAlgorithmId(algorithmId int64) ([]dto.FileDto, error) {
	algorithmPath := getAlgorithmDir(algorithmId)
	err := generateAlgorithmDirectory(algorithmPath, algorithmId)
	if err != nil {
		return nil, err
	}
	return getNodeInfo(algorithmPath)
}

// 清除算法缓存
func (f *fileService) CleanCache(algorithmId int64) {
	var dir = getAlgorithmDir(algorithmId)
	os.RemoveAll(dir)
}

func (f *fileService) CleanCacheByOperatorId(operatorId int64) {
	var dir = getOperatorDir(operatorId)
	os.RemoveAll(dir)
}

func (f *fileService) GetStartupPath(algorithmId int64) string {
	algorithmDir := getAlgorithmDir(algorithmId)
	generateAlgorithmDirectory(algorithmDir, algorithmId)
	return path.Join(algorithmDir, STARTUP_FILENAME)
}

func (f *fileService) UploadFile(file *multipart.FileHeader) (string, error) {
	if file.Size > 1024*1024*20 {
		return "", errors.New("文件大小不能超过20MB")
	}
	fileExt := path.Ext(file.Filename)
	if fileExt != ".zip" && fileExt != ".py" {
		return "", errors.New("仅支持上传zip,py文件")
	}
	filename := uuid.New().String()
	fileKey := filename + ".zip"

	if fileExt == ".zip" {
		if fileReader, err := file.Open(); err != nil {
			return "", err
		} else {
			defer fileReader.Close()
			if err = utils.NewEosClient().Upload(fileKey, fileReader); err != nil {
				return "", err
			}
		}
	} else {
		tempDir := path.Join(ROOT_PATH, TEMP_PATH, filename)
		if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
			return "", err
		}
		filePath := path.Join(tempDir, file.Filename)
		tempFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return "", err
		}
		defer tempFile.Close()
		fileReader, err := file.Open()
		if err != nil {
			return "", err
		}
		defer fileReader.Close()
		bytes := make([]byte, file.Size)

		_, err = fileReader.Read(bytes)
		if err != nil {
			return "", err
		}
		_, err = tempFile.Write(bytes)
		if err != nil {
			return "", err
		}

		zipPath := path.Join(ROOT_PATH, TEMP_PATH, fileKey)
		err = utils.ZipUtil.Compress(tempDir, zipPath)
		if err != nil {
			return "", nil
		}
		zipReader, err := os.Open(zipPath)
		if err != nil {
			return "", err
		}
		defer zipReader.Close()
		err = utils.NewEosClient().Upload(fileKey, zipReader)
		if err != nil {
			return "", err
		}
	}
	return fileKey, nil
}

func (f *fileService) GetPackageFilePath(algorithmPackage *models.AlgorithmPackage) (string, error) {
	fileName := algorithmPackage.Name + "_" + algorithmPackage.CreateDate.Format("20060102150405")
	filekey := fileName + ".zip"
	packageDir := path.Join(ROOT_PATH, ALGORITHM_PACKAGE_PATH, fileName)
	if !existDir(packageDir) {

		zipPath := path.Join(ROOT_PATH, ALGORITHM_PACKAGE_PATH, filekey)
		if err := generateAlgorithmDirectory(packageDir, algorithmPackage.AlgorithmId); err != nil {
			return "", err
		}

		para, err := AlgorithmBuilder.BuildForPacket(algorithmPackage)
		if err != nil {
			return "", err
		}

		config := dto.AlgoithmPackageParams{
			CompanyId: algorithmPackage.CompanyId,
			CraneId:   algorithmPackage.CraneId,
			Cron:      algorithmPackage.Cron,
			Perild:    algorithmPackage.Period,
			Para:      *para,
		}
		configStr, _ := json.Marshal(config)
		configPath := path.Join(packageDir, "config.json")
		configFile, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return "", err
		}
		defer configFile.Close()

		configFile.WriteString(string(configStr))
		if err := utils.ZipUtil.CompressEncryption(packageDir, zipPath, "LSAREWQR24214JFLSAF2"); err != nil {
			return "", err
		}
		//utils.ZipUtil.Compress(packageDir, zipPath)
		zipFile, err := os.Open(zipPath)
		if err != nil {
			return "", err
		}
		defer zipFile.Close()
		utils.NewEosClient().Upload(filekey, zipFile)

		os.Remove(zipPath)
		os.RemoveAll(packageDir)
	}
	return filekey, nil
}

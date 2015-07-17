package bitrise

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	log "github.com/Sirupsen/logrus"
	models "github.com/bitrise-io/bitrise-cli/models/models_1_0_0"
	"github.com/bitrise-io/go-pathutil/pathutil"
)

// NewError ...
func NewError(a ...interface{}) error {
	errStr := fmt.Sprint(a...)
	return errors.New(errStr)
}

// NewErrorf ...
func NewErrorf(format string, a ...interface{}) error {
	errStr := fmt.Sprintf(format, a...)
	return errors.New(errStr)
}

// ReadBitriseConfigYML ...
func ReadBitriseConfigYML(pth string) (models.BitriseConfigModel, error) {
	if isExists, err := pathutil.IsPathExists(pth); err != nil {
		return models.BitriseConfigModel{}, err
	} else if !isExists {
		return models.BitriseConfigModel{}, NewErrorf("No file found at path", pth)
	}

	bytes, err := ioutil.ReadFile(pth)
	if err != nil {
		return models.BitriseConfigModel{}, err
	}
	var bitriseConfigYML models.BitriseConfigYMLModel
	if err := yaml.Unmarshal(bytes, &bitriseConfigYML); err != nil {
		return models.BitriseConfigModel{}, err
	}

	return bitriseConfigYML.ToBitriseConfigModel()
}

// ReadSpecStep ...
func ReadSpecStep(pth string) (models.StepModel, error) {
	if isExists, err := pathutil.IsPathExists(pth); err != nil {
		return models.StepModel{}, err
	} else if !isExists {
		return models.StepModel{}, NewErrorf("No file found at path", pth)
	}

	bytes, err := ioutil.ReadFile(pth)
	if err != nil {
		return models.StepModel{}, err
	}
	var specStep models.StepModel
	if err := yaml.Unmarshal(bytes, &specStep); err != nil {
		return models.StepModel{}, err
	}

	return specStep, nil
}

// WriteStringToFile ...
func WriteStringToFile(pth string, fileCont string) error {
	return WriteBytesToFile(pth, []byte(fileCont))
}

// WriteBytesToFile ...
func WriteBytesToFile(pth string, fileCont []byte) error {
	if pth == "" {
		return errors.New("No path provided")
	}

	file, err := os.Create(pth)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorln("Failed to close file:", err)
		}
	}()

	if _, err := file.Write(fileCont); err != nil {
		return err
	}

	return nil
}

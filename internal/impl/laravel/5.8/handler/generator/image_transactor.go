package generator

import (
	"asher/internal/api/codebuilder/php/core"
	"fmt"
	"strings"
)

type ImageTransactor struct {
	transactorGen *TransactorGenerator
}

func NewImageTransactor(transactorGen *TransactorGenerator) *ImageTransactor {
	return &ImageTransactor{transactorGen: transactorGen}
}

func (imageTransactor *ImageTransactor) AddDefaults() *ImageTransactor {
	imageTransactor.transactorGen.AppendImports([]string{ImageUploadHelperPath}).
		AddParentConstructorCallArgs(core.NewParameter(NewImageUploadHelper)).
		AddTransactorMember(core.NewSimpleStatement(fmt.Sprintf(`%s const BASE_PATH = "%s"`,
			VisibilityPrivate, strings.ToLower(imageTransactor.transactorGen.identifier)))).
		AddTransactorMember(core.NewSimpleStatement(ImageValidationRules))
	return imageTransactor
}

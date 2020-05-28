package generator

type TransactorModel struct {
	modelGen *ModelGenerator
}

func NewTransactorModel(modelGen *ModelGenerator) *TransactorModel {
	return &TransactorModel{modelGen: modelGen}
}

func (transactorModelGen *TransactorModel) AddFileUrlsToFillAbles() *TransactorModel {
	transactorModelGen.modelGen.AddFillable("file_urls")
	return transactorModelGen
}

func (transactorModelGen *TransactorModel) AddFileUrlsValidationRules() *TransactorModel {
	transactorModelGen.modelGen.AddCreateValidationRule("file_urls", "sometimes|required", "").
		AddCreateValidationRule("file_urls.urls", "array", "").
		AddUpdateValidationRule("file_urls", "sometimes|required", "").
		AddUpdateValidationRule("file_urls.urls", "array", "")
	return transactorModelGen
}

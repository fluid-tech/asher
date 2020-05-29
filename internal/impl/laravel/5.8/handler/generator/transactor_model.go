package generator

/*CONSTANTS*/
const (
	fileUrls = `file_urls`
	urlsOffFileUrls = `file_urls.urls`
	fileUrlsValidationRule = `sometimes|required`
	urlsOffFileUrlsValidationRule = `array`
)

/*Model related activities for transactor will be handled in this class*/
type TransactorModel struct {
	modelGen *ModelGenerator
}

/**
Returns a New Instance of TransactorModel
Parameters:
	- modelGen: model generator on which transactor specific activities will be carried out
Returns:
	- instance of the generator object
*/
func NewTransactorModel(modelGen *ModelGenerator) *TransactorModel {
	return &TransactorModel{modelGen: modelGen}
}

/**
this method will add file_urls to fill able for mass insertion
Returns:
	- instance of the generator object
*/
func (transactorModelGen *TransactorModel) AddFileUrlsToFillAbles() *TransactorModel {
	transactorModelGen.modelGen.AddFillable("file_urls")
	return transactorModelGen
}

/**
Adds validation for file urls in the model update and create validation rules
Returns:
	- instance of the generator object
*/
func (transactorModelGen *TransactorModel) AddFileUrlsValidationRules() *TransactorModel {
	transactorModelGen.modelGen.
		AddCreateValidationRule(fileUrls, fileUrlsValidationRule, "").
		AddCreateValidationRule(urlsOffFileUrls, urlsOffFileUrlsValidationRule, "").
		AddUpdateValidationRule(fileUrls, fileUrlsValidationRule, "").
		AddUpdateValidationRule(urlsOffFileUrls, urlsOffFileUrlsValidationRule, "")
	return transactorModelGen
}

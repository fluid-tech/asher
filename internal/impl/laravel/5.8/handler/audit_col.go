package handler

import "asher/internal/api"

type AuditCol struct{
	api.Handler
}

func NewAuditColHandler() AuditCol {
	return AuditCol{}
}

func (auditColHandler AuditCol) Handle()  (api.EmitterFile, error){

	return api.EmitterFile{}, nil
}

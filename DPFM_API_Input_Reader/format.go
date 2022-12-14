package dpfm_api_input_reader

import (
	"data-platform-api-accounting-document-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToAccountingDocumentItem() *requests.AccountingDocumentItem {
	data := sdc.AccountingDocument
	return &requests.AccountingDocumentItem{
		AccountingDocument:             data.AccountingDocument,
		AccountingDocumentItem:         data.AccountingDocumentItem,
		FiscalYear:                     data.FiscalYear,
		ChartOfAccounts:                data.ChartOfAccounts,
		AccountingDocumentItemType:     data.AccountingDocumentItemType,
		ClearingDate:                   data.ClearingDate,
		ClearingCreationDate:           data.ClearingCreationDate,
		ClearingAccountingDocument:     data.ClearingAccountingDocument,
		IsCleared:                      data.IsCleared,
		PostingKey:                     data.PostingKey,
		DebitCreditCode:                data.DebitCreditCode,
		BusinessArea:                   data.BusinessArea,
		TaxCode:                        data.TaxCode,
		TaxType:                        data.TaxType,
		DocumentItemText:               data.DocumentItemText,
		OrderID:                        data.OrderID,
		OrderItem:                      data.OrderItem,
		OrderType:                      data.OrderType,
		ContractType:                   data.ContractType,
		InvoiceDocument:                data.InvoiceDocument,
		InvoiceDocumentItem:            data.InvoiceDocumentItem,
		ScheduleLine:                   data.ScheduleLine,
		IsOpenItemManaged:              data.IsOpenItemManaged,
		IsAutomaticallyCreated:         data.IsAutomaticallyCreated,
		GLAccount:                      data.GLAccount,
		GLAccountName:                  data.GLAccountName,
		Customer:                       data.Customer,
		Supplier:                       data.Supplier,
		IsBalanceSheetAccount:          data.IsBalanceSheetAccount,
		ProfitLossAccountType:          data.ProfitLossAccountType,
		DueCalculationBaseDate:         data.DueCalculationBaseDate,
		PaymentTerms:                   data.PaymentTerms,
		NetPaymentDays:                 data.NetPaymentDays,
		PaymentMethod:                  data.PaymentMethod,
		PaymentBlockStatus:             data.PaymentBlockStatus,
		HouseBank:                      data.HouseBank,
		FinInstIdentification:          data.FinInstIdentification,
		InvoiceReference:               data.InvoiceReference,
		InvoiceReferenceFiscalYear:     data.InvoiceReferenceFiscalYear,
		InvoiceItemReference:           data.InvoiceItemReference,
		SupplyingCountry:               data.SupplyingCountry,
		Material:                       data.Material,
		Plant:                          data.Plant,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		InventoryValuationType:         data.InventoryValuationType,
		ClearingIsReversed:             data.ClearingIsReversed,
		CreditControlArea:              data.CreditControlArea,
		ClearingItem:                   data.ClearingItem,
		ClearingDocFiscalYear:          data.ClearingDocFiscalYear,
		HouseBankAccount:               data.HouseBankAccount,
		ReferenceDocumentType:          data.ReferenceDocumentType,
		OriginalReferenceDocument:      data.OriginalReferenceDocument,
		AccountingDocumentItemRef:      data.AccountingDocumentItemRef,
		FiscalPeriod:                   data.FiscalPeriod,
		PostingDate:                    data.PostingDate,
		DocumentDate:                   data.DocumentDate,
		AccountingDocumentType:         data.AccountingDocumentType,
		NetDueDate:                     data.NetDueDate,
		TransactionCurrency:            data.TransactionCurrency,
		AmountInTransactionCurrency:    data.AmountInTransactionCurrency,
		OriginalTaxBaseAmount:          data.OriginalTaxBaseAmount,
		TaxAmount:                      data.TaxAmount,
		TaxBaseAmountInTransCrcy:       data.TaxBaseAmountInTransCrcy,
		BaseUnit:                       data.BaseUnit,
		Quantity:                       data.Quantity,
		GoodsMovementEntryUnit:         data.GoodsMovementEntryUnit,
		QuantityInEntryUnit:            data.QuantityInEntryUnit,
		AccountingDocumentCreationDate: data.AccountingDocumentCreationDate,
		CreationTime:                   data.CreationTime,
		LastChangeDate:                 data.LastChangeDate,
		ExchangeRateDate:               data.ExchangeRateDate,
		IntercompanyTransaction:        data.IntercompanyTransaction,
		DocumentReferenceID:            data.DocumentReferenceID,
		ReverseDocument:                data.ReverseDocument,
		ReverseDocumentFiscalYear:      data.ReverseDocumentFiscalYear,
		AccountingDocumentHeaderText:   data.AccountingDocumentHeaderText,
		ExchangeRate:                   data.ExchangeRate,
		TaxIsCalculatedAutomatically:   data.TaxIsCalculatedAutomatically,
		Reference1InDocumentHeader:     data.Reference1InDocumentHeader,
		Reference2InDocumentHeader:     data.Reference2InDocumentHeader,
		IsReversed:                     data.IsReversed,
	}
}

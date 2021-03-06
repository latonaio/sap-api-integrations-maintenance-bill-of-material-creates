package sap_api_input_reader

import (
	"sap-api-integrations-maintenance-bill-of-material-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.MaintenanceBillOfMaterial
	return &requests.Header{
		BillOfMaterial:               data.BillOfMaterial,
		BillOfMaterialCategory:       data.BillOfMaterialCategory,
		BillOfMaterialVariant:        data.BillOfMaterialVariant,
		BillOfMaterialVersion:        data.BillOfMaterialVersion,
		TechnicalObject:              data.TechnicalObject,
		Plant:                        data.Plant,
		EngineeringChangeDocument:    data.EngineeringChangeDocument,
		BillOfMaterialVariantUsage:   data.BillOfMaterialVariantUsage,
		BillOfMaterialHeaderUUID:     data.BillOfMaterialHeaderUUID,
		IsMultipleBOMAlt:             data.IsMultipleBOMAlt,
		BOMHeaderInternalChangeCount: data.BOMHeaderInternalChangeCount,
		BOMUsagePriority:             data.BOMUsagePriority,
		BillOfMaterialAuthsnGrp:      data.BillOfMaterialAuthsnGrp,
		BOMVersionStatus:             data.BOMVersionStatus,
		IsVersionBillOfMaterial:      data.IsVersionBillOfMaterial,
		IsLatestBOMVersion:           data.IsLatestBOMVersion,
		IsConfiguredMaterial:         data.IsConfiguredMaterial,
		BOMTechnicalType:             data.BOMTechnicalType,
		BOMGroup:                     data.BOMGroup,
		BOMHeaderText:                data.BOMHeaderText,
		BOMAlternativeText:           data.BOMAlternativeText,
		BillOfMaterialStatus:         data.BillOfMaterialStatus,
		HeaderValidityStartDate:      data.HeaderValidityStartDate,
		HeaderValidityEndDate:        data.HeaderValidityEndDate,
		ChgToEngineeringChgDocument:  data.ChgToEngineeringChgDocument,
		IsMarkedForDeletion:          data.IsMarkedForDeletion,
		BOMHeaderBaseUnit:            data.BOMHeaderBaseUnit,
		BOMHeaderQuantityInBaseUnit:  data.BOMHeaderQuantityInBaseUnit,
		RecordCreationDate:           data.RecordCreationDate,
		LastChangeDate:               data.LastChangeDate,
		BOMIsToBeDeleted:             data.BOMIsToBeDeleted,
		// ToItem: &struct {
		// 	ToItemResults []*requests.Item `json:"results"`
		// }{
		// 	ToItemResults: []*requests.Item{
		// 		sdc.ConvertToItem(),
		// 	},
		// },
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	data := sdc.MaintenanceBillOfMaterial.MaintenanceBillOfMaterialItem
	return &requests.Item{
		TechnicalObject:               data.TechnicalObject,
		Plant:                         data.Plant,
		EngineeringChangeDocument:     data.EngineeringChangeDocument,
		BillOfMaterialVariantUsage:    data.BillOfMaterialVariantUsage,
		BillOfMaterialItemNodeNumber:  data.BillOfMaterialItemNodeNumber,
		BillOfMaterialItemUUID:        data.BillOfMaterialItemUUID,
		BOMItemInternalChangeCount:    data.BOMItemInternalChangeCount,
		ValidityStartDate:             data.ValidityStartDate,
		ValidityEndDate:               data.ValidityEndDate,
		ChgToEngineeringChgDocument:   data.ChgToEngineeringChgDocument,
		InheritedNodeNumberForBOMItem: data.InheritedNodeNumberForBOMItem,
		BOMItemRecordCreationDate:     data.BOMItemRecordCreationDate,
		BOMItemLastChangeDate:         data.BOMItemLastChangeDate,
		BillOfMaterialComponent:       data.BillOfMaterialComponent,
		BillOfMaterialItemCategory:    data.BillOfMaterialItemCategory,
		BillOfMaterialItemNumber:      data.BillOfMaterialItemNumber,
		BillOfMaterialItemUnit:        data.BillOfMaterialItemUnit,
		BillOfMaterialItemQuantity:    data.BillOfMaterialItemQuantity,
		IsAssembly:                    data.IsAssembly,
		IsSubItem:                     data.IsSubItem,
		BOMItemSorter:                 data.BOMItemSorter,
		PurchasingGroup:               data.PurchasingGroup,
		Currency:                      data.Currency,
		MaterialComponentPrice:        data.MaterialComponentPrice,
		IdentifierBOMItem:             data.IdentifierBOMItem,
		MaterialPriceUnitQty:          data.MaterialPriceUnitQty,
		ComponentScrapInPercent:       data.ComponentScrapInPercent,
		OperationScrapInPercent:       data.OperationScrapInPercent,
		IsNetScrap:                    data.IsNetScrap,
		QuantityVariableSizeItem:      data.QuantityVariableSizeItem,
		FormulaKey:                    data.FormulaKey,
		ComponentDescription:          data.ComponentDescription,
		BOMItemDescription:            data.BOMItemDescription,
		BOMItemText2:                  data.BOMItemText2,
		MaterialGroup:                 data.MaterialGroup,
		DocumentType:                  data.DocumentType,
		DocNumber:                     data.DocNumber,
		DocumentVersion:               data.DocumentVersion,
		DocumentPart:                  data.DocumentPart,
		ClassNumber:                   data.ClassNumber,
		ClassType:                     data.ClassType,
		ResultingItemCategory:         data.ResultingItemCategory,
		DependencyObjectNumber:        data.DependencyObjectNumber,
		ObjectType:                    data.ObjectType,
		IsClassificationRelevant:      data.IsClassificationRelevant,
		IsBulkMaterial:                data.IsBulkMaterial,
		BOMItemIsSparePart:            data.BOMItemIsSparePart,
		BOMItemIsSalesRelevant:        data.BOMItemIsSalesRelevant,
		IsProductionRelevant:          data.IsProductionRelevant,
		BOMItemIsPlantMaintRelevant:   data.BOMItemIsPlantMaintRelevant,
		BOMItemIsCostingRelevant:      data.BOMItemIsCostingRelevant,
		IsEngineeringRelevant:         data.IsEngineeringRelevant,
		SpecialProcurementType:        data.SpecialProcurementType,
		IsBOMRecursiveAllowed:         data.IsBOMRecursiveAllowed,
		OperationLeadTimeOffset:       data.OperationLeadTimeOffset,
		OpsLeadTimeOffsetUnit:         data.OpsLeadTimeOffsetUnit,
		IsMaterialProvision:           data.IsMaterialProvision,
		BOMIsRecursive:                data.BOMIsRecursive,
		DocumentIsCreatedByCAD:        data.DocumentIsCreatedByCAD,
		DistrKeyCompConsumption:       data.DistrKeyCompConsumption,
		DeliveryDurationInDays:        data.DeliveryDurationInDays,
		Creditor:                      data.Creditor,
		CostElement:                   data.CostElement,
		Size1:                         data.Size1,
		Size2:                         data.Size2,
		Size3:                         data.Size3,
		UnitOfMeasureForSize1To3:      data.UnitOfMeasureForSize1To3,
		GoodsReceiptDuration:          data.GoodsReceiptDuration,
		PurchasingOrganization:        data.PurchasingOrganization,
		RequiredComponent:             data.RequiredComponent,
		MultipleSelectionAllowed:      data.MultipleSelectionAllowed,
		ProdOrderIssueLocation:        data.ProdOrderIssueLocation,
		MaterialIsCoProduct:           data.MaterialIsCoProduct,
		ExplosionType:                 data.ExplosionType,
		AlternativeItemGroup:          data.AlternativeItemGroup,
		FollowUpGroup:                 data.FollowUpGroup,
		DiscontinuationGroup:          data.DiscontinuationGroup,
		IsConfigurableBOM:             data.IsConfigurableBOM,
		ReferencePoint:                data.ReferencePoint,
		LeadTimeOffset:                data.LeadTimeOffset,
		ProductionSupplyArea:          data.ProductionSupplyArea,
		IsDeleted:                     data.IsDeleted,
		BillOfMaterialHeaderUUID:      data.BillOfMaterialHeaderUUID,
	}
}

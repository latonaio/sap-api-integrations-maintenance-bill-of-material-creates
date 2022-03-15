package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-maintenance-bill-of-material-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	header := &Header{
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
		// ToItem:                         data.ToItem.ToItemResults[0]
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		BillOfMaterial:                data.BillOfMaterial,
		BillOfMaterialCategory:        data.BillOfMaterialCategory,
		BillOfMaterialVariant:         data.BillOfMaterialVariant,
		BillOfMaterialVersion:         data.BillOfMaterialVersion,
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

	return item, nil
}

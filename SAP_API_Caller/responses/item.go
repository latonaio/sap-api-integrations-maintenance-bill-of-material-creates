package responses

type Item struct {
	D struct {
            BillOfMaterial                string      `json:"BillOfMaterial"`
			BillOfMaterialCategory        string      `json:"BillOfMaterialCategory"`
			BillOfMaterialVariant         string      `json:"BillOfMaterialVariant"`
			BillOfMaterialVersion         string      `json:"BillOfMaterialVersion"`
			TechnicalObject               string      `json:"TechnicalObject"`
			Plant                         string      `json:"Plant"`
			EngineeringChangeDocument     string      `json:"EngineeringChangeDocument"`
			BillOfMaterialVariantUsage    string      `json:"BillOfMaterialVariantUsage"`
			BillOfMaterialItemNodeNumber  string      `json:"BillOfMaterialItemNodeNumber"`
			BillOfMaterialItemUUID        string      `json:"BillOfMaterialItemUUID"`
			BOMItemInternalChangeCount    string      `json:"BOMItemInternalChangeCount"`
			ValidityStartDate             string      `json:"ValidityStartDate"`
			ValidityEndDate               string      `json:"ValidityEndDate"`
			ChgToEngineeringChgDocument   string      `json:"ChgToEngineeringChgDocument"`
			InheritedNodeNumberForBOMItem string      `json:"InheritedNodeNumberForBOMItem"`
			BOMItemRecordCreationDate     string      `json:"BOMItemRecordCreationDate"`
			BOMItemLastChangeDate         string      `json:"BOMItemLastChangeDate"`
			BillOfMaterialComponent       string      `json:"BillOfMaterialComponent"`
			BillOfMaterialItemCategory    string      `json:"BillOfMaterialItemCategory"`
			BillOfMaterialItemNumber      string      `json:"BillOfMaterialItemNumber"`
			BillOfMaterialItemUnit        string      `json:"BillOfMaterialItemUnit"`
			BillOfMaterialItemQuantity    string      `json:"BillOfMaterialItemQuantity"`
			IsAssembly                    string      `json:"IsAssembly"`
			IsSubItem                     bool        `json:"IsSubItem"`
			BOMItemSorter                 string      `json:"BOMItemSorter"`
			PurchasingGroup               string      `json:"PurchasingGroup"`
			Currency                      string      `json:"Currency"`
			MaterialComponentPrice        string      `json:"MaterialComponentPrice"`
			IdentifierBOMItem             string      `json:"IdentifierBOMItem"`
			MaterialPriceUnitQty          string      `json:"MaterialPriceUnitQty"`
			ComponentScrapInPercent       string      `json:"ComponentScrapInPercent"`
			OperationScrapInPercent       string      `json:"OperationScrapInPercent"`
			IsNetScrap                    bool        `json:"IsNetScrap"`
			QuantityVariableSizeItem      string      `json:"QuantityVariableSizeItem"`
			FormulaKey                    string      `json:"FormulaKey"`
			ComponentDescription          string      `json:"ComponentDescription"`
			BOMItemDescription            string      `json:"BOMItemDescription"`
			BOMItemText2                  string      `json:"BOMItemText2"`
			MaterialGroup                 string      `json:"MaterialGroup"`
			DocumentType                  string      `json:"DocumentType"`
			DocNumber                     string      `json:"DocNumber"`
			DocumentVersion               string      `json:"DocumentVersion"`
			DocumentPart                  string      `json:"DocumentPart"`
			ClassNumber                   string      `json:"ClassNumber"`
			ClassType                     string      `json:"ClassType"`
			ResultingItemCategory         string      `json:"ResultingItemCategory"`
			DependencyObjectNumber        string      `json:"DependencyObjectNumber"`
			ObjectType                    string      `json:"ObjectType"`
			IsClassificationRelevant      bool        `json:"IsClassificationRelevant"`
			IsBulkMaterial                bool        `json:"IsBulkMaterial"`
			BOMItemIsSparePart            string      `json:"BOMItemIsSparePart"`
			BOMItemIsSalesRelevant        string      `json:"BOMItemIsSalesRelevant"`
			IsProductionRelevant          bool        `json:"IsProductionRelevant"`
			BOMItemIsPlantMaintRelevant   bool        `json:"BOMItemIsPlantMaintRelevant"`
			BOMItemIsCostingRelevant      string      `json:"BOMItemIsCostingRelevant"`
			IsEngineeringRelevant         bool        `json:"IsEngineeringRelevant"`
			SpecialProcurementType        string      `json:"SpecialProcurementType"`
			IsBOMRecursiveAllowed         bool        `json:"IsBOMRecursiveAllowed"`
			OperationLeadTimeOffset       string      `json:"OperationLeadTimeOffset"`
			OpsLeadTimeOffsetUnit         string      `json:"OpsLeadTimeOffsetUnit"`
			IsMaterialProvision           string      `json:"IsMaterialProvision"`
			BOMIsRecursive                bool        `json:"BOMIsRecursive"`
			DocumentIsCreatedByCAD        bool        `json:"DocumentIsCreatedByCAD"`
			DistrKeyCompConsumption       string      `json:"DistrKeyCompConsumption"`
			DeliveryDurationInDays        string      `json:"DeliveryDurationInDays"`
			Creditor                      string      `json:"Creditor"`
			CostElement                   string      `json:"CostElement"`
			Size1                         string      `json:"Size1"`
			Size2                         string      `json:"Size2"`
			Size3                         string      `json:"Size3"`
			UnitOfMeasureForSize1To3      string      `json:"UnitOfMeasureForSize1To3"`
			GoodsReceiptDuration          string      `json:"GoodsReceiptDuration"`
			PurchasingOrganization        string      `json:"PurchasingOrganization"`
			RequiredComponent             bool        `json:"RequiredComponent"`
			MultipleSelectionAllowed      bool        `json:"MultipleSelectionAllowed"`
			ProdOrderIssueLocation        string      `json:"ProdOrderIssueLocation"`
			MaterialIsCoProduct           bool        `json:"MaterialIsCoProduct"`
			ExplosionType                 string      `json:"ExplosionType"`
			AlternativeItemGroup          string      `json:"AlternativeItemGroup"`
			FollowUpGroup                 string      `json:"FollowUpGroup"`
			DiscontinuationGroup          string      `json:"DiscontinuationGroup"`
			IsConfigurableBOM             string      `json:"IsConfigurableBOM"`
			ReferencePoint                string      `json:"ReferencePoint"`
			LeadTimeOffset                string      `json:"LeadTimeOffset"`
			ProductionSupplyArea          string      `json:"ProductionSupplyArea"`
			IsDeleted                     bool        `json:"IsDeleted"`
			BillOfMaterialHeaderUUID      string      `json:"BillOfMaterialHeaderUUID"`
	} `json:"d"`
}

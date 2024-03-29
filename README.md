# sap-api-integrations-maintenance-bill-of-material-creates  
sap-api-integrations-maintenance-bill-of-material-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、保全部品表（BOM）データを登録するマイクロサービスです。  
sap-api-integrations-maintenance-bill-of-material-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-maintenance-bill-of-material-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_MAINTENANCE_BOM_0001/overview  

## 動作環境  
sap-api-integrations-maintenance-bill-of-material-creates  は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-maintenance-bill-of-material-creates  は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-maintenance-bill-of-material-creates  が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_MAINTENANCE_BOM_0001/overview  
* APIサービス名(=baseURL): API_MAINTENANCEBOM

## 本レポジトリ に 含まれる API名
sap-api-integrations-maintenance-bill-of-material-creates  には、次の API をコールするためのリソースが含まれています。  

* BOMHeader（保全部品表 - ヘッダ）
* BOMItem（保全部品表 - 明細）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
"api_schema": "SAPMaintenanceBillOfMaterialCreates",
"accepter": ["Header"],
"technical_object": "210100091",
"plant": "DE10",
"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
"api_schema": "SAPMaintenanceBillOfMaterialCreates",
"accepter": ["All"],
"technical_object": "210100091",
"plant": "DE10",
"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostMaintenanceBillOfMaterial(
	header            *requests.Header,
	item              *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
    wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、golang-logging-library-for-sap により、以下のようなデータがJSON形式で出力されます。
以下の sample.json の例は、SAP 保全部品表 の ヘッダ が登録された結果の JSON の例です。
以下の項目のうち、"BillOfMaterial" ～ "to_MaintBillOfMaterialItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。
```
{
    "cursor": "/Users/latona2/bitbucket/sap-api-integrations-maintenance-bill-of-material-creates-reads/SAP_API_Caller/caller.go#L73",
    "function": "sap-api-integrations-maintenance-bill-of-material-creates-reads/SAP_API_Caller.(*SAPAPICaller).Header",
    "level": "INFO",
    "message": [
        {
            "BillOfMaterial": "00000001",
            "BillOfMaterialCategory": "E",
            "BillOfMaterialVariant": "1",
            "BillOfMaterialVersion": "",
            "TechnicalObject": "210100091",
            "Plant": "DE10",
            "EngineeringChangeDocument": "",
            "BillOfMaterialVariantUsage": "4",
            "BillOfMaterialHeaderUUID": "00163e19-8846-1ed6-8f85-6c999d39c7b1",
            "IsMultipleBOMAlt": false,
            "BOMHeaderInternalChangeCount": "1",
            "BOMUsagePriority": "",
            "BillOfMaterialAuthsnGrp": "",
            "BOMVersionStatus": "",
            "IsVersionBillOfMaterial": false,
            "IsLatestBOMVersion": false,
            "IsConfiguredMaterial": false,
            "BOMTechnicalType": "",
            "BOMGroup": "",
            "BOMHeaderText": "Manual Isolation Valve",
            "BOMAlternativeText": "",
            "BillOfMaterialStatus": "1",
            "HeaderValidityStartDate": "2016-06-27T09:00:00+09:00",
            "HeaderValidityEndDate": "9999-12-31T09:00:00+09:00",
            "ChgToEngineeringChgDocument": "",
            "IsMarkedForDeletion": false,
            "BOMHeaderBaseUnit": "PC",
            "BOMHeaderQuantityInBaseUnit": "1",
            "RecordCreationDate": "2016-06-27T09:00:00+09:00",
            "LastChangeDate": "",
            "BOMIsToBeDeleted": "",
            "to_MaintBillOfMaterialItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_MAINTENANCEBOM/BOMHeader(BillOfMaterial='00000001',BillOfMaterialCategory='E',BillOfMaterialVariant='1',BillOfMaterialVersion='',TechnicalObject='210100091',Plant='DE10',EngineeringChangeDocument='',BillOfMaterialVariantUsage='4')/to_MaintBillOfMaterialItem"
        }
    ],
    "time": "2022-01-28T17:05:29+09:00"
}
```

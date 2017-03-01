package main

import (
	"encoding/xml"
	"reflect"
	"strings"
	"time"

	"gitlab.com/garanteka/parsexsd/xsd"
)

// Export is generated from an XSD element
type Export struct {
	BankGuarantee                                ZfcsBankGuaranteeType                             `xml:"bankGuarantee,omitempty" bson:"bankGuarantee,omitempty"`
	BankGuaranteeInvalid                         ZfcsBankGuaranteeInvalidType                      `xml:"bankGuaranteeInvalid,omitempty" bson:"bankGuaranteeInvalid,omitempty"`
	BankGuaranteeRefusal                         ZfcsBankGuaranteeRefusalType                      `xml:"bankGuaranteeRefusal,omitempty" bson:"bankGuaranteeRefusal,omitempty"`
	BankGuaranteeRefusalInvalid                  ZfcsBankGuaranteeRefusalInvalidType               `xml:"bankGuaranteeRefusalInvalid,omitempty" bson:"bankGuaranteeRefusalInvalid,omitempty"`
	BankGuaranteeTermination                     ZfcsBankGuaranteeTerminationType                  `xml:"bankGuaranteeTermination,omitempty" bson:"bankGuaranteeTermination,omitempty"`
	BankGuaranteeTerminationInvalid              ZfcsBankGuaranteeTerminationInvalidType           `xml:"bankGuaranteeTerminationInvalid,omitempty" bson:"bankGuaranteeTerminationInvalid,omitempty"`
	BankGuaranteeReturn                          ZfcsBankGuaranteeReturnType                       `xml:"bankGuaranteeReturn,omitempty" bson:"bankGuaranteeReturn,omitempty"`
	BankGuaranteeReturnInvalid                   ZfcsBankGuaranteeReturnInvalidType                `xml:"bankGuaranteeReturnInvalid,omitempty" bson:"bankGuaranteeReturnInvalid,omitempty"`
	Contract                                     []ZfcsContract2015Type                            `xml:"contract,omitempty" bson:"contract,omitempty"`
	ContractCancel                               []ZfcsContractCancel2015Type                      `xml:"contractCancel,omitempty" bson:"contractCancel,omitempty"`
	ContractProcedure                            []ZfcsContractProcedure2015Type                   `xml:"contractProcedure,omitempty" bson:"contractProcedure,omitempty"`
	ContractProcedureCancel                      []ZfcsContractProcedureCancel2015Type             `xml:"contractProcedureCancel,omitempty" bson:"contractProcedureCancel,omitempty"`
	FcsClarification                             ZfcsClarificationType                             `xml:"fcsClarification,omitempty" bson:"fcsClarification,omitempty"`
	FcsContractSign                              ZfcsContractSignType                              `xml:"fcsContractSign,omitempty" bson:"fcsContractSign,omitempty"`
	FcsNotificationEF                            ZfcsNotificationEFType                            `xml:"fcsNotificationEF,omitempty" bson:"fcsNotificationEF,omitempty"`
	FcsNotificationEFDateChange                  ZfcsNotificationEFDateChangeType                  `xml:"fcs_notificationEFDateChange,omitempty" bson:"fcs_notificationEFDateChange,omitempty"`
	FcsNotificationEP                            ZfcsNotificationEPType                            `xml:"fcsNotificationEP,omitempty" bson:"fcsNotificationEP,omitempty"`
	FcsNotificationISM                           ZfcsNotificationISMType                           `xml:"fcsNotificationISM,omitempty" bson:"fcsNotificationISM,omitempty"`
	FcsNotificationISO                           ZfcsNotificationISOType                           `xml:"fcsNotificationISO,omitempty" bson:"fcsNotificationISO,omitempty"`
	FcsNotificationOK                            ZfcsNotificationOKType                            `xml:"fcsNotificationOK,omitempty" bson:"fcsNotificationOK,omitempty"`
	FcsNotificationOKD                           ZfcsNotificationOKDType                           `xml:"fcsNotificationOKD,omitempty" bson:"fcsNotificationOKD,omitempty"`
	FcsNotificationOKOU                          ZfcsNotificationOKOUType                          `xml:"fcsNotificationOKOU,omitempty" bson:"fcsNotificationOKOU,omitempty"`
	FcsNotificationPO                            ZfcsNotificationPOType                            `xml:"fcsNotificationPO,omitempty" bson:"fcsNotificationPO,omitempty"`
	FcsNotificationZakA                          ZfcsNotificationZakAType                          `xml:"fcsNotificationZakA,omitempty" bson:"fcsNotificationZakA,omitempty"`
	FcsNotificationZakKD                         ZfcsNotificationZakKDType                         `xml:"fcsNotificationZakKD,omitempty" bson:"fcsNotificationZakKD,omitempty"`
	FcsNotificationZakKOU                        ZfcsNotificationZakKOUType                        `xml:"fcsNotificationZakKOU,omitempty" bson:"fcsNotificationZakKOU,omitempty"`
	FcsNotificationZakK                          ZfcsNotificationZakKType                          `xml:"fcsNotificationZakK,omitempty" bson:"fcsNotificationZakK,omitempty"`
	FcsNotificationZK                            ZfcsNotificationZKType                            `xml:"fcsNotificationZK,omitempty" bson:"fcsNotificationZK,omitempty"`
	FcsNotificationZP                            ZfcsNotificationZPType                            `xml:"fcsNotificationZP,omitempty" bson:"fcsNotificationZP,omitempty"`
	FcsNotificationLotChange                     ZfcsNotificationLotChangeType                     `xml:"fcsNotificationLotChange,omitempty" bson:"fcsNotificationLotChange,omitempty"`
	FcsNotificationCancel                        ZfcsNotificationCancelType                        `xml:"fcsNotificationCancel,omitempty" bson:"fcsNotificationCancel,omitempty"`
	FcsNotificationCancelFailure                 ZfcsNotificationCancelFailureType                 `xml:"fcsNotificationCancelFailure,omitempty" bson:"fcsNotificationCancelFailure,omitempty"`
	FcsNotificationLotCancel                     ZfcsNotificationLotCancelType                     `xml:"fcsNotificationLotCancel,omitempty" bson:"fcsNotificationLotCancel,omitempty"`
	FcsNotificationOrgChange                     ZfcsNotificationOrgChangeType                     `xml:"fcsNotificationOrgChange,omitempty" bson:"fcsNotificationOrgChange,omitempty"`
	FcsPurchaseDocument                          ZfcsPurchaseDocumentType                          `xml:"fcsPurchaseDocument,omitempty" bson:"fcsPurchaseDocument,omitempty"`
	FcsPlacementResult                           ZfcsPlacementResultType                           `xml:"fcsPlacementResult,omitempty" bson:"fcsPlacementResult,omitempty"`
	FcsPurchaseDocumentCancel                    ZfcsPurchaseDocumentCancelType                    `xml:"fcsPurchaseDocumentCancel,omitempty" bson:"fcsPurchaseDocumentCancel,omitempty"`
	FcsPurchaseProlongationOK                    ZfcsPurchaseProlongationOKType                    `xml:"fcsPurchaseProlongationOK,omitempty" bson:"fcsPurchaseProlongationOK,omitempty"`
	FcsPurchaseProlongationZK                    ZfcsPurchaseProlongationZKType                    `xml:"fcsPurchaseProlongationZK,omitempty" bson:"fcsPurchaseProlongationZK,omitempty"`
	FcsProtocolOK1                               ZfcsProtocolOK1Type                               `xml:"fcsProtocolOK1,omitempty" bson:"fcsProtocolOK1,omitempty"`
	FcsProtocolOK2                               ZfcsProtocolOK2Type                               `xml:"fcsProtocolOK2,omitempty" bson:"fcsProtocolOK2,omitempty"`
	FcsProtocolOKSingleApp                       ZfcsProtocolOKSingleAppType                       `xml:"fcsProtocolOKSingleApp,omitempty" bson:"fcsProtocolOKSingleApp,omitempty"`
	FcsProtocolOKD1                              ZfcsProtocolOKD1Type                              `xml:"fcsProtocolOKD1,omitempty" bson:"fcsProtocolOKD1,omitempty"`
	FcsProtocolOKD2                              ZfcsProtocolOKD2Type                              `xml:"fcsProtocolOKD2,omitempty" bson:"fcsProtocolOKD2,omitempty"`
	FcsProtocolOKD3                              ZfcsProtocolOKD3Type                              `xml:"fcsProtocolOKD3,omitempty" bson:"fcsProtocolOKD3,omitempty"`
	FcsProtocolOKD4                              ZfcsProtocolOKD4Type                              `xml:"fcsProtocolOKD4,omitempty" bson:"fcsProtocolOKD4,omitempty"`
	FcsProtocolOKD5                              ZfcsProtocolOKD5Type                              `xml:"fcsProtocolOKD5,omitempty" bson:"fcsProtocolOKD5,omitempty"`
	FcsProtocolOKDSingleApp                      ZfcsProtocolOKDSingleAppType                      `xml:"fcsProtocolOKDSingleApp,omitempty" bson:"fcsProtocolOKDSingleApp,omitempty"`
	FcsProtocolOKOU1                             ZfcsProtocolOKOU1Type                             `xml:"fcsProtocolOKOU1,omitempty" bson:"fcsProtocolOKOU1,omitempty"`
	FcsProtocolOKOU2                             ZfcsProtocolOKOU2Type                             `xml:"fcsProtocolOKOU2,omitempty" bson:"fcsProtocolOKOU2,omitempty"`
	FcsProtocolOKOU3                             ZfcsProtocolOKOU3Type                             `xml:"fcsProtocolOKOU3,omitempty" bson:"fcsProtocolOKOU3,omitempty"`
	FcsProtocolOKOUSingleApp                     ZfcsProtocolOKOUSingleAppType                     `xml:"fcsProtocolOKOUSingleApp,omitempty" bson:"fcsProtocolOKOUSingleApp,omitempty"`
	FcsProtocolPO                                ZfcsProtocolPOType                                `xml:"fcsProtocolPO,omitempty" bson:"fcsProtocolPO,omitempty"`
	FcsProtocolZKAfterProlong                    ZfcsProtocolZKAfterProlongType                    `xml:"fcsProtocolZKAfterProlong,omitempty" bson:"fcsProtocolZKAfterProlong,omitempty"`
	FcsProtocolZK                                ZfcsProtocolZKType                                `xml:"fcsProtocolZK,omitempty" bson:"fcsProtocolZK,omitempty"`
	FcsProtocolZKBI                              ZfcsProtocolZKBIType                              `xml:"fcsProtocolZKBI,omitempty" bson:"fcsProtocolZKBI,omitempty"`
	FcsProtocolZPExtract                         ZfcsProtocolZPExtractType                         `xml:"fcsProtocolZPExtract,omitempty" bson:"fcsProtocolZPExtract,omitempty"`
	FcsProtocolZPFinal                           ZfcsProtocolZPFinalType                           `xml:"fcsProtocolZPFinal,omitempty" bson:"fcsProtocolZPFinal,omitempty"`
	FcsProtocolZP                                ZfcsProtocolZPType                                `xml:"fcsProtocolZP,omitempty" bson:"fcsProtocolZP,omitempty"`
	FcsProtocolEvasion                           ZfcsProtocolEvasionType                           `xml:"fcsProtocolEvasion,omitempty" bson:"fcsProtocolEvasion,omitempty"`
	FcsProtocolDeviation                         ZfcsProtocolDeviationType                         `xml:"fcsProtocolDeviation,omitempty" bson:"fcsProtocolDeviation,omitempty"`
	FcsProtocolEF1                               ZfcsProtocolEF1Type                               `xml:"fcsProtocolEF1,omitempty" bson:"fcsProtocolEF1,omitempty"`
	FcsProtocolEF2                               ZfcsProtocolEF2Type                               `xml:"fcsProtocolEF2,omitempty" bson:"fcsProtocolEF2,omitempty"`
	FcsProtocolEF3                               ZfcsProtocolEF3Type                               `xml:"fcsProtocolEF3,omitempty" bson:"fcsProtocolEF3,omitempty"`
	FcsProtocolEFInvalidation                    ZfcsProtocolEFInvalidationType                    `xml:"fcsProtocolEFInvalidation,omitempty" bson:"fcsProtocolEFInvalidation,omitempty"`
	FcsProtocolEFSingleApp                       ZfcsProtocolEFSingleAppType                       `xml:"fcsProtocolEFSingleApp,omitempty" bson:"fcsProtocolEFSingleApp,omitempty"`
	FcsProtocolEFSinglePart                      ZfcsProtocolEFSinglePartType                      `xml:"fcsProtocolEFSinglePart,omitempty" bson:"fcsProtocolEFSinglePart,omitempty"`
	FcsProtocolCancel                            ZfcsProtocolCancelType                            `xml:"fcsProtocolCancel,omitempty" bson:"fcsProtocolCancel,omitempty"`
	FcsAddInfo                                   ZfcsAddInfoType                                   `xml:"fcsAddInfo,omitempty" bson:"fcsAddInfo,omitempty"`
	FcsAddInfoInvalid                            ZfcsAddInfoInvalidType                            `xml:"fcsAddInfoInvalid,omitempty" bson:"fcsAddInfoInvalid,omitempty"`
	FcsСustomerReportContractExecution           ZfcsCustomerReportContractExecutionType           `xml:"fcsСustomerReportContractExecution,omitempty" bson:"fcsСustomerReportContractExecution,omitempty"`
	FcsСustomerReportContractExecutionInvalid    ZfcsCustomerReportContractExecutionInvalidType    `xml:"fcsСustomerReportContractExecutionInvalid,omitempty" bson:"fcsСustomerReportContractExecutionInvalid,omitempty"`
	FcsСustomerReportSmallScaleBusiness          ZfcsCustomerReportSmallScaleBusinessType          `xml:"fcsСustomerReportSmallScaleBusiness,omitempty" bson:"fcsСustomerReportSmallScaleBusiness,omitempty"`
	FcsСustomerReportSmallScaleBusinessInvalid   ZfcsCustomerReportSmallScaleBusinessInvalidType   `xml:"fcsСustomerReportSmallScaleBusinessInvalid,omitempty" bson:"fcsСustomerReportSmallScaleBusinessInvalid,omitempty"`
	FcsСustomerReportSingleContractor            ZfcsCustomerReportSingleContractorType            `xml:"fcsСustomerReportSingleContractor,omitempty" bson:"fcsСustomerReportSingleContractor,omitempty"`
	FcsСustomerReportSingleContractorInvalid     ZfcsCustomerReportSingleContractorInvalidType     `xml:"fcsСustomerReportSingleContractorInvalid,omitempty" bson:"fcsСustomerReportSingleContractorInvalid,omitempty"`
	FcsСustomerReportBigProjectMonitoring        ZfcsCustomerReportBigProjectMonitoringType        `xml:"fcsСustomerReportBigProjectMonitoring,omitempty" bson:"fcsСustomerReportBigProjectMonitoring,omitempty"`
	FcsСustomerReportBigProjectMonitoringInvalid ZfcsCustomerReportBigProjectMonitoringInvalidType `xml:"fcsСustomerReportBigProjectMonitoringInvalid,omitempty" bson:"fcsСustomerReportBigProjectMonitoringInvalid,omitempty"`
	FcsStandardContract                          ZfcsStandardContractType                          `xml:"fcsStandardContract,omitempty" bson:"fcsStandardContract,omitempty"`
	FcsStandardContractInvalid                   ZfcsStandardContractInvalidType                   `xml:"fcsStandardContractInvalid,omitempty" bson:"fcsStandardContractInvalid,omitempty"`
	FcsPublicDiscussionLargePurchase             ZfcsPublicDiscussionLargePurchaseType             `xml:"fcsPublicDiscussionLargePurchase,omitempty" bson:"fcsPublicDiscussionLargePurchase,omitempty"`
	FcsPublicDiscussionComment                   ZfcsPublicDiscussionCommentType                   `xml:"fcsPublicDiscussionComment,omitempty" bson:"fcsPublicDiscussionComment,omitempty"`
	FcsPublicDiscussionForm                      ZfcsPublicDiscussionFormType                      `xml:"fcsPublicDiscussionForm,omitempty" bson:"fcsPublicDiscussionForm,omitempty"`
	FcsPublicDiscussionAnswer                    ZfcsPublicDiscussionAnswerType                    `xml:"fcsPublicDiscussionAnswer,omitempty" bson:"fcsPublicDiscussionAnswer,omitempty"`
	FcsPublicDiscussionProtocol                  ZfcsPublicDiscussionProtocolType                  `xml:"fcsPublicDiscussionProtocol,omitempty" bson:"fcsPublicDiscussionProtocol,omitempty"`
	FcsRegulationRules                           ZfcsRegulationRulesType                           `xml:"fcsRegulationRules,omitempty" bson:"fcsRegulationRules,omitempty"`
	FcsRegulationRulesInvalid                    ZfcsRegulationRulesInvalidType                    `xml:"fcsRegulationRulesInvalid,omitempty" bson:"fcsRegulationRulesInvalid,omitempty"`
	FcsRequestForQuotation                       ZfcsRequestForQuotationType                       `xml:"fcsRequestForQuotation,omitempty" bson:"fcsRequestForQuotation,omitempty"`
	FcsRequestForQuotationCancel                 ZfcsRequestForQuotationCancelType                 `xml:"fcsRequestForQuotationCancel,omitempty" bson:"fcsRequestForQuotationCancel,omitempty"`
	FcsAuditResult                               ZfcsAuditResultType                               `xml:"fcsAuditResult,omitempty" bson:"fcsAuditResult,omitempty"`
	PublicDiscussionResult                       ZfcsPublicDiscussionResultType                    `xml:"publicDiscussionResult,omitempty" bson:"publicDiscussionResult,omitempty"`
	SketchPlan                                   ZfcsSketchPlanType                                `xml:"sketchPlan,omitempty" bson:"sketchPlan,omitempty"`
	SketchPlanExecution                          ZfcsSketchPlanExecutionType                       `xml:"sketchPlanExecution,omitempty" bson:"sketchPlanExecution,omitempty"`
	PurchasePlan                                 ZfcsPurchasePlanType                              `xml:"purchasePlan,omitempty" bson:"purchasePlan,omitempty"`
	TenderPlan                                   ZfcsTenderPlanType                                `xml:"tenderPlan,omitempty" bson:"tenderPlan,omitempty"`
	TenderPlanUnstructured                       ZfcsTenderPlanUnstructuredType                    `xml:"tenderPlanUnstructured,omitempty" bson:"tenderPlanUnstructured,omitempty"`
	TenderPlanCancel                             ZfcsTenderPlanCancelType                          `xml:"tenderPlanCancel,omitempty" bson:"tenderPlanCancel,omitempty"`
	Complaint                                    ZfcsComplaintType                                 `xml:"complaint,omitempty" bson:"complaint,omitempty"`
	ComplaintGroup                               ZfcsComplaintGroupType                            `xml:"complaintGroup,omitempty" bson:"complaintGroup,omitempty"`
	ComplaintCancel                              ZfcsComplaintCancelType                           `xml:"complaintCancel,omitempty" bson:"complaintCancel,omitempty"`
	TenderSuspension                             ZfcsTenderSuspensionType                          `xml:"tenderSuspension,omitempty" bson:"tenderSuspension,omitempty"`
	CheckPlan                                    ZfcsCheckPlanType                                 `xml:"checkPlan,omitempty" bson:"checkPlan,omitempty"`
	EventPlan                                    ZfcsEventPlanType                                 `xml:"eventPlan,omitempty" bson:"eventPlan,omitempty"`
	UnfairSupplier                               ZfcsUnfairSupplierType                            `xml:"unfairSupplier,omitempty" bson:"unfairSupplier,omitempty"`
	UnplannedCheck                               ZfcsUnplannedCheckType                            `xml:"unplannedCheck,omitempty" bson:"unplannedCheck,omitempty"`
	UnplannedEvent                               ZfcsUnplannedEventType                            `xml:"unplannedEvent,omitempty" bson:"unplannedEvent,omitempty"`
	UnplannedCheckCancel                         ZfcsUnplannedCheckCancelType                      `xml:"unplannedCheckCancel,omitempty" bson:"unplannedCheckCancel,omitempty"`
	UnplannedEventCancel                         ZfcsUnplannedEventCancelType                      `xml:"unplannedEventCancel,omitempty" bson:"unplannedEventCancel,omitempty"`
	CheckResult                                  ZfcsCheckResultType                               `xml:"checkResult,omitempty" bson:"checkResult,omitempty"`
	EventResult                                  ZfcsEventResultType                               `xml:"eventResult,omitempty" bson:"eventResult,omitempty"`
	CheckResultCancel                            ZfcsCheckResultCancelType                         `xml:"checkResultCancel,omitempty" bson:"checkResultCancel,omitempty"`
	EventResultCancel                            ZfcsEventResultCancelType                         `xml:"eventResultCancel,omitempty" bson:"eventResultCancel,omitempty"`
	NsiBankGuaranteeRefusalReasonList            NsiBankGuaranteeRefusalReasonList                 `xml:"nsiBankGuaranteeRefusalReasonList,omitempty" bson:"nsiBankGuaranteeRefusalReasonList,omitempty"`
	NsiBudgetList                                NsiBudgetList                                     `xml:"nsiBudgetList,omitempty" bson:"nsiBudgetList,omitempty"`
	NsiCalendarDayList                           NsiCalendarDayList                                `xml:"nsiCalendarDayList,omitempty" bson:"nsiCalendarDayList,omitempty"`
	NsiCommissionList                            NsiCommissionList                                 `xml:"nsiCommissionList,omitempty" bson:"nsiCommissionList,omitempty"`
	NsiCommissionRoleList                        NsiCommissionRoleList                             `xml:"nsiCommissionRoleList,omitempty" bson:"nsiCommissionRoleList,omitempty"`
	NsiContractPriceChangeReasonList             NsiContractPriceChangeReasonList                  `xml:"nsiContractPriceChangeReasonList,omitempty" bson:"nsiContractPriceChangeReasonList,omitempty"`
	NsiContractRefusalReasonList                 NsiContractRefusalReasonList                      `xml:"nsiContractRefusalReasonList,omitempty" bson:"nsiContractRefusalReasonList,omitempty"`
	NsiContractSingleCustomerReasonList          NsiContractSingleCustomerReasonList               `xml:"nsiContractSingleCustomerReasonList,omitempty" bson:"nsiContractSingleCustomerReasonList,omitempty"`
	NsiContractTerminationReason                 NsiContractTerminationReason                      `xml:"nsiContractTerminationReason,omitempty" bson:"nsiContractTerminationReason,omitempty"`
	NsiContractModificationReason                NsiContractModificationReason                     `xml:"nsiContractModificationReason,omitempty" bson:"nsiContractModificationReason,omitempty"`
	NsiContractExecutionDoc                      NsiContractExecutionDoc                           `xml:"nsiContractExecutionDoc,omitempty" bson:"nsiContractExecutionDoc,omitempty"`
	NsiContractReparationDoc                     NsiContractReparationDoc                          `xml:"nsiContractReparationDoc,omitempty" bson:"nsiContractReparationDoc,omitempty"`
	NsiContractPenaltyReason                     NsiContractPenaltyReason                          `xml:"nsiContractPenaltyReason,omitempty" bson:"nsiContractPenaltyReason,omitempty"`
	NsiContractOKOPFExtraBudget                  NsiContractOKOPFExtraBudget                       `xml:"nsiContractOKOPFExtraBudget,omitempty" bson:"nsiContractOKOPFExtraBudget,omitempty"`
	NsiContractCurrencyCBRF                      NsiContractCurrencyCBRF                           `xml:"nsiContractCurrencyCBRF,omitempty" bson:"nsiContractCurrencyCBRF,omitempty"`
	NsiCurrencyList                              NsiCurrencyList                                   `xml:"nsiCurrencyList,omitempty" bson:"nsiCurrencyList,omitempty"`
	NsiEvalCriterionList                         NsiEvalCriterionList                              `xml:"nsiEvalCriterionList,omitempty" bson:"nsiEvalCriterionList,omitempty"`
	NsiKBKBudgetList                             NsiKBKBudgetList                                  `xml:"nsiKBKBudgetList,omitempty" bson:"nsiKBKBudgetList,omitempty"`
	NsiKOSGUList                                 NsiKOSGUList                                      `xml:"nsiKOSGUList,omitempty" bson:"nsiKOSGUList,omitempty"`
	NsiOffBudgetList                             NsiOffBudgetList                                  `xml:"nsiOffBudgetList,omitempty" bson:"nsiOffBudgetList,omitempty"`
	NsiOKEIList                                  NsiOKEIList                                       `xml:"nsiOKEIList,omitempty" bson:"nsiOKEIList,omitempty"`
	NsiOKOPFList                                 NsiOKOPFList                                      `xml:"nsiOKOPFList,omitempty" bson:"nsiOKOPFList,omitempty"`
	NsiOKPDList                                  NsiOKPDList                                       `xml:"nsiOKPDList,omitempty" bson:"nsiOKPDList,omitempty"`
	NsiOKSMList                                  NsiOKSMList                                       `xml:"nsiOKSMList,omitempty" bson:"nsiOKSMList,omitempty"`
	NsiOKVEDList                                 NsiOKVEDList                                      `xml:"nsiOKVEDList,omitempty" bson:"nsiOKVEDList,omitempty"`
	NsiOKTMOList                                 NsiOKTMOList                                      `xml:"nsiOKTMOList,omitempty" bson:"nsiOKTMOList,omitempty"`
	NsiOrganizationList                          NsiOrganizationList                               `xml:"nsiOrganizationList,omitempty" bson:"nsiOrganizationList,omitempty"`
	NsiOrganizationRightsList                    NsiOrganizationRightsList                         `xml:"nsiOrganizationRightsList,omitempty" bson:"nsiOrganizationRightsList,omitempty"`
	NsiOrganizationTypesList                     NsiOrganizationTypesList                          `xml:"nsiOrganizationTypesList,omitempty" bson:"nsiOrganizationTypesList,omitempty"`
	NsiPlacingWayList                            NsiPlacingWayList                                 `xml:"nsiPlacingWayList,omitempty" bson:"nsiPlacingWayList,omitempty"`
	NsiPlanPositionChangeReasonList              NsiPlanPositionChangeReasonList                   `xml:"nsiPlanPositionChangeReasonList,omitempty" bson:"nsiPlanPositionChangeReasonList,omitempty"`
	NsiPurchaseDocumentTypesList                 NsiPurchaseDocumentTypesList                      `xml:"nsiPurchaseDocumentTypesList,omitempty" bson:"nsiPurchaseDocumentTypesList,omitempty"`
	NsiPurchasePreferenceList                    NsiPurchasePreferenceList                         `xml:"nsiPurchasePreferenceList,omitempty" bson:"nsiPurchasePreferenceList,omitempty"`
	NsiPurchaseRejectReasonList                  NsiPurchaseRejectReasonList                       `xml:"nsiPurchaseRejectReasonList,omitempty" bson:"nsiPurchaseRejectReasonList,omitempty"`
	NsiUserList                                  NsiUserList                                       `xml:"nsiUserList,omitempty" bson:"nsiUserList,omitempty"`
	NsiAbandonedReasonList                       NsiAbandonedReasonList                            `xml:"nsiAbandonedReasonList,omitempty" bson:"nsiAbandonedReasonList,omitempty"`
	NsiOKPD2List                                 NsiOKPD2List                                      `xml:"nsiOKPD2List,omitempty" bson:"nsiOKPD2List,omitempty"`
	NsiOKVED2List                                NsiOKVED2List                                     `xml:"nsiOKVED2List,omitempty" bson:"nsiOKVED2List,omitempty"`
	NsiDeviationFactFoundationList               NsiDeviationFactFoundationList                    `xml:"nsiDeviationFactFoundationList,omitempty" bson:"nsiDeviationFactFoundationList,omitempty"`
	NsiPublicDiscussionQuestionnaries            NsiPublicDiscussionQuestionnaries                 `xml:"nsiPublicDiscussionQuestionnaries,omitempty" bson:"nsiPublicDiscussionQuestionnaries,omitempty"`
	NsiPublicDiscussionDecisions                 NsiPublicDiscussionDecisions                      `xml:"nsiPublicDiscussionDecisions,omitempty" bson:"nsiPublicDiscussionDecisions,omitempty"`
	NsiAuditActionSubjects                       NsiAuditActionSubjects                            `xml:"nsiAuditActionSubjects,omitempty" bson:"nsiAuditActionSubjects,omitempty"`
	NsiBusinessControls                          NsiBusinessControls                               `xml:"nsiBusinessControls,omitempty" bson:"nsiBusinessControls,omitempty"`
}

// NsiBankGuaranteeRefusalReasonList is generated from an XSD element
type NsiBankGuaranteeRefusalReasonList struct {
	NsiBankGuaranteeRefusalReason []ZfcsNsiBankGuaranteeRefusalReasonType `xml:"nsiBankGuaranteeRefusalReason,omitempty" bson:"nsiBankGuaranteeRefusalReason,omitempty"`
}

// NsiBudgetList is generated from an XSD element
type NsiBudgetList struct {
	NsiBudget []ZfcsNsiBudgetType `xml:"nsiBudget,omitempty" bson:"nsiBudget,omitempty"`
}

// NsiCalendarDayList is generated from an XSD element
type NsiCalendarDayList struct {
	NsiCalendarDay []ZfcsNsiCalendarDaysType `xml:"nsiCalendarDay,omitempty" bson:"nsiCalendarDay,omitempty"`
}

// NsiCommissionList is generated from an XSD element
type NsiCommissionList struct {
	NsiCommission []ZfcsNsiCommissionType `xml:"nsiCommission,omitempty" bson:"nsiCommission,omitempty"`
}

// NsiCommissionRoleList is generated from an XSD element
type NsiCommissionRoleList struct {
	NsiCommissionRole []ZfcsNsiCommissionRoleType `xml:"nsiCommissionRole,omitempty" bson:"nsiCommissionRole,omitempty"`
}

// NsiContractPriceChangeReasonList is generated from an XSD element
type NsiContractPriceChangeReasonList struct {
	NsiContractPriceChangeReason []ZfcsNsiContractPriceChangeReasonType `xml:"nsiContractPriceChangeReason,omitempty" bson:"nsiContractPriceChangeReason,omitempty"`
}

// NsiContractRefusalReasonList is generated from an XSD element
type NsiContractRefusalReasonList struct {
	NsiContractRefusalReason []ZfcsNsiContractRefusalReasonType `xml:"nsiContractRefusalReason,omitempty" bson:"nsiContractRefusalReason,omitempty"`
}

// NsiContractSingleCustomerReasonList is generated from an XSD element
type NsiContractSingleCustomerReasonList struct {
	NsiContractSingleCustomerReason []ZfcsNsiContractSingleCustomerReasonType `xml:"nsiContractSingleCustomerReason,omitempty" bson:"nsiContractSingleCustomerReason,omitempty"`
}

// NsiContractTerminationReason is generated from an XSD element
type NsiContractTerminationReason struct {
	NsiContractTerminationReason []ZfcsNsiContractTerminationReasonType `xml:"nsiContractTerminationReason,omitempty" bson:"nsiContractTerminationReason,omitempty"`
}

// NsiContractModificationReason is generated from an XSD element
type NsiContractModificationReason struct {
	NsiContractModificationReason []ZfcsNsiContractModificationReasonType `xml:"nsiContractModificationReason,omitempty" bson:"nsiContractModificationReason,omitempty"`
}

// NsiContractExecutionDoc is generated from an XSD element
type NsiContractExecutionDoc struct {
	NsiContractExecutionDoc []ZfcsNsiContractExecutionDocType `xml:"nsiContractExecutionDoc,omitempty" bson:"nsiContractExecutionDoc,omitempty"`
}

// NsiContractReparationDoc is generated from an XSD element
type NsiContractReparationDoc struct {
	NsiContractReparationDoc []ZfcsNsiContractReparationDocType `xml:"nsiContractReparationDoc,omitempty" bson:"nsiContractReparationDoc,omitempty"`
}

// NsiContractPenaltyReason is generated from an XSD element
type NsiContractPenaltyReason struct {
	NsiContractPenaltyReason []ZfcsNsiContractPenaltyReasonType `xml:"nsiContractPenaltyReason,omitempty" bson:"nsiContractPenaltyReason,omitempty"`
}

// NsiContractOKOPFExtraBudget is generated from an XSD element
type NsiContractOKOPFExtraBudget struct {
	NsiContractOKOPFExtraBudget []ZfcsNsiContractOKOPFExtraBudgetType `xml:"nsiContractOKOPFExtraBudget,omitempty" bson:"nsiContractOKOPFExtraBudget,omitempty"`
}

// NsiContractCurrencyCBRF is generated from an XSD element
type NsiContractCurrencyCBRF struct {
	NsiContractCurrencyCBRF []ZfcsNsiContractCurrencyCBRFType `xml:"nsiContractCurrencyCBRF,omitempty" bson:"nsiContractCurrencyCBRF,omitempty"`
}

// NsiCurrencyList is generated from an XSD element
type NsiCurrencyList struct {
	NsiCurrency []ZfcsNsiCurrencyType `xml:"nsiCurrency,omitempty" bson:"nsiCurrency,omitempty"`
}

// NsiEvalCriterionList is generated from an XSD element
type NsiEvalCriterionList struct {
	NsiEvalCriterion []ZfcsNsiEvalCriterionType `xml:"nsiEvalCriterion,omitempty" bson:"nsiEvalCriterion,omitempty"`
}

// NsiKBKBudgetList is generated from an XSD element
type NsiKBKBudgetList struct {
	NsiKBKBudget []ZfcsNsiKBKBudgetType `xml:"nsiKBKBudget,omitempty" bson:"nsiKBKBudget,omitempty"`
}

// NsiKOSGUList is generated from an XSD element
type NsiKOSGUList struct {
	NsiKOSGU []ZfcsNsiKOSGUType `xml:"nsiKOSGU,omitempty" bson:"nsiKOSGU,omitempty"`
}

// NsiOffBudgetList is generated from an XSD element
type NsiOffBudgetList struct {
	NsiOffBudget []ZfcsNsiOffBudgetType `xml:"nsiOffBudget,omitempty" bson:"nsiOffBudget,omitempty"`
}

// NsiOKEIList is generated from an XSD element
type NsiOKEIList struct {
	NsiOKEI []ZfcsNsiOKEIType `xml:"nsiOKEI,omitempty" bson:"nsiOKEI,omitempty"`
}

// NsiOKOPFList is generated from an XSD element
type NsiOKOPFList struct {
	NsiOKOPF []ZfcsNsiOKOPFType `xml:"nsiOKOPF,omitempty" bson:"nsiOKOPF,omitempty"`
}

// NsiOKPDList is generated from an XSD element
type NsiOKPDList struct {
	NsiOKPD []ZfcsNsiOKPDType `xml:"nsiOKPD,omitempty" bson:"nsiOKPD,omitempty"`
}

// NsiOKSMList is generated from an XSD element
type NsiOKSMList struct {
	NsiOKSM []ZfcsNsiOKSMType `xml:"nsiOKSM,omitempty" bson:"nsiOKSM,omitempty"`
}

// NsiOKVEDList is generated from an XSD element
type NsiOKVEDList struct {
	NsiOKVED []ZfcsNsiOKVEDType `xml:"nsiOKVED,omitempty" bson:"nsiOKVED,omitempty"`
}

// NsiOKTMOList is generated from an XSD element
type NsiOKTMOList struct {
	NsiOKTMO []ZfcsNsiOKTMOType `xml:"nsiOKTMO,omitempty" bson:"nsiOKTMO,omitempty"`
}

// NsiOrganizationList is generated from an XSD element
type NsiOrganizationList struct {
	NsiOrganization []ZfcsNsiOrganizationType `xml:"nsiOrganization,omitempty" bson:"nsiOrganization,omitempty"`
}

// NsiOrganizationRightsList is generated from an XSD element
type NsiOrganizationRightsList struct {
	NsiOrganizationRights []ZfcsNsiOrganizationRightsType `xml:"nsiOrganizationRights,omitempty" bson:"nsiOrganizationRights,omitempty"`
}

// NsiOrganizationTypesList is generated from an XSD element
type NsiOrganizationTypesList struct {
	NsiOrganizationType []ZfcsNsiOrganizationTypesType `xml:"nsiOrganizationType,omitempty" bson:"nsiOrganizationType,omitempty"`
}

// NsiPlacingWayList is generated from an XSD element
type NsiPlacingWayList struct {
	NsiPlacingWay []ZfcsNsiPlacingWayType `xml:"nsiPlacingWay,omitempty" bson:"nsiPlacingWay,omitempty"`
}

// NsiPlanPositionChangeReasonList is generated from an XSD element
type NsiPlanPositionChangeReasonList struct {
	NsiPlanPositionChangeReason []ZfcsNsiPlanPositionChangeReasonType `xml:"nsiPlanPositionChangeReason,omitempty" bson:"nsiPlanPositionChangeReason,omitempty"`
}

// NsiPurchaseDocumentTypesList is generated from an XSD element
type NsiPurchaseDocumentTypesList struct {
	NsiPurchaseDocumentTypes []ZfcsNsiPurchaseDocumentTypesType `xml:"nsiPurchaseDocumentTypes,omitempty" bson:"nsiPurchaseDocumentTypes,omitempty"`
}

// NsiPurchasePreferenceList is generated from an XSD element
type NsiPurchasePreferenceList struct {
	NsiPurchasePreference []ZfcsNsiPurchasePreferencesType `xml:"nsiPurchasePreference,omitempty" bson:"nsiPurchasePreference,omitempty"`
}

// NsiPurchaseRejectReasonList is generated from an XSD element
type NsiPurchaseRejectReasonList struct {
	NsiPurchaseRejectReason []ZfcsNsiPurchaseRejectReasonType `xml:"nsiPurchaseRejectReason,omitempty" bson:"nsiPurchaseRejectReason,omitempty"`
}

// NsiUserList is generated from an XSD element
type NsiUserList struct {
	NsiUser []ZfcsNsiUserType `xml:"nsiUser,omitempty" bson:"nsiUser,omitempty"`
}

// NsiAbandonedReasonList is generated from an XSD element
type NsiAbandonedReasonList struct {
	NsiAbandonedReason []ZfcsNsiAbandonedReasonType `xml:"nsiAbandonedReason,omitempty" bson:"nsiAbandonedReason,omitempty"`
}

// NsiOKPD2List is generated from an XSD element
type NsiOKPD2List struct {
	NsiOKPD2 []ZfcsNsiOKPD2Type `xml:"nsiOKPD2,omitempty" bson:"nsiOKPD2,omitempty"`
}

// NsiOKVED2List is generated from an XSD element
type NsiOKVED2List struct {
	NsiOKVED2 []ZfcsNsiOKVED2Type `xml:"nsiOKVED2,omitempty" bson:"nsiOKVED2,omitempty"`
}

// NsiDeviationFactFoundationList is generated from an XSD element
type NsiDeviationFactFoundationList struct {
	NsiDeviationFactFoundation []ZfcsNsiDeviationFactFoundationType `xml:"nsiDeviationFactFoundation,omitempty" bson:"nsiDeviationFactFoundation,omitempty"`
}

// NsiPublicDiscussionQuestionnaries is generated from an XSD element
type NsiPublicDiscussionQuestionnaries struct {
	NsiPublicDiscussionQuestionnarie []ZfcsNsiPublicDiscussionQuestionnarieType `xml:"nsiPublicDiscussionQuestionnarie,omitempty" bson:"nsiPublicDiscussionQuestionnarie,omitempty"`
}

// NsiPublicDiscussionDecisions is generated from an XSD element
type NsiPublicDiscussionDecisions struct {
	NsiPublicDiscussionDecision []ZfcsNsiPublicDiscussionDecisionsType `xml:"nsiPublicDiscussionDecision,omitempty" bson:"nsiPublicDiscussionDecision,omitempty"`
}

// NsiAuditActionSubjects is generated from an XSD element
type NsiAuditActionSubjects struct {
	NsiAuditActionSubject []ZfcsNsiAuditActionSubjectsType `xml:"nsiAuditActionSubject,omitempty" bson:"nsiAuditActionSubject,omitempty"`
}

// NsiBusinessControls is generated from an XSD element
type NsiBusinessControls struct {
	NsiBusinessControls []ZfcsNsiBusinessControlType `xml:"nsiBusinessControls,omitempty" bson:"nsiBusinessControls,omitempty"`
}

// ZfcsOKTMORef is generated from an XSD element
type ZfcsOKTMORef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ContractSignType is generated from an XSD element
type ContractSignType struct {
	ID            int64           `xml:"id,omitempty" bson:"id,omitempty"`
	Number        string          `xml:"number,omitempty" bson:"number,omitempty"`
	SignDate      xsd.Date        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	Foundation    Foundation      `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Customer      OrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
	ProtocolDate  xsd.Date        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	Suppliers     Suppliers       `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	Scan          DocumentList    `xml:"scan,omitempty" bson:"scan,omitempty"`
	DocumentMetas DocumentList    `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// Foundation is generated from an XSD element
type Foundation struct {
	Order Order `xml:"order,omitempty" bson:"order,omitempty"`
}

// Order is generated from an XSD element
type Order struct {
	NotificationNumber       string `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationProtocolNumber string `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
}

// Suppliers is generated from an XSD element
type Suppliers struct {
	Supplier []ParticipantType `xml:"supplier,omitempty" bson:"supplier,omitempty"`
}

// NotificationEFType is generated from an XSD element
type NotificationEFType struct {
	ID                           int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber           string                 `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber string                 `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                int64                  `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                   time.Time              `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                   PlacingWay             `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                    string                 `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                        Order                  `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                  ContactInfoType        `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                    Document               `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                DocumentList           `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                  time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate           time.Time              `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                 ModificationType       `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                         string                 `xml:"href,omitempty" bson:"href,omitempty"`
	NotificationCommission       NotificationCommission `xml:"notificationCommission,omitempty" bson:"notificationCommission,omitempty"`
	Lots                         Lots                   `xml:"lots,omitempty" bson:"lots,omitempty"`
	EP                           EPType                 `xml:"EP,omitempty" bson:"EP,omitempty"`
}

// PlacingWay is generated from an XSD element
type PlacingWay struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// NotificationCommission is generated from an XSD element
type NotificationCommission struct {
	P1Date time.Time `xml:"p1Date,omitempty" bson:"p1Date,omitempty"`
	P2Date time.Time `xml:"p2Date,omitempty" bson:"p2Date,omitempty"`
	P3Date time.Time `xml:"p3Date,omitempty" bson:"p3Date,omitempty"`
}

// Lots is generated from an XSD element
type Lots struct {
	Lot Lot `xml:"lot,omitempty" bson:"lot,omitempty"`
}

// Lot is generated from an XSD element
type Lot struct {
	Sid                  int64                `xml:"sid,omitempty" bson:"sid,omitempty"`
	OrdinalNumber        int64                `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	Subject              string               `xml:"subject,omitempty" bson:"subject,omitempty"`
	Currency             CurrencyType         `xml:"currency,omitempty" bson:"currency,omitempty"`
	Products             ProductsType         `xml:"products,omitempty" bson:"products,omitempty"`
	NotificationFeatures NotificationFeatures `xml:"notificationFeatures,omitempty" bson:"notificationFeatures,omitempty"`
	CustomerRequirements CustomerRequirements `xml:"customerRequirements,omitempty" bson:"customerRequirements,omitempty"`
	AuctionItems         AuctionItemsType     `xml:"auctionItems,omitempty" bson:"auctionItems,omitempty"`
	DocumentRequirements DocumentRequirements `xml:"documentRequirements,omitempty" bson:"documentRequirements,omitempty"`
	AuctionProducts      AuctionProducts      `xml:"auctionProducts,omitempty" bson:"auctionProducts,omitempty"`
	LotDocRequirements   LotDocRequirements   `xml:"lotDocRequirements,omitempty" bson:"lotDocRequirements,omitempty"`
}

// NotificationFeatures is generated from an XSD element
type NotificationFeatures struct {
	NotificationFeature []NotificationFeatureType `xml:"notificationFeature,omitempty" bson:"notificationFeature,omitempty"`
}

// CustomerRequirements is generated from an XSD element
type CustomerRequirements struct {
	CustomerRequirement []CustomerRequirement `xml:"customerRequirement,omitempty" bson:"customerRequirement,omitempty"`
}

// CustomerRequirement is generated from an XSD element
type CustomerRequirement struct {
	Sid               int64             `xml:"sid,omitempty" bson:"sid,omitempty"`
	Quantity          string            `xml:"quantity,omitempty" bson:"quantity,omitempty"`
	MaxPrice          float64           `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	Customer          OrganizationRef   `xml:"customer,omitempty" bson:"customer,omitempty"`
	DeliveryPlace     string            `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	DeliveryTerm      string            `xml:"deliveryTerm,omitempty" bson:"deliveryTerm,omitempty"`
	GuaranteeApp      GuaranteeApp      `xml:"guaranteeApp,omitempty" bson:"guaranteeApp,omitempty"`
	GuaranteeContract GuaranteeContract `xml:"guaranteeContract,omitempty" bson:"guaranteeContract,omitempty"`
	AdditionalInfo    string            `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	KBK               KBKType           `xml:"KBK,omitempty" bson:"KBK,omitempty"`
	TenderPlanInfo    TendePlanInfoType `xml:"tenderPlanInfo,omitempty" bson:"tenderPlanInfo,omitempty"`
}

// GuaranteeApp is generated from an XSD element
type GuaranteeApp struct {
	Procedure         string  `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string  `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string  `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string  `xml:"bik,omitempty" bson:"bik,omitempty"`
	Amount            float64 `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// GuaranteeContract is generated from an XSD element
type GuaranteeContract struct {
	Procedure         string  `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string  `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string  `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string  `xml:"bik,omitempty" bson:"bik,omitempty"`
	IsBail            bool    `xml:"isBail,omitempty" bson:"isBail,omitempty"`
	Amount            float64 `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// DocumentRequirements is generated from an XSD element
type DocumentRequirements struct {
	DocumentRequirement []DocumentRequirement `xml:"documentRequirement,omitempty" bson:"documentRequirement,omitempty"`
}

// DocumentRequirement is generated from an XSD element
type DocumentRequirement struct {
	Sid           int64  `xml:"sid,omitempty" bson:"sid,omitempty"`
	OrdinalNumber int64  `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	ReqValue      string `xml:"reqValue,omitempty" bson:"reqValue,omitempty"`
	DocName       string `xml:"docName,omitempty" bson:"docName,omitempty"`
}

// AuctionProducts is generated from an XSD element
type AuctionProducts struct {
	AuctionProduct []AuctionProduct `xml:"auctionProduct,omitempty" bson:"auctionProduct,omitempty"`
}

// AuctionProduct is generated from an XSD element
type AuctionProduct struct {
	Sid                int64                `xml:"sid,omitempty" bson:"sid,omitempty"`
	OrdinalNumber      int64                `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	ProductName        string               `xml:"productName,omitempty" bson:"productName,omitempty"`
	TradeMark          string               `xml:"tradeMark,omitempty" bson:"tradeMark,omitempty"`
	ProductRequirement []ProductRequirement `xml:"productRequirement,omitempty" bson:"productRequirement,omitempty"`
	EquivalenceParam   []EquivalenceParam   `xml:"equivalenceParam,omitempty" bson:"equivalenceParam,omitempty"`
}

// ProductRequirement is generated from an XSD element
type ProductRequirement struct {
	Sid           int64  `xml:"sid,omitempty" bson:"sid,omitempty"`
	OrdinalNumber int64  `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	Requirement   string `xml:"requirement,omitempty" bson:"requirement,omitempty"`
}

// EquivalenceParam is generated from an XSD element
type EquivalenceParam struct {
	Sid           int64  `xml:"sid,omitempty" bson:"sid,omitempty"`
	OrdinalNumber int64  `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	Name          string `xml:"name,omitempty" bson:"name,omitempty"`
	ParamType     string `xml:"paramType,omitempty" bson:"paramType,omitempty"`
	ParamValue    string `xml:"paramValue,omitempty" bson:"paramValue,omitempty"`
	Modifiable    bool   `xml:"modifiable,omitempty" bson:"modifiable,omitempty"`
}

// LotDocRequirements is generated from an XSD element
type LotDocRequirements struct {
	DocReq1111 DocReqType `xml:"docReq-1.1.11,omitempty" bson:"docReq-1.1.11,omitempty"`
	DocReq1211 DocReqType `xml:"docReq-1.2.11,omitempty" bson:"docReq-1.2.11,omitempty"`
	DocReq2111 DocReqType `xml:"docReq-2.1.11,omitempty" bson:"docReq-2.1.11,omitempty"`
}

// ZfcsContactInfoType is generated from an XSD element
type ZfcsContactInfoType struct {
	OrgPostAddress string                `xml:"orgPostAddress,omitempty" bson:"orgPostAddress,omitempty"`
	OrgFactAddress string                `xml:"orgFactAddress,omitempty" bson:"orgFactAddress,omitempty"`
	ContactPerson  ZfcsContactPersonType `xml:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	ContactEMail   string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone   string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax     string                `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AddInfo        string                `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsProtocolOKD1Type is generated from an XSD element
type ZfcsProtocolOKD1Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                 `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	IsPrequalification       bool                         `xml:"isPrequalification,omitempty" bson:"isPrequalification,omitempty"`
}

// ProtocolPublisher is generated from an XSD element
type ProtocolPublisher struct {
	PublisherOrg  ZfcsPurchaseOrganizationType `xml:"publisherOrg,omitempty" bson:"publisherOrg,omitempty"`
	PublisherRole string                       `xml:"publisherRole,omitempty" bson:"publisherRole,omitempty"`
}

// PurchaseInfo is generated from an XSD element
type PurchaseInfo struct {
	PurchaseResponsible    PurchaseResponsible `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay             ZfcsPlacingWayType  `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	PublishDate            time.Time           `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PurchaseObjectInfo     string              `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	NotificationFullNumber string              `xml:"notificationFullNumber,omitempty" bson:"notificationFullNumber,omitempty"`
	NotificationFullName   string              `xml:"notificationFullName,omitempty" bson:"notificationFullName,omitempty"`
}

// PurchaseResponsible is generated from an XSD element
type PurchaseResponsible struct {
	ResponsibleOrg  ZfcsOrganizationRef `xml:"responsibleOrg,omitempty" bson:"responsibleOrg,omitempty"`
	ResponsibleRole string              `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// ProtocolLots is generated from an XSD element
type ProtocolLots struct {
	ProtocolLot []ProtocolLot `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ProtocolLot is generated from an XSD element
type ProtocolLot struct {
	LotNumber            uint64                  `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	LotInfo              ZfcsLotInfoType         `xml:"lotInfo,omitempty" bson:"lotInfo,omitempty"`
	DocumentRequirements DocumentRequirements    `xml:"documentRequirements,omitempty" bson:"documentRequirements,omitempty"`
	Applications         Applications            `xml:"applications,omitempty" bson:"applications,omitempty"`
	AbandonedReason      ZfcsAbandonedReasonType `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
}

// Applications is generated from an XSD element
type Applications struct {
	Application []Application `xml:"application,omitempty" bson:"application,omitempty"`
}

// Application is generated from an XSD element
type Application struct {
	JournalNumber       string              `xml:"journalNumber,omitempty" bson:"journalNumber,omitempty"`
	AppDate             time.Time           `xml:"appDate,omitempty" bson:"appDate,omitempty"`
	AppParticipants     AppParticipants     `xml:"appParticipants,omitempty" bson:"appParticipants,omitempty"`
	DocumentCompliances DocumentCompliances `xml:"documentCompliances,omitempty" bson:"documentCompliances,omitempty"`
}

// AppParticipants is generated from an XSD element
type AppParticipants struct {
	AppParticipant []ZfcsParticipantType `xml:"appParticipant,omitempty" bson:"appParticipant,omitempty"`
}

// DocumentCompliances is generated from an XSD element
type DocumentCompliances struct {
	DocumentCompliance []DocumentCompliance `xml:"documentCompliance,omitempty" bson:"documentCompliance,omitempty"`
}

// DocumentCompliance is generated from an XSD element
type DocumentCompliance struct {
	Number     uint64 `xml:"number,omitempty" bson:"number,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Mandatory  bool   `xml:"mandatory,omitempty" bson:"mandatory,omitempty"`
	Compliance string `xml:"compliance,omitempty" bson:"compliance,omitempty"`
	OtherInfo  string `xml:"otherInfo,omitempty" bson:"otherInfo,omitempty"`
}

// ZfcsProtocolOKD2Type is generated from an XSD element
type ZfcsProtocolOKD2Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsElectronicComplaintType is generated from an XSD element
type ZfcsElectronicComplaintType struct {
	SchemeVersion string                            `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID            int64                             `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID    string                            `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	CommonInfo    CommonInfo                        `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Indicted      []ZfcsComplaintProjectSubjectType `xml:"indicted,omitempty" bson:"indicted,omitempty"`
	Applicant     ZfcsApplicantType                 `xml:"applicant,omitempty" bson:"applicant,omitempty"`
	Object        ZfcsComplaintObjectType           `xml:"object,omitempty" bson:"object,omitempty"`
	Text          string                            `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm     ZfcsPrintFormType                 `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType              `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType            `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReturnInfo    ZfcsComplaintReturnInfoType       `xml:"returnInfo,omitempty" bson:"returnInfo,omitempty"`
}

// CommonInfo is generated from an XSD element
type CommonInfo struct {
	RegDate          time.Time           `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	ComplaintNumber  string              `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	VersionNumber    int64               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PlanDecisionDate time.Time           `xml:"planDecisionDate,omitempty" bson:"planDecisionDate,omitempty"`
	DecisionPlace    string              `xml:"decisionPlace,omitempty" bson:"decisionPlace,omitempty"`
	RegistrationKO   ZfcsOrganizationRef `xml:"registrationKO,omitempty" bson:"registrationKO,omitempty"`
	ConsiderationKO  ZfcsOrganizationRef `xml:"considerationKO,omitempty" bson:"considerationKO,omitempty"`
	RegType          string              `xml:"regType,omitempty" bson:"regType,omitempty"`
}

// ZfcsComplaintCancelInfoType is generated from an XSD element
type ZfcsComplaintCancelInfoType struct {
	Name        string                 `xml:"name,omitempty" bson:"name,omitempty"`
	Text        string                 `xml:"text,omitempty" bson:"text,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsLotISType is generated from an XSD element
type ZfcsLotISType struct {
	LotNumber        uint64                   `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	MaxPrice         string                   `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	Currency         ZfcsCurrencyRef          `xml:"currency,omitempty" bson:"currency,omitempty"`
	Preferenses      Preferenses              `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements     Requirements             `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	AddInfo          string                   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	PublicDiscussion ZfcsPublicDiscussionType `xml:"publicDiscussion,omitempty" bson:"publicDiscussion,omitempty"`
	OKPD             []ZfcsOKPDRef            `xml:"OKPD,omitempty" bson:"OKPD,omitempty"`
	OKPD2            []ZfcsOKPDRef            `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
}

// Preferenses is generated from an XSD element
type Preferenses struct {
	Preferense []ZfcsPreferenseType `xml:"preferense,omitempty" bson:"preferense,omitempty"`
}

// Requirements is generated from an XSD element
type Requirements struct {
	Requirement []ZfcsRequirementType `xml:"requirement,omitempty" bson:"requirement,omitempty"`
}

// ZfcsNsiPurchasePreferencesType is generated from an XSD element
type ZfcsNsiPurchasePreferencesType struct {
	ID              int64       `xml:"id,omitempty" bson:"id,omitempty"`
	Name            string      `xml:"name,omitempty" bson:"name,omitempty"`
	ShortName       string      `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	Type            string      `xml:"type,omitempty" bson:"type,omitempty"`
	PrefEstimateApp bool        `xml:"prefEstimateApp,omitempty" bson:"prefEstimateApp,omitempty"`
	PrefValue       float64     `xml:"prefValue,omitempty" bson:"prefValue,omitempty"`
	PlacingWays     PlacingWays `xml:"placingWays,omitempty" bson:"placingWays,omitempty"`
	Actual          bool        `xml:"actual,omitempty" bson:"actual,omitempty"`
	UseTenderPlans  bool        `xml:"useTenderPlans,omitempty" bson:"useTenderPlans,omitempty"`
}

// PlacingWays is generated from an XSD element
type PlacingWays struct {
	Code []string `xml:"code,omitempty" bson:"code,omitempty"`
}

// ZfcsTOFKRef is generated from an XSD element
type ZfcsTOFKRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsUserTenderPlanType is generated from an XSD element
type ZfcsUserTenderPlanType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	Phone      string `xml:"phone,omitempty" bson:"phone,omitempty"`
	Fax        string `xml:"fax,omitempty" bson:"fax,omitempty"`
	Email      string `xml:"email,omitempty" bson:"email,omitempty"`
}

// ZfcsNsiAbandonedReasonType is generated from an XSD element
type ZfcsNsiAbandonedReasonType struct {
	ID         int64              `xml:"id,omitempty" bson:"id,omitempty"`
	Code       string             `xml:"code,omitempty" bson:"code,omitempty"`
	Name       string             `xml:"name,omitempty" bson:"name,omitempty"`
	ObjectName string             `xml:"objectName,omitempty" bson:"objectName,omitempty"`
	Type       string             `xml:"type,omitempty" bson:"type,omitempty"`
	DocType    ZfcsDocType        `xml:"docType,omitempty" bson:"docType,omitempty"`
	PlacingWay ZfcsPlacingWayType `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Actual     bool               `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsCommissionMemberType is generated from an XSD element
type ZfcsCommissionMemberType struct {
	MemberNumber uint64                 `xml:"memberNumber,omitempty" bson:"memberNumber,omitempty"`
	LastName     string                 `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName    string                 `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName   string                 `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	Role         ZfcsCommissionRoleType `xml:"role,omitempty" bson:"role,omitempty"`
}

// ZfcsNsiContractTerminationReasonType is generated from an XSD element
type ZfcsNsiContractTerminationReasonType struct {
	ID                   int64                `xml:"id,omitempty" bson:"id,omitempty"`
	Code                 string               `xml:"code,omitempty" bson:"code,omitempty"`
	Name                 string               `xml:"name,omitempty" bson:"name,omitempty"`
	SubsystemType        string               `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
	Actual               bool                 `xml:"actual,omitempty" bson:"actual,omitempty"`
	Documents            Documents            `xml:"documents,omitempty" bson:"documents,omitempty"`
	ReparationsDocuments ReparationsDocuments `xml:"reparationsDocuments,omitempty" bson:"reparationsDocuments,omitempty"`
}

// Documents is generated from an XSD element
type Documents struct {
	Document []Document `xml:"document,omitempty" bson:"document,omitempty"`
}

// Document is generated from an XSD element
type Document struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ReparationsDocuments is generated from an XSD element
type ReparationsDocuments struct {
	Document []Document `xml:"document,omitempty" bson:"document,omitempty"`
}

// ZfcsNsiOKTMOType is generated from an XSD element
type ZfcsNsiOKTMOType struct {
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	ParentCode string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	Section    string `xml:"section,omitempty" bson:"section,omitempty"`
	FullName   string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNotificationZakKOUType is generated from an XSD element
type ZfcsNotificationZakKOUType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKOUType        `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsPurchaseApprovalType is generated from an XSD element
type ZfcsPurchaseApprovalType struct {
	SchemeVersion  string              `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ReceiptNumber  string              `xml:"receiptNumber,omitempty" bson:"receiptNumber,omitempty"`
	Authority      ZfcsOrganizationRef `xml:"authority,omitempty" bson:"authority,omitempty"`
	ApprovalDate   time.Time           `xml:"approvalDate,omitempty" bson:"approvalDate,omitempty"`
	ApprovalResult bool                `xml:"approvalResult,omitempty" bson:"approvalResult,omitempty"`
	Reason         string              `xml:"reason,omitempty" bson:"reason,omitempty"`
	AddInfo        string              `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	PublishDate    time.Time           `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Purchase       Purchase            `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Documents      Documents           `xml:"documents,omitempty" bson:"documents,omitempty"`
}

// Purchase is generated from an XSD element
type Purchase struct {
	PurchaseNumber     string              `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber          string              `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href               string              `xml:"href,omitempty" bson:"href,omitempty"`
	ResponsibleOrg     ZfcsOrganizationRef `xml:"responsibleOrg,omitempty" bson:"responsibleOrg,omitempty"`
	PurchaseObjectInfo string              `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PlacingWay         ZfcsPlacingWayType  `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
}

// ZfcsFinanceResourcesType is generated from an XSD element
type ZfcsFinanceResourcesType struct {
	CurrentYear string `xml:"currentYear,omitempty" bson:"currentYear,omitempty"`
	FirstYear   string `xml:"firstYear,omitempty" bson:"firstYear,omitempty"`
	SecondYear  string `xml:"secondYear,omitempty" bson:"secondYear,omitempty"`
	SubsecYears string `xml:"subsecYears,omitempty" bson:"subsecYears,omitempty"`
	Total       string `xml:"total,omitempty" bson:"total,omitempty"`
}

// ZfcsPublicDiscussionFormType is generated from an XSD element
type ZfcsPublicDiscussionFormType struct {
	SchemeVersion          string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID             string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID                     int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	PublicDiscussionNum    string                               `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	VersionNumber          string                               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate         time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	FormNumber             string                               `xml:"formNumber,omitempty" bson:"formNumber,omitempty"`
	Author                 Author                               `xml:"author,omitempty" bson:"author,omitempty"`
	Phase                  string                               `xml:"phase,omitempty" bson:"phase,omitempty"`
	PublicDiscussionFacets PublicDiscussionFacets               `xml:"publicDiscussionFacets,omitempty" bson:"publicDiscussionFacets,omitempty"`
	Purchase               ZfcsPublicDiscussionPurchaseInfoType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
}

// Author is generated from an XSD element
type Author struct {
	Name  string `xml:"name,omitempty" bson:"name,omitempty"`
	EMail string `xml:"eMail,omitempty" bson:"eMail,omitempty"`
}

// PublicDiscussionFacets is generated from an XSD element
type PublicDiscussionFacets struct {
	PublicDiscussionFacet []PublicDiscussionFacet `xml:"publicDiscussionFacet,omitempty" bson:"publicDiscussionFacet,omitempty"`
}

// PublicDiscussionFacet is generated from an XSD element
type PublicDiscussionFacet struct {
	Code                     string                     `xml:"code,omitempty" bson:"code,omitempty"`
	FacetName                string                     `xml:"facetName,omitempty" bson:"facetName,omitempty"`
	PublicDiscussionQuestion []PublicDiscussionQuestion `xml:"publicDiscussionQuestion,omitempty" bson:"publicDiscussionQuestion,omitempty"`
}

// PublicDiscussionQuestion is generated from an XSD element
type PublicDiscussionQuestion struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	UserAnswer  string `xml:"userAnswer,omitempty" bson:"userAnswer,omitempty"`
	UserComment string `xml:"userComment,omitempty" bson:"userComment,omitempty"`
}

// ZfcsCheckPlanType is generated from an XSD element
type ZfcsCheckPlanType struct {
	SchemeVersion string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo           `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	StartStage    ZfcsStageType        `xml:"startStage,omitempty" bson:"startStage,omitempty"`
	EndStage      ZfcsStageType        `xml:"endStage,omitempty" bson:"endStage,omitempty"`
	CheckList     CheckList            `xml:"checkList,omitempty" bson:"checkList,omitempty"`
	PrintForm     ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// CheckList is generated from an XSD element
type CheckList struct {
	CheckInfo []CheckInfo `xml:"checkInfo,omitempty" bson:"checkInfo,omitempty"`
}

// CheckInfo is generated from an XSD element
type CheckInfo struct {
	CheckNumber      string                   `xml:"checkNumber,omitempty" bson:"checkNumber,omitempty"`
	CheckSubject     ZfcsCheckSubjectPlanType `xml:"checkSubject,omitempty" bson:"checkSubject,omitempty"`
	CheckStartStage  ZfcsStageType            `xml:"checkStartStage,omitempty" bson:"checkStartStage,omitempty"`
	CheckPublishDate time.Time                `xml:"checkPublishDate,omitempty" bson:"checkPublishDate,omitempty"`
	Base             string                   `xml:"base,omitempty" bson:"base,omitempty"`
}

// ZfcsContract2015DocumentInfo is generated from an XSD element
type ZfcsContract2015DocumentInfo struct {
	DocumentName string   `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
}

// ZfcsCustomerReportContractExecutionType is generated from an XSD element
type ZfcsCustomerReportContractExecutionType struct {
	SchemeVersion         string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate               time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate        time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber         string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg            ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href                  string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments           ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID              string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	ModificationReason    string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	Signer                Signer                           `xml:"signer,omitempty" bson:"signer,omitempty"`
	Customer              Customer                         `xml:"customer,omitempty" bson:"customer,omitempty"`
	ContractInfo          ContractInfo                     `xml:"contractInfo,omitempty" bson:"contractInfo,omitempty"`
	Suppliers             Suppliers                        `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	ExecutionInfo         ExecutionInfo                    `xml:"executionInfo,omitempty" bson:"executionInfo,omitempty"`
	ImproperExecutionInfo ImproperExecutionInfo            `xml:"improperExecutionInfo,omitempty" bson:"improperExecutionInfo,omitempty"`
	ModifyTerminateInfo   ModifyTerminateInfo              `xml:"modifyTerminateInfo,omitempty" bson:"modifyTerminateInfo,omitempty"`
}

// Signer is generated from an XSD element
type Signer struct {
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	Position   string `xml:"position,omitempty" bson:"position,omitempty"`
}

// Customer is generated from an XSD element
type Customer struct {
	RegNum          string       `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string       `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string       `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string       `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	INN             string       `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string       `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	Adress          string       `xml:"adress,omitempty" bson:"adress,omitempty"`
	Phone           string       `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email           string       `xml:"email,omitempty" bson:"email,omitempty"`
	OKTMO           ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	OKPO            ZfcsOKPORef  `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	OKOPF           ZfcsOkopfRef `xml:"OKOPF,omitempty" bson:"OKOPF,omitempty"`
	PPOName         string       `xml:"PPOName,omitempty" bson:"PPOName,omitempty"`
	OKFS            ZfcsOKFSRef  `xml:"OKFS,omitempty" bson:"OKFS,omitempty"`
}

// ContractInfo is generated from an XSD element
type ContractInfo struct {
	ContractRegNum string    `xml:"contractRegNum,omitempty" bson:"contractRegNum,omitempty"`
	ContractNumber string    `xml:"contractNumber,omitempty" bson:"contractNumber,omitempty"`
	Product        []Product `xml:"product,omitempty" bson:"product,omitempty"`
	PurchaseCode   string    `xml:"purchaseCode,omitempty" bson:"purchaseCode,omitempty"`
	Okpd2okved2    bool      `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// Product is generated from an XSD element
type Product struct {
	Name           string           `xml:"name,omitempty" bson:"name,omitempty"`
	FinanceSources []FinanceSources `xml:"financeSources,omitempty" bson:"financeSources,omitempty"`
	OKPD           ZfcsOKPDRef      `xml:"OKPD,omitempty" bson:"OKPD,omitempty"`
	OKPD2          ZfcsOKPDRef      `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
}

// FinanceSources is generated from an XSD element
type FinanceSources struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ExecutionInfo is generated from an XSD element
type ExecutionInfo struct {
	ProvideByContract ZfcsContractExecutionType `xml:"provideByContract,omitempty" bson:"provideByContract,omitempty"`
	Executed          ZfcsContractExecutionType `xml:"executed,omitempty" bson:"executed,omitempty"`
	Price             Price                     `xml:"price,omitempty" bson:"price,omitempty"`
	Quantity          Quantity                  `xml:"quantity,omitempty" bson:"quantity,omitempty"`
	AdvanceMarker     bool                      `xml:"advanceMarker,omitempty" bson:"advanceMarker,omitempty"`
	Advance           Advance                   `xml:"advance,omitempty" bson:"advance,omitempty"`
}

// Price is generated from an XSD element
type Price struct {
	PriceByContract string `xml:"priceByContract,omitempty" bson:"priceByContract,omitempty"`
	PriceExecution  string `xml:"priceExecution,omitempty" bson:"priceExecution,omitempty"`
	PriceDocument   string `xml:"priceDocument,omitempty" bson:"priceDocument,omitempty"`
	PriceNote       string `xml:"priceNote,omitempty" bson:"priceNote,omitempty"`
}

// Quantity is generated from an XSD element
type Quantity struct {
	QuantityItem     []QuantityItem `xml:"quantityItem,omitempty" bson:"quantityItem,omitempty"`
	QuantityDocument string         `xml:"quantityDocument,omitempty" bson:"quantityDocument,omitempty"`
	QuantityNote     string         `xml:"quantityNote,omitempty" bson:"quantityNote,omitempty"`
}

// QuantityItem is generated from an XSD element
type QuantityItem struct {
	OKEIByContract     ZfcsOKEIRef `xml:"OKEIByContract,omitempty" bson:"OKEIByContract,omitempty"`
	QuantityByContract string      `xml:"quantityByContract,omitempty" bson:"quantityByContract,omitempty"`
	OKEIExecution      ZfcsOKEIRef `xml:"OKEIExecution,omitempty" bson:"OKEIExecution,omitempty"`
	QuantityExecution  string      `xml:"quantityExecution,omitempty" bson:"quantityExecution,omitempty"`
}

// Advance is generated from an XSD element
type Advance struct {
	AdvanceItem     []AdvanceItem `xml:"advanceItem,omitempty" bson:"advanceItem,omitempty"`
	AdvanceDocument string        `xml:"advanceDocument,omitempty" bson:"advanceDocument,omitempty"`
	AdvanceNote     string        `xml:"advanceNote,omitempty" bson:"advanceNote,omitempty"`
}

// AdvanceItem is generated from an XSD element
type AdvanceItem struct {
	AdvanceByContract string   `xml:"advanceByContract,omitempty" bson:"advanceByContract,omitempty"`
	AdvanceDate       xsd.Date `xml:"advanceDate,omitempty" bson:"advanceDate,omitempty"`
	AdvanceExecution  string   `xml:"advanceExecution,omitempty" bson:"advanceExecution,omitempty"`
	ExecutionDate     xsd.Date `xml:"executionDate,omitempty" bson:"executionDate,omitempty"`
}

// ImproperExecutionInfo is generated from an XSD element
type ImproperExecutionInfo struct {
	ConformContractMarker bool                `xml:"conformContractMarker,omitempty" bson:"conformContractMarker,omitempty"`
	ImproperExecution     []ImproperExecution `xml:"improperExecution,omitempty" bson:"improperExecution,omitempty"`
}

// ImproperExecution is generated from an XSD element
type ImproperExecution struct {
	NameObligations  string `xml:"nameObligations,omitempty" bson:"nameObligations,omitempty"`
	EssenceViolation string `xml:"essenceViolation,omitempty" bson:"essenceViolation,omitempty"`
	PenaltyInfo      string `xml:"penaltyInfo,omitempty" bson:"penaltyInfo,omitempty"`
	PenaltyDoc       string `xml:"penaltyDoc,omitempty" bson:"penaltyDoc,omitempty"`
	Note             string `xml:"note,omitempty" bson:"note,omitempty"`
	Indicator        string `xml:"indicator,omitempty" bson:"indicator,omitempty"`
	IndicatorName    string `xml:"indicatorName,omitempty" bson:"indicatorName,omitempty"`
}

// ModifyTerminateInfo is generated from an XSD element
type ModifyTerminateInfo struct {
	ModifyContract    []ModifyContract  `xml:"modifyContract,omitempty" bson:"modifyContract,omitempty"`
	TerminateContract TerminateContract `xml:"terminateContract,omitempty" bson:"terminateContract,omitempty"`
}

// ModifyContract is generated from an XSD element
type ModifyContract struct {
	EventDate xsd.Date `xml:"eventDate,omitempty" bson:"eventDate,omitempty"`
	DocInfo   string   `xml:"docInfo,omitempty" bson:"docInfo,omitempty"`
	Reason    Reason   `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// Reason is generated from an XSD element
type Reason struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// TerminateContract is generated from an XSD element
type TerminateContract struct {
	EventDate xsd.Date `xml:"eventDate,omitempty" bson:"eventDate,omitempty"`
	DocInfo   string   `xml:"docInfo,omitempty" bson:"docInfo,omitempty"`
	Reason    Reason   `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsPurchaseProcedureCollectingType is generated from an XSD element
type ZfcsPurchaseProcedureCollectingType struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	Place     string    `xml:"place,omitempty" bson:"place,omitempty"`
	Order     string    `xml:"order,omitempty" bson:"order,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ZfcsNsiOrganizationTypesType is generated from an XSD element
type ZfcsNsiOrganizationTypesType struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsBigProjectMemberType is generated from an XSD element
type ZfcsBigProjectMemberType struct {
	Name    string `xml:"name,omitempty" bson:"name,omitempty"`
	INN     string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP     string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	Address string `xml:"address,omitempty" bson:"address,omitempty"`
}

// ZfcsNotificationZakKType is generated from an XSD element
type ZfcsNotificationZakKType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKType          `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsPublicDiscussionFacetRef is generated from an XSD element
type ZfcsPublicDiscussionFacetRef struct {
	Code      string `xml:"code,omitempty" bson:"code,omitempty"`
	FacetName string `xml:"facetName,omitempty" bson:"facetName,omitempty"`
}

// ZfcsProtocolZPFinalType is generated from an XSD element
type ZfcsProtocolZPFinalType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	FoundationProtocolName   string                       `xml:"foundationProtocolName,omitempty" bson:"foundationProtocolName,omitempty"`
}

// ZfcsNsiCommissionType is generated from an XSD element
type ZfcsNsiCommissionType struct {
	RegNumber         int64               `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	CommissionName    string              `xml:"commissionName,omitempty" bson:"commissionName,omitempty"`
	CommissionMembers CommissionMembers   `xml:"commissionMembers,omitempty" bson:"commissionMembers,omitempty"`
	Owner             ZfcsOrganizationRef `xml:"owner,omitempty" bson:"owner,omitempty"`
	Competent         bool                `xml:"competent,omitempty" bson:"competent,omitempty"`
	AddInfo           string              `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Actual            bool                `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// CommissionMembers is generated from an XSD element
type CommissionMembers struct {
	CommissionMember []CommissionMember `xml:"commissionMember,omitempty" bson:"commissionMember,omitempty"`
}

// CommissionMember is generated from an XSD element
type CommissionMember struct {
	LastName   string                      `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string                      `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string                      `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	ID         int64                       `xml:"id,omitempty" bson:"id,omitempty"`
	Role       []ZfcsNsiCommissionRoleType `xml:"role,omitempty" bson:"role,omitempty"`
}

// ZfcsNsiContractOKOPFExtraBudgetType is generated from an XSD element
type ZfcsNsiContractOKOPFExtraBudgetType struct {
	Extrabudget  ZfcsExtraBudgetFundsContract2015 `xml:"extrabudget,omitempty" bson:"extrabudget,omitempty"`
	LegalFormOld LegalFormOld                     `xml:"legalFormOld,omitempty" bson:"legalFormOld,omitempty"`
	LegalFormNew OkopfType                        `xml:"legalFormNew,omitempty" bson:"legalFormNew,omitempty"`
	Actual       bool                             `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// LegalFormOld is generated from an XSD element
type LegalFormOld struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ContactInfoType is generated from an XSD element
type ContactInfoType struct {
	OrgName        string            `xml:"orgName,omitempty" bson:"orgName,omitempty"`
	OrgFactAddress string            `xml:"orgFactAddress,omitempty" bson:"orgFactAddress,omitempty"`
	OrgPostAddress string            `xml:"orgPostAddress,omitempty" bson:"orgPostAddress,omitempty"`
	ContactPerson  ContactPersonType `xml:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	ContactEMail   string            `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone   string            `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax     string            `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
}

// ZfcsConfirmationType is generated from an XSD element
type ZfcsConfirmationType struct {
	SchemeVersion    string           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	LoadID           string           `xml:"loadId,omitempty" bson:"loadId,omitempty"`
	Result           string           `xml:"result,omitempty" bson:"result,omitempty"`
	Violations       Violations       `xml:"violations,omitempty" bson:"violations,omitempty"`
	LoadURL          string           `xml:"loadUrl,omitempty" bson:"loadUrl,omitempty"`
	RefID            string           `xml:"refId,omitempty" bson:"refId,omitempty"`
	MissingIndexNums MissingIndexNums `xml:"missingIndexNums,omitempty" bson:"missingIndexNums,omitempty"`
}

// Violations is generated from an XSD element
type Violations struct {
	Violation []ZfcsViolationType `xml:"violation,omitempty" bson:"violation,omitempty"`
}

// MissingIndexNums is generated from an XSD element
type MissingIndexNums struct {
	IndexNum []int64 `xml:"indexNum,omitempty" bson:"indexNum,omitempty"`
}

// ZfcsKladrType is generated from an XSD element
type ZfcsKladrType struct {
	KladrType string `xml:"kladrType,omitempty" bson:"kladrType,omitempty"`
	KladrCode string `xml:"kladrCode,omitempty" bson:"kladrCode,omitempty"`
	FullName  string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsTenderPlanTotalPositionKBKsType is generated from an XSD element
type ZfcsTenderPlanTotalPositionKBKsType struct {
	KBK []KBK `xml:"KBK,omitempty" bson:"KBK,omitempty"`
}

// KBK is generated from an XSD element
type KBK struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Amount string `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// ZfcsPurchasePlanPositionFinancingsType is generated from an XSD element
type ZfcsPurchasePlanPositionFinancingsType struct {
	KVR []KVR `xml:"KVR,omitempty" bson:"KVR,omitempty"`
	KBK []KBK `xml:"KBK,omitempty" bson:"KBK,omitempty"`
}

// KVR is generated from an XSD element
type KVR struct {
	Code    string                   `xml:"code,omitempty" bson:"code,omitempty"`
	Amounts ZfcsFinanceResourcesType `xml:"amounts,omitempty" bson:"amounts,omitempty"`
}

// ZfcsAddInfoInvalidType is generated from an XSD element
type ZfcsAddInfoInvalidType struct {
	SchemeVersion    string                    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID       string                    `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID               int64                     `xml:"id,omitempty" bson:"id,omitempty"`
	RegistryNum      string                    `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	PublishOrg       ZfcsOrganizationInfoType  `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	VersionNumber    string                    `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate   time.Time                 `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	FirstPublishDate time.Time                 `xml:"firstPublishDate,omitempty" bson:"firstPublishDate,omitempty"`
	Href             string                    `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm        ZfcsContractPrintFormType `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	InfoType         string                    `xml:"infoType,omitempty" bson:"infoType,omitempty"`
	InvalidityInfo   InvalidityInfo            `xml:"invalidityInfo,omitempty" bson:"invalidityInfo,omitempty"`
	Attachments      ZfcsAttachmentListType    `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ExtPrintForm     ZfcsExtPrintFormType      `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Contract         Contract                  `xml:"contract,omitempty" bson:"contract,omitempty"`
	Purchase         Purchase                  `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Contractor       Contractor                `xml:"contractor,omitempty" bson:"contractor,omitempty"`
}

// InvalidityInfo is generated from an XSD element
type InvalidityInfo struct {
	Date   time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Reason string    `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// Contract is generated from an XSD element
type Contract struct {
	RegNum   string                   `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	Number   string                   `xml:"number,omitempty" bson:"number,omitempty"`
	SignDate xsd.Date                 `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	Customer ZfcsOrganizationInfoType `xml:"customer,omitempty" bson:"customer,omitempty"`
}

// Contractor is generated from an XSD element
type Contractor struct {
	OrganizationRF           OrganizationRF           `xml:"organizationRF,omitempty" bson:"organizationRF,omitempty"`
	PersonRF                 PersonRF                 `xml:"personRF,omitempty" bson:"personRF,omitempty"`
	OrganizationForeignState OrganizationForeignState `xml:"organizationForeignState,omitempty" bson:"organizationForeignState,omitempty"`
	PersonForeignState       PersonForeignState       `xml:"personForeignState,omitempty" bson:"personForeignState,omitempty"`
}

// OrganizationRF is generated from an XSD element
type OrganizationRF struct {
	FullName string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	FirmName string `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Inn      string `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp      string `xml:"kpp,omitempty" bson:"kpp,omitempty"`
}

// PersonRF is generated from an XSD element
type PersonRF struct {
	Name ZfcsContactPersonType `xml:"name,omitempty" bson:"name,omitempty"`
	Inn  string                `xml:"inn,omitempty" bson:"inn,omitempty"`
}

// OrganizationForeignState is generated from an XSD element
type OrganizationForeignState struct {
	FullName string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	FirmName string `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Inn      string `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp      string `xml:"kpp,omitempty" bson:"kpp,omitempty"`
}

// PersonForeignState is generated from an XSD element
type PersonForeignState struct {
	Name ZfcsContactPersonType `xml:"name,omitempty" bson:"name,omitempty"`
	Inn  string                `xml:"inn,omitempty" bson:"inn,omitempty"`
}

// ZfcsCheckResultPrescriptionType is generated from an XSD element
type ZfcsCheckResultPrescriptionType struct {
	PrescriptionNumber string                 `xml:"prescriptionNumber,omitempty" bson:"prescriptionNumber,omitempty"`
	PrescriptionDate   xsd.Date               `xml:"prescriptionDate,omitempty" bson:"prescriptionDate,omitempty"`
	Attachments        ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsNsiDeviationFactFoundationType is generated from an XSD element
type ZfcsNsiDeviationFactFoundationType struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsTenderSuspensionType is generated from an XSD element
type ZfcsTenderSuspensionType struct {
	SchemeVersion   string              `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ComplaintNumber string              `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	RegDate         time.Time           `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	PublishDate     time.Time           `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Action          string              `xml:"action,omitempty" bson:"action,omitempty"`
	KO              ZfcsOrganizationRef `xml:"KO,omitempty" bson:"KO,omitempty"`
	TendersInfo     TendersInfo         `xml:"tendersInfo,omitempty" bson:"tendersInfo,omitempty"`
}

// TendersInfo is generated from an XSD element
type TendersInfo struct {
	Purchase ZfcsComplaintPurchaseType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Order    ZfcsComplaintOrderType    `xml:"order,omitempty" bson:"order,omitempty"`
}

// ZfcsTenderPlanPositionType is generated from an XSD element
type ZfcsTenderPlanPositionType struct {
	CommonInfo         CommonInfo         `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Products           Products           `xml:"products,omitempty" bson:"products,omitempty"`
	PurchaseConditions PurchaseConditions `xml:"purchaseConditions,omitempty" bson:"purchaseConditions,omitempty"`
}

// Products is generated from an XSD element
type Products struct {
	Product []Product `xml:"product,omitempty" bson:"product,omitempty"`
}

// PurchaseConditions is generated from an XSD element
type PurchaseConditions struct {
	PurchaseFinCondition   PurchaseFinCondition   `xml:"purchaseFinCondition,omitempty" bson:"purchaseFinCondition,omitempty"`
	ContractFinCondition   ContractFinCondition   `xml:"contractFinCondition,omitempty" bson:"contractFinCondition,omitempty"`
	Advance                string                 `xml:"advance,omitempty" bson:"advance,omitempty"`
	PurchaseGraph          PurchaseGraph          `xml:"purchaseGraph,omitempty" bson:"purchaseGraph,omitempty"`
	Prohibitions           string                 `xml:"prohibitions,omitempty" bson:"prohibitions,omitempty"`
	PreferensesRequirement PreferensesRequirement `xml:"preferensesRequirement,omitempty" bson:"preferensesRequirement,omitempty"`
}

// PurchaseFinCondition is generated from an XSD element
type PurchaseFinCondition struct {
	Procedure string `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	Amount    string `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// ContractFinCondition is generated from an XSD element
type ContractFinCondition struct {
	Procedure string `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	Amount    string `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// PurchaseGraph is generated from an XSD element
type PurchaseGraph struct {
	PurchasePlacingTerm     ZfcsStageType `xml:"purchasePlacingTerm,omitempty" bson:"purchasePlacingTerm,omitempty"`
	ContractExecutionTerm   ZfcsStageType `xml:"contractExecutionTerm,omitempty" bson:"contractExecutionTerm,omitempty"`
	ContractExecutionStages string        `xml:"contractExecutionStages,omitempty" bson:"contractExecutionStages,omitempty"`
	Periodicity             string        `xml:"periodicity,omitempty" bson:"periodicity,omitempty"`
}

// PreferensesRequirement is generated from an XSD element
type PreferensesRequirement struct {
	Preferenses  Preferenses  `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements Requirements `xml:"requirements,omitempty" bson:"requirements,omitempty"`
}

// ZfcsPurchaseCancelReasonType is generated from an XSD element
type ZfcsPurchaseCancelReasonType struct {
	ResponsibleDecision   ResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
	CourtDecision         CourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	DiscussionResult      DiscussionResult      `xml:"discussionResult,omitempty" bson:"discussionResult,omitempty"`
}

// ResponsibleDecision is generated from an XSD element
type ResponsibleDecision struct {
	DecisionDate time.Time `xml:"decisionDate,omitempty" bson:"decisionDate,omitempty"`
}

// AuthorityPrescription is generated from an XSD element
type AuthorityPrescription struct {
	ReestrPrescription   ReestrPrescription   `xml:"reestrPrescription,omitempty" bson:"reestrPrescription,omitempty"`
	ExternalPrescription ExternalPrescription `xml:"externalPrescription,omitempty" bson:"externalPrescription,omitempty"`
}

// ReestrPrescription is generated from an XSD element
type ReestrPrescription struct {
	CheckResultNumber  string    `xml:"checkResultNumber,omitempty" bson:"checkResultNumber,omitempty"`
	PrescriptionNumber string    `xml:"prescriptionNumber,omitempty" bson:"prescriptionNumber,omitempty"`
	Foundation         string    `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	AuthorityName      string    `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	DocDate            time.Time `xml:"docDate,omitempty" bson:"docDate,omitempty"`
}

// ExternalPrescription is generated from an XSD element
type ExternalPrescription struct {
	AuthorityName string    `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	AuthorityType string    `xml:"authorityType,omitempty" bson:"authorityType,omitempty"`
	DocName       string    `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate       time.Time `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber     string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// CourtDecision is generated from an XSD element
type CourtDecision struct {
	CourtName string    `xml:"courtName,omitempty" bson:"courtName,omitempty"`
	DocName   string    `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate   time.Time `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// DiscussionResult is generated from an XSD element
type DiscussionResult struct {
	DocName   string    `xml:"docName,omitempty" bson:"docName,omitempty"`
	DocDate   time.Time `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
}

// CommonAttributesType is generated from an XSD element
type CommonAttributesType struct {
	MaxPrice            MaxPrice            `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	LifeTime            LifeTime            `xml:"lifeTime,omitempty" bson:"lifeTime,omitempty"`
	TimeRanges          TimeRanges          `xml:"timeRanges,omitempty" bson:"timeRanges,omitempty"`
	MinTime             MinTime             `xml:"minTime,omitempty" bson:"minTime,omitempty"`
	GuaranteeTime       GuaranteeTime       `xml:"guaranteeTime,omitempty" bson:"guaranteeTime,omitempty"`
	IncludingExpences   IncludingExpences   `xml:"includingExpences,omitempty" bson:"includingExpences,omitempty"`
	MinVolume           MinVolume           `xml:"minVolume,omitempty" bson:"minVolume,omitempty"`
	GuaranteePrice      GuaranteePrice      `xml:"guaranteePrice,omitempty" bson:"guaranteePrice,omitempty"`
	MustExecuteContract MustExecuteContract `xml:"mustExecuteContract,omitempty" bson:"mustExecuteContract,omitempty"`
}

// MaxPrice is generated from an XSD element
type MaxPrice struct {
	Value float64 `xml:"value,omitempty" bson:"value,omitempty"`
}

// LifeTime is generated from an XSD element
type LifeTime struct {
	Value    float64 `xml:"value,omitempty" bson:"value,omitempty"`
	TimeUnit string  `xml:"timeUnit,omitempty" bson:"timeUnit,omitempty"`
}

// TimeRanges is generated from an XSD element
type TimeRanges struct {
	TimeRange []TimeRange `xml:"timeRange,omitempty" bson:"timeRange,omitempty"`
	TimeUnit  string      `xml:"timeUnit,omitempty" bson:"timeUnit,omitempty"`
}

// TimeRange is generated from an XSD element
type TimeRange struct {
	OrdinalNumber int64   `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	From          float64 `xml:"from,omitempty" bson:"from,omitempty"`
	To            float64 `xml:"to,omitempty" bson:"to,omitempty"`
}

// MinTime is generated from an XSD element
type MinTime struct {
	Value    float64 `xml:"value,omitempty" bson:"value,omitempty"`
	TimeUnit string  `xml:"timeUnit,omitempty" bson:"timeUnit,omitempty"`
}

// GuaranteeTime is generated from an XSD element
type GuaranteeTime struct {
	Value    float64 `xml:"value,omitempty" bson:"value,omitempty"`
	TimeUnit string  `xml:"timeUnit,omitempty" bson:"timeUnit,omitempty"`
}

// IncludingExpences is generated from an XSD element
type IncludingExpences struct {
	Value bool `xml:"value,omitempty" bson:"value,omitempty"`
}

// MinVolume is generated from an XSD element
type MinVolume struct {
	Value float64 `xml:"value,omitempty" bson:"value,omitempty"`
}

// GuaranteePrice is generated from an XSD element
type GuaranteePrice struct {
	Value float64 `xml:"value,omitempty" bson:"value,omitempty"`
}

// MustExecuteContract is generated from an XSD element
type MustExecuteContract struct {
	Value bool `xml:"value,omitempty" bson:"value,omitempty"`
}

// ZfcsUnplannedEventType is generated from an XSD element
type ZfcsUnplannedEventType struct {
	SchemeVersion      string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo         CommonInfo             `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	UnplannedEventType UnplannedEventType     `xml:"unplannedEventType,omitempty" bson:"unplannedEventType,omitempty"`
	Period             ZfcsPeriodType         `xml:"period,omitempty" bson:"period,omitempty"`
	Inspector          ZfcsOrganizationRef    `xml:"inspector,omitempty" bson:"inspector,omitempty"`
	CheckedSubject     []ZfcsEventSubjectType `xml:"checkedSubject,omitempty" bson:"checkedSubject,omitempty"`
	Base               Base                   `xml:"base,omitempty" bson:"base,omitempty"`
	CheckedObject      CheckedObject          `xml:"checkedObject,omitempty" bson:"checkedObject,omitempty"`
	PrintForm          ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments        ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// UnplannedEventType is generated from an XSD element
type UnplannedEventType struct {
	UnplannedCheck    UnplannedCheck  `xml:"unplannedCheck,omitempty" bson:"unplannedCheck,omitempty"`
	UnplannedRevision bool            `xml:"unplannedRevision,omitempty" bson:"unplannedRevision,omitempty"`
	UnplannedSurvey   UnplannedSurvey `xml:"unplannedSurvey,omitempty" bson:"unplannedSurvey,omitempty"`
}

// UnplannedCheck is generated from an XSD element
type UnplannedCheck struct {
	InspectionType string    `xml:"inspectionType,omitempty" bson:"inspectionType,omitempty"`
	EventLink      EventLink `xml:"eventLink,omitempty" bson:"eventLink,omitempty"`
}

// EventLink is generated from an XSD element
type EventLink struct {
	ParentEventPlanNumber      string `xml:"parentEventPlanNumber,omitempty" bson:"parentEventPlanNumber,omitempty"`
	ParentUnplannedEventNumber string `xml:"parentUnplannedEventNumber,omitempty" bson:"parentUnplannedEventNumber,omitempty"`
}

// UnplannedSurvey is generated from an XSD element
type UnplannedSurvey struct {
	SurveyType string    `xml:"surveyType,omitempty" bson:"surveyType,omitempty"`
	EventLink  EventLink `xml:"eventLink,omitempty" bson:"eventLink,omitempty"`
}

// Base is generated from an XSD element
type Base struct {
	ViolationInfo       ViolationInfo       `xml:"violationInfo,omitempty" bson:"violationInfo,omitempty"`
	PrescriptionControl PrescriptionControl `xml:"prescriptionControl,omitempty" bson:"prescriptionControl,omitempty"`
	BaseOther           string              `xml:"baseOther,omitempty" bson:"baseOther,omitempty"`
}

// ViolationInfo is generated from an XSD element
type ViolationInfo struct {
	InfoSource        string    `xml:"infoSource,omitempty" bson:"infoSource,omitempty"`
	InfoReceivingDate time.Time `xml:"infoReceivingDate,omitempty" bson:"infoReceivingDate,omitempty"`
}

// PrescriptionControl is generated from an XSD element
type PrescriptionControl struct {
	PrescriptionInfo PrescriptionInfo `xml:"prescriptionInfo,omitempty" bson:"prescriptionInfo,omitempty"`
	DecisionInfo     DecisionInfo     `xml:"decisionInfo,omitempty" bson:"decisionInfo,omitempty"`
}

// PrescriptionInfo is generated from an XSD element
type PrescriptionInfo struct {
	CheckResultNumber  string `xml:"checkResultNumber,omitempty" bson:"checkResultNumber,omitempty"`
	PrescriptionNumber string `xml:"prescriptionNumber,omitempty" bson:"prescriptionNumber,omitempty"`
}

// DecisionInfo is generated from an XSD element
type DecisionInfo struct {
	AuthorityType string    `xml:"authorityType,omitempty" bson:"authorityType,omitempty"`
	AuthorityName string    `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	DesDate       time.Time `xml:"desDate,omitempty" bson:"desDate,omitempty"`
	DesNumber     string    `xml:"desNumber,omitempty" bson:"desNumber,omitempty"`
}

// CheckedObject is generated from an XSD element
type CheckedObject struct {
	Info            string         `xml:"info,omitempty" bson:"info,omitempty"`
	CheckedOrder    []CheckedOrder `xml:"checkedOrder,omitempty" bson:"checkedOrder,omitempty"`
	ObjectOtherInfo string         `xml:"objectOtherInfo,omitempty" bson:"objectOtherInfo,omitempty"`
}

// CheckedOrder is generated from an XSD element
type CheckedOrder struct {
	Info           string   `xml:"info,omitempty" bson:"info,omitempty"`
	Order          Order    `xml:"order,omitempty" bson:"order,omitempty"`
	Purchase       Purchase `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	SingleCustomer bool     `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
}

// ZfcsProtocolZKType is generated from an XSD element
type ZfcsProtocolZKType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsNsiOrganizationType is generated from an XSD element
type ZfcsNsiOrganizationType struct {
	RegNumber         string                `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	ConsRegistryNum   string                `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	ShortName         string                `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	FullName          string                `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	GRBSCode          string                `xml:"GRBSCode,omitempty" bson:"GRBSCode,omitempty"`
	BIK               string                `xml:"BIK,omitempty" bson:"BIK,omitempty"`
	NomBank           string                `xml:"nomBank,omitempty" bson:"nomBank,omitempty"`
	FactualAddress    FactualAddress        `xml:"factualAddress,omitempty" bson:"factualAddress,omitempty"`
	PostalAddress     string                `xml:"postalAddress,omitempty" bson:"postalAddress,omitempty"`
	Email             string                `xml:"email,omitempty" bson:"email,omitempty"`
	Phone             string                `xml:"phone,omitempty" bson:"phone,omitempty"`
	Fax               string                `xml:"fax,omitempty" bson:"fax,omitempty"`
	ContactPerson     ZfcsContactPersonType `xml:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	Accounts          Accounts              `xml:"accounts,omitempty" bson:"accounts,omitempty"`
	Budgets           Budgets               `xml:"budgets,omitempty" bson:"budgets,omitempty"`
	HeadAgency        ZfcsOrganizationRef   `xml:"headAgency,omitempty" bson:"headAgency,omitempty"`
	OrderingAgency    ZfcsOrganizationRef   `xml:"orderingAgency,omitempty" bson:"orderingAgency,omitempty"`
	INN               string                `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP               string                `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	RegistrationDate  xsd.Date              `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	UBPCode           string                `xml:"UBPCode,omitempty" bson:"UBPCode,omitempty"`
	IKUInfo           IKUInfo               `xml:"IKUInfo,omitempty" bson:"IKUInfo,omitempty"`
	OGRN              string                `xml:"OGRN,omitempty" bson:"OGRN,omitempty"`
	OKOPF             OKOPF                 `xml:"OKOPF,omitempty" bson:"OKOPF,omitempty"`
	OKPO              string                `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	OKVED             string                `xml:"OKVED,omitempty" bson:"OKVED,omitempty"`
	OKOGU             OKOGU                 `xml:"OKOGU,omitempty" bson:"OKOGU,omitempty"`
	OrganizationRole  []string              `xml:"organizationRole,omitempty" bson:"organizationRole,omitempty"`
	OrganizationType  OrganizationType      `xml:"organizationType,omitempty" bson:"organizationType,omitempty"`
	OKTMO             ZfcsOKTMORef          `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	OKTMOPPO          ZfcsOKTMORef          `xml:"OKTMOPPO,omitempty" bson:"OKTMOPPO,omitempty"`
	SubordinationType string                `xml:"subordinationType,omitempty" bson:"subordinationType,omitempty"`
	URL               string                `xml:"url,omitempty" bson:"url,omitempty"`
	TimeZone          int64                 `xml:"timeZone,omitempty" bson:"timeZone,omitempty"`
	TimeZoneUtcOffset string                `xml:"timeZoneUtcOffset,omitempty" bson:"timeZoneUtcOffset,omitempty"`
	TimeZoneOlson     string                `xml:"timeZoneOlson,omitempty" bson:"timeZoneOlson,omitempty"`
	Actual            bool                  `xml:"actual,omitempty" bson:"actual,omitempty"`
	Register          bool                  `xml:"register,omitempty" bson:"register,omitempty"`
}

// FactualAddress is generated from an XSD element
type FactualAddress struct {
	OKATO          string         `xml:"OKATO,omitempty" bson:"OKATO,omitempty"`
	AddressLine    string         `xml:"addressLine,omitempty" bson:"addressLine,omitempty"`
	Area           ZfcsKladrType  `xml:"area,omitempty" bson:"area,omitempty"`
	Building       string         `xml:"building,omitempty" bson:"building,omitempty"`
	Country        ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
	FilledManually bool           `xml:"filledManually,omitempty" bson:"filledManually,omitempty"`
	Office         string         `xml:"office,omitempty" bson:"office,omitempty"`
	Region         ZfcsKladrType  `xml:"region,omitempty" bson:"region,omitempty"`
	Settlement     ZfcsKladrType  `xml:"settlement,omitempty" bson:"settlement,omitempty"`
	City           ZfcsKladrType  `xml:"city,omitempty" bson:"city,omitempty"`
	ShortStreet    string         `xml:"shortStreet,omitempty" bson:"shortStreet,omitempty"`
	Street         ZfcsKladrType  `xml:"street,omitempty" bson:"street,omitempty"`
	Zip            string         `xml:"zip,omitempty" bson:"zip,omitempty"`
}

// Accounts is generated from an XSD element
type Accounts struct {
	Account []ZfcsAccountType `xml:"account,omitempty" bson:"account,omitempty"`
}

// Budgets is generated from an XSD element
type Budgets struct {
	Budget []Budget `xml:"budget,omitempty" bson:"budget,omitempty"`
}

// Budget is generated from an XSD element
type Budget struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// IKUInfo is generated from an XSD element
type IKUInfo struct {
	IKU        string    `xml:"IKU,omitempty" bson:"IKU,omitempty"`
	DateStIKU  time.Time `xml:"dateStIKU,omitempty" bson:"dateStIKU,omitempty"`
	DateEndIKU time.Time `xml:"dateEndIKU,omitempty" bson:"dateEndIKU,omitempty"`
}

// OKOPF is generated from an XSD element
type OKOPF struct {
	Code     string `xml:"code,omitempty" bson:"code,omitempty"`
	FullName string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// OKOGU is generated from an XSD element
type OKOGU struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// OrganizationType is generated from an XSD element
type OrganizationType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsPublicDiscussionResultType is generated from an XSD element
type ZfcsPublicDiscussionResultType struct {
	SchemeVersion        string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	PublicDiscussionNum  string                       `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	RevisionNumber       string                       `xml:"revisionNumber,omitempty" bson:"revisionNumber,omitempty"`
	PublishDate          time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	UpdatePublishDate    time.Time                    `xml:"UpdatePublishDate,omitempty" bson:"UpdatePublishDate,omitempty"`
	Phase                string                       `xml:"phase,omitempty" bson:"phase,omitempty"`
	Customer             ZfcsPurchaseOrganizationType `xml:"customer,omitempty" bson:"customer,omitempty"`
	Stages               Stages                       `xml:"stages,omitempty" bson:"stages,omitempty"`
	Place                string                       `xml:"place,omitempty" bson:"place,omitempty"`
	Date                 time.Time                    `xml:"date,omitempty" bson:"date,omitempty"`
	Attachments          ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PlanNumber           string                       `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	PositionNumber       string                       `xml:"positionNumber,omitempty" bson:"positionNumber,omitempty"`
	PlanObjectInfo       string                       `xml:"planObjectInfo,omitempty" bson:"planObjectInfo,omitempty"`
	Year                 int64                        `xml:"year,omitempty" bson:"year,omitempty"`
	PlanContractMaxPrice string                       `xml:"planContractMaxPrice,omitempty" bson:"planContractMaxPrice,omitempty"`
	PurchaseNumber       string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	LotNumber            uint64                       `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	ObjectInfo           string                       `xml:"objectInfo,omitempty" bson:"objectInfo,omitempty"`
	ContractMaxPrice     string                       `xml:"contractMaxPrice,omitempty" bson:"contractMaxPrice,omitempty"`
}

// Stages is generated from an XSD element
type Stages struct {
	Stage Stage `xml:"stage,omitempty" bson:"stage,omitempty"`
}

// Stage is generated from an XSD element
type Stage struct {
	Number            uint64                 `xml:"number,omitempty" bson:"number,omitempty"`
	PublishDate       time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	RevisionNumber    string                 `xml:"revisionNumber,omitempty" bson:"revisionNumber,omitempty"`
	UpdatePublishDate time.Time              `xml:"updatePublishDate,omitempty" bson:"updatePublishDate,omitempty"`
	Period            Period                 `xml:"period,omitempty" bson:"period,omitempty"`
	DecisionType      ZfcsDecisionRef        `xml:"decisionType,omitempty" bson:"decisionType,omitempty"`
	BaseType          BaseType               `xml:"baseType,omitempty" bson:"baseType,omitempty"`
	Attachments       ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// Period is generated from an XSD element
type Period struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// BaseType is generated from an XSD element
type BaseType struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsDocType is generated from an XSD element
type ZfcsDocType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNotificationOKDType is generated from an XSD element
type ZfcsNotificationOKDType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKDType         `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsProtocolCancelType is generated from an XSD element
type ZfcsProtocolCancelType struct {
	SchemeVersion   string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID              int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID      string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber  string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber  string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	DocNumber       string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate         time.Time                    `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate  time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href            string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm       ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm    ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	AddInfo         string                       `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Attachments     ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	CancelReason    ZfcsProtocolCancelReasonType `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
	CancelProtocols CancelProtocols              `xml:"cancelProtocols,omitempty" bson:"cancelProtocols,omitempty"`
	CancelOrg       CancelOrg                    `xml:"cancelOrg,omitempty" bson:"cancelOrg,omitempty"`
}

// CancelProtocols is generated from an XSD element
type CancelProtocols struct {
	CancelProtocol []CancelProtocol `xml:"cancelProtocol,omitempty" bson:"cancelProtocol,omitempty"`
}

// CancelProtocol is generated from an XSD element
type CancelProtocol struct {
	ProtocolName        string    `xml:"protocolName,omitempty" bson:"protocolName,omitempty"`
	ProtocolNumber      string    `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	ProtocolPublishDate time.Time `xml:"protocolPublishDate,omitempty" bson:"protocolPublishDate,omitempty"`
}

// CancelOrg is generated from an XSD element
type CancelOrg struct {
	CancelOrg     ZfcsPurchaseOrganizationType `xml:"cancelOrg,omitempty" bson:"cancelOrg,omitempty"`
	CancelOrgRole string                       `xml:"cancelOrgRole,omitempty" bson:"cancelOrgRole,omitempty"`
}

// ZfcsPlacingWayType is generated from an XSD element
type ZfcsPlacingWayType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsStageType is generated from an XSD element
type ZfcsStageType struct {
	Month byte  `xml:"month,omitempty" bson:"month,omitempty"`
	Year  int64 `xml:"year,omitempty" bson:"year,omitempty"`
}

// ZfcsPurchaseDocumentCancelType is generated from an XSD element
type ZfcsPurchaseDocumentCancelType struct {
	SchemeVersion  string                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                   `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                  `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                  `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string                  `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time               `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time               `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string                  `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType       `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType    `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	AddInfo        string                  `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Attachments    ZfcsAttachmentListType  `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	CancelReason   ZfcsDocCancelReasonType `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
}

// ZfcsProtocolModificationReasonType is generated from an XSD element
type ZfcsProtocolModificationReasonType struct {
	ResponsibleDecision   ResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	CourtDecision         CourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
}

// ContractConditionAttributesType is generated from an XSD element
type ContractConditionAttributesType struct {
	MaxPrice            MaxPrice            `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	TimeRanges          TimeRanges          `xml:"timeRanges,omitempty" bson:"timeRanges,omitempty"`
	MinTime             MinTime             `xml:"minTime,omitempty" bson:"minTime,omitempty"`
	MinVolume           MinVolume           `xml:"minVolume,omitempty" bson:"minVolume,omitempty"`
	MustExecuteContract MustExecuteContract `xml:"mustExecuteContract,omitempty" bson:"mustExecuteContract,omitempty"`
}

// ZfcsNotificationEFType is generated from an XSD element
type ZfcsNotificationEFType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ETP                 ZfcsETPType                      `xml:"ETP,omitempty" bson:"ETP,omitempty"`
	ProcedureInfo       ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                 Lot                              `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ProcedureInfo is generated from an XSD element
type ProcedureInfo struct {
	Collecting ZfcsPurchaseProcedureCollectingType `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Scoring    Scoring                             `xml:"scoring,omitempty" bson:"scoring,omitempty"`
	Bidding    Bidding                             `xml:"bidding,omitempty" bson:"bidding,omitempty"`
}

// Scoring is generated from an XSD element
type Scoring struct {
	Date time.Time `xml:"date,omitempty" bson:"date,omitempty"`
}

// Bidding is generated from an XSD element
type Bidding struct {
	Date    time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	AddInfo string    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsApplicantType is generated from an XSD element
type ZfcsApplicantType struct {
	ApplicantType    string `xml:"applicantType,omitempty" bson:"applicantType,omitempty"`
	OrganizationName string `xml:"organizationName,omitempty" bson:"organizationName,omitempty"`
	FactualAddress   string `xml:"factualAddress,omitempty" bson:"factualAddress,omitempty"`
	PostAddress      string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	ContactEMail     string `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax       string `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
}

// ZfcsNotificationOKOUType is generated from an XSD element
type ZfcsNotificationOKOUType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKOUType        `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsLotOKType is generated from an XSD element
type ZfcsLotOKType struct {
	LotNumber              uint64                   `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	LotObjectInfo          string                   `xml:"lotObjectInfo,omitempty" bson:"lotObjectInfo,omitempty"`
	MaxPrice               string                   `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	PriceFormula           string                   `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	StandardContractNumber string                   `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	Currency               ZfcsCurrencyRef          `xml:"currency,omitempty" bson:"currency,omitempty"`
	FinanceSource          string                   `xml:"financeSource,omitempty" bson:"financeSource,omitempty"`
	QuantityUndefined      bool                     `xml:"quantityUndefined,omitempty" bson:"quantityUndefined,omitempty"`
	CustomerRequirements   CustomerRequirements     `xml:"customerRequirements,omitempty" bson:"customerRequirements,omitempty"`
	PurchaseObjects        PurchaseObjects          `xml:"purchaseObjects,omitempty" bson:"purchaseObjects,omitempty"`
	Preferenses            Preferenses              `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements           Requirements             `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	RestrictInfo           string                   `xml:"restrictInfo,omitempty" bson:"restrictInfo,omitempty"`
	RestrictForeignsInfo   string                   `xml:"restrictForeignsInfo,omitempty" bson:"restrictForeignsInfo,omitempty"`
	AddInfo                string                   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	PublicDiscussion       ZfcsPublicDiscussionType `xml:"publicDiscussion,omitempty" bson:"publicDiscussion,omitempty"`
}

// PurchaseObjects is generated from an XSD element
type PurchaseObjects struct {
	PurchaseObject []PurchaseObject `xml:"purchaseObject,omitempty" bson:"purchaseObject,omitempty"`
	TotalSum       string           `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// PurchaseObject is generated from an XSD element
type PurchaseObject struct {
	Name               string               `xml:"name,omitempty" bson:"name,omitempty"`
	OKEI               ZfcsContractOKEIType `xml:"OKEI,omitempty" bson:"OKEI,omitempty"`
	CustomerQuantities CustomerQuantities   `xml:"customerQuantities,omitempty" bson:"customerQuantities,omitempty"`
	Price              string               `xml:"price,omitempty" bson:"price,omitempty"`
	Quantity           Quantity             `xml:"quantity,omitempty" bson:"quantity,omitempty"`
	Sum                string               `xml:"sum,omitempty" bson:"sum,omitempty"`
	OKPD               ZfcsOKPDRef          `xml:"OKPD,omitempty" bson:"OKPD,omitempty"`
	OKPD2              ZfcsOKPDRef          `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
}

// CustomerQuantities is generated from an XSD element
type CustomerQuantities struct {
	CustomerQuantity []CustomerQuantity `xml:"customerQuantity,omitempty" bson:"customerQuantity,omitempty"`
}

// CustomerQuantity is generated from an XSD element
type CustomerQuantity struct {
	Customer ZfcsOrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
	Quantity float64             `xml:"quantity,omitempty" bson:"quantity,omitempty"`
}

// ZfcsPreferenseType is generated from an XSD element
type ZfcsPreferenseType struct {
	Code      int64   `xml:"code,omitempty" bson:"code,omitempty"`
	Name      string  `xml:"name,omitempty" bson:"name,omitempty"`
	PrefValue float64 `xml:"prefValue,omitempty" bson:"prefValue,omitempty"`
}

// ZfcsPlacementResultType is generated from an XSD element
type ZfcsPlacementResultType struct {
	SchemeVersion            string                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	PurchaseNumber           string                  `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                  `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	LotNumber                uint64                  `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	FoundationProtocolNumber string                  `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	VersionNumber            int64                   `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate               time.Time               `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	ProcedurelFailed         bool                    `xml:"procedurelFailed,omitempty" bson:"procedurelFailed,omitempty"`
	AbandonedReason          ZfcsAbandonedReasonType `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
	Applications             Applications            `xml:"applications,omitempty" bson:"applications,omitempty"`
	Result                   string                  `xml:"result,omitempty" bson:"result,omitempty"`
}

// HrefType is generated from an XSD element
type HrefType struct {
	Open    string `xml:"open,omitempty" bson:"open,omitempty"`
	Private string `xml:"private,omitempty" bson:"private,omitempty"`
}

// ZfcsComplaintReturnInfoType is generated from an XSD element
type ZfcsComplaintReturnInfoType struct {
	Base        string                 `xml:"base,omitempty" bson:"base,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsExtraBudgetFundsContract2015 is generated from an XSD element
type ZfcsExtraBudgetFundsContract2015 struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsLinkUser is generated from an XSD element
type ZfcsLinkUser struct {
	User        string `xml:"user,omitempty" bson:"user,omitempty"`
	BlockStatus string `xml:"blockStatus,omitempty" bson:"blockStatus,omitempty"`
}

// ZfcsPositionKOSGUsYearsType is generated from an XSD element
type ZfcsPositionKOSGUsYearsType struct {
	KOSGU []KOSGU `xml:"KOSGU,omitempty" bson:"KOSGU,omitempty"`
}

// KOSGU is generated from an XSD element
type KOSGU struct {
	Code      string      `xml:"code,omitempty" bson:"code,omitempty"`
	YearsList []YearsList `xml:"yearsList,omitempty" bson:"yearsList,omitempty"`
}

// YearsList is generated from an XSD element
type YearsList struct {
	Year       int64  `xml:"year,omitempty" bson:"year,omitempty"`
	YearAmount string `xml:"yearAmount,omitempty" bson:"yearAmount,omitempty"`
}

// ZfcsUnplannedCheckObjectType is generated from an XSD element
type ZfcsUnplannedCheckObjectType struct {
	Order    Order    `xml:"order,omitempty" bson:"order,omitempty"`
	Purchase Purchase `xml:"purchase,omitempty" bson:"purchase,omitempty"`
}

// ZfcsCustomerReportSingleContractorType is generated from an XSD element
type ZfcsCustomerReportSingleContractorType struct {
	SchemeVersion      string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate            time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate     time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber      string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg         ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments        ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID           string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	ModificationReason string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	Signer             Signer                           `xml:"signer,omitempty" bson:"signer,omitempty"`
	Customer           ZfcsOrganizationInfoWithOgrnType `xml:"customer,omitempty" bson:"customer,omitempty"`
	PurchaseNumber     string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ContractRegNum     string                           `xml:"contractRegNum,omitempty" bson:"contractRegNum,omitempty"`
	LotNumber          int64                            `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	Reason             Reason                           `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsClarificationType is generated from an XSD element
type ZfcsClarificationType struct {
	SchemeVersion  string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                 `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	RequestNumber  string                 `xml:"requestNumber,omitempty" bson:"requestNumber,omitempty"`
	DocNumber      string                 `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string                 `xml:"href,omitempty" bson:"href,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Question       string                 `xml:"question,omitempty" bson:"question,omitempty"`
	Topic          string                 `xml:"topic,omitempty" bson:"topic,omitempty"`
	Attachments    ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsNotificationCancelFailureType is generated from an XSD element
type ZfcsNotificationCancelFailureType struct {
	SchemeVersion                string                             `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                           int64                              `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                   string                             `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber               string                             `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber                    string                             `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate                      time.Time                          `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate               time.Time                          `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href                         string                             `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                    ZfcsPrintFormType                  `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm                 ZfcsExtPrintFormType               `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	NotificationCancelInfo       string                             `xml:"notificationCancelInfo,omitempty" bson:"notificationCancelInfo,omitempty"`
	PlacingWay                   ZfcsPlacingWayType                 `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	PurchaseObjectInfo           string                             `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	Lot                          Lot                                `xml:"lot,omitempty" bson:"lot,omitempty"`
	RecoveryToStage              string                             `xml:"recoveryToStage,omitempty" bson:"recoveryToStage,omitempty"`
	NotificationCancelFailureOrg NotificationCancelFailureOrg       `xml:"notificationCancelFailureOrg,omitempty" bson:"notificationCancelFailureOrg,omitempty"`
	CancelReason                 ZfcsProtocolModificationReasonType `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
	AddInfo                      string                             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Attachments                  ZfcsAttachmentListType             `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// NotificationCancelFailureOrg is generated from an XSD element
type NotificationCancelFailureOrg struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// ZfcsProtocolZKAfterProlongType is generated from an XSD element
type ZfcsProtocolZKAfterProlongType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsPublicDiscussionQuestionRef is generated from an XSD element
type ZfcsPublicDiscussionQuestionRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsOKATORef is generated from an XSD element
type ZfcsOKATORef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// GuaranteeContractType is generated from an XSD element
type GuaranteeContractType struct {
	Procedure         string `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string `xml:"bik,omitempty" bson:"bik,omitempty"`
	IsBail            bool   `xml:"isBail,omitempty" bson:"isBail,omitempty"`
}

// ZfcsNsiOKEIType is generated from an XSD element
type ZfcsNsiOKEIType struct {
	Code                string  `xml:"code,omitempty" bson:"code,omitempty"`
	FullName            string  `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	Section             Section `xml:"section,omitempty" bson:"section,omitempty"`
	Group               Group   `xml:"group,omitempty" bson:"group,omitempty"`
	LocalName           string  `xml:"localName,omitempty" bson:"localName,omitempty"`
	InternationalName   string  `xml:"internationalName,omitempty" bson:"internationalName,omitempty"`
	LocalSymbol         string  `xml:"localSymbol,omitempty" bson:"localSymbol,omitempty"`
	InternationalSymbol string  `xml:"internationalSymbol,omitempty" bson:"internationalSymbol,omitempty"`
	Actual              bool    `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Section is generated from an XSD element
type Section struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// Group is generated from an XSD element
type Group struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsUnplannedCheckSubjectPlanType is generated from an XSD element
type ZfcsUnplannedCheckSubjectPlanType struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	EP                     ZfcsOrganizationRef    `xml:"EP,omitempty" bson:"EP,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	OOSAuthority           ZfcsOrganizationRef    `xml:"OOS_authority,omitempty" bson:"OOS_authority,omitempty"`
	OOSSupport             ZfcsOrganizationRef    `xml:"OOS_support,omitempty" bson:"OOS_support,omitempty"`
	RCAuthority            ZfcsOrganizationRef    `xml:"RC_authority,omitempty" bson:"RC_authority,omitempty"`
	FCAuthority            ZfcsOrganizationRef    `xml:"FC_authority,omitempty" bson:"FC_authority,omitempty"`
	LegalEntity44          ZfcsOrganizationRef    `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         ZfcsOrganizationRef    `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
	Other                  ZfcsOrganizationRef    `xml:"other,omitempty" bson:"other,omitempty"`
}

// Commission44 is generated from an XSD element
type Commission44 struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// Commission94 is generated from an XSD element
type Commission94 struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// ContractServiceOfficer is generated from an XSD element
type ContractServiceOfficer struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// ContractService is generated from an XSD element
type ContractService struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// ZfcsFoundationProtocolInfoType is generated from an XSD element
type ZfcsFoundationProtocolInfoType struct {
	Name  string    `xml:"name,omitempty" bson:"name,omitempty"`
	Date  time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place string    `xml:"place,omitempty" bson:"place,omitempty"`
}

// ZfcsProtocolOKOU3Type is generated from an XSD element
type ZfcsProtocolOKOU3Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
	OpeningProtocol          ZfcsFoundationProtocolInfoType `xml:"openingProtocol,omitempty" bson:"openingProtocol,omitempty"`
}

// ZfcsSpecialPurchaseRef is generated from an XSD element
type ZfcsSpecialPurchaseRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ProtocolEF2Type is generated from an XSD element
type ProtocolEF2Type struct {
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modifications            ModificationType `xml:"modifications,omitempty" bson:"modifications,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsComplaintGroupType is generated from an XSD element
type ZfcsComplaintGroupType struct {
	SchemeVersion       string                      `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo          ZfcsComplaintCommonInfoType `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Indicted            []Indicted                  `xml:"indicted,omitempty" bson:"indicted,omitempty"`
	Text                string                      `xml:"text,omitempty" bson:"text,omitempty"`
	ComplaintGroupItems []ComplaintGroupItems       `xml:"complaintGroupItems,omitempty" bson:"complaintGroupItems,omitempty"`
	PrintForm           ZfcsPrintFormType           `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType        `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments         ZfcsAttachmentListType      `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReturnInfo          ReturnInfo                  `xml:"returnInfo,omitempty" bson:"returnInfo,omitempty"`
}

// Indicted is generated from an XSD element
type Indicted struct {
	EPFailure ZfcsOrganizationRef `xml:"EP_failure,omitempty" bson:"EP_failure,omitempty"`
}

// ComplaintGroupItems is generated from an XSD element
type ComplaintGroupItems struct {
	ItemComplaintNumber string                    `xml:"itemComplaintNumber,omitempty" bson:"itemComplaintNumber,omitempty"`
	RegDate             time.Time                 `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	Applicant           ZfcsApplicantType         `xml:"applicant,omitempty" bson:"applicant,omitempty"`
	ReturnInfo          ReturnInfo                `xml:"returnInfo,omitempty" bson:"returnInfo,omitempty"`
	Attachments         ZfcsAttachmentListType    `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Order               ZfcsComplaintOrderType    `xml:"order,omitempty" bson:"order,omitempty"`
	Purchase            ZfcsComplaintPurchaseType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
}

// ReturnInfo is generated from an XSD element
type ReturnInfo struct {
	Base        string                 `xml:"base,omitempty" bson:"base,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsPurchaseNotificationType is generated from an XSD element
type ZfcsPurchaseNotificationType struct {
	SchemeVersion       string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// ZfcsProtocolDeviationType is generated from an XSD element
type ZfcsProtocolDeviationType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	FoundationProtocolName   string                       `xml:"foundationProtocolName,omitempty" bson:"foundationProtocolName,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsPublicDiscussionAnswerType is generated from an XSD element
type ZfcsPublicDiscussionAnswerType struct {
	SchemeVersion       string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PublicDiscussionNum string                               `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	VersionNumber       string                               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate      time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	PublishOrg          ZfcsOrganizationInfoType             `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	CommentNumber       int64                                `xml:"commentNumber,omitempty" bson:"commentNumber,omitempty"`
	Comment             string                               `xml:"comment,omitempty" bson:"comment,omitempty"`
	AnswerNumber        int64                                `xml:"answerNumber,omitempty" bson:"answerNumber,omitempty"`
	Answer              string                               `xml:"answer,omitempty" bson:"answer,omitempty"`
	Phase               string                               `xml:"phase,omitempty" bson:"phase,omitempty"`
	Purchase            ZfcsPublicDiscussionPurchaseInfoType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Attachments         ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsBigProjectFinancingsType is generated from an XSD element
type ZfcsBigProjectFinancingsType struct {
	Sources Sources `xml:"sources,omitempty" bson:"sources,omitempty"`
	Cost    string  `xml:"cost,omitempty" bson:"cost,omitempty"`
}

// Sources is generated from an XSD element
type Sources struct {
	Federal   string `xml:"federal,omitempty" bson:"federal,omitempty"`
	Region    string `xml:"region,omitempty" bson:"region,omitempty"`
	Self      string `xml:"self,omitempty" bson:"self,omitempty"`
	Nonbudget string `xml:"nonbudget,omitempty" bson:"nonbudget,omitempty"`
}

// ZfcsNotificationModificationType is generated from an XSD element
type ZfcsNotificationModificationType struct {
	ModificationNumber int64                  `xml:"modificationNumber,omitempty" bson:"modificationNumber,omitempty"`
	Info               string                 `xml:"info,omitempty" bson:"info,omitempty"`
	AddInfo            string                 `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Reason             ZfcsPurchaseChangeType `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsProtocolOKOUSingleAppType is generated from an XSD element
type ZfcsProtocolOKOUSingleAppType struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
	OpeningProtocol          ZfcsFoundationProtocolInfoType `xml:"openingProtocol,omitempty" bson:"openingProtocol,omitempty"`
}

// ZfcsSubjectRFRef is generated from an XSD element
type ZfcsSubjectRFRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNsiContractPenaltyReasonType is generated from an XSD element
type ZfcsNsiContractPenaltyReasonType struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	PenaltyType string `xml:"penaltyType,omitempty" bson:"penaltyType,omitempty"`
	Actual      bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNotificationEFDateChangeType is generated from an XSD element
type ZfcsNotificationEFDateChangeType struct {
	SchemeVersion       string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber           string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate      time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	PurchaseResponsible PurchaseResponsible  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PurchaseObjectInfo  string               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PlacingWay          PlacingWay           `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	AuctionTime         time.Time            `xml:"auctionTime,omitempty" bson:"auctionTime,omitempty"`
	NewAuctionDate      time.Time            `xml:"newAuctionDate,omitempty" bson:"newAuctionDate,omitempty"`
	Reason              Reason               `xml:"reason,omitempty" bson:"reason,omitempty"`
	Href                string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsPurchaseProcedureScoringPlaceType is generated from an XSD element
type ZfcsPurchaseProcedureScoringPlaceType struct {
	Date    time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place   string    `xml:"place,omitempty" bson:"place,omitempty"`
	AddInfo string    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsProtocolOK2Type is generated from an XSD element
type ZfcsProtocolOK2Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsCheckResultDecisionType is generated from an XSD element
type ZfcsCheckResultDecisionType struct {
	DecisionNumber string                 `xml:"decisionNumber,omitempty" bson:"decisionNumber,omitempty"`
	DecisionDate   xsd.Date               `xml:"decisionDate,omitempty" bson:"decisionDate,omitempty"`
	Attachments    ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsApplicationAdmittedInfoType is generated from an XSD element
type ZfcsApplicationAdmittedInfoType struct {
	AppRejectedReason []ZfcsAppRejectedReasonType `xml:"appRejectedReason,omitempty" bson:"appRejectedReason,omitempty"`
	Admitted          bool                        `xml:"admitted,omitempty" bson:"admitted,omitempty"`
	Score             float64                     `xml:"score,omitempty" bson:"score,omitempty"`
	AppRating         int64                       `xml:"appRating,omitempty" bson:"appRating,omitempty"`
	ConditionsScoring ConditionsScoring           `xml:"conditionsScoring,omitempty" bson:"conditionsScoring,omitempty"`
}

// ConditionsScoring is generated from an XSD element
type ConditionsScoring struct {
	ConditionScoring []ConditionScoring `xml:"conditionScoring,omitempty" bson:"conditionScoring,omitempty"`
}

// ConditionScoring is generated from an XSD element
type ConditionScoring struct {
	CostCriterion        CostCriterion        `xml:"costCriterion,omitempty" bson:"costCriterion,omitempty"`
	QualitativeCriterion QualitativeCriterion `xml:"qualitativeCriterion,omitempty" bson:"qualitativeCriterion,omitempty"`
}

// CostCriterion is generated from an XSD element
type CostCriterion struct {
	CriterionCode    string  `xml:"criterionCode,omitempty" bson:"criterionCode,omitempty"`
	Score            float64 `xml:"score,omitempty" bson:"score,omitempty"`
	NormedScore      float64 `xml:"normedScore,omitempty" bson:"normedScore,omitempty"`
	Value            float64 `xml:"value,omitempty" bson:"value,omitempty"`
	AddInfo          string  `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Limit            float64 `xml:"limit,omitempty" bson:"limit,omitempty"`
	MeasurementOrder string  `xml:"measurementOrder,omitempty" bson:"measurementOrder,omitempty"`
	Offer            string  `xml:"offer,omitempty" bson:"offer,omitempty"`
	SpelledOffer     string  `xml:"spelledOffer,omitempty" bson:"spelledOffer,omitempty"`
	OfferInfo        string  `xml:"offerInfo,omitempty" bson:"offerInfo,omitempty"`
}

// QualitativeCriterion is generated from an XSD element
type QualitativeCriterion struct {
	CriterionCode          string             `xml:"criterionCode,omitempty" bson:"criterionCode,omitempty"`
	AddInfo                string             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Value                  float64            `xml:"value,omitempty" bson:"value,omitempty"`
	Score                  float64            `xml:"score,omitempty" bson:"score,omitempty"`
	NormedScore            float64            `xml:"normedScore,omitempty" bson:"normedScore,omitempty"`
	Limit                  float64            `xml:"limit,omitempty" bson:"limit,omitempty"`
	CriterionHaveIndicator bool               `xml:"criterionHaveIndicator,omitempty" bson:"criterionHaveIndicator,omitempty"`
	Indicator              []Indicator        `xml:"indicator,omitempty" bson:"indicator,omitempty"`
	CriterionScoring       []CriterionScoring `xml:"criterionScoring,omitempty" bson:"criterionScoring,omitempty"`
}

// Indicator is generated from an XSD element
type Indicator struct {
	IndicatorCode    int64              `xml:"indicatorCode,omitempty" bson:"indicatorCode,omitempty"`
	IndicatorScoring []IndicatorScoring `xml:"indicatorScoring,omitempty" bson:"indicatorScoring,omitempty"`
	Score            float64            `xml:"score,omitempty" bson:"score,omitempty"`
	NormedScore      float64            `xml:"normedScore,omitempty" bson:"normedScore,omitempty"`
	Name             string             `xml:"name,omitempty" bson:"name,omitempty"`
	Value            float64            `xml:"value,omitempty" bson:"value,omitempty"`
	Limit            float64            `xml:"limit,omitempty" bson:"limit,omitempty"`
	MeasurementOrder string             `xml:"measurementOrder,omitempty" bson:"measurementOrder,omitempty"`
	Offer            float64            `xml:"offer,omitempty" bson:"offer,omitempty"`
	OfferInfo        string             `xml:"offerInfo,omitempty" bson:"offerInfo,omitempty"`
}

// IndicatorScoring is generated from an XSD element
type IndicatorScoring struct {
	ProtocolCommissionMember ZfcsCommissionMemberInAppType `xml:"protocolCommissionMember,omitempty" bson:"protocolCommissionMember,omitempty"`
	Score                    float64                       `xml:"score,omitempty" bson:"score,omitempty"`
	NormedScore              float64                       `xml:"normedScore,omitempty" bson:"normedScore,omitempty"`
}

// CriterionScoring is generated from an XSD element
type CriterionScoring struct {
	ProtocolCommissionMember ZfcsCommissionMemberInAppType `xml:"protocolCommissionMember,omitempty" bson:"protocolCommissionMember,omitempty"`
	Score                    float64                       `xml:"score,omitempty" bson:"score,omitempty"`
	NormedScore              float64                       `xml:"normedScore,omitempty" bson:"normedScore,omitempty"`
}

// NotificationFeatureType is generated from an XSD element
type NotificationFeatureType struct {
	PrefValue        float64          `xml:"prefValue,omitempty" bson:"prefValue,omitempty"`
	PlacementFeature PlacementFeature `xml:"placementFeature,omitempty" bson:"placementFeature,omitempty"`
}

// PlacementFeature is generated from an XSD element
type PlacementFeature struct {
	Code int64  `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// RefusalFactFoundation is generated from an XSD element
type RefusalFactFoundation struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsBankGuaranteeOrganizationType is generated from an XSD element
type ZfcsBankGuaranteeOrganizationType struct {
	RegNum           string           `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum  string           `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName         string           `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName        string           `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress      string           `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress      string           `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN              string           `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP              string           `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	Location         string           `xml:"location,omitempty" bson:"location,omitempty"`
	LegalForm        OkopfType        `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	SubjectRF        ZfcsSubjectRFRef `xml:"subjectRF,omitempty" bson:"subjectRF,omitempty"`
	OKTMO            ZfcsOKTMORef     `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	RegistrationDate xsd.Date         `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	IKU              string           `xml:"IKU,omitempty" bson:"IKU,omitempty"`
}

// ZfcsContract2015SupplierType is generated from an XSD element
type ZfcsContract2015SupplierType struct {
	LegalEntityRF                LegalEntityRF                `xml:"legalEntityRF,omitempty" bson:"legalEntityRF,omitempty"`
	LegalEntityForeignState      LegalEntityForeignState      `xml:"legalEntityForeignState,omitempty" bson:"legalEntityForeignState,omitempty"`
	IndividualPersonRF           IndividualPersonRF           `xml:"individualPersonRF,omitempty" bson:"individualPersonRF,omitempty"`
	IndividualPersonForeignState IndividualPersonForeignState `xml:"individualPersonForeignState,omitempty" bson:"individualPersonForeignState,omitempty"`
}

// LegalEntityRF is generated from an XSD element
type LegalEntityRF struct {
	LegalForm        ZfcsOkopfRef          `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	FullName         string                `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName        string                `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	FirmName         string                `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Status           string                `xml:"status,omitempty" bson:"status,omitempty"`
	ContractPrice    string                `xml:"contractPrice,omitempty" bson:"contractPrice,omitempty"`
	OKPO             string                `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	INN              string                `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP              string                `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	RegistrationDate xsd.Date              `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	OKTMO            ZfcsOKTMORef          `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address          string                `xml:"address,omitempty" bson:"address,omitempty"`
	ContactInfo      ZfcsContactPersonType `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	ContactEMail     string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// LegalEntityForeignState is generated from an XSD element
type LegalEntityForeignState struct {
	FullName                string                  `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName               string                  `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	FirmName                string                  `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	FullNameLat             string                  `xml:"fullNameLat,omitempty" bson:"fullNameLat,omitempty"`
	TaxPayerCode            string                  `xml:"taxPayerCode,omitempty" bson:"taxPayerCode,omitempty"`
	RegisterInRFTaxBodies   RegisterInRFTaxBodies   `xml:"registerInRFTaxBodies,omitempty" bson:"registerInRFTaxBodies,omitempty"`
	PlaceOfStayInRegCountry PlaceOfStayInRegCountry `xml:"placeOfStayInRegCountry,omitempty" bson:"placeOfStayInRegCountry,omitempty"`
	PlaceOfStayInRF         PlaceOfStayInRF         `xml:"placeOfStayInRF,omitempty" bson:"placeOfStayInRF,omitempty"`
}

// RegisterInRFTaxBodies is generated from an XSD element
type RegisterInRFTaxBodies struct {
	INN              string   `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP              string   `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	RegistrationDate xsd.Date `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
}

// PlaceOfStayInRegCountry is generated from an XSD element
type PlaceOfStayInRegCountry struct {
	Country      ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
	Address      string         `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string         `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string         `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// PlaceOfStayInRF is generated from an XSD element
type PlaceOfStayInRF struct {
	OKTMO        ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address      string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
}

// IndividualPersonRF is generated from an XSD element
type IndividualPersonRF struct {
	LastName         string       `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName        string       `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName       string       `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	INN              string       `xml:"INN,omitempty" bson:"INN,omitempty"`
	IsIP             bool         `xml:"isIP,omitempty" bson:"isIP,omitempty"`
	RegistrationDate xsd.Date     `xml:"registrationDate,omitempty" bson:"registrationDate,omitempty"`
	OKTMO            ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	Address          string       `xml:"address,omitempty" bson:"address,omitempty"`
	ContactEMail     string       `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone     string       `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	IsCulture        bool         `xml:"isCulture,omitempty" bson:"isCulture,omitempty"`
}

// IndividualPersonForeignState is generated from an XSD element
type IndividualPersonForeignState struct {
	LastName                string                  `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName               string                  `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName              string                  `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	LastNameLat             string                  `xml:"lastNameLat,omitempty" bson:"lastNameLat,omitempty"`
	FirstNameLat            string                  `xml:"firstNameLat,omitempty" bson:"firstNameLat,omitempty"`
	MiddleNameLat           string                  `xml:"middleNameLat,omitempty" bson:"middleNameLat,omitempty"`
	TaxPayerCode            string                  `xml:"taxPayerCode,omitempty" bson:"taxPayerCode,omitempty"`
	RegisterInRFTaxBodies   RegisterInRFTaxBodies   `xml:"registerInRFTaxBodies,omitempty" bson:"registerInRFTaxBodies,omitempty"`
	PlaceOfStayInRegCountry PlaceOfStayInRegCountry `xml:"placeOfStayInRegCountry,omitempty" bson:"placeOfStayInRegCountry,omitempty"`
	PlaceOfStayInRF         PlaceOfStayInRF         `xml:"placeOfStayInRF,omitempty" bson:"placeOfStayInRF,omitempty"`
	IsCulture               bool                    `xml:"isCulture,omitempty" bson:"isCulture,omitempty"`
}

// ZfcsNotificationISMType is generated from an XSD element
type ZfcsNotificationISMType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ProcedureInfo       ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                Lots                             `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsPurchaseProcedureSelectingType is generated from an XSD element
type ZfcsPurchaseProcedureSelectingType struct {
	Date  time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place string    `xml:"place,omitempty" bson:"place,omitempty"`
}

// ZfcsIndicatorType is generated from an XSD element
type ZfcsIndicatorType struct {
	Code             int64   `xml:"code,omitempty" bson:"code,omitempty"`
	Name             string  `xml:"name,omitempty" bson:"name,omitempty"`
	Value            float64 `xml:"value,omitempty" bson:"value,omitempty"`
	Limit            float64 `xml:"limit,omitempty" bson:"limit,omitempty"`
	MeasurementOrder string  `xml:"measurementOrder,omitempty" bson:"measurementOrder,omitempty"`
}

// EvaluationResult is generated from an XSD element
type EvaluationResult struct {
	CommissionMemberID int64   `xml:"commissionMemberId,omitempty" bson:"commissionMemberId,omitempty"`
	EvalResult         float64 `xml:"evalResult,omitempty" bson:"evalResult,omitempty"`
}

// ZfcsBaseRef is generated from an XSD element
type ZfcsBaseRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsAccountType is generated from an XSD element
type ZfcsAccountType struct {
	BankAddress     string `xml:"bankAddress,omitempty" bson:"bankAddress,omitempty"`
	BankName        string `xml:"bankName,omitempty" bson:"bankName,omitempty"`
	Bik             string `xml:"bik,omitempty" bson:"bik,omitempty"`
	CorrAccount     string `xml:"corrAccount,omitempty" bson:"corrAccount,omitempty"`
	PaymentAccount  string `xml:"paymentAccount,omitempty" bson:"paymentAccount,omitempty"`
	PersonalAccount string `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
}

// ZfcsBankGuaranteeRefusalType is generated from an XSD element
type ZfcsBankGuaranteeRefusalType struct {
	SchemeVersion    string                            `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID               int64                             `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID       string                            `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNumber        string                            `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber        string                            `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	VersionNumber    int64                             `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate   time.Time                         `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Bank             ZfcsBankGuaranteeOrganizationType `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo     ZfcsBankGuaranteeParticipantType  `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee        ZfcsBankGuaranteeInfoType         `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	Href             string                            `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm        ZfcsPrintFormType                 `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm     ZfcsExtPrintFormType              `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	RefusalInfo      RefusalInfo                       `xml:"refusalInfo,omitempty" bson:"refusalInfo,omitempty"`
	Attachments      ZfcsAttachmentListType            `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationInfo string                            `xml:"modificationInfo,omitempty" bson:"modificationInfo,omitempty"`
}

// RefusalInfo is generated from an XSD element
type RefusalInfo struct {
	DocDate        time.Time      `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocNumber      string         `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocName        string         `xml:"docName,omitempty" bson:"docName,omitempty"`
	RefusalReasons RefusalReasons `xml:"refusalReasons,omitempty" bson:"refusalReasons,omitempty"`
}

// RefusalReasons is generated from an XSD element
type RefusalReasons struct {
	RefusalReason []ZfcsBankGuaranteeRefReasonType `xml:"refusalReason,omitempty" bson:"refusalReason,omitempty"`
}

// ZfcsBankGuaranteeTerminationType is generated from an XSD element
type ZfcsBankGuaranteeTerminationType struct {
	SchemeVersion            string                                                `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                                                 `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                                                `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNumber                string                                                `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber                string                                                `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	VersionNumber            int64                                                 `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate           time.Time                                             `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Bank                     ZfcsBankGuaranteeOrganizationType                     `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo             ZfcsBankGuaranteeParticipantType                      `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee                ZfcsBankGuaranteeInfoType                             `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	RegNum                   string                                                `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	BankGuaranteeTermination ZfcsContractProcedure2015BankGuaranteeTerminationType `xml:"bankGuaranteeTermination,omitempty" bson:"bankGuaranteeTermination,omitempty"`
	Href                     string                                                `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType                                     `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType                                  `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments              ZfcsAttachmentListType                                `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationInfo         string                                                `xml:"modificationInfo,omitempty" bson:"modificationInfo,omitempty"`
}

// ZfcsComplaintCancelType is generated from an XSD element
type ZfcsComplaintCancelType struct {
	SchemeVersion   string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ComplaintNumber string                 `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	IsGroupItem     bool                   `xml:"isGroupItem,omitempty" bson:"isGroupItem,omitempty"`
	RegDate         time.Time              `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	RegistrationKO  ZfcsOrganizationRef    `xml:"registrationKO,omitempty" bson:"registrationKO,omitempty"`
	Name            string                 `xml:"name,omitempty" bson:"name,omitempty"`
	RegType         string                 `xml:"regType,omitempty" bson:"regType,omitempty"`
	Text            string                 `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm       ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm    ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments     ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsPurchasePlanAddInfoType is generated from an XSD element
type ZfcsPurchasePlanAddInfoType struct {
	AddInfo    string             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Attachment ZfcsAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsRefusalFactFoundation is generated from an XSD element
type ZfcsRefusalFactFoundation struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNotificationI111Type is generated from an XSD element
type ZfcsNotificationI111Type struct {
	SchemeVersion             string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                        int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber            string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate            time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber                 string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                      string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                 ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	PurchaseObjectInfo        string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible       PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay                ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2               bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ProcedureInfo             ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                      Lots                             `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments               ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification              ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ParticularsActProcurement string                           `xml:"particularsActProcurement,omitempty" bson:"particularsActProcurement,omitempty"`
}

// ZfcsProtocolCancelReasonType is generated from an XSD element
type ZfcsProtocolCancelReasonType struct {
	CourtDecision         CourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
}

// ZfcsNsiContractPriceChangeReasonType is generated from an XSD element
type ZfcsNsiContractPriceChangeReasonType struct {
	ID            int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name          string `xml:"name,omitempty" bson:"name,omitempty"`
	SubsystemType string `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
	Actual        bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNsiPublicDiscussionDecisionsType is generated from an XSD element
type ZfcsNsiPublicDiscussionDecisionsType struct {
	ID          int64       `xml:"id,omitempty" bson:"id,omitempty"`
	Code        string      `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string      `xml:"name,omitempty" bson:"name,omitempty"`
	Foundations Foundations `xml:"foundations,omitempty" bson:"foundations,omitempty"`
	Type        string      `xml:"type,omitempty" bson:"type,omitempty"`
	Phase       string      `xml:"phase,omitempty" bson:"phase,omitempty"`
	Actual      bool        `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Foundations is generated from an XSD element
type Foundations struct {
	Foundation []Foundation `xml:"foundation,omitempty" bson:"foundation,omitempty"`
}

// ZfcsProtocolZKBIAfterProlongType is generated from an XSD element
type ZfcsProtocolZKBIAfterProlongType struct {
	SchemeVersion  string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	SignDate       time.Time                        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	DocNumber      string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href           string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolLot    ZfcsInfoProtocolZKBIType         `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	Modification   ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// DocumentList is generated from an XSD element
type DocumentList struct {
	DocumentMeta []Document `xml:"documentMeta,omitempty" bson:"documentMeta,omitempty"`
}

// ModificationType is generated from an XSD element
type ModificationType struct {
	InitiativeType        string                `xml:"initiativeType,omitempty" bson:"initiativeType,omitempty"`
	ModificationDate      time.Time             `xml:"modificationDate,omitempty" bson:"modificationDate,omitempty"`
	Info                  string                `xml:"info,omitempty" bson:"info,omitempty"`
	AuthorityType         string                `xml:"authorityType,omitempty" bson:"authorityType,omitempty"`
	AuthorityName         string                `xml:"authorityName,omitempty" bson:"authorityName,omitempty"`
	CourtName             string                `xml:"courtName,omitempty" bson:"courtName,omitempty"`
	CourtDesNumber        string                `xml:"courtDesNumber,omitempty" bson:"courtDesNumber,omitempty"`
	DesNumber             string                `xml:"desNumber,omitempty" bson:"desNumber,omitempty"`
	AdditionalInfo        string                `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
	DocumentMetas         DocumentList          `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// ZfcsPublicDiscussionLargePurchaseType is generated from an XSD element
type ZfcsPublicDiscussionLargePurchaseType struct {
	SchemeVersion       string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID          string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID                  int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	PublicDiscussionNum string                               `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	VersionNumber       string                               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate      time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	PublishOrg          ZfcsOrganizationInfoType             `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Topic               string                               `xml:"topic,omitempty" bson:"topic,omitempty"`
	Phase               string                               `xml:"phase,omitempty" bson:"phase,omitempty"`
	Purchase            ZfcsPublicDiscussionPurchaseInfoType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Hearing             Hearing                              `xml:"hearing,omitempty" bson:"hearing,omitempty"`
	Decision            Decision                             `xml:"decision,omitempty" bson:"decision,omitempty"`
	Attachments         ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// Hearing is generated from an XSD element
type Hearing struct {
	Date  time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place string    `xml:"place,omitempty" bson:"place,omitempty"`
}

// Decision is generated from an XSD element
type Decision struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ZfcsComplaintPurchaseType is generated from an XSD element
type ZfcsComplaintPurchaseType struct {
	PurchaseNumber string `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	Lots           Lots   `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// ZfcsContract2015BankGuaranteeReturnType is generated from an XSD element
type ZfcsContract2015BankGuaranteeReturnType struct {
	GuaranteeReturn []GuaranteeReturn `xml:"guaranteeReturn,omitempty" bson:"guaranteeReturn,omitempty"`
}

// GuaranteeReturn is generated from an XSD element
type GuaranteeReturn struct {
	BankGuaranteeReturn BankGuaranteeReturn `xml:"bankGuaranteeReturn,omitempty" bson:"bankGuaranteeReturn,omitempty"`
	WaiverNotice        WaiverNotice        `xml:"waiverNotice,omitempty" bson:"waiverNotice,omitempty"`
}

// BankGuaranteeReturn is generated from an XSD element
type BankGuaranteeReturn struct {
	RegNumber         string    `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber         string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	ReturnDate        xsd.Date  `xml:"returnDate,omitempty" bson:"returnDate,omitempty"`
	ReturnReason      string    `xml:"returnReason,omitempty" bson:"returnReason,omitempty"`
	ReturnPublishDate time.Time `xml:"returnPublishDate,omitempty" bson:"returnPublishDate,omitempty"`
}

// WaiverNotice is generated from an XSD element
type WaiverNotice struct {
	RegNumber         string    `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber         string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	NoticeDate        xsd.Date  `xml:"noticeDate,omitempty" bson:"noticeDate,omitempty"`
	NoticeNumber      string    `xml:"noticeNumber,omitempty" bson:"noticeNumber,omitempty"`
	NoticeReason      string    `xml:"noticeReason,omitempty" bson:"noticeReason,omitempty"`
	NoticePublishDate time.Time `xml:"noticePublishDate,omitempty" bson:"noticePublishDate,omitempty"`
}

// ZfcsCustomerReportBigProjectMonitoringInvalidType is generated from an XSD element
type ZfcsCustomerReportBigProjectMonitoringInvalidType struct {
	SchemeVersion     string                                     `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                int64                                      `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID        string                                     `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate           time.Time                                  `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate    time.Time                                  `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber     string                                     `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg        ZfcsOrganizationInfoWithOgrnType           `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href              string                                     `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm         ZfcsPrintFormType                          `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType                       `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments       ZfcsAttachmentListType                     `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID          string                                     `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	InvalidReportInfo ZfcsInvalidReportType                      `xml:"invalidReportInfo,omitempty" bson:"invalidReportInfo,omitempty"`
	Report            ZfcsCustomerReportBigProjectMonitoringType `xml:"report,omitempty" bson:"report,omitempty"`
}

// ZfcsPurchaseProcedureOKType is generated from an XSD element
type ZfcsPurchaseProcedureOKType struct {
	Collecting ZfcsPurchaseProcedureCollectingType `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Opening    ZfcsPurchaseProcedureOpeningType    `xml:"opening,omitempty" bson:"opening,omitempty"`
	Scoring    ZfcsPurchaseProcedureScoringType    `xml:"scoring,omitempty" bson:"scoring,omitempty"`
}

// NotificationPOType is generated from an XSD element
type NotificationPOType struct {
	ID                           int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber           string                 `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber string                 `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                int64                  `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                   time.Time              `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                   PlacingWay             `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                    string                 `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                        Order                  `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                  ContactInfoType        `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                    Document               `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                DocumentList           `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                  time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate           time.Time              `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                 ModificationType       `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                         string                 `xml:"href,omitempty" bson:"href,omitempty"`
	NotificationCommission       NotificationCommission `xml:"notificationCommission,omitempty" bson:"notificationCommission,omitempty"`
	Lots                         Lots                   `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// ZfcsCheckSubjectPlanType is generated from an XSD element
type ZfcsCheckSubjectPlanType struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	EP                     ZfcsOrganizationRef    `xml:"EP,omitempty" bson:"EP,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	OOSAuthority           ZfcsOrganizationRef    `xml:"OOS_authority,omitempty" bson:"OOS_authority,omitempty"`
	RCAuthority            ZfcsOrganizationRef    `xml:"RC_authority,omitempty" bson:"RC_authority,omitempty"`
	FCAuthority            ZfcsOrganizationRef    `xml:"FC_authority,omitempty" bson:"FC_authority,omitempty"`
	LegalEntity44          ZfcsOrganizationRef    `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         ZfcsOrganizationRef    `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
}

// ZfcsNsiEvalCriterionType is generated from an XSD element
type ZfcsNsiEvalCriterionType struct {
	ID              int64           `xml:"id,omitempty" bson:"id,omitempty"`
	Name            string          `xml:"name,omitempty" bson:"name,omitempty"`
	CriterionGroups CriterionGroups `xml:"criterionGroups,omitempty" bson:"criterionGroups,omitempty"`
	Code            string          `xml:"code,omitempty" bson:"code,omitempty"`
	CriterionCode   string          `xml:"criterionCode,omitempty" bson:"criterionCode,omitempty"`
	Description     string          `xml:"description,omitempty" bson:"description,omitempty"`
	NumericalCode   int64           `xml:"numericalCode,omitempty" bson:"numericalCode,omitempty"`
	Order           int64           `xml:"order,omitempty" bson:"order,omitempty"`
	Actual          bool            `xml:"actual,omitempty" bson:"actual,omitempty"`
	NeedExpertEval  bool            `xml:"needExpertEval,omitempty" bson:"needExpertEval,omitempty"`
}

// CriterionGroups is generated from an XSD element
type CriterionGroups struct {
	CriterionGroup []CriterionGroup `xml:"criterionGroup,omitempty" bson:"criterionGroup,omitempty"`
}

// CriterionGroup is generated from an XSD element
type CriterionGroup struct {
	Code        string      `xml:"code,omitempty" bson:"code,omitempty"`
	Name        string      `xml:"name,omitempty" bson:"name,omitempty"`
	PlacingWays PlacingWays `xml:"placingWays,omitempty" bson:"placingWays,omitempty"`
}

// ZfcsTenderPlanChangeType is generated from an XSD element
type ZfcsTenderPlanChangeType struct {
	SchemeVersion          string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                     int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	PlanNumber             string                 `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	VersionNumber          int64                  `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Description            string                 `xml:"description,omitempty" bson:"description,omitempty"`
	ConfirmDate            time.Time              `xml:"confirmDate,omitempty" bson:"confirmDate,omitempty"`
	ResponsibleContactInfo ZfcsUserTenderPlanType `xml:"responsibleContactInfo,omitempty" bson:"responsibleContactInfo,omitempty"`
	ProvidedPurchases      ProvidedPurchases      `xml:"providedPurchases,omitempty" bson:"providedPurchases,omitempty"`
	ExtPrintForm           ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ProvidedPurchases is generated from an XSD element
type ProvidedPurchases struct {
	Positions      Positions                        `xml:"positions,omitempty" bson:"positions,omitempty"`
	FinalPositions ZfcsTenderPlanFinalPositionsType `xml:"finalPositions,omitempty" bson:"finalPositions,omitempty"`
}

// Positions is generated from an XSD element
type Positions struct {
	Position []ZfcsTenderPlanPositionType `xml:"position,omitempty" bson:"position,omitempty"`
}

// CountryType is generated from an XSD element
type CountryType struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsBankGuaranteeParticipantType is generated from an XSD element
type ZfcsBankGuaranteeParticipantType struct {
	LegalEntityRF                LegalEntityRF                `xml:"legalEntityRF,omitempty" bson:"legalEntityRF,omitempty"`
	LegalEntityForeignState      LegalEntityForeignState      `xml:"legalEntityForeignState,omitempty" bson:"legalEntityForeignState,omitempty"`
	IndividualPersonRF           IndividualPersonRF           `xml:"individualPersonRF,omitempty" bson:"individualPersonRF,omitempty"`
	IndividualPersonForeignState IndividualPersonForeignState `xml:"individualPersonForeignState,omitempty" bson:"individualPersonForeignState,omitempty"`
}

// ZfcsNotificationISOType is generated from an XSD element
type ZfcsNotificationISOType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ProcedureInfo       ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                 ZfcsLotISType                    `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ClarificationType is generated from an XSD element
type ClarificationType struct {
	ID                 int64        `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber string       `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	RequestRegNumber   string       `xml:"requestRegNumber,omitempty" bson:"requestRegNumber,omitempty"`
	RegNumber          string       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	VersionNumber      int64        `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate         time.Time    `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PublishDate        time.Time    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Href               string       `xml:"href,omitempty" bson:"href,omitempty"`
	Question           string       `xml:"question,omitempty" bson:"question,omitempty"`
	Topic              string       `xml:"topic,omitempty" bson:"topic,omitempty"`
	DocumentMetas      DocumentList `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// ZfcsCustomerReportType is generated from an XSD element
type ZfcsCustomerReportType struct {
	SchemeVersion      string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate            time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate     time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber      string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg         ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments        ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID           string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	ModificationReason string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	Signer             Signer                           `xml:"signer,omitempty" bson:"signer,omitempty"`
}

// ZfcsCommissionMemberInAppType is generated from an XSD element
type ZfcsCommissionMemberInAppType struct {
	MemberNumber uint64 `xml:"memberNumber,omitempty" bson:"memberNumber,omitempty"`
}

// ZfcsReleasePurchaseDocumentationType is generated from an XSD element
type ZfcsReleasePurchaseDocumentationType struct {
	GrantStartDate time.Time       `xml:"grantStartDate,omitempty" bson:"grantStartDate,omitempty"`
	GrantPlace     string          `xml:"grantPlace,omitempty" bson:"grantPlace,omitempty"`
	GrantOrder     string          `xml:"grantOrder,omitempty" bson:"grantOrder,omitempty"`
	Languages      string          `xml:"languages,omitempty" bson:"languages,omitempty"`
	GrantMeans     string          `xml:"grantMeans,omitempty" bson:"grantMeans,omitempty"`
	GrantEndDate   time.Time       `xml:"grantEndDate,omitempty" bson:"grantEndDate,omitempty"`
	PayCurrency    ZfcsCurrencyRef `xml:"payCurrency,omitempty" bson:"payCurrency,omitempty"`
	PayInfo        PayInfo         `xml:"payInfo,omitempty" bson:"payInfo,omitempty"`
}

// PayInfo is generated from an XSD element
type PayInfo struct {
	Amount            string          `xml:"amount,omitempty" bson:"amount,omitempty"`
	Part              float64         `xml:"part,omitempty" bson:"part,omitempty"`
	ProcedureInfo     string          `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	SettlementAccount string          `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string          `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string          `xml:"bik,omitempty" bson:"bik,omitempty"`
	PayCurrency       ZfcsCurrencyRef `xml:"payCurrency,omitempty" bson:"payCurrency,omitempty"`
}

// NotificationSZType is generated from an XSD element
type NotificationSZType struct {
	ID                           int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber           string                 `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber string                 `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                int64                  `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                   time.Time              `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                   PlacingWay             `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                    string                 `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                        Order                  `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                  ContactInfoType        `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                    Document               `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                DocumentList           `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                  time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate           time.Time              `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                 ModificationType       `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                         string                 `xml:"href,omitempty" bson:"href,omitempty"`
	NotificationCommission       NotificationCommission `xml:"notificationCommission,omitempty" bson:"notificationCommission,omitempty"`
	Lots                         Lots                   `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// ProtocolOK3Type is generated from an XSD element
type ProtocolOK3Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// Commission is generated from an XSD element
type Commission struct {
	RegNumber                 int64                     `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	Name                      string                    `xml:"name,omitempty" bson:"name,omitempty"`
	ProtocolCommissionMembers ProtocolCommissionMembers `xml:"protocolCommissionMembers,omitempty" bson:"protocolCommissionMembers,omitempty"`
}

// ProtocolCommissionMembers is generated from an XSD element
type ProtocolCommissionMembers struct {
	ProtocolCommissionMember []ProtocolCommissionMember `xml:"protocolCommissionMember,omitempty" bson:"protocolCommissionMember,omitempty"`
}

// ProtocolCommissionMember is generated from an XSD element
type ProtocolCommissionMember struct {
	ID      int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name    string `xml:"name,omitempty" bson:"name,omitempty"`
	Role    Role   `xml:"role,omitempty" bson:"role,omitempty"`
	Present bool   `xml:"present,omitempty" bson:"present,omitempty"`
}

// Role is generated from an XSD element
type Role struct {
	RoleID int64  `xml:"roleId,omitempty" bson:"roleId,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsBankGuaranteeTerminationInvalidType is generated from an XSD element
type ZfcsBankGuaranteeTerminationInvalidType struct {
	SchemeVersion                string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                           int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                   string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	TerminationDocNumber         string                       `xml:"terminationDocNumber,omitempty" bson:"terminationDocNumber,omitempty"`
	DocNumber                    string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate               time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	BankGuaranteeTerminationInfo BankGuaranteeTerminationInfo `xml:"bankGuaranteeTerminationInfo,omitempty" bson:"bankGuaranteeTerminationInfo,omitempty"`
	Href                         string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                    ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm                 ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments                  ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Reason                       string                       `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// BankGuaranteeTerminationInfo is generated from an XSD element
type BankGuaranteeTerminationInfo struct {
	Bank                     ZfcsBankGuaranteeOrganizationType `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo             ZfcsBankGuaranteeParticipantType  `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee                ZfcsBankGuaranteeInfoType         `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	BankGuaranteeTermination BankGuaranteeTermination          `xml:"bankGuaranteeTermination,omitempty" bson:"bankGuaranteeTermination,omitempty"`
}

// BankGuaranteeTermination is generated from an XSD element
type BankGuaranteeTermination struct {
	RegNumber              string    `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber              string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	TerminationDate        xsd.Date  `xml:"terminationDate,omitempty" bson:"terminationDate,omitempty"`
	TerminationReason      string    `xml:"terminationReason,omitempty" bson:"terminationReason,omitempty"`
	TerminationPublishDate time.Time `xml:"terminationPublishDate,omitempty" bson:"terminationPublishDate,omitempty"`
}

// ZfcsProtocolZKBIType is generated from an XSD element
type ZfcsProtocolZKBIType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	SignDate            time.Time                        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          PlacingWay                       `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Lot                 Lot                              `xml:"lot,omitempty" bson:"lot,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ProtocolLot         ZfcsInfoProtocolZKBIType         `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsSignIncomingReqConfirmationType is generated from an XSD element
type ZfcsSignIncomingReqConfirmationType struct {
	SchemeVersion string `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RefPacketUID  string `xml:"refPacketUid,omitempty" bson:"refPacketUid,omitempty"`
}

// ZfcsContractAttachmentListType is generated from an XSD element
type ZfcsContractAttachmentListType struct {
	Attachment []ZfcsContractAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsCustomerReportBaseType is generated from an XSD element
type ZfcsCustomerReportBaseType struct {
	SchemeVersion  string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate        time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber  string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg     ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href           string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments    ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsCriterionType is generated from an XSD element
type ZfcsCriterionType struct {
	Code             string     `xml:"code,omitempty" bson:"code,omitempty"`
	Value            float64    `xml:"value,omitempty" bson:"value,omitempty"`
	AddInfo          string     `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Indicators       Indicators `xml:"indicators,omitempty" bson:"indicators,omitempty"`
	Limit            float64    `xml:"limit,omitempty" bson:"limit,omitempty"`
	MeasurementOrder string     `xml:"measurementOrder,omitempty" bson:"measurementOrder,omitempty"`
}

// Indicators is generated from an XSD element
type Indicators struct {
	Indicator []ZfcsIndicatorType `xml:"indicator,omitempty" bson:"indicator,omitempty"`
}

// ZfcsOrganizationInfoType is generated from an XSD element
type ZfcsOrganizationInfoType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
}

// ZfcsClarificationRequestType is generated from an XSD element
type ZfcsClarificationRequestType struct {
	SchemeVersion  string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID     string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                 `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string                 `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time              `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	Topic          string                 `xml:"topic,omitempty" bson:"topic,omitempty"`
	Participant    Participant            `xml:"participant,omitempty" bson:"participant,omitempty"`
	Attachments    ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// Participant is generated from an XSD element
type Participant struct {
	Name  string `xml:"name,omitempty" bson:"name,omitempty"`
	Email string `xml:"email,omitempty" bson:"email,omitempty"`
}

// ZfcsPurchaseProtocolEFType is generated from an XSD element
type ZfcsPurchaseProtocolEFType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsContractCancelType is generated from an XSD element
type ZfcsContractCancelType struct {
	SchemeVersion        string    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RegNum               string    `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	CancelDate           time.Time `xml:"cancelDate,omitempty" bson:"cancelDate,omitempty"`
	DocumentBase         string    `xml:"documentBase,omitempty" bson:"documentBase,omitempty"`
	CurrentContractStage string    `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// ZfcsProtocolEF3Type is generated from an XSD element
type ZfcsProtocolEF3Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsNsiOKVEDType is generated from an XSD element
type ZfcsNsiOKVEDType struct {
	ID         int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	Section    string `xml:"section,omitempty" bson:"section,omitempty"`
	Subsection string `xml:"subsection,omitempty" bson:"subsection,omitempty"`
	ParentID   int64  `xml:"parentId,omitempty" bson:"parentId,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsGuaranteeType is generated from an XSD element
type ZfcsGuaranteeType struct {
	Amount            string  `xml:"amount,omitempty" bson:"amount,omitempty"`
	Part              float64 `xml:"part,omitempty" bson:"part,omitempty"`
	ProcedureInfo     string  `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	SettlementAccount string  `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string  `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string  `xml:"bik,omitempty" bson:"bik,omitempty"`
}

// ZfcsPurchaseProtocolType is generated from an XSD element
type ZfcsPurchaseProtocolType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsOKFSRef is generated from an XSD element
type ZfcsOKFSRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsPublicDiscussionFoundationRef is generated from an XSD element
type ZfcsPublicDiscussionFoundationRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsTenderPlanPositionKBK2016sType is generated from an XSD element
type ZfcsTenderPlanPositionKBK2016sType struct {
	KBK2016 []KBK2016 `xml:"KBK2016,omitempty" bson:"KBK2016,omitempty"`
}

// KBK2016 is generated from an XSD element
type KBK2016 struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Amount string `xml:"amount,omitempty" bson:"amount,omitempty"`
}

// ZfcsInvalidReportType is generated from an XSD element
type ZfcsInvalidReportType struct {
	PublishDate xsd.Date               `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Reason      string                 `xml:"reason,omitempty" bson:"reason,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsAppRejectedReasonNotIDType is generated from an XSD element
type ZfcsAppRejectedReasonNotIDType struct {
	NsiRejectReason NsiRejectReason `xml:"nsiRejectReason,omitempty" bson:"nsiRejectReason,omitempty"`
	Explanation     string          `xml:"explanation,omitempty" bson:"explanation,omitempty"`
}

// NsiRejectReason is generated from an XSD element
type NsiRejectReason struct {
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Reason string `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsOKPDRef is generated from an XSD element
type ZfcsOKPDRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// GuaranteeType is generated from an XSD element
type GuaranteeType struct {
	Procedure         string       `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string       `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string       `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string       `xml:"bik,omitempty" bson:"bik,omitempty"`
	Amount            float64      `xml:"amount,omitempty" bson:"amount,omitempty"`
	Currency          CurrencyType `xml:"currency,omitempty" bson:"currency,omitempty"`
}

// ZfcsComplaintType is generated from an XSD element
type ZfcsComplaintType struct {
	SchemeVersion string                      `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo                  `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Indicted      []ZfcsComplaintSubjectType  `xml:"indicted,omitempty" bson:"indicted,omitempty"`
	Applicant     ZfcsApplicantType           `xml:"applicant,omitempty" bson:"applicant,omitempty"`
	Object        ZfcsComplaintObjectType     `xml:"object,omitempty" bson:"object,omitempty"`
	Text          string                      `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm     ZfcsPrintFormType           `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType        `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType      `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReturnInfo    ZfcsComplaintReturnInfoType `xml:"returnInfo,omitempty" bson:"returnInfo,omitempty"`
}

// ZfcsContract2015Type is generated from an XSD element
type ZfcsContract2015Type struct {
	SchemeVersion             string                                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                        int64                                   `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                string                                  `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PublishDate               time.Time                               `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	VersionNumber             int64                                   `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Foundation                Foundation                              `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Customer                  Customer                                `xml:"customer,omitempty" bson:"customer,omitempty"`
	Placer                    Placer                                  `xml:"placer,omitempty" bson:"placer,omitempty"`
	Finances                  Finances                                `xml:"finances,omitempty" bson:"finances,omitempty"`
	ProtocolDate              xsd.Date                                `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	DocumentBase              string                                  `xml:"documentBase,omitempty" bson:"documentBase,omitempty"`
	DocumentCode              string                                  `xml:"documentCode,omitempty" bson:"documentCode,omitempty"`
	SignDate                  xsd.Date                                `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	RegNum                    string                                  `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	Number                    string                                  `xml:"number,omitempty" bson:"number,omitempty"`
	PriceInfo                 PriceInfo                               `xml:"priceInfo,omitempty" bson:"priceInfo,omitempty"`
	SubContractorsSum         SubContractorsSum                       `xml:"subContractorsSum,omitempty" bson:"subContractorsSum,omitempty"`
	ExecutionPeriod           ExecutionPeriod                         `xml:"executionPeriod,omitempty" bson:"executionPeriod,omitempty"`
	Enforcement               Enforcement                             `xml:"enforcement,omitempty" bson:"enforcement,omitempty"`
	GuaranteeReturns          ZfcsContract2015BankGuaranteeReturnType `xml:"guaranteeReturns,omitempty" bson:"guaranteeReturns,omitempty"`
	EnergyServiceContractInfo string                                  `xml:"energyServiceContractInfo,omitempty" bson:"energyServiceContractInfo,omitempty"`
	Products                  Products                                `xml:"products,omitempty" bson:"products,omitempty"`
	Suppliers                 Suppliers                               `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	Href                      string                                  `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                 ZfcsContractPrintFormType               `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm              ZfcsExtPrintFormType                    `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ScanDocuments             ZfcsContractAttachmentListType          `xml:"scanDocuments,omitempty" bson:"scanDocuments,omitempty"`
	MedicalDocuments          ZfcsContractAttachmentListType          `xml:"medicalDocuments,omitempty" bson:"medicalDocuments,omitempty"`
	Attachments               ZfcsContractAttachmentListType          `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification              Modification                            `xml:"modification,omitempty" bson:"modification,omitempty"`
	CurrentContractStage      string                                  `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
	Okpd2okved2               bool                                    `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// Placer is generated from an XSD element
type Placer struct {
	ResponsibleOrg  ZfcsOrganizationRef `xml:"responsibleOrg,omitempty" bson:"responsibleOrg,omitempty"`
	ResponsibleRole string              `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// Finances is generated from an XSD element
type Finances struct {
	BudgetFunds      BudgetFunds      `xml:"budgetFunds,omitempty" bson:"budgetFunds,omitempty"`
	ExtrabudgetFunds ExtrabudgetFunds `xml:"extrabudgetFunds,omitempty" bson:"extrabudgetFunds,omitempty"`
}

// BudgetFunds is generated from an XSD element
type BudgetFunds struct {
	Budget      ZfcsBudgetFundsContract2015 `xml:"budget,omitempty" bson:"budget,omitempty"`
	OKTMO       ZfcsOKTMORef                `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	BudgetLevel string                      `xml:"budgetLevel,omitempty" bson:"budgetLevel,omitempty"`
	Stages      []Stages                    `xml:"stages,omitempty" bson:"stages,omitempty"`
}

// ExtrabudgetFunds is generated from an XSD element
type ExtrabudgetFunds struct {
	Extrabudget ZfcsExtraBudgetFundsContract2015 `xml:"extrabudget,omitempty" bson:"extrabudget,omitempty"`
	Stages      []Stages                         `xml:"stages,omitempty" bson:"stages,omitempty"`
}

// PriceInfo is generated from an XSD element
type PriceInfo struct {
	Price        string                       `xml:"price,omitempty" bson:"price,omitempty"`
	PriceType    string                       `xml:"priceType,omitempty" bson:"priceType,omitempty"`
	PriceFormula string                       `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	PriceRUR     string                       `xml:"priceRUR,omitempty" bson:"priceRUR,omitempty"`
}

// SubContractorsSum is generated from an XSD element
type SubContractorsSum struct {
	SumInPercents float64 `xml:"sumInPercents,omitempty" bson:"sumInPercents,omitempty"`
	PriceValueRUR string  `xml:"priceValueRUR,omitempty" bson:"priceValueRUR,omitempty"`
}

// ExecutionPeriod is generated from an XSD element
type ExecutionPeriod struct {
	StartDate xsd.Date `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	Stages    []Stages `xml:"stages,omitempty" bson:"stages,omitempty"`
	EndDate   xsd.Date `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// Enforcement is generated from an XSD element
type Enforcement struct {
	BankGuarantee BankGuarantee `xml:"bankGuarantee,omitempty" bson:"bankGuarantee,omitempty"`
	CashAccount   CashAccount   `xml:"cashAccount,omitempty" bson:"cashAccount,omitempty"`
}

// BankGuarantee is generated from an XSD element
type BankGuarantee struct {
	RegNumber          string                       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber          string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Currency           ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	GuaranteeAmount    string                       `xml:"guaranteeAmount,omitempty" bson:"guaranteeAmount,omitempty"`
	CurrencyRate       ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	GuaranteeAmountRUR string                       `xml:"guaranteeAmountRUR,omitempty" bson:"guaranteeAmountRUR,omitempty"`
}

// CashAccount is generated from an XSD element
type CashAccount struct {
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// Modification is generated from an XSD element
type Modification struct {
	Attachments     ZfcsContractAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ContractChange  ContractChange                 `xml:"contractChange,omitempty" bson:"contractChange,omitempty"`
	ErrorCorrection ErrorCorrection                `xml:"errorCorrection,omitempty" bson:"errorCorrection,omitempty"`
}

// ContractChange is generated from an XSD element
type ContractChange struct {
	Reason         Reason           `xml:"reason,omitempty" bson:"reason,omitempty"`
	Document       Document         `xml:"document,omitempty" bson:"document,omitempty"`
	DamagePayments []DamagePayments `xml:"damagePayments,omitempty" bson:"damagePayments,omitempty"`
}

// DamagePayments is generated from an XSD element
type DamagePayments struct {
	Document     ZfcsContract2015DocumentInfo `xml:"document,omitempty" bson:"document,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	Amount       string                       `xml:"amount,omitempty" bson:"amount,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	AmountRUR    string                       `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// ErrorCorrection is generated from an XSD element
type ErrorCorrection struct {
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsOrganizationInfoExtendedType is generated from an XSD element
type ZfcsOrganizationInfoExtendedType struct {
	RegNum          string       `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string       `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string       `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string       `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	INN             string       `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string       `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	Adress          string       `xml:"adress,omitempty" bson:"adress,omitempty"`
	Phone           string       `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email           string       `xml:"email,omitempty" bson:"email,omitempty"`
	OKTMO           ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
	OKPO            ZfcsOKPORef  `xml:"OKPO,omitempty" bson:"OKPO,omitempty"`
	OKOPF           ZfcsOkopfRef `xml:"OKOPF,omitempty" bson:"OKOPF,omitempty"`
}

// ZfcsNotificationOKType is generated from an XSD element
type ZfcsNotificationOKType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKType          `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsProtocolEFSingleAppType is generated from an XSD element
type ZfcsProtocolEFSingleAppType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	AbandonedReason          ZfcsAbandonedReasonType      `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
}

// EvalCriterion is generated from an XSD element
type EvalCriterion struct {
	ID    int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name  string `xml:"name,omitempty" bson:"name,omitempty"`
	Order int64  `xml:"order,omitempty" bson:"order,omitempty"`
}

// ProtocolEF3Type is generated from an XSD element
type ProtocolEF3Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsProtocolOKOU1Type is generated from an XSD element
type ZfcsProtocolOKOU1Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                 `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsPurchasePlanOrganizationType is generated from an XSD element
type ZfcsPurchasePlanOrganizationType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	Phone           string `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email           string `xml:"email,omitempty" bson:"email,omitempty"`
}

// PlacementResultType is generated from an XSD element
type PlacementResultType struct {
	NotificationNumber            string       `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber                string       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	VersionNumber                 int64        `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Status                        string       `xml:"status,omitempty" bson:"status,omitempty"`
	CreateDate                    time.Time    `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	LastUpdateDate                time.Time    `xml:"lastUpdateDate,omitempty" bson:"lastUpdateDate,omitempty"`
	LotNumber                     int64        `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	OrderPlacementCancel          bool         `xml:"orderPlacementCancel,omitempty" bson:"orderPlacementCancel,omitempty"`
	ChangePossible                string       `xml:"changePossible,omitempty" bson:"changePossible,omitempty"`
	RepeatedPlacement             string       `xml:"repeatedPlacement,omitempty" bson:"repeatedPlacement,omitempty"`
	ContractWithParticipant       bool         `xml:"contractWithParticipant,omitempty" bson:"contractWithParticipant,omitempty"`
	ContractWithSingleApplication string       `xml:"contractWithSingleApplication,omitempty" bson:"contractWithSingleApplication,omitempty"`
	Applications                  Applications `xml:"applications,omitempty" bson:"applications,omitempty"`
}

// ZfcsRequestForQuotationCancelType is generated from an XSD element
type ZfcsRequestForQuotationCancelType struct {
	SchemeVersion      string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocPublishDate     time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	RegistryNum        string                           `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	VersionNumber      string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	State              string                           `xml:"state,omitempty" bson:"state,omitempty"`
	Href               string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PublishOrg         ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	RequestObjectInfo  string                           `xml:"requestObjectInfo,omitempty" bson:"requestObjectInfo,omitempty"`
	ResponsibleInfo    ResponsibleInfo                  `xml:"responsibleInfo,omitempty" bson:"responsibleInfo,omitempty"`
	ProcedureInfo      ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Products           Products                         `xml:"products,omitempty" bson:"products,omitempty"`
	Conditions         Conditions                       `xml:"conditions,omitempty" bson:"conditions,omitempty"`
	Attachments        ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationReason string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	CancelReason       string                           `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
}

// ResponsibleInfo is generated from an XSD element
type ResponsibleInfo struct {
	Place         string                `xml:"place,omitempty" bson:"place,omitempty"`
	ContactPerson ZfcsContactPersonType `xml:"contactPerson,omitempty" bson:"contactPerson,omitempty"`
	ContactEMail  string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone  string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax    string                `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AddInfo       string                `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// Conditions is generated from an XSD element
type Conditions struct {
	Main              string `xml:"main,omitempty" bson:"main,omitempty"`
	Payment           string `xml:"payment,omitempty" bson:"payment,omitempty"`
	ContractGuarantee string `xml:"contractGuarantee,omitempty" bson:"contractGuarantee,omitempty"`
	Warranty          string `xml:"warranty,omitempty" bson:"warranty,omitempty"`
	Delivery          string `xml:"delivery,omitempty" bson:"delivery,omitempty"`
	AddInfo           string `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsNsiOffBudgetType is generated from an XSD element
type ZfcsNsiOffBudgetType struct {
	Code          string `xml:"code,omitempty" bson:"code,omitempty"`
	Name          string `xml:"name,omitempty" bson:"name,omitempty"`
	SubsystemType string `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
	Actual        bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsTenderPlanTotalPositionKOSGUsType is generated from an XSD element
type ZfcsTenderPlanTotalPositionKOSGUsType struct {
	KOSGU []KOSGU `xml:"KOSGU,omitempty" bson:"KOSGU,omitempty"`
}

// ZfcsSignIncomingConfirmationType is generated from an XSD element
type ZfcsSignIncomingConfirmationType struct {
	SchemeVersion string     `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RefPacketUID  string     `xml:"refPacketUid,omitempty" bson:"refPacketUid,omitempty"`
	Success       Success    `xml:"success,omitempty" bson:"success,omitempty"`
	Violations    Violations `xml:"violations,omitempty" bson:"violations,omitempty"`
	Processing    bool       `xml:"processing,omitempty" bson:"processing,omitempty"`
}

// Success is generated from an XSD element
type Success struct {
	RegistrationInfo RegistrationInfo `xml:"registrationInfo,omitempty" bson:"registrationInfo,omitempty"`
	LoadURL          string           `xml:"loadUrl,omitempty" bson:"loadUrl,omitempty"`
	Warnings         Warnings         `xml:"warnings,omitempty" bson:"warnings,omitempty"`
}

// RegistrationInfo is generated from an XSD element
type RegistrationInfo struct {
	LoadID      string    `xml:"loadId,omitempty" bson:"loadId,omitempty"`
	DocumentUID string    `xml:"documentUid,omitempty" bson:"documentUid,omitempty"`
	RegNumber   string    `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	PublishDate time.Time `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
}

// Warnings is generated from an XSD element
type Warnings struct {
	Warning []ZfcsViolationType `xml:"warning,omitempty" bson:"warning,omitempty"`
}

// ZfcsContractPrintFormType is generated from an XSD element
type ZfcsContractPrintFormType struct {
	URL          string      `xml:"url,omitempty" bson:"url,omitempty"`
	DocRegNumber string      `xml:"docRegNumber,omitempty" bson:"docRegNumber,omitempty"`
	Signature    []Signature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// Signature is generated from an XSD element
type Signature struct {
	Type      string `xml:"type,attr" bson:"type,omitempty"`
	Signature []byte `xml:",chardata" bson:",omitempty"`
}

// ZfcsCustomerReportSingleContractorInvalidType is generated from an XSD element
type ZfcsCustomerReportSingleContractorInvalidType struct {
	SchemeVersion     string                                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                int64                                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID        string                                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate           time.Time                              `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate    time.Time                              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber     string                                 `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg        ZfcsOrganizationInfoWithOgrnType       `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href              string                                 `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm         ZfcsPrintFormType                      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType                   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments       ZfcsAttachmentListType                 `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID          string                                 `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	InvalidReportInfo ZfcsInvalidReportType                  `xml:"invalidReportInfo,omitempty" bson:"invalidReportInfo,omitempty"`
	Report            ZfcsCustomerReportSingleContractorType `xml:"report,omitempty" bson:"report,omitempty"`
}

// ZfcsPurchaseProlongationZKType is generated from an XSD element
type ZfcsPurchaseProlongationZKType struct {
	SchemeVersion              string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                         int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID                 string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber             string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber                  string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate                    time.Time            `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate             time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href                       string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                  ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm               ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	DocType                    ZfcsDocType          `xml:"docType,omitempty" bson:"docType,omitempty"`
	CollectingEndDate          time.Time            `xml:"collectingEndDate,omitempty" bson:"collectingEndDate,omitempty"`
	CollectingProlongationDate time.Time            `xml:"collectingProlongationDate,omitempty" bson:"collectingProlongationDate,omitempty"`
	OpeningDate                time.Time            `xml:"openingDate,omitempty" bson:"openingDate,omitempty"`
	OpeningProlongationDate    time.Time            `xml:"openingProlongationDate,omitempty" bson:"openingProlongationDate,omitempty"`
}

// ZfcsTenderPlanType is generated from an XSD element
type ZfcsTenderPlanType struct {
	SchemeVersion          string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo             ZfcsTenderPlanCommonInfoType `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CustomerInfo           CustomerInfo                 `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	ResponsibleContactInfo ResponsibleContactInfo       `xml:"responsibleContactInfo,omitempty" bson:"responsibleContactInfo,omitempty"`
	ProvidedPurchases      ProvidedPurchases            `xml:"providedPurchases,omitempty" bson:"providedPurchases,omitempty"`
	PrintForm              ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm           ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// CustomerInfo is generated from an XSD element
type CustomerInfo struct {
	Customer ZfcsPurchaseOrganizationType `xml:"customer,omitempty" bson:"customer,omitempty"`
	OKTMO    ZfcsOKTMORef                 `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
}

// ResponsibleContactInfo is generated from an XSD element
type ResponsibleContactInfo struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	Phone      string `xml:"phone,omitempty" bson:"phone,omitempty"`
	Fax        string `xml:"fax,omitempty" bson:"fax,omitempty"`
	Email      string `xml:"email,omitempty" bson:"email,omitempty"`
}

// ZfcsComplaintProjectSubjectType is generated from an XSD element
type ZfcsComplaintProjectSubjectType struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	LegalEntity44          LegalEntity44          `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         LegalEntity307         `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
}

// LegalEntity44 is generated from an XSD element
type LegalEntity44 struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// LegalEntity307 is generated from an XSD element
type LegalEntity307 struct {
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Info         string              `xml:"info,omitempty" bson:"info,omitempty"`
}

// ZfcsContractCancel2015Type is generated from an XSD element
type ZfcsContractCancel2015Type struct {
	SchemeVersion        string    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RegNum               string    `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	CancelDate           time.Time `xml:"cancelDate,omitempty" bson:"cancelDate,omitempty"`
	DocumentBase         string    `xml:"documentBase,omitempty" bson:"documentBase,omitempty"`
	CurrentContractStage string    `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// ZfcsNotificationCancelType is generated from an XSD element
type ZfcsNotificationCancelType struct {
	SchemeVersion  string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time                    `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	CancelReason   ZfcsPurchaseCancelReasonType `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
	AddInfo        string                       `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsTenderPlanPositionKBKsType is generated from an XSD element
type ZfcsTenderPlanPositionKBKsType struct {
	KBK []KBK `xml:"KBK,omitempty" bson:"KBK,omitempty"`
}

// ZfcsComplaintSubjectType is generated from an XSD element
type ZfcsComplaintSubjectType struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	EP                     ZfcsOrganizationRef    `xml:"EP,omitempty" bson:"EP,omitempty"`
	EPRefuse               ZfcsOrganizationRef    `xml:"EP_refuse,omitempty" bson:"EP_refuse,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	LegalEntity44          LegalEntity44          `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         LegalEntity307         `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
}

// ZfcsEventResultPrescriptionType is generated from an XSD element
type ZfcsEventResultPrescriptionType struct {
	PrescriptionNumber string                 `xml:"prescriptionNumber,omitempty" bson:"prescriptionNumber,omitempty"`
	PrescriptionDate   xsd.Date               `xml:"prescriptionDate,omitempty" bson:"prescriptionDate,omitempty"`
	Attachments        ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsLotZKType is generated from an XSD element
type ZfcsLotZKType struct {
	MaxPrice               string                   `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	MaxPriceInfo           string                   `xml:"maxPriceInfo,omitempty" bson:"maxPriceInfo,omitempty"`
	PriceFormula           string                   `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	StandardContractNumber string                   `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	Currency               ZfcsCurrencyRef          `xml:"currency,omitempty" bson:"currency,omitempty"`
	FinanceSource          string                   `xml:"financeSource,omitempty" bson:"financeSource,omitempty"`
	QuantityUndefined      bool                     `xml:"quantityUndefined,omitempty" bson:"quantityUndefined,omitempty"`
	CustomerRequirements   CustomerRequirements     `xml:"customerRequirements,omitempty" bson:"customerRequirements,omitempty"`
	PurchaseObjects        PurchaseObjects          `xml:"purchaseObjects,omitempty" bson:"purchaseObjects,omitempty"`
	Preferenses            Preferenses              `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements           Requirements             `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	RestrictInfo           string                   `xml:"restrictInfo,omitempty" bson:"restrictInfo,omitempty"`
	AddInfo                string                   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	PublicDiscussion       ZfcsPublicDiscussionType `xml:"publicDiscussion,omitempty" bson:"publicDiscussion,omitempty"`
}

// ZfcsRegulationRulesInvalidType is generated from an XSD element
type ZfcsRegulationRulesInvalidType struct {
	DocPublishDate     time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	ID                 int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	RegistryNum        string                 `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	RevisionNumber     string                 `xml:"revisionNumber,omitempty" bson:"revisionNumber,omitempty"`
	PublishOrg         PublishOrg             `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Type               string                 `xml:"type,omitempty" bson:"type,omitempty"`
	State              string                 `xml:"state,omitempty" bson:"state,omitempty"`
	TermsControl       bool                   `xml:"termsControl,omitempty" bson:"termsControl,omitempty"`
	ApprovedFrom       ApprovedFrom           `xml:"approvedFrom,omitempty" bson:"approvedFrom,omitempty"`
	ApproveFor         ApproveFor             `xml:"approveFor,omitempty" bson:"approveFor,omitempty"`
	BaseDocument       BaseDocument           `xml:"baseDocument,omitempty" bson:"baseDocument,omitempty"`
	Regions            Regions                `xml:"regions,omitempty" bson:"regions,omitempty"`
	AddInfo            string                 `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Documents          Documents              `xml:"documents,omitempty" bson:"documents,omitempty"`
	PrintFormDocuments PrintFormDocuments     `xml:"printFormDocuments,omitempty" bson:"printFormDocuments,omitempty"`
	Discussion         Discussion             `xml:"discussion,omitempty" bson:"discussion,omitempty"`
	InvalidityInfo     InvalidityInfo         `xml:"invalidityInfo,omitempty" bson:"invalidityInfo,omitempty"`
	Attachments        ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// PublishOrg is generated from an XSD element
type PublishOrg struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	Ogrn            string `xml:"ogrn,omitempty" bson:"ogrn,omitempty"`
	Address         string `xml:"address,omitempty" bson:"address,omitempty"`
	Phone           string `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email           string `xml:"email,omitempty" bson:"email,omitempty"`
}

// ApprovedFrom is generated from an XSD element
type ApprovedFrom struct {
	Government bool       `xml:"government,omitempty" bson:"government,omitempty"`
	PublishOrg PublishOrg `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
}

// ApproveFor is generated from an XSD element
type ApproveFor struct {
	Central     bool `xml:"central,omitempty" bson:"central,omitempty"`
	Territorial bool `xml:"territorial,omitempty" bson:"territorial,omitempty"`
	Treasury    bool `xml:"treasury,omitempty" bson:"treasury,omitempty"`
	Budgetary   bool `xml:"budgetary,omitempty" bson:"budgetary,omitempty"`
}

// BaseDocument is generated from an XSD element
type BaseDocument struct {
	Name   string    `xml:"name,omitempty" bson:"name,omitempty"`
	Number string    `xml:"number,omitempty" bson:"number,omitempty"`
	Date   time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Type   string    `xml:"type,omitempty" bson:"type,omitempty"`
}

// Regions is generated from an XSD element
type Regions struct {
	Region []Region `xml:"region,omitempty" bson:"region,omitempty"`
}

// Region is generated from an XSD element
type Region struct {
	Code           string         `xml:"code,omitempty" bson:"code,omitempty"`
	Name           string         `xml:"name,omitempty" bson:"name,omitempty"`
	Municipalities Municipalities `xml:"municipalities,omitempty" bson:"municipalities,omitempty"`
}

// Municipalities is generated from an XSD element
type Municipalities struct {
	Municipality []ZfcsOKTMORef `xml:"municipality,omitempty" bson:"municipality,omitempty"`
}

// PrintFormDocuments is generated from an XSD element
type PrintFormDocuments struct {
	PrintFormDocument []PrintFormDocument `xml:"printFormDocument,omitempty" bson:"printFormDocument,omitempty"`
}

// PrintFormDocument is generated from an XSD element
type PrintFormDocument struct {
	Type         string    `xml:"type,omitempty" bson:"type,omitempty"`
	PlacingDate  time.Time `xml:"placingDate,omitempty" bson:"placingDate,omitempty"`
	ChangingDate time.Time `xml:"changingDate,omitempty" bson:"changingDate,omitempty"`
}

// Discussion is generated from an XSD element
type Discussion struct {
	Term        Term   `xml:"term,omitempty" bson:"term,omitempty"`
	PostAddress string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	EMail       string `xml:"eMail,omitempty" bson:"eMail,omitempty"`
	Decision    string `xml:"decision,omitempty" bson:"decision,omitempty"`
	AddInfo     string `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// Term is generated from an XSD element
type Term struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
}

// ZfcsPublicDiscussionDecisionRef is generated from an XSD element
type ZfcsPublicDiscussionDecisionRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// NotificationType is generated from an XSD element
type NotificationType struct {
	ID                           int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber           string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber string           `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                   time.Time        `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                   PlacingWay       `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                    string           `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                        Order            `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                  ContactInfoType  `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                    Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                  time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate           time.Time        `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                 ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                         string           `xml:"href,omitempty" bson:"href,omitempty"`
}

// ZfcsPeriodType is generated from an XSD element
type ZfcsPeriodType struct {
	Start xsd.Date `xml:"start,omitempty" bson:"start,omitempty"`
	End   xsd.Date `xml:"end,omitempty" bson:"end,omitempty"`
}

// ZfcsCurrencyRateContract2015 is generated from an XSD element
type ZfcsCurrencyRateContract2015 struct {
	Rate    float64 `xml:"rate,omitempty" bson:"rate,omitempty"`
	Raiting int64   `xml:"raiting,omitempty" bson:"raiting,omitempty"`
}

// ZfcsTimeEFType is generated from an XSD element
type ZfcsTimeEFType struct {
	SchemeVersion                  string    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	PurchaseNumber                 string    `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber                      string    `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	AuctionTime                    time.Time `xml:"auctionTime,omitempty" bson:"auctionTime,omitempty"`
	NotificationModificationNumber int64     `xml:"notificationModificationNumber,omitempty" bson:"notificationModificationNumber,omitempty"`
}

// ZfcsNsiOKPD2Type is generated from an XSD element
type ZfcsNsiOKPD2Type struct {
	ID         int64  `xml:"id,omitempty" bson:"id,omitempty"`
	ParentID   int64  `xml:"parentId,omitempty" bson:"parentId,omitempty"`
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	ParentCode string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Comment    string `xml:"comment,omitempty" bson:"comment,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsSketchPlanType is generated from an XSD element
type ZfcsSketchPlanType struct {
	SchemeVersion string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo                   `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CustomerInfo  ZfcsQuickRefOrganizationType `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	Attachments   ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PrintForm     ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsTenderPlanContextType is generated from an XSD element
type ZfcsTenderPlanContextType struct {
	AmountKBKs     ZfcsTenderPlanTotalPositionKBKsType     `xml:"amountKBKs,omitempty" bson:"amountKBKs,omitempty"`
	AmountKBKs2016 ZfcsTenderPlanTotalPositionKBK2016sType `xml:"amountKBKs2016,omitempty" bson:"amountKBKs2016,omitempty"`
	AmountKOSGUs   ZfcsTenderPlanTotalPositionKOSGUsType   `xml:"amountKOSGUs,omitempty" bson:"amountKOSGUs,omitempty"`
	AmountKVRs     ZfcsTenderPlanTotalPositionKVRsType     `xml:"amountKVRs,omitempty" bson:"amountKVRs,omitempty"`
}

// TendePlanInfoType is generated from an XSD element
type TendePlanInfoType struct {
	PlanNumber         string `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	PlanPositionNumber string `xml:"planPositionNumber,omitempty" bson:"planPositionNumber,omitempty"`
}

// ZfcsContract2015DocDictRef is generated from an XSD element
type ZfcsContract2015DocDictRef struct {
	Code         string   `xml:"code,omitempty" bson:"code,omitempty"`
	Name         string   `xml:"name,omitempty" bson:"name,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
}

// ZfcsNsiContractSingleCustomerReasonType is generated from an XSD element
type ZfcsNsiContractSingleCustomerReasonType struct {
	ID            int64     `xml:"id,omitempty" bson:"id,omitempty"`
	Code          string    `xml:"code,omitempty" bson:"code,omitempty"`
	Name          string    `xml:"name,omitempty" bson:"name,omitempty"`
	PointLaw      string    `xml:"pointLaw,omitempty" bson:"pointLaw,omitempty"`
	SubsystemType string    `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
	Actual        bool      `xml:"actual,omitempty" bson:"actual,omitempty"`
	Documents     Documents `xml:"documents,omitempty" bson:"documents,omitempty"`
}

// ZfcsBankGuaranteeReturnType is generated from an XSD element
type ZfcsBankGuaranteeReturnType struct {
	SchemeVersion    string                                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID               int64                                   `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID       string                                  `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNumber        string                                  `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber        string                                  `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	VersionNumber    int64                                   `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate   time.Time                               `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Bank             ZfcsBankGuaranteeOrganizationType       `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo     ZfcsBankGuaranteeParticipantType        `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee        ZfcsBankGuaranteeInfoType               `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	RegNum           string                                  `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	GuaranteeReturns ZfcsContract2015BankGuaranteeReturnType `xml:"guaranteeReturns,omitempty" bson:"guaranteeReturns,omitempty"`
	Href             string                                  `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm        ZfcsPrintFormType                       `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm     ZfcsExtPrintFormType                    `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments      ZfcsAttachmentListType                  `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationInfo string                                  `xml:"modificationInfo,omitempty" bson:"modificationInfo,omitempty"`
}

// ZfcsModifyTerminateContractType is generated from an XSD element
type ZfcsModifyTerminateContractType struct {
	EventDate xsd.Date `xml:"eventDate,omitempty" bson:"eventDate,omitempty"`
	DocInfo   string   `xml:"docInfo,omitempty" bson:"docInfo,omitempty"`
}

// ZfcsNotificationEPType is generated from an XSD element
type ZfcsNotificationEPType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Lot                 Lot                              `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// ZfcsRefusalFact is generated from an XSD element
type ZfcsRefusalFact struct {
	VoucherEntry string                    `xml:"voucherEntry,omitempty" bson:"voucherEntry,omitempty"`
	Explanation  string                    `xml:"explanation,omitempty" bson:"explanation,omitempty"`
	Foundation   ZfcsRefusalFactFoundation `xml:"foundation,omitempty" bson:"foundation,omitempty"`
}

// ZfcsPositionKBKsYearsType is generated from an XSD element
type ZfcsPositionKBKsYearsType struct {
	KBK []KBK `xml:"KBK,omitempty" bson:"KBK,omitempty"`
}

// ZfcsProtocolOKD4Type is generated from an XSD element
type ZfcsProtocolOKD4Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsTenderPlanUnstructuredType is generated from an XSD element
type ZfcsTenderPlanUnstructuredType struct {
	SchemeVersion          string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo             ZfcsTenderPlanCommonInfoType `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	PublicDiscussionNum    []string                     `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	CustomerInfo           CustomerInfo                 `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	ResponsibleContactInfo ResponsibleContactInfo       `xml:"responsibleContactInfo,omitempty" bson:"responsibleContactInfo,omitempty"`
	ProvidedNotPurchases   bool                         `xml:"providedNotPurchases,omitempty" bson:"providedNotPurchases,omitempty"`
	Attachments            ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PrintForm              ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
}

// OkopfType is generated from an XSD element
type OkopfType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ProtocolPO1Type is generated from an XSD element
type ProtocolPO1Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// TimeEFType is generated from an XSD element
type TimeEFType struct {
	NotificationNumber string    `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	AuctionTime        time.Time `xml:"auctionTime,omitempty" bson:"auctionTime,omitempty"`
}

// ZfcsProtocolZPExtractType is generated from an XSD element
type ZfcsProtocolZPExtractType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	CommissionName           string                       `xml:"commissionName,omitempty" bson:"commissionName,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ProtocolZK5Type is generated from an XSD element
type ProtocolZK5Type struct {
	OrderName                string           `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                    Order            `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo              ContactInfoType  `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	Currency                 CurrencyType     `xml:"currency,omitempty" bson:"currency,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
}

// ZfcsCustomerReportBigProjectMonitoringType is generated from an XSD element
type ZfcsCustomerReportBigProjectMonitoringType struct {
	SchemeVersion      string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate            time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate     time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber      string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg         ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments        ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID           string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	ModificationReason string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	Signer             Signer                           `xml:"signer,omitempty" bson:"signer,omitempty"`
	NeedInContract     bool                             `xml:"needInContract,omitempty" bson:"needInContract,omitempty"`
	Customer           ZfcsOrganizationInfoWithOgrnType `xml:"customer,omitempty" bson:"customer,omitempty"`
	ContractsInfo      ContractsInfo                    `xml:"contractsInfo,omitempty" bson:"contractsInfo,omitempty"`
	Constructor        Constructor                      `xml:"constructor,omitempty" bson:"constructor,omitempty"`
	ProjectInfo        ProjectInfo                      `xml:"projectInfo,omitempty" bson:"projectInfo,omitempty"`
	Contractors        Contractors                      `xml:"contractors,omitempty" bson:"contractors,omitempty"`
	Participants       Participants                     `xml:"participants,omitempty" bson:"participants,omitempty"`
	Documents          Documents                        `xml:"documents,omitempty" bson:"documents,omitempty"`
	Cost               Cost                             `xml:"cost,omitempty" bson:"cost,omitempty"`
	Financings         Financings                       `xml:"financings,omitempty" bson:"financings,omitempty"`
	Tenders            Tenders                          `xml:"tenders,omitempty" bson:"tenders,omitempty"`
	Realization        Realization                      `xml:"realization,omitempty" bson:"realization,omitempty"`
	Changes            Changes                          `xml:"changes,omitempty" bson:"changes,omitempty"`
}

// ContractsInfo is generated from an XSD element
type ContractsInfo struct {
	ContractInfo []ContractInfo `xml:"contractInfo,omitempty" bson:"contractInfo,omitempty"`
}

// Constructor is generated from an XSD element
type Constructor struct {
	FullName  string       `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName string       `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	OKOPF     ZfcsOkopfRef `xml:"OKOPF,omitempty" bson:"OKOPF,omitempty"`
	Head      string       `xml:"head,omitempty" bson:"head,omitempty"`
	Position  string       `xml:"position,omitempty" bson:"position,omitempty"`
}

// ProjectInfo is generated from an XSD element
type ProjectInfo struct {
	Name           string     `xml:"name,omitempty" bson:"name,omitempty"`
	Purpose        string     `xml:"purpose,omitempty" bson:"purpose,omitempty"`
	Period         Period     `xml:"period,omitempty" bson:"period,omitempty"`
	Direction      []string   `xml:"direction,omitempty" bson:"direction,omitempty"`
	OtherDirection string     `xml:"otherDirection,omitempty" bson:"otherDirection,omitempty"`
	Mechanism      string     `xml:"mechanism,omitempty" bson:"mechanism,omitempty"`
	OtherMechanism string     `xml:"otherMechanism,omitempty" bson:"otherMechanism,omitempty"`
	Grbs           string     `xml:"grbs,omitempty" bson:"grbs,omitempty"`
	Indicators     Indicators `xml:"indicators,omitempty" bson:"indicators,omitempty"`
}

// Contractors is generated from an XSD element
type Contractors struct {
	Contractor []ZfcsBigProjectMemberType `xml:"contractor,omitempty" bson:"contractor,omitempty"`
}

// Participants is generated from an XSD element
type Participants struct {
	Participant []ZfcsBigProjectMemberType `xml:"participant,omitempty" bson:"participant,omitempty"`
}

// Cost is generated from an XSD element
type Cost struct {
	Act          Act          `xml:"act,omitempty" bson:"act,omitempty"`
	Authenticity Authenticity `xml:"authenticity,omitempty" bson:"authenticity,omitempty"`
	Contract     Contract     `xml:"contract,omitempty" bson:"contract,omitempty"`
	Economy      string       `xml:"economy,omitempty" bson:"economy,omitempty"`
}

// Act is generated from an XSD element
type Act struct {
	PriceYear []ZfcsBigProjectCostType `xml:"priceYear,omitempty" bson:"priceYear,omitempty"`
	Vat       bool                     `xml:"vat,omitempty" bson:"vat,omitempty"`
}

// Authenticity is generated from an XSD element
type Authenticity struct {
	Cost      string                   `xml:"cost,omitempty" bson:"cost,omitempty"`
	Year      int64                    `xml:"year,omitempty" bson:"year,omitempty"`
	PriceYear []ZfcsBigProjectCostType `xml:"priceYear,omitempty" bson:"priceYear,omitempty"`
	Vat       bool                     `xml:"vat,omitempty" bson:"vat,omitempty"`
}

// Financings is generated from an XSD element
type Financings struct {
	Years  ZfcsBigProjectFinancingYearsType `xml:"years,omitempty" bson:"years,omitempty"`
	Stages Stages                           `xml:"stages,omitempty" bson:"stages,omitempty"`
	Site   string                           `xml:"site,omitempty" bson:"site,omitempty"`
}

// Tenders is generated from an XSD element
type Tenders struct {
	Tender []Tender `xml:"tender,omitempty" bson:"tender,omitempty"`
	Total  Total    `xml:"total,omitempty" bson:"total,omitempty"`
}

// Tender is generated from an XSD element
type Tender struct {
	PurchaseNumber string          `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	LotNumber      int64           `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	PlacingWay     string          `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Name           string          `xml:"name,omitempty" bson:"name,omitempty"`
	Currency       ZfcsCurrencyRef `xml:"currency,omitempty" bson:"currency,omitempty"`
	StartCost      string          `xml:"startCost,omitempty" bson:"startCost,omitempty"`
	StartDate      time.Time       `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate        time.Time       `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	Participants   Participants    `xml:"participants,omitempty" bson:"participants,omitempty"`
	Winners        Winners         `xml:"winners,omitempty" bson:"winners,omitempty"`
	FinalCost      string          `xml:"finalCost,omitempty" bson:"finalCost,omitempty"`
	Salary         string          `xml:"salary,omitempty" bson:"salary,omitempty"`
	Funds          string          `xml:"funds,omitempty" bson:"funds,omitempty"`
	AddInfo        string          `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// Winners is generated from an XSD element
type Winners struct {
	Winner []Winner `xml:"winner,omitempty" bson:"winner,omitempty"`
}

// Winner is generated from an XSD element
type Winner struct {
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	StartCost string `xml:"startCost,omitempty" bson:"startCost,omitempty"`
}

// Total is generated from an XSD element
type Total struct {
	FinalCost string `xml:"finalCost,omitempty" bson:"finalCost,omitempty"`
	Salary    string `xml:"salary,omitempty" bson:"salary,omitempty"`
	Funds     string `xml:"funds,omitempty" bson:"funds,omitempty"`
}

// Realization is generated from an XSD element
type Realization struct {
	Year []Year `xml:"year,omitempty" bson:"year,omitempty"`
}

// Year is generated from an XSD element
type Year struct {
	Year        int64                   `xml:"year,omitempty" bson:"year,omitempty"`
	Total       ZfcsBigProjectValueType `xml:"total,omitempty" bson:"total,omitempty"`
	Plan        ZfcsBigProjectValueType `xml:"plan,omitempty" bson:"plan,omitempty"`
	PlanAdvance ZfcsBigProjectValueType `xml:"planAdvance,omitempty" bson:"planAdvance,omitempty"`
	Paid        ZfcsBigProjectValueType `xml:"paid,omitempty" bson:"paid,omitempty"`
	PaidAdvance ZfcsBigProjectValueType `xml:"paidAdvance,omitempty" bson:"paidAdvance,omitempty"`
	Federal     ZfcsBigProjectValueType `xml:"federal,omitempty" bson:"federal,omitempty"`
	Region      ZfcsBigProjectValueType `xml:"region,omitempty" bson:"region,omitempty"`
	Self        ZfcsBigProjectValueType `xml:"self,omitempty" bson:"self,omitempty"`
	Nonbudget   ZfcsBigProjectValueType `xml:"nonbudget,omitempty" bson:"nonbudget,omitempty"`
	Salary      ZfcsBigProjectValueType `xml:"salary,omitempty" bson:"salary,omitempty"`
	Funds       ZfcsBigProjectValueType `xml:"funds,omitempty" bson:"funds,omitempty"`
}

// Changes is generated from an XSD element
type Changes struct {
	Year []Year `xml:"year,omitempty" bson:"year,omitempty"`
}

// ZfcsMasterDataType is generated from an XSD element
type ZfcsMasterDataType struct {
	SchemeVersion                   string                                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	NsiBankGuaranteeRefusalReason   ZfcsNsiBankGuaranteeRefusalReasonType   `xml:"nsiBankGuaranteeRefusalReason,omitempty" bson:"nsiBankGuaranteeRefusalReason,omitempty"`
	NsiBudget                       ZfcsNsiBudgetType                       `xml:"nsiBudget,omitempty" bson:"nsiBudget,omitempty"`
	NsiCalendarDays                 ZfcsNsiCalendarDaysType                 `xml:"nsiCalendarDays,omitempty" bson:"nsiCalendarDays,omitempty"`
	NsiCommission                   ZfcsNsiCommissionType                   `xml:"nsiCommission,omitempty" bson:"nsiCommission,omitempty"`
	NsiCommissionRole               ZfcsNsiCommissionRoleType               `xml:"nsiCommissionRole,omitempty" bson:"nsiCommissionRole,omitempty"`
	NsiContractPriceChangeReason    ZfcsNsiContractPriceChangeReasonType    `xml:"nsiContractPriceChangeReason,omitempty" bson:"nsiContractPriceChangeReason,omitempty"`
	NsiContractRefusalReason        ZfcsNsiContractRefusalReasonType        `xml:"nsiContractRefusalReason,omitempty" bson:"nsiContractRefusalReason,omitempty"`
	NsiContractSingleCustomerReason ZfcsNsiContractSingleCustomerReasonType `xml:"nsiContractSingleCustomerReason,omitempty" bson:"nsiContractSingleCustomerReason,omitempty"`
	NsiContractTerminationReason    ZfcsNsiContractTerminationReasonType    `xml:"nsiContractTerminationReason,omitempty" bson:"nsiContractTerminationReason,omitempty"`
	NsiContractModificationReason   ZfcsNsiContractModificationReasonType   `xml:"nsiContractModificationReason,omitempty" bson:"nsiContractModificationReason,omitempty"`
	NsiContractExecutionDoc         ZfcsNsiContractExecutionDocType         `xml:"nsiContractExecutionDoc,omitempty" bson:"nsiContractExecutionDoc,omitempty"`
	NsiContractReparationDoc        ZfcsNsiContractReparationDocType        `xml:"nsiContractReparationDoc,omitempty" bson:"nsiContractReparationDoc,omitempty"`
	NsiContractPenaltyReason        ZfcsNsiContractPenaltyReasonType        `xml:"nsiContractPenaltyReason,omitempty" bson:"nsiContractPenaltyReason,omitempty"`
	NsiContractOKOPFExtraBudget     ZfcsNsiContractOKOPFExtraBudgetType     `xml:"nsiContractOKOPFExtraBudget,omitempty" bson:"nsiContractOKOPFExtraBudget,omitempty"`
	NsiContractCurrencyCBRF         ZfcsNsiContractCurrencyCBRFType         `xml:"nsiContractCurrencyCBRF,omitempty" bson:"nsiContractCurrencyCBRF,omitempty"`
	NsiEvalCriterion                ZfcsNsiEvalCriterionType                `xml:"nsiEvalCriterion,omitempty" bson:"nsiEvalCriterion,omitempty"`
	NsiKBKBudget                    ZfcsNsiKBKBudgetType                    `xml:"nsiKBKBudget,omitempty" bson:"nsiKBKBudget,omitempty"`
	NsiKOSGU                        ZfcsNsiKOSGUType                        `xml:"nsiKOSGU,omitempty" bson:"nsiKOSGU,omitempty"`
	NsiOKOPF                        ZfcsNsiOKOPFType                        `xml:"nsiOKOPF,omitempty" bson:"nsiOKOPF,omitempty"`
	NsiOKPD                         ZfcsNsiOKPDType                         `xml:"nsiOKPD,omitempty" bson:"nsiOKPD,omitempty"`
	NsiOKVED                        ZfcsNsiOKVEDType                        `xml:"nsiOKVED,omitempty" bson:"nsiOKVED,omitempty"`
	NsiOrganization                 ZfcsNsiOrganizationType                 `xml:"nsiOrganization,omitempty" bson:"nsiOrganization,omitempty"`
	NsiOrganizationRights           ZfcsNsiOrganizationRightsType           `xml:"nsiOrganizationRights,omitempty" bson:"nsiOrganizationRights,omitempty"`
	NsiOrganizationType             ZfcsNsiOrganizationTypesType            `xml:"nsiOrganizationType,omitempty" bson:"nsiOrganizationType,omitempty"`
	NsiPlacingWay                   ZfcsNsiPlacingWayType                   `xml:"nsiPlacingWay,omitempty" bson:"nsiPlacingWay,omitempty"`
	NsiPlanPositionChangeReason     ZfcsNsiPlanPositionChangeReasonType     `xml:"nsiPlanPositionChangeReason,omitempty" bson:"nsiPlanPositionChangeReason,omitempty"`
	NsiPurchaseDocumentTypes        ZfcsNsiPurchaseDocumentTypesType        `xml:"nsiPurchaseDocumentTypes,omitempty" bson:"nsiPurchaseDocumentTypes,omitempty"`
	NsiPurchasePreferences          ZfcsNsiPurchasePreferencesType          `xml:"nsiPurchasePreferences,omitempty" bson:"nsiPurchasePreferences,omitempty"`
	NsiPurchaseRejectReason         ZfcsNsiPurchaseRejectReasonType         `xml:"nsiPurchaseRejectReason,omitempty" bson:"nsiPurchaseRejectReason,omitempty"`
	NsiUser                         ZfcsNsiUserType                         `xml:"nsiUser,omitempty" bson:"nsiUser,omitempty"`
	NsiAbandonedReason              ZfcsNsiAbandonedReasonType              `xml:"nsiAbandonedReason,omitempty" bson:"nsiAbandonedReason,omitempty"`
}

// QuestionType is generated from an XSD element
type QuestionType struct {
	NotificationNumber string       `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	RegNumber          string       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	RegDate            time.Time    `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	Subject            string       `xml:"subject,omitempty" bson:"subject,omitempty"`
	Href               string       `xml:"href,omitempty" bson:"href,omitempty"`
	DocumentMetas      DocumentList `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// ZfcsExtPrintFormType is generated from an XSD element
type ZfcsExtPrintFormType struct {
	Signature                Signature                `xml:"signature,omitempty" bson:"signature,omitempty"`
	FileType                 string                   `xml:"fileType,omitempty" bson:"fileType,omitempty"`
	ControlPersonalSignature ControlPersonalSignature `xml:"controlPersonalSignature,omitempty" bson:"controlPersonalSignature,omitempty"`
	Content                  []byte                   `xml:"content,omitempty" bson:"content,omitempty"`
	URL                      string                   `xml:"url,omitempty" bson:"url,omitempty"`
}

// ControlPersonalSignature is generated from an XSD element
type ControlPersonalSignature struct {
	Type                     string `xml:"type,attr" bson:"type,omitempty"`
	ControlPersonalSignature []byte `xml:",chardata" bson:",omitempty"`
}

// ZfcsComplaintOrderType is generated from an XSD element
type ZfcsComplaintOrderType struct {
	NotificationNumber string `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	Lots               Lots   `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// ZfcsCountryRef is generated from an XSD element
type ZfcsCountryRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsTendePlanInfoType is generated from an XSD element
type ZfcsTendePlanInfoType struct {
	PlanNumber     string `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	PositionNumber string `xml:"positionNumber,omitempty" bson:"positionNumber,omitempty"`
}

// AdmissionResults is generated from an XSD element
type AdmissionResults struct {
	AdmissionResult []AdmissionResult `xml:"admissionResult,omitempty" bson:"admissionResult,omitempty"`
}

// AdmissionResult is generated from an XSD element
type AdmissionResult struct {
	ProtocolCommissionMember CommissionMemberType    `xml:"protocolCommissionMember,omitempty" bson:"protocolCommissionMember,omitempty"`
	Admitted                 bool                    `xml:"admitted,omitempty" bson:"admitted,omitempty"`
	AppRejectedReason        []AppRejectedReasonType `xml:"appRejectedReason,omitempty" bson:"appRejectedReason,omitempty"`
}

// ApplicationFeaturesCorrespondence is generated from an XSD element
type ApplicationFeaturesCorrespondence struct {
	Compatible          bool                    `xml:"compatible,omitempty" bson:"compatible,omitempty"`
	NotificationFeature NotificationFeatureType `xml:"notificationFeature,omitempty" bson:"notificationFeature,omitempty"`
}

// ZfcsPaymentInfoType is generated from an XSD element
type ZfcsPaymentInfoType struct {
	Amount            string  `xml:"amount,omitempty" bson:"amount,omitempty"`
	Part              float64 `xml:"part,omitempty" bson:"part,omitempty"`
	ProcedureInfo     string  `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	SettlementAccount string  `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string  `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string  `xml:"bik,omitempty" bson:"bik,omitempty"`
}

// ZfcsProtocolEF2Type is generated from an XSD element
type ZfcsProtocolEF2Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsTenderPlanCancelType is generated from an XSD element
type ZfcsTenderPlanCancelType struct {
	SchemeVersion string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID            int64                `xml:"id,omitempty" bson:"id,omitempty"`
	PlanNumber    string               `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	CustomerInfo  CustomerInfo         `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	Year          int64                `xml:"year,omitempty" bson:"year,omitempty"`
	Description   string               `xml:"description,omitempty" bson:"description,omitempty"`
	VersionNumber int64                `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CancelDate    time.Time            `xml:"cancelDate,omitempty" bson:"cancelDate,omitempty"`
	PrintForm     ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsBankGuaranteeRefusalInvalidType is generated from an XSD element
type ZfcsBankGuaranteeRefusalInvalidType struct {
	SchemeVersion            string                   `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                    `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                   `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RefusalDocNumber         string                   `xml:"refusalDocNumber,omitempty" bson:"refusalDocNumber,omitempty"`
	DocNumber                string                   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate           time.Time                `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	BankGuaranteeRefusalInfo BankGuaranteeRefusalInfo `xml:"bankGuaranteeRefusalInfo,omitempty" bson:"bankGuaranteeRefusalInfo,omitempty"`
	Href                     string                   `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType     `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments              ZfcsAttachmentListType   `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Reason                   string                   `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// BankGuaranteeRefusalInfo is generated from an XSD element
type BankGuaranteeRefusalInfo struct {
	Bank         ZfcsBankGuaranteeOrganizationType `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo ZfcsBankGuaranteeParticipantType  `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee    ZfcsBankGuaranteeInfoType         `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	RefusalInfo  RefusalInfo                       `xml:"refusalInfo,omitempty" bson:"refusalInfo,omitempty"`
}

// ZfcsBankGuaranteeInvalidType is generated from an XSD element
type ZfcsBankGuaranteeInvalidType struct {
	SchemeVersion  string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNumber      string                 `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber      string                 `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Reason         string                 `xml:"reason,omitempty" bson:"reason,omitempty"`
	GuaranteeInfo  GuaranteeInfo          `xml:"guaranteeInfo,omitempty" bson:"guaranteeInfo,omitempty"`
	Href           string                 `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	Documents      ZfcsAttachmentListType `xml:"documents,omitempty" bson:"documents,omitempty"`
}

// GuaranteeInfo is generated from an XSD element
type GuaranteeInfo struct {
	Bank         ZfcsBankGuaranteeOrganizationType `xml:"bank,omitempty" bson:"bank,omitempty"`
	PlacingOrg   ZfcsBankGuaranteeOrganizationType `xml:"placingOrg,omitempty" bson:"placingOrg,omitempty"`
	SupplierInfo ZfcsBankGuaranteeParticipantType  `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee    ZfcsBankGuaranteeInfoType         `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
}

// ZfcsInfoProtocolZKBIType is generated from an XSD element
type ZfcsInfoProtocolZKBIType struct {
	Execution       Execution               `xml:"execution,omitempty" bson:"execution,omitempty"`
	Commission      Commission              `xml:"commission,omitempty" bson:"commission,omitempty"`
	PurchaseInfo    PurchaseInfo            `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	LotInfo         ZfcsLotInfoType         `xml:"lotInfo,omitempty" bson:"lotInfo,omitempty"`
	Applications    Applications            `xml:"applications,omitempty" bson:"applications,omitempty"`
	AbandonedReason ZfcsAbandonedReasonType `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
}

// Execution is generated from an XSD element
type Execution struct {
	Place       string    `xml:"place,omitempty" bson:"place,omitempty"`
	OpeningDate time.Time `xml:"openingDate,omitempty" bson:"openingDate,omitempty"`
	ScoringDate time.Time `xml:"scoringDate,omitempty" bson:"scoringDate,omitempty"`
	SignDate    time.Time `xml:"signDate,omitempty" bson:"signDate,omitempty"`
}

// ZfcsNsiPlacingWayType is generated from an XSD element
type ZfcsNsiPlacingWayType struct {
	PlacingWayID  int64     `xml:"placingWayId,omitempty" bson:"placingWayId,omitempty"`
	Code          string    `xml:"code,omitempty" bson:"code,omitempty"`
	Name          string    `xml:"name,omitempty" bson:"name,omitempty"`
	Type          string    `xml:"type,omitempty" bson:"type,omitempty"`
	SubsystemType string    `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
	Actual        bool      `xml:"actual,omitempty" bson:"actual,omitempty"`
	Documents     Documents `xml:"documents,omitempty" bson:"documents,omitempty"`
}

// ZfcsBankGuaranteeRefReasonType is generated from an XSD element
type ZfcsBankGuaranteeRefReasonType struct {
	Code int64  `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsRequestForQuotationType is generated from an XSD element
type ZfcsRequestForQuotationType struct {
	SchemeVersion      string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocPublishDate     time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	RegistryNum        string                 `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	VersionNumber      string                 `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	State              string                 `xml:"state,omitempty" bson:"state,omitempty"`
	PublishOrg         PublishOrg             `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                 `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	RequestObjectInfo  string                 `xml:"requestObjectInfo,omitempty" bson:"requestObjectInfo,omitempty"`
	ResponsibleInfo    ResponsibleInfo        `xml:"responsibleInfo,omitempty" bson:"responsibleInfo,omitempty"`
	ProcedureInfo      ProcedureInfo          `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Products           Products               `xml:"products,omitempty" bson:"products,omitempty"`
	Conditions         Conditions             `xml:"conditions,omitempty" bson:"conditions,omitempty"`
	Attachments        ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationReason string                 `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
}

// ZfcsNsiContractReparationDocType is generated from an XSD element
type ZfcsNsiContractReparationDocType struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNsiPlanPositionChangeReasonType is generated from an XSD element
type ZfcsNsiPlanPositionChangeReasonType struct {
	ID          int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
	Actual      bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Criterion is generated from an XSD element
type Criterion struct {
	NsiEvalCriterion EvalCriterion          `xml:"nsiEvalCriterion,omitempty" bson:"nsiEvalCriterion,omitempty"`
	CriterionValue   float64                `xml:"criterionValue,omitempty" bson:"criterionValue,omitempty"`
	ChildrenCriteria []ChildrenCriteriaType `xml:"childrenCriteria,omitempty" bson:"childrenCriteria,omitempty"`
	EvalValue        string                 `xml:"evalValue,omitempty" bson:"evalValue,omitempty"`
}

// ZfcsNsiRMISType is generated from an XSD element
type ZfcsNsiRMISType struct {
	RegNumber            string                 `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	CreateDate           time.Time              `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PublishDate          time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	NameRMIS             string                 `xml:"nameRMIS,omitempty" bson:"nameRMIS,omitempty"`
	TypeRMIS             string                 `xml:"typeRMIS,omitempty" bson:"typeRMIS,omitempty"`
	DescriptionRMIS      string                 `xml:"descriptionRMIS,omitempty" bson:"descriptionRMIS,omitempty"`
	StartDateRMIS        xsd.Date               `xml:"startDateRMIS,omitempty" bson:"startDateRMIS,omitempty"`
	RequisitesRMIS       string                 `xml:"requisitesRMIS,omitempty" bson:"requisitesRMIS,omitempty"`
	URLRMIS              string                 `xml:"urlRMIS,omitempty" bson:"urlRMIS,omitempty"`
	PPO                  ZfcsPPORef             `xml:"PPO,omitempty" bson:"PPO,omitempty"`
	TOFK                 ZfcsTOFKRef            `xml:"TOFK,omitempty" bson:"TOFK,omitempty"`
	FullName             string                 `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	CodeSPZ              string                 `xml:"codeSPZ,omitempty" bson:"codeSPZ,omitempty"`
	CodeSvodReestr       string                 `xml:"codeSvodReestr,omitempty" bson:"codeSvodReestr,omitempty"`
	OGRN                 string                 `xml:"OGRN,omitempty" bson:"OGRN,omitempty"`
	INN                  string                 `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP                  string                 `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	AuthorizedPersonInfo AuthorizedPersonInfo   `xml:"authorizedPersonInfo,omitempty" bson:"authorizedPersonInfo,omitempty"`
	CertificateInfo      CertificateInfo        `xml:"certificateInfo,omitempty" bson:"certificateInfo,omitempty"`
	Actual               bool                   `xml:"actual,omitempty" bson:"actual,omitempty"`
	Attachments          ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// AuthorizedPersonInfo is generated from an XSD element
type AuthorizedPersonInfo struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	Position   string `xml:"position,omitempty" bson:"position,omitempty"`
	Address    string `xml:"address,omitempty" bson:"address,omitempty"`
	Phone      string `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email      string `xml:"email,omitempty" bson:"email,omitempty"`
}

// CertificateInfo is generated from an XSD element
type CertificateInfo struct {
	Content []byte `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsGuaranteeAttachmentListType is generated from an XSD element
type ZfcsGuaranteeAttachmentListType struct {
	Attachment []ZfcsGuaranteeAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsPurchaseProcedureOKDType is generated from an XSD element
type ZfcsPurchaseProcedureOKDType struct {
	StageOne StageOne `xml:"stageOne,omitempty" bson:"stageOne,omitempty"`
	StageTwo StageTwo `xml:"stageTwo,omitempty" bson:"stageTwo,omitempty"`
}

// StageOne is generated from an XSD element
type StageOne struct {
	Collecting       ZfcsPurchaseProcedureCollectingType       `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Opening          ZfcsPurchaseProcedureOpeningType          `xml:"opening,omitempty" bson:"opening,omitempty"`
	Scoring          ZfcsPurchaseProcedureScoringPlaceType     `xml:"scoring,omitempty" bson:"scoring,omitempty"`
	Prequalification ZfcsPurchaseProcedurePrequalificationType `xml:"prequalification,omitempty" bson:"prequalification,omitempty"`
}

// StageTwo is generated from an XSD element
type StageTwo struct {
	Collecting ZfcsPurchaseProcedureCollectingType `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Opening    ZfcsPurchaseProcedureOpeningType    `xml:"opening,omitempty" bson:"opening,omitempty"`
	Scoring    ZfcsPurchaseProcedureScoringType    `xml:"scoring,omitempty" bson:"scoring,omitempty"`
}

// ZfcsBidType is generated from an XSD element
type ZfcsBidType struct {
	Price                string    `xml:"price,omitempty" bson:"price,omitempty"`
	Date                 time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	IncreaseInitialPrice bool      `xml:"increaseInitialPrice,omitempty" bson:"increaseInitialPrice,omitempty"`
}

// ZfcsOKVEDRef is generated from an XSD element
type ZfcsOKVEDRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNsiAuditActionSubjectsType is generated from an XSD element
type ZfcsNsiAuditActionSubjectsType struct {
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsOrganizationBudgetsType is generated from an XSD element
type ZfcsOrganizationBudgetsType struct {
	RegNumber string  `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	Budgets   Budgets `xml:"budgets,omitempty" bson:"budgets,omitempty"`
}

// ZfcsTenderPlanCommonInfoType is generated from an XSD element
type ZfcsTenderPlanCommonInfoType struct {
	ID            int64     `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID    string    `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PlanNumber    string    `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	Year          int64     `xml:"year,omitempty" bson:"year,omitempty"`
	VersionNumber int64     `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Owner         Owner     `xml:"owner,omitempty" bson:"owner,omitempty"`
	CreateDate    time.Time `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	Description   string    `xml:"description,omitempty" bson:"description,omitempty"`
	ConfirmDate   time.Time `xml:"confirmDate,omitempty" bson:"confirmDate,omitempty"`
	PublishDate   time.Time `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
}

// Owner is generated from an XSD element
type Owner struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// ZfcsAddInfoType is generated from an XSD element
type ZfcsAddInfoType struct {
	SchemeVersion      string                    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID         string                    `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID                 int64                     `xml:"id,omitempty" bson:"id,omitempty"`
	RegistryNum        string                    `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	PublishOrg         ZfcsOrganizationInfoType  `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	VersionNumber      string                    `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate     time.Time                 `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	FirstPublishDate   time.Time                 `xml:"firstPublishDate,omitempty" bson:"firstPublishDate,omitempty"`
	Href               string                    `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsContractPrintFormType `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	InfoType           string                    `xml:"infoType,omitempty" bson:"infoType,omitempty"`
	Attachments        ZfcsAttachmentListType    `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationReason string                    `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType      `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Contract           Contract                  `xml:"contract,omitempty" bson:"contract,omitempty"`
	Purchase           Purchase                  `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Contractor         Contractor                `xml:"contractor,omitempty" bson:"contractor,omitempty"`
}

// ZfcsTenderPlanTotalPositionKBK2016sType is generated from an XSD element
type ZfcsTenderPlanTotalPositionKBK2016sType struct {
	KBK2016 []KBK2016 `xml:"KBK2016,omitempty" bson:"KBK2016,omitempty"`
}

// ZfcsUnfairSupplierType is generated from an XSD element
type ZfcsUnfairSupplierType struct {
	SchemeVersion  string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RegistryNum    string                       `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	PublishDate    time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	ApproveDate    time.Time                    `xml:"approveDate,omitempty" bson:"approveDate,omitempty"`
	State          string                       `xml:"state,omitempty" bson:"state,omitempty"`
	PublishOrg     ZfcsOrganizationRef          `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	CreateReason   string                       `xml:"createReason,omitempty" bson:"createReason,omitempty"`
	ApproveReason  string                       `xml:"approveReason,omitempty" bson:"approveReason,omitempty"`
	Customer       ZfcsPurchaseOrganizationType `xml:"customer,omitempty" bson:"customer,omitempty"`
	UnfairSupplier UnfairSupplier               `xml:"unfairSupplier,omitempty" bson:"unfairSupplier,omitempty"`
	Purchase       Purchase                     `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	NotOosPurchase bool                         `xml:"notOosPurchase,omitempty" bson:"notOosPurchase,omitempty"`
	Contract       Contract                     `xml:"contract,omitempty" bson:"contract,omitempty"`
	Exclude        Exclude                      `xml:"exclude,omitempty" bson:"exclude,omitempty"`
}

// UnfairSupplier is generated from an XSD element
type UnfairSupplier struct {
	FullName string     `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	Type     string     `xml:"type,omitempty" bson:"type,omitempty"`
	FirmName string     `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Inn      string     `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp      string     `xml:"kpp,omitempty" bson:"kpp,omitempty"`
	Place    Place      `xml:"place,omitempty" bson:"place,omitempty"`
	Founders []Founders `xml:"founders,omitempty" bson:"founders,omitempty"`
}

// Place is generated from an XSD element
type Place struct {
	Zip     string         `xml:"zip,omitempty" bson:"zip,omitempty"`
	Place   string         `xml:"place,omitempty" bson:"place,omitempty"`
	Email   string         `xml:"email,omitempty" bson:"email,omitempty"`
	Kladr   Kladr          `xml:"kladr,omitempty" bson:"kladr,omitempty"`
	Country ZfcsCountryRef `xml:"country,omitempty" bson:"country,omitempty"`
}

// Kladr is generated from an XSD element
type Kladr struct {
	KladrType string `xml:"kladrType,omitempty" bson:"kladrType,omitempty"`
	KladrCode string `xml:"kladrCode,omitempty" bson:"kladrCode,omitempty"`
	FullName  string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	SubjectRF string `xml:"subjectRF,omitempty" bson:"subjectRF,omitempty"`
	Area      string `xml:"area,omitempty" bson:"area,omitempty"`
	City      string `xml:"city,omitempty" bson:"city,omitempty"`
	Street    string `xml:"street,omitempty" bson:"street,omitempty"`
	Building  string `xml:"building,omitempty" bson:"building,omitempty"`
	Office    string `xml:"office,omitempty" bson:"office,omitempty"`
}

// Founders is generated from an XSD element
type Founders struct {
	Names string `xml:"names,omitempty" bson:"names,omitempty"`
	Inn   string `xml:"inn,omitempty" bson:"inn,omitempty"`
}

// Exclude is generated from an XSD element
type Exclude struct {
	ExcludeDate time.Time `xml:"excludeDate,omitempty" bson:"excludeDate,omitempty"`
	Name        string    `xml:"name,omitempty" bson:"name,omitempty"`
	Date        xsd.Date  `xml:"date,omitempty" bson:"date,omitempty"`
	Number      string    `xml:"number,omitempty" bson:"number,omitempty"`
	Type        string    `xml:"type,omitempty" bson:"type,omitempty"`
}

// LotRefType is generated from an XSD element
type LotRefType struct {
	Currency CurrencyType `xml:"currency,omitempty" bson:"currency,omitempty"`
	Subject  string       `xml:"subject,omitempty" bson:"subject,omitempty"`
}

// ZfcsNsiKBKBudgetType is generated from an XSD element
type ZfcsNsiKBKBudgetType struct {
	Kbk    string `xml:"kbk,omitempty" bson:"kbk,omitempty"`
	Budget string `xml:"budget,omitempty" bson:"budget,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNsiOKOPFType is generated from an XSD element
type ZfcsNsiOKOPFType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	ParentCode   string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	FullName     string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
	Actual       bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsPurchaseDocumentCommonType is generated from an XSD element
type ZfcsPurchaseDocumentCommonType struct {
	SchemeVersion  string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time            `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsDocumentRequirementType is generated from an XSD element
type ZfcsDocumentRequirementType struct {
	Number    uint64 `xml:"number,omitempty" bson:"number,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	Mandatory bool   `xml:"mandatory,omitempty" bson:"mandatory,omitempty"`
}

// ZfcsPositionKVRsYearsType is generated from an XSD element
type ZfcsPositionKVRsYearsType struct {
	KVR []KVR `xml:"KVR,omitempty" bson:"KVR,omitempty"`
}

// ZfcsContractProcedureType is generated from an XSD element
type ZfcsContractProcedureType struct {
	SchemeVersion          string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                     int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID             string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNum                 string                         `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	PublishDate            time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	VersionNumber          int64                          `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Executions             Executions                     `xml:"executions,omitempty" bson:"executions,omitempty"`
	Terminations           Terminations                   `xml:"terminations,omitempty" bson:"terminations,omitempty"`
	Penalties              Penalties                      `xml:"penalties,omitempty" bson:"penalties,omitempty"`
	PrintForm              ZfcsContractPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	PaymentDocuments       ZfcsContractAttachmentListType `xml:"paymentDocuments,omitempty" bson:"paymentDocuments,omitempty"`
	ReceiptDocuments       ZfcsContractAttachmentListType `xml:"receiptDocuments,omitempty" bson:"receiptDocuments,omitempty"`
	ProductOriginDocuments ZfcsContractAttachmentListType `xml:"productOriginDocuments,omitempty" bson:"productOriginDocuments,omitempty"`
	Reason                 string                         `xml:"reason,omitempty" bson:"reason,omitempty"`
	CurrentContractStage   string                         `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// Executions is generated from an XSD element
type Executions struct {
	Stage               ZfcsStageType `xml:"stage,omitempty" bson:"stage,omitempty"`
	OrdinalNumber       int64         `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	FinalStageExecution bool          `xml:"finalStageExecution,omitempty" bson:"finalStageExecution,omitempty"`
	Execution           []Execution   `xml:"execution,omitempty" bson:"execution,omitempty"`
}

// Terminations is generated from an XSD element
type Terminations struct {
	Termination Termination `xml:"termination,omitempty" bson:"termination,omitempty"`
}

// Termination is generated from an XSD element
type Termination struct {
	Paid              string            `xml:"paid,omitempty" bson:"paid,omitempty"`
	TerminationDate   xsd.Date          `xml:"terminationDate,omitempty" bson:"terminationDate,omitempty"`
	Reason            string            `xml:"reason,omitempty" bson:"reason,omitempty"`
	TerminationReason TerminationReason `xml:"terminationReason,omitempty" bson:"terminationReason,omitempty"`
}

// TerminationReason is generated from an XSD element
type TerminationReason struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// Penalties is generated from an XSD element
type Penalties struct {
	Penalty   []Penalty                 `xml:"penalty,omitempty" bson:"penalty,omitempty"`
	PrintForm ZfcsContractPrintFormType `xml:"printForm,omitempty" bson:"printForm,omitempty"`
}

// Penalty is generated from an XSD element
type Penalty struct {
	ContractParty  string         `xml:"contractParty,omitempty" bson:"contractParty,omitempty"`
	PenaltyType    string         `xml:"penaltyType,omitempty" bson:"penaltyType,omitempty"`
	Foundation     string         `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	PenaltyAccrual PenaltyAccrual `xml:"penaltyAccrual,omitempty" bson:"penaltyAccrual,omitempty"`
	PenaltyChange  PenaltyChange  `xml:"penaltyChange,omitempty" bson:"penaltyChange,omitempty"`
}

// PenaltyAccrual is generated from an XSD element
type PenaltyAccrual struct {
	AccrualDecisionDate time.Time                   `xml:"accrualDecisionDate,omitempty" bson:"accrualDecisionDate,omitempty"`
	DocumentInfoList    ZfcsPenaltyDocumentInfoList `xml:"documentInfoList,omitempty" bson:"documentInfoList,omitempty"`
	AccrualAmount       string                      `xml:"accrualAmount,omitempty" bson:"accrualAmount,omitempty"`
	PenaltyFact         PenaltyFact                 `xml:"penaltyFact,omitempty" bson:"penaltyFact,omitempty"`
}

// PenaltyFact is generated from an XSD element
type PenaltyFact struct {
	DocumentInfoList          ZfcsPenaltyDocumentInfoList `xml:"documentInfoList,omitempty" bson:"documentInfoList,omitempty"`
	CollectedAmount           string                      `xml:"collectedAmount,omitempty" bson:"collectedAmount,omitempty"`
	CollectedAmountInPercents float64                     `xml:"collectedAmountInPercents,omitempty" bson:"collectedAmountInPercents,omitempty"`
}

// PenaltyChange is generated from an XSD element
type PenaltyChange struct {
	PenaltyChangeType  string                      `xml:"penaltyChangeType,omitempty" bson:"penaltyChangeType,omitempty"`
	ChangeDecisionDate time.Time                   `xml:"changeDecisionDate,omitempty" bson:"changeDecisionDate,omitempty"`
	DocumentInfoList   ZfcsPenaltyDocumentInfoList `xml:"documentInfoList,omitempty" bson:"documentInfoList,omitempty"`
	PenaltyChangeFact  PenaltyChangeFact           `xml:"penaltyChangeFact,omitempty" bson:"penaltyChangeFact,omitempty"`
}

// PenaltyChangeFact is generated from an XSD element
type PenaltyChangeFact struct {
	DocumentInfoList         ZfcsPenaltyDocumentInfoList `xml:"documentInfoList,omitempty" bson:"documentInfoList,omitempty"`
	ReturnedAmount           string                      `xml:"returnedAmount,omitempty" bson:"returnedAmount,omitempty"`
	ReturnedAmountInPercents float64                     `xml:"returnedAmountInPercents,omitempty" bson:"returnedAmountInPercents,omitempty"`
}

// ZfcsPurchaseProcedurePrequalificationType is generated from an XSD element
type ZfcsPurchaseProcedurePrequalificationType struct {
	Date  time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place string    `xml:"place,omitempty" bson:"place,omitempty"`
}

// OrganizationRef is generated from an XSD element
type OrganizationRef struct {
	RegNum   string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	FullName string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsNotificationLotChangeType is generated from an XSD element
type ZfcsNotificationLotChangeType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	Lot                   Lot                                  `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsProtocolOKD3Type is generated from an XSD element
type ZfcsProtocolOKD3Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       FoundationProtocol             `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
	OpeningProtocol          ZfcsFoundationProtocolInfoType `xml:"openingProtocol,omitempty" bson:"openingProtocol,omitempty"`
}

// FoundationProtocol is generated from an XSD element
type FoundationProtocol struct {
	Name  string    `xml:"name,omitempty" bson:"name,omitempty"`
	Date  time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place string    `xml:"place,omitempty" bson:"place,omitempty"`
	IsPPO bool      `xml:"isPPO,omitempty" bson:"isPPO,omitempty"`
}

// ZfcsCurrencyFullRef is generated from an XSD element
type ZfcsCurrencyFullRef struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	DigitalCode string `xml:"digitalCode,omitempty" bson:"digitalCode,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsAuditResultType is generated from an XSD element
type ZfcsAuditResultType struct {
	SchemeVersion      string                    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                     `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                    `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	VersionNumber      string                    `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	RegistryNum        string                    `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	DocPublishDate     time.Time                 `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	FirstPublishDate   time.Time                 `xml:"firstPublishDate,omitempty" bson:"firstPublishDate,omitempty"`
	PublishOrg         ZfcsOrganizationInfoType  `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                    `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsContractPrintFormType `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType      `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Type               string                    `xml:"type,omitempty" bson:"type,omitempty"`
	Name               string                    `xml:"name,omitempty" bson:"name,omitempty"`
	Document           Document                  `xml:"document,omitempty" bson:"document,omitempty"`
	Period             Period                    `xml:"period,omitempty" bson:"period,omitempty"`
	AddInfo            string                    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Action             Action                    `xml:"action,omitempty" bson:"action,omitempty"`
	Attachments        ZfcsAttachmentListType    `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ModificationReason string                    `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
}

// Action is generated from an XSD element
type Action struct {
	AuditOrg AuditOrg `xml:"auditOrg,omitempty" bson:"auditOrg,omitempty"`
	Subjects Subjects `xml:"subjects,omitempty" bson:"subjects,omitempty"`
	Objects  Objects  `xml:"objects,omitempty" bson:"objects,omitempty"`
	Period   Period   `xml:"period,omitempty" bson:"period,omitempty"`
}

// AuditOrg is generated from an XSD element
type AuditOrg struct {
	Registered    ZfcsOrganizationInfoWithOgrnType `xml:"registered,omitempty" bson:"registered,omitempty"`
	Nonregistered Nonregistered                    `xml:"nonregistered,omitempty" bson:"nonregistered,omitempty"`
}

// Nonregistered is generated from an XSD element
type Nonregistered struct {
	FullName string       `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	OGRN     string       `xml:"OGRN,omitempty" bson:"OGRN,omitempty"`
	OKTMO    ZfcsOKTMORef `xml:"OKTMO,omitempty" bson:"OKTMO,omitempty"`
}

// Subjects is generated from an XSD element
type Subjects struct {
	Subject []ZfcsAuditActionSubjectsRef `xml:"subject,omitempty" bson:"subject,omitempty"`
}

// Objects is generated from an XSD element
type Objects struct {
	Object []Object `xml:"object,omitempty" bson:"object,omitempty"`
}

// Object is generated from an XSD element
type Object struct {
	Violations    Violations                       `xml:"violations,omitempty" bson:"violations,omitempty"`
	Registered    ZfcsOrganizationInfoWithOgrnType `xml:"registered,omitempty" bson:"registered,omitempty"`
	Nonregistered Nonregistered                    `xml:"nonregistered,omitempty" bson:"nonregistered,omitempty"`
}

// ZfcsNsiContractRefusalReasonType is generated from an XSD element
type ZfcsNsiContractRefusalReasonType struct {
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// GuaranteeAppType is generated from an XSD element
type GuaranteeAppType struct {
	Procedure         string `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string `xml:"bik,omitempty" bson:"bik,omitempty"`
}

// NotificationOKType is generated from an XSD element
type NotificationOKType struct {
	ID                              int64                           `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber              string                          `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber    string                          `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                   int64                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                      time.Time                       `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                      PlacingWay                      `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                       string                          `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                           Order                           `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                     ContactInfoType                 `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                       Document                        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                   DocumentList                    `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                     time.Time                       `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate              time.Time                       `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                    ModificationType                `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                            string                          `xml:"href,omitempty" bson:"href,omitempty"`
	NotificationPlacement           NotificationPlacement           `xml:"notificationPlacement,omitempty" bson:"notificationPlacement,omitempty"`
	CompetitiveDocumentProvisioning CompetitiveDocumentProvisioning `xml:"competitiveDocumentProvisioning,omitempty" bson:"competitiveDocumentProvisioning,omitempty"`
	NotificationCommission          NotificationCommission          `xml:"notificationCommission,omitempty" bson:"notificationCommission,omitempty"`
	Lots                            Lots                            `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// NotificationPlacement is generated from an XSD element
type NotificationPlacement struct {
	AdditionalInfo       string                `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	GuaranteeApp         GuaranteeAppType      `xml:"guaranteeApp,omitempty" bson:"guaranteeApp,omitempty"`
	GuaranteeContract    GuaranteeContractType `xml:"guaranteeContract,omitempty" bson:"guaranteeContract,omitempty"`
	NotificationFeatures NotificationFeatures  `xml:"notificationFeatures,omitempty" bson:"notificationFeatures,omitempty"`
}

// CompetitiveDocumentProvisioning is generated from an XSD element
type CompetitiveDocumentProvisioning struct {
	DeliveryTerm      time.Time `xml:"deliveryTerm,omitempty" bson:"deliveryTerm,omitempty"`
	DeliveryTerm2     time.Time `xml:"deliveryTerm2,omitempty" bson:"deliveryTerm2,omitempty"`
	DeliveryPlace     string    `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	DeliveryProcedure string    `xml:"deliveryProcedure,omitempty" bson:"deliveryProcedure,omitempty"`
	Www               string    `xml:"www,omitempty" bson:"www,omitempty"`
	Guarantee         Guarantee `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
}

// Guarantee is generated from an XSD element
type Guarantee struct {
	Procedure         string       `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	SettlementAccount string       `xml:"settlementAccount,omitempty" bson:"settlementAccount,omitempty"`
	PersonalAccount   string       `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
	Bik               string       `xml:"bik,omitempty" bson:"bik,omitempty"`
	Amount            float64      `xml:"amount,omitempty" bson:"amount,omitempty"`
	Currency          CurrencyType `xml:"currency,omitempty" bson:"currency,omitempty"`
}

// ZfcsDecisionRef is generated from an XSD element
type ZfcsDecisionRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsUnplannedCheckCancelType is generated from an XSD element
type ZfcsUnplannedCheckCancelType struct {
	SchemeVersion string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CheckNumber   string                 `xml:"checkNumber,omitempty" bson:"checkNumber,omitempty"`
	PublishDate   time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Text          string                 `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm     ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsNsiBankGuaranteeRefusalReasonType is generated from an XSD element
type ZfcsNsiBankGuaranteeRefusalReasonType struct {
	ID     int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ProductsType is generated from an XSD element
type ProductsType struct {
	Product []ProductType `xml:"product,omitempty" bson:"product,omitempty"`
}

// ZfcsComplaintObjectType is generated from an XSD element
type ZfcsComplaintObjectType struct {
	Purchase   ZfcsComplaintPurchaseType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Order      ZfcsComplaintOrderType    `xml:"order,omitempty" bson:"order,omitempty"`
	SketchPlan SketchPlan                `xml:"sketchPlan,omitempty" bson:"sketchPlan,omitempty"`
	TenderPlan TenderPlan                `xml:"tenderPlan,omitempty" bson:"tenderPlan,omitempty"`
}

// SketchPlan is generated from an XSD element
type SketchPlan struct {
	PlanDescription string `xml:"planDescription,omitempty" bson:"planDescription,omitempty"`
	YearPlan        int64  `xml:"yearPlan,omitempty" bson:"yearPlan,omitempty"`
}

// TenderPlan is generated from an XSD element
type TenderPlan struct {
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsContractProcedure2015Type is generated from an XSD element
type ZfcsContractProcedure2015Type struct {
	SchemeVersion            string                                                `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                                                 `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                                                `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNum                   string                                                `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	PublishDate              time.Time                                             `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	VersionNumber            int64                                                 `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Executions               Executions                                            `xml:"executions,omitempty" bson:"executions,omitempty"`
	Termination              Termination                                           `xml:"termination,omitempty" bson:"termination,omitempty"`
	BankGuaranteeTermination ZfcsContractProcedure2015BankGuaranteeTerminationType `xml:"bankGuaranteeTermination,omitempty" bson:"bankGuaranteeTermination,omitempty"`
	Penalties                []Penalties                                           `xml:"penalties,omitempty" bson:"penalties,omitempty"`
	DelayWriteOffPenalties   DelayWriteOffPenalties                                `xml:"delayWriteOffPenalties,omitempty" bson:"delayWriteOffPenalties,omitempty"`
	BankGuaranteePayment     BankGuaranteePayment                                  `xml:"bankGuaranteePayment,omitempty" bson:"bankGuaranteePayment,omitempty"`
	HoldCashEnforcement      HoldCashEnforcement                                   `xml:"holdCashEnforcement,omitempty" bson:"holdCashEnforcement,omitempty"`
	PrintForm                ZfcsContractPrintFormType                             `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType                                  `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PaymentDocuments         ZfcsContractAttachmentListType                        `xml:"paymentDocuments,omitempty" bson:"paymentDocuments,omitempty"`
	ReceiptDocuments         ZfcsContractAttachmentListType                        `xml:"receiptDocuments,omitempty" bson:"receiptDocuments,omitempty"`
	ProductOriginDocuments   ZfcsContractAttachmentListType                        `xml:"productOriginDocuments,omitempty" bson:"productOriginDocuments,omitempty"`
	ModificationReason       string                                                `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	CurrentContractStage     string                                                `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
	Okpd2okved2              bool                                                  `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// DelayWriteOffPenalties is generated from an XSD element
type DelayWriteOffPenalties struct {
	TotalAmount             string            `xml:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	DelayPenaltiesInProcent float64           `xml:"delayPenaltiesInProcent,omitempty" bson:"delayPenaltiesInProcent,omitempty"`
	DelayPenalties          DelayPenalties    `xml:"delayPenalties,omitempty" bson:"delayPenalties,omitempty"`
	WriteOffPenalties       WriteOffPenalties `xml:"writeOffPenalties,omitempty" bson:"writeOffPenalties,omitempty"`
}

// DelayPenalties is generated from an XSD element
type DelayPenalties struct {
	DelayDate      xsd.Date                     `xml:"delayDate,omitempty" bson:"delayDate,omitempty"`
	Currency       ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	DelayAmount    string                       `xml:"delayAmount,omitempty" bson:"delayAmount,omitempty"`
	CurrencyRate   ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	DelayAmountRUR string                       `xml:"delayAmountRUR,omitempty" bson:"delayAmountRUR,omitempty"`
	DelayPeriod    xsd.Date                     `xml:"delayPeriod,omitempty" bson:"delayPeriod,omitempty"`
	NoticeDetails  NoticeDetails                `xml:"noticeDetails,omitempty" bson:"noticeDetails,omitempty"`
}

// NoticeDetails is generated from an XSD element
type NoticeDetails struct {
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
}

// WriteOffPenalties is generated from an XSD element
type WriteOffPenalties struct {
	WriteOffDate      xsd.Date                     `xml:"writeOffDate,omitempty" bson:"writeOffDate,omitempty"`
	Currency          ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	WriteOffAmount    string                       `xml:"writeOffAmount,omitempty" bson:"writeOffAmount,omitempty"`
	CurrencyRate      ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	WriteOffAmountRUR string                       `xml:"writeOffAmountRUR,omitempty" bson:"writeOffAmountRUR,omitempty"`
	NoticeDetails     NoticeDetails                `xml:"noticeDetails,omitempty" bson:"noticeDetails,omitempty"`
}

// BankGuaranteePayment is generated from an XSD element
type BankGuaranteePayment struct {
	RegNumber                    string       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber                    string       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	ImproperExecInfo             string       `xml:"improperExecInfo,omitempty" bson:"improperExecInfo,omitempty"`
	Requirements                 Requirements `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	Paid                         Paid         `xml:"paid,omitempty" bson:"paid,omitempty"`
	BankCancelDetails            string       `xml:"bankCancelDetails,omitempty" bson:"bankCancelDetails,omitempty"`
	ImproperGuaranteePaymentInfo string       `xml:"improperGuaranteePaymentInfo,omitempty" bson:"improperGuaranteePaymentInfo,omitempty"`
	Restructure                  Restructure  `xml:"restructure,omitempty" bson:"restructure,omitempty"`
}

// Paid is generated from an XSD element
type Paid struct {
	Name         string                       `xml:"name,omitempty" bson:"name,omitempty"`
	DocumentDate xsd.Date                     `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	DocumentNum  string                       `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	Currency     ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	BankPaid     string                       `xml:"bankPaid,omitempty" bson:"bankPaid,omitempty"`
	CurrencyRate ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	BankPaidRUR  string                       `xml:"bankPaidRUR,omitempty" bson:"bankPaidRUR,omitempty"`
}

// Restructure is generated from an XSD element
type Restructure struct {
	RestructureDate   xsd.Date `xml:"restructureDate,omitempty" bson:"restructureDate,omitempty"`
	RestructureAmount string   `xml:"restructureAmount,omitempty" bson:"restructureAmount,omitempty"`
	RepaymentSchedule string   `xml:"repaymentSchedule,omitempty" bson:"repaymentSchedule,omitempty"`
}

// HoldCashEnforcement is generated from an XSD element
type HoldCashEnforcement struct {
	ImproperSupplierInfo string                       `xml:"improperSupplierInfo,omitempty" bson:"improperSupplierInfo,omitempty"`
	Currency             ZfcsCurrencyRef              `xml:"currency,omitempty" bson:"currency,omitempty"`
	HoldAmount           string                       `xml:"holdAmount,omitempty" bson:"holdAmount,omitempty"`
	HoldDate             xsd.Date                     `xml:"holdDate,omitempty" bson:"holdDate,omitempty"`
	CurrencyRate         ZfcsCurrencyRateContract2015 `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	HoldAmountRUR        string                       `xml:"holdAmountRUR,omitempty" bson:"holdAmountRUR,omitempty"`
}

// ZfcsStandardContractType is generated from an XSD element
type ZfcsStandardContractType struct {
	ID                     int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	DocPublishDate         time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	StandardContractNumber string                       `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	Type                   string                       `xml:"type,omitempty" bson:"type,omitempty"`
	ApproveInfo            ApproveInfo                  `xml:"approveInfo,omitempty" bson:"approveInfo,omitempty"`
	PlacerOrganization     ZfcsPurchaseOrganizationType `xml:"placerOrganization,omitempty" bson:"placerOrganization,omitempty"`
	Indications            Indications                  `xml:"indications,omitempty" bson:"indications,omitempty"`
	UseCases               UseCases                     `xml:"useCases,omitempty" bson:"useCases,omitempty"`
	PrintForm              ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	Attachments            ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification           Modification                 `xml:"modification,omitempty" bson:"modification,omitempty"`
	Okpd2okved2            bool                         `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// ApproveInfo is generated from an XSD element
type ApproveInfo struct {
	Organization ZfcsPurchaseOrganizationType `xml:"organization,omitempty" bson:"organization,omitempty"`
	Date         time.Time                    `xml:"date,omitempty" bson:"date,omitempty"`
	Document     Document                     `xml:"document,omitempty" bson:"document,omitempty"`
}

// Indications is generated from an XSD element
type Indications struct {
	PurchaseObjects PurchaseObjects `xml:"purchaseObjects,omitempty" bson:"purchaseObjects,omitempty"`
	ContractPrice   ContractPrice   `xml:"contractPrice,omitempty" bson:"contractPrice,omitempty"`
	OtherIndicators string          `xml:"otherIndicators,omitempty" bson:"otherIndicators,omitempty"`
}

// ContractPrice is generated from an XSD element
type ContractPrice struct {
	MinPrice string `xml:"minPrice,omitempty" bson:"minPrice,omitempty"`
	MaxPrice string `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
}

// UseCases is generated from an XSD element
type UseCases struct {
	Type          []string `xml:"type,omitempty" bson:"type,omitempty"`
	UseTerms      string   `xml:"useTerms,omitempty" bson:"useTerms,omitempty"`
	AddInfo       string   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	RequiredTerms string   `xml:"requiredTerms,omitempty" bson:"requiredTerms,omitempty"`
}

// ZfcsOKEIRef is generated from an XSD element
type ZfcsOKEIRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsContractProcedure2015BankGuaranteeTerminationType is generated from an XSD element
type ZfcsContractProcedure2015BankGuaranteeTerminationType struct {
	RegNumber         string   `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber         string   `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	TerminationDate   xsd.Date `xml:"terminationDate,omitempty" bson:"terminationDate,omitempty"`
	TerminationReason string   `xml:"terminationReason,omitempty" bson:"terminationReason,omitempty"`
}

// ZfcsPurchaseNotificationISType is generated from an XSD element
type ZfcsPurchaseNotificationISType struct {
	SchemeVersion       string              `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64               `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string              `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string              `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time           `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string              `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string              `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType   `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	PurchaseObjectInfo  string              `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType  `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
}

// ZfcsPlanPositionChangeReasonRef is generated from an XSD element
type ZfcsPlanPositionChangeReasonRef struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// BidType is generated from an XSD element
type BidType struct {
	Price                float64   `xml:"price,omitempty" bson:"price,omitempty"`
	Date                 time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	IncreaseInitialPrice bool      `xml:"increaseInitialPrice,omitempty" bson:"increaseInitialPrice,omitempty"`
}

// NotificationZKType is generated from an XSD element
type NotificationZKType struct {
	ID                              int64                           `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber              string                          `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	FoundationNotificationNumber    string                          `xml:"foundationNotificationNumber,omitempty" bson:"foundationNotificationNumber,omitempty"`
	VersionNumber                   int64                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate                      time.Time                       `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PlacingWay                      PlacingWay                      `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	OrderName                       string                          `xml:"orderName,omitempty" bson:"orderName,omitempty"`
	Order                           Order                           `xml:"order,omitempty" bson:"order,omitempty"`
	ContactInfo                     ContactInfoType                 `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	PrintForm                       Document                        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas                   DocumentList                    `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	PublishDate                     time.Time                       `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PublishVersionDate              time.Time                       `xml:"publishVersionDate,omitempty" bson:"publishVersionDate,omitempty"`
	Modification                    ModificationType                `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                            string                          `xml:"href,omitempty" bson:"href,omitempty"`
	CompetitiveDocumentProvisioning CompetitiveDocumentProvisioning `xml:"competitiveDocumentProvisioning,omitempty" bson:"competitiveDocumentProvisioning,omitempty"`
	NotificationCommission          NotificationCommission          `xml:"notificationCommission,omitempty" bson:"notificationCommission,omitempty"`
	Lots                            Lots                            `xml:"lots,omitempty" bson:"lots,omitempty"`
}

// ZfcsPublicDiscussionPurchaseInfoType is generated from an XSD element
type ZfcsPublicDiscussionPurchaseInfoType struct {
	PlanNumber           string                       `xml:"planNumber,omitempty" bson:"planNumber,omitempty"`
	PositionNumber       string                       `xml:"positionNumber,omitempty" bson:"positionNumber,omitempty"`
	PurchaseNumber       string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	LotNumber            int64                        `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	PlanObjectInfo       string                       `xml:"planObjectInfo,omitempty" bson:"planObjectInfo,omitempty"`
	PlacingWay           string                       `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Year                 int64                        `xml:"year,omitempty" bson:"year,omitempty"`
	PlanContractMaxPrice string                       `xml:"planContractMaxPrice,omitempty" bson:"planContractMaxPrice,omitempty"`
	Customer             ZfcsPurchaseOrganizationType `xml:"customer,omitempty" bson:"customer,omitempty"`
}

// ZfcsSubStageType is generated from an XSD element
type ZfcsSubStageType struct {
	Month         byte  `xml:"month,omitempty" bson:"month,omitempty"`
	Year          int64 `xml:"year,omitempty" bson:"year,omitempty"`
	SubstageMonth byte  `xml:"substageMonth,omitempty" bson:"substageMonth,omitempty"`
	SubstageYear  int64 `xml:"substageYear,omitempty" bson:"substageYear,omitempty"`
}

// ZfcsCustomerReportSmallScaleBusinessType is generated from an XSD element
type ZfcsCustomerReportSmallScaleBusinessType struct {
	SchemeVersion      string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate            time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate     time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber      string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg         ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href               string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm       ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments        ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID           string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	ModificationReason string                           `xml:"modificationReason,omitempty" bson:"modificationReason,omitempty"`
	Signer             Signer                           `xml:"signer,omitempty" bson:"signer,omitempty"`
	Customer           ZfcsOrganizationInfoExtendedType `xml:"customer,omitempty" bson:"customer,omitempty"`
	ReportingPeriod    int64                            `xml:"reportingPeriod,omitempty" bson:"reportingPeriod,omitempty"`
	QuantityPurchase   QuantityPurchase                 `xml:"quantityPurchase,omitempty" bson:"quantityPurchase,omitempty"`
	ContractsInfo      ContractsInfo                    `xml:"contractsInfo,omitempty" bson:"contractsInfo,omitempty"`
}

// QuantityPurchase is generated from an XSD element
type QuantityPurchase struct {
	Privacy          float64 `xml:"privacy,omitempty" bson:"privacy,omitempty"`
	AnnualVolumeSt30 float64 `xml:"annualVolumeSt30,omitempty" bson:"annualVolumeSt30,omitempty"`
	Lending          float64 `xml:"lending,omitempty" bson:"lending,omitempty"`
	SingleSupplier   float64 `xml:"singleSupplier,omitempty" bson:"singleSupplier,omitempty"`
	NuclearEnergy    float64 `xml:"nuclearEnergy,omitempty" bson:"nuclearEnergy,omitempty"`
	ZK               float64 `xml:"ZK,omitempty" bson:"ZK,omitempty"`
	AnnualVolume     float64 `xml:"annualVolume,omitempty" bson:"annualVolume,omitempty"`
	Percent15        float64 `xml:"percent15,omitempty" bson:"percent15,omitempty"`
	OnlySMP          float64 `xml:"onlySMP,omitempty" bson:"onlySMP,omitempty"`
	NotSMP           float64 `xml:"notSMP,omitempty" bson:"notSMP,omitempty"`
	AnnualVolumeSMP  float64 `xml:"annualVolumeSMP,omitempty" bson:"annualVolumeSMP,omitempty"`
	RateSMP          float64 `xml:"rateSMP,omitempty" bson:"rateSMP,omitempty"`
	AbandonedSum     float64 `xml:"abandonedSum,omitempty" bson:"abandonedSum,omitempty"`
}

// ZfcsPurchaseProcedureOpeningType is generated from an XSD element
type ZfcsPurchaseProcedureOpeningType struct {
	Date    time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place   string    `xml:"place,omitempty" bson:"place,omitempty"`
	AddInfo string    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ZfcsEtpPrivilege is generated from an XSD element
type ZfcsEtpPrivilege struct {
	Etp          ZfcsETPType         `xml:"etp,omitempty" bson:"etp,omitempty"`
	EtpAction    string              `xml:"etpAction,omitempty" bson:"etpAction,omitempty"`
	Organization ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	Status       string              `xml:"status,omitempty" bson:"status,omitempty"`
}

// ZfcsNonbudgetFinancingsType is generated from an XSD element
type ZfcsNonbudgetFinancingsType struct {
	NonbudgetFinancing []NonbudgetFinancing `xml:"nonbudgetFinancing,omitempty" bson:"nonbudgetFinancing,omitempty"`
	TotalSum           string               `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// NonbudgetFinancing is generated from an XSD element
type NonbudgetFinancing struct {
	Year      int64  `xml:"year,omitempty" bson:"year,omitempty"`
	Sum       string `xml:"sum,omitempty" bson:"sum,omitempty"`
	KosguCode string `xml:"kosguCode,omitempty" bson:"kosguCode,omitempty"`
	KvrCode   string `xml:"kvrCode,omitempty" bson:"kvrCode,omitempty"`
}

// ZfcsLotI111Type is generated from an XSD element
type ZfcsLotI111Type struct {
	LotNumber              uint64                   `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	MaxPrice               string                   `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	Currency               ZfcsCurrencyRef          `xml:"currency,omitempty" bson:"currency,omitempty"`
	OKPD2                  []ZfcsOKPDRef            `xml:"OKPD2,omitempty" bson:"OKPD2,omitempty"`
	Preferenses            Preferenses              `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements           Requirements             `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	AddInfo                string                   `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	PublicDiscussion       ZfcsPublicDiscussionType `xml:"publicDiscussion,omitempty" bson:"publicDiscussion,omitempty"`
	MaxCostDefinitionOrder string                   `xml:"maxCostDefinitionOrder,omitempty" bson:"maxCostDefinitionOrder,omitempty"`
}

// ZfcsPurchaseDocumentType is generated from an XSD element
type ZfcsPurchaseDocumentType struct {
	SchemeVersion  string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                 `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string                 `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time              `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string                 `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	DocType        ZfcsDocType            `xml:"docType,omitempty" bson:"docType,omitempty"`
	AddInfo        string                 `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Attachments    ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification   Modification           `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsNsiOrganizationRightsType is generated from an XSD element
type ZfcsNsiOrganizationRightsType struct {
	RegNumber         string            `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	OrganizationLinks OrganizationLinks `xml:"organizationLinks,omitempty" bson:"organizationLinks,omitempty"`
}

// OrganizationLinks is generated from an XSD element
type OrganizationLinks struct {
	OrganizationLink []ZfcsOrganizationLink `xml:"organizationLink,omitempty" bson:"organizationLink,omitempty"`
}

// ZfcsPurchasePlanType is generated from an XSD element
type ZfcsPurchasePlanType struct {
	SchemeVersion     string                      `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo        CommonInfo                  `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CustomerInfo      CustomerInfo                `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	LocalInfo         ZfcsPurchasePlanAddInfoType `xml:"localInfo,omitempty" bson:"localInfo,omitempty"`
	ProvidedPurchases ProvidedPurchases           `xml:"providedPurchases,omitempty" bson:"providedPurchases,omitempty"`
	Attachments       ZfcsAttachmentListType      `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PrintForm         ZfcsPrintFormType           `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType        `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// DocReqType is generated from an XSD element
type DocReqType struct {
	DocumentRequirement []DocumentRequirement `xml:"documentRequirement,omitempty" bson:"documentRequirement,omitempty"`
}

// ParticipantType is generated from an XSD element
type ParticipantType struct {
	ParticipantType   string            `xml:"participantType,omitempty" bson:"participantType,omitempty"`
	Inn               string            `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp               string            `xml:"kpp,omitempty" bson:"kpp,omitempty"`
	OrganizationForm  string            `xml:"organizationForm,omitempty" bson:"organizationForm,omitempty"`
	LegalForm         OkopfType         `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	IDNumber          string            `xml:"idNumber,omitempty" bson:"idNumber,omitempty"`
	IDNumberExtension string            `xml:"idNumberExtension,omitempty" bson:"idNumberExtension,omitempty"`
	OrganizationName  string            `xml:"organizationName,omitempty" bson:"organizationName,omitempty"`
	Country           CountryType       `xml:"country,omitempty" bson:"country,omitempty"`
	FactualAddress    string            `xml:"factualAddress,omitempty" bson:"factualAddress,omitempty"`
	PostAddress       string            `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	ContactInfo       ContactPersonType `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	ContactEMail      string            `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone      string            `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax        string            `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AdditionalInfo    string            `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	Status            string            `xml:"status,omitempty" bson:"status,omitempty"`
}

// ProtocolEvasionType is generated from an XSD element
type ProtocolEvasionType struct {
	ID                        int64           `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber        string          `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber            string          `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber  string          `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber      string          `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                     string          `xml:"place,omitempty" bson:"place,omitempty"`
	VersionNumber             int64           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	ProtocolDate              time.Time       `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                  time.Time       `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate               time.Time       `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	PrintForm                 Document        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas             DocumentList    `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Href                      string          `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots              ProtocolLots    `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	RefusalFacts              RefusalFacts    `xml:"refusalFacts,omitempty" bson:"refusalFacts,omitempty"`
	NewProtocolTypeIndication bool            `xml:"newProtocolTypeIndication,omitempty" bson:"newProtocolTypeIndication,omitempty"`
	Customer                  OrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
}

// RefusalFacts is generated from an XSD element
type RefusalFacts struct {
	RefusalFact []RefusalFact `xml:"refusalFact,omitempty" bson:"refusalFact,omitempty"`
}

// ZfcsEventResultCancelType is generated from an XSD element
type ZfcsEventResultCancelType struct {
	SchemeVersion string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo             `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CancelType    string                 `xml:"cancelType,omitempty" bson:"cancelType,omitempty"`
	Info          string                 `xml:"info,omitempty" bson:"info,omitempty"`
	DocumentName  string                 `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentDate  time.Time              `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	PrintForm     ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsProtocolEvasionType is generated from an XSD element
type ZfcsProtocolEvasionType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	FoundationProtocolName   string                       `xml:"foundationProtocolName,omitempty" bson:"foundationProtocolName,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// AppRejectedReasonType is generated from an XSD element
type AppRejectedReasonType struct {
	NsiRejectReason NsiRejectReason `xml:"nsiRejectReason,omitempty" bson:"nsiRejectReason,omitempty"`
	Explanation     string          `xml:"explanation,omitempty" bson:"explanation,omitempty"`
}

// ZfcsComplaintCommonInfoType is generated from an XSD element
type ZfcsComplaintCommonInfoType struct {
	ComplaintNumber  string              `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	VersionNumber    int64               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PlanDecisionDate time.Time           `xml:"planDecisionDate,omitempty" bson:"planDecisionDate,omitempty"`
	DecisionPlace    string              `xml:"decisionPlace,omitempty" bson:"decisionPlace,omitempty"`
	RegistrationKO   ZfcsOrganizationRef `xml:"registrationKO,omitempty" bson:"registrationKO,omitempty"`
	ConsiderationKO  ZfcsOrganizationRef `xml:"considerationKO,omitempty" bson:"considerationKO,omitempty"`
	RegType          string              `xml:"regType,omitempty" bson:"regType,omitempty"`
}

// ZfcsPurchaseProtocolEFNoCommissionType is generated from an XSD element
type ZfcsPurchaseProtocolEFNoCommissionType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsCommissionRoleType is generated from an XSD element
type ZfcsCommissionRoleType struct {
	ID        int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	RightVote bool   `xml:"rightVote,omitempty" bson:"rightVote,omitempty"`
}

// ZfcsParticipantType is generated from an XSD element
type ZfcsParticipantType struct {
	ParticipantType   string                `xml:"participantType,omitempty" bson:"participantType,omitempty"`
	Inn               string                `xml:"inn,omitempty" bson:"inn,omitempty"`
	Kpp               string                `xml:"kpp,omitempty" bson:"kpp,omitempty"`
	Ogrn              string                `xml:"ogrn,omitempty" bson:"ogrn,omitempty"`
	LegalForm         ZfcsOkopfRef          `xml:"legalForm,omitempty" bson:"legalForm,omitempty"`
	IDNumber          string                `xml:"idNumber,omitempty" bson:"idNumber,omitempty"`
	IDNumberExtension string                `xml:"idNumberExtension,omitempty" bson:"idNumberExtension,omitempty"`
	OrganizationName  string                `xml:"organizationName,omitempty" bson:"organizationName,omitempty"`
	FirmName          string                `xml:"firmName,omitempty" bson:"firmName,omitempty"`
	Country           ZfcsCountryRef        `xml:"country,omitempty" bson:"country,omitempty"`
	FactualAddress    string                `xml:"factualAddress,omitempty" bson:"factualAddress,omitempty"`
	PostAddress       string                `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	ContactInfo       ZfcsContactPersonType `xml:"contactInfo,omitempty" bson:"contactInfo,omitempty"`
	ContactEMail      string                `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
	ContactPhone      string                `xml:"contactPhone,omitempty" bson:"contactPhone,omitempty"`
	ContactFax        string                `xml:"contactFax,omitempty" bson:"contactFax,omitempty"`
	AdditionalInfo    string                `xml:"additionalInfo,omitempty" bson:"additionalInfo,omitempty"`
	Status            string                `xml:"status,omitempty" bson:"status,omitempty"`
}

// ZfcsOKPORef is generated from an XSD element
type ZfcsOKPORef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsPurchasePlanPositionType is generated from an XSD element
type ZfcsPurchasePlanPositionType struct {
	CommonInfo  CommonInfo                  `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	FinanceInfo FinanceInfo                 `xml:"financeInfo,omitempty" bson:"financeInfo,omitempty"`
	LocalInfo   ZfcsPurchasePlanAddInfoType `xml:"localInfo,omitempty" bson:"localInfo,omitempty"`
}

// FinanceInfo is generated from an XSD element
type FinanceInfo struct {
	PublishYear  int64                                  `xml:"publishYear,omitempty" bson:"publishYear,omitempty"`
	Finances     ZfcsPurchasePlanPositionFinancingsType `xml:"finances,omitempty" bson:"finances,omitempty"`
	YearFinances ZfcsFinanceResourcesType               `xml:"yearFinances,omitempty" bson:"yearFinances,omitempty"`
}

// ZfcsOkopfRef is generated from an XSD element
type ZfcsOkopfRef struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	SingularName string `xml:"singularName,omitempty" bson:"singularName,omitempty"`
}

// ZfcsTenderPlanFinalPositionsType is generated from an XSD element
type ZfcsTenderPlanFinalPositionsType struct {
	Purchase83        Purchase83        `xml:"purchase83,omitempty" bson:"purchase83,omitempty"`
	Purchase83st544   Purchase83st544   `xml:"purchase83st544,omitempty" bson:"purchase83st544,omitempty"`
	Purchase93        Purchase93        `xml:"purchase93,omitempty" bson:"purchase93,omitempty"`
	OutcomeIndicators OutcomeIndicators `xml:"outcomeIndicators,omitempty" bson:"outcomeIndicators,omitempty"`
}

// Purchase83 is generated from an XSD element
type Purchase83 struct {
	TeachingService TeachingService `xml:"teachingService,omitempty" bson:"teachingService,omitempty"`
	GUIDeService    GUIDeService    `xml:"guideService,omitempty" bson:"guideService,omitempty"`
	Medicine        Medicine        `xml:"medicine,omitempty" bson:"medicine,omitempty"`
}

// TeachingService is generated from an XSD element
type TeachingService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// GUIDeService is generated from an XSD element
type GUIDeService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// Medicine is generated from an XSD element
type Medicine struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// Purchase83st544 is generated from an XSD element
type Purchase83st544 struct {
	Medicine Medicine `xml:"medicine,omitempty" bson:"medicine,omitempty"`
}

// Purchase93 is generated from an XSD element
type Purchase93 struct {
	PurchaseAmountLess100      PurchaseAmountLess100      `xml:"purchaseAmountLess100,omitempty" bson:"purchaseAmountLess100,omitempty"`
	PurchaseAmountLess400      PurchaseAmountLess400      `xml:"purchaseAmountLess400,omitempty" bson:"purchaseAmountLess400,omitempty"`
	MaintenanceRepairService   MaintenanceRepairService   `xml:"maintenanceRepairService,omitempty" bson:"maintenanceRepairService,omitempty"`
	BusinessTripService        BusinessTripService        `xml:"businessTripService,omitempty" bson:"businessTripService,omitempty"`
	TeachingService            TeachingService            `xml:"teachingService,omitempty" bson:"teachingService,omitempty"`
	GUIDeService               GUIDeService               `xml:"guideService,omitempty" bson:"guideService,omitempty"`
	CollectionStatisticService CollectionStatisticService `xml:"collectionStatisticService,omitempty" bson:"collectionStatisticService,omitempty"`
	AccessDBService            AccessDBService            `xml:"accessDBService,omitempty" bson:"accessDBService,omitempty"`
}

// PurchaseAmountLess100 is generated from an XSD element
type PurchaseAmountLess100 struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// PurchaseAmountLess400 is generated from an XSD element
type PurchaseAmountLess400 struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// MaintenanceRepairService is generated from an XSD element
type MaintenanceRepairService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// BusinessTripService is generated from an XSD element
type BusinessTripService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// CollectionStatisticService is generated from an XSD element
type CollectionStatisticService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// AccessDBService is generated from an XSD element
type AccessDBService struct {
	Finances ZfcsTenderPlanContextType `xml:"finances,omitempty" bson:"finances,omitempty"`
}

// OutcomeIndicators is generated from an XSD element
type OutcomeIndicators struct {
	SumPushaseSingleSupplier4 string `xml:"sumPushaseSingleSupplier4,omitempty" bson:"sumPushaseSingleSupplier4,omitempty"`
	SumPushaseSingleSupplier5 string `xml:"sumPushaseSingleSupplier5,omitempty" bson:"sumPushaseSingleSupplier5,omitempty"`
	SumPushaseSmallBusiness   string `xml:"sumPushaseSmallBusiness,omitempty" bson:"sumPushaseSmallBusiness,omitempty"`
	SumPushaseRequest         string `xml:"sumPushaseRequest,omitempty" bson:"sumPushaseRequest,omitempty"`
	SumContractMaxPrice       string `xml:"sumContractMaxPrice,omitempty" bson:"sumContractMaxPrice,omitempty"`
	SumPaymentsTotal          string `xml:"sumPaymentsTotal,omitempty" bson:"sumPaymentsTotal,omitempty"`
}

// ZfcsSketchPlanExecutionType is generated from an XSD element
type ZfcsSketchPlanExecutionType struct {
	SchemeVersion string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo                   `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CustomerInfo  ZfcsQuickRefOrganizationType `xml:"customerInfo,omitempty" bson:"customerInfo,omitempty"`
	CreateDate    time.Time                    `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	VersionNumber int64                        `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishDate   time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Description   string                       `xml:"description,omitempty" bson:"description,omitempty"`
	Attachments   ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	PrintForm     ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ProtocolEF1Type is generated from an XSD element
type ProtocolEF1Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsProtocolModificationType is generated from an XSD element
type ZfcsProtocolModificationType struct {
	ModificationNumber int64                              `xml:"modificationNumber,omitempty" bson:"modificationNumber,omitempty"`
	Info               string                             `xml:"info,omitempty" bson:"info,omitempty"`
	AddInfo            string                             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Reason             ZfcsProtocolModificationReasonType `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ProtocolCancelType is generated from an XSD element
type ProtocolCancelType struct {
	NotificationNumber string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber     string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	CreateDate         time.Time        `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PublishDate        time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	ProtocolLots       ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	PrintForm          Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas      DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification       ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href               string           `xml:"href,omitempty" bson:"href,omitempty"`
}

// ProtocolZK1Type is generated from an XSD element
type ProtocolZK1Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsEventObjectType is generated from an XSD element
type ZfcsEventObjectType struct {
	Order    Order    `xml:"order,omitempty" bson:"order,omitempty"`
	Purchase Purchase `xml:"purchase,omitempty" bson:"purchase,omitempty"`
}

// ZfcsETPType is generated from an XSD element
type ZfcsETPType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
	URL  string `xml:"url,omitempty" bson:"url,omitempty"`
}

// ZfcsTenderPlanPositionKVRsType is generated from an XSD element
type ZfcsTenderPlanPositionKVRsType struct {
	KVR []KVR `xml:"KVR,omitempty" bson:"KVR,omitempty"`
}

// ZfcsDeviationFactFoundation is generated from an XSD element
type ZfcsDeviationFactFoundation struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// CommissionMemberType is generated from an XSD element
type CommissionMemberType struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
	Role Role   `xml:"role,omitempty" bson:"role,omitempty"`
}

// ZfcsPublicDiscussionProtocolType is generated from an XSD element
type ZfcsPublicDiscussionProtocolType struct {
	SchemeVersion       string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID          string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID                  int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	PublicDiscussionNum string                               `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	VersionNumber       string                               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate      time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	PublishOrg          ZfcsOrganizationInfoType             `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Topic               string                               `xml:"topic,omitempty" bson:"topic,omitempty"`
	Phase               string                               `xml:"phase,omitempty" bson:"phase,omitempty"`
	Decision            ZfcsPublicDiscussionDecisionRef      `xml:"decision,omitempty" bson:"decision,omitempty"`
	Foundation          ZfcsPublicDiscussionFoundationRef    `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Purchase            ZfcsPublicDiscussionPurchaseInfoType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Attachments         ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// ZfcsProtocolOK1Type is generated from an XSD element
type ZfcsProtocolOK1Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                 `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsPrintFormType is generated from an XSD element
type ZfcsPrintFormType struct {
	URL       string    `xml:"url,omitempty" bson:"url,omitempty"`
	Signature Signature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsStandardContractInvalidType is generated from an XSD element
type ZfcsStandardContractInvalidType struct {
	ID                     int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	DocPublishDate         time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	StandardContractNumber string                       `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	PlacerOrganization     ZfcsPurchaseOrganizationType `xml:"placerOrganization,omitempty" bson:"placerOrganization,omitempty"`
	Type                   string                       `xml:"type,omitempty" bson:"type,omitempty"`
	ApproveInfo            ApproveInfo                  `xml:"approveInfo,omitempty" bson:"approveInfo,omitempty"`
	Indications            Indications                  `xml:"indications,omitempty" bson:"indications,omitempty"`
	UseCases               UseCases                     `xml:"useCases,omitempty" bson:"useCases,omitempty"`
	Okpd2okved2            bool                         `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PrintForm              ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	Attachments            ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	InvalidityInfo         InvalidityInfo               `xml:"invalidityInfo,omitempty" bson:"invalidityInfo,omitempty"`
}

// ZfcsPublicDiscussionType is generated from an XSD element
type ZfcsPublicDiscussionType struct {
	Place  string `xml:"place,omitempty" bson:"place,omitempty"`
	Number string `xml:"number,omitempty" bson:"number,omitempty"`
	Href   string `xml:"href,omitempty" bson:"href,omitempty"`
}

// ZfcsNsiPurchaseRejectReasonType is generated from an XSD element
type ZfcsNsiPurchaseRejectReasonType struct {
	ID            int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Code          string `xml:"code,omitempty" bson:"code,omitempty"`
	Reason        string `xml:"reason,omitempty" bson:"reason,omitempty"`
	Actual        bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
	SubsystemType string `xml:"subsystemType,omitempty" bson:"subsystemType,omitempty"`
}

// ZfcsBankGuaranteeReturnInvalidType is generated from an XSD element
type ZfcsBankGuaranteeReturnInvalidType struct {
	SchemeVersion           string                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                      int64                   `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID              string                  `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ReturnDocNumber         string                  `xml:"returnDocNumber,omitempty" bson:"returnDocNumber,omitempty"`
	DocNumber               string                  `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocPublishDate          time.Time               `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	BankGuaranteeReturnInfo BankGuaranteeReturnInfo `xml:"bankGuaranteeReturnInfo,omitempty" bson:"bankGuaranteeReturnInfo,omitempty"`
	Href                    string                  `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm               ZfcsPrintFormType       `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm            ZfcsExtPrintFormType    `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments             ZfcsAttachmentListType  `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Reason                  string                  `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// BankGuaranteeReturnInfo is generated from an XSD element
type BankGuaranteeReturnInfo struct {
	Bank             ZfcsBankGuaranteeOrganizationType       `xml:"bank,omitempty" bson:"bank,omitempty"`
	SupplierInfo     ZfcsBankGuaranteeParticipantType        `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee        ZfcsBankGuaranteeInfoType               `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	GuaranteeReturns ZfcsContract2015BankGuaranteeReturnType `xml:"guaranteeReturns,omitempty" bson:"guaranteeReturns,omitempty"`
}

// ZfcsPurchaseChangeType is generated from an XSD element
type ZfcsPurchaseChangeType struct {
	ResponsibleDecision   ResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
	CourtDecision         CourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	DiscussionResult      DiscussionResult      `xml:"discussionResult,omitempty" bson:"discussionResult,omitempty"`
}

// ZfcsPurchaseProlongationCommonType is generated from an XSD element
type ZfcsPurchaseProlongationCommonType struct {
	SchemeVersion  string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time            `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	DocType        ZfcsDocType          `xml:"docType,omitempty" bson:"docType,omitempty"`
}

// ZfcsProtocolOKSingleAppType is generated from an XSD element
type ZfcsProtocolOKSingleAppType struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsContractMultiType is generated from an XSD element
type ZfcsContractMultiType struct {
	ContractCount uint64 `xml:"contractCount,omitempty" bson:"contractCount,omitempty"`
}

// ZfcsContractExecutionType is generated from an XSD element
type ZfcsContractExecutionType struct {
	ContractDate  xsd.Date `xml:"contractDate,omitempty" bson:"contractDate,omitempty"`
	ExecutionDate xsd.Date `xml:"executionDate,omitempty" bson:"executionDate,omitempty"`
	Note          string   `xml:"note,omitempty" bson:"note,omitempty"`
}

// ZfcsNsiOKSMType is generated from an XSD element
type ZfcsNsiOKSMType struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
	Actual          bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ProtocolOK1Type is generated from an XSD element
type ProtocolOK1Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// ZfcsAbandonedReasonType is generated from an XSD element
type ZfcsAbandonedReasonType struct {
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	ObjectName string `xml:"objectName,omitempty" bson:"objectName,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Type       string `xml:"type,omitempty" bson:"type,omitempty"`
}

// ZfcsGuaranteeAttachmentType is generated from an XSD element
type ZfcsGuaranteeAttachmentType struct {
	PublishedContentID string      `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string      `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	FileSize           string      `xml:"fileSize,omitempty" bson:"fileSize,omitempty"`
	DocDescription     string      `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	RegDocNumber       string      `xml:"regDocNumber,omitempty" bson:"regDocNumber,omitempty"`
	CryptoSigns        CryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string      `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string      `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte      `xml:"content,omitempty" bson:"content,omitempty"`
}

// CryptoSigns is generated from an XSD element
type CryptoSigns struct {
	Signature []Signature `xml:"signature,omitempty" bson:"signature,omitempty"`
}

// ZfcsEventResultDecisionType is generated from an XSD element
type ZfcsEventResultDecisionType struct {
	DecisionNumber string                 `xml:"decisionNumber,omitempty" bson:"decisionNumber,omitempty"`
	DecisionDate   xsd.Date               `xml:"decisionDate,omitempty" bson:"decisionDate,omitempty"`
	Attachments    ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsCustomerReportSmallScaleBusinessInvalidType is generated from an XSD element
type ZfcsCustomerReportSmallScaleBusinessInvalidType struct {
	SchemeVersion     string                                   `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                int64                                    `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID        string                                   `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate           time.Time                                `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate    time.Time                                `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber     string                                   `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg        ZfcsOrganizationInfoWithOgrnType         `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href              string                                   `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm         ZfcsPrintFormType                        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType                     `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments       ZfcsAttachmentListType                   `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID          string                                   `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	InvalidReportInfo ZfcsInvalidReportType                    `xml:"invalidReportInfo,omitempty" bson:"invalidReportInfo,omitempty"`
	Report            ZfcsCustomerReportSmallScaleBusinessType `xml:"report,omitempty" bson:"report,omitempty"`
}

// ZfcsKladrPlacesType is generated from an XSD element
type ZfcsKladrPlacesType struct {
	KladrPlace []KladrPlace `xml:"kladrPlace,omitempty" bson:"kladrPlace,omitempty"`
}

// KladrPlace is generated from an XSD element
type KladrPlace struct {
	DeliveryPlace              string                     `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	NoKladrForRegionSettlement NoKladrForRegionSettlement `xml:"noKladrForRegionSettlement,omitempty" bson:"noKladrForRegionSettlement,omitempty"`
	Kladr                      ZfcsKladrType              `xml:"kladr,omitempty" bson:"kladr,omitempty"`
	Country                    ZfcsCountryRef             `xml:"country,omitempty" bson:"country,omitempty"`
}

// NoKladrForRegionSettlement is generated from an XSD element
type NoKladrForRegionSettlement struct {
	Region     string `xml:"region,omitempty" bson:"region,omitempty"`
	Settlement string `xml:"settlement,omitempty" bson:"settlement,omitempty"`
}

// ZfcsPurchaseProcedureScoringType is generated from an XSD element
type ZfcsPurchaseProcedureScoringType struct {
	Date    time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place   string    `xml:"place,omitempty" bson:"place,omitempty"`
	AddInfo string    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// OKEIType is generated from an XSD element
type OKEIType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsUnplannedEventCancelType is generated from an XSD element
type ZfcsUnplannedEventCancelType struct {
	SchemeVersion string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	EventNumber   string                 `xml:"eventNumber,omitempty" bson:"eventNumber,omitempty"`
	PublishDate   time.Time              `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Text          string                 `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm     ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsNsiBudgetType is generated from an XSD element
type ZfcsNsiBudgetType struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// NotificationCancelType is generated from an XSD element
type NotificationCancelType struct {
	ID                        int64                     `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber        string                    `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	CreateDate                time.Time                 `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PublishDate               time.Time                 `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	VersionNumber             int64                     `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Lots                      Lots                      `xml:"lots,omitempty" bson:"lots,omitempty"`
	PrintForm                 Document                  `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas             DocumentList              `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification              ModificationType          `xml:"modification,omitempty" bson:"modification,omitempty"`
	NotificationCancelFailure NotificationCancelFailure `xml:"notificationCancelFailure,omitempty" bson:"notificationCancelFailure,omitempty"`
}

// NotificationCancelFailure is generated from an XSD element
type NotificationCancelFailure struct {
	RecoveryToStage RecoveryToStage `xml:"recoveryToStage,omitempty" bson:"recoveryToStage,omitempty"`
	NotRecovery     NotRecovery     `xml:"notRecovery,omitempty" bson:"notRecovery,omitempty"`
}

// RecoveryToStage is generated from an XSD element
type RecoveryToStage struct {
	Stage string `xml:"stage,omitempty" bson:"stage,omitempty"`
}

// NotRecovery is generated from an XSD element
type NotRecovery struct {
	SubstitutedOrder SubstitutedOrder `xml:"substitutedOrder,omitempty" bson:"substitutedOrder,omitempty"`
}

// SubstitutedOrder is generated from an XSD element
type SubstitutedOrder struct {
	NotificationNumber string `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	LotNumber          int64  `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
}

// ZfcsBudgetFundsContract2015 is generated from an XSD element
type ZfcsBudgetFundsContract2015 struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsViolationType is generated from an XSD element
type ZfcsViolationType struct {
	ErrCode     string `xml:"errCode,omitempty" bson:"errCode,omitempty"`
	Level       string `xml:"level,omitempty" bson:"level,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsBankGuaranteeType is generated from an XSD element
type ZfcsBankGuaranteeType struct {
	SchemeVersion      string                            `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                 int64                             `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID         string                            `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNumber          string                            `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	DocNumber          string                            `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	ExtendedDocNumber  string                            `xml:"extendedDocNumber,omitempty" bson:"extendedDocNumber,omitempty"`
	VersionNumber      int64                             `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate     time.Time                         `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Bank               ZfcsBankGuaranteeOrganizationType `xml:"bank,omitempty" bson:"bank,omitempty"`
	PlacingOrg         ZfcsBankGuaranteeOrganizationType `xml:"placingOrg,omitempty" bson:"placingOrg,omitempty"`
	Supplier           ZfcsParticipantType               `xml:"supplier,omitempty" bson:"supplier,omitempty"`
	SupplierInfo       ZfcsBankGuaranteeParticipantType  `xml:"supplierInfo,omitempty" bson:"supplierInfo,omitempty"`
	Guarantee          Guarantee                         `xml:"guarantee,omitempty" bson:"guarantee,omitempty"`
	Href               string                            `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          ZfcsPrintFormType                 `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	AgreementDocuments ZfcsGuaranteeAttachmentListType   `xml:"agreementDocuments,omitempty" bson:"agreementDocuments,omitempty"`
	Modification       Modification                      `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsBigProjectCostType is generated from an XSD element
type ZfcsBigProjectCostType struct {
	Cost string `xml:"cost,omitempty" bson:"cost,omitempty"`
	Year int64  `xml:"year,omitempty" bson:"year,omitempty"`
}

// ZfcsNotificationLotCancelType is generated from an XSD element
type ZfcsNotificationLotCancelType struct {
	SchemeVersion  string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID             int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID     string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber      string                       `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate        time.Time                    `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate time.Time                    `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href           string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm      ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm   ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Lot            Lot                          `xml:"lot,omitempty" bson:"lot,omitempty"`
	CancelReason   ZfcsPurchaseCancelReasonType `xml:"cancelReason,omitempty" bson:"cancelReason,omitempty"`
	AddInfo        string                       `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// CommissionType is generated from an XSD element
type CommissionType struct {
	RegNumber         int64             `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	CommissionName    string            `xml:"commissionName,omitempty" bson:"commissionName,omitempty"`
	CommissionMembers CommissionMembers `xml:"commissionMembers,omitempty" bson:"commissionMembers,omitempty"`
	Owner             OrganizationRef   `xml:"owner,omitempty" bson:"owner,omitempty"`
	Actual            bool              `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ProductType is generated from an XSD element
type ProductType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsDocCancelReasonType is generated from an XSD element
type ZfcsDocCancelReasonType struct {
	ResponsibleDecision   ResponsibleDecision   `xml:"responsibleDecision,omitempty" bson:"responsibleDecision,omitempty"`
	AuthorityPrescription AuthorityPrescription `xml:"authorityPrescription,omitempty" bson:"authorityPrescription,omitempty"`
	CourtDecision         CourtDecision         `xml:"courtDecision,omitempty" bson:"courtDecision,omitempty"`
	DiscussionResult      DiscussionResult      `xml:"discussionResult,omitempty" bson:"discussionResult,omitempty"`
}

// RefusalFact is generated from an XSD element
type RefusalFact struct {
	VoucherEntry string                `xml:"voucherEntry,omitempty" bson:"voucherEntry,omitempty"`
	Explanation  string                `xml:"explanation,omitempty" bson:"explanation,omitempty"`
	Foundation   RefusalFactFoundation `xml:"foundation,omitempty" bson:"foundation,omitempty"`
}

// ZfcsNotificationOrgChangeType is generated from an XSD element
type ZfcsNotificationOrgChangeType struct {
	SchemeVersion   string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID              int64                  `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID      string                 `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocNumber       string                 `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate         time.Time              `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate  time.Time              `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	BaseChange      BaseChange             `xml:"baseChange,omitempty" bson:"baseChange,omitempty"`
	NotifRespOrg    NotifRespOrg           `xml:"notifRespOrg,omitempty" bson:"notifRespOrg,omitempty"`
	Purchase        Purchase               `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	PreviousRespOrg PreviousRespOrg        `xml:"previousRespOrg,omitempty" bson:"previousRespOrg,omitempty"`
	NewRespOrg      NewRespOrg             `xml:"newRespOrg,omitempty" bson:"newRespOrg,omitempty"`
	Href            string                 `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm       ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm    ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments     ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// BaseChange is generated from an XSD element
type BaseChange struct {
	ChangeType string `xml:"changeType,omitempty" bson:"changeType,omitempty"`
	AddInfo    string `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// NotifRespOrg is generated from an XSD element
type NotifRespOrg struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// PreviousRespOrg is generated from an XSD element
type PreviousRespOrg struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// NewRespOrg is generated from an XSD element
type NewRespOrg struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
	SpecialOrg      bool   `xml:"specialOrg,omitempty" bson:"specialOrg,omitempty"`
}

// ZfcsAdmissionResults is generated from an XSD element
type ZfcsAdmissionResults struct {
	AdmissionResult []AdmissionResult `xml:"admissionResult,omitempty" bson:"admissionResult,omitempty"`
}

// ZfcsCurrencyRef is generated from an XSD element
type ZfcsCurrencyRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNsiPublicDiscussionQuestionnarieType is generated from an XSD element
type ZfcsNsiPublicDiscussionQuestionnarieType struct {
	ID        int64     `xml:"id,omitempty" bson:"id,omitempty"`
	Code      string    `xml:"code,omitempty" bson:"code,omitempty"`
	FacetName string    `xml:"facetName,omitempty" bson:"facetName,omitempty"`
	Questions Questions `xml:"questions,omitempty" bson:"questions,omitempty"`
	Type      string    `xml:"type,omitempty" bson:"type,omitempty"`
	Actual    bool      `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Questions is generated from an XSD element
type Questions struct {
	Question []Question `xml:"question,omitempty" bson:"question,omitempty"`
}

// Question is generated from an XSD element
type Question struct {
	ID      int64   `xml:"id,omitempty" bson:"id,omitempty"`
	Code    string  `xml:"code,omitempty" bson:"code,omitempty"`
	Name    string  `xml:"name,omitempty" bson:"name,omitempty"`
	Answers Answers `xml:"answers,omitempty" bson:"answers,omitempty"`
	Actual  bool    `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Answers is generated from an XSD element
type Answers struct {
	Answer []string `xml:"answer,omitempty" bson:"answer,omitempty"`
}

// ProtocolType is generated from an XSD element
type ProtocolType struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
}

// ZfcsNsiUserType is generated from an XSD element
type ZfcsNsiUserType struct {
	Login                string              `xml:"login,omitempty" bson:"login,omitempty"`
	Password             []byte              `xml:"password,omitempty" bson:"password,omitempty"`
	FirstName            string              `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName           string              `xml:"middleName,omitempty" bson:"middleName,omitempty"`
	LastName             string              `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	CodePhrase           string              `xml:"codePhrase,omitempty" bson:"codePhrase,omitempty"`
	Position             string              `xml:"position,omitempty" bson:"position,omitempty"`
	Phone                string              `xml:"phone,omitempty" bson:"phone,omitempty"`
	Email                string              `xml:"email,omitempty" bson:"email,omitempty"`
	Organization         ZfcsOrganizationRef `xml:"organization,omitempty" bson:"organization,omitempty"`
	CertificateSN        string              `xml:"certificateSN,omitempty" bson:"certificateSN,omitempty"`
	CertificateMask      string              `xml:"certificateMask,omitempty" bson:"certificateMask,omitempty"`
	EsIssuerDN           string              `xml:"esIssuerDN,omitempty" bson:"esIssuerDN,omitempty"`
	EsIssuerSN           string              `xml:"esIssuerSN,omitempty" bson:"esIssuerSN,omitempty"`
	UserOrganizationRole string              `xml:"userOrganizationRole,omitempty" bson:"userOrganizationRole,omitempty"`
	Status               string              `xml:"status,omitempty" bson:"status,omitempty"`
	ETPPrivileges        []ZfcsEtpPrivilege  `xml:"ETPPrivileges,omitempty" bson:"ETPPrivileges,omitempty"`
}

// ZfcsAttachmentType is generated from an XSD element
type ZfcsAttachmentType struct {
	PublishedContentID string      `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string      `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	FileSize           string      `xml:"fileSize,omitempty" bson:"fileSize,omitempty"`
	DocDescription     string      `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	CryptoSigns        CryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string      `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string      `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte      `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsContractProcedureCancel2015Type is generated from an XSD element
type ZfcsContractProcedureCancel2015Type struct {
	SchemeVersion        string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CancelledProcedureID int64                `xml:"cancelledProcedureId,omitempty" bson:"cancelledProcedureId,omitempty"`
	RegNum               string               `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	CancelDate           time.Time            `xml:"cancelDate,omitempty" bson:"cancelDate,omitempty"`
	Reason               string               `xml:"reason,omitempty" bson:"reason,omitempty"`
	ExtPrintForm         ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	CurrentContractStage string               `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// ZfcsPurchaseProcedureZakAType is generated from an XSD element
type ZfcsPurchaseProcedureZakAType struct {
	Collecting ZfcsPurchaseProcedureCollectingType `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Opening    ZfcsPurchaseProcedureOpeningType    `xml:"opening,omitempty" bson:"opening,omitempty"`
	Scoring    ZfcsPurchaseProcedureScoringType    `xml:"scoring,omitempty" bson:"scoring,omitempty"`
	Bidding    ZfcsPurchaseProcedureBiddingType    `xml:"bidding,omitempty" bson:"bidding,omitempty"`
}

// ZfcsPurchaseProlongationOKType is generated from an XSD element
type ZfcsPurchaseProlongationOKType struct {
	SchemeVersion           string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                      int64                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID              string               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber          string               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocNumber               string               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	DocDate                 time.Time            `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate          time.Time            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	Href                    string               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm               ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm            ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	DocType                 ZfcsDocType          `xml:"docType,omitempty" bson:"docType,omitempty"`
	Lot                     Lot                  `xml:"lot,omitempty" bson:"lot,omitempty"`
	ScoringDate             time.Time            `xml:"scoringDate,omitempty" bson:"scoringDate,omitempty"`
	ScoringProlongationDate time.Time            `xml:"scoringProlongationDate,omitempty" bson:"scoringProlongationDate,omitempty"`
}

// ZfcsProtocolOKDSingleAppType is generated from an XSD element
type ZfcsProtocolOKDSingleAppType struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsOKSMRef is generated from an XSD element
type ZfcsOKSMRef struct {
	CountryCode     string `xml:"countryCode,omitempty" bson:"countryCode,omitempty"`
	CountryFullName string `xml:"countryFullName,omitempty" bson:"countryFullName,omitempty"`
}

// ZfcsPublicDiscussionCommentType is generated from an XSD element
type ZfcsPublicDiscussionCommentType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	PublicDiscussionNum   string                               `xml:"publicDiscussionNum,omitempty" bson:"publicDiscussionNum,omitempty"`
	VersionNumber         string                               `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	PublicDiscussionFacet ZfcsPublicDiscussionFacetRef         `xml:"publicDiscussionFacet,omitempty" bson:"publicDiscussionFacet,omitempty"`
	Author                Author                               `xml:"author,omitempty" bson:"author,omitempty"`
	CommentNumber         int64                                `xml:"commentNumber,omitempty" bson:"commentNumber,omitempty"`
	Comment               string                               `xml:"comment,omitempty" bson:"comment,omitempty"`
	Phase                 string                               `xml:"phase,omitempty" bson:"phase,omitempty"`
	Purchase              ZfcsPublicDiscussionPurchaseInfoType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsRegulationRulesType is generated from an XSD element
type ZfcsRegulationRulesType struct {
	ID                 int64              `xml:"id,omitempty" bson:"id,omitempty"`
	DocPublishDate     time.Time          `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	RegistryNum        string             `xml:"registryNum,omitempty" bson:"registryNum,omitempty"`
	PublishOrg         PublishOrg         `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Type               string             `xml:"type,omitempty" bson:"type,omitempty"`
	State              string             `xml:"state,omitempty" bson:"state,omitempty"`
	TermsControl       bool               `xml:"termsControl,omitempty" bson:"termsControl,omitempty"`
	ApprovedFrom       ApprovedFrom       `xml:"approvedFrom,omitempty" bson:"approvedFrom,omitempty"`
	ApproveFor         ApproveFor         `xml:"approveFor,omitempty" bson:"approveFor,omitempty"`
	BaseDocument       BaseDocument       `xml:"baseDocument,omitempty" bson:"baseDocument,omitempty"`
	Regions            Regions            `xml:"regions,omitempty" bson:"regions,omitempty"`
	AddInfo            string             `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
	Discussion         Discussion         `xml:"discussion,omitempty" bson:"discussion,omitempty"`
	Documents          Documents          `xml:"documents,omitempty" bson:"documents,omitempty"`
	PrintFormDocuments PrintFormDocuments `xml:"printFormDocuments,omitempty" bson:"printFormDocuments,omitempty"`
	Modification       Modification       `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsPenaltyDocumentInfoList is generated from an XSD element
type ZfcsPenaltyDocumentInfoList struct {
	DocumentInfo []DocumentInfo `xml:"documentInfo,omitempty" bson:"documentInfo,omitempty"`
}

// DocumentInfo is generated from an XSD element
type DocumentInfo struct {
	DocumentDate time.Time `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	DocumentName string    `xml:"documentName,omitempty" bson:"documentName,omitempty"`
}

// ZfcsBigProjectFinancingYearsType is generated from an XSD element
type ZfcsBigProjectFinancingYearsType struct {
	Total ZfcsBigProjectFinancingsType `xml:"total,omitempty" bson:"total,omitempty"`
	Year  []Year                       `xml:"year,omitempty" bson:"year,omitempty"`
}

// ZfcsNsiContractModificationReasonType is generated from an XSD element
type ZfcsNsiContractModificationReasonType struct {
	Code      string    `xml:"code,omitempty" bson:"code,omitempty"`
	Name      string    `xml:"name,omitempty" bson:"name,omitempty"`
	Actual    bool      `xml:"actual,omitempty" bson:"actual,omitempty"`
	Documents Documents `xml:"documents,omitempty" bson:"documents,omitempty"`
}

// NotificationPlacerChangeType is generated from an XSD element
type NotificationPlacerChangeType struct {
	ID                 int64           `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber string          `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	VersionNumber      int64           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	CreateDate         time.Time       `xml:"createDate,omitempty" bson:"createDate,omitempty"`
	PublishDate        time.Time       `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	ChangeType         string          `xml:"changeType,omitempty" bson:"changeType,omitempty"`
	PlacerChange       PlacerChange    `xml:"placerChange,omitempty" bson:"placerChange,omitempty"`
	InitiatorChange    InitiatorChange `xml:"initiatorChange,omitempty" bson:"initiatorChange,omitempty"`
	Comment            string          `xml:"comment,omitempty" bson:"comment,omitempty"`
	Href               string          `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm          Document        `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas      DocumentList    `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// PlacerChange is generated from an XSD element
type PlacerChange struct {
	CurrentPlacer OrganizationRef `xml:"currentPlacer,omitempty" bson:"currentPlacer,omitempty"`
	NewPlacer     OrganizationRef `xml:"newPlacer,omitempty" bson:"newPlacer,omitempty"`
}

// InitiatorChange is generated from an XSD element
type InitiatorChange struct {
	CurrentInitiator OrganizationRef `xml:"currentInitiator,omitempty" bson:"currentInitiator,omitempty"`
	NewInitiator     OrganizationRef `xml:"newInitiator,omitempty" bson:"newInitiator,omitempty"`
}

// ZfcsNsiCommissionRoleType is generated from an XSD element
type ZfcsNsiCommissionRoleType struct {
	ID        int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name      string `xml:"name,omitempty" bson:"name,omitempty"`
	Order     int64  `xml:"order,omitempty" bson:"order,omitempty"`
	RightVote bool   `xml:"rightVote,omitempty" bson:"rightVote,omitempty"`
	Actual    bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsNsiOKPDType is generated from an XSD element
type ZfcsNsiOKPDType struct {
	ID         int64  `xml:"id,omitempty" bson:"id,omitempty"`
	ParentID   int64  `xml:"parentId,omitempty" bson:"parentId,omitempty"`
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	ParentCode string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Comment    string `xml:"comment,omitempty" bson:"comment,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// Account is generated from an XSD element
type Account struct {
	BankAddress     string `xml:"bankAddress,omitempty" bson:"bankAddress,omitempty"`
	BankName        string `xml:"bankName,omitempty" bson:"bankName,omitempty"`
	Bik             string `xml:"bik,omitempty" bson:"bik,omitempty"`
	CorrAccount     string `xml:"corrAccount,omitempty" bson:"corrAccount,omitempty"`
	PaymentAccount  string `xml:"paymentAccount,omitempty" bson:"paymentAccount,omitempty"`
	PersonalAccount string `xml:"personalAccount,omitempty" bson:"personalAccount,omitempty"`
}

// ZfcsCheckResultActType is generated from an XSD element
type ZfcsCheckResultActType struct {
	ActNumber   string                 `xml:"actNumber,omitempty" bson:"actNumber,omitempty"`
	ActDate     xsd.Date               `xml:"actDate,omitempty" bson:"actDate,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsContract2015PayDocInfo is generated from an XSD element
type ZfcsContract2015PayDocInfo struct {
	DocumentName string   `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentNum  string   `xml:"documentNum,omitempty" bson:"documentNum,omitempty"`
	DocumentDate xsd.Date `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	Amount       string   `xml:"amount,omitempty" bson:"amount,omitempty"`
	AmountRUR    string   `xml:"amountRUR,omitempty" bson:"amountRUR,omitempty"`
}

// ZfcsContractProcedureCancelType is generated from an XSD element
type ZfcsContractProcedureCancelType struct {
	SchemeVersion        string    `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CancelledProcedureID int64     `xml:"cancelledProcedureId,omitempty" bson:"cancelledProcedureId,omitempty"`
	RegNum               string    `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	CancelDate           time.Time `xml:"cancelDate,omitempty" bson:"cancelDate,omitempty"`
	Reason               string    `xml:"reason,omitempty" bson:"reason,omitempty"`
	CurrentContractStage string    `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// ZfcsAttachmentListType is generated from an XSD element
type ZfcsAttachmentListType struct {
	Attachment []ZfcsAttachmentType `xml:"attachment,omitempty" bson:"attachment,omitempty"`
}

// ZfcsContractAttachmentType is generated from an XSD element
type ZfcsContractAttachmentType struct {
	PublishedContentID string      `xml:"publishedContentId,omitempty" bson:"publishedContentId,omitempty"`
	FileName           string      `xml:"fileName,omitempty" bson:"fileName,omitempty"`
	DocDescription     string      `xml:"docDescription,omitempty" bson:"docDescription,omitempty"`
	DocRegNumber       string      `xml:"docRegNumber,omitempty" bson:"docRegNumber,omitempty"`
	CryptoSigns        CryptoSigns `xml:"cryptoSigns,omitempty" bson:"cryptoSigns,omitempty"`
	URL                string      `xml:"url,omitempty" bson:"url,omitempty"`
	ContentID          string      `xml:"contentId,omitempty" bson:"contentId,omitempty"`
	Content            []byte      `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsNsiOKVED2Type is generated from an XSD element
type ZfcsNsiOKVED2Type struct {
	ID         int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	Section    string `xml:"section,omitempty" bson:"section,omitempty"`
	ParentCode string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Comment    string `xml:"comment,omitempty" bson:"comment,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// EPType is generated from an XSD element
type EPType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
	URL  string `xml:"url,omitempty" bson:"url,omitempty"`
}

// ZfcsContractType is generated from an XSD element
type ZfcsContractType struct {
	SchemeVersion        string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                   int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID           string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	RegNum               string                         `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	Number               string                         `xml:"number,omitempty" bson:"number,omitempty"`
	PublishDate          time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	SignDate             xsd.Date                       `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	VersionNumber        int64                          `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Foundation           Foundation                     `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Customer             Customer                       `xml:"customer,omitempty" bson:"customer,omitempty"`
	ProtocolDate         xsd.Date                       `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	DocumentBase         string                         `xml:"documentBase,omitempty" bson:"documentBase,omitempty"`
	Price                string                         `xml:"price,omitempty" bson:"price,omitempty"`
	Currency             ZfcsCurrencyRef                `xml:"currency,omitempty" bson:"currency,omitempty"`
	SingleCustomerReason SingleCustomerReason           `xml:"singleCustomerReason,omitempty" bson:"singleCustomerReason,omitempty"`
	PriceChangeReason    PriceChangeReason              `xml:"priceChangeReason,omitempty" bson:"priceChangeReason,omitempty"`
	ExecutionDate        ZfcsStageType                  `xml:"executionDate,omitempty" bson:"executionDate,omitempty"`
	Finances             Finances                       `xml:"finances,omitempty" bson:"finances,omitempty"`
	Products             Products                       `xml:"products,omitempty" bson:"products,omitempty"`
	Suppliers            Suppliers                      `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	Href                 string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm            ZfcsContractPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ScanDocuments        ZfcsContractAttachmentListType `xml:"scanDocuments,omitempty" bson:"scanDocuments,omitempty"`
	MedicalDocuments     ZfcsContractAttachmentListType `xml:"medicalDocuments,omitempty" bson:"medicalDocuments,omitempty"`
	Attachments          ZfcsContractAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification         Modification                   `xml:"modification,omitempty" bson:"modification,omitempty"`
	CurrentContractStage string                         `xml:"currentContractStage,omitempty" bson:"currentContractStage,omitempty"`
}

// SingleCustomerReason is generated from an XSD element
type SingleCustomerReason struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// PriceChangeReason is generated from an XSD element
type PriceChangeReason struct {
	ID      int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name    string `xml:"name,omitempty" bson:"name,omitempty"`
	Comment string `xml:"comment,omitempty" bson:"comment,omitempty"`
}

// ZfcsPurchaseProcedureCollectingWithFormType is generated from an XSD element
type ZfcsPurchaseProcedureCollectingWithFormType struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	Place     string    `xml:"place,omitempty" bson:"place,omitempty"`
	Order     string    `xml:"order,omitempty" bson:"order,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	Form      string    `xml:"form,omitempty" bson:"form,omitempty"`
}

// CurrencyType is generated from an XSD element
type CurrencyType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsUnplannedCheckType is generated from an XSD element
type ZfcsUnplannedCheckType struct {
	SchemeVersion   string                              `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo      CommonInfo                          `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Period          ZfcsPeriodType                      `xml:"period,omitempty" bson:"period,omitempty"`
	Inspector       ZfcsOrganizationRef                 `xml:"inspector,omitempty" bson:"inspector,omitempty"`
	InspectionDate  time.Time                           `xml:"inspectionDate,omitempty" bson:"inspectionDate,omitempty"`
	InspectionPlace string                              `xml:"inspectionPlace,omitempty" bson:"inspectionPlace,omitempty"`
	CheckedSubject  []ZfcsUnplannedCheckSubjectPlanType `xml:"checkedSubject,omitempty" bson:"checkedSubject,omitempty"`
	Base            Base                                `xml:"base,omitempty" bson:"base,omitempty"`
	CheckedObject   CheckedObject                       `xml:"checkedObject,omitempty" bson:"checkedObject,omitempty"`
	PrintForm       ZfcsPrintFormType                   `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm    ZfcsExtPrintFormType                `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments     ZfcsAttachmentListType              `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsBigProjectValueType is generated from an XSD element
type ZfcsBigProjectValueType struct {
	Cost       string  `xml:"cost,omitempty" bson:"cost,omitempty"`
	Percentage float64 `xml:"percentage,omitempty" bson:"percentage,omitempty"`
}

// ZfcsProtocolEFSinglePartType is generated from an XSD element
type ZfcsProtocolEFSinglePartType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
	AbandonedReason          ZfcsAbandonedReasonType      `xml:"abandonedReason,omitempty" bson:"abandonedReason,omitempty"`
}

// ZfcsPurchaseProcedureContractingType is generated from an XSD element
type ZfcsPurchaseProcedureContractingType struct {
	ContractingTerm string `xml:"contractingTerm,omitempty" bson:"contractingTerm,omitempty"`
	EvadeConditions string `xml:"evadeConditions,omitempty" bson:"evadeConditions,omitempty"`
}

// ZfcsProtocolEF1Type is generated from an XSD element
type ZfcsProtocolEF1Type struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsNsiCalendarDaysType is generated from an XSD element
type ZfcsNsiCalendarDaysType struct {
	StartDate time.Time `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate   time.Time `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	Days      Days      `xml:"days,omitempty" bson:"days,omitempty"`
}

// Days is generated from an XSD element
type Days struct {
	Day []Day `xml:"day,omitempty" bson:"day,omitempty"`
}

// Day is generated from an XSD element
type Day struct {
	DayDate   time.Time `xml:"dayDate,omitempty" bson:"dayDate,omitempty"`
	Regions   Regions   `xml:"regions,omitempty" bson:"regions,omitempty"`
	IsWorkDay bool      `xml:"isWorkDay,omitempty" bson:"isWorkDay,omitempty"`
}

// ZfcsPositionKBK2016sYearsType is generated from an XSD element
type ZfcsPositionKBK2016sYearsType struct {
	KBK2016 []KBK2016 `xml:"KBK2016,omitempty" bson:"KBK2016,omitempty"`
}

// ZfcsNotificationPOType is generated from an XSD element
type ZfcsNotificationPOType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ContractServiceInfo string                           `xml:"contractServiceInfo,omitempty" bson:"contractServiceInfo,omitempty"`
	ProcedureInfo       ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                 Lot                              `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsApplicationCorrespondence is generated from an XSD element
type ZfcsApplicationCorrespondence struct {
	Compatible   bool                `xml:"compatible,omitempty" bson:"compatible,omitempty"`
	OverallValue float64             `xml:"overallValue,omitempty" bson:"overallValue,omitempty"`
	Preferense   ZfcsPreferenseType  `xml:"preferense,omitempty" bson:"preferense,omitempty"`
	Requirement  ZfcsRequirementType `xml:"requirement,omitempty" bson:"requirement,omitempty"`
}

// ZfcsProtocolOKOU2Type is generated from an XSD element
type ZfcsProtocolOKOU2Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsNsiContractExecutionDocType is generated from an XSD element
type ZfcsNsiContractExecutionDocType struct {
	Code   string `xml:"code,omitempty" bson:"code,omitempty"`
	Name   string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// KBKType is generated from an XSD element
type KBKType struct {
	Kbk1 string `xml:"kbk1,omitempty" bson:"kbk1,omitempty"`
	Kbk2 string `xml:"kbk2,omitempty" bson:"kbk2,omitempty"`
	Kbk3 string `xml:"kbk3,omitempty" bson:"kbk3,omitempty"`
	Kbk4 string `xml:"kbk4,omitempty" bson:"kbk4,omitempty"`
	Kbk5 string `xml:"kbk5,omitempty" bson:"kbk5,omitempty"`
}

// ZfcsNotificationZakAType is generated from an XSD element
type ZfcsNotificationZakAType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureZakAType        `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsPurchaseOrganizationType is generated from an XSD element
type ZfcsPurchaseOrganizationType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	PostAddress     string `xml:"postAddress,omitempty" bson:"postAddress,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
}

// ZfcsPurchaseProcedureOKOUType is generated from an XSD element
type ZfcsPurchaseProcedureOKOUType struct {
	Collecting       ZfcsPurchaseProcedureCollectingType       `xml:"collecting,omitempty" bson:"collecting,omitempty"`
	Opening          ZfcsPurchaseProcedureOpeningType          `xml:"opening,omitempty" bson:"opening,omitempty"`
	Prequalification ZfcsPurchaseProcedurePrequalificationType `xml:"prequalification,omitempty" bson:"prequalification,omitempty"`
	Scoring          ZfcsPurchaseProcedureScoringType          `xml:"scoring,omitempty" bson:"scoring,omitempty"`
}

// ZfcsPurchaseProcedureBiddingType is generated from an XSD element
type ZfcsPurchaseProcedureBiddingType struct {
	Date    time.Time `xml:"date,omitempty" bson:"date,omitempty"`
	Place   string    `xml:"place,omitempty" bson:"place,omitempty"`
	AddInfo string    `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ChildrenCriteriaType is generated from an XSD element
type ChildrenCriteriaType struct {
	ID             int64   `xml:"id,omitempty" bson:"id,omitempty"`
	Name           string  `xml:"name,omitempty" bson:"name,omitempty"`
	CriterionValue float64 `xml:"criterionValue,omitempty" bson:"criterionValue,omitempty"`
	EvalValue      string  `xml:"evalValue,omitempty" bson:"evalValue,omitempty"`
}

// ZfcsEventSubjectType is generated from an XSD element
type ZfcsEventSubjectType struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	OOSAuthority           ZfcsOrganizationRef    `xml:"OOS_authority,omitempty" bson:"OOS_authority,omitempty"`
	RCAuthority            ZfcsOrganizationRef    `xml:"RC_authority,omitempty" bson:"RC_authority,omitempty"`
	FCAuthority            ZfcsOrganizationRef    `xml:"FC_authority,omitempty" bson:"FC_authority,omitempty"`
	LegalEntity44          ZfcsOrganizationRef    `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         ZfcsOrganizationRef    `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
	Other                  ZfcsOrganizationRef    `xml:"other,omitempty" bson:"other,omitempty"`
}

// ProtocolOK2Type is generated from an XSD element
type ProtocolOK2Type struct {
	ID                       int64            `xml:"id,omitempty" bson:"id,omitempty"`
	NotificationNumber       string           `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	ProtocolNumber           string           `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string           `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string           `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	VersionNumber            int64            `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	Place                    string           `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time        `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time        `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time        `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               Commission       `xml:"commission,omitempty" bson:"commission,omitempty"`
	PrintForm                Document         `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	DocumentMetas            DocumentList     `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
	Modification             ModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	Href                     string           `xml:"href,omitempty" bson:"href,omitempty"`
	ProtocolLots             ProtocolLots     `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
}

// RequirementCompliances is generated from an XSD element
type RequirementCompliances struct {
	RequirementCompliance []RequirementCompliance `xml:"requirementCompliance,omitempty" bson:"requirementCompliance,omitempty"`
}

// RequirementCompliance is generated from an XSD element
type RequirementCompliance struct {
	OrdinalNumber    int64  `xml:"ordinalNumber,omitempty" bson:"ordinalNumber,omitempty"`
	AvailabilityType string `xml:"availabilityType,omitempty" bson:"availabilityType,omitempty"`
	Reason           string `xml:"reason,omitempty" bson:"reason,omitempty"`
}

// ZfcsOrganizationInfoWithOgrnType is generated from an XSD element
type ZfcsOrganizationInfoWithOgrnType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	ShortName       string `xml:"shortName,omitempty" bson:"shortName,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	OGRN            string `xml:"OGRN,omitempty" bson:"OGRN,omitempty"`
	ResponsibleRole string `xml:"responsibleRole,omitempty" bson:"responsibleRole,omitempty"`
}

// ZfcsBudgetFinancingsType is generated from an XSD element
type ZfcsBudgetFinancingsType struct {
	BudgetFinancing []BudgetFinancing `xml:"budgetFinancing,omitempty" bson:"budgetFinancing,omitempty"`
	TotalSum        string            `xml:"totalSum,omitempty" bson:"totalSum,omitempty"`
}

// BudgetFinancing is generated from an XSD element
type BudgetFinancing struct {
	Year        int64  `xml:"year,omitempty" bson:"year,omitempty"`
	Sum         string `xml:"sum,omitempty" bson:"sum,omitempty"`
	KbkCode     string `xml:"kbkCode,omitempty" bson:"kbkCode,omitempty"`
	KbkCode2016 string `xml:"kbkCode2016,omitempty" bson:"kbkCode2016,omitempty"`
}

// ZfcsBankGuaranteeInfoType is generated from an XSD element
type ZfcsBankGuaranteeInfoType struct {
	Customer                ZfcsBankGuaranteeOrganizationType `xml:"customer,omitempty" bson:"customer,omitempty"`
	GuaranteeDate           time.Time                         `xml:"guaranteeDate,omitempty" bson:"guaranteeDate,omitempty"`
	GuaranteeGrantDate      time.Time                         `xml:"guaranteeGrantDate,omitempty" bson:"guaranteeGrantDate,omitempty"`
	GuaranteePublishDate    time.Time                         `xml:"guaranteePublishDate,omitempty" bson:"guaranteePublishDate,omitempty"`
	GuaranteeNumber         string                            `xml:"guaranteeNumber,omitempty" bson:"guaranteeNumber,omitempty"`
	ExpireDate              time.Time                         `xml:"expireDate,omitempty" bson:"expireDate,omitempty"`
	GuaranteeAmount         string                            `xml:"guaranteeAmount,omitempty" bson:"guaranteeAmount,omitempty"`
	Currency                ZfcsCurrencyFullRef               `xml:"currency,omitempty" bson:"currency,omitempty"`
	EntryForceDate          time.Time                         `xml:"entryForceDate,omitempty" bson:"entryForceDate,omitempty"`
	Procedure               string                            `xml:"procedure,omitempty" bson:"procedure,omitempty"`
	GuaranteeAmountRUR      string                            `xml:"guaranteeAmountRUR,omitempty" bson:"guaranteeAmountRUR,omitempty"`
	CurrencyRate            float64                           `xml:"currencyRate,omitempty" bson:"currencyRate,omitempty"`
	PurchaseRequestEnsure   PurchaseRequestEnsure             `xml:"purchaseRequestEnsure,omitempty" bson:"purchaseRequestEnsure,omitempty"`
	ContractExecutionEnsure ContractExecutionEnsure           `xml:"contractExecutionEnsure,omitempty" bson:"contractExecutionEnsure,omitempty"`
}

// PurchaseRequestEnsure is generated from an XSD element
type PurchaseRequestEnsure struct {
	LotNumber          int64  `xml:"lotNumber,omitempty" bson:"lotNumber,omitempty"`
	MLots              bool   `xml:"mLots,omitempty" bson:"mLots,omitempty"`
	PurchaseNumber     string `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	NotificationNumber string `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
}

// ContractExecutionEnsure is generated from an XSD element
type ContractExecutionEnsure struct {
	RegNum         string   `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	Purchase       Purchase `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	MLots          bool     `xml:"mLots,omitempty" bson:"mLots,omitempty"`
	SingleSupplier bool     `xml:"singleSupplier,omitempty" bson:"singleSupplier,omitempty"`
}

// ZfcsContractSignType is generated from an XSD element
type ZfcsContractSignType struct {
	SchemeVersion         string              `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64               `xml:"id,omitempty" bson:"id,omitempty"`
	Number                string              `xml:"number,omitempty" bson:"number,omitempty"`
	SignDate              xsd.Date            `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	Foundation            Foundation          `xml:"foundation,omitempty" bson:"foundation,omitempty"`
	Customer              ZfcsOrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
	Price                 string              `xml:"price,omitempty" bson:"price,omitempty"`
	PriceRUR              string              `xml:"priceRUR,omitempty" bson:"priceRUR,omitempty"`
	Currency              ZfcsCurrencyRef     `xml:"currency,omitempty" bson:"currency,omitempty"`
	ConcludeContractRight bool                `xml:"concludeContractRight,omitempty" bson:"concludeContractRight,omitempty"`
	ProtocolDate          xsd.Date            `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	Suppliers             Suppliers           `xml:"suppliers,omitempty" bson:"suppliers,omitempty"`
	Document              Document            `xml:"document,omitempty" bson:"document,omitempty"`
}

// ZfcsTenderPlanPositionKOSGUsType is generated from an XSD element
type ZfcsTenderPlanPositionKOSGUsType struct {
	KOSGU []KOSGU `xml:"KOSGU,omitempty" bson:"KOSGU,omitempty"`
}

// ZfcsNotificationZPType is generated from an XSD element
type ZfcsNotificationZPType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ProcedureInfo                        `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                   Lot                                  `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsProtocolZPType is generated from an XSD element
type ZfcsProtocolZPType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsAuditActionSubjectsRef is generated from an XSD element
type ZfcsAuditActionSubjectsRef struct {
	ID   int64  `xml:"id,omitempty" bson:"id,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsTimelineViolationType is generated from an XSD element
type ZfcsTimelineViolationType struct {
	SchemeVersion string `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	RefID         string `xml:"refId,omitempty" bson:"refId,omitempty"`
}

// ZfcsImproperContractExecutionType is generated from an XSD element
type ZfcsImproperContractExecutionType struct {
	NameObligations  string `xml:"nameObligations,omitempty" bson:"nameObligations,omitempty"`
	EssenceViolation string `xml:"essenceViolation,omitempty" bson:"essenceViolation,omitempty"`
	PenaltyInfo      string `xml:"penaltyInfo,omitempty" bson:"penaltyInfo,omitempty"`
	PenaltyDoc       string `xml:"penaltyDoc,omitempty" bson:"penaltyDoc,omitempty"`
	Note             string `xml:"note,omitempty" bson:"note,omitempty"`
}

// ZfcsProtocolOKD5Type is generated from an XSD element
type ZfcsProtocolOKD5Type struct {
	SchemeVersion            string                         `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                          `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                         `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                         `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                         `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                         `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                         `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                         `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                      `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                      `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                      `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType             `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                         `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType              `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType           `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher              `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                   `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType         `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType   `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLots             ProtocolLots                   `xml:"protocolLots,omitempty" bson:"protocolLots,omitempty"`
	FoundationProtocol       ZfcsFoundationProtocolInfoType `xml:"foundationProtocol,omitempty" bson:"foundationProtocol,omitempty"`
}

// ZfcsPPORef is generated from an XSD element
type ZfcsPPORef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsNsiKOSGUType is generated from an XSD element
type ZfcsNsiKOSGUType struct {
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	ParentCode string `xml:"parentCode,omitempty" bson:"parentCode,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsTenderPlanTotalPositionKVRsType is generated from an XSD element
type ZfcsTenderPlanTotalPositionKVRsType struct {
	KVR []KVR `xml:"KVR,omitempty" bson:"KVR,omitempty"`
}

// ZfcsCheckResultType is generated from an XSD element
type ZfcsCheckResultType struct {
	SchemeVersion string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo           `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	StartDate     xsd.Date             `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate       xsd.Date             `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Base          Base                 `xml:"base,omitempty" bson:"base,omitempty"`
	Complaint     Complaint            `xml:"complaint,omitempty" bson:"complaint,omitempty"`
}

// Complaint is generated from an XSD element
type Complaint struct {
	ComplaintNumber string                          `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	PublishDate     time.Time                       `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	CheckSubjects   []CheckSubjects                 `xml:"checkSubjects,omitempty" bson:"checkSubjects,omitempty"`
	CheckedObject   []ZfcsComplaintObjectType       `xml:"checkedObject,omitempty" bson:"checkedObject,omitempty"`
	Decision        ZfcsCheckResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	Prescription    ZfcsCheckResultPrescriptionType `xml:"prescription,omitempty" bson:"prescription,omitempty"`
}

// CheckSubjects is generated from an XSD element
type CheckSubjects struct {
	SubjectComplaint      SubjectComplaint      `xml:"subjectComplaint,omitempty" bson:"subjectComplaint,omitempty"`
	SubjectComplaintGroup SubjectComplaintGroup `xml:"subjectComplaintGroup,omitempty" bson:"subjectComplaintGroup,omitempty"`
}

// SubjectComplaint is generated from an XSD element
type SubjectComplaint struct {
	Customer               ZfcsOrganizationRef    `xml:"customer,omitempty" bson:"customer,omitempty"`
	Authority              ZfcsOrganizationRef    `xml:"authority,omitempty" bson:"authority,omitempty"`
	AuthorityAgency        ZfcsOrganizationRef    `xml:"authorityAgency,omitempty" bson:"authorityAgency,omitempty"`
	Specialized            ZfcsOrganizationRef    `xml:"specialized,omitempty" bson:"specialized,omitempty"`
	EP                     ZfcsOrganizationRef    `xml:"EP,omitempty" bson:"EP,omitempty"`
	EPRefuse               ZfcsOrganizationRef    `xml:"EP_refuse,omitempty" bson:"EP_refuse,omitempty"`
	Commission44           Commission44           `xml:"commission44,omitempty" bson:"commission44,omitempty"`
	Commission94           Commission94           `xml:"commission94,omitempty" bson:"commission94,omitempty"`
	ContractServiceOfficer ContractServiceOfficer `xml:"contractServiceOfficer,omitempty" bson:"contractServiceOfficer,omitempty"`
	ContractService        ContractService        `xml:"contractService,omitempty" bson:"contractService,omitempty"`
	LegalEntity44          LegalEntity44          `xml:"legalEntity44,omitempty" bson:"legalEntity44,omitempty"`
	LegalEntity307         LegalEntity307         `xml:"legalEntity307,omitempty" bson:"legalEntity307,omitempty"`
}

// SubjectComplaintGroup is generated from an XSD element
type SubjectComplaintGroup struct {
	EPFailure ZfcsOrganizationRef `xml:"EP_failure,omitempty" bson:"EP_failure,omitempty"`
}

// ZfcsNotificationZakKDType is generated from an XSD element
type ZfcsNotificationZakKDType struct {
	SchemeVersion         string                               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                    int64                                `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID            string                               `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber        string                               `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate        time.Time                            `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber             string                               `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                  string                               `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm             ZfcsPrintFormType                    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm          ZfcsExtPrintFormType                 `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo    string                               `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible   PurchaseResponsible                  `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay            ZfcsPlacingWayType                   `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2           bool                                 `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	PurchaseDocumentation ZfcsReleasePurchaseDocumentationType `xml:"purchaseDocumentation,omitempty" bson:"purchaseDocumentation,omitempty"`
	ProcedureInfo         ZfcsPurchaseProcedureOKDType         `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lots                  Lots                                 `xml:"lots,omitempty" bson:"lots,omitempty"`
	Attachments           ZfcsAttachmentListType               `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification          ZfcsNotificationModificationType     `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsEventPlanType is generated from an XSD element
type ZfcsEventPlanType struct {
	SchemeVersion string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo           `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	StartStage    ZfcsStageType        `xml:"startStage,omitempty" bson:"startStage,omitempty"`
	EndStage      ZfcsStageType        `xml:"endStage,omitempty" bson:"endStage,omitempty"`
	EventList     EventList            `xml:"eventList,omitempty" bson:"eventList,omitempty"`
	PrintForm     ZfcsPrintFormType    `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
}

// EventList is generated from an XSD element
type EventList struct {
	EventInfo []EventInfo `xml:"eventInfo,omitempty" bson:"eventInfo,omitempty"`
}

// EventInfo is generated from an XSD element
type EventInfo struct {
	EventNumber           string               `xml:"eventNumber,omitempty" bson:"eventNumber,omitempty"`
	EventType             string               `xml:"eventType,omitempty" bson:"eventType,omitempty"`
	InspectionType        string               `xml:"inspectionType,omitempty" bson:"inspectionType,omitempty"`
	SurveyType            string               `xml:"surveyType,omitempty" bson:"surveyType,omitempty"`
	EventSubject          ZfcsEventSubjectType `xml:"eventSubject,omitempty" bson:"eventSubject,omitempty"`
	Period                ZfcsPeriodType       `xml:"period,omitempty" bson:"period,omitempty"`
	EventStartStage       ZfcsStageType        `xml:"eventStartStage,omitempty" bson:"eventStartStage,omitempty"`
	EventPublishDate      time.Time            `xml:"eventPublishDate,omitempty" bson:"eventPublishDate,omitempty"`
	Base                  string               `xml:"base,omitempty" bson:"base,omitempty"`
	ParentEventPlanNumber string               `xml:"parentEventPlanNumber,omitempty" bson:"parentEventPlanNumber,omitempty"`
}

// ZfcsNsiCurrencyType is generated from an XSD element
type ZfcsNsiCurrencyType struct {
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	DigitalCode string `xml:"digitalCode,omitempty" bson:"digitalCode,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual      bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsRequirementType is generated from an XSD element
type ZfcsRequirementType struct {
	Code    int64  `xml:"code,omitempty" bson:"code,omitempty"`
	Name    string `xml:"name,omitempty" bson:"name,omitempty"`
	Content string `xml:"content,omitempty" bson:"content,omitempty"`
}

// ZfcsNsiBusinessControlsType is generated from an XSD element
type ZfcsNsiBusinessControlsType struct {
	NsiBusinessControl []ZfcsNsiBusinessControlType `xml:"nsiBusinessControl,omitempty" bson:"nsiBusinessControl,omitempty"`
}

// ZfcsEventResultType is generated from an XSD element
type ZfcsEventResultType struct {
	SchemeVersion string               `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo           `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	StartDate     xsd.Date             `xml:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate       xsd.Date             `xml:"endDate,omitempty" bson:"endDate,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	TypeEvent     TypeEvent            `xml:"typeEvent,omitempty" bson:"typeEvent,omitempty"`
	Complaint     Complaint            `xml:"complaint,omitempty" bson:"complaint,omitempty"`
}

// TypeEvent is generated from an XSD element
type TypeEvent struct {
	UnplannedCheckComplaint UnplannedCheckComplaint `xml:"unplannedCheckComplaint,omitempty" bson:"unplannedCheckComplaint,omitempty"`
	UnplannedCheck          UnplannedCheck          `xml:"unplannedCheck,omitempty" bson:"unplannedCheck,omitempty"`
	UnplannedRevision       UnplannedRevision       `xml:"unplannedRevision,omitempty" bson:"unplannedRevision,omitempty"`
	UnplannedSurvey         UnplannedSurvey         `xml:"unplannedSurvey,omitempty" bson:"unplannedSurvey,omitempty"`
	PlannedCheck            PlannedCheck            `xml:"plannedCheck,omitempty" bson:"plannedCheck,omitempty"`
	PlannedRevision         PlannedRevision         `xml:"plannedRevision,omitempty" bson:"plannedRevision,omitempty"`
	PlannedSurvey           PlannedSurvey           `xml:"plannedSurvey,omitempty" bson:"plannedSurvey,omitempty"`
}

// UnplannedCheckComplaint is generated from an XSD element
type UnplannedCheckComplaint struct {
	ComplaintNumber string                          `xml:"complaintNumber,omitempty" bson:"complaintNumber,omitempty"`
	CheckSubjects   []CheckSubjects                 `xml:"checkSubjects,omitempty" bson:"checkSubjects,omitempty"`
	CheckedObject   []ZfcsComplaintObjectType       `xml:"checkedObject,omitempty" bson:"checkedObject,omitempty"`
	Info            string                          `xml:"info,omitempty" bson:"info,omitempty"`
	Decision        ZfcsEventResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	Prescription    ZfcsEventResultPrescriptionType `xml:"prescription,omitempty" bson:"prescription,omitempty"`
}

// UnplannedRevision is generated from an XSD element
type UnplannedRevision struct {
	UnplannedEventNumber string                          `xml:"unplannedEventNumber,omitempty" bson:"unplannedEventNumber,omitempty"`
	EventSubjects        []ZfcsEventSubjectType          `xml:"eventSubjects,omitempty" bson:"eventSubjects,omitempty"`
	Info                 string                          `xml:"info,omitempty" bson:"info,omitempty"`
	Attachments          ZfcsAttachmentListType          `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Orders               []Orders                        `xml:"orders,omitempty" bson:"orders,omitempty"`
	ObjectOtherInfo      string                          `xml:"objectOtherInfo,omitempty" bson:"objectOtherInfo,omitempty"`
	Prescription         ZfcsEventResultPrescriptionType `xml:"prescription,omitempty" bson:"prescription,omitempty"`
	Act                  ZfcsEventResultActType          `xml:"act,omitempty" bson:"act,omitempty"`
	Decision             ZfcsEventResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	Conclusion           ZfcsEventResultDecisionType     `xml:"conclusion,omitempty" bson:"conclusion,omitempty"`
}

// Orders is generated from an XSD element
type Orders struct {
	Info           string                    `xml:"info,omitempty" bson:"info,omitempty"`
	Purchase       ZfcsComplaintPurchaseType `xml:"purchase,omitempty" bson:"purchase,omitempty"`
	Order          ZfcsComplaintOrderType    `xml:"order,omitempty" bson:"order,omitempty"`
	SingleCustomer bool                      `xml:"singleCustomer,omitempty" bson:"singleCustomer,omitempty"`
}

// PlannedCheck is generated from an XSD element
type PlannedCheck struct {
	EventNumber          string                          `xml:"eventNumber,omitempty" bson:"eventNumber,omitempty"`
	InspectionType       string                          `xml:"inspectionType,omitempty" bson:"inspectionType,omitempty"`
	CheckSubject         ZfcsEventSubjectType            `xml:"checkSubject,omitempty" bson:"checkSubject,omitempty"`
	Orders               []ZfcsEventObjectType           `xml:"orders,omitempty" bson:"orders,omitempty"`
	ObjectOtherInfo      string                          `xml:"objectOtherInfo,omitempty" bson:"objectOtherInfo,omitempty"`
	Info                 string                          `xml:"info,omitempty" bson:"info,omitempty"`
	Act                  ZfcsEventResultActType          `xml:"act,omitempty" bson:"act,omitempty"`
	ActPrescription      ZfcsEventResultPrescriptionType `xml:"actPrescription,omitempty" bson:"actPrescription,omitempty"`
	Decision             ZfcsEventResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	DecisionPrescription ZfcsEventResultPrescriptionType `xml:"decisionPrescription,omitempty" bson:"decisionPrescription,omitempty"`
	Conclusion           ZfcsEventResultDecisionType     `xml:"conclusion,omitempty" bson:"conclusion,omitempty"`
}

// PlannedRevision is generated from an XSD element
type PlannedRevision struct {
	EventNumber          string                          `xml:"eventNumber,omitempty" bson:"eventNumber,omitempty"`
	CheckSubject         ZfcsEventSubjectType            `xml:"checkSubject,omitempty" bson:"checkSubject,omitempty"`
	Orders               []ZfcsEventObjectType           `xml:"orders,omitempty" bson:"orders,omitempty"`
	ObjectOtherInfo      string                          `xml:"objectOtherInfo,omitempty" bson:"objectOtherInfo,omitempty"`
	Info                 string                          `xml:"info,omitempty" bson:"info,omitempty"`
	Act                  ZfcsEventResultActType          `xml:"act,omitempty" bson:"act,omitempty"`
	ActPrescription      ZfcsEventResultPrescriptionType `xml:"actPrescription,omitempty" bson:"actPrescription,omitempty"`
	Decision             ZfcsEventResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	DecisionPrescription ZfcsEventResultPrescriptionType `xml:"decisionPrescription,omitempty" bson:"decisionPrescription,omitempty"`
	Conclusion           ZfcsEventResultDecisionType     `xml:"conclusion,omitempty" bson:"conclusion,omitempty"`
}

// PlannedSurvey is generated from an XSD element
type PlannedSurvey struct {
	EventNumber          string                          `xml:"eventNumber,omitempty" bson:"eventNumber,omitempty"`
	SurveyType           string                          `xml:"surveyType,omitempty" bson:"surveyType,omitempty"`
	CheckSubject         ZfcsEventSubjectType            `xml:"checkSubject,omitempty" bson:"checkSubject,omitempty"`
	Orders               []ZfcsEventObjectType           `xml:"orders,omitempty" bson:"orders,omitempty"`
	ObjectOtherInfo      string                          `xml:"objectOtherInfo,omitempty" bson:"objectOtherInfo,omitempty"`
	Info                 string                          `xml:"info,omitempty" bson:"info,omitempty"`
	Act                  ZfcsEventResultActType          `xml:"act,omitempty" bson:"act,omitempty"`
	ActPrescription      ZfcsEventResultPrescriptionType `xml:"actPrescription,omitempty" bson:"actPrescription,omitempty"`
	Decision             ZfcsEventResultDecisionType     `xml:"decision,omitempty" bson:"decision,omitempty"`
	DecisionPrescription ZfcsEventResultPrescriptionType `xml:"decisionPrescription,omitempty" bson:"decisionPrescription,omitempty"`
	Conclusion           ZfcsEventResultDecisionType     `xml:"conclusion,omitempty" bson:"conclusion,omitempty"`
}

// ZfcsAppRejectedReasonType is generated from an XSD element
type ZfcsAppRejectedReasonType struct {
	NsiRejectReason NsiRejectReason `xml:"nsiRejectReason,omitempty" bson:"nsiRejectReason,omitempty"`
	Explanation     string          `xml:"explanation,omitempty" bson:"explanation,omitempty"`
}

// ZfcsNsiBusinessControlType is generated from an XSD element
type ZfcsNsiBusinessControlType struct {
	GUID        string `xml:"GUID,omitempty" bson:"GUID,omitempty"`
	Code        string `xml:"code,omitempty" bson:"code,omitempty"`
	Order       uint64 `xml:"order,omitempty" bson:"order,omitempty"`
	Name        string `xml:"name,omitempty" bson:"name,omitempty"`
	Actual      bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
	SubSystem   string `xml:"subSystem,omitempty" bson:"subSystem,omitempty"`
	Document    string `xml:"document,omitempty" bson:"document,omitempty"`
	Action      string `xml:"action,omitempty" bson:"action,omitempty"`
	Level       string `xml:"level,omitempty" bson:"level,omitempty"`
	Description string `xml:"description,omitempty" bson:"description,omitempty"`
}

// ZfcsContactPersonType is generated from an XSD element
type ZfcsContactPersonType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
}

// ZfcsContract2015SingleCustomerType is generated from an XSD element
type ZfcsContract2015SingleCustomerType struct {
	Reason      Reason                         `xml:"reason,omitempty" bson:"reason,omitempty"`
	Document    Document                       `xml:"document,omitempty" bson:"document,omitempty"`
	ReportBase  string                         `xml:"reportBase,omitempty" bson:"reportBase,omitempty"`
	ReportCode  string                         `xml:"reportCode,omitempty" bson:"reportCode,omitempty"`
	Attachments ZfcsContractAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsCustomerReportContractExecutionInvalidType is generated from an XSD element
type ZfcsCustomerReportContractExecutionInvalidType struct {
	SchemeVersion     string                                  `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                int64                                   `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID        string                                  `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate           time.Time                               `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate    time.Time                               `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber     string                                  `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg        ZfcsOrganizationInfoWithOgrnType        `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href              string                                  `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm         ZfcsPrintFormType                       `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType                    `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments       ZfcsAttachmentListType                  `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID          string                                  `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	InvalidReportInfo ZfcsInvalidReportType                   `xml:"invalidReportInfo,omitempty" bson:"invalidReportInfo,omitempty"`
	Report            ZfcsCustomerReportContractExecutionType `xml:"report,omitempty" bson:"report,omitempty"`
}

// ZfcsContractTerminationReasonType is generated from an XSD element
type ZfcsContractTerminationReasonType struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsQuickRefOrganizationType is generated from an XSD element
type ZfcsQuickRefOrganizationType struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
	FactAddress     string `xml:"factAddress,omitempty" bson:"factAddress,omitempty"`
	INN             string `xml:"INN,omitempty" bson:"INN,omitempty"`
	KPP             string `xml:"KPP,omitempty" bson:"KPP,omitempty"`
	ContactEMail    string `xml:"contactEMail,omitempty" bson:"contactEMail,omitempty"`
}

// ZfcsCustomerReportInvalidType is generated from an XSD element
type ZfcsCustomerReportInvalidType struct {
	SchemeVersion     string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID        string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	DocDate           time.Time                        `xml:"docDate,omitempty" bson:"docDate,omitempty"`
	DocPublishDate    time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	VersionNumber     string                           `xml:"versionNumber,omitempty" bson:"versionNumber,omitempty"`
	PublishOrg        ZfcsOrganizationInfoWithOgrnType `xml:"publishOrg,omitempty" bson:"publishOrg,omitempty"`
	Href              string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm         ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm      ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments       ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReportID          string                           `xml:"reportId,omitempty" bson:"reportId,omitempty"`
	InvalidReportInfo ZfcsInvalidReportType            `xml:"invalidReportInfo,omitempty" bson:"invalidReportInfo,omitempty"`
}

// ZfcsNotificationZKType is generated from an XSD element
type ZfcsNotificationZKType struct {
	SchemeVersion       string                           `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                  int64                            `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID          string                           `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber      string                           `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	DocPublishDate      time.Time                        `xml:"docPublishDate,omitempty" bson:"docPublishDate,omitempty"`
	DocNumber           string                           `xml:"docNumber,omitempty" bson:"docNumber,omitempty"`
	Href                string                           `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm           ZfcsPrintFormType                `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm        ZfcsExtPrintFormType             `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	PurchaseObjectInfo  string                           `xml:"purchaseObjectInfo,omitempty" bson:"purchaseObjectInfo,omitempty"`
	PurchaseResponsible PurchaseResponsible              `xml:"purchaseResponsible,omitempty" bson:"purchaseResponsible,omitempty"`
	PlacingWay          ZfcsPlacingWayType               `xml:"placingWay,omitempty" bson:"placingWay,omitempty"`
	Okpd2okved2         bool                             `xml:"okpd2okved2,omitempty" bson:"okpd2okved2,omitempty"`
	ContractServiceInfo string                           `xml:"contractServiceInfo,omitempty" bson:"contractServiceInfo,omitempty"`
	ProcedureInfo       ProcedureInfo                    `xml:"procedureInfo,omitempty" bson:"procedureInfo,omitempty"`
	Lot                 Lot                              `xml:"lot,omitempty" bson:"lot,omitempty"`
	Attachments         ZfcsAttachmentListType           `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification        ZfcsNotificationModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
}

// ZfcsLotInfoType is generated from an XSD element
type ZfcsLotInfoType struct {
	LotObjectInfo          string               `xml:"lotObjectInfo,omitempty" bson:"lotObjectInfo,omitempty"`
	Currency               ZfcsCurrencyRef      `xml:"currency,omitempty" bson:"currency,omitempty"`
	MaxPrice               string               `xml:"maxPrice,omitempty" bson:"maxPrice,omitempty"`
	SpelledMaxPrice        string               `xml:"spelledMaxPrice,omitempty" bson:"spelledMaxPrice,omitempty"`
	UnitPrice              string               `xml:"unitPrice,omitempty" bson:"unitPrice,omitempty"`
	SpelledUnitPrice       string               `xml:"spelledUnitPrice,omitempty" bson:"spelledUnitPrice,omitempty"`
	StandardContractNumber string               `xml:"standardContractNumber,omitempty" bson:"standardContractNumber,omitempty"`
	PriceFormula           string               `xml:"priceFormula,omitempty" bson:"priceFormula,omitempty"`
	FinanceSource          string               `xml:"financeSource,omitempty" bson:"financeSource,omitempty"`
	DeliveryTerm           string               `xml:"deliveryTerm,omitempty" bson:"deliveryTerm,omitempty"`
	Customers              Customers            `xml:"customers,omitempty" bson:"customers,omitempty"`
	Preferenses            Preferenses          `xml:"preferenses,omitempty" bson:"preferenses,omitempty"`
	Requirements           Requirements         `xml:"requirements,omitempty" bson:"requirements,omitempty"`
	PurchaseProlongation   PurchaseProlongation `xml:"purchaseProlongation,omitempty" bson:"purchaseProlongation,omitempty"`
	SpelledInvalidAppCount string               `xml:"spelledInvalidAppCount,omitempty" bson:"spelledInvalidAppCount,omitempty"`
	SpelledAppCount        string               `xml:"spelledAppCount,omitempty" bson:"spelledAppCount,omitempty"`
	SpelledValidAppCount   string               `xml:"spelledValidAppCount,omitempty" bson:"spelledValidAppCount,omitempty"`
	DeliveryPlace          string               `xml:"deliveryPlace,omitempty" bson:"deliveryPlace,omitempty"`
	KladrPlaces            ZfcsKladrPlacesType  `xml:"kladrPlaces,omitempty" bson:"kladrPlaces,omitempty"`
}

// Customers is generated from an XSD element
type Customers struct {
	Customer []ZfcsOrganizationRef `xml:"customer,omitempty" bson:"customer,omitempty"`
}

// PurchaseProlongation is generated from an XSD element
type PurchaseProlongation struct {
	ProlongationNumber string    `xml:"prolongationNumber,omitempty" bson:"prolongationNumber,omitempty"`
	PublishDate        time.Time `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
}

// ContactPersonType is generated from an XSD element
type ContactPersonType struct {
	LastName   string `xml:"lastName,omitempty" bson:"lastName,omitempty"`
	FirstName  string `xml:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName string `xml:"middleName,omitempty" bson:"middleName,omitempty"`
}

// ZfcsNsiContractCurrencyCBRFType is generated from an XSD element
type ZfcsNsiContractCurrencyCBRFType struct {
	Currency ZfcsCurrencyFullRef `xml:"currency,omitempty" bson:"currency,omitempty"`
	Actual   bool                `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsPurchasePlanPositionChangeReasonRef is generated from an XSD element
type ZfcsPurchasePlanPositionChangeReasonRef struct {
	Code string `xml:"code,omitempty" bson:"code,omitempty"`
	Name string `xml:"name,omitempty" bson:"name,omitempty"`
}

// ZfcsCommissionType is generated from an XSD element
type ZfcsCommissionType struct {
	CommissionName    string            `xml:"commissionName,omitempty" bson:"commissionName,omitempty"`
	CommissionMembers CommissionMembers `xml:"commissionMembers,omitempty" bson:"commissionMembers,omitempty"`
	Competent         bool              `xml:"competent,omitempty" bson:"competent,omitempty"`
	AddInfo           string            `xml:"addInfo,omitempty" bson:"addInfo,omitempty"`
}

// ClarificationRequestType is generated from an XSD element
type ClarificationRequestType struct {
	NotificationNumber string       `xml:"notificationNumber,omitempty" bson:"notificationNumber,omitempty"`
	RegNumber          string       `xml:"regNumber,omitempty" bson:"regNumber,omitempty"`
	RegDate            time.Time    `xml:"regDate,omitempty" bson:"regDate,omitempty"`
	Topic              string       `xml:"topic,omitempty" bson:"topic,omitempty"`
	Participant        Participant  `xml:"participant,omitempty" bson:"participant,omitempty"`
	DocumentMetas      DocumentList `xml:"documentMetas,omitempty" bson:"documentMetas,omitempty"`
}

// CriterionRef is generated from an XSD element
type CriterionRef struct {
	ID            int64  `xml:"id,omitempty" bson:"id,omitempty"`
	CriterionType string `xml:"criterionType,omitempty" bson:"criterionType,omitempty"`
}

// ZfcsComplaintProjectType is generated from an XSD element
type ZfcsComplaintProjectType struct {
	SchemeVersion string                            `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID            int64                             `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID    string                            `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	CommonInfo    CommonInfo                        `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	Indicted      []ZfcsComplaintProjectSubjectType `xml:"indicted,omitempty" bson:"indicted,omitempty"`
	Applicant     ZfcsApplicantType                 `xml:"applicant,omitempty" bson:"applicant,omitempty"`
	Object        ZfcsComplaintObjectType           `xml:"object,omitempty" bson:"object,omitempty"`
	Text          string                            `xml:"text,omitempty" bson:"text,omitempty"`
	PrintForm     ZfcsPrintFormType                 `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType              `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType            `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	ReturnInfo    ZfcsComplaintReturnInfoType       `xml:"returnInfo,omitempty" bson:"returnInfo,omitempty"`
}

// ZfcsEventResultActType is generated from an XSD element
type ZfcsEventResultActType struct {
	ActNumber   string                 `xml:"actNumber,omitempty" bson:"actNumber,omitempty"`
	ActDate     xsd.Date               `xml:"actDate,omitempty" bson:"actDate,omitempty"`
	Attachments ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// AuctionItemsType is generated from an XSD element
type AuctionItemsType struct {
	AuctionItem []AuctionItem `xml:"auctionItem,omitempty" bson:"auctionItem,omitempty"`
}

// AuctionItem is generated from an XSD element
type AuctionItem struct {
	Sid         int64   `xml:"sid,omitempty" bson:"sid,omitempty"`
	Description string  `xml:"description,omitempty" bson:"description,omitempty"`
	Price       float64 `xml:"price,omitempty" bson:"price,omitempty"`
}

// ZfcsOrganizationRef is generated from an XSD element
type ZfcsOrganizationRef struct {
	RegNum          string `xml:"regNum,omitempty" bson:"regNum,omitempty"`
	ConsRegistryNum string `xml:"consRegistryNum,omitempty" bson:"consRegistryNum,omitempty"`
	FullName        string `xml:"fullName,omitempty" bson:"fullName,omitempty"`
}

// ZfcsContractOKEIType is generated from an XSD element
type ZfcsContractOKEIType struct {
	Code         string `xml:"code,omitempty" bson:"code,omitempty"`
	NationalCode string `xml:"nationalCode,omitempty" bson:"nationalCode,omitempty"`
}

// ZfcsProtocolPOType is generated from an XSD element
type ZfcsProtocolPOType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	PurchaseInfo             PurchaseInfo                 `xml:"purchaseInfo,omitempty" bson:"purchaseInfo,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

// ZfcsOrganizationLink is generated from an XSD element
type ZfcsOrganizationLink struct {
	ID                    int64               `xml:"id,omitempty" bson:"id,omitempty"`
	ActiveUntil           time.Time           `xml:"activeUntil,omitempty" bson:"activeUntil,omitempty"`
	BlockStatus           string              `xml:"blockStatus,omitempty" bson:"blockStatus,omitempty"`
	DependentOrganization ZfcsOrganizationRef `xml:"dependentOrganization,omitempty" bson:"dependentOrganization,omitempty"`
	LinkUsers             []ZfcsLinkUser      `xml:"linkUsers,omitempty" bson:"linkUsers,omitempty"`
	OrdersVisibilityType  string              `xml:"ordersVisibilityType,omitempty" bson:"ordersVisibilityType,omitempty"`
	LastModifyDate        time.Time           `xml:"lastModifyDate,omitempty" bson:"lastModifyDate,omitempty"`
}

// ZfcsNsiPurchaseDocumentTypesType is generated from an XSD element
type ZfcsNsiPurchaseDocumentTypesType struct {
	PlacingWayCode string        `xml:"placingWayCode,omitempty" bson:"placingWayCode,omitempty"`
	PlacingWayType string        `xml:"placingWayType,omitempty" bson:"placingWayType,omitempty"`
	PlacingWayName string        `xml:"placingWayName,omitempty" bson:"placingWayName,omitempty"`
	Actual         bool          `xml:"actual,omitempty" bson:"actual,omitempty"`
	DocumentTypes  DocumentTypes `xml:"documentTypes,omitempty" bson:"documentTypes,omitempty"`
}

// DocumentTypes is generated from an XSD element
type DocumentTypes struct {
	DocumentType []DocumentType `xml:"documentType,omitempty" bson:"documentType,omitempty"`
}

// DocumentType is generated from an XSD element
type DocumentType struct {
	Code       string `xml:"code,omitempty" bson:"code,omitempty"`
	Name       string `xml:"name,omitempty" bson:"name,omitempty"`
	Structured bool   `xml:"structured,omitempty" bson:"structured,omitempty"`
	Actual     bool   `xml:"actual,omitempty" bson:"actual,omitempty"`
}

// ZfcsCheckResultCancelType is generated from an XSD element
type ZfcsCheckResultCancelType struct {
	SchemeVersion string                 `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	CommonInfo    CommonInfo             `xml:"commonInfo,omitempty" bson:"commonInfo,omitempty"`
	CancelType    string                 `xml:"cancelType,omitempty" bson:"cancelType,omitempty"`
	Info          string                 `xml:"info,omitempty" bson:"info,omitempty"`
	DocumentName  string                 `xml:"documentName,omitempty" bson:"documentName,omitempty"`
	DocumentDate  time.Time              `xml:"documentDate,omitempty" bson:"documentDate,omitempty"`
	PrintForm     ZfcsPrintFormType      `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm  ZfcsExtPrintFormType   `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	Attachments   ZfcsAttachmentListType `xml:"attachments,omitempty" bson:"attachments,omitempty"`
}

// ZfcsProtocolEFInvalidationType is generated from an XSD element
type ZfcsProtocolEFInvalidationType struct {
	SchemeVersion            string                       `xml:"schemeVersion,attr" bson:"schemeVersion,omitempty"`
	ID                       int64                        `xml:"id,omitempty" bson:"id,omitempty"`
	ExternalID               string                       `xml:"externalId,omitempty" bson:"externalId,omitempty"`
	PurchaseNumber           string                       `xml:"purchaseNumber,omitempty" bson:"purchaseNumber,omitempty"`
	ProtocolNumber           string                       `xml:"protocolNumber,omitempty" bson:"protocolNumber,omitempty"`
	FoundationProtocolNumber string                       `xml:"foundationProtocolNumber,omitempty" bson:"foundationProtocolNumber,omitempty"`
	ParentProtocolNumber     string                       `xml:"parentProtocolNumber,omitempty" bson:"parentProtocolNumber,omitempty"`
	Place                    string                       `xml:"place,omitempty" bson:"place,omitempty"`
	ProtocolDate             time.Time                    `xml:"protocolDate,omitempty" bson:"protocolDate,omitempty"`
	SignDate                 time.Time                    `xml:"signDate,omitempty" bson:"signDate,omitempty"`
	PublishDate              time.Time                    `xml:"publishDate,omitempty" bson:"publishDate,omitempty"`
	Commission               ZfcsCommissionType           `xml:"commission,omitempty" bson:"commission,omitempty"`
	Href                     string                       `xml:"href,omitempty" bson:"href,omitempty"`
	PrintForm                ZfcsPrintFormType            `xml:"printForm,omitempty" bson:"printForm,omitempty"`
	ExtPrintForm             ZfcsExtPrintFormType         `xml:"extPrintForm,omitempty" bson:"extPrintForm,omitempty"`
	ProtocolPublisher        ProtocolPublisher            `xml:"protocolPublisher,omitempty" bson:"protocolPublisher,omitempty"`
	Attachments              ZfcsAttachmentListType       `xml:"attachments,omitempty" bson:"attachments,omitempty"`
	Modification             ZfcsProtocolModificationType `xml:"modification,omitempty" bson:"modification,omitempty"`
	ProtocolLot              ProtocolLot                  `xml:"protocolLot,omitempty" bson:"protocolLot,omitempty"`
}

func GetExportInterface(xmlBts []byte, title string) interface{} {
	export := Export{}
	xml.Unmarshal(xmlBts, &export)
	expVal := reflect.ValueOf(export)

	for i := 0; i < expVal.NumField(); i++ {
		fieldValue := expVal.Field(i)
		fieldType := expVal.Type().Field(i)
		tag := fieldType.Tag.Get("xml")

		if strings.Split(tag, ",")[0] == title {
			return fieldValue.Interface()
		}
	}
	return nil
}

// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/a_w_s_connector"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/a_w_s_evaluations"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/assessment_reports"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/azure_connector"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/azure_evaluations"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/connector_groups_management"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/g_c_p_connector"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/g_c_p_evaluations"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/remediation_activity"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/reports"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/response_actions"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/response_notifications"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/response_rules"
	"github.com/form3tech-oss/terraform-provider-qualys/swagger/server/operations/user_access_management"
)

//go:generate swagger generate server --target ../../swagger --name CloudViewAPIs --spec ../swagger.json --server-package server --principal interface{}

func configureFlags(api *operations.CloudViewAPIsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CloudViewAPIsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.EmptyProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented(" producer has not yet been implemented")
	})
	api.BinProducer = runtime.ByteStreamProducer()

	// Applies when the Authorization header is set with the Basic scheme
	if api.BasicAuthAuth == nil {
		api.BasicAuthAuth = func(user string, pass string) (interface{}, error) {
			return nil, errors.NotImplemented("basic auth  (basicAuth) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.ResponseRulesActivateRuleUsingPOSTHandler == nil {
		api.ResponseRulesActivateRuleUsingPOSTHandler = response_rules.ActivateRuleUsingPOSTHandlerFunc(func(params response_rules.ActivateRuleUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.ActivateRuleUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsCreateEmailActionUsingPOSTHandler == nil {
		api.ResponseActionsCreateEmailActionUsingPOSTHandler = response_actions.CreateEmailActionUsingPOSTHandlerFunc(func(params response_actions.CreateEmailActionUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.CreateEmailActionUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsCreatePagerActionUsingPOSTHandler == nil {
		api.ResponseActionsCreatePagerActionUsingPOSTHandler = response_actions.CreatePagerActionUsingPOSTHandlerFunc(func(params response_actions.CreatePagerActionUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.CreatePagerActionUsingPOST has not yet been implemented")
		})
	}
	if api.AssessmentReportsCreateReportUsingPOSTHandler == nil {
		api.AssessmentReportsCreateReportUsingPOSTHandler = assessment_reports.CreateReportUsingPOSTHandlerFunc(func(params assessment_reports.CreateReportUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation assessment_reports.CreateReportUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseRulesCreateRuleUsingPOSTHandler == nil {
		api.ResponseRulesCreateRuleUsingPOSTHandler = response_rules.CreateRuleUsingPOSTHandlerFunc(func(params response_rules.CreateRuleUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.CreateRuleUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsCreateSlackActionUsingPOSTHandler == nil {
		api.ResponseActionsCreateSlackActionUsingPOSTHandler = response_actions.CreateSlackActionUsingPOSTHandlerFunc(func(params response_actions.CreateSlackActionUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.CreateSlackActionUsingPOST has not yet been implemented")
		})
	}
	if api.AzureConnectorCreateUsingPOSTHandler == nil {
		api.AzureConnectorCreateUsingPOSTHandler = azure_connector.CreateUsingPOSTHandlerFunc(func(params azure_connector.CreateUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.CreateUsingPOST has not yet been implemented")
		})
	}
	if api.AwsConnectorCreateUsingPOST1Handler == nil {
		api.AwsConnectorCreateUsingPOST1Handler = a_w_s_connector.CreateUsingPOST1HandlerFunc(func(params a_w_s_connector.CreateUsingPOST1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.CreateUsingPOST1 has not yet been implemented")
		})
	}
	if api.GcpConnectorCreateUsingPOST2Handler == nil {
		api.GcpConnectorCreateUsingPOST2Handler = g_c_p_connector.CreateUsingPOST2HandlerFunc(func(params g_c_p_connector.CreateUsingPOST2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.CreateUsingPOST2 has not yet been implemented")
		})
	}
	if api.ConnectorGroupsManagementCreateUsingPOST3Handler == nil {
		api.ConnectorGroupsManagementCreateUsingPOST3Handler = connector_groups_management.CreateUsingPOST3HandlerFunc(func(params connector_groups_management.CreateUsingPOST3Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation connector_groups_management.CreateUsingPOST3 has not yet been implemented")
		})
	}
	if api.ReportsCreateUsingPOST4Handler == nil {
		api.ReportsCreateUsingPOST4Handler = reports.CreateUsingPOST4HandlerFunc(func(params reports.CreateUsingPOST4Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.CreateUsingPOST4 has not yet been implemented")
		})
	}
	if api.ResponseActionsDeleteActionUsingPOSTHandler == nil {
		api.ResponseActionsDeleteActionUsingPOSTHandler = response_actions.DeleteActionUsingPOSTHandlerFunc(func(params response_actions.DeleteActionUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.DeleteActionUsingPOST has not yet been implemented")
		})
	}
	if api.AzureConnectorDeleteConnectorsUsingDELETEHandler == nil {
		api.AzureConnectorDeleteConnectorsUsingDELETEHandler = azure_connector.DeleteConnectorsUsingDELETEHandlerFunc(func(params azure_connector.DeleteConnectorsUsingDELETEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.DeleteConnectorsUsingDELETE has not yet been implemented")
		})
	}
	if api.AwsConnectorDeleteConnectorsUsingDELETE1Handler == nil {
		api.AwsConnectorDeleteConnectorsUsingDELETE1Handler = a_w_s_connector.DeleteConnectorsUsingDELETE1HandlerFunc(func(params a_w_s_connector.DeleteConnectorsUsingDELETE1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.DeleteConnectorsUsingDELETE1 has not yet been implemented")
		})
	}
	if api.GcpConnectorDeleteConnectorsUsingDELETE2Handler == nil {
		api.GcpConnectorDeleteConnectorsUsingDELETE2Handler = g_c_p_connector.DeleteConnectorsUsingDELETE2HandlerFunc(func(params g_c_p_connector.DeleteConnectorsUsingDELETE2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.DeleteConnectorsUsingDELETE2 has not yet been implemented")
		})
	}
	if api.ResponseRulesDeleteRuleUsingPOSTHandler == nil {
		api.ResponseRulesDeleteRuleUsingPOSTHandler = response_rules.DeleteRuleUsingPOSTHandlerFunc(func(params response_rules.DeleteRuleUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.DeleteRuleUsingPOST has not yet been implemented")
		})
	}
	if api.ReportsDeleteUsingDELETEHandler == nil {
		api.ReportsDeleteUsingDELETEHandler = reports.DeleteUsingDELETEHandlerFunc(func(params reports.DeleteUsingDELETEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.DeleteUsingDELETE has not yet been implemented")
		})
	}
	if api.AzureConnectorDisableConnectorsUsingPATCHHandler == nil {
		api.AzureConnectorDisableConnectorsUsingPATCHHandler = azure_connector.DisableConnectorsUsingPATCHHandlerFunc(func(params azure_connector.DisableConnectorsUsingPATCHParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.DisableConnectorsUsingPATCH has not yet been implemented")
		})
	}
	if api.AwsConnectorDisableConnectorsUsingPATCH1Handler == nil {
		api.AwsConnectorDisableConnectorsUsingPATCH1Handler = a_w_s_connector.DisableConnectorsUsingPATCH1HandlerFunc(func(params a_w_s_connector.DisableConnectorsUsingPATCH1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.DisableConnectorsUsingPATCH1 has not yet been implemented")
		})
	}
	if api.GcpConnectorDisableConnectorsUsingPATCH2Handler == nil {
		api.GcpConnectorDisableConnectorsUsingPATCH2Handler = g_c_p_connector.DisableConnectorsUsingPATCH2HandlerFunc(func(params g_c_p_connector.DisableConnectorsUsingPATCH2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.DisableConnectorsUsingPATCH2 has not yet been implemented")
		})
	}
	if api.ResponseRulesDisableRuleUsingPOSTHandler == nil {
		api.ResponseRulesDisableRuleUsingPOSTHandler = response_rules.DisableRuleUsingPOSTHandlerFunc(func(params response_rules.DisableRuleUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.DisableRuleUsingPOST has not yet been implemented")
		})
	}
	if api.AwsConnectorDownloadAwsCloudFormationTemplateUsingGETHandler == nil {
		api.AwsConnectorDownloadAwsCloudFormationTemplateUsingGETHandler = a_w_s_connector.DownloadAwsCloudFormationTemplateUsingGETHandlerFunc(func(params a_w_s_connector.DownloadAwsCloudFormationTemplateUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.DownloadAwsCloudFormationTemplateUsingGET has not yet been implemented")
		})
	}
	if api.AssessmentReportsDownloadReportUsingGETHandler == nil {
		api.AssessmentReportsDownloadReportUsingGETHandler = assessment_reports.DownloadReportUsingGETHandlerFunc(func(params assessment_reports.DownloadReportUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation assessment_reports.DownloadReportUsingGET has not yet been implemented")
		})
	}
	if api.AzureConnectorEnableConnectorsUsingPATCHHandler == nil {
		api.AzureConnectorEnableConnectorsUsingPATCHHandler = azure_connector.EnableConnectorsUsingPATCHHandlerFunc(func(params azure_connector.EnableConnectorsUsingPATCHParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.EnableConnectorsUsingPATCH has not yet been implemented")
		})
	}
	if api.AwsConnectorEnableConnectorsUsingPATCH1Handler == nil {
		api.AwsConnectorEnableConnectorsUsingPATCH1Handler = a_w_s_connector.EnableConnectorsUsingPATCH1HandlerFunc(func(params a_w_s_connector.EnableConnectorsUsingPATCH1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.EnableConnectorsUsingPATCH1 has not yet been implemented")
		})
	}
	if api.GcpConnectorEnableConnectorsUsingPATCH2Handler == nil {
		api.GcpConnectorEnableConnectorsUsingPATCH2Handler = g_c_p_connector.EnableConnectorsUsingPATCH2HandlerFunc(func(params g_c_p_connector.EnableConnectorsUsingPATCH2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.EnableConnectorsUsingPATCH2 has not yet been implemented")
		})
	}
	if api.ResponseActionsGetActionByIDUsingGETHandler == nil {
		api.ResponseActionsGetActionByIDUsingGETHandler = response_actions.GetActionByIDUsingGETHandlerFunc(func(params response_actions.GetActionByIDUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.GetActionByIDUsingGET has not yet been implemented")
		})
	}
	if api.ResponseActionsGetActionsTypesUsingGETHandler == nil {
		api.ResponseActionsGetActionsTypesUsingGETHandler = response_actions.GetActionsTypesUsingGETHandlerFunc(func(params response_actions.GetActionsTypesUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.GetActionsTypesUsingGET has not yet been implemented")
		})
	}
	if api.ResponseActionsGetAllActionsByFilterUsingGETHandler == nil {
		api.ResponseActionsGetAllActionsByFilterUsingGETHandler = response_actions.GetAllActionsByFilterUsingGETHandlerFunc(func(params response_actions.GetAllActionsByFilterUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.GetAllActionsByFilterUsingGET has not yet been implemented")
		})
	}
	if api.ReportsGetAllMandatesUsingGETHandler == nil {
		api.ReportsGetAllMandatesUsingGETHandler = reports.GetAllMandatesUsingGETHandlerFunc(func(params reports.GetAllMandatesUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.GetAllMandatesUsingGET has not yet been implemented")
		})
	}
	if api.ResponseNotificationsGetAllNotificationsUsingGETHandler == nil {
		api.ResponseNotificationsGetAllNotificationsUsingGETHandler = response_notifications.GetAllNotificationsUsingGETHandlerFunc(func(params response_notifications.GetAllNotificationsUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_notifications.GetAllNotificationsUsingGET has not yet been implemented")
		})
	}
	if api.ReportsGetAllPoliciesUsingGETHandler == nil {
		api.ReportsGetAllPoliciesUsingGETHandler = reports.GetAllPoliciesUsingGETHandlerFunc(func(params reports.GetAllPoliciesUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.GetAllPoliciesUsingGET has not yet been implemented")
		})
	}
	if api.ResponseRulesGetAllRulesByFilterUsingGETHandler == nil {
		api.ResponseRulesGetAllRulesByFilterUsingGETHandler = response_rules.GetAllRulesByFilterUsingGETHandlerFunc(func(params response_rules.GetAllRulesByFilterUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.GetAllRulesByFilterUsingGET has not yet been implemented")
		})
	}
	if api.AwsConnectorGetAwsAccountIDUsingGETHandler == nil {
		api.AwsConnectorGetAwsAccountIDUsingGETHandler = a_w_s_connector.GetAwsAccountIDUsingGETHandlerFunc(func(params a_w_s_connector.GetAwsAccountIDUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.GetAwsAccountIDUsingGET has not yet been implemented")
		})
	}
	if api.AzureConnectorGetAzureErrorsListUsingGETHandler == nil {
		api.AzureConnectorGetAzureErrorsListUsingGETHandler = azure_connector.GetAzureErrorsListUsingGETHandlerFunc(func(params azure_connector.GetAzureErrorsListUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.GetAzureErrorsListUsingGET has not yet been implemented")
		})
	}
	if api.ReportsGetDataUsingGETHandler == nil {
		api.ReportsGetDataUsingGETHandler = reports.GetDataUsingGETHandlerFunc(func(params reports.GetDataUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.GetDataUsingGET has not yet been implemented")
		})
	}
	if api.AzureConnectorGetDetailsUsingGETHandler == nil {
		api.AzureConnectorGetDetailsUsingGETHandler = azure_connector.GetDetailsUsingGETHandlerFunc(func(params azure_connector.GetDetailsUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.GetDetailsUsingGET has not yet been implemented")
		})
	}
	if api.AwsConnectorGetDetailsUsingGET1Handler == nil {
		api.AwsConnectorGetDetailsUsingGET1Handler = a_w_s_connector.GetDetailsUsingGET1HandlerFunc(func(params a_w_s_connector.GetDetailsUsingGET1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.GetDetailsUsingGET1 has not yet been implemented")
		})
	}
	if api.GcpConnectorGetDetailsUsingGET2Handler == nil {
		api.GcpConnectorGetDetailsUsingGET2Handler = g_c_p_connector.GetDetailsUsingGET2HandlerFunc(func(params g_c_p_connector.GetDetailsUsingGET2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.GetDetailsUsingGET2 has not yet been implemented")
		})
	}
	if api.AwsConnectorGetErrorsListUsingGETHandler == nil {
		api.AwsConnectorGetErrorsListUsingGETHandler = a_w_s_connector.GetErrorsListUsingGETHandlerFunc(func(params a_w_s_connector.GetErrorsListUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.GetErrorsListUsingGET has not yet been implemented")
		})
	}
	if api.AwsEvaluationsGetEvaluatedControlsUsingGETHandler == nil {
		api.AwsEvaluationsGetEvaluatedControlsUsingGETHandler = a_w_s_evaluations.GetEvaluatedControlsUsingGETHandlerFunc(func(params a_w_s_evaluations.GetEvaluatedControlsUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_evaluations.GetEvaluatedControlsUsingGET has not yet been implemented")
		})
	}
	if api.AzureEvaluationsGetEvaluatedControlsUsingGET1Handler == nil {
		api.AzureEvaluationsGetEvaluatedControlsUsingGET1Handler = azure_evaluations.GetEvaluatedControlsUsingGET1HandlerFunc(func(params azure_evaluations.GetEvaluatedControlsUsingGET1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_evaluations.GetEvaluatedControlsUsingGET1 has not yet been implemented")
		})
	}
	if api.GcpEvaluationsGetEvaluatedControlsUsingGET2Handler == nil {
		api.GcpEvaluationsGetEvaluatedControlsUsingGET2Handler = g_c_p_evaluations.GetEvaluatedControlsUsingGET2HandlerFunc(func(params g_c_p_evaluations.GetEvaluatedControlsUsingGET2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_evaluations.GetEvaluatedControlsUsingGET2 has not yet been implemented")
		})
	}
	if api.AwsEvaluationsGetEvaluatedResourcesUsingGETHandler == nil {
		api.AwsEvaluationsGetEvaluatedResourcesUsingGETHandler = a_w_s_evaluations.GetEvaluatedResourcesUsingGETHandlerFunc(func(params a_w_s_evaluations.GetEvaluatedResourcesUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_evaluations.GetEvaluatedResourcesUsingGET has not yet been implemented")
		})
	}
	if api.AzureEvaluationsGetEvaluatedResourcesUsingGET1Handler == nil {
		api.AzureEvaluationsGetEvaluatedResourcesUsingGET1Handler = azure_evaluations.GetEvaluatedResourcesUsingGET1HandlerFunc(func(params azure_evaluations.GetEvaluatedResourcesUsingGET1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_evaluations.GetEvaluatedResourcesUsingGET1 has not yet been implemented")
		})
	}
	if api.GcpEvaluationsGetEvaluatedResourcesUsingGET2Handler == nil {
		api.GcpEvaluationsGetEvaluatedResourcesUsingGET2Handler = g_c_p_evaluations.GetEvaluatedResourcesUsingGET2HandlerFunc(func(params g_c_p_evaluations.GetEvaluatedResourcesUsingGET2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_evaluations.GetEvaluatedResourcesUsingGET2 has not yet been implemented")
		})
	}
	if api.AwsEvaluationsGetEvaluationsStatsUsingGETHandler == nil {
		api.AwsEvaluationsGetEvaluationsStatsUsingGETHandler = a_w_s_evaluations.GetEvaluationsStatsUsingGETHandlerFunc(func(params a_w_s_evaluations.GetEvaluationsStatsUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_evaluations.GetEvaluationsStatsUsingGET has not yet been implemented")
		})
	}
	if api.AzureEvaluationsGetEvaluationsStatsUsingGET1Handler == nil {
		api.AzureEvaluationsGetEvaluationsStatsUsingGET1Handler = azure_evaluations.GetEvaluationsStatsUsingGET1HandlerFunc(func(params azure_evaluations.GetEvaluationsStatsUsingGET1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_evaluations.GetEvaluationsStatsUsingGET1 has not yet been implemented")
		})
	}
	if api.GcpEvaluationsGetEvaluationsStatsUsingGET2Handler == nil {
		api.GcpEvaluationsGetEvaluationsStatsUsingGET2Handler = g_c_p_evaluations.GetEvaluationsStatsUsingGET2HandlerFunc(func(params g_c_p_evaluations.GetEvaluationsStatsUsingGET2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_evaluations.GetEvaluationsStatsUsingGET2 has not yet been implemented")
		})
	}
	if api.GcpConnectorGetGcpErrorsListUsingGETHandler == nil {
		api.GcpConnectorGetGcpErrorsListUsingGETHandler = g_c_p_connector.GetGcpErrorsListUsingGETHandlerFunc(func(params g_c_p_connector.GetGcpErrorsListUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.GetGcpErrorsListUsingGET has not yet been implemented")
		})
	}
	if api.ConnectorGroupsManagementGetGroupUsingGETHandler == nil {
		api.ConnectorGroupsManagementGetGroupUsingGETHandler = connector_groups_management.GetGroupUsingGETHandlerFunc(func(params connector_groups_management.GetGroupUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation connector_groups_management.GetGroupUsingGET has not yet been implemented")
		})
	}
	if api.AzureConnectorGetListUsingGETHandler == nil {
		api.AzureConnectorGetListUsingGETHandler = azure_connector.GetListUsingGETHandlerFunc(func(params azure_connector.GetListUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.GetListUsingGET has not yet been implemented")
		})
	}
	if api.AwsConnectorGetListUsingGET1Handler == nil {
		api.AwsConnectorGetListUsingGET1Handler = a_w_s_connector.GetListUsingGET1HandlerFunc(func(params a_w_s_connector.GetListUsingGET1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.GetListUsingGET1 has not yet been implemented")
		})
	}
	if api.GcpConnectorGetListUsingGET2Handler == nil {
		api.GcpConnectorGetListUsingGET2Handler = g_c_p_connector.GetListUsingGET2HandlerFunc(func(params g_c_p_connector.GetListUsingGET2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.GetListUsingGET2 has not yet been implemented")
		})
	}
	if api.ReportsGetListUsingGET3Handler == nil {
		api.ReportsGetListUsingGET3Handler = reports.GetListUsingGET3HandlerFunc(func(params reports.GetListUsingGET3Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.GetListUsingGET3 has not yet been implemented")
		})
	}
	if api.AssessmentReportsGetListUsingGET4Handler == nil {
		api.AssessmentReportsGetListUsingGET4Handler = assessment_reports.GetListUsingGET4HandlerFunc(func(params assessment_reports.GetListUsingGET4Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation assessment_reports.GetListUsingGET4 has not yet been implemented")
		})
	}
	if api.ResponseNotificationsGetNotificationByIDUsingGETHandler == nil {
		api.ResponseNotificationsGetNotificationByIDUsingGETHandler = response_notifications.GetNotificationByIDUsingGETHandlerFunc(func(params response_notifications.GetNotificationByIDUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_notifications.GetNotificationByIDUsingGET has not yet been implemented")
		})
	}
	if api.RemediationActivityGetRemediationHistoryUsingGETHandler == nil {
		api.RemediationActivityGetRemediationHistoryUsingGETHandler = remediation_activity.GetRemediationHistoryUsingGETHandlerFunc(func(params remediation_activity.GetRemediationHistoryUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation remediation_activity.GetRemediationHistoryUsingGET has not yet been implemented")
		})
	}
	if api.ResponseRulesGetRuleByIDUsingGETHandler == nil {
		api.ResponseRulesGetRuleByIDUsingGETHandler = response_rules.GetRuleByIDUsingGETHandlerFunc(func(params response_rules.GetRuleByIDUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.GetRuleByIDUsingGET has not yet been implemented")
		})
	}
	if api.UserAccessManagementGetUserScopeUsingGETHandler == nil {
		api.UserAccessManagementGetUserScopeUsingGETHandler = user_access_management.GetUserScopeUsingGETHandlerFunc(func(params user_access_management.GetUserScopeUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user_access_management.GetUserScopeUsingGET has not yet been implemented")
		})
	}
	if api.ReportsGetUsingGETHandler == nil {
		api.ReportsGetUsingGETHandler = reports.GetUsingGETHandlerFunc(func(params reports.GetUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.GetUsingGET has not yet been implemented")
		})
	}
	if api.ConnectorGroupsManagementListGroupsUsingGETHandler == nil {
		api.ConnectorGroupsManagementListGroupsUsingGETHandler = connector_groups_management.ListGroupsUsingGETHandlerFunc(func(params connector_groups_management.ListGroupsUsingGETParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation connector_groups_management.ListGroupsUsingGET has not yet been implemented")
		})
	}
	if api.AssessmentReportsRerunReportUsingPOSTHandler == nil {
		api.AssessmentReportsRerunReportUsingPOSTHandler = assessment_reports.RerunReportUsingPOSTHandlerFunc(func(params assessment_reports.RerunReportUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation assessment_reports.RerunReportUsingPOST has not yet been implemented")
		})
	}
	if api.AzureConnectorRunConnectorUsingPOSTHandler == nil {
		api.AzureConnectorRunConnectorUsingPOSTHandler = azure_connector.RunConnectorUsingPOSTHandlerFunc(func(params azure_connector.RunConnectorUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.RunConnectorUsingPOST has not yet been implemented")
		})
	}
	if api.AwsConnectorRunConnectorUsingPOST1Handler == nil {
		api.AwsConnectorRunConnectorUsingPOST1Handler = a_w_s_connector.RunConnectorUsingPOST1HandlerFunc(func(params a_w_s_connector.RunConnectorUsingPOST1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.RunConnectorUsingPOST1 has not yet been implemented")
		})
	}
	if api.GcpConnectorRunConnectorUsingPOST2Handler == nil {
		api.GcpConnectorRunConnectorUsingPOST2Handler = g_c_p_connector.RunConnectorUsingPOST2HandlerFunc(func(params g_c_p_connector.RunConnectorUsingPOST2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.RunConnectorUsingPOST2 has not yet been implemented")
		})
	}
	if api.ResponseActionsTestSlackActionUsingPOSTHandler == nil {
		api.ResponseActionsTestSlackActionUsingPOSTHandler = response_actions.TestSlackActionUsingPOSTHandlerFunc(func(params response_actions.TestSlackActionUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.TestSlackActionUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsTestSlackActionUsingPOST1Handler == nil {
		api.ResponseActionsTestSlackActionUsingPOST1Handler = response_actions.TestSlackActionUsingPOST1HandlerFunc(func(params response_actions.TestSlackActionUsingPOST1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.TestSlackActionUsingPOST1 has not yet been implemented")
		})
	}
	if api.AzureConnectorUpdateConnectorUsingPUTHandler == nil {
		api.AzureConnectorUpdateConnectorUsingPUTHandler = azure_connector.UpdateConnectorUsingPUTHandlerFunc(func(params azure_connector.UpdateConnectorUsingPUTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation azure_connector.UpdateConnectorUsingPUT has not yet been implemented")
		})
	}
	if api.AwsConnectorUpdateConnectorUsingPUT1Handler == nil {
		api.AwsConnectorUpdateConnectorUsingPUT1Handler = a_w_s_connector.UpdateConnectorUsingPUT1HandlerFunc(func(params a_w_s_connector.UpdateConnectorUsingPUT1Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation a_w_s_connector.UpdateConnectorUsingPUT1 has not yet been implemented")
		})
	}
	if api.GcpConnectorUpdateConnectorUsingPUT2Handler == nil {
		api.GcpConnectorUpdateConnectorUsingPUT2Handler = g_c_p_connector.UpdateConnectorUsingPUT2HandlerFunc(func(params g_c_p_connector.UpdateConnectorUsingPUT2Params, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation g_c_p_connector.UpdateConnectorUsingPUT2 has not yet been implemented")
		})
	}
	if api.ResponseActionsUpdateEmailActionUsingPUTHandler == nil {
		api.ResponseActionsUpdateEmailActionUsingPUTHandler = response_actions.UpdateEmailActionUsingPUTHandlerFunc(func(params response_actions.UpdateEmailActionUsingPUTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.UpdateEmailActionUsingPUT has not yet been implemented")
		})
	}
	if api.UserAccessManagementUpdateGroupScopeForUserUsingPOSTHandler == nil {
		api.UserAccessManagementUpdateGroupScopeForUserUsingPOSTHandler = user_access_management.UpdateGroupScopeForUserUsingPOSTHandlerFunc(func(params user_access_management.UpdateGroupScopeForUserUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user_access_management.UpdateGroupScopeForUserUsingPOST has not yet been implemented")
		})
	}
	if api.ConnectorGroupsManagementUpdateGroupUsingPOSTHandler == nil {
		api.ConnectorGroupsManagementUpdateGroupUsingPOSTHandler = connector_groups_management.UpdateGroupUsingPOSTHandlerFunc(func(params connector_groups_management.UpdateGroupUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation connector_groups_management.UpdateGroupUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsUpdatePagerActionUsingPUTHandler == nil {
		api.ResponseActionsUpdatePagerActionUsingPUTHandler = response_actions.UpdatePagerActionUsingPUTHandlerFunc(func(params response_actions.UpdatePagerActionUsingPUTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.UpdatePagerActionUsingPUT has not yet been implemented")
		})
	}
	if api.ResponseRulesUpdateRuleUsingPUTHandler == nil {
		api.ResponseRulesUpdateRuleUsingPUTHandler = response_rules.UpdateRuleUsingPUTHandlerFunc(func(params response_rules.UpdateRuleUsingPUTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_rules.UpdateRuleUsingPUT has not yet been implemented")
		})
	}
	if api.UserAccessManagementUpdateScopeForUserUsingPOSTHandler == nil {
		api.UserAccessManagementUpdateScopeForUserUsingPOSTHandler = user_access_management.UpdateScopeForUserUsingPOSTHandlerFunc(func(params user_access_management.UpdateScopeForUserUsingPOSTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user_access_management.UpdateScopeForUserUsingPOST has not yet been implemented")
		})
	}
	if api.ResponseActionsUpdateSlackActionUsingPUTHandler == nil {
		api.ResponseActionsUpdateSlackActionUsingPUTHandler = response_actions.UpdateSlackActionUsingPUTHandlerFunc(func(params response_actions.UpdateSlackActionUsingPUTParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation response_actions.UpdateSlackActionUsingPUT has not yet been implemented")
		})
	}
	if api.ReportsUpdateUsingPATCHHandler == nil {
		api.ReportsUpdateUsingPATCHHandler = reports.UpdateUsingPATCHHandlerFunc(func(params reports.UpdateUsingPATCHParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation reports.UpdateUsingPATCH has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

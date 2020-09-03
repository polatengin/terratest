output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "loganalytics_workspace_name" {
    value = azurerm_log_analytics_workspace.logws.name
}

output "loganalytics_workspace_sku" {
    value = azurerm_log_analytics_workspace.logws.sku
}

output "loganalytics_workspace_retention" {
    value = azurerm_log_analytics_workspace.logws.retention_in_days
}

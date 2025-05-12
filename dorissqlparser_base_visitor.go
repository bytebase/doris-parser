// Code generated from DorisSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DorisSQLParser
import "github.com/antlr4-go/antlr/v4"

type BaseDorisSQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDorisSQLParserVisitor) VisitSqlStatements(ctx *SqlStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSingleStatement(ctx *SingleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitEmptyStatement(ctx *EmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropDbStatement(ctx *DropDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDataStmt(ctx *ShowDataStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnDesc(ctx *ColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDefaultDesc(ctx *DefaultDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIndexDesc(ctx *IndexDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitEngineDesc(ctx *EngineDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCharsetDesc(ctx *CharsetDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCollateDesc(ctx *CollateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitKeyDesc(ctx *KeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOrderByDesc(ctx *OrderByDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnNullable(ctx *ColumnNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAggStateDesc(ctx *AggStateDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAggDesc(ctx *AggDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRollupDesc(ctx *RollupDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRollupItem(ctx *RollupItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDupKeys(ctx *DupKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFromRollup(ctx *FromRollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOrReplace(ctx *OrReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropTableStatement(ctx *DropTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTableStatement(ctx *ShowTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDescTableStatement(ctx *DescTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropViewStatement(ctx *DropViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTaskClause(ctx *TaskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitKillStatement(ctx *KillStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSyncStatement(ctx *SyncStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTypeDesc(ctx *TypeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLocationsDesc(ctx *LocationsDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDictionaryName(ctx *DictionaryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterClause(ctx *AlterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddBackendClause(ctx *AddBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropBackendClause(ctx *DropBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateImageClause(ctx *CreateImageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropIndexClause(ctx *DropIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableRenameClause(ctx *TableRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSwapTableClause(ctx *SwapTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOptimizeRange(ctx *OptimizeRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOptimizeClause(ctx *OptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddColumnClause(ctx *AddColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropColumnClause(ctx *DropColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCompactionClause(ctx *CompactionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubfieldName(ctx *SubfieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNestedFieldName(ctx *NestedFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddFieldClause(ctx *AddFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropFieldClause(ctx *DropFieldClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropBranchClause(ctx *DropBranchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropTagClause(ctx *DropTagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableOperationClause(ctx *TableOperationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTagOptions(ctx *TagOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBranchOptions(ctx *BranchOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRefRetain(ctx *RefRetainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSnapshotId(ctx *SnapshotIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTimeUnit(ctx *TimeUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInteger_list(ctx *Integer_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataSource(ctx *DataSourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLoadProperties(ctx *LoadPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitImportColumns(ctx *ImportColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnProperties(ctx *ColumnPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitJobProperties(ctx *JobPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRegularColumns(ctx *RegularColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAllColumns(ctx *AllColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPredicateColumns(ctx *PredicateColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitHistogramStatement(ctx *HistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitClassifier(ctx *ClassifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInlineFunction(ctx *InlineFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataDescList(ctx *DataDescListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataDesc(ctx *DataDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFormatProps(ctx *FormatPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBrokerDesc(ctx *BrokerDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitResourceDesc(ctx *ResourceDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowProcStatement(ctx *ShowProcStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitHelpStatement(ctx *HelpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropUserStatement(ctx *DropUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowUserStatement(ctx *ShowUserStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnUser(ctx *GrantOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGrantOnAll(ctx *GrantOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivObjectName(ctx *PrivObjectNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBackupStatement(ctx *BackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRestoreStatement(ctx *RestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowExportStatement(ctx *ShowExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropFileStatement(ctx *DropFileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetStatement(ctx *SetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetNames(ctx *SetNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetUserVar(ctx *SetUserVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetSystemVar(ctx *SetSystemVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIsolation_level(ctx *Isolation_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIsolation_types(ctx *Isolation_typesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRoleList(ctx *RoleListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLock_item(ctx *Lock_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCommitStatement(ctx *CommitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTranslateStatement(ctx *TranslateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDialect(ctx *DialectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTranslateSQL(ctx *TranslateSQLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryStatement(ctx *QueryStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryRelation(ctx *QueryRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryNoWith(ctx *QueryNoWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryPeriod(ctx *QueryPeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPeriodType(ctx *PeriodTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetOperation(ctx *SetOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRowConstructor(ctx *RowConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSortItem(ctx *SortItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLimitConstExpr(ctx *LimitConstExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLimitElement(ctx *LimitElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFrom(ctx *FromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDual(ctx *DualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRollup(ctx *RollupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCube(ctx *CubeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGroupingSet(ctx *GroupingSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSetQuantifier(ctx *SetQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSelectSingle(ctx *SelectSingleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSelectAll(ctx *SelectAllContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExcludeClause(ctx *ExcludeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRelations(ctx *RelationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRelation(ctx *RelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableAtom(ctx *TableAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInlineTable(ctx *InlineTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableFunction(ctx *TableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFileTableFunction(ctx *FileTableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPivotValue(ctx *PivotValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNamedArguments(ctx *NamedArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitJoinRelation(ctx *JoinRelationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBracketHint(ctx *BracketHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitHintMap(ctx *HintMapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitJoinCriteria(ctx *JoinCriteriaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionNames(ctx *PartitionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTabletList(ctx *TabletListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPrepareSql(ctx *PrepareSqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitReplicaList(ctx *ReplicaListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMapExpressionList(ctx *MapExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMapExpression(ctx *MapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLogicalNot(ctx *LogicalNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLogicalBinary(ctx *LogicalBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIsNull(ctx *IsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInSubquery(ctx *InSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInList(ctx *InListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBetween(ctx *BetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLike(ctx *LikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDereference(ctx *DereferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMatchExpr(ctx *MatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitConvert(ctx *ConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCast(ctx *CastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSimpleCase(ctx *SimpleCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArrowExpression(ctx *ArrowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitConcat(ctx *ConcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCollate(ctx *CollateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArrayConstructor(ctx *ArrayConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMapConstructor(ctx *MapConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArraySlice(ctx *ArraySliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExists(ctx *ExistsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSearchedCase(ctx *SearchedCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNullLiteral(ctx *NullLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDateLiteral(ctx *DateLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExtract(ctx *ExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitGroupingOperation(ctx *GroupingOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInformationFunction(ctx *InformationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWindowFunction(ctx *WindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWhenClause(ctx *WhenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOver(ctx *OverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWindowFrame(ctx *WindowFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBoundedFrame(ctx *BoundedFrameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableDesc(ctx *TableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExplainDesc(ctx *ExplainDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionExpr(ctx *PartitionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionDesc(ctx *PartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitListPartitionValue(ctx *ListPartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionValueList(ctx *PartitionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitKeyPartition(ctx *KeyPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPartitionValue(ctx *PartitionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDistributionClause(ctx *DistributionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDistributionDesc(ctx *DistributionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStatusDesc(ctx *StatusDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitProperties(ctx *PropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitExtProperties(ctx *ExtPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserPropertyList(ctx *UserPropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInlineProperties(ctx *InlinePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInlineProperty(ctx *InlinePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitVarType(ctx *VarTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitOutfile(ctx *OutfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitFileFormat(ctx *FileFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBooleanValue(ctx *BooleanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTaskInterval(ctx *TaskIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnitBoundary(ctx *UnitBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubfieldDesc(ctx *SubfieldDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitStructType(ctx *StructTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTypeParameter(ctx *TypeParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBaseType(ctx *BaseTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDecimalType(ctx *DecimalTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitQualifiedName(ctx *QualifiedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitWriteBranch(ctx *WriteBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserWithHost(ctx *UserWithHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitAssignmentList(ctx *AssignmentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDecimalValue(ctx *DecimalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitDoubleValue(ctx *DoubleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitIntegerValue(ctx *IntegerValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDorisSQLParserVisitor) VisitNonReserved(ctx *NonReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

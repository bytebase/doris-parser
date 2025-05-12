// Code generated from DorisSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DorisSQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DorisSQLParserParser.
type DorisSQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DorisSQLParserParser#sqlStatements.
	VisitSqlStatements(ctx *SqlStatementsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#singleStatement.
	VisitSingleStatement(ctx *SingleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#useDatabaseStatement.
	VisitUseDatabaseStatement(ctx *UseDatabaseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#useCatalogStatement.
	VisitUseCatalogStatement(ctx *UseCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setCatalogStatement.
	VisitSetCatalogStatement(ctx *SetCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDatabasesStatement.
	VisitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterDbQuotaStatement.
	VisitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createDbStatement.
	VisitCreateDbStatement(ctx *CreateDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropDbStatement.
	VisitDropDbStatement(ctx *DropDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateDbStatement.
	VisitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterDatabaseRenameStatement.
	VisitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#recoverDbStmt.
	VisitRecoverDbStmt(ctx *RecoverDbStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDataStmt.
	VisitShowDataStmt(ctx *ShowDataStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDataDistributionStmt.
	VisitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createTableStatement.
	VisitCreateTableStatement(ctx *CreateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnDesc.
	VisitColumnDesc(ctx *ColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#defaultDesc.
	VisitDefaultDesc(ctx *DefaultDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#generatedColumnDesc.
	VisitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#indexDesc.
	VisitIndexDesc(ctx *IndexDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#engineDesc.
	VisitEngineDesc(ctx *EngineDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#charsetDesc.
	VisitCharsetDesc(ctx *CharsetDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#collateDesc.
	VisitCollateDesc(ctx *CollateDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#keyDesc.
	VisitKeyDesc(ctx *KeyDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#orderByDesc.
	VisitOrderByDesc(ctx *OrderByDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnNullable.
	VisitColumnNullable(ctx *ColumnNullableContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#typeWithNullable.
	VisitTypeWithNullable(ctx *TypeWithNullableContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#aggStateDesc.
	VisitAggStateDesc(ctx *AggStateDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#aggDesc.
	VisitAggDesc(ctx *AggDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rollupDesc.
	VisitRollupDesc(ctx *RollupDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rollupItem.
	VisitRollupItem(ctx *RollupItemContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dupKeys.
	VisitDupKeys(ctx *DupKeysContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#fromRollup.
	VisitFromRollup(ctx *FromRollupContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#orReplace.
	VisitOrReplace(ctx *OrReplaceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createTableAsSelectStatement.
	VisitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropTableStatement.
	VisitDropTableStatement(ctx *DropTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cleanTemporaryTableStatement.
	VisitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterTableStatement.
	VisitAlterTableStatement(ctx *AlterTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createIndexStatement.
	VisitCreateIndexStatement(ctx *CreateIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropIndexStatement.
	VisitDropIndexStatement(ctx *DropIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTableStatement.
	VisitShowTableStatement(ctx *ShowTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTemporaryTablesStatement.
	VisitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateTableStatement.
	VisitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showColumnStatement.
	VisitShowColumnStatement(ctx *ShowColumnStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTableStatusStatement.
	VisitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#refreshTableStatement.
	VisitRefreshTableStatement(ctx *RefreshTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showAlterStatement.
	VisitShowAlterStatement(ctx *ShowAlterStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#descTableStatement.
	VisitDescTableStatement(ctx *DescTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createTableLikeStatement.
	VisitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showIndexStatement.
	VisitShowIndexStatement(ctx *ShowIndexStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#recoverTableStatement.
	VisitRecoverTableStatement(ctx *RecoverTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#truncateTableStatement.
	VisitTruncateTableStatement(ctx *TruncateTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelAlterTableStatement.
	VisitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showPartitionsStatement.
	VisitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#recoverPartitionStatement.
	VisitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createViewStatement.
	VisitCreateViewStatement(ctx *CreateViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterViewStatement.
	VisitAlterViewStatement(ctx *AlterViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropViewStatement.
	VisitDropViewStatement(ctx *DropViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnNameWithComment.
	VisitColumnNameWithComment(ctx *ColumnNameWithCommentContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#submitTaskStatement.
	VisitSubmitTaskStatement(ctx *SubmitTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#taskClause.
	VisitTaskClause(ctx *TaskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropTaskStatement.
	VisitDropTaskStatement(ctx *DropTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#taskScheduleDesc.
	VisitTaskScheduleDesc(ctx *TaskScheduleDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createMaterializedViewStatement.
	VisitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#mvPartitionExprs.
	VisitMvPartitionExprs(ctx *MvPartitionExprsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#materializedViewDesc.
	VisitMaterializedViewDesc(ctx *MaterializedViewDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showMaterializedViewsStatement.
	VisitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropMaterializedViewStatement.
	VisitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterMaterializedViewStatement.
	VisitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#refreshMaterializedViewStatement.
	VisitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelRefreshMaterializedViewStatement.
	VisitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminSetConfigStatement.
	VisitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminSetReplicaStatusStatement.
	VisitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminShowConfigStatement.
	VisitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminShowReplicaDistributionStatement.
	VisitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminShowReplicaStatusStatement.
	VisitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminRepairTableStatement.
	VisitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminCancelRepairTableStatement.
	VisitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminCheckTabletsStatement.
	VisitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminSetPartitionVersion.
	VisitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#killStatement.
	VisitKillStatement(ctx *KillStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#syncStatement.
	VisitSyncStatement(ctx *SyncStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminSetAutomatedSnapshotOnStatement.
	VisitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#adminSetAutomatedSnapshotOffStatement.
	VisitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterSystemStatement.
	VisitAlterSystemStatement(ctx *AlterSystemStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelAlterSystemStatement.
	VisitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showComputeNodesStatement.
	VisitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createExternalCatalogStatement.
	VisitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateExternalCatalogStatement.
	VisitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropExternalCatalogStatement.
	VisitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCatalogsStatement.
	VisitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterCatalogStatement.
	VisitAlterCatalogStatement(ctx *AlterCatalogStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createStorageVolumeStatement.
	VisitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#typeDesc.
	VisitTypeDesc(ctx *TypeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#locationsDesc.
	VisitLocationsDesc(ctx *LocationsDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showStorageVolumesStatement.
	VisitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropStorageVolumeStatement.
	VisitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterStorageVolumeStatement.
	VisitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterStorageVolumeClause.
	VisitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyStorageVolumePropertiesClause.
	VisitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyStorageVolumeCommentClause.
	VisitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#descStorageVolumeStatement.
	VisitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setDefaultStorageVolumeStatement.
	VisitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#updateFailPointStatusStatement.
	VisitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showFailPointStatement.
	VisitShowFailPointStatement(ctx *ShowFailPointStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createDictionaryStatement.
	VisitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropDictionaryStatement.
	VisitDropDictionaryStatement(ctx *DropDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#refreshDictionaryStatement.
	VisitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDictionaryStatement.
	VisitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelRefreshDictionaryStatement.
	VisitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dictionaryColumnDesc.
	VisitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dictionaryName.
	VisitDictionaryName(ctx *DictionaryNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterClause.
	VisitAlterClause(ctx *AlterClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addFrontendClause.
	VisitAddFrontendClause(ctx *AddFrontendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropFrontendClause.
	VisitDropFrontendClause(ctx *DropFrontendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyFrontendHostClause.
	VisitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addBackendClause.
	VisitAddBackendClause(ctx *AddBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropBackendClause.
	VisitDropBackendClause(ctx *DropBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#decommissionBackendClause.
	VisitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyBackendClause.
	VisitModifyBackendClause(ctx *ModifyBackendClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addComputeNodeClause.
	VisitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropComputeNodeClause.
	VisitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyBrokerClause.
	VisitModifyBrokerClause(ctx *ModifyBrokerClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterLoadErrorUrlClause.
	VisitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createImageClause.
	VisitCreateImageClause(ctx *CreateImageClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cleanTabletSchedQClause.
	VisitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#decommissionDiskClause.
	VisitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelDecommissionDiskClause.
	VisitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#disableDiskClause.
	VisitDisableDiskClause(ctx *DisableDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelDisableDiskClause.
	VisitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createIndexClause.
	VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropIndexClause.
	VisitDropIndexClause(ctx *DropIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableRenameClause.
	VisitTableRenameClause(ctx *TableRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#swapTableClause.
	VisitSwapTableClause(ctx *SwapTableClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyPropertiesClause.
	VisitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyCommentClause.
	VisitModifyCommentClause(ctx *ModifyCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#optimizeRange.
	VisitOptimizeRange(ctx *OptimizeRangeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#optimizeClause.
	VisitOptimizeClause(ctx *OptimizeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addColumnClause.
	VisitAddColumnClause(ctx *AddColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addColumnsClause.
	VisitAddColumnsClause(ctx *AddColumnsClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropColumnClause.
	VisitDropColumnClause(ctx *DropColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyColumnClause.
	VisitModifyColumnClause(ctx *ModifyColumnClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyColumnCommentClause.
	VisitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnRenameClause.
	VisitColumnRenameClause(ctx *ColumnRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#reorderColumnsClause.
	VisitReorderColumnsClause(ctx *ReorderColumnsClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rollupRenameClause.
	VisitRollupRenameClause(ctx *RollupRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#compactionClause.
	VisitCompactionClause(ctx *CompactionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subfieldName.
	VisitSubfieldName(ctx *SubfieldNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#nestedFieldName.
	VisitNestedFieldName(ctx *NestedFieldNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addFieldClause.
	VisitAddFieldClause(ctx *AddFieldClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropFieldClause.
	VisitDropFieldClause(ctx *DropFieldClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createOrReplaceTagClause.
	VisitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createOrReplaceBranchClause.
	VisitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropBranchClause.
	VisitDropBranchClause(ctx *DropBranchClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropTagClause.
	VisitDropTagClause(ctx *DropTagClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableOperationClause.
	VisitTableOperationClause(ctx *TableOperationClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tagOptions.
	VisitTagOptions(ctx *TagOptionsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#branchOptions.
	VisitBranchOptions(ctx *BranchOptionsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#snapshotRetention.
	VisitSnapshotRetention(ctx *SnapshotRetentionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#refRetain.
	VisitRefRetain(ctx *RefRetainContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#maxSnapshotAge.
	VisitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#minSnapshotsToKeep.
	VisitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#snapshotId.
	VisitSnapshotId(ctx *SnapshotIdContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#timeUnit.
	VisitTimeUnit(ctx *TimeUnitContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#integer_list.
	VisitInteger_list(ctx *Integer_listContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropPersistentIndexClause.
	VisitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addPartitionClause.
	VisitAddPartitionClause(ctx *AddPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropPartitionClause.
	VisitDropPartitionClause(ctx *DropPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#truncatePartitionClause.
	VisitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#modifyPartitionClause.
	VisitModifyPartitionClause(ctx *ModifyPartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#replacePartitionClause.
	VisitReplacePartitionClause(ctx *ReplacePartitionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionRenameClause.
	VisitPartitionRenameClause(ctx *PartitionRenameClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#insertLabelOrColumnAliases.
	VisitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnAliasesOrByName.
	VisitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createRoutineLoadStatement.
	VisitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterRoutineLoadStatement.
	VisitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataSource.
	VisitDataSource(ctx *DataSourceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#loadProperties.
	VisitLoadProperties(ctx *LoadPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#colSeparatorProperty.
	VisitColSeparatorProperty(ctx *ColSeparatorPropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rowDelimiterProperty.
	VisitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#importColumns.
	VisitImportColumns(ctx *ImportColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnProperties.
	VisitColumnProperties(ctx *ColumnPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#jobProperties.
	VisitJobProperties(ctx *JobPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataSourceProperties.
	VisitDataSourceProperties(ctx *DataSourcePropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#stopRoutineLoadStatement.
	VisitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#resumeRoutineLoadStatement.
	VisitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#pauseRoutineLoadStatement.
	VisitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRoutineLoadStatement.
	VisitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRoutineLoadTaskStatement.
	VisitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateRoutineLoadStatement.
	VisitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showStreamLoadStatement.
	VisitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#analyzeStatement.
	VisitAnalyzeStatement(ctx *AnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#regularColumns.
	VisitRegularColumns(ctx *RegularColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#allColumns.
	VisitAllColumns(ctx *AllColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#predicateColumns.
	VisitPredicateColumns(ctx *PredicateColumnsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#multiColumnSet.
	VisitMultiColumnSet(ctx *MultiColumnSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropStatsStatement.
	VisitDropStatsStatement(ctx *DropStatsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#histogramStatement.
	VisitHistogramStatement(ctx *HistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#analyzeHistogramStatement.
	VisitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropHistogramStatement.
	VisitDropHistogramStatement(ctx *DropHistogramStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createAnalyzeStatement.
	VisitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropAnalyzeJobStatement.
	VisitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showAnalyzeStatement.
	VisitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showStatsMetaStatement.
	VisitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showHistogramMetaStatement.
	VisitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#killAnalyzeStatement.
	VisitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#analyzeProfileStatement.
	VisitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createBaselinePlanStatement.
	VisitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropBaselinePlanStatement.
	VisitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showBaselinePlanStatement.
	VisitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createResourceGroupStatement.
	VisitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropResourceGroupStatement.
	VisitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterResourceGroupStatement.
	VisitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showResourceGroupStatement.
	VisitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showResourceGroupUsageStatement.
	VisitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createResourceStatement.
	VisitCreateResourceStatement(ctx *CreateResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterResourceStatement.
	VisitAlterResourceStatement(ctx *AlterResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropResourceStatement.
	VisitDropResourceStatement(ctx *DropResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showResourceStatement.
	VisitShowResourceStatement(ctx *ShowResourceStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#classifier.
	VisitClassifier(ctx *ClassifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showFunctionsStatement.
	VisitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropFunctionStatement.
	VisitDropFunctionStatement(ctx *DropFunctionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createFunctionStatement.
	VisitCreateFunctionStatement(ctx *CreateFunctionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inlineFunction.
	VisitInlineFunction(ctx *InlineFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#typeList.
	VisitTypeList(ctx *TypeListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataDescList.
	VisitDataDescList(ctx *DataDescListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataDesc.
	VisitDataDesc(ctx *DataDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#formatProps.
	VisitFormatProps(ctx *FormatPropsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#brokerDesc.
	VisitBrokerDesc(ctx *BrokerDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#resourceDesc.
	VisitResourceDesc(ctx *ResourceDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showLoadStatement.
	VisitShowLoadStatement(ctx *ShowLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showLoadWarningsStatement.
	VisitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelLoadStatement.
	VisitCancelLoadStatement(ctx *CancelLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterLoadStatement.
	VisitAlterLoadStatement(ctx *AlterLoadStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelCompactionStatement.
	VisitCancelCompactionStatement(ctx *CancelCompactionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showAuthorStatement.
	VisitShowAuthorStatement(ctx *ShowAuthorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showBackendsStatement.
	VisitShowBackendsStatement(ctx *ShowBackendsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showBrokerStatement.
	VisitShowBrokerStatement(ctx *ShowBrokerStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCharsetStatement.
	VisitShowCharsetStatement(ctx *ShowCharsetStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCollationStatement.
	VisitShowCollationStatement(ctx *ShowCollationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDeleteStatement.
	VisitShowDeleteStatement(ctx *ShowDeleteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDynamicPartitionStatement.
	VisitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showEventsStatement.
	VisitShowEventsStatement(ctx *ShowEventsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showEnginesStatement.
	VisitShowEnginesStatement(ctx *ShowEnginesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showFrontendsStatement.
	VisitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showPluginsStatement.
	VisitShowPluginsStatement(ctx *ShowPluginsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRepositoriesStatement.
	VisitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showOpenTableStatement.
	VisitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showPrivilegesStatement.
	VisitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showProcedureStatement.
	VisitShowProcedureStatement(ctx *ShowProcedureStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showProcStatement.
	VisitShowProcStatement(ctx *ShowProcStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showProcesslistStatement.
	VisitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showProfilelistStatement.
	VisitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRunningQueriesStatement.
	VisitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showStatusStatement.
	VisitShowStatusStatement(ctx *ShowStatusStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTabletStatement.
	VisitShowTabletStatement(ctx *ShowTabletStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTransactionStatement.
	VisitShowTransactionStatement(ctx *ShowTransactionStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showTriggersStatement.
	VisitShowTriggersStatement(ctx *ShowTriggersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showUserPropertyStatement.
	VisitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showVariablesStatement.
	VisitShowVariablesStatement(ctx *ShowVariablesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showWarningStatement.
	VisitShowWarningStatement(ctx *ShowWarningStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#helpStatement.
	VisitHelpStatement(ctx *HelpStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createUserStatement.
	VisitCreateUserStatement(ctx *CreateUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropUserStatement.
	VisitDropUserStatement(ctx *DropUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterUserStatement.
	VisitAlterUserStatement(ctx *AlterUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showUserStatement.
	VisitShowUserStatement(ctx *ShowUserStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showAllAuthentication.
	VisitShowAllAuthentication(ctx *ShowAllAuthenticationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showAuthenticationForUser.
	VisitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#executeAsStatement.
	VisitExecuteAsStatement(ctx *ExecuteAsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createRoleStatement.
	VisitCreateRoleStatement(ctx *CreateRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterRoleStatement.
	VisitAlterRoleStatement(ctx *AlterRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropRoleStatement.
	VisitDropRoleStatement(ctx *DropRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRolesStatement.
	VisitShowRolesStatement(ctx *ShowRolesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantRoleToUser.
	VisitGrantRoleToUser(ctx *GrantRoleToUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantRoleToRole.
	VisitGrantRoleToRole(ctx *GrantRoleToRoleContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeRoleFromUser.
	VisitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeRoleFromRole.
	VisitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setRoleStatement.
	VisitSetRoleStatement(ctx *SetRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setDefaultRoleStatement.
	VisitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantRevokeClause.
	VisitGrantRevokeClause(ctx *GrantRevokeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnUser.
	VisitGrantOnUser(ctx *GrantOnUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnTableBrief.
	VisitGrantOnTableBrief(ctx *GrantOnTableBriefContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnFunc.
	VisitGrantOnFunc(ctx *GrantOnFuncContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnSystem.
	VisitGrantOnSystem(ctx *GrantOnSystemContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnPrimaryObj.
	VisitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#grantOnAll.
	VisitGrantOnAll(ctx *GrantOnAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnUser.
	VisitRevokeOnUser(ctx *RevokeOnUserContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnTableBrief.
	VisitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnFunc.
	VisitRevokeOnFunc(ctx *RevokeOnFuncContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnSystem.
	VisitRevokeOnSystem(ctx *RevokeOnSystemContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnPrimaryObj.
	VisitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#revokeOnAll.
	VisitRevokeOnAll(ctx *RevokeOnAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showGrantsStatement.
	VisitShowGrantsStatement(ctx *ShowGrantsStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#authWithoutPlugin.
	VisitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#authWithPlugin.
	VisitAuthWithPlugin(ctx *AuthWithPluginContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privObjectName.
	VisitPrivObjectName(ctx *PrivObjectNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privObjectNameList.
	VisitPrivObjectNameList(ctx *PrivObjectNameListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privFunctionObjectNameList.
	VisitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privilegeTypeList.
	VisitPrivilegeTypeList(ctx *PrivilegeTypeListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privObjectType.
	VisitPrivObjectType(ctx *PrivObjectTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#privObjectTypePlural.
	VisitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createSecurityIntegrationStatement.
	VisitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterSecurityIntegrationStatement.
	VisitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropSecurityIntegrationStatement.
	VisitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showSecurityIntegrationStatement.
	VisitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateSecurityIntegrationStatement.
	VisitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createGroupProviderStatement.
	VisitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropGroupProviderStatement.
	VisitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showGroupProvidersStatement.
	VisitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showCreateGroupProviderStatement.
	VisitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#backupStatement.
	VisitBackupStatement(ctx *BackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelBackupStatement.
	VisitCancelBackupStatement(ctx *CancelBackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showBackupStatement.
	VisitShowBackupStatement(ctx *ShowBackupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#restoreStatement.
	VisitRestoreStatement(ctx *RestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelRestoreStatement.
	VisitCancelRestoreStatement(ctx *CancelRestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showRestoreStatement.
	VisitShowRestoreStatement(ctx *ShowRestoreStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showSnapshotStatement.
	VisitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createRepositoryStatement.
	VisitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropRepositoryStatement.
	VisitDropRepositoryStatement(ctx *DropRepositoryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addSqlBlackListStatement.
	VisitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#delSqlBlackListStatement.
	VisitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showSqlBlackListStatement.
	VisitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showWhiteListStatement.
	VisitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#addBackendBlackListStatement.
	VisitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#delBackendBlackListStatement.
	VisitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showBackendBlackListStatement.
	VisitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataCacheTarget.
	VisitDataCacheTarget(ctx *DataCacheTargetContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createDataCacheRuleStatement.
	VisitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showDataCacheRulesStatement.
	VisitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropDataCacheRuleStatement.
	VisitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#clearDataCacheRulesStatement.
	VisitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dataCacheSelectStatement.
	VisitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#exportStatement.
	VisitExportStatement(ctx *ExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cancelExportStatement.
	VisitCancelExportStatement(ctx *CancelExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showExportStatement.
	VisitShowExportStatement(ctx *ShowExportStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#installPluginStatement.
	VisitInstallPluginStatement(ctx *InstallPluginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#uninstallPluginStatement.
	VisitUninstallPluginStatement(ctx *UninstallPluginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createFileStatement.
	VisitCreateFileStatement(ctx *CreateFileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropFileStatement.
	VisitDropFileStatement(ctx *DropFileStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showSmallFilesStatement.
	VisitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createPipeStatement.
	VisitCreatePipeStatement(ctx *CreatePipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropPipeStatement.
	VisitDropPipeStatement(ctx *DropPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterPipeClause.
	VisitAlterPipeClause(ctx *AlterPipeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterPipeStatement.
	VisitAlterPipeStatement(ctx *AlterPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#descPipeStatement.
	VisitDescPipeStatement(ctx *DescPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showPipeStatement.
	VisitShowPipeStatement(ctx *ShowPipeStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setStatement.
	VisitSetStatement(ctx *SetStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setNames.
	VisitSetNames(ctx *SetNamesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setUserVar.
	VisitSetUserVar(ctx *SetUserVarContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setSystemVar.
	VisitSetSystemVar(ctx *SetSystemVarContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#transaction_characteristics.
	VisitTransaction_characteristics(ctx *Transaction_characteristicsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#transaction_access_mode.
	VisitTransaction_access_mode(ctx *Transaction_access_modeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#isolation_level.
	VisitIsolation_level(ctx *Isolation_levelContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#isolation_types.
	VisitIsolation_types(ctx *Isolation_typesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setUserPropertyStatement.
	VisitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#roleList.
	VisitRoleList(ctx *RoleListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#executeScriptStatement.
	VisitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unsupportedStatement.
	VisitUnsupportedStatement(ctx *UnsupportedStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#lock_item.
	VisitLock_item(ctx *Lock_itemContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#lock_type.
	VisitLock_type(ctx *Lock_typeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterPlanAdvisorAddStatement.
	VisitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#truncatePlanAdvisorStatement.
	VisitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterPlanAdvisorDropStatement.
	VisitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showPlanAdvisorStatement.
	VisitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createWarehouseStatement.
	VisitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropWarehouseStatement.
	VisitDropWarehouseStatement(ctx *DropWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#suspendWarehouseStatement.
	VisitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#resumeWarehouseStatement.
	VisitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setWarehouseStatement.
	VisitSetWarehouseStatement(ctx *SetWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showWarehousesStatement.
	VisitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showClustersStatement.
	VisitShowClustersStatement(ctx *ShowClustersStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#showNodesStatement.
	VisitShowNodesStatement(ctx *ShowNodesStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterWarehouseStatement.
	VisitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#createCNGroupStatement.
	VisitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dropCNGroupStatement.
	VisitDropCNGroupStatement(ctx *DropCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#enableCNGroupStatement.
	VisitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#disableCNGroupStatement.
	VisitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#alterCNGroupStatement.
	VisitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#commitStatement.
	VisitCommitStatement(ctx *CommitStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#translateStatement.
	VisitTranslateStatement(ctx *TranslateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dialect.
	VisitDialect(ctx *DialectContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#translateSQL.
	VisitTranslateSQL(ctx *TranslateSQLContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryStatement.
	VisitQueryStatement(ctx *QueryStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryRelation.
	VisitQueryRelation(ctx *QueryRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryNoWith.
	VisitQueryNoWith(ctx *QueryNoWithContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryPeriod.
	VisitQueryPeriod(ctx *QueryPeriodContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#periodType.
	VisitPeriodType(ctx *PeriodTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryWithParentheses.
	VisitQueryWithParentheses(ctx *QueryWithParenthesesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setOperation.
	VisitSetOperation(ctx *SetOperationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#queryPrimaryDefault.
	VisitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rowConstructor.
	VisitRowConstructor(ctx *RowConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#sortItem.
	VisitSortItem(ctx *SortItemContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#limitConstExpr.
	VisitLimitConstExpr(ctx *LimitConstExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#limitElement.
	VisitLimitElement(ctx *LimitElementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#from.
	VisitFrom(ctx *FromContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dual.
	VisitDual(ctx *DualContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rollup.
	VisitRollup(ctx *RollupContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cube.
	VisitCube(ctx *CubeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#multipleGroupingSets.
	VisitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#singleGroupingSet.
	VisitSingleGroupingSet(ctx *SingleGroupingSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#groupingSet.
	VisitGroupingSet(ctx *GroupingSetContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#commonTableExpression.
	VisitCommonTableExpression(ctx *CommonTableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#setQuantifier.
	VisitSetQuantifier(ctx *SetQuantifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#selectSingle.
	VisitSelectSingle(ctx *SelectSingleContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#selectAll.
	VisitSelectAll(ctx *SelectAllContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#excludeClause.
	VisitExcludeClause(ctx *ExcludeClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#relations.
	VisitRelations(ctx *RelationsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#relation.
	VisitRelation(ctx *RelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableAtom.
	VisitTableAtom(ctx *TableAtomContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inlineTable.
	VisitInlineTable(ctx *InlineTableContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subqueryWithAlias.
	VisitSubqueryWithAlias(ctx *SubqueryWithAliasContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableFunction.
	VisitTableFunction(ctx *TableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#normalizedTableFunction.
	VisitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#fileTableFunction.
	VisitFileTableFunction(ctx *FileTableFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#parenthesizedRelation.
	VisitParenthesizedRelation(ctx *ParenthesizedRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#pivotClause.
	VisitPivotClause(ctx *PivotClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#pivotAggregationExpression.
	VisitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#pivotValue.
	VisitPivotValue(ctx *PivotValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#namedArgumentList.
	VisitNamedArgumentList(ctx *NamedArgumentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#namedArguments.
	VisitNamedArguments(ctx *NamedArgumentsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#joinRelation.
	VisitJoinRelation(ctx *JoinRelationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#crossOrInnerJoinType.
	VisitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#outerAndSemiJoinType.
	VisitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#bracketHint.
	VisitBracketHint(ctx *BracketHintContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#hintMap.
	VisitHintMap(ctx *HintMapContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#joinCriteria.
	VisitJoinCriteria(ctx *JoinCriteriaContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnAliases.
	VisitColumnAliases(ctx *ColumnAliasesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionNames.
	VisitPartitionNames(ctx *PartitionNamesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#keyPartitionList.
	VisitKeyPartitionList(ctx *KeyPartitionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tabletList.
	VisitTabletList(ctx *TabletListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#prepareSql.
	VisitPrepareSql(ctx *PrepareSqlContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#deallocateStatement.
	VisitDeallocateStatement(ctx *DeallocateStatementContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#replicaList.
	VisitReplicaList(ctx *ReplicaListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#expressionsWithDefault.
	VisitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#mapExpressionList.
	VisitMapExpressionList(ctx *MapExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#mapExpression.
	VisitMapExpression(ctx *MapExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#expressionSingleton.
	VisitExpressionSingleton(ctx *ExpressionSingletonContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#expressionDefault.
	VisitExpressionDefault(ctx *ExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#logicalNot.
	VisitLogicalNot(ctx *LogicalNotContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#logicalBinary.
	VisitLogicalBinary(ctx *LogicalBinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#booleanExpressionDefault.
	VisitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#isNull.
	VisitIsNull(ctx *IsNullContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#scalarSubquery.
	VisitScalarSubquery(ctx *ScalarSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tupleInSubquery.
	VisitTupleInSubquery(ctx *TupleInSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inSubquery.
	VisitInSubquery(ctx *InSubqueryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inList.
	VisitInList(ctx *InListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#between.
	VisitBetween(ctx *BetweenContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#like.
	VisitLike(ctx *LikeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#valueExpressionDefault.
	VisitValueExpressionDefault(ctx *ValueExpressionDefaultContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arithmeticBinary.
	VisitArithmeticBinary(ctx *ArithmeticBinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dereference.
	VisitDereference(ctx *DereferenceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#odbcFunctionCallExpression.
	VisitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#matchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#convert.
	VisitConvert(ctx *ConvertContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#collectionSubscript.
	VisitCollectionSubscript(ctx *CollectionSubscriptContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#cast.
	VisitCast(ctx *CastContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#parenthesizedExpression.
	VisitParenthesizedExpression(ctx *ParenthesizedExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userVariableExpression.
	VisitUserVariableExpression(ctx *UserVariableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#functionCallExpression.
	VisitFunctionCallExpression(ctx *FunctionCallExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#simpleCase.
	VisitSimpleCase(ctx *SimpleCaseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arrowExpression.
	VisitArrowExpression(ctx *ArrowExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#systemVariableExpression.
	VisitSystemVariableExpression(ctx *SystemVariableExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#concat.
	VisitConcat(ctx *ConcatContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subqueryExpression.
	VisitSubqueryExpression(ctx *SubqueryExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#lambdaFunctionExpr.
	VisitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dictionaryGetExpr.
	VisitDictionaryGetExpr(ctx *DictionaryGetExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#collate.
	VisitCollate(ctx *CollateContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arrayConstructor.
	VisitArrayConstructor(ctx *ArrayConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#mapConstructor.
	VisitMapConstructor(ctx *MapConstructorContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arraySlice.
	VisitArraySlice(ctx *ArraySliceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#exists.
	VisitExists(ctx *ExistsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#searchedCase.
	VisitSearchedCase(ctx *SearchedCaseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arithmeticUnary.
	VisitArithmeticUnary(ctx *ArithmeticUnaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#nullLiteral.
	VisitNullLiteral(ctx *NullLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#dateLiteral.
	VisitDateLiteral(ctx *DateLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#intervalLiteral.
	VisitIntervalLiteral(ctx *IntervalLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unitBoundaryLiteral.
	VisitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#binaryLiteral.
	VisitBinaryLiteral(ctx *BinaryLiteralContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#Parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#extract.
	VisitExtract(ctx *ExtractContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#groupingOperation.
	VisitGroupingOperation(ctx *GroupingOperationContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#informationFunction.
	VisitInformationFunction(ctx *InformationFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#specialDateTime.
	VisitSpecialDateTime(ctx *SpecialDateTimeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#aggregationFunctionCall.
	VisitAggregationFunctionCall(ctx *AggregationFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#windowFunctionCall.
	VisitWindowFunctionCall(ctx *WindowFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#translateFunctionCall.
	VisitTranslateFunctionCall(ctx *TranslateFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#informationFunctionExpression.
	VisitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#specialDateTimeExpression.
	VisitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#specialFunctionExpression.
	VisitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#whenClause.
	VisitWhenClause(ctx *WhenClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#over.
	VisitOver(ctx *OverContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#ignoreNulls.
	VisitIgnoreNulls(ctx *IgnoreNullsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#windowFrame.
	VisitWindowFrame(ctx *WindowFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unboundedFrame.
	VisitUnboundedFrame(ctx *UnboundedFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#currentRowBound.
	VisitCurrentRowBound(ctx *CurrentRowBoundContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#boundedFrame.
	VisitBoundedFrame(ctx *BoundedFrameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#backupRestoreObjectDesc.
	VisitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableDesc.
	VisitTableDesc(ctx *TableDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#backupRestoreTableDesc.
	VisitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#explainDesc.
	VisitExplainDesc(ctx *ExplainDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#optimizerTrace.
	VisitOptimizerTrace(ctx *OptimizerTraceContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionExpr.
	VisitPartitionExpr(ctx *PartitionExprContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionDesc.
	VisitPartitionDesc(ctx *PartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#listPartitionDesc.
	VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#singleItemListPartitionDesc.
	VisitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#multiItemListPartitionDesc.
	VisitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#multiListPartitionValues.
	VisitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#singleListPartitionValues.
	VisitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#listPartitionValues.
	VisitListPartitionValues(ctx *ListPartitionValuesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#listPartitionValue.
	VisitListPartitionValue(ctx *ListPartitionValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#literalExpressionList.
	VisitLiteralExpressionList(ctx *LiteralExpressionListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#rangePartitionDesc.
	VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#singleRangePartition.
	VisitSingleRangePartition(ctx *SingleRangePartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#multiRangePartition.
	VisitMultiRangePartition(ctx *MultiRangePartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionRangeDesc.
	VisitPartitionRangeDesc(ctx *PartitionRangeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionKeyDesc.
	VisitPartitionKeyDesc(ctx *PartitionKeyDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionValueList.
	VisitPartitionValueList(ctx *PartitionValueListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#keyPartition.
	VisitKeyPartition(ctx *KeyPartitionContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#partitionValue.
	VisitPartitionValue(ctx *PartitionValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#distributionClause.
	VisitDistributionClause(ctx *DistributionClauseContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#distributionDesc.
	VisitDistributionDesc(ctx *DistributionDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#refreshSchemeDesc.
	VisitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#statusDesc.
	VisitStatusDesc(ctx *StatusDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#properties.
	VisitProperties(ctx *PropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#extProperties.
	VisitExtProperties(ctx *ExtPropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#propertyList.
	VisitPropertyList(ctx *PropertyListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userPropertyList.
	VisitUserPropertyList(ctx *UserPropertyListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inlineProperties.
	VisitInlineProperties(ctx *InlinePropertiesContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#inlineProperty.
	VisitInlineProperty(ctx *InlinePropertyContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#varType.
	VisitVarType(ctx *VarTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#outfile.
	VisitOutfile(ctx *OutfileContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#fileFormat.
	VisitFileFormat(ctx *FileFormatContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#booleanValue.
	VisitBooleanValue(ctx *BooleanValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#taskInterval.
	VisitTaskInterval(ctx *TaskIntervalContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#taskUnitIdentifier.
	VisitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unitIdentifier.
	VisitUnitIdentifier(ctx *UnitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unitBoundary.
	VisitUnitBoundary(ctx *UnitBoundaryContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subfieldDesc.
	VisitSubfieldDesc(ctx *SubfieldDescContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#subfieldDescs.
	VisitSubfieldDescs(ctx *SubfieldDescsContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#structType.
	VisitStructType(ctx *StructTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#typeParameter.
	VisitTypeParameter(ctx *TypeParameterContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#decimalType.
	VisitDecimalType(ctx *DecimalTypeContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#qualifiedName.
	VisitQualifiedName(ctx *QualifiedNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#writeBranch.
	VisitWriteBranch(ctx *WriteBranchContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#unquotedIdentifier.
	VisitUnquotedIdentifier(ctx *UnquotedIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#digitIdentifier.
	VisitDigitIdentifier(ctx *DigitIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#backQuotedIdentifier.
	VisitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierWithAlias.
	VisitIdentifierWithAlias(ctx *IdentifierWithAliasContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierWithAliasList.
	VisitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierOrString.
	VisitIdentifierOrString(ctx *IdentifierOrStringContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierOrStringList.
	VisitIdentifierOrStringList(ctx *IdentifierOrStringListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#identifierOrStringOrStar.
	VisitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userWithoutHost.
	VisitUserWithoutHost(ctx *UserWithoutHostContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userWithHost.
	VisitUserWithHost(ctx *UserWithHostContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#userWithHostAndBlanket.
	VisitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#assignmentList.
	VisitAssignmentList(ctx *AssignmentListContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#decimalValue.
	VisitDecimalValue(ctx *DecimalValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#doubleValue.
	VisitDoubleValue(ctx *DoubleValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#integerValue.
	VisitIntegerValue(ctx *IntegerValueContext) interface{}

	// Visit a parse tree produced by DorisSQLParserParser#nonReserved.
	VisitNonReserved(ctx *NonReservedContext) interface{}
}

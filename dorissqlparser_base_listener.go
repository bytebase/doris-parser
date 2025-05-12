// Code generated from DorisSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // DorisSQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseDorisSQLParserListener is a complete listener for a parse tree produced by DorisSQLParserParser.
type BaseDorisSQLParserListener struct{}

var _ DorisSQLParserListener = &BaseDorisSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDorisSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDorisSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDorisSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDorisSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSqlStatements is called when production sqlStatements is entered.
func (s *BaseDorisSQLParserListener) EnterSqlStatements(ctx *SqlStatementsContext) {}

// ExitSqlStatements is called when production sqlStatements is exited.
func (s *BaseDorisSQLParserListener) ExitSqlStatements(ctx *SqlStatementsContext) {}

// EnterSingleStatement is called when production singleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSingleStatement(ctx *SingleStatementContext) {}

// ExitSingleStatement is called when production singleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSingleStatement(ctx *SingleStatementContext) {}

// EnterEmptyStatement is called when production emptyStatement is entered.
func (s *BaseDorisSQLParserListener) EnterEmptyStatement(ctx *EmptyStatementContext) {}

// ExitEmptyStatement is called when production emptyStatement is exited.
func (s *BaseDorisSQLParserListener) ExitEmptyStatement(ctx *EmptyStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDorisSQLParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDorisSQLParserListener) ExitStatement(ctx *StatementContext) {}

// EnterUseDatabaseStatement is called when production useDatabaseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// ExitUseDatabaseStatement is called when production useDatabaseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUseDatabaseStatement(ctx *UseDatabaseStatementContext) {}

// EnterUseCatalogStatement is called when production useCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// ExitUseCatalogStatement is called when production useCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUseCatalogStatement(ctx *UseCatalogStatementContext) {}

// EnterSetCatalogStatement is called when production setCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// ExitSetCatalogStatement is called when production setCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetCatalogStatement(ctx *SetCatalogStatementContext) {}

// EnterShowDatabasesStatement is called when production showDatabasesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {
}

// ExitShowDatabasesStatement is called when production showDatabasesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowDatabasesStatement(ctx *ShowDatabasesStatementContext) {}

// EnterAlterDbQuotaStatement is called when production alterDbQuotaStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// ExitAlterDbQuotaStatement is called when production alterDbQuotaStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterDbQuotaStatement(ctx *AlterDbQuotaStatementContext) {}

// EnterCreateDbStatement is called when production createDbStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateDbStatement(ctx *CreateDbStatementContext) {}

// ExitCreateDbStatement is called when production createDbStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateDbStatement(ctx *CreateDbStatementContext) {}

// EnterDropDbStatement is called when production dropDbStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropDbStatement(ctx *DropDbStatementContext) {}

// ExitDropDbStatement is called when production dropDbStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropDbStatement(ctx *DropDbStatementContext) {}

// EnterShowCreateDbStatement is called when production showCreateDbStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// ExitShowCreateDbStatement is called when production showCreateDbStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateDbStatement(ctx *ShowCreateDbStatementContext) {}

// EnterAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// ExitAlterDatabaseRenameStatement is called when production alterDatabaseRenameStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterDatabaseRenameStatement(ctx *AlterDatabaseRenameStatementContext) {
}

// EnterRecoverDbStmt is called when production recoverDbStmt is entered.
func (s *BaseDorisSQLParserListener) EnterRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// ExitRecoverDbStmt is called when production recoverDbStmt is exited.
func (s *BaseDorisSQLParserListener) ExitRecoverDbStmt(ctx *RecoverDbStmtContext) {}

// EnterShowDataStmt is called when production showDataStmt is entered.
func (s *BaseDorisSQLParserListener) EnterShowDataStmt(ctx *ShowDataStmtContext) {}

// ExitShowDataStmt is called when production showDataStmt is exited.
func (s *BaseDorisSQLParserListener) ExitShowDataStmt(ctx *ShowDataStmtContext) {}

// EnterShowDataDistributionStmt is called when production showDataDistributionStmt is entered.
func (s *BaseDorisSQLParserListener) EnterShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// ExitShowDataDistributionStmt is called when production showDataDistributionStmt is exited.
func (s *BaseDorisSQLParserListener) ExitShowDataDistributionStmt(ctx *ShowDataDistributionStmtContext) {
}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterColumnDesc is called when production columnDesc is entered.
func (s *BaseDorisSQLParserListener) EnterColumnDesc(ctx *ColumnDescContext) {}

// ExitColumnDesc is called when production columnDesc is exited.
func (s *BaseDorisSQLParserListener) ExitColumnDesc(ctx *ColumnDescContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseDorisSQLParserListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseDorisSQLParserListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterDefaultDesc is called when production defaultDesc is entered.
func (s *BaseDorisSQLParserListener) EnterDefaultDesc(ctx *DefaultDescContext) {}

// ExitDefaultDesc is called when production defaultDesc is exited.
func (s *BaseDorisSQLParserListener) ExitDefaultDesc(ctx *DefaultDescContext) {}

// EnterGeneratedColumnDesc is called when production generatedColumnDesc is entered.
func (s *BaseDorisSQLParserListener) EnterGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// ExitGeneratedColumnDesc is called when production generatedColumnDesc is exited.
func (s *BaseDorisSQLParserListener) ExitGeneratedColumnDesc(ctx *GeneratedColumnDescContext) {}

// EnterIndexDesc is called when production indexDesc is entered.
func (s *BaseDorisSQLParserListener) EnterIndexDesc(ctx *IndexDescContext) {}

// ExitIndexDesc is called when production indexDesc is exited.
func (s *BaseDorisSQLParserListener) ExitIndexDesc(ctx *IndexDescContext) {}

// EnterEngineDesc is called when production engineDesc is entered.
func (s *BaseDorisSQLParserListener) EnterEngineDesc(ctx *EngineDescContext) {}

// ExitEngineDesc is called when production engineDesc is exited.
func (s *BaseDorisSQLParserListener) ExitEngineDesc(ctx *EngineDescContext) {}

// EnterCharsetDesc is called when production charsetDesc is entered.
func (s *BaseDorisSQLParserListener) EnterCharsetDesc(ctx *CharsetDescContext) {}

// ExitCharsetDesc is called when production charsetDesc is exited.
func (s *BaseDorisSQLParserListener) ExitCharsetDesc(ctx *CharsetDescContext) {}

// EnterCollateDesc is called when production collateDesc is entered.
func (s *BaseDorisSQLParserListener) EnterCollateDesc(ctx *CollateDescContext) {}

// ExitCollateDesc is called when production collateDesc is exited.
func (s *BaseDorisSQLParserListener) ExitCollateDesc(ctx *CollateDescContext) {}

// EnterKeyDesc is called when production keyDesc is entered.
func (s *BaseDorisSQLParserListener) EnterKeyDesc(ctx *KeyDescContext) {}

// ExitKeyDesc is called when production keyDesc is exited.
func (s *BaseDorisSQLParserListener) ExitKeyDesc(ctx *KeyDescContext) {}

// EnterOrderByDesc is called when production orderByDesc is entered.
func (s *BaseDorisSQLParserListener) EnterOrderByDesc(ctx *OrderByDescContext) {}

// ExitOrderByDesc is called when production orderByDesc is exited.
func (s *BaseDorisSQLParserListener) ExitOrderByDesc(ctx *OrderByDescContext) {}

// EnterColumnNullable is called when production columnNullable is entered.
func (s *BaseDorisSQLParserListener) EnterColumnNullable(ctx *ColumnNullableContext) {}

// ExitColumnNullable is called when production columnNullable is exited.
func (s *BaseDorisSQLParserListener) ExitColumnNullable(ctx *ColumnNullableContext) {}

// EnterTypeWithNullable is called when production typeWithNullable is entered.
func (s *BaseDorisSQLParserListener) EnterTypeWithNullable(ctx *TypeWithNullableContext) {}

// ExitTypeWithNullable is called when production typeWithNullable is exited.
func (s *BaseDorisSQLParserListener) ExitTypeWithNullable(ctx *TypeWithNullableContext) {}

// EnterAggStateDesc is called when production aggStateDesc is entered.
func (s *BaseDorisSQLParserListener) EnterAggStateDesc(ctx *AggStateDescContext) {}

// ExitAggStateDesc is called when production aggStateDesc is exited.
func (s *BaseDorisSQLParserListener) ExitAggStateDesc(ctx *AggStateDescContext) {}

// EnterAggDesc is called when production aggDesc is entered.
func (s *BaseDorisSQLParserListener) EnterAggDesc(ctx *AggDescContext) {}

// ExitAggDesc is called when production aggDesc is exited.
func (s *BaseDorisSQLParserListener) ExitAggDesc(ctx *AggDescContext) {}

// EnterRollupDesc is called when production rollupDesc is entered.
func (s *BaseDorisSQLParserListener) EnterRollupDesc(ctx *RollupDescContext) {}

// ExitRollupDesc is called when production rollupDesc is exited.
func (s *BaseDorisSQLParserListener) ExitRollupDesc(ctx *RollupDescContext) {}

// EnterRollupItem is called when production rollupItem is entered.
func (s *BaseDorisSQLParserListener) EnterRollupItem(ctx *RollupItemContext) {}

// ExitRollupItem is called when production rollupItem is exited.
func (s *BaseDorisSQLParserListener) ExitRollupItem(ctx *RollupItemContext) {}

// EnterDupKeys is called when production dupKeys is entered.
func (s *BaseDorisSQLParserListener) EnterDupKeys(ctx *DupKeysContext) {}

// ExitDupKeys is called when production dupKeys is exited.
func (s *BaseDorisSQLParserListener) ExitDupKeys(ctx *DupKeysContext) {}

// EnterFromRollup is called when production fromRollup is entered.
func (s *BaseDorisSQLParserListener) EnterFromRollup(ctx *FromRollupContext) {}

// ExitFromRollup is called when production fromRollup is exited.
func (s *BaseDorisSQLParserListener) ExitFromRollup(ctx *FromRollupContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseDorisSQLParserListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseDorisSQLParserListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseDorisSQLParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseDorisSQLParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterCreateTableAsSelectStatement is called when production createTableAsSelectStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// ExitCreateTableAsSelectStatement is called when production createTableAsSelectStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateTableAsSelectStatement(ctx *CreateTableAsSelectStatementContext) {
}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// ExitCleanTemporaryTableStatement is called when production cleanTemporaryTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCleanTemporaryTableStatement(ctx *CleanTemporaryTableStatementContext) {
}

// EnterAlterTableStatement is called when production alterTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterTableStatement(ctx *AlterTableStatementContext) {}

// ExitAlterTableStatement is called when production alterTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterTableStatement(ctx *AlterTableStatementContext) {}

// EnterCreateIndexStatement is called when production createIndexStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// ExitCreateIndexStatement is called when production createIndexStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateIndexStatement(ctx *CreateIndexStatementContext) {}

// EnterDropIndexStatement is called when production dropIndexStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropIndexStatement(ctx *DropIndexStatementContext) {}

// ExitDropIndexStatement is called when production dropIndexStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropIndexStatement(ctx *DropIndexStatementContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseDorisSQLParserListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseDorisSQLParserListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterShowTableStatement is called when production showTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTableStatement(ctx *ShowTableStatementContext) {}

// ExitShowTableStatement is called when production showTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTableStatement(ctx *ShowTableStatementContext) {}

// EnterShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// ExitShowTemporaryTablesStatement is called when production showTemporaryTablesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTemporaryTablesStatement(ctx *ShowTemporaryTablesStatementContext) {
}

// EnterShowCreateTableStatement is called when production showCreateTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// ExitShowCreateTableStatement is called when production showCreateTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateTableStatement(ctx *ShowCreateTableStatementContext) {
}

// EnterShowColumnStatement is called when production showColumnStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowColumnStatement(ctx *ShowColumnStatementContext) {}

// ExitShowColumnStatement is called when production showColumnStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowColumnStatement(ctx *ShowColumnStatementContext) {}

// EnterShowTableStatusStatement is called when production showTableStatusStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// ExitShowTableStatusStatement is called when production showTableStatusStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTableStatusStatement(ctx *ShowTableStatusStatementContext) {
}

// EnterRefreshTableStatement is called when production refreshTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// ExitRefreshTableStatement is called when production refreshTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRefreshTableStatement(ctx *RefreshTableStatementContext) {}

// EnterShowAlterStatement is called when production showAlterStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowAlterStatement(ctx *ShowAlterStatementContext) {}

// ExitShowAlterStatement is called when production showAlterStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowAlterStatement(ctx *ShowAlterStatementContext) {}

// EnterDescTableStatement is called when production descTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDescTableStatement(ctx *DescTableStatementContext) {}

// ExitDescTableStatement is called when production descTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDescTableStatement(ctx *DescTableStatementContext) {}

// EnterCreateTableLikeStatement is called when production createTableLikeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// ExitCreateTableLikeStatement is called when production createTableLikeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateTableLikeStatement(ctx *CreateTableLikeStatementContext) {
}

// EnterShowIndexStatement is called when production showIndexStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowIndexStatement(ctx *ShowIndexStatementContext) {}

// ExitShowIndexStatement is called when production showIndexStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowIndexStatement(ctx *ShowIndexStatementContext) {}

// EnterRecoverTableStatement is called when production recoverTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// ExitRecoverTableStatement is called when production recoverTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRecoverTableStatement(ctx *RecoverTableStatementContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {
}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// EnterCancelAlterTableStatement is called when production cancelAlterTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// ExitCancelAlterTableStatement is called when production cancelAlterTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelAlterTableStatement(ctx *CancelAlterTableStatementContext) {
}

// EnterShowPartitionsStatement is called when production showPartitionsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {
}

// ExitShowPartitionsStatement is called when production showPartitionsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowPartitionsStatement(ctx *ShowPartitionsStatementContext) {
}

// EnterRecoverPartitionStatement is called when production recoverPartitionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// ExitRecoverPartitionStatement is called when production recoverPartitionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRecoverPartitionStatement(ctx *RecoverPartitionStatementContext) {
}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterAlterViewStatement is called when production alterViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterViewStatement(ctx *AlterViewStatementContext) {}

// ExitAlterViewStatement is called when production alterViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterViewStatement(ctx *AlterViewStatementContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterColumnNameWithComment is called when production columnNameWithComment is entered.
func (s *BaseDorisSQLParserListener) EnterColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// ExitColumnNameWithComment is called when production columnNameWithComment is exited.
func (s *BaseDorisSQLParserListener) ExitColumnNameWithComment(ctx *ColumnNameWithCommentContext) {}

// EnterSubmitTaskStatement is called when production submitTaskStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// ExitSubmitTaskStatement is called when production submitTaskStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSubmitTaskStatement(ctx *SubmitTaskStatementContext) {}

// EnterTaskClause is called when production taskClause is entered.
func (s *BaseDorisSQLParserListener) EnterTaskClause(ctx *TaskClauseContext) {}

// ExitTaskClause is called when production taskClause is exited.
func (s *BaseDorisSQLParserListener) ExitTaskClause(ctx *TaskClauseContext) {}

// EnterDropTaskStatement is called when production dropTaskStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropTaskStatement(ctx *DropTaskStatementContext) {}

// ExitDropTaskStatement is called when production dropTaskStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropTaskStatement(ctx *DropTaskStatementContext) {}

// EnterTaskScheduleDesc is called when production taskScheduleDesc is entered.
func (s *BaseDorisSQLParserListener) EnterTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// ExitTaskScheduleDesc is called when production taskScheduleDesc is exited.
func (s *BaseDorisSQLParserListener) ExitTaskScheduleDesc(ctx *TaskScheduleDescContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterMvPartitionExprs is called when production mvPartitionExprs is entered.
func (s *BaseDorisSQLParserListener) EnterMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// ExitMvPartitionExprs is called when production mvPartitionExprs is exited.
func (s *BaseDorisSQLParserListener) ExitMvPartitionExprs(ctx *MvPartitionExprsContext) {}

// EnterMaterializedViewDesc is called when production materializedViewDesc is entered.
func (s *BaseDorisSQLParserListener) EnterMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// ExitMaterializedViewDesc is called when production materializedViewDesc is exited.
func (s *BaseDorisSQLParserListener) ExitMaterializedViewDesc(ctx *MaterializedViewDescContext) {}

// EnterShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// ExitShowMaterializedViewsStatement is called when production showMaterializedViewsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowMaterializedViewsStatement(ctx *ShowMaterializedViewsStatementContext) {
}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// ExitAlterMaterializedViewStatement is called when production alterMaterializedViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterMaterializedViewStatement(ctx *AlterMaterializedViewStatementContext) {
}

// EnterRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// ExitRefreshMaterializedViewStatement is called when production refreshMaterializedViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRefreshMaterializedViewStatement(ctx *RefreshMaterializedViewStatementContext) {
}

// EnterCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// ExitCancelRefreshMaterializedViewStatement is called when production cancelRefreshMaterializedViewStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelRefreshMaterializedViewStatement(ctx *CancelRefreshMaterializedViewStatementContext) {
}

// EnterAdminSetConfigStatement is called when production adminSetConfigStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {
}

// ExitAdminSetConfigStatement is called when production adminSetConfigStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminSetConfigStatement(ctx *AdminSetConfigStatementContext) {
}

// EnterAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// ExitAdminSetReplicaStatusStatement is called when production adminSetReplicaStatusStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminSetReplicaStatusStatement(ctx *AdminSetReplicaStatusStatementContext) {
}

// EnterAdminShowConfigStatement is called when production adminShowConfigStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// ExitAdminShowConfigStatement is called when production adminShowConfigStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminShowConfigStatement(ctx *AdminShowConfigStatementContext) {
}

// EnterAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// ExitAdminShowReplicaDistributionStatement is called when production adminShowReplicaDistributionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminShowReplicaDistributionStatement(ctx *AdminShowReplicaDistributionStatementContext) {
}

// EnterAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// ExitAdminShowReplicaStatusStatement is called when production adminShowReplicaStatusStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminShowReplicaStatusStatement(ctx *AdminShowReplicaStatusStatementContext) {
}

// EnterAdminRepairTableStatement is called when production adminRepairTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// ExitAdminRepairTableStatement is called when production adminRepairTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminRepairTableStatement(ctx *AdminRepairTableStatementContext) {
}

// EnterAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// ExitAdminCancelRepairTableStatement is called when production adminCancelRepairTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminCancelRepairTableStatement(ctx *AdminCancelRepairTableStatementContext) {
}

// EnterAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// ExitAdminCheckTabletsStatement is called when production adminCheckTabletsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminCheckTabletsStatement(ctx *AdminCheckTabletsStatementContext) {
}

// EnterAdminSetPartitionVersion is called when production adminSetPartitionVersion is entered.
func (s *BaseDorisSQLParserListener) EnterAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// ExitAdminSetPartitionVersion is called when production adminSetPartitionVersion is exited.
func (s *BaseDorisSQLParserListener) ExitAdminSetPartitionVersion(ctx *AdminSetPartitionVersionContext) {
}

// EnterKillStatement is called when production killStatement is entered.
func (s *BaseDorisSQLParserListener) EnterKillStatement(ctx *KillStatementContext) {}

// ExitKillStatement is called when production killStatement is exited.
func (s *BaseDorisSQLParserListener) ExitKillStatement(ctx *KillStatementContext) {}

// EnterSyncStatement is called when production syncStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSyncStatement(ctx *SyncStatementContext) {}

// ExitSyncStatement is called when production syncStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSyncStatement(ctx *SyncStatementContext) {}

// EnterAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOnStatement is called when production adminSetAutomatedSnapshotOnStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminSetAutomatedSnapshotOnStatement(ctx *AdminSetAutomatedSnapshotOnStatementContext) {
}

// EnterAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// ExitAdminSetAutomatedSnapshotOffStatement is called when production adminSetAutomatedSnapshotOffStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAdminSetAutomatedSnapshotOffStatement(ctx *AdminSetAutomatedSnapshotOffStatementContext) {
}

// EnterAlterSystemStatement is called when production alterSystemStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// ExitAlterSystemStatement is called when production alterSystemStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterSystemStatement(ctx *AlterSystemStatementContext) {}

// EnterCancelAlterSystemStatement is called when production cancelAlterSystemStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// ExitCancelAlterSystemStatement is called when production cancelAlterSystemStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelAlterSystemStatement(ctx *CancelAlterSystemStatementContext) {
}

// EnterShowComputeNodesStatement is called when production showComputeNodesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// ExitShowComputeNodesStatement is called when production showComputeNodesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowComputeNodesStatement(ctx *ShowComputeNodesStatementContext) {
}

// EnterCreateExternalCatalogStatement is called when production createExternalCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// ExitCreateExternalCatalogStatement is called when production createExternalCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateExternalCatalogStatement(ctx *CreateExternalCatalogStatementContext) {
}

// EnterShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// ExitShowCreateExternalCatalogStatement is called when production showCreateExternalCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateExternalCatalogStatement(ctx *ShowCreateExternalCatalogStatementContext) {
}

// EnterDropExternalCatalogStatement is called when production dropExternalCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// ExitDropExternalCatalogStatement is called when production dropExternalCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropExternalCatalogStatement(ctx *DropExternalCatalogStatementContext) {
}

// EnterShowCatalogsStatement is called when production showCatalogsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// ExitShowCatalogsStatement is called when production showCatalogsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCatalogsStatement(ctx *ShowCatalogsStatementContext) {}

// EnterAlterCatalogStatement is called when production alterCatalogStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// ExitAlterCatalogStatement is called when production alterCatalogStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterCatalogStatement(ctx *AlterCatalogStatementContext) {}

// EnterCreateStorageVolumeStatement is called when production createStorageVolumeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// ExitCreateStorageVolumeStatement is called when production createStorageVolumeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateStorageVolumeStatement(ctx *CreateStorageVolumeStatementContext) {
}

// EnterTypeDesc is called when production typeDesc is entered.
func (s *BaseDorisSQLParserListener) EnterTypeDesc(ctx *TypeDescContext) {}

// ExitTypeDesc is called when production typeDesc is exited.
func (s *BaseDorisSQLParserListener) ExitTypeDesc(ctx *TypeDescContext) {}

// EnterLocationsDesc is called when production locationsDesc is entered.
func (s *BaseDorisSQLParserListener) EnterLocationsDesc(ctx *LocationsDescContext) {}

// ExitLocationsDesc is called when production locationsDesc is exited.
func (s *BaseDorisSQLParserListener) ExitLocationsDesc(ctx *LocationsDescContext) {}

// EnterShowStorageVolumesStatement is called when production showStorageVolumesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// ExitShowStorageVolumesStatement is called when production showStorageVolumesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowStorageVolumesStatement(ctx *ShowStorageVolumesStatementContext) {
}

// EnterDropStorageVolumeStatement is called when production dropStorageVolumeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// ExitDropStorageVolumeStatement is called when production dropStorageVolumeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropStorageVolumeStatement(ctx *DropStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// ExitAlterStorageVolumeStatement is called when production alterStorageVolumeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterStorageVolumeStatement(ctx *AlterStorageVolumeStatementContext) {
}

// EnterAlterStorageVolumeClause is called when production alterStorageVolumeClause is entered.
func (s *BaseDorisSQLParserListener) EnterAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// ExitAlterStorageVolumeClause is called when production alterStorageVolumeClause is exited.
func (s *BaseDorisSQLParserListener) ExitAlterStorageVolumeClause(ctx *AlterStorageVolumeClauseContext) {
}

// EnterModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// ExitModifyStorageVolumePropertiesClause is called when production modifyStorageVolumePropertiesClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyStorageVolumePropertiesClause(ctx *ModifyStorageVolumePropertiesClauseContext) {
}

// EnterModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// ExitModifyStorageVolumeCommentClause is called when production modifyStorageVolumeCommentClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyStorageVolumeCommentClause(ctx *ModifyStorageVolumeCommentClauseContext) {
}

// EnterDescStorageVolumeStatement is called when production descStorageVolumeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// ExitDescStorageVolumeStatement is called when production descStorageVolumeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDescStorageVolumeStatement(ctx *DescStorageVolumeStatementContext) {
}

// EnterSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// ExitSetDefaultStorageVolumeStatement is called when production setDefaultStorageVolumeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetDefaultStorageVolumeStatement(ctx *SetDefaultStorageVolumeStatementContext) {
}

// EnterUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// ExitUpdateFailPointStatusStatement is called when production updateFailPointStatusStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUpdateFailPointStatusStatement(ctx *UpdateFailPointStatusStatementContext) {
}

// EnterShowFailPointStatement is called when production showFailPointStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowFailPointStatement(ctx *ShowFailPointStatementContext) {
}

// ExitShowFailPointStatement is called when production showFailPointStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowFailPointStatement(ctx *ShowFailPointStatementContext) {}

// EnterCreateDictionaryStatement is called when production createDictionaryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// ExitCreateDictionaryStatement is called when production createDictionaryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateDictionaryStatement(ctx *CreateDictionaryStatementContext) {
}

// EnterDropDictionaryStatement is called when production dropDictionaryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropDictionaryStatement(ctx *DropDictionaryStatementContext) {
}

// ExitDropDictionaryStatement is called when production dropDictionaryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropDictionaryStatement(ctx *DropDictionaryStatementContext) {
}

// EnterRefreshDictionaryStatement is called when production refreshDictionaryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// ExitRefreshDictionaryStatement is called when production refreshDictionaryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRefreshDictionaryStatement(ctx *RefreshDictionaryStatementContext) {
}

// EnterShowDictionaryStatement is called when production showDictionaryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {
}

// ExitShowDictionaryStatement is called when production showDictionaryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowDictionaryStatement(ctx *ShowDictionaryStatementContext) {
}

// EnterCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// ExitCancelRefreshDictionaryStatement is called when production cancelRefreshDictionaryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelRefreshDictionaryStatement(ctx *CancelRefreshDictionaryStatementContext) {
}

// EnterDictionaryColumnDesc is called when production dictionaryColumnDesc is entered.
func (s *BaseDorisSQLParserListener) EnterDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// ExitDictionaryColumnDesc is called when production dictionaryColumnDesc is exited.
func (s *BaseDorisSQLParserListener) ExitDictionaryColumnDesc(ctx *DictionaryColumnDescContext) {}

// EnterDictionaryName is called when production dictionaryName is entered.
func (s *BaseDorisSQLParserListener) EnterDictionaryName(ctx *DictionaryNameContext) {}

// ExitDictionaryName is called when production dictionaryName is exited.
func (s *BaseDorisSQLParserListener) ExitDictionaryName(ctx *DictionaryNameContext) {}

// EnterAlterClause is called when production alterClause is entered.
func (s *BaseDorisSQLParserListener) EnterAlterClause(ctx *AlterClauseContext) {}

// ExitAlterClause is called when production alterClause is exited.
func (s *BaseDorisSQLParserListener) ExitAlterClause(ctx *AlterClauseContext) {}

// EnterAddFrontendClause is called when production addFrontendClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddFrontendClause(ctx *AddFrontendClauseContext) {}

// ExitAddFrontendClause is called when production addFrontendClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddFrontendClause(ctx *AddFrontendClauseContext) {}

// EnterDropFrontendClause is called when production dropFrontendClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropFrontendClause(ctx *DropFrontendClauseContext) {}

// ExitDropFrontendClause is called when production dropFrontendClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropFrontendClause(ctx *DropFrontendClauseContext) {}

// EnterModifyFrontendHostClause is called when production modifyFrontendHostClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// ExitModifyFrontendHostClause is called when production modifyFrontendHostClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyFrontendHostClause(ctx *ModifyFrontendHostClauseContext) {
}

// EnterAddBackendClause is called when production addBackendClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddBackendClause(ctx *AddBackendClauseContext) {}

// ExitAddBackendClause is called when production addBackendClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddBackendClause(ctx *AddBackendClauseContext) {}

// EnterDropBackendClause is called when production dropBackendClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropBackendClause(ctx *DropBackendClauseContext) {}

// ExitDropBackendClause is called when production dropBackendClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropBackendClause(ctx *DropBackendClauseContext) {}

// EnterDecommissionBackendClause is called when production decommissionBackendClause is entered.
func (s *BaseDorisSQLParserListener) EnterDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// ExitDecommissionBackendClause is called when production decommissionBackendClause is exited.
func (s *BaseDorisSQLParserListener) ExitDecommissionBackendClause(ctx *DecommissionBackendClauseContext) {
}

// EnterModifyBackendClause is called when production modifyBackendClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// ExitModifyBackendClause is called when production modifyBackendClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyBackendClause(ctx *ModifyBackendClauseContext) {}

// EnterAddComputeNodeClause is called when production addComputeNodeClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// ExitAddComputeNodeClause is called when production addComputeNodeClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddComputeNodeClause(ctx *AddComputeNodeClauseContext) {}

// EnterDropComputeNodeClause is called when production dropComputeNodeClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// ExitDropComputeNodeClause is called when production dropComputeNodeClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropComputeNodeClause(ctx *DropComputeNodeClauseContext) {}

// EnterModifyBrokerClause is called when production modifyBrokerClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// ExitModifyBrokerClause is called when production modifyBrokerClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyBrokerClause(ctx *ModifyBrokerClauseContext) {}

// EnterAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is entered.
func (s *BaseDorisSQLParserListener) EnterAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {
}

// ExitAlterLoadErrorUrlClause is called when production alterLoadErrorUrlClause is exited.
func (s *BaseDorisSQLParserListener) ExitAlterLoadErrorUrlClause(ctx *AlterLoadErrorUrlClauseContext) {
}

// EnterCreateImageClause is called when production createImageClause is entered.
func (s *BaseDorisSQLParserListener) EnterCreateImageClause(ctx *CreateImageClauseContext) {}

// ExitCreateImageClause is called when production createImageClause is exited.
func (s *BaseDorisSQLParserListener) ExitCreateImageClause(ctx *CreateImageClauseContext) {}

// EnterCleanTabletSchedQClause is called when production cleanTabletSchedQClause is entered.
func (s *BaseDorisSQLParserListener) EnterCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {
}

// ExitCleanTabletSchedQClause is called when production cleanTabletSchedQClause is exited.
func (s *BaseDorisSQLParserListener) ExitCleanTabletSchedQClause(ctx *CleanTabletSchedQClauseContext) {
}

// EnterDecommissionDiskClause is called when production decommissionDiskClause is entered.
func (s *BaseDorisSQLParserListener) EnterDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {
}

// ExitDecommissionDiskClause is called when production decommissionDiskClause is exited.
func (s *BaseDorisSQLParserListener) ExitDecommissionDiskClause(ctx *DecommissionDiskClauseContext) {}

// EnterCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is entered.
func (s *BaseDorisSQLParserListener) EnterCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// ExitCancelDecommissionDiskClause is called when production cancelDecommissionDiskClause is exited.
func (s *BaseDorisSQLParserListener) ExitCancelDecommissionDiskClause(ctx *CancelDecommissionDiskClauseContext) {
}

// EnterDisableDiskClause is called when production disableDiskClause is entered.
func (s *BaseDorisSQLParserListener) EnterDisableDiskClause(ctx *DisableDiskClauseContext) {}

// ExitDisableDiskClause is called when production disableDiskClause is exited.
func (s *BaseDorisSQLParserListener) ExitDisableDiskClause(ctx *DisableDiskClauseContext) {}

// EnterCancelDisableDiskClause is called when production cancelDisableDiskClause is entered.
func (s *BaseDorisSQLParserListener) EnterCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {
}

// ExitCancelDisableDiskClause is called when production cancelDisableDiskClause is exited.
func (s *BaseDorisSQLParserListener) ExitCancelDisableDiskClause(ctx *CancelDisableDiskClauseContext) {
}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseDorisSQLParserListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseDorisSQLParserListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterDropIndexClause is called when production dropIndexClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropIndexClause(ctx *DropIndexClauseContext) {}

// ExitDropIndexClause is called when production dropIndexClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropIndexClause(ctx *DropIndexClauseContext) {}

// EnterTableRenameClause is called when production tableRenameClause is entered.
func (s *BaseDorisSQLParserListener) EnterTableRenameClause(ctx *TableRenameClauseContext) {}

// ExitTableRenameClause is called when production tableRenameClause is exited.
func (s *BaseDorisSQLParserListener) ExitTableRenameClause(ctx *TableRenameClauseContext) {}

// EnterSwapTableClause is called when production swapTableClause is entered.
func (s *BaseDorisSQLParserListener) EnterSwapTableClause(ctx *SwapTableClauseContext) {}

// ExitSwapTableClause is called when production swapTableClause is exited.
func (s *BaseDorisSQLParserListener) ExitSwapTableClause(ctx *SwapTableClauseContext) {}

// EnterModifyPropertiesClause is called when production modifyPropertiesClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {
}

// ExitModifyPropertiesClause is called when production modifyPropertiesClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyPropertiesClause(ctx *ModifyPropertiesClauseContext) {}

// EnterModifyCommentClause is called when production modifyCommentClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// ExitModifyCommentClause is called when production modifyCommentClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyCommentClause(ctx *ModifyCommentClauseContext) {}

// EnterOptimizeRange is called when production optimizeRange is entered.
func (s *BaseDorisSQLParserListener) EnterOptimizeRange(ctx *OptimizeRangeContext) {}

// ExitOptimizeRange is called when production optimizeRange is exited.
func (s *BaseDorisSQLParserListener) ExitOptimizeRange(ctx *OptimizeRangeContext) {}

// EnterOptimizeClause is called when production optimizeClause is entered.
func (s *BaseDorisSQLParserListener) EnterOptimizeClause(ctx *OptimizeClauseContext) {}

// ExitOptimizeClause is called when production optimizeClause is exited.
func (s *BaseDorisSQLParserListener) ExitOptimizeClause(ctx *OptimizeClauseContext) {}

// EnterAddColumnClause is called when production addColumnClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddColumnClause(ctx *AddColumnClauseContext) {}

// ExitAddColumnClause is called when production addColumnClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddColumnClause(ctx *AddColumnClauseContext) {}

// EnterAddColumnsClause is called when production addColumnsClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddColumnsClause(ctx *AddColumnsClauseContext) {}

// ExitAddColumnsClause is called when production addColumnsClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddColumnsClause(ctx *AddColumnsClauseContext) {}

// EnterDropColumnClause is called when production dropColumnClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropColumnClause(ctx *DropColumnClauseContext) {}

// ExitDropColumnClause is called when production dropColumnClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropColumnClause(ctx *DropColumnClauseContext) {}

// EnterModifyColumnClause is called when production modifyColumnClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// ExitModifyColumnClause is called when production modifyColumnClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyColumnClause(ctx *ModifyColumnClauseContext) {}

// EnterModifyColumnCommentClause is called when production modifyColumnCommentClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// ExitModifyColumnCommentClause is called when production modifyColumnCommentClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyColumnCommentClause(ctx *ModifyColumnCommentClauseContext) {
}

// EnterColumnRenameClause is called when production columnRenameClause is entered.
func (s *BaseDorisSQLParserListener) EnterColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// ExitColumnRenameClause is called when production columnRenameClause is exited.
func (s *BaseDorisSQLParserListener) ExitColumnRenameClause(ctx *ColumnRenameClauseContext) {}

// EnterReorderColumnsClause is called when production reorderColumnsClause is entered.
func (s *BaseDorisSQLParserListener) EnterReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// ExitReorderColumnsClause is called when production reorderColumnsClause is exited.
func (s *BaseDorisSQLParserListener) ExitReorderColumnsClause(ctx *ReorderColumnsClauseContext) {}

// EnterRollupRenameClause is called when production rollupRenameClause is entered.
func (s *BaseDorisSQLParserListener) EnterRollupRenameClause(ctx *RollupRenameClauseContext) {}

// ExitRollupRenameClause is called when production rollupRenameClause is exited.
func (s *BaseDorisSQLParserListener) ExitRollupRenameClause(ctx *RollupRenameClauseContext) {}

// EnterCompactionClause is called when production compactionClause is entered.
func (s *BaseDorisSQLParserListener) EnterCompactionClause(ctx *CompactionClauseContext) {}

// ExitCompactionClause is called when production compactionClause is exited.
func (s *BaseDorisSQLParserListener) ExitCompactionClause(ctx *CompactionClauseContext) {}

// EnterSubfieldName is called when production subfieldName is entered.
func (s *BaseDorisSQLParserListener) EnterSubfieldName(ctx *SubfieldNameContext) {}

// ExitSubfieldName is called when production subfieldName is exited.
func (s *BaseDorisSQLParserListener) ExitSubfieldName(ctx *SubfieldNameContext) {}

// EnterNestedFieldName is called when production nestedFieldName is entered.
func (s *BaseDorisSQLParserListener) EnterNestedFieldName(ctx *NestedFieldNameContext) {}

// ExitNestedFieldName is called when production nestedFieldName is exited.
func (s *BaseDorisSQLParserListener) ExitNestedFieldName(ctx *NestedFieldNameContext) {}

// EnterAddFieldClause is called when production addFieldClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddFieldClause(ctx *AddFieldClauseContext) {}

// ExitAddFieldClause is called when production addFieldClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddFieldClause(ctx *AddFieldClauseContext) {}

// EnterDropFieldClause is called when production dropFieldClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropFieldClause(ctx *DropFieldClauseContext) {}

// ExitDropFieldClause is called when production dropFieldClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropFieldClause(ctx *DropFieldClauseContext) {}

// EnterCreateOrReplaceTagClause is called when production createOrReplaceTagClause is entered.
func (s *BaseDorisSQLParserListener) EnterCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// ExitCreateOrReplaceTagClause is called when production createOrReplaceTagClause is exited.
func (s *BaseDorisSQLParserListener) ExitCreateOrReplaceTagClause(ctx *CreateOrReplaceTagClauseContext) {
}

// EnterCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is entered.
func (s *BaseDorisSQLParserListener) EnterCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// ExitCreateOrReplaceBranchClause is called when production createOrReplaceBranchClause is exited.
func (s *BaseDorisSQLParserListener) ExitCreateOrReplaceBranchClause(ctx *CreateOrReplaceBranchClauseContext) {
}

// EnterDropBranchClause is called when production dropBranchClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropBranchClause(ctx *DropBranchClauseContext) {}

// ExitDropBranchClause is called when production dropBranchClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropBranchClause(ctx *DropBranchClauseContext) {}

// EnterDropTagClause is called when production dropTagClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropTagClause(ctx *DropTagClauseContext) {}

// ExitDropTagClause is called when production dropTagClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropTagClause(ctx *DropTagClauseContext) {}

// EnterTableOperationClause is called when production tableOperationClause is entered.
func (s *BaseDorisSQLParserListener) EnterTableOperationClause(ctx *TableOperationClauseContext) {}

// ExitTableOperationClause is called when production tableOperationClause is exited.
func (s *BaseDorisSQLParserListener) ExitTableOperationClause(ctx *TableOperationClauseContext) {}

// EnterTagOptions is called when production tagOptions is entered.
func (s *BaseDorisSQLParserListener) EnterTagOptions(ctx *TagOptionsContext) {}

// ExitTagOptions is called when production tagOptions is exited.
func (s *BaseDorisSQLParserListener) ExitTagOptions(ctx *TagOptionsContext) {}

// EnterBranchOptions is called when production branchOptions is entered.
func (s *BaseDorisSQLParserListener) EnterBranchOptions(ctx *BranchOptionsContext) {}

// ExitBranchOptions is called when production branchOptions is exited.
func (s *BaseDorisSQLParserListener) ExitBranchOptions(ctx *BranchOptionsContext) {}

// EnterSnapshotRetention is called when production snapshotRetention is entered.
func (s *BaseDorisSQLParserListener) EnterSnapshotRetention(ctx *SnapshotRetentionContext) {}

// ExitSnapshotRetention is called when production snapshotRetention is exited.
func (s *BaseDorisSQLParserListener) ExitSnapshotRetention(ctx *SnapshotRetentionContext) {}

// EnterRefRetain is called when production refRetain is entered.
func (s *BaseDorisSQLParserListener) EnterRefRetain(ctx *RefRetainContext) {}

// ExitRefRetain is called when production refRetain is exited.
func (s *BaseDorisSQLParserListener) ExitRefRetain(ctx *RefRetainContext) {}

// EnterMaxSnapshotAge is called when production maxSnapshotAge is entered.
func (s *BaseDorisSQLParserListener) EnterMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// ExitMaxSnapshotAge is called when production maxSnapshotAge is exited.
func (s *BaseDorisSQLParserListener) ExitMaxSnapshotAge(ctx *MaxSnapshotAgeContext) {}

// EnterMinSnapshotsToKeep is called when production minSnapshotsToKeep is entered.
func (s *BaseDorisSQLParserListener) EnterMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// ExitMinSnapshotsToKeep is called when production minSnapshotsToKeep is exited.
func (s *BaseDorisSQLParserListener) ExitMinSnapshotsToKeep(ctx *MinSnapshotsToKeepContext) {}

// EnterSnapshotId is called when production snapshotId is entered.
func (s *BaseDorisSQLParserListener) EnterSnapshotId(ctx *SnapshotIdContext) {}

// ExitSnapshotId is called when production snapshotId is exited.
func (s *BaseDorisSQLParserListener) ExitSnapshotId(ctx *SnapshotIdContext) {}

// EnterTimeUnit is called when production timeUnit is entered.
func (s *BaseDorisSQLParserListener) EnterTimeUnit(ctx *TimeUnitContext) {}

// ExitTimeUnit is called when production timeUnit is exited.
func (s *BaseDorisSQLParserListener) ExitTimeUnit(ctx *TimeUnitContext) {}

// EnterInteger_list is called when production integer_list is entered.
func (s *BaseDorisSQLParserListener) EnterInteger_list(ctx *Integer_listContext) {}

// ExitInteger_list is called when production integer_list is exited.
func (s *BaseDorisSQLParserListener) ExitInteger_list(ctx *Integer_listContext) {}

// EnterDropPersistentIndexClause is called when production dropPersistentIndexClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// ExitDropPersistentIndexClause is called when production dropPersistentIndexClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropPersistentIndexClause(ctx *DropPersistentIndexClauseContext) {
}

// EnterAddPartitionClause is called when production addPartitionClause is entered.
func (s *BaseDorisSQLParserListener) EnterAddPartitionClause(ctx *AddPartitionClauseContext) {}

// ExitAddPartitionClause is called when production addPartitionClause is exited.
func (s *BaseDorisSQLParserListener) ExitAddPartitionClause(ctx *AddPartitionClauseContext) {}

// EnterDropPartitionClause is called when production dropPartitionClause is entered.
func (s *BaseDorisSQLParserListener) EnterDropPartitionClause(ctx *DropPartitionClauseContext) {}

// ExitDropPartitionClause is called when production dropPartitionClause is exited.
func (s *BaseDorisSQLParserListener) ExitDropPartitionClause(ctx *DropPartitionClauseContext) {}

// EnterTruncatePartitionClause is called when production truncatePartitionClause is entered.
func (s *BaseDorisSQLParserListener) EnterTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {
}

// ExitTruncatePartitionClause is called when production truncatePartitionClause is exited.
func (s *BaseDorisSQLParserListener) ExitTruncatePartitionClause(ctx *TruncatePartitionClauseContext) {
}

// EnterModifyPartitionClause is called when production modifyPartitionClause is entered.
func (s *BaseDorisSQLParserListener) EnterModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// ExitModifyPartitionClause is called when production modifyPartitionClause is exited.
func (s *BaseDorisSQLParserListener) ExitModifyPartitionClause(ctx *ModifyPartitionClauseContext) {}

// EnterReplacePartitionClause is called when production replacePartitionClause is entered.
func (s *BaseDorisSQLParserListener) EnterReplacePartitionClause(ctx *ReplacePartitionClauseContext) {
}

// ExitReplacePartitionClause is called when production replacePartitionClause is exited.
func (s *BaseDorisSQLParserListener) ExitReplacePartitionClause(ctx *ReplacePartitionClauseContext) {}

// EnterPartitionRenameClause is called when production partitionRenameClause is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// ExitPartitionRenameClause is called when production partitionRenameClause is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionRenameClause(ctx *PartitionRenameClauseContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseDorisSQLParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseDorisSQLParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is entered.
func (s *BaseDorisSQLParserListener) EnterInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// ExitInsertLabelOrColumnAliases is called when production insertLabelOrColumnAliases is exited.
func (s *BaseDorisSQLParserListener) ExitInsertLabelOrColumnAliases(ctx *InsertLabelOrColumnAliasesContext) {
}

// EnterColumnAliasesOrByName is called when production columnAliasesOrByName is entered.
func (s *BaseDorisSQLParserListener) EnterColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// ExitColumnAliasesOrByName is called when production columnAliasesOrByName is exited.
func (s *BaseDorisSQLParserListener) ExitColumnAliasesOrByName(ctx *ColumnAliasesOrByNameContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterCreateRoutineLoadStatement is called when production createRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// ExitCreateRoutineLoadStatement is called when production createRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateRoutineLoadStatement(ctx *CreateRoutineLoadStatementContext) {
}

// EnterAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// ExitAlterRoutineLoadStatement is called when production alterRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterRoutineLoadStatement(ctx *AlterRoutineLoadStatementContext) {
}

// EnterDataSource is called when production dataSource is entered.
func (s *BaseDorisSQLParserListener) EnterDataSource(ctx *DataSourceContext) {}

// ExitDataSource is called when production dataSource is exited.
func (s *BaseDorisSQLParserListener) ExitDataSource(ctx *DataSourceContext) {}

// EnterLoadProperties is called when production loadProperties is entered.
func (s *BaseDorisSQLParserListener) EnterLoadProperties(ctx *LoadPropertiesContext) {}

// ExitLoadProperties is called when production loadProperties is exited.
func (s *BaseDorisSQLParserListener) ExitLoadProperties(ctx *LoadPropertiesContext) {}

// EnterColSeparatorProperty is called when production colSeparatorProperty is entered.
func (s *BaseDorisSQLParserListener) EnterColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// ExitColSeparatorProperty is called when production colSeparatorProperty is exited.
func (s *BaseDorisSQLParserListener) ExitColSeparatorProperty(ctx *ColSeparatorPropertyContext) {}

// EnterRowDelimiterProperty is called when production rowDelimiterProperty is entered.
func (s *BaseDorisSQLParserListener) EnterRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// ExitRowDelimiterProperty is called when production rowDelimiterProperty is exited.
func (s *BaseDorisSQLParserListener) ExitRowDelimiterProperty(ctx *RowDelimiterPropertyContext) {}

// EnterImportColumns is called when production importColumns is entered.
func (s *BaseDorisSQLParserListener) EnterImportColumns(ctx *ImportColumnsContext) {}

// ExitImportColumns is called when production importColumns is exited.
func (s *BaseDorisSQLParserListener) ExitImportColumns(ctx *ImportColumnsContext) {}

// EnterColumnProperties is called when production columnProperties is entered.
func (s *BaseDorisSQLParserListener) EnterColumnProperties(ctx *ColumnPropertiesContext) {}

// ExitColumnProperties is called when production columnProperties is exited.
func (s *BaseDorisSQLParserListener) ExitColumnProperties(ctx *ColumnPropertiesContext) {}

// EnterJobProperties is called when production jobProperties is entered.
func (s *BaseDorisSQLParserListener) EnterJobProperties(ctx *JobPropertiesContext) {}

// ExitJobProperties is called when production jobProperties is exited.
func (s *BaseDorisSQLParserListener) ExitJobProperties(ctx *JobPropertiesContext) {}

// EnterDataSourceProperties is called when production dataSourceProperties is entered.
func (s *BaseDorisSQLParserListener) EnterDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// ExitDataSourceProperties is called when production dataSourceProperties is exited.
func (s *BaseDorisSQLParserListener) ExitDataSourceProperties(ctx *DataSourcePropertiesContext) {}

// EnterStopRoutineLoadStatement is called when production stopRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// ExitStopRoutineLoadStatement is called when production stopRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitStopRoutineLoadStatement(ctx *StopRoutineLoadStatementContext) {
}

// EnterResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// ExitResumeRoutineLoadStatement is called when production resumeRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitResumeRoutineLoadStatement(ctx *ResumeRoutineLoadStatementContext) {
}

// EnterPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// ExitPauseRoutineLoadStatement is called when production pauseRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitPauseRoutineLoadStatement(ctx *PauseRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadStatement is called when production showRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// ExitShowRoutineLoadStatement is called when production showRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRoutineLoadStatement(ctx *ShowRoutineLoadStatementContext) {
}

// EnterShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// ExitShowRoutineLoadTaskStatement is called when production showRoutineLoadTaskStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRoutineLoadTaskStatement(ctx *ShowRoutineLoadTaskStatementContext) {
}

// EnterShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// ExitShowCreateRoutineLoadStatement is called when production showCreateRoutineLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateRoutineLoadStatement(ctx *ShowCreateRoutineLoadStatementContext) {
}

// EnterShowStreamLoadStatement is called when production showStreamLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {
}

// ExitShowStreamLoadStatement is called when production showStreamLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowStreamLoadStatement(ctx *ShowStreamLoadStatementContext) {
}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterRegularColumns is called when production regularColumns is entered.
func (s *BaseDorisSQLParserListener) EnterRegularColumns(ctx *RegularColumnsContext) {}

// ExitRegularColumns is called when production regularColumns is exited.
func (s *BaseDorisSQLParserListener) ExitRegularColumns(ctx *RegularColumnsContext) {}

// EnterAllColumns is called when production allColumns is entered.
func (s *BaseDorisSQLParserListener) EnterAllColumns(ctx *AllColumnsContext) {}

// ExitAllColumns is called when production allColumns is exited.
func (s *BaseDorisSQLParserListener) ExitAllColumns(ctx *AllColumnsContext) {}

// EnterPredicateColumns is called when production predicateColumns is entered.
func (s *BaseDorisSQLParserListener) EnterPredicateColumns(ctx *PredicateColumnsContext) {}

// ExitPredicateColumns is called when production predicateColumns is exited.
func (s *BaseDorisSQLParserListener) ExitPredicateColumns(ctx *PredicateColumnsContext) {}

// EnterMultiColumnSet is called when production multiColumnSet is entered.
func (s *BaseDorisSQLParserListener) EnterMultiColumnSet(ctx *MultiColumnSetContext) {}

// ExitMultiColumnSet is called when production multiColumnSet is exited.
func (s *BaseDorisSQLParserListener) ExitMultiColumnSet(ctx *MultiColumnSetContext) {}

// EnterDropStatsStatement is called when production dropStatsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropStatsStatement(ctx *DropStatsStatementContext) {}

// ExitDropStatsStatement is called when production dropStatsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropStatsStatement(ctx *DropStatsStatementContext) {}

// EnterHistogramStatement is called when production histogramStatement is entered.
func (s *BaseDorisSQLParserListener) EnterHistogramStatement(ctx *HistogramStatementContext) {}

// ExitHistogramStatement is called when production histogramStatement is exited.
func (s *BaseDorisSQLParserListener) ExitHistogramStatement(ctx *HistogramStatementContext) {}

// EnterAnalyzeHistogramStatement is called when production analyzeHistogramStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// ExitAnalyzeHistogramStatement is called when production analyzeHistogramStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAnalyzeHistogramStatement(ctx *AnalyzeHistogramStatementContext) {
}

// EnterDropHistogramStatement is called when production dropHistogramStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropHistogramStatement(ctx *DropHistogramStatementContext) {
}

// ExitDropHistogramStatement is called when production dropHistogramStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropHistogramStatement(ctx *DropHistogramStatementContext) {}

// EnterCreateAnalyzeStatement is called when production createAnalyzeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {
}

// ExitCreateAnalyzeStatement is called when production createAnalyzeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateAnalyzeStatement(ctx *CreateAnalyzeStatementContext) {}

// EnterDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {
}

// ExitDropAnalyzeJobStatement is called when production dropAnalyzeJobStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropAnalyzeJobStatement(ctx *DropAnalyzeJobStatementContext) {
}

// EnterShowAnalyzeStatement is called when production showAnalyzeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// ExitShowAnalyzeStatement is called when production showAnalyzeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowAnalyzeStatement(ctx *ShowAnalyzeStatementContext) {}

// EnterShowStatsMetaStatement is called when production showStatsMetaStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {
}

// ExitShowStatsMetaStatement is called when production showStatsMetaStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowStatsMetaStatement(ctx *ShowStatsMetaStatementContext) {}

// EnterShowHistogramMetaStatement is called when production showHistogramMetaStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// ExitShowHistogramMetaStatement is called when production showHistogramMetaStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowHistogramMetaStatement(ctx *ShowHistogramMetaStatementContext) {
}

// EnterKillAnalyzeStatement is called when production killAnalyzeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// ExitKillAnalyzeStatement is called when production killAnalyzeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitKillAnalyzeStatement(ctx *KillAnalyzeStatementContext) {}

// EnterAnalyzeProfileStatement is called when production analyzeProfileStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {
}

// ExitAnalyzeProfileStatement is called when production analyzeProfileStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAnalyzeProfileStatement(ctx *AnalyzeProfileStatementContext) {
}

// EnterCreateBaselinePlanStatement is called when production createBaselinePlanStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// ExitCreateBaselinePlanStatement is called when production createBaselinePlanStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateBaselinePlanStatement(ctx *CreateBaselinePlanStatementContext) {
}

// EnterDropBaselinePlanStatement is called when production dropBaselinePlanStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// ExitDropBaselinePlanStatement is called when production dropBaselinePlanStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropBaselinePlanStatement(ctx *DropBaselinePlanStatementContext) {
}

// EnterShowBaselinePlanStatement is called when production showBaselinePlanStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// ExitShowBaselinePlanStatement is called when production showBaselinePlanStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowBaselinePlanStatement(ctx *ShowBaselinePlanStatementContext) {
}

// EnterCreateResourceGroupStatement is called when production createResourceGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// ExitCreateResourceGroupStatement is called when production createResourceGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateResourceGroupStatement(ctx *CreateResourceGroupStatementContext) {
}

// EnterDropResourceGroupStatement is called when production dropResourceGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// ExitDropResourceGroupStatement is called when production dropResourceGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropResourceGroupStatement(ctx *DropResourceGroupStatementContext) {
}

// EnterAlterResourceGroupStatement is called when production alterResourceGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// ExitAlterResourceGroupStatement is called when production alterResourceGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterResourceGroupStatement(ctx *AlterResourceGroupStatementContext) {
}

// EnterShowResourceGroupStatement is called when production showResourceGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// ExitShowResourceGroupStatement is called when production showResourceGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowResourceGroupStatement(ctx *ShowResourceGroupStatementContext) {
}

// EnterShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// ExitShowResourceGroupUsageStatement is called when production showResourceGroupUsageStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowResourceGroupUsageStatement(ctx *ShowResourceGroupUsageStatementContext) {
}

// EnterCreateResourceStatement is called when production createResourceStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateResourceStatement(ctx *CreateResourceStatementContext) {
}

// ExitCreateResourceStatement is called when production createResourceStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateResourceStatement(ctx *CreateResourceStatementContext) {
}

// EnterAlterResourceStatement is called when production alterResourceStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterResourceStatement(ctx *AlterResourceStatementContext) {
}

// ExitAlterResourceStatement is called when production alterResourceStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterResourceStatement(ctx *AlterResourceStatementContext) {}

// EnterDropResourceStatement is called when production dropResourceStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropResourceStatement(ctx *DropResourceStatementContext) {}

// ExitDropResourceStatement is called when production dropResourceStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropResourceStatement(ctx *DropResourceStatementContext) {}

// EnterShowResourceStatement is called when production showResourceStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowResourceStatement(ctx *ShowResourceStatementContext) {}

// ExitShowResourceStatement is called when production showResourceStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowResourceStatement(ctx *ShowResourceStatementContext) {}

// EnterClassifier is called when production classifier is entered.
func (s *BaseDorisSQLParserListener) EnterClassifier(ctx *ClassifierContext) {}

// ExitClassifier is called when production classifier is exited.
func (s *BaseDorisSQLParserListener) ExitClassifier(ctx *ClassifierContext) {}

// EnterShowFunctionsStatement is called when production showFunctionsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {
}

// ExitShowFunctionsStatement is called when production showFunctionsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowFunctionsStatement(ctx *ShowFunctionsStatementContext) {}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {
}

// EnterInlineFunction is called when production inlineFunction is entered.
func (s *BaseDorisSQLParserListener) EnterInlineFunction(ctx *InlineFunctionContext) {}

// ExitInlineFunction is called when production inlineFunction is exited.
func (s *BaseDorisSQLParserListener) ExitInlineFunction(ctx *InlineFunctionContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseDorisSQLParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseDorisSQLParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseDorisSQLParserListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseDorisSQLParserListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterDataDescList is called when production dataDescList is entered.
func (s *BaseDorisSQLParserListener) EnterDataDescList(ctx *DataDescListContext) {}

// ExitDataDescList is called when production dataDescList is exited.
func (s *BaseDorisSQLParserListener) ExitDataDescList(ctx *DataDescListContext) {}

// EnterDataDesc is called when production dataDesc is entered.
func (s *BaseDorisSQLParserListener) EnterDataDesc(ctx *DataDescContext) {}

// ExitDataDesc is called when production dataDesc is exited.
func (s *BaseDorisSQLParserListener) ExitDataDesc(ctx *DataDescContext) {}

// EnterFormatProps is called when production formatProps is entered.
func (s *BaseDorisSQLParserListener) EnterFormatProps(ctx *FormatPropsContext) {}

// ExitFormatProps is called when production formatProps is exited.
func (s *BaseDorisSQLParserListener) ExitFormatProps(ctx *FormatPropsContext) {}

// EnterBrokerDesc is called when production brokerDesc is entered.
func (s *BaseDorisSQLParserListener) EnterBrokerDesc(ctx *BrokerDescContext) {}

// ExitBrokerDesc is called when production brokerDesc is exited.
func (s *BaseDorisSQLParserListener) ExitBrokerDesc(ctx *BrokerDescContext) {}

// EnterResourceDesc is called when production resourceDesc is entered.
func (s *BaseDorisSQLParserListener) EnterResourceDesc(ctx *ResourceDescContext) {}

// ExitResourceDesc is called when production resourceDesc is exited.
func (s *BaseDorisSQLParserListener) ExitResourceDesc(ctx *ResourceDescContext) {}

// EnterShowLoadStatement is called when production showLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowLoadStatement(ctx *ShowLoadStatementContext) {}

// ExitShowLoadStatement is called when production showLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowLoadStatement(ctx *ShowLoadStatementContext) {}

// EnterShowLoadWarningsStatement is called when production showLoadWarningsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// ExitShowLoadWarningsStatement is called when production showLoadWarningsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowLoadWarningsStatement(ctx *ShowLoadWarningsStatementContext) {
}

// EnterCancelLoadStatement is called when production cancelLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// ExitCancelLoadStatement is called when production cancelLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelLoadStatement(ctx *CancelLoadStatementContext) {}

// EnterAlterLoadStatement is called when production alterLoadStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// ExitAlterLoadStatement is called when production alterLoadStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterLoadStatement(ctx *AlterLoadStatementContext) {}

// EnterCancelCompactionStatement is called when production cancelCompactionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// ExitCancelCompactionStatement is called when production cancelCompactionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelCompactionStatement(ctx *CancelCompactionStatementContext) {
}

// EnterShowAuthorStatement is called when production showAuthorStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// ExitShowAuthorStatement is called when production showAuthorStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowAuthorStatement(ctx *ShowAuthorStatementContext) {}

// EnterShowBackendsStatement is called when production showBackendsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// ExitShowBackendsStatement is called when production showBackendsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowBackendsStatement(ctx *ShowBackendsStatementContext) {}

// EnterShowBrokerStatement is called when production showBrokerStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// ExitShowBrokerStatement is called when production showBrokerStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowBrokerStatement(ctx *ShowBrokerStatementContext) {}

// EnterShowCharsetStatement is called when production showCharsetStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// ExitShowCharsetStatement is called when production showCharsetStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCharsetStatement(ctx *ShowCharsetStatementContext) {}

// EnterShowCollationStatement is called when production showCollationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCollationStatement(ctx *ShowCollationStatementContext) {
}

// ExitShowCollationStatement is called when production showCollationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCollationStatement(ctx *ShowCollationStatementContext) {}

// EnterShowDeleteStatement is called when production showDeleteStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// ExitShowDeleteStatement is called when production showDeleteStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowDeleteStatement(ctx *ShowDeleteStatementContext) {}

// EnterShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// ExitShowDynamicPartitionStatement is called when production showDynamicPartitionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowDynamicPartitionStatement(ctx *ShowDynamicPartitionStatementContext) {
}

// EnterShowEventsStatement is called when production showEventsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowEventsStatement(ctx *ShowEventsStatementContext) {}

// ExitShowEventsStatement is called when production showEventsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowEventsStatement(ctx *ShowEventsStatementContext) {}

// EnterShowEnginesStatement is called when production showEnginesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// ExitShowEnginesStatement is called when production showEnginesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowEnginesStatement(ctx *ShowEnginesStatementContext) {}

// EnterShowFrontendsStatement is called when production showFrontendsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {
}

// ExitShowFrontendsStatement is called when production showFrontendsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowFrontendsStatement(ctx *ShowFrontendsStatementContext) {}

// EnterShowPluginsStatement is called when production showPluginsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// ExitShowPluginsStatement is called when production showPluginsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowPluginsStatement(ctx *ShowPluginsStatementContext) {}

// EnterShowRepositoriesStatement is called when production showRepositoriesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// ExitShowRepositoriesStatement is called when production showRepositoriesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRepositoriesStatement(ctx *ShowRepositoriesStatementContext) {
}

// EnterShowOpenTableStatement is called when production showOpenTableStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {
}

// ExitShowOpenTableStatement is called when production showOpenTableStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowOpenTableStatement(ctx *ShowOpenTableStatementContext) {}

// EnterShowPrivilegesStatement is called when production showPrivilegesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {
}

// ExitShowPrivilegesStatement is called when production showPrivilegesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowPrivilegesStatement(ctx *ShowPrivilegesStatementContext) {
}

// EnterShowProcedureStatement is called when production showProcedureStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowProcedureStatement(ctx *ShowProcedureStatementContext) {
}

// ExitShowProcedureStatement is called when production showProcedureStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowProcedureStatement(ctx *ShowProcedureStatementContext) {}

// EnterShowProcStatement is called when production showProcStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowProcStatement(ctx *ShowProcStatementContext) {}

// ExitShowProcStatement is called when production showProcStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowProcStatement(ctx *ShowProcStatementContext) {}

// EnterShowProcesslistStatement is called when production showProcesslistStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// ExitShowProcesslistStatement is called when production showProcesslistStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowProcesslistStatement(ctx *ShowProcesslistStatementContext) {
}

// EnterShowProfilelistStatement is called when production showProfilelistStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// ExitShowProfilelistStatement is called when production showProfilelistStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowProfilelistStatement(ctx *ShowProfilelistStatementContext) {
}

// EnterShowRunningQueriesStatement is called when production showRunningQueriesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// ExitShowRunningQueriesStatement is called when production showRunningQueriesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRunningQueriesStatement(ctx *ShowRunningQueriesStatementContext) {
}

// EnterShowStatusStatement is called when production showStatusStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowStatusStatement(ctx *ShowStatusStatementContext) {}

// ExitShowStatusStatement is called when production showStatusStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowStatusStatement(ctx *ShowStatusStatementContext) {}

// EnterShowTabletStatement is called when production showTabletStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTabletStatement(ctx *ShowTabletStatementContext) {}

// ExitShowTabletStatement is called when production showTabletStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTabletStatement(ctx *ShowTabletStatementContext) {}

// EnterShowTransactionStatement is called when production showTransactionStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// ExitShowTransactionStatement is called when production showTransactionStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTransactionStatement(ctx *ShowTransactionStatementContext) {
}

// EnterShowTriggersStatement is called when production showTriggersStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// ExitShowTriggersStatement is called when production showTriggersStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowTriggersStatement(ctx *ShowTriggersStatementContext) {}

// EnterShowUserPropertyStatement is called when production showUserPropertyStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// ExitShowUserPropertyStatement is called when production showUserPropertyStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowUserPropertyStatement(ctx *ShowUserPropertyStatementContext) {
}

// EnterShowVariablesStatement is called when production showVariablesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowVariablesStatement(ctx *ShowVariablesStatementContext) {
}

// ExitShowVariablesStatement is called when production showVariablesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowVariablesStatement(ctx *ShowVariablesStatementContext) {}

// EnterShowWarningStatement is called when production showWarningStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowWarningStatement(ctx *ShowWarningStatementContext) {}

// ExitShowWarningStatement is called when production showWarningStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowWarningStatement(ctx *ShowWarningStatementContext) {}

// EnterHelpStatement is called when production helpStatement is entered.
func (s *BaseDorisSQLParserListener) EnterHelpStatement(ctx *HelpStatementContext) {}

// ExitHelpStatement is called when production helpStatement is exited.
func (s *BaseDorisSQLParserListener) ExitHelpStatement(ctx *HelpStatementContext) {}

// EnterCreateUserStatement is called when production createUserStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateUserStatement(ctx *CreateUserStatementContext) {}

// ExitCreateUserStatement is called when production createUserStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateUserStatement(ctx *CreateUserStatementContext) {}

// EnterDropUserStatement is called when production dropUserStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropUserStatement(ctx *DropUserStatementContext) {}

// ExitDropUserStatement is called when production dropUserStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropUserStatement(ctx *DropUserStatementContext) {}

// EnterAlterUserStatement is called when production alterUserStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterUserStatement(ctx *AlterUserStatementContext) {}

// ExitAlterUserStatement is called when production alterUserStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterUserStatement(ctx *AlterUserStatementContext) {}

// EnterShowUserStatement is called when production showUserStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowUserStatement(ctx *ShowUserStatementContext) {}

// ExitShowUserStatement is called when production showUserStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowUserStatement(ctx *ShowUserStatementContext) {}

// EnterShowAllAuthentication is called when production showAllAuthentication is entered.
func (s *BaseDorisSQLParserListener) EnterShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// ExitShowAllAuthentication is called when production showAllAuthentication is exited.
func (s *BaseDorisSQLParserListener) ExitShowAllAuthentication(ctx *ShowAllAuthenticationContext) {}

// EnterShowAuthenticationForUser is called when production showAuthenticationForUser is entered.
func (s *BaseDorisSQLParserListener) EnterShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// ExitShowAuthenticationForUser is called when production showAuthenticationForUser is exited.
func (s *BaseDorisSQLParserListener) ExitShowAuthenticationForUser(ctx *ShowAuthenticationForUserContext) {
}

// EnterExecuteAsStatement is called when production executeAsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// ExitExecuteAsStatement is called when production executeAsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitExecuteAsStatement(ctx *ExecuteAsStatementContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterAlterRoleStatement is called when production alterRoleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// ExitAlterRoleStatement is called when production alterRoleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterRoleStatement(ctx *AlterRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterShowRolesStatement is called when production showRolesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRolesStatement(ctx *ShowRolesStatementContext) {}

// ExitShowRolesStatement is called when production showRolesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRolesStatement(ctx *ShowRolesStatementContext) {}

// EnterGrantRoleToUser is called when production grantRoleToUser is entered.
func (s *BaseDorisSQLParserListener) EnterGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// ExitGrantRoleToUser is called when production grantRoleToUser is exited.
func (s *BaseDorisSQLParserListener) ExitGrantRoleToUser(ctx *GrantRoleToUserContext) {}

// EnterGrantRoleToRole is called when production grantRoleToRole is entered.
func (s *BaseDorisSQLParserListener) EnterGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// ExitGrantRoleToRole is called when production grantRoleToRole is exited.
func (s *BaseDorisSQLParserListener) ExitGrantRoleToRole(ctx *GrantRoleToRoleContext) {}

// EnterRevokeRoleFromUser is called when production revokeRoleFromUser is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// ExitRevokeRoleFromUser is called when production revokeRoleFromUser is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeRoleFromUser(ctx *RevokeRoleFromUserContext) {}

// EnterRevokeRoleFromRole is called when production revokeRoleFromRole is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// ExitRevokeRoleFromRole is called when production revokeRoleFromRole is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeRoleFromRole(ctx *RevokeRoleFromRoleContext) {}

// EnterSetRoleStatement is called when production setRoleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetRoleStatement(ctx *SetRoleStatementContext) {}

// ExitSetRoleStatement is called when production setRoleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetRoleStatement(ctx *SetRoleStatementContext) {}

// EnterSetDefaultRoleStatement is called when production setDefaultRoleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {
}

// ExitSetDefaultRoleStatement is called when production setDefaultRoleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetDefaultRoleStatement(ctx *SetDefaultRoleStatementContext) {
}

// EnterGrantRevokeClause is called when production grantRevokeClause is entered.
func (s *BaseDorisSQLParserListener) EnterGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// ExitGrantRevokeClause is called when production grantRevokeClause is exited.
func (s *BaseDorisSQLParserListener) ExitGrantRevokeClause(ctx *GrantRevokeClauseContext) {}

// EnterGrantOnUser is called when production grantOnUser is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnUser(ctx *GrantOnUserContext) {}

// ExitGrantOnUser is called when production grantOnUser is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnUser(ctx *GrantOnUserContext) {}

// EnterGrantOnTableBrief is called when production grantOnTableBrief is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// ExitGrantOnTableBrief is called when production grantOnTableBrief is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnTableBrief(ctx *GrantOnTableBriefContext) {}

// EnterGrantOnFunc is called when production grantOnFunc is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnFunc(ctx *GrantOnFuncContext) {}

// ExitGrantOnFunc is called when production grantOnFunc is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnFunc(ctx *GrantOnFuncContext) {}

// EnterGrantOnSystem is called when production grantOnSystem is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnSystem(ctx *GrantOnSystemContext) {}

// ExitGrantOnSystem is called when production grantOnSystem is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnSystem(ctx *GrantOnSystemContext) {}

// EnterGrantOnPrimaryObj is called when production grantOnPrimaryObj is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// ExitGrantOnPrimaryObj is called when production grantOnPrimaryObj is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnPrimaryObj(ctx *GrantOnPrimaryObjContext) {}

// EnterGrantOnAll is called when production grantOnAll is entered.
func (s *BaseDorisSQLParserListener) EnterGrantOnAll(ctx *GrantOnAllContext) {}

// ExitGrantOnAll is called when production grantOnAll is exited.
func (s *BaseDorisSQLParserListener) ExitGrantOnAll(ctx *GrantOnAllContext) {}

// EnterRevokeOnUser is called when production revokeOnUser is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnUser(ctx *RevokeOnUserContext) {}

// ExitRevokeOnUser is called when production revokeOnUser is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnUser(ctx *RevokeOnUserContext) {}

// EnterRevokeOnTableBrief is called when production revokeOnTableBrief is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// ExitRevokeOnTableBrief is called when production revokeOnTableBrief is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnTableBrief(ctx *RevokeOnTableBriefContext) {}

// EnterRevokeOnFunc is called when production revokeOnFunc is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// ExitRevokeOnFunc is called when production revokeOnFunc is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnFunc(ctx *RevokeOnFuncContext) {}

// EnterRevokeOnSystem is called when production revokeOnSystem is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// ExitRevokeOnSystem is called when production revokeOnSystem is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnSystem(ctx *RevokeOnSystemContext) {}

// EnterRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// ExitRevokeOnPrimaryObj is called when production revokeOnPrimaryObj is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnPrimaryObj(ctx *RevokeOnPrimaryObjContext) {}

// EnterRevokeOnAll is called when production revokeOnAll is entered.
func (s *BaseDorisSQLParserListener) EnterRevokeOnAll(ctx *RevokeOnAllContext) {}

// ExitRevokeOnAll is called when production revokeOnAll is exited.
func (s *BaseDorisSQLParserListener) ExitRevokeOnAll(ctx *RevokeOnAllContext) {}

// EnterShowGrantsStatement is called when production showGrantsStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// ExitShowGrantsStatement is called when production showGrantsStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowGrantsStatement(ctx *ShowGrantsStatementContext) {}

// EnterAuthWithoutPlugin is called when production authWithoutPlugin is entered.
func (s *BaseDorisSQLParserListener) EnterAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// ExitAuthWithoutPlugin is called when production authWithoutPlugin is exited.
func (s *BaseDorisSQLParserListener) ExitAuthWithoutPlugin(ctx *AuthWithoutPluginContext) {}

// EnterAuthWithPlugin is called when production authWithPlugin is entered.
func (s *BaseDorisSQLParserListener) EnterAuthWithPlugin(ctx *AuthWithPluginContext) {}

// ExitAuthWithPlugin is called when production authWithPlugin is exited.
func (s *BaseDorisSQLParserListener) ExitAuthWithPlugin(ctx *AuthWithPluginContext) {}

// EnterPrivObjectName is called when production privObjectName is entered.
func (s *BaseDorisSQLParserListener) EnterPrivObjectName(ctx *PrivObjectNameContext) {}

// ExitPrivObjectName is called when production privObjectName is exited.
func (s *BaseDorisSQLParserListener) ExitPrivObjectName(ctx *PrivObjectNameContext) {}

// EnterPrivObjectNameList is called when production privObjectNameList is entered.
func (s *BaseDorisSQLParserListener) EnterPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// ExitPrivObjectNameList is called when production privObjectNameList is exited.
func (s *BaseDorisSQLParserListener) ExitPrivObjectNameList(ctx *PrivObjectNameListContext) {}

// EnterPrivFunctionObjectNameList is called when production privFunctionObjectNameList is entered.
func (s *BaseDorisSQLParserListener) EnterPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// ExitPrivFunctionObjectNameList is called when production privFunctionObjectNameList is exited.
func (s *BaseDorisSQLParserListener) ExitPrivFunctionObjectNameList(ctx *PrivFunctionObjectNameListContext) {
}

// EnterPrivilegeTypeList is called when production privilegeTypeList is entered.
func (s *BaseDorisSQLParserListener) EnterPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// ExitPrivilegeTypeList is called when production privilegeTypeList is exited.
func (s *BaseDorisSQLParserListener) ExitPrivilegeTypeList(ctx *PrivilegeTypeListContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseDorisSQLParserListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseDorisSQLParserListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterPrivObjectType is called when production privObjectType is entered.
func (s *BaseDorisSQLParserListener) EnterPrivObjectType(ctx *PrivObjectTypeContext) {}

// ExitPrivObjectType is called when production privObjectType is exited.
func (s *BaseDorisSQLParserListener) ExitPrivObjectType(ctx *PrivObjectTypeContext) {}

// EnterPrivObjectTypePlural is called when production privObjectTypePlural is entered.
func (s *BaseDorisSQLParserListener) EnterPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// ExitPrivObjectTypePlural is called when production privObjectTypePlural is exited.
func (s *BaseDorisSQLParserListener) ExitPrivObjectTypePlural(ctx *PrivObjectTypePluralContext) {}

// EnterCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// ExitCreateSecurityIntegrationStatement is called when production createSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateSecurityIntegrationStatement(ctx *CreateSecurityIntegrationStatementContext) {
}

// EnterAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// ExitAlterSecurityIntegrationStatement is called when production alterSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterSecurityIntegrationStatement(ctx *AlterSecurityIntegrationStatementContext) {
}

// EnterDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// ExitDropSecurityIntegrationStatement is called when production dropSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropSecurityIntegrationStatement(ctx *DropSecurityIntegrationStatementContext) {
}

// EnterShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// ExitShowSecurityIntegrationStatement is called when production showSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowSecurityIntegrationStatement(ctx *ShowSecurityIntegrationStatementContext) {
}

// EnterShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// ExitShowCreateSecurityIntegrationStatement is called when production showCreateSecurityIntegrationStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateSecurityIntegrationStatement(ctx *ShowCreateSecurityIntegrationStatementContext) {
}

// EnterCreateGroupProviderStatement is called when production createGroupProviderStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// ExitCreateGroupProviderStatement is called when production createGroupProviderStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateGroupProviderStatement(ctx *CreateGroupProviderStatementContext) {
}

// EnterDropGroupProviderStatement is called when production dropGroupProviderStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// ExitDropGroupProviderStatement is called when production dropGroupProviderStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropGroupProviderStatement(ctx *DropGroupProviderStatementContext) {
}

// EnterShowGroupProvidersStatement is called when production showGroupProvidersStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// ExitShowGroupProvidersStatement is called when production showGroupProvidersStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowGroupProvidersStatement(ctx *ShowGroupProvidersStatementContext) {
}

// EnterShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// ExitShowCreateGroupProviderStatement is called when production showCreateGroupProviderStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowCreateGroupProviderStatement(ctx *ShowCreateGroupProviderStatementContext) {
}

// EnterBackupStatement is called when production backupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterBackupStatement(ctx *BackupStatementContext) {}

// ExitBackupStatement is called when production backupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitBackupStatement(ctx *BackupStatementContext) {}

// EnterCancelBackupStatement is called when production cancelBackupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// ExitCancelBackupStatement is called when production cancelBackupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelBackupStatement(ctx *CancelBackupStatementContext) {}

// EnterShowBackupStatement is called when production showBackupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowBackupStatement(ctx *ShowBackupStatementContext) {}

// ExitShowBackupStatement is called when production showBackupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowBackupStatement(ctx *ShowBackupStatementContext) {}

// EnterRestoreStatement is called when production restoreStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRestoreStatement(ctx *RestoreStatementContext) {}

// ExitRestoreStatement is called when production restoreStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRestoreStatement(ctx *RestoreStatementContext) {}

// EnterCancelRestoreStatement is called when production cancelRestoreStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelRestoreStatement(ctx *CancelRestoreStatementContext) {
}

// ExitCancelRestoreStatement is called when production cancelRestoreStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelRestoreStatement(ctx *CancelRestoreStatementContext) {}

// EnterShowRestoreStatement is called when production showRestoreStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// ExitShowRestoreStatement is called when production showRestoreStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowRestoreStatement(ctx *ShowRestoreStatementContext) {}

// EnterShowSnapshotStatement is called when production showSnapshotStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// ExitShowSnapshotStatement is called when production showSnapshotStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowSnapshotStatement(ctx *ShowSnapshotStatementContext) {}

// EnterCreateRepositoryStatement is called when production createRepositoryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// ExitCreateRepositoryStatement is called when production createRepositoryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateRepositoryStatement(ctx *CreateRepositoryStatementContext) {
}

// EnterDropRepositoryStatement is called when production dropRepositoryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropRepositoryStatement(ctx *DropRepositoryStatementContext) {
}

// ExitDropRepositoryStatement is called when production dropRepositoryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropRepositoryStatement(ctx *DropRepositoryStatementContext) {
}

// EnterAddSqlBlackListStatement is called when production addSqlBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// ExitAddSqlBlackListStatement is called when production addSqlBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAddSqlBlackListStatement(ctx *AddSqlBlackListStatementContext) {
}

// EnterDelSqlBlackListStatement is called when production delSqlBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// ExitDelSqlBlackListStatement is called when production delSqlBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDelSqlBlackListStatement(ctx *DelSqlBlackListStatementContext) {
}

// EnterShowSqlBlackListStatement is called when production showSqlBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// ExitShowSqlBlackListStatement is called when production showSqlBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowSqlBlackListStatement(ctx *ShowSqlBlackListStatementContext) {
}

// EnterShowWhiteListStatement is called when production showWhiteListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {
}

// ExitShowWhiteListStatement is called when production showWhiteListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowWhiteListStatement(ctx *ShowWhiteListStatementContext) {}

// EnterAddBackendBlackListStatement is called when production addBackendBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// ExitAddBackendBlackListStatement is called when production addBackendBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAddBackendBlackListStatement(ctx *AddBackendBlackListStatementContext) {
}

// EnterDelBackendBlackListStatement is called when production delBackendBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// ExitDelBackendBlackListStatement is called when production delBackendBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDelBackendBlackListStatement(ctx *DelBackendBlackListStatementContext) {
}

// EnterShowBackendBlackListStatement is called when production showBackendBlackListStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// ExitShowBackendBlackListStatement is called when production showBackendBlackListStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowBackendBlackListStatement(ctx *ShowBackendBlackListStatementContext) {
}

// EnterDataCacheTarget is called when production dataCacheTarget is entered.
func (s *BaseDorisSQLParserListener) EnterDataCacheTarget(ctx *DataCacheTargetContext) {}

// ExitDataCacheTarget is called when production dataCacheTarget is exited.
func (s *BaseDorisSQLParserListener) ExitDataCacheTarget(ctx *DataCacheTargetContext) {}

// EnterCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// ExitCreateDataCacheRuleStatement is called when production createDataCacheRuleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateDataCacheRuleStatement(ctx *CreateDataCacheRuleStatementContext) {
}

// EnterShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// ExitShowDataCacheRulesStatement is called when production showDataCacheRulesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowDataCacheRulesStatement(ctx *ShowDataCacheRulesStatementContext) {
}

// EnterDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// ExitDropDataCacheRuleStatement is called when production dropDataCacheRuleStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropDataCacheRuleStatement(ctx *DropDataCacheRuleStatementContext) {
}

// EnterClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// ExitClearDataCacheRulesStatement is called when production clearDataCacheRulesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitClearDataCacheRulesStatement(ctx *ClearDataCacheRulesStatementContext) {
}

// EnterDataCacheSelectStatement is called when production dataCacheSelectStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// ExitDataCacheSelectStatement is called when production dataCacheSelectStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDataCacheSelectStatement(ctx *DataCacheSelectStatementContext) {
}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseDorisSQLParserListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseDorisSQLParserListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterCancelExportStatement is called when production cancelExportStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCancelExportStatement(ctx *CancelExportStatementContext) {}

// ExitCancelExportStatement is called when production cancelExportStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCancelExportStatement(ctx *CancelExportStatementContext) {}

// EnterShowExportStatement is called when production showExportStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowExportStatement(ctx *ShowExportStatementContext) {}

// ExitShowExportStatement is called when production showExportStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowExportStatement(ctx *ShowExportStatementContext) {}

// EnterInstallPluginStatement is called when production installPluginStatement is entered.
func (s *BaseDorisSQLParserListener) EnterInstallPluginStatement(ctx *InstallPluginStatementContext) {
}

// ExitInstallPluginStatement is called when production installPluginStatement is exited.
func (s *BaseDorisSQLParserListener) ExitInstallPluginStatement(ctx *InstallPluginStatementContext) {}

// EnterUninstallPluginStatement is called when production uninstallPluginStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// ExitUninstallPluginStatement is called when production uninstallPluginStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUninstallPluginStatement(ctx *UninstallPluginStatementContext) {
}

// EnterCreateFileStatement is called when production createFileStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateFileStatement(ctx *CreateFileStatementContext) {}

// ExitCreateFileStatement is called when production createFileStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateFileStatement(ctx *CreateFileStatementContext) {}

// EnterDropFileStatement is called when production dropFileStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropFileStatement(ctx *DropFileStatementContext) {}

// ExitDropFileStatement is called when production dropFileStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropFileStatement(ctx *DropFileStatementContext) {}

// EnterShowSmallFilesStatement is called when production showSmallFilesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {
}

// ExitShowSmallFilesStatement is called when production showSmallFilesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowSmallFilesStatement(ctx *ShowSmallFilesStatementContext) {
}

// EnterCreatePipeStatement is called when production createPipeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// ExitCreatePipeStatement is called when production createPipeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreatePipeStatement(ctx *CreatePipeStatementContext) {}

// EnterDropPipeStatement is called when production dropPipeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropPipeStatement(ctx *DropPipeStatementContext) {}

// ExitDropPipeStatement is called when production dropPipeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropPipeStatement(ctx *DropPipeStatementContext) {}

// EnterAlterPipeClause is called when production alterPipeClause is entered.
func (s *BaseDorisSQLParserListener) EnterAlterPipeClause(ctx *AlterPipeClauseContext) {}

// ExitAlterPipeClause is called when production alterPipeClause is exited.
func (s *BaseDorisSQLParserListener) ExitAlterPipeClause(ctx *AlterPipeClauseContext) {}

// EnterAlterPipeStatement is called when production alterPipeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// ExitAlterPipeStatement is called when production alterPipeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterPipeStatement(ctx *AlterPipeStatementContext) {}

// EnterDescPipeStatement is called when production descPipeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDescPipeStatement(ctx *DescPipeStatementContext) {}

// ExitDescPipeStatement is called when production descPipeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDescPipeStatement(ctx *DescPipeStatementContext) {}

// EnterShowPipeStatement is called when production showPipeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowPipeStatement(ctx *ShowPipeStatementContext) {}

// ExitShowPipeStatement is called when production showPipeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowPipeStatement(ctx *ShowPipeStatementContext) {}

// EnterSetStatement is called when production setStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetStatement(ctx *SetStatementContext) {}

// ExitSetStatement is called when production setStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetStatement(ctx *SetStatementContext) {}

// EnterSetNames is called when production setNames is entered.
func (s *BaseDorisSQLParserListener) EnterSetNames(ctx *SetNamesContext) {}

// ExitSetNames is called when production setNames is exited.
func (s *BaseDorisSQLParserListener) ExitSetNames(ctx *SetNamesContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseDorisSQLParserListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseDorisSQLParserListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterSetUserVar is called when production setUserVar is entered.
func (s *BaseDorisSQLParserListener) EnterSetUserVar(ctx *SetUserVarContext) {}

// ExitSetUserVar is called when production setUserVar is exited.
func (s *BaseDorisSQLParserListener) ExitSetUserVar(ctx *SetUserVarContext) {}

// EnterSetSystemVar is called when production setSystemVar is entered.
func (s *BaseDorisSQLParserListener) EnterSetSystemVar(ctx *SetSystemVarContext) {}

// ExitSetSystemVar is called when production setSystemVar is exited.
func (s *BaseDorisSQLParserListener) ExitSetSystemVar(ctx *SetSystemVarContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseDorisSQLParserListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseDorisSQLParserListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterTransaction_characteristics is called when production transaction_characteristics is entered.
func (s *BaseDorisSQLParserListener) EnterTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// ExitTransaction_characteristics is called when production transaction_characteristics is exited.
func (s *BaseDorisSQLParserListener) ExitTransaction_characteristics(ctx *Transaction_characteristicsContext) {
}

// EnterTransaction_access_mode is called when production transaction_access_mode is entered.
func (s *BaseDorisSQLParserListener) EnterTransaction_access_mode(ctx *Transaction_access_modeContext) {
}

// ExitTransaction_access_mode is called when production transaction_access_mode is exited.
func (s *BaseDorisSQLParserListener) ExitTransaction_access_mode(ctx *Transaction_access_modeContext) {
}

// EnterIsolation_level is called when production isolation_level is entered.
func (s *BaseDorisSQLParserListener) EnterIsolation_level(ctx *Isolation_levelContext) {}

// ExitIsolation_level is called when production isolation_level is exited.
func (s *BaseDorisSQLParserListener) ExitIsolation_level(ctx *Isolation_levelContext) {}

// EnterIsolation_types is called when production isolation_types is entered.
func (s *BaseDorisSQLParserListener) EnterIsolation_types(ctx *Isolation_typesContext) {}

// ExitIsolation_types is called when production isolation_types is exited.
func (s *BaseDorisSQLParserListener) ExitIsolation_types(ctx *Isolation_typesContext) {}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseDorisSQLParserListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseDorisSQLParserListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterSetUserPropertyStatement is called when production setUserPropertyStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// ExitSetUserPropertyStatement is called when production setUserPropertyStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetUserPropertyStatement(ctx *SetUserPropertyStatementContext) {
}

// EnterRoleList is called when production roleList is entered.
func (s *BaseDorisSQLParserListener) EnterRoleList(ctx *RoleListContext) {}

// ExitRoleList is called when production roleList is exited.
func (s *BaseDorisSQLParserListener) ExitRoleList(ctx *RoleListContext) {}

// EnterExecuteScriptStatement is called when production executeScriptStatement is entered.
func (s *BaseDorisSQLParserListener) EnterExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {
}

// ExitExecuteScriptStatement is called when production executeScriptStatement is exited.
func (s *BaseDorisSQLParserListener) ExitExecuteScriptStatement(ctx *ExecuteScriptStatementContext) {}

// EnterUnsupportedStatement is called when production unsupportedStatement is entered.
func (s *BaseDorisSQLParserListener) EnterUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// ExitUnsupportedStatement is called when production unsupportedStatement is exited.
func (s *BaseDorisSQLParserListener) ExitUnsupportedStatement(ctx *UnsupportedStatementContext) {}

// EnterLock_item is called when production lock_item is entered.
func (s *BaseDorisSQLParserListener) EnterLock_item(ctx *Lock_itemContext) {}

// ExitLock_item is called when production lock_item is exited.
func (s *BaseDorisSQLParserListener) ExitLock_item(ctx *Lock_itemContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseDorisSQLParserListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseDorisSQLParserListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// ExitAlterPlanAdvisorAddStatement is called when production alterPlanAdvisorAddStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterPlanAdvisorAddStatement(ctx *AlterPlanAdvisorAddStatementContext) {
}

// EnterTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is entered.
func (s *BaseDorisSQLParserListener) EnterTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// ExitTruncatePlanAdvisorStatement is called when production truncatePlanAdvisorStatement is exited.
func (s *BaseDorisSQLParserListener) ExitTruncatePlanAdvisorStatement(ctx *TruncatePlanAdvisorStatementContext) {
}

// EnterAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// ExitAlterPlanAdvisorDropStatement is called when production alterPlanAdvisorDropStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterPlanAdvisorDropStatement(ctx *AlterPlanAdvisorDropStatementContext) {
}

// EnterShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// ExitShowPlanAdvisorStatement is called when production showPlanAdvisorStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowPlanAdvisorStatement(ctx *ShowPlanAdvisorStatementContext) {
}

// EnterCreateWarehouseStatement is called when production createWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// ExitCreateWarehouseStatement is called when production createWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateWarehouseStatement(ctx *CreateWarehouseStatementContext) {
}

// EnterDropWarehouseStatement is called when production dropWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropWarehouseStatement(ctx *DropWarehouseStatementContext) {
}

// ExitDropWarehouseStatement is called when production dropWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropWarehouseStatement(ctx *DropWarehouseStatementContext) {}

// EnterSuspendWarehouseStatement is called when production suspendWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// ExitSuspendWarehouseStatement is called when production suspendWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSuspendWarehouseStatement(ctx *SuspendWarehouseStatementContext) {
}

// EnterResumeWarehouseStatement is called when production resumeWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// ExitResumeWarehouseStatement is called when production resumeWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitResumeWarehouseStatement(ctx *ResumeWarehouseStatementContext) {
}

// EnterSetWarehouseStatement is called when production setWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// ExitSetWarehouseStatement is called when production setWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitSetWarehouseStatement(ctx *SetWarehouseStatementContext) {}

// EnterShowWarehousesStatement is called when production showWarehousesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {
}

// ExitShowWarehousesStatement is called when production showWarehousesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowWarehousesStatement(ctx *ShowWarehousesStatementContext) {
}

// EnterShowClustersStatement is called when production showClustersStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowClustersStatement(ctx *ShowClustersStatementContext) {}

// ExitShowClustersStatement is called when production showClustersStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowClustersStatement(ctx *ShowClustersStatementContext) {}

// EnterShowNodesStatement is called when production showNodesStatement is entered.
func (s *BaseDorisSQLParserListener) EnterShowNodesStatement(ctx *ShowNodesStatementContext) {}

// ExitShowNodesStatement is called when production showNodesStatement is exited.
func (s *BaseDorisSQLParserListener) ExitShowNodesStatement(ctx *ShowNodesStatementContext) {}

// EnterAlterWarehouseStatement is called when production alterWarehouseStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {
}

// ExitAlterWarehouseStatement is called when production alterWarehouseStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterWarehouseStatement(ctx *AlterWarehouseStatementContext) {
}

// EnterCreateCNGroupStatement is called when production createCNGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {
}

// ExitCreateCNGroupStatement is called when production createCNGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCreateCNGroupStatement(ctx *CreateCNGroupStatementContext) {}

// EnterDropCNGroupStatement is called when production dropCNGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// ExitDropCNGroupStatement is called when production dropCNGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDropCNGroupStatement(ctx *DropCNGroupStatementContext) {}

// EnterEnableCNGroupStatement is called when production enableCNGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {
}

// ExitEnableCNGroupStatement is called when production enableCNGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitEnableCNGroupStatement(ctx *EnableCNGroupStatementContext) {}

// EnterDisableCNGroupStatement is called when production disableCNGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {
}

// ExitDisableCNGroupStatement is called when production disableCNGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDisableCNGroupStatement(ctx *DisableCNGroupStatementContext) {
}

// EnterAlterCNGroupStatement is called when production alterCNGroupStatement is entered.
func (s *BaseDorisSQLParserListener) EnterAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// ExitAlterCNGroupStatement is called when production alterCNGroupStatement is exited.
func (s *BaseDorisSQLParserListener) ExitAlterCNGroupStatement(ctx *AlterCNGroupStatementContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseDorisSQLParserListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseDorisSQLParserListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseDorisSQLParserListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseDorisSQLParserListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseDorisSQLParserListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseDorisSQLParserListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterTranslateStatement is called when production translateStatement is entered.
func (s *BaseDorisSQLParserListener) EnterTranslateStatement(ctx *TranslateStatementContext) {}

// ExitTranslateStatement is called when production translateStatement is exited.
func (s *BaseDorisSQLParserListener) ExitTranslateStatement(ctx *TranslateStatementContext) {}

// EnterDialect is called when production dialect is entered.
func (s *BaseDorisSQLParserListener) EnterDialect(ctx *DialectContext) {}

// ExitDialect is called when production dialect is exited.
func (s *BaseDorisSQLParserListener) ExitDialect(ctx *DialectContext) {}

// EnterTranslateSQL is called when production translateSQL is entered.
func (s *BaseDorisSQLParserListener) EnterTranslateSQL(ctx *TranslateSQLContext) {}

// ExitTranslateSQL is called when production translateSQL is exited.
func (s *BaseDorisSQLParserListener) ExitTranslateSQL(ctx *TranslateSQLContext) {}

// EnterQueryStatement is called when production queryStatement is entered.
func (s *BaseDorisSQLParserListener) EnterQueryStatement(ctx *QueryStatementContext) {}

// ExitQueryStatement is called when production queryStatement is exited.
func (s *BaseDorisSQLParserListener) ExitQueryStatement(ctx *QueryStatementContext) {}

// EnterQueryRelation is called when production queryRelation is entered.
func (s *BaseDorisSQLParserListener) EnterQueryRelation(ctx *QueryRelationContext) {}

// ExitQueryRelation is called when production queryRelation is exited.
func (s *BaseDorisSQLParserListener) ExitQueryRelation(ctx *QueryRelationContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseDorisSQLParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseDorisSQLParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterQueryNoWith is called when production queryNoWith is entered.
func (s *BaseDorisSQLParserListener) EnterQueryNoWith(ctx *QueryNoWithContext) {}

// ExitQueryNoWith is called when production queryNoWith is exited.
func (s *BaseDorisSQLParserListener) ExitQueryNoWith(ctx *QueryNoWithContext) {}

// EnterQueryPeriod is called when production queryPeriod is entered.
func (s *BaseDorisSQLParserListener) EnterQueryPeriod(ctx *QueryPeriodContext) {}

// ExitQueryPeriod is called when production queryPeriod is exited.
func (s *BaseDorisSQLParserListener) ExitQueryPeriod(ctx *QueryPeriodContext) {}

// EnterPeriodType is called when production periodType is entered.
func (s *BaseDorisSQLParserListener) EnterPeriodType(ctx *PeriodTypeContext) {}

// ExitPeriodType is called when production periodType is exited.
func (s *BaseDorisSQLParserListener) ExitPeriodType(ctx *PeriodTypeContext) {}

// EnterQueryWithParentheses is called when production queryWithParentheses is entered.
func (s *BaseDorisSQLParserListener) EnterQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// ExitQueryWithParentheses is called when production queryWithParentheses is exited.
func (s *BaseDorisSQLParserListener) ExitQueryWithParentheses(ctx *QueryWithParenthesesContext) {}

// EnterSetOperation is called when production setOperation is entered.
func (s *BaseDorisSQLParserListener) EnterSetOperation(ctx *SetOperationContext) {}

// ExitSetOperation is called when production setOperation is exited.
func (s *BaseDorisSQLParserListener) ExitSetOperation(ctx *SetOperationContext) {}

// EnterQueryPrimaryDefault is called when production queryPrimaryDefault is entered.
func (s *BaseDorisSQLParserListener) EnterQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// ExitQueryPrimaryDefault is called when production queryPrimaryDefault is exited.
func (s *BaseDorisSQLParserListener) ExitQueryPrimaryDefault(ctx *QueryPrimaryDefaultContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseDorisSQLParserListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseDorisSQLParserListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterRowConstructor is called when production rowConstructor is entered.
func (s *BaseDorisSQLParserListener) EnterRowConstructor(ctx *RowConstructorContext) {}

// ExitRowConstructor is called when production rowConstructor is exited.
func (s *BaseDorisSQLParserListener) ExitRowConstructor(ctx *RowConstructorContext) {}

// EnterSortItem is called when production sortItem is entered.
func (s *BaseDorisSQLParserListener) EnterSortItem(ctx *SortItemContext) {}

// ExitSortItem is called when production sortItem is exited.
func (s *BaseDorisSQLParserListener) ExitSortItem(ctx *SortItemContext) {}

// EnterLimitConstExpr is called when production limitConstExpr is entered.
func (s *BaseDorisSQLParserListener) EnterLimitConstExpr(ctx *LimitConstExprContext) {}

// ExitLimitConstExpr is called when production limitConstExpr is exited.
func (s *BaseDorisSQLParserListener) ExitLimitConstExpr(ctx *LimitConstExprContext) {}

// EnterLimitElement is called when production limitElement is entered.
func (s *BaseDorisSQLParserListener) EnterLimitElement(ctx *LimitElementContext) {}

// ExitLimitElement is called when production limitElement is exited.
func (s *BaseDorisSQLParserListener) ExitLimitElement(ctx *LimitElementContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseDorisSQLParserListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseDorisSQLParserListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseDorisSQLParserListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseDorisSQLParserListener) ExitFrom(ctx *FromContext) {}

// EnterDual is called when production dual is entered.
func (s *BaseDorisSQLParserListener) EnterDual(ctx *DualContext) {}

// ExitDual is called when production dual is exited.
func (s *BaseDorisSQLParserListener) ExitDual(ctx *DualContext) {}

// EnterRollup is called when production rollup is entered.
func (s *BaseDorisSQLParserListener) EnterRollup(ctx *RollupContext) {}

// ExitRollup is called when production rollup is exited.
func (s *BaseDorisSQLParserListener) ExitRollup(ctx *RollupContext) {}

// EnterCube is called when production cube is entered.
func (s *BaseDorisSQLParserListener) EnterCube(ctx *CubeContext) {}

// ExitCube is called when production cube is exited.
func (s *BaseDorisSQLParserListener) ExitCube(ctx *CubeContext) {}

// EnterMultipleGroupingSets is called when production multipleGroupingSets is entered.
func (s *BaseDorisSQLParserListener) EnterMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// ExitMultipleGroupingSets is called when production multipleGroupingSets is exited.
func (s *BaseDorisSQLParserListener) ExitMultipleGroupingSets(ctx *MultipleGroupingSetsContext) {}

// EnterSingleGroupingSet is called when production singleGroupingSet is entered.
func (s *BaseDorisSQLParserListener) EnterSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// ExitSingleGroupingSet is called when production singleGroupingSet is exited.
func (s *BaseDorisSQLParserListener) ExitSingleGroupingSet(ctx *SingleGroupingSetContext) {}

// EnterGroupingSet is called when production groupingSet is entered.
func (s *BaseDorisSQLParserListener) EnterGroupingSet(ctx *GroupingSetContext) {}

// ExitGroupingSet is called when production groupingSet is exited.
func (s *BaseDorisSQLParserListener) ExitGroupingSet(ctx *GroupingSetContext) {}

// EnterCommonTableExpression is called when production commonTableExpression is entered.
func (s *BaseDorisSQLParserListener) EnterCommonTableExpression(ctx *CommonTableExpressionContext) {}

// ExitCommonTableExpression is called when production commonTableExpression is exited.
func (s *BaseDorisSQLParserListener) ExitCommonTableExpression(ctx *CommonTableExpressionContext) {}

// EnterSetQuantifier is called when production setQuantifier is entered.
func (s *BaseDorisSQLParserListener) EnterSetQuantifier(ctx *SetQuantifierContext) {}

// ExitSetQuantifier is called when production setQuantifier is exited.
func (s *BaseDorisSQLParserListener) ExitSetQuantifier(ctx *SetQuantifierContext) {}

// EnterSelectSingle is called when production selectSingle is entered.
func (s *BaseDorisSQLParserListener) EnterSelectSingle(ctx *SelectSingleContext) {}

// ExitSelectSingle is called when production selectSingle is exited.
func (s *BaseDorisSQLParserListener) ExitSelectSingle(ctx *SelectSingleContext) {}

// EnterSelectAll is called when production selectAll is entered.
func (s *BaseDorisSQLParserListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production selectAll is exited.
func (s *BaseDorisSQLParserListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterExcludeClause is called when production excludeClause is entered.
func (s *BaseDorisSQLParserListener) EnterExcludeClause(ctx *ExcludeClauseContext) {}

// ExitExcludeClause is called when production excludeClause is exited.
func (s *BaseDorisSQLParserListener) ExitExcludeClause(ctx *ExcludeClauseContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseDorisSQLParserListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseDorisSQLParserListener) ExitRelations(ctx *RelationsContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseDorisSQLParserListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseDorisSQLParserListener) ExitRelation(ctx *RelationContext) {}

// EnterTableAtom is called when production tableAtom is entered.
func (s *BaseDorisSQLParserListener) EnterTableAtom(ctx *TableAtomContext) {}

// ExitTableAtom is called when production tableAtom is exited.
func (s *BaseDorisSQLParserListener) ExitTableAtom(ctx *TableAtomContext) {}

// EnterInlineTable is called when production inlineTable is entered.
func (s *BaseDorisSQLParserListener) EnterInlineTable(ctx *InlineTableContext) {}

// ExitInlineTable is called when production inlineTable is exited.
func (s *BaseDorisSQLParserListener) ExitInlineTable(ctx *InlineTableContext) {}

// EnterSubqueryWithAlias is called when production subqueryWithAlias is entered.
func (s *BaseDorisSQLParserListener) EnterSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// ExitSubqueryWithAlias is called when production subqueryWithAlias is exited.
func (s *BaseDorisSQLParserListener) ExitSubqueryWithAlias(ctx *SubqueryWithAliasContext) {}

// EnterTableFunction is called when production tableFunction is entered.
func (s *BaseDorisSQLParserListener) EnterTableFunction(ctx *TableFunctionContext) {}

// ExitTableFunction is called when production tableFunction is exited.
func (s *BaseDorisSQLParserListener) ExitTableFunction(ctx *TableFunctionContext) {}

// EnterNormalizedTableFunction is called when production normalizedTableFunction is entered.
func (s *BaseDorisSQLParserListener) EnterNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {
}

// ExitNormalizedTableFunction is called when production normalizedTableFunction is exited.
func (s *BaseDorisSQLParserListener) ExitNormalizedTableFunction(ctx *NormalizedTableFunctionContext) {
}

// EnterFileTableFunction is called when production fileTableFunction is entered.
func (s *BaseDorisSQLParserListener) EnterFileTableFunction(ctx *FileTableFunctionContext) {}

// ExitFileTableFunction is called when production fileTableFunction is exited.
func (s *BaseDorisSQLParserListener) ExitFileTableFunction(ctx *FileTableFunctionContext) {}

// EnterParenthesizedRelation is called when production parenthesizedRelation is entered.
func (s *BaseDorisSQLParserListener) EnterParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// ExitParenthesizedRelation is called when production parenthesizedRelation is exited.
func (s *BaseDorisSQLParserListener) ExitParenthesizedRelation(ctx *ParenthesizedRelationContext) {}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseDorisSQLParserListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseDorisSQLParserListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotAggregationExpression is called when production pivotAggregationExpression is entered.
func (s *BaseDorisSQLParserListener) EnterPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// ExitPivotAggregationExpression is called when production pivotAggregationExpression is exited.
func (s *BaseDorisSQLParserListener) ExitPivotAggregationExpression(ctx *PivotAggregationExpressionContext) {
}

// EnterPivotValue is called when production pivotValue is entered.
func (s *BaseDorisSQLParserListener) EnterPivotValue(ctx *PivotValueContext) {}

// ExitPivotValue is called when production pivotValue is exited.
func (s *BaseDorisSQLParserListener) ExitPivotValue(ctx *PivotValueContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseDorisSQLParserListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseDorisSQLParserListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseDorisSQLParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseDorisSQLParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterNamedArgumentList is called when production namedArgumentList is entered.
func (s *BaseDorisSQLParserListener) EnterNamedArgumentList(ctx *NamedArgumentListContext) {}

// ExitNamedArgumentList is called when production namedArgumentList is exited.
func (s *BaseDorisSQLParserListener) ExitNamedArgumentList(ctx *NamedArgumentListContext) {}

// EnterNamedArguments is called when production namedArguments is entered.
func (s *BaseDorisSQLParserListener) EnterNamedArguments(ctx *NamedArgumentsContext) {}

// ExitNamedArguments is called when production namedArguments is exited.
func (s *BaseDorisSQLParserListener) ExitNamedArguments(ctx *NamedArgumentsContext) {}

// EnterJoinRelation is called when production joinRelation is entered.
func (s *BaseDorisSQLParserListener) EnterJoinRelation(ctx *JoinRelationContext) {}

// ExitJoinRelation is called when production joinRelation is exited.
func (s *BaseDorisSQLParserListener) ExitJoinRelation(ctx *JoinRelationContext) {}

// EnterCrossOrInnerJoinType is called when production crossOrInnerJoinType is entered.
func (s *BaseDorisSQLParserListener) EnterCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// ExitCrossOrInnerJoinType is called when production crossOrInnerJoinType is exited.
func (s *BaseDorisSQLParserListener) ExitCrossOrInnerJoinType(ctx *CrossOrInnerJoinTypeContext) {}

// EnterOuterAndSemiJoinType is called when production outerAndSemiJoinType is entered.
func (s *BaseDorisSQLParserListener) EnterOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// ExitOuterAndSemiJoinType is called when production outerAndSemiJoinType is exited.
func (s *BaseDorisSQLParserListener) ExitOuterAndSemiJoinType(ctx *OuterAndSemiJoinTypeContext) {}

// EnterBracketHint is called when production bracketHint is entered.
func (s *BaseDorisSQLParserListener) EnterBracketHint(ctx *BracketHintContext) {}

// ExitBracketHint is called when production bracketHint is exited.
func (s *BaseDorisSQLParserListener) ExitBracketHint(ctx *BracketHintContext) {}

// EnterHintMap is called when production hintMap is entered.
func (s *BaseDorisSQLParserListener) EnterHintMap(ctx *HintMapContext) {}

// ExitHintMap is called when production hintMap is exited.
func (s *BaseDorisSQLParserListener) ExitHintMap(ctx *HintMapContext) {}

// EnterJoinCriteria is called when production joinCriteria is entered.
func (s *BaseDorisSQLParserListener) EnterJoinCriteria(ctx *JoinCriteriaContext) {}

// ExitJoinCriteria is called when production joinCriteria is exited.
func (s *BaseDorisSQLParserListener) ExitJoinCriteria(ctx *JoinCriteriaContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseDorisSQLParserListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseDorisSQLParserListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterPartitionNames is called when production partitionNames is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionNames(ctx *PartitionNamesContext) {}

// ExitPartitionNames is called when production partitionNames is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionNames(ctx *PartitionNamesContext) {}

// EnterKeyPartitionList is called when production keyPartitionList is entered.
func (s *BaseDorisSQLParserListener) EnterKeyPartitionList(ctx *KeyPartitionListContext) {}

// ExitKeyPartitionList is called when production keyPartitionList is exited.
func (s *BaseDorisSQLParserListener) ExitKeyPartitionList(ctx *KeyPartitionListContext) {}

// EnterTabletList is called when production tabletList is entered.
func (s *BaseDorisSQLParserListener) EnterTabletList(ctx *TabletListContext) {}

// ExitTabletList is called when production tabletList is exited.
func (s *BaseDorisSQLParserListener) ExitTabletList(ctx *TabletListContext) {}

// EnterPrepareStatement is called when production prepareStatement is entered.
func (s *BaseDorisSQLParserListener) EnterPrepareStatement(ctx *PrepareStatementContext) {}

// ExitPrepareStatement is called when production prepareStatement is exited.
func (s *BaseDorisSQLParserListener) ExitPrepareStatement(ctx *PrepareStatementContext) {}

// EnterPrepareSql is called when production prepareSql is entered.
func (s *BaseDorisSQLParserListener) EnterPrepareSql(ctx *PrepareSqlContext) {}

// ExitPrepareSql is called when production prepareSql is exited.
func (s *BaseDorisSQLParserListener) ExitPrepareSql(ctx *PrepareSqlContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseDorisSQLParserListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseDorisSQLParserListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterDeallocateStatement is called when production deallocateStatement is entered.
func (s *BaseDorisSQLParserListener) EnterDeallocateStatement(ctx *DeallocateStatementContext) {}

// ExitDeallocateStatement is called when production deallocateStatement is exited.
func (s *BaseDorisSQLParserListener) ExitDeallocateStatement(ctx *DeallocateStatementContext) {}

// EnterReplicaList is called when production replicaList is entered.
func (s *BaseDorisSQLParserListener) EnterReplicaList(ctx *ReplicaListContext) {}

// ExitReplicaList is called when production replicaList is exited.
func (s *BaseDorisSQLParserListener) ExitReplicaList(ctx *ReplicaListContext) {}

// EnterExpressionsWithDefault is called when production expressionsWithDefault is entered.
func (s *BaseDorisSQLParserListener) EnterExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {
}

// ExitExpressionsWithDefault is called when production expressionsWithDefault is exited.
func (s *BaseDorisSQLParserListener) ExitExpressionsWithDefault(ctx *ExpressionsWithDefaultContext) {}

// EnterExpressionOrDefault is called when production expressionOrDefault is entered.
func (s *BaseDorisSQLParserListener) EnterExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// ExitExpressionOrDefault is called when production expressionOrDefault is exited.
func (s *BaseDorisSQLParserListener) ExitExpressionOrDefault(ctx *ExpressionOrDefaultContext) {}

// EnterMapExpressionList is called when production mapExpressionList is entered.
func (s *BaseDorisSQLParserListener) EnterMapExpressionList(ctx *MapExpressionListContext) {}

// ExitMapExpressionList is called when production mapExpressionList is exited.
func (s *BaseDorisSQLParserListener) ExitMapExpressionList(ctx *MapExpressionListContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseDorisSQLParserListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseDorisSQLParserListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterExpressionSingleton is called when production expressionSingleton is entered.
func (s *BaseDorisSQLParserListener) EnterExpressionSingleton(ctx *ExpressionSingletonContext) {}

// ExitExpressionSingleton is called when production expressionSingleton is exited.
func (s *BaseDorisSQLParserListener) ExitExpressionSingleton(ctx *ExpressionSingletonContext) {}

// EnterExpressionDefault is called when production expressionDefault is entered.
func (s *BaseDorisSQLParserListener) EnterExpressionDefault(ctx *ExpressionDefaultContext) {}

// ExitExpressionDefault is called when production expressionDefault is exited.
func (s *BaseDorisSQLParserListener) ExitExpressionDefault(ctx *ExpressionDefaultContext) {}

// EnterLogicalNot is called when production logicalNot is entered.
func (s *BaseDorisSQLParserListener) EnterLogicalNot(ctx *LogicalNotContext) {}

// ExitLogicalNot is called when production logicalNot is exited.
func (s *BaseDorisSQLParserListener) ExitLogicalNot(ctx *LogicalNotContext) {}

// EnterLogicalBinary is called when production logicalBinary is entered.
func (s *BaseDorisSQLParserListener) EnterLogicalBinary(ctx *LogicalBinaryContext) {}

// ExitLogicalBinary is called when production logicalBinary is exited.
func (s *BaseDorisSQLParserListener) ExitLogicalBinary(ctx *LogicalBinaryContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseDorisSQLParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseDorisSQLParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseDorisSQLParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseDorisSQLParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterBooleanExpressionDefault is called when production booleanExpressionDefault is entered.
func (s *BaseDorisSQLParserListener) EnterBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// ExitBooleanExpressionDefault is called when production booleanExpressionDefault is exited.
func (s *BaseDorisSQLParserListener) ExitBooleanExpressionDefault(ctx *BooleanExpressionDefaultContext) {
}

// EnterIsNull is called when production isNull is entered.
func (s *BaseDorisSQLParserListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production isNull is exited.
func (s *BaseDorisSQLParserListener) ExitIsNull(ctx *IsNullContext) {}

// EnterScalarSubquery is called when production scalarSubquery is entered.
func (s *BaseDorisSQLParserListener) EnterScalarSubquery(ctx *ScalarSubqueryContext) {}

// ExitScalarSubquery is called when production scalarSubquery is exited.
func (s *BaseDorisSQLParserListener) ExitScalarSubquery(ctx *ScalarSubqueryContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseDorisSQLParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseDorisSQLParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterTupleInSubquery is called when production tupleInSubquery is entered.
func (s *BaseDorisSQLParserListener) EnterTupleInSubquery(ctx *TupleInSubqueryContext) {}

// ExitTupleInSubquery is called when production tupleInSubquery is exited.
func (s *BaseDorisSQLParserListener) ExitTupleInSubquery(ctx *TupleInSubqueryContext) {}

// EnterInSubquery is called when production inSubquery is entered.
func (s *BaseDorisSQLParserListener) EnterInSubquery(ctx *InSubqueryContext) {}

// ExitInSubquery is called when production inSubquery is exited.
func (s *BaseDorisSQLParserListener) ExitInSubquery(ctx *InSubqueryContext) {}

// EnterInList is called when production inList is entered.
func (s *BaseDorisSQLParserListener) EnterInList(ctx *InListContext) {}

// ExitInList is called when production inList is exited.
func (s *BaseDorisSQLParserListener) ExitInList(ctx *InListContext) {}

// EnterBetween is called when production between is entered.
func (s *BaseDorisSQLParserListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production between is exited.
func (s *BaseDorisSQLParserListener) ExitBetween(ctx *BetweenContext) {}

// EnterLike is called when production like is entered.
func (s *BaseDorisSQLParserListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production like is exited.
func (s *BaseDorisSQLParserListener) ExitLike(ctx *LikeContext) {}

// EnterValueExpressionDefault is called when production valueExpressionDefault is entered.
func (s *BaseDorisSQLParserListener) EnterValueExpressionDefault(ctx *ValueExpressionDefaultContext) {
}

// ExitValueExpressionDefault is called when production valueExpressionDefault is exited.
func (s *BaseDorisSQLParserListener) ExitValueExpressionDefault(ctx *ValueExpressionDefaultContext) {}

// EnterArithmeticBinary is called when production arithmeticBinary is entered.
func (s *BaseDorisSQLParserListener) EnterArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// ExitArithmeticBinary is called when production arithmeticBinary is exited.
func (s *BaseDorisSQLParserListener) ExitArithmeticBinary(ctx *ArithmeticBinaryContext) {}

// EnterDereference is called when production dereference is entered.
func (s *BaseDorisSQLParserListener) EnterDereference(ctx *DereferenceContext) {}

// ExitDereference is called when production dereference is exited.
func (s *BaseDorisSQLParserListener) ExitDereference(ctx *DereferenceContext) {}

// EnterOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is entered.
func (s *BaseDorisSQLParserListener) EnterOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// ExitOdbcFunctionCallExpression is called when production odbcFunctionCallExpression is exited.
func (s *BaseDorisSQLParserListener) ExitOdbcFunctionCallExpression(ctx *OdbcFunctionCallExpressionContext) {
}

// EnterMatchExpr is called when production matchExpr is entered.
func (s *BaseDorisSQLParserListener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production matchExpr is exited.
func (s *BaseDorisSQLParserListener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseDorisSQLParserListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseDorisSQLParserListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterConvert is called when production convert is entered.
func (s *BaseDorisSQLParserListener) EnterConvert(ctx *ConvertContext) {}

// ExitConvert is called when production convert is exited.
func (s *BaseDorisSQLParserListener) ExitConvert(ctx *ConvertContext) {}

// EnterCollectionSubscript is called when production collectionSubscript is entered.
func (s *BaseDorisSQLParserListener) EnterCollectionSubscript(ctx *CollectionSubscriptContext) {}

// ExitCollectionSubscript is called when production collectionSubscript is exited.
func (s *BaseDorisSQLParserListener) ExitCollectionSubscript(ctx *CollectionSubscriptContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDorisSQLParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDorisSQLParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCast is called when production cast is entered.
func (s *BaseDorisSQLParserListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BaseDorisSQLParserListener) ExitCast(ctx *CastContext) {}

// EnterParenthesizedExpression is called when production parenthesizedExpression is entered.
func (s *BaseDorisSQLParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production parenthesizedExpression is exited.
func (s *BaseDorisSQLParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterUserVariableExpression is called when production userVariableExpression is entered.
func (s *BaseDorisSQLParserListener) EnterUserVariableExpression(ctx *UserVariableExpressionContext) {
}

// ExitUserVariableExpression is called when production userVariableExpression is exited.
func (s *BaseDorisSQLParserListener) ExitUserVariableExpression(ctx *UserVariableExpressionContext) {}

// EnterFunctionCallExpression is called when production functionCallExpression is entered.
func (s *BaseDorisSQLParserListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {
}

// ExitFunctionCallExpression is called when production functionCallExpression is exited.
func (s *BaseDorisSQLParserListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterSimpleCase is called when production simpleCase is entered.
func (s *BaseDorisSQLParserListener) EnterSimpleCase(ctx *SimpleCaseContext) {}

// ExitSimpleCase is called when production simpleCase is exited.
func (s *BaseDorisSQLParserListener) ExitSimpleCase(ctx *SimpleCaseContext) {}

// EnterArrowExpression is called when production arrowExpression is entered.
func (s *BaseDorisSQLParserListener) EnterArrowExpression(ctx *ArrowExpressionContext) {}

// ExitArrowExpression is called when production arrowExpression is exited.
func (s *BaseDorisSQLParserListener) ExitArrowExpression(ctx *ArrowExpressionContext) {}

// EnterSystemVariableExpression is called when production systemVariableExpression is entered.
func (s *BaseDorisSQLParserListener) EnterSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// ExitSystemVariableExpression is called when production systemVariableExpression is exited.
func (s *BaseDorisSQLParserListener) ExitSystemVariableExpression(ctx *SystemVariableExpressionContext) {
}

// EnterConcat is called when production concat is entered.
func (s *BaseDorisSQLParserListener) EnterConcat(ctx *ConcatContext) {}

// ExitConcat is called when production concat is exited.
func (s *BaseDorisSQLParserListener) ExitConcat(ctx *ConcatContext) {}

// EnterSubqueryExpression is called when production subqueryExpression is entered.
func (s *BaseDorisSQLParserListener) EnterSubqueryExpression(ctx *SubqueryExpressionContext) {}

// ExitSubqueryExpression is called when production subqueryExpression is exited.
func (s *BaseDorisSQLParserListener) ExitSubqueryExpression(ctx *SubqueryExpressionContext) {}

// EnterLambdaFunctionExpr is called when production lambdaFunctionExpr is entered.
func (s *BaseDorisSQLParserListener) EnterLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// ExitLambdaFunctionExpr is called when production lambdaFunctionExpr is exited.
func (s *BaseDorisSQLParserListener) ExitLambdaFunctionExpr(ctx *LambdaFunctionExprContext) {}

// EnterDictionaryGetExpr is called when production dictionaryGetExpr is entered.
func (s *BaseDorisSQLParserListener) EnterDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// ExitDictionaryGetExpr is called when production dictionaryGetExpr is exited.
func (s *BaseDorisSQLParserListener) ExitDictionaryGetExpr(ctx *DictionaryGetExprContext) {}

// EnterCollate is called when production collate is entered.
func (s *BaseDorisSQLParserListener) EnterCollate(ctx *CollateContext) {}

// ExitCollate is called when production collate is exited.
func (s *BaseDorisSQLParserListener) ExitCollate(ctx *CollateContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BaseDorisSQLParserListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BaseDorisSQLParserListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterMapConstructor is called when production mapConstructor is entered.
func (s *BaseDorisSQLParserListener) EnterMapConstructor(ctx *MapConstructorContext) {}

// ExitMapConstructor is called when production mapConstructor is exited.
func (s *BaseDorisSQLParserListener) ExitMapConstructor(ctx *MapConstructorContext) {}

// EnterArraySlice is called when production arraySlice is entered.
func (s *BaseDorisSQLParserListener) EnterArraySlice(ctx *ArraySliceContext) {}

// ExitArraySlice is called when production arraySlice is exited.
func (s *BaseDorisSQLParserListener) ExitArraySlice(ctx *ArraySliceContext) {}

// EnterExists is called when production exists is entered.
func (s *BaseDorisSQLParserListener) EnterExists(ctx *ExistsContext) {}

// ExitExists is called when production exists is exited.
func (s *BaseDorisSQLParserListener) ExitExists(ctx *ExistsContext) {}

// EnterSearchedCase is called when production searchedCase is entered.
func (s *BaseDorisSQLParserListener) EnterSearchedCase(ctx *SearchedCaseContext) {}

// ExitSearchedCase is called when production searchedCase is exited.
func (s *BaseDorisSQLParserListener) ExitSearchedCase(ctx *SearchedCaseContext) {}

// EnterArithmeticUnary is called when production arithmeticUnary is entered.
func (s *BaseDorisSQLParserListener) EnterArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// ExitArithmeticUnary is called when production arithmeticUnary is exited.
func (s *BaseDorisSQLParserListener) ExitArithmeticUnary(ctx *ArithmeticUnaryContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitNullLiteral(ctx *NullLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterUnitBoundaryLiteral is called when production unitBoundaryLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// ExitUnitBoundaryLiteral is called when production unitBoundaryLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitUnitBoundaryLiteral(ctx *UnitBoundaryLiteralContext) {}

// EnterBinaryLiteral is called when production binaryLiteral is entered.
func (s *BaseDorisSQLParserListener) EnterBinaryLiteral(ctx *BinaryLiteralContext) {}

// ExitBinaryLiteral is called when production binaryLiteral is exited.
func (s *BaseDorisSQLParserListener) ExitBinaryLiteral(ctx *BinaryLiteralContext) {}

// EnterParameter is called when production Parameter is entered.
func (s *BaseDorisSQLParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production Parameter is exited.
func (s *BaseDorisSQLParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterExtract is called when production extract is entered.
func (s *BaseDorisSQLParserListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BaseDorisSQLParserListener) ExitExtract(ctx *ExtractContext) {}

// EnterGroupingOperation is called when production groupingOperation is entered.
func (s *BaseDorisSQLParserListener) EnterGroupingOperation(ctx *GroupingOperationContext) {}

// ExitGroupingOperation is called when production groupingOperation is exited.
func (s *BaseDorisSQLParserListener) ExitGroupingOperation(ctx *GroupingOperationContext) {}

// EnterInformationFunction is called when production informationFunction is entered.
func (s *BaseDorisSQLParserListener) EnterInformationFunction(ctx *InformationFunctionContext) {}

// ExitInformationFunction is called when production informationFunction is exited.
func (s *BaseDorisSQLParserListener) ExitInformationFunction(ctx *InformationFunctionContext) {}

// EnterSpecialDateTime is called when production specialDateTime is entered.
func (s *BaseDorisSQLParserListener) EnterSpecialDateTime(ctx *SpecialDateTimeContext) {}

// ExitSpecialDateTime is called when production specialDateTime is exited.
func (s *BaseDorisSQLParserListener) ExitSpecialDateTime(ctx *SpecialDateTimeContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseDorisSQLParserListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseDorisSQLParserListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterAggregationFunctionCall is called when production aggregationFunctionCall is entered.
func (s *BaseDorisSQLParserListener) EnterAggregationFunctionCall(ctx *AggregationFunctionCallContext) {
}

// ExitAggregationFunctionCall is called when production aggregationFunctionCall is exited.
func (s *BaseDorisSQLParserListener) ExitAggregationFunctionCall(ctx *AggregationFunctionCallContext) {
}

// EnterWindowFunctionCall is called when production windowFunctionCall is entered.
func (s *BaseDorisSQLParserListener) EnterWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// ExitWindowFunctionCall is called when production windowFunctionCall is exited.
func (s *BaseDorisSQLParserListener) ExitWindowFunctionCall(ctx *WindowFunctionCallContext) {}

// EnterTranslateFunctionCall is called when production translateFunctionCall is entered.
func (s *BaseDorisSQLParserListener) EnterTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// ExitTranslateFunctionCall is called when production translateFunctionCall is exited.
func (s *BaseDorisSQLParserListener) ExitTranslateFunctionCall(ctx *TranslateFunctionCallContext) {}

// EnterSimpleFunctionCall is called when production simpleFunctionCall is entered.
func (s *BaseDorisSQLParserListener) EnterSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// ExitSimpleFunctionCall is called when production simpleFunctionCall is exited.
func (s *BaseDorisSQLParserListener) ExitSimpleFunctionCall(ctx *SimpleFunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseDorisSQLParserListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseDorisSQLParserListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseDorisSQLParserListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseDorisSQLParserListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseDorisSQLParserListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseDorisSQLParserListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterColumnReference is called when production columnReference is entered.
func (s *BaseDorisSQLParserListener) EnterColumnReference(ctx *ColumnReferenceContext) {}

// ExitColumnReference is called when production columnReference is exited.
func (s *BaseDorisSQLParserListener) ExitColumnReference(ctx *ColumnReferenceContext) {}

// EnterInformationFunctionExpression is called when production informationFunctionExpression is entered.
func (s *BaseDorisSQLParserListener) EnterInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// ExitInformationFunctionExpression is called when production informationFunctionExpression is exited.
func (s *BaseDorisSQLParserListener) ExitInformationFunctionExpression(ctx *InformationFunctionExpressionContext) {
}

// EnterSpecialDateTimeExpression is called when production specialDateTimeExpression is entered.
func (s *BaseDorisSQLParserListener) EnterSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// ExitSpecialDateTimeExpression is called when production specialDateTimeExpression is exited.
func (s *BaseDorisSQLParserListener) ExitSpecialDateTimeExpression(ctx *SpecialDateTimeExpressionContext) {
}

// EnterSpecialFunctionExpression is called when production specialFunctionExpression is entered.
func (s *BaseDorisSQLParserListener) EnterSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// ExitSpecialFunctionExpression is called when production specialFunctionExpression is exited.
func (s *BaseDorisSQLParserListener) ExitSpecialFunctionExpression(ctx *SpecialFunctionExpressionContext) {
}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseDorisSQLParserListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseDorisSQLParserListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWhenClause is called when production whenClause is entered.
func (s *BaseDorisSQLParserListener) EnterWhenClause(ctx *WhenClauseContext) {}

// ExitWhenClause is called when production whenClause is exited.
func (s *BaseDorisSQLParserListener) ExitWhenClause(ctx *WhenClauseContext) {}

// EnterOver is called when production over is entered.
func (s *BaseDorisSQLParserListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BaseDorisSQLParserListener) ExitOver(ctx *OverContext) {}

// EnterIgnoreNulls is called when production ignoreNulls is entered.
func (s *BaseDorisSQLParserListener) EnterIgnoreNulls(ctx *IgnoreNullsContext) {}

// ExitIgnoreNulls is called when production ignoreNulls is exited.
func (s *BaseDorisSQLParserListener) ExitIgnoreNulls(ctx *IgnoreNullsContext) {}

// EnterWindowFrame is called when production windowFrame is entered.
func (s *BaseDorisSQLParserListener) EnterWindowFrame(ctx *WindowFrameContext) {}

// ExitWindowFrame is called when production windowFrame is exited.
func (s *BaseDorisSQLParserListener) ExitWindowFrame(ctx *WindowFrameContext) {}

// EnterUnboundedFrame is called when production unboundedFrame is entered.
func (s *BaseDorisSQLParserListener) EnterUnboundedFrame(ctx *UnboundedFrameContext) {}

// ExitUnboundedFrame is called when production unboundedFrame is exited.
func (s *BaseDorisSQLParserListener) ExitUnboundedFrame(ctx *UnboundedFrameContext) {}

// EnterCurrentRowBound is called when production currentRowBound is entered.
func (s *BaseDorisSQLParserListener) EnterCurrentRowBound(ctx *CurrentRowBoundContext) {}

// ExitCurrentRowBound is called when production currentRowBound is exited.
func (s *BaseDorisSQLParserListener) ExitCurrentRowBound(ctx *CurrentRowBoundContext) {}

// EnterBoundedFrame is called when production boundedFrame is entered.
func (s *BaseDorisSQLParserListener) EnterBoundedFrame(ctx *BoundedFrameContext) {}

// ExitBoundedFrame is called when production boundedFrame is exited.
func (s *BaseDorisSQLParserListener) ExitBoundedFrame(ctx *BoundedFrameContext) {}

// EnterBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is entered.
func (s *BaseDorisSQLParserListener) EnterBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {
}

// ExitBackupRestoreObjectDesc is called when production backupRestoreObjectDesc is exited.
func (s *BaseDorisSQLParserListener) ExitBackupRestoreObjectDesc(ctx *BackupRestoreObjectDescContext) {
}

// EnterTableDesc is called when production tableDesc is entered.
func (s *BaseDorisSQLParserListener) EnterTableDesc(ctx *TableDescContext) {}

// ExitTableDesc is called when production tableDesc is exited.
func (s *BaseDorisSQLParserListener) ExitTableDesc(ctx *TableDescContext) {}

// EnterBackupRestoreTableDesc is called when production backupRestoreTableDesc is entered.
func (s *BaseDorisSQLParserListener) EnterBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {
}

// ExitBackupRestoreTableDesc is called when production backupRestoreTableDesc is exited.
func (s *BaseDorisSQLParserListener) ExitBackupRestoreTableDesc(ctx *BackupRestoreTableDescContext) {}

// EnterExplainDesc is called when production explainDesc is entered.
func (s *BaseDorisSQLParserListener) EnterExplainDesc(ctx *ExplainDescContext) {}

// ExitExplainDesc is called when production explainDesc is exited.
func (s *BaseDorisSQLParserListener) ExitExplainDesc(ctx *ExplainDescContext) {}

// EnterOptimizerTrace is called when production optimizerTrace is entered.
func (s *BaseDorisSQLParserListener) EnterOptimizerTrace(ctx *OptimizerTraceContext) {}

// ExitOptimizerTrace is called when production optimizerTrace is exited.
func (s *BaseDorisSQLParserListener) ExitOptimizerTrace(ctx *OptimizerTraceContext) {}

// EnterPartitionExpr is called when production partitionExpr is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionExpr(ctx *PartitionExprContext) {}

// ExitPartitionExpr is called when production partitionExpr is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionExpr(ctx *PartitionExprContext) {}

// EnterPartitionDesc is called when production partitionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionDesc(ctx *PartitionDescContext) {}

// ExitPartitionDesc is called when production partitionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionDesc(ctx *PartitionDescContext) {}

// EnterListPartitionDesc is called when production listPartitionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterListPartitionDesc(ctx *ListPartitionDescContext) {}

// ExitListPartitionDesc is called when production listPartitionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitListPartitionDesc(ctx *ListPartitionDescContext) {}

// EnterSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// ExitSingleItemListPartitionDesc is called when production singleItemListPartitionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitSingleItemListPartitionDesc(ctx *SingleItemListPartitionDescContext) {
}

// EnterMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// ExitMultiItemListPartitionDesc is called when production multiItemListPartitionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitMultiItemListPartitionDesc(ctx *MultiItemListPartitionDescContext) {
}

// EnterMultiListPartitionValues is called when production multiListPartitionValues is entered.
func (s *BaseDorisSQLParserListener) EnterMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// ExitMultiListPartitionValues is called when production multiListPartitionValues is exited.
func (s *BaseDorisSQLParserListener) ExitMultiListPartitionValues(ctx *MultiListPartitionValuesContext) {
}

// EnterSingleListPartitionValues is called when production singleListPartitionValues is entered.
func (s *BaseDorisSQLParserListener) EnterSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// ExitSingleListPartitionValues is called when production singleListPartitionValues is exited.
func (s *BaseDorisSQLParserListener) ExitSingleListPartitionValues(ctx *SingleListPartitionValuesContext) {
}

// EnterListPartitionValues is called when production listPartitionValues is entered.
func (s *BaseDorisSQLParserListener) EnterListPartitionValues(ctx *ListPartitionValuesContext) {}

// ExitListPartitionValues is called when production listPartitionValues is exited.
func (s *BaseDorisSQLParserListener) ExitListPartitionValues(ctx *ListPartitionValuesContext) {}

// EnterListPartitionValue is called when production listPartitionValue is entered.
func (s *BaseDorisSQLParserListener) EnterListPartitionValue(ctx *ListPartitionValueContext) {}

// ExitListPartitionValue is called when production listPartitionValue is exited.
func (s *BaseDorisSQLParserListener) ExitListPartitionValue(ctx *ListPartitionValueContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseDorisSQLParserListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseDorisSQLParserListener) ExitStringList(ctx *StringListContext) {}

// EnterLiteralExpressionList is called when production literalExpressionList is entered.
func (s *BaseDorisSQLParserListener) EnterLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// ExitLiteralExpressionList is called when production literalExpressionList is exited.
func (s *BaseDorisSQLParserListener) ExitLiteralExpressionList(ctx *LiteralExpressionListContext) {}

// EnterRangePartitionDesc is called when production rangePartitionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterRangePartitionDesc(ctx *RangePartitionDescContext) {}

// ExitRangePartitionDesc is called when production rangePartitionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitRangePartitionDesc(ctx *RangePartitionDescContext) {}

// EnterSingleRangePartition is called when production singleRangePartition is entered.
func (s *BaseDorisSQLParserListener) EnterSingleRangePartition(ctx *SingleRangePartitionContext) {}

// ExitSingleRangePartition is called when production singleRangePartition is exited.
func (s *BaseDorisSQLParserListener) ExitSingleRangePartition(ctx *SingleRangePartitionContext) {}

// EnterMultiRangePartition is called when production multiRangePartition is entered.
func (s *BaseDorisSQLParserListener) EnterMultiRangePartition(ctx *MultiRangePartitionContext) {}

// ExitMultiRangePartition is called when production multiRangePartition is exited.
func (s *BaseDorisSQLParserListener) ExitMultiRangePartition(ctx *MultiRangePartitionContext) {}

// EnterPartitionRangeDesc is called when production partitionRangeDesc is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// ExitPartitionRangeDesc is called when production partitionRangeDesc is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionRangeDesc(ctx *PartitionRangeDescContext) {}

// EnterPartitionKeyDesc is called when production partitionKeyDesc is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// ExitPartitionKeyDesc is called when production partitionKeyDesc is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionKeyDesc(ctx *PartitionKeyDescContext) {}

// EnterPartitionValueList is called when production partitionValueList is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionValueList(ctx *PartitionValueListContext) {}

// ExitPartitionValueList is called when production partitionValueList is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionValueList(ctx *PartitionValueListContext) {}

// EnterKeyPartition is called when production keyPartition is entered.
func (s *BaseDorisSQLParserListener) EnterKeyPartition(ctx *KeyPartitionContext) {}

// ExitKeyPartition is called when production keyPartition is exited.
func (s *BaseDorisSQLParserListener) ExitKeyPartition(ctx *KeyPartitionContext) {}

// EnterPartitionValue is called when production partitionValue is entered.
func (s *BaseDorisSQLParserListener) EnterPartitionValue(ctx *PartitionValueContext) {}

// ExitPartitionValue is called when production partitionValue is exited.
func (s *BaseDorisSQLParserListener) ExitPartitionValue(ctx *PartitionValueContext) {}

// EnterDistributionClause is called when production distributionClause is entered.
func (s *BaseDorisSQLParserListener) EnterDistributionClause(ctx *DistributionClauseContext) {}

// ExitDistributionClause is called when production distributionClause is exited.
func (s *BaseDorisSQLParserListener) ExitDistributionClause(ctx *DistributionClauseContext) {}

// EnterDistributionDesc is called when production distributionDesc is entered.
func (s *BaseDorisSQLParserListener) EnterDistributionDesc(ctx *DistributionDescContext) {}

// ExitDistributionDesc is called when production distributionDesc is exited.
func (s *BaseDorisSQLParserListener) ExitDistributionDesc(ctx *DistributionDescContext) {}

// EnterRefreshSchemeDesc is called when production refreshSchemeDesc is entered.
func (s *BaseDorisSQLParserListener) EnterRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// ExitRefreshSchemeDesc is called when production refreshSchemeDesc is exited.
func (s *BaseDorisSQLParserListener) ExitRefreshSchemeDesc(ctx *RefreshSchemeDescContext) {}

// EnterStatusDesc is called when production statusDesc is entered.
func (s *BaseDorisSQLParserListener) EnterStatusDesc(ctx *StatusDescContext) {}

// ExitStatusDesc is called when production statusDesc is exited.
func (s *BaseDorisSQLParserListener) ExitStatusDesc(ctx *StatusDescContext) {}

// EnterProperties is called when production properties is entered.
func (s *BaseDorisSQLParserListener) EnterProperties(ctx *PropertiesContext) {}

// ExitProperties is called when production properties is exited.
func (s *BaseDorisSQLParserListener) ExitProperties(ctx *PropertiesContext) {}

// EnterExtProperties is called when production extProperties is entered.
func (s *BaseDorisSQLParserListener) EnterExtProperties(ctx *ExtPropertiesContext) {}

// ExitExtProperties is called when production extProperties is exited.
func (s *BaseDorisSQLParserListener) ExitExtProperties(ctx *ExtPropertiesContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseDorisSQLParserListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseDorisSQLParserListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterUserPropertyList is called when production userPropertyList is entered.
func (s *BaseDorisSQLParserListener) EnterUserPropertyList(ctx *UserPropertyListContext) {}

// ExitUserPropertyList is called when production userPropertyList is exited.
func (s *BaseDorisSQLParserListener) ExitUserPropertyList(ctx *UserPropertyListContext) {}

// EnterProperty is called when production property is entered.
func (s *BaseDorisSQLParserListener) EnterProperty(ctx *PropertyContext) {}

// ExitProperty is called when production property is exited.
func (s *BaseDorisSQLParserListener) ExitProperty(ctx *PropertyContext) {}

// EnterInlineProperties is called when production inlineProperties is entered.
func (s *BaseDorisSQLParserListener) EnterInlineProperties(ctx *InlinePropertiesContext) {}

// ExitInlineProperties is called when production inlineProperties is exited.
func (s *BaseDorisSQLParserListener) ExitInlineProperties(ctx *InlinePropertiesContext) {}

// EnterInlineProperty is called when production inlineProperty is entered.
func (s *BaseDorisSQLParserListener) EnterInlineProperty(ctx *InlinePropertyContext) {}

// ExitInlineProperty is called when production inlineProperty is exited.
func (s *BaseDorisSQLParserListener) ExitInlineProperty(ctx *InlinePropertyContext) {}

// EnterVarType is called when production varType is entered.
func (s *BaseDorisSQLParserListener) EnterVarType(ctx *VarTypeContext) {}

// ExitVarType is called when production varType is exited.
func (s *BaseDorisSQLParserListener) ExitVarType(ctx *VarTypeContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseDorisSQLParserListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseDorisSQLParserListener) ExitComment(ctx *CommentContext) {}

// EnterOutfile is called when production outfile is entered.
func (s *BaseDorisSQLParserListener) EnterOutfile(ctx *OutfileContext) {}

// ExitOutfile is called when production outfile is exited.
func (s *BaseDorisSQLParserListener) ExitOutfile(ctx *OutfileContext) {}

// EnterFileFormat is called when production fileFormat is entered.
func (s *BaseDorisSQLParserListener) EnterFileFormat(ctx *FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *BaseDorisSQLParserListener) ExitFileFormat(ctx *FileFormatContext) {}

// EnterString is called when production string is entered.
func (s *BaseDorisSQLParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseDorisSQLParserListener) ExitString(ctx *StringContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseDorisSQLParserListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseDorisSQLParserListener) ExitBinary(ctx *BinaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseDorisSQLParserListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseDorisSQLParserListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseDorisSQLParserListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseDorisSQLParserListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseDorisSQLParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseDorisSQLParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterTaskInterval is called when production taskInterval is entered.
func (s *BaseDorisSQLParserListener) EnterTaskInterval(ctx *TaskIntervalContext) {}

// ExitTaskInterval is called when production taskInterval is exited.
func (s *BaseDorisSQLParserListener) ExitTaskInterval(ctx *TaskIntervalContext) {}

// EnterTaskUnitIdentifier is called when production taskUnitIdentifier is entered.
func (s *BaseDorisSQLParserListener) EnterTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// ExitTaskUnitIdentifier is called when production taskUnitIdentifier is exited.
func (s *BaseDorisSQLParserListener) ExitTaskUnitIdentifier(ctx *TaskUnitIdentifierContext) {}

// EnterUnitIdentifier is called when production unitIdentifier is entered.
func (s *BaseDorisSQLParserListener) EnterUnitIdentifier(ctx *UnitIdentifierContext) {}

// ExitUnitIdentifier is called when production unitIdentifier is exited.
func (s *BaseDorisSQLParserListener) ExitUnitIdentifier(ctx *UnitIdentifierContext) {}

// EnterUnitBoundary is called when production unitBoundary is entered.
func (s *BaseDorisSQLParserListener) EnterUnitBoundary(ctx *UnitBoundaryContext) {}

// ExitUnitBoundary is called when production unitBoundary is exited.
func (s *BaseDorisSQLParserListener) ExitUnitBoundary(ctx *UnitBoundaryContext) {}

// EnterType is called when production type is entered.
func (s *BaseDorisSQLParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDorisSQLParserListener) ExitType(ctx *TypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseDorisSQLParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseDorisSQLParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseDorisSQLParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseDorisSQLParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterSubfieldDesc is called when production subfieldDesc is entered.
func (s *BaseDorisSQLParserListener) EnterSubfieldDesc(ctx *SubfieldDescContext) {}

// ExitSubfieldDesc is called when production subfieldDesc is exited.
func (s *BaseDorisSQLParserListener) ExitSubfieldDesc(ctx *SubfieldDescContext) {}

// EnterSubfieldDescs is called when production subfieldDescs is entered.
func (s *BaseDorisSQLParserListener) EnterSubfieldDescs(ctx *SubfieldDescsContext) {}

// ExitSubfieldDescs is called when production subfieldDescs is exited.
func (s *BaseDorisSQLParserListener) ExitSubfieldDescs(ctx *SubfieldDescsContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseDorisSQLParserListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseDorisSQLParserListener) ExitStructType(ctx *StructTypeContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseDorisSQLParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseDorisSQLParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseDorisSQLParserListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseDorisSQLParserListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterDecimalType is called when production decimalType is entered.
func (s *BaseDorisSQLParserListener) EnterDecimalType(ctx *DecimalTypeContext) {}

// ExitDecimalType is called when production decimalType is exited.
func (s *BaseDorisSQLParserListener) ExitDecimalType(ctx *DecimalTypeContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseDorisSQLParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseDorisSQLParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseDorisSQLParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseDorisSQLParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterWriteBranch is called when production writeBranch is entered.
func (s *BaseDorisSQLParserListener) EnterWriteBranch(ctx *WriteBranchContext) {}

// ExitWriteBranch is called when production writeBranch is exited.
func (s *BaseDorisSQLParserListener) ExitWriteBranch(ctx *WriteBranchContext) {}

// EnterUnquotedIdentifier is called when production unquotedIdentifier is entered.
func (s *BaseDorisSQLParserListener) EnterUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// ExitUnquotedIdentifier is called when production unquotedIdentifier is exited.
func (s *BaseDorisSQLParserListener) ExitUnquotedIdentifier(ctx *UnquotedIdentifierContext) {}

// EnterDigitIdentifier is called when production digitIdentifier is entered.
func (s *BaseDorisSQLParserListener) EnterDigitIdentifier(ctx *DigitIdentifierContext) {}

// ExitDigitIdentifier is called when production digitIdentifier is exited.
func (s *BaseDorisSQLParserListener) ExitDigitIdentifier(ctx *DigitIdentifierContext) {}

// EnterBackQuotedIdentifier is called when production backQuotedIdentifier is entered.
func (s *BaseDorisSQLParserListener) EnterBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// ExitBackQuotedIdentifier is called when production backQuotedIdentifier is exited.
func (s *BaseDorisSQLParserListener) ExitBackQuotedIdentifier(ctx *BackQuotedIdentifierContext) {}

// EnterIdentifierWithAlias is called when production identifierWithAlias is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// ExitIdentifierWithAlias is called when production identifierWithAlias is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierWithAlias(ctx *IdentifierWithAliasContext) {}

// EnterIdentifierWithAliasList is called when production identifierWithAliasList is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {
}

// ExitIdentifierWithAliasList is called when production identifierWithAliasList is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierWithAliasList(ctx *IdentifierWithAliasListContext) {
}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterIdentifierOrString is called when production identifierOrString is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierOrString(ctx *IdentifierOrStringContext) {}

// ExitIdentifierOrString is called when production identifierOrString is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierOrString(ctx *IdentifierOrStringContext) {}

// EnterIdentifierOrStringList is called when production identifierOrStringList is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierOrStringList(ctx *IdentifierOrStringListContext) {
}

// ExitIdentifierOrStringList is called when production identifierOrStringList is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierOrStringList(ctx *IdentifierOrStringListContext) {}

// EnterIdentifierOrStringOrStar is called when production identifierOrStringOrStar is entered.
func (s *BaseDorisSQLParserListener) EnterIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// ExitIdentifierOrStringOrStar is called when production identifierOrStringOrStar is exited.
func (s *BaseDorisSQLParserListener) ExitIdentifierOrStringOrStar(ctx *IdentifierOrStringOrStarContext) {
}

// EnterUserWithoutHost is called when production userWithoutHost is entered.
func (s *BaseDorisSQLParserListener) EnterUserWithoutHost(ctx *UserWithoutHostContext) {}

// ExitUserWithoutHost is called when production userWithoutHost is exited.
func (s *BaseDorisSQLParserListener) ExitUserWithoutHost(ctx *UserWithoutHostContext) {}

// EnterUserWithHost is called when production userWithHost is entered.
func (s *BaseDorisSQLParserListener) EnterUserWithHost(ctx *UserWithHostContext) {}

// ExitUserWithHost is called when production userWithHost is exited.
func (s *BaseDorisSQLParserListener) ExitUserWithHost(ctx *UserWithHostContext) {}

// EnterUserWithHostAndBlanket is called when production userWithHostAndBlanket is entered.
func (s *BaseDorisSQLParserListener) EnterUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {
}

// ExitUserWithHostAndBlanket is called when production userWithHostAndBlanket is exited.
func (s *BaseDorisSQLParserListener) ExitUserWithHostAndBlanket(ctx *UserWithHostAndBlanketContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseDorisSQLParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseDorisSQLParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentList is called when production assignmentList is entered.
func (s *BaseDorisSQLParserListener) EnterAssignmentList(ctx *AssignmentListContext) {}

// ExitAssignmentList is called when production assignmentList is exited.
func (s *BaseDorisSQLParserListener) ExitAssignmentList(ctx *AssignmentListContext) {}

// EnterDecimalValue is called when production decimalValue is entered.
func (s *BaseDorisSQLParserListener) EnterDecimalValue(ctx *DecimalValueContext) {}

// ExitDecimalValue is called when production decimalValue is exited.
func (s *BaseDorisSQLParserListener) ExitDecimalValue(ctx *DecimalValueContext) {}

// EnterDoubleValue is called when production doubleValue is entered.
func (s *BaseDorisSQLParserListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production doubleValue is exited.
func (s *BaseDorisSQLParserListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseDorisSQLParserListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseDorisSQLParserListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseDorisSQLParserListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseDorisSQLParserListener) ExitNonReserved(ctx *NonReservedContext) {}

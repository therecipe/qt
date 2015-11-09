#ifdef __cplusplus
extern "C" {
#endif

int QTextDocument_BlockCount(void* ptr);
char* QTextDocument_DefaultStyleSheet(void* ptr);
double QTextDocument_DocumentMargin(void* ptr);
double QTextDocument_IndentWidth(void* ptr);
int QTextDocument_IsModified(void* ptr);
int QTextDocument_IsUndoRedoEnabled(void* ptr);
void QTextDocument_MarkContentsDirty(void* ptr, int position, int length);
int QTextDocument_MaximumBlockCount(void* ptr);
void QTextDocument_SetBaseUrl(void* ptr, void* url);
void QTextDocument_SetDefaultStyleSheet(void* ptr, char* sheet);
void QTextDocument_SetDocumentMargin(void* ptr, double margin);
void QTextDocument_SetMaximumBlockCount(void* ptr, int maximum);
void QTextDocument_SetModified(void* ptr, int m);
void QTextDocument_SetPageSize(void* ptr, void* size);
void QTextDocument_SetTextWidth(void* ptr, double width);
void QTextDocument_SetUndoRedoEnabled(void* ptr, int enable);
void QTextDocument_SetUseDesignMetrics(void* ptr, int b);
double QTextDocument_TextWidth(void* ptr);
int QTextDocument_UseDesignMetrics(void* ptr);
void* QTextDocument_NewQTextDocument(void* parent);
void* QTextDocument_NewQTextDocument2(char* text, void* parent);
void QTextDocument_AddResource(void* ptr, int ty, void* name, void* resource);
void QTextDocument_AdjustSize(void* ptr);
int QTextDocument_AvailableRedoSteps(void* ptr);
int QTextDocument_AvailableUndoSteps(void* ptr);
void QTextDocument_ConnectBlockCountChanged(void* ptr);
void QTextDocument_DisconnectBlockCountChanged(void* ptr);
int QTextDocument_CharacterCount(void* ptr);
void QTextDocument_Clear(void* ptr);
void QTextDocument_ClearUndoRedoStacks(void* ptr, int stacksToClear);
void* QTextDocument_Clone(void* ptr, void* parent);
void QTextDocument_ConnectContentsChange(void* ptr);
void QTextDocument_DisconnectContentsChange(void* ptr);
void QTextDocument_ConnectContentsChanged(void* ptr);
void QTextDocument_DisconnectContentsChanged(void* ptr);
int QTextDocument_DefaultCursorMoveStyle(void* ptr);
void* QTextDocument_DocumentLayout(void* ptr);
void QTextDocument_ConnectDocumentLayoutChanged(void* ptr);
void QTextDocument_DisconnectDocumentLayoutChanged(void* ptr);
void QTextDocument_DrawContents(void* ptr, void* p, void* rect);
double QTextDocument_IdealWidth(void* ptr);
int QTextDocument_IsEmpty(void* ptr);
int QTextDocument_IsRedoAvailable(void* ptr);
int QTextDocument_IsUndoAvailable(void* ptr);
int QTextDocument_LineCount(void* ptr);
char* QTextDocument_MetaInformation(void* ptr, int info);
void QTextDocument_ConnectModificationChanged(void* ptr);
void QTextDocument_DisconnectModificationChanged(void* ptr);
void* QTextDocument_Object(void* ptr, int objectIndex);
void* QTextDocument_ObjectForFormat(void* ptr, void* f);
int QTextDocument_PageCount(void* ptr);
void QTextDocument_Print(void* ptr, void* printer);
void QTextDocument_Redo2(void* ptr);
void QTextDocument_Redo(void* ptr, void* cursor);
void QTextDocument_ConnectRedoAvailable(void* ptr);
void QTextDocument_DisconnectRedoAvailable(void* ptr);
void* QTextDocument_Resource(void* ptr, int ty, void* name);
int QTextDocument_Revision(void* ptr);
void* QTextDocument_RootFrame(void* ptr);
void QTextDocument_SetDefaultCursorMoveStyle(void* ptr, int style);
void QTextDocument_SetDefaultFont(void* ptr, void* font);
void QTextDocument_SetDefaultTextOption(void* ptr, void* option);
void QTextDocument_SetDocumentLayout(void* ptr, void* layout);
void QTextDocument_SetHtml(void* ptr, char* html);
void QTextDocument_SetIndentWidth(void* ptr, double width);
void QTextDocument_SetMetaInformation(void* ptr, int info, char* stri);
void QTextDocument_SetPlainText(void* ptr, char* text);
char* QTextDocument_ToHtml(void* ptr, void* encoding);
char* QTextDocument_ToPlainText(void* ptr);
void QTextDocument_Undo2(void* ptr);
void QTextDocument_Undo(void* ptr, void* cursor);
void QTextDocument_ConnectUndoAvailable(void* ptr);
void QTextDocument_DisconnectUndoAvailable(void* ptr);
void QTextDocument_ConnectUndoCommandAdded(void* ptr);
void QTextDocument_DisconnectUndoCommandAdded(void* ptr);
void QTextDocument_DestroyQTextDocument(void* ptr);

#ifdef __cplusplus
}
#endif
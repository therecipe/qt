#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QTextDocument_BaseUrl(QtObjectPtr ptr);
int QTextDocument_BlockCount(QtObjectPtr ptr);
char* QTextDocument_DefaultStyleSheet(QtObjectPtr ptr);
int QTextDocument_IsModified(QtObjectPtr ptr);
int QTextDocument_IsUndoRedoEnabled(QtObjectPtr ptr);
void QTextDocument_MarkContentsDirty(QtObjectPtr ptr, int position, int length);
int QTextDocument_MaximumBlockCount(QtObjectPtr ptr);
void QTextDocument_SetBaseUrl(QtObjectPtr ptr, char* url);
void QTextDocument_SetDefaultStyleSheet(QtObjectPtr ptr, char* sheet);
void QTextDocument_SetMaximumBlockCount(QtObjectPtr ptr, int maximum);
void QTextDocument_SetModified(QtObjectPtr ptr, int m);
void QTextDocument_SetPageSize(QtObjectPtr ptr, QtObjectPtr size);
void QTextDocument_SetUndoRedoEnabled(QtObjectPtr ptr, int enable);
void QTextDocument_SetUseDesignMetrics(QtObjectPtr ptr, int b);
int QTextDocument_UseDesignMetrics(QtObjectPtr ptr);
QtObjectPtr QTextDocument_NewQTextDocument(QtObjectPtr parent);
QtObjectPtr QTextDocument_NewQTextDocument2(char* text, QtObjectPtr parent);
void QTextDocument_AddResource(QtObjectPtr ptr, int ty, char* name, char* resource);
void QTextDocument_AdjustSize(QtObjectPtr ptr);
int QTextDocument_AvailableRedoSteps(QtObjectPtr ptr);
int QTextDocument_AvailableUndoSteps(QtObjectPtr ptr);
void QTextDocument_ConnectBaseUrlChanged(QtObjectPtr ptr);
void QTextDocument_DisconnectBaseUrlChanged(QtObjectPtr ptr);
void QTextDocument_ConnectBlockCountChanged(QtObjectPtr ptr);
void QTextDocument_DisconnectBlockCountChanged(QtObjectPtr ptr);
int QTextDocument_CharacterCount(QtObjectPtr ptr);
void QTextDocument_Clear(QtObjectPtr ptr);
void QTextDocument_ClearUndoRedoStacks(QtObjectPtr ptr, int stacksToClear);
QtObjectPtr QTextDocument_Clone(QtObjectPtr ptr, QtObjectPtr parent);
void QTextDocument_ConnectContentsChange(QtObjectPtr ptr);
void QTextDocument_DisconnectContentsChange(QtObjectPtr ptr);
void QTextDocument_ConnectContentsChanged(QtObjectPtr ptr);
void QTextDocument_DisconnectContentsChanged(QtObjectPtr ptr);
int QTextDocument_DefaultCursorMoveStyle(QtObjectPtr ptr);
QtObjectPtr QTextDocument_DocumentLayout(QtObjectPtr ptr);
void QTextDocument_ConnectDocumentLayoutChanged(QtObjectPtr ptr);
void QTextDocument_DisconnectDocumentLayoutChanged(QtObjectPtr ptr);
void QTextDocument_DrawContents(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr rect);
int QTextDocument_IsEmpty(QtObjectPtr ptr);
int QTextDocument_IsRedoAvailable(QtObjectPtr ptr);
int QTextDocument_IsUndoAvailable(QtObjectPtr ptr);
int QTextDocument_LineCount(QtObjectPtr ptr);
char* QTextDocument_MetaInformation(QtObjectPtr ptr, int info);
void QTextDocument_ConnectModificationChanged(QtObjectPtr ptr);
void QTextDocument_DisconnectModificationChanged(QtObjectPtr ptr);
QtObjectPtr QTextDocument_Object(QtObjectPtr ptr, int objectIndex);
QtObjectPtr QTextDocument_ObjectForFormat(QtObjectPtr ptr, QtObjectPtr f);
int QTextDocument_PageCount(QtObjectPtr ptr);
void QTextDocument_Print(QtObjectPtr ptr, QtObjectPtr printer);
void QTextDocument_Redo2(QtObjectPtr ptr);
void QTextDocument_Redo(QtObjectPtr ptr, QtObjectPtr cursor);
void QTextDocument_ConnectRedoAvailable(QtObjectPtr ptr);
void QTextDocument_DisconnectRedoAvailable(QtObjectPtr ptr);
char* QTextDocument_Resource(QtObjectPtr ptr, int ty, char* name);
int QTextDocument_Revision(QtObjectPtr ptr);
QtObjectPtr QTextDocument_RootFrame(QtObjectPtr ptr);
void QTextDocument_SetDefaultCursorMoveStyle(QtObjectPtr ptr, int style);
void QTextDocument_SetDefaultFont(QtObjectPtr ptr, QtObjectPtr font);
void QTextDocument_SetDefaultTextOption(QtObjectPtr ptr, QtObjectPtr option);
void QTextDocument_SetDocumentLayout(QtObjectPtr ptr, QtObjectPtr layout);
void QTextDocument_SetHtml(QtObjectPtr ptr, char* html);
void QTextDocument_SetMetaInformation(QtObjectPtr ptr, int info, char* stri);
void QTextDocument_SetPlainText(QtObjectPtr ptr, char* text);
char* QTextDocument_ToHtml(QtObjectPtr ptr, QtObjectPtr encoding);
char* QTextDocument_ToPlainText(QtObjectPtr ptr);
void QTextDocument_Undo2(QtObjectPtr ptr);
void QTextDocument_Undo(QtObjectPtr ptr, QtObjectPtr cursor);
void QTextDocument_ConnectUndoAvailable(QtObjectPtr ptr);
void QTextDocument_DisconnectUndoAvailable(QtObjectPtr ptr);
void QTextDocument_ConnectUndoCommandAdded(QtObjectPtr ptr);
void QTextDocument_DisconnectUndoCommandAdded(QtObjectPtr ptr);
void QTextDocument_DestroyQTextDocument(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
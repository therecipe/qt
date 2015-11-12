#include "qtextcursor.h"
#include <QTextList>
#include <QTextListFormat>
#include <QTextBlockFormat>
#include <QTextBlock>
#include <QTextCharFormat>
#include <QString>
#include <QImage>
#include <QTextDocument>
#include <QVariant>
#include <QModelIndex>
#include <QTextTableFormat>
#include <QTextFrameFormat>
#include <QTextTable>
#include <QTextImageFormat>
#include <QUrl>
#include <QTextDocumentFragment>
#include <QTextFrame>
#include <QTextCursor>
#include "_cgo_export.h"

class MyQTextCursor: public QTextCursor {
public:
};

void QTextCursor_InsertBlock3(void* ptr, void* format, void* charFormat){
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format), *static_cast<QTextCharFormat*>(charFormat));
}

void* QTextCursor_InsertTable2(void* ptr, int rows, int columns){
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns);
}

void* QTextCursor_InsertTable(void* ptr, int rows, int columns, void* format){
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns, *static_cast<QTextTableFormat*>(format));
}

void QTextCursor_InsertText2(void* ptr, char* text, void* format){
	static_cast<QTextCursor*>(ptr)->insertText(QString(text), *static_cast<QTextCharFormat*>(format));
}

int QTextCursor_MovePosition(void* ptr, int operation, int mode, int n){
	return static_cast<QTextCursor*>(ptr)->movePosition(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode), n);
}

void* QTextCursor_NewQTextCursor(){
	return new QTextCursor();
}

void* QTextCursor_NewQTextCursor2(void* document){
	return new QTextCursor(static_cast<QTextDocument*>(document));
}

void* QTextCursor_NewQTextCursor4(void* frame){
	return new QTextCursor(static_cast<QTextFrame*>(frame));
}

void* QTextCursor_NewQTextCursor5(void* block){
	return new QTextCursor(*static_cast<QTextBlock*>(block));
}

void* QTextCursor_NewQTextCursor7(void* cursor){
	return new QTextCursor(*static_cast<QTextCursor*>(cursor));
}

int QTextCursor_Anchor(void* ptr){
	return static_cast<QTextCursor*>(ptr)->anchor();
}

int QTextCursor_AtBlockEnd(void* ptr){
	return static_cast<QTextCursor*>(ptr)->atBlockEnd();
}

int QTextCursor_AtBlockStart(void* ptr){
	return static_cast<QTextCursor*>(ptr)->atBlockStart();
}

int QTextCursor_AtEnd(void* ptr){
	return static_cast<QTextCursor*>(ptr)->atEnd();
}

int QTextCursor_AtStart(void* ptr){
	return static_cast<QTextCursor*>(ptr)->atStart();
}

void QTextCursor_BeginEditBlock(void* ptr){
	static_cast<QTextCursor*>(ptr)->beginEditBlock();
}

int QTextCursor_BlockNumber(void* ptr){
	return static_cast<QTextCursor*>(ptr)->blockNumber();
}

void QTextCursor_ClearSelection(void* ptr){
	static_cast<QTextCursor*>(ptr)->clearSelection();
}

int QTextCursor_ColumnNumber(void* ptr){
	return static_cast<QTextCursor*>(ptr)->columnNumber();
}

void* QTextCursor_CreateList2(void* ptr, int style){
	return static_cast<QTextCursor*>(ptr)->createList(static_cast<QTextListFormat::Style>(style));
}

void* QTextCursor_CreateList(void* ptr, void* format){
	return static_cast<QTextCursor*>(ptr)->createList(*static_cast<QTextListFormat*>(format));
}

void* QTextCursor_CurrentFrame(void* ptr){
	return static_cast<QTextCursor*>(ptr)->currentFrame();
}

void* QTextCursor_CurrentList(void* ptr){
	return static_cast<QTextCursor*>(ptr)->currentList();
}

void* QTextCursor_CurrentTable(void* ptr){
	return static_cast<QTextCursor*>(ptr)->currentTable();
}

void QTextCursor_DeleteChar(void* ptr){
	static_cast<QTextCursor*>(ptr)->deleteChar();
}

void QTextCursor_DeletePreviousChar(void* ptr){
	static_cast<QTextCursor*>(ptr)->deletePreviousChar();
}

void* QTextCursor_Document(void* ptr){
	return static_cast<QTextCursor*>(ptr)->document();
}

void QTextCursor_EndEditBlock(void* ptr){
	static_cast<QTextCursor*>(ptr)->endEditBlock();
}

int QTextCursor_HasComplexSelection(void* ptr){
	return static_cast<QTextCursor*>(ptr)->hasComplexSelection();
}

int QTextCursor_HasSelection(void* ptr){
	return static_cast<QTextCursor*>(ptr)->hasSelection();
}

void QTextCursor_InsertBlock(void* ptr){
	static_cast<QTextCursor*>(ptr)->insertBlock();
}

void QTextCursor_InsertBlock2(void* ptr, void* format){
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_InsertFragment(void* ptr, void* fragment){
	static_cast<QTextCursor*>(ptr)->insertFragment(*static_cast<QTextDocumentFragment*>(fragment));
}

void* QTextCursor_InsertFrame(void* ptr, void* format){
	return static_cast<QTextCursor*>(ptr)->insertFrame(*static_cast<QTextFrameFormat*>(format));
}

void QTextCursor_InsertHtml(void* ptr, char* html){
	static_cast<QTextCursor*>(ptr)->insertHtml(QString(html));
}

void QTextCursor_InsertImage4(void* ptr, void* image, char* name){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QImage*>(image), QString(name));
}

void QTextCursor_InsertImage3(void* ptr, char* name){
	static_cast<QTextCursor*>(ptr)->insertImage(QString(name));
}

void QTextCursor_InsertImage(void* ptr, void* format){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format));
}

void QTextCursor_InsertImage2(void* ptr, void* format, int alignment){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format), static_cast<QTextFrameFormat::Position>(alignment));
}

void* QTextCursor_InsertList2(void* ptr, int style){
	return static_cast<QTextCursor*>(ptr)->insertList(static_cast<QTextListFormat::Style>(style));
}

void* QTextCursor_InsertList(void* ptr, void* format){
	return static_cast<QTextCursor*>(ptr)->insertList(*static_cast<QTextListFormat*>(format));
}

void QTextCursor_InsertText(void* ptr, char* text){
	static_cast<QTextCursor*>(ptr)->insertText(QString(text));
}

int QTextCursor_IsCopyOf(void* ptr, void* other){
	return static_cast<QTextCursor*>(ptr)->isCopyOf(*static_cast<QTextCursor*>(other));
}

int QTextCursor_IsNull(void* ptr){
	return static_cast<QTextCursor*>(ptr)->isNull();
}

void QTextCursor_JoinPreviousEditBlock(void* ptr){
	static_cast<QTextCursor*>(ptr)->joinPreviousEditBlock();
}

int QTextCursor_KeepPositionOnInsert(void* ptr){
	return static_cast<QTextCursor*>(ptr)->keepPositionOnInsert();
}

void QTextCursor_MergeBlockCharFormat(void* ptr, void* modifier){
	static_cast<QTextCursor*>(ptr)->mergeBlockCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextCursor_MergeBlockFormat(void* ptr, void* modifier){
	static_cast<QTextCursor*>(ptr)->mergeBlockFormat(*static_cast<QTextBlockFormat*>(modifier));
}

void QTextCursor_MergeCharFormat(void* ptr, void* modifier){
	static_cast<QTextCursor*>(ptr)->mergeCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

int QTextCursor_Position(void* ptr){
	return static_cast<QTextCursor*>(ptr)->position();
}

int QTextCursor_PositionInBlock(void* ptr){
	return static_cast<QTextCursor*>(ptr)->positionInBlock();
}

void QTextCursor_RemoveSelectedText(void* ptr){
	static_cast<QTextCursor*>(ptr)->removeSelectedText();
}

void QTextCursor_Select(void* ptr, int selection){
	static_cast<QTextCursor*>(ptr)->select(static_cast<QTextCursor::SelectionType>(selection));
}

void QTextCursor_SelectedTableCells(void* ptr, int firstRow, int numRows, int firstColumn, int numColumns){
	static_cast<QTextCursor*>(ptr)->selectedTableCells(&firstRow, &numRows, &firstColumn, &numColumns);
}

char* QTextCursor_SelectedText(void* ptr){
	return static_cast<QTextCursor*>(ptr)->selectedText().toUtf8().data();
}

int QTextCursor_SelectionEnd(void* ptr){
	return static_cast<QTextCursor*>(ptr)->selectionEnd();
}

int QTextCursor_SelectionStart(void* ptr){
	return static_cast<QTextCursor*>(ptr)->selectionStart();
}

void QTextCursor_SetBlockCharFormat(void* ptr, void* format){
	static_cast<QTextCursor*>(ptr)->setBlockCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetBlockFormat(void* ptr, void* format){
	static_cast<QTextCursor*>(ptr)->setBlockFormat(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_SetCharFormat(void* ptr, void* format){
	static_cast<QTextCursor*>(ptr)->setCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetKeepPositionOnInsert(void* ptr, int b){
	static_cast<QTextCursor*>(ptr)->setKeepPositionOnInsert(b != 0);
}

void QTextCursor_SetPosition(void* ptr, int pos, int m){
	static_cast<QTextCursor*>(ptr)->setPosition(pos, static_cast<QTextCursor::MoveMode>(m));
}

void QTextCursor_SetVerticalMovementX(void* ptr, int x){
	static_cast<QTextCursor*>(ptr)->setVerticalMovementX(x);
}

void QTextCursor_SetVisualNavigation(void* ptr, int b){
	static_cast<QTextCursor*>(ptr)->setVisualNavigation(b != 0);
}

void QTextCursor_Swap(void* ptr, void* other){
	static_cast<QTextCursor*>(ptr)->swap(*static_cast<QTextCursor*>(other));
}

int QTextCursor_VerticalMovementX(void* ptr){
	return static_cast<QTextCursor*>(ptr)->verticalMovementX();
}

int QTextCursor_VisualNavigation(void* ptr){
	return static_cast<QTextCursor*>(ptr)->visualNavigation();
}

void QTextCursor_DestroyQTextCursor(void* ptr){
	static_cast<QTextCursor*>(ptr)->~QTextCursor();
}


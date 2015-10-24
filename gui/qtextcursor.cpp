#include "qtextcursor.h"
#include <QTextList>
#include <QTextFrameFormat>
#include <QTextBlockFormat>
#include <QTextDocument>
#include <QTextDocumentFragment>
#include <QVariant>
#include <QUrl>
#include <QTextTable>
#include <QTextBlock>
#include <QImage>
#include <QTextListFormat>
#include <QString>
#include <QTextImageFormat>
#include <QTextCharFormat>
#include <QModelIndex>
#include <QTextTableFormat>
#include <QTextFrame>
#include <QTextCursor>
#include "_cgo_export.h"

class MyQTextCursor: public QTextCursor {
public:
};

void QTextCursor_InsertBlock3(QtObjectPtr ptr, QtObjectPtr format, QtObjectPtr charFormat){
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format), *static_cast<QTextCharFormat*>(charFormat));
}

QtObjectPtr QTextCursor_InsertTable2(QtObjectPtr ptr, int rows, int columns){
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns);
}

QtObjectPtr QTextCursor_InsertTable(QtObjectPtr ptr, int rows, int columns, QtObjectPtr format){
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns, *static_cast<QTextTableFormat*>(format));
}

void QTextCursor_InsertText2(QtObjectPtr ptr, char* text, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->insertText(QString(text), *static_cast<QTextCharFormat*>(format));
}

int QTextCursor_MovePosition(QtObjectPtr ptr, int operation, int mode, int n){
	return static_cast<QTextCursor*>(ptr)->movePosition(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode), n);
}

QtObjectPtr QTextCursor_NewQTextCursor(){
	return new QTextCursor();
}

QtObjectPtr QTextCursor_NewQTextCursor2(QtObjectPtr document){
	return new QTextCursor(static_cast<QTextDocument*>(document));
}

QtObjectPtr QTextCursor_NewQTextCursor4(QtObjectPtr frame){
	return new QTextCursor(static_cast<QTextFrame*>(frame));
}

QtObjectPtr QTextCursor_NewQTextCursor5(QtObjectPtr block){
	return new QTextCursor(*static_cast<QTextBlock*>(block));
}

QtObjectPtr QTextCursor_NewQTextCursor7(QtObjectPtr cursor){
	return new QTextCursor(*static_cast<QTextCursor*>(cursor));
}

int QTextCursor_Anchor(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->anchor();
}

int QTextCursor_AtBlockEnd(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->atBlockEnd();
}

int QTextCursor_AtBlockStart(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->atBlockStart();
}

int QTextCursor_AtEnd(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->atEnd();
}

int QTextCursor_AtStart(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->atStart();
}

void QTextCursor_BeginEditBlock(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->beginEditBlock();
}

int QTextCursor_BlockNumber(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->blockNumber();
}

void QTextCursor_ClearSelection(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->clearSelection();
}

int QTextCursor_ColumnNumber(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->columnNumber();
}

QtObjectPtr QTextCursor_CreateList2(QtObjectPtr ptr, int style){
	return static_cast<QTextCursor*>(ptr)->createList(static_cast<QTextListFormat::Style>(style));
}

QtObjectPtr QTextCursor_CreateList(QtObjectPtr ptr, QtObjectPtr format){
	return static_cast<QTextCursor*>(ptr)->createList(*static_cast<QTextListFormat*>(format));
}

QtObjectPtr QTextCursor_CurrentFrame(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->currentFrame();
}

QtObjectPtr QTextCursor_CurrentList(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->currentList();
}

QtObjectPtr QTextCursor_CurrentTable(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->currentTable();
}

void QTextCursor_DeleteChar(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->deleteChar();
}

void QTextCursor_DeletePreviousChar(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->deletePreviousChar();
}

QtObjectPtr QTextCursor_Document(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->document();
}

void QTextCursor_EndEditBlock(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->endEditBlock();
}

int QTextCursor_HasComplexSelection(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->hasComplexSelection();
}

int QTextCursor_HasSelection(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->hasSelection();
}

void QTextCursor_InsertBlock(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->insertBlock();
}

void QTextCursor_InsertBlock2(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_InsertFragment(QtObjectPtr ptr, QtObjectPtr fragment){
	static_cast<QTextCursor*>(ptr)->insertFragment(*static_cast<QTextDocumentFragment*>(fragment));
}

QtObjectPtr QTextCursor_InsertFrame(QtObjectPtr ptr, QtObjectPtr format){
	return static_cast<QTextCursor*>(ptr)->insertFrame(*static_cast<QTextFrameFormat*>(format));
}

void QTextCursor_InsertHtml(QtObjectPtr ptr, char* html){
	static_cast<QTextCursor*>(ptr)->insertHtml(QString(html));
}

void QTextCursor_InsertImage4(QtObjectPtr ptr, QtObjectPtr image, char* name){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QImage*>(image), QString(name));
}

void QTextCursor_InsertImage3(QtObjectPtr ptr, char* name){
	static_cast<QTextCursor*>(ptr)->insertImage(QString(name));
}

void QTextCursor_InsertImage(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format));
}

void QTextCursor_InsertImage2(QtObjectPtr ptr, QtObjectPtr format, int alignment){
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format), static_cast<QTextFrameFormat::Position>(alignment));
}

QtObjectPtr QTextCursor_InsertList2(QtObjectPtr ptr, int style){
	return static_cast<QTextCursor*>(ptr)->insertList(static_cast<QTextListFormat::Style>(style));
}

QtObjectPtr QTextCursor_InsertList(QtObjectPtr ptr, QtObjectPtr format){
	return static_cast<QTextCursor*>(ptr)->insertList(*static_cast<QTextListFormat*>(format));
}

void QTextCursor_InsertText(QtObjectPtr ptr, char* text){
	static_cast<QTextCursor*>(ptr)->insertText(QString(text));
}

int QTextCursor_IsCopyOf(QtObjectPtr ptr, QtObjectPtr other){
	return static_cast<QTextCursor*>(ptr)->isCopyOf(*static_cast<QTextCursor*>(other));
}

int QTextCursor_IsNull(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->isNull();
}

void QTextCursor_JoinPreviousEditBlock(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->joinPreviousEditBlock();
}

int QTextCursor_KeepPositionOnInsert(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->keepPositionOnInsert();
}

void QTextCursor_MergeBlockCharFormat(QtObjectPtr ptr, QtObjectPtr modifier){
	static_cast<QTextCursor*>(ptr)->mergeBlockCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextCursor_MergeBlockFormat(QtObjectPtr ptr, QtObjectPtr modifier){
	static_cast<QTextCursor*>(ptr)->mergeBlockFormat(*static_cast<QTextBlockFormat*>(modifier));
}

void QTextCursor_MergeCharFormat(QtObjectPtr ptr, QtObjectPtr modifier){
	static_cast<QTextCursor*>(ptr)->mergeCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

int QTextCursor_Position(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->position();
}

int QTextCursor_PositionInBlock(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->positionInBlock();
}

void QTextCursor_RemoveSelectedText(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->removeSelectedText();
}

void QTextCursor_Select(QtObjectPtr ptr, int selection){
	static_cast<QTextCursor*>(ptr)->select(static_cast<QTextCursor::SelectionType>(selection));
}

void QTextCursor_SelectedTableCells(QtObjectPtr ptr, int firstRow, int numRows, int firstColumn, int numColumns){
	static_cast<QTextCursor*>(ptr)->selectedTableCells(&firstRow, &numRows, &firstColumn, &numColumns);
}

char* QTextCursor_SelectedText(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->selectedText().toUtf8().data();
}

int QTextCursor_SelectionEnd(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->selectionEnd();
}

int QTextCursor_SelectionStart(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->selectionStart();
}

void QTextCursor_SetBlockCharFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->setBlockCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetBlockFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->setBlockFormat(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_SetCharFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextCursor*>(ptr)->setCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetKeepPositionOnInsert(QtObjectPtr ptr, int b){
	static_cast<QTextCursor*>(ptr)->setKeepPositionOnInsert(b != 0);
}

void QTextCursor_SetPosition(QtObjectPtr ptr, int pos, int m){
	static_cast<QTextCursor*>(ptr)->setPosition(pos, static_cast<QTextCursor::MoveMode>(m));
}

void QTextCursor_SetVerticalMovementX(QtObjectPtr ptr, int x){
	static_cast<QTextCursor*>(ptr)->setVerticalMovementX(x);
}

void QTextCursor_SetVisualNavigation(QtObjectPtr ptr, int b){
	static_cast<QTextCursor*>(ptr)->setVisualNavigation(b != 0);
}

void QTextCursor_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QTextCursor*>(ptr)->swap(*static_cast<QTextCursor*>(other));
}

int QTextCursor_VerticalMovementX(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->verticalMovementX();
}

int QTextCursor_VisualNavigation(QtObjectPtr ptr){
	return static_cast<QTextCursor*>(ptr)->visualNavigation();
}

void QTextCursor_DestroyQTextCursor(QtObjectPtr ptr){
	static_cast<QTextCursor*>(ptr)->~QTextCursor();
}


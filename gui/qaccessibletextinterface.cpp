#include "qaccessibletextinterface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QAccessible>
#include <QString>
#include <QAccessibleTextInterface>
#include "_cgo_export.h"

class MyQAccessibleTextInterface: public QAccessibleTextInterface {
public:
};

void QAccessibleTextInterface_AddSelection(QtObjectPtr ptr, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->addSelection(startOffset, endOffset);
}

char* QAccessibleTextInterface_Attributes(QtObjectPtr ptr, int offset, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->attributes(offset, &startOffset, &endOffset).toUtf8().data();
}

int QAccessibleTextInterface_CharacterCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->characterCount();
}

int QAccessibleTextInterface_CursorPosition(QtObjectPtr ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->cursorPosition();
}

int QAccessibleTextInterface_OffsetAtPoint(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QAccessibleTextInterface*>(ptr)->offsetAtPoint(*static_cast<QPoint*>(point));
}

void QAccessibleTextInterface_RemoveSelection(QtObjectPtr ptr, int selectionIndex){
	static_cast<QAccessibleTextInterface*>(ptr)->removeSelection(selectionIndex);
}

void QAccessibleTextInterface_ScrollToSubstring(QtObjectPtr ptr, int startIndex, int endIndex){
	static_cast<QAccessibleTextInterface*>(ptr)->scrollToSubstring(startIndex, endIndex);
}

void QAccessibleTextInterface_Selection(QtObjectPtr ptr, int selectionIndex, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->selection(selectionIndex, &startOffset, &endOffset);
}

int QAccessibleTextInterface_SelectionCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->selectionCount();
}

void QAccessibleTextInterface_SetCursorPosition(QtObjectPtr ptr, int position){
	static_cast<QAccessibleTextInterface*>(ptr)->setCursorPosition(position);
}

void QAccessibleTextInterface_SetSelection(QtObjectPtr ptr, int selectionIndex, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->setSelection(selectionIndex, startOffset, endOffset);
}

char* QAccessibleTextInterface_Text(QtObjectPtr ptr, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->text(startOffset, endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAfterOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textAfterOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAtOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textAtOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextBeforeOffset(QtObjectPtr ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textBeforeOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

void QAccessibleTextInterface_DestroyQAccessibleTextInterface(QtObjectPtr ptr){
	static_cast<QAccessibleTextInterface*>(ptr)->~QAccessibleTextInterface();
}


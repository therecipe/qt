#include "qaccessibletextinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QPoint>
#include <QAccessibleTextInterface>
#include "_cgo_export.h"

class MyQAccessibleTextInterface: public QAccessibleTextInterface {
public:
};

void QAccessibleTextInterface_AddSelection(void* ptr, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->addSelection(startOffset, endOffset);
}

char* QAccessibleTextInterface_Attributes(void* ptr, int offset, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->attributes(offset, &startOffset, &endOffset).toUtf8().data();
}

int QAccessibleTextInterface_CharacterCount(void* ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->characterCount();
}

int QAccessibleTextInterface_CursorPosition(void* ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->cursorPosition();
}

int QAccessibleTextInterface_OffsetAtPoint(void* ptr, void* point){
	return static_cast<QAccessibleTextInterface*>(ptr)->offsetAtPoint(*static_cast<QPoint*>(point));
}

void QAccessibleTextInterface_RemoveSelection(void* ptr, int selectionIndex){
	static_cast<QAccessibleTextInterface*>(ptr)->removeSelection(selectionIndex);
}

void QAccessibleTextInterface_ScrollToSubstring(void* ptr, int startIndex, int endIndex){
	static_cast<QAccessibleTextInterface*>(ptr)->scrollToSubstring(startIndex, endIndex);
}

void QAccessibleTextInterface_Selection(void* ptr, int selectionIndex, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->selection(selectionIndex, &startOffset, &endOffset);
}

int QAccessibleTextInterface_SelectionCount(void* ptr){
	return static_cast<QAccessibleTextInterface*>(ptr)->selectionCount();
}

void QAccessibleTextInterface_SetCursorPosition(void* ptr, int position){
	static_cast<QAccessibleTextInterface*>(ptr)->setCursorPosition(position);
}

void QAccessibleTextInterface_SetSelection(void* ptr, int selectionIndex, int startOffset, int endOffset){
	static_cast<QAccessibleTextInterface*>(ptr)->setSelection(selectionIndex, startOffset, endOffset);
}

char* QAccessibleTextInterface_Text(void* ptr, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->text(startOffset, endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAfterOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textAfterOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAtOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textAtOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextBeforeOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset){
	return static_cast<QAccessibleTextInterface*>(ptr)->textBeforeOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

void QAccessibleTextInterface_DestroyQAccessibleTextInterface(void* ptr){
	static_cast<QAccessibleTextInterface*>(ptr)->~QAccessibleTextInterface();
}


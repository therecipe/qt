#include "qheaderview.h"
#include <QVariant>
#include <QUrl>
#include <QMetaObject>
#include <QByteArray>
#include <QString>
#include <QModelIndex>
#include <QPoint>
#include <QObject>
#include <QAbstractItemModel>
#include <QHeaderView>
#include "_cgo_export.h"

class MyQHeaderView: public QHeaderView {
public:
void Signal_GeometriesChanged(){callbackQHeaderViewGeometriesChanged(this->objectName().toUtf8().data());};
void Signal_SectionClicked(int logicalIndex){callbackQHeaderViewSectionClicked(this->objectName().toUtf8().data(), logicalIndex);};
void Signal_SectionCountChanged(int oldCount, int newCount){callbackQHeaderViewSectionCountChanged(this->objectName().toUtf8().data(), oldCount, newCount);};
void Signal_SectionDoubleClicked(int logicalIndex){callbackQHeaderViewSectionDoubleClicked(this->objectName().toUtf8().data(), logicalIndex);};
void Signal_SectionEntered(int logicalIndex){callbackQHeaderViewSectionEntered(this->objectName().toUtf8().data(), logicalIndex);};
void Signal_SectionHandleDoubleClicked(int logicalIndex){callbackQHeaderViewSectionHandleDoubleClicked(this->objectName().toUtf8().data(), logicalIndex);};
void Signal_SectionMoved(int logicalIndex, int oldVisualIndex, int newVisualIndex){callbackQHeaderViewSectionMoved(this->objectName().toUtf8().data(), logicalIndex, oldVisualIndex, newVisualIndex);};
void Signal_SectionPressed(int logicalIndex){callbackQHeaderViewSectionPressed(this->objectName().toUtf8().data(), logicalIndex);};
void Signal_SectionResized(int logicalIndex, int oldSize, int newSize){callbackQHeaderViewSectionResized(this->objectName().toUtf8().data(), logicalIndex, oldSize, newSize);};
void Signal_SortIndicatorChanged(int logicalIndex, Qt::SortOrder order){callbackQHeaderViewSortIndicatorChanged(this->objectName().toUtf8().data(), logicalIndex, order);};
};

int QHeaderView_CascadingSectionResizes(void* ptr){
	return static_cast<QHeaderView*>(ptr)->cascadingSectionResizes();
}

int QHeaderView_DefaultAlignment(void* ptr){
	return static_cast<QHeaderView*>(ptr)->defaultAlignment();
}

int QHeaderView_DefaultSectionSize(void* ptr){
	return static_cast<QHeaderView*>(ptr)->defaultSectionSize();
}

int QHeaderView_HighlightSections(void* ptr){
	return static_cast<QHeaderView*>(ptr)->highlightSections();
}

int QHeaderView_IsSortIndicatorShown(void* ptr){
	return static_cast<QHeaderView*>(ptr)->isSortIndicatorShown();
}

int QHeaderView_MaximumSectionSize(void* ptr){
	return static_cast<QHeaderView*>(ptr)->maximumSectionSize();
}

int QHeaderView_MinimumSectionSize(void* ptr){
	return static_cast<QHeaderView*>(ptr)->minimumSectionSize();
}

void QHeaderView_ResetDefaultSectionSize(void* ptr){
	static_cast<QHeaderView*>(ptr)->resetDefaultSectionSize();
}

void QHeaderView_ResizeSection(void* ptr, int logicalIndex, int size){
	static_cast<QHeaderView*>(ptr)->resizeSection(logicalIndex, size);
}

void QHeaderView_SetCascadingSectionResizes(void* ptr, int enable){
	static_cast<QHeaderView*>(ptr)->setCascadingSectionResizes(enable != 0);
}

void QHeaderView_SetDefaultAlignment(void* ptr, int alignment){
	static_cast<QHeaderView*>(ptr)->setDefaultAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QHeaderView_SetDefaultSectionSize(void* ptr, int size){
	static_cast<QHeaderView*>(ptr)->setDefaultSectionSize(size);
}

void QHeaderView_SetHighlightSections(void* ptr, int highlight){
	static_cast<QHeaderView*>(ptr)->setHighlightSections(highlight != 0);
}

void QHeaderView_SetMaximumSectionSize(void* ptr, int size){
	static_cast<QHeaderView*>(ptr)->setMaximumSectionSize(size);
}

void QHeaderView_SetMinimumSectionSize(void* ptr, int size){
	static_cast<QHeaderView*>(ptr)->setMinimumSectionSize(size);
}

void QHeaderView_SetOffset(void* ptr, int offset){
	QMetaObject::invokeMethod(static_cast<QHeaderView*>(ptr), "setOffset", Q_ARG(int, offset));
}

void QHeaderView_SetSortIndicatorShown(void* ptr, int show){
	static_cast<QHeaderView*>(ptr)->setSortIndicatorShown(show != 0);
}

void QHeaderView_SetStretchLastSection(void* ptr, int stretch){
	static_cast<QHeaderView*>(ptr)->setStretchLastSection(stretch != 0);
}

int QHeaderView_StretchLastSection(void* ptr){
	return static_cast<QHeaderView*>(ptr)->stretchLastSection();
}

int QHeaderView_Count(void* ptr){
	return static_cast<QHeaderView*>(ptr)->count();
}

void QHeaderView_ConnectGeometriesChanged(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)()>(&QHeaderView::geometriesChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)()>(&MyQHeaderView::Signal_GeometriesChanged));;
}

void QHeaderView_DisconnectGeometriesChanged(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)()>(&QHeaderView::geometriesChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)()>(&MyQHeaderView::Signal_GeometriesChanged));;
}

void QHeaderView_HeaderDataChanged(void* ptr, int orientation, int logicalFirst, int logicalLast){
	QMetaObject::invokeMethod(static_cast<QHeaderView*>(ptr), "headerDataChanged", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(orientation)), Q_ARG(int, logicalFirst), Q_ARG(int, logicalLast));
}

int QHeaderView_HiddenSectionCount(void* ptr){
	return static_cast<QHeaderView*>(ptr)->hiddenSectionCount();
}

void QHeaderView_HideSection(void* ptr, int logicalIndex){
	static_cast<QHeaderView*>(ptr)->hideSection(logicalIndex);
}

int QHeaderView_IsSectionHidden(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->isSectionHidden(logicalIndex);
}

int QHeaderView_Length(void* ptr){
	return static_cast<QHeaderView*>(ptr)->length();
}

int QHeaderView_LogicalIndex(void* ptr, int visualIndex){
	return static_cast<QHeaderView*>(ptr)->logicalIndex(visualIndex);
}

int QHeaderView_LogicalIndexAt3(void* ptr, void* pos){
	return static_cast<QHeaderView*>(ptr)->logicalIndexAt(*static_cast<QPoint*>(pos));
}

int QHeaderView_LogicalIndexAt(void* ptr, int position){
	return static_cast<QHeaderView*>(ptr)->logicalIndexAt(position);
}

int QHeaderView_LogicalIndexAt2(void* ptr, int x, int y){
	return static_cast<QHeaderView*>(ptr)->logicalIndexAt(x, y);
}

void QHeaderView_MoveSection(void* ptr, int from, int to){
	static_cast<QHeaderView*>(ptr)->moveSection(from, to);
}

int QHeaderView_Offset(void* ptr){
	return static_cast<QHeaderView*>(ptr)->offset();
}

int QHeaderView_Orientation(void* ptr){
	return static_cast<QHeaderView*>(ptr)->orientation();
}

void QHeaderView_Reset(void* ptr){
	static_cast<QHeaderView*>(ptr)->reset();
}

int QHeaderView_ResizeContentsPrecision(void* ptr){
	return static_cast<QHeaderView*>(ptr)->resizeContentsPrecision();
}

void QHeaderView_ResizeSections(void* ptr, int mode){
	static_cast<QHeaderView*>(ptr)->resizeSections(static_cast<QHeaderView::ResizeMode>(mode));
}

int QHeaderView_RestoreState(void* ptr, void* state){
	return static_cast<QHeaderView*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void* QHeaderView_SaveState(void* ptr){
	return new QByteArray(static_cast<QHeaderView*>(ptr)->saveState());
}

void QHeaderView_ConnectSectionClicked(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionClicked));;
}

void QHeaderView_DisconnectSectionClicked(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionClicked));;
}

void QHeaderView_ConnectSectionCountChanged(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int)>(&QHeaderView::sectionCountChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int)>(&MyQHeaderView::Signal_SectionCountChanged));;
}

void QHeaderView_DisconnectSectionCountChanged(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int)>(&QHeaderView::sectionCountChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int)>(&MyQHeaderView::Signal_SectionCountChanged));;
}

void QHeaderView_ConnectSectionDoubleClicked(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionDoubleClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionDoubleClicked));;
}

void QHeaderView_DisconnectSectionDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionDoubleClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionDoubleClicked));;
}

void QHeaderView_ConnectSectionEntered(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionEntered), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionEntered));;
}

void QHeaderView_DisconnectSectionEntered(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionEntered), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionEntered));;
}

void QHeaderView_ConnectSectionHandleDoubleClicked(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionHandleDoubleClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionHandleDoubleClicked));;
}

void QHeaderView_DisconnectSectionHandleDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionHandleDoubleClicked), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionHandleDoubleClicked));;
}

void QHeaderView_ConnectSectionMoved(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int, int)>(&QHeaderView::sectionMoved), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int, int)>(&MyQHeaderView::Signal_SectionMoved));;
}

void QHeaderView_DisconnectSectionMoved(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int, int)>(&QHeaderView::sectionMoved), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int, int)>(&MyQHeaderView::Signal_SectionMoved));;
}

int QHeaderView_SectionPosition(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->sectionPosition(logicalIndex);
}

void QHeaderView_ConnectSectionPressed(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionPressed), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionPressed));;
}

void QHeaderView_DisconnectSectionPressed(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int)>(&QHeaderView::sectionPressed), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int)>(&MyQHeaderView::Signal_SectionPressed));;
}

int QHeaderView_SectionResizeMode(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->sectionResizeMode(logicalIndex);
}

void QHeaderView_ConnectSectionResized(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int, int)>(&QHeaderView::sectionResized), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int, int)>(&MyQHeaderView::Signal_SectionResized));;
}

void QHeaderView_DisconnectSectionResized(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, int, int)>(&QHeaderView::sectionResized), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, int, int)>(&MyQHeaderView::Signal_SectionResized));;
}

int QHeaderView_SectionSize(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->sectionSize(logicalIndex);
}

int QHeaderView_SectionSizeHint(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->sectionSizeHint(logicalIndex);
}

int QHeaderView_SectionViewportPosition(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->sectionViewportPosition(logicalIndex);
}

int QHeaderView_SectionsClickable(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sectionsClickable();
}

int QHeaderView_SectionsHidden(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sectionsHidden();
}

int QHeaderView_SectionsMovable(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sectionsMovable();
}

int QHeaderView_SectionsMoved(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sectionsMoved();
}

void QHeaderView_SetModel(void* ptr, void* model){
	static_cast<QHeaderView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHeaderView_SetOffsetToLastSection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHeaderView*>(ptr), "setOffsetToLastSection");
}

void QHeaderView_SetOffsetToSectionPosition(void* ptr, int visualSectionNumber){
	QMetaObject::invokeMethod(static_cast<QHeaderView*>(ptr), "setOffsetToSectionPosition", Q_ARG(int, visualSectionNumber));
}

void QHeaderView_SetResizeContentsPrecision(void* ptr, int precision){
	static_cast<QHeaderView*>(ptr)->setResizeContentsPrecision(precision);
}

void QHeaderView_SetSectionHidden(void* ptr, int logicalIndex, int hide){
	static_cast<QHeaderView*>(ptr)->setSectionHidden(logicalIndex, hide != 0);
}

void QHeaderView_SetSectionResizeMode(void* ptr, int mode){
	static_cast<QHeaderView*>(ptr)->setSectionResizeMode(static_cast<QHeaderView::ResizeMode>(mode));
}

void QHeaderView_SetSectionResizeMode2(void* ptr, int logicalIndex, int mode){
	static_cast<QHeaderView*>(ptr)->setSectionResizeMode(logicalIndex, static_cast<QHeaderView::ResizeMode>(mode));
}

void QHeaderView_SetSectionsClickable(void* ptr, int clickable){
	static_cast<QHeaderView*>(ptr)->setSectionsClickable(clickable != 0);
}

void QHeaderView_SetSectionsMovable(void* ptr, int movable){
	static_cast<QHeaderView*>(ptr)->setSectionsMovable(movable != 0);
}

void QHeaderView_SetSortIndicator(void* ptr, int logicalIndex, int order){
	static_cast<QHeaderView*>(ptr)->setSortIndicator(logicalIndex, static_cast<Qt::SortOrder>(order));
}

void QHeaderView_SetVisible(void* ptr, int v){
	static_cast<QHeaderView*>(ptr)->setVisible(v != 0);
}

void QHeaderView_ShowSection(void* ptr, int logicalIndex){
	static_cast<QHeaderView*>(ptr)->showSection(logicalIndex);
}

void QHeaderView_ConnectSortIndicatorChanged(void* ptr){
	QObject::connect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, Qt::SortOrder)>(&QHeaderView::sortIndicatorChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, Qt::SortOrder)>(&MyQHeaderView::Signal_SortIndicatorChanged));;
}

void QHeaderView_DisconnectSortIndicatorChanged(void* ptr){
	QObject::disconnect(static_cast<QHeaderView*>(ptr), static_cast<void (QHeaderView::*)(int, Qt::SortOrder)>(&QHeaderView::sortIndicatorChanged), static_cast<MyQHeaderView*>(ptr), static_cast<void (MyQHeaderView::*)(int, Qt::SortOrder)>(&MyQHeaderView::Signal_SortIndicatorChanged));;
}

int QHeaderView_SortIndicatorOrder(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sortIndicatorOrder();
}

int QHeaderView_SortIndicatorSection(void* ptr){
	return static_cast<QHeaderView*>(ptr)->sortIndicatorSection();
}

int QHeaderView_StretchSectionCount(void* ptr){
	return static_cast<QHeaderView*>(ptr)->stretchSectionCount();
}

void QHeaderView_SwapSections(void* ptr, int first, int second){
	static_cast<QHeaderView*>(ptr)->swapSections(first, second);
}

int QHeaderView_VisualIndex(void* ptr, int logicalIndex){
	return static_cast<QHeaderView*>(ptr)->visualIndex(logicalIndex);
}

int QHeaderView_VisualIndexAt(void* ptr, int position){
	return static_cast<QHeaderView*>(ptr)->visualIndexAt(position);
}

void QHeaderView_DestroyQHeaderView(void* ptr){
	static_cast<QHeaderView*>(ptr)->~QHeaderView();
}


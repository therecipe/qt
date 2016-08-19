// +build !minimal

#define protected public
#define private public

#include "help.h"
#include "_cgo_export.h"

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QAction>
#include <QActionEvent>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHelpContentItem>
#include <QHelpContentModel>
#include <QHelpContentWidget>
#include <QHelpEngine>
#include <QHelpEngineCore>
#include <QHelpIndexModel>
#include <QHelpIndexWidget>
#include <QHelpSearchEngine>
#include <QHelpSearchQuery>
#include <QHelpSearchQueryWidget>
#include <QHelpSearchResultWidget>
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QKeyEvent>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QStringList>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

void* QHelpContentItem_Child(void* ptr, int row)
{
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child)
{
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

void* QHelpContentItem_Parent(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

void* QHelpContentItem_Url(void* ptr)
{
	return new QUrl(static_cast<QHelpContentItem*>(ptr)->url());
}

void QHelpContentItem_DestroyQHelpContentItem(void* ptr)
{
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

class MyQHelpContentModel: public QHelpContentModel
{
public:
	void Signal_ContentsCreated() { callbackQHelpContentModel_ContentsCreated(this, this->objectName().toUtf8().data()); };
	void Signal_ContentsCreationStarted() { callbackQHelpContentModel_ContentsCreationStarted(this, this->objectName().toUtf8().data()); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Sibling(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), row, column, const_cast<QModelIndex*>(&index))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Buddy(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQHelpContentModel_CanDropMimeData(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQHelpContentModel_CanFetchMore(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)) != 0; };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQHelpContentModel_DropMimeData(this, this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQHelpContentModel_FetchMore(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQHelpContentModel_Flags(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQHelpContentModel_HasChildren(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQHelpContentModel_HeaderData(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), section, orientation, role)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpContentModel_InsertColumns(this, this->objectName().toUtf8().data(), column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQHelpContentModel_InsertRows(this, this->objectName().toUtf8().data(), row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QStringList mimeTypes() const { return QString(callbackQHelpContentModel_MimeTypes(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpContentModel_MoveColumns(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpContentModel_MoveRows(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpContentModel_RemoveColumns(this, this->objectName().toUtf8().data(), column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQHelpContentModel_RemoveRows(this, this->objectName().toUtf8().data(), row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQHelpContentModel_ResetInternalData(this, this->objectName().toUtf8().data()); };
	void revert() { callbackQHelpContentModel_Revert(this, this->objectName().toUtf8().data()); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQHelpContentModel_SetData(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQHelpContentModel_SetHeaderData(this, this->objectName().toUtf8().data(), section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQHelpContentModel_Sort(this, this->objectName().toUtf8().data(), column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQHelpContentModel_Span(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQHelpContentModel_Submit(this, this->objectName().toUtf8().data()) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQHelpContentModel_SupportedDragActions(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data())); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQHelpContentModel_SupportedDropActions(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQHelpContentModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpContentModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpContentModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpContentModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpContentModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpContentModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQHelpContentModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpContentModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpContentModel_MetaObject(const_cast<MyQHelpContentModel*>(this), this->objectName().toUtf8().data())); };
};

int QHelpContentModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel_ContentItemAt(void* ptr, void* index)
{
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void QHelpContentModel_ConnectContentsCreated(void* ptr)
{
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));
}

void QHelpContentModel_DisconnectContentsCreated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));
}

void QHelpContentModel_ContentsCreated(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->contentsCreated();
}

void QHelpContentModel_ConnectContentsCreationStarted(void* ptr)
{
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));
}

void QHelpContentModel_DisconnectContentsCreationStarted(void* ptr)
{
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));
}

void QHelpContentModel_ContentsCreationStarted(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->contentsCreationStarted();
}

void QHelpContentModel_CreateContents(void* ptr, char* customFilterName)
{
	static_cast<QHelpContentModel*>(ptr)->createContents(QString(customFilterName));
}

void* QHelpContentModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QHelpContentModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QHelpContentModel_IsCreatingContents(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

void* QHelpContentModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

int QHelpContentModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_DestroyQHelpContentModel(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

void* QHelpContentModel_Sibling(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_SiblingDefault(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::buddy(*static_cast<QModelIndex*>(index)));
}

int QHelpContentModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_CanDropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_DropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QHelpContentModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_Flags(void* ptr, void* index)
{
	return static_cast<QHelpContentModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QHelpContentModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::flags(*static_cast<QModelIndex*>(index));
}

int QHelpContentModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel_HeaderData(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QHelpContentModel_HeaderDataDefault(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

int QHelpContentModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char* QHelpContentModel_MimeTypes(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

char* QHelpContentModel_MimeTypesDefault(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::mimeTypes().join("|").toUtf8().data();
}

int QHelpContentModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpContentModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpContentModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpContentModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpContentModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpContentModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpContentModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentModel*>(ptr), "resetInternalData");
}

void QHelpContentModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::resetInternalData();
}

void QHelpContentModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentModel*>(ptr), "revert");
}

void QHelpContentModel_RevertDefault(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::revert();
}

int QHelpContentModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QHelpContentModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QHelpContentModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QHelpContentModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QHelpContentModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

int QHelpContentModel_SetHeaderDataDefault(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QHelpContentModel_Sort(void* ptr, int column, int order)
{
	static_cast<QHelpContentModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpContentModel_SortDefault(void* ptr, int column, int order)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QHelpContentModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QHelpContentModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QHelpContentModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpContentModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpContentModel_SubmitDefault(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::submit();
}

int QHelpContentModel_SupportedDragActions(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->supportedDragActions();
}

int QHelpContentModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::supportedDragActions();
}

int QHelpContentModel_SupportedDropActions(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->supportedDropActions();
}

int QHelpContentModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::supportedDropActions();
}

void QHelpContentModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpContentModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpContentModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::customEvent(static_cast<QEvent*>(event));
}

void QHelpContentModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentModel*>(ptr), "deleteLater");
}

void QHelpContentModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::deleteLater();
}

void QHelpContentModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpContentModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpContentModel_Event(void* ptr, void* e)
{
	return static_cast<QHelpContentModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHelpContentModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::event(static_cast<QEvent*>(e));
}

int QHelpContentModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpContentModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpContentModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpContentModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpContentModel*>(ptr)->metaObject());
}

void* QHelpContentModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::metaObject());
}

class MyQHelpContentWidget: public QHelpContentWidget
{
public:
	void Signal_LinkActivated(const QUrl & link) { callbackQHelpContentWidget_LinkActivated(this, this->objectName().toUtf8().data(), const_cast<QUrl*>(&link)); };
	void collapse(const QModelIndex & index) { callbackQHelpContentWidget_Collapse(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void expand(const QModelIndex & index) { callbackQHelpContentWidget_Expand(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void collapseAll() { callbackQHelpContentWidget_CollapseAll(this, this->objectName().toUtf8().data()); };
	void columnCountChanged(int oldCount, int newCount) { callbackQHelpContentWidget_ColumnCountChanged(this, this->objectName().toUtf8().data(), oldCount, newCount); };
	void columnMoved() { callbackQHelpContentWidget_ColumnMoved(this, this->objectName().toUtf8().data()); };
	void columnResized(int column, int oldSize, int newSize) { callbackQHelpContentWidget_ColumnResized(this, this->objectName().toUtf8().data(), column, oldSize, newSize); };
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQHelpContentWidget_CurrentChanged(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&current), const_cast<QModelIndex*>(&previous)); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpContentWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void drawBranches(QPainter * painter, const QRect & rect, const QModelIndex & index) const { callbackQHelpContentWidget_DrawBranches(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), painter, const_cast<QRect*>(&rect), const_cast<QModelIndex*>(&index)); };
	void drawRow(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQHelpContentWidget_DrawRow(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void expandAll() { callbackQHelpContentWidget_ExpandAll(this, this->objectName().toUtf8().data()); };
	void expandToDepth(int depth) { callbackQHelpContentWidget_ExpandToDepth(this, this->objectName().toUtf8().data(), depth); };
	void hideColumn(int column) { callbackQHelpContentWidget_HideColumn(this, this->objectName().toUtf8().data(), column); };
	int horizontalOffset() const { return callbackQHelpContentWidget_HorizontalOffset(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data()); };
	QModelIndex indexAt(const QPoint & point) const { return *static_cast<QModelIndex*>(callbackQHelpContentWidget_IndexAt(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), const_cast<QPoint*>(&point))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQHelpContentWidget_IsIndexHidden(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpContentWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyboardSearch(const QString & search) { callbackQHelpContentWidget_KeyboardSearch(this, this->objectName().toUtf8().data(), search.toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpContentWidget_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQHelpContentWidget_MoveCursor(this, this->objectName().toUtf8().data(), cursorAction, modifiers)); };
	void paintEvent(QPaintEvent * event) { callbackQHelpContentWidget_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void reset() { callbackQHelpContentWidget_Reset(this, this->objectName().toUtf8().data()); };
	void resizeColumnToContents(int column) { callbackQHelpContentWidget_ResizeColumnToContents(this, this->objectName().toUtf8().data(), column); };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsAboutToBeRemoved(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent), start, end); };
	void rowsInserted(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsInserted(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent), start, end); };
	void rowsRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsRemoved(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent), start, end); };
	void scrollContentsBy(int dx, int dy) { callbackQHelpContentWidget_ScrollContentsBy(this, this->objectName().toUtf8().data(), dx, dy); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQHelpContentWidget_ScrollTo(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), hint); };
	void selectAll() { callbackQHelpContentWidget_SelectAll(this, this->objectName().toUtf8().data()); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQHelpContentWidget_SelectionChanged(this, this->objectName().toUtf8().data(), const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setModel(QAbstractItemModel * model) { callbackQHelpContentWidget_SetModel(this, this->objectName().toUtf8().data(), model); };
	void setRootIndex(const QModelIndex & index) { callbackQHelpContentWidget_SetRootIndex(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpContentWidget_SetSelection(this, this->objectName().toUtf8().data(), const_cast<QRect*>(&rect), command); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpContentWidget_SetSelectionModel(this, this->objectName().toUtf8().data(), selectionModel); };
	void showColumn(int column) { callbackQHelpContentWidget_ShowColumn(this, this->objectName().toUtf8().data(), column); };
	int sizeHintForColumn(int column) const { return callbackQHelpContentWidget_SizeHintForColumn(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), column); };
	void updateGeometries() { callbackQHelpContentWidget_UpdateGeometries(this, this->objectName().toUtf8().data()); };
	int verticalOffset() const { return callbackQHelpContentWidget_VerticalOffset(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data()); };
	bool viewportEvent(QEvent * event) { return callbackQHelpContentWidget_ViewportEvent(this, this->objectName().toUtf8().data(), event) != 0; };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_ViewportSizeHint(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data())); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQHelpContentWidget_VisualRect(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQHelpContentWidget_VisualRegionForSelection(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), const_cast<QItemSelection*>(&selection))); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpContentWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void clearSelection() { callbackQHelpContentWidget_ClearSelection(this, this->objectName().toUtf8().data()); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackQHelpContentWidget_CloseEditor(this, this->objectName().toUtf8().data(), editor, hint); };
	void commitData(QWidget * editor) { callbackQHelpContentWidget_CommitData(this, this->objectName().toUtf8().data(), editor); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpContentWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpContentWidget_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void edit(const QModelIndex & index) { callbackQHelpContentWidget_Edit(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQHelpContentWidget_Edit2(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	void editorDestroyed(QObject * editor) { callbackQHelpContentWidget_EditorDestroyed(this, this->objectName().toUtf8().data(), editor); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpContentWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQHelpContentWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpContentWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpContentWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpContentWidget_InputMethodQuery(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpContentWidget_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void scrollToBottom() { callbackQHelpContentWidget_ScrollToBottom(this, this->objectName().toUtf8().data()); };
	void scrollToTop() { callbackQHelpContentWidget_ScrollToTop(this, this->objectName().toUtf8().data()); };
	QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex & index, const QEvent * event) const { return static_cast<QItemSelectionModel::SelectionFlag>(callbackQHelpContentWidget_SelectionCommand(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), const_cast<QEvent*>(event))); };
	void setCurrentIndex(const QModelIndex & index) { callbackQHelpContentWidget_SetCurrentIndex(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	int sizeHintForRow(int row) const { return callbackQHelpContentWidget_SizeHintForRow(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), row); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpContentWidget_StartDrag(this, this->objectName().toUtf8().data(), supportedActions); };
	void update(const QModelIndex & index) { callbackQHelpContentWidget_Update(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	QStyleOptionViewItem viewOptions() const { return *static_cast<QStyleOptionViewItem*>(callbackQHelpContentWidget_ViewOptions(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data())); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpContentWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_MinimumSizeHint(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data())); };
	void setupViewport(QWidget * viewport) { callbackQHelpContentWidget_SetupViewport(this, this->objectName().toUtf8().data(), viewport); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_SizeHint(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data())); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpContentWidget_WheelEvent(this, this->objectName().toUtf8().data(), e); };
	void changeEvent(QEvent * ev) { callbackQHelpContentWidget_ChangeEvent(this, this->objectName().toUtf8().data(), ev); };
	void actionEvent(QActionEvent * event) { callbackQHelpContentWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpContentWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpContentWidget_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpContentWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpContentWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQHelpContentWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQHelpContentWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQHelpContentWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQHelpContentWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQHelpContentWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQHelpContentWidget_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQHelpContentWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQHelpContentWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	bool hasHeightForWidth() const { return callbackQHelpContentWidget_HasHeightForWidth(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQHelpContentWidget_HeightForWidth(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQHelpContentWidget_Hide(this, this->objectName().toUtf8().data()); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpContentWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQHelpContentWidget_Lower(this, this->objectName().toUtf8().data()); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpContentWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQHelpContentWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQHelpContentWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQHelpContentWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQHelpContentWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQHelpContentWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQHelpContentWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQHelpContentWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQHelpContentWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQHelpContentWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQHelpContentWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpContentWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void updateMicroFocus() { callbackQHelpContentWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void childEvent(QChildEvent * event) { callbackQHelpContentWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpContentWidget_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpContentWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpContentWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpContentWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpContentWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpContentWidget_MetaObject(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data())); };
};

void* QHelpContentWidget_IndexOf(void* ptr, void* link)
{
	return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->indexOf(*static_cast<QUrl*>(link)));
}

void QHelpContentWidget_ConnectLinkActivated(void* ptr)
{
	QObject::connect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));
}

void QHelpContentWidget_DisconnectLinkActivated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));
}

void QHelpContentWidget_LinkActivated(void* ptr, void* link)
{
	static_cast<QHelpContentWidget*>(ptr)->linkActivated(*static_cast<QUrl*>(link));
}

void QHelpContentWidget_Collapse(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "collapse", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpContentWidget_CollapseDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::collapse(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_Expand(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "expand", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpContentWidget_ExpandDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expand(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_CollapseAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "collapseAll");
}

void QHelpContentWidget_CollapseAllDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::collapseAll();
}

void QHelpContentWidget_ColumnCountChanged(void* ptr, int oldCount, int newCount)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "columnCountChanged", Q_ARG(int, oldCount), Q_ARG(int, newCount));
}

void QHelpContentWidget_ColumnCountChangedDefault(void* ptr, int oldCount, int newCount)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnCountChanged(oldCount, newCount);
}

void QHelpContentWidget_ColumnMoved(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "columnMoved");
}

void QHelpContentWidget_ColumnMovedDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnMoved();
}

void QHelpContentWidget_ColumnResized(void* ptr, int column, int oldSize, int newSize)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "columnResized", Q_ARG(int, column), Q_ARG(int, oldSize), Q_ARG(int, newSize));
}

void QHelpContentWidget_ColumnResizedDefault(void* ptr, int column, int oldSize, int newSize)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnResized(column, oldSize, newSize);
}

void QHelpContentWidget_CurrentChanged(void* ptr, void* current, void* previous)
{
	static_cast<QHelpContentWidget*>(ptr)->currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpContentWidget_CurrentChangedDefault(void* ptr, void* current, void* previous)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpContentWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpContentWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpContentWidget_DrawBranches(void* ptr, void* painter, void* rect, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->drawBranches(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), *static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_DrawBranchesDefault(void* ptr, void* painter, void* rect, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::drawBranches(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), *static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_DrawRow(void* ptr, void* painter, void* option, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->drawRow(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_DrawRowDefault(void* ptr, void* painter, void* option, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::drawRow(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_ExpandAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "expandAll");
}

void QHelpContentWidget_ExpandAllDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expandAll();
}

void QHelpContentWidget_ExpandToDepth(void* ptr, int depth)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "expandToDepth", Q_ARG(int, depth));
}

void QHelpContentWidget_ExpandToDepthDefault(void* ptr, int depth)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expandToDepth(depth);
}

void QHelpContentWidget_HideColumn(void* ptr, int column)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "hideColumn", Q_ARG(int, column));
}

void QHelpContentWidget_HideColumnDefault(void* ptr, int column)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hideColumn(column);
}

int QHelpContentWidget_HorizontalOffset(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->horizontalOffset();
}

int QHelpContentWidget_HorizontalOffsetDefault(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::horizontalOffset();
}

void* QHelpContentWidget_IndexAt(void* ptr, void* point)
{
	return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->indexAt(*static_cast<QPoint*>(point)));
}

void* QHelpContentWidget_IndexAtDefault(void* ptr, void* point)
{
	return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::indexAt(*static_cast<QPoint*>(point)));
}

int QHelpContentWidget_IsIndexHidden(void* ptr, void* index)
{
	return static_cast<QHelpContentWidget*>(ptr)->isIndexHidden(*static_cast<QModelIndex*>(index));
}

int QHelpContentWidget_IsIndexHiddenDefault(void* ptr, void* index)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::isIndexHidden(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyboardSearch(void* ptr, char* search)
{
	static_cast<QHelpContentWidget*>(ptr)->keyboardSearch(QString(search));
}

void QHelpContentWidget_KeyboardSearchDefault(void* ptr, char* search)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyboardSearch(QString(search));
}

void QHelpContentWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void* QHelpContentWidget_MoveCursor(void* ptr, int cursorAction, int modifiers)
{
	return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void* QHelpContentWidget_MoveCursorDefault(void* ptr, int cursorAction, int modifiers)
{
	return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void QHelpContentWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpContentWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpContentWidget_Reset(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->reset();
}

void QHelpContentWidget_ResetDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::reset();
}

void QHelpContentWidget_ResizeColumnToContents(void* ptr, int column)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QHelpContentWidget_ResizeColumnToContentsDefault(void* ptr, int column)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::resizeColumnToContents(column);
}

void QHelpContentWidget_RowsAboutToBeRemoved(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpContentWidget*>(ptr)->rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsAboutToBeRemovedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsInserted(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpContentWidget*>(ptr)->rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsInsertedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsRemoved(void* ptr, void* parent, int start, int end)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "rowsRemoved", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(parent)), Q_ARG(int, start), Q_ARG(int, end));
}

void QHelpContentWidget_RowsRemovedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_ScrollContentsBy(void* ptr, int dx, int dy)
{
	static_cast<QHelpContentWidget*>(ptr)->scrollContentsBy(dx, dy);
}

void QHelpContentWidget_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollContentsBy(dx, dy);
}

void QHelpContentWidget_ScrollTo(void* ptr, void* index, int hint)
{
	static_cast<QHelpContentWidget*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpContentWidget_ScrollToDefault(void* ptr, void* index, int hint)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpContentWidget_SelectAll(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->selectAll();
}

void QHelpContentWidget_SelectAllDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectAll();
}

void QHelpContentWidget_SelectionChanged(void* ptr, void* selected, void* deselected)
{
	static_cast<QHelpContentWidget*>(ptr)->selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpContentWidget_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpContentWidget_SetModel(void* ptr, void* model)
{
	static_cast<QHelpContentWidget*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpContentWidget_SetModelDefault(void* ptr, void* model)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpContentWidget_SetRootIndex(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_SetRootIndexDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setRootIndex(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_SetSelection(void* ptr, void* rect, int command)
{
	static_cast<QHelpContentWidget*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpContentWidget_SetSelectionDefault(void* ptr, void* rect, int command)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpContentWidget_SetSelectionModel(void* ptr, void* selectionModel)
{
	static_cast<QHelpContentWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpContentWidget_SetSelectionModelDefault(void* ptr, void* selectionModel)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpContentWidget_ShowColumn(void* ptr, int column)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "showColumn", Q_ARG(int, column));
}

void QHelpContentWidget_ShowColumnDefault(void* ptr, int column)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showColumn(column);
}

int QHelpContentWidget_SizeHintForColumn(void* ptr, int column)
{
	return static_cast<QHelpContentWidget*>(ptr)->sizeHintForColumn(column);
}

int QHelpContentWidget_SizeHintForColumnDefault(void* ptr, int column)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHintForColumn(column);
}

void QHelpContentWidget_UpdateGeometries(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->updateGeometries();
}

void QHelpContentWidget_UpdateGeometriesDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::updateGeometries();
}

int QHelpContentWidget_VerticalOffset(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->verticalOffset();
}

int QHelpContentWidget_VerticalOffsetDefault(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::verticalOffset();
}

int QHelpContentWidget_ViewportEvent(void* ptr, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->viewportEvent(static_cast<QEvent*>(event));
}

int QHelpContentWidget_ViewportEventDefault(void* ptr, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewportEvent(static_cast<QEvent*>(event));
}

void* QHelpContentWidget_ViewportSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_ViewportSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_VisualRect(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QHelpContentWidget*>(ptr)->visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_VisualRectDefault(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_VisualRegionForSelection(void* ptr, void* selection)
{
	return new QRegion(static_cast<QHelpContentWidget*>(ptr)->visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QHelpContentWidget_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
	return new QRegion(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void QHelpContentWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpContentWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpContentWidget_ClearSelection(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "clearSelection");
}

void QHelpContentWidget_ClearSelectionDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::clearSelection();
}

void QHelpContentWidget_CloseEditor(void* ptr, void* editor, int hint)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpContentWidget_CloseEditorDefault(void* ptr, void* editor, int hint)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::closeEditor(static_cast<QWidget*>(editor), static_cast<QAbstractItemDelegate::EndEditHint>(hint));
}

void QHelpContentWidget_CommitData(void* ptr, void* editor)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpContentWidget_CommitDataDefault(void* ptr, void* editor)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::commitData(static_cast<QWidget*>(editor));
}

void QHelpContentWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpContentWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpContentWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpContentWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpContentWidget_Edit(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "edit", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpContentWidget_EditDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::edit(*static_cast<QModelIndex*>(index));
}

int QHelpContentWidget_Edit2(void* ptr, void* index, int trigger, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

int QHelpContentWidget_Edit2Default(void* ptr, void* index, int trigger, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

void QHelpContentWidget_EditorDestroyed(void* ptr, void* editor)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpContentWidget_EditorDestroyedDefault(void* ptr, void* editor)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::editorDestroyed(static_cast<QObject*>(editor));
}

void QHelpContentWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

int QHelpContentWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QHelpContentWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QHelpContentWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusNextPrevChild(next != 0);
}

void QHelpContentWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpContentWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QHelpContentWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpContentWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QHelpContentWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QHelpContentWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpContentWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpContentWidget_ScrollToBottom(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "scrollToBottom");
}

void QHelpContentWidget_ScrollToBottomDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollToBottom();
}

void QHelpContentWidget_ScrollToTop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "scrollToTop");
}

void QHelpContentWidget_ScrollToTopDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollToTop();
}

int QHelpContentWidget_SelectionCommand(void* ptr, void* index, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

int QHelpContentWidget_SelectionCommandDefault(void* ptr, void* index, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

void QHelpContentWidget_SetCurrentIndex(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpContentWidget_SetCurrentIndexDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setCurrentIndex(*static_cast<QModelIndex*>(index));
}

int QHelpContentWidget_SizeHintForRow(void* ptr, int row)
{
	return static_cast<QHelpContentWidget*>(ptr)->sizeHintForRow(row);
}

int QHelpContentWidget_SizeHintForRowDefault(void* ptr, int row)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHintForRow(row);
}

void QHelpContentWidget_StartDrag(void* ptr, int supportedActions)
{
	static_cast<QHelpContentWidget*>(ptr)->startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpContentWidget_StartDragDefault(void* ptr, int supportedActions)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpContentWidget_Update(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "update", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpContentWidget_UpdateDefault(void* ptr, void* index)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::update(*static_cast<QModelIndex*>(index));
}

void* QHelpContentWidget_ViewOptions(void* ptr)
{
	return new QStyleOptionViewItem(static_cast<QHelpContentWidget*>(ptr)->viewOptions());
}

void* QHelpContentWidget_ViewOptionsDefault(void* ptr)
{
	return new QStyleOptionViewItem(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewOptions());
}

void QHelpContentWidget_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QHelpContentWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpContentWidget_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void* QHelpContentWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpContentWidget_SetupViewport(void* ptr, void* viewport)
{
	static_cast<QHelpContentWidget*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpContentWidget_SetupViewportDefault(void* ptr, void* viewport)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void* QHelpContentWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpContentWidget_WheelEvent(void* ptr, void* e)
{
	static_cast<QHelpContentWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpContentWidget_WheelEventDefault(void* ptr, void* e)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpContentWidget_ChangeEvent(void* ptr, void* ev)
{
	static_cast<QHelpContentWidget*>(ptr)->changeEvent(static_cast<QEvent*>(ev));
}

void QHelpContentWidget_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::changeEvent(static_cast<QEvent*>(ev));
}

void QHelpContentWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpContentWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpContentWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpContentWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpContentWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpContentWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpContentWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QHelpContentWidget_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setEnabled(vbo != 0);
}

void QHelpContentWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QHelpContentWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setStyleSheet(QString(styleSheet));
}

void QHelpContentWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpContentWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setVisible(visible != 0);
}

void QHelpContentWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QHelpContentWidget_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setWindowModified(vbo != 0);
}

void QHelpContentWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QHelpContentWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setWindowTitle(QString(vqs));
}

void QHelpContentWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpContentWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showEvent(static_cast<QShowEvent*>(event));
}

int QHelpContentWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpContentWidget_CloseDefault(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::close();
}

void QHelpContentWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpContentWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

int QHelpContentWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->hasHeightForWidth();
}

int QHelpContentWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hasHeightForWidth();
}

int QHelpContentWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QHelpContentWidget*>(ptr)->heightForWidth(w);
}

int QHelpContentWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::heightForWidth(w);
}

void QHelpContentWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "hide");
}

void QHelpContentWidget_HideDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hide();
}

void QHelpContentWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "lower");
}

void QHelpContentWidget_LowerDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::lower();
}

int QHelpContentWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpContentWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QHelpContentWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QHelpContentWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "raise");
}

void QHelpContentWidget_RaiseDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::raise();
}

void QHelpContentWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "repaint");
}

void QHelpContentWidget_RepaintDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::repaint();
}

void QHelpContentWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QHelpContentWidget_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setDisabled(disable != 0);
}

void QHelpContentWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setFocus");
}

void QHelpContentWidget_SetFocus2Default(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setFocus();
}

void QHelpContentWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QHelpContentWidget_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setHidden(hidden != 0);
}

void QHelpContentWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "show");
}

void QHelpContentWidget_ShowDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::show();
}

void QHelpContentWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "showFullScreen");
}

void QHelpContentWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showFullScreen();
}

void QHelpContentWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "showMaximized");
}

void QHelpContentWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showMaximized();
}

void QHelpContentWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "showMinimized");
}

void QHelpContentWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showMinimized();
}

void QHelpContentWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "showNormal");
}

void QHelpContentWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showNormal();
}

void QHelpContentWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpContentWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpContentWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "updateMicroFocus");
}

void QHelpContentWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::updateMicroFocus();
}

void QHelpContentWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpContentWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "deleteLater");
}

void QHelpContentWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::deleteLater();
}

void QHelpContentWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpContentWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpContentWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpContentWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpContentWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpContentWidget*>(ptr)->metaObject());
}

void* QHelpContentWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::metaObject());
}

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent)
{
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

void* QHelpEngine_ContentModel(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_ContentWidget(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_IndexModel(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

void* QHelpEngine_IndexWidget(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

void* QHelpEngine_SearchEngine(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(void* ptr)
{
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

void QHelpEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::customEvent(static_cast<QEvent*>(event));
}

void QHelpEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpEngine*>(ptr), "deleteLater");
}

void QHelpEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::deleteLater();
}

void QHelpEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpEngine_Event(void* ptr, void* e)
{
	return static_cast<QHelpEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHelpEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QHelpEngine*>(ptr)->QHelpEngine::event(static_cast<QEvent*>(e));
}

int QHelpEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpEngine*>(ptr)->QHelpEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpEngine*>(ptr)->metaObject());
}

void* QHelpEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpEngine*>(ptr)->QHelpEngine::metaObject());
}

class MyQHelpEngineCore: public QHelpEngineCore
{
public:
	MyQHelpEngineCore(const QString &collectionFile, QObject *parent) : QHelpEngineCore(collectionFile, parent) {};
	void Signal_CurrentFilterChanged(const QString & newFilter) { callbackQHelpEngineCore_CurrentFilterChanged(this, this->objectName().toUtf8().data(), newFilter.toUtf8().data()); };
	void Signal_ReadersAboutToBeInvalidated() { callbackQHelpEngineCore_ReadersAboutToBeInvalidated(this, this->objectName().toUtf8().data()); };
	void Signal_SetupFinished() { callbackQHelpEngineCore_SetupFinished(this, this->objectName().toUtf8().data()); };
	void Signal_SetupStarted() { callbackQHelpEngineCore_SetupStarted(this, this->objectName().toUtf8().data()); };
	void Signal_Warning(const QString & msg) { callbackQHelpEngineCore_Warning(this, this->objectName().toUtf8().data(), msg.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQHelpEngineCore_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpEngineCore_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpEngineCore_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpEngineCore_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQHelpEngineCore_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpEngineCore_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpEngineCore_MetaObject(const_cast<MyQHelpEngineCore*>(this), this->objectName().toUtf8().data())); };
};

int QHelpEngineCore_AutoSaveFilter(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

char* QHelpEngineCore_CollectionFile(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8().data();
}

char* QHelpEngineCore_CurrentFilter(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8().data();
}

void QHelpEngineCore_SetAutoSaveFilter(void* ptr, int save)
{
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(void* ptr, char* fileName)
{
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString(fileName));
}

void QHelpEngineCore_SetCurrentFilter(void* ptr, char* filterName)
{
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString(filterName));
}

void* QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, void* parent)
{
	return new MyQHelpEngineCore(QString(collectionFile), static_cast<QObject*>(parent));
}

int QHelpEngineCore_AddCustomFilter(void* ptr, char* filterName, char* attributes)
{
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString(filterName), QString(attributes).split("|", QString::SkipEmptyParts));
}

int QHelpEngineCore_CopyCollectionFile(void* ptr, char* fileName)
{
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString(fileName));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));
}

void QHelpEngineCore_CurrentFilterChanged(void* ptr, char* newFilter)
{
	static_cast<QHelpEngineCore*>(ptr)->currentFilterChanged(QString(newFilter));
}

char* QHelpEngineCore_CustomFilters(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->customFilters().join("|").toUtf8().data();
}

void* QHelpEngineCore_CustomValue(void* ptr, char* key, void* defaultValue)
{
	return new QVariant(static_cast<QHelpEngineCore*>(ptr)->customValue(QString(key), *static_cast<QVariant*>(defaultValue)));
}

char* QHelpEngineCore_DocumentationFileName(void* ptr, char* namespaceName)
{
	return static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString(namespaceName)).toUtf8().data();
}

char* QHelpEngineCore_Error(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->error().toUtf8().data();
}

char* QHelpEngineCore_FileData(void* ptr, void* url)
{
	return QString(static_cast<QHelpEngineCore*>(ptr)->fileData(*static_cast<QUrl*>(url))).toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join("|").toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes2(void* ptr, char* filterName)
{
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString(filterName)).join("|").toUtf8().data();
}

void* QHelpEngineCore_FindFile(void* ptr, void* url)
{
	return new QUrl(static_cast<QHelpEngineCore*>(ptr)->findFile(*static_cast<QUrl*>(url)));
}

void* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name)
{
	return new QVariant(QHelpEngineCore::metaData(QString(documentationFileName), QString(name)));
}

char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName)
{
	return QHelpEngineCore::namespaceName(QString(documentationFileName)).toUtf8().data();
}

void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));
}

void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));
}

void QHelpEngineCore_ReadersAboutToBeInvalidated(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->readersAboutToBeInvalidated();
}

int QHelpEngineCore_RegisterDocumentation(void* ptr, char* documentationFileName)
{
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString(documentationFileName));
}

char* QHelpEngineCore_RegisteredDocumentations(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join("|").toUtf8().data();
}

int QHelpEngineCore_RemoveCustomFilter(void* ptr, char* filterName)
{
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString(filterName));
}

int QHelpEngineCore_RemoveCustomValue(void* ptr, char* key)
{
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString(key));
}

int QHelpEngineCore_SetCustomValue(void* ptr, char* key, void* value)
{
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString(key), *static_cast<QVariant*>(value));
}

int QHelpEngineCore_SetupData(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

void QHelpEngineCore_ConnectSetupFinished(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));
}

void QHelpEngineCore_DisconnectSetupFinished(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));
}

void QHelpEngineCore_SetupFinished(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->setupFinished();
}

void QHelpEngineCore_ConnectSetupStarted(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));
}

void QHelpEngineCore_DisconnectSetupStarted(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));
}

void QHelpEngineCore_SetupStarted(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->setupStarted();
}

int QHelpEngineCore_UnregisterDocumentation(void* ptr, char* namespaceName)
{
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString(namespaceName));
}

void QHelpEngineCore_ConnectWarning(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));
}

void QHelpEngineCore_DisconnectWarning(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));
}

void QHelpEngineCore_Warning(void* ptr, char* msg)
{
	static_cast<QHelpEngineCore*>(ptr)->warning(QString(msg));
}

void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

void QHelpEngineCore_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngineCore_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngineCore_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngineCore_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngineCore_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpEngineCore*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngineCore_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngineCore_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpEngineCore_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::customEvent(static_cast<QEvent*>(event));
}

void QHelpEngineCore_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpEngineCore*>(ptr), "deleteLater");
}

void QHelpEngineCore_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::deleteLater();
}

void QHelpEngineCore_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpEngineCore*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpEngineCore_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpEngineCore_Event(void* ptr, void* e)
{
	return static_cast<QHelpEngineCore*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHelpEngineCore_EventDefault(void* ptr, void* e)
{
	return static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::event(static_cast<QEvent*>(e));
}

int QHelpEngineCore_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpEngineCore*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpEngineCore_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpEngineCore_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpEngineCore*>(ptr)->metaObject());
}

void* QHelpEngineCore_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::metaObject());
}

class MyQHelpIndexModel: public QHelpIndexModel
{
public:
	void Signal_IndexCreated() { callbackQHelpIndexModel_IndexCreated(this, this->objectName().toUtf8().data()); };
	void Signal_IndexCreationStarted() { callbackQHelpIndexModel_IndexCreationStarted(this, this->objectName().toUtf8().data()); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQHelpIndexModel_Data(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), role)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQHelpIndexModel_Flags(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_InsertRows(this, this->objectName().toUtf8().data(), row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_RemoveRows(this, this->objectName().toUtf8().data(), row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	int rowCount(const QModelIndex & parent) const { return callbackQHelpIndexModel_RowCount(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQHelpIndexModel_SetData(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Sibling(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), row, column, const_cast<QModelIndex*>(&idx))); };
	void sort(int column, Qt::SortOrder order) { callbackQHelpIndexModel_Sort(this, this->objectName().toUtf8().data(), column, order); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQHelpIndexModel_SupportedDropActions(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data())); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Index(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), row, column, const_cast<QModelIndex*>(&parent))); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQHelpIndexModel_DropMimeData(this, this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Buddy(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQHelpIndexModel_CanDropMimeData(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQHelpIndexModel_CanFetchMore(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQHelpIndexModel_ColumnCount(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)); };
	void fetchMore(const QModelIndex & parent) { callbackQHelpIndexModel_FetchMore(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQHelpIndexModel_HasChildren(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQHelpIndexModel_HeaderData(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), section, orientation, role)); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_InsertColumns(this, this->objectName().toUtf8().data(), column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QStringList mimeTypes() const { return QString(callbackQHelpIndexModel_MimeTypes(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpIndexModel_MoveColumns(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpIndexModel_MoveRows(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Parent(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_RemoveColumns(this, this->objectName().toUtf8().data(), column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQHelpIndexModel_ResetInternalData(this, this->objectName().toUtf8().data()); };
	void revert() { callbackQHelpIndexModel_Revert(this, this->objectName().toUtf8().data()); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQHelpIndexModel_SetHeaderData(this, this->objectName().toUtf8().data(), section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQHelpIndexModel_Span(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQHelpIndexModel_Submit(this, this->objectName().toUtf8().data()) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQHelpIndexModel_SupportedDragActions(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQHelpIndexModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpIndexModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpIndexModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpIndexModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpIndexModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpIndexModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQHelpIndexModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpIndexModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpIndexModel_MetaObject(const_cast<MyQHelpIndexModel*>(this), this->objectName().toUtf8().data())); };
};

void QHelpIndexModel_CreateIndex(void* ptr, char* customFilterName)
{
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString(customFilterName));
}

void* QHelpIndexModel_Filter(void* ptr, char* filter, char* wildcard)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->filter(QString(filter), QString(wildcard)));
}

void QHelpIndexModel_ConnectIndexCreated(void* ptr)
{
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));
}

void QHelpIndexModel_DisconnectIndexCreated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));
}

void QHelpIndexModel_IndexCreated(void* ptr)
{
	static_cast<QHelpIndexModel*>(ptr)->indexCreated();
}

void QHelpIndexModel_ConnectIndexCreationStarted(void* ptr)
{
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));
}

void QHelpIndexModel_DisconnectIndexCreationStarted(void* ptr)
{
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));
}

void QHelpIndexModel_IndexCreationStarted(void* ptr)
{
	static_cast<QHelpIndexModel*>(ptr)->indexCreationStarted();
}

int QHelpIndexModel_IsCreatingIndex(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

void* QHelpIndexModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QHelpIndexModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QHelpIndexModel_DataDefault(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::data(*static_cast<QModelIndex*>(index), role));
}

int QHelpIndexModel_Flags(void* ptr, void* index)
{
	return static_cast<QHelpIndexModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QHelpIndexModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::flags(*static_cast<QModelIndex*>(index));
}

int QHelpIndexModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_RowCountDefault(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::rowCount(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QHelpIndexModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QHelpIndexModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void* QHelpIndexModel_Sibling(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QHelpIndexModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QHelpIndexModel_Sort(void* ptr, int column, int order)
{
	static_cast<QHelpIndexModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpIndexModel_SortDefault(void* ptr, int column, int order)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::sort(column, static_cast<Qt::SortOrder>(order));
}

int QHelpIndexModel_SupportedDropActions(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->supportedDropActions();
}

int QHelpIndexModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::supportedDropActions();
}

void* QHelpIndexModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QHelpIndexModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

int QHelpIndexModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_DropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QHelpIndexModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QHelpIndexModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::buddy(*static_cast<QModelIndex*>(index)));
}

int QHelpIndexModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_CanDropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void QHelpIndexModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QHelpIndexModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QHelpIndexModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QHelpIndexModel_HeaderData(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QHelpIndexModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QHelpIndexModel_HeaderDataDefault(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

int QHelpIndexModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char* QHelpIndexModel_MimeTypes(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

char* QHelpIndexModel_MimeTypesDefault(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::mimeTypes().join("|").toUtf8().data();
}

int QHelpIndexModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpIndexModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpIndexModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpIndexModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpIndexModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QHelpIndexModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QHelpIndexModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void* QHelpIndexModel_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::parent(*static_cast<QModelIndex*>(index)));
}

int QHelpIndexModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

void QHelpIndexModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexModel*>(ptr), "resetInternalData");
}

void QHelpIndexModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::resetInternalData();
}

void QHelpIndexModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexModel*>(ptr), "revert");
}

void QHelpIndexModel_RevertDefault(void* ptr)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::revert();
}

int QHelpIndexModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QHelpIndexModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

int QHelpIndexModel_SetHeaderDataDefault(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void* QHelpIndexModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexModel_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QHelpIndexModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpIndexModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpIndexModel_SubmitDefault(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::submit();
}

int QHelpIndexModel_SupportedDragActions(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->supportedDragActions();
}

int QHelpIndexModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::supportedDragActions();
}

void QHelpIndexModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpIndexModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpIndexModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpIndexModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexModel*>(ptr), "deleteLater");
}

void QHelpIndexModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::deleteLater();
}

void QHelpIndexModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpIndexModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpIndexModel_Event(void* ptr, void* e)
{
	return static_cast<QHelpIndexModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHelpIndexModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::event(static_cast<QEvent*>(e));
}

int QHelpIndexModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpIndexModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpIndexModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpIndexModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpIndexModel*>(ptr)->metaObject());
}

void* QHelpIndexModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::metaObject());
}

class MyQHelpIndexWidget: public QHelpIndexWidget
{
public:
	void activateCurrentItem() { callbackQHelpIndexWidget_ActivateCurrentItem(this, this->objectName().toUtf8().data()); };
	void filterIndices(const QString & filter, const QString & wildcard) { callbackQHelpIndexWidget_FilterIndices(this, this->objectName().toUtf8().data(), filter.toUtf8().data(), wildcard.toUtf8().data()); };
	void Signal_LinkActivated(const QUrl & link, const QString & keyword) { callbackQHelpIndexWidget_LinkActivated(this, this->objectName().toUtf8().data(), const_cast<QUrl*>(&link), keyword.toUtf8().data()); };
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQHelpIndexWidget_CurrentChanged(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&current), const_cast<QModelIndex*>(&previous)); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQHelpIndexWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQHelpIndexWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void dropEvent(QDropEvent * e) { callbackQHelpIndexWidget_DropEvent(this, this->objectName().toUtf8().data(), e); };
	int horizontalOffset() const { return callbackQHelpIndexWidget_HorizontalOffset(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data()); };
	QModelIndex indexAt(const QPoint & p) const { return *static_cast<QModelIndex*>(callbackQHelpIndexWidget_IndexAt(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), const_cast<QPoint*>(&p))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQHelpIndexWidget_IsIndexHidden(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)) != 0; };
	void mouseMoveEvent(QMouseEvent * e) { callbackQHelpIndexWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQHelpIndexWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQHelpIndexWidget_MoveCursor(this, this->objectName().toUtf8().data(), cursorAction, modifiers)); };
	void paintEvent(QPaintEvent * e) { callbackQHelpIndexWidget_PaintEvent(this, this->objectName().toUtf8().data(), e); };
	void resizeEvent(QResizeEvent * e) { callbackQHelpIndexWidget_ResizeEvent(this, this->objectName().toUtf8().data(), e); };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpIndexWidget_RowsAboutToBeRemoved(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent), start, end); };
	void rowsInserted(const QModelIndex & parent, int start, int end) { callbackQHelpIndexWidget_RowsInserted(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&parent), start, end); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQHelpIndexWidget_ScrollTo(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), hint); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQHelpIndexWidget_SelectionChanged(this, this->objectName().toUtf8().data(), const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpIndexWidget_SetSelection(this, this->objectName().toUtf8().data(), const_cast<QRect*>(&rect), command); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpIndexWidget_StartDrag(this, this->objectName().toUtf8().data(), supportedActions); };
	void updateGeometries() { callbackQHelpIndexWidget_UpdateGeometries(this, this->objectName().toUtf8().data()); };
	int verticalOffset() const { return callbackQHelpIndexWidget_VerticalOffset(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data()); };
	QStyleOptionViewItem viewOptions() const { return *static_cast<QStyleOptionViewItem*>(callbackQHelpIndexWidget_ViewOptions(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data())); };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_ViewportSizeHint(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data())); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQHelpIndexWidget_VisualRect(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQHelpIndexWidget_VisualRegionForSelection(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), const_cast<QItemSelection*>(&selection))); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpIndexWidget_WheelEvent(this, this->objectName().toUtf8().data(), e); };
	bool viewportEvent(QEvent * event) { return callbackQHelpIndexWidget_ViewportEvent(this, this->objectName().toUtf8().data(), event) != 0; };
	void clearSelection() { callbackQHelpIndexWidget_ClearSelection(this, this->objectName().toUtf8().data()); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackQHelpIndexWidget_CloseEditor(this, this->objectName().toUtf8().data(), editor, hint); };
	void commitData(QWidget * editor) { callbackQHelpIndexWidget_CommitData(this, this->objectName().toUtf8().data(), editor); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpIndexWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void edit(const QModelIndex & index) { callbackQHelpIndexWidget_Edit(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQHelpIndexWidget_Edit2(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	void editorDestroyed(QObject * editor) { callbackQHelpIndexWidget_EditorDestroyed(this, this->objectName().toUtf8().data(), editor); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpIndexWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQHelpIndexWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpIndexWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpIndexWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpIndexWidget_InputMethodQuery(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpIndexWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyboardSearch(const QString & search) { callbackQHelpIndexWidget_KeyboardSearch(this, this->objectName().toUtf8().data(), search.toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpIndexWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpIndexWidget_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void reset() { callbackQHelpIndexWidget_Reset(this, this->objectName().toUtf8().data()); };
	void scrollToBottom() { callbackQHelpIndexWidget_ScrollToBottom(this, this->objectName().toUtf8().data()); };
	void scrollToTop() { callbackQHelpIndexWidget_ScrollToTop(this, this->objectName().toUtf8().data()); };
	void selectAll() { callbackQHelpIndexWidget_SelectAll(this, this->objectName().toUtf8().data()); };
	QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex & index, const QEvent * event) const { return static_cast<QItemSelectionModel::SelectionFlag>(callbackQHelpIndexWidget_SelectionCommand(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index), const_cast<QEvent*>(event))); };
	void setCurrentIndex(const QModelIndex & index) { callbackQHelpIndexWidget_SetCurrentIndex(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void setModel(QAbstractItemModel * model) { callbackQHelpIndexWidget_SetModel(this, this->objectName().toUtf8().data(), model); };
	void setRootIndex(const QModelIndex & index) { callbackQHelpIndexWidget_SetRootIndex(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpIndexWidget_SetSelectionModel(this, this->objectName().toUtf8().data(), selectionModel); };
	int sizeHintForColumn(int column) const { return callbackQHelpIndexWidget_SizeHintForColumn(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), column); };
	int sizeHintForRow(int row) const { return callbackQHelpIndexWidget_SizeHintForRow(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), row); };
	void update(const QModelIndex & index) { callbackQHelpIndexWidget_Update(this, this->objectName().toUtf8().data(), const_cast<QModelIndex*>(&index)); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpIndexWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_MinimumSizeHint(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data())); };
	void scrollContentsBy(int dx, int dy) { callbackQHelpIndexWidget_ScrollContentsBy(this, this->objectName().toUtf8().data(), dx, dy); };
	void setupViewport(QWidget * viewport) { callbackQHelpIndexWidget_SetupViewport(this, this->objectName().toUtf8().data(), viewport); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_SizeHint(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * ev) { callbackQHelpIndexWidget_ChangeEvent(this, this->objectName().toUtf8().data(), ev); };
	void actionEvent(QActionEvent * event) { callbackQHelpIndexWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpIndexWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpIndexWidget_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpIndexWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpIndexWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQHelpIndexWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQHelpIndexWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQHelpIndexWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQHelpIndexWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQHelpIndexWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQHelpIndexWidget_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQHelpIndexWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQHelpIndexWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	bool hasHeightForWidth() const { return callbackQHelpIndexWidget_HasHeightForWidth(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQHelpIndexWidget_HeightForWidth(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQHelpIndexWidget_Hide(this, this->objectName().toUtf8().data()); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpIndexWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQHelpIndexWidget_Lower(this, this->objectName().toUtf8().data()); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpIndexWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQHelpIndexWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQHelpIndexWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void setDisabled(bool disable) { callbackQHelpIndexWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQHelpIndexWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQHelpIndexWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQHelpIndexWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQHelpIndexWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQHelpIndexWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQHelpIndexWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQHelpIndexWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpIndexWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void updateMicroFocus() { callbackQHelpIndexWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void childEvent(QChildEvent * event) { callbackQHelpIndexWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpIndexWidget_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpIndexWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpIndexWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpIndexWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpIndexWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpIndexWidget_MetaObject(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data())); };
};

void QHelpIndexWidget_ActivateCurrentItem(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(void* ptr, char* filter, char* wildcard)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

void QHelpIndexWidget_ConnectLinkActivated(void* ptr)
{
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));
}

void QHelpIndexWidget_DisconnectLinkActivated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));
}

void QHelpIndexWidget_LinkActivated(void* ptr, void* link, char* keyword)
{
	static_cast<QHelpIndexWidget*>(ptr)->linkActivated(*static_cast<QUrl*>(link), QString(keyword));
}

void QHelpIndexWidget_CurrentChanged(void* ptr, void* current, void* previous)
{
	static_cast<QHelpIndexWidget*>(ptr)->currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpIndexWidget_CurrentChangedDefault(void* ptr, void* current, void* previous)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpIndexWidget_DragLeaveEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QHelpIndexWidget_DragLeaveEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QHelpIndexWidget_DragMoveEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QHelpIndexWidget_DragMoveEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QHelpIndexWidget_DropEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(e));
}

void QHelpIndexWidget_DropEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dropEvent(static_cast<QDropEvent*>(e));
}

int QHelpIndexWidget_HorizontalOffset(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->horizontalOffset();
}

int QHelpIndexWidget_HorizontalOffsetDefault(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::horizontalOffset();
}

void* QHelpIndexWidget_IndexAt(void* ptr, void* p)
{
	return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->indexAt(*static_cast<QPoint*>(p)));
}

void* QHelpIndexWidget_IndexAtDefault(void* ptr, void* p)
{
	return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::indexAt(*static_cast<QPoint*>(p)));
}

int QHelpIndexWidget_IsIndexHidden(void* ptr, void* index)
{
	return static_cast<QHelpIndexWidget*>(ptr)->isIndexHidden(*static_cast<QModelIndex*>(index));
}

int QHelpIndexWidget_IsIndexHiddenDefault(void* ptr, void* index)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::isIndexHidden(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_MouseMoveEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseMoveEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseReleaseEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseReleaseEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void* QHelpIndexWidget_MoveCursor(void* ptr, int cursorAction, int modifiers)
{
	return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void* QHelpIndexWidget_MoveCursorDefault(void* ptr, int cursorAction, int modifiers)
{
	return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

void QHelpIndexWidget_PaintEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QHelpIndexWidget_PaintEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::paintEvent(static_cast<QPaintEvent*>(e));
}

void QHelpIndexWidget_ResizeEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(e));
}

void QHelpIndexWidget_ResizeEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void QHelpIndexWidget_RowsAboutToBeRemoved(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpIndexWidget*>(ptr)->rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_RowsAboutToBeRemovedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_RowsInserted(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpIndexWidget*>(ptr)->rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_RowsInsertedDefault(void* ptr, void* parent, int start, int end)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_ScrollTo(void* ptr, void* index, int hint)
{
	static_cast<QHelpIndexWidget*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpIndexWidget_ScrollToDefault(void* ptr, void* index, int hint)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpIndexWidget_SelectionChanged(void* ptr, void* selected, void* deselected)
{
	static_cast<QHelpIndexWidget*>(ptr)->selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpIndexWidget_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpIndexWidget_SetSelection(void* ptr, void* rect, int command)
{
	static_cast<QHelpIndexWidget*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpIndexWidget_SetSelectionDefault(void* ptr, void* rect, int command)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpIndexWidget_StartDrag(void* ptr, int supportedActions)
{
	static_cast<QHelpIndexWidget*>(ptr)->startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpIndexWidget_StartDragDefault(void* ptr, int supportedActions)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpIndexWidget_UpdateGeometries(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->updateGeometries();
}

void QHelpIndexWidget_UpdateGeometriesDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::updateGeometries();
}

int QHelpIndexWidget_VerticalOffset(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->verticalOffset();
}

int QHelpIndexWidget_VerticalOffsetDefault(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::verticalOffset();
}

void* QHelpIndexWidget_ViewOptions(void* ptr)
{
	return new QStyleOptionViewItem(static_cast<QHelpIndexWidget*>(ptr)->viewOptions());
}

void* QHelpIndexWidget_ViewOptionsDefault(void* ptr)
{
	return new QStyleOptionViewItem(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewOptions());
}

void* QHelpIndexWidget_ViewportSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_ViewportSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_VisualRect(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QHelpIndexWidget*>(ptr)->visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_VisualRectDefault(void* ptr, void* index)
{
	return ({ QRect tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_VisualRegionForSelection(void* ptr, void* selection)
{
	return new QRegion(static_cast<QHelpIndexWidget*>(ptr)->visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QHelpIndexWidget_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
	return new QRegion(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void QHelpIndexWidget_WheelEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpIndexWidget_WheelEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

int QHelpIndexWidget_ViewportEvent(void* ptr, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->viewportEvent(static_cast<QEvent*>(event));
}

int QHelpIndexWidget_ViewportEventDefault(void* ptr, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewportEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_ClearSelection(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "clearSelection");
}

void QHelpIndexWidget_ClearSelectionDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::clearSelection();
}

void QHelpIndexWidget_CloseEditor(void* ptr, void* editor, int hint)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpIndexWidget_CloseEditorDefault(void* ptr, void* editor, int hint)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::closeEditor(static_cast<QWidget*>(editor), static_cast<QAbstractItemDelegate::EndEditHint>(hint));
}

void QHelpIndexWidget_CommitData(void* ptr, void* editor)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpIndexWidget_CommitDataDefault(void* ptr, void* editor)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::commitData(static_cast<QWidget*>(editor));
}

void QHelpIndexWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpIndexWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpIndexWidget_Edit(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "edit", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpIndexWidget_EditDefault(void* ptr, void* index)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::edit(*static_cast<QModelIndex*>(index));
}

int QHelpIndexWidget_Edit2(void* ptr, void* index, int trigger, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

int QHelpIndexWidget_Edit2Default(void* ptr, void* index, int trigger, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

void QHelpIndexWidget_EditorDestroyed(void* ptr, void* editor)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpIndexWidget_EditorDestroyedDefault(void* ptr, void* editor)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::editorDestroyed(static_cast<QObject*>(editor));
}

void QHelpIndexWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

int QHelpIndexWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QHelpIndexWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QHelpIndexWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusNextPrevChild(next != 0);
}

void QHelpIndexWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpIndexWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QHelpIndexWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpIndexWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QHelpIndexWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QHelpIndexWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyboardSearch(void* ptr, char* search)
{
	static_cast<QHelpIndexWidget*>(ptr)->keyboardSearch(QString(search));
}

void QHelpIndexWidget_KeyboardSearchDefault(void* ptr, char* search)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyboardSearch(QString(search));
}

void QHelpIndexWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_Reset(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "reset");
}

void QHelpIndexWidget_ResetDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::reset();
}

void QHelpIndexWidget_ScrollToBottom(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "scrollToBottom");
}

void QHelpIndexWidget_ScrollToBottomDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollToBottom();
}

void QHelpIndexWidget_ScrollToTop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "scrollToTop");
}

void QHelpIndexWidget_ScrollToTopDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollToTop();
}

void QHelpIndexWidget_SelectAll(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "selectAll");
}

void QHelpIndexWidget_SelectAllDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectAll();
}

int QHelpIndexWidget_SelectionCommand(void* ptr, void* index, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

int QHelpIndexWidget_SelectionCommandDefault(void* ptr, void* index, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

void QHelpIndexWidget_SetCurrentIndex(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpIndexWidget_SetCurrentIndexDefault(void* ptr, void* index)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setCurrentIndex(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_SetModel(void* ptr, void* model)
{
	static_cast<QHelpIndexWidget*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpIndexWidget_SetModelDefault(void* ptr, void* model)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpIndexWidget_SetRootIndex(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setRootIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpIndexWidget_SetRootIndexDefault(void* ptr, void* index)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setRootIndex(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_SetSelectionModel(void* ptr, void* selectionModel)
{
	static_cast<QHelpIndexWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpIndexWidget_SetSelectionModelDefault(void* ptr, void* selectionModel)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QHelpIndexWidget_SizeHintForColumn(void* ptr, int column)
{
	return static_cast<QHelpIndexWidget*>(ptr)->sizeHintForColumn(column);
}

int QHelpIndexWidget_SizeHintForColumnDefault(void* ptr, int column)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHintForColumn(column);
}

int QHelpIndexWidget_SizeHintForRow(void* ptr, int row)
{
	return static_cast<QHelpIndexWidget*>(ptr)->sizeHintForRow(row);
}

int QHelpIndexWidget_SizeHintForRowDefault(void* ptr, int row)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHintForRow(row);
}

void QHelpIndexWidget_Update(void* ptr, void* index)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "update", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QHelpIndexWidget_UpdateDefault(void* ptr, void* index)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::update(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_ContextMenuEvent(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpIndexWidget_ContextMenuEventDefault(void* ptr, void* e)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void* QHelpIndexWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpIndexWidget_ScrollContentsBy(void* ptr, int dx, int dy)
{
	static_cast<QHelpIndexWidget*>(ptr)->scrollContentsBy(dx, dy);
}

void QHelpIndexWidget_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollContentsBy(dx, dy);
}

void QHelpIndexWidget_SetupViewport(void* ptr, void* viewport)
{
	static_cast<QHelpIndexWidget*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpIndexWidget_SetupViewportDefault(void* ptr, void* viewport)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void* QHelpIndexWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpIndexWidget_ChangeEvent(void* ptr, void* ev)
{
	static_cast<QHelpIndexWidget*>(ptr)->changeEvent(static_cast<QEvent*>(ev));
}

void QHelpIndexWidget_ChangeEventDefault(void* ptr, void* ev)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::changeEvent(static_cast<QEvent*>(ev));
}

void QHelpIndexWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpIndexWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpIndexWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpIndexWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpIndexWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpIndexWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpIndexWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QHelpIndexWidget_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setEnabled(vbo != 0);
}

void QHelpIndexWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QHelpIndexWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setStyleSheet(QString(styleSheet));
}

void QHelpIndexWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpIndexWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setVisible(visible != 0);
}

void QHelpIndexWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QHelpIndexWidget_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setWindowModified(vbo != 0);
}

void QHelpIndexWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QHelpIndexWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setWindowTitle(QString(vqs));
}

void QHelpIndexWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpIndexWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showEvent(static_cast<QShowEvent*>(event));
}

int QHelpIndexWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpIndexWidget_CloseDefault(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::close();
}

void QHelpIndexWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpIndexWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

int QHelpIndexWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->hasHeightForWidth();
}

int QHelpIndexWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hasHeightForWidth();
}

int QHelpIndexWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QHelpIndexWidget*>(ptr)->heightForWidth(w);
}

int QHelpIndexWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::heightForWidth(w);
}

void QHelpIndexWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "hide");
}

void QHelpIndexWidget_HideDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hide();
}

void QHelpIndexWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "lower");
}

void QHelpIndexWidget_LowerDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::lower();
}

int QHelpIndexWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpIndexWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QHelpIndexWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QHelpIndexWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "raise");
}

void QHelpIndexWidget_RaiseDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::raise();
}

void QHelpIndexWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "repaint");
}

void QHelpIndexWidget_RepaintDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::repaint();
}

void QHelpIndexWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QHelpIndexWidget_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setDisabled(disable != 0);
}

void QHelpIndexWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setFocus");
}

void QHelpIndexWidget_SetFocus2Default(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setFocus();
}

void QHelpIndexWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QHelpIndexWidget_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setHidden(hidden != 0);
}

void QHelpIndexWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "show");
}

void QHelpIndexWidget_ShowDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::show();
}

void QHelpIndexWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "showFullScreen");
}

void QHelpIndexWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showFullScreen();
}

void QHelpIndexWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "showMaximized");
}

void QHelpIndexWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showMaximized();
}

void QHelpIndexWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "showMinimized");
}

void QHelpIndexWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showMinimized();
}

void QHelpIndexWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "showNormal");
}

void QHelpIndexWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showNormal();
}

void QHelpIndexWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpIndexWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpIndexWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "updateMicroFocus");
}

void QHelpIndexWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::updateMicroFocus();
}

void QHelpIndexWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpIndexWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "deleteLater");
}

void QHelpIndexWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::deleteLater();
}

void QHelpIndexWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpIndexWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpIndexWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpIndexWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpIndexWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpIndexWidget*>(ptr)->metaObject());
}

void* QHelpIndexWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::metaObject());
}

class MyQHelpSearchEngine: public QHelpSearchEngine
{
public:
	MyQHelpSearchEngine(QHelpEngineCore *helpEngine, QObject *parent) : QHelpSearchEngine(helpEngine, parent) {};
	void cancelIndexing() { callbackQHelpSearchEngine_CancelIndexing(this, this->objectName().toUtf8().data()); };
	void cancelSearching() { callbackQHelpSearchEngine_CancelSearching(this, this->objectName().toUtf8().data()); };
	void Signal_IndexingFinished() { callbackQHelpSearchEngine_IndexingFinished(this, this->objectName().toUtf8().data()); };
	void Signal_IndexingStarted() { callbackQHelpSearchEngine_IndexingStarted(this, this->objectName().toUtf8().data()); };
	void reindexDocumentation() { callbackQHelpSearchEngine_ReindexDocumentation(this, this->objectName().toUtf8().data()); };
	void Signal_SearchingFinished(int hits) { callbackQHelpSearchEngine_SearchingFinished(this, this->objectName().toUtf8().data(), hits); };
	void Signal_SearchingStarted() { callbackQHelpSearchEngine_SearchingStarted(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchEngine_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchEngine_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchEngine_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchEngine_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpSearchEngine_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchEngine_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQHelpSearchEngine_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchEngine_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchEngine_MetaObject(const_cast<MyQHelpSearchEngine*>(this), this->objectName().toUtf8().data())); };
};

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent)
{
	return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));
}

void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));
}

void QHelpSearchEngine_IndexingFinished(void* ptr)
{
	static_cast<QHelpSearchEngine*>(ptr)->indexingFinished();
}

void QHelpSearchEngine_ConnectIndexingStarted(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));
}

void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));
}

void QHelpSearchEngine_IndexingStarted(void* ptr)
{
	static_cast<QHelpSearchEngine*>(ptr)->indexingStarted();
}

void* QHelpSearchEngine_QueryWidget(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void* QHelpSearchEngine_ResultWidget(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));
}

void QHelpSearchEngine_SearchingFinished(void* ptr, int hits)
{
	static_cast<QHelpSearchEngine*>(ptr)->searchingFinished(hits);
}

void QHelpSearchEngine_ConnectSearchingStarted(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));
}

void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));
}

void QHelpSearchEngine_SearchingStarted(void* ptr)
{
	static_cast<QHelpSearchEngine*>(ptr)->searchingStarted();
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr)
{
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

void QHelpSearchEngine_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchEngine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchEngine_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchEngine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchEngine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchEngine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchEngine_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchEngine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchEngine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "deleteLater");
}

void QHelpSearchEngine_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::deleteLater();
}

void QHelpSearchEngine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchEngine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpSearchEngine_Event(void* ptr, void* e)
{
	return static_cast<QHelpSearchEngine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHelpSearchEngine_EventDefault(void* ptr, void* e)
{
	return static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::event(static_cast<QEvent*>(e));
}

int QHelpSearchEngine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchEngine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpSearchEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpSearchEngine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchEngine*>(ptr)->metaObject());
}

void* QHelpSearchEngine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::metaObject());
}

void* QHelpSearchQuery_NewQHelpSearchQuery()
{
	return new QHelpSearchQuery();
}

void* QHelpSearchQuery_NewQHelpSearchQuery2(int field, char* wordList)
{
	return new QHelpSearchQuery(static_cast<QHelpSearchQuery::FieldName>(field), QString(wordList).split("|", QString::SkipEmptyParts));
}

int QHelpSearchQuery_FieldName(void* ptr)
{
	return static_cast<QHelpSearchQuery*>(ptr)->fieldName;
}

void QHelpSearchQuery_SetFieldName(void* ptr, int vfi)
{
	static_cast<QHelpSearchQuery*>(ptr)->fieldName = static_cast<QHelpSearchQuery::FieldName>(vfi);
}

char* QHelpSearchQuery_WordList(void* ptr)
{
	return static_cast<QHelpSearchQuery*>(ptr)->wordList.join("|").toUtf8().data();
}

void QHelpSearchQuery_SetWordList(void* ptr, char* vqs)
{
	static_cast<QHelpSearchQuery*>(ptr)->wordList = QString(vqs).split("|", QString::SkipEmptyParts);
}

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget
{
public:
	MyQHelpSearchQueryWidget(QWidget *parent) : QHelpSearchQueryWidget(parent) {};
	void Signal_Search() { callbackQHelpSearchQueryWidget_Search(this, this->objectName().toUtf8().data()); };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchQueryWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchQueryWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchQueryWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchQueryWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchQueryWidget_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchQueryWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpSearchQueryWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchQueryWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchQueryWidget_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchQueryWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchQueryWidget_MinimumSizeHint(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchQueryWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchQueryWidget_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQHelpSearchQueryWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQHelpSearchQueryWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQHelpSearchQueryWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQHelpSearchQueryWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQHelpSearchQueryWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQHelpSearchQueryWidget_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchQueryWidget_SizeHint(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQHelpSearchQueryWidget_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQHelpSearchQueryWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchQueryWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchQueryWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQHelpSearchQueryWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQHelpSearchQueryWidget_HasHeightForWidth(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQHelpSearchQueryWidget_HeightForWidth(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQHelpSearchQueryWidget_Hide(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchQueryWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpSearchQueryWidget_InputMethodQuery(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQHelpSearchQueryWidget_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpSearchQueryWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQHelpSearchQueryWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQHelpSearchQueryWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchQueryWidget_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQHelpSearchQueryWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQHelpSearchQueryWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQHelpSearchQueryWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQHelpSearchQueryWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQHelpSearchQueryWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQHelpSearchQueryWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQHelpSearchQueryWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQHelpSearchQueryWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchQueryWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQHelpSearchQueryWidget_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQHelpSearchQueryWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchQueryWidget_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchQueryWidget_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchQueryWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchQueryWidget_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchQueryWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpSearchQueryWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchQueryWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchQueryWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchQueryWidget_MetaObject(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data())); };
};

int QHelpSearchQueryWidget_IsCompactMode(void* ptr)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(void* parent)
{
	return new MyQHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));
}

void QHelpSearchQueryWidget_DisconnectSearch(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));
}

void QHelpSearchQueryWidget_Search(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->search();
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

void QHelpSearchQueryWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchQueryWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchQueryWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchQueryWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchQueryWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchQueryWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchQueryWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchQueryWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchQueryWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchQueryWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchQueryWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchQueryWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchQueryWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* QHelpSearchQueryWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchQueryWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpSearchQueryWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchQueryWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchQueryWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchQueryWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchQueryWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QHelpSearchQueryWidget_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setEnabled(vbo != 0);
}

void QHelpSearchQueryWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QHelpSearchQueryWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setStyleSheet(QString(styleSheet));
}

void QHelpSearchQueryWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchQueryWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setVisible(visible != 0);
}

void QHelpSearchQueryWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QHelpSearchQueryWidget_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setWindowModified(vbo != 0);
}

void QHelpSearchQueryWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QHelpSearchQueryWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setWindowTitle(QString(vqs));
}

void QHelpSearchQueryWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchQueryWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QHelpSearchQueryWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchQueryWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpSearchQueryWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::changeEvent(static_cast<QEvent*>(event));
}

int QHelpSearchQueryWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpSearchQueryWidget_CloseDefault(void* ptr)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::close();
}

void QHelpSearchQueryWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchQueryWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchQueryWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchQueryWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QHelpSearchQueryWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QHelpSearchQueryWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusNextPrevChild(next != 0);
}

int QHelpSearchQueryWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->hasHeightForWidth();
}

int QHelpSearchQueryWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hasHeightForWidth();
}

int QHelpSearchQueryWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->heightForWidth(w);
}

int QHelpSearchQueryWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::heightForWidth(w);
}

void QHelpSearchQueryWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "hide");
}

void QHelpSearchQueryWidget_HideDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hide();
}

void QHelpSearchQueryWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchQueryWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QHelpSearchQueryWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpSearchQueryWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QHelpSearchQueryWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QHelpSearchQueryWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "lower");
}

void QHelpSearchQueryWidget_LowerDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::lower();
}

void QHelpSearchQueryWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QHelpSearchQueryWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QHelpSearchQueryWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QHelpSearchQueryWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "raise");
}

void QHelpSearchQueryWidget_RaiseDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::raise();
}

void QHelpSearchQueryWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "repaint");
}

void QHelpSearchQueryWidget_RepaintDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::repaint();
}

void QHelpSearchQueryWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchQueryWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchQueryWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QHelpSearchQueryWidget_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setDisabled(disable != 0);
}

void QHelpSearchQueryWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setFocus");
}

void QHelpSearchQueryWidget_SetFocus2Default(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setFocus();
}

void QHelpSearchQueryWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QHelpSearchQueryWidget_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setHidden(hidden != 0);
}

void QHelpSearchQueryWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "show");
}

void QHelpSearchQueryWidget_ShowDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::show();
}

void QHelpSearchQueryWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "showFullScreen");
}

void QHelpSearchQueryWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showFullScreen();
}

void QHelpSearchQueryWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "showMaximized");
}

void QHelpSearchQueryWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showMaximized();
}

void QHelpSearchQueryWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "showMinimized");
}

void QHelpSearchQueryWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showMinimized();
}

void QHelpSearchQueryWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "showNormal");
}

void QHelpSearchQueryWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showNormal();
}

void QHelpSearchQueryWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchQueryWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchQueryWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "update");
}

void QHelpSearchQueryWidget_UpdateDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::update();
}

void QHelpSearchQueryWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "updateMicroFocus");
}

void QHelpSearchQueryWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::updateMicroFocus();
}

void QHelpSearchQueryWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchQueryWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchQueryWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchQueryWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchQueryWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchQueryWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchQueryWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchQueryWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchQueryWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "deleteLater");
}

void QHelpSearchQueryWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::deleteLater();
}

void QHelpSearchQueryWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchQueryWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpSearchQueryWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpSearchQueryWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpSearchQueryWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchQueryWidget*>(ptr)->metaObject());
}

void* QHelpSearchQueryWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::metaObject());
}

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget
{
public:
	void Signal_RequestShowLink(const QUrl & link) { callbackQHelpSearchResultWidget_RequestShowLink(this, this->objectName().toUtf8().data(), const_cast<QUrl*>(&link)); };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchResultWidget_ActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchResultWidget_DragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchResultWidget_DragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchResultWidget_DragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchResultWidget_DropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchResultWidget_EnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpSearchResultWidget_FocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchResultWidget_FocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchResultWidget_HideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchResultWidget_LeaveEvent(this, this->objectName().toUtf8().data(), event); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchResultWidget_MinimumSizeHint(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data())); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchResultWidget_MoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchResultWidget_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setEnabled(bool vbo) { callbackQHelpSearchResultWidget_SetEnabled(this, this->objectName().toUtf8().data(), vbo); };
	void setStyleSheet(const QString & styleSheet) { callbackQHelpSearchResultWidget_SetStyleSheet(this, this->objectName().toUtf8().data(), styleSheet.toUtf8().data()); };
	void setVisible(bool visible) { callbackQHelpSearchResultWidget_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWindowModified(bool vbo) { callbackQHelpSearchResultWidget_SetWindowModified(this, this->objectName().toUtf8().data(), vbo); };
	void setWindowTitle(const QString & vqs) { callbackQHelpSearchResultWidget_SetWindowTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void showEvent(QShowEvent * event) { callbackQHelpSearchResultWidget_ShowEvent(this, this->objectName().toUtf8().data(), event); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchResultWidget_SizeHint(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data())); };
	void changeEvent(QEvent * event) { callbackQHelpSearchResultWidget_ChangeEvent(this, this->objectName().toUtf8().data(), event); };
	bool close() { return callbackQHelpSearchResultWidget_Close(this, this->objectName().toUtf8().data()) != 0; };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchResultWidget_CloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchResultWidget_ContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	bool focusNextPrevChild(bool next) { return callbackQHelpSearchResultWidget_FocusNextPrevChild(this, this->objectName().toUtf8().data(), next) != 0; };
	bool hasHeightForWidth() const { return callbackQHelpSearchResultWidget_HasHeightForWidth(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data()) != 0; };
	int heightForWidth(int w) const { return callbackQHelpSearchResultWidget_HeightForWidth(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data(), w); };
	void hide() { callbackQHelpSearchResultWidget_Hide(this, this->objectName().toUtf8().data()); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchResultWidget_InputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpSearchResultWidget_InputMethodQuery(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data(), query)); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchResultWidget_KeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchResultWidget_KeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void lower() { callbackQHelpSearchResultWidget_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpSearchResultWidget_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQHelpSearchResultWidget_Raise(this, this->objectName().toUtf8().data()); };
	void repaint() { callbackQHelpSearchResultWidget_Repaint(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchResultWidget_ResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void setDisabled(bool disable) { callbackQHelpSearchResultWidget_SetDisabled(this, this->objectName().toUtf8().data(), disable); };
	void setFocus() { callbackQHelpSearchResultWidget_SetFocus2(this, this->objectName().toUtf8().data()); };
	void setHidden(bool hidden) { callbackQHelpSearchResultWidget_SetHidden(this, this->objectName().toUtf8().data(), hidden); };
	void show() { callbackQHelpSearchResultWidget_Show(this, this->objectName().toUtf8().data()); };
	void showFullScreen() { callbackQHelpSearchResultWidget_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQHelpSearchResultWidget_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQHelpSearchResultWidget_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQHelpSearchResultWidget_ShowNormal(this, this->objectName().toUtf8().data()); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchResultWidget_TabletEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQHelpSearchResultWidget_Update(this, this->objectName().toUtf8().data()); };
	void updateMicroFocus() { callbackQHelpSearchResultWidget_UpdateMicroFocus(this, this->objectName().toUtf8().data()); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchResultWidget_WheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchResultWidget_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchResultWidget_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchResultWidget_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchResultWidget_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQHelpSearchResultWidget_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchResultWidget_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchResultWidget_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchResultWidget_MetaObject(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data())); };
};

void* QHelpSearchResultWidget_LinkAt(void* ptr, void* point)
{
	return new QUrl(static_cast<QHelpSearchResultWidget*>(ptr)->linkAt(*static_cast<QPoint*>(point)));
}

void QHelpSearchResultWidget_ConnectRequestShowLink(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));
}

void QHelpSearchResultWidget_DisconnectRequestShowLink(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));
}

void QHelpSearchResultWidget_RequestShowLink(void* ptr, void* link)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->requestShowLink(*static_cast<QUrl*>(link));
}

void QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->~QHelpSearchResultWidget();
}

void QHelpSearchResultWidget_ActionEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchResultWidget_ActionEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchResultWidget_DragEnterEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchResultWidget_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchResultWidget_DragLeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchResultWidget_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchResultWidget_DragMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchResultWidget_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchResultWidget_DropEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchResultWidget_DropEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchResultWidget_EnterEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_EnterEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_FocusInEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusInEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusOutEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusOutEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_HideEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchResultWidget_HideEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchResultWidget_LeaveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_LeaveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::leaveEvent(static_cast<QEvent*>(event));
}

void* QHelpSearchResultWidget_MinimumSizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchResultWidget_MinimumSizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpSearchResultWidget_MoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchResultWidget_MoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchResultWidget_PaintEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchResultWidget_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchResultWidget_SetEnabled(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setEnabled", Q_ARG(bool, vbo != 0));
}

void QHelpSearchResultWidget_SetEnabledDefault(void* ptr, int vbo)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setEnabled(vbo != 0);
}

void QHelpSearchResultWidget_SetStyleSheet(void* ptr, char* styleSheet)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QHelpSearchResultWidget_SetStyleSheetDefault(void* ptr, char* styleSheet)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setStyleSheet(QString(styleSheet));
}

void QHelpSearchResultWidget_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchResultWidget_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setVisible(visible != 0);
}

void QHelpSearchResultWidget_SetWindowModified(void* ptr, int vbo)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setWindowModified", Q_ARG(bool, vbo != 0));
}

void QHelpSearchResultWidget_SetWindowModifiedDefault(void* ptr, int vbo)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setWindowModified(vbo != 0);
}

void QHelpSearchResultWidget_SetWindowTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(vqs)));
}

void QHelpSearchResultWidget_SetWindowTitleDefault(void* ptr, char* vqs)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setWindowTitle(QString(vqs));
}

void QHelpSearchResultWidget_ShowEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchResultWidget_ShowEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showEvent(static_cast<QShowEvent*>(event));
}

void* QHelpSearchResultWidget_SizeHint(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchResultWidget_SizeHintDefault(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpSearchResultWidget_ChangeEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->changeEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_ChangeEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::changeEvent(static_cast<QEvent*>(event));
}

int QHelpSearchResultWidget_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHelpSearchResultWidget_CloseDefault(void* ptr)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::close();
}

void QHelpSearchResultWidget_CloseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchResultWidget_CloseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchResultWidget_ContextMenuEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchResultWidget_ContextMenuEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

int QHelpSearchResultWidget_FocusNextPrevChild(void* ptr, int next)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->focusNextPrevChild(next != 0);
}

int QHelpSearchResultWidget_FocusNextPrevChildDefault(void* ptr, int next)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusNextPrevChild(next != 0);
}

int QHelpSearchResultWidget_HasHeightForWidth(void* ptr)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->hasHeightForWidth();
}

int QHelpSearchResultWidget_HasHeightForWidthDefault(void* ptr)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hasHeightForWidth();
}

int QHelpSearchResultWidget_HeightForWidth(void* ptr, int w)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->heightForWidth(w);
}

int QHelpSearchResultWidget_HeightForWidthDefault(void* ptr, int w)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::heightForWidth(w);
}

void QHelpSearchResultWidget_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "hide");
}

void QHelpSearchResultWidget_HideDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hide();
}

void QHelpSearchResultWidget_InputMethodEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchResultWidget_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* QHelpSearchResultWidget_InputMethodQuery(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpSearchResultWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QHelpSearchResultWidget_InputMethodQueryDefault(void* ptr, int query)
{
	return new QVariant(static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QHelpSearchResultWidget_KeyPressEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "lower");
}

void QHelpSearchResultWidget_LowerDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::lower();
}

void QHelpSearchResultWidget_MouseDoubleClickEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseMoveEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MousePressEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseReleaseEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

int QHelpSearchResultWidget_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QHelpSearchResultWidget_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::nativeEvent(QByteArray(eventType), message, &result);
}

void QHelpSearchResultWidget_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "raise");
}

void QHelpSearchResultWidget_RaiseDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::raise();
}

void QHelpSearchResultWidget_Repaint(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "repaint");
}

void QHelpSearchResultWidget_RepaintDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::repaint();
}

void QHelpSearchResultWidget_ResizeEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchResultWidget_ResizeEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchResultWidget_SetDisabled(void* ptr, int disable)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QHelpSearchResultWidget_SetDisabledDefault(void* ptr, int disable)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setDisabled(disable != 0);
}

void QHelpSearchResultWidget_SetFocus2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setFocus");
}

void QHelpSearchResultWidget_SetFocus2Default(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setFocus();
}

void QHelpSearchResultWidget_SetHidden(void* ptr, int hidden)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QHelpSearchResultWidget_SetHiddenDefault(void* ptr, int hidden)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setHidden(hidden != 0);
}

void QHelpSearchResultWidget_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "show");
}

void QHelpSearchResultWidget_ShowDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::show();
}

void QHelpSearchResultWidget_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "showFullScreen");
}

void QHelpSearchResultWidget_ShowFullScreenDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showFullScreen();
}

void QHelpSearchResultWidget_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "showMaximized");
}

void QHelpSearchResultWidget_ShowMaximizedDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showMaximized();
}

void QHelpSearchResultWidget_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "showMinimized");
}

void QHelpSearchResultWidget_ShowMinimizedDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showMinimized();
}

void QHelpSearchResultWidget_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "showNormal");
}

void QHelpSearchResultWidget_ShowNormalDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showNormal();
}

void QHelpSearchResultWidget_TabletEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchResultWidget_TabletEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchResultWidget_Update(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "update");
}

void QHelpSearchResultWidget_UpdateDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::update();
}

void QHelpSearchResultWidget_UpdateMicroFocus(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "updateMicroFocus");
}

void QHelpSearchResultWidget_UpdateMicroFocusDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::updateMicroFocus();
}

void QHelpSearchResultWidget_WheelEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchResultWidget_WheelEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchResultWidget_TimerEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchResultWidget_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchResultWidget_ChildEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchResultWidget_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchResultWidget_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchResultWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchResultWidget_CustomEvent(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "deleteLater");
}

void QHelpSearchResultWidget_DeleteLaterDefault(void* ptr)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::deleteLater();
}

void QHelpSearchResultWidget_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchResultWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHelpSearchResultWidget_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHelpSearchResultWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHelpSearchResultWidget_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchResultWidget*>(ptr)->metaObject());
}

void* QHelpSearchResultWidget_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::metaObject());
}


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
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QFocusEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QHelpContentItem>
#include <QHelpContentModel>
#include <QHelpContentWidget>
#include <QHelpEngine>
#include <QHelpEngineCore>
#include <QHelpIndexModel>
#include <QHelpIndexWidget>
#include <QHelpSearchEngine>
#include <QHelpSearchQueryWidget>
#include <QHelpSearchResult>
#include <QHelpSearchResultWidget>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QKeyEvent>
#include <QLayout>
#include <QList>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QPoint>
#include <QQuickItem>
#include <QRadioData>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSignalSpy>
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
#include <QVector>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>

void QHelpContentItem_DestroyQHelpContentItem(void* ptr)
{
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

void* QHelpContentItem_Child(void* ptr, int row)
{
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

void* QHelpContentItem_Parent(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

struct QtHelp_PackedString QHelpContentItem_Title(void* ptr)
{
	return ({ QByteArray t03003b = static_cast<QHelpContentItem*>(ptr)->title().toUtf8(); QtHelp_PackedString { const_cast<char*>(t03003b.prepend("WHITESPACE").constData()+10), t03003b.size()-10 }; });
}

void* QHelpContentItem_Url(void* ptr)
{
	return new QUrl(static_cast<QHelpContentItem*>(ptr)->url());
}

int QHelpContentItem_ChildCount(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child)
{
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

int QHelpContentItem_Row(void* ptr)
{
	return static_cast<QHelpContentItem*>(ptr)->row();
}

class MyQHelpContentModel: public QHelpContentModel
{
public:
	void Signal_ContentsCreated() { callbackQHelpContentModel_ContentsCreated(this); };
	void Signal_ContentsCreationStarted() { callbackQHelpContentModel_ContentsCreationStarted(this); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQHelpContentModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	int columnCount(const QModelIndex & parent) const { return callbackQHelpContentModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	int rowCount(const QModelIndex & parent) const { return callbackQHelpContentModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQHelpContentModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpContentModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQHelpContentModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpContentModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpContentModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpContentModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQHelpContentModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQHelpContentModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQHelpContentModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQHelpContentModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackQHelpContentModel_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQHelpContentModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQHelpContentModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQHelpContentModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackQHelpContentModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQHelpContentModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQHelpContentModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); QtHelp_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQHelpContentModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); QtHelp_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackQHelpContentModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQHelpContentModel_ModelReset(this); };
	void resetInternalData() { callbackQHelpContentModel_ResetInternalData(this); };
	void revert() { callbackQHelpContentModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQHelpContentModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQHelpContentModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQHelpContentModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpContentModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void sort(int column, Qt::SortOrder order) { callbackQHelpContentModel_Sort(this, column, order); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackQHelpContentModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackQHelpContentModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQHelpContentModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); QtHelp_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpContentModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackQHelpContentModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQHelpContentModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ QtHelp_PackedString tempVal = callbackQHelpContentModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQHelpContentModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQHelpContentModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQHelpContentModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQHelpContentModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQHelpContentModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQHelpContentModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackQHelpContentModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool event(QEvent * e) { return callbackQHelpContentModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpContentModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpContentModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpContentModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpContentModel_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpContentModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpContentModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpContentModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpContentModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpContentModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpContentModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpContentModel*)

int QHelpContentModel_QHelpContentModel_QRegisterMetaType(){qRegisterMetaType<QHelpContentModel*>(); return qRegisterMetaType<MyQHelpContentModel*>();}

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

void QHelpContentModel_CreateContents(void* ptr, struct QtHelp_PackedString customFilterName)
{
	static_cast<QHelpContentModel*>(ptr)->createContents(QString::fromUtf8(customFilterName.data, customFilterName.len));
}

void QHelpContentModel_DestroyQHelpContentModel(void* ptr)
{
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

void* QHelpContentModel_ContentItemAt(void* ptr, void* index)
{
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void* QHelpContentModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QHelpContentModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QHelpContentModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_ParentDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::parent(*static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QHelpContentModel_DataDefault(void* ptr, void* index, int role)
{
		return new QVariant(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::data(*static_cast<QModelIndex*>(index), role));
}

char QHelpContentModel_IsCreatingContents(void* ptr)
{
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

int QHelpContentModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::columnCount(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QHelpContentModel_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::rowCount(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QHelpContentModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QHelpContentModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct QtHelp_PackedList QHelpContentModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpContentModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpContentModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int QHelpContentModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QHelpContentModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QHelpContentModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* QHelpContentModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QHelpContentModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* QHelpContentModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QHelpContentModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* QHelpContentModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void QHelpContentModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QHelpContentModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>;
}

struct QtHelp_PackedList QHelpContentModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpContentModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QHelpContentModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QHelpContentModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct QtHelp_PackedList QHelpContentModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpContentModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpContentModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpContentModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int QHelpContentModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpContentModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpContentModel_____setItemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpContentModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpContentModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpContentModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpContentModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpContentModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpContentModel_____setRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpContentModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpContentModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpContentModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpContentModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpContentModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpContentModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* QHelpContentModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpContentModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpContentModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpContentModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentModel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpContentModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpContentModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QHelpContentModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QHelpContentModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QHelpContentModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QHelpContentModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QHelpContentModel_SubmitDefault(void* ptr)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::submit();
}

void QHelpContentModel_FetchMoreDefault(void* ptr, void* parent)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_ResetInternalDataDefault(void* ptr)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::resetInternalData();
}

void QHelpContentModel_RevertDefault(void* ptr)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::revert();
}

void QHelpContentModel_SortDefault(void* ptr, int column, long long order)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::sort(column, static_cast<Qt::SortOrder>(order));
}

struct QtHelp_PackedList QHelpContentModel_RoleNamesDefault(void* ptr)
{
		return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::roleNames()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedList QHelpContentModel_ItemDataDefault(void* ptr, void* index)
{
		return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::itemData(*static_cast<QModelIndex*>(index))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpContentModel_MimeDataDefault(void* ptr, void* indexes)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* QHelpContentModel_BuddyDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* QHelpContentModel_SiblingDefault(void* ptr, int row, int column, void* index)
{
		return new QModelIndex(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

struct QtHelp_PackedList QHelpContentModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpContentModel_SpanDefault(void* ptr, void* index)
{
		return ({ QSize tmpValue = static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtHelp_PackedString QHelpContentModel_MimeTypesDefault(void* ptr)
{
		return ({ QByteArray tf8824d = static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::mimeTypes().join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(tf8824d.prepend("WHITESPACE").constData()+10), tf8824d.size()-10 }; });
}

void* QHelpContentModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
		return new QVariant(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long QHelpContentModel_SupportedDragActionsDefault(void* ptr)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::supportedDragActions();
}

long long QHelpContentModel_SupportedDropActionsDefault(void* ptr)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::supportedDropActions();
}

long long QHelpContentModel_FlagsDefault(void* ptr, void* index)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::flags(*static_cast<QModelIndex*>(index));
}

char QHelpContentModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_CanFetchMoreDefault(void* ptr, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

char QHelpContentModel_EventDefault(void* ptr, void* e)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::event(static_cast<QEvent*>(e));
}

char QHelpContentModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpContentModel_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentModel_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentModel_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::customEvent(static_cast<QEvent*>(event));
}

void QHelpContentModel_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::deleteLater();
}

void QHelpContentModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentModel_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHelpContentModel_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::metaObject());
}

class MyQHelpContentWidget: public QHelpContentWidget
{
public:
	void Signal_LinkActivated(const QUrl & link) { callbackQHelpContentWidget_LinkActivated(this, const_cast<QUrl*>(&link)); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQHelpContentWidget_MoveCursor(this, cursorAction, modifiers)); };
	bool viewportEvent(QEvent * event) { return callbackQHelpContentWidget_ViewportEvent(this, event) != 0; };
	void collapse(const QModelIndex & index) { callbackQHelpContentWidget_Collapse(this, const_cast<QModelIndex*>(&index)); };
	void collapseAll() { callbackQHelpContentWidget_CollapseAll(this); };
	void Signal_Collapsed(const QModelIndex & index) { callbackQHelpContentWidget_Collapsed(this, const_cast<QModelIndex*>(&index)); };
	void columnCountChanged(int oldCount, int newCount) { callbackQHelpContentWidget_ColumnCountChanged(this, oldCount, newCount); };
	void columnMoved() { callbackQHelpContentWidget_ColumnMoved(this); };
	void columnResized(int column, int oldSize, int newSize) { callbackQHelpContentWidget_ColumnResized(this, column, oldSize, newSize); };
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQHelpContentWidget_CurrentChanged(this, const_cast<QModelIndex*>(&current), const_cast<QModelIndex*>(&previous)); };
	void dataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQHelpContentWidget_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpContentWidget_DragMoveEvent(this, event); };
	void expand(const QModelIndex & index) { callbackQHelpContentWidget_Expand(this, const_cast<QModelIndex*>(&index)); };
	void expandAll() { callbackQHelpContentWidget_ExpandAll(this); };
	void expandToDepth(int depth) { callbackQHelpContentWidget_ExpandToDepth(this, depth); };
	void Signal_Expanded(const QModelIndex & index) { callbackQHelpContentWidget_Expanded(this, const_cast<QModelIndex*>(&index)); };
	void hideColumn(int column) { callbackQHelpContentWidget_HideColumn(this, column); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpContentWidget_KeyPressEvent(this, event); };
	void keyboardSearch(const QString & search) { QByteArray t3559d7 = search.toUtf8(); QtHelp_PackedString searchPacked = { const_cast<char*>(t3559d7.prepend("WHITESPACE").constData()+10), t3559d7.size()-10 };callbackQHelpContentWidget_KeyboardSearch(this, searchPacked); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpContentWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpContentWidget_MouseReleaseEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpContentWidget_PaintEvent(this, event); };
	void reset() { callbackQHelpContentWidget_Reset(this); };
	void resizeColumnToContents(int column) { callbackQHelpContentWidget_ResizeColumnToContents(this, column); };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), start, end); };
	void rowsInserted(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void rowsRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpContentWidget_RowsRemoved(this, const_cast<QModelIndex*>(&parent), start, end); };
	void scrollContentsBy(int dx, int dy) { callbackQHelpContentWidget_ScrollContentsBy(this, dx, dy); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQHelpContentWidget_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void selectAll() { callbackQHelpContentWidget_SelectAll(this); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQHelpContentWidget_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setModel(QAbstractItemModel * model) { callbackQHelpContentWidget_SetModel(this, model); };
	void setRootIndex(const QModelIndex & index) { callbackQHelpContentWidget_SetRootIndex(this, const_cast<QModelIndex*>(&index)); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpContentWidget_SetSelection(this, const_cast<QRect*>(&rect), command); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpContentWidget_SetSelectionModel(this, selectionModel); };
	void showColumn(int column) { callbackQHelpContentWidget_ShowColumn(this, column); };
	void updateGeometries() { callbackQHelpContentWidget_UpdateGeometries(this); };
	QModelIndex indexAt(const QPoint & point) const { return *static_cast<QModelIndex*>(callbackQHelpContentWidget_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&point))); };
	
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQHelpContentWidget_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQHelpContentWidget_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQHelpContentWidget_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	int horizontalOffset() const { return callbackQHelpContentWidget_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	int sizeHintForColumn(int column) const { return callbackQHelpContentWidget_SizeHintForColumn(const_cast<void*>(static_cast<const void*>(this)), column); };
	int verticalOffset() const { return callbackQHelpContentWidget_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	void drawBranches(QPainter * painter, const QRect & rect, const QModelIndex & index) const { callbackQHelpContentWidget_DrawBranches(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QRect*>(&rect), const_cast<QModelIndex*>(&index)); };
	void drawRow(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackQHelpContentWidget_DrawRow(const_cast<void*>(static_cast<const void*>(this)), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQHelpContentWidget_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQHelpContentWidget_FocusNextPrevChild(this, next) != 0; };
	void Signal_Activated(const QModelIndex & index) { callbackQHelpContentWidget_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackQHelpContentWidget_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackQHelpContentWidget_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackQHelpContentWidget_CloseEditor(this, editor, hint); };
	void commitData(QWidget * editor) { callbackQHelpContentWidget_CommitData(this, editor); };
	void Signal_DoubleClicked(const QModelIndex & index) { callbackQHelpContentWidget_DoubleClicked(this, const_cast<QModelIndex*>(&index)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpContentWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpContentWidget_DragLeaveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQHelpContentWidget_DropEvent(this, event); };
	void edit(const QModelIndex & index) { callbackQHelpContentWidget_Edit(this, const_cast<QModelIndex*>(&index)); };
	void editorDestroyed(QObject * editor) { callbackQHelpContentWidget_EditorDestroyed(this, editor); };
	void Signal_Entered(const QModelIndex & index) { callbackQHelpContentWidget_Entered(this, const_cast<QModelIndex*>(&index)); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpContentWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpContentWidget_FocusOutEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & size) { callbackQHelpContentWidget_IconSizeChanged(this, const_cast<QSize*>(&size)); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpContentWidget_InputMethodEvent(this, event); };
	void Signal_Pressed(const QModelIndex & index) { callbackQHelpContentWidget_Pressed(this, const_cast<QModelIndex*>(&index)); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpContentWidget_ResizeEvent(this, event); };
	void scrollToBottom() { callbackQHelpContentWidget_ScrollToBottom(this); };
	void scrollToTop() { callbackQHelpContentWidget_ScrollToTop(this); };
	void setCurrentIndex(const QModelIndex & index) { callbackQHelpContentWidget_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpContentWidget_StartDrag(this, supportedActions); };
	void update(const QModelIndex & index) { callbackQHelpContentWidget_Update(this, const_cast<QModelIndex*>(&index)); };
	void Signal_ViewportEntered() { callbackQHelpContentWidget_ViewportEntered(this); };
	QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex & index, const QEvent * event) const { return static_cast<QItemSelectionModel::SelectionFlag>(callbackQHelpContentWidget_SelectionCommand(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), const_cast<QEvent*>(event))); };
	QStyleOptionViewItem viewOptions() const { return *static_cast<QStyleOptionViewItem*>(callbackQHelpContentWidget_ViewOptions(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpContentWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	int sizeHintForRow(int row) const { return callbackQHelpContentWidget_SizeHintForRow(const_cast<void*>(static_cast<const void*>(this)), row); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpContentWidget_ContextMenuEvent(this, e); };
	void setupViewport(QWidget * viewport) { callbackQHelpContentWidget_SetupViewport(this, viewport); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpContentWidget_WheelEvent(this, e); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpContentWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void changeEvent(QEvent * ev) { callbackQHelpContentWidget_ChangeEvent(this, ev); };
	bool close() { return callbackQHelpContentWidget_Close(this) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpContentWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQHelpContentWidget_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpContentWidget_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQHelpContentWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQHelpContentWidget_EnterEvent(this, event); };
	void hide() { callbackQHelpContentWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQHelpContentWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpContentWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQHelpContentWidget_LeaveEvent(this, event); };
	void lower() { callbackQHelpContentWidget_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackQHelpContentWidget_MoveEvent(this, event); };
	void raise() { callbackQHelpContentWidget_Raise(this); };
	void repaint() { callbackQHelpContentWidget_Repaint(this); };
	void setDisabled(bool disable) { callbackQHelpContentWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQHelpContentWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQHelpContentWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQHelpContentWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtHelp_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQHelpContentWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQHelpContentWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQHelpContentWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtHelp_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQHelpContentWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQHelpContentWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQHelpContentWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQHelpContentWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQHelpContentWidget_ShowMaximized(this); };
	void showMinimized() { callbackQHelpContentWidget_ShowMinimized(this); };
	void showNormal() { callbackQHelpContentWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpContentWidget_TabletEvent(this, event); };
	void updateMicroFocus() { callbackQHelpContentWidget_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQHelpContentWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtHelp_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQHelpContentWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQHelpContentWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQHelpContentWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQHelpContentWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQHelpContentWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpContentWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpContentWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpContentWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpContentWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpContentWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpContentWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpContentWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpContentWidget_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpContentWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpContentWidget*)

int QHelpContentWidget_QHelpContentWidget_QRegisterMetaType(){qRegisterMetaType<QHelpContentWidget*>(); return qRegisterMetaType<MyQHelpContentWidget*>();}

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

int QHelpContentWidget___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QHelpContentWidget___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QHelpContentWidget___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* QHelpContentWidget___selectedIndexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpContentWidget___selectedIndexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpContentWidget___selectedIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpContentWidget___scrollBarWidgets_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void QHelpContentWidget___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QHelpContentWidget___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>;
}

void* QHelpContentWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpContentWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpContentWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpContentWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpContentWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpContentWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpContentWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpContentWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpContentWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpContentWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpContentWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpContentWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpContentWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpContentWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpContentWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpContentWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpContentWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QHelpContentWidget_MoveCursorDefault(void* ptr, long long cursorAction, long long modifiers)
{
		return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

char QHelpContentWidget_ViewportEventDefault(void* ptr, void* event)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewportEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_CollapseDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::collapse(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_CollapseAllDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::collapseAll();
}

void QHelpContentWidget_ColumnCountChangedDefault(void* ptr, int oldCount, int newCount)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnCountChanged(oldCount, newCount);
}

void QHelpContentWidget_ColumnMovedDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnMoved();
}

void QHelpContentWidget_ColumnResizedDefault(void* ptr, int column, int oldSize, int newSize)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::columnResized(column, oldSize, newSize);
}

void QHelpContentWidget_CurrentChangedDefault(void* ptr, void* current, void* previous)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpContentWidget_DataChangedDefault(void* ptr, void* topLeft, void* bottomRight, void* roles)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dataChanged(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight), *static_cast<QVector<int>*>(roles));
}

void QHelpContentWidget_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpContentWidget_ExpandDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expand(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_ExpandAllDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expandAll();
}

void QHelpContentWidget_ExpandToDepthDefault(void* ptr, int depth)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::expandToDepth(depth);
}

void QHelpContentWidget_HideColumnDefault(void* ptr, int column)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hideColumn(column);
}

void QHelpContentWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyboardSearchDefault(void* ptr, struct QtHelp_PackedString search)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyboardSearch(QString::fromUtf8(search.data, search.len));
}

void QHelpContentWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpContentWidget_ResetDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::reset();
}

void QHelpContentWidget_ResizeColumnToContentsDefault(void* ptr, int column)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::resizeColumnToContents(column);
}

void QHelpContentWidget_RowsAboutToBeRemovedDefault(void* ptr, void* parent, int start, int end)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsInsertedDefault(void* ptr, void* parent, int start, int end)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_RowsRemovedDefault(void* ptr, void* parent, int start, int end)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::rowsRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpContentWidget_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollContentsBy(dx, dy);
}

void QHelpContentWidget_ScrollToDefault(void* ptr, void* index, long long hint)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpContentWidget_SelectAllDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectAll();
}

void QHelpContentWidget_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpContentWidget_SetModelDefault(void* ptr, void* model)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpContentWidget_SetRootIndexDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setRootIndex(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_SetSelectionDefault(void* ptr, void* rect, long long command)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpContentWidget_SetSelectionModelDefault(void* ptr, void* selectionModel)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpContentWidget_ShowColumnDefault(void* ptr, int column)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showColumn(column);
}

void QHelpContentWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentWidget_UpdateGeometriesDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::updateGeometries();
}

void* QHelpContentWidget_IndexAtDefault(void* ptr, void* point)
{
		return new QModelIndex(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::indexAt(*static_cast<QPoint*>(point)));
}



void* QHelpContentWidget_VisualRectDefault(void* ptr, void* index)
{
		return ({ QRect tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
		return new QRegion(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QHelpContentWidget_ViewportSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QHelpContentWidget_IsIndexHiddenDefault(void* ptr, void* index)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::isIndexHidden(*static_cast<QModelIndex*>(index));
}

int QHelpContentWidget_HorizontalOffsetDefault(void* ptr)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::horizontalOffset();
}

int QHelpContentWidget_SizeHintForColumnDefault(void* ptr, int column)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHintForColumn(column);
}

int QHelpContentWidget_VerticalOffsetDefault(void* ptr)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::verticalOffset();
}

void QHelpContentWidget_DrawBranchesDefault(void* ptr, void* painter, void* rect, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::drawBranches(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), *static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_DrawRowDefault(void* ptr, void* painter, void* option, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::drawRow(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char QHelpContentWidget_Edit2Default(void* ptr, void* index, long long trigger, void* event)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

char QHelpContentWidget_EventDefault(void* ptr, void* event)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::event(static_cast<QEvent*>(event));
}

char QHelpContentWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusNextPrevChild(next != 0);
}

void QHelpContentWidget_ClearSelectionDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::clearSelection();
}

void QHelpContentWidget_CloseEditorDefault(void* ptr, void* editor, long long hint)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::closeEditor(static_cast<QWidget*>(editor), static_cast<QAbstractItemDelegate::EndEditHint>(hint));
}

void QHelpContentWidget_CommitDataDefault(void* ptr, void* editor)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::commitData(static_cast<QWidget*>(editor));
}

void QHelpContentWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpContentWidget_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpContentWidget_DropEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpContentWidget_EditDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::edit(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_EditorDestroyedDefault(void* ptr, void* editor)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::editorDestroyed(static_cast<QObject*>(editor));
}

void QHelpContentWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpContentWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpContentWidget_ScrollToBottomDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollToBottom();
}

void QHelpContentWidget_ScrollToTopDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollToTop();
}

void QHelpContentWidget_SetCurrentIndexDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setCurrentIndex(*static_cast<QModelIndex*>(index));
}

void QHelpContentWidget_StartDragDefault(void* ptr, long long supportedActions)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpContentWidget_UpdateDefault(void* ptr, void* index)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::update(*static_cast<QModelIndex*>(index));
}

long long QHelpContentWidget_SelectionCommandDefault(void* ptr, void* index, void* event)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

void* QHelpContentWidget_ViewOptionsDefault(void* ptr)
{
		return new QStyleOptionViewItem(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::viewOptions());
}

void* QHelpContentWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int QHelpContentWidget_SizeHintForRowDefault(void* ptr, int row)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHintForRow(row);
}

void QHelpContentWidget_ContextMenuEventDefault(void* ptr, void* e)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpContentWidget_SetupViewportDefault(void* ptr, void* viewport)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpContentWidget_WheelEventDefault(void* ptr, void* e)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void* QHelpContentWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpContentWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpContentWidget_ChangeEventDefault(void* ptr, void* ev)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::changeEvent(static_cast<QEvent*>(ev));
}

char QHelpContentWidget_CloseDefault(void* ptr)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::close();
}

char QHelpContentWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QHelpContentWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpContentWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpContentWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_HideDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hide();
}

void QHelpContentWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpContentWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_LowerDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::lower();
}

void QHelpContentWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpContentWidget_RaiseDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::raise();
}

void QHelpContentWidget_RepaintDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::repaint();
}

void QHelpContentWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setDisabled(disable != 0);
}

void QHelpContentWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setEnabled(vbo != 0);
}

void QHelpContentWidget_SetFocus2Default(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setFocus();
}

void QHelpContentWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setHidden(hidden != 0);
}

void QHelpContentWidget_SetStyleSheetDefault(void* ptr, struct QtHelp_PackedString styleSheet)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QHelpContentWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setVisible(visible != 0);
}

void QHelpContentWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setWindowModified(vbo != 0);
}

void QHelpContentWidget_SetWindowTitleDefault(void* ptr, struct QtHelp_PackedString vqs)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QHelpContentWidget_ShowDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::show();
}

void QHelpContentWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpContentWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showFullScreen();
}

void QHelpContentWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showMaximized();
}

void QHelpContentWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showMinimized();
}

void QHelpContentWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showNormal();
}

void QHelpContentWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpContentWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::updateMicroFocus();
}

void* QHelpContentWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::paintEngine();
}

char QHelpContentWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hasHeightForWidth();
}

int QHelpContentWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::heightForWidth(w);
}

int QHelpContentWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QHelpContentWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpContentWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpContentWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::deleteLater();
}

void QHelpContentWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void* QHelpContentWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::metaObject());
}

class MyQHelpEngine: public QHelpEngine
{
public:
	MyQHelpEngine(const QString &collectionFile, QObject *parent = Q_NULLPTR) : QHelpEngine(collectionFile, parent) {QHelpEngine_QHelpEngine_QRegisterMetaType();};
	void Signal_CurrentFilterChanged(const QString & newFilter) { QByteArray t56548b = newFilter.toUtf8(); QtHelp_PackedString newFilterPacked = { const_cast<char*>(t56548b.prepend("WHITESPACE").constData()+10), t56548b.size()-10 };callbackQHelpEngineCore_CurrentFilterChanged(this, newFilterPacked); };
	void Signal_ReadersAboutToBeInvalidated() { callbackQHelpEngineCore_ReadersAboutToBeInvalidated(this); };
	void Signal_SetupFinished() { callbackQHelpEngineCore_SetupFinished(this); };
	void Signal_SetupStarted() { callbackQHelpEngineCore_SetupStarted(this); };
	void Signal_Warning(const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtHelp_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQHelpEngineCore_Warning(this, msgPacked); };
	bool event(QEvent * e) { return callbackQHelpEngineCore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpEngineCore_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpEngineCore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpEngineCore_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpEngineCore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpEngineCore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpEngineCore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpEngineCore_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpEngineCore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpEngine*)

int QHelpEngine_QHelpEngine_QRegisterMetaType(){qRegisterMetaType<QHelpEngine*>(); return qRegisterMetaType<MyQHelpEngine*>();}

void* QHelpEngine_ContentWidget(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_NewQHelpEngine(struct QtHelp_PackedString collectionFile, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQHelpEngine(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QObject*>(parent));
	}
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

void* QHelpEngine_ContentModel(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_IndexModel(void* ptr)
{
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

class MyQHelpEngineCore: public QHelpEngineCore
{
public:
	MyQHelpEngineCore(const QString &collectionFile, QObject *parent = Q_NULLPTR) : QHelpEngineCore(collectionFile, parent) {QHelpEngineCore_QHelpEngineCore_QRegisterMetaType();};
	void Signal_CurrentFilterChanged(const QString & newFilter) { QByteArray t56548b = newFilter.toUtf8(); QtHelp_PackedString newFilterPacked = { const_cast<char*>(t56548b.prepend("WHITESPACE").constData()+10), t56548b.size()-10 };callbackQHelpEngineCore_CurrentFilterChanged(this, newFilterPacked); };
	void Signal_ReadersAboutToBeInvalidated() { callbackQHelpEngineCore_ReadersAboutToBeInvalidated(this); };
	void Signal_SetupFinished() { callbackQHelpEngineCore_SetupFinished(this); };
	void Signal_SetupStarted() { callbackQHelpEngineCore_SetupStarted(this); };
	void Signal_Warning(const QString & msg) { QByteArray t19f34e = msg.toUtf8(); QtHelp_PackedString msgPacked = { const_cast<char*>(t19f34e.prepend("WHITESPACE").constData()+10), t19f34e.size()-10 };callbackQHelpEngineCore_Warning(this, msgPacked); };
	 ~MyQHelpEngineCore() { callbackQHelpEngineCore_DestroyQHelpEngineCore(this); };
	bool event(QEvent * e) { return callbackQHelpEngineCore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpEngineCore_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpEngineCore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpEngineCore_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpEngineCore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpEngineCore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpEngineCore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpEngineCore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpEngineCore_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpEngineCore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpEngineCore*)

int QHelpEngineCore_QHelpEngineCore_QRegisterMetaType(){qRegisterMetaType<QHelpEngineCore*>(); return qRegisterMetaType<MyQHelpEngineCore*>();}

void* QHelpEngineCore_NewQHelpEngineCore(struct QtHelp_PackedString collectionFile, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQHelpEngineCore(QString::fromUtf8(collectionFile.data, collectionFile.len), static_cast<QObject*>(parent));
	}
}

struct QtHelp_PackedList QHelpEngineCore_Files(void* ptr, struct QtHelp_PackedString namespaceName, struct QtHelp_PackedString filterAttributes, struct QtHelp_PackedString extensionFilter)
{
	return ({ QList<QUrl>* tmpValue = new QList<QUrl>(static_cast<QHelpEngineCore*>(ptr)->files(QString::fromUtf8(namespaceName.data, namespaceName.len), QString::fromUtf8(filterAttributes.data, filterAttributes.len).split("|", QString::SkipEmptyParts), QString::fromUtf8(extensionFilter.data, extensionFilter.len))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedString QHelpEngineCore_DocumentationFileName(void* ptr, struct QtHelp_PackedString namespaceName)
{
	return ({ QByteArray t454496 = static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString::fromUtf8(namespaceName.data, namespaceName.len)).toUtf8(); QtHelp_PackedString { const_cast<char*>(t454496.prepend("WHITESPACE").constData()+10), t454496.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_QHelpEngineCore_NamespaceName(struct QtHelp_PackedString documentationFileName)
{
	return ({ QByteArray te8bcfa = QHelpEngineCore::namespaceName(QString::fromUtf8(documentationFileName.data, documentationFileName.len)).toUtf8(); QtHelp_PackedString { const_cast<char*>(te8bcfa.prepend("WHITESPACE").constData()+10), te8bcfa.size()-10 }; });
}

void* QHelpEngineCore_QHelpEngineCore_MetaData(struct QtHelp_PackedString documentationFileName, struct QtHelp_PackedString name)
{
	return new QVariant(QHelpEngineCore::metaData(QString::fromUtf8(documentationFileName.data, documentationFileName.len), QString::fromUtf8(name.data, name.len)));
}

char QHelpEngineCore_AddCustomFilter(void* ptr, struct QtHelp_PackedString filterName, struct QtHelp_PackedString attributes)
{
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString::fromUtf8(filterName.data, filterName.len), QString::fromUtf8(attributes.data, attributes.len).split("|", QString::SkipEmptyParts));
}

char QHelpEngineCore_CopyCollectionFile(void* ptr, struct QtHelp_PackedString fileName)
{
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString::fromUtf8(fileName.data, fileName.len));
}

char QHelpEngineCore_RegisterDocumentation(void* ptr, struct QtHelp_PackedString documentationFileName)
{
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString::fromUtf8(documentationFileName.data, documentationFileName.len));
}

char QHelpEngineCore_RemoveCustomFilter(void* ptr, struct QtHelp_PackedString filterName)
{
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString::fromUtf8(filterName.data, filterName.len));
}

char QHelpEngineCore_RemoveCustomValue(void* ptr, struct QtHelp_PackedString key)
{
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString::fromUtf8(key.data, key.len));
}

char QHelpEngineCore_SetCustomValue(void* ptr, struct QtHelp_PackedString key, void* value)
{
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(value));
}

char QHelpEngineCore_SetupData(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

char QHelpEngineCore_UnregisterDocumentation(void* ptr, struct QtHelp_PackedString namespaceName)
{
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString::fromUtf8(namespaceName.data, namespaceName.len));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));
}

void QHelpEngineCore_CurrentFilterChanged(void* ptr, struct QtHelp_PackedString newFilter)
{
	static_cast<QHelpEngineCore*>(ptr)->currentFilterChanged(QString::fromUtf8(newFilter.data, newFilter.len));
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

void QHelpEngineCore_SetAutoSaveFilter(void* ptr, char save)
{
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(void* ptr, struct QtHelp_PackedString fileName)
{
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString::fromUtf8(fileName.data, fileName.len));
}

void QHelpEngineCore_SetCurrentFilter(void* ptr, struct QtHelp_PackedString filterName)
{
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString::fromUtf8(filterName.data, filterName.len));
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

void QHelpEngineCore_ConnectWarning(void* ptr)
{
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));
}

void QHelpEngineCore_DisconnectWarning(void* ptr)
{
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));
}

void QHelpEngineCore_Warning(void* ptr, struct QtHelp_PackedString msg)
{
	static_cast<QHelpEngineCore*>(ptr)->warning(QString::fromUtf8(msg.data, msg.len));
}

void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr)
{
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

void QHelpEngineCore_DestroyQHelpEngineCoreDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QHelpEngineCore_FileData(void* ptr, void* url)
{
	return new QByteArray(static_cast<QHelpEngineCore*>(ptr)->fileData(*static_cast<QUrl*>(url)));
}

struct QtHelp_PackedList QHelpEngineCore_LinksForIdentifier(void* ptr, struct QtHelp_PackedString id)
{
	return ({ QMap<QString, QUrl>* tmpValue = new QMap<QString, QUrl>(static_cast<QHelpEngineCore*>(ptr)->linksForIdentifier(QString::fromUtf8(id.data, id.len))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedList QHelpEngineCore_LinksForKeyword(void* ptr, struct QtHelp_PackedString keyword)
{
	return ({ QMap<QString, QUrl>* tmpValue = new QMap<QString, QUrl>(static_cast<QHelpEngineCore*>(ptr)->linksForKeyword(QString::fromUtf8(keyword.data, keyword.len))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedString QHelpEngineCore_CollectionFile(void* ptr)
{
	return ({ QByteArray tbb1d46 = static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8(); QtHelp_PackedString { const_cast<char*>(tbb1d46.prepend("WHITESPACE").constData()+10), tbb1d46.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_CurrentFilter(void* ptr)
{
	return ({ QByteArray taf4fbf = static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8(); QtHelp_PackedString { const_cast<char*>(taf4fbf.prepend("WHITESPACE").constData()+10), taf4fbf.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_Error(void* ptr)
{
	return ({ QByteArray t952e50 = static_cast<QHelpEngineCore*>(ptr)->error().toUtf8(); QtHelp_PackedString { const_cast<char*>(t952e50.prepend("WHITESPACE").constData()+10), t952e50.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_CustomFilters(void* ptr)
{
	return ({ QByteArray t2b7197 = static_cast<QHelpEngineCore*>(ptr)->customFilters().join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t2b7197.prepend("WHITESPACE").constData()+10), t2b7197.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_FilterAttributes(void* ptr)
{
	return ({ QByteArray t22f9b0 = static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t22f9b0.prepend("WHITESPACE").constData()+10), t22f9b0.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_FilterAttributes2(void* ptr, struct QtHelp_PackedString filterName)
{
	return ({ QByteArray t8bd21d = static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString::fromUtf8(filterName.data, filterName.len)).join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t8bd21d.prepend("WHITESPACE").constData()+10), t8bd21d.size()-10 }; });
}

struct QtHelp_PackedString QHelpEngineCore_RegisteredDocumentations(void* ptr)
{
	return ({ QByteArray t7c5252 = static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t7c5252.prepend("WHITESPACE").constData()+10), t7c5252.size()-10 }; });
}

void* QHelpEngineCore_FindFile(void* ptr, void* url)
{
	return new QUrl(static_cast<QHelpEngineCore*>(ptr)->findFile(*static_cast<QUrl*>(url)));
}

void* QHelpEngineCore_CustomValue(void* ptr, struct QtHelp_PackedString key, void* defaultValue)
{
	return new QVariant(static_cast<QHelpEngineCore*>(ptr)->customValue(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(defaultValue)));
}

char QHelpEngineCore_AutoSaveFilter(void* ptr)
{
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

void* QHelpEngineCore___files_atList(void* ptr, int i)
{
	return new QUrl(static_cast<QList<QUrl>*>(ptr)->at(i));
}

void QHelpEngineCore___files_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QHelpEngineCore___files_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>;
}

struct QtHelp_PackedString QHelpEngineCore___filterAttributeSets_atList(void* ptr, int i)
{
	return ({ QByteArray t5372d5 = static_cast<QList<QStringList>*>(ptr)->at(i).join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t5372d5.prepend("WHITESPACE").constData()+10), t5372d5.size()-10 }; });
}

void QHelpEngineCore___filterAttributeSets_setList(void* ptr, struct QtHelp_PackedString i)
{
	static_cast<QList<QStringList>*>(ptr)->append(QString::fromUtf8(i.data, i.len).split("|", QString::SkipEmptyParts));
}

void* QHelpEngineCore___filterAttributeSets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QStringList>;
}

void* QHelpEngineCore___linksForIdentifier_atList(void* ptr, struct QtHelp_PackedString i)
{
	return new QUrl(static_cast<QMap<QString, QUrl>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QHelpEngineCore___linksForIdentifier_setList(void* ptr, struct QtHelp_PackedString key, void* i)
{
	static_cast<QMap<QString, QUrl>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QUrl*>(i));
}

void* QHelpEngineCore___linksForIdentifier_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QUrl>;
}

struct QtHelp_PackedList QHelpEngineCore___linksForIdentifier_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QUrl>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpEngineCore___linksForKeyword_atList(void* ptr, struct QtHelp_PackedString i)
{
	return new QUrl(static_cast<QMap<QString, QUrl>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QHelpEngineCore___linksForKeyword_setList(void* ptr, struct QtHelp_PackedString key, void* i)
{
	static_cast<QMap<QString, QUrl>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QUrl*>(i));
}

void* QHelpEngineCore___linksForKeyword_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QUrl>;
}

struct QtHelp_PackedList QHelpEngineCore___linksForKeyword_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QUrl>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedString QHelpEngineCore_____linksForIdentifier_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtHelp_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QHelpEngineCore_____linksForIdentifier_keyList_setList(void* ptr, struct QtHelp_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QHelpEngineCore_____linksForIdentifier_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtHelp_PackedString QHelpEngineCore_____linksForKeyword_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtHelp_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QHelpEngineCore_____linksForKeyword_keyList_setList(void* ptr, struct QtHelp_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QHelpEngineCore_____linksForKeyword_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QHelpEngineCore___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpEngineCore___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpEngineCore___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpEngineCore___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpEngineCore___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpEngineCore___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpEngineCore___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpEngineCore___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpEngineCore___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpEngineCore___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpEngineCore___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpEngineCore___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpEngineCore___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpEngineCore___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpEngineCore___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpEngineCore_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHelpEngine*>(ptr)->QHelpEngine::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::event(static_cast<QEvent*>(e));
	}
}

char QHelpEngineCore_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHelpEngine*>(ptr)->QHelpEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QHelpEngineCore_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QHelpEngineCore_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QHelpEngineCore_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::customEvent(static_cast<QEvent*>(event));
	}
}

void QHelpEngineCore_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::deleteLater();
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::deleteLater();
	}
}

void QHelpEngineCore_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QHelpEngineCore_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QHelpEngine*>(ptr)->QHelpEngine::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QHelpEngineCore_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QHelpEngine*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHelpEngine*>(ptr)->QHelpEngine::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::metaObject());
	}
}

class MyQHelpIndexModel: public QHelpIndexModel
{
public:
	void Signal_IndexCreated() { callbackQHelpIndexModel_IndexCreated(this); };
	void Signal_IndexCreationStarted() { callbackQHelpIndexModel_IndexCreationStarted(this); };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQHelpIndexModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQHelpIndexModel_Sort(this, column, order); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQHelpIndexModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQHelpIndexModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQHelpIndexModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	int rowCount(const QModelIndex & parent) const { return callbackQHelpIndexModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQHelpIndexModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpIndexModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQHelpIndexModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQHelpIndexModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQHelpIndexModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQHelpIndexModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = const_cast<QMap<int, QVariant>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	bool submit() { return callbackQHelpIndexModel_Submit(this) != 0; };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQHelpIndexModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQHelpIndexModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQHelpIndexModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackQHelpIndexModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQHelpIndexModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQHelpIndexModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); QtHelp_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQHelpIndexModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = const_cast<QList<QPersistentModelIndex>*>(&parents); QtHelp_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackQHelpIndexModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQHelpIndexModel_ModelReset(this); };
	void resetInternalData() { callbackQHelpIndexModel_ResetInternalData(this); };
	void revert() { callbackQHelpIndexModel_Revert(this); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQHelpIndexModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQHelpIndexModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQHelpIndexModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQHelpIndexModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QHash<int, QByteArray> roleNames() const { return *static_cast<QHash<int, QByteArray>*>(callbackQHelpIndexModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return *static_cast<QMap<int, QVariant>*>(callbackQHelpIndexModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQHelpIndexModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); QtHelp_PackedList { tmpValue, tmpValue->size() }; }))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQHelpIndexModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return *static_cast<QList<QModelIndex>*>(callbackQHelpIndexModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQHelpIndexModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QStringList mimeTypes() const { return ({ QtHelp_PackedString tempVal = callbackQHelpIndexModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQHelpIndexModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQHelpIndexModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQHelpIndexModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQHelpIndexModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	bool hasChildren(const QModelIndex & parent) const { return callbackQHelpIndexModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQHelpIndexModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool event(QEvent * e) { return callbackQHelpIndexModel_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpIndexModel_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpIndexModel_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpIndexModel_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpIndexModel_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpIndexModel_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpIndexModel_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpIndexModel_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpIndexModel_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpIndexModel_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpIndexModel_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpIndexModel*)

int QHelpIndexModel_QHelpIndexModel_QRegisterMetaType(){qRegisterMetaType<QHelpIndexModel*>(); return qRegisterMetaType<MyQHelpIndexModel*>();}

void* QHelpIndexModel_Filter(void* ptr, struct QtHelp_PackedString filter, struct QtHelp_PackedString wildcard)
{
	return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->filter(QString::fromUtf8(filter.data, filter.len), QString::fromUtf8(wildcard.data, wildcard.len)));
}

void QHelpIndexModel_CreateIndex(void* ptr, struct QtHelp_PackedString customFilterName)
{
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString::fromUtf8(customFilterName.data, customFilterName.len));
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

char QHelpIndexModel_IsCreatingIndex(void* ptr)
{
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

void* QHelpIndexModel___linksForKeyword_atList(void* ptr, struct QtHelp_PackedString i)
{
	return new QUrl(static_cast<QMap<QString, QUrl>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QHelpIndexModel___linksForKeyword_setList(void* ptr, struct QtHelp_PackedString key, void* i)
{
	static_cast<QMap<QString, QUrl>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QUrl*>(i));
}

void* QHelpIndexModel___linksForKeyword_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QUrl>;
}

struct QtHelp_PackedList QHelpIndexModel___linksForKeyword_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QUrl>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedString QHelpIndexModel_____linksForKeyword_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtHelp_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QHelpIndexModel_____linksForKeyword_keyList_setList(void* ptr, struct QtHelp_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QHelpIndexModel_____linksForKeyword_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

int QHelpIndexModel_____setItemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpIndexModel_____setItemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpIndexModel_____setItemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpIndexModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpIndexModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpIndexModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpIndexModel_____itemData_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpIndexModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpIndexModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* QHelpIndexModel___setItemData_roles_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QHelpIndexModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QHelpIndexModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct QtHelp_PackedList QHelpIndexModel___setItemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpIndexModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int QHelpIndexModel___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QHelpIndexModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QHelpIndexModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* QHelpIndexModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QHelpIndexModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* QHelpIndexModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QHelpIndexModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>;
}

void* QHelpIndexModel___roleNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QHash<int, QByteArray>*>(ptr)->value(i));
}

void QHelpIndexModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QHelpIndexModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>;
}

struct QtHelp_PackedList QHelpIndexModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexModel___itemData_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QMap<int, QVariant>*>(ptr)->value(i));
}

void QHelpIndexModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QHelpIndexModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>;
}

struct QtHelp_PackedList QHelpIndexModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpIndexModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpIndexModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

int QHelpIndexModel_____doSetRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpIndexModel_____doSetRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpIndexModel_____doSetRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

int QHelpIndexModel_____setRoleNames_keyList_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QHelpIndexModel_____setRoleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QHelpIndexModel_____setRoleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* QHelpIndexModel___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpIndexModel___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpIndexModel___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpIndexModel___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexModel___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexModel___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexModel___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexModel___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexModel___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexModel___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexModel___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexModel___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexModel___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpIndexModel___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexModel___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpIndexModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

void QHelpIndexModel_SortDefault(void* ptr, int column, long long order)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QHelpIndexModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
		return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QHelpIndexModel_DataDefault(void* ptr, void* index, int role)
{
		return new QVariant(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::data(*static_cast<QModelIndex*>(index), role));
}

long long QHelpIndexModel_SupportedDropActionsDefault(void* ptr)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::supportedDropActions();
}

long long QHelpIndexModel_FlagsDefault(void* ptr, void* index)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::flags(*static_cast<QModelIndex*>(index));
}

int QHelpIndexModel_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::rowCount(*static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QHelpIndexModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QHelpIndexModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QHelpIndexModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QHelpIndexModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QHelpIndexModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QHelpIndexModel_SubmitDefault(void* ptr)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::submit();
}

void QHelpIndexModel_FetchMoreDefault(void* ptr, void* parent)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

void QHelpIndexModel_ResetInternalDataDefault(void* ptr)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::resetInternalData();
}

void QHelpIndexModel_RevertDefault(void* ptr)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::revert();
}

struct QtHelp_PackedList QHelpIndexModel_RoleNamesDefault(void* ptr)
{
		return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::roleNames()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedList QHelpIndexModel_ItemDataDefault(void* ptr, void* index)
{
		return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::itemData(*static_cast<QModelIndex*>(index))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexModel_MimeDataDefault(void* ptr, void* indexes)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::mimeData(*static_cast<QList<QModelIndex>*>(indexes));
}

void* QHelpIndexModel_BuddyDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::buddy(*static_cast<QModelIndex*>(index)));
}

void* QHelpIndexModel_ParentDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::parent(*static_cast<QModelIndex*>(index)));
}

struct QtHelp_PackedList QHelpIndexModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexModel_SpanDefault(void* ptr, void* index)
{
		return ({ QSize tmpValue = static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtHelp_PackedString QHelpIndexModel_MimeTypesDefault(void* ptr)
{
		return ({ QByteArray t6c8862 = static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::mimeTypes().join("|").toUtf8(); QtHelp_PackedString { const_cast<char*>(t6c8862.prepend("WHITESPACE").constData()+10), t6c8862.size()-10 }; });
}

void* QHelpIndexModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
		return new QVariant(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

long long QHelpIndexModel_SupportedDragActionsDefault(void* ptr)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::supportedDragActions();
}

char QHelpIndexModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_CanFetchMoreDefault(void* ptr, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

int QHelpIndexModel_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::columnCount(*static_cast<QModelIndex*>(parent));
}

char QHelpIndexModel_EventDefault(void* ptr, void* e)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::event(static_cast<QEvent*>(e));
}

char QHelpIndexModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpIndexModel_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexModel_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexModel_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexModel_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::deleteLater();
}

void QHelpIndexModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexModel_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHelpIndexModel_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::metaObject());
}

class MyQHelpIndexWidget: public QHelpIndexWidget
{
public:
	void activateCurrentItem() { callbackQHelpIndexWidget_ActivateCurrentItem(this); };
	void filterIndices(const QString & filter, const QString & wildcard) { QByteArray t4bb4ca = filter.toUtf8(); QtHelp_PackedString filterPacked = { const_cast<char*>(t4bb4ca.prepend("WHITESPACE").constData()+10), t4bb4ca.size()-10 };QByteArray t08654e = wildcard.toUtf8(); QtHelp_PackedString wildcardPacked = { const_cast<char*>(t08654e.prepend("WHITESPACE").constData()+10), t08654e.size()-10 };callbackQHelpIndexWidget_FilterIndices(this, filterPacked, wildcardPacked); };
	void Signal_LinkActivated(const QUrl & link, const QString & keyword) { QByteArray ta43aa2 = keyword.toUtf8(); QtHelp_PackedString keywordPacked = { const_cast<char*>(ta43aa2.prepend("WHITESPACE").constData()+10), ta43aa2.size()-10 };callbackQHelpIndexWidget_LinkActivated(this, const_cast<QUrl*>(&link), keywordPacked); };
	void Signal_LinksActivated(const QMap<QString, QUrl> & links, const QString & keyword) { QByteArray ta43aa2 = keyword.toUtf8(); QtHelp_PackedString keywordPacked = { const_cast<char*>(ta43aa2.prepend("WHITESPACE").constData()+10), ta43aa2.size()-10 };callbackQHelpIndexWidget_LinksActivated(this, ({ QMap<QString, QUrl>* tmpValue = const_cast<QMap<QString, QUrl>*>(&links); QtHelp_PackedList { tmpValue, tmpValue->size() }; }), keywordPacked); };
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, Qt::KeyboardModifiers modifiers) { return *static_cast<QModelIndex*>(callbackQHelpIndexWidget_MoveCursor(this, cursorAction, modifiers)); };
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { callbackQHelpIndexWidget_CurrentChanged(this, const_cast<QModelIndex*>(&current), const_cast<QModelIndex*>(&previous)); };
	void dataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQHelpIndexWidget_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = const_cast<QVector<int>*>(&roles); QtHelp_PackedList { tmpValue, tmpValue->size() }; })); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQHelpIndexWidget_DragLeaveEvent(this, e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQHelpIndexWidget_DragMoveEvent(this, e); };
	void dropEvent(QDropEvent * e) { callbackQHelpIndexWidget_DropEvent(this, e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQHelpIndexWidget_MouseMoveEvent(this, e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQHelpIndexWidget_MouseReleaseEvent(this, e); };
	void paintEvent(QPaintEvent * e) { callbackQHelpIndexWidget_PaintEvent(this, e); };
	void resizeEvent(QResizeEvent * e) { callbackQHelpIndexWidget_ResizeEvent(this, e); };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { callbackQHelpIndexWidget_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), start, end); };
	void rowsInserted(const QModelIndex & parent, int start, int end) { callbackQHelpIndexWidget_RowsInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { callbackQHelpIndexWidget_ScrollTo(this, const_cast<QModelIndex*>(&index), hint); };
	void selectionChanged(const QItemSelection & selected, const QItemSelection & deselected) { callbackQHelpIndexWidget_SelectionChanged(this, const_cast<QItemSelection*>(&selected), const_cast<QItemSelection*>(&deselected)); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpIndexWidget_SetSelection(this, const_cast<QRect*>(&rect), command); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpIndexWidget_StartDrag(this, supportedActions); };
	void updateGeometries() { callbackQHelpIndexWidget_UpdateGeometries(this); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpIndexWidget_WheelEvent(this, e); };
	QModelIndex indexAt(const QPoint & p) const { return *static_cast<QModelIndex*>(callbackQHelpIndexWidget_IndexAt(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPoint*>(&p))); };
	QList<QModelIndex> selectedIndexes() const { return *static_cast<QList<QModelIndex>*>(callbackQHelpIndexWidget_SelectedIndexes(const_cast<void*>(static_cast<const void*>(this)))); };
	QRect visualRect(const QModelIndex & index) const { return *static_cast<QRect*>(callbackQHelpIndexWidget_VisualRect(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QRegion visualRegionForSelection(const QItemSelection & selection) const { return *static_cast<QRegion*>(callbackQHelpIndexWidget_VisualRegionForSelection(const_cast<void*>(static_cast<const void*>(this)), const_cast<QItemSelection*>(&selection))); };
	QSize viewportSizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_ViewportSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QStyleOptionViewItem viewOptions() const { return *static_cast<QStyleOptionViewItem*>(callbackQHelpIndexWidget_ViewOptions(const_cast<void*>(static_cast<const void*>(this)))); };
	bool isIndexHidden(const QModelIndex & index) const { return callbackQHelpIndexWidget_IsIndexHidden(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index)) != 0; };
	int horizontalOffset() const { return callbackQHelpIndexWidget_HorizontalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	int verticalOffset() const { return callbackQHelpIndexWidget_VerticalOffset(const_cast<void*>(static_cast<const void*>(this))); };
	bool edit(const QModelIndex & index, QAbstractItemView::EditTrigger trigger, QEvent * event) { return callbackQHelpIndexWidget_Edit2(this, const_cast<QModelIndex*>(&index), trigger, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQHelpIndexWidget_FocusNextPrevChild(this, next) != 0; };
	bool viewportEvent(QEvent * event) { return callbackQHelpIndexWidget_ViewportEvent(this, event) != 0; };
	void Signal_Activated(const QModelIndex & index) { callbackQHelpIndexWidget_Activated(this, const_cast<QModelIndex*>(&index)); };
	void clearSelection() { callbackQHelpIndexWidget_ClearSelection(this); };
	void Signal_Clicked(const QModelIndex & index) { callbackQHelpIndexWidget_Clicked(this, const_cast<QModelIndex*>(&index)); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { callbackQHelpIndexWidget_CloseEditor(this, editor, hint); };
	void commitData(QWidget * editor) { callbackQHelpIndexWidget_CommitData(this, editor); };
	void Signal_DoubleClicked(const QModelIndex & index) { callbackQHelpIndexWidget_DoubleClicked(this, const_cast<QModelIndex*>(&index)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpIndexWidget_DragEnterEvent(this, event); };
	void edit(const QModelIndex & index) { callbackQHelpIndexWidget_Edit(this, const_cast<QModelIndex*>(&index)); };
	void editorDestroyed(QObject * editor) { callbackQHelpIndexWidget_EditorDestroyed(this, editor); };
	void Signal_Entered(const QModelIndex & index) { callbackQHelpIndexWidget_Entered(this, const_cast<QModelIndex*>(&index)); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpIndexWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpIndexWidget_FocusOutEvent(this, event); };
	void Signal_IconSizeChanged(const QSize & size) { callbackQHelpIndexWidget_IconSizeChanged(this, const_cast<QSize*>(&size)); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpIndexWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpIndexWidget_KeyPressEvent(this, event); };
	void keyboardSearch(const QString & search) { QByteArray t3559d7 = search.toUtf8(); QtHelp_PackedString searchPacked = { const_cast<char*>(t3559d7.prepend("WHITESPACE").constData()+10), t3559d7.size()-10 };callbackQHelpIndexWidget_KeyboardSearch(this, searchPacked); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpIndexWidget_MouseDoubleClickEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpIndexWidget_MousePressEvent(this, event); };
	void Signal_Pressed(const QModelIndex & index) { callbackQHelpIndexWidget_Pressed(this, const_cast<QModelIndex*>(&index)); };
	void reset() { callbackQHelpIndexWidget_Reset(this); };
	void scrollToBottom() { callbackQHelpIndexWidget_ScrollToBottom(this); };
	void scrollToTop() { callbackQHelpIndexWidget_ScrollToTop(this); };
	void selectAll() { callbackQHelpIndexWidget_SelectAll(this); };
	void setCurrentIndex(const QModelIndex & index) { callbackQHelpIndexWidget_SetCurrentIndex(this, const_cast<QModelIndex*>(&index)); };
	void setModel(QAbstractItemModel * model) { callbackQHelpIndexWidget_SetModel(this, model); };
	void setRootIndex(const QModelIndex & index) { callbackQHelpIndexWidget_SetRootIndex(this, const_cast<QModelIndex*>(&index)); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpIndexWidget_SetSelectionModel(this, selectionModel); };
	void update(const QModelIndex & index) { callbackQHelpIndexWidget_Update(this, const_cast<QModelIndex*>(&index)); };
	void Signal_ViewportEntered() { callbackQHelpIndexWidget_ViewportEntered(this); };
	QItemSelectionModel::SelectionFlags selectionCommand(const QModelIndex & index, const QEvent * event) const { return static_cast<QItemSelectionModel::SelectionFlag>(callbackQHelpIndexWidget_SelectionCommand(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), const_cast<QEvent*>(event))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpIndexWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	int sizeHintForColumn(int column) const { return callbackQHelpIndexWidget_SizeHintForColumn(const_cast<void*>(static_cast<const void*>(this)), column); };
	int sizeHintForRow(int row) const { return callbackQHelpIndexWidget_SizeHintForRow(const_cast<void*>(static_cast<const void*>(this)), row); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpIndexWidget_ContextMenuEvent(this, e); };
	void scrollContentsBy(int dx, int dy) { callbackQHelpIndexWidget_ScrollContentsBy(this, dx, dy); };
	void setupViewport(QWidget * viewport) { callbackQHelpIndexWidget_SetupViewport(this, viewport); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpIndexWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	void changeEvent(QEvent * ev) { callbackQHelpIndexWidget_ChangeEvent(this, ev); };
	bool close() { return callbackQHelpIndexWidget_Close(this) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpIndexWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQHelpIndexWidget_ActionEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpIndexWidget_CloseEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQHelpIndexWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void enterEvent(QEvent * event) { callbackQHelpIndexWidget_EnterEvent(this, event); };
	void hide() { callbackQHelpIndexWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQHelpIndexWidget_HideEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpIndexWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQHelpIndexWidget_LeaveEvent(this, event); };
	void lower() { callbackQHelpIndexWidget_Lower(this); };
	void moveEvent(QMoveEvent * event) { callbackQHelpIndexWidget_MoveEvent(this, event); };
	void raise() { callbackQHelpIndexWidget_Raise(this); };
	void repaint() { callbackQHelpIndexWidget_Repaint(this); };
	void setDisabled(bool disable) { callbackQHelpIndexWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQHelpIndexWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQHelpIndexWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQHelpIndexWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtHelp_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQHelpIndexWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQHelpIndexWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQHelpIndexWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtHelp_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQHelpIndexWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQHelpIndexWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQHelpIndexWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQHelpIndexWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQHelpIndexWidget_ShowMaximized(this); };
	void showMinimized() { callbackQHelpIndexWidget_ShowMinimized(this); };
	void showNormal() { callbackQHelpIndexWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpIndexWidget_TabletEvent(this, event); };
	void updateMicroFocus() { callbackQHelpIndexWidget_UpdateMicroFocus(this); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQHelpIndexWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtHelp_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQHelpIndexWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQHelpIndexWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	bool hasHeightForWidth() const { return callbackQHelpIndexWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQHelpIndexWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQHelpIndexWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpIndexWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpIndexWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpIndexWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpIndexWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpIndexWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpIndexWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpIndexWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpIndexWidget_ObjectNameChanged(this, objectNamePacked); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpIndexWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpIndexWidget*)

int QHelpIndexWidget_QHelpIndexWidget_QRegisterMetaType(){qRegisterMetaType<QHelpIndexWidget*>(); return qRegisterMetaType<MyQHelpIndexWidget*>();}

void QHelpIndexWidget_ActivateCurrentItem(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_ActivateCurrentItemDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::activateCurrentItem();
}

void QHelpIndexWidget_FilterIndices(void* ptr, struct QtHelp_PackedString filter, struct QtHelp_PackedString wildcard)
{
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(const QString, QString::fromUtf8(filter.data, filter.len)), Q_ARG(const QString, QString::fromUtf8(wildcard.data, wildcard.len)));
}

void QHelpIndexWidget_FilterIndicesDefault(void* ptr, struct QtHelp_PackedString filter, struct QtHelp_PackedString wildcard)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::filterIndices(QString::fromUtf8(filter.data, filter.len), QString::fromUtf8(wildcard.data, wildcard.len));
}

void QHelpIndexWidget_ConnectLinkActivated(void* ptr)
{
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));
}

void QHelpIndexWidget_DisconnectLinkActivated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));
}

void QHelpIndexWidget_LinkActivated(void* ptr, void* link, struct QtHelp_PackedString keyword)
{
	static_cast<QHelpIndexWidget*>(ptr)->linkActivated(*static_cast<QUrl*>(link), QString::fromUtf8(keyword.data, keyword.len));
}

void QHelpIndexWidget_ConnectLinksActivated(void* ptr)
{
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QMap<QString, QUrl> &, const QString &)>(&QHelpIndexWidget::linksActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QMap<QString, QUrl> &, const QString &)>(&MyQHelpIndexWidget::Signal_LinksActivated));
}

void QHelpIndexWidget_DisconnectLinksActivated(void* ptr)
{
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QMap<QString, QUrl> &, const QString &)>(&QHelpIndexWidget::linksActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QMap<QString, QUrl> &, const QString &)>(&MyQHelpIndexWidget::Signal_LinksActivated));
}

void QHelpIndexWidget_LinksActivated(void* ptr, void* links, struct QtHelp_PackedString keyword)
{
	static_cast<QHelpIndexWidget*>(ptr)->linksActivated(*static_cast<QMap<QString, QUrl>*>(links), QString::fromUtf8(keyword.data, keyword.len));
}

void* QHelpIndexWidget___linksActivated_links_atList(void* ptr, struct QtHelp_PackedString i)
{
	return new QUrl(static_cast<QMap<QString, QUrl>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QHelpIndexWidget___linksActivated_links_setList(void* ptr, struct QtHelp_PackedString key, void* i)
{
	static_cast<QMap<QString, QUrl>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QUrl*>(i));
}

void* QHelpIndexWidget___linksActivated_links_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QUrl>;
}

struct QtHelp_PackedList QHelpIndexWidget___linksActivated_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QUrl>*>(ptr)->keys()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtHelp_PackedString QHelpIndexWidget_____linksActivated_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtHelp_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QHelpIndexWidget_____linksActivated_keyList_setList(void* ptr, struct QtHelp_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QHelpIndexWidget_____linksActivated_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

int QHelpIndexWidget___dataChanged_roles_atList(void* ptr, int i)
{
	return static_cast<QVector<int>*>(ptr)->at(i);
}

void QHelpIndexWidget___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QHelpIndexWidget___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>;
}

void* QHelpIndexWidget___indexesMoved_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexWidget___indexesMoved_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexWidget___indexesMoved_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpIndexWidget___selectedIndexes_atList(void* ptr, int i)
{
	return new QModelIndex(static_cast<QList<QModelIndex>*>(ptr)->at(i));
}

void QHelpIndexWidget___selectedIndexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QHelpIndexWidget___selectedIndexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>;
}

void* QHelpIndexWidget___scrollBarWidgets_atList(void* ptr, int i)
{
	return const_cast<QWidget*>(static_cast<QList<QWidget *>*>(ptr)->at(i));
}

void QHelpIndexWidget___scrollBarWidgets_setList(void* ptr, void* i)
{
	static_cast<QList<QWidget *>*>(ptr)->append(static_cast<QWidget*>(i));
}

void* QHelpIndexWidget___scrollBarWidgets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWidget *>;
}

void* QHelpIndexWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpIndexWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpIndexWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpIndexWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpIndexWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpIndexWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpIndexWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpIndexWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpIndexWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpIndexWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpIndexWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpIndexWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpIndexWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpIndexWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpIndexWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpIndexWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpIndexWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

void* QHelpIndexWidget_MoveCursorDefault(void* ptr, long long cursorAction, long long modifiers)
{
		return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::moveCursor(static_cast<QAbstractItemView::CursorAction>(cursorAction), static_cast<Qt::KeyboardModifier>(modifiers)));
}

char QHelpIndexWidget_EventDefault(void* ptr, void* e)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::event(static_cast<QEvent*>(e));
}

void QHelpIndexWidget_CurrentChangedDefault(void* ptr, void* current, void* previous)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::currentChanged(*static_cast<QModelIndex*>(current), *static_cast<QModelIndex*>(previous));
}

void QHelpIndexWidget_DataChangedDefault(void* ptr, void* topLeft, void* bottomRight, void* roles)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dataChanged(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight), *static_cast<QVector<int>*>(roles));
}

void QHelpIndexWidget_DragLeaveEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QHelpIndexWidget_DragMoveEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QHelpIndexWidget_DropEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dropEvent(static_cast<QDropEvent*>(e));
}

void QHelpIndexWidget_MouseMoveEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseReleaseEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_PaintEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::paintEvent(static_cast<QPaintEvent*>(e));
}

void QHelpIndexWidget_ResizeEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void QHelpIndexWidget_RowsAboutToBeRemovedDefault(void* ptr, void* parent, int start, int end)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::rowsAboutToBeRemoved(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_RowsInsertedDefault(void* ptr, void* parent, int start, int end)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::rowsInserted(*static_cast<QModelIndex*>(parent), start, end);
}

void QHelpIndexWidget_ScrollToDefault(void* ptr, void* index, long long hint)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QHelpIndexWidget_SelectionChangedDefault(void* ptr, void* selected, void* deselected)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectionChanged(*static_cast<QItemSelection*>(selected), *static_cast<QItemSelection*>(deselected));
}

void QHelpIndexWidget_SetSelectionDefault(void* ptr, void* rect, long long command)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpIndexWidget_StartDragDefault(void* ptr, long long supportedActions)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpIndexWidget_TimerEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::timerEvent(static_cast<QTimerEvent*>(e));
}

void QHelpIndexWidget_UpdateGeometriesDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::updateGeometries();
}

void QHelpIndexWidget_WheelEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void* QHelpIndexWidget_IndexAtDefault(void* ptr, void* p)
{
		return new QModelIndex(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::indexAt(*static_cast<QPoint*>(p)));
}

struct QtHelp_PackedList QHelpIndexWidget_SelectedIndexesDefault(void* ptr)
{
		return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectedIndexes()); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

void* QHelpIndexWidget_VisualRectDefault(void* ptr, void* index)
{
		return ({ QRect tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::visualRect(*static_cast<QModelIndex*>(index)); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_VisualRegionForSelectionDefault(void* ptr, void* selection)
{
		return new QRegion(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::visualRegionForSelection(*static_cast<QItemSelection*>(selection)));
}

void* QHelpIndexWidget_ViewportSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewportSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_ViewOptionsDefault(void* ptr)
{
		return new QStyleOptionViewItem(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewOptions());
}

char QHelpIndexWidget_IsIndexHiddenDefault(void* ptr, void* index)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::isIndexHidden(*static_cast<QModelIndex*>(index));
}

int QHelpIndexWidget_HorizontalOffsetDefault(void* ptr)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::horizontalOffset();
}

int QHelpIndexWidget_VerticalOffsetDefault(void* ptr)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::verticalOffset();
}

char QHelpIndexWidget_Edit2Default(void* ptr, void* index, long long trigger, void* event)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::edit(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::EditTrigger>(trigger), static_cast<QEvent*>(event));
}

char QHelpIndexWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusNextPrevChild(next != 0);
}

char QHelpIndexWidget_ViewportEventDefault(void* ptr, void* event)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::viewportEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_ClearSelectionDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::clearSelection();
}

void QHelpIndexWidget_CloseEditorDefault(void* ptr, void* editor, long long hint)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::closeEditor(static_cast<QWidget*>(editor), static_cast<QAbstractItemDelegate::EndEditHint>(hint));
}

void QHelpIndexWidget_CommitDataDefault(void* ptr, void* editor)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::commitData(static_cast<QWidget*>(editor));
}

void QHelpIndexWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpIndexWidget_EditDefault(void* ptr, void* index)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::edit(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_EditorDestroyedDefault(void* ptr, void* editor)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::editorDestroyed(static_cast<QObject*>(editor));
}

void QHelpIndexWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpIndexWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyboardSearchDefault(void* ptr, struct QtHelp_PackedString search)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyboardSearch(QString::fromUtf8(search.data, search.len));
}

void QHelpIndexWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_ResetDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::reset();
}

void QHelpIndexWidget_ScrollToBottomDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollToBottom();
}

void QHelpIndexWidget_ScrollToTopDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollToTop();
}

void QHelpIndexWidget_SelectAllDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectAll();
}

void QHelpIndexWidget_SetCurrentIndexDefault(void* ptr, void* index)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setCurrentIndex(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_SetModelDefault(void* ptr, void* model)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpIndexWidget_SetRootIndexDefault(void* ptr, void* index)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setRootIndex(*static_cast<QModelIndex*>(index));
}

void QHelpIndexWidget_SetSelectionModelDefault(void* ptr, void* selectionModel)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpIndexWidget_UpdateDefault(void* ptr, void* index)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::update(*static_cast<QModelIndex*>(index));
}

long long QHelpIndexWidget_SelectionCommandDefault(void* ptr, void* index, void* event)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::selectionCommand(*static_cast<QModelIndex*>(index), static_cast<QEvent*>(event));
}

void* QHelpIndexWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

int QHelpIndexWidget_SizeHintForColumnDefault(void* ptr, int column)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHintForColumn(column);
}

int QHelpIndexWidget_SizeHintForRowDefault(void* ptr, int row)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHintForRow(row);
}

void QHelpIndexWidget_ContextMenuEventDefault(void* ptr, void* e)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpIndexWidget_ScrollContentsByDefault(void* ptr, int dx, int dy)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollContentsBy(dx, dy);
}

void QHelpIndexWidget_SetupViewportDefault(void* ptr, void* viewport)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void* QHelpIndexWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpIndexWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QHelpIndexWidget_ChangeEventDefault(void* ptr, void* ev)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::changeEvent(static_cast<QEvent*>(ev));
}

char QHelpIndexWidget_CloseDefault(void* ptr)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::close();
}

char QHelpIndexWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QHelpIndexWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpIndexWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpIndexWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_HideDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hide();
}

void QHelpIndexWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpIndexWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_LowerDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::lower();
}

void QHelpIndexWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpIndexWidget_RaiseDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::raise();
}

void QHelpIndexWidget_RepaintDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::repaint();
}

void QHelpIndexWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setDisabled(disable != 0);
}

void QHelpIndexWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setEnabled(vbo != 0);
}

void QHelpIndexWidget_SetFocus2Default(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setFocus();
}

void QHelpIndexWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setHidden(hidden != 0);
}

void QHelpIndexWidget_SetStyleSheetDefault(void* ptr, struct QtHelp_PackedString styleSheet)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QHelpIndexWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setVisible(visible != 0);
}

void QHelpIndexWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setWindowModified(vbo != 0);
}

void QHelpIndexWidget_SetWindowTitleDefault(void* ptr, struct QtHelp_PackedString vqs)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QHelpIndexWidget_ShowDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::show();
}

void QHelpIndexWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpIndexWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showFullScreen();
}

void QHelpIndexWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showMaximized();
}

void QHelpIndexWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showMinimized();
}

void QHelpIndexWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showNormal();
}

void QHelpIndexWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpIndexWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::updateMicroFocus();
}

void* QHelpIndexWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::paintEngine();
}

char QHelpIndexWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hasHeightForWidth();
}

int QHelpIndexWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::heightForWidth(w);
}

int QHelpIndexWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QHelpIndexWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpIndexWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpIndexWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::deleteLater();
}

void QHelpIndexWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void* QHelpIndexWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::metaObject());
}

class MyQHelpSearchEngine: public QHelpSearchEngine
{
public:
	MyQHelpSearchEngine(QHelpEngineCore *helpEngine, QObject *parent = Q_NULLPTR) : QHelpSearchEngine(helpEngine, parent) {QHelpSearchEngine_QHelpSearchEngine_QRegisterMetaType();};
	void cancelIndexing() { callbackQHelpSearchEngine_CancelIndexing(this); };
	void cancelSearching() { callbackQHelpSearchEngine_CancelSearching(this); };
	void Signal_IndexingFinished() { callbackQHelpSearchEngine_IndexingFinished(this); };
	void Signal_IndexingStarted() { callbackQHelpSearchEngine_IndexingStarted(this); };
	void reindexDocumentation() { callbackQHelpSearchEngine_ReindexDocumentation(this); };
	void search(const QString & searchInput) { QByteArray t978613 = searchInput.toUtf8(); QtHelp_PackedString searchInputPacked = { const_cast<char*>(t978613.prepend("WHITESPACE").constData()+10), t978613.size()-10 };callbackQHelpSearchEngine_Search(this, searchInputPacked); };
	void Signal_SearchingFinished(int searchResultCount) { callbackQHelpSearchEngine_SearchingFinished(this, searchResultCount); };
	void Signal_SearchingStarted() { callbackQHelpSearchEngine_SearchingStarted(this); };
	bool event(QEvent * e) { return callbackQHelpSearchEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchEngine_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpSearchEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpSearchEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpSearchEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpSearchEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchEngine_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpSearchEngine*)

int QHelpSearchEngine_QHelpSearchEngine_QRegisterMetaType(){qRegisterMetaType<QHelpSearchEngine*>(); return qRegisterMetaType<MyQHelpSearchEngine*>();}

void QHelpSearchEngine_CancelIndexing(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelIndexingDefault(void* ptr)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::cancelIndexing();
}

void QHelpSearchEngine_CancelSearching(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

void QHelpSearchEngine_CancelSearchingDefault(void* ptr)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::cancelSearching();
}

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QWindow*>(parent));
	} else {
		return new MyQHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
	}
}

void* QHelpSearchEngine_QueryWidget(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void* QHelpSearchEngine_ResultWidget(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
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

void QHelpSearchEngine_ReindexDocumentation(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void QHelpSearchEngine_ReindexDocumentationDefault(void* ptr)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::reindexDocumentation();
}

void QHelpSearchEngine_Search(void* ptr, struct QtHelp_PackedString searchInput)
{
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "search", Q_ARG(const QString, QString::fromUtf8(searchInput.data, searchInput.len)));
}

void QHelpSearchEngine_SearchDefault(void* ptr, struct QtHelp_PackedString searchInput)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::search(QString::fromUtf8(searchInput.data, searchInput.len));
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr)
{
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr)
{
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));
}

void QHelpSearchEngine_SearchingFinished(void* ptr, int searchResultCount)
{
	static_cast<QHelpSearchEngine*>(ptr)->searchingFinished(searchResultCount);
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

struct QtHelp_PackedString QHelpSearchEngine_SearchInput(void* ptr)
{
	return ({ QByteArray t530fa8 = static_cast<QHelpSearchEngine*>(ptr)->searchInput().toUtf8(); QtHelp_PackedString { const_cast<char*>(t530fa8.prepend("WHITESPACE").constData()+10), t530fa8.size()-10 }; });
}

struct QtHelp_PackedList QHelpSearchEngine_SearchResults(void* ptr, int start, int end)
{
	return ({ QVector<QHelpSearchResult>* tmpValue = new QVector<QHelpSearchResult>(static_cast<QHelpSearchEngine*>(ptr)->searchResults(start, end)); QtHelp_PackedList { tmpValue, tmpValue->size() }; });
}

int QHelpSearchEngine_SearchResultCount(void* ptr)
{
	return static_cast<QHelpSearchEngine*>(ptr)->searchResultCount();
}

void* QHelpSearchEngine___search_queryList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHelpSearchQuery>;
}

void* QHelpSearchEngine___query_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHelpSearchQuery>;
}

void* QHelpSearchEngine___searchResults_atList(void* ptr, int i)
{
	return new QHelpSearchResult(static_cast<QVector<QHelpSearchResult>*>(ptr)->at(i));
}

void QHelpSearchEngine___searchResults_setList(void* ptr, void* i)
{
	static_cast<QVector<QHelpSearchResult>*>(ptr)->append(*static_cast<QHelpSearchResult*>(i));
}

void* QHelpSearchEngine___searchResults_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QHelpSearchResult>;
}

void* QHelpSearchEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpSearchEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpSearchEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpSearchEngine___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchEngine___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchEngine___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchEngine___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchEngine___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchEngine___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpSearchEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpSearchEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::event(static_cast<QEvent*>(e));
}

char QHelpSearchEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpSearchEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::deleteLater();
}

void QHelpSearchEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHelpSearchEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::metaObject());
}

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget
{
public:
	MyQHelpSearchQueryWidget(QWidget *parent = Q_NULLPTR) : QHelpSearchQueryWidget(parent) {QHelpSearchQueryWidget_QHelpSearchQueryWidget_QRegisterMetaType();};
	void Signal_Search() { callbackQHelpSearchQueryWidget_Search(this); };
	bool close() { return callbackQHelpSearchQueryWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQHelpSearchQueryWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQHelpSearchQueryWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpSearchQueryWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchQueryWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQHelpSearchQueryWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchQueryWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchQueryWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQHelpSearchQueryWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchQueryWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchQueryWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchQueryWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchQueryWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchQueryWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpSearchQueryWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchQueryWidget_FocusOutEvent(this, event); };
	void hide() { callbackQHelpSearchQueryWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchQueryWidget_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchQueryWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchQueryWidget_LeaveEvent(this, event); };
	void lower() { callbackQHelpSearchQueryWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidget_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchQueryWidget_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchQueryWidget_PaintEvent(this, event); };
	void raise() { callbackQHelpSearchQueryWidget_Raise(this); };
	void repaint() { callbackQHelpSearchQueryWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchQueryWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQHelpSearchQueryWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQHelpSearchQueryWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQHelpSearchQueryWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQHelpSearchQueryWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtHelp_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQHelpSearchQueryWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQHelpSearchQueryWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQHelpSearchQueryWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtHelp_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQHelpSearchQueryWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQHelpSearchQueryWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQHelpSearchQueryWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQHelpSearchQueryWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQHelpSearchQueryWidget_ShowMaximized(this); };
	void showMinimized() { callbackQHelpSearchQueryWidget_ShowMinimized(this); };
	void showNormal() { callbackQHelpSearchQueryWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchQueryWidget_TabletEvent(this, event); };
	void update() { callbackQHelpSearchQueryWidget_Update(this); };
	void updateMicroFocus() { callbackQHelpSearchQueryWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchQueryWidget_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQHelpSearchQueryWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtHelp_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQHelpSearchQueryWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQHelpSearchQueryWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchQueryWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchQueryWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpSearchQueryWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQHelpSearchQueryWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQHelpSearchQueryWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQHelpSearchQueryWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchQueryWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpSearchQueryWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchQueryWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchQueryWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpSearchQueryWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpSearchQueryWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchQueryWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpSearchQueryWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchQueryWidget_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchQueryWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpSearchQueryWidget*)

int QHelpSearchQueryWidget_QHelpSearchQueryWidget_QRegisterMetaType(){qRegisterMetaType<QHelpSearchQueryWidget*>(); return qRegisterMetaType<MyQHelpSearchQueryWidget*>();}

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

void QHelpSearchQueryWidget_SetSearchInput(void* ptr, struct QtHelp_PackedString searchInput)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->setSearchInput(QString::fromUtf8(searchInput.data, searchInput.len));
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr)
{
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

struct QtHelp_PackedString QHelpSearchQueryWidget_SearchInput(void* ptr)
{
	return ({ QByteArray t285afe = static_cast<QHelpSearchQueryWidget*>(ptr)->searchInput().toUtf8(); QtHelp_PackedString { const_cast<char*>(t285afe.prepend("WHITESPACE").constData()+10), t285afe.size()-10 }; });
}

char QHelpSearchQueryWidget_IsCompactMode(void* ptr)
{
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget___setQuery_queryList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHelpSearchQuery>;
}

void* QHelpSearchQueryWidget___query_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QHelpSearchQuery>;
}

void* QHelpSearchQueryWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchQueryWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchQueryWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchQueryWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchQueryWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchQueryWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchQueryWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpSearchQueryWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpSearchQueryWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchQueryWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchQueryWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchQueryWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchQueryWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchQueryWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchQueryWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpSearchQueryWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchQueryWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpSearchQueryWidget_CloseDefault(void* ptr)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::close();
}

char QHelpSearchQueryWidget_EventDefault(void* ptr, void* event)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::event(static_cast<QEvent*>(event));
}

char QHelpSearchQueryWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusNextPrevChild(next != 0);
}

char QHelpSearchQueryWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QHelpSearchQueryWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchQueryWidget_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::changeEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchQueryWidget_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchQueryWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchQueryWidget_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchQueryWidget_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchQueryWidget_DropEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchQueryWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_HideDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hide();
}

void QHelpSearchQueryWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchQueryWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchQueryWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_LowerDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::lower();
}

void QHelpSearchQueryWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchQueryWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchQueryWidget_RaiseDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::raise();
}

void QHelpSearchQueryWidget_RepaintDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::repaint();
}

void QHelpSearchQueryWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchQueryWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setDisabled(disable != 0);
}

void QHelpSearchQueryWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setEnabled(vbo != 0);
}

void QHelpSearchQueryWidget_SetFocus2Default(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setFocus();
}

void QHelpSearchQueryWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setHidden(hidden != 0);
}

void QHelpSearchQueryWidget_SetStyleSheetDefault(void* ptr, struct QtHelp_PackedString styleSheet)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QHelpSearchQueryWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setVisible(visible != 0);
}

void QHelpSearchQueryWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setWindowModified(vbo != 0);
}

void QHelpSearchQueryWidget_SetWindowTitleDefault(void* ptr, struct QtHelp_PackedString vqs)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QHelpSearchQueryWidget_ShowDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::show();
}

void QHelpSearchQueryWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchQueryWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showFullScreen();
}

void QHelpSearchQueryWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showMaximized();
}

void QHelpSearchQueryWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showMinimized();
}

void QHelpSearchQueryWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showNormal();
}

void QHelpSearchQueryWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchQueryWidget_UpdateDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::update();
}

void QHelpSearchQueryWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::updateMicroFocus();
}

void QHelpSearchQueryWidget_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QHelpSearchQueryWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::paintEngine();
}

void* QHelpSearchQueryWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchQueryWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchQueryWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QHelpSearchQueryWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hasHeightForWidth();
}

int QHelpSearchQueryWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::heightForWidth(w);
}

int QHelpSearchQueryWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QHelpSearchQueryWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchQueryWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchQueryWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::deleteLater();
}

void QHelpSearchQueryWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchQueryWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHelpSearchQueryWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::metaObject());
}

void* QHelpSearchResult_NewQHelpSearchResult()
{
	return new QHelpSearchResult();
}

void* QHelpSearchResult_NewQHelpSearchResult2(void* other)
{
	return new QHelpSearchResult(*static_cast<QHelpSearchResult*>(other));
}

void* QHelpSearchResult_NewQHelpSearchResult3(void* url, struct QtHelp_PackedString title, struct QtHelp_PackedString snippet)
{
	return new QHelpSearchResult(*static_cast<QUrl*>(url), QString::fromUtf8(title.data, title.len), QString::fromUtf8(snippet.data, snippet.len));
}

void QHelpSearchResult_DestroyQHelpSearchResult(void* ptr)
{
	static_cast<QHelpSearchResult*>(ptr)->~QHelpSearchResult();
}

struct QtHelp_PackedString QHelpSearchResult_Snippet(void* ptr)
{
	return ({ QByteArray tf76589 = static_cast<QHelpSearchResult*>(ptr)->snippet().toUtf8(); QtHelp_PackedString { const_cast<char*>(tf76589.prepend("WHITESPACE").constData()+10), tf76589.size()-10 }; });
}

struct QtHelp_PackedString QHelpSearchResult_Title(void* ptr)
{
	return ({ QByteArray t34a6b4 = static_cast<QHelpSearchResult*>(ptr)->title().toUtf8(); QtHelp_PackedString { const_cast<char*>(t34a6b4.prepend("WHITESPACE").constData()+10), t34a6b4.size()-10 }; });
}

void* QHelpSearchResult_Url(void* ptr)
{
	return new QUrl(static_cast<QHelpSearchResult*>(ptr)->url());
}

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget
{
public:
	void Signal_RequestShowLink(const QUrl & link) { callbackQHelpSearchResultWidget_RequestShowLink(this, const_cast<QUrl*>(&link)); };
	bool close() { return callbackQHelpSearchResultWidget_Close(this) != 0; };
	bool event(QEvent * event) { return callbackQHelpSearchResultWidget_Event(this, event) != 0; };
	bool focusNextPrevChild(bool next) { return callbackQHelpSearchResultWidget_FocusNextPrevChild(this, next) != 0; };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQHelpSearchResultWidget_NativeEvent(this, const_cast<QByteArray*>(&eventType), message, *result) != 0; };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchResultWidget_ActionEvent(this, event); };
	void changeEvent(QEvent * event) { callbackQHelpSearchResultWidget_ChangeEvent(this, event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchResultWidget_CloseEvent(this, event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchResultWidget_ContextMenuEvent(this, event); };
	void Signal_CustomContextMenuRequested(const QPoint & pos) { callbackQHelpSearchResultWidget_CustomContextMenuRequested(this, const_cast<QPoint*>(&pos)); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchResultWidget_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchResultWidget_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchResultWidget_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchResultWidget_DropEvent(this, event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchResultWidget_EnterEvent(this, event); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpSearchResultWidget_FocusInEvent(this, event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchResultWidget_FocusOutEvent(this, event); };
	void hide() { callbackQHelpSearchResultWidget_Hide(this); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchResultWidget_HideEvent(this, event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchResultWidget_InputMethodEvent(this, event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchResultWidget_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchResultWidget_KeyReleaseEvent(this, event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchResultWidget_LeaveEvent(this, event); };
	void lower() { callbackQHelpSearchResultWidget_Lower(this); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchResultWidget_MouseReleaseEvent(this, event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchResultWidget_MoveEvent(this, event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchResultWidget_PaintEvent(this, event); };
	void raise() { callbackQHelpSearchResultWidget_Raise(this); };
	void repaint() { callbackQHelpSearchResultWidget_Repaint(this); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchResultWidget_ResizeEvent(this, event); };
	void setDisabled(bool disable) { callbackQHelpSearchResultWidget_SetDisabled(this, disable); };
	void setEnabled(bool vbo) { callbackQHelpSearchResultWidget_SetEnabled(this, vbo); };
	void setFocus() { callbackQHelpSearchResultWidget_SetFocus2(this); };
	void setHidden(bool hidden) { callbackQHelpSearchResultWidget_SetHidden(this, hidden); };
	void setStyleSheet(const QString & styleSheet) { QByteArray t728ae7 = styleSheet.toUtf8(); QtHelp_PackedString styleSheetPacked = { const_cast<char*>(t728ae7.prepend("WHITESPACE").constData()+10), t728ae7.size()-10 };callbackQHelpSearchResultWidget_SetStyleSheet(this, styleSheetPacked); };
	void setVisible(bool visible) { callbackQHelpSearchResultWidget_SetVisible(this, visible); };
	void setWindowModified(bool vbo) { callbackQHelpSearchResultWidget_SetWindowModified(this, vbo); };
	void setWindowTitle(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); QtHelp_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackQHelpSearchResultWidget_SetWindowTitle(this, vqsPacked); };
	void show() { callbackQHelpSearchResultWidget_Show(this); };
	void showEvent(QShowEvent * event) { callbackQHelpSearchResultWidget_ShowEvent(this, event); };
	void showFullScreen() { callbackQHelpSearchResultWidget_ShowFullScreen(this); };
	void showMaximized() { callbackQHelpSearchResultWidget_ShowMaximized(this); };
	void showMinimized() { callbackQHelpSearchResultWidget_ShowMinimized(this); };
	void showNormal() { callbackQHelpSearchResultWidget_ShowNormal(this); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchResultWidget_TabletEvent(this, event); };
	void update() { callbackQHelpSearchResultWidget_Update(this); };
	void updateMicroFocus() { callbackQHelpSearchResultWidget_UpdateMicroFocus(this); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchResultWidget_WheelEvent(this, event); };
	void Signal_WindowIconChanged(const QIcon & icon) { callbackQHelpSearchResultWidget_WindowIconChanged(this, const_cast<QIcon*>(&icon)); };
	void Signal_WindowTitleChanged(const QString & title) { QByteArray t3c6de1 = title.toUtf8(); QtHelp_PackedString titlePacked = { const_cast<char*>(t3c6de1.prepend("WHITESPACE").constData()+10), t3c6de1.size()-10 };callbackQHelpSearchResultWidget_WindowTitleChanged(this, titlePacked); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQHelpSearchResultWidget_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize minimumSizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchResultWidget_MinimumSizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize sizeHint() const { return *static_cast<QSize*>(callbackQHelpSearchResultWidget_SizeHint(const_cast<void*>(static_cast<const void*>(this)))); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackQHelpSearchResultWidget_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool hasHeightForWidth() const { return callbackQHelpSearchResultWidget_HasHeightForWidth(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	int heightForWidth(int w) const { return callbackQHelpSearchResultWidget_HeightForWidth(const_cast<void*>(static_cast<const void*>(this)), w); };
	int metric(QPaintDevice::PaintDeviceMetric m) const { return callbackQHelpSearchResultWidget_Metric(const_cast<void*>(static_cast<const void*>(this)), m); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQHelpSearchResultWidget_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQHelpSearchResultWidget_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQHelpSearchResultWidget_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQHelpSearchResultWidget_CustomEvent(this, event); };
	void deleteLater() { callbackQHelpSearchResultWidget_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQHelpSearchResultWidget_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQHelpSearchResultWidget_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtHelp_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQHelpSearchResultWidget_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchResultWidget_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQHelpSearchResultWidget_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHelpSearchResultWidget*)

int QHelpSearchResultWidget_QHelpSearchResultWidget_QRegisterMetaType(){qRegisterMetaType<QHelpSearchResultWidget*>(); return qRegisterMetaType<MyQHelpSearchResultWidget*>();}

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

void* QHelpSearchResultWidget___addActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___addActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchResultWidget___addActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchResultWidget___insertActions_actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___insertActions_actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchResultWidget___insertActions_actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchResultWidget___actions_atList(void* ptr, int i)
{
	return const_cast<QAction*>(static_cast<QList<QAction *>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___actions_setList(void* ptr, void* i)
{
	static_cast<QList<QAction *>*>(ptr)->append(static_cast<QAction*>(i));
}

void* QHelpSearchResultWidget___actions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAction *>;
}

void* QHelpSearchResultWidget___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QHelpSearchResultWidget___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QHelpSearchResultWidget___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchResultWidget___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchResultWidget___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchResultWidget___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchResultWidget___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchResultWidget___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QHelpSearchResultWidget___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QHelpSearchResultWidget___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QHelpSearchResultWidget___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QHelpSearchResultWidget_CloseDefault(void* ptr)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::close();
}

char QHelpSearchResultWidget_EventDefault(void* ptr, void* event)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::event(static_cast<QEvent*>(event));
}

char QHelpSearchResultWidget_FocusNextPrevChildDefault(void* ptr, char next)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusNextPrevChild(next != 0);
}

char QHelpSearchResultWidget_NativeEventDefault(void* ptr, void* eventType, void* message, long result)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::nativeEvent(*static_cast<QByteArray*>(eventType), message, &result);
}

void QHelpSearchResultWidget_ActionEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchResultWidget_ChangeEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::changeEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_CloseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchResultWidget_ContextMenuEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchResultWidget_DragEnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchResultWidget_DragLeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchResultWidget_DragMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchResultWidget_DropEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchResultWidget_EnterEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_FocusInEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusOutEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_HideDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hide();
}

void QHelpSearchResultWidget_HideEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchResultWidget_InputMethodEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchResultWidget_KeyPressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_LeaveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_LowerDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::lower();
}

void QHelpSearchResultWidget_MouseDoubleClickEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseMoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MousePressEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseReleaseEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MoveEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchResultWidget_PaintEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchResultWidget_RaiseDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::raise();
}

void QHelpSearchResultWidget_RepaintDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::repaint();
}

void QHelpSearchResultWidget_ResizeEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchResultWidget_SetDisabledDefault(void* ptr, char disable)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setDisabled(disable != 0);
}

void QHelpSearchResultWidget_SetEnabledDefault(void* ptr, char vbo)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setEnabled(vbo != 0);
}

void QHelpSearchResultWidget_SetFocus2Default(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setFocus();
}

void QHelpSearchResultWidget_SetHiddenDefault(void* ptr, char hidden)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setHidden(hidden != 0);
}

void QHelpSearchResultWidget_SetStyleSheetDefault(void* ptr, struct QtHelp_PackedString styleSheet)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setStyleSheet(QString::fromUtf8(styleSheet.data, styleSheet.len));
}

void QHelpSearchResultWidget_SetVisibleDefault(void* ptr, char visible)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setVisible(visible != 0);
}

void QHelpSearchResultWidget_SetWindowModifiedDefault(void* ptr, char vbo)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setWindowModified(vbo != 0);
}

void QHelpSearchResultWidget_SetWindowTitleDefault(void* ptr, struct QtHelp_PackedString vqs)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::setWindowTitle(QString::fromUtf8(vqs.data, vqs.len));
}

void QHelpSearchResultWidget_ShowDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::show();
}

void QHelpSearchResultWidget_ShowEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchResultWidget_ShowFullScreenDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showFullScreen();
}

void QHelpSearchResultWidget_ShowMaximizedDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showMaximized();
}

void QHelpSearchResultWidget_ShowMinimizedDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showMinimized();
}

void QHelpSearchResultWidget_ShowNormalDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showNormal();
}

void QHelpSearchResultWidget_TabletEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchResultWidget_UpdateDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::update();
}

void QHelpSearchResultWidget_UpdateMicroFocusDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::updateMicroFocus();
}

void QHelpSearchResultWidget_WheelEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void* QHelpSearchResultWidget_PaintEngineDefault(void* ptr)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::paintEngine();
}

void* QHelpSearchResultWidget_MinimumSizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::minimumSizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchResultWidget_SizeHintDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::sizeHint(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QHelpSearchResultWidget_InputMethodQueryDefault(void* ptr, long long query)
{
		return new QVariant(static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char QHelpSearchResultWidget_HasHeightForWidthDefault(void* ptr)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hasHeightForWidth();
}

int QHelpSearchResultWidget_HeightForWidthDefault(void* ptr, int w)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::heightForWidth(w);
}

int QHelpSearchResultWidget_MetricDefault(void* ptr, long long m)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::metric(static_cast<QPaintDevice::PaintDeviceMetric>(m));
}

char QHelpSearchResultWidget_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchResultWidget_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchResultWidget_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_DeleteLaterDefault(void* ptr)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::deleteLater();
}

void QHelpSearchResultWidget_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHelpSearchResultWidget_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QHelpSearchResultWidget_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::metaObject());
}


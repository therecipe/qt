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
#include <QMetaObject>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

void* QHelpContentItem_Child(void* ptr, int row){
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child){
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

void* QHelpContentItem_Parent(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

void* QHelpContentItem_Url(void* ptr){
	return new QUrl(static_cast<QHelpContentItem*>(ptr)->url());
}

void QHelpContentItem_DestroyQHelpContentItem(void* ptr){
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

class MyQHelpContentModel: public QHelpContentModel {
public:
	void Signal_ContentsCreated() { callbackQHelpContentModelContentsCreated(this->objectName().toUtf8().data()); };
	void Signal_ContentsCreationStarted() { callbackQHelpContentModelContentsCreationStarted(this->objectName().toUtf8().data()); };
	void fetchMore(const QModelIndex & parent) { if (!callbackQHelpContentModelFetchMore(this->objectName().toUtf8().data(), parent.internalPointer())) { QHelpContentModel::fetchMore(parent); }; };
	void revert() { if (!callbackQHelpContentModelRevert(this->objectName().toUtf8().data())) { QHelpContentModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { if (!callbackQHelpContentModelSort(this->objectName().toUtf8().data(), column, order)) { QHelpContentModel::sort(column, order); }; };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpContentModelTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpContentModel::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpContentModelChildEvent(this->objectName().toUtf8().data(), event)) { QHelpContentModel::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpContentModelCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpContentModel::customEvent(event); }; };
};

int QHelpContentModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel_ContentItemAt(void* ptr, void* index){
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void QHelpContentModel_ConnectContentsCreated(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_DisconnectContentsCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_ConnectContentsCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_DisconnectContentsCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_CreateContents(void* ptr, char* customFilterName){
	static_cast<QHelpContentModel*>(ptr)->createContents(QString(customFilterName));
}

void* QHelpContentModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QHelpContentModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QHelpContentModel_IsCreatingContents(void* ptr){
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

void* QHelpContentModel_Parent(void* ptr, void* index){
	return static_cast<QHelpContentModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QHelpContentModel_RowCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_DestroyQHelpContentModel(void* ptr){
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

class MyQHelpContentWidget: public QHelpContentWidget {
public:
	void Signal_LinkActivated(const QUrl & link) { callbackQHelpContentWidgetLinkActivated(this->objectName().toUtf8().data(), new QUrl(link)); };
	void keyboardSearch(const QString & search) { if (!callbackQHelpContentWidgetKeyboardSearch(this->objectName().toUtf8().data(), search.toUtf8().data())) { QHelpContentWidget::keyboardSearch(search); }; };
	void reset() { if (!callbackQHelpContentWidgetReset(this->objectName().toUtf8().data())) { QHelpContentWidget::reset(); }; };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { if (!callbackQHelpContentWidgetScrollTo(this->objectName().toUtf8().data(), index.internalPointer(), hint)) { QHelpContentWidget::scrollTo(index, hint); }; };
	void selectAll() { if (!callbackQHelpContentWidgetSelectAll(this->objectName().toUtf8().data())) { QHelpContentWidget::selectAll(); }; };
	void setModel(QAbstractItemModel * model) { if (!callbackQHelpContentWidgetSetModel(this->objectName().toUtf8().data(), model)) { QHelpContentWidget::setModel(model); }; };
	void setRootIndex(const QModelIndex & index) { if (!callbackQHelpContentWidgetSetRootIndex(this->objectName().toUtf8().data(), index.internalPointer())) { QHelpContentWidget::setRootIndex(index); }; };
	void setSelectionModel(QItemSelectionModel * selectionModel) { if (!callbackQHelpContentWidgetSetSelectionModel(this->objectName().toUtf8().data(), selectionModel)) { QHelpContentWidget::setSelectionModel(selectionModel); }; };
	void setupViewport(QWidget * viewport) { if (!callbackQHelpContentWidgetSetupViewport(this->objectName().toUtf8().data(), viewport)) { QHelpContentWidget::setupViewport(viewport); }; };
	void setVisible(bool visible) { if (!callbackQHelpContentWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QHelpContentWidget::setVisible(visible); }; };
protected:
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { if (!callbackQHelpContentWidgetCurrentChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer())) { QHelpContentWidget::currentChanged(current, previous); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQHelpContentWidgetDragMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::dragMoveEvent(event); }; };
	void drawBranches(QPainter * painter, const QRect & rect, const QModelIndex & index) const { if (!callbackQHelpContentWidgetDrawBranches(this->objectName().toUtf8().data(), painter, new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), index.internalPointer())) { QHelpContentWidget::drawBranches(painter, rect, index); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQHelpContentWidgetKeyPressEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::keyPressEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQHelpContentWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQHelpContentWidgetMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQHelpContentWidgetMousePressEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQHelpContentWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::mouseReleaseEvent(event); }; };
	void paintEvent(QPaintEvent * event) { if (!callbackQHelpContentWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::paintEvent(event); }; };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { if (!callbackQHelpContentWidgetRowsAboutToBeRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), start, end)) { QHelpContentWidget::rowsAboutToBeRemoved(parent, start, end); }; };
	void rowsInserted(const QModelIndex & parent, int start, int end) { if (!callbackQHelpContentWidgetRowsInserted(this->objectName().toUtf8().data(), parent.internalPointer(), start, end)) { QHelpContentWidget::rowsInserted(parent, start, end); }; };
	void scrollContentsBy(int dx, int dy) { if (!callbackQHelpContentWidgetScrollContentsBy(this->objectName().toUtf8().data(), dx, dy)) { QHelpContentWidget::scrollContentsBy(dx, dy); }; };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { if (!callbackQHelpContentWidgetSetSelection(this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), command)) { QHelpContentWidget::setSelection(rect, command); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpContentWidgetTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::timerEvent(event); }; };
	void updateGeometries() { if (!callbackQHelpContentWidgetUpdateGeometries(this->objectName().toUtf8().data())) { QHelpContentWidget::updateGeometries(); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQHelpContentWidgetDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::dragLeaveEvent(event); }; };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { if (!callbackQHelpContentWidgetCloseEditor(this->objectName().toUtf8().data(), editor, hint)) { QHelpContentWidget::closeEditor(editor, hint); }; };
	void commitData(QWidget * editor) { if (!callbackQHelpContentWidgetCommitData(this->objectName().toUtf8().data(), editor)) { QHelpContentWidget::commitData(editor); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQHelpContentWidgetDragEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::dragEnterEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQHelpContentWidgetDropEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::dropEvent(event); }; };
	void editorDestroyed(QObject * editor) { if (!callbackQHelpContentWidgetEditorDestroyed(this->objectName().toUtf8().data(), editor)) { QHelpContentWidget::editorDestroyed(editor); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQHelpContentWidgetFocusInEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQHelpContentWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::focusOutEvent(event); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQHelpContentWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::inputMethodEvent(event); }; };
	void resizeEvent(QResizeEvent * event) { if (!callbackQHelpContentWidgetResizeEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::resizeEvent(event); }; };
	void startDrag(Qt::DropActions supportedActions) { if (!callbackQHelpContentWidgetStartDrag(this->objectName().toUtf8().data(), supportedActions)) { QHelpContentWidget::startDrag(supportedActions); }; };
	void contextMenuEvent(QContextMenuEvent * e) { if (!callbackQHelpContentWidgetContextMenuEvent(this->objectName().toUtf8().data(), e)) { QHelpContentWidget::contextMenuEvent(e); }; };
	void wheelEvent(QWheelEvent * e) { if (!callbackQHelpContentWidgetWheelEvent(this->objectName().toUtf8().data(), e)) { QHelpContentWidget::wheelEvent(e); }; };
	void changeEvent(QEvent * ev) { if (!callbackQHelpContentWidgetChangeEvent(this->objectName().toUtf8().data(), ev)) { QHelpContentWidget::changeEvent(ev); }; };
	void actionEvent(QActionEvent * event) { if (!callbackQHelpContentWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::actionEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQHelpContentWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::enterEvent(event); }; };
	void hideEvent(QHideEvent * event) { if (!callbackQHelpContentWidgetHideEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::hideEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQHelpContentWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQHelpContentWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::moveEvent(event); }; };
	void showEvent(QShowEvent * event) { if (!callbackQHelpContentWidgetShowEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::showEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQHelpContentWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::closeEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQHelpContentWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QHelpContentWidget::initPainter(painter); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQHelpContentWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::keyReleaseEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQHelpContentWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::tabletEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpContentWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpContentWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpContentWidget::customEvent(event); }; };
};

void* QHelpContentWidget_IndexOf(void* ptr, void* link){
	return static_cast<QHelpContentWidget*>(ptr)->indexOf(*static_cast<QUrl*>(link)).internalPointer();
}

void QHelpContentWidget_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

void QHelpContentWidget_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent){
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

void* QHelpEngine_ContentModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_ContentWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_IndexModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

void* QHelpEngine_IndexWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

void* QHelpEngine_SearchEngine(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(void* ptr){
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

class MyQHelpEngineCore: public QHelpEngineCore {
public:
	MyQHelpEngineCore(const QString &collectionFile, QObject *parent) : QHelpEngineCore(collectionFile, parent) {};
	void Signal_CurrentFilterChanged(const QString & newFilter) { callbackQHelpEngineCoreCurrentFilterChanged(this->objectName().toUtf8().data(), newFilter.toUtf8().data()); };
	void Signal_ReadersAboutToBeInvalidated() { callbackQHelpEngineCoreReadersAboutToBeInvalidated(this->objectName().toUtf8().data()); };
	void Signal_SetupFinished() { callbackQHelpEngineCoreSetupFinished(this->objectName().toUtf8().data()); };
	void Signal_SetupStarted() { callbackQHelpEngineCoreSetupStarted(this->objectName().toUtf8().data()); };
	void Signal_Warning(const QString & msg) { callbackQHelpEngineCoreWarning(this->objectName().toUtf8().data(), msg.toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpEngineCoreTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpEngineCore::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpEngineCoreChildEvent(this->objectName().toUtf8().data(), event)) { QHelpEngineCore::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpEngineCoreCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpEngineCore::customEvent(event); }; };
};

int QHelpEngineCore_AutoSaveFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

char* QHelpEngineCore_CollectionFile(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8().data();
}

char* QHelpEngineCore_CurrentFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8().data();
}

void QHelpEngineCore_SetAutoSaveFilter(void* ptr, int save){
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(void* ptr, char* fileName){
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString(fileName));
}

void QHelpEngineCore_SetCurrentFilter(void* ptr, char* filterName){
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString(filterName));
}

void* QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, void* parent){
	return new MyQHelpEngineCore(QString(collectionFile), static_cast<QObject*>(parent));
}

int QHelpEngineCore_AddCustomFilter(void* ptr, char* filterName, char* attributes){
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString(filterName), QString(attributes).split(",,,", QString::SkipEmptyParts));
}

int QHelpEngineCore_CopyCollectionFile(void* ptr, char* fileName){
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString(fileName));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

char* QHelpEngineCore_CustomFilters(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->customFilters().join(",,,").toUtf8().data();
}

void* QHelpEngineCore_CustomValue(void* ptr, char* key, void* defaultValue){
	return new QVariant(static_cast<QHelpEngineCore*>(ptr)->customValue(QString(key), *static_cast<QVariant*>(defaultValue)));
}

char* QHelpEngineCore_DocumentationFileName(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString(namespaceName)).toUtf8().data();
}

char* QHelpEngineCore_Error(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->error().toUtf8().data();
}

void* QHelpEngineCore_FileData(void* ptr, void* url){
	return new QByteArray(static_cast<QHelpEngineCore*>(ptr)->fileData(*static_cast<QUrl*>(url)));
}

char* QHelpEngineCore_FilterAttributes(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join(",,,").toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes2(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString(filterName)).join(",,,").toUtf8().data();
}

void* QHelpEngineCore_FindFile(void* ptr, void* url){
	return new QUrl(static_cast<QHelpEngineCore*>(ptr)->findFile(*static_cast<QUrl*>(url)));
}

void* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name){
	return new QVariant(QHelpEngineCore::metaData(QString(documentationFileName), QString(name)));
}

char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName){
	return QHelpEngineCore::namespaceName(QString(documentationFileName)).toUtf8().data();
}

void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

int QHelpEngineCore_RegisterDocumentation(void* ptr, char* documentationFileName){
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString(documentationFileName));
}

char* QHelpEngineCore_RegisteredDocumentations(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join(",,,").toUtf8().data();
}

int QHelpEngineCore_RemoveCustomFilter(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString(filterName));
}

int QHelpEngineCore_RemoveCustomValue(void* ptr, char* key){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString(key));
}

int QHelpEngineCore_SetCustomValue(void* ptr, char* key, void* value){
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString(key), *static_cast<QVariant*>(value));
}

int QHelpEngineCore_SetupData(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

void QHelpEngineCore_ConnectSetupFinished(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_DisconnectSetupFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_ConnectSetupStarted(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

void QHelpEngineCore_DisconnectSetupStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

int QHelpEngineCore_UnregisterDocumentation(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString(namespaceName));
}

void QHelpEngineCore_ConnectWarning(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DisconnectWarning(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

class MyQHelpIndexModel: public QHelpIndexModel {
public:
	void Signal_IndexCreated() { callbackQHelpIndexModelIndexCreated(this->objectName().toUtf8().data()); };
	void Signal_IndexCreationStarted() { callbackQHelpIndexModelIndexCreationStarted(this->objectName().toUtf8().data()); };
	void sort(int column, Qt::SortOrder order) { if (!callbackQHelpIndexModelSort(this->objectName().toUtf8().data(), column, order)) { QHelpIndexModel::sort(column, order); }; };
	void fetchMore(const QModelIndex & parent) { if (!callbackQHelpIndexModelFetchMore(this->objectName().toUtf8().data(), parent.internalPointer())) { QHelpIndexModel::fetchMore(parent); }; };
	void revert() { if (!callbackQHelpIndexModelRevert(this->objectName().toUtf8().data())) { QHelpIndexModel::revert(); }; };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpIndexModelTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexModel::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpIndexModelChildEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexModel::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpIndexModelCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexModel::customEvent(event); }; };
};

void QHelpIndexModel_CreateIndex(void* ptr, char* customFilterName){
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString(customFilterName));
}

void* QHelpIndexModel_Filter(void* ptr, char* filter, char* wildcard){
	return static_cast<QHelpIndexModel*>(ptr)->filter(QString(filter), QString(wildcard)).internalPointer();
}

void QHelpIndexModel_ConnectIndexCreated(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_DisconnectIndexCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_ConnectIndexCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

void QHelpIndexModel_DisconnectIndexCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

int QHelpIndexModel_IsCreatingIndex(void* ptr){
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

class MyQHelpIndexWidget: public QHelpIndexWidget {
public:
	void Signal_LinkActivated(const QUrl & link, const QString & keyword) { callbackQHelpIndexWidgetLinkActivated(this->objectName().toUtf8().data(), new QUrl(link), keyword.toUtf8().data()); };
	void scrollTo(const QModelIndex & index, QAbstractItemView::ScrollHint hint) { if (!callbackQHelpIndexWidgetScrollTo(this->objectName().toUtf8().data(), index.internalPointer(), hint)) { QHelpIndexWidget::scrollTo(index, hint); }; };
	void keyboardSearch(const QString & search) { if (!callbackQHelpIndexWidgetKeyboardSearch(this->objectName().toUtf8().data(), search.toUtf8().data())) { QHelpIndexWidget::keyboardSearch(search); }; };
	void reset() { if (!callbackQHelpIndexWidgetReset(this->objectName().toUtf8().data())) { QHelpIndexWidget::reset(); }; };
	void selectAll() { if (!callbackQHelpIndexWidgetSelectAll(this->objectName().toUtf8().data())) { QHelpIndexWidget::selectAll(); }; };
	void setModel(QAbstractItemModel * model) { if (!callbackQHelpIndexWidgetSetModel(this->objectName().toUtf8().data(), model)) { QHelpIndexWidget::setModel(model); }; };
	void setRootIndex(const QModelIndex & index) { if (!callbackQHelpIndexWidgetSetRootIndex(this->objectName().toUtf8().data(), index.internalPointer())) { QHelpIndexWidget::setRootIndex(index); }; };
	void setSelectionModel(QItemSelectionModel * selectionModel) { if (!callbackQHelpIndexWidgetSetSelectionModel(this->objectName().toUtf8().data(), selectionModel)) { QHelpIndexWidget::setSelectionModel(selectionModel); }; };
	void setupViewport(QWidget * viewport) { if (!callbackQHelpIndexWidgetSetupViewport(this->objectName().toUtf8().data(), viewport)) { QHelpIndexWidget::setupViewport(viewport); }; };
	void setVisible(bool visible) { if (!callbackQHelpIndexWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QHelpIndexWidget::setVisible(visible); }; };
protected:
	void currentChanged(const QModelIndex & current, const QModelIndex & previous) { if (!callbackQHelpIndexWidgetCurrentChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer())) { QHelpIndexWidget::currentChanged(current, previous); }; };
	void dragLeaveEvent(QDragLeaveEvent * e) { if (!callbackQHelpIndexWidgetDragLeaveEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::dragLeaveEvent(e); }; };
	void dragMoveEvent(QDragMoveEvent * e) { if (!callbackQHelpIndexWidgetDragMoveEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::dragMoveEvent(e); }; };
	void dropEvent(QDropEvent * e) { if (!callbackQHelpIndexWidgetDropEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::dropEvent(e); }; };
	void mouseMoveEvent(QMouseEvent * e) { if (!callbackQHelpIndexWidgetMouseMoveEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::mouseMoveEvent(e); }; };
	void mouseReleaseEvent(QMouseEvent * e) { if (!callbackQHelpIndexWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::mouseReleaseEvent(e); }; };
	void paintEvent(QPaintEvent * e) { if (!callbackQHelpIndexWidgetPaintEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::paintEvent(e); }; };
	void resizeEvent(QResizeEvent * e) { if (!callbackQHelpIndexWidgetResizeEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::resizeEvent(e); }; };
	void rowsAboutToBeRemoved(const QModelIndex & parent, int start, int end) { if (!callbackQHelpIndexWidgetRowsAboutToBeRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), start, end)) { QHelpIndexWidget::rowsAboutToBeRemoved(parent, start, end); }; };
	void rowsInserted(const QModelIndex & parent, int start, int end) { if (!callbackQHelpIndexWidgetRowsInserted(this->objectName().toUtf8().data(), parent.internalPointer(), start, end)) { QHelpIndexWidget::rowsInserted(parent, start, end); }; };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { if (!callbackQHelpIndexWidgetSetSelection(this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), command)) { QHelpIndexWidget::setSelection(rect, command); }; };
	void startDrag(Qt::DropActions supportedActions) { if (!callbackQHelpIndexWidgetStartDrag(this->objectName().toUtf8().data(), supportedActions)) { QHelpIndexWidget::startDrag(supportedActions); }; };
	void timerEvent(QTimerEvent * e) { if (!callbackQHelpIndexWidgetTimerEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::timerEvent(e); }; };
	void updateGeometries() { if (!callbackQHelpIndexWidgetUpdateGeometries(this->objectName().toUtf8().data())) { QHelpIndexWidget::updateGeometries(); }; };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { if (!callbackQHelpIndexWidgetCloseEditor(this->objectName().toUtf8().data(), editor, hint)) { QHelpIndexWidget::closeEditor(editor, hint); }; };
	void commitData(QWidget * editor) { if (!callbackQHelpIndexWidgetCommitData(this->objectName().toUtf8().data(), editor)) { QHelpIndexWidget::commitData(editor); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQHelpIndexWidgetDragEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::dragEnterEvent(event); }; };
	void editorDestroyed(QObject * editor) { if (!callbackQHelpIndexWidgetEditorDestroyed(this->objectName().toUtf8().data(), editor)) { QHelpIndexWidget::editorDestroyed(editor); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQHelpIndexWidgetFocusInEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQHelpIndexWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::focusOutEvent(event); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQHelpIndexWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQHelpIndexWidgetKeyPressEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::keyPressEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQHelpIndexWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::mouseDoubleClickEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQHelpIndexWidgetMousePressEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::mousePressEvent(event); }; };
	void contextMenuEvent(QContextMenuEvent * e) { if (!callbackQHelpIndexWidgetContextMenuEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::contextMenuEvent(e); }; };
	void scrollContentsBy(int dx, int dy) { if (!callbackQHelpIndexWidgetScrollContentsBy(this->objectName().toUtf8().data(), dx, dy)) { QHelpIndexWidget::scrollContentsBy(dx, dy); }; };
	void wheelEvent(QWheelEvent * e) { if (!callbackQHelpIndexWidgetWheelEvent(this->objectName().toUtf8().data(), e)) { QHelpIndexWidget::wheelEvent(e); }; };
	void changeEvent(QEvent * ev) { if (!callbackQHelpIndexWidgetChangeEvent(this->objectName().toUtf8().data(), ev)) { QHelpIndexWidget::changeEvent(ev); }; };
	void actionEvent(QActionEvent * event) { if (!callbackQHelpIndexWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::actionEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQHelpIndexWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::enterEvent(event); }; };
	void hideEvent(QHideEvent * event) { if (!callbackQHelpIndexWidgetHideEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::hideEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQHelpIndexWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQHelpIndexWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::moveEvent(event); }; };
	void showEvent(QShowEvent * event) { if (!callbackQHelpIndexWidgetShowEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::showEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQHelpIndexWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::closeEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQHelpIndexWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QHelpIndexWidget::initPainter(painter); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQHelpIndexWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::keyReleaseEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQHelpIndexWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::tabletEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpIndexWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpIndexWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpIndexWidget::customEvent(event); }; };
};

void QHelpIndexWidget_ActivateCurrentItem(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(void* ptr, char* filter, char* wildcard){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

void QHelpIndexWidget_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

void QHelpIndexWidget_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

class MyQHelpSearchEngine: public QHelpSearchEngine {
public:
	void Signal_IndexingFinished() { callbackQHelpSearchEngineIndexingFinished(this->objectName().toUtf8().data()); };
	void Signal_IndexingStarted() { callbackQHelpSearchEngineIndexingStarted(this->objectName().toUtf8().data()); };
	void Signal_SearchingFinished(int hits) { callbackQHelpSearchEngineSearchingFinished(this->objectName().toUtf8().data(), hits); };
	void Signal_SearchingStarted() { callbackQHelpSearchEngineSearchingStarted(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpSearchEngineTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchEngine::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpSearchEngineChildEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchEngine::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpSearchEngineCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchEngine::customEvent(event); }; };
};

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent){
	return new QHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_ConnectIndexingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void* QHelpSearchEngine_QueryWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void* QHelpSearchEngine_ResultWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_ConnectSearchingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

void* QHelpSearchQuery_NewQHelpSearchQuery(){
	return new QHelpSearchQuery();
}

void* QHelpSearchQuery_NewQHelpSearchQuery2(int field, char* wordList){
	return new QHelpSearchQuery(static_cast<QHelpSearchQuery::FieldName>(field), QString(wordList).split(",,,", QString::SkipEmptyParts));
}

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget {
public:
	void Signal_Search() { callbackQHelpSearchQueryWidgetSearch(this->objectName().toUtf8().data()); };
	void setVisible(bool visible) { if (!callbackQHelpSearchQueryWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QHelpSearchQueryWidget::setVisible(visible); }; };
protected:
	void actionEvent(QActionEvent * event) { if (!callbackQHelpSearchQueryWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::actionEvent(event); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQHelpSearchQueryWidgetDragEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQHelpSearchQueryWidgetDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQHelpSearchQueryWidgetDragMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQHelpSearchQueryWidgetDropEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::dropEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQHelpSearchQueryWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::enterEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQHelpSearchQueryWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::focusOutEvent(event); }; };
	void hideEvent(QHideEvent * event) { if (!callbackQHelpSearchQueryWidgetHideEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::hideEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQHelpSearchQueryWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQHelpSearchQueryWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::moveEvent(event); }; };
	void paintEvent(QPaintEvent * event) { if (!callbackQHelpSearchQueryWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::paintEvent(event); }; };
	void showEvent(QShowEvent * event) { if (!callbackQHelpSearchQueryWidgetShowEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::showEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQHelpSearchQueryWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::closeEvent(event); }; };
	void contextMenuEvent(QContextMenuEvent * event) { if (!callbackQHelpSearchQueryWidgetContextMenuEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::contextMenuEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQHelpSearchQueryWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QHelpSearchQueryWidget::initPainter(painter); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQHelpSearchQueryWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQHelpSearchQueryWidgetKeyPressEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQHelpSearchQueryWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQHelpSearchQueryWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQHelpSearchQueryWidgetMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQHelpSearchQueryWidgetMousePressEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQHelpSearchQueryWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::mouseReleaseEvent(event); }; };
	void resizeEvent(QResizeEvent * event) { if (!callbackQHelpSearchQueryWidgetResizeEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::resizeEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQHelpSearchQueryWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::tabletEvent(event); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQHelpSearchQueryWidgetWheelEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpSearchQueryWidgetTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpSearchQueryWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpSearchQueryWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchQueryWidget::customEvent(event); }; };
};

int QHelpSearchQueryWidget_IsCompactMode(void* ptr){
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(void* parent){
	return new QHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(void* ptr){
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DisconnectSearch(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget {
public:
	void Signal_RequestShowLink(const QUrl & link) { callbackQHelpSearchResultWidgetRequestShowLink(this->objectName().toUtf8().data(), new QUrl(link)); };
	void setVisible(bool visible) { if (!callbackQHelpSearchResultWidgetSetVisible(this->objectName().toUtf8().data(), visible)) { QHelpSearchResultWidget::setVisible(visible); }; };
protected:
	void actionEvent(QActionEvent * event) { if (!callbackQHelpSearchResultWidgetActionEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::actionEvent(event); }; };
	void dragEnterEvent(QDragEnterEvent * event) { if (!callbackQHelpSearchResultWidgetDragEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::dragEnterEvent(event); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { if (!callbackQHelpSearchResultWidgetDragLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::dragLeaveEvent(event); }; };
	void dragMoveEvent(QDragMoveEvent * event) { if (!callbackQHelpSearchResultWidgetDragMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::dragMoveEvent(event); }; };
	void dropEvent(QDropEvent * event) { if (!callbackQHelpSearchResultWidgetDropEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::dropEvent(event); }; };
	void enterEvent(QEvent * event) { if (!callbackQHelpSearchResultWidgetEnterEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::enterEvent(event); }; };
	void focusInEvent(QFocusEvent * event) { if (!callbackQHelpSearchResultWidgetFocusInEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::focusInEvent(event); }; };
	void focusOutEvent(QFocusEvent * event) { if (!callbackQHelpSearchResultWidgetFocusOutEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::focusOutEvent(event); }; };
	void hideEvent(QHideEvent * event) { if (!callbackQHelpSearchResultWidgetHideEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::hideEvent(event); }; };
	void leaveEvent(QEvent * event) { if (!callbackQHelpSearchResultWidgetLeaveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::leaveEvent(event); }; };
	void moveEvent(QMoveEvent * event) { if (!callbackQHelpSearchResultWidgetMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::moveEvent(event); }; };
	void paintEvent(QPaintEvent * event) { if (!callbackQHelpSearchResultWidgetPaintEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::paintEvent(event); }; };
	void showEvent(QShowEvent * event) { if (!callbackQHelpSearchResultWidgetShowEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::showEvent(event); }; };
	void closeEvent(QCloseEvent * event) { if (!callbackQHelpSearchResultWidgetCloseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::closeEvent(event); }; };
	void contextMenuEvent(QContextMenuEvent * event) { if (!callbackQHelpSearchResultWidgetContextMenuEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::contextMenuEvent(event); }; };
	void initPainter(QPainter * painter) const { if (!callbackQHelpSearchResultWidgetInitPainter(this->objectName().toUtf8().data(), painter)) { QHelpSearchResultWidget::initPainter(painter); }; };
	void inputMethodEvent(QInputMethodEvent * event) { if (!callbackQHelpSearchResultWidgetInputMethodEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::inputMethodEvent(event); }; };
	void keyPressEvent(QKeyEvent * event) { if (!callbackQHelpSearchResultWidgetKeyPressEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::keyPressEvent(event); }; };
	void keyReleaseEvent(QKeyEvent * event) { if (!callbackQHelpSearchResultWidgetKeyReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::keyReleaseEvent(event); }; };
	void mouseDoubleClickEvent(QMouseEvent * event) { if (!callbackQHelpSearchResultWidgetMouseDoubleClickEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::mouseDoubleClickEvent(event); }; };
	void mouseMoveEvent(QMouseEvent * event) { if (!callbackQHelpSearchResultWidgetMouseMoveEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::mouseMoveEvent(event); }; };
	void mousePressEvent(QMouseEvent * event) { if (!callbackQHelpSearchResultWidgetMousePressEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::mousePressEvent(event); }; };
	void mouseReleaseEvent(QMouseEvent * event) { if (!callbackQHelpSearchResultWidgetMouseReleaseEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::mouseReleaseEvent(event); }; };
	void resizeEvent(QResizeEvent * event) { if (!callbackQHelpSearchResultWidgetResizeEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::resizeEvent(event); }; };
	void tabletEvent(QTabletEvent * event) { if (!callbackQHelpSearchResultWidgetTabletEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::tabletEvent(event); }; };
	void wheelEvent(QWheelEvent * event) { if (!callbackQHelpSearchResultWidgetWheelEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::wheelEvent(event); }; };
	void timerEvent(QTimerEvent * event) { if (!callbackQHelpSearchResultWidgetTimerEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQHelpSearchResultWidgetChildEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQHelpSearchResultWidgetCustomEvent(this->objectName().toUtf8().data(), event)) { QHelpSearchResultWidget::customEvent(event); }; };
};

void* QHelpSearchResultWidget_LinkAt(void* ptr, void* point){
	return new QUrl(static_cast<QHelpSearchResultWidget*>(ptr)->linkAt(*static_cast<QPoint*>(point)));
}

void QHelpSearchResultWidget_ConnectRequestShowLink(void* ptr){
	QObject::connect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_DisconnectRequestShowLink(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(void* ptr){
	static_cast<QHelpSearchResultWidget*>(ptr)->~QHelpSearchResultWidget();
}


#include "qfilesystemmodel.h"
#include <QDir>
#include <QDateTime>
#include <QDate>
#include <QFile>
#include <QString>
#include <QModelIndex>
#include <QFileIconProvider>
#include <QMimeData>
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QFileSystemModel>
#include "_cgo_export.h"

class MyQFileSystemModel: public QFileSystemModel {
public:
void Signal_DirectoryLoaded(const QString & path){callbackQFileSystemModelDirectoryLoaded(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_FileRenamed(const QString & path, const QString & oldName, const QString & newName){callbackQFileSystemModelFileRenamed(this->objectName().toUtf8().data(), path.toUtf8().data(), oldName.toUtf8().data(), newName.toUtf8().data());};
void Signal_RootPathChanged(const QString & newPath){callbackQFileSystemModelRootPathChanged(this->objectName().toUtf8().data(), newPath.toUtf8().data());};
};

int QFileSystemModel_FilePathRole_Type(){
	return QFileSystemModel::FilePathRole;
}

int QFileSystemModel_FileNameRole_Type(){
	return QFileSystemModel::FileNameRole;
}

int QFileSystemModel_FilePermissions_Type(){
	return QFileSystemModel::FilePermissions;
}

int QFileSystemModel_IsReadOnly(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->isReadOnly();
}

int QFileSystemModel_NameFilterDisables(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilterDisables();
}

int QFileSystemModel_ResolveSymlinks(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->resolveSymlinks();
}

int QFileSystemModel_Rmdir(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->rmdir(*static_cast<QModelIndex*>(index));
}

void QFileSystemModel_SetNameFilterDisables(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setNameFilterDisables(enable != 0);
}

void QFileSystemModel_SetReadOnly(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setReadOnly(enable != 0);
}

void QFileSystemModel_SetResolveSymlinks(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setResolveSymlinks(enable != 0);
}

void* QFileSystemModel_NewQFileSystemModel(void* parent){
	return new QFileSystemModel(static_cast<QObject*>(parent));
}

int QFileSystemModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QFileSystemModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void QFileSystemModel_ConnectDirectoryLoaded(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

void QFileSystemModel_DisconnectDirectoryLoaded(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

int QFileSystemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QFileSystemModel_FetchMore(void* ptr, void* parent){
	static_cast<QFileSystemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

char* QFileSystemModel_FileName(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->fileName(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

char* QFileSystemModel_FilePath(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->filePath(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_ConnectFileRenamed(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

void QFileSystemModel_DisconnectFileRenamed(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

int QFileSystemModel_Filter(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->filter();
}

int QFileSystemModel_Flags(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QFileSystemModel_HasChildren(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QFileSystemModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QFileSystemModel_IconProvider(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->iconProvider();
}

void* QFileSystemModel_Index2(void* ptr, char* path, int column){
	return static_cast<QFileSystemModel*>(ptr)->index(QString(path), column).internalPointer();
}

void* QFileSystemModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QFileSystemModel_IsDir(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->isDir(*static_cast<QModelIndex*>(index));
}

void* QFileSystemModel_LastModified(void* ptr, void* index){
	return new QDateTime(static_cast<QFileSystemModel*>(ptr)->lastModified(*static_cast<QModelIndex*>(index)));
}

char* QFileSystemModel_MimeTypes(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void* QFileSystemModel_Mkdir(void* ptr, void* parent, char* name){
	return static_cast<QFileSystemModel*>(ptr)->mkdir(*static_cast<QModelIndex*>(parent), QString(name)).internalPointer();
}

void* QFileSystemModel_MyComputer(void* ptr, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->myComputer(role));
}

char* QFileSystemModel_NameFilters(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilters().join("|").toUtf8().data();
}

void* QFileSystemModel_Parent(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

void* QFileSystemModel_RootDirectory(void* ptr){
	return new QDir(static_cast<QFileSystemModel*>(ptr)->rootDirectory());
}

char* QFileSystemModel_RootPath(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->rootPath().toUtf8().data();
}

void QFileSystemModel_ConnectRootPathChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

void QFileSystemModel_DisconnectRootPathChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

int QFileSystemModel_RowCount(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_SetData(void* ptr, void* idx, void* value, int role){
	return static_cast<QFileSystemModel*>(ptr)->setData(*static_cast<QModelIndex*>(idx), *static_cast<QVariant*>(value), role);
}

void QFileSystemModel_SetFilter(void* ptr, int filters){
	static_cast<QFileSystemModel*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileSystemModel_SetIconProvider(void* ptr, void* provider){
	static_cast<QFileSystemModel*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileSystemModel_SetNameFilters(void* ptr, char* filters){
	static_cast<QFileSystemModel*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void* QFileSystemModel_SetRootPath(void* ptr, char* newPath){
	return static_cast<QFileSystemModel*>(ptr)->setRootPath(QString(newPath)).internalPointer();
}

void QFileSystemModel_Sort(void* ptr, int column, int order){
	static_cast<QFileSystemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QFileSystemModel_SupportedDropActions(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->supportedDropActions();
}

char* QFileSystemModel_Type(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->type(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_DestroyQFileSystemModel(void* ptr){
	static_cast<QFileSystemModel*>(ptr)->~QFileSystemModel();
}

#include "qgraphicsanchor.h"
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSizePolicy>
#include <QGraphicsAnchor>
#include "_cgo_export.h"

class MyQGraphicsAnchor: public QGraphicsAnchor {
public:
};

void QGraphicsAnchor_SetSizePolicy(void* ptr, int policy){
	static_cast<QGraphicsAnchor*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QGraphicsAnchor_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchor*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

int QGraphicsAnchor_SizePolicy(void* ptr){
	return static_cast<QGraphicsAnchor*>(ptr)->sizePolicy();
}

double QGraphicsAnchor_Spacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchor*>(ptr)->spacing());
}

void QGraphicsAnchor_UnsetSpacing(void* ptr){
	static_cast<QGraphicsAnchor*>(ptr)->unsetSpacing();
}

void QGraphicsAnchor_DestroyQGraphicsAnchor(void* ptr){
	static_cast<QGraphicsAnchor*>(ptr)->~QGraphicsAnchor();
}

#include "qmessagebox.h"
#include <QAbstractButton>
#include <QString>
#include <QObject>
#include <QWidget>
#include <QCheckBox>
#include <QPixmap>
#include <QPushButton>
#include <QMetaObject>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMessageBox>
#include "_cgo_export.h"

class MyQMessageBox: public QMessageBox {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQMessageBoxButtonClicked(this->objectName().toUtf8().data(), button);};
};

int QMessageBox_ButtonMask_Type(){
	return QMessageBox::ButtonMask;
}

char* QMessageBox_DetailedText(void* ptr){
	return static_cast<QMessageBox*>(ptr)->detailedText().toUtf8().data();
}

int QMessageBox_Icon(void* ptr){
	return static_cast<QMessageBox*>(ptr)->icon();
}

char* QMessageBox_InformativeText(void* ptr){
	return static_cast<QMessageBox*>(ptr)->informativeText().toUtf8().data();
}

void QMessageBox_SetDetailedText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setDetailedText(QString(text));
}

void QMessageBox_SetIcon(void* ptr, int v){
	static_cast<QMessageBox*>(ptr)->setIcon(static_cast<QMessageBox::Icon>(v));
}

void QMessageBox_SetIconPixmap(void* ptr, void* pixmap){
	static_cast<QMessageBox*>(ptr)->setIconPixmap(*static_cast<QPixmap*>(pixmap));
}

void QMessageBox_SetInformativeText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setInformativeText(QString(text));
}

void QMessageBox_SetStandardButtons(void* ptr, int buttons){
	static_cast<QMessageBox*>(ptr)->setStandardButtons(static_cast<QMessageBox::StandardButton>(buttons));
}

void QMessageBox_SetText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setText(QString(text));
}

void QMessageBox_SetTextFormat(void* ptr, int format){
	static_cast<QMessageBox*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(format));
}

void QMessageBox_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QMessageBox*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

int QMessageBox_StandardButtons(void* ptr){
	return static_cast<QMessageBox*>(ptr)->standardButtons();
}

char* QMessageBox_Text(void* ptr){
	return static_cast<QMessageBox*>(ptr)->text().toUtf8().data();
}

int QMessageBox_TextFormat(void* ptr){
	return static_cast<QMessageBox*>(ptr)->textFormat();
}

int QMessageBox_TextInteractionFlags(void* ptr){
	return static_cast<QMessageBox*>(ptr)->textInteractionFlags();
}

void* QMessageBox_NewQMessageBox2(int icon, char* title, char* text, int buttons, void* parent, int f){
	return new QMessageBox(static_cast<QMessageBox::Icon>(icon), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QMessageBox_NewQMessageBox(void* parent){
	return new QMessageBox(static_cast<QWidget*>(parent));
}

void QMessageBox_QMessageBox_About(void* parent, char* title, char* text){
	QMessageBox::about(static_cast<QWidget*>(parent), QString(title), QString(text));
}

void QMessageBox_QMessageBox_AboutQt(void* parent, char* title){
	QMessageBox::aboutQt(static_cast<QWidget*>(parent), QString(title));
}

void* QMessageBox_AddButton3(void* ptr, int button){
	return static_cast<QMessageBox*>(ptr)->addButton(static_cast<QMessageBox::StandardButton>(button));
}

void* QMessageBox_AddButton2(void* ptr, char* text, int role){
	return static_cast<QMessageBox*>(ptr)->addButton(QString(text), static_cast<QMessageBox::ButtonRole>(role));
}

void QMessageBox_AddButton(void* ptr, void* button, int role){
	static_cast<QMessageBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QMessageBox::ButtonRole>(role));
}

void* QMessageBox_Button(void* ptr, int which){
	return static_cast<QMessageBox*>(ptr)->button(static_cast<QMessageBox::StandardButton>(which));
}

void QMessageBox_ConnectButtonClicked(void* ptr){
	QObject::connect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

void QMessageBox_DisconnectButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

int QMessageBox_ButtonRole(void* ptr, void* button){
	return static_cast<QMessageBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

void* QMessageBox_CheckBox(void* ptr){
	return static_cast<QMessageBox*>(ptr)->checkBox();
}

void* QMessageBox_ClickedButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->clickedButton();
}

int QMessageBox_QMessageBox_Critical(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::critical(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void* QMessageBox_DefaultButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->defaultButton();
}

void* QMessageBox_EscapeButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->escapeButton();
}

int QMessageBox_Exec(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QMessageBox*>(ptr), "exec");
}

int QMessageBox_QMessageBox_Information(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::information(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_Open(void* ptr, void* receiver, char* member){
	static_cast<QMessageBox*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QMessageBox_QMessageBox_Question(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::question(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_RemoveButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetCheckBox(void* ptr, void* cb){
	static_cast<QMessageBox*>(ptr)->setCheckBox(static_cast<QCheckBox*>(cb));
}

void QMessageBox_SetDefaultButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QPushButton*>(button));
}

void QMessageBox_SetDefaultButton2(void* ptr, int button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetEscapeButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetEscapeButton2(void* ptr, int button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetVisible(void* ptr, int visible){
	static_cast<QMessageBox*>(ptr)->setVisible(visible != 0);
}

void QMessageBox_SetWindowModality(void* ptr, int windowModality){
	static_cast<QMessageBox*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QMessageBox_SetWindowTitle(void* ptr, char* title){
	static_cast<QMessageBox*>(ptr)->setWindowTitle(QString(title));
}

int QMessageBox_StandardButton(void* ptr, void* button){
	return static_cast<QMessageBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

int QMessageBox_QMessageBox_Warning(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::warning(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_DestroyQMessageBox(void* ptr){
	static_cast<QMessageBox*>(ptr)->~QMessageBox();
}

#include "qdatetimeedit.h"
#include <QEvent>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QCalendarWidget>
#include <QDate>
#include <QVariant>
#include <QObject>
#include <QTime>
#include <QDateTime>
#include <QMetaObject>
#include <QDateTimeEdit>
#include "_cgo_export.h"

class MyQDateTimeEdit: public QDateTimeEdit {
public:
void Signal_DateTimeChanged(const QDateTime & datetime){callbackQDateTimeEditDateTimeChanged(this->objectName().toUtf8().data(), new QDateTime(datetime));};
};

void* QDateTimeEdit_NewQDateTimeEdit3(void* date, void* parent){
	return new QDateTimeEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_NewQDateTimeEdit4(void* time, void* parent){
	return new QDateTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

int QDateTimeEdit_CalendarPopup(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarPopup();
}

void QDateTimeEdit_ClearMaximumDate(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDate();
}

void QDateTimeEdit_ClearMaximumDateTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumDateTime();
}

void QDateTimeEdit_ClearMaximumTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMaximumTime();
}

void QDateTimeEdit_ClearMinimumDate(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDate();
}

void QDateTimeEdit_ClearMinimumDateTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumDateTime();
}

void QDateTimeEdit_ClearMinimumTime(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clearMinimumTime();
}

int QDateTimeEdit_CurrentSection(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSection();
}

int QDateTimeEdit_CurrentSectionIndex(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->currentSectionIndex();
}

void* QDateTimeEdit_DateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->dateTime());
}

char* QDateTimeEdit_DisplayFormat(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayFormat().toUtf8().data();
}

int QDateTimeEdit_DisplayedSections(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->displayedSections();
}

void* QDateTimeEdit_MaximumDateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->maximumDateTime());
}

void* QDateTimeEdit_MinimumDateTime(void* ptr){
	return new QDateTime(static_cast<QDateTimeEdit*>(ptr)->minimumDateTime());
}

int QDateTimeEdit_SectionCount(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->sectionCount();
}

char* QDateTimeEdit_SectionText(void* ptr, int section){
	return static_cast<QDateTimeEdit*>(ptr)->sectionText(static_cast<QDateTimeEdit::Section>(section)).toUtf8().data();
}

void QDateTimeEdit_SetCalendarPopup(void* ptr, int enable){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarPopup(enable != 0);
}

void QDateTimeEdit_SetCurrentSection(void* ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetCurrentSectionIndex(void* ptr, int index){
	static_cast<QDateTimeEdit*>(ptr)->setCurrentSectionIndex(index);
}

void QDateTimeEdit_SetDate(void* ptr, void* date){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QDateTimeEdit_SetDateTime(void* ptr, void* dateTime){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setDateTime", Q_ARG(QDateTime, *static_cast<QDateTime*>(dateTime)));
}

void QDateTimeEdit_SetDisplayFormat(void* ptr, char* format){
	static_cast<QDateTimeEdit*>(ptr)->setDisplayFormat(QString(format));
}

void QDateTimeEdit_SetMaximumDate(void* ptr, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDate(*static_cast<QDate*>(max));
}

void QDateTimeEdit_SetMaximumDateTime(void* ptr, void* dt){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMaximumTime(void* ptr, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setMaximumTime(*static_cast<QTime*>(max));
}

void QDateTimeEdit_SetMinimumDate(void* ptr, void* min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDate(*static_cast<QDate*>(min));
}

void QDateTimeEdit_SetMinimumDateTime(void* ptr, void* dt){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumDateTime(*static_cast<QDateTime*>(dt));
}

void QDateTimeEdit_SetMinimumTime(void* ptr, void* min){
	static_cast<QDateTimeEdit*>(ptr)->setMinimumTime(*static_cast<QTime*>(min));
}

void QDateTimeEdit_SetTime(void* ptr, void* time){
	QMetaObject::invokeMethod(static_cast<QDateTimeEdit*>(ptr), "setTime", Q_ARG(QTime, *static_cast<QTime*>(time)));
}

void QDateTimeEdit_SetTimeSpec(void* ptr, int spec){
	static_cast<QDateTimeEdit*>(ptr)->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}

int QDateTimeEdit_TimeSpec(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->timeSpec();
}

void* QDateTimeEdit_NewQDateTimeEdit(void* parent){
	return new QDateTimeEdit(static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_NewQDateTimeEdit2(void* datetime, void* parent){
	return new QDateTimeEdit(*static_cast<QDateTime*>(datetime), static_cast<QWidget*>(parent));
}

void* QDateTimeEdit_CalendarWidget(void* ptr){
	return static_cast<QDateTimeEdit*>(ptr)->calendarWidget();
}

void QDateTimeEdit_Clear(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->clear();
}

void QDateTimeEdit_ConnectDateTimeChanged(void* ptr){
	QObject::connect(static_cast<QDateTimeEdit*>(ptr), static_cast<void (QDateTimeEdit::*)(const QDateTime &)>(&QDateTimeEdit::dateTimeChanged), static_cast<MyQDateTimeEdit*>(ptr), static_cast<void (MyQDateTimeEdit::*)(const QDateTime &)>(&MyQDateTimeEdit::Signal_DateTimeChanged));;
}

void QDateTimeEdit_DisconnectDateTimeChanged(void* ptr){
	QObject::disconnect(static_cast<QDateTimeEdit*>(ptr), static_cast<void (QDateTimeEdit::*)(const QDateTime &)>(&QDateTimeEdit::dateTimeChanged), static_cast<MyQDateTimeEdit*>(ptr), static_cast<void (MyQDateTimeEdit::*)(const QDateTime &)>(&MyQDateTimeEdit::Signal_DateTimeChanged));;
}

int QDateTimeEdit_Event(void* ptr, void* event){
	return static_cast<QDateTimeEdit*>(ptr)->event(static_cast<QEvent*>(event));
}

int QDateTimeEdit_SectionAt(void* ptr, int index){
	return static_cast<QDateTimeEdit*>(ptr)->sectionAt(index);
}

void QDateTimeEdit_SetCalendarWidget(void* ptr, void* calendarWidget){
	static_cast<QDateTimeEdit*>(ptr)->setCalendarWidget(static_cast<QCalendarWidget*>(calendarWidget));
}

void QDateTimeEdit_SetDateRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setDateRange(*static_cast<QDate*>(min), *static_cast<QDate*>(max));
}

void QDateTimeEdit_SetDateTimeRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setDateTimeRange(*static_cast<QDateTime*>(min), *static_cast<QDateTime*>(max));
}

void QDateTimeEdit_SetSelectedSection(void* ptr, int section){
	static_cast<QDateTimeEdit*>(ptr)->setSelectedSection(static_cast<QDateTimeEdit::Section>(section));
}

void QDateTimeEdit_SetTimeRange(void* ptr, void* min, void* max){
	static_cast<QDateTimeEdit*>(ptr)->setTimeRange(*static_cast<QTime*>(min), *static_cast<QTime*>(max));
}

void QDateTimeEdit_StepBy(void* ptr, int steps){
	static_cast<QDateTimeEdit*>(ptr)->stepBy(steps);
}

void QDateTimeEdit_DestroyQDateTimeEdit(void* ptr){
	static_cast<QDateTimeEdit*>(ptr)->~QDateTimeEdit();
}

#include "qfontdialog.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QFontDialog>
#include "_cgo_export.h"

class MyQFontDialog: public QFontDialog {
public:
};

int QFontDialog_Options(void* ptr){
	return static_cast<QFontDialog*>(ptr)->options();
}

void QFontDialog_SetOptions(void* ptr, int options){
	static_cast<QFontDialog*>(ptr)->setOptions(static_cast<QFontDialog::FontDialogOption>(options));
}

void* QFontDialog_NewQFontDialog(void* parent){
	return new QFontDialog(static_cast<QWidget*>(parent));
}

void* QFontDialog_NewQFontDialog2(void* initial, void* parent){
	return new QFontDialog(*static_cast<QFont*>(initial), static_cast<QWidget*>(parent));
}

void QFontDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QFontDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QFontDialog_SetCurrentFont(void* ptr, void* font){
	static_cast<QFontDialog*>(ptr)->setCurrentFont(*static_cast<QFont*>(font));
}

void QFontDialog_SetOption(void* ptr, int option, int on){
	static_cast<QFontDialog*>(ptr)->setOption(static_cast<QFontDialog::FontDialogOption>(option), on != 0);
}

void QFontDialog_SetVisible(void* ptr, int visible){
	static_cast<QFontDialog*>(ptr)->setVisible(visible != 0);
}

int QFontDialog_TestOption(void* ptr, int option){
	return static_cast<QFontDialog*>(ptr)->testOption(static_cast<QFontDialog::FontDialogOption>(option));
}

#include "qundogroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QUndoStack>
#include <QMetaObject>
#include <QUndoGroup>
#include "_cgo_export.h"

class MyQUndoGroup: public QUndoGroup {
public:
void Signal_ActiveStackChanged(QUndoStack * stack){callbackQUndoGroupActiveStackChanged(this->objectName().toUtf8().data(), stack);};
void Signal_CanRedoChanged(bool canRedo){callbackQUndoGroupCanRedoChanged(this->objectName().toUtf8().data(), canRedo);};
void Signal_CanUndoChanged(bool canUndo){callbackQUndoGroupCanUndoChanged(this->objectName().toUtf8().data(), canUndo);};
void Signal_CleanChanged(bool clean){callbackQUndoGroupCleanChanged(this->objectName().toUtf8().data(), clean);};
void Signal_IndexChanged(int idx){callbackQUndoGroupIndexChanged(this->objectName().toUtf8().data(), idx);};
void Signal_RedoTextChanged(const QString & redoText){callbackQUndoGroupRedoTextChanged(this->objectName().toUtf8().data(), redoText.toUtf8().data());};
void Signal_UndoTextChanged(const QString & undoText){callbackQUndoGroupUndoTextChanged(this->objectName().toUtf8().data(), undoText.toUtf8().data());};
};

void* QUndoGroup_NewQUndoGroup(void* parent){
	return new QUndoGroup(static_cast<QObject*>(parent));
}

void* QUndoGroup_ActiveStack(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->activeStack();
}

void QUndoGroup_ConnectActiveStackChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(QUndoStack *)>(&QUndoGroup::activeStackChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(QUndoStack *)>(&MyQUndoGroup::Signal_ActiveStackChanged));;
}

void QUndoGroup_DisconnectActiveStackChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(QUndoStack *)>(&QUndoGroup::activeStackChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(QUndoStack *)>(&MyQUndoGroup::Signal_ActiveStackChanged));;
}

void QUndoGroup_AddStack(void* ptr, void* stack){
	static_cast<QUndoGroup*>(ptr)->addStack(static_cast<QUndoStack*>(stack));
}

int QUndoGroup_CanRedo(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->canRedo();
}

void QUndoGroup_ConnectCanRedoChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canRedoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanRedoChanged));;
}

void QUndoGroup_DisconnectCanRedoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canRedoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanRedoChanged));;
}

int QUndoGroup_CanUndo(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->canUndo();
}

void QUndoGroup_ConnectCanUndoChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canUndoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanUndoChanged));;
}

void QUndoGroup_DisconnectCanUndoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canUndoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanUndoChanged));;
}

void QUndoGroup_ConnectCleanChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::cleanChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CleanChanged));;
}

void QUndoGroup_DisconnectCleanChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::cleanChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CleanChanged));;
}

void* QUndoGroup_CreateRedoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoGroup*>(ptr)->createRedoAction(static_cast<QObject*>(parent), QString(prefix));
}

void* QUndoGroup_CreateUndoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoGroup*>(ptr)->createUndoAction(static_cast<QObject*>(parent), QString(prefix));
}

void QUndoGroup_ConnectIndexChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(int)>(&QUndoGroup::indexChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(int)>(&MyQUndoGroup::Signal_IndexChanged));;
}

void QUndoGroup_DisconnectIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(int)>(&QUndoGroup::indexChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(int)>(&MyQUndoGroup::Signal_IndexChanged));;
}

int QUndoGroup_IsClean(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->isClean();
}

void QUndoGroup_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "redo");
}

char* QUndoGroup_RedoText(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->redoText().toUtf8().data();
}

void QUndoGroup_ConnectRedoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::redoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_RedoTextChanged));;
}

void QUndoGroup_DisconnectRedoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::redoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_RedoTextChanged));;
}

void QUndoGroup_RemoveStack(void* ptr, void* stack){
	static_cast<QUndoGroup*>(ptr)->removeStack(static_cast<QUndoStack*>(stack));
}

void QUndoGroup_SetActiveStack(void* ptr, void* stack){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "setActiveStack", Q_ARG(QUndoStack*, static_cast<QUndoStack*>(stack)));
}

void QUndoGroup_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "undo");
}

char* QUndoGroup_UndoText(void* ptr){
	return static_cast<QUndoGroup*>(ptr)->undoText().toUtf8().data();
}

void QUndoGroup_ConnectUndoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::undoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_UndoTextChanged));;
}

void QUndoGroup_DisconnectUndoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::undoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_UndoTextChanged));;
}

void QUndoGroup_DestroyQUndoGroup(void* ptr){
	static_cast<QUndoGroup*>(ptr)->~QUndoGroup();
}

#include "qtimeedit.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTime>
#include <QWidget>
#include <QTimeEdit>
#include "_cgo_export.h"

class MyQTimeEdit: public QTimeEdit {
public:
};

void* QTimeEdit_NewQTimeEdit(void* parent){
	return new QTimeEdit(static_cast<QWidget*>(parent));
}

void* QTimeEdit_NewQTimeEdit2(void* time, void* parent){
	return new QTimeEdit(*static_cast<QTime*>(time), static_cast<QWidget*>(parent));
}

void QTimeEdit_DestroyQTimeEdit(void* ptr){
	static_cast<QTimeEdit*>(ptr)->~QTimeEdit();
}

#include "qboxlayout.h"
#include <QString>
#include <QVariant>
#include <QWidget>
#include <QLayoutItem>
#include <QRect>
#include <QUrl>
#include <QModelIndex>
#include <QLayout>
#include <QSpacerItem>
#include <QBoxLayout>
#include "_cgo_export.h"

class MyQBoxLayout: public QBoxLayout {
public:
};

int QBoxLayout_Direction(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->direction();
}

void* QBoxLayout_NewQBoxLayout(int dir, void* parent){
	return new QBoxLayout(static_cast<QBoxLayout::Direction>(dir), static_cast<QWidget*>(parent));
}

void QBoxLayout_AddItem(void* ptr, void* item){
	static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QBoxLayout_AddLayout(void* ptr, void* layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_AddSpacerItem(void* ptr, void* spacerItem){
	static_cast<QBoxLayout*>(ptr)->addSpacerItem(static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_AddSpacing(void* ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addSpacing(size);
}

void QBoxLayout_AddStretch(void* ptr, int stretch){
	static_cast<QBoxLayout*>(ptr)->addStretch(stretch);
}

void QBoxLayout_AddStrut(void* ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addStrut(size);
}

void QBoxLayout_AddWidget(void* ptr, void* widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

int QBoxLayout_Count(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->count();
}

int QBoxLayout_ExpandingDirections(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->expandingDirections();
}

int QBoxLayout_HasHeightForWidth(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->hasHeightForWidth();
}

int QBoxLayout_HeightForWidth(void* ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->heightForWidth(w);
}

void QBoxLayout_InsertItem(void* ptr, int index, void* item){
	static_cast<QBoxLayout*>(ptr)->insertItem(index, static_cast<QLayoutItem*>(item));
}

void QBoxLayout_InsertLayout(void* ptr, int index, void* layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertLayout(index, static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_InsertSpacerItem(void* ptr, int index, void* spacerItem){
	static_cast<QBoxLayout*>(ptr)->insertSpacerItem(index, static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_InsertSpacing(void* ptr, int index, int size){
	static_cast<QBoxLayout*>(ptr)->insertSpacing(index, size);
}

void QBoxLayout_InsertStretch(void* ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertStretch(index, stretch);
}

void QBoxLayout_InsertWidget(void* ptr, int index, void* widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

void QBoxLayout_Invalidate(void* ptr){
	static_cast<QBoxLayout*>(ptr)->invalidate();
}

void* QBoxLayout_ItemAt(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->itemAt(index);
}

int QBoxLayout_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->minimumHeightForWidth(w);
}

void QBoxLayout_SetDirection(void* ptr, int direction){
	static_cast<QBoxLayout*>(ptr)->setDirection(static_cast<QBoxLayout::Direction>(direction));
}

void QBoxLayout_SetGeometry(void* ptr, void* r){
	static_cast<QBoxLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QBoxLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QBoxLayout*>(ptr)->setSpacing(spacing);
}

void QBoxLayout_SetStretch(void* ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->setStretch(index, stretch);
}

int QBoxLayout_SetStretchFactor2(void* ptr, void* layout, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QLayout*>(layout), stretch);
}

int QBoxLayout_SetStretchFactor(void* ptr, void* widget, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QWidget*>(widget), stretch);
}

int QBoxLayout_Spacing(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->spacing();
}

int QBoxLayout_Stretch(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->stretch(index);
}

void* QBoxLayout_TakeAt(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->takeAt(index);
}

void QBoxLayout_DestroyQBoxLayout(void* ptr){
	static_cast<QBoxLayout*>(ptr)->~QBoxLayout();
}

#include "qstyleoptionfocusrect.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionFocusRect>
#include "_cgo_export.h"

class MyQStyleOptionFocusRect: public QStyleOptionFocusRect {
public:
};

void* QStyleOptionFocusRect_NewQStyleOptionFocusRect(){
	return new QStyleOptionFocusRect();
}

void* QStyleOptionFocusRect_NewQStyleOptionFocusRect2(void* other){
	return new QStyleOptionFocusRect(*static_cast<QStyleOptionFocusRect*>(other));
}

#include "qgraphicsrectitem.h"
#include <QStyleOption>
#include <QStyle>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsItem>
#include <QRect>
#include <QRectF>
#include <QPainter>
#include <QString>
#include <QModelIndex>
#include <QPointF>
#include <QGraphicsRectItem>
#include "_cgo_export.h"

class MyQGraphicsRectItem: public QGraphicsRectItem {
public:
};

void QGraphicsRectItem_SetRect(void* ptr, void* rectangle){
	static_cast<QGraphicsRectItem*>(ptr)->setRect(*static_cast<QRectF*>(rectangle));
}

int QGraphicsRectItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsRectItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsRectItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsRectItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsRectItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsRectItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsRectItem_SetRect2(void* ptr, double x, double y, double width, double height){
	static_cast<QGraphicsRectItem*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

int QGraphicsRectItem_Type(void* ptr){
	return static_cast<QGraphicsRectItem*>(ptr)->type();
}

void QGraphicsRectItem_DestroyQGraphicsRectItem(void* ptr){
	static_cast<QGraphicsRectItem*>(ptr)->~QGraphicsRectItem();
}

#include "qwhatsthis.h"
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QWhatsThis>
#include "_cgo_export.h"

class MyQWhatsThis: public QWhatsThis {
public:
};

void* QWhatsThis_QWhatsThis_CreateAction(void* parent){
	return QWhatsThis::createAction(static_cast<QObject*>(parent));
}

void QWhatsThis_QWhatsThis_EnterWhatsThisMode(){
	QWhatsThis::enterWhatsThisMode();
}

void QWhatsThis_QWhatsThis_HideText(){
	QWhatsThis::hideText();
}

int QWhatsThis_QWhatsThis_InWhatsThisMode(){
	return QWhatsThis::inWhatsThisMode();
}

void QWhatsThis_QWhatsThis_LeaveWhatsThisMode(){
	QWhatsThis::leaveWhatsThisMode();
}

void QWhatsThis_QWhatsThis_ShowText(void* pos, char* text, void* w){
	QWhatsThis::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w));
}

#include "qstyleoptionrubberband.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionRubberBand>
#include "_cgo_export.h"

class MyQStyleOptionRubberBand: public QStyleOptionRubberBand {
public:
};

void* QStyleOptionRubberBand_NewQStyleOptionRubberBand(){
	return new QStyleOptionRubberBand();
}

void* QStyleOptionRubberBand_NewQStyleOptionRubberBand2(void* other){
	return new QStyleOptionRubberBand(*static_cast<QStyleOptionRubberBand*>(other));
}

#include "qdial.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QString>
#include <QDial>
#include "_cgo_export.h"

class MyQDial: public QDial {
public:
};

int QDial_NotchSize(void* ptr){
	return static_cast<QDial*>(ptr)->notchSize();
}

double QDial_NotchTarget(void* ptr){
	return static_cast<double>(static_cast<QDial*>(ptr)->notchTarget());
}

int QDial_NotchesVisible(void* ptr){
	return static_cast<QDial*>(ptr)->notchesVisible();
}

void QDial_SetNotchesVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setNotchesVisible", Q_ARG(bool, visible != 0));
}

void QDial_SetWrapping(void* ptr, int on){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setWrapping", Q_ARG(bool, on != 0));
}

int QDial_Wrapping(void* ptr){
	return static_cast<QDial*>(ptr)->wrapping();
}

void* QDial_NewQDial(void* parent){
	return new QDial(static_cast<QWidget*>(parent));
}

void QDial_DestroyQDial(void* ptr){
	static_cast<QDial*>(ptr)->~QDial();
}

#include "qlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QLayoutItem>
#include <QRect>
#include <QMargins>
#include <QLayout>
#include "_cgo_export.h"

class MyQLayout: public QLayout {
public:
};

void QLayout_SetSizeConstraint(void* ptr, int v){
	static_cast<QLayout*>(ptr)->setSizeConstraint(static_cast<QLayout::SizeConstraint>(v));
}

void QLayout_SetSpacing(void* ptr, int v){
	static_cast<QLayout*>(ptr)->setSpacing(v);
}

int QLayout_SizeConstraint(void* ptr){
	return static_cast<QLayout*>(ptr)->sizeConstraint();
}

int QLayout_Spacing(void* ptr){
	return static_cast<QLayout*>(ptr)->spacing();
}

int QLayout_Activate(void* ptr){
	return static_cast<QLayout*>(ptr)->activate();
}

void QLayout_AddItem(void* ptr, void* item){
	static_cast<QLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QLayout_AddWidget(void* ptr, void* w){
	static_cast<QLayout*>(ptr)->addWidget(static_cast<QWidget*>(w));
}

int QLayout_ControlTypes(void* ptr){
	return static_cast<QLayout*>(ptr)->controlTypes();
}

int QLayout_Count(void* ptr){
	return static_cast<QLayout*>(ptr)->count();
}

int QLayout_ExpandingDirections(void* ptr){
	return static_cast<QLayout*>(ptr)->expandingDirections();
}

void QLayout_GetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

int QLayout_IndexOf(void* ptr, void* widget){
	return static_cast<QLayout*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

void QLayout_Invalidate(void* ptr){
	static_cast<QLayout*>(ptr)->invalidate();
}

int QLayout_IsEmpty(void* ptr){
	return static_cast<QLayout*>(ptr)->isEmpty();
}

int QLayout_IsEnabled(void* ptr){
	return static_cast<QLayout*>(ptr)->isEnabled();
}

void* QLayout_ItemAt(void* ptr, int index){
	return static_cast<QLayout*>(ptr)->itemAt(index);
}

void* QLayout_Layout(void* ptr){
	return static_cast<QLayout*>(ptr)->layout();
}

void* QLayout_MenuBar(void* ptr){
	return static_cast<QLayout*>(ptr)->menuBar();
}

void* QLayout_ParentWidget(void* ptr){
	return static_cast<QLayout*>(ptr)->parentWidget();
}

void QLayout_RemoveItem(void* ptr, void* item){
	static_cast<QLayout*>(ptr)->removeItem(static_cast<QLayoutItem*>(item));
}

void QLayout_RemoveWidget(void* ptr, void* widget){
	static_cast<QLayout*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void* QLayout_ReplaceWidget(void* ptr, void* from, void* to, int options){
	return static_cast<QLayout*>(ptr)->replaceWidget(static_cast<QWidget*>(from), static_cast<QWidget*>(to), static_cast<Qt::FindChildOption>(options));
}

int QLayout_SetAlignment2(void* ptr, void* l, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QLayout*>(l), static_cast<Qt::AlignmentFlag>(alignment));
}

int QLayout_SetAlignment(void* ptr, void* w, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QWidget*>(w), static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayout_SetContentsMargins2(void* ptr, void* margins){
	static_cast<QLayout*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QLayout_SetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QLayout_SetEnabled(void* ptr, int enable){
	static_cast<QLayout*>(ptr)->setEnabled(enable != 0);
}

void QLayout_SetGeometry(void* ptr, void* r){
	static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QLayout_SetMenuBar(void* ptr, void* widget){
	static_cast<QLayout*>(ptr)->setMenuBar(static_cast<QWidget*>(widget));
}

void* QLayout_TakeAt(void* ptr, int index){
	return static_cast<QLayout*>(ptr)->takeAt(index);
}

void QLayout_Update(void* ptr){
	static_cast<QLayout*>(ptr)->update();
}

#include "qstylefactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleFactory>
#include "_cgo_export.h"

class MyQStyleFactory: public QStyleFactory {
public:
};

void* QStyleFactory_QStyleFactory_Create(char* key){
	return QStyleFactory::create(QString(key));
}

char* QStyleFactory_QStyleFactory_Keys(){
	return QStyleFactory::keys().join("|").toUtf8().data();
}

#include "qmdisubwindow.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QMenu>
#include <QMdiSubWindow>
#include "_cgo_export.h"

class MyQMdiSubWindow: public QMdiSubWindow {
public:
void Signal_AboutToActivate(){callbackQMdiSubWindowAboutToActivate(this->objectName().toUtf8().data());};
void Signal_WindowStateChanged(Qt::WindowStates oldState, Qt::WindowStates newState){callbackQMdiSubWindowWindowStateChanged(this->objectName().toUtf8().data(), oldState, newState);};
};

int QMdiSubWindow_KeyboardPageStep(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardPageStep();
}

int QMdiSubWindow_KeyboardSingleStep(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->keyboardSingleStep();
}

void QMdiSubWindow_SetKeyboardPageStep(void* ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardPageStep(step);
}

void QMdiSubWindow_SetKeyboardSingleStep(void* ptr, int step){
	static_cast<QMdiSubWindow*>(ptr)->setKeyboardSingleStep(step);
}

void* QMdiSubWindow_NewQMdiSubWindow(void* parent, int flags){
	return new QMdiSubWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMdiSubWindow_ConnectAboutToActivate(void* ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

void QMdiSubWindow_DisconnectAboutToActivate(void* ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)()>(&QMdiSubWindow::aboutToActivate), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)()>(&MyQMdiSubWindow::Signal_AboutToActivate));;
}

int QMdiSubWindow_IsShaded(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->isShaded();
}

void* QMdiSubWindow_MdiArea(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->mdiArea();
}

void QMdiSubWindow_SetOption(void* ptr, int option, int on){
	static_cast<QMdiSubWindow*>(ptr)->setOption(static_cast<QMdiSubWindow::SubWindowOption>(option), on != 0);
}

void QMdiSubWindow_SetSystemMenu(void* ptr, void* systemMenu){
	static_cast<QMdiSubWindow*>(ptr)->setSystemMenu(static_cast<QMenu*>(systemMenu));
}

void QMdiSubWindow_SetWidget(void* ptr, void* widget){
	static_cast<QMdiSubWindow*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QMdiSubWindow_ShowShaded(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showShaded");
}

void QMdiSubWindow_ShowSystemMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiSubWindow*>(ptr), "showSystemMenu");
}

void* QMdiSubWindow_SystemMenu(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->systemMenu();
}

int QMdiSubWindow_TestOption(void* ptr, int option){
	return static_cast<QMdiSubWindow*>(ptr)->testOption(static_cast<QMdiSubWindow::SubWindowOption>(option));
}

void* QMdiSubWindow_Widget(void* ptr){
	return static_cast<QMdiSubWindow*>(ptr)->widget();
}

void QMdiSubWindow_ConnectWindowStateChanged(void* ptr){
	QObject::connect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DisconnectWindowStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMdiSubWindow*>(ptr), static_cast<void (QMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&QMdiSubWindow::windowStateChanged), static_cast<MyQMdiSubWindow*>(ptr), static_cast<void (MyQMdiSubWindow::*)(Qt::WindowStates, Qt::WindowStates)>(&MyQMdiSubWindow::Signal_WindowStateChanged));;
}

void QMdiSubWindow_DestroyQMdiSubWindow(void* ptr){
	static_cast<QMdiSubWindow*>(ptr)->~QMdiSubWindow();
}

#include "qstyleoptionheader.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionHeader>
#include "_cgo_export.h"

class MyQStyleOptionHeader: public QStyleOptionHeader {
public:
};

void* QStyleOptionHeader_NewQStyleOptionHeader(){
	return new QStyleOptionHeader();
}

void* QStyleOptionHeader_NewQStyleOptionHeader2(void* other){
	return new QStyleOptionHeader(*static_cast<QStyleOptionHeader*>(other));
}

#include "qmenubar.h"
#include <QIcon>
#include <QMenu>
#include <QVariant>
#include <QUrl>
#include <QPoint>
#include <QObject>
#include <QWidget>
#include <QAction>
#include <QString>
#include <QModelIndex>
#include <QMetaObject>
#include <QMenuBar>
#include "_cgo_export.h"

class MyQMenuBar: public QMenuBar {
public:
void Signal_Hovered(QAction * action){callbackQMenuBarHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuBarTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenuBar_IsDefaultUp(void* ptr){
	return static_cast<QMenuBar*>(ptr)->isDefaultUp();
}

int QMenuBar_IsNativeMenuBar(void* ptr){
	return static_cast<QMenuBar*>(ptr)->isNativeMenuBar();
}

void QMenuBar_SetCornerWidget(void* ptr, void* widget, int corner){
	static_cast<QMenuBar*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QMenuBar_SetDefaultUp(void* ptr, int v){
	static_cast<QMenuBar*>(ptr)->setDefaultUp(v != 0);
}

void QMenuBar_SetNativeMenuBar(void* ptr, int nativeMenuBar){
	static_cast<QMenuBar*>(ptr)->setNativeMenuBar(nativeMenuBar != 0);
}

void* QMenuBar_NewQMenuBar(void* parent){
	return new QMenuBar(static_cast<QWidget*>(parent));
}

void* QMenuBar_ActionAt(void* ptr, void* pt){
	return static_cast<QMenuBar*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

void* QMenuBar_ActiveAction(void* ptr){
	return static_cast<QMenuBar*>(ptr)->activeAction();
}

void* QMenuBar_AddAction(void* ptr, char* text){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text));
}

void* QMenuBar_AddAction2(void* ptr, char* text, void* receiver, char* member){
	return static_cast<QMenuBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QMenuBar_AddMenu(void* ptr, void* menu){
	return static_cast<QMenuBar*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenuBar_AddMenu3(void* ptr, void* icon, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

void* QMenuBar_AddMenu2(void* ptr, char* title){
	return static_cast<QMenuBar*>(ptr)->addMenu(QString(title));
}

void* QMenuBar_AddSeparator(void* ptr){
	return static_cast<QMenuBar*>(ptr)->addSeparator();
}

void QMenuBar_Clear(void* ptr){
	static_cast<QMenuBar*>(ptr)->clear();
}

void* QMenuBar_CornerWidget(void* ptr, int corner){
	return static_cast<QMenuBar*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

int QMenuBar_HeightForWidth(void* ptr, int v){
	return static_cast<QMenuBar*>(ptr)->heightForWidth(v);
}

void QMenuBar_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

void QMenuBar_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::hovered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Hovered));;
}

void* QMenuBar_InsertMenu(void* ptr, void* before, void* menu){
	return static_cast<QMenuBar*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

void* QMenuBar_InsertSeparator(void* ptr, void* before){
	return static_cast<QMenuBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

void QMenuBar_SetActiveAction(void* ptr, void* act){
	static_cast<QMenuBar*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenuBar_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QMenuBar*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QMenuBar_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QMenuBar*>(ptr), static_cast<void (QMenuBar::*)(QAction *)>(&QMenuBar::triggered), static_cast<MyQMenuBar*>(ptr), static_cast<void (MyQMenuBar::*)(QAction *)>(&MyQMenuBar::Signal_Triggered));;
}

void QMenuBar_DestroyQMenuBar(void* ptr){
	static_cast<QMenuBar*>(ptr)->~QMenuBar();
}

#include "qabstractscrollarea.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QScrollBar>
#include <QAbstractScrollArea>
#include "_cgo_export.h"

class MyQAbstractScrollArea: public QAbstractScrollArea {
public:
};

int QAbstractScrollArea_HorizontalScrollBarPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBarPolicy();
}

void QAbstractScrollArea_SetHorizontalScrollBarPolicy(void* ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

void QAbstractScrollArea_SetSizeAdjustPolicy(void* ptr, int policy){
	static_cast<QAbstractScrollArea*>(ptr)->setSizeAdjustPolicy(static_cast<QAbstractScrollArea::SizeAdjustPolicy>(policy));
}

void QAbstractScrollArea_SetVerticalScrollBarPolicy(void* ptr, int v){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBarPolicy(static_cast<Qt::ScrollBarPolicy>(v));
}

int QAbstractScrollArea_SizeAdjustPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->sizeAdjustPolicy();
}

int QAbstractScrollArea_VerticalScrollBarPolicy(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBarPolicy();
}

void* QAbstractScrollArea_NewQAbstractScrollArea(void* parent){
	return new QAbstractScrollArea(static_cast<QWidget*>(parent));
}

void QAbstractScrollArea_AddScrollBarWidget(void* ptr, void* widget, int alignment){
	static_cast<QAbstractScrollArea*>(ptr)->addScrollBarWidget(static_cast<QWidget*>(widget), static_cast<Qt::AlignmentFlag>(alignment));
}

void* QAbstractScrollArea_CornerWidget(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->cornerWidget();
}

void* QAbstractScrollArea_HorizontalScrollBar(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->horizontalScrollBar();
}

void QAbstractScrollArea_SetCornerWidget(void* ptr, void* widget){
	static_cast<QAbstractScrollArea*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetHorizontalScrollBar(void* ptr, void* scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setHorizontalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetVerticalScrollBar(void* ptr, void* scrollBar){
	static_cast<QAbstractScrollArea*>(ptr)->setVerticalScrollBar(static_cast<QScrollBar*>(scrollBar));
}

void QAbstractScrollArea_SetViewport(void* ptr, void* widget){
	static_cast<QAbstractScrollArea*>(ptr)->setViewport(static_cast<QWidget*>(widget));
}

void QAbstractScrollArea_SetupViewport(void* ptr, void* viewport){
	static_cast<QAbstractScrollArea*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void* QAbstractScrollArea_VerticalScrollBar(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->verticalScrollBar();
}

void* QAbstractScrollArea_Viewport(void* ptr){
	return static_cast<QAbstractScrollArea*>(ptr)->viewport();
}

void QAbstractScrollArea_DestroyQAbstractScrollArea(void* ptr){
	static_cast<QAbstractScrollArea*>(ptr)->~QAbstractScrollArea();
}

#include "qwidget.h"
#include <QByteArray>
#include <QUrl>
#include <QWindow>
#include <QIcon>
#include <QSizePolicy>
#include <QRect>
#include <QPalette>
#include <QKeySequence>
#include <QPaintDevice>
#include <QLocale>
#include <QAction>
#include <QBitmap>
#include <QPainter>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QSize>
#include <QLayout>
#include <QFont>
#include <QRegion>
#include <QCursor>
#include <QPoint>
#include <QGraphicsEffect>
#include <QMargins>
#include <QModelIndex>
#include <QObject>
#include <QStyle>
#include <QWidget>
#include "_cgo_export.h"

class MyQWidget: public QWidget {
public:
void Signal_WindowIconTextChanged(const QString & iconText){callbackQWidgetWindowIconTextChanged(this->objectName().toUtf8().data(), iconText.toUtf8().data());};
void Signal_WindowTitleChanged(const QString & title){callbackQWidgetWindowTitleChanged(this->objectName().toUtf8().data(), title.toUtf8().data());};
};

int QWidget_AcceptDrops(void* ptr){
	return static_cast<QWidget*>(ptr)->acceptDrops();
}

char* QWidget_AccessibleDescription(void* ptr){
	return static_cast<QWidget*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QWidget_AccessibleName(void* ptr){
	return static_cast<QWidget*>(ptr)->accessibleName().toUtf8().data();
}

void QWidget_ActivateWindow(void* ptr){
	static_cast<QWidget*>(ptr)->activateWindow();
}

int QWidget_AutoFillBackground(void* ptr){
	return static_cast<QWidget*>(ptr)->autoFillBackground();
}

void* QWidget_ChildrenRegion(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->childrenRegion());
}

void QWidget_ClearMask(void* ptr){
	static_cast<QWidget*>(ptr)->clearMask();
}

int QWidget_ContextMenuPolicy(void* ptr){
	return static_cast<QWidget*>(ptr)->contextMenuPolicy();
}

int QWidget_FocusPolicy(void* ptr){
	return static_cast<QWidget*>(ptr)->focusPolicy();
}

void QWidget_GrabKeyboard(void* ptr){
	static_cast<QWidget*>(ptr)->grabKeyboard();
}

void QWidget_GrabMouse(void* ptr){
	static_cast<QWidget*>(ptr)->grabMouse();
}

void QWidget_GrabMouse2(void* ptr, void* cursor){
	static_cast<QWidget*>(ptr)->grabMouse(*static_cast<QCursor*>(cursor));
}

int QWidget_HasFocus(void* ptr){
	return static_cast<QWidget*>(ptr)->hasFocus();
}

int QWidget_InputMethodHints(void* ptr){
	return static_cast<QWidget*>(ptr)->inputMethodHints();
}

int QWidget_IsActiveWindow(void* ptr){
	return static_cast<QWidget*>(ptr)->isActiveWindow();
}

int QWidget_IsFullScreen(void* ptr){
	return static_cast<QWidget*>(ptr)->isFullScreen();
}

int QWidget_IsMaximized(void* ptr){
	return static_cast<QWidget*>(ptr)->isMaximized();
}

int QWidget_IsMinimized(void* ptr){
	return static_cast<QWidget*>(ptr)->isMinimized();
}

int QWidget_IsWindowModified(void* ptr){
	return static_cast<QWidget*>(ptr)->isWindowModified();
}

void* QWidget_QWidget_KeyboardGrabber(){
	return QWidget::keyboardGrabber();
}

int QWidget_LayoutDirection(void* ptr){
	return static_cast<QWidget*>(ptr)->layoutDirection();
}

void* QWidget_QWidget_MouseGrabber(){
	return QWidget::mouseGrabber();
}

void QWidget_Move(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->move(*static_cast<QPoint*>(v));
}

void* QWidget_PaintEngine(void* ptr){
	return static_cast<QWidget*>(ptr)->paintEngine();
}

void QWidget_ReleaseKeyboard(void* ptr){
	static_cast<QWidget*>(ptr)->releaseKeyboard();
}

void QWidget_ReleaseMouse(void* ptr){
	static_cast<QWidget*>(ptr)->releaseMouse();
}

void QWidget_Resize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->resize(*static_cast<QSize*>(v));
}

void QWidget_SetAcceptDrops(void* ptr, int on){
	static_cast<QWidget*>(ptr)->setAcceptDrops(on != 0);
}

void QWidget_SetAccessibleDescription(void* ptr, char* description){
	static_cast<QWidget*>(ptr)->setAccessibleDescription(QString(description));
}

void QWidget_SetAccessibleName(void* ptr, char* name){
	static_cast<QWidget*>(ptr)->setAccessibleName(QString(name));
}

void QWidget_SetAutoFillBackground(void* ptr, int enabled){
	static_cast<QWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QWidget_SetContextMenuPolicy(void* ptr, int policy){
	static_cast<QWidget*>(ptr)->setContextMenuPolicy(static_cast<Qt::ContextMenuPolicy>(policy));
}

void QWidget_SetCursor(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setCursor(*static_cast<QCursor*>(v));
}

void QWidget_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QWidget_SetFixedSize2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setFixedSize(w, h);
}

void QWidget_SetFocusPolicy(void* ptr, int policy){
	static_cast<QWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QWidget_SetFont(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setFont(*static_cast<QFont*>(v));
}

void QWidget_SetGeometry(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setGeometry(*static_cast<QRect*>(v));
}

void QWidget_SetInputMethodHints(void* ptr, int hints){
	static_cast<QWidget*>(ptr)->setInputMethodHints(static_cast<Qt::InputMethodHint>(hints));
}

void QWidget_SetLayout(void* ptr, void* layout){
	static_cast<QWidget*>(ptr)->setLayout(static_cast<QLayout*>(layout));
}

void QWidget_SetLayoutDirection(void* ptr, int direction){
	static_cast<QWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QWidget_SetLocale(void* ptr, void* locale){
	static_cast<QWidget*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QWidget_SetMask(void* ptr, void* bitmap){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QBitmap*>(bitmap));
}

void QWidget_SetMask2(void* ptr, void* region){
	static_cast<QWidget*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWidget_SetMaximumHeight(void* ptr, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumHeight(maxh);
}

void QWidget_SetMaximumWidth(void* ptr, int maxw){
	static_cast<QWidget*>(ptr)->setMaximumWidth(maxw);
}

void QWidget_SetMinimumHeight(void* ptr, int minh){
	static_cast<QWidget*>(ptr)->setMinimumHeight(minh);
}

void QWidget_SetMinimumWidth(void* ptr, int minw){
	static_cast<QWidget*>(ptr)->setMinimumWidth(minw);
}

void QWidget_SetPalette(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setPalette(*static_cast<QPalette*>(v));
}

void QWidget_SetSizePolicy(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setSizePolicy(*static_cast<QSizePolicy*>(v));
}

void QWidget_SetStatusTip(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setStatusTip(QString(v));
}

void QWidget_SetStyleSheet(void* ptr, char* styleSheet){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setStyleSheet", Q_ARG(QString, QString(styleSheet)));
}

void QWidget_SetToolTip(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setToolTip(QString(v));
}

void QWidget_SetToolTipDuration(void* ptr, int msec){
	static_cast<QWidget*>(ptr)->setToolTipDuration(msec);
}

void QWidget_SetUpdatesEnabled(void* ptr, int enable){
	static_cast<QWidget*>(ptr)->setUpdatesEnabled(enable != 0);
}

void QWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWidget_SetWhatsThis(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setWhatsThis(QString(v));
}

void QWidget_SetWindowFilePath(void* ptr, char* filePath){
	static_cast<QWidget*>(ptr)->setWindowFilePath(QString(filePath));
}

void QWidget_SetWindowFlags(void* ptr, int ty){
	static_cast<QWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(ty));
}

void QWidget_SetWindowIcon(void* ptr, void* icon){
	static_cast<QWidget*>(ptr)->setWindowIcon(*static_cast<QIcon*>(icon));
}

void QWidget_SetWindowIconText(void* ptr, char* v){
	static_cast<QWidget*>(ptr)->setWindowIconText(QString(v));
}

void QWidget_SetWindowModality(void* ptr, int windowModality){
	static_cast<QWidget*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QWidget_SetWindowModified(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowModified", Q_ARG(bool, v != 0));
}

void QWidget_SetWindowOpacity(void* ptr, double level){
	static_cast<QWidget*>(ptr)->setWindowOpacity(static_cast<qreal>(level));
}

void QWidget_SetWindowState(void* ptr, int windowState){
	static_cast<QWidget*>(ptr)->setWindowState(static_cast<Qt::WindowState>(windowState));
}

void QWidget_SetWindowTitle(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setWindowTitle", Q_ARG(QString, QString(v)));
}

char* QWidget_StatusTip(void* ptr){
	return static_cast<QWidget*>(ptr)->statusTip().toUtf8().data();
}

char* QWidget_StyleSheet(void* ptr){
	return static_cast<QWidget*>(ptr)->styleSheet().toUtf8().data();
}

char* QWidget_ToolTip(void* ptr){
	return static_cast<QWidget*>(ptr)->toolTip().toUtf8().data();
}

int QWidget_ToolTipDuration(void* ptr){
	return static_cast<QWidget*>(ptr)->toolTipDuration();
}

void QWidget_UnsetCursor(void* ptr){
	static_cast<QWidget*>(ptr)->unsetCursor();
}

void QWidget_UnsetLayoutDirection(void* ptr){
	static_cast<QWidget*>(ptr)->unsetLayoutDirection();
}

void QWidget_UnsetLocale(void* ptr){
	static_cast<QWidget*>(ptr)->unsetLocale();
}

char* QWidget_WhatsThis(void* ptr){
	return static_cast<QWidget*>(ptr)->whatsThis().toUtf8().data();
}

char* QWidget_WindowFilePath(void* ptr){
	return static_cast<QWidget*>(ptr)->windowFilePath().toUtf8().data();
}

char* QWidget_WindowIconText(void* ptr){
	return static_cast<QWidget*>(ptr)->windowIconText().toUtf8().data();
}

int QWidget_WindowModality(void* ptr){
	return static_cast<QWidget*>(ptr)->windowModality();
}

double QWidget_WindowOpacity(void* ptr){
	return static_cast<double>(static_cast<QWidget*>(ptr)->windowOpacity());
}

char* QWidget_WindowTitle(void* ptr){
	return static_cast<QWidget*>(ptr)->windowTitle().toUtf8().data();
}

int QWidget_X(void* ptr){
	return static_cast<QWidget*>(ptr)->x();
}

int QWidget_Y(void* ptr){
	return static_cast<QWidget*>(ptr)->y();
}

void* QWidget_NewQWidget(void* parent, int f){
	return new QWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_AddAction(void* ptr, void* action){
	static_cast<QWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QWidget_AdjustSize(void* ptr){
	static_cast<QWidget*>(ptr)->adjustSize();
}

int QWidget_BackgroundRole(void* ptr){
	return static_cast<QWidget*>(ptr)->backgroundRole();
}

void* QWidget_BackingStore(void* ptr){
	return static_cast<QWidget*>(ptr)->backingStore();
}

void* QWidget_ChildAt2(void* ptr, void* p){
	return static_cast<QWidget*>(ptr)->childAt(*static_cast<QPoint*>(p));
}

void* QWidget_ChildAt(void* ptr, int x, int y){
	return static_cast<QWidget*>(ptr)->childAt(x, y);
}

void QWidget_ClearFocus(void* ptr){
	static_cast<QWidget*>(ptr)->clearFocus();
}

int QWidget_Close(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "close");
}

void QWidget_EnsurePolished(void* ptr){
	static_cast<QWidget*>(ptr)->ensurePolished();
}

void* QWidget_FocusProxy(void* ptr){
	return static_cast<QWidget*>(ptr)->focusProxy();
}

void* QWidget_FocusWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->focusWidget();
}

int QWidget_ForegroundRole(void* ptr){
	return static_cast<QWidget*>(ptr)->foregroundRole();
}

void QWidget_GetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

void QWidget_GrabGesture(void* ptr, int gesture, int flags){
	static_cast<QWidget*>(ptr)->grabGesture(static_cast<Qt::GestureType>(gesture), static_cast<Qt::GestureFlag>(flags));
}

int QWidget_GrabShortcut(void* ptr, void* key, int context){
	return static_cast<QWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(key), static_cast<Qt::ShortcutContext>(context));
}

void* QWidget_GraphicsEffect(void* ptr){
	return static_cast<QWidget*>(ptr)->graphicsEffect();
}

void* QWidget_GraphicsProxyWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->graphicsProxyWidget();
}

int QWidget_HasHeightForWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->hasHeightForWidth();
}

int QWidget_HasMouseTracking(void* ptr){
	return static_cast<QWidget*>(ptr)->hasMouseTracking();
}

int QWidget_Height(void* ptr){
	return static_cast<QWidget*>(ptr)->height();
}

int QWidget_HeightForWidth(void* ptr, int w){
	return static_cast<QWidget*>(ptr)->heightForWidth(w);
}

void QWidget_Hide(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "hide");
}

void* QWidget_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QWidget*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QWidget_InsertAction(void* ptr, void* before, void* action){
	static_cast<QWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QWidget_IsAncestorOf(void* ptr, void* child){
	return static_cast<QWidget*>(ptr)->isAncestorOf(static_cast<QWidget*>(child));
}

int QWidget_IsEnabled(void* ptr){
	return static_cast<QWidget*>(ptr)->isEnabled();
}

int QWidget_IsEnabledTo(void* ptr, void* ancestor){
	return static_cast<QWidget*>(ptr)->isEnabledTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsHidden(void* ptr){
	return static_cast<QWidget*>(ptr)->isHidden();
}

int QWidget_IsModal(void* ptr){
	return static_cast<QWidget*>(ptr)->isModal();
}

int QWidget_IsVisible(void* ptr){
	return static_cast<QWidget*>(ptr)->isVisible();
}

int QWidget_IsVisibleTo(void* ptr, void* ancestor){
	return static_cast<QWidget*>(ptr)->isVisibleTo(static_cast<QWidget*>(ancestor));
}

int QWidget_IsWindow(void* ptr){
	return static_cast<QWidget*>(ptr)->isWindow();
}

void* QWidget_Layout(void* ptr){
	return static_cast<QWidget*>(ptr)->layout();
}

void QWidget_Lower(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "lower");
}

void* QWidget_Mask(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->mask());
}

int QWidget_MaximumHeight(void* ptr){
	return static_cast<QWidget*>(ptr)->maximumHeight();
}

int QWidget_MaximumWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->maximumWidth();
}

int QWidget_MinimumHeight(void* ptr){
	return static_cast<QWidget*>(ptr)->minimumHeight();
}

int QWidget_MinimumWidth(void* ptr){
	return static_cast<QWidget*>(ptr)->minimumWidth();
}

void QWidget_Move2(void* ptr, int x, int y){
	static_cast<QWidget*>(ptr)->move(x, y);
}

void* QWidget_NativeParentWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->nativeParentWidget();
}

void* QWidget_NextInFocusChain(void* ptr){
	return static_cast<QWidget*>(ptr)->nextInFocusChain();
}

void QWidget_OverrideWindowFlags(void* ptr, int flags){
	static_cast<QWidget*>(ptr)->overrideWindowFlags(static_cast<Qt::WindowType>(flags));
}

void* QWidget_ParentWidget(void* ptr){
	return static_cast<QWidget*>(ptr)->parentWidget();
}

void* QWidget_PreviousInFocusChain(void* ptr){
	return static_cast<QWidget*>(ptr)->previousInFocusChain();
}

void QWidget_Raise(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "raise");
}

void QWidget_ReleaseShortcut(void* ptr, int id){
	static_cast<QWidget*>(ptr)->releaseShortcut(id);
}

void QWidget_RemoveAction(void* ptr, void* action){
	static_cast<QWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QWidget_Render(void* ptr, void* target, void* targetOffset, void* sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPaintDevice*>(target), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Render2(void* ptr, void* painter, void* targetOffset, void* sourceRegion, int renderFlags){
	static_cast<QWidget*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QPoint*>(targetOffset), *static_cast<QRegion*>(sourceRegion), static_cast<QWidget::RenderFlag>(renderFlags));
}

void QWidget_Repaint(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "repaint");
}

void QWidget_Repaint3(void* ptr, void* rect){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRect*>(rect));
}

void QWidget_Repaint4(void* ptr, void* rgn){
	static_cast<QWidget*>(ptr)->repaint(*static_cast<QRegion*>(rgn));
}

void QWidget_Repaint2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->repaint(x, y, w, h);
}

void QWidget_Resize2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->resize(w, h);
}

int QWidget_RestoreGeometry(void* ptr, void* geometry){
	return static_cast<QWidget*>(ptr)->restoreGeometry(*static_cast<QByteArray*>(geometry));
}

void* QWidget_SaveGeometry(void* ptr){
	return new QByteArray(static_cast<QWidget*>(ptr)->saveGeometry());
}

void QWidget_Scroll(void* ptr, int dx, int dy){
	static_cast<QWidget*>(ptr)->scroll(dx, dy);
}

void QWidget_Scroll2(void* ptr, int dx, int dy, void* r){
	static_cast<QWidget*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(r));
}

void QWidget_SetAttribute(void* ptr, int attribute, int on){
	static_cast<QWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QWidget_SetBackgroundRole(void* ptr, int role){
	static_cast<QWidget*>(ptr)->setBackgroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetBaseSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setBaseSize(*static_cast<QSize*>(v));
}

void QWidget_SetBaseSize2(void* ptr, int basew, int baseh){
	static_cast<QWidget*>(ptr)->setBaseSize(basew, baseh);
}

void QWidget_SetContentsMargins2(void* ptr, void* margins){
	static_cast<QWidget*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QWidget_SetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QWidget*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QWidget_SetDisabled(void* ptr, int disable){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setDisabled", Q_ARG(bool, disable != 0));
}

void QWidget_SetFixedHeight(void* ptr, int h){
	static_cast<QWidget*>(ptr)->setFixedHeight(h);
}

void QWidget_SetFixedSize(void* ptr, void* s){
	static_cast<QWidget*>(ptr)->setFixedSize(*static_cast<QSize*>(s));
}

void QWidget_SetFixedWidth(void* ptr, int w){
	static_cast<QWidget*>(ptr)->setFixedWidth(w);
}

void QWidget_SetFocus2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setFocus");
}

void QWidget_SetFocus(void* ptr, int reason){
	static_cast<QWidget*>(ptr)->setFocus(static_cast<Qt::FocusReason>(reason));
}

void QWidget_SetFocusProxy(void* ptr, void* w){
	static_cast<QWidget*>(ptr)->setFocusProxy(static_cast<QWidget*>(w));
}

void QWidget_SetForegroundRole(void* ptr, int role){
	static_cast<QWidget*>(ptr)->setForegroundRole(static_cast<QPalette::ColorRole>(role));
}

void QWidget_SetGeometry2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->setGeometry(x, y, w, h);
}

void QWidget_SetGraphicsEffect(void* ptr, void* effect){
	static_cast<QWidget*>(ptr)->setGraphicsEffect(static_cast<QGraphicsEffect*>(effect));
}

void QWidget_SetHidden(void* ptr, int hidden){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "setHidden", Q_ARG(bool, hidden != 0));
}

void QWidget_SetMaximumSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setMaximumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMaximumSize2(void* ptr, int maxw, int maxh){
	static_cast<QWidget*>(ptr)->setMaximumSize(maxw, maxh);
}

void QWidget_SetMinimumSize(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setMinimumSize(*static_cast<QSize*>(v));
}

void QWidget_SetMinimumSize2(void* ptr, int minw, int minh){
	static_cast<QWidget*>(ptr)->setMinimumSize(minw, minh);
}

void QWidget_SetMouseTracking(void* ptr, int enable){
	static_cast<QWidget*>(ptr)->setMouseTracking(enable != 0);
}

void QWidget_SetParent(void* ptr, void* parent){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent));
}

void QWidget_SetParent2(void* ptr, void* parent, int f){
	static_cast<QWidget*>(ptr)->setParent(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QWidget_SetShortcutAutoRepeat(void* ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutAutoRepeat(id, enable != 0);
}

void QWidget_SetShortcutEnabled(void* ptr, int id, int enable){
	static_cast<QWidget*>(ptr)->setShortcutEnabled(id, enable != 0);
}

void QWidget_SetSizeIncrement(void* ptr, void* v){
	static_cast<QWidget*>(ptr)->setSizeIncrement(*static_cast<QSize*>(v));
}

void QWidget_SetSizeIncrement2(void* ptr, int w, int h){
	static_cast<QWidget*>(ptr)->setSizeIncrement(w, h);
}

void QWidget_SetSizePolicy2(void* ptr, int horizontal, int vertical){
	static_cast<QWidget*>(ptr)->setSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical));
}

void QWidget_SetStyle(void* ptr, void* style){
	static_cast<QWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QWidget_QWidget_SetTabOrder(void* first, void* second){
	QWidget::setTabOrder(static_cast<QWidget*>(first), static_cast<QWidget*>(second));
}

void QWidget_SetWindowRole(void* ptr, char* role){
	static_cast<QWidget*>(ptr)->setWindowRole(QString(role));
}

void QWidget_Show(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "show");
}

void QWidget_ShowFullScreen(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showFullScreen");
}

void QWidget_ShowMaximized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMaximized");
}

void QWidget_ShowMinimized(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showMinimized");
}

void QWidget_ShowNormal(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "showNormal");
}

void QWidget_StackUnder(void* ptr, void* w){
	static_cast<QWidget*>(ptr)->stackUnder(static_cast<QWidget*>(w));
}

void* QWidget_Style(void* ptr){
	return static_cast<QWidget*>(ptr)->style();
}

int QWidget_TestAttribute(void* ptr, int attribute){
	return static_cast<QWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QWidget_UnderMouse(void* ptr){
	return static_cast<QWidget*>(ptr)->underMouse();
}

void QWidget_UngrabGesture(void* ptr, int gesture){
	static_cast<QWidget*>(ptr)->ungrabGesture(static_cast<Qt::GestureType>(gesture));
}

void QWidget_Update(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWidget*>(ptr), "update");
}

void QWidget_Update3(void* ptr, void* rect){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QWidget_Update4(void* ptr, void* rgn){
	static_cast<QWidget*>(ptr)->update(*static_cast<QRegion*>(rgn));
}

void QWidget_Update2(void* ptr, int x, int y, int w, int h){
	static_cast<QWidget*>(ptr)->update(x, y, w, h);
}

void QWidget_UpdateGeometry(void* ptr){
	static_cast<QWidget*>(ptr)->updateGeometry();
}

int QWidget_UpdatesEnabled(void* ptr){
	return static_cast<QWidget*>(ptr)->updatesEnabled();
}

void* QWidget_VisibleRegion(void* ptr){
	return new QRegion(static_cast<QWidget*>(ptr)->visibleRegion());
}

int QWidget_Width(void* ptr){
	return static_cast<QWidget*>(ptr)->width();
}

void* QWidget_Window(void* ptr){
	return static_cast<QWidget*>(ptr)->window();
}

int QWidget_WindowFlags(void* ptr){
	return static_cast<QWidget*>(ptr)->windowFlags();
}

void* QWidget_WindowHandle(void* ptr){
	return static_cast<QWidget*>(ptr)->windowHandle();
}

void QWidget_ConnectWindowIconTextChanged(void* ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

void QWidget_DisconnectWindowIconTextChanged(void* ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowIconTextChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowIconTextChanged));;
}

char* QWidget_WindowRole(void* ptr){
	return static_cast<QWidget*>(ptr)->windowRole().toUtf8().data();
}

int QWidget_WindowState(void* ptr){
	return static_cast<QWidget*>(ptr)->windowState();
}

void QWidget_ConnectWindowTitleChanged(void* ptr){
	QObject::connect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

void QWidget_DisconnectWindowTitleChanged(void* ptr){
	QObject::disconnect(static_cast<QWidget*>(ptr), static_cast<void (QWidget::*)(const QString &)>(&QWidget::windowTitleChanged), static_cast<MyQWidget*>(ptr), static_cast<void (MyQWidget::*)(const QString &)>(&MyQWidget::Signal_WindowTitleChanged));;
}

int QWidget_WindowType(void* ptr){
	return static_cast<QWidget*>(ptr)->windowType();
}

void QWidget_DestroyQWidget(void* ptr){
	static_cast<QWidget*>(ptr)->~QWidget();
}

void* QWidget_QWidget_CreateWindowContainer(void* window, void* parent, int flags){
	return QWidget::createWindowContainer(static_cast<QWindow*>(window), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

#include "qabstractgraphicsshapeitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBrush>
#include <QGraphicsItem>
#include <QPen>
#include <QAbstractGraphicsShapeItem>
#include "_cgo_export.h"

class MyQAbstractGraphicsShapeItem: public QAbstractGraphicsShapeItem {
public:
};

void* QAbstractGraphicsShapeItem_Brush(void* ptr){
	return new QBrush(static_cast<QAbstractGraphicsShapeItem*>(ptr)->brush());
}

int QAbstractGraphicsShapeItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QAbstractGraphicsShapeItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QAbstractGraphicsShapeItem_SetBrush(void* ptr, void* brush){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QAbstractGraphicsShapeItem_SetPen(void* ptr, void* pen){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QAbstractGraphicsShapeItem_DestroyQAbstractGraphicsShapeItem(void* ptr){
	static_cast<QAbstractGraphicsShapeItem*>(ptr)->~QAbstractGraphicsShapeItem();
}

#include "qgraphicssceneresizeevent.h"
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsSceneResizeEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneResizeEvent: public QGraphicsSceneResizeEvent {
public:
};

void* QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent(){
	return new QGraphicsSceneResizeEvent();
}

void QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(void* ptr){
	static_cast<QGraphicsSceneResizeEvent*>(ptr)->~QGraphicsSceneResizeEvent();
}

#include "qgraphicsgridlayout.h"
#include <QGraphicsLayoutItem>
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsGridLayout>
#include "_cgo_export.h"

class MyQGraphicsGridLayout: public QGraphicsGridLayout {
public:
};

void* QGraphicsGridLayout_NewQGraphicsGridLayout(void* parent){
	return new QGraphicsGridLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsGridLayout_AddItem2(void* ptr, void* item, int row, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGraphicsGridLayout_Alignment(void* ptr, void* item){
	return static_cast<QGraphicsGridLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_ColumnAlignment(void* ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnAlignment(column);
}

int QGraphicsGridLayout_ColumnCount(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnCount();
}

double QGraphicsGridLayout_ColumnMaximumWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnMaximumWidth(column));
}

double QGraphicsGridLayout_ColumnMinimumWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnMinimumWidth(column));
}

double QGraphicsGridLayout_ColumnPreferredWidth(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnPreferredWidth(column));
}

double QGraphicsGridLayout_ColumnSpacing(void* ptr, int column){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->columnSpacing(column));
}

int QGraphicsGridLayout_ColumnStretchFactor(void* ptr, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->columnStretchFactor(column);
}

int QGraphicsGridLayout_Count(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->count();
}

double QGraphicsGridLayout_HorizontalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->horizontalSpacing());
}

void QGraphicsGridLayout_Invalidate(void* ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->invalidate();
}

void* QGraphicsGridLayout_ItemAt2(void* ptr, int index){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(index);
}

void* QGraphicsGridLayout_ItemAt(void* ptr, int row, int column){
	return static_cast<QGraphicsGridLayout*>(ptr)->itemAt(row, column);
}

void QGraphicsGridLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsGridLayout*>(ptr)->removeAt(index);
}

void QGraphicsGridLayout_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsGridLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsGridLayout_RowAlignment(void* ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowAlignment(row);
}

int QGraphicsGridLayout_RowCount(void* ptr){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowCount();
}

double QGraphicsGridLayout_RowMaximumHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowMaximumHeight(row));
}

double QGraphicsGridLayout_RowMinimumHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowMinimumHeight(row));
}

double QGraphicsGridLayout_RowPreferredHeight(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowPreferredHeight(row));
}

double QGraphicsGridLayout_RowSpacing(void* ptr, int row){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->rowSpacing(row));
}

int QGraphicsGridLayout_RowStretchFactor(void* ptr, int row){
	return static_cast<QGraphicsGridLayout*>(ptr)->rowStretchFactor(row);
}

void QGraphicsGridLayout_SetAlignment(void* ptr, void* item, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnAlignment(void* ptr, int column, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnAlignment(column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetColumnFixedWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnFixedWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnMaximumWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnMaximumWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnMinimumWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnMinimumWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnPreferredWidth(void* ptr, int column, double width){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnPreferredWidth(column, static_cast<qreal>(width));
}

void QGraphicsGridLayout_SetColumnSpacing(void* ptr, int column, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnSpacing(column, static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetColumnStretchFactor(void* ptr, int column, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setColumnStretchFactor(column, stretch);
}

void QGraphicsGridLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsGridLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsGridLayout_SetHorizontalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setHorizontalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetRowAlignment(void* ptr, int row, int alignment){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowAlignment(row, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsGridLayout_SetRowFixedHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowFixedHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowMaximumHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowMaximumHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowMinimumHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowMinimumHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowPreferredHeight(void* ptr, int row, double height){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowPreferredHeight(row, static_cast<qreal>(height));
}

void QGraphicsGridLayout_SetRowSpacing(void* ptr, int row, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowSpacing(row, static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetRowStretchFactor(void* ptr, int row, int stretch){
	static_cast<QGraphicsGridLayout*>(ptr)->setRowStretchFactor(row, stretch);
}

void QGraphicsGridLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsGridLayout_SetVerticalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsGridLayout*>(ptr)->setVerticalSpacing(static_cast<qreal>(spacing));
}

double QGraphicsGridLayout_VerticalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsGridLayout*>(ptr)->verticalSpacing());
}

void QGraphicsGridLayout_DestroyQGraphicsGridLayout(void* ptr){
	static_cast<QGraphicsGridLayout*>(ptr)->~QGraphicsGridLayout();
}

#include "qgraphicsitemgroup.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QPainter>
#include <QStyle>
#include <QGraphicsItem>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsItemGroup>
#include "_cgo_export.h"

class MyQGraphicsItemGroup: public QGraphicsItemGroup {
public:
};

void* QGraphicsItemGroup_NewQGraphicsItemGroup(void* parent){
	return new QGraphicsItemGroup(static_cast<QGraphicsItem*>(parent));
}

void QGraphicsItemGroup_AddToGroup(void* ptr, void* item){
	static_cast<QGraphicsItemGroup*>(ptr)->addToGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsItemGroup*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsItemGroup_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsItemGroup*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsItemGroup_RemoveFromGroup(void* ptr, void* item){
	static_cast<QGraphicsItemGroup*>(ptr)->removeFromGroup(static_cast<QGraphicsItem*>(item));
}

int QGraphicsItemGroup_Type(void* ptr){
	return static_cast<QGraphicsItemGroup*>(ptr)->type();
}

void QGraphicsItemGroup_DestroyQGraphicsItemGroup(void* ptr){
	static_cast<QGraphicsItemGroup*>(ptr)->~QGraphicsItemGroup();
}

#include "qstyleoptionbutton.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionButton>
#include "_cgo_export.h"

class MyQStyleOptionButton: public QStyleOptionButton {
public:
};

void* QStyleOptionButton_NewQStyleOptionButton(){
	return new QStyleOptionButton();
}

void* QStyleOptionButton_NewQStyleOptionButton2(void* other){
	return new QStyleOptionButton(*static_cast<QStyleOptionButton*>(other));
}

#include "qcompleter.h"
#include <QUrl>
#include <QObject>
#include <QWidget>
#include <QMetaObject>
#include <QAbstractItemModel>
#include <QString>
#include <QModelIndex>
#include <QAbstractItemView>
#include <QRect>
#include <QVariant>
#include <QCompleter>
#include "_cgo_export.h"

class MyQCompleter: public QCompleter {
public:
void Signal_Activated(const QString & text){callbackQCompleterActivated(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_Highlighted(const QString & text){callbackQCompleterHighlighted(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

int QCompleter_CaseSensitivity(void* ptr){
	return static_cast<QCompleter*>(ptr)->caseSensitivity();
}

int QCompleter_CompletionColumn(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionColumn();
}

int QCompleter_CompletionMode(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionMode();
}

char* QCompleter_CompletionPrefix(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionPrefix().toUtf8().data();
}

int QCompleter_CompletionRole(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionRole();
}

int QCompleter_FilterMode(void* ptr){
	return static_cast<QCompleter*>(ptr)->filterMode();
}

int QCompleter_MaxVisibleItems(void* ptr){
	return static_cast<QCompleter*>(ptr)->maxVisibleItems();
}

int QCompleter_ModelSorting(void* ptr){
	return static_cast<QCompleter*>(ptr)->modelSorting();
}

void QCompleter_SetCaseSensitivity(void* ptr, int caseSensitivity){
	static_cast<QCompleter*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(caseSensitivity));
}

void QCompleter_SetCompletionColumn(void* ptr, int column){
	static_cast<QCompleter*>(ptr)->setCompletionColumn(column);
}

void QCompleter_SetCompletionMode(void* ptr, int mode){
	static_cast<QCompleter*>(ptr)->setCompletionMode(static_cast<QCompleter::CompletionMode>(mode));
}

void QCompleter_SetCompletionPrefix(void* ptr, char* prefix){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setCompletionPrefix", Q_ARG(QString, QString(prefix)));
}

void QCompleter_SetCompletionRole(void* ptr, int role){
	static_cast<QCompleter*>(ptr)->setCompletionRole(role);
}

void QCompleter_SetFilterMode(void* ptr, int filterMode){
	static_cast<QCompleter*>(ptr)->setFilterMode(static_cast<Qt::MatchFlag>(filterMode));
}

void QCompleter_SetMaxVisibleItems(void* ptr, int maxItems){
	static_cast<QCompleter*>(ptr)->setMaxVisibleItems(maxItems);
}

void QCompleter_SetModelSorting(void* ptr, int sorting){
	static_cast<QCompleter*>(ptr)->setModelSorting(static_cast<QCompleter::ModelSorting>(sorting));
}

void QCompleter_SetWrapAround(void* ptr, int wrap){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setWrapAround", Q_ARG(bool, wrap != 0));
}

int QCompleter_WrapAround(void* ptr){
	return static_cast<QCompleter*>(ptr)->wrapAround();
}

void* QCompleter_NewQCompleter2(void* model, void* parent){
	return new QCompleter(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

void* QCompleter_NewQCompleter(void* parent){
	return new QCompleter(static_cast<QObject*>(parent));
}

void* QCompleter_NewQCompleter3(char* list, void* parent){
	return new QCompleter(QString(list).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

void QCompleter_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_Complete(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "complete", Q_ARG(QRect, *static_cast<QRect*>(rect)));
}

int QCompleter_CompletionCount(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionCount();
}

void* QCompleter_CompletionModel(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionModel();
}

char* QCompleter_CurrentCompletion(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentCompletion().toUtf8().data();
}

void* QCompleter_CurrentIndex(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentIndex().internalPointer();
}

int QCompleter_CurrentRow(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentRow();
}

void QCompleter_ConnectHighlighted(void* ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

void QCompleter_DisconnectHighlighted(void* ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

void* QCompleter_Model(void* ptr){
	return static_cast<QCompleter*>(ptr)->model();
}

char* QCompleter_PathFromIndex(void* ptr, void* index){
	return static_cast<QCompleter*>(ptr)->pathFromIndex(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void* QCompleter_Popup(void* ptr){
	return static_cast<QCompleter*>(ptr)->popup();
}

int QCompleter_SetCurrentRow(void* ptr, int row){
	return static_cast<QCompleter*>(ptr)->setCurrentRow(row);
}

void QCompleter_SetModel(void* ptr, void* model){
	static_cast<QCompleter*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QCompleter_SetPopup(void* ptr, void* popup){
	static_cast<QCompleter*>(ptr)->setPopup(static_cast<QAbstractItemView*>(popup));
}

void QCompleter_SetWidget(void* ptr, void* widget){
	static_cast<QCompleter*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

char* QCompleter_SplitPath(void* ptr, char* path){
	return static_cast<QCompleter*>(ptr)->splitPath(QString(path)).join("|").toUtf8().data();
}

void* QCompleter_Widget(void* ptr){
	return static_cast<QCompleter*>(ptr)->widget();
}

void QCompleter_DestroyQCompleter(void* ptr){
	static_cast<QCompleter*>(ptr)->~QCompleter();
}

#include "qgraphicsscenemoveevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneMoveEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMoveEvent: public QGraphicsSceneMoveEvent {
public:
};

void* QGraphicsSceneMoveEvent_NewQGraphicsSceneMoveEvent(){
	return new QGraphicsSceneMoveEvent();
}

void QGraphicsSceneMoveEvent_DestroyQGraphicsSceneMoveEvent(void* ptr){
	static_cast<QGraphicsSceneMoveEvent*>(ptr)->~QGraphicsSceneMoveEvent();
}

#include "qtreewidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QTreeWidgetItem>
#include <QItemSelectionModel>
#include <QItemSelection>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QAbstractItemView>
#include <QMetaObject>
#include <QPoint>
#include <QObject>
#include <QTreeWidget>
#include "_cgo_export.h"

class MyQTreeWidget: public QTreeWidget {
public:
void Signal_CurrentItemChanged(QTreeWidgetItem * current, QTreeWidgetItem * previous){callbackQTreeWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_ItemActivated(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemActivated(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemChanged(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemChanged(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemCollapsed(QTreeWidgetItem * item){callbackQTreeWidgetItemCollapsed(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemEntered(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemEntered(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemExpanded(QTreeWidgetItem * item){callbackQTreeWidgetItemExpanded(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemPressed(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemSelectionChanged(){callbackQTreeWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QTreeWidget_ColumnCount(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->columnCount();
}

void QTreeWidget_SetColumnCount(void* ptr, int columns){
	static_cast<QTreeWidget*>(ptr)->setColumnCount(columns);
}

int QTreeWidget_TopLevelItemCount(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->topLevelItemCount();
}

void* QTreeWidget_NewQTreeWidget(void* parent){
	return new QTreeWidget(static_cast<QWidget*>(parent));
}

void QTreeWidget_AddTopLevelItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->addTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "clear");
}

void QTreeWidget_ClosePersistentEditor(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->closePersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_CollapseItem(void* ptr, void* item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "collapseItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

int QTreeWidget_CurrentColumn(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->currentColumn();
}

void* QTreeWidget_CurrentItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->currentItem();
}

void QTreeWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_EditItem(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->editItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ExpandItem(void* ptr, void* item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "expandItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

void* QTreeWidget_HeaderItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->headerItem();
}

int QTreeWidget_IndexOfTopLevelItem(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->indexOfTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_InsertTopLevelItem(void* ptr, int index, void* item){
	static_cast<QTreeWidget*>(ptr)->insertTopLevelItem(index, static_cast<QTreeWidgetItem*>(item));
}

void* QTreeWidget_InvisibleRootItem(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->invisibleRootItem();
}

int QTreeWidget_IsFirstItemColumnSpanned(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->isFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item));
}

void* QTreeWidget_ItemAbove(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->itemAbove(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

void QTreeWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

void* QTreeWidget_ItemAt(void* ptr, void* p){
	return static_cast<QTreeWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

void* QTreeWidget_ItemAt2(void* ptr, int x, int y){
	return static_cast<QTreeWidget*>(ptr)->itemAt(x, y);
}

void* QTreeWidget_ItemBelow(void* ptr, void* item){
	return static_cast<QTreeWidget*>(ptr)->itemBelow(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_ConnectItemCollapsed(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_DisconnectItemCollapsed(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_ConnectItemExpanded(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_DisconnectItemExpanded(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

void QTreeWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

void* QTreeWidget_ItemWidget(void* ptr, void* item, int column){
	return static_cast<QTreeWidget*>(ptr)->itemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_OpenPersistentEditor(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->openPersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_RemoveItemWidget(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->removeItemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "scrollToItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTreeWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetCurrentItem2(void* ptr, void* item, int column){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_SetCurrentItem3(void* ptr, void* item, int column, int command){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTreeWidget_SetFirstItemColumnSpanned(void* ptr, void* item, int span){
	static_cast<QTreeWidget*>(ptr)->setFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item), span != 0);
}

void QTreeWidget_SetHeaderItem(void* ptr, void* item){
	static_cast<QTreeWidget*>(ptr)->setHeaderItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetHeaderLabel(void* ptr, char* label){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabel(QString(label));
}

void QTreeWidget_SetHeaderLabels(void* ptr, char* labels){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTreeWidget_SetItemWidget(void* ptr, void* item, int column, void* widget){
	static_cast<QTreeWidget*>(ptr)->setItemWidget(static_cast<QTreeWidgetItem*>(item), column, static_cast<QWidget*>(widget));
}

void QTreeWidget_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTreeWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QTreeWidget_SortColumn(void* ptr){
	return static_cast<QTreeWidget*>(ptr)->sortColumn();
}

void QTreeWidget_SortItems(void* ptr, int column, int order){
	static_cast<QTreeWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

void* QTreeWidget_TakeTopLevelItem(void* ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->takeTopLevelItem(index);
}

void* QTreeWidget_TopLevelItem(void* ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->topLevelItem(index);
}

void QTreeWidget_DestroyQTreeWidget(void* ptr){
	static_cast<QTreeWidget*>(ptr)->~QTreeWidget();
}

#include "qformlayout.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLayoutItem>
#include <QRect>
#include <QLayout>
#include <QFormLayout>
#include "_cgo_export.h"

class MyQFormLayout: public QFormLayout {
public:
};

int QFormLayout_FieldGrowthPolicy(void* ptr){
	return static_cast<QFormLayout*>(ptr)->fieldGrowthPolicy();
}

int QFormLayout_FormAlignment(void* ptr){
	return static_cast<QFormLayout*>(ptr)->formAlignment();
}

int QFormLayout_HorizontalSpacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->horizontalSpacing();
}

int QFormLayout_LabelAlignment(void* ptr){
	return static_cast<QFormLayout*>(ptr)->labelAlignment();
}

int QFormLayout_RowWrapPolicy(void* ptr){
	return static_cast<QFormLayout*>(ptr)->rowWrapPolicy();
}

void QFormLayout_SetFieldGrowthPolicy(void* ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setFieldGrowthPolicy(static_cast<QFormLayout::FieldGrowthPolicy>(policy));
}

void QFormLayout_SetFormAlignment(void* ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setFormAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetHorizontalSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QFormLayout_SetLabelAlignment(void* ptr, int alignment){
	static_cast<QFormLayout*>(ptr)->setLabelAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QFormLayout_SetRowWrapPolicy(void* ptr, int policy){
	static_cast<QFormLayout*>(ptr)->setRowWrapPolicy(static_cast<QFormLayout::RowWrapPolicy>(policy));
}

void QFormLayout_SetVerticalSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QFormLayout_VerticalSpacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->verticalSpacing();
}

void* QFormLayout_NewQFormLayout(void* parent){
	return new QFormLayout(static_cast<QWidget*>(parent));
}

void QFormLayout_AddItem(void* ptr, void* item){
	static_cast<QFormLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QFormLayout_AddRow6(void* ptr, void* layout){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QLayout*>(layout));
}

void QFormLayout_AddRow2(void* ptr, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow(void* ptr, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_AddRow5(void* ptr, void* widget){
	static_cast<QFormLayout*>(ptr)->addRow(static_cast<QWidget*>(widget));
}

void QFormLayout_AddRow4(void* ptr, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_AddRow3(void* ptr, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->addRow(QString(labelText), static_cast<QWidget*>(field));
}

int QFormLayout_Count(void* ptr){
	return static_cast<QFormLayout*>(ptr)->count();
}

int QFormLayout_ExpandingDirections(void* ptr){
	return static_cast<QFormLayout*>(ptr)->expandingDirections();
}

int QFormLayout_HasHeightForWidth(void* ptr){
	return static_cast<QFormLayout*>(ptr)->hasHeightForWidth();
}

int QFormLayout_HeightForWidth(void* ptr, int width){
	return static_cast<QFormLayout*>(ptr)->heightForWidth(width);
}

void QFormLayout_InsertRow6(void* ptr, int row, void* layout){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QLayout*>(layout));
}

void QFormLayout_InsertRow2(void* ptr, int row, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow(void* ptr, int row, void* label, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(label), static_cast<QWidget*>(field));
}

void QFormLayout_InsertRow5(void* ptr, int row, void* widget){
	static_cast<QFormLayout*>(ptr)->insertRow(row, static_cast<QWidget*>(widget));
}

void QFormLayout_InsertRow4(void* ptr, int row, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QLayout*>(field));
}

void QFormLayout_InsertRow3(void* ptr, int row, char* labelText, void* field){
	static_cast<QFormLayout*>(ptr)->insertRow(row, QString(labelText), static_cast<QWidget*>(field));
}

void QFormLayout_Invalidate(void* ptr){
	static_cast<QFormLayout*>(ptr)->invalidate();
}

void* QFormLayout_ItemAt2(void* ptr, int index){
	return static_cast<QFormLayout*>(ptr)->itemAt(index);
}

void* QFormLayout_ItemAt(void* ptr, int row, int role){
	return static_cast<QFormLayout*>(ptr)->itemAt(row, static_cast<QFormLayout::ItemRole>(role));
}

void* QFormLayout_LabelForField2(void* ptr, void* field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QLayout*>(field));
}

void* QFormLayout_LabelForField(void* ptr, void* field){
	return static_cast<QFormLayout*>(ptr)->labelForField(static_cast<QWidget*>(field));
}

int QFormLayout_RowCount(void* ptr){
	return static_cast<QFormLayout*>(ptr)->rowCount();
}

void QFormLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QFormLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QFormLayout_SetItem(void* ptr, int row, int role, void* item){
	static_cast<QFormLayout*>(ptr)->setItem(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayoutItem*>(item));
}

void QFormLayout_SetLayout(void* ptr, int row, int role, void* layout){
	static_cast<QFormLayout*>(ptr)->setLayout(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QLayout*>(layout));
}

void QFormLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QFormLayout*>(ptr)->setSpacing(spacing);
}

void QFormLayout_SetWidget(void* ptr, int row, int role, void* widget){
	static_cast<QFormLayout*>(ptr)->setWidget(row, static_cast<QFormLayout::ItemRole>(role), static_cast<QWidget*>(widget));
}

int QFormLayout_Spacing(void* ptr){
	return static_cast<QFormLayout*>(ptr)->spacing();
}

void* QFormLayout_TakeAt(void* ptr, int index){
	return static_cast<QFormLayout*>(ptr)->takeAt(index);
}

void QFormLayout_DestroyQFormLayout(void* ptr){
	static_cast<QFormLayout*>(ptr)->~QFormLayout();
}

#include "qkeyeventtransition.h"
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QObject>
#include <QKeyEvent>
#include <QState>
#include <QString>
#include <QVariant>
#include <QKeyEventTransition>
#include "_cgo_export.h"

class MyQKeyEventTransition: public QKeyEventTransition {
public:
};

void* QKeyEventTransition_NewQKeyEventTransition2(void* object, int ty, int key, void* sourceState){
	return new QKeyEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), key, static_cast<QState*>(sourceState));
}

void* QKeyEventTransition_NewQKeyEventTransition(void* sourceState){
	return new QKeyEventTransition(static_cast<QState*>(sourceState));
}

int QKeyEventTransition_Key(void* ptr){
	return static_cast<QKeyEventTransition*>(ptr)->key();
}

int QKeyEventTransition_ModifierMask(void* ptr){
	return static_cast<QKeyEventTransition*>(ptr)->modifierMask();
}

void QKeyEventTransition_SetKey(void* ptr, int key){
	static_cast<QKeyEventTransition*>(ptr)->setKey(key);
}

void QKeyEventTransition_SetModifierMask(void* ptr, int modifierMask){
	static_cast<QKeyEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QKeyEventTransition_DestroyQKeyEventTransition(void* ptr){
	static_cast<QKeyEventTransition*>(ptr)->~QKeyEventTransition();
}

#include "qframe.h"
#include <QModelIndex>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFrame>
#include "_cgo_export.h"

class MyQFrame: public QFrame {
public:
};

int QFrame_FrameShadow(void* ptr){
	return static_cast<QFrame*>(ptr)->frameShadow();
}

int QFrame_FrameShape(void* ptr){
	return static_cast<QFrame*>(ptr)->frameShape();
}

int QFrame_FrameWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->frameWidth();
}

int QFrame_LineWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->lineWidth();
}

int QFrame_MidLineWidth(void* ptr){
	return static_cast<QFrame*>(ptr)->midLineWidth();
}

void QFrame_SetFrameRect(void* ptr, void* v){
	static_cast<QFrame*>(ptr)->setFrameRect(*static_cast<QRect*>(v));
}

void QFrame_SetFrameShadow(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShadow(static_cast<QFrame::Shadow>(v));
}

void QFrame_SetFrameShape(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setFrameShape(static_cast<QFrame::Shape>(v));
}

void QFrame_SetLineWidth(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setLineWidth(v);
}

void QFrame_SetMidLineWidth(void* ptr, int v){
	static_cast<QFrame*>(ptr)->setMidLineWidth(v);
}

void* QFrame_NewQFrame(void* parent, int f){
	return new QFrame(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QFrame_FrameStyle(void* ptr){
	return static_cast<QFrame*>(ptr)->frameStyle();
}

void QFrame_SetFrameStyle(void* ptr, int style){
	static_cast<QFrame*>(ptr)->setFrameStyle(style);
}

void QFrame_DestroyQFrame(void* ptr){
	static_cast<QFrame*>(ptr)->~QFrame();
}

#include "qgraphicssceneevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneEvent: public QGraphicsSceneEvent {
public:
};

void* QGraphicsSceneEvent_Widget(void* ptr){
	return static_cast<QGraphicsSceneEvent*>(ptr)->widget();
}

void QGraphicsSceneEvent_DestroyQGraphicsSceneEvent(void* ptr){
	static_cast<QGraphicsSceneEvent*>(ptr)->~QGraphicsSceneEvent();
}

#include "qlistview.h"
#include <QList>
#include <QPoint>
#include <QList>
#include <QUrl>
#include <QVariant>
#include <QModelIndex>
#include <QSize>
#include <QWidget>
#include <QAbstractItemView>
#include <QString>
#include <QListView>
#include "_cgo_export.h"

class MyQListView: public QListView {
public:
};

int QListView_BatchSize(void* ptr){
	return static_cast<QListView*>(ptr)->batchSize();
}

int QListView_Flow(void* ptr){
	return static_cast<QListView*>(ptr)->flow();
}

int QListView_IsSelectionRectVisible(void* ptr){
	return static_cast<QListView*>(ptr)->isSelectionRectVisible();
}

int QListView_IsWrapping(void* ptr){
	return static_cast<QListView*>(ptr)->isWrapping();
}

int QListView_LayoutMode(void* ptr){
	return static_cast<QListView*>(ptr)->layoutMode();
}

int QListView_ModelColumn(void* ptr){
	return static_cast<QListView*>(ptr)->modelColumn();
}

int QListView_Movement(void* ptr){
	return static_cast<QListView*>(ptr)->movement();
}

int QListView_ResizeMode(void* ptr){
	return static_cast<QListView*>(ptr)->resizeMode();
}

void QListView_SetBatchSize(void* ptr, int batchSize){
	static_cast<QListView*>(ptr)->setBatchSize(batchSize);
}

void QListView_SetFlow(void* ptr, int flow){
	static_cast<QListView*>(ptr)->setFlow(static_cast<QListView::Flow>(flow));
}

void QListView_SetGridSize(void* ptr, void* size){
	static_cast<QListView*>(ptr)->setGridSize(*static_cast<QSize*>(size));
}

void QListView_SetLayoutMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setLayoutMode(static_cast<QListView::LayoutMode>(mode));
}

void QListView_SetModelColumn(void* ptr, int column){
	static_cast<QListView*>(ptr)->setModelColumn(column);
}

void QListView_SetMovement(void* ptr, int movement){
	static_cast<QListView*>(ptr)->setMovement(static_cast<QListView::Movement>(movement));
}

void QListView_SetResizeMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setResizeMode(static_cast<QListView::ResizeMode>(mode));
}

void QListView_SetSelectionRectVisible(void* ptr, int show){
	static_cast<QListView*>(ptr)->setSelectionRectVisible(show != 0);
}

void QListView_SetSpacing(void* ptr, int space){
	static_cast<QListView*>(ptr)->setSpacing(space);
}

void QListView_SetUniformItemSizes(void* ptr, int enable){
	static_cast<QListView*>(ptr)->setUniformItemSizes(enable != 0);
}

void QListView_SetViewMode(void* ptr, int mode){
	static_cast<QListView*>(ptr)->setViewMode(static_cast<QListView::ViewMode>(mode));
}

void QListView_SetWordWrap(void* ptr, int on){
	static_cast<QListView*>(ptr)->setWordWrap(on != 0);
}

void QListView_SetWrapping(void* ptr, int enable){
	static_cast<QListView*>(ptr)->setWrapping(enable != 0);
}

int QListView_Spacing(void* ptr){
	return static_cast<QListView*>(ptr)->spacing();
}

int QListView_UniformItemSizes(void* ptr){
	return static_cast<QListView*>(ptr)->uniformItemSizes();
}

int QListView_ViewMode(void* ptr){
	return static_cast<QListView*>(ptr)->viewMode();
}

int QListView_WordWrap(void* ptr){
	return static_cast<QListView*>(ptr)->wordWrap();
}

void* QListView_NewQListView(void* parent){
	return new QListView(static_cast<QWidget*>(parent));
}

void QListView_ClearPropertyFlags(void* ptr){
	static_cast<QListView*>(ptr)->clearPropertyFlags();
}

void* QListView_IndexAt(void* ptr, void* p){
	return static_cast<QListView*>(ptr)->indexAt(*static_cast<QPoint*>(p)).internalPointer();
}

int QListView_IsRowHidden(void* ptr, int row){
	return static_cast<QListView*>(ptr)->isRowHidden(row);
}

void QListView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QListView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QListView_SetRowHidden(void* ptr, int row, int hide){
	static_cast<QListView*>(ptr)->setRowHidden(row, hide != 0);
}

void QListView_DestroyQListView(void* ptr){
	static_cast<QListView*>(ptr)->~QListView();
}

#include "qfileiconprovider.h"
#include <QFileInfo>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFile>
#include <QFileIconProvider>
#include "_cgo_export.h"

class MyQFileIconProvider: public QFileIconProvider {
public:
};

void* QFileIconProvider_NewQFileIconProvider(){
	return new QFileIconProvider();
}

int QFileIconProvider_Options(void* ptr){
	return static_cast<QFileIconProvider*>(ptr)->options();
}

void QFileIconProvider_SetOptions(void* ptr, int options){
	static_cast<QFileIconProvider*>(ptr)->setOptions(static_cast<QFileIconProvider::Option>(options));
}

char* QFileIconProvider_Type(void* ptr, void* info){
	return static_cast<QFileIconProvider*>(ptr)->type(*static_cast<QFileInfo*>(info)).toUtf8().data();
}

void QFileIconProvider_DestroyQFileIconProvider(void* ptr){
	static_cast<QFileIconProvider*>(ptr)->~QFileIconProvider();
}

#include "qwidgetaction.h"
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidgetAction>
#include "_cgo_export.h"

class MyQWidgetAction: public QWidgetAction {
public:
};

void* QWidgetAction_NewQWidgetAction(void* parent){
	return new QWidgetAction(static_cast<QObject*>(parent));
}

void* QWidgetAction_DefaultWidget(void* ptr){
	return static_cast<QWidgetAction*>(ptr)->defaultWidget();
}

void QWidgetAction_ReleaseWidget(void* ptr, void* widget){
	static_cast<QWidgetAction*>(ptr)->releaseWidget(static_cast<QWidget*>(widget));
}

void* QWidgetAction_RequestWidget(void* ptr, void* parent){
	return static_cast<QWidgetAction*>(ptr)->requestWidget(static_cast<QWidget*>(parent));
}

void QWidgetAction_SetDefaultWidget(void* ptr, void* widget){
	static_cast<QWidgetAction*>(ptr)->setDefaultWidget(static_cast<QWidget*>(widget));
}

void QWidgetAction_DestroyQWidgetAction(void* ptr){
	static_cast<QWidgetAction*>(ptr)->~QWidgetAction();
}

#include "qundocommand.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QUndoCommand>
#include "_cgo_export.h"

class MyQUndoCommand: public QUndoCommand {
public:
};

void* QUndoCommand_NewQUndoCommand(void* parent){
	return new QUndoCommand(static_cast<QUndoCommand*>(parent));
}

void* QUndoCommand_NewQUndoCommand2(char* text, void* parent){
	return new QUndoCommand(QString(text), static_cast<QUndoCommand*>(parent));
}

char* QUndoCommand_ActionText(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->actionText().toUtf8().data();
}

void* QUndoCommand_Child(void* ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoCommand*>(ptr)->child(index));
}

int QUndoCommand_ChildCount(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->childCount();
}

int QUndoCommand_Id(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->id();
}

int QUndoCommand_MergeWith(void* ptr, void* command){
	return static_cast<QUndoCommand*>(ptr)->mergeWith(static_cast<QUndoCommand*>(command));
}

void QUndoCommand_Redo(void* ptr){
	static_cast<QUndoCommand*>(ptr)->redo();
}

void QUndoCommand_SetText(void* ptr, char* text){
	static_cast<QUndoCommand*>(ptr)->setText(QString(text));
}

char* QUndoCommand_Text(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->text().toUtf8().data();
}

void QUndoCommand_Undo(void* ptr){
	static_cast<QUndoCommand*>(ptr)->undo();
}

void QUndoCommand_DestroyQUndoCommand(void* ptr){
	static_cast<QUndoCommand*>(ptr)->~QUndoCommand();
}

#include "qstyleoptiontoolbox.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionToolBox>
#include "_cgo_export.h"

class MyQStyleOptionToolBox: public QStyleOptionToolBox {
public:
};

void* QStyleOptionToolBox_NewQStyleOptionToolBox(){
	return new QStyleOptionToolBox();
}

void* QStyleOptionToolBox_NewQStyleOptionToolBox2(void* other){
	return new QStyleOptionToolBox(*static_cast<QStyleOptionToolBox*>(other));
}

#include "qbuttongroup.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractButton>
#include <QButtonGroup>
#include "_cgo_export.h"

class MyQButtonGroup: public QButtonGroup {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQButtonGroupButtonClicked(this->objectName().toUtf8().data(), button);};
void Signal_ButtonPressed(QAbstractButton * button){callbackQButtonGroupButtonPressed(this->objectName().toUtf8().data(), button);};
void Signal_ButtonReleased(QAbstractButton * button){callbackQButtonGroupButtonReleased(this->objectName().toUtf8().data(), button);};
void Signal_ButtonToggled(QAbstractButton * button, bool checked){callbackQButtonGroupButtonToggled(this->objectName().toUtf8().data(), button, checked);};
};

void* QButtonGroup_NewQButtonGroup(void* parent){
	return new QButtonGroup(static_cast<QObject*>(parent));
}

void QButtonGroup_AddButton(void* ptr, void* button, int id){
	static_cast<QButtonGroup*>(ptr)->addButton(static_cast<QAbstractButton*>(button), id);
}

void* QButtonGroup_Button(void* ptr, int id){
	return static_cast<QButtonGroup*>(ptr)->button(id);
}

void* QButtonGroup_CheckedButton(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedButton();
}

int QButtonGroup_CheckedId(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->checkedId();
}

int QButtonGroup_Exclusive(void* ptr){
	return static_cast<QButtonGroup*>(ptr)->exclusive();
}

int QButtonGroup_Id(void* ptr, void* button){
	return static_cast<QButtonGroup*>(ptr)->id(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_RemoveButton(void* ptr, void* button){
	static_cast<QButtonGroup*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QButtonGroup_SetExclusive(void* ptr, int v){
	static_cast<QButtonGroup*>(ptr)->setExclusive(v != 0);
}

void QButtonGroup_SetId(void* ptr, void* button, int id){
	static_cast<QButtonGroup*>(ptr)->setId(static_cast<QAbstractButton*>(button), id);
}

void QButtonGroup_DestroyQButtonGroup(void* ptr){
	static_cast<QButtonGroup*>(ptr)->~QButtonGroup();
}

void QButtonGroup_ConnectButtonClicked(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_DisconnectButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonClicked), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonClicked));;
}

void QButtonGroup_ConnectButtonPressed(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_DisconnectButtonPressed(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonPressed), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonPressed));;
}

void QButtonGroup_ConnectButtonReleased(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_DisconnectButtonReleased(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *)>(&QButtonGroup::buttonReleased), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *)>(&MyQButtonGroup::Signal_ButtonReleased));;
}

void QButtonGroup_ConnectButtonToggled(void* ptr){
	QObject::connect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

void QButtonGroup_DisconnectButtonToggled(void* ptr){
	QObject::disconnect(static_cast<QButtonGroup*>(ptr), static_cast<void (QButtonGroup::*)(QAbstractButton *, bool)>(&QButtonGroup::buttonToggled), static_cast<MyQButtonGroup*>(ptr), static_cast<void (MyQButtonGroup::*)(QAbstractButton *, bool)>(&MyQButtonGroup::Signal_ButtonToggled));;
}

#include "qwizard.h"
#include <QPixmap>
#include <QWizardPage>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QAbstractButton>
#include <QWizard>
#include "_cgo_export.h"

class MyQWizard: public QWizard {
public:
void Signal_CurrentIdChanged(int id){callbackQWizardCurrentIdChanged(this->objectName().toUtf8().data(), id);};
void Signal_CustomButtonClicked(int which){callbackQWizardCustomButtonClicked(this->objectName().toUtf8().data(), which);};
void Signal_HelpRequested(){callbackQWizardHelpRequested(this->objectName().toUtf8().data());};
void Signal_PageAdded(int id){callbackQWizardPageAdded(this->objectName().toUtf8().data(), id);};
void Signal_PageRemoved(int id){callbackQWizardPageRemoved(this->objectName().toUtf8().data(), id);};
};

int QWizard_CurrentId(void* ptr){
	return static_cast<QWizard*>(ptr)->currentId();
}

int QWizard_HasVisitedPage(void* ptr, int id){
	return static_cast<QWizard*>(ptr)->hasVisitedPage(id);
}

int QWizard_Options(void* ptr){
	return static_cast<QWizard*>(ptr)->options();
}

void* QWizard_Page(void* ptr, int id){
	return static_cast<QWizard*>(ptr)->page(id);
}

void QWizard_SetOptions(void* ptr, int options){
	static_cast<QWizard*>(ptr)->setOptions(static_cast<QWizard::WizardOption>(options));
}

void QWizard_SetPage(void* ptr, int id, void* page){
	static_cast<QWizard*>(ptr)->setPage(id, static_cast<QWizardPage*>(page));
}

void QWizard_SetStartId(void* ptr, int id){
	static_cast<QWizard*>(ptr)->setStartId(id);
}

void QWizard_SetSubTitleFormat(void* ptr, int format){
	static_cast<QWizard*>(ptr)->setSubTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetTitleFormat(void* ptr, int format){
	static_cast<QWizard*>(ptr)->setTitleFormat(static_cast<Qt::TextFormat>(format));
}

void QWizard_SetWizardStyle(void* ptr, int style){
	static_cast<QWizard*>(ptr)->setWizardStyle(static_cast<QWizard::WizardStyle>(style));
}

int QWizard_StartId(void* ptr){
	return static_cast<QWizard*>(ptr)->startId();
}

int QWizard_SubTitleFormat(void* ptr){
	return static_cast<QWizard*>(ptr)->subTitleFormat();
}

int QWizard_TitleFormat(void* ptr){
	return static_cast<QWizard*>(ptr)->titleFormat();
}

int QWizard_WizardStyle(void* ptr){
	return static_cast<QWizard*>(ptr)->wizardStyle();
}

void* QWizard_NewQWizard(void* parent, int flags){
	return new QWizard(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QWizard_AddPage(void* ptr, void* page){
	return static_cast<QWizard*>(ptr)->addPage(static_cast<QWizardPage*>(page));
}

void QWizard_Back(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "back");
}

void* QWizard_Button(void* ptr, int which){
	return static_cast<QWizard*>(ptr)->button(static_cast<QWizard::WizardButton>(which));
}

char* QWizard_ButtonText(void* ptr, int which){
	return static_cast<QWizard*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizard_ConnectCurrentIdChanged(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

void QWizard_DisconnectCurrentIdChanged(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::currentIdChanged), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CurrentIdChanged));;
}

void* QWizard_CurrentPage(void* ptr){
	return static_cast<QWizard*>(ptr)->currentPage();
}

void QWizard_ConnectCustomButtonClicked(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

void QWizard_DisconnectCustomButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::customButtonClicked), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_CustomButtonClicked));;
}

void* QWizard_Field(void* ptr, char* name){
	return new QVariant(static_cast<QWizard*>(ptr)->field(QString(name)));
}

void QWizard_ConnectHelpRequested(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_DisconnectHelpRequested(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)()>(&QWizard::helpRequested), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)()>(&MyQWizard::Signal_HelpRequested));;
}

void QWizard_Next(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "next");
}

int QWizard_NextId(void* ptr){
	return static_cast<QWizard*>(ptr)->nextId();
}

void QWizard_ConnectPageAdded(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_DisconnectPageAdded(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageAdded), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageAdded));;
}

void QWizard_ConnectPageRemoved(void* ptr){
	QObject::connect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_DisconnectPageRemoved(void* ptr){
	QObject::disconnect(static_cast<QWizard*>(ptr), static_cast<void (QWizard::*)(int)>(&QWizard::pageRemoved), static_cast<MyQWizard*>(ptr), static_cast<void (MyQWizard::*)(int)>(&MyQWizard::Signal_PageRemoved));;
}

void QWizard_RemovePage(void* ptr, int id){
	static_cast<QWizard*>(ptr)->removePage(id);
}

void QWizard_Restart(void* ptr){
	QMetaObject::invokeMethod(static_cast<QWizard*>(ptr), "restart");
}

void QWizard_SetButton(void* ptr, int which, void* button){
	static_cast<QWizard*>(ptr)->setButton(static_cast<QWizard::WizardButton>(which), static_cast<QAbstractButton*>(button));
}

void QWizard_SetButtonText(void* ptr, int which, char* text){
	static_cast<QWizard*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizard_SetDefaultProperty(void* ptr, char* className, char* property, char* changedSignal){
	static_cast<QWizard*>(ptr)->setDefaultProperty(const_cast<const char*>(className), const_cast<const char*>(property), const_cast<const char*>(changedSignal));
}

void QWizard_SetField(void* ptr, char* name, void* value){
	static_cast<QWizard*>(ptr)->setField(QString(name), *static_cast<QVariant*>(value));
}

void QWizard_SetOption(void* ptr, int option, int on){
	static_cast<QWizard*>(ptr)->setOption(static_cast<QWizard::WizardOption>(option), on != 0);
}

void QWizard_SetPixmap(void* ptr, int which, void* pixmap){
	static_cast<QWizard*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

void QWizard_SetSideWidget(void* ptr, void* widget){
	static_cast<QWizard*>(ptr)->setSideWidget(static_cast<QWidget*>(widget));
}

void QWizard_SetVisible(void* ptr, int visible){
	static_cast<QWizard*>(ptr)->setVisible(visible != 0);
}

void* QWizard_SideWidget(void* ptr){
	return static_cast<QWizard*>(ptr)->sideWidget();
}

int QWizard_TestOption(void* ptr, int option){
	return static_cast<QWizard*>(ptr)->testOption(static_cast<QWizard::WizardOption>(option));
}

int QWizard_ValidateCurrentPage(void* ptr){
	return static_cast<QWizard*>(ptr)->validateCurrentPage();
}

void QWizard_DestroyQWizard(void* ptr){
	static_cast<QWizard*>(ptr)->~QWizard();
}

#include "qgraphicslayout.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsLayout>
#include "_cgo_export.h"

#include "qitemeditorfactory.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QItemEditorCreator>
#include <QByteArray>
#include <QItemEditorCreatorBase>
#include <QString>
#include <QItemEditorFactory>
#include "_cgo_export.h"

class MyQItemEditorFactory: public QItemEditorFactory {
public:
};

void* QItemEditorFactory_NewQItemEditorFactory(){
	return new QItemEditorFactory();
}

void* QItemEditorFactory_CreateEditor(void* ptr, int userType, void* parent){
	return static_cast<QItemEditorFactory*>(ptr)->createEditor(userType, static_cast<QWidget*>(parent));
}

void* QItemEditorFactory_QItemEditorFactory_DefaultFactory(){
	return const_cast<QItemEditorFactory*>(QItemEditorFactory::defaultFactory());
}

void QItemEditorFactory_RegisterEditor(void* ptr, int userType, void* creator){
	static_cast<QItemEditorFactory*>(ptr)->registerEditor(userType, static_cast<QItemEditorCreatorBase*>(creator));
}

void QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(void* factory){
	QItemEditorFactory::setDefaultFactory(static_cast<QItemEditorFactory*>(factory));
}

void* QItemEditorFactory_ValuePropertyName(void* ptr, int userType){
	return new QByteArray(static_cast<QItemEditorFactory*>(ptr)->valuePropertyName(userType));
}

void QItemEditorFactory_DestroyQItemEditorFactory(void* ptr){
	static_cast<QItemEditorFactory*>(ptr)->~QItemEditorFactory();
}

#include "qstandarditemeditorcreator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include "_cgo_export.h"

#include "qscrollerproperties.h"
#include <QScroller>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScrollerProperties>
#include "_cgo_export.h"

class MyQScrollerProperties: public QScrollerProperties {
public:
};

void* QScrollerProperties_NewQScrollerProperties(){
	return new QScrollerProperties();
}

void* QScrollerProperties_NewQScrollerProperties2(void* sp){
	return new QScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

void* QScrollerProperties_ScrollMetric(void* ptr, int metric){
	return new QVariant(static_cast<QScrollerProperties*>(ptr)->scrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric)));
}

void QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(void* sp){
	QScrollerProperties::setDefaultScrollerProperties(*static_cast<QScrollerProperties*>(sp));
}

void QScrollerProperties_SetScrollMetric(void* ptr, int metric, void* value){
	static_cast<QScrollerProperties*>(ptr)->setScrollMetric(static_cast<QScrollerProperties::ScrollMetric>(metric), *static_cast<QVariant*>(value));
}

void QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties(){
	QScrollerProperties::unsetDefaultScrollerProperties();
}

void QScrollerProperties_DestroyQScrollerProperties(void* ptr){
	static_cast<QScrollerProperties*>(ptr)->~QScrollerProperties();
}

#include "qgraphicsopacityeffect.h"
#include <QBrush>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsOpacityEffect>
#include "_cgo_export.h"

class MyQGraphicsOpacityEffect: public QGraphicsOpacityEffect {
public:
void Signal_OpacityMaskChanged(const QBrush & mask){callbackQGraphicsOpacityEffectOpacityMaskChanged(this->objectName().toUtf8().data(), new QBrush(mask));};
};

double QGraphicsOpacityEffect_Opacity(void* ptr){
	return static_cast<double>(static_cast<QGraphicsOpacityEffect*>(ptr)->opacity());
}

void* QGraphicsOpacityEffect_OpacityMask(void* ptr){
	return new QBrush(static_cast<QGraphicsOpacityEffect*>(ptr)->opacityMask());
}

void QGraphicsOpacityEffect_SetOpacity(void* ptr, double opacity){
	QMetaObject::invokeMethod(static_cast<QGraphicsOpacityEffect*>(ptr), "setOpacity", Q_ARG(qreal, static_cast<qreal>(opacity)));
}

void QGraphicsOpacityEffect_SetOpacityMask(void* ptr, void* mask){
	QMetaObject::invokeMethod(static_cast<QGraphicsOpacityEffect*>(ptr), "setOpacityMask", Q_ARG(QBrush, *static_cast<QBrush*>(mask)));
}

void* QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(void* parent){
	return new QGraphicsOpacityEffect(static_cast<QObject*>(parent));
}

void QGraphicsOpacityEffect_ConnectOpacityMaskChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsOpacityEffect*>(ptr), static_cast<void (QGraphicsOpacityEffect::*)(const QBrush &)>(&QGraphicsOpacityEffect::opacityMaskChanged), static_cast<MyQGraphicsOpacityEffect*>(ptr), static_cast<void (MyQGraphicsOpacityEffect::*)(const QBrush &)>(&MyQGraphicsOpacityEffect::Signal_OpacityMaskChanged));;
}

void QGraphicsOpacityEffect_DisconnectOpacityMaskChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsOpacityEffect*>(ptr), static_cast<void (QGraphicsOpacityEffect::*)(const QBrush &)>(&QGraphicsOpacityEffect::opacityMaskChanged), static_cast<MyQGraphicsOpacityEffect*>(ptr), static_cast<void (MyQGraphicsOpacityEffect::*)(const QBrush &)>(&MyQGraphicsOpacityEffect::Signal_OpacityMaskChanged));;
}

void QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(void* ptr){
	static_cast<QGraphicsOpacityEffect*>(ptr)->~QGraphicsOpacityEffect();
}

#include "qrubberband.h"
#include <QSize>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QPoint>
#include <QRubberBand>
#include "_cgo_export.h"

class MyQRubberBand: public QRubberBand {
public:
};

void QRubberBand_SetGeometry(void* ptr, void* rect){
	static_cast<QRubberBand*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QRubberBand_NewQRubberBand(int s, void* p){
	return new QRubberBand(static_cast<QRubberBand::Shape>(s), static_cast<QWidget*>(p));
}

void QRubberBand_Move2(void* ptr, void* p){
	static_cast<QRubberBand*>(ptr)->move(*static_cast<QPoint*>(p));
}

void QRubberBand_Move(void* ptr, int x, int y){
	static_cast<QRubberBand*>(ptr)->move(x, y);
}

void QRubberBand_Resize2(void* ptr, void* size){
	static_cast<QRubberBand*>(ptr)->resize(*static_cast<QSize*>(size));
}

void QRubberBand_Resize(void* ptr, int width, int height){
	static_cast<QRubberBand*>(ptr)->resize(width, height);
}

void QRubberBand_SetGeometry2(void* ptr, int x, int y, int width, int height){
	static_cast<QRubberBand*>(ptr)->setGeometry(x, y, width, height);
}

int QRubberBand_Shape(void* ptr){
	return static_cast<QRubberBand*>(ptr)->shape();
}

void QRubberBand_DestroyQRubberBand(void* ptr){
	static_cast<QRubberBand*>(ptr)->~QRubberBand();
}

#include "qgraphicslayoutitem.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QString>
#include <QGraphicsLayoutItem>
#include "_cgo_export.h"

#include "qmainwindow.h"
#include <QWidget>
#include <QToolBar>
#include <QMenuBar>
#include <QStatusBar>
#include <QByteArray>
#include <QSize>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QTabWidget>
#include <QMenu>
#include <QDockWidget>
#include <QModelIndex>
#include <QMetaObject>
#include <QUrl>
#include <QMainWindow>
#include "_cgo_export.h"

class MyQMainWindow: public QMainWindow {
public:
void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle){callbackQMainWindowToolButtonStyleChanged(this->objectName().toUtf8().data(), toolButtonStyle);};
};

int QMainWindow_DockOptions(void* ptr){
	return static_cast<QMainWindow*>(ptr)->dockOptions();
}

int QMainWindow_DocumentMode(void* ptr){
	return static_cast<QMainWindow*>(ptr)->documentMode();
}

int QMainWindow_IsAnimated(void* ptr){
	return static_cast<QMainWindow*>(ptr)->isAnimated();
}

int QMainWindow_IsDockNestingEnabled(void* ptr){
	return static_cast<QMainWindow*>(ptr)->isDockNestingEnabled();
}

void QMainWindow_SetAnimated(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setAnimated", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockNestingEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setDockNestingEnabled", Q_ARG(bool, enabled != 0));
}

void QMainWindow_SetDockOptions(void* ptr, int options){
	static_cast<QMainWindow*>(ptr)->setDockOptions(static_cast<QMainWindow::DockOption>(options));
}

void QMainWindow_SetDocumentMode(void* ptr, int enabled){
	static_cast<QMainWindow*>(ptr)->setDocumentMode(enabled != 0);
}

void QMainWindow_SetIconSize(void* ptr, void* iconSize){
	static_cast<QMainWindow*>(ptr)->setIconSize(*static_cast<QSize*>(iconSize));
}

void QMainWindow_SetTabShape(void* ptr, int tabShape){
	static_cast<QMainWindow*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(tabShape));
}

void QMainWindow_SetToolButtonStyle(void* ptr, int toolButtonStyle){
	static_cast<QMainWindow*>(ptr)->setToolButtonStyle(static_cast<Qt::ToolButtonStyle>(toolButtonStyle));
}

void QMainWindow_SetUnifiedTitleAndToolBarOnMac(void* ptr, int set){
	QMetaObject::invokeMethod(static_cast<QMainWindow*>(ptr), "setUnifiedTitleAndToolBarOnMac", Q_ARG(bool, set != 0));
}

void QMainWindow_SplitDockWidget(void* ptr, void* first, void* second, int orientation){
	static_cast<QMainWindow*>(ptr)->splitDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second), static_cast<Qt::Orientation>(orientation));
}

int QMainWindow_TabShape(void* ptr){
	return static_cast<QMainWindow*>(ptr)->tabShape();
}

void QMainWindow_TabifyDockWidget(void* ptr, void* first, void* second){
	static_cast<QMainWindow*>(ptr)->tabifyDockWidget(static_cast<QDockWidget*>(first), static_cast<QDockWidget*>(second));
}

int QMainWindow_ToolButtonStyle(void* ptr){
	return static_cast<QMainWindow*>(ptr)->toolButtonStyle();
}

int QMainWindow_UnifiedTitleAndToolBarOnMac(void* ptr){
	return static_cast<QMainWindow*>(ptr)->unifiedTitleAndToolBarOnMac();
}

void* QMainWindow_NewQMainWindow(void* parent, int flags){
	return new QMainWindow(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QMainWindow_AddDockWidget(void* ptr, int area, void* dockwidget){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_AddDockWidget2(void* ptr, int area, void* dockwidget, int orientation){
	static_cast<QMainWindow*>(ptr)->addDockWidget(static_cast<Qt::DockWidgetArea>(area), static_cast<QDockWidget*>(dockwidget), static_cast<Qt::Orientation>(orientation));
}

void* QMainWindow_AddToolBar3(void* ptr, char* title){
	return static_cast<QMainWindow*>(ptr)->addToolBar(QString(title));
}

void QMainWindow_AddToolBar2(void* ptr, void* toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBar(void* ptr, int area, void* toolbar){
	static_cast<QMainWindow*>(ptr)->addToolBar(static_cast<Qt::ToolBarArea>(area), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_AddToolBarBreak(void* ptr, int area){
	static_cast<QMainWindow*>(ptr)->addToolBarBreak(static_cast<Qt::ToolBarArea>(area));
}

void* QMainWindow_CentralWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->centralWidget();
}

int QMainWindow_Corner(void* ptr, int corner){
	return static_cast<QMainWindow*>(ptr)->corner(static_cast<Qt::Corner>(corner));
}

void* QMainWindow_CreatePopupMenu(void* ptr){
	return static_cast<QMainWindow*>(ptr)->createPopupMenu();
}

int QMainWindow_DockWidgetArea(void* ptr, void* dockwidget){
	return static_cast<QMainWindow*>(ptr)->dockWidgetArea(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_InsertToolBar(void* ptr, void* before, void* toolbar){
	static_cast<QMainWindow*>(ptr)->insertToolBar(static_cast<QToolBar*>(before), static_cast<QToolBar*>(toolbar));
}

void QMainWindow_InsertToolBarBreak(void* ptr, void* before){
	static_cast<QMainWindow*>(ptr)->insertToolBarBreak(static_cast<QToolBar*>(before));
}

void* QMainWindow_MenuBar(void* ptr){
	return static_cast<QMainWindow*>(ptr)->menuBar();
}

void* QMainWindow_MenuWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->menuWidget();
}

void QMainWindow_RemoveDockWidget(void* ptr, void* dockwidget){
	static_cast<QMainWindow*>(ptr)->removeDockWidget(static_cast<QDockWidget*>(dockwidget));
}

void QMainWindow_RemoveToolBar(void* ptr, void* toolbar){
	static_cast<QMainWindow*>(ptr)->removeToolBar(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_RemoveToolBarBreak(void* ptr, void* before){
	static_cast<QMainWindow*>(ptr)->removeToolBarBreak(static_cast<QToolBar*>(before));
}

int QMainWindow_RestoreDockWidget(void* ptr, void* dockwidget){
	return static_cast<QMainWindow*>(ptr)->restoreDockWidget(static_cast<QDockWidget*>(dockwidget));
}

int QMainWindow_RestoreState(void* ptr, void* state, int version){
	return static_cast<QMainWindow*>(ptr)->restoreState(*static_cast<QByteArray*>(state), version);
}

void* QMainWindow_SaveState(void* ptr, int version){
	return new QByteArray(static_cast<QMainWindow*>(ptr)->saveState(version));
}

void QMainWindow_SetCentralWidget(void* ptr, void* widget){
	static_cast<QMainWindow*>(ptr)->setCentralWidget(static_cast<QWidget*>(widget));
}

void QMainWindow_SetCorner(void* ptr, int corner, int area){
	static_cast<QMainWindow*>(ptr)->setCorner(static_cast<Qt::Corner>(corner), static_cast<Qt::DockWidgetArea>(area));
}

void QMainWindow_SetMenuBar(void* ptr, void* menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuBar(static_cast<QMenuBar*>(menuBar));
}

void QMainWindow_SetMenuWidget(void* ptr, void* menuBar){
	static_cast<QMainWindow*>(ptr)->setMenuWidget(static_cast<QWidget*>(menuBar));
}

void QMainWindow_SetStatusBar(void* ptr, void* statusbar){
	static_cast<QMainWindow*>(ptr)->setStatusBar(static_cast<QStatusBar*>(statusbar));
}

void QMainWindow_SetTabPosition(void* ptr, int areas, int tabPosition){
	static_cast<QMainWindow*>(ptr)->setTabPosition(static_cast<Qt::DockWidgetArea>(areas), static_cast<QTabWidget::TabPosition>(tabPosition));
}

void* QMainWindow_StatusBar(void* ptr){
	return static_cast<QMainWindow*>(ptr)->statusBar();
}

int QMainWindow_TabPosition(void* ptr, int area){
	return static_cast<QMainWindow*>(ptr)->tabPosition(static_cast<Qt::DockWidgetArea>(area));
}

void* QMainWindow_TakeCentralWidget(void* ptr){
	return static_cast<QMainWindow*>(ptr)->takeCentralWidget();
}

int QMainWindow_ToolBarArea(void* ptr, void* toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarArea(static_cast<QToolBar*>(toolbar));
}

int QMainWindow_ToolBarBreak(void* ptr, void* toolbar){
	return static_cast<QMainWindow*>(ptr)->toolBarBreak(static_cast<QToolBar*>(toolbar));
}

void QMainWindow_ConnectToolButtonStyleChanged(void* ptr){
	QObject::connect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DisconnectToolButtonStyleChanged(void* ptr){
	QObject::disconnect(static_cast<QMainWindow*>(ptr), static_cast<void (QMainWindow::*)(Qt::ToolButtonStyle)>(&QMainWindow::toolButtonStyleChanged), static_cast<MyQMainWindow*>(ptr), static_cast<void (MyQMainWindow::*)(Qt::ToolButtonStyle)>(&MyQMainWindow::Signal_ToolButtonStyleChanged));;
}

void QMainWindow_DestroyQMainWindow(void* ptr){
	static_cast<QMainWindow*>(ptr)->~QMainWindow();
}

#include "qaction.h"
#include <QModelIndex>
#include <QObject>
#include <QIcon>
#include <QMenu>
#include <QMetaObject>
#include <QKeySequence>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QActionGroup>
#include <QAction>
#include "_cgo_export.h"

class MyQAction: public QAction {
public:
void Signal_Changed(){callbackQActionChanged(this->objectName().toUtf8().data());};
void Signal_Hovered(){callbackQActionHovered(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQActionToggled(this->objectName().toUtf8().data(), checked);};
void Signal_Triggered(bool checked){callbackQActionTriggered(this->objectName().toUtf8().data(), checked);};
};

int QAction_AutoRepeat(void* ptr){
	return static_cast<QAction*>(ptr)->autoRepeat();
}

char* QAction_IconText(void* ptr){
	return static_cast<QAction*>(ptr)->iconText().toUtf8().data();
}

int QAction_IsCheckable(void* ptr){
	return static_cast<QAction*>(ptr)->isCheckable();
}

int QAction_IsChecked(void* ptr){
	return static_cast<QAction*>(ptr)->isChecked();
}

int QAction_IsEnabled(void* ptr){
	return static_cast<QAction*>(ptr)->isEnabled();
}

int QAction_IsIconVisibleInMenu(void* ptr){
	return static_cast<QAction*>(ptr)->isIconVisibleInMenu();
}

int QAction_IsVisible(void* ptr){
	return static_cast<QAction*>(ptr)->isVisible();
}

int QAction_MenuRole(void* ptr){
	return static_cast<QAction*>(ptr)->menuRole();
}

int QAction_Priority(void* ptr){
	return static_cast<QAction*>(ptr)->priority();
}

void QAction_SetAutoRepeat(void* ptr, int v){
	static_cast<QAction*>(ptr)->setAutoRepeat(v != 0);
}

void QAction_SetCheckable(void* ptr, int v){
	static_cast<QAction*>(ptr)->setCheckable(v != 0);
}

void QAction_SetChecked(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAction_SetData(void* ptr, void* userData){
	static_cast<QAction*>(ptr)->setData(*static_cast<QVariant*>(userData));
}

void QAction_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QAction_SetFont(void* ptr, void* font){
	static_cast<QAction*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QAction_SetIcon(void* ptr, void* icon){
	static_cast<QAction*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAction_SetIconText(void* ptr, char* text){
	static_cast<QAction*>(ptr)->setIconText(QString(text));
}

void QAction_SetIconVisibleInMenu(void* ptr, int visible){
	static_cast<QAction*>(ptr)->setIconVisibleInMenu(visible != 0);
}

void QAction_SetMenuRole(void* ptr, int menuRole){
	static_cast<QAction*>(ptr)->setMenuRole(static_cast<QAction::MenuRole>(menuRole));
}

void QAction_SetPriority(void* ptr, int priority){
	static_cast<QAction*>(ptr)->setPriority(static_cast<QAction::Priority>(priority));
}

void QAction_SetShortcut(void* ptr, void* shortcut){
	static_cast<QAction*>(ptr)->setShortcut(*static_cast<QKeySequence*>(shortcut));
}

void QAction_SetShortcutContext(void* ptr, int context){
	static_cast<QAction*>(ptr)->setShortcutContext(static_cast<Qt::ShortcutContext>(context));
}

void QAction_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QAction*>(ptr)->setStatusTip(QString(statusTip));
}

void QAction_SetText(void* ptr, char* text){
	static_cast<QAction*>(ptr)->setText(QString(text));
}

void QAction_SetToolTip(void* ptr, char* tip){
	static_cast<QAction*>(ptr)->setToolTip(QString(tip));
}

void QAction_SetVisible(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

void QAction_SetWhatsThis(void* ptr, char* what){
	static_cast<QAction*>(ptr)->setWhatsThis(QString(what));
}

int QAction_ShortcutContext(void* ptr){
	return static_cast<QAction*>(ptr)->shortcutContext();
}

char* QAction_StatusTip(void* ptr){
	return static_cast<QAction*>(ptr)->statusTip().toUtf8().data();
}

char* QAction_Text(void* ptr){
	return static_cast<QAction*>(ptr)->text().toUtf8().data();
}

void QAction_Toggle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "toggle");
}

char* QAction_ToolTip(void* ptr){
	return static_cast<QAction*>(ptr)->toolTip().toUtf8().data();
}

char* QAction_WhatsThis(void* ptr){
	return static_cast<QAction*>(ptr)->whatsThis().toUtf8().data();
}

void* QAction_NewQAction(void* parent){
	return new QAction(static_cast<QObject*>(parent));
}

void* QAction_NewQAction3(void* icon, char* text, void* parent){
	return new QAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(parent));
}

void* QAction_NewQAction2(char* text, void* parent){
	return new QAction(QString(text), static_cast<QObject*>(parent));
}

void* QAction_ActionGroup(void* ptr){
	return static_cast<QAction*>(ptr)->actionGroup();
}

void QAction_Activate(void* ptr, int event){
	static_cast<QAction*>(ptr)->activate(static_cast<QAction::ActionEvent>(event));
}

void QAction_ConnectChanged(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

void QAction_DisconnectChanged(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::changed), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Changed));;
}

void* QAction_Data(void* ptr){
	return new QVariant(static_cast<QAction*>(ptr)->data());
}

void QAction_Hover(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "hover");
}

void QAction_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

void QAction_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)()>(&QAction::hovered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)()>(&MyQAction::Signal_Hovered));;
}

int QAction_IsSeparator(void* ptr){
	return static_cast<QAction*>(ptr)->isSeparator();
}

void* QAction_Menu(void* ptr){
	return static_cast<QAction*>(ptr)->menu();
}

void* QAction_ParentWidget(void* ptr){
	return static_cast<QAction*>(ptr)->parentWidget();
}

void QAction_SetActionGroup(void* ptr, void* group){
	static_cast<QAction*>(ptr)->setActionGroup(static_cast<QActionGroup*>(group));
}

void QAction_SetDisabled(void* ptr, int b){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QAction_SetMenu(void* ptr, void* menu){
	static_cast<QAction*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QAction_SetSeparator(void* ptr, int b){
	static_cast<QAction*>(ptr)->setSeparator(b != 0);
}

void QAction_SetShortcuts2(void* ptr, int key){
	static_cast<QAction*>(ptr)->setShortcuts(static_cast<QKeySequence::StandardKey>(key));
}

int QAction_ShowStatusText(void* ptr, void* widget){
	return static_cast<QAction*>(ptr)->showStatusText(static_cast<QWidget*>(widget));
}

void QAction_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::toggled), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Toggled));;
}

void QAction_Trigger(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAction*>(ptr), "trigger");
}

void QAction_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QAction*>(ptr), static_cast<void (QAction::*)(bool)>(&QAction::triggered), static_cast<MyQAction*>(ptr), static_cast<void (MyQAction::*)(bool)>(&MyQAction::Signal_Triggered));;
}

void QAction_DestroyQAction(void* ptr){
	static_cast<QAction*>(ptr)->~QAction();
}

#include "qgraphicsscenehoverevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneHoverEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneHoverEvent: public QGraphicsSceneHoverEvent {
public:
};

int QGraphicsSceneHoverEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneHoverEvent*>(ptr)->modifiers();
}

void QGraphicsSceneHoverEvent_DestroyQGraphicsSceneHoverEvent(void* ptr){
	static_cast<QGraphicsSceneHoverEvent*>(ptr)->~QGraphicsSceneHoverEvent();
}

#include "qstyleoptiongraphicsitem.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QTransform>
#include <QString>
#include <QVariant>
#include <QStyleOptionGraphicsItem>
#include "_cgo_export.h"

class MyQStyleOptionGraphicsItem: public QStyleOptionGraphicsItem {
public:
};

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem(){
	return new QStyleOptionGraphicsItem();
}

void* QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(void* other){
	return new QStyleOptionGraphicsItem(*static_cast<QStyleOptionGraphicsItem*>(other));
}

double QStyleOptionGraphicsItem_QStyleOptionGraphicsItem_LevelOfDetailFromTransform(void* worldTransform){
	return static_cast<double>(QStyleOptionGraphicsItem::levelOfDetailFromTransform(*static_cast<QTransform*>(worldTransform)));
}

#include "qscrollbar.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QEvent>
#include <QString>
#include <QScrollBar>
#include "_cgo_export.h"

class MyQScrollBar: public QScrollBar {
public:
};

void* QScrollBar_NewQScrollBar(void* parent){
	return new QScrollBar(static_cast<QWidget*>(parent));
}

void* QScrollBar_NewQScrollBar2(int orientation, void* parent){
	return new QScrollBar(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QScrollBar_Event(void* ptr, void* event){
	return static_cast<QScrollBar*>(ptr)->event(static_cast<QEvent*>(event));
}

void QScrollBar_DestroyQScrollBar(void* ptr){
	static_cast<QScrollBar*>(ptr)->~QScrollBar();
}

#include "qgraphicslinearlayout.h"
#include <QGraphicsLayoutItem>
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsLinearLayout>
#include "_cgo_export.h"

class MyQGraphicsLinearLayout: public QGraphicsLinearLayout {
public:
};

void* QGraphicsLinearLayout_NewQGraphicsLinearLayout(void* parent){
	return new QGraphicsLinearLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void* QGraphicsLinearLayout_NewQGraphicsLinearLayout2(int orientation, void* parent){
	return new QGraphicsLinearLayout(static_cast<Qt::Orientation>(orientation), static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsLinearLayout_AddItem(void* ptr, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_AddStretch(void* ptr, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->addStretch(stretch);
}

int QGraphicsLinearLayout_Alignment(void* ptr, void* item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsLinearLayout_Count(void* ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->count();
}

void QGraphicsLinearLayout_InsertItem(void* ptr, int index, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertItem(index, static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_InsertStretch(void* ptr, int index, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertStretch(index, stretch);
}

void QGraphicsLinearLayout_Invalidate(void* ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->invalidate();
}

void* QGraphicsLinearLayout_ItemAt(void* ptr, int index){
	return static_cast<QGraphicsLinearLayout*>(ptr)->itemAt(index);
}

double QGraphicsLinearLayout_ItemSpacing(void* ptr, int index){
	return static_cast<double>(static_cast<QGraphicsLinearLayout*>(ptr)->itemSpacing(index));
}

int QGraphicsLinearLayout_Orientation(void* ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->orientation();
}

void QGraphicsLinearLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeAt(index);
}

void QGraphicsLinearLayout_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_SetAlignment(void* ptr, void* item, int alignment){
	static_cast<QGraphicsLinearLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsLinearLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsLinearLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsLinearLayout_SetItemSpacing(void* ptr, int index, double spacing){
	static_cast<QGraphicsLinearLayout*>(ptr)->setItemSpacing(index, static_cast<qreal>(spacing));
}

void QGraphicsLinearLayout_SetOrientation(void* ptr, int orientation){
	static_cast<QGraphicsLinearLayout*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QGraphicsLinearLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsLinearLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsLinearLayout_SetStretchFactor(void* ptr, void* item, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->setStretchFactor(static_cast<QGraphicsLayoutItem*>(item), stretch);
}

double QGraphicsLinearLayout_Spacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsLinearLayout*>(ptr)->spacing());
}

int QGraphicsLinearLayout_StretchFactor(void* ptr, void* item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->stretchFactor(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(void* ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->~QGraphicsLinearLayout();
}

#include "qitemeditorcreator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemEditorCreator>
#include "_cgo_export.h"

#include "qspinbox.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QSpinBox>
#include "_cgo_export.h"

class MyQSpinBox: public QSpinBox {
public:
void Signal_ValueChanged(int i){callbackQSpinBoxValueChanged(this->objectName().toUtf8().data(), i);};
};

char* QSpinBox_CleanText(void* ptr){
	return static_cast<QSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QSpinBox_DisplayIntegerBase(void* ptr){
	return static_cast<QSpinBox*>(ptr)->displayIntegerBase();
}

int QSpinBox_Maximum(void* ptr){
	return static_cast<QSpinBox*>(ptr)->maximum();
}

int QSpinBox_Minimum(void* ptr){
	return static_cast<QSpinBox*>(ptr)->minimum();
}

char* QSpinBox_Prefix(void* ptr){
	return static_cast<QSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QSpinBox_SetDisplayIntegerBase(void* ptr, int base){
	static_cast<QSpinBox*>(ptr)->setDisplayIntegerBase(base);
}

void QSpinBox_SetMaximum(void* ptr, int max){
	static_cast<QSpinBox*>(ptr)->setMaximum(max);
}

void QSpinBox_SetMinimum(void* ptr, int min){
	static_cast<QSpinBox*>(ptr)->setMinimum(min);
}

void QSpinBox_SetPrefix(void* ptr, char* prefix){
	static_cast<QSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QSpinBox_SetSingleStep(void* ptr, int val){
	static_cast<QSpinBox*>(ptr)->setSingleStep(val);
}

void QSpinBox_SetSuffix(void* ptr, char* suffix){
	static_cast<QSpinBox*>(ptr)->setSuffix(QString(suffix));
}

void QSpinBox_SetValue(void* ptr, int val){
	QMetaObject::invokeMethod(static_cast<QSpinBox*>(ptr), "setValue", Q_ARG(int, val));
}

int QSpinBox_SingleStep(void* ptr){
	return static_cast<QSpinBox*>(ptr)->singleStep();
}

char* QSpinBox_Suffix(void* ptr){
	return static_cast<QSpinBox*>(ptr)->suffix().toUtf8().data();
}

int QSpinBox_Value(void* ptr){
	return static_cast<QSpinBox*>(ptr)->value();
}

void* QSpinBox_NewQSpinBox(void* parent){
	return new QSpinBox(static_cast<QWidget*>(parent));
}

void QSpinBox_SetRange(void* ptr, int minimum, int maximum){
	static_cast<QSpinBox*>(ptr)->setRange(minimum, maximum);
}

void QSpinBox_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QSpinBox*>(ptr), static_cast<void (QSpinBox::*)(int)>(&QSpinBox::valueChanged), static_cast<MyQSpinBox*>(ptr), static_cast<void (MyQSpinBox::*)(int)>(&MyQSpinBox::Signal_ValueChanged));;
}

void QSpinBox_DestroyQSpinBox(void* ptr){
	static_cast<QSpinBox*>(ptr)->~QSpinBox();
}

#include "qmenu.h"
#include <QPoint>
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QIcon>
#include <QAction>
#include <QKeySequence>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QMenu>
#include "_cgo_export.h"

class MyQMenu: public QMenu {
public:
void Signal_AboutToHide(){callbackQMenuAboutToHide(this->objectName().toUtf8().data());};
void Signal_AboutToShow(){callbackQMenuAboutToShow(this->objectName().toUtf8().data());};
void Signal_Hovered(QAction * action){callbackQMenuHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQMenuTriggered(this->objectName().toUtf8().data(), action);};
};

int QMenu_IsTearOffEnabled(void* ptr){
	return static_cast<QMenu*>(ptr)->isTearOffEnabled();
}

int QMenu_SeparatorsCollapsible(void* ptr){
	return static_cast<QMenu*>(ptr)->separatorsCollapsible();
}

void QMenu_SetIcon(void* ptr, void* icon){
	static_cast<QMenu*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QMenu_SetSeparatorsCollapsible(void* ptr, int collapse){
	static_cast<QMenu*>(ptr)->setSeparatorsCollapsible(collapse != 0);
}

void QMenu_SetTearOffEnabled(void* ptr, int v){
	static_cast<QMenu*>(ptr)->setTearOffEnabled(v != 0);
}

void QMenu_SetTitle(void* ptr, char* title){
	static_cast<QMenu*>(ptr)->setTitle(QString(title));
}

void QMenu_SetToolTipsVisible(void* ptr, int visible){
	static_cast<QMenu*>(ptr)->setToolTipsVisible(visible != 0);
}

char* QMenu_Title(void* ptr){
	return static_cast<QMenu*>(ptr)->title().toUtf8().data();
}

int QMenu_ToolTipsVisible(void* ptr){
	return static_cast<QMenu*>(ptr)->toolTipsVisible();
}

void* QMenu_NewQMenu(void* parent){
	return new QMenu(static_cast<QWidget*>(parent));
}

void* QMenu_NewQMenu2(char* title, void* parent){
	return new QMenu(QString(title), static_cast<QWidget*>(parent));
}

void QMenu_ConnectAboutToHide(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_DisconnectAboutToHide(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToHide), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToHide));;
}

void QMenu_ConnectAboutToShow(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

void QMenu_DisconnectAboutToShow(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)()>(&QMenu::aboutToShow), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)()>(&MyQMenu::Signal_AboutToShow));;
}

void* QMenu_ActionAt(void* ptr, void* pt){
	return static_cast<QMenu*>(ptr)->actionAt(*static_cast<QPoint*>(pt));
}

void* QMenu_ActiveAction(void* ptr){
	return static_cast<QMenu*>(ptr)->activeAction();
}

void* QMenu_AddAction2(void* ptr, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member, void* shortcut){
	return static_cast<QMenu*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

void* QMenu_AddAction(void* ptr, char* text){
	return static_cast<QMenu*>(ptr)->addAction(QString(text));
}

void* QMenu_AddAction3(void* ptr, char* text, void* receiver, char* member, void* shortcut){
	return static_cast<QMenu*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member), *static_cast<QKeySequence*>(shortcut));
}

void* QMenu_AddMenu(void* ptr, void* menu){
	return static_cast<QMenu*>(ptr)->addMenu(static_cast<QMenu*>(menu));
}

void* QMenu_AddMenu3(void* ptr, void* icon, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(*static_cast<QIcon*>(icon), QString(title));
}

void* QMenu_AddMenu2(void* ptr, char* title){
	return static_cast<QMenu*>(ptr)->addMenu(QString(title));
}

void* QMenu_AddSection2(void* ptr, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->addSection(*static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_AddSection(void* ptr, char* text){
	return static_cast<QMenu*>(ptr)->addSection(QString(text));
}

void* QMenu_AddSeparator(void* ptr){
	return static_cast<QMenu*>(ptr)->addSeparator();
}

void QMenu_Clear(void* ptr){
	static_cast<QMenu*>(ptr)->clear();
}

void* QMenu_Exec(void* ptr){
	return static_cast<QMenu*>(ptr)->exec();
}

void* QMenu_Exec2(void* ptr, void* p, void* action){
	return static_cast<QMenu*>(ptr)->exec(*static_cast<QPoint*>(p), static_cast<QAction*>(action));
}

void QMenu_HideTearOffMenu(void* ptr){
	static_cast<QMenu*>(ptr)->hideTearOffMenu();
}

void QMenu_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

void QMenu_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::hovered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Hovered));;
}

void* QMenu_InsertMenu(void* ptr, void* before, void* menu){
	return static_cast<QMenu*>(ptr)->insertMenu(static_cast<QAction*>(before), static_cast<QMenu*>(menu));
}

void* QMenu_InsertSection2(void* ptr, void* before, void* icon, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), *static_cast<QIcon*>(icon), QString(text));
}

void* QMenu_InsertSection(void* ptr, void* before, char* text){
	return static_cast<QMenu*>(ptr)->insertSection(static_cast<QAction*>(before), QString(text));
}

void* QMenu_InsertSeparator(void* ptr, void* before){
	return static_cast<QMenu*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

int QMenu_IsEmpty(void* ptr){
	return static_cast<QMenu*>(ptr)->isEmpty();
}

int QMenu_IsTearOffMenuVisible(void* ptr){
	return static_cast<QMenu*>(ptr)->isTearOffMenuVisible();
}

void* QMenu_MenuAction(void* ptr){
	return static_cast<QMenu*>(ptr)->menuAction();
}

void QMenu_Popup(void* ptr, void* p, void* atAction){
	static_cast<QMenu*>(ptr)->popup(*static_cast<QPoint*>(p), static_cast<QAction*>(atAction));
}

void QMenu_SetActiveAction(void* ptr, void* act){
	static_cast<QMenu*>(ptr)->setActiveAction(static_cast<QAction*>(act));
}

void QMenu_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QMenu*>(ptr), static_cast<void (QMenu::*)(QAction *)>(&QMenu::triggered), static_cast<MyQMenu*>(ptr), static_cast<void (MyQMenu::*)(QAction *)>(&MyQMenu::Signal_Triggered));;
}

void QMenu_DestroyQMenu(void* ptr){
	static_cast<QMenu*>(ptr)->~QMenu();
}

#include "qgraphicsscenehelpevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneHelpEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneHelpEvent: public QGraphicsSceneHelpEvent {
public:
};

void QGraphicsSceneHelpEvent_DestroyQGraphicsSceneHelpEvent(void* ptr){
	static_cast<QGraphicsSceneHelpEvent*>(ptr)->~QGraphicsSceneHelpEvent();
}

#include "qmdiarea.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QBrush>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QMdiSubWindow>
#include <QUrl>
#include <QTabWidget>
#include <QMdiArea>
#include "_cgo_export.h"

class MyQMdiArea: public QMdiArea {
public:
void Signal_SubWindowActivated(QMdiSubWindow * window){callbackQMdiAreaSubWindowActivated(this->objectName().toUtf8().data(), window);};
};

int QMdiArea_ActivationOrder(void* ptr){
	return static_cast<QMdiArea*>(ptr)->activationOrder();
}

void* QMdiArea_Background(void* ptr){
	return new QBrush(static_cast<QMdiArea*>(ptr)->background());
}

int QMdiArea_DocumentMode(void* ptr){
	return static_cast<QMdiArea*>(ptr)->documentMode();
}

void QMdiArea_SetActivationOrder(void* ptr, int order){
	static_cast<QMdiArea*>(ptr)->setActivationOrder(static_cast<QMdiArea::WindowOrder>(order));
}

void QMdiArea_SetBackground(void* ptr, void* background){
	static_cast<QMdiArea*>(ptr)->setBackground(*static_cast<QBrush*>(background));
}

void QMdiArea_SetDocumentMode(void* ptr, int enabled){
	static_cast<QMdiArea*>(ptr)->setDocumentMode(enabled != 0);
}

void QMdiArea_SetTabPosition(void* ptr, int position){
	static_cast<QMdiArea*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(position));
}

void QMdiArea_SetTabShape(void* ptr, int shape){
	static_cast<QMdiArea*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(shape));
}

void QMdiArea_SetTabsClosable(void* ptr, int closable){
	static_cast<QMdiArea*>(ptr)->setTabsClosable(closable != 0);
}

void QMdiArea_SetTabsMovable(void* ptr, int movable){
	static_cast<QMdiArea*>(ptr)->setTabsMovable(movable != 0);
}

void QMdiArea_SetViewMode(void* ptr, int mode){
	static_cast<QMdiArea*>(ptr)->setViewMode(static_cast<QMdiArea::ViewMode>(mode));
}

int QMdiArea_TabPosition(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabPosition();
}

int QMdiArea_TabShape(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabShape();
}

int QMdiArea_TabsClosable(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabsClosable();
}

int QMdiArea_TabsMovable(void* ptr){
	return static_cast<QMdiArea*>(ptr)->tabsMovable();
}

int QMdiArea_ViewMode(void* ptr){
	return static_cast<QMdiArea*>(ptr)->viewMode();
}

void* QMdiArea_NewQMdiArea(void* parent){
	return new QMdiArea(static_cast<QWidget*>(parent));
}

void QMdiArea_ActivateNextSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activateNextSubWindow");
}

void QMdiArea_ActivatePreviousSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "activatePreviousSubWindow");
}

void* QMdiArea_ActiveSubWindow(void* ptr){
	return static_cast<QMdiArea*>(ptr)->activeSubWindow();
}

void* QMdiArea_AddSubWindow(void* ptr, void* widget, int windowFlags){
	return static_cast<QMdiArea*>(ptr)->addSubWindow(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(windowFlags));
}

void QMdiArea_CascadeSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "cascadeSubWindows");
}

void QMdiArea_CloseActiveSubWindow(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeActiveSubWindow");
}

void QMdiArea_CloseAllSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "closeAllSubWindows");
}

void* QMdiArea_CurrentSubWindow(void* ptr){
	return static_cast<QMdiArea*>(ptr)->currentSubWindow();
}

void QMdiArea_RemoveSubWindow(void* ptr, void* widget){
	static_cast<QMdiArea*>(ptr)->removeSubWindow(static_cast<QWidget*>(widget));
}

void QMdiArea_SetActiveSubWindow(void* ptr, void* window){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "setActiveSubWindow", Q_ARG(QMdiSubWindow*, static_cast<QMdiSubWindow*>(window)));
}

void QMdiArea_SetOption(void* ptr, int option, int on){
	static_cast<QMdiArea*>(ptr)->setOption(static_cast<QMdiArea::AreaOption>(option), on != 0);
}

void QMdiArea_ConnectSubWindowActivated(void* ptr){
	QObject::connect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

void QMdiArea_DisconnectSubWindowActivated(void* ptr){
	QObject::disconnect(static_cast<QMdiArea*>(ptr), static_cast<void (QMdiArea::*)(QMdiSubWindow *)>(&QMdiArea::subWindowActivated), static_cast<MyQMdiArea*>(ptr), static_cast<void (MyQMdiArea::*)(QMdiSubWindow *)>(&MyQMdiArea::Signal_SubWindowActivated));;
}

int QMdiArea_TestOption(void* ptr, int option){
	return static_cast<QMdiArea*>(ptr)->testOption(static_cast<QMdiArea::AreaOption>(option));
}

void QMdiArea_TileSubWindows(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMdiArea*>(ptr), "tileSubWindows");
}

void QMdiArea_DestroyQMdiArea(void* ptr){
	static_cast<QMdiArea*>(ptr)->~QMdiArea();
}

#include "qtableview.h"
#include <QModelIndex>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QAbstractItemModel>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHeaderView>
#include <QMetaObject>
#include <QTableView>
#include "_cgo_export.h"

class MyQTableView: public QTableView {
public:
};

int QTableView_GridStyle(void* ptr){
	return static_cast<QTableView*>(ptr)->gridStyle();
}

int QTableView_IsCornerButtonEnabled(void* ptr){
	return static_cast<QTableView*>(ptr)->isCornerButtonEnabled();
}

int QTableView_IsSortingEnabled(void* ptr){
	return static_cast<QTableView*>(ptr)->isSortingEnabled();
}

void QTableView_SetCornerButtonEnabled(void* ptr, int enable){
	static_cast<QTableView*>(ptr)->setCornerButtonEnabled(enable != 0);
}

void QTableView_SetGridStyle(void* ptr, int style){
	static_cast<QTableView*>(ptr)->setGridStyle(static_cast<Qt::PenStyle>(style));
}

void QTableView_SetShowGrid(void* ptr, int show){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "setShowGrid", Q_ARG(bool, show != 0));
}

void QTableView_SetSpan(void* ptr, int row, int column, int rowSpanCount, int columnSpanCount){
	static_cast<QTableView*>(ptr)->setSpan(row, column, rowSpanCount, columnSpanCount);
}

void QTableView_SetWordWrap(void* ptr, int on){
	static_cast<QTableView*>(ptr)->setWordWrap(on != 0);
}

int QTableView_ShowGrid(void* ptr){
	return static_cast<QTableView*>(ptr)->showGrid();
}

int QTableView_WordWrap(void* ptr){
	return static_cast<QTableView*>(ptr)->wordWrap();
}

void QTableView_ClearSpans(void* ptr){
	static_cast<QTableView*>(ptr)->clearSpans();
}

int QTableView_ColumnAt(void* ptr, int x){
	return static_cast<QTableView*>(ptr)->columnAt(x);
}

int QTableView_ColumnSpan(void* ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->columnSpan(row, column);
}

int QTableView_ColumnViewportPosition(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->columnViewportPosition(column);
}

int QTableView_ColumnWidth(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->columnWidth(column);
}

void QTableView_HideColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideColumn", Q_ARG(int, column));
}

void QTableView_HideRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideRow", Q_ARG(int, row));
}

void* QTableView_HorizontalHeader(void* ptr){
	return static_cast<QTableView*>(ptr)->horizontalHeader();
}

void* QTableView_IndexAt(void* ptr, void* pos){
	return static_cast<QTableView*>(ptr)->indexAt(*static_cast<QPoint*>(pos)).internalPointer();
}

int QTableView_IsColumnHidden(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->isColumnHidden(column);
}

int QTableView_IsRowHidden(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->isRowHidden(row);
}

void QTableView_ResizeColumnToContents(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTableView_ResizeColumnsToContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnsToContents");
}

void QTableView_ResizeRowToContents(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowToContents", Q_ARG(int, row));
}

void QTableView_ResizeRowsToContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowsToContents");
}

int QTableView_RowAt(void* ptr, int y){
	return static_cast<QTableView*>(ptr)->rowAt(y);
}

int QTableView_RowHeight(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->rowHeight(row);
}

int QTableView_RowSpan(void* ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->rowSpan(row, column);
}

int QTableView_RowViewportPosition(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->rowViewportPosition(row);
}

void QTableView_SelectColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectColumn", Q_ARG(int, column));
}

void QTableView_SelectRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectRow", Q_ARG(int, row));
}

void QTableView_SetColumnHidden(void* ptr, int column, int hide){
	static_cast<QTableView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTableView_SetColumnWidth(void* ptr, int column, int width){
	static_cast<QTableView*>(ptr)->setColumnWidth(column, width);
}

void QTableView_SetHorizontalHeader(void* ptr, void* header){
	static_cast<QTableView*>(ptr)->setHorizontalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_SetModel(void* ptr, void* model){
	static_cast<QTableView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTableView_SetRootIndex(void* ptr, void* index){
	static_cast<QTableView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTableView_SetRowHeight(void* ptr, int row, int height){
	static_cast<QTableView*>(ptr)->setRowHeight(row, height);
}

void QTableView_SetRowHidden(void* ptr, int row, int hide){
	static_cast<QTableView*>(ptr)->setRowHidden(row, hide != 0);
}

void QTableView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTableView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTableView_SetSortingEnabled(void* ptr, int enable){
	static_cast<QTableView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTableView_SetVerticalHeader(void* ptr, void* header){
	static_cast<QTableView*>(ptr)->setVerticalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_ShowColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTableView_ShowRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showRow", Q_ARG(int, row));
}

void QTableView_SortByColumn(void* ptr, int column, int order){
	static_cast<QTableView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

void* QTableView_VerticalHeader(void* ptr){
	return static_cast<QTableView*>(ptr)->verticalHeader();
}

void QTableView_DestroyQTableView(void* ptr){
	static_cast<QTableView*>(ptr)->~QTableView();
}

#include "qstyleoptionslider.h"
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyleOptionSlider>
#include "_cgo_export.h"

class MyQStyleOptionSlider: public QStyleOptionSlider {
public:
};

void* QStyleOptionSlider_NewQStyleOptionSlider(){
	return new QStyleOptionSlider();
}

void* QStyleOptionSlider_NewQStyleOptionSlider2(void* other){
	return new QStyleOptionSlider(*static_cast<QStyleOptionSlider*>(other));
}

#include "qproxystyle.h"
#include <QStyleOptionComplex>
#include <QStyle>
#include <QVariant>
#include <QApplication>
#include <QSize>
#include <QString>
#include <QRect>
#include <QPixmap>
#include <QStyleHintReturn>
#include <QPoint>
#include <QWidget>
#include <QUrl>
#include <QSizePolicy>
#include <QPainter>
#include <QPalette>
#include <QModelIndex>
#include <QStyleOption>
#include <QProxyStyle>
#include "_cgo_export.h"

class MyQProxyStyle: public QProxyStyle {
public:
};

void* QProxyStyle_BaseStyle(void* ptr){
	return static_cast<QProxyStyle*>(ptr)->baseStyle();
}

void QProxyStyle_DrawComplexControl(void* ptr, int control, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawControl(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QProxyStyle_DrawItemPixmap(void* ptr, void* painter, void* rect, int alignment, void* pixmap){
	static_cast<QProxyStyle*>(ptr)->drawItemPixmap(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), alignment, *static_cast<QPixmap*>(pixmap));
}

void QProxyStyle_DrawItemText(void* ptr, void* painter, void* rect, int flags, void* pal, int enabled, char* text, int textRole){
	static_cast<QProxyStyle*>(ptr)->drawItemText(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QProxyStyle_DrawPrimitive(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QProxyStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

int QProxyStyle_HitTestComplexControl(void* ptr, int control, void* option, void* pos, void* widget){
	return static_cast<QProxyStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), *static_cast<QPoint*>(pos), static_cast<QWidget*>(widget));
}

int QProxyStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget){
	return static_cast<QProxyStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QProxyStyle_PixelMetric(void* ptr, int metric, void* option, void* widget){
	return static_cast<QProxyStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(metric), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void QProxyStyle_Polish3(void* ptr, void* app){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QProxyStyle_Polish2(void* ptr, void* pal){
	static_cast<QProxyStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QProxyStyle_Polish(void* ptr, void* widget){
	static_cast<QProxyStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

void QProxyStyle_SetBaseStyle(void* ptr, void* style){
	static_cast<QProxyStyle*>(ptr)->setBaseStyle(static_cast<QStyle*>(style));
}

int QProxyStyle_StyleHint(void* ptr, int hint, void* option, void* widget, void* returnData){
	return static_cast<QProxyStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(hint), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(returnData));
}

void QProxyStyle_Unpolish2(void* ptr, void* app){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QApplication*>(app));
}

void QProxyStyle_Unpolish(void* ptr, void* widget){
	static_cast<QProxyStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QProxyStyle_DestroyQProxyStyle(void* ptr){
	static_cast<QProxyStyle*>(ptr)->~QProxyStyle();
}

#include "qprogressdialog.h"
#include <QString>
#include <QUrl>
#include <QLabel>
#include <QMetaObject>
#include <QProgressBar>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QPushButton>
#include <QProgressDialog>
#include "_cgo_export.h"

class MyQProgressDialog: public QProgressDialog {
public:
void Signal_Canceled(){callbackQProgressDialogCanceled(this->objectName().toUtf8().data());};
};

int QProgressDialog_AutoClose(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->autoClose();
}

int QProgressDialog_AutoReset(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->autoReset();
}

char* QProgressDialog_LabelText(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->labelText().toUtf8().data();
}

int QProgressDialog_Maximum(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->maximum();
}

int QProgressDialog_Minimum(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->minimum();
}

int QProgressDialog_MinimumDuration(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->minimumDuration();
}

void QProgressDialog_SetAutoClose(void* ptr, int close){
	static_cast<QProgressDialog*>(ptr)->setAutoClose(close != 0);
}

void QProgressDialog_SetAutoReset(void* ptr, int reset){
	static_cast<QProgressDialog*>(ptr)->setAutoReset(reset != 0);
}

void QProgressDialog_SetLabelText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setLabelText", Q_ARG(QString, QString(text)));
}

void QProgressDialog_SetMaximum(void* ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressDialog_SetMinimum(void* ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressDialog_SetMinimumDuration(void* ptr, int ms){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimumDuration", Q_ARG(int, ms));
}

void QProgressDialog_SetValue(void* ptr, int progress){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setValue", Q_ARG(int, progress));
}

int QProgressDialog_Value(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->value();
}

int QProgressDialog_WasCanceled(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->wasCanceled();
}

void* QProgressDialog_NewQProgressDialog(void* parent, int f){
	return new QProgressDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QProgressDialog_NewQProgressDialog2(char* labelText, char* cancelButtonText, int minimum, int maximum, void* parent, int f){
	return new QProgressDialog(QString(labelText), QString(cancelButtonText), minimum, maximum, static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QProgressDialog_Cancel(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "cancel");
}

void QProgressDialog_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QProgressDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QProgressDialog_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "reset");
}

void QProgressDialog_SetBar(void* ptr, void* bar){
	static_cast<QProgressDialog*>(ptr)->setBar(static_cast<QProgressBar*>(bar));
}

void QProgressDialog_SetCancelButton(void* ptr, void* cancelButton){
	static_cast<QProgressDialog*>(ptr)->setCancelButton(static_cast<QPushButton*>(cancelButton));
}

void QProgressDialog_SetCancelButtonText(void* ptr, char* cancelButtonText){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setCancelButtonText", Q_ARG(QString, QString(cancelButtonText)));
}

void QProgressDialog_SetLabel(void* ptr, void* label){
	static_cast<QProgressDialog*>(ptr)->setLabel(static_cast<QLabel*>(label));
}

void QProgressDialog_SetRange(void* ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressDialog_DestroyQProgressDialog(void* ptr){
	static_cast<QProgressDialog*>(ptr)->~QProgressDialog();
}

#include "qtextbrowser.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QTextBrowser>
#include "_cgo_export.h"

class MyQTextBrowser: public QTextBrowser {
public:
void Signal_BackwardAvailable(bool available){callbackQTextBrowserBackwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_ForwardAvailable(bool available){callbackQTextBrowserForwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_HistoryChanged(){callbackQTextBrowserHistoryChanged(this->objectName().toUtf8().data());};
};

int QTextBrowser_OpenExternalLinks(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->openExternalLinks();
}

int QTextBrowser_OpenLinks(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->openLinks();
}

char* QTextBrowser_SearchPaths(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->searchPaths().join("|").toUtf8().data();
}

void QTextBrowser_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenExternalLinks(open != 0);
}

void QTextBrowser_SetOpenLinks(void* ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenLinks(open != 0);
}

void QTextBrowser_SetSearchPaths(void* ptr, char* paths){
	static_cast<QTextBrowser*>(ptr)->setSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QTextBrowser_SetSource(void* ptr, void* name){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(name)));
}

void* QTextBrowser_NewQTextBrowser(void* parent){
	return new QTextBrowser(static_cast<QWidget*>(parent));
}

void QTextBrowser_Backward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "backward");
}

void QTextBrowser_ConnectBackwardAvailable(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

void QTextBrowser_DisconnectBackwardAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

int QTextBrowser_BackwardHistoryCount(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->backwardHistoryCount();
}

void QTextBrowser_ClearHistory(void* ptr){
	static_cast<QTextBrowser*>(ptr)->clearHistory();
}

void QTextBrowser_Forward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "forward");
}

void QTextBrowser_ConnectForwardAvailable(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

void QTextBrowser_DisconnectForwardAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

int QTextBrowser_ForwardHistoryCount(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->forwardHistoryCount();
}

void QTextBrowser_ConnectHistoryChanged(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

void QTextBrowser_DisconnectHistoryChanged(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

char* QTextBrowser_HistoryTitle(void* ptr, int i){
	return static_cast<QTextBrowser*>(ptr)->historyTitle(i).toUtf8().data();
}

void QTextBrowser_Home(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "home");
}

int QTextBrowser_IsBackwardAvailable(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->isBackwardAvailable();
}

int QTextBrowser_IsForwardAvailable(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->isForwardAvailable();
}

void* QTextBrowser_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QTextBrowser*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

void QTextBrowser_Reload(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "reload");
}

#include "qdatawidgetmapper.h"
#include <QAbstractItemModel>
#include <QAbstractItemDelegate>
#include <QString>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QVariant>
#include <QUrl>
#include <QByteArray>
#include <QDataWidgetMapper>
#include "_cgo_export.h"

class MyQDataWidgetMapper: public QDataWidgetMapper {
public:
void Signal_CurrentIndexChanged(int index){callbackQDataWidgetMapperCurrentIndexChanged(this->objectName().toUtf8().data(), index);};
};

int QDataWidgetMapper_CurrentIndex(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->currentIndex();
}

int QDataWidgetMapper_Orientation(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->orientation();
}

void QDataWidgetMapper_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QDataWidgetMapper_SetOrientation(void* ptr, int aOrientation){
	static_cast<QDataWidgetMapper*>(ptr)->setOrientation(static_cast<Qt::Orientation>(aOrientation));
}

void QDataWidgetMapper_SetSubmitPolicy(void* ptr, int policy){
	static_cast<QDataWidgetMapper*>(ptr)->setSubmitPolicy(static_cast<QDataWidgetMapper::SubmitPolicy>(policy));
}

int QDataWidgetMapper_SubmitPolicy(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->submitPolicy();
}

void* QDataWidgetMapper_NewQDataWidgetMapper(void* parent){
	return new QDataWidgetMapper(static_cast<QObject*>(parent));
}

void QDataWidgetMapper_AddMapping(void* ptr, void* widget, int section){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section);
}

void QDataWidgetMapper_AddMapping2(void* ptr, void* widget, int section, void* propertyName){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section, *static_cast<QByteArray*>(propertyName));
}

void QDataWidgetMapper_ClearMapping(void* ptr){
	static_cast<QDataWidgetMapper*>(ptr)->clearMapping();
}

void QDataWidgetMapper_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

void QDataWidgetMapper_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

void* QDataWidgetMapper_ItemDelegate(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->itemDelegate();
}

void* QDataWidgetMapper_MappedPropertyName(void* ptr, void* widget){
	return new QByteArray(static_cast<QDataWidgetMapper*>(ptr)->mappedPropertyName(static_cast<QWidget*>(widget)));
}

int QDataWidgetMapper_MappedSection(void* ptr, void* widget){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedSection(static_cast<QWidget*>(widget));
}

void* QDataWidgetMapper_MappedWidgetAt(void* ptr, int section){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedWidgetAt(section);
}

void* QDataWidgetMapper_Model(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->model();
}

void QDataWidgetMapper_RemoveMapping(void* ptr, void* widget){
	static_cast<QDataWidgetMapper*>(ptr)->removeMapping(static_cast<QWidget*>(widget));
}

void QDataWidgetMapper_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "revert");
}

void* QDataWidgetMapper_RootIndex(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->rootIndex().internalPointer();
}

void QDataWidgetMapper_SetCurrentModelIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentModelIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QDataWidgetMapper_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QDataWidgetMapper*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QDataWidgetMapper_SetModel(void* ptr, void* model){
	static_cast<QDataWidgetMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QDataWidgetMapper_SetRootIndex(void* ptr, void* index){
	static_cast<QDataWidgetMapper*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

int QDataWidgetMapper_Submit(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "submit");
}

void QDataWidgetMapper_ToFirst(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toFirst");
}

void QDataWidgetMapper_ToLast(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toLast");
}

void QDataWidgetMapper_ToNext(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toNext");
}

void QDataWidgetMapper_ToPrevious(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toPrevious");
}

void QDataWidgetMapper_DestroyQDataWidgetMapper(void* ptr){
	static_cast<QDataWidgetMapper*>(ptr)->~QDataWidgetMapper();
}

#include "qstyleoption.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QWidget>
#include <QStyleOption>
#include "_cgo_export.h"

class MyQStyleOption: public QStyleOption {
public:
};

int QStyleOption_SO_Slider_Type(){
	return QStyleOption::SO_Slider;
}

int QStyleOption_SO_SpinBox_Type(){
	return QStyleOption::SO_SpinBox;
}

int QStyleOption_SO_ToolButton_Type(){
	return QStyleOption::SO_ToolButton;
}

int QStyleOption_SO_ComboBox_Type(){
	return QStyleOption::SO_ComboBox;
}

int QStyleOption_SO_TitleBar_Type(){
	return QStyleOption::SO_TitleBar;
}

int QStyleOption_SO_GroupBox_Type(){
	return QStyleOption::SO_GroupBox;
}

int QStyleOption_SO_SizeGrip_Type(){
	return QStyleOption::SO_SizeGrip;
}

void* QStyleOption_NewQStyleOption2(void* other){
	return new QStyleOption(*static_cast<QStyleOption*>(other));
}

void* QStyleOption_NewQStyleOption(int version, int ty){
	return new QStyleOption(version, ty);
}

void QStyleOption_InitFrom(void* ptr, void* widget){
	static_cast<QStyleOption*>(ptr)->initFrom(static_cast<QWidget*>(widget));
}

void QStyleOption_DestroyQStyleOption(void* ptr){
	static_cast<QStyleOption*>(ptr)->~QStyleOption();
}

#include "qcheckbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QCheckBox>
#include "_cgo_export.h"

class MyQCheckBox: public QCheckBox {
public:
void Signal_StateChanged(int state){callbackQCheckBoxStateChanged(this->objectName().toUtf8().data(), state);};
};

int QCheckBox_IsTristate(void* ptr){
	return static_cast<QCheckBox*>(ptr)->isTristate();
}

void QCheckBox_SetTristate(void* ptr, int y){
	static_cast<QCheckBox*>(ptr)->setTristate(y != 0);
}

void* QCheckBox_NewQCheckBox(void* parent){
	return new QCheckBox(static_cast<QWidget*>(parent));
}

void* QCheckBox_NewQCheckBox2(char* text, void* parent){
	return new QCheckBox(QString(text), static_cast<QWidget*>(parent));
}

int QCheckBox_CheckState(void* ptr){
	return static_cast<QCheckBox*>(ptr)->checkState();
}

void QCheckBox_SetCheckState(void* ptr, int state){
	static_cast<QCheckBox*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QCheckBox_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QCheckBox*>(ptr), static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), static_cast<MyQCheckBox*>(ptr), static_cast<void (MyQCheckBox::*)(int)>(&MyQCheckBox::Signal_StateChanged));;
}

void QCheckBox_DestroyQCheckBox(void* ptr){
	static_cast<QCheckBox*>(ptr)->~QCheckBox();
}

#include "qdockwidget.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QDockWidget>
#include "_cgo_export.h"

class MyQDockWidget: public QDockWidget {
public:
void Signal_AllowedAreasChanged(Qt::DockWidgetAreas allowedAreas){callbackQDockWidgetAllowedAreasChanged(this->objectName().toUtf8().data(), allowedAreas);};
void Signal_DockLocationChanged(Qt::DockWidgetArea area){callbackQDockWidgetDockLocationChanged(this->objectName().toUtf8().data(), area);};
void Signal_FeaturesChanged(QDockWidget::DockWidgetFeatures features){callbackQDockWidgetFeaturesChanged(this->objectName().toUtf8().data(), features);};
void Signal_TopLevelChanged(bool topLevel){callbackQDockWidgetTopLevelChanged(this->objectName().toUtf8().data(), topLevel);};
void Signal_VisibilityChanged(bool visible){callbackQDockWidgetVisibilityChanged(this->objectName().toUtf8().data(), visible);};
};

int QDockWidget_AllowedAreas(void* ptr){
	return static_cast<QDockWidget*>(ptr)->allowedAreas();
}

int QDockWidget_Features(void* ptr){
	return static_cast<QDockWidget*>(ptr)->features();
}

void QDockWidget_SetAllowedAreas(void* ptr, int areas){
	static_cast<QDockWidget*>(ptr)->setAllowedAreas(static_cast<Qt::DockWidgetArea>(areas));
}

void QDockWidget_SetFeatures(void* ptr, int features){
	static_cast<QDockWidget*>(ptr)->setFeatures(static_cast<QDockWidget::DockWidgetFeature>(features));
}

void QDockWidget_SetFloating(void* ptr, int floating){
	static_cast<QDockWidget*>(ptr)->setFloating(floating != 0);
}

void* QDockWidget_NewQDockWidget2(void* parent, int flags){
	return new QDockWidget(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void* QDockWidget_NewQDockWidget(char* title, void* parent, int flags){
	return new QDockWidget(QString(title), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QDockWidget_ConnectAllowedAreasChanged(void* ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetAreas)>(&QDockWidget::allowedAreasChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetAreas)>(&MyQDockWidget::Signal_AllowedAreasChanged));;
}

void QDockWidget_DisconnectAllowedAreasChanged(void* ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetAreas)>(&QDockWidget::allowedAreasChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetAreas)>(&MyQDockWidget::Signal_AllowedAreasChanged));;
}

void QDockWidget_ConnectDockLocationChanged(void* ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetArea)>(&QDockWidget::dockLocationChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetArea)>(&MyQDockWidget::Signal_DockLocationChanged));;
}

void QDockWidget_DisconnectDockLocationChanged(void* ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(Qt::DockWidgetArea)>(&QDockWidget::dockLocationChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(Qt::DockWidgetArea)>(&MyQDockWidget::Signal_DockLocationChanged));;
}

void QDockWidget_ConnectFeaturesChanged(void* ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&QDockWidget::featuresChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&MyQDockWidget::Signal_FeaturesChanged));;
}

void QDockWidget_DisconnectFeaturesChanged(void* ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&QDockWidget::featuresChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(QDockWidget::DockWidgetFeatures)>(&MyQDockWidget::Signal_FeaturesChanged));;
}

int QDockWidget_IsAreaAllowed(void* ptr, int area){
	return static_cast<QDockWidget*>(ptr)->isAreaAllowed(static_cast<Qt::DockWidgetArea>(area));
}

int QDockWidget_IsFloating(void* ptr){
	return static_cast<QDockWidget*>(ptr)->isFloating();
}

void QDockWidget_SetTitleBarWidget(void* ptr, void* widget){
	static_cast<QDockWidget*>(ptr)->setTitleBarWidget(static_cast<QWidget*>(widget));
}

void QDockWidget_SetWidget(void* ptr, void* widget){
	static_cast<QDockWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void* QDockWidget_TitleBarWidget(void* ptr){
	return static_cast<QDockWidget*>(ptr)->titleBarWidget();
}

void* QDockWidget_ToggleViewAction(void* ptr){
	return static_cast<QDockWidget*>(ptr)->toggleViewAction();
}

void QDockWidget_ConnectTopLevelChanged(void* ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::topLevelChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_TopLevelChanged));;
}

void QDockWidget_DisconnectTopLevelChanged(void* ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::topLevelChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_TopLevelChanged));;
}

void QDockWidget_ConnectVisibilityChanged(void* ptr){
	QObject::connect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::visibilityChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_VisibilityChanged));;
}

void QDockWidget_DisconnectVisibilityChanged(void* ptr){
	QObject::disconnect(static_cast<QDockWidget*>(ptr), static_cast<void (QDockWidget::*)(bool)>(&QDockWidget::visibilityChanged), static_cast<MyQDockWidget*>(ptr), static_cast<void (MyQDockWidget::*)(bool)>(&MyQDockWidget::Signal_VisibilityChanged));;
}

void* QDockWidget_Widget(void* ptr){
	return static_cast<QDockWidget*>(ptr)->widget();
}

void QDockWidget_DestroyQDockWidget(void* ptr){
	static_cast<QDockWidget*>(ptr)->~QDockWidget();
}

#include "qdialogbuttonbox.h"
#include <QModelIndex>
#include <QAbstractButton>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDialog>
#include <QDial>
#include <QObject>
#include <QDialogButtonBox>
#include "_cgo_export.h"

class MyQDialogButtonBox: public QDialogButtonBox {
public:
void Signal_Accepted(){callbackQDialogButtonBoxAccepted(this->objectName().toUtf8().data());};
void Signal_Clicked(QAbstractButton * button){callbackQDialogButtonBoxClicked(this->objectName().toUtf8().data(), button);};
void Signal_HelpRequested(){callbackQDialogButtonBoxHelpRequested(this->objectName().toUtf8().data());};
void Signal_Rejected(){callbackQDialogButtonBoxRejected(this->objectName().toUtf8().data());};
};

int QDialogButtonBox_CenterButtons(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->centerButtons();
}

int QDialogButtonBox_Orientation(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->orientation();
}

void QDialogButtonBox_SetCenterButtons(void* ptr, int center){
	static_cast<QDialogButtonBox*>(ptr)->setCenterButtons(center != 0);
}

void QDialogButtonBox_SetOrientation(void* ptr, int orientation){
	static_cast<QDialogButtonBox*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QDialogButtonBox_SetStandardButtons(void* ptr, int buttons){
	static_cast<QDialogButtonBox*>(ptr)->setStandardButtons(static_cast<QDialogButtonBox::StandardButton>(buttons));
}

int QDialogButtonBox_StandardButtons(void* ptr){
	return static_cast<QDialogButtonBox*>(ptr)->standardButtons();
}

void* QDialogButtonBox_NewQDialogButtonBox(void* parent){
	return new QDialogButtonBox(static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox2(int orientation, void* parent){
	return new QDialogButtonBox(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox3(int buttons, void* parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<QWidget*>(parent));
}

void* QDialogButtonBox_NewQDialogButtonBox4(int buttons, int orientation, void* parent){
	return new QDialogButtonBox(static_cast<QDialogButtonBox::StandardButton>(buttons), static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QDialogButtonBox_ConnectAccepted(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

void QDialogButtonBox_DisconnectAccepted(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Accepted));;
}

void* QDialogButtonBox_AddButton3(void* ptr, int button){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QDialogButtonBox::StandardButton>(button));
}

void* QDialogButtonBox_AddButton2(void* ptr, char* text, int role){
	return static_cast<QDialogButtonBox*>(ptr)->addButton(QString(text), static_cast<QDialogButtonBox::ButtonRole>(role));
}

void QDialogButtonBox_AddButton(void* ptr, void* button, int role){
	static_cast<QDialogButtonBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QDialogButtonBox::ButtonRole>(role));
}

void* QDialogButtonBox_Button(void* ptr, int which){
	return static_cast<QDialogButtonBox*>(ptr)->button(static_cast<QDialogButtonBox::StandardButton>(which));
}

int QDialogButtonBox_ButtonRole(void* ptr, void* button){
	return static_cast<QDialogButtonBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_Clear(void* ptr){
	static_cast<QDialogButtonBox*>(ptr)->clear();
}

void QDialogButtonBox_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)(QAbstractButton *)>(&QDialogButtonBox::clicked), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)(QAbstractButton *)>(&MyQDialogButtonBox::Signal_Clicked));;
}

void QDialogButtonBox_ConnectHelpRequested(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_DisconnectHelpRequested(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_HelpRequested));;
}

void QDialogButtonBox_ConnectRejected(void* ptr){
	QObject::connect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_DisconnectRejected(void* ptr){
	QObject::disconnect(static_cast<QDialogButtonBox*>(ptr), static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), static_cast<MyQDialogButtonBox*>(ptr), static_cast<void (MyQDialogButtonBox::*)()>(&MyQDialogButtonBox::Signal_Rejected));;
}

void QDialogButtonBox_RemoveButton(void* ptr, void* button){
	static_cast<QDialogButtonBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

int QDialogButtonBox_StandardButton(void* ptr, void* button){
	return static_cast<QDialogButtonBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

void QDialogButtonBox_DestroyQDialogButtonBox(void* ptr){
	static_cast<QDialogButtonBox*>(ptr)->~QDialogButtonBox();
}

#include "qerrormessage.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QErrorMessage>
#include "_cgo_export.h"

class MyQErrorMessage: public QErrorMessage {
public:
};

void* QErrorMessage_NewQErrorMessage(void* parent){
	return new QErrorMessage(static_cast<QWidget*>(parent));
}

void* QErrorMessage_QErrorMessage_QtHandler(){
	return QErrorMessage::qtHandler();
}

void QErrorMessage_ShowMessage(void* ptr, char* message){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)));
}

void QErrorMessage_ShowMessage2(void* ptr, char* message, char* ty){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(QString, QString(ty)));
}

void QErrorMessage_DestroyQErrorMessage(void* ptr){
	static_cast<QErrorMessage*>(ptr)->~QErrorMessage();
}

#include "qabstractbutton.h"
#include <QMetaObject>
#include <QModelIndex>
#include <QIcon>
#include <QKeySequence>
#include <QObject>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractButton>
#include "_cgo_export.h"

class MyQAbstractButton: public QAbstractButton {
public:
void Signal_Clicked(bool checked){callbackQAbstractButtonClicked(this->objectName().toUtf8().data(), checked);};
void Signal_Pressed(){callbackQAbstractButtonPressed(this->objectName().toUtf8().data());};
void Signal_Released(){callbackQAbstractButtonReleased(this->objectName().toUtf8().data());};
void Signal_Toggled(bool checked){callbackQAbstractButtonToggled(this->objectName().toUtf8().data(), checked);};
};

int QAbstractButton_AutoExclusive(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoExclusive();
}

int QAbstractButton_AutoRepeat(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeat();
}

int QAbstractButton_AutoRepeatDelay(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatDelay();
}

int QAbstractButton_AutoRepeatInterval(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->autoRepeatInterval();
}

int QAbstractButton_IsCheckable(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isCheckable();
}

int QAbstractButton_IsChecked(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isChecked();
}

int QAbstractButton_IsDown(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->isDown();
}

void QAbstractButton_SetAutoExclusive(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoExclusive(v != 0);
}

void QAbstractButton_SetAutoRepeat(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeat(v != 0);
}

void QAbstractButton_SetAutoRepeatDelay(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatDelay(v);
}

void QAbstractButton_SetAutoRepeatInterval(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setAutoRepeatInterval(v);
}

void QAbstractButton_SetCheckable(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setCheckable(v != 0);
}

void QAbstractButton_SetChecked(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setChecked", Q_ARG(bool, v != 0));
}

void QAbstractButton_SetDown(void* ptr, int v){
	static_cast<QAbstractButton*>(ptr)->setDown(v != 0);
}

void QAbstractButton_SetIcon(void* ptr, void* icon){
	static_cast<QAbstractButton*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QAbstractButton_SetIconSize(void* ptr, void* size){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(size)));
}

void QAbstractButton_SetShortcut(void* ptr, void* key){
	static_cast<QAbstractButton*>(ptr)->setShortcut(*static_cast<QKeySequence*>(key));
}

void QAbstractButton_SetText(void* ptr, char* text){
	static_cast<QAbstractButton*>(ptr)->setText(QString(text));
}

char* QAbstractButton_Text(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->text().toUtf8().data();
}

void QAbstractButton_Toggle(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "toggle");
}

void QAbstractButton_AnimateClick(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "animateClick", Q_ARG(int, msec));
}

void QAbstractButton_Click(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractButton*>(ptr), "click");
}

void QAbstractButton_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

void QAbstractButton_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Clicked));;
}

void* QAbstractButton_Group(void* ptr){
	return static_cast<QAbstractButton*>(ptr)->group();
}

void QAbstractButton_ConnectPressed(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_DisconnectPressed(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Pressed));;
}

void QAbstractButton_ConnectReleased(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_DisconnectReleased(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)()>(&MyQAbstractButton::Signal_Released));;
}

void QAbstractButton_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QAbstractButton*>(ptr), static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), static_cast<MyQAbstractButton*>(ptr), static_cast<void (MyQAbstractButton::*)(bool)>(&MyQAbstractButton::Signal_Toggled));;
}

void QAbstractButton_DestroyQAbstractButton(void* ptr){
	static_cast<QAbstractButton*>(ptr)->~QAbstractButton();
}

#include "qabstractitemdelegate.h"
#include <QAbstractItemView>
#include <QStyleOptionViewItem>
#include <QHelpEvent>
#include <QUrl>
#include <QWidget>
#include <QAbstractItemModel>
#include <QEvent>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QStyleOption>
#include <QStyle>
#include <QAbstractItemDelegate>
#include "_cgo_export.h"

class MyQAbstractItemDelegate: public QAbstractItemDelegate {
public:
void Signal_CloseEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint){callbackQAbstractItemDelegateCloseEditor(this->objectName().toUtf8().data(), editor, hint);};
void Signal_CommitData(QWidget * editor){callbackQAbstractItemDelegateCommitData(this->objectName().toUtf8().data(), editor);};
void Signal_SizeHintChanged(const QModelIndex & index){callbackQAbstractItemDelegateSizeHintChanged(this->objectName().toUtf8().data(), index.internalPointer());};
};

void QAbstractItemDelegate_ConnectCloseEditor(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_DisconnectCloseEditor(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_ConnectCommitData(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

void QAbstractItemDelegate_DisconnectCommitData(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

void* QAbstractItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyEditor(void* ptr, void* editor, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_ConnectSizeHintChanged(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_DisconnectSizeHintChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyQAbstractItemDelegate(void* ptr){
	static_cast<QAbstractItemDelegate*>(ptr)->~QAbstractItemDelegate();
}

#include "qgraphicsview.h"
#include <QModelIndex>
#include <QString>
#include <QWidget>
#include <QPointF>
#include <QBrush>
#include <QMetaObject>
#include <QGraphicsItem>
#include <QUrl>
#include <QRect>
#include <QRectF>
#include <QPainter>
#include <QTransform>
#include <QGraphicsScene>
#include <QPoint>
#include <QVariant>
#include <QGraphicsView>
#include "_cgo_export.h"

class MyQGraphicsView: public QGraphicsView {
public:
};

int QGraphicsView_Alignment(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->alignment();
}

void* QGraphicsView_BackgroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsView*>(ptr)->backgroundBrush());
}

int QGraphicsView_CacheMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->cacheMode();
}

int QGraphicsView_DragMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->dragMode();
}

void* QGraphicsView_ForegroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsView*>(ptr)->foregroundBrush());
}

int QGraphicsView_IsInteractive(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->isInteractive();
}

int QGraphicsView_OptimizationFlags(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->optimizationFlags();
}

int QGraphicsView_RenderHints(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->renderHints();
}

int QGraphicsView_ResizeAnchor(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->resizeAnchor();
}

int QGraphicsView_RubberBandSelectionMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->rubberBandSelectionMode();
}

void QGraphicsView_SetAlignment(void* ptr, int alignment){
	static_cast<QGraphicsView*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsView_SetBackgroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsView*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetCacheMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setCacheMode(static_cast<QGraphicsView::CacheModeFlag>(mode));
}

void QGraphicsView_SetDragMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setDragMode(static_cast<QGraphicsView::DragMode>(mode));
}

void QGraphicsView_SetForegroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsView*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsView_SetInteractive(void* ptr, int allowed){
	static_cast<QGraphicsView*>(ptr)->setInteractive(allowed != 0);
}

void QGraphicsView_SetOptimizationFlags(void* ptr, int flags){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlags(static_cast<QGraphicsView::OptimizationFlag>(flags));
}

void QGraphicsView_SetRenderHints(void* ptr, int hints){
	static_cast<QGraphicsView*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints));
}

void QGraphicsView_SetResizeAnchor(void* ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setResizeAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetRubberBandSelectionMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setRubberBandSelectionMode(static_cast<Qt::ItemSelectionMode>(mode));
}

void QGraphicsView_SetSceneRect(void* ptr, void* rect){
	static_cast<QGraphicsView*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsView_SetTransformationAnchor(void* ptr, int anchor){
	static_cast<QGraphicsView*>(ptr)->setTransformationAnchor(static_cast<QGraphicsView::ViewportAnchor>(anchor));
}

void QGraphicsView_SetViewportUpdateMode(void* ptr, int mode){
	static_cast<QGraphicsView*>(ptr)->setViewportUpdateMode(static_cast<QGraphicsView::ViewportUpdateMode>(mode));
}

int QGraphicsView_TransformationAnchor(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->transformationAnchor();
}

int QGraphicsView_ViewportUpdateMode(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->viewportUpdateMode();
}

void* QGraphicsView_NewQGraphicsView2(void* scene, void* parent){
	return new QGraphicsView(static_cast<QGraphicsScene*>(scene), static_cast<QWidget*>(parent));
}

void* QGraphicsView_NewQGraphicsView(void* parent){
	return new QGraphicsView(static_cast<QWidget*>(parent));
}

void QGraphicsView_CenterOn3(void* ptr, void* item){
	static_cast<QGraphicsView*>(ptr)->centerOn(static_cast<QGraphicsItem*>(item));
}

void QGraphicsView_CenterOn(void* ptr, void* pos){
	static_cast<QGraphicsView*>(ptr)->centerOn(*static_cast<QPointF*>(pos));
}

void QGraphicsView_CenterOn2(void* ptr, double x, double y){
	static_cast<QGraphicsView*>(ptr)->centerOn(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QGraphicsView_EnsureVisible3(void* ptr, void* item, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(static_cast<QGraphicsItem*>(item), xmargin, ymargin);
}

void QGraphicsView_EnsureVisible(void* ptr, void* rect, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(*static_cast<QRectF*>(rect), xmargin, ymargin);
}

void QGraphicsView_EnsureVisible2(void* ptr, double x, double y, double w, double h, int xmargin, int ymargin){
	static_cast<QGraphicsView*>(ptr)->ensureVisible(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), xmargin, ymargin);
}

void QGraphicsView_FitInView3(void* ptr, void* item, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(static_cast<QGraphicsItem*>(item), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_FitInView(void* ptr, void* rect, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(*static_cast<QRectF*>(rect), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_FitInView2(void* ptr, double x, double y, double w, double h, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->fitInView(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void* QGraphicsView_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QGraphicsView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QGraphicsView_InvalidateScene(void* ptr, void* rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "invalidateScene", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

int QGraphicsView_IsTransformed(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->isTransformed();
}

void* QGraphicsView_ItemAt(void* ptr, void* pos){
	return static_cast<QGraphicsView*>(ptr)->itemAt(*static_cast<QPoint*>(pos));
}

void* QGraphicsView_ItemAt2(void* ptr, int x, int y){
	return static_cast<QGraphicsView*>(ptr)->itemAt(x, y);
}

void QGraphicsView_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode){
	static_cast<QGraphicsView*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRect*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsView_ResetCachedContent(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetCachedContent();
}

void QGraphicsView_ResetMatrix(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetMatrix();
}

void QGraphicsView_ResetTransform(void* ptr){
	static_cast<QGraphicsView*>(ptr)->resetTransform();
}

void QGraphicsView_Rotate(void* ptr, double angle){
	static_cast<QGraphicsView*>(ptr)->rotate(static_cast<qreal>(angle));
}

void QGraphicsView_Scale(void* ptr, double sx, double sy){
	static_cast<QGraphicsView*>(ptr)->scale(static_cast<qreal>(sx), static_cast<qreal>(sy));
}

void* QGraphicsView_Scene(void* ptr){
	return static_cast<QGraphicsView*>(ptr)->scene();
}

void QGraphicsView_SetOptimizationFlag(void* ptr, int flag, int enabled){
	static_cast<QGraphicsView*>(ptr)->setOptimizationFlag(static_cast<QGraphicsView::OptimizationFlag>(flag), enabled != 0);
}

void QGraphicsView_SetRenderHint(void* ptr, int hint, int enabled){
	static_cast<QGraphicsView*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), enabled != 0);
}

void QGraphicsView_SetScene(void* ptr, void* scene){
	static_cast<QGraphicsView*>(ptr)->setScene(static_cast<QGraphicsScene*>(scene));
}

void QGraphicsView_SetSceneRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsView*>(ptr)->setSceneRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsView_SetTransform(void* ptr, void* matrix, int combine){
	static_cast<QGraphicsView*>(ptr)->setTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QGraphicsView_Shear(void* ptr, double sh, double sv){
	static_cast<QGraphicsView*>(ptr)->shear(static_cast<qreal>(sh), static_cast<qreal>(sv));
}

void QGraphicsView_Translate(void* ptr, double dx, double dy){
	static_cast<QGraphicsView*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QGraphicsView_UpdateSceneRect(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsView*>(ptr), "updateSceneRect", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

void QGraphicsView_DestroyQGraphicsView(void* ptr){
	static_cast<QGraphicsView*>(ptr)->~QGraphicsView();
}

#include "qscrollarea.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScrollArea>
#include "_cgo_export.h"

class MyQScrollArea: public QScrollArea {
public:
};

int QScrollArea_Alignment(void* ptr){
	return static_cast<QScrollArea*>(ptr)->alignment();
}

void QScrollArea_SetAlignment(void* ptr, int v){
	static_cast<QScrollArea*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QScrollArea_SetWidget(void* ptr, void* widget){
	static_cast<QScrollArea*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void QScrollArea_SetWidgetResizable(void* ptr, int resizable){
	static_cast<QScrollArea*>(ptr)->setWidgetResizable(resizable != 0);
}

int QScrollArea_WidgetResizable(void* ptr){
	return static_cast<QScrollArea*>(ptr)->widgetResizable();
}

void* QScrollArea_NewQScrollArea(void* parent){
	return new QScrollArea(static_cast<QWidget*>(parent));
}

void QScrollArea_EnsureVisible(void* ptr, int x, int y, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureVisible(x, y, xmargin, ymargin);
}

void QScrollArea_EnsureWidgetVisible(void* ptr, void* childWidget, int xmargin, int ymargin){
	static_cast<QScrollArea*>(ptr)->ensureWidgetVisible(static_cast<QWidget*>(childWidget), xmargin, ymargin);
}

int QScrollArea_FocusNextPrevChild(void* ptr, int next){
	return static_cast<QScrollArea*>(ptr)->focusNextPrevChild(next != 0);
}

void* QScrollArea_TakeWidget(void* ptr){
	return static_cast<QScrollArea*>(ptr)->takeWidget();
}

void* QScrollArea_Widget(void* ptr){
	return static_cast<QScrollArea*>(ptr)->widget();
}

void QScrollArea_DestroyQScrollArea(void* ptr){
	static_cast<QScrollArea*>(ptr)->~QScrollArea();
}

#include "qsizepolicy.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizePolicy>
#include "_cgo_export.h"

class MyQSizePolicy: public QSizePolicy {
public:
};

void* QSizePolicy_NewQSizePolicy(){
	return new QSizePolicy();
}

void* QSizePolicy_NewQSizePolicy2(int horizontal, int vertical, int ty){
	return new QSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical), static_cast<QSizePolicy::ControlType>(ty));
}

int QSizePolicy_ControlType(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->controlType();
}

int QSizePolicy_ExpandingDirections(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->expandingDirections();
}

int QSizePolicy_HasHeightForWidth(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->hasHeightForWidth();
}

int QSizePolicy_HasWidthForHeight(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->hasWidthForHeight();
}

int QSizePolicy_HorizontalPolicy(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalPolicy();
}

int QSizePolicy_HorizontalStretch(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalStretch();
}

int QSizePolicy_RetainSizeWhenHidden(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->retainSizeWhenHidden();
}

void QSizePolicy_SetControlType(void* ptr, int ty){
	static_cast<QSizePolicy*>(ptr)->setControlType(static_cast<QSizePolicy::ControlType>(ty));
}

void QSizePolicy_SetHeightForWidth(void* ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setHeightForWidth(dependent != 0);
}

void QSizePolicy_SetHorizontalPolicy(void* ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setHorizontalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetHorizontalStretch(void* ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setHorizontalStretch(stretchFactor);
}

void QSizePolicy_SetRetainSizeWhenHidden(void* ptr, int retainSize){
	static_cast<QSizePolicy*>(ptr)->setRetainSizeWhenHidden(retainSize != 0);
}

void QSizePolicy_SetVerticalPolicy(void* ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setVerticalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetVerticalStretch(void* ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setVerticalStretch(stretchFactor);
}

void QSizePolicy_SetWidthForHeight(void* ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setWidthForHeight(dependent != 0);
}

void QSizePolicy_Transpose(void* ptr){
	static_cast<QSizePolicy*>(ptr)->transpose();
}

int QSizePolicy_VerticalPolicy(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalPolicy();
}

int QSizePolicy_VerticalStretch(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalStretch();
}

#include "qgraphicsitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsItem>
#include "_cgo_export.h"

#include "qheaderview.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QByteArray>
#include <QPoint>
#include <QUrl>
#include <QMetaObject>
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

#include "qstackedlayout.h"
#include <QRect>
#include <QString>
#include <QUrl>
#include <QLayout>
#include <QObject>
#include <QLayoutItem>
#include <QVariant>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QStack>
#include <QStackedLayout>
#include "_cgo_export.h"

class MyQStackedLayout: public QStackedLayout {
public:
void Signal_CurrentChanged(int index){callbackQStackedLayoutCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedLayoutWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedLayout_Count(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->count();
}

int QStackedLayout_CurrentIndex(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->currentIndex();
}

void QStackedLayout_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedLayout_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QStackedLayout*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QStackedLayout_SetStackingMode(void* ptr, int stackingMode){
	static_cast<QStackedLayout*>(ptr)->setStackingMode(static_cast<QStackedLayout::StackingMode>(stackingMode));
}

int QStackedLayout_StackingMode(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->stackingMode();
}

void* QStackedLayout_NewQStackedLayout(){
	return new QStackedLayout();
}

void* QStackedLayout_NewQStackedLayout3(void* parentLayout){
	return new QStackedLayout(static_cast<QLayout*>(parentLayout));
}

void* QStackedLayout_NewQStackedLayout2(void* parent){
	return new QStackedLayout(static_cast<QWidget*>(parent));
}

void QStackedLayout_AddItem(void* ptr, void* item){
	static_cast<QStackedLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

int QStackedLayout_AddWidget(void* ptr, void* widget){
	return static_cast<QStackedLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedLayout_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

void QStackedLayout_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::currentChanged), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_CurrentChanged));;
}

void* QStackedLayout_CurrentWidget(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->currentWidget();
}

int QStackedLayout_HasHeightForWidth(void* ptr){
	return static_cast<QStackedLayout*>(ptr)->hasHeightForWidth();
}

int QStackedLayout_HeightForWidth(void* ptr, int width){
	return static_cast<QStackedLayout*>(ptr)->heightForWidth(width);
}

int QStackedLayout_InsertWidget(void* ptr, int index, void* widget){
	return static_cast<QStackedLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

void* QStackedLayout_ItemAt(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->itemAt(index);
}

void QStackedLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QStackedLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QStackedLayout_TakeAt(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->takeAt(index);
}

void* QStackedLayout_Widget(void* ptr, int index){
	return static_cast<QStackedLayout*>(ptr)->widget(index);
}

void QStackedLayout_ConnectWidgetRemoved(void* ptr){
	QObject::connect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DisconnectWidgetRemoved(void* ptr){
	QObject::disconnect(static_cast<QStackedLayout*>(ptr), static_cast<void (QStackedLayout::*)(int)>(&QStackedLayout::widgetRemoved), static_cast<MyQStackedLayout*>(ptr), static_cast<void (MyQStackedLayout::*)(int)>(&MyQStackedLayout::Signal_WidgetRemoved));;
}

void QStackedLayout_DestroyQStackedLayout(void* ptr){
	static_cast<QStackedLayout*>(ptr)->~QStackedLayout();
}

#include "qgraphicslineitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QGraphicsItem>
#include <QLineF>
#include <QPainter>
#include <QStyleOptionGraphicsItem>
#include <QPointF>
#include <QStyleOption>
#include <QLine>
#include <QPoint>
#include <QPen>
#include <QStyle>
#include <QGraphicsLineItem>
#include "_cgo_export.h"

class MyQGraphicsLineItem: public QGraphicsLineItem {
public:
};

void* QGraphicsLineItem_NewQGraphicsLineItem(void* parent){
	return new QGraphicsLineItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsLineItem_NewQGraphicsLineItem2(void* line, void* parent){
	return new QGraphicsLineItem(*static_cast<QLineF*>(line), static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsLineItem_NewQGraphicsLineItem3(double x1, double y1, double x2, double y2, void* parent){
	return new QGraphicsLineItem(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsLineItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsLineItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsLineItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsLineItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsLineItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsLineItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsLineItem_SetLine(void* ptr, void* line){
	static_cast<QGraphicsLineItem*>(ptr)->setLine(*static_cast<QLineF*>(line));
}

void QGraphicsLineItem_SetLine2(void* ptr, double x1, double y1, double x2, double y2){
	static_cast<QGraphicsLineItem*>(ptr)->setLine(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2));
}

void QGraphicsLineItem_SetPen(void* ptr, void* pen){
	static_cast<QGraphicsLineItem*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

int QGraphicsLineItem_Type(void* ptr){
	return static_cast<QGraphicsLineItem*>(ptr)->type();
}

void QGraphicsLineItem_DestroyQGraphicsLineItem(void* ptr){
	static_cast<QGraphicsLineItem*>(ptr)->~QGraphicsLineItem();
}

#include "qgraphicsrotation.h"
#include <QMatrix4x4>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVector3D>
#include <QVector>
#include <QGraphicsRotation>
#include "_cgo_export.h"

class MyQGraphicsRotation: public QGraphicsRotation {
public:
void Signal_AngleChanged(){callbackQGraphicsRotationAngleChanged(this->objectName().toUtf8().data());};
void Signal_AxisChanged(){callbackQGraphicsRotationAxisChanged(this->objectName().toUtf8().data());};
void Signal_OriginChanged(){callbackQGraphicsRotationOriginChanged(this->objectName().toUtf8().data());};
};

double QGraphicsRotation_Angle(void* ptr){
	return static_cast<double>(static_cast<QGraphicsRotation*>(ptr)->angle());
}

void QGraphicsRotation_SetAngle(void* ptr, double v){
	static_cast<QGraphicsRotation*>(ptr)->setAngle(static_cast<qreal>(v));
}

void QGraphicsRotation_SetAxis2(void* ptr, int axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(static_cast<Qt::Axis>(axis));
}

void QGraphicsRotation_SetAxis(void* ptr, void* axis){
	static_cast<QGraphicsRotation*>(ptr)->setAxis(*static_cast<QVector3D*>(axis));
}

void QGraphicsRotation_SetOrigin(void* ptr, void* point){
	static_cast<QGraphicsRotation*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

void* QGraphicsRotation_NewQGraphicsRotation(void* parent){
	return new QGraphicsRotation(static_cast<QObject*>(parent));
}

void QGraphicsRotation_ConnectAngleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_DisconnectAngleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::angleChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AngleChanged));;
}

void QGraphicsRotation_ApplyTo(void* ptr, void* matrix){
	static_cast<QGraphicsRotation*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsRotation_ConnectAxisChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_DisconnectAxisChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::axisChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_AxisChanged));;
}

void QGraphicsRotation_ConnectOriginChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DisconnectOriginChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsRotation*>(ptr), static_cast<void (QGraphicsRotation::*)()>(&QGraphicsRotation::originChanged), static_cast<MyQGraphicsRotation*>(ptr), static_cast<void (MyQGraphicsRotation::*)()>(&MyQGraphicsRotation::Signal_OriginChanged));;
}

void QGraphicsRotation_DestroyQGraphicsRotation(void* ptr){
	static_cast<QGraphicsRotation*>(ptr)->~QGraphicsRotation();
}

#include "qtreewidgetitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBrush>
#include <QSize>
#include <QFont>
#include <QTreeWidget>
#include <QIcon>
#include <QDataStream>
#include <QTreeWidgetItem>
#include "_cgo_export.h"

class MyQTreeWidgetItem: public QTreeWidgetItem {
public:
};

void* QTreeWidgetItem_NewQTreeWidgetItem5(void* parent, void* preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem4(void* parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem3(void* parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidget*>(parent), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem8(void* parent, void* preceding, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), static_cast<QTreeWidgetItem*>(preceding), ty);
}

int QTreeWidgetItem_Flags(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->flags();
}

void QTreeWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QTreeWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QTreeWidgetItem_NewQTreeWidgetItem7(void* parent, char* strin, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem6(void* parent, int ty){
	return new QTreeWidgetItem(static_cast<QTreeWidgetItem*>(parent), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem2(char* strin, int ty){
	return new QTreeWidgetItem(QString(strin).split("|", QString::SkipEmptyParts), ty);
}

void* QTreeWidgetItem_NewQTreeWidgetItem9(void* other){
	return new QTreeWidgetItem(*static_cast<QTreeWidgetItem*>(other));
}

void* QTreeWidgetItem_NewQTreeWidgetItem(int ty){
	return new QTreeWidgetItem(ty);
}

void QTreeWidgetItem_AddChild(void* ptr, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->addChild(static_cast<QTreeWidgetItem*>(child));
}

void* QTreeWidgetItem_Background(void* ptr, int column){
	return new QBrush(static_cast<QTreeWidgetItem*>(ptr)->background(column));
}

int QTreeWidgetItem_CheckState(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->checkState(column);
}

void* QTreeWidgetItem_Child(void* ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->child(index);
}

int QTreeWidgetItem_ChildCount(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childCount();
}

int QTreeWidgetItem_ChildIndicatorPolicy(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->childIndicatorPolicy();
}

int QTreeWidgetItem_ColumnCount(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->columnCount();
}

void* QTreeWidgetItem_Data(void* ptr, int column, int role){
	return new QVariant(static_cast<QTreeWidgetItem*>(ptr)->data(column, role));
}

void* QTreeWidgetItem_Clone(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->clone();
}

void* QTreeWidgetItem_Foreground(void* ptr, int column){
	return new QBrush(static_cast<QTreeWidgetItem*>(ptr)->foreground(column));
}

int QTreeWidgetItem_IndexOfChild(void* ptr, void* child){
	return static_cast<QTreeWidgetItem*>(ptr)->indexOfChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_InsertChild(void* ptr, int index, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->insertChild(index, static_cast<QTreeWidgetItem*>(child));
}

int QTreeWidgetItem_IsDisabled(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isDisabled();
}

int QTreeWidgetItem_IsExpanded(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isExpanded();
}

int QTreeWidgetItem_IsFirstColumnSpanned(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isFirstColumnSpanned();
}

int QTreeWidgetItem_IsHidden(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isHidden();
}

int QTreeWidgetItem_IsSelected(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->isSelected();
}

void* QTreeWidgetItem_Parent(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->parent();
}

void QTreeWidgetItem_Read(void* ptr, void* in){
	static_cast<QTreeWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QTreeWidgetItem_RemoveChild(void* ptr, void* child){
	static_cast<QTreeWidgetItem*>(ptr)->removeChild(static_cast<QTreeWidgetItem*>(child));
}

void QTreeWidgetItem_SetBackground(void* ptr, int column, void* brush){
	static_cast<QTreeWidgetItem*>(ptr)->setBackground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetCheckState(void* ptr, int column, int state){
	static_cast<QTreeWidgetItem*>(ptr)->setCheckState(column, static_cast<Qt::CheckState>(state));
}

void QTreeWidgetItem_SetChildIndicatorPolicy(void* ptr, int policy){
	static_cast<QTreeWidgetItem*>(ptr)->setChildIndicatorPolicy(static_cast<QTreeWidgetItem::ChildIndicatorPolicy>(policy));
}

void QTreeWidgetItem_SetData(void* ptr, int column, int role, void* value){
	static_cast<QTreeWidgetItem*>(ptr)->setData(column, role, *static_cast<QVariant*>(value));
}

void QTreeWidgetItem_SetDisabled(void* ptr, int disabled){
	static_cast<QTreeWidgetItem*>(ptr)->setDisabled(disabled != 0);
}

void QTreeWidgetItem_SetExpanded(void* ptr, int expand){
	static_cast<QTreeWidgetItem*>(ptr)->setExpanded(expand != 0);
}

void QTreeWidgetItem_SetFirstColumnSpanned(void* ptr, int span){
	static_cast<QTreeWidgetItem*>(ptr)->setFirstColumnSpanned(span != 0);
}

void QTreeWidgetItem_SetFont(void* ptr, int column, void* font){
	static_cast<QTreeWidgetItem*>(ptr)->setFont(column, *static_cast<QFont*>(font));
}

void QTreeWidgetItem_SetForeground(void* ptr, int column, void* brush){
	static_cast<QTreeWidgetItem*>(ptr)->setForeground(column, *static_cast<QBrush*>(brush));
}

void QTreeWidgetItem_SetHidden(void* ptr, int hide){
	static_cast<QTreeWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QTreeWidgetItem_SetIcon(void* ptr, int column, void* icon){
	static_cast<QTreeWidgetItem*>(ptr)->setIcon(column, *static_cast<QIcon*>(icon));
}

void QTreeWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QTreeWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTreeWidgetItem_SetSizeHint(void* ptr, int column, void* size){
	static_cast<QTreeWidgetItem*>(ptr)->setSizeHint(column, *static_cast<QSize*>(size));
}

void QTreeWidgetItem_SetStatusTip(void* ptr, int column, char* statusTip){
	static_cast<QTreeWidgetItem*>(ptr)->setStatusTip(column, QString(statusTip));
}

void QTreeWidgetItem_SetText(void* ptr, int column, char* text){
	static_cast<QTreeWidgetItem*>(ptr)->setText(column, QString(text));
}

void QTreeWidgetItem_SetTextAlignment(void* ptr, int column, int alignment){
	static_cast<QTreeWidgetItem*>(ptr)->setTextAlignment(column, alignment);
}

void QTreeWidgetItem_SetToolTip(void* ptr, int column, char* toolTip){
	static_cast<QTreeWidgetItem*>(ptr)->setToolTip(column, QString(toolTip));
}

void QTreeWidgetItem_SetWhatsThis(void* ptr, int column, char* whatsThis){
	static_cast<QTreeWidgetItem*>(ptr)->setWhatsThis(column, QString(whatsThis));
}

void QTreeWidgetItem_SortChildren(void* ptr, int column, int order){
	static_cast<QTreeWidgetItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QTreeWidgetItem_StatusTip(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->statusTip(column).toUtf8().data();
}

void* QTreeWidgetItem_TakeChild(void* ptr, int index){
	return static_cast<QTreeWidgetItem*>(ptr)->takeChild(index);
}

char* QTreeWidgetItem_Text(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->text(column).toUtf8().data();
}

int QTreeWidgetItem_TextAlignment(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->textAlignment(column);
}

char* QTreeWidgetItem_ToolTip(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->toolTip(column).toUtf8().data();
}

void* QTreeWidgetItem_TreeWidget(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->treeWidget();
}

int QTreeWidgetItem_Type(void* ptr){
	return static_cast<QTreeWidgetItem*>(ptr)->type();
}

char* QTreeWidgetItem_WhatsThis(void* ptr, int column){
	return static_cast<QTreeWidgetItem*>(ptr)->whatsThis(column).toUtf8().data();
}

void QTreeWidgetItem_Write(void* ptr, void* out){
	static_cast<QTreeWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTreeWidgetItem_DestroyQTreeWidgetItem(void* ptr){
	static_cast<QTreeWidgetItem*>(ptr)->~QTreeWidgetItem();
}

#include "qabstractitemview.h"
#include <QAbstractItemDelegate>
#include <QItemSelectionModel>
#include <QAbstractItemModel>
#include <QItemSelection>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QPoint>
#include <QSize>
#include <QAbstractItemView>
#include "_cgo_export.h"

class MyQAbstractItemView: public QAbstractItemView {
public:
void Signal_Activated(const QModelIndex & index){callbackQAbstractItemViewActivated(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Clicked(const QModelIndex & index){callbackQAbstractItemViewClicked(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_DoubleClicked(const QModelIndex & index){callbackQAbstractItemViewDoubleClicked(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Entered(const QModelIndex & index){callbackQAbstractItemViewEntered(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Pressed(const QModelIndex & index){callbackQAbstractItemViewPressed(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_ViewportEntered(){callbackQAbstractItemViewViewportEntered(this->objectName().toUtf8().data());};
};

int QAbstractItemView_AlternatingRowColors(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->alternatingRowColors();
}

int QAbstractItemView_AutoScrollMargin(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->autoScrollMargin();
}

int QAbstractItemView_DefaultDropAction(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->defaultDropAction();
}

int QAbstractItemView_DragDropMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropMode();
}

int QAbstractItemView_DragDropOverwriteMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragDropOverwriteMode();
}

int QAbstractItemView_DragEnabled(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->dragEnabled();
}

int QAbstractItemView_EditTriggers(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->editTriggers();
}

int QAbstractItemView_HasAutoScroll(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->hasAutoScroll();
}

int QAbstractItemView_HorizontalScrollMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->horizontalScrollMode();
}

int QAbstractItemView_SelectionBehavior(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionBehavior();
}

int QAbstractItemView_SelectionMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionMode();
}

void QAbstractItemView_SetAlternatingRowColors(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAlternatingRowColors(enable != 0);
}

void QAbstractItemView_SetAutoScroll(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setAutoScroll(enable != 0);
}

void QAbstractItemView_SetAutoScrollMargin(void* ptr, int margin){
	static_cast<QAbstractItemView*>(ptr)->setAutoScrollMargin(margin);
}

void QAbstractItemView_SetDefaultDropAction(void* ptr, int dropAction){
	static_cast<QAbstractItemView*>(ptr)->setDefaultDropAction(static_cast<Qt::DropAction>(dropAction));
}

void QAbstractItemView_SetDragDropMode(void* ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setDragDropMode(static_cast<QAbstractItemView::DragDropMode>(behavior));
}

void QAbstractItemView_SetDragDropOverwriteMode(void* ptr, int overwrite){
	static_cast<QAbstractItemView*>(ptr)->setDragDropOverwriteMode(overwrite != 0);
}

void QAbstractItemView_SetDragEnabled(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDragEnabled(enable != 0);
}

void QAbstractItemView_SetDropIndicatorShown(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setDropIndicatorShown(enable != 0);
}

void QAbstractItemView_SetEditTriggers(void* ptr, int triggers){
	static_cast<QAbstractItemView*>(ptr)->setEditTriggers(static_cast<QAbstractItemView::EditTrigger>(triggers));
}

void QAbstractItemView_SetHorizontalScrollMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setHorizontalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

void QAbstractItemView_SetIconSize(void* ptr, void* size){
	static_cast<QAbstractItemView*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QAbstractItemView_SetSelectionBehavior(void* ptr, int behavior){
	static_cast<QAbstractItemView*>(ptr)->setSelectionBehavior(static_cast<QAbstractItemView::SelectionBehavior>(behavior));
}

void QAbstractItemView_SetSelectionMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setSelectionMode(static_cast<QAbstractItemView::SelectionMode>(mode));
}

void QAbstractItemView_SetTabKeyNavigation(void* ptr, int enable){
	static_cast<QAbstractItemView*>(ptr)->setTabKeyNavigation(enable != 0);
}

void QAbstractItemView_SetTextElideMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setTextElideMode(static_cast<Qt::TextElideMode>(mode));
}

void QAbstractItemView_SetVerticalScrollMode(void* ptr, int mode){
	static_cast<QAbstractItemView*>(ptr)->setVerticalScrollMode(static_cast<QAbstractItemView::ScrollMode>(mode));
}

int QAbstractItemView_ShowDropIndicator(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->showDropIndicator();
}

int QAbstractItemView_TabKeyNavigation(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->tabKeyNavigation();
}

int QAbstractItemView_TextElideMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->textElideMode();
}

int QAbstractItemView_VerticalScrollMode(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->verticalScrollMode();
}

void QAbstractItemView_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::activated), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Activated));;
}

void QAbstractItemView_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "clearSelection");
}

void QAbstractItemView_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::clicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Clicked));;
}

void QAbstractItemView_ClosePersistentEditor(void* ptr, void* index){
	static_cast<QAbstractItemView*>(ptr)->closePersistentEditor(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_CurrentIndex(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->currentIndex().internalPointer();
}

void QAbstractItemView_ConnectDoubleClicked(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_DisconnectDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::doubleClicked), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_DoubleClicked));;
}

void QAbstractItemView_Edit(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "edit", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectEntered(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

void QAbstractItemView_DisconnectEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::entered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Entered));;
}

void* QAbstractItemView_IndexAt(void* ptr, void* point){
	return static_cast<QAbstractItemView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QAbstractItemView_IndexWidget(void* ptr, void* index){
	return static_cast<QAbstractItemView*>(ptr)->indexWidget(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QAbstractItemView*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void* QAbstractItemView_ItemDelegate(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate();
}

void* QAbstractItemView_ItemDelegate2(void* ptr, void* index){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegate(*static_cast<QModelIndex*>(index));
}

void* QAbstractItemView_ItemDelegateForColumn(void* ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForColumn(column);
}

void* QAbstractItemView_ItemDelegateForRow(void* ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->itemDelegateForRow(row);
}

void QAbstractItemView_KeyboardSearch(void* ptr, char* search){
	static_cast<QAbstractItemView*>(ptr)->keyboardSearch(QString(search));
}

void* QAbstractItemView_Model(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->model();
}

void QAbstractItemView_OpenPersistentEditor(void* ptr, void* index){
	static_cast<QAbstractItemView*>(ptr)->openPersistentEditor(*static_cast<QModelIndex*>(index));
}

void QAbstractItemView_ConnectPressed(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_DisconnectPressed(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)(const QModelIndex &)>(&QAbstractItemView::pressed), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)(const QModelIndex &)>(&MyQAbstractItemView::Signal_Pressed));;
}

void QAbstractItemView_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "reset");
}

void* QAbstractItemView_RootIndex(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->rootIndex().internalPointer();
}

void QAbstractItemView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QAbstractItemView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QAbstractItemView_ScrollToBottom(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToBottom");
}

void QAbstractItemView_ScrollToTop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "scrollToTop");
}

void QAbstractItemView_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "selectAll");
}

void* QAbstractItemView_SelectionModel(void* ptr){
	return static_cast<QAbstractItemView*>(ptr)->selectionModel();
}

void QAbstractItemView_SetCurrentIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetIndexWidget(void* ptr, void* index, void* widget){
	static_cast<QAbstractItemView*>(ptr)->setIndexWidget(*static_cast<QModelIndex*>(index), static_cast<QWidget*>(widget));
}

void QAbstractItemView_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForColumn(void* ptr, int column, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForColumn(column, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetItemDelegateForRow(void* ptr, int row, void* delegate){
	static_cast<QAbstractItemView*>(ptr)->setItemDelegateForRow(row, static_cast<QAbstractItemDelegate*>(delegate));
}

void QAbstractItemView_SetModel(void* ptr, void* model){
	static_cast<QAbstractItemView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QAbstractItemView_SetRootIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "setRootIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QAbstractItemView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QAbstractItemView_SizeHintForColumn(void* ptr, int column){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForColumn(column);
}

int QAbstractItemView_SizeHintForRow(void* ptr, int row){
	return static_cast<QAbstractItemView*>(ptr)->sizeHintForRow(row);
}

void QAbstractItemView_Update(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QAbstractItemView*>(ptr), "update", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemView_ConnectViewportEntered(void* ptr){
	QObject::connect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DisconnectViewportEntered(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemView*>(ptr), static_cast<void (QAbstractItemView::*)()>(&QAbstractItemView::viewportEntered), static_cast<MyQAbstractItemView*>(ptr), static_cast<void (MyQAbstractItemView::*)()>(&MyQAbstractItemView::Signal_ViewportEntered));;
}

void QAbstractItemView_DestroyQAbstractItemView(void* ptr){
	static_cast<QAbstractItemView*>(ptr)->~QAbstractItemView();
}

#include "qtreeview.h"
#include <QMetaObject>
#include <QAbstractItemModel>
#include <QPoint>
#include <QVariant>
#include <QModelIndex>
#include <QItemSelection>
#include <QObject>
#include <QAbstractItemView>
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QHeaderView>
#include <QItemSelectionModel>
#include <QTreeView>
#include "_cgo_export.h"

class MyQTreeView: public QTreeView {
public:
void Signal_Collapsed(const QModelIndex & index){callbackQTreeViewCollapsed(this->objectName().toUtf8().data(), index.internalPointer());};
void Signal_Expanded(const QModelIndex & index){callbackQTreeViewExpanded(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QTreeView_AllColumnsShowFocus(void* ptr){
	return static_cast<QTreeView*>(ptr)->allColumnsShowFocus();
}

int QTreeView_AutoExpandDelay(void* ptr){
	return static_cast<QTreeView*>(ptr)->autoExpandDelay();
}

void QTreeView_Collapse(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapse", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QTreeView_Expand(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expand", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

int QTreeView_ExpandsOnDoubleClick(void* ptr){
	return static_cast<QTreeView*>(ptr)->expandsOnDoubleClick();
}

int QTreeView_Indentation(void* ptr){
	return static_cast<QTreeView*>(ptr)->indentation();
}

int QTreeView_IsAnimated(void* ptr){
	return static_cast<QTreeView*>(ptr)->isAnimated();
}

int QTreeView_IsExpanded(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->isExpanded(*static_cast<QModelIndex*>(index));
}

int QTreeView_IsHeaderHidden(void* ptr){
	return static_cast<QTreeView*>(ptr)->isHeaderHidden();
}

int QTreeView_IsSortingEnabled(void* ptr){
	return static_cast<QTreeView*>(ptr)->isSortingEnabled();
}

int QTreeView_ItemsExpandable(void* ptr){
	return static_cast<QTreeView*>(ptr)->itemsExpandable();
}

void QTreeView_ResetIndentation(void* ptr){
	static_cast<QTreeView*>(ptr)->resetIndentation();
}

int QTreeView_RootIsDecorated(void* ptr){
	return static_cast<QTreeView*>(ptr)->rootIsDecorated();
}

void QTreeView_SetAllColumnsShowFocus(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAllColumnsShowFocus(enable != 0);
}

void QTreeView_SetAnimated(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setAnimated(enable != 0);
}

void QTreeView_SetAutoExpandDelay(void* ptr, int delay){
	static_cast<QTreeView*>(ptr)->setAutoExpandDelay(delay);
}

void QTreeView_SetExpandsOnDoubleClick(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setExpandsOnDoubleClick(enable != 0);
}

void QTreeView_SetHeaderHidden(void* ptr, int hide){
	static_cast<QTreeView*>(ptr)->setHeaderHidden(hide != 0);
}

void QTreeView_SetIndentation(void* ptr, int i){
	static_cast<QTreeView*>(ptr)->setIndentation(i);
}

void QTreeView_SetItemsExpandable(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setItemsExpandable(enable != 0);
}

void QTreeView_SetRootIsDecorated(void* ptr, int show){
	static_cast<QTreeView*>(ptr)->setRootIsDecorated(show != 0);
}

void QTreeView_SetSortingEnabled(void* ptr, int enable){
	static_cast<QTreeView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTreeView_SetUniformRowHeights(void* ptr, int uniform){
	static_cast<QTreeView*>(ptr)->setUniformRowHeights(uniform != 0);
}

void QTreeView_SetWordWrap(void* ptr, int on){
	static_cast<QTreeView*>(ptr)->setWordWrap(on != 0);
}

int QTreeView_UniformRowHeights(void* ptr){
	return static_cast<QTreeView*>(ptr)->uniformRowHeights();
}

int QTreeView_WordWrap(void* ptr){
	return static_cast<QTreeView*>(ptr)->wordWrap();
}

void* QTreeView_NewQTreeView(void* parent){
	return new QTreeView(static_cast<QWidget*>(parent));
}

void QTreeView_CollapseAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "collapseAll");
}

void QTreeView_ConnectCollapsed(void* ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

void QTreeView_DisconnectCollapsed(void* ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::collapsed), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Collapsed));;
}

int QTreeView_ColumnAt(void* ptr, int x){
	return static_cast<QTreeView*>(ptr)->columnAt(x);
}

int QTreeView_ColumnViewportPosition(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnViewportPosition(column);
}

int QTreeView_ColumnWidth(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->columnWidth(column);
}

void QTreeView_ExpandAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandAll");
}

void QTreeView_ExpandToDepth(void* ptr, int depth){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "expandToDepth", Q_ARG(int, depth));
}

void QTreeView_ConnectExpanded(void* ptr){
	QObject::connect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

void QTreeView_DisconnectExpanded(void* ptr){
	QObject::disconnect(static_cast<QTreeView*>(ptr), static_cast<void (QTreeView::*)(const QModelIndex &)>(&QTreeView::expanded), static_cast<MyQTreeView*>(ptr), static_cast<void (MyQTreeView::*)(const QModelIndex &)>(&MyQTreeView::Signal_Expanded));;
}

void* QTreeView_Header(void* ptr){
	return static_cast<QTreeView*>(ptr)->header();
}

void QTreeView_HideColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "hideColumn", Q_ARG(int, column));
}

void* QTreeView_IndexAbove(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->indexAbove(*static_cast<QModelIndex*>(index)).internalPointer();
}

void* QTreeView_IndexAt(void* ptr, void* point){
	return static_cast<QTreeView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QTreeView_IndexBelow(void* ptr, void* index){
	return static_cast<QTreeView*>(ptr)->indexBelow(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QTreeView_IsColumnHidden(void* ptr, int column){
	return static_cast<QTreeView*>(ptr)->isColumnHidden(column);
}

int QTreeView_IsFirstColumnSpanned(void* ptr, int row, void* parent){
	return static_cast<QTreeView*>(ptr)->isFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent));
}

int QTreeView_IsRowHidden(void* ptr, int row, void* parent){
	return static_cast<QTreeView*>(ptr)->isRowHidden(row, *static_cast<QModelIndex*>(parent));
}

void QTreeView_KeyboardSearch(void* ptr, char* search){
	static_cast<QTreeView*>(ptr)->keyboardSearch(QString(search));
}

void QTreeView_Reset(void* ptr){
	static_cast<QTreeView*>(ptr)->reset();
}

void QTreeView_ResizeColumnToContents(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTreeView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QTreeView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QTreeView_SelectAll(void* ptr){
	static_cast<QTreeView*>(ptr)->selectAll();
}

void QTreeView_SetColumnHidden(void* ptr, int column, int hide){
	static_cast<QTreeView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTreeView_SetColumnWidth(void* ptr, int column, int width){
	static_cast<QTreeView*>(ptr)->setColumnWidth(column, width);
}

void QTreeView_SetExpanded(void* ptr, void* index, int expanded){
	static_cast<QTreeView*>(ptr)->setExpanded(*static_cast<QModelIndex*>(index), expanded != 0);
}

void QTreeView_SetFirstColumnSpanned(void* ptr, int row, void* parent, int span){
	static_cast<QTreeView*>(ptr)->setFirstColumnSpanned(row, *static_cast<QModelIndex*>(parent), span != 0);
}

void QTreeView_SetHeader(void* ptr, void* header){
	static_cast<QTreeView*>(ptr)->setHeader(static_cast<QHeaderView*>(header));
}

void QTreeView_SetModel(void* ptr, void* model){
	static_cast<QTreeView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTreeView_SetRootIndex(void* ptr, void* index){
	static_cast<QTreeView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTreeView_SetRowHidden(void* ptr, int row, void* parent, int hide){
	static_cast<QTreeView*>(ptr)->setRowHidden(row, *static_cast<QModelIndex*>(parent), hide != 0);
}

void QTreeView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTreeView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTreeView_SetTreePosition(void* ptr, int index){
	static_cast<QTreeView*>(ptr)->setTreePosition(index);
}

void QTreeView_ShowColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTreeView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTreeView_SortByColumn(void* ptr, int column, int order){
	static_cast<QTreeView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

int QTreeView_TreePosition(void* ptr){
	return static_cast<QTreeView*>(ptr)->treePosition();
}

void QTreeView_DestroyQTreeView(void* ptr){
	static_cast<QTreeView*>(ptr)->~QTreeView();
}

#include "qwidgetitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QWidget>
#include <QWidgetItem>
#include "_cgo_export.h"

class MyQWidgetItem: public QWidgetItem {
public:
};

void* QWidgetItem_NewQWidgetItem(void* widget){
	return new QWidgetItem(static_cast<QWidget*>(widget));
}

int QWidgetItem_ControlTypes(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->controlTypes();
}

int QWidgetItem_ExpandingDirections(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->expandingDirections();
}

int QWidgetItem_HasHeightForWidth(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->hasHeightForWidth();
}

int QWidgetItem_HeightForWidth(void* ptr, int w){
	return static_cast<QWidgetItem*>(ptr)->heightForWidth(w);
}

int QWidgetItem_IsEmpty(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->isEmpty();
}

void QWidgetItem_SetGeometry(void* ptr, void* rect){
	static_cast<QWidgetItem*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void* QWidgetItem_Widget(void* ptr){
	return static_cast<QWidgetItem*>(ptr)->widget();
}

void QWidgetItem_DestroyQWidgetItem(void* ptr){
	static_cast<QWidgetItem*>(ptr)->~QWidgetItem();
}

#include "qstyleditemdelegate.h"
#include <QPainter>
#include <QStyle>
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QLocale>
#include <QStyleOptionViewItem>
#include <QAbstractItemModel>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QItemEditorFactory>
#include <QStyleOption>
#include <QStyledItemDelegate>
#include "_cgo_export.h"

class MyQStyledItemDelegate: public QStyledItemDelegate {
public:
};

void* QStyledItemDelegate_NewQStyledItemDelegate(void* parent){
	return new QStyledItemDelegate(static_cast<QObject*>(parent));
}

void* QStyledItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QStyledItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char* QStyledItemDelegate_DisplayText(void* ptr, void* value, void* locale){
	return static_cast<QStyledItemDelegate*>(ptr)->displayText(*static_cast<QVariant*>(value), *static_cast<QLocale*>(locale)).toUtf8().data();
}

void* QStyledItemDelegate_ItemEditorFactory(void* ptr){
	return static_cast<QStyledItemDelegate*>(ptr)->itemEditorFactory();
}

void QStyledItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetItemEditorFactory(void* ptr, void* factory){
	static_cast<QStyledItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QStyledItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_DestroyQStyledItemDelegate(void* ptr){
	static_cast<QStyledItemDelegate*>(ptr)->~QStyledItemDelegate();
}

#include "qvboxlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QVBoxLayout>
#include "_cgo_export.h"

class MyQVBoxLayout: public QVBoxLayout {
public:
};

void* QVBoxLayout_NewQVBoxLayout(){
	return new QVBoxLayout();
}

void* QVBoxLayout_NewQVBoxLayout2(void* parent){
	return new QVBoxLayout(static_cast<QWidget*>(parent));
}

void QVBoxLayout_DestroyQVBoxLayout(void* ptr){
	static_cast<QVBoxLayout*>(ptr)->~QVBoxLayout();
}

#include "qstylehintreturnvariant.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleHintReturn>
#include <QStyle>
#include <QString>
#include <QStyleHintReturnVariant>
#include "_cgo_export.h"

class MyQStyleHintReturnVariant: public QStyleHintReturnVariant {
public:
};

void* QStyleHintReturnVariant_NewQStyleHintReturnVariant(){
	return new QStyleHintReturnVariant();
}

void QStyleHintReturnVariant_DestroyQStyleHintReturnVariant(void* ptr){
	static_cast<QStyleHintReturnVariant*>(ptr)->~QStyleHintReturnVariant();
}

#include "qgraphicsanchorlayout.h"
#include <QRectF>
#include <QGraphicsAnchor>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QGraphicsAnchorLayout>
#include "_cgo_export.h"

class MyQGraphicsAnchorLayout: public QGraphicsAnchorLayout {
public:
};

void* QGraphicsAnchorLayout_NewQGraphicsAnchorLayout(void* parent){
	return new QGraphicsAnchorLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void* QGraphicsAnchorLayout_AddAnchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

void QGraphicsAnchorLayout_AddAnchors(void* ptr, void* firstItem, void* secondItem, int orientations){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Orientation>(orientations));
}

void QGraphicsAnchorLayout_AddCornerAnchors(void* ptr, void* firstItem, int firstCorner, void* secondItem, int secondCorner){
	static_cast<QGraphicsAnchorLayout*>(ptr)->addCornerAnchors(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::Corner>(firstCorner), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::Corner>(secondCorner));
}

void* QGraphicsAnchorLayout_Anchor(void* ptr, void* firstItem, int firstEdge, void* secondItem, int secondEdge){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->anchor(static_cast<QGraphicsLayoutItem*>(firstItem), static_cast<Qt::AnchorPoint>(firstEdge), static_cast<QGraphicsLayoutItem*>(secondItem), static_cast<Qt::AnchorPoint>(secondEdge));
}

int QGraphicsAnchorLayout_Count(void* ptr){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->count();
}

double QGraphicsAnchorLayout_HorizontalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchorLayout*>(ptr)->horizontalSpacing());
}

void QGraphicsAnchorLayout_Invalidate(void* ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->invalidate();
}

void* QGraphicsAnchorLayout_ItemAt(void* ptr, int index){
	return static_cast<QGraphicsAnchorLayout*>(ptr)->itemAt(index);
}

void QGraphicsAnchorLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsAnchorLayout*>(ptr)->removeAt(index);
}

void QGraphicsAnchorLayout_SetGeometry(void* ptr, void* geom){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(geom));
}

void QGraphicsAnchorLayout_SetHorizontalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setHorizontalSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsAnchorLayout_SetVerticalSpacing(void* ptr, double spacing){
	static_cast<QGraphicsAnchorLayout*>(ptr)->setVerticalSpacing(static_cast<qreal>(spacing));
}

double QGraphicsAnchorLayout_VerticalSpacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsAnchorLayout*>(ptr)->verticalSpacing());
}

void QGraphicsAnchorLayout_DestroyQGraphicsAnchorLayout(void* ptr){
	static_cast<QGraphicsAnchorLayout*>(ptr)->~QGraphicsAnchorLayout();
}

#include "qundoview.h"
#include <QUrl>
#include <QIcon>
#include <QUndoStack>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QUndoGroup>
#include <QWidget>
#include <QUndoView>
#include "_cgo_export.h"

class MyQUndoView: public QUndoView {
public:
};

char* QUndoView_EmptyLabel(void* ptr){
	return static_cast<QUndoView*>(ptr)->emptyLabel().toUtf8().data();
}

void QUndoView_SetCleanIcon(void* ptr, void* icon){
	static_cast<QUndoView*>(ptr)->setCleanIcon(*static_cast<QIcon*>(icon));
}

void QUndoView_SetEmptyLabel(void* ptr, char* label){
	static_cast<QUndoView*>(ptr)->setEmptyLabel(QString(label));
}

void* QUndoView_NewQUndoView3(void* group, void* parent){
	return new QUndoView(static_cast<QUndoGroup*>(group), static_cast<QWidget*>(parent));
}

void* QUndoView_NewQUndoView2(void* stack, void* parent){
	return new QUndoView(static_cast<QUndoStack*>(stack), static_cast<QWidget*>(parent));
}

void* QUndoView_NewQUndoView(void* parent){
	return new QUndoView(static_cast<QWidget*>(parent));
}

void* QUndoView_Group(void* ptr){
	return static_cast<QUndoView*>(ptr)->group();
}

void QUndoView_SetGroup(void* ptr, void* group){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setGroup", Q_ARG(QUndoGroup*, static_cast<QUndoGroup*>(group)));
}

void QUndoView_SetStack(void* ptr, void* stack){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setStack", Q_ARG(QUndoStack*, static_cast<QUndoStack*>(stack)));
}

void* QUndoView_Stack(void* ptr){
	return static_cast<QUndoView*>(ptr)->stack();
}

void QUndoView_DestroyQUndoView(void* ptr){
	static_cast<QUndoView*>(ptr)->~QUndoView();
}

#include "qabstractspinbox.h"
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QMetaObject>
#include <QAbstractSpinBox>
#include "_cgo_export.h"

class MyQAbstractSpinBox: public QAbstractSpinBox {
public:
void Signal_EditingFinished(){callbackQAbstractSpinBoxEditingFinished(this->objectName().toUtf8().data());};
};

int QAbstractSpinBox_Alignment(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->alignment();
}

int QAbstractSpinBox_ButtonSymbols(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->buttonSymbols();
}

int QAbstractSpinBox_CorrectionMode(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->correctionMode();
}

int QAbstractSpinBox_HasAcceptableInput(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasAcceptableInput();
}

int QAbstractSpinBox_HasFrame(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasFrame();
}

int QAbstractSpinBox_IsAccelerated(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isAccelerated();
}

int QAbstractSpinBox_IsGroupSeparatorShown(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isGroupSeparatorShown();
}

int QAbstractSpinBox_IsReadOnly(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isReadOnly();
}

int QAbstractSpinBox_KeyboardTracking(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->keyboardTracking();
}

void QAbstractSpinBox_SetAccelerated(void* ptr, int on){
	static_cast<QAbstractSpinBox*>(ptr)->setAccelerated(on != 0);
}

void QAbstractSpinBox_SetAlignment(void* ptr, int flag){
	static_cast<QAbstractSpinBox*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QAbstractSpinBox_SetButtonSymbols(void* ptr, int bs){
	static_cast<QAbstractSpinBox*>(ptr)->setButtonSymbols(static_cast<QAbstractSpinBox::ButtonSymbols>(bs));
}

void QAbstractSpinBox_SetCorrectionMode(void* ptr, int cm){
	static_cast<QAbstractSpinBox*>(ptr)->setCorrectionMode(static_cast<QAbstractSpinBox::CorrectionMode>(cm));
}

void QAbstractSpinBox_SetFrame(void* ptr, int v){
	static_cast<QAbstractSpinBox*>(ptr)->setFrame(v != 0);
}

void QAbstractSpinBox_SetGroupSeparatorShown(void* ptr, int shown){
	static_cast<QAbstractSpinBox*>(ptr)->setGroupSeparatorShown(shown != 0);
}

void QAbstractSpinBox_SetKeyboardTracking(void* ptr, int kt){
	static_cast<QAbstractSpinBox*>(ptr)->setKeyboardTracking(kt != 0);
}

void QAbstractSpinBox_SetReadOnly(void* ptr, int r){
	static_cast<QAbstractSpinBox*>(ptr)->setReadOnly(r != 0);
}

void QAbstractSpinBox_SetSpecialValueText(void* ptr, char* txt){
	static_cast<QAbstractSpinBox*>(ptr)->setSpecialValueText(QString(txt));
}

void QAbstractSpinBox_SetWrapping(void* ptr, int w){
	static_cast<QAbstractSpinBox*>(ptr)->setWrapping(w != 0);
}

char* QAbstractSpinBox_SpecialValueText(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->specialValueText().toUtf8().data();
}

char* QAbstractSpinBox_Text(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->text().toUtf8().data();
}

int QAbstractSpinBox_Wrapping(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->wrapping();
}

void* QAbstractSpinBox_NewQAbstractSpinBox(void* parent){
	return new QAbstractSpinBox(static_cast<QWidget*>(parent));
}

void QAbstractSpinBox_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "clear");
}

void QAbstractSpinBox_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

void QAbstractSpinBox_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

int QAbstractSpinBox_Event(void* ptr, void* event){
	return static_cast<QAbstractSpinBox*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QAbstractSpinBox_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QAbstractSpinBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QAbstractSpinBox_InterpretText(void* ptr){
	static_cast<QAbstractSpinBox*>(ptr)->interpretText();
}

void QAbstractSpinBox_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "selectAll");
}

void QAbstractSpinBox_StepBy(void* ptr, int steps){
	static_cast<QAbstractSpinBox*>(ptr)->stepBy(steps);
}

void QAbstractSpinBox_StepDown(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepDown");
}

void QAbstractSpinBox_StepUp(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepUp");
}

void QAbstractSpinBox_DestroyQAbstractSpinBox(void* ptr){
	static_cast<QAbstractSpinBox*>(ptr)->~QAbstractSpinBox();
}

#include "qfontcombobox.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QFont>
#include <QMetaObject>
#include <QFontDatabase>
#include <QString>
#include <QVariant>
#include <QFontComboBox>
#include "_cgo_export.h"

class MyQFontComboBox: public QFontComboBox {
public:
};

int QFontComboBox_FontFilters(void* ptr){
	return static_cast<QFontComboBox*>(ptr)->fontFilters();
}

void QFontComboBox_SetCurrentFont(void* ptr, void* font){
	QMetaObject::invokeMethod(static_cast<QFontComboBox*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(font)));
}

void QFontComboBox_SetFontFilters(void* ptr, int filters){
	static_cast<QFontComboBox*>(ptr)->setFontFilters(static_cast<QFontComboBox::FontFilter>(filters));
}

void QFontComboBox_SetWritingSystem(void* ptr, int script){
	static_cast<QFontComboBox*>(ptr)->setWritingSystem(static_cast<QFontDatabase::WritingSystem>(script));
}

int QFontComboBox_WritingSystem(void* ptr){
	return static_cast<QFontComboBox*>(ptr)->writingSystem();
}

void* QFontComboBox_NewQFontComboBox(void* parent){
	return new QFontComboBox(static_cast<QWidget*>(parent));
}

void QFontComboBox_DestroyQFontComboBox(void* ptr){
	static_cast<QFontComboBox*>(ptr)->~QFontComboBox();
}

#include "qtablewidget.h"
#include <QModelIndex>
#include <QItemSelection>
#include <QWidget>
#include <QAbstractItemView>
#include <QItemSelectionModel>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QTableWidgetSelectionRange>
#include <QTableWidgetItem>
#include <QPoint>
#include <QUrl>
#include <QObject>
#include <QTableWidget>
#include "_cgo_export.h"

class MyQTableWidget: public QTableWidget {
public:
void Signal_CellActivated(int row, int column){callbackQTableWidgetCellActivated(this->objectName().toUtf8().data(), row, column);};
void Signal_CellChanged(int row, int column){callbackQTableWidgetCellChanged(this->objectName().toUtf8().data(), row, column);};
void Signal_CellClicked(int row, int column){callbackQTableWidgetCellClicked(this->objectName().toUtf8().data(), row, column);};
void Signal_CellDoubleClicked(int row, int column){callbackQTableWidgetCellDoubleClicked(this->objectName().toUtf8().data(), row, column);};
void Signal_CellEntered(int row, int column){callbackQTableWidgetCellEntered(this->objectName().toUtf8().data(), row, column);};
void Signal_CellPressed(int row, int column){callbackQTableWidgetCellPressed(this->objectName().toUtf8().data(), row, column);};
void Signal_CurrentCellChanged(int currentRow, int currentColumn, int previousRow, int previousColumn){callbackQTableWidgetCurrentCellChanged(this->objectName().toUtf8().data(), currentRow, currentColumn, previousRow, previousColumn);};
void Signal_CurrentItemChanged(QTableWidgetItem * current, QTableWidgetItem * previous){callbackQTableWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_ItemActivated(QTableWidgetItem * item){callbackQTableWidgetItemActivated(this->objectName().toUtf8().data(), item);};
void Signal_ItemChanged(QTableWidgetItem * item){callbackQTableWidgetItemChanged(this->objectName().toUtf8().data(), item);};
void Signal_ItemClicked(QTableWidgetItem * item){callbackQTableWidgetItemClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QTableWidgetItem * item){callbackQTableWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemEntered(QTableWidgetItem * item){callbackQTableWidgetItemEntered(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QTableWidgetItem * item){callbackQTableWidgetItemPressed(this->objectName().toUtf8().data(), item);};
void Signal_ItemSelectionChanged(){callbackQTableWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

void* QTableWidget_ItemAt(void* ptr, void* point){
	return static_cast<QTableWidget*>(ptr)->itemAt(*static_cast<QPoint*>(point));
}

void QTableWidget_ConnectCellActivated(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_DisconnectCellActivated(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_ConnectCellChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_DisconnectCellChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_ConnectCellClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_DisconnectCellClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_ConnectCellDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_DisconnectCellDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_ConnectCellEntered(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_DisconnectCellEntered(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_ConnectCellPressed(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

void QTableWidget_DisconnectCellPressed(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

void* QTableWidget_CellWidget(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->cellWidget(row, column);
}

void QTableWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clear");
}

void QTableWidget_ClearContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clearContents");
}

void QTableWidget_ClosePersistentEditor(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->closePersistentEditor(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_Column(void* ptr, void* item){
	return static_cast<QTableWidget*>(ptr)->column(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_ColumnCount(void* ptr){
	return static_cast<QTableWidget*>(ptr)->columnCount();
}

void QTableWidget_ConnectCurrentCellChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

void QTableWidget_DisconnectCurrentCellChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

int QTableWidget_CurrentColumn(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentColumn();
}

void* QTableWidget_CurrentItem(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentItem();
}

void QTableWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

void QTableWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

int QTableWidget_CurrentRow(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentRow();
}

void QTableWidget_EditItem(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->editItem(static_cast<QTableWidgetItem*>(item));
}

void* QTableWidget_HorizontalHeaderItem(void* ptr, int column){
	return static_cast<QTableWidget*>(ptr)->horizontalHeaderItem(column);
}

void QTableWidget_InsertColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertColumn", Q_ARG(int, column));
}

void QTableWidget_InsertRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertRow", Q_ARG(int, row));
}

void* QTableWidget_Item(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->item(row, column);
}

void QTableWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

void QTableWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

void* QTableWidget_ItemAt2(void* ptr, int ax, int ay){
	return static_cast<QTableWidget*>(ptr)->itemAt(ax, ay);
}

void QTableWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

void QTableWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

void* QTableWidget_ItemPrototype(void* ptr){
	return const_cast<QTableWidgetItem*>(static_cast<QTableWidget*>(ptr)->itemPrototype());
}

void QTableWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_OpenPersistentEditor(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->openPersistentEditor(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_RemoveCellWidget(void* ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->removeCellWidget(row, column);
}

void QTableWidget_RemoveColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeColumn", Q_ARG(int, column));
}

void QTableWidget_RemoveRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeRow", Q_ARG(int, row));
}

int QTableWidget_Row(void* ptr, void* item){
	return static_cast<QTableWidget*>(ptr)->row(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_RowCount(void* ptr){
	return static_cast<QTableWidget*>(ptr)->rowCount();
}

void QTableWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "scrollToItem", Q_ARG(QTableWidgetItem*, static_cast<QTableWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTableWidget_SetCellWidget(void* ptr, int row, int column, void* widget){
	static_cast<QTableWidget*>(ptr)->setCellWidget(row, column, static_cast<QWidget*>(widget));
}

void QTableWidget_SetColumnCount(void* ptr, int columns){
	static_cast<QTableWidget*>(ptr)->setColumnCount(columns);
}

void QTableWidget_SetCurrentCell(void* ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column);
}

void QTableWidget_SetCurrentCell2(void* ptr, int row, int column, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetCurrentItem2(void* ptr, void* item, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetHorizontalHeaderItem(void* ptr, int column, void* item){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderItem(column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetHorizontalHeaderLabels(void* ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SetItem(void* ptr, int row, int column, void* item){
	static_cast<QTableWidget*>(ptr)->setItem(row, column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetItemPrototype(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->setItemPrototype(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetRangeSelected(void* ptr, void* ran, int sele){
	static_cast<QTableWidget*>(ptr)->setRangeSelected(*static_cast<QTableWidgetSelectionRange*>(ran), sele != 0);
}

void QTableWidget_SetRowCount(void* ptr, int rows){
	static_cast<QTableWidget*>(ptr)->setRowCount(rows);
}

void QTableWidget_SetVerticalHeaderItem(void* ptr, int row, void* item){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderItem(row, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetVerticalHeaderLabels(void* ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SortItems(void* ptr, int column, int order){
	static_cast<QTableWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

void* QTableWidget_TakeHorizontalHeaderItem(void* ptr, int column){
	return static_cast<QTableWidget*>(ptr)->takeHorizontalHeaderItem(column);
}

void* QTableWidget_TakeItem(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->takeItem(row, column);
}

void* QTableWidget_TakeVerticalHeaderItem(void* ptr, int row){
	return static_cast<QTableWidget*>(ptr)->takeVerticalHeaderItem(row);
}

void* QTableWidget_VerticalHeaderItem(void* ptr, int row){
	return static_cast<QTableWidget*>(ptr)->verticalHeaderItem(row);
}

int QTableWidget_VisualColumn(void* ptr, int logicalColumn){
	return static_cast<QTableWidget*>(ptr)->visualColumn(logicalColumn);
}

int QTableWidget_VisualRow(void* ptr, int logicalRow){
	return static_cast<QTableWidget*>(ptr)->visualRow(logicalRow);
}

void QTableWidget_DestroyQTableWidget(void* ptr){
	static_cast<QTableWidget*>(ptr)->~QTableWidget();
}

#include "qcalendarwidget.h"
#include <QVariant>
#include <QUrl>
#include <QMetaObject>
#include <QDate>
#include <QObject>
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QTextCharFormat>
#include <QCalendarWidget>
#include "_cgo_export.h"

class MyQCalendarWidget: public QCalendarWidget {
public:
void Signal_CurrentPageChanged(int year, int month){callbackQCalendarWidgetCurrentPageChanged(this->objectName().toUtf8().data(), year, month);};
void Signal_SelectionChanged(){callbackQCalendarWidgetSelectionChanged(this->objectName().toUtf8().data());};
};

int QCalendarWidget_DateEditAcceptDelay(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->dateEditAcceptDelay();
}

int QCalendarWidget_FirstDayOfWeek(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->firstDayOfWeek();
}

int QCalendarWidget_HorizontalHeaderFormat(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->horizontalHeaderFormat();
}

int QCalendarWidget_IsDateEditEnabled(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isDateEditEnabled();
}

int QCalendarWidget_IsGridVisible(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isGridVisible();
}

int QCalendarWidget_IsNavigationBarVisible(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->isNavigationBarVisible();
}

int QCalendarWidget_SelectionMode(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->selectionMode();
}

void QCalendarWidget_SetDateEditAcceptDelay(void* ptr, int delay){
	static_cast<QCalendarWidget*>(ptr)->setDateEditAcceptDelay(delay);
}

void QCalendarWidget_SetDateEditEnabled(void* ptr, int enable){
	static_cast<QCalendarWidget*>(ptr)->setDateEditEnabled(enable != 0);
}

void QCalendarWidget_SetFirstDayOfWeek(void* ptr, int dayOfWeek){
	static_cast<QCalendarWidget*>(ptr)->setFirstDayOfWeek(static_cast<Qt::DayOfWeek>(dayOfWeek));
}

void QCalendarWidget_SetGridVisible(void* ptr, int show){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setGridVisible", Q_ARG(bool, show != 0));
}

void QCalendarWidget_SetHorizontalHeaderFormat(void* ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setHorizontalHeaderFormat(static_cast<QCalendarWidget::HorizontalHeaderFormat>(format));
}

void QCalendarWidget_SetMaximumDate(void* ptr, void* date){
	static_cast<QCalendarWidget*>(ptr)->setMaximumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetMinimumDate(void* ptr, void* date){
	static_cast<QCalendarWidget*>(ptr)->setMinimumDate(*static_cast<QDate*>(date));
}

void QCalendarWidget_SetNavigationBarVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setNavigationBarVisible", Q_ARG(bool, visible != 0));
}

void QCalendarWidget_SetSelectedDate(void* ptr, void* date){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setSelectedDate", Q_ARG(QDate, *static_cast<QDate*>(date)));
}

void QCalendarWidget_SetSelectionMode(void* ptr, int mode){
	static_cast<QCalendarWidget*>(ptr)->setSelectionMode(static_cast<QCalendarWidget::SelectionMode>(mode));
}

void QCalendarWidget_SetVerticalHeaderFormat(void* ptr, int format){
	static_cast<QCalendarWidget*>(ptr)->setVerticalHeaderFormat(static_cast<QCalendarWidget::VerticalHeaderFormat>(format));
}

int QCalendarWidget_VerticalHeaderFormat(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->verticalHeaderFormat();
}

void* QCalendarWidget_NewQCalendarWidget(void* parent){
	return new QCalendarWidget(static_cast<QWidget*>(parent));
}

void QCalendarWidget_ConnectCurrentPageChanged(void* ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

void QCalendarWidget_DisconnectCurrentPageChanged(void* ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)(int, int)>(&MyQCalendarWidget::Signal_CurrentPageChanged));;
}

int QCalendarWidget_MonthShown(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->monthShown();
}

void QCalendarWidget_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QCalendarWidget*>(ptr), static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), static_cast<MyQCalendarWidget*>(ptr), static_cast<void (MyQCalendarWidget::*)()>(&MyQCalendarWidget::Signal_SelectionChanged));;
}

void QCalendarWidget_SetCurrentPage(void* ptr, int year, int month){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setCurrentPage", Q_ARG(int, year), Q_ARG(int, month));
}

void QCalendarWidget_SetDateRange(void* ptr, void* min, void* max){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "setDateRange", Q_ARG(QDate, *static_cast<QDate*>(min)), Q_ARG(QDate, *static_cast<QDate*>(max)));
}

void QCalendarWidget_SetDateTextFormat(void* ptr, void* date, void* format){
	static_cast<QCalendarWidget*>(ptr)->setDateTextFormat(*static_cast<QDate*>(date), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetHeaderTextFormat(void* ptr, void* format){
	static_cast<QCalendarWidget*>(ptr)->setHeaderTextFormat(*static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_SetWeekdayTextFormat(void* ptr, int dayOfWeek, void* format){
	static_cast<QCalendarWidget*>(ptr)->setWeekdayTextFormat(static_cast<Qt::DayOfWeek>(dayOfWeek), *static_cast<QTextCharFormat*>(format));
}

void QCalendarWidget_ShowNextMonth(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextMonth");
}

void QCalendarWidget_ShowNextYear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showNextYear");
}

void QCalendarWidget_ShowPreviousMonth(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousMonth");
}

void QCalendarWidget_ShowPreviousYear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showPreviousYear");
}

void QCalendarWidget_ShowSelectedDate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showSelectedDate");
}

void QCalendarWidget_ShowToday(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCalendarWidget*>(ptr), "showToday");
}

int QCalendarWidget_YearShown(void* ptr){
	return static_cast<QCalendarWidget*>(ptr)->yearShown();
}

void QCalendarWidget_DestroyQCalendarWidget(void* ptr){
	static_cast<QCalendarWidget*>(ptr)->~QCalendarWidget();
}

#include "qgraphicsscale.h"
#include <QModelIndex>
#include <QObject>
#include <QVector3D>
#include <QVector>
#include <QMatrix4x4>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsScale>
#include "_cgo_export.h"

class MyQGraphicsScale: public QGraphicsScale {
public:
void Signal_OriginChanged(){callbackQGraphicsScaleOriginChanged(this->objectName().toUtf8().data());};
void Signal_ScaleChanged(){callbackQGraphicsScaleScaleChanged(this->objectName().toUtf8().data());};
void Signal_XScaleChanged(){callbackQGraphicsScaleXScaleChanged(this->objectName().toUtf8().data());};
void Signal_YScaleChanged(){callbackQGraphicsScaleYScaleChanged(this->objectName().toUtf8().data());};
void Signal_ZScaleChanged(){callbackQGraphicsScaleZScaleChanged(this->objectName().toUtf8().data());};
};

void QGraphicsScale_SetOrigin(void* ptr, void* point){
	static_cast<QGraphicsScale*>(ptr)->setOrigin(*static_cast<QVector3D*>(point));
}

void QGraphicsScale_SetXScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setXScale(static_cast<qreal>(v));
}

void QGraphicsScale_SetYScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setYScale(static_cast<qreal>(v));
}

void QGraphicsScale_SetZScale(void* ptr, double v){
	static_cast<QGraphicsScale*>(ptr)->setZScale(static_cast<qreal>(v));
}

double QGraphicsScale_XScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->xScale());
}

double QGraphicsScale_YScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->yScale());
}

double QGraphicsScale_ZScale(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScale*>(ptr)->zScale());
}

void* QGraphicsScale_NewQGraphicsScale(void* parent){
	return new QGraphicsScale(static_cast<QObject*>(parent));
}

void QGraphicsScale_ApplyTo(void* ptr, void* matrix){
	static_cast<QGraphicsScale*>(ptr)->applyTo(static_cast<QMatrix4x4*>(matrix));
}

void QGraphicsScale_ConnectOriginChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_DisconnectOriginChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::originChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_OriginChanged));;
}

void QGraphicsScale_ConnectScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_DisconnectScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::scaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ScaleChanged));;
}

void QGraphicsScale_ConnectXScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_DisconnectXScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::xScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_XScaleChanged));;
}

void QGraphicsScale_ConnectYScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_DisconnectYScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::yScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_YScaleChanged));;
}

void QGraphicsScale_ConnectZScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DisconnectZScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScale*>(ptr), static_cast<void (QGraphicsScale::*)()>(&QGraphicsScale::zScaleChanged), static_cast<MyQGraphicsScale*>(ptr), static_cast<void (MyQGraphicsScale::*)()>(&MyQGraphicsScale::Signal_ZScaleChanged));;
}

void QGraphicsScale_DestroyQGraphicsScale(void* ptr){
	static_cast<QGraphicsScale*>(ptr)->~QGraphicsScale();
}

#include "qstyleoptiontoolbutton.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleOptionToolButton>
#include "_cgo_export.h"

class MyQStyleOptionToolButton: public QStyleOptionToolButton {
public:
};

void* QStyleOptionToolButton_NewQStyleOptionToolButton(){
	return new QStyleOptionToolButton();
}

void* QStyleOptionToolButton_NewQStyleOptionToolButton2(void* other){
	return new QStyleOptionToolButton(*static_cast<QStyleOptionToolButton*>(other));
}

#include "qstyleoptiontab.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QStyleOptionTab>
#include "_cgo_export.h"

class MyQStyleOptionTab: public QStyleOptionTab {
public:
};

void* QStyleOptionTab_NewQStyleOptionTab(){
	return new QStyleOptionTab();
}

void* QStyleOptionTab_NewQStyleOptionTab2(void* other){
	return new QStyleOptionTab(*static_cast<QStyleOptionTab*>(other));
}

#include "qstyleoptionviewitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include "_cgo_export.h"

class MyQStyleOptionViewItem: public QStyleOptionViewItem {
public:
};

void* QStyleOptionViewItem_NewQStyleOptionViewItem(){
	return new QStyleOptionViewItem();
}

void* QStyleOptionViewItem_NewQStyleOptionViewItem2(void* other){
	return new QStyleOptionViewItem(*static_cast<QStyleOptionViewItem*>(other));
}

#include "qapplication.h"
#include <QFont>
#include <QEvent>
#include <QByteArray>
#include <QWidget>
#include <QPoint>
#include <QSize>
#include <QObject>
#include <QIcon>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QList>
#include <QMetaObject>
#include <QPalette>
#include <QApplication>
#include "_cgo_export.h"

class MyQApplication: public QApplication {
public:
void Signal_FocusChanged(QWidget * old, QWidget * now){callbackQApplicationFocusChanged(this->objectName().toUtf8().data(), old, now);};
};

void QApplication_QApplication_Alert(void* widget, int msec){
	QApplication::alert(static_cast<QWidget*>(widget), msec);
}

int QApplication_AutoMaximizeThreshold(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "autoMaximizeThreshold");
}

int QApplication_AutoSipEnabled(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "autoSipEnabled");
}

void QApplication_QApplication_Beep(){
	QApplication::beep();
}

int QApplication_QApplication_CursorFlashTime(){
	return QApplication::cursorFlashTime();
}

int QApplication_QApplication_DoubleClickInterval(){
	return QApplication::doubleClickInterval();
}

int QApplication_QApplication_IsEffectEnabled(int effect){
	return QApplication::isEffectEnabled(static_cast<Qt::UIEffect>(effect));
}

int QApplication_QApplication_KeyboardInputInterval(){
	return QApplication::keyboardInputInterval();
}

void QApplication_QApplication_SetActiveWindow(void* active){
	QApplication::setActiveWindow(static_cast<QWidget*>(active));
}

void QApplication_SetAutoMaximizeThreshold(void* ptr, int threshold){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setAutoMaximizeThreshold", Q_ARG(int, threshold));
}

void QApplication_SetAutoSipEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setAutoSipEnabled", Q_ARG(bool, enabled != 0));
}

void QApplication_QApplication_SetCursorFlashTime(int v){
	QApplication::setCursorFlashTime(v);
}

void QApplication_QApplication_SetDoubleClickInterval(int v){
	QApplication::setDoubleClickInterval(v);
}

void QApplication_QApplication_SetEffectEnabled(int effect, int enable){
	QApplication::setEffectEnabled(static_cast<Qt::UIEffect>(effect), enable != 0);
}

void QApplication_QApplication_SetGlobalStrut(void* v){
	QApplication::setGlobalStrut(*static_cast<QSize*>(v));
}

void QApplication_QApplication_SetKeyboardInputInterval(int v){
	QApplication::setKeyboardInputInterval(v);
}

void QApplication_QApplication_SetStartDragDistance(int l){
	QApplication::setStartDragDistance(l);
}

void QApplication_QApplication_SetStartDragTime(int ms){
	QApplication::setStartDragTime(ms);
}

void QApplication_SetStyleSheet(void* ptr, char* sheet){
	QMetaObject::invokeMethod(static_cast<QApplication*>(ptr), "setStyleSheet", Q_ARG(QString, QString(sheet)));
}

void QApplication_QApplication_SetWheelScrollLines(int v){
	QApplication::setWheelScrollLines(v);
}

void QApplication_QApplication_SetWindowIcon(void* icon){
	QApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

int QApplication_QApplication_StartDragDistance(){
	return QApplication::startDragDistance();
}

int QApplication_QApplication_StartDragTime(){
	return QApplication::startDragTime();
}

char* QApplication_StyleSheet(void* ptr){
	return static_cast<QApplication*>(ptr)->styleSheet().toUtf8().data();
}

void* QApplication_QApplication_TopLevelAt(void* point){
	return QApplication::topLevelAt(*static_cast<QPoint*>(point));
}

int QApplication_QApplication_WheelScrollLines(){
	return QApplication::wheelScrollLines();
}

void* QApplication_QApplication_WidgetAt(void* point){
	return QApplication::widgetAt(*static_cast<QPoint*>(point));
}

void* QApplication_NewQApplication(int argc, char* argv){
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = aList[i].data();

	return new QApplication(argcs, argvs);
}

void QApplication_QApplication_AboutQt(){
	QMetaObject::invokeMethod(QApplication::instance(), "aboutQt");
}

void* QApplication_QApplication_ActiveModalWidget(){
	return QApplication::activeModalWidget();
}

void* QApplication_QApplication_ActivePopupWidget(){
	return QApplication::activePopupWidget();
}

void* QApplication_QApplication_ActiveWindow(){
	return QApplication::activeWindow();
}

void QApplication_QApplication_CloseAllWindows(){
	QMetaObject::invokeMethod(QApplication::instance(), "closeAllWindows");
}

int QApplication_QApplication_ColorSpec(){
	return QApplication::colorSpec();
}

void* QApplication_QApplication_Desktop(){
	return QApplication::desktop();
}

int QApplication_QApplication_Exec(){
	return QApplication::exec();
}

void QApplication_ConnectFocusChanged(void* ptr){
	QObject::connect(static_cast<QApplication*>(ptr), static_cast<void (QApplication::*)(QWidget *, QWidget *)>(&QApplication::focusChanged), static_cast<MyQApplication*>(ptr), static_cast<void (MyQApplication::*)(QWidget *, QWidget *)>(&MyQApplication::Signal_FocusChanged));;
}

void QApplication_DisconnectFocusChanged(void* ptr){
	QObject::disconnect(static_cast<QApplication*>(ptr), static_cast<void (QApplication::*)(QWidget *, QWidget *)>(&QApplication::focusChanged), static_cast<MyQApplication*>(ptr), static_cast<void (MyQApplication::*)(QWidget *, QWidget *)>(&MyQApplication::Signal_FocusChanged));;
}

void* QApplication_QApplication_FocusWidget(){
	return QApplication::focusWidget();
}

int QApplication_Notify(void* ptr, void* receiver, void* e){
	return static_cast<QApplication*>(ptr)->notify(static_cast<QObject*>(receiver), static_cast<QEvent*>(e));
}

void QApplication_QApplication_SetColorSpec(int spec){
	QApplication::setColorSpec(spec);
}

void QApplication_QApplication_SetFont(void* font, char* className){
	QApplication::setFont(*static_cast<QFont*>(font), const_cast<const char*>(className));
}

void QApplication_QApplication_SetPalette(void* palette, char* className){
	QApplication::setPalette(*static_cast<QPalette*>(palette), const_cast<const char*>(className));
}

void* QApplication_QApplication_SetStyle2(char* style){
	return QApplication::setStyle(QString(style));
}

void QApplication_QApplication_SetStyle(void* style){
	QApplication::setStyle(static_cast<QStyle*>(style));
}

void* QApplication_QApplication_Style(){
	return QApplication::style();
}

void* QApplication_QApplication_TopLevelAt2(int x, int y){
	return QApplication::topLevelAt(x, y);
}

void* QApplication_QApplication_WidgetAt2(int x, int y){
	return QApplication::widgetAt(x, y);
}

void QApplication_DestroyQApplication(void* ptr){
	static_cast<QApplication*>(ptr)->~QApplication();
}

#include "qcombobox.h"
#include <QModelIndex>
#include <QIcon>
#include <QAbstractItemDelegate>
#include <QLineEdit>
#include <QUrl>
#include <QCompleter>
#include <QAbstractItemView>
#include <QEvent>
#include <QLine>
#include <QSize>
#include <QValidator>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QAbstractItemModel>
#include <QObject>
#include <QWidget>
#include <QComboBox>
#include "_cgo_export.h"

class MyQComboBox: public QComboBox {
public:
void Signal_Activated(int index){callbackQComboBoxActivated(this->objectName().toUtf8().data(), index);};
void Signal_CurrentIndexChanged(int index){callbackQComboBoxCurrentIndexChanged(this->objectName().toUtf8().data(), index);};
void Signal_CurrentTextChanged(const QString & text){callbackQComboBoxCurrentTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_EditTextChanged(const QString & text){callbackQComboBoxEditTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_Highlighted(int index){callbackQComboBoxHighlighted(this->objectName().toUtf8().data(), index);};
};

int QComboBox_Count(void* ptr){
	return static_cast<QComboBox*>(ptr)->count();
}

void* QComboBox_CurrentData(void* ptr, int role){
	return new QVariant(static_cast<QComboBox*>(ptr)->currentData(role));
}

int QComboBox_CurrentIndex(void* ptr){
	return static_cast<QComboBox*>(ptr)->currentIndex();
}

char* QComboBox_CurrentText(void* ptr){
	return static_cast<QComboBox*>(ptr)->currentText().toUtf8().data();
}

int QComboBox_DuplicatesEnabled(void* ptr){
	return static_cast<QComboBox*>(ptr)->duplicatesEnabled();
}

int QComboBox_HasFrame(void* ptr){
	return static_cast<QComboBox*>(ptr)->hasFrame();
}

int QComboBox_InsertPolicy(void* ptr){
	return static_cast<QComboBox*>(ptr)->insertPolicy();
}

int QComboBox_IsEditable(void* ptr){
	return static_cast<QComboBox*>(ptr)->isEditable();
}

int QComboBox_MaxCount(void* ptr){
	return static_cast<QComboBox*>(ptr)->maxCount();
}

int QComboBox_MaxVisibleItems(void* ptr){
	return static_cast<QComboBox*>(ptr)->maxVisibleItems();
}

int QComboBox_MinimumContentsLength(void* ptr){
	return static_cast<QComboBox*>(ptr)->minimumContentsLength();
}

int QComboBox_ModelColumn(void* ptr){
	return static_cast<QComboBox*>(ptr)->modelColumn();
}

void QComboBox_SetCompleter(void* ptr, void* completer){
	static_cast<QComboBox*>(ptr)->setCompleter(static_cast<QCompleter*>(completer));
}

void QComboBox_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QComboBox_SetCurrentText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetDuplicatesEnabled(void* ptr, int enable){
	static_cast<QComboBox*>(ptr)->setDuplicatesEnabled(enable != 0);
}

void QComboBox_SetEditable(void* ptr, int editable){
	static_cast<QComboBox*>(ptr)->setEditable(editable != 0);
}

void QComboBox_SetFrame(void* ptr, int v){
	static_cast<QComboBox*>(ptr)->setFrame(v != 0);
}

void QComboBox_SetIconSize(void* ptr, void* size){
	static_cast<QComboBox*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QComboBox_SetInsertPolicy(void* ptr, int policy){
	static_cast<QComboBox*>(ptr)->setInsertPolicy(static_cast<QComboBox::InsertPolicy>(policy));
}

void QComboBox_SetMaxCount(void* ptr, int max){
	static_cast<QComboBox*>(ptr)->setMaxCount(max);
}

void QComboBox_SetMaxVisibleItems(void* ptr, int maxItems){
	static_cast<QComboBox*>(ptr)->setMaxVisibleItems(maxItems);
}

void QComboBox_SetMinimumContentsLength(void* ptr, int characters){
	static_cast<QComboBox*>(ptr)->setMinimumContentsLength(characters);
}

void QComboBox_SetModelColumn(void* ptr, int visibleColumn){
	static_cast<QComboBox*>(ptr)->setModelColumn(visibleColumn);
}

void QComboBox_SetSizeAdjustPolicy(void* ptr, int policy){
	static_cast<QComboBox*>(ptr)->setSizeAdjustPolicy(static_cast<QComboBox::SizeAdjustPolicy>(policy));
}

void QComboBox_SetValidator(void* ptr, void* validator){
	static_cast<QComboBox*>(ptr)->setValidator(static_cast<QValidator*>(validator));
}

int QComboBox_SizeAdjustPolicy(void* ptr){
	return static_cast<QComboBox*>(ptr)->sizeAdjustPolicy();
}

void* QComboBox_NewQComboBox(void* parent){
	return new QComboBox(static_cast<QWidget*>(parent));
}

void QComboBox_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_AddItem2(void* ptr, void* icon, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItem(void* ptr, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->addItem(QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItems(void* ptr, char* texts){
	static_cast<QComboBox*>(ptr)->addItems(QString(texts).split("|", QString::SkipEmptyParts));
}

void QComboBox_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clear");
}

void QComboBox_ClearEditText(void* ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clearEditText");
}

void* QComboBox_Completer(void* ptr){
	return static_cast<QComboBox*>(ptr)->completer();
}

void QComboBox_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_ConnectCurrentTextChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_DisconnectCurrentTextChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_ConnectEditTextChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

void QComboBox_DisconnectEditTextChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

int QComboBox_Event(void* ptr, void* event){
	return static_cast<QComboBox*>(ptr)->event(static_cast<QEvent*>(event));
}

int QComboBox_FindData(void* ptr, void* data, int role, int flags){
	return static_cast<QComboBox*>(ptr)->findData(*static_cast<QVariant*>(data), role, static_cast<Qt::MatchFlag>(flags));
}

int QComboBox_FindText(void* ptr, char* text, int flags){
	return static_cast<QComboBox*>(ptr)->findText(QString(text), static_cast<Qt::MatchFlag>(flags));
}

void QComboBox_HidePopup(void* ptr){
	static_cast<QComboBox*>(ptr)->hidePopup();
}

void QComboBox_ConnectHighlighted(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

void QComboBox_DisconnectHighlighted(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

void* QComboBox_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QComboBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QComboBox_InsertItem2(void* ptr, int index, void* icon, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, *static_cast<QIcon*>(icon), QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_InsertItem(void* ptr, int index, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_InsertItems(void* ptr, int index, char* list){
	static_cast<QComboBox*>(ptr)->insertItems(index, QString(list).split("|", QString::SkipEmptyParts));
}

void QComboBox_InsertSeparator(void* ptr, int index){
	static_cast<QComboBox*>(ptr)->insertSeparator(index);
}

void* QComboBox_ItemData(void* ptr, int index, int role){
	return new QVariant(static_cast<QComboBox*>(ptr)->itemData(index, role));
}

void* QComboBox_ItemDelegate(void* ptr){
	return static_cast<QComboBox*>(ptr)->itemDelegate();
}

char* QComboBox_ItemText(void* ptr, int index){
	return static_cast<QComboBox*>(ptr)->itemText(index).toUtf8().data();
}

void* QComboBox_LineEdit(void* ptr){
	return static_cast<QComboBox*>(ptr)->lineEdit();
}

void* QComboBox_Model(void* ptr){
	return static_cast<QComboBox*>(ptr)->model();
}

void QComboBox_RemoveItem(void* ptr, int index){
	static_cast<QComboBox*>(ptr)->removeItem(index);
}

void* QComboBox_RootModelIndex(void* ptr){
	return static_cast<QComboBox*>(ptr)->rootModelIndex().internalPointer();
}

void QComboBox_SetEditText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setEditText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetItemData(void* ptr, int index, void* value, int role){
	static_cast<QComboBox*>(ptr)->setItemData(index, *static_cast<QVariant*>(value), role);
}

void QComboBox_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QComboBox*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QComboBox_SetItemIcon(void* ptr, int index, void* icon){
	static_cast<QComboBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QComboBox_SetItemText(void* ptr, int index, char* text){
	static_cast<QComboBox*>(ptr)->setItemText(index, QString(text));
}

void QComboBox_SetLineEdit(void* ptr, void* edit){
	static_cast<QComboBox*>(ptr)->setLineEdit(static_cast<QLineEdit*>(edit));
}

void QComboBox_SetModel(void* ptr, void* model){
	static_cast<QComboBox*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QComboBox_SetRootModelIndex(void* ptr, void* index){
	static_cast<QComboBox*>(ptr)->setRootModelIndex(*static_cast<QModelIndex*>(index));
}

void QComboBox_SetView(void* ptr, void* itemView){
	static_cast<QComboBox*>(ptr)->setView(static_cast<QAbstractItemView*>(itemView));
}

void QComboBox_ShowPopup(void* ptr){
	static_cast<QComboBox*>(ptr)->showPopup();
}

void* QComboBox_Validator(void* ptr){
	return const_cast<QValidator*>(static_cast<QComboBox*>(ptr)->validator());
}

void* QComboBox_View(void* ptr){
	return static_cast<QComboBox*>(ptr)->view();
}

void QComboBox_DestroyQComboBox(void* ptr){
	static_cast<QComboBox*>(ptr)->~QComboBox();
}

#include "qgraphicsproxywidget.h"
#include <QWidget>
#include <QStyleOption>
#include <QGraphicsItem>
#include <QRect>
#include <QRectF>
#include <QStyle>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionGraphicsItem>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QGraphicsProxyWidget>
#include "_cgo_export.h"

class MyQGraphicsProxyWidget: public QGraphicsProxyWidget {
public:
};

void* QGraphicsProxyWidget_NewQGraphicsProxyWidget(void* parent, int wFlags){
	return new QGraphicsProxyWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

void* QGraphicsProxyWidget_CreateProxyForChildWidget(void* ptr, void* child){
	return static_cast<QGraphicsProxyWidget*>(ptr)->createProxyForChildWidget(static_cast<QWidget*>(child));
}

void QGraphicsProxyWidget_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsProxyWidget_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsProxyWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsProxyWidget_SetWidget(void* ptr, void* widget){
	static_cast<QGraphicsProxyWidget*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

int QGraphicsProxyWidget_Type(void* ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->type();
}

void* QGraphicsProxyWidget_Widget(void* ptr){
	return static_cast<QGraphicsProxyWidget*>(ptr)->widget();
}

void QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(void* ptr){
	static_cast<QGraphicsProxyWidget*>(ptr)->~QGraphicsProxyWidget();
}

#include "qtapgesture.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QTapGesture>
#include "_cgo_export.h"

class MyQTapGesture: public QTapGesture {
public:
};

void QTapGesture_SetPosition(void* ptr, void* pos){
	static_cast<QTapGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapGesture_DestroyQTapGesture(void* ptr){
	static_cast<QTapGesture*>(ptr)->~QTapGesture();
}

#include "qpinchgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QPinchGesture>
#include "_cgo_export.h"

class MyQPinchGesture: public QPinchGesture {
public:
};

int QPinchGesture_ChangeFlags(void* ptr){
	return static_cast<QPinchGesture*>(ptr)->changeFlags();
}

double QPinchGesture_LastRotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->lastRotationAngle());
}

double QPinchGesture_LastScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->lastScaleFactor());
}

double QPinchGesture_RotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->rotationAngle());
}

double QPinchGesture_ScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->scaleFactor());
}

void QPinchGesture_SetCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetChangeFlags(void* ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

void QPinchGesture_SetLastCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setLastCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetLastRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setLastRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetLastScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setLastScaleFactor(static_cast<qreal>(value));
}

void QPinchGesture_SetRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setScaleFactor(static_cast<qreal>(value));
}

void QPinchGesture_SetStartCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setStartCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetTotalChangeFlags(void* ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setTotalChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

void QPinchGesture_SetTotalRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setTotalRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetTotalScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setTotalScaleFactor(static_cast<qreal>(value));
}

int QPinchGesture_TotalChangeFlags(void* ptr){
	return static_cast<QPinchGesture*>(ptr)->totalChangeFlags();
}

double QPinchGesture_TotalRotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->totalRotationAngle());
}

double QPinchGesture_TotalScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->totalScaleFactor());
}

void QPinchGesture_DestroyQPinchGesture(void* ptr){
	static_cast<QPinchGesture*>(ptr)->~QPinchGesture();
}

#include "qcolordialog.h"
#include <QObject>
#include <QWidget>
#include <QColor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColorDialog>
#include "_cgo_export.h"

class MyQColorDialog: public QColorDialog {
public:
void Signal_ColorSelected(const QColor & color){callbackQColorDialogColorSelected(this->objectName().toUtf8().data(), new QColor(color));};
void Signal_CurrentColorChanged(const QColor & color){callbackQColorDialogCurrentColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

void* QColorDialog_CurrentColor(void* ptr){
	return new QColor(static_cast<QColorDialog*>(ptr)->currentColor());
}

int QColorDialog_Options(void* ptr){
	return static_cast<QColorDialog*>(ptr)->options();
}

void QColorDialog_SetCurrentColor(void* ptr, void* color){
	static_cast<QColorDialog*>(ptr)->setCurrentColor(*static_cast<QColor*>(color));
}

void QColorDialog_SetOptions(void* ptr, int options){
	static_cast<QColorDialog*>(ptr)->setOptions(static_cast<QColorDialog::ColorDialogOption>(options));
}

void* QColorDialog_NewQColorDialog(void* parent){
	return new QColorDialog(static_cast<QWidget*>(parent));
}

void* QColorDialog_NewQColorDialog2(void* initial, void* parent){
	return new QColorDialog(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent));
}

void QColorDialog_ConnectColorSelected(void* ptr){
	QObject::connect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::colorSelected), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_ColorSelected));;
}

void QColorDialog_DisconnectColorSelected(void* ptr){
	QObject::disconnect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::colorSelected), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_ColorSelected));;
}

void QColorDialog_ConnectCurrentColorChanged(void* ptr){
	QObject::connect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::currentColorChanged), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_CurrentColorChanged));;
}

void QColorDialog_DisconnectCurrentColorChanged(void* ptr){
	QObject::disconnect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::currentColorChanged), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_CurrentColorChanged));;
}

void* QColorDialog_QColorDialog_CustomColor(int index){
	return new QColor(QColorDialog::customColor(index));
}

int QColorDialog_QColorDialog_CustomCount(){
	return QColorDialog::customCount();
}

void* QColorDialog_QColorDialog_GetColor(void* initial, void* parent, char* title, int options){
	return new QColor(QColorDialog::getColor(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent), QString(title), static_cast<QColorDialog::ColorDialogOption>(options)));
}

void QColorDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QColorDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QColorDialog_SelectedColor(void* ptr){
	return new QColor(static_cast<QColorDialog*>(ptr)->selectedColor());
}

void QColorDialog_QColorDialog_SetCustomColor(int index, void* color){
	QColorDialog::setCustomColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetOption(void* ptr, int option, int on){
	static_cast<QColorDialog*>(ptr)->setOption(static_cast<QColorDialog::ColorDialogOption>(option), on != 0);
}

void QColorDialog_QColorDialog_SetStandardColor(int index, void* color){
	QColorDialog::setStandardColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetVisible(void* ptr, int visible){
	static_cast<QColorDialog*>(ptr)->setVisible(visible != 0);
}

void* QColorDialog_QColorDialog_StandardColor(int index){
	return new QColor(QColorDialog::standardColor(index));
}

int QColorDialog_TestOption(void* ptr, int option){
	return static_cast<QColorDialog*>(ptr)->testOption(static_cast<QColorDialog::ColorDialogOption>(option));
}

void QColorDialog_DestroyQColorDialog(void* ptr){
	static_cast<QColorDialog*>(ptr)->~QColorDialog();
}

#include "qpushbutton.h"
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QIcon>
#include <QPushButton>
#include "_cgo_export.h"

class MyQPushButton: public QPushButton {
public:
};

int QPushButton_AutoDefault(void* ptr){
	return static_cast<QPushButton*>(ptr)->autoDefault();
}

int QPushButton_IsDefault(void* ptr){
	return static_cast<QPushButton*>(ptr)->isDefault();
}

int QPushButton_IsFlat(void* ptr){
	return static_cast<QPushButton*>(ptr)->isFlat();
}

void QPushButton_SetAutoDefault(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setAutoDefault(v != 0);
}

void QPushButton_SetDefault(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setDefault(v != 0);
}

void QPushButton_SetFlat(void* ptr, int v){
	static_cast<QPushButton*>(ptr)->setFlat(v != 0);
}

void* QPushButton_NewQPushButton(void* parent){
	return new QPushButton(static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton3(void* icon, char* text, void* parent){
	return new QPushButton(*static_cast<QIcon*>(icon), QString(text), static_cast<QWidget*>(parent));
}

void* QPushButton_NewQPushButton2(char* text, void* parent){
	return new QPushButton(QString(text), static_cast<QWidget*>(parent));
}

void* QPushButton_Menu(void* ptr){
	return static_cast<QPushButton*>(ptr)->menu();
}

void QPushButton_SetMenu(void* ptr, void* menu){
	static_cast<QPushButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QPushButton_ShowMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPushButton*>(ptr), "showMenu");
}

void QPushButton_DestroyQPushButton(void* ptr){
	static_cast<QPushButton*>(ptr)->~QPushButton();
}

#include "qstyleoptiontoolbar.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionToolBar>
#include "_cgo_export.h"

class MyQStyleOptionToolBar: public QStyleOptionToolBar {
public:
};

void* QStyleOptionToolBar_NewQStyleOptionToolBar(){
	return new QStyleOptionToolBar();
}

void* QStyleOptionToolBar_NewQStyleOptionToolBar2(void* other){
	return new QStyleOptionToolBar(*static_cast<QStyleOptionToolBar*>(other));
}

#include "qundostack.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QUndoCommand>
#include <QMetaObject>
#include <QUndoStack>
#include "_cgo_export.h"

class MyQUndoStack: public QUndoStack {
public:
void Signal_CanRedoChanged(bool canRedo){callbackQUndoStackCanRedoChanged(this->objectName().toUtf8().data(), canRedo);};
void Signal_CanUndoChanged(bool canUndo){callbackQUndoStackCanUndoChanged(this->objectName().toUtf8().data(), canUndo);};
void Signal_CleanChanged(bool clean){callbackQUndoStackCleanChanged(this->objectName().toUtf8().data(), clean);};
void Signal_IndexChanged(int idx){callbackQUndoStackIndexChanged(this->objectName().toUtf8().data(), idx);};
void Signal_RedoTextChanged(const QString & redoText){callbackQUndoStackRedoTextChanged(this->objectName().toUtf8().data(), redoText.toUtf8().data());};
void Signal_UndoTextChanged(const QString & undoText){callbackQUndoStackUndoTextChanged(this->objectName().toUtf8().data(), undoText.toUtf8().data());};
};

int QUndoStack_IsActive(void* ptr){
	return static_cast<QUndoStack*>(ptr)->isActive();
}

void QUndoStack_SetActive(void* ptr, int active){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setActive", Q_ARG(bool, active != 0));
}

void QUndoStack_SetUndoLimit(void* ptr, int limit){
	static_cast<QUndoStack*>(ptr)->setUndoLimit(limit);
}

int QUndoStack_UndoLimit(void* ptr){
	return static_cast<QUndoStack*>(ptr)->undoLimit();
}

void* QUndoStack_NewQUndoStack(void* parent){
	return new QUndoStack(static_cast<QObject*>(parent));
}

void QUndoStack_BeginMacro(void* ptr, char* text){
	static_cast<QUndoStack*>(ptr)->beginMacro(QString(text));
}

int QUndoStack_CanRedo(void* ptr){
	return static_cast<QUndoStack*>(ptr)->canRedo();
}

void QUndoStack_ConnectCanRedoChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

void QUndoStack_DisconnectCanRedoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

int QUndoStack_CanUndo(void* ptr){
	return static_cast<QUndoStack*>(ptr)->canUndo();
}

void QUndoStack_ConnectCanUndoChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_DisconnectCanUndoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_ConnectCleanChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

void QUndoStack_DisconnectCleanChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

int QUndoStack_CleanIndex(void* ptr){
	return static_cast<QUndoStack*>(ptr)->cleanIndex();
}

void QUndoStack_Clear(void* ptr){
	static_cast<QUndoStack*>(ptr)->clear();
}

void* QUndoStack_Command(void* ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoStack*>(ptr)->command(index));
}

int QUndoStack_Count(void* ptr){
	return static_cast<QUndoStack*>(ptr)->count();
}

void* QUndoStack_CreateRedoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createRedoAction(static_cast<QObject*>(parent), QString(prefix));
}

void* QUndoStack_CreateUndoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createUndoAction(static_cast<QObject*>(parent), QString(prefix));
}

void QUndoStack_EndMacro(void* ptr){
	static_cast<QUndoStack*>(ptr)->endMacro();
}

int QUndoStack_Index(void* ptr){
	return static_cast<QUndoStack*>(ptr)->index();
}

void QUndoStack_ConnectIndexChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

void QUndoStack_DisconnectIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

int QUndoStack_IsClean(void* ptr){
	return static_cast<QUndoStack*>(ptr)->isClean();
}

void QUndoStack_Push(void* ptr, void* cmd){
	static_cast<QUndoStack*>(ptr)->push(static_cast<QUndoCommand*>(cmd));
}

void QUndoStack_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "redo");
}

char* QUndoStack_RedoText(void* ptr){
	return static_cast<QUndoStack*>(ptr)->redoText().toUtf8().data();
}

void QUndoStack_ConnectRedoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_DisconnectRedoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_SetClean(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setClean");
}

void QUndoStack_SetIndex(void* ptr, int idx){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setIndex", Q_ARG(int, idx));
}

char* QUndoStack_Text(void* ptr, int idx){
	return static_cast<QUndoStack*>(ptr)->text(idx).toUtf8().data();
}

void QUndoStack_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "undo");
}

char* QUndoStack_UndoText(void* ptr){
	return static_cast<QUndoStack*>(ptr)->undoText().toUtf8().data();
}

void QUndoStack_ConnectUndoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DisconnectUndoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DestroyQUndoStack(void* ptr){
	static_cast<QUndoStack*>(ptr)->~QUndoStack();
}

#include "qlabel.h"
#include <QMovie>
#include <QMetaObject>
#include <QPicture>
#include <QWidget>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QPixmap>
#include <QString>
#include <QUrl>
#include <QLabel>
#include "_cgo_export.h"

class MyQLabel: public QLabel {
public:
void Signal_LinkActivated(const QString & link){callbackQLabelLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQLabelLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QLabel_Alignment(void* ptr){
	return static_cast<QLabel*>(ptr)->alignment();
}

int QLabel_HasScaledContents(void* ptr){
	return static_cast<QLabel*>(ptr)->hasScaledContents();
}

int QLabel_HasSelectedText(void* ptr){
	return static_cast<QLabel*>(ptr)->hasSelectedText();
}

int QLabel_Indent(void* ptr){
	return static_cast<QLabel*>(ptr)->indent();
}

int QLabel_Margin(void* ptr){
	return static_cast<QLabel*>(ptr)->margin();
}

int QLabel_OpenExternalLinks(void* ptr){
	return static_cast<QLabel*>(ptr)->openExternalLinks();
}

void* QLabel_Pixmap(void* ptr){
	return const_cast<QPixmap*>(static_cast<QLabel*>(ptr)->pixmap());
}

char* QLabel_SelectedText(void* ptr){
	return static_cast<QLabel*>(ptr)->selectedText().toUtf8().data();
}

void QLabel_SetAlignment(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(v));
}

void QLabel_SetIndent(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setIndent(v);
}

void QLabel_SetMargin(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setMargin(v);
}

void QLabel_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QLabel*>(ptr)->setOpenExternalLinks(open != 0);
}

void QLabel_SetPixmap(void* ptr, void* v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPixmap", Q_ARG(QPixmap, *static_cast<QPixmap*>(v)));
}

void QLabel_SetScaledContents(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setScaledContents(v != 0);
}

void QLabel_SetText(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

void QLabel_SetTextFormat(void* ptr, int v){
	static_cast<QLabel*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(v));
}

void QLabel_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QLabel*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QLabel_SetWordWrap(void* ptr, int on){
	static_cast<QLabel*>(ptr)->setWordWrap(on != 0);
}

char* QLabel_Text(void* ptr){
	return static_cast<QLabel*>(ptr)->text().toUtf8().data();
}

int QLabel_TextFormat(void* ptr){
	return static_cast<QLabel*>(ptr)->textFormat();
}

int QLabel_TextInteractionFlags(void* ptr){
	return static_cast<QLabel*>(ptr)->textInteractionFlags();
}

int QLabel_WordWrap(void* ptr){
	return static_cast<QLabel*>(ptr)->wordWrap();
}

void* QLabel_NewQLabel(void* parent, int f){
	return new QLabel(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QLabel_NewQLabel2(char* text, void* parent, int f){
	return new QLabel(QString(text), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QLabel_Buddy(void* ptr){
	return static_cast<QLabel*>(ptr)->buddy();
}

void QLabel_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "clear");
}

int QLabel_HeightForWidth(void* ptr, int w){
	return static_cast<QLabel*>(ptr)->heightForWidth(w);
}

void QLabel_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkActivated), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkActivated));;
}

void QLabel_ConnectLinkHovered(void* ptr){
	QObject::connect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

void QLabel_DisconnectLinkHovered(void* ptr){
	QObject::disconnect(static_cast<QLabel*>(ptr), static_cast<void (QLabel::*)(const QString &)>(&QLabel::linkHovered), static_cast<MyQLabel*>(ptr), static_cast<void (MyQLabel::*)(const QString &)>(&MyQLabel::Signal_LinkHovered));;
}

void* QLabel_Movie(void* ptr){
	return static_cast<QLabel*>(ptr)->movie();
}

void* QLabel_Picture(void* ptr){
	return const_cast<QPicture*>(static_cast<QLabel*>(ptr)->picture());
}

int QLabel_SelectionStart(void* ptr){
	return static_cast<QLabel*>(ptr)->selectionStart();
}

void QLabel_SetBuddy(void* ptr, void* buddy){
	static_cast<QLabel*>(ptr)->setBuddy(static_cast<QWidget*>(buddy));
}

void QLabel_SetMovie(void* ptr, void* movie){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setMovie", Q_ARG(QMovie*, static_cast<QMovie*>(movie)));
}

void QLabel_SetNum(void* ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setNum", Q_ARG(int, num));
}

void QLabel_SetPicture(void* ptr, void* picture){
	QMetaObject::invokeMethod(static_cast<QLabel*>(ptr), "setPicture", Q_ARG(QPicture, *static_cast<QPicture*>(picture)));
}

void QLabel_SetSelection(void* ptr, int start, int length){
	static_cast<QLabel*>(ptr)->setSelection(start, length);
}

void QLabel_DestroyQLabel(void* ptr){
	static_cast<QLabel*>(ptr)->~QLabel();
}

#include "qactiongroup.h"
#include <QAction>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QIcon>
#include <QActionGroup>
#include "_cgo_export.h"

class MyQActionGroup: public QActionGroup {
public:
void Signal_Hovered(QAction * action){callbackQActionGroupHovered(this->objectName().toUtf8().data(), action);};
void Signal_Triggered(QAction * action){callbackQActionGroupTriggered(this->objectName().toUtf8().data(), action);};
};

void* QActionGroup_AddAction(void* ptr, void* action){
	return static_cast<QActionGroup*>(ptr)->addAction(static_cast<QAction*>(action));
}

int QActionGroup_IsEnabled(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isEnabled();
}

int QActionGroup_IsExclusive(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isExclusive();
}

int QActionGroup_IsVisible(void* ptr){
	return static_cast<QActionGroup*>(ptr)->isVisible();
}

void QActionGroup_SetEnabled(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setEnabled", Q_ARG(bool, v != 0));
}

void QActionGroup_SetExclusive(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setExclusive", Q_ARG(bool, v != 0));
}

void QActionGroup_SetVisible(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setVisible", Q_ARG(bool, v != 0));
}

void* QActionGroup_NewQActionGroup(void* parent){
	return new QActionGroup(static_cast<QObject*>(parent));
}

void* QActionGroup_AddAction3(void* ptr, void* icon, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QActionGroup_AddAction2(void* ptr, char* text){
	return static_cast<QActionGroup*>(ptr)->addAction(QString(text));
}

void* QActionGroup_CheckedAction(void* ptr){
	return static_cast<QActionGroup*>(ptr)->checkedAction();
}

void QActionGroup_ConnectHovered(void* ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_DisconnectHovered(void* ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::hovered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Hovered));;
}

void QActionGroup_RemoveAction(void* ptr, void* action){
	static_cast<QActionGroup*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QActionGroup_SetDisabled(void* ptr, int b){
	QMetaObject::invokeMethod(static_cast<QActionGroup*>(ptr), "setDisabled", Q_ARG(bool, b != 0));
}

void QActionGroup_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QActionGroup*>(ptr), static_cast<void (QActionGroup::*)(QAction *)>(&QActionGroup::triggered), static_cast<MyQActionGroup*>(ptr), static_cast<void (MyQActionGroup::*)(QAction *)>(&MyQActionGroup::Signal_Triggered));;
}

void QActionGroup_DestroyQActionGroup(void* ptr){
	static_cast<QActionGroup*>(ptr)->~QActionGroup();
}

#include "qsizegrip.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QSize>
#include <QSizeGrip>
#include "_cgo_export.h"

class MyQSizeGrip: public QSizeGrip {
public:
};

void* QSizeGrip_NewQSizeGrip(void* parent){
	return new QSizeGrip(static_cast<QWidget*>(parent));
}

void QSizeGrip_SetVisible(void* ptr, int visible){
	static_cast<QSizeGrip*>(ptr)->setVisible(visible != 0);
}

void QSizeGrip_DestroyQSizeGrip(void* ptr){
	static_cast<QSizeGrip*>(ptr)->~QSizeGrip();
}

#include "qshortcut.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QKeySequence>
#include <QShortcut>
#include "_cgo_export.h"

class MyQShortcut: public QShortcut {
public:
void Signal_Activated(){callbackQShortcutActivated(this->objectName().toUtf8().data());};
void Signal_ActivatedAmbiguously(){callbackQShortcutActivatedAmbiguously(this->objectName().toUtf8().data());};
};

int QShortcut_AutoRepeat(void* ptr){
	return static_cast<QShortcut*>(ptr)->autoRepeat();
}

int QShortcut_Context(void* ptr){
	return static_cast<QShortcut*>(ptr)->context();
}

int QShortcut_IsEnabled(void* ptr){
	return static_cast<QShortcut*>(ptr)->isEnabled();
}

void QShortcut_SetAutoRepeat(void* ptr, int on){
	static_cast<QShortcut*>(ptr)->setAutoRepeat(on != 0);
}

void QShortcut_SetContext(void* ptr, int context){
	static_cast<QShortcut*>(ptr)->setContext(static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_SetEnabled(void* ptr, int enable){
	static_cast<QShortcut*>(ptr)->setEnabled(enable != 0);
}

void QShortcut_SetKey(void* ptr, void* key){
	static_cast<QShortcut*>(ptr)->setKey(*static_cast<QKeySequence*>(key));
}

void QShortcut_SetWhatsThis(void* ptr, char* text){
	static_cast<QShortcut*>(ptr)->setWhatsThis(QString(text));
}

char* QShortcut_WhatsThis(void* ptr){
	return static_cast<QShortcut*>(ptr)->whatsThis().toUtf8().data();
}

void* QShortcut_NewQShortcut(void* parent){
	return new QShortcut(static_cast<QWidget*>(parent));
}

void* QShortcut_NewQShortcut2(void* key, void* parent, char* member, char* ambiguousMember, int context){
	return new QShortcut(*static_cast<QKeySequence*>(key), static_cast<QWidget*>(parent), const_cast<const char*>(member), const_cast<const char*>(ambiguousMember), static_cast<Qt::ShortcutContext>(context));
}

void QShortcut_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activated), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_Activated));;
}

void QShortcut_ConnectActivatedAmbiguously(void* ptr){
	QObject::connect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

void QShortcut_DisconnectActivatedAmbiguously(void* ptr){
	QObject::disconnect(static_cast<QShortcut*>(ptr), static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), static_cast<MyQShortcut*>(ptr), static_cast<void (MyQShortcut::*)()>(&MyQShortcut::Signal_ActivatedAmbiguously));;
}

int QShortcut_Id(void* ptr){
	return static_cast<QShortcut*>(ptr)->id();
}

void* QShortcut_ParentWidget(void* ptr){
	return static_cast<QShortcut*>(ptr)->parentWidget();
}

void QShortcut_DestroyQShortcut(void* ptr){
	static_cast<QShortcut*>(ptr)->~QShortcut();
}

#include "qgraphicseffect.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsEffect>
#include "_cgo_export.h"

#include "qitemdelegate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QStyle>
#include <QItemEditorFactory>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QPainter>
#include <QObject>
#include <QWidget>
#include <QItemDelegate>
#include "_cgo_export.h"

class MyQItemDelegate: public QItemDelegate {
public:
};

int QItemDelegate_HasClipping(void* ptr){
	return static_cast<QItemDelegate*>(ptr)->hasClipping();
}

void QItemDelegate_SetClipping(void* ptr, int clip){
	static_cast<QItemDelegate*>(ptr)->setClipping(clip != 0);
}

void* QItemDelegate_NewQItemDelegate(void* parent){
	return new QItemDelegate(static_cast<QObject*>(parent));
}

void* QItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void* QItemDelegate_ItemEditorFactory(void* ptr){
	return static_cast<QItemDelegate*>(ptr)->itemEditorFactory();
}

void QItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetItemEditorFactory(void* ptr, void* factory){
	static_cast<QItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_DestroyQItemDelegate(void* ptr){
	static_cast<QItemDelegate*>(ptr)->~QItemDelegate();
}

#include "qtablewidgetitem.h"
#include <QVariant>
#include <QUrl>
#include <QBrush>
#include <QSize>
#include <QIcon>
#include <QDataStream>
#include <QString>
#include <QModelIndex>
#include <QTableWidget>
#include <QFont>
#include <QTableWidgetItem>
#include "_cgo_export.h"

class MyQTableWidgetItem: public QTableWidgetItem {
public:
};

void QTableWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QTableWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QTableWidgetItem_NewQTableWidgetItem3(void* icon, char* text, int ty){
	return new QTableWidgetItem(*static_cast<QIcon*>(icon), QString(text), ty);
}

void* QTableWidgetItem_NewQTableWidgetItem2(char* text, int ty){
	return new QTableWidgetItem(QString(text), ty);
}

void* QTableWidgetItem_NewQTableWidgetItem4(void* other){
	return new QTableWidgetItem(*static_cast<QTableWidgetItem*>(other));
}

void* QTableWidgetItem_NewQTableWidgetItem(int ty){
	return new QTableWidgetItem(ty);
}

void* QTableWidgetItem_Background(void* ptr){
	return new QBrush(static_cast<QTableWidgetItem*>(ptr)->background());
}

int QTableWidgetItem_CheckState(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->checkState();
}

void* QTableWidgetItem_Clone(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->clone();
}

int QTableWidgetItem_Column(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->column();
}

void* QTableWidgetItem_Data(void* ptr, int role){
	return new QVariant(static_cast<QTableWidgetItem*>(ptr)->data(role));
}

int QTableWidgetItem_Flags(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->flags();
}

void* QTableWidgetItem_Foreground(void* ptr){
	return new QBrush(static_cast<QTableWidgetItem*>(ptr)->foreground());
}

int QTableWidgetItem_IsSelected(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->isSelected();
}

void QTableWidgetItem_Read(void* ptr, void* in){
	static_cast<QTableWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

int QTableWidgetItem_Row(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->row();
}

void QTableWidgetItem_SetBackground(void* ptr, void* brush){
	static_cast<QTableWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetCheckState(void* ptr, int state){
	static_cast<QTableWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QTableWidgetItem_SetData(void* ptr, int role, void* value){
	static_cast<QTableWidgetItem*>(ptr)->setData(role, *static_cast<QVariant*>(value));
}

void QTableWidgetItem_SetFont(void* ptr, void* font){
	static_cast<QTableWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTableWidgetItem_SetForeground(void* ptr, void* brush){
	static_cast<QTableWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTableWidgetItem_SetIcon(void* ptr, void* icon){
	static_cast<QTableWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QTableWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QTableWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QTableWidgetItem_SetSizeHint(void* ptr, void* size){
	static_cast<QTableWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QTableWidgetItem_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QTableWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QTableWidgetItem_SetText(void* ptr, char* text){
	static_cast<QTableWidgetItem*>(ptr)->setText(QString(text));
}

void QTableWidgetItem_SetTextAlignment(void* ptr, int alignment){
	static_cast<QTableWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QTableWidgetItem_SetToolTip(void* ptr, char* toolTip){
	static_cast<QTableWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QTableWidgetItem_SetWhatsThis(void* ptr, char* whatsThis){
	static_cast<QTableWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QTableWidgetItem_StatusTip(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

void* QTableWidgetItem_TableWidget(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->tableWidget();
}

char* QTableWidgetItem_Text(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->text().toUtf8().data();
}

int QTableWidgetItem_TextAlignment(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->textAlignment();
}

char* QTableWidgetItem_ToolTip(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QTableWidgetItem_Type(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->type();
}

char* QTableWidgetItem_WhatsThis(void* ptr){
	return static_cast<QTableWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QTableWidgetItem_Write(void* ptr, void* out){
	static_cast<QTableWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QTableWidgetItem_DestroyQTableWidgetItem(void* ptr){
	static_cast<QTableWidgetItem*>(ptr)->~QTableWidgetItem();
}

#include "qmacnativewidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMacNativeWidget>
#include "_cgo_export.h"

class MyQMacNativeWidget: public QMacNativeWidget {
public:
};

void QMacNativeWidget_DestroyQMacNativeWidget(void* ptr){
	static_cast<QMacNativeWidget*>(ptr)->~QMacNativeWidget();
}

#include "qhboxlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QHBoxLayout>
#include "_cgo_export.h"

class MyQHBoxLayout: public QHBoxLayout {
public:
};

void* QHBoxLayout_NewQHBoxLayout(){
	return new QHBoxLayout();
}

void* QHBoxLayout_NewQHBoxLayout2(void* parent){
	return new QHBoxLayout(static_cast<QWidget*>(parent));
}

void QHBoxLayout_DestroyQHBoxLayout(void* ptr){
	static_cast<QHBoxLayout*>(ptr)->~QHBoxLayout();
}

#include "qgraphicsobject.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QGraphicsObject>
#include "_cgo_export.h"

class MyQGraphicsObject: public QGraphicsObject {
public:
void Signal_EnabledChanged(){callbackQGraphicsObjectEnabledChanged(this->objectName().toUtf8().data());};
void Signal_OpacityChanged(){callbackQGraphicsObjectOpacityChanged(this->objectName().toUtf8().data());};
void Signal_ParentChanged(){callbackQGraphicsObjectParentChanged(this->objectName().toUtf8().data());};
void Signal_RotationChanged(){callbackQGraphicsObjectRotationChanged(this->objectName().toUtf8().data());};
void Signal_ScaleChanged(){callbackQGraphicsObjectScaleChanged(this->objectName().toUtf8().data());};
void Signal_VisibleChanged(){callbackQGraphicsObjectVisibleChanged(this->objectName().toUtf8().data());};
void Signal_XChanged(){callbackQGraphicsObjectXChanged(this->objectName().toUtf8().data());};
void Signal_YChanged(){callbackQGraphicsObjectYChanged(this->objectName().toUtf8().data());};
void Signal_ZChanged(){callbackQGraphicsObjectZChanged(this->objectName().toUtf8().data());};
};

void QGraphicsObject_ConnectEnabledChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::enabledChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_EnabledChanged));;
}

void QGraphicsObject_DisconnectEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::enabledChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_EnabledChanged));;
}

void QGraphicsObject_GrabGesture(void* ptr, int gesture, int flags){
	static_cast<QGraphicsObject*>(ptr)->grabGesture(static_cast<Qt::GestureType>(gesture), static_cast<Qt::GestureFlag>(flags));
}

void QGraphicsObject_ConnectOpacityChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::opacityChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_OpacityChanged));;
}

void QGraphicsObject_DisconnectOpacityChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::opacityChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_OpacityChanged));;
}

void QGraphicsObject_ConnectParentChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::parentChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ParentChanged));;
}

void QGraphicsObject_DisconnectParentChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::parentChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ParentChanged));;
}

void QGraphicsObject_ConnectRotationChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::rotationChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_RotationChanged));;
}

void QGraphicsObject_DisconnectRotationChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::rotationChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_RotationChanged));;
}

void QGraphicsObject_ConnectScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::scaleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ScaleChanged));;
}

void QGraphicsObject_DisconnectScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::scaleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ScaleChanged));;
}

void QGraphicsObject_UngrabGesture(void* ptr, int gesture){
	static_cast<QGraphicsObject*>(ptr)->ungrabGesture(static_cast<Qt::GestureType>(gesture));
}

void QGraphicsObject_ConnectVisibleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::visibleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_VisibleChanged));;
}

void QGraphicsObject_DisconnectVisibleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::visibleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_VisibleChanged));;
}

void QGraphicsObject_ConnectXChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::xChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_XChanged));;
}

void QGraphicsObject_DisconnectXChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::xChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_XChanged));;
}

void QGraphicsObject_ConnectYChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::yChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_YChanged));;
}

void QGraphicsObject_DisconnectYChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::yChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_YChanged));;
}

void QGraphicsObject_ConnectZChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::zChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ZChanged));;
}

void QGraphicsObject_DisconnectZChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::zChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ZChanged));;
}

void QGraphicsObject_DestroyQGraphicsObject(void* ptr){
	static_cast<QGraphicsObject*>(ptr)->~QGraphicsObject();
}

#include "qfiledialog.h"
#include <QString>
#include <QUrl>
#include <QWidget>
#include <QFileIconProvider>
#include <QAbstractProxyModel>
#include <QByteArray>
#include <QFile>
#include <QVariant>
#include <QModelIndex>
#include <QObject>
#include <QStringList>
#include <QAbstractItemDelegate>
#include <QDir>
#include <QFileDialog>
#include "_cgo_export.h"

class MyQFileDialog: public QFileDialog {
public:
void Signal_CurrentChanged(const QString & path){callbackQFileDialogCurrentChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_DirectoryEntered(const QString & directory){callbackQFileDialogDirectoryEntered(this->objectName().toUtf8().data(), directory.toUtf8().data());};
void Signal_FileSelected(const QString & file){callbackQFileDialogFileSelected(this->objectName().toUtf8().data(), file.toUtf8().data());};
void Signal_FilesSelected(const QStringList & selected){callbackQFileDialogFilesSelected(this->objectName().toUtf8().data(), selected.join("|").toUtf8().data());};
void Signal_FilterSelected(const QString & filter){callbackQFileDialogFilterSelected(this->objectName().toUtf8().data(), filter.toUtf8().data());};
};

void* QFileDialog_NewQFileDialog(void* parent, int flags){
	return new QFileDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QFileDialog_AcceptMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->acceptMode();
}

int QFileDialog_ConfirmOverwrite(void* ptr){
	return static_cast<QFileDialog*>(ptr)->confirmOverwrite();
}

char* QFileDialog_DefaultSuffix(void* ptr){
	return static_cast<QFileDialog*>(ptr)->defaultSuffix().toUtf8().data();
}

int QFileDialog_FileMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->fileMode();
}

int QFileDialog_IsNameFilterDetailsVisible(void* ptr){
	return static_cast<QFileDialog*>(ptr)->isNameFilterDetailsVisible();
}

int QFileDialog_IsReadOnly(void* ptr){
	return static_cast<QFileDialog*>(ptr)->isReadOnly();
}

int QFileDialog_Options(void* ptr){
	return static_cast<QFileDialog*>(ptr)->options();
}

int QFileDialog_ResolveSymlinks(void* ptr){
	return static_cast<QFileDialog*>(ptr)->resolveSymlinks();
}

void QFileDialog_SetAcceptMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setAcceptMode(static_cast<QFileDialog::AcceptMode>(mode));
}

void QFileDialog_SetConfirmOverwrite(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setConfirmOverwrite(enabled != 0);
}

void QFileDialog_SetDefaultSuffix(void* ptr, char* suffix){
	static_cast<QFileDialog*>(ptr)->setDefaultSuffix(QString(suffix));
}

void QFileDialog_SetFileMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setFileMode(static_cast<QFileDialog::FileMode>(mode));
}

void QFileDialog_SetNameFilterDetailsVisible(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setNameFilterDetailsVisible(enabled != 0);
}

void QFileDialog_SetOptions(void* ptr, int options){
	static_cast<QFileDialog*>(ptr)->setOptions(static_cast<QFileDialog::Option>(options));
}

void QFileDialog_SetReadOnly(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setReadOnly(enabled != 0);
}

void QFileDialog_SetResolveSymlinks(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setResolveSymlinks(enabled != 0);
}

void QFileDialog_SetViewMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setViewMode(static_cast<QFileDialog::ViewMode>(mode));
}

int QFileDialog_ViewMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->viewMode();
}

void* QFileDialog_NewQFileDialog2(void* parent, char* caption, char* directory, char* filter){
	return new QFileDialog(static_cast<QWidget*>(parent), QString(caption), QString(directory), QString(filter));
}

void QFileDialog_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void QFileDialog_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void* QFileDialog_Directory(void* ptr){
	return new QDir(static_cast<QFileDialog*>(ptr)->directory());
}

void QFileDialog_ConnectDirectoryEntered(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

void QFileDialog_DisconnectDirectoryEntered(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

void QFileDialog_ConnectFileSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_DisconnectFileSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_ConnectFilesSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

void QFileDialog_DisconnectFilesSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

int QFileDialog_Filter(void* ptr){
	return static_cast<QFileDialog*>(ptr)->filter();
}

void QFileDialog_ConnectFilterSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

void QFileDialog_DisconnectFilterSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

char* QFileDialog_QFileDialog_GetExistingDirectory(void* parent, char* caption, char* dir, int options){
	return QFileDialog::getExistingDirectory(static_cast<QWidget*>(parent), QString(caption), QString(dir), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileNames(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileNames(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).join("|").toUtf8().data();
}

char* QFileDialog_QFileDialog_GetSaveFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getSaveFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_History(void* ptr){
	return static_cast<QFileDialog*>(ptr)->history().join("|").toUtf8().data();
}

void* QFileDialog_IconProvider(void* ptr){
	return static_cast<QFileDialog*>(ptr)->iconProvider();
}

void* QFileDialog_ItemDelegate(void* ptr){
	return static_cast<QFileDialog*>(ptr)->itemDelegate();
}

char* QFileDialog_LabelText(void* ptr, int label){
	return static_cast<QFileDialog*>(ptr)->labelText(static_cast<QFileDialog::DialogLabel>(label)).toUtf8().data();
}

char* QFileDialog_MimeTypeFilters(void* ptr){
	return static_cast<QFileDialog*>(ptr)->mimeTypeFilters().join("|").toUtf8().data();
}

char* QFileDialog_NameFilters(void* ptr){
	return static_cast<QFileDialog*>(ptr)->nameFilters().join("|").toUtf8().data();
}

void QFileDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QFileDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QFileDialog_ProxyModel(void* ptr){
	return static_cast<QFileDialog*>(ptr)->proxyModel();
}

int QFileDialog_RestoreState(void* ptr, void* state){
	return static_cast<QFileDialog*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void* QFileDialog_SaveState(void* ptr){
	return new QByteArray(static_cast<QFileDialog*>(ptr)->saveState());
}

void QFileDialog_SelectFile(void* ptr, char* filename){
	static_cast<QFileDialog*>(ptr)->selectFile(QString(filename));
}

void QFileDialog_SelectMimeTypeFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectMimeTypeFilter(QString(filter));
}

void QFileDialog_SelectNameFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectNameFilter(QString(filter));
}

void QFileDialog_SelectUrl(void* ptr, void* url){
	static_cast<QFileDialog*>(ptr)->selectUrl(*static_cast<QUrl*>(url));
}

char* QFileDialog_SelectedFiles(void* ptr){
	return static_cast<QFileDialog*>(ptr)->selectedFiles().join("|").toUtf8().data();
}

char* QFileDialog_SelectedNameFilter(void* ptr){
	return static_cast<QFileDialog*>(ptr)->selectedNameFilter().toUtf8().data();
}

void QFileDialog_SetDirectory2(void* ptr, void* directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(*static_cast<QDir*>(directory));
}

void QFileDialog_SetDirectory(void* ptr, char* directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(QString(directory));
}

void QFileDialog_SetDirectoryUrl(void* ptr, void* directory){
	static_cast<QFileDialog*>(ptr)->setDirectoryUrl(*static_cast<QUrl*>(directory));
}

void QFileDialog_SetFilter(void* ptr, int filters){
	static_cast<QFileDialog*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileDialog_SetHistory(void* ptr, char* paths){
	static_cast<QFileDialog*>(ptr)->setHistory(QString(paths).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetIconProvider(void* ptr, void* provider){
	static_cast<QFileDialog*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileDialog_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QFileDialog*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QFileDialog_SetLabelText(void* ptr, int label, char* text){
	static_cast<QFileDialog*>(ptr)->setLabelText(static_cast<QFileDialog::DialogLabel>(label), QString(text));
}

void QFileDialog_SetMimeTypeFilters(void* ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setMimeTypeFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetNameFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->setNameFilter(QString(filter));
}

void QFileDialog_SetNameFilters(void* ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetOption(void* ptr, int option, int on){
	static_cast<QFileDialog*>(ptr)->setOption(static_cast<QFileDialog::Option>(option), on != 0);
}

void QFileDialog_SetProxyModel(void* ptr, void* proxyModel){
	static_cast<QFileDialog*>(ptr)->setProxyModel(static_cast<QAbstractProxyModel*>(proxyModel));
}

void QFileDialog_SetVisible(void* ptr, int visible){
	static_cast<QFileDialog*>(ptr)->setVisible(visible != 0);
}

int QFileDialog_TestOption(void* ptr, int option){
	return static_cast<QFileDialog*>(ptr)->testOption(static_cast<QFileDialog::Option>(option));
}

void QFileDialog_DestroyQFileDialog(void* ptr){
	static_cast<QFileDialog*>(ptr)->~QFileDialog();
}

#include "qtextedit.h"
#include <QString>
#include <QUrl>
#include <QFont>
#include <QTextCharFormat>
#include <QPoint>
#include <QTextCursor>
#include <QObject>
#include <QColor>
#include <QPagedPaintDevice>
#include <QVariant>
#include <QWidget>
#include <QModelIndex>
#include <QMetaObject>
#include <QTextDocument>
#include <QTextOption>
#include <QTextEdit>
#include "_cgo_export.h"

class MyQTextEdit: public QTextEdit {
public:
void Signal_CopyAvailable(bool yes){callbackQTextEditCopyAvailable(this->objectName().toUtf8().data(), yes);};
void Signal_CursorPositionChanged(){callbackQTextEditCursorPositionChanged(this->objectName().toUtf8().data());};
void Signal_RedoAvailable(bool available){callbackQTextEditRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_SelectionChanged(){callbackQTextEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(){callbackQTextEditTextChanged(this->objectName().toUtf8().data());};
void Signal_UndoAvailable(bool available){callbackQTextEditUndoAvailable(this->objectName().toUtf8().data(), available);};
};

int QTextEdit_AcceptRichText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->acceptRichText();
}

int QTextEdit_AutoFormatting(void* ptr){
	return static_cast<QTextEdit*>(ptr)->autoFormatting();
}

int QTextEdit_CursorWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->cursorWidth();
}

void* QTextEdit_Document(void* ptr){
	return static_cast<QTextEdit*>(ptr)->document();
}

int QTextEdit_IsReadOnly(void* ptr){
	return static_cast<QTextEdit*>(ptr)->isReadOnly();
}

int QTextEdit_LineWrapColumnOrWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapColumnOrWidth();
}

int QTextEdit_LineWrapMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->lineWrapMode();
}

int QTextEdit_OverwriteMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->overwriteMode();
}

char* QTextEdit_PlaceholderText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QTextEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "redo");
}

void QTextEdit_SetAcceptRichText(void* ptr, int accept){
	static_cast<QTextEdit*>(ptr)->setAcceptRichText(accept != 0);
}

void QTextEdit_SetAutoFormatting(void* ptr, int features){
	static_cast<QTextEdit*>(ptr)->setAutoFormatting(static_cast<QTextEdit::AutoFormattingFlag>(features));
}

void QTextEdit_SetCursorWidth(void* ptr, int width){
	static_cast<QTextEdit*>(ptr)->setCursorWidth(width);
}

void QTextEdit_SetDocument(void* ptr, void* document){
	static_cast<QTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QTextEdit_SetFontWeight(void* ptr, int weight){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontWeight", Q_ARG(int, weight));
}

void QTextEdit_SetHtml(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetLineWrapColumnOrWidth(void* ptr, int w){
	static_cast<QTextEdit*>(ptr)->setLineWrapColumnOrWidth(w);
}

void QTextEdit_SetLineWrapMode(void* ptr, int mode){
	static_cast<QTextEdit*>(ptr)->setLineWrapMode(static_cast<QTextEdit::LineWrapMode>(mode));
}

void QTextEdit_SetOverwriteMode(void* ptr, int overwrite){
	static_cast<QTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QTextEdit_SetPlaceholderText(void* ptr, char* placeholderText){
	static_cast<QTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QTextEdit_SetReadOnly(void* ptr, int ro){
	static_cast<QTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QTextEdit_SetTabChangesFocus(void* ptr, int b){
	static_cast<QTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QTextEdit_SetTabStopWidth(void* ptr, int width){
	static_cast<QTextEdit*>(ptr)->setTabStopWidth(width);
}

void QTextEdit_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QTextEdit_SetWordWrapMode(void* ptr, int policy){
	static_cast<QTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QTextEdit_TabChangesFocus(void* ptr){
	return static_cast<QTextEdit*>(ptr)->tabChangesFocus();
}

int QTextEdit_TabStopWidth(void* ptr){
	return static_cast<QTextEdit*>(ptr)->tabStopWidth();
}

int QTextEdit_TextInteractionFlags(void* ptr){
	return static_cast<QTextEdit*>(ptr)->textInteractionFlags();
}

char* QTextEdit_ToHtml(void* ptr){
	return static_cast<QTextEdit*>(ptr)->toHtml().toUtf8().data();
}

int QTextEdit_WordWrapMode(void* ptr){
	return static_cast<QTextEdit*>(ptr)->wordWrapMode();
}

void QTextEdit_ZoomIn(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QTextEdit_ZoomOut(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

void* QTextEdit_NewQTextEdit(void* parent){
	return new QTextEdit(static_cast<QWidget*>(parent));
}

void* QTextEdit_NewQTextEdit2(char* text, void* parent){
	return new QTextEdit(QString(text), static_cast<QWidget*>(parent));
}

int QTextEdit_Alignment(void* ptr){
	return static_cast<QTextEdit*>(ptr)->alignment();
}

char* QTextEdit_AnchorAt(void* ptr, void* pos){
	return static_cast<QTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QTextEdit_Append(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "append", Q_ARG(QString, QString(text)));
}

int QTextEdit_CanPaste(void* ptr){
	return static_cast<QTextEdit*>(ptr)->canPaste();
}

void QTextEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "clear");
}

void QTextEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "copy");
}

void QTextEdit_ConnectCopyAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

void QTextEdit_DisconnectCopyAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::copyAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_CopyAvailable));;
}

void* QTextEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu();
}

void* QTextEdit_CreateStandardContextMenu2(void* ptr, void* position){
	return static_cast<QTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QTextEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::cursorPositionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_CursorPositionChanged));;
}

void QTextEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "cut");
}

char* QTextEdit_DocumentTitle(void* ptr){
	return static_cast<QTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QTextEdit_EnsureCursorVisible(void* ptr){
	static_cast<QTextEdit*>(ptr)->ensureCursorVisible();
}

char* QTextEdit_FontFamily(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontFamily().toUtf8().data();
}

int QTextEdit_FontItalic(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontItalic();
}

double QTextEdit_FontPointSize(void* ptr){
	return static_cast<double>(static_cast<QTextEdit*>(ptr)->fontPointSize());
}

int QTextEdit_FontUnderline(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontUnderline();
}

int QTextEdit_FontWeight(void* ptr){
	return static_cast<QTextEdit*>(ptr)->fontWeight();
}

void* QTextEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QTextEdit_InsertHtml(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertHtml", Q_ARG(QString, QString(text)));
}

void QTextEdit_InsertPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QTextEdit_IsUndoRedoEnabled(void* ptr){
	return static_cast<QTextEdit*>(ptr)->isUndoRedoEnabled();
}

void* QTextEdit_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QTextEdit*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

void QTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier){
	static_cast<QTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextEdit_MoveCursor(void* ptr, int operation, int mode){
	static_cast<QTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QTextEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "paste");
}

void QTextEdit_Print(void* ptr, void* printer){
	static_cast<QTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextEdit_ConnectRedoAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_DisconnectRedoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::redoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_RedoAvailable));;
}

void QTextEdit_ScrollToAnchor(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "scrollToAnchor", Q_ARG(QString, QString(name)));
}

void QTextEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "selectAll");
}

void QTextEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::selectionChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_SelectionChanged));;
}

void QTextEdit_SetAlignment(void* ptr, int a){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setAlignment", Q_ARG(Qt::AlignmentFlag, static_cast<Qt::AlignmentFlag>(a)));
}

void QTextEdit_SetCurrentCharFormat(void* ptr, void* format){
	static_cast<QTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextEdit_SetCurrentFont(void* ptr, void* f){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setCurrentFont", Q_ARG(QFont, *static_cast<QFont*>(f)));
}

void QTextEdit_SetDocumentTitle(void* ptr, char* title){
	static_cast<QTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QTextEdit_SetFontFamily(void* ptr, char* fontFamily){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontFamily", Q_ARG(QString, QString(fontFamily)));
}

void QTextEdit_SetFontItalic(void* ptr, int italic){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontItalic", Q_ARG(bool, italic != 0));
}

void QTextEdit_SetFontPointSize(void* ptr, double s){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontPointSize", Q_ARG(qreal, static_cast<qreal>(s)));
}

void QTextEdit_SetFontUnderline(void* ptr, int underline){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setFontUnderline", Q_ARG(bool, underline != 0));
}

void QTextEdit_SetPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setText", Q_ARG(QString, QString(text)));
}

void QTextEdit_SetTextBackgroundColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextBackgroundColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "setTextColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QTextEdit_SetTextCursor(void* ptr, void* cursor){
	static_cast<QTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QTextEdit_SetUndoRedoEnabled(void* ptr, int enable){
	static_cast<QTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void* QTextEdit_TextBackgroundColor(void* ptr){
	return new QColor(static_cast<QTextEdit*>(ptr)->textBackgroundColor());
}

void QTextEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

void QTextEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)()>(&QTextEdit::textChanged), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)()>(&MyQTextEdit::Signal_TextChanged));;
}

void* QTextEdit_TextColor(void* ptr){
	return new QColor(static_cast<QTextEdit*>(ptr)->textColor());
}

char* QTextEdit_ToPlainText(void* ptr){
	return static_cast<QTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QTextEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextEdit*>(ptr), "undo");
}

void QTextEdit_ConnectUndoAvailable(void* ptr){
	QObject::connect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DisconnectUndoAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextEdit*>(ptr), static_cast<void (QTextEdit::*)(bool)>(&QTextEdit::undoAvailable), static_cast<MyQTextEdit*>(ptr), static_cast<void (MyQTextEdit::*)(bool)>(&MyQTextEdit::Signal_UndoAvailable));;
}

void QTextEdit_DestroyQTextEdit(void* ptr){
	static_cast<QTextEdit*>(ptr)->~QTextEdit();
}

#include "qaccessiblewidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QAccessible>
#include <QAccessibleInterface>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QAccessibleWidget>
#include "_cgo_export.h"

class MyQAccessibleWidget: public QAccessibleWidget {
public:
};

void* QAccessibleWidget_NewQAccessibleWidget(void* w, int role, char* name){
	return new QAccessibleWidget(static_cast<QWidget*>(w), static_cast<QAccessible::Role>(role), QString(name));
}

char* QAccessibleWidget_ActionNames(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->actionNames().join("|").toUtf8().data();
}

void* QAccessibleWidget_BackgroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleWidget*>(ptr)->backgroundColor());
}

void* QAccessibleWidget_Child(void* ptr, int index){
	return static_cast<QAccessibleWidget*>(ptr)->child(index);
}

int QAccessibleWidget_ChildCount(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->childCount();
}

void QAccessibleWidget_DoAction(void* ptr, char* actionName){
	static_cast<QAccessibleWidget*>(ptr)->doAction(QString(actionName));
}

void* QAccessibleWidget_FocusChild(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->focusChild();
}

void* QAccessibleWidget_ForegroundColor(void* ptr){
	return new QColor(static_cast<QAccessibleWidget*>(ptr)->foregroundColor());
}

int QAccessibleWidget_IndexOfChild(void* ptr, void* child){
	return static_cast<QAccessibleWidget*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void* QAccessibleWidget_Interface_cast(void* ptr, int t){
	return static_cast<QAccessibleWidget*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(t));
}

int QAccessibleWidget_IsValid(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->isValid();
}

char* QAccessibleWidget_KeyBindingsForAction(void* ptr, char* actionName){
	return static_cast<QAccessibleWidget*>(ptr)->keyBindingsForAction(QString(actionName)).join("|").toUtf8().data();
}

void* QAccessibleWidget_Parent(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->parent();
}

int QAccessibleWidget_Role(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->role();
}

char* QAccessibleWidget_Text(void* ptr, int t){
	return static_cast<QAccessibleWidget*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

void* QAccessibleWidget_Window(void* ptr){
	return static_cast<QAccessibleWidget*>(ptr)->window();
}

#include "qgraphicsblureffect.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsBlurEffect>
#include "_cgo_export.h"

class MyQGraphicsBlurEffect: public QGraphicsBlurEffect {
public:
void Signal_BlurHintsChanged(QGraphicsBlurEffect::BlurHints hints){callbackQGraphicsBlurEffectBlurHintsChanged(this->objectName().toUtf8().data(), hints);};
};

int QGraphicsBlurEffect_BlurHints(void* ptr){
	return static_cast<QGraphicsBlurEffect*>(ptr)->blurHints();
}

double QGraphicsBlurEffect_BlurRadius(void* ptr){
	return static_cast<double>(static_cast<QGraphicsBlurEffect*>(ptr)->blurRadius());
}

void QGraphicsBlurEffect_SetBlurHints(void* ptr, int hints){
	QMetaObject::invokeMethod(static_cast<QGraphicsBlurEffect*>(ptr), "setBlurHints", Q_ARG(QGraphicsBlurEffect::BlurHint, static_cast<QGraphicsBlurEffect::BlurHint>(hints)));
}

void QGraphicsBlurEffect_SetBlurRadius(void* ptr, double blurRadius){
	QMetaObject::invokeMethod(static_cast<QGraphicsBlurEffect*>(ptr), "setBlurRadius", Q_ARG(qreal, static_cast<qreal>(blurRadius)));
}

void* QGraphicsBlurEffect_NewQGraphicsBlurEffect(void* parent){
	return new QGraphicsBlurEffect(static_cast<QObject*>(parent));
}

void QGraphicsBlurEffect_ConnectBlurHintsChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DisconnectBlurHintsChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsBlurEffect*>(ptr), static_cast<void (QGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&QGraphicsBlurEffect::blurHintsChanged), static_cast<MyQGraphicsBlurEffect*>(ptr), static_cast<void (MyQGraphicsBlurEffect::*)(QGraphicsBlurEffect::BlurHints)>(&MyQGraphicsBlurEffect::Signal_BlurHintsChanged));;
}

void QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(void* ptr){
	static_cast<QGraphicsBlurEffect*>(ptr)->~QGraphicsBlurEffect();
}

#include "qstatusbar.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QStatusBar>
#include "_cgo_export.h"

class MyQStatusBar: public QStatusBar {
public:
void Signal_MessageChanged(const QString & message){callbackQStatusBarMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

int QStatusBar_IsSizeGripEnabled(void* ptr){
	return static_cast<QStatusBar*>(ptr)->isSizeGripEnabled();
}

void QStatusBar_SetSizeGripEnabled(void* ptr, int v){
	static_cast<QStatusBar*>(ptr)->setSizeGripEnabled(v != 0);
}

void* QStatusBar_NewQStatusBar(void* parent){
	return new QStatusBar(static_cast<QWidget*>(parent));
}

void QStatusBar_AddPermanentWidget(void* ptr, void* widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addPermanentWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_AddWidget(void* ptr, void* widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ClearMessage(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "clearMessage");
}

char* QStatusBar_CurrentMessage(void* ptr){
	return static_cast<QStatusBar*>(ptr)->currentMessage().toUtf8().data();
}

int QStatusBar_InsertPermanentWidget(void* ptr, int index, void* widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertPermanentWidget(index, static_cast<QWidget*>(widget), stretch);
}

int QStatusBar_InsertWidget(void* ptr, int index, void* widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ConnectMessageChanged(void* ptr){
	QObject::connect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_DisconnectMessageChanged(void* ptr){
	QObject::disconnect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_RemoveWidget(void* ptr, void* widget){
	static_cast<QStatusBar*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void QStatusBar_ShowMessage(void* ptr, char* message, int timeout){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, timeout));
}

void QStatusBar_DestroyQStatusBar(void* ptr){
	static_cast<QStatusBar*>(ptr)->~QStatusBar();
}

#include "qtoolbutton.h"
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QAction>
#include <QMenu>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QVariant>
#include <QToolButton>
#include "_cgo_export.h"

class MyQToolButton: public QToolButton {
public:
void Signal_Triggered(QAction * action){callbackQToolButtonTriggered(this->objectName().toUtf8().data(), action);};
};

int QToolButton_ArrowType(void* ptr){
	return static_cast<QToolButton*>(ptr)->arrowType();
}

int QToolButton_AutoRaise(void* ptr){
	return static_cast<QToolButton*>(ptr)->autoRaise();
}

int QToolButton_PopupMode(void* ptr){
	return static_cast<QToolButton*>(ptr)->popupMode();
}

void QToolButton_SetArrowType(void* ptr, int ty){
	static_cast<QToolButton*>(ptr)->setArrowType(static_cast<Qt::ArrowType>(ty));
}

void QToolButton_SetAutoRaise(void* ptr, int enable){
	static_cast<QToolButton*>(ptr)->setAutoRaise(enable != 0);
}

void QToolButton_SetPopupMode(void* ptr, int mode){
	static_cast<QToolButton*>(ptr)->setPopupMode(static_cast<QToolButton::ToolButtonPopupMode>(mode));
}

void QToolButton_SetToolButtonStyle(void* ptr, int style){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(style)));
}

int QToolButton_ToolButtonStyle(void* ptr){
	return static_cast<QToolButton*>(ptr)->toolButtonStyle();
}

void* QToolButton_NewQToolButton(void* parent){
	return new QToolButton(static_cast<QWidget*>(parent));
}

void* QToolButton_Menu(void* ptr){
	return static_cast<QToolButton*>(ptr)->menu();
}

void QToolButton_SetMenu(void* ptr, void* menu){
	static_cast<QToolButton*>(ptr)->setMenu(static_cast<QMenu*>(menu));
}

void QToolButton_ShowMenu(void* ptr){
	QMetaObject::invokeMethod(static_cast<QToolButton*>(ptr), "showMenu");
}

void QToolButton_ConnectTriggered(void* ptr){
	QObject::connect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DisconnectTriggered(void* ptr){
	QObject::disconnect(static_cast<QToolButton*>(ptr), static_cast<void (QToolButton::*)(QAction *)>(&QToolButton::triggered), static_cast<MyQToolButton*>(ptr), static_cast<void (MyQToolButton::*)(QAction *)>(&MyQToolButton::Signal_Triggered));;
}

void QToolButton_DestroyQToolButton(void* ptr){
	static_cast<QToolButton*>(ptr)->~QToolButton();
}

#include "qstyleoptionprogressbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionProgressBar>
#include "_cgo_export.h"

class MyQStyleOptionProgressBar: public QStyleOptionProgressBar {
public:
};

void* QStyleOptionProgressBar_NewQStyleOptionProgressBar(){
	return new QStyleOptionProgressBar();
}

void* QStyleOptionProgressBar_NewQStyleOptionProgressBar2(void* other){
	return new QStyleOptionProgressBar(*static_cast<QStyleOptionProgressBar*>(other));
}

#include "qgesture.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QObject>
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QGesture>
#include "_cgo_export.h"

class MyQGesture: public QGesture {
public:
};

int QGesture_GestureCancelPolicy(void* ptr){
	return static_cast<QGesture*>(ptr)->gestureCancelPolicy();
}

int QGesture_GestureType(void* ptr){
	return static_cast<QGesture*>(ptr)->gestureType();
}

int QGesture_HasHotSpot(void* ptr){
	return static_cast<QGesture*>(ptr)->hasHotSpot();
}

void QGesture_SetGestureCancelPolicy(void* ptr, int policy){
	static_cast<QGesture*>(ptr)->setGestureCancelPolicy(static_cast<QGesture::GestureCancelPolicy>(policy));
}

void QGesture_SetHotSpot(void* ptr, void* value){
	static_cast<QGesture*>(ptr)->setHotSpot(*static_cast<QPointF*>(value));
}

void QGesture_UnsetHotSpot(void* ptr){
	static_cast<QGesture*>(ptr)->unsetHotSpot();
}

void* QGesture_NewQGesture(void* parent){
	return new QGesture(static_cast<QObject*>(parent));
}

void QGesture_DestroyQGesture(void* ptr){
	static_cast<QGesture*>(ptr)->~QGesture();
}

#include "qlayoutitem.h"
#include <QModelIndex>
#include <QRect>
#include <QLayout>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLayoutItem>
#include "_cgo_export.h"

class MyQLayoutItem: public QLayoutItem {
public:
};

int QLayoutItem_Alignment(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->alignment();
}

int QLayoutItem_ControlTypes(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->controlTypes();
}

int QLayoutItem_ExpandingDirections(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->expandingDirections();
}

int QLayoutItem_HasHeightForWidth(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->hasHeightForWidth();
}

int QLayoutItem_HeightForWidth(void* ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->heightForWidth(w);
}

void QLayoutItem_Invalidate(void* ptr){
	static_cast<QLayoutItem*>(ptr)->invalidate();
}

int QLayoutItem_IsEmpty(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->isEmpty();
}

void* QLayoutItem_Layout(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->layout();
}

int QLayoutItem_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QLayoutItem*>(ptr)->minimumHeightForWidth(w);
}

void QLayoutItem_SetAlignment(void* ptr, int alignment){
	static_cast<QLayoutItem*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayoutItem_SetGeometry(void* ptr, void* r){
	static_cast<QLayoutItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void* QLayoutItem_SpacerItem(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->spacerItem();
}

void* QLayoutItem_Widget(void* ptr){
	return static_cast<QLayoutItem*>(ptr)->widget();
}

void QLayoutItem_DestroyQLayoutItem(void* ptr){
	static_cast<QLayoutItem*>(ptr)->~QLayoutItem();
}

#include "qstylepainter.h"
#include <QString>
#include <QVariant>
#include <QStyle>
#include <QPalette>
#include <QPixmap>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionComplex>
#include <QWidget>
#include <QStyleOption>
#include <QPaintDevice>
#include <QRect>
#include <QStylePainter>
#include "_cgo_export.h"

class MyQStylePainter: public QStylePainter {
public:
};

void* QStylePainter_NewQStylePainter(){
	return new QStylePainter();
}

void* QStylePainter_NewQStylePainter3(void* pd, void* widget){
	return new QStylePainter(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

void* QStylePainter_NewQStylePainter2(void* widget){
	return new QStylePainter(static_cast<QWidget*>(widget));
}

int QStylePainter_Begin2(void* ptr, void* pd, void* widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QPaintDevice*>(pd), static_cast<QWidget*>(widget));
}

int QStylePainter_Begin(void* ptr, void* widget){
	return static_cast<QStylePainter*>(ptr)->begin(static_cast<QWidget*>(widget));
}

void QStylePainter_DrawComplexControl(void* ptr, int cc, void* option){
	static_cast<QStylePainter*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), *static_cast<QStyleOptionComplex*>(option));
}

void QStylePainter_DrawControl(void* ptr, int ce, void* option){
	static_cast<QStylePainter*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(ce), *static_cast<QStyleOption*>(option));
}

void QStylePainter_DrawItemPixmap(void* ptr, void* rect, int flags, void* pixmap){
	static_cast<QStylePainter*>(ptr)->drawItemPixmap(*static_cast<QRect*>(rect), flags, *static_cast<QPixmap*>(pixmap));
}

void QStylePainter_DrawItemText(void* ptr, void* rect, int flags, void* pal, int enabled, char* text, int textRole){
	static_cast<QStylePainter*>(ptr)->drawItemText(*static_cast<QRect*>(rect), flags, *static_cast<QPalette*>(pal), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QStylePainter_DrawPrimitive(void* ptr, int pe, void* option){
	static_cast<QStylePainter*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), *static_cast<QStyleOption*>(option));
}

void* QStylePainter_Style(void* ptr){
	return static_cast<QStylePainter*>(ptr)->style();
}

#include "qgestureevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGesture>
#include <QGestureEvent>
#include "_cgo_export.h"

class MyQGestureEvent: public QGestureEvent {
public:
};

void QGestureEvent_Accept(void* ptr, void* gesture){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Accept2(void* ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<Qt::GestureType>(gestureType));
}

void* QGestureEvent_Gesture(void* ptr, int ty){
	return static_cast<QGestureEvent*>(ptr)->gesture(static_cast<Qt::GestureType>(ty));
}

void QGestureEvent_Ignore(void* ptr, void* gesture){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Ignore2(void* ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<Qt::GestureType>(gestureType));
}

int QGestureEvent_IsAccepted(void* ptr, void* gesture){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<QGesture*>(gesture));
}

int QGestureEvent_IsAccepted2(void* ptr, int gestureType){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<Qt::GestureType>(gestureType));
}

void QGestureEvent_SetAccepted(void* ptr, void* gesture, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<QGesture*>(gesture), value != 0);
}

void QGestureEvent_SetAccepted2(void* ptr, int gestureType, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<Qt::GestureType>(gestureType), value != 0);
}

void* QGestureEvent_Widget(void* ptr){
	return static_cast<QGestureEvent*>(ptr)->widget();
}

void QGestureEvent_DestroyQGestureEvent(void* ptr){
	static_cast<QGestureEvent*>(ptr)->~QGestureEvent();
}

#include "qswipegesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSwipeGesture>
#include "_cgo_export.h"

class MyQSwipeGesture: public QSwipeGesture {
public:
};

int QSwipeGesture_HorizontalDirection(void* ptr){
	return static_cast<QSwipeGesture*>(ptr)->horizontalDirection();
}

void QSwipeGesture_SetSwipeAngle(void* ptr, double value){
	static_cast<QSwipeGesture*>(ptr)->setSwipeAngle(static_cast<qreal>(value));
}

double QSwipeGesture_SwipeAngle(void* ptr){
	return static_cast<double>(static_cast<QSwipeGesture*>(ptr)->swipeAngle());
}

int QSwipeGesture_VerticalDirection(void* ptr){
	return static_cast<QSwipeGesture*>(ptr)->verticalDirection();
}

void QSwipeGesture_DestroyQSwipeGesture(void* ptr){
	static_cast<QSwipeGesture*>(ptr)->~QSwipeGesture();
}

#include "qdesktopwidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QDesktopWidget>
#include "_cgo_export.h"

class MyQDesktopWidget: public QDesktopWidget {
public:
void Signal_Resized(int screen){callbackQDesktopWidgetResized(this->objectName().toUtf8().data(), screen);};
void Signal_ScreenCountChanged(int newCount){callbackQDesktopWidgetScreenCountChanged(this->objectName().toUtf8().data(), newCount);};
void Signal_WorkAreaResized(int screen){callbackQDesktopWidgetWorkAreaResized(this->objectName().toUtf8().data(), screen);};
};

int QDesktopWidget_IsVirtualDesktop(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->isVirtualDesktop();
}

int QDesktopWidget_PrimaryScreen(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->primaryScreen();
}

void* QDesktopWidget_Screen(void* ptr, int screen){
	return static_cast<QDesktopWidget*>(ptr)->screen(screen);
}

int QDesktopWidget_ScreenNumber2(void* ptr, void* point){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(*static_cast<QPoint*>(point));
}

int QDesktopWidget_ScreenNumber(void* ptr, void* widget){
	return static_cast<QDesktopWidget*>(ptr)->screenNumber(static_cast<QWidget*>(widget));
}

void QDesktopWidget_ConnectResized(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

void QDesktopWidget_DisconnectResized(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::resized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_Resized));;
}

int QDesktopWidget_ScreenCount(void* ptr){
	return static_cast<QDesktopWidget*>(ptr)->screenCount();
}

void QDesktopWidget_ConnectScreenCountChanged(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_DisconnectScreenCountChanged(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::screenCountChanged), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_ScreenCountChanged));;
}

void QDesktopWidget_ConnectWorkAreaResized(void* ptr){
	QObject::connect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

void QDesktopWidget_DisconnectWorkAreaResized(void* ptr){
	QObject::disconnect(static_cast<QDesktopWidget*>(ptr), static_cast<void (QDesktopWidget::*)(int)>(&QDesktopWidget::workAreaResized), static_cast<MyQDesktopWidget*>(ptr), static_cast<void (MyQDesktopWidget::*)(int)>(&MyQDesktopWidget::Signal_WorkAreaResized));;
}

#include "qlineedit.h"
#include <QEvent>
#include <QModelIndex>
#include <QLine>
#include <QWidget>
#include <QCompleter>
#include <QIcon>
#include <QValidator>
#include <QMargins>
#include <QString>
#include <QUrl>
#include <QPoint>
#include <QVariant>
#include <QMetaObject>
#include <QObject>
#include <QAction>
#include <QLineEdit>
#include "_cgo_export.h"

class MyQLineEdit: public QLineEdit {
public:
void Signal_CursorPositionChanged(int old, int n){callbackQLineEditCursorPositionChanged(this->objectName().toUtf8().data(), old, n);};
void Signal_EditingFinished(){callbackQLineEditEditingFinished(this->objectName().toUtf8().data());};
void Signal_ReturnPressed(){callbackQLineEditReturnPressed(this->objectName().toUtf8().data());};
void Signal_SelectionChanged(){callbackQLineEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(const QString & text){callbackQLineEditTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_TextEdited(const QString & text){callbackQLineEditTextEdited(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

int QLineEdit_Alignment(void* ptr){
	return static_cast<QLineEdit*>(ptr)->alignment();
}

int QLineEdit_CursorMoveStyle(void* ptr){
	return static_cast<QLineEdit*>(ptr)->cursorMoveStyle();
}

int QLineEdit_CursorPosition(void* ptr){
	return static_cast<QLineEdit*>(ptr)->cursorPosition();
}

char* QLineEdit_DisplayText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->displayText().toUtf8().data();
}

int QLineEdit_DragEnabled(void* ptr){
	return static_cast<QLineEdit*>(ptr)->dragEnabled();
}

int QLineEdit_EchoMode(void* ptr){
	return static_cast<QLineEdit*>(ptr)->echoMode();
}

int QLineEdit_HasAcceptableInput(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasAcceptableInput();
}

int QLineEdit_HasFrame(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasFrame();
}

int QLineEdit_HasSelectedText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->hasSelectedText();
}

char* QLineEdit_InputMask(void* ptr){
	return static_cast<QLineEdit*>(ptr)->inputMask().toUtf8().data();
}

int QLineEdit_IsClearButtonEnabled(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isClearButtonEnabled();
}

int QLineEdit_IsModified(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isModified();
}

int QLineEdit_IsReadOnly(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isReadOnly();
}

int QLineEdit_IsRedoAvailable(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isRedoAvailable();
}

int QLineEdit_IsUndoAvailable(void* ptr){
	return static_cast<QLineEdit*>(ptr)->isUndoAvailable();
}

int QLineEdit_MaxLength(void* ptr){
	return static_cast<QLineEdit*>(ptr)->maxLength();
}

char* QLineEdit_PlaceholderText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->placeholderText().toUtf8().data();
}

char* QLineEdit_SelectedText(void* ptr){
	return static_cast<QLineEdit*>(ptr)->selectedText().toUtf8().data();
}

void QLineEdit_SetAlignment(void* ptr, int flag){
	static_cast<QLineEdit*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QLineEdit_SetClearButtonEnabled(void* ptr, int enable){
	static_cast<QLineEdit*>(ptr)->setClearButtonEnabled(enable != 0);
}

void QLineEdit_SetCursorMoveStyle(void* ptr, int style){
	static_cast<QLineEdit*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QLineEdit_SetCursorPosition(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setCursorPosition(v);
}

void QLineEdit_SetDragEnabled(void* ptr, int b){
	static_cast<QLineEdit*>(ptr)->setDragEnabled(b != 0);
}

void QLineEdit_SetEchoMode(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setEchoMode(static_cast<QLineEdit::EchoMode>(v));
}

void QLineEdit_SetFrame(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setFrame(v != 0);
}

void QLineEdit_SetInputMask(void* ptr, char* inputMask){
	static_cast<QLineEdit*>(ptr)->setInputMask(QString(inputMask));
}

void QLineEdit_SetMaxLength(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setMaxLength(v);
}

void QLineEdit_SetModified(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setModified(v != 0);
}

void QLineEdit_SetPlaceholderText(void* ptr, char* v){
	static_cast<QLineEdit*>(ptr)->setPlaceholderText(QString(v));
}

void QLineEdit_SetReadOnly(void* ptr, int v){
	static_cast<QLineEdit*>(ptr)->setReadOnly(v != 0);
}

void QLineEdit_SetText(void* ptr, char* v){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "setText", Q_ARG(QString, QString(v)));
}

char* QLineEdit_Text(void* ptr){
	return static_cast<QLineEdit*>(ptr)->text().toUtf8().data();
}

void* QLineEdit_NewQLineEdit(void* parent){
	return new QLineEdit(static_cast<QWidget*>(parent));
}

void* QLineEdit_NewQLineEdit2(char* contents, void* parent){
	return new QLineEdit(QString(contents), static_cast<QWidget*>(parent));
}

void* QLineEdit_AddAction2(void* ptr, void* icon, int position){
	return static_cast<QLineEdit*>(ptr)->addAction(*static_cast<QIcon*>(icon), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_AddAction(void* ptr, void* action, int position){
	static_cast<QLineEdit*>(ptr)->addAction(static_cast<QAction*>(action), static_cast<QLineEdit::ActionPosition>(position));
}

void QLineEdit_Backspace(void* ptr){
	static_cast<QLineEdit*>(ptr)->backspace();
}

void QLineEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "clear");
}

void* QLineEdit_Completer(void* ptr){
	return static_cast<QLineEdit*>(ptr)->completer();
}

void QLineEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "copy");
}

void* QLineEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QLineEdit*>(ptr)->createStandardContextMenu();
}

void QLineEdit_CursorBackward(void* ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorBackward(mark != 0, steps);
}

void QLineEdit_CursorForward(void* ptr, int mark, int steps){
	static_cast<QLineEdit*>(ptr)->cursorForward(mark != 0, steps);
}

int QLineEdit_CursorPositionAt(void* ptr, void* pos){
	return static_cast<QLineEdit*>(ptr)->cursorPositionAt(*static_cast<QPoint*>(pos));
}

void QLineEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(int, int)>(&QLineEdit::cursorPositionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(int, int)>(&MyQLineEdit::Signal_CursorPositionChanged));;
}

void QLineEdit_CursorWordBackward(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordBackward(mark != 0);
}

void QLineEdit_CursorWordForward(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->cursorWordForward(mark != 0);
}

void QLineEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "cut");
}

void QLineEdit_Del(void* ptr){
	static_cast<QLineEdit*>(ptr)->del();
}

void QLineEdit_Deselect(void* ptr){
	static_cast<QLineEdit*>(ptr)->deselect();
}

void QLineEdit_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::editingFinished), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_EditingFinished));;
}

void QLineEdit_End(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->end(mark != 0);
}

int QLineEdit_Event(void* ptr, void* e){
	return static_cast<QLineEdit*>(ptr)->event(static_cast<QEvent*>(e));
}

void QLineEdit_GetTextMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->getTextMargins(&left, &top, &right, &bottom);
}

void QLineEdit_Home(void* ptr, int mark){
	static_cast<QLineEdit*>(ptr)->home(mark != 0);
}

void* QLineEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QLineEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QLineEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "paste");
}

void QLineEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "redo");
}

void QLineEdit_ConnectReturnPressed(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_DisconnectReturnPressed(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::returnPressed), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_ReturnPressed));;
}

void QLineEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "selectAll");
}

void QLineEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

void QLineEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)()>(&QLineEdit::selectionChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)()>(&MyQLineEdit::Signal_SelectionChanged));;
}

int QLineEdit_SelectionStart(void* ptr){
	return static_cast<QLineEdit*>(ptr)->selectionStart();
}

void QLineEdit_SetCompleter(void* ptr, void* c){
	static_cast<QLineEdit*>(ptr)->setCompleter(static_cast<QCompleter*>(c));
}

void QLineEdit_SetSelection(void* ptr, int start, int length){
	static_cast<QLineEdit*>(ptr)->setSelection(start, length);
}

void QLineEdit_SetTextMargins2(void* ptr, void* margins){
	static_cast<QLineEdit*>(ptr)->setTextMargins(*static_cast<QMargins*>(margins));
}

void QLineEdit_SetTextMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLineEdit*>(ptr)->setTextMargins(left, top, right, bottom);
}

void QLineEdit_SetValidator(void* ptr, void* v){
	static_cast<QLineEdit*>(ptr)->setValidator(static_cast<QValidator*>(v));
}

void QLineEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textChanged), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextChanged));;
}

void QLineEdit_ConnectTextEdited(void* ptr){
	QObject::connect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_DisconnectTextEdited(void* ptr){
	QObject::disconnect(static_cast<QLineEdit*>(ptr), static_cast<void (QLineEdit::*)(const QString &)>(&QLineEdit::textEdited), static_cast<MyQLineEdit*>(ptr), static_cast<void (MyQLineEdit::*)(const QString &)>(&MyQLineEdit::Signal_TextEdited));;
}

void QLineEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLineEdit*>(ptr), "undo");
}

void* QLineEdit_Validator(void* ptr){
	return const_cast<QValidator*>(static_cast<QLineEdit*>(ptr)->validator());
}

void QLineEdit_DestroyQLineEdit(void* ptr){
	static_cast<QLineEdit*>(ptr)->~QLineEdit();
}

#include "qcommonstyle.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QWidget>
#include <QSizePolicy>
#include <QStyleHintReturn>
#include <QPainter>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QApplication>
#include <QStyleOptionComplex>
#include <QSize>
#include <QString>
#include <QPalette>
#include <QStyle>
#include <QCommonStyle>
#include "_cgo_export.h"

class MyQCommonStyle: public QCommonStyle {
public:
};

void QCommonStyle_DrawControl(void* ptr, int element, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawPrimitive(void* ptr, int pe, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(pe), static_cast<QStyleOption*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

void QCommonStyle_DrawComplexControl(void* ptr, int cc, void* opt, void* p, void* widget){
	static_cast<QCommonStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), static_cast<QPainter*>(p), static_cast<QWidget*>(widget));
}

int QCommonStyle_HitTestComplexControl(void* ptr, int cc, void* opt, void* pt, void* widget){
	return static_cast<QCommonStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(cc), static_cast<QStyleOptionComplex*>(opt), *static_cast<QPoint*>(pt), static_cast<QWidget*>(widget));
}

int QCommonStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget){
	return static_cast<QCommonStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QCommonStyle_PixelMetric(void* ptr, int m, void* opt, void* widget){
	return static_cast<QCommonStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(m), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget));
}

void QCommonStyle_Polish2(void* ptr, void* app){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QApplication*>(app));
}

void QCommonStyle_Polish(void* ptr, void* pal){
	static_cast<QCommonStyle*>(ptr)->polish(*static_cast<QPalette*>(pal));
}

void QCommonStyle_Polish3(void* ptr, void* widget){
	static_cast<QCommonStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

int QCommonStyle_StyleHint(void* ptr, int sh, void* opt, void* widget, void* hret){
	return static_cast<QCommonStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(sh), static_cast<QStyleOption*>(opt), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(hret));
}

void QCommonStyle_Unpolish2(void* ptr, void* application){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QApplication*>(application));
}

void QCommonStyle_Unpolish(void* ptr, void* widget){
	static_cast<QCommonStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

void QCommonStyle_DestroyQCommonStyle(void* ptr){
	static_cast<QCommonStyle*>(ptr)->~QCommonStyle();
}

#include "qabstractslider.h"
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractSlider>
#include "_cgo_export.h"

class MyQAbstractSlider: public QAbstractSlider {
public:
void Signal_ActionTriggered(int action){callbackQAbstractSliderActionTriggered(this->objectName().toUtf8().data(), action);};
void Signal_RangeChanged(int min, int max){callbackQAbstractSliderRangeChanged(this->objectName().toUtf8().data(), min, max);};
void Signal_SliderMoved(int value){callbackQAbstractSliderSliderMoved(this->objectName().toUtf8().data(), value);};
void Signal_SliderPressed(){callbackQAbstractSliderSliderPressed(this->objectName().toUtf8().data());};
void Signal_SliderReleased(){callbackQAbstractSliderSliderReleased(this->objectName().toUtf8().data());};
void Signal_ValueChanged(int value){callbackQAbstractSliderValueChanged(this->objectName().toUtf8().data(), value);};
};

int QAbstractSlider_HasTracking(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->hasTracking();
}

int QAbstractSlider_InvertedAppearance(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->invertedAppearance();
}

int QAbstractSlider_InvertedControls(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->invertedControls();
}

int QAbstractSlider_IsSliderDown(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->isSliderDown();
}

int QAbstractSlider_Maximum(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->maximum();
}

int QAbstractSlider_Minimum(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->minimum();
}

int QAbstractSlider_Orientation(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->orientation();
}

int QAbstractSlider_PageStep(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->pageStep();
}

void QAbstractSlider_SetInvertedAppearance(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setInvertedAppearance(v != 0);
}

void QAbstractSlider_SetInvertedControls(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setInvertedControls(v != 0);
}

void QAbstractSlider_SetMaximum(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setMaximum(v);
}

void QAbstractSlider_SetMinimum(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setMinimum(v);
}

void QAbstractSlider_SetOrientation(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(v)));
}

void QAbstractSlider_SetPageStep(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setPageStep(v);
}

void QAbstractSlider_SetSingleStep(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSingleStep(v);
}

void QAbstractSlider_SetSliderDown(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSliderDown(v != 0);
}

void QAbstractSlider_SetSliderPosition(void* ptr, int v){
	static_cast<QAbstractSlider*>(ptr)->setSliderPosition(v);
}

void QAbstractSlider_SetTracking(void* ptr, int enable){
	static_cast<QAbstractSlider*>(ptr)->setTracking(enable != 0);
}

void QAbstractSlider_SetValue(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setValue", Q_ARG(int, v));
}

int QAbstractSlider_SingleStep(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->singleStep();
}

int QAbstractSlider_SliderPosition(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->sliderPosition();
}

int QAbstractSlider_Value(void* ptr){
	return static_cast<QAbstractSlider*>(ptr)->value();
}

void* QAbstractSlider_NewQAbstractSlider(void* parent){
	return new QAbstractSlider(static_cast<QWidget*>(parent));
}

void QAbstractSlider_ConnectActionTriggered(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::actionTriggered), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ActionTriggered));;
}

void QAbstractSlider_DisconnectActionTriggered(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::actionTriggered), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ActionTriggered));;
}

void QAbstractSlider_ConnectRangeChanged(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int, int)>(&QAbstractSlider::rangeChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int, int)>(&MyQAbstractSlider::Signal_RangeChanged));;
}

void QAbstractSlider_DisconnectRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int, int)>(&QAbstractSlider::rangeChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int, int)>(&MyQAbstractSlider::Signal_RangeChanged));;
}

void QAbstractSlider_SetRange(void* ptr, int min, int max){
	QMetaObject::invokeMethod(static_cast<QAbstractSlider*>(ptr), "setRange", Q_ARG(int, min), Q_ARG(int, max));
}

void QAbstractSlider_ConnectSliderMoved(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::sliderMoved), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_SliderMoved));;
}

void QAbstractSlider_DisconnectSliderMoved(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::sliderMoved), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_SliderMoved));;
}

void QAbstractSlider_ConnectSliderPressed(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderPressed), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderPressed));;
}

void QAbstractSlider_DisconnectSliderPressed(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderPressed), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderPressed));;
}

void QAbstractSlider_ConnectSliderReleased(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderReleased), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderReleased));;
}

void QAbstractSlider_DisconnectSliderReleased(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)()>(&QAbstractSlider::sliderReleased), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)()>(&MyQAbstractSlider::Signal_SliderReleased));;
}

void QAbstractSlider_TriggerAction(void* ptr, int action){
	static_cast<QAbstractSlider*>(ptr)->triggerAction(static_cast<QAbstractSlider::SliderAction>(action));
}

void QAbstractSlider_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::valueChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ValueChanged));;
}

void QAbstractSlider_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractSlider*>(ptr), static_cast<void (QAbstractSlider::*)(int)>(&QAbstractSlider::valueChanged), static_cast<MyQAbstractSlider*>(ptr), static_cast<void (MyQAbstractSlider::*)(int)>(&MyQAbstractSlider::Signal_ValueChanged));;
}

void QAbstractSlider_DestroyQAbstractSlider(void* ptr){
	static_cast<QAbstractSlider*>(ptr)->~QAbstractSlider();
}

#include "qinputdialog.h"
#include <QWidget>
#include <QLineEdit>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLine>
#include <QObject>
#include <QInputDialog>
#include "_cgo_export.h"

class MyQInputDialog: public QInputDialog {
public:
void Signal_IntValueChanged(int value){callbackQInputDialogIntValueChanged(this->objectName().toUtf8().data(), value);};
void Signal_IntValueSelected(int value){callbackQInputDialogIntValueSelected(this->objectName().toUtf8().data(), value);};
void Signal_TextValueChanged(const QString & text){callbackQInputDialogTextValueChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_TextValueSelected(const QString & text){callbackQInputDialogTextValueSelected(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

char* QInputDialog_CancelButtonText(void* ptr){
	return static_cast<QInputDialog*>(ptr)->cancelButtonText().toUtf8().data();
}

char* QInputDialog_ComboBoxItems(void* ptr){
	return static_cast<QInputDialog*>(ptr)->comboBoxItems().join("|").toUtf8().data();
}

int QInputDialog_DoubleDecimals(void* ptr){
	return static_cast<QInputDialog*>(ptr)->doubleDecimals();
}

int QInputDialog_InputMode(void* ptr){
	return static_cast<QInputDialog*>(ptr)->inputMode();
}

int QInputDialog_IntMaximum(void* ptr){
	return static_cast<QInputDialog*>(ptr)->intMaximum();
}

int QInputDialog_IntMinimum(void* ptr){
	return static_cast<QInputDialog*>(ptr)->intMinimum();
}

int QInputDialog_IntStep(void* ptr){
	return static_cast<QInputDialog*>(ptr)->intStep();
}

int QInputDialog_IntValue(void* ptr){
	return static_cast<QInputDialog*>(ptr)->intValue();
}

int QInputDialog_IsComboBoxEditable(void* ptr){
	return static_cast<QInputDialog*>(ptr)->isComboBoxEditable();
}

char* QInputDialog_LabelText(void* ptr){
	return static_cast<QInputDialog*>(ptr)->labelText().toUtf8().data();
}

char* QInputDialog_OkButtonText(void* ptr){
	return static_cast<QInputDialog*>(ptr)->okButtonText().toUtf8().data();
}

int QInputDialog_Options(void* ptr){
	return static_cast<QInputDialog*>(ptr)->options();
}

void QInputDialog_SetCancelButtonText(void* ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setCancelButtonText(QString(text));
}

void QInputDialog_SetComboBoxEditable(void* ptr, int editable){
	static_cast<QInputDialog*>(ptr)->setComboBoxEditable(editable != 0);
}

void QInputDialog_SetComboBoxItems(void* ptr, char* items){
	static_cast<QInputDialog*>(ptr)->setComboBoxItems(QString(items).split("|", QString::SkipEmptyParts));
}

void QInputDialog_SetDoubleDecimals(void* ptr, int decimals){
	static_cast<QInputDialog*>(ptr)->setDoubleDecimals(decimals);
}

void QInputDialog_SetInputMode(void* ptr, int mode){
	static_cast<QInputDialog*>(ptr)->setInputMode(static_cast<QInputDialog::InputMode>(mode));
}

void QInputDialog_SetIntMaximum(void* ptr, int max){
	static_cast<QInputDialog*>(ptr)->setIntMaximum(max);
}

void QInputDialog_SetIntMinimum(void* ptr, int min){
	static_cast<QInputDialog*>(ptr)->setIntMinimum(min);
}

void QInputDialog_SetIntStep(void* ptr, int step){
	static_cast<QInputDialog*>(ptr)->setIntStep(step);
}

void QInputDialog_SetIntValue(void* ptr, int value){
	static_cast<QInputDialog*>(ptr)->setIntValue(value);
}

void QInputDialog_SetLabelText(void* ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setLabelText(QString(text));
}

void QInputDialog_SetOkButtonText(void* ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setOkButtonText(QString(text));
}

void QInputDialog_SetOptions(void* ptr, int options){
	static_cast<QInputDialog*>(ptr)->setOptions(static_cast<QInputDialog::InputDialogOption>(options));
}

void QInputDialog_SetTextEchoMode(void* ptr, int mode){
	static_cast<QInputDialog*>(ptr)->setTextEchoMode(static_cast<QLineEdit::EchoMode>(mode));
}

void QInputDialog_SetTextValue(void* ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setTextValue(QString(text));
}

int QInputDialog_TextEchoMode(void* ptr){
	return static_cast<QInputDialog*>(ptr)->textEchoMode();
}

char* QInputDialog_TextValue(void* ptr){
	return static_cast<QInputDialog*>(ptr)->textValue().toUtf8().data();
}

void* QInputDialog_NewQInputDialog(void* parent, int flags){
	return new QInputDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QInputDialog_Done(void* ptr, int result){
	static_cast<QInputDialog*>(ptr)->done(result);
}

int QInputDialog_QInputDialog_GetInt(void* parent, char* title, char* label, int value, int min, int max, int step, int ok, int flags){
	return QInputDialog::getInt(static_cast<QWidget*>(parent), QString(title), QString(label), value, min, max, step, NULL, static_cast<Qt::WindowType>(flags));
}

char* QInputDialog_QInputDialog_GetItem(void* parent, char* title, char* label, char* items, int current, int editable, int ok, int flags, int inputMethodHints){
	return QInputDialog::getItem(static_cast<QWidget*>(parent), QString(title), QString(label), QString(items).split("|", QString::SkipEmptyParts), current, editable != 0, NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

char* QInputDialog_QInputDialog_GetMultiLineText(void* parent, char* title, char* label, char* text, int ok, int flags, int inputMethodHints){
	return QInputDialog::getMultiLineText(static_cast<QWidget*>(parent), QString(title), QString(label), QString(text), NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

char* QInputDialog_QInputDialog_GetText(void* parent, char* title, char* label, int mode, char* text, int ok, int flags, int inputMethodHints){
	return QInputDialog::getText(static_cast<QWidget*>(parent), QString(title), QString(label), static_cast<QLineEdit::EchoMode>(mode), QString(text), NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

void QInputDialog_ConnectIntValueChanged(void* ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueChanged));;
}

void QInputDialog_DisconnectIntValueChanged(void* ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueChanged));;
}

void QInputDialog_ConnectIntValueSelected(void* ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueSelected));;
}

void QInputDialog_DisconnectIntValueSelected(void* ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueSelected));;
}

void QInputDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QInputDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QInputDialog_SetIntRange(void* ptr, int min, int max){
	static_cast<QInputDialog*>(ptr)->setIntRange(min, max);
}

void QInputDialog_SetOption(void* ptr, int option, int on){
	static_cast<QInputDialog*>(ptr)->setOption(static_cast<QInputDialog::InputDialogOption>(option), on != 0);
}

void QInputDialog_SetVisible(void* ptr, int visible){
	static_cast<QInputDialog*>(ptr)->setVisible(visible != 0);
}

int QInputDialog_TestOption(void* ptr, int option){
	return static_cast<QInputDialog*>(ptr)->testOption(static_cast<QInputDialog::InputDialogOption>(option));
}

void QInputDialog_ConnectTextValueChanged(void* ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueChanged));;
}

void QInputDialog_DisconnectTextValueChanged(void* ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueChanged));;
}

void QInputDialog_ConnectTextValueSelected(void* ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueSelected));;
}

void QInputDialog_DisconnectTextValueSelected(void* ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueSelected));;
}

void QInputDialog_DestroyQInputDialog(void* ptr){
	static_cast<QInputDialog*>(ptr)->~QInputDialog();
}

#include "qlistwidgetitem.h"
#include <QList>
#include <QListWidget>
#include <QDataStream>
#include <QFont>
#include <QBrush>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QList>
#include <QSize>
#include <QIcon>
#include <QListWidgetItem>
#include "_cgo_export.h"

class MyQListWidgetItem: public QListWidgetItem {
public:
};

void* QListWidgetItem_NewQListWidgetItem(void* parent, int ty){
	return new QListWidgetItem(static_cast<QListWidget*>(parent), ty);
}

void* QListWidgetItem_NewQListWidgetItem3(void* icon, char* text, void* parent, int ty){
	return new QListWidgetItem(*static_cast<QIcon*>(icon), QString(text), static_cast<QListWidget*>(parent), ty);
}

void* QListWidgetItem_NewQListWidgetItem2(char* text, void* parent, int ty){
	return new QListWidgetItem(QString(text), static_cast<QListWidget*>(parent), ty);
}

void QListWidgetItem_SetFlags(void* ptr, int flags){
	static_cast<QListWidgetItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void* QListWidgetItem_NewQListWidgetItem4(void* other){
	return new QListWidgetItem(*static_cast<QListWidgetItem*>(other));
}

void* QListWidgetItem_Background(void* ptr){
	return new QBrush(static_cast<QListWidgetItem*>(ptr)->background());
}

int QListWidgetItem_CheckState(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->checkState();
}

void* QListWidgetItem_Clone(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->clone();
}

void* QListWidgetItem_Data(void* ptr, int role){
	return new QVariant(static_cast<QListWidgetItem*>(ptr)->data(role));
}

int QListWidgetItem_Flags(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->flags();
}

void* QListWidgetItem_Foreground(void* ptr){
	return new QBrush(static_cast<QListWidgetItem*>(ptr)->foreground());
}

int QListWidgetItem_IsHidden(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->isHidden();
}

int QListWidgetItem_IsSelected(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->isSelected();
}

void* QListWidgetItem_ListWidget(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->listWidget();
}

void QListWidgetItem_Read(void* ptr, void* in){
	static_cast<QListWidgetItem*>(ptr)->read(*static_cast<QDataStream*>(in));
}

void QListWidgetItem_SetBackground(void* ptr, void* brush){
	static_cast<QListWidgetItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetCheckState(void* ptr, int state){
	static_cast<QListWidgetItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QListWidgetItem_SetData(void* ptr, int role, void* value){
	static_cast<QListWidgetItem*>(ptr)->setData(role, *static_cast<QVariant*>(value));
}

void QListWidgetItem_SetFont(void* ptr, void* font){
	static_cast<QListWidgetItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QListWidgetItem_SetForeground(void* ptr, void* brush){
	static_cast<QListWidgetItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QListWidgetItem_SetHidden(void* ptr, int hide){
	static_cast<QListWidgetItem*>(ptr)->setHidden(hide != 0);
}

void QListWidgetItem_SetIcon(void* ptr, void* icon){
	static_cast<QListWidgetItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QListWidgetItem_SetSelected(void* ptr, int sele){
	static_cast<QListWidgetItem*>(ptr)->setSelected(sele != 0);
}

void QListWidgetItem_SetSizeHint(void* ptr, void* size){
	static_cast<QListWidgetItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QListWidgetItem_SetStatusTip(void* ptr, char* statusTip){
	static_cast<QListWidgetItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QListWidgetItem_SetText(void* ptr, char* text){
	static_cast<QListWidgetItem*>(ptr)->setText(QString(text));
}

void QListWidgetItem_SetTextAlignment(void* ptr, int alignment){
	static_cast<QListWidgetItem*>(ptr)->setTextAlignment(alignment);
}

void QListWidgetItem_SetToolTip(void* ptr, char* toolTip){
	static_cast<QListWidgetItem*>(ptr)->setToolTip(QString(toolTip));
}

void QListWidgetItem_SetWhatsThis(void* ptr, char* whatsThis){
	static_cast<QListWidgetItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

char* QListWidgetItem_StatusTip(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->statusTip().toUtf8().data();
}

char* QListWidgetItem_Text(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->text().toUtf8().data();
}

int QListWidgetItem_TextAlignment(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->textAlignment();
}

char* QListWidgetItem_ToolTip(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->toolTip().toUtf8().data();
}

int QListWidgetItem_Type(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->type();
}

char* QListWidgetItem_WhatsThis(void* ptr){
	return static_cast<QListWidgetItem*>(ptr)->whatsThis().toUtf8().data();
}

void QListWidgetItem_Write(void* ptr, void* out){
	static_cast<QListWidgetItem*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QListWidgetItem_DestroyQListWidgetItem(void* ptr){
	static_cast<QListWidgetItem*>(ptr)->~QListWidgetItem();
}

#include "qstyleoptiontitlebar.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionTitleBar>
#include "_cgo_export.h"

class MyQStyleOptionTitleBar: public QStyleOptionTitleBar {
public:
};

void* QStyleOptionTitleBar_NewQStyleOptionTitleBar(){
	return new QStyleOptionTitleBar();
}

void* QStyleOptionTitleBar_NewQStyleOptionTitleBar2(void* other){
	return new QStyleOptionTitleBar(*static_cast<QStyleOptionTitleBar*>(other));
}

#include "qcolumnview.h"
#include <QObject>
#include <QItemSelectionModel>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemSelection>
#include <QWidget>
#include <QAbstractItemView>
#include <QString>
#include <QAbstractItemModel>
#include <QPoint>
#include <QColumnView>
#include "_cgo_export.h"

class MyQColumnView: public QColumnView {
public:
void Signal_UpdatePreviewWidget(const QModelIndex & index){callbackQColumnViewUpdatePreviewWidget(this->objectName().toUtf8().data(), index.internalPointer());};
};

int QColumnView_ResizeGripsVisible(void* ptr){
	return static_cast<QColumnView*>(ptr)->resizeGripsVisible();
}

void QColumnView_SetResizeGripsVisible(void* ptr, int visible){
	static_cast<QColumnView*>(ptr)->setResizeGripsVisible(visible != 0);
}

void* QColumnView_NewQColumnView(void* parent){
	return new QColumnView(static_cast<QWidget*>(parent));
}

void* QColumnView_IndexAt(void* ptr, void* point){
	return static_cast<QColumnView*>(ptr)->indexAt(*static_cast<QPoint*>(point)).internalPointer();
}

void* QColumnView_PreviewWidget(void* ptr){
	return static_cast<QColumnView*>(ptr)->previewWidget();
}

void QColumnView_ScrollTo(void* ptr, void* index, int hint){
	static_cast<QColumnView*>(ptr)->scrollTo(*static_cast<QModelIndex*>(index), static_cast<QAbstractItemView::ScrollHint>(hint));
}

void QColumnView_SelectAll(void* ptr){
	static_cast<QColumnView*>(ptr)->selectAll();
}

void QColumnView_SetModel(void* ptr, void* model){
	static_cast<QColumnView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QColumnView_SetPreviewWidget(void* ptr, void* widget){
	static_cast<QColumnView*>(ptr)->setPreviewWidget(static_cast<QWidget*>(widget));
}

void QColumnView_SetRootIndex(void* ptr, void* index){
	static_cast<QColumnView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QColumnView_SetSelectionModel(void* ptr, void* newSelectionModel){
	static_cast<QColumnView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(newSelectionModel));
}

void QColumnView_ConnectUpdatePreviewWidget(void* ptr){
	QObject::connect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DisconnectUpdatePreviewWidget(void* ptr){
	QObject::disconnect(static_cast<QColumnView*>(ptr), static_cast<void (QColumnView::*)(const QModelIndex &)>(&QColumnView::updatePreviewWidget), static_cast<MyQColumnView*>(ptr), static_cast<void (MyQColumnView::*)(const QModelIndex &)>(&MyQColumnView::Signal_UpdatePreviewWidget));;
}

void QColumnView_DestroyQColumnView(void* ptr){
	static_cast<QColumnView*>(ptr)->~QColumnView();
}

#include "qstyleoptiontabbarbase.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionTab>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionTabBarBase>
#include "_cgo_export.h"

class MyQStyleOptionTabBarBase: public QStyleOptionTabBarBase {
public:
};

void* QStyleOptionTabBarBase_NewQStyleOptionTabBarBase(){
	return new QStyleOptionTabBarBase();
}

void* QStyleOptionTabBarBase_NewQStyleOptionTabBarBase2(void* other){
	return new QStyleOptionTabBarBase(*static_cast<QStyleOptionTabBarBase*>(other));
}

#include "qlcdnumber.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QLCDNumber>
#include "_cgo_export.h"

class MyQLCDNumber: public QLCDNumber {
public:
void Signal_Overflow(){callbackQLCDNumberOverflow(this->objectName().toUtf8().data());};
};

int QLCDNumber_IntValue(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->intValue();
}

int QLCDNumber_Mode(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->mode();
}

int QLCDNumber_SegmentStyle(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->segmentStyle();
}

void QLCDNumber_SetMode(void* ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setMode(static_cast<QLCDNumber::Mode>(v));
}

void QLCDNumber_SetSegmentStyle(void* ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setSegmentStyle(static_cast<QLCDNumber::SegmentStyle>(v));
}

void QLCDNumber_SetSmallDecimalPoint(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setSmallDecimalPoint", Q_ARG(bool, v != 0));
}

int QLCDNumber_SmallDecimalPoint(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->smallDecimalPoint();
}

void* QLCDNumber_NewQLCDNumber(void* parent){
	return new QLCDNumber(static_cast<QWidget*>(parent));
}

int QLCDNumber_CheckOverflow2(void* ptr, int num){
	return static_cast<QLCDNumber*>(ptr)->checkOverflow(num);
}

int QLCDNumber_DigitCount(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->digitCount();
}

void QLCDNumber_Display(void* ptr, char* s){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(QString, QString(s)));
}

void QLCDNumber_Display3(void* ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(int, num));
}

void QLCDNumber_ConnectOverflow(void* ptr){
	QObject::connect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_DisconnectOverflow(void* ptr){
	QObject::disconnect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_SetBinMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setBinMode");
}

void QLCDNumber_SetDecMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setDecMode");
}

void QLCDNumber_SetDigitCount(void* ptr, int numDigits){
	static_cast<QLCDNumber*>(ptr)->setDigitCount(numDigits);
}

void QLCDNumber_SetHexMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setHexMode");
}

void QLCDNumber_SetOctMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setOctMode");
}

void QLCDNumber_DestroyQLCDNumber(void* ptr){
	static_cast<QLCDNumber*>(ptr)->~QLCDNumber();
}

#include "qopenglwidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include "_cgo_export.h"

#include "qgraphicsscenedragdropevent.h"
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QGraphicsSceneDragDropEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneDragDropEvent: public QGraphicsSceneDragDropEvent {
public:
};

void QGraphicsSceneDragDropEvent_AcceptProposedAction(void* ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->acceptProposedAction();
}

int QGraphicsSceneDragDropEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->buttons();
}

int QGraphicsSceneDragDropEvent_DropAction(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->dropAction();
}

void* QGraphicsSceneDragDropEvent_MimeData(void* ptr){
	return const_cast<QMimeData*>(static_cast<QGraphicsSceneDragDropEvent*>(ptr)->mimeData());
}

int QGraphicsSceneDragDropEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->modifiers();
}

int QGraphicsSceneDragDropEvent_PossibleActions(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->possibleActions();
}

int QGraphicsSceneDragDropEvent_ProposedAction(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->proposedAction();
}

void QGraphicsSceneDragDropEvent_SetDropAction(void* ptr, int action){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

void* QGraphicsSceneDragDropEvent_Source(void* ptr){
	return static_cast<QGraphicsSceneDragDropEvent*>(ptr)->source();
}

void QGraphicsSceneDragDropEvent_DestroyQGraphicsSceneDragDropEvent(void* ptr){
	static_cast<QGraphicsSceneDragDropEvent*>(ptr)->~QGraphicsSceneDragDropEvent();
}

#include "qstyleoptionframe.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionFrame>
#include "_cgo_export.h"

class MyQStyleOptionFrame: public QStyleOptionFrame {
public:
};

void* QStyleOptionFrame_NewQStyleOptionFrame(){
	return new QStyleOptionFrame();
}

void* QStyleOptionFrame_NewQStyleOptionFrame2(void* other){
	return new QStyleOptionFrame(*static_cast<QStyleOptionFrame*>(other));
}

#include "qstyleoptioncomplex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionComplex>
#include "_cgo_export.h"

class MyQStyleOptionComplex: public QStyleOptionComplex {
public:
};

void* QStyleOptionComplex_NewQStyleOptionComplex2(void* other){
	return new QStyleOptionComplex(*static_cast<QStyleOptionComplex*>(other));
}

void* QStyleOptionComplex_NewQStyleOptionComplex(int version, int ty){
	return new QStyleOptionComplex(version, ty);
}

#include "qstyleoptionmenuitem.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOptionMenuItem>
#include "_cgo_export.h"

class MyQStyleOptionMenuItem: public QStyleOptionMenuItem {
public:
};

void* QStyleOptionMenuItem_NewQStyleOptionMenuItem(){
	return new QStyleOptionMenuItem();
}

void* QStyleOptionMenuItem_NewQStyleOptionMenuItem2(void* other){
	return new QStyleOptionMenuItem(*static_cast<QStyleOptionMenuItem*>(other));
}

#include "qpangesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QPanGesture>
#include "_cgo_export.h"

class MyQPanGesture: public QPanGesture {
public:
};

double QPanGesture_Acceleration(void* ptr){
	return static_cast<double>(static_cast<QPanGesture*>(ptr)->acceleration());
}

void QPanGesture_SetAcceleration(void* ptr, double value){
	static_cast<QPanGesture*>(ptr)->setAcceleration(static_cast<qreal>(value));
}

void QPanGesture_SetLastOffset(void* ptr, void* value){
	static_cast<QPanGesture*>(ptr)->setLastOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_SetOffset(void* ptr, void* value){
	static_cast<QPanGesture*>(ptr)->setOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_DestroyQPanGesture(void* ptr){
	static_cast<QPanGesture*>(ptr)->~QPanGesture();
}

#include "qgraphicsdropshadoweffect.h"
#include <QModelIndex>
#include <QObject>
#include <QPointF>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QColor>
#include <QMetaObject>
#include <QString>
#include <QGraphicsDropShadowEffect>
#include "_cgo_export.h"

class MyQGraphicsDropShadowEffect: public QGraphicsDropShadowEffect {
public:
void Signal_ColorChanged(const QColor & color){callbackQGraphicsDropShadowEffectColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

double QGraphicsDropShadowEffect_BlurRadius(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->blurRadius());
}

void* QGraphicsDropShadowEffect_Color(void* ptr){
	return new QColor(static_cast<QGraphicsDropShadowEffect*>(ptr)->color());
}

void QGraphicsDropShadowEffect_SetBlurRadius(void* ptr, double blurRadius){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setBlurRadius", Q_ARG(qreal, static_cast<qreal>(blurRadius)));
}

void QGraphicsDropShadowEffect_SetColor(void* ptr, void* color){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QGraphicsDropShadowEffect_SetOffset(void* ptr, void* ofs){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(QPointF, *static_cast<QPointF*>(ofs)));
}

void* QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(void* parent){
	return new QGraphicsDropShadowEffect(static_cast<QObject*>(parent));
}

void QGraphicsDropShadowEffect_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsDropShadowEffect*>(ptr), static_cast<void (QGraphicsDropShadowEffect::*)(const QColor &)>(&QGraphicsDropShadowEffect::colorChanged), static_cast<MyQGraphicsDropShadowEffect*>(ptr), static_cast<void (MyQGraphicsDropShadowEffect::*)(const QColor &)>(&MyQGraphicsDropShadowEffect::Signal_ColorChanged));;
}

void QGraphicsDropShadowEffect_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsDropShadowEffect*>(ptr), static_cast<void (QGraphicsDropShadowEffect::*)(const QColor &)>(&QGraphicsDropShadowEffect::colorChanged), static_cast<MyQGraphicsDropShadowEffect*>(ptr), static_cast<void (MyQGraphicsDropShadowEffect::*)(const QColor &)>(&MyQGraphicsDropShadowEffect::Signal_ColorChanged));;
}

void QGraphicsDropShadowEffect_SetOffset3(void* ptr, double d){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(qreal, static_cast<qreal>(d)));
}

void QGraphicsDropShadowEffect_SetOffset2(void* ptr, double dx, double dy){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setOffset", Q_ARG(qreal, static_cast<qreal>(dx)), Q_ARG(qreal, static_cast<qreal>(dy)));
}

void QGraphicsDropShadowEffect_SetXOffset(void* ptr, double dx){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setXOffset", Q_ARG(qreal, static_cast<qreal>(dx)));
}

void QGraphicsDropShadowEffect_SetYOffset(void* ptr, double dy){
	QMetaObject::invokeMethod(static_cast<QGraphicsDropShadowEffect*>(ptr), "setYOffset", Q_ARG(qreal, static_cast<qreal>(dy)));
}

double QGraphicsDropShadowEffect_XOffset(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->xOffset());
}

double QGraphicsDropShadowEffect_YOffset(void* ptr){
	return static_cast<double>(static_cast<QGraphicsDropShadowEffect*>(ptr)->yOffset());
}

void QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(void* ptr){
	static_cast<QGraphicsDropShadowEffect*>(ptr)->~QGraphicsDropShadowEffect();
}

#include "qsplashscreen.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QColor>
#include <QPixmap>
#include <QSplashScreen>
#include "_cgo_export.h"

class MyQSplashScreen: public QSplashScreen {
public:
void Signal_MessageChanged(const QString & message){callbackQSplashScreenMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

void* QSplashScreen_NewQSplashScreen2(void* parent, void* pixmap, int f){
	return new QSplashScreen(static_cast<QWidget*>(parent), *static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

void* QSplashScreen_NewQSplashScreen(void* pixmap, int f){
	return new QSplashScreen(*static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

void QSplashScreen_ClearMessage(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "clearMessage");
}

void QSplashScreen_Finish(void* ptr, void* mainWin){
	static_cast<QSplashScreen*>(ptr)->finish(static_cast<QWidget*>(mainWin));
}

char* QSplashScreen_Message(void* ptr){
	return static_cast<QSplashScreen*>(ptr)->message().toUtf8().data();
}

void QSplashScreen_ConnectMessageChanged(void* ptr){
	QObject::connect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_DisconnectMessageChanged(void* ptr){
	QObject::disconnect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_Repaint(void* ptr){
	static_cast<QSplashScreen*>(ptr)->repaint();
}

void QSplashScreen_SetPixmap(void* ptr, void* pixmap){
	static_cast<QSplashScreen*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QSplashScreen_ShowMessage(void* ptr, char* message, int alignment, void* color){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, alignment), Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QSplashScreen_DestroyQSplashScreen(void* ptr){
	static_cast<QSplashScreen*>(ptr)->~QSplashScreen();
}

#include "qstyleoptiontabwidgetframe.h"
#include <QStyleOptionTab>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyleOptionTabWidgetFrame>
#include "_cgo_export.h"

class MyQStyleOptionTabWidgetFrame: public QStyleOptionTabWidgetFrame {
public:
};

void* QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame(){
	return new QStyleOptionTabWidgetFrame();
}

void* QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame2(void* other){
	return new QStyleOptionTabWidgetFrame(*static_cast<QStyleOptionTabWidgetFrame*>(other));
}

#include "qstylehintreturnmask.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleHintReturn>
#include <QStyle>
#include <QStyleHintReturnMask>
#include "_cgo_export.h"

class MyQStyleHintReturnMask: public QStyleHintReturnMask {
public:
};

void* QStyleHintReturnMask_NewQStyleHintReturnMask(){
	return new QStyleHintReturnMask();
}

void QStyleHintReturnMask_DestroyQStyleHintReturnMask(void* ptr){
	static_cast<QStyleHintReturnMask*>(ptr)->~QStyleHintReturnMask();
}

#include "qplaintextedit.h"
#include <QVariant>
#include <QTextOption>
#include <QTextCharFormat>
#include <QPoint>
#include <QMetaObject>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QTextCursor>
#include <QObject>
#include <QWidget>
#include <QPagedPaintDevice>
#include <QTextDocument>
#include <QPlainTextEdit>
#include "_cgo_export.h"

class MyQPlainTextEdit: public QPlainTextEdit {
public:
void Signal_BlockCountChanged(int newBlockCount){callbackQPlainTextEditBlockCountChanged(this->objectName().toUtf8().data(), newBlockCount);};
void Signal_CopyAvailable(bool yes){callbackQPlainTextEditCopyAvailable(this->objectName().toUtf8().data(), yes);};
void Signal_CursorPositionChanged(){callbackQPlainTextEditCursorPositionChanged(this->objectName().toUtf8().data());};
void Signal_ModificationChanged(bool changed){callbackQPlainTextEditModificationChanged(this->objectName().toUtf8().data(), changed);};
void Signal_RedoAvailable(bool available){callbackQPlainTextEditRedoAvailable(this->objectName().toUtf8().data(), available);};
void Signal_SelectionChanged(){callbackQPlainTextEditSelectionChanged(this->objectName().toUtf8().data());};
void Signal_TextChanged(){callbackQPlainTextEditTextChanged(this->objectName().toUtf8().data());};
void Signal_UndoAvailable(bool available){callbackQPlainTextEditUndoAvailable(this->objectName().toUtf8().data(), available);};
};

int QPlainTextEdit_BackgroundVisible(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->backgroundVisible();
}

int QPlainTextEdit_BlockCount(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->blockCount();
}

int QPlainTextEdit_CenterOnScroll(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->centerOnScroll();
}

int QPlainTextEdit_CursorWidth(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->cursorWidth();
}

int QPlainTextEdit_IsReadOnly(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isReadOnly();
}

int QPlainTextEdit_LineWrapMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->lineWrapMode();
}

int QPlainTextEdit_OverwriteMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->overwriteMode();
}

char* QPlainTextEdit_PlaceholderText(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->placeholderText().toUtf8().data();
}

void QPlainTextEdit_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "redo");
}

void QPlainTextEdit_SetBackgroundVisible(void* ptr, int visible){
	static_cast<QPlainTextEdit*>(ptr)->setBackgroundVisible(visible != 0);
}

void QPlainTextEdit_SetCenterOnScroll(void* ptr, int enabled){
	static_cast<QPlainTextEdit*>(ptr)->setCenterOnScroll(enabled != 0);
}

void QPlainTextEdit_SetCursorWidth(void* ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setCursorWidth(width);
}

void QPlainTextEdit_SetLineWrapMode(void* ptr, int mode){
	static_cast<QPlainTextEdit*>(ptr)->setLineWrapMode(static_cast<QPlainTextEdit::LineWrapMode>(mode));
}

void QPlainTextEdit_SetOverwriteMode(void* ptr, int overwrite){
	static_cast<QPlainTextEdit*>(ptr)->setOverwriteMode(overwrite != 0);
}

void QPlainTextEdit_SetPlaceholderText(void* ptr, char* placeholderText){
	static_cast<QPlainTextEdit*>(ptr)->setPlaceholderText(QString(placeholderText));
}

void QPlainTextEdit_SetReadOnly(void* ptr, int ro){
	static_cast<QPlainTextEdit*>(ptr)->setReadOnly(ro != 0);
}

void QPlainTextEdit_SetTabChangesFocus(void* ptr, int b){
	static_cast<QPlainTextEdit*>(ptr)->setTabChangesFocus(b != 0);
}

void QPlainTextEdit_SetTabStopWidth(void* ptr, int width){
	static_cast<QPlainTextEdit*>(ptr)->setTabStopWidth(width);
}

void QPlainTextEdit_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QPlainTextEdit*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QPlainTextEdit_SetWordWrapMode(void* ptr, int policy){
	static_cast<QPlainTextEdit*>(ptr)->setWordWrapMode(static_cast<QTextOption::WrapMode>(policy));
}

int QPlainTextEdit_TabChangesFocus(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabChangesFocus();
}

int QPlainTextEdit_TabStopWidth(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->tabStopWidth();
}

int QPlainTextEdit_TextInteractionFlags(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->textInteractionFlags();
}

int QPlainTextEdit_WordWrapMode(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->wordWrapMode();
}

void QPlainTextEdit_ZoomIn(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomIn", Q_ARG(int, ran));
}

void QPlainTextEdit_ZoomOut(void* ptr, int ran){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "zoomOut", Q_ARG(int, ran));
}

void* QPlainTextEdit_NewQPlainTextEdit(void* parent){
	return new QPlainTextEdit(static_cast<QWidget*>(parent));
}

void* QPlainTextEdit_NewQPlainTextEdit2(char* text, void* parent){
	return new QPlainTextEdit(QString(text), static_cast<QWidget*>(parent));
}

char* QPlainTextEdit_AnchorAt(void* ptr, void* pos){
	return static_cast<QPlainTextEdit*>(ptr)->anchorAt(*static_cast<QPoint*>(pos)).toUtf8().data();
}

void QPlainTextEdit_AppendPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_CenterCursor(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "centerCursor");
}

void QPlainTextEdit_AppendHtml(void* ptr, char* html){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "appendHtml", Q_ARG(QString, QString(html)));
}

void QPlainTextEdit_ConnectBlockCountChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

void QPlainTextEdit_DisconnectBlockCountChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(int)>(&QPlainTextEdit::blockCountChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(int)>(&MyQPlainTextEdit::Signal_BlockCountChanged));;
}

int QPlainTextEdit_CanPaste(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->canPaste();
}

void QPlainTextEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "clear");
}

void QPlainTextEdit_Copy(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "copy");
}

void QPlainTextEdit_ConnectCopyAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

void QPlainTextEdit_DisconnectCopyAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::copyAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_CopyAvailable));;
}

void* QPlainTextEdit_CreateStandardContextMenu(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu();
}

void* QPlainTextEdit_CreateStandardContextMenu2(void* ptr, void* position){
	return static_cast<QPlainTextEdit*>(ptr)->createStandardContextMenu(*static_cast<QPoint*>(position));
}

void QPlainTextEdit_ConnectCursorPositionChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_DisconnectCursorPositionChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::cursorPositionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_CursorPositionChanged));;
}

void QPlainTextEdit_Cut(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "cut");
}

void* QPlainTextEdit_Document(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->document();
}

char* QPlainTextEdit_DocumentTitle(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->documentTitle().toUtf8().data();
}

void QPlainTextEdit_EnsureCursorVisible(void* ptr){
	static_cast<QPlainTextEdit*>(ptr)->ensureCursorVisible();
}

void* QPlainTextEdit_InputMethodQuery(void* ptr, int property){
	return new QVariant(static_cast<QPlainTextEdit*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(property)));
}

void QPlainTextEdit_InsertPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "insertPlainText", Q_ARG(QString, QString(text)));
}

int QPlainTextEdit_IsUndoRedoEnabled(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->isUndoRedoEnabled();
}

void* QPlainTextEdit_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QPlainTextEdit*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

int QPlainTextEdit_MaximumBlockCount(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->maximumBlockCount();
}

void QPlainTextEdit_MergeCurrentCharFormat(void* ptr, void* modifier){
	static_cast<QPlainTextEdit*>(ptr)->mergeCurrentCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QPlainTextEdit_ConnectModificationChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_DisconnectModificationChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::modificationChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_ModificationChanged));;
}

void QPlainTextEdit_MoveCursor(void* ptr, int operation, int mode){
	static_cast<QPlainTextEdit*>(ptr)->moveCursor(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode));
}

void QPlainTextEdit_Paste(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "paste");
}

void QPlainTextEdit_Print(void* ptr, void* printer){
	static_cast<QPlainTextEdit*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QPlainTextEdit_ConnectRedoAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_DisconnectRedoAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::redoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_RedoAvailable));;
}

void QPlainTextEdit_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "selectAll");
}

void QPlainTextEdit_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::selectionChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_SelectionChanged));;
}

void QPlainTextEdit_SetCurrentCharFormat(void* ptr, void* format){
	static_cast<QPlainTextEdit*>(ptr)->setCurrentCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QPlainTextEdit_SetDocument(void* ptr, void* document){
	static_cast<QPlainTextEdit*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QPlainTextEdit_SetDocumentTitle(void* ptr, char* title){
	static_cast<QPlainTextEdit*>(ptr)->setDocumentTitle(QString(title));
}

void QPlainTextEdit_SetMaximumBlockCount(void* ptr, int maximum){
	static_cast<QPlainTextEdit*>(ptr)->setMaximumBlockCount(maximum);
}

void QPlainTextEdit_SetPlainText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "setPlainText", Q_ARG(QString, QString(text)));
}

void QPlainTextEdit_SetTextCursor(void* ptr, void* cursor){
	static_cast<QPlainTextEdit*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QPlainTextEdit_SetUndoRedoEnabled(void* ptr, int enable){
	static_cast<QPlainTextEdit*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QPlainTextEdit_ConnectTextChanged(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

void QPlainTextEdit_DisconnectTextChanged(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)()>(&QPlainTextEdit::textChanged), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)()>(&MyQPlainTextEdit::Signal_TextChanged));;
}

char* QPlainTextEdit_ToPlainText(void* ptr){
	return static_cast<QPlainTextEdit*>(ptr)->toPlainText().toUtf8().data();
}

void QPlainTextEdit_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QPlainTextEdit*>(ptr), "undo");
}

void QPlainTextEdit_ConnectUndoAvailable(void* ptr){
	QObject::connect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DisconnectUndoAvailable(void* ptr){
	QObject::disconnect(static_cast<QPlainTextEdit*>(ptr), static_cast<void (QPlainTextEdit::*)(bool)>(&QPlainTextEdit::undoAvailable), static_cast<MyQPlainTextEdit*>(ptr), static_cast<void (MyQPlainTextEdit::*)(bool)>(&MyQPlainTextEdit::Signal_UndoAvailable));;
}

void QPlainTextEdit_DestroyQPlainTextEdit(void* ptr){
	static_cast<QPlainTextEdit*>(ptr)->~QPlainTextEdit();
}

#include "qstylehintreturn.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleHintReturn>
#include "_cgo_export.h"

class MyQStyleHintReturn: public QStyleHintReturn {
public:
};

int QStyleHintReturn_SH_Mask_Type(){
	return QStyleHintReturn::SH_Mask;
}

int QStyleHintReturn_SH_Variant_Type(){
	return QStyleHintReturn::SH_Variant;
}

void* QStyleHintReturn_NewQStyleHintReturn(int version, int ty){
	return new QStyleHintReturn(version, ty);
}

#include "qcommandlinkbutton.h"
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QCommandLinkButton>
#include "_cgo_export.h"

class MyQCommandLinkButton: public QCommandLinkButton {
public:
};

char* QCommandLinkButton_Description(void* ptr){
	return static_cast<QCommandLinkButton*>(ptr)->description().toUtf8().data();
}

void QCommandLinkButton_SetDescription(void* ptr, char* description){
	static_cast<QCommandLinkButton*>(ptr)->setDescription(QString(description));
}

void* QCommandLinkButton_NewQCommandLinkButton(void* parent){
	return new QCommandLinkButton(static_cast<QWidget*>(parent));
}

void* QCommandLinkButton_NewQCommandLinkButton2(char* text, void* parent){
	return new QCommandLinkButton(QString(text), static_cast<QWidget*>(parent));
}

void* QCommandLinkButton_NewQCommandLinkButton3(char* text, char* description, void* parent){
	return new QCommandLinkButton(QString(text), QString(description), static_cast<QWidget*>(parent));
}

void QCommandLinkButton_DestroyQCommandLinkButton(void* ptr){
	static_cast<QCommandLinkButton*>(ptr)->~QCommandLinkButton();
}

#include "qdateedit.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDate>
#include <QWidget>
#include <QDateEdit>
#include "_cgo_export.h"

class MyQDateEdit: public QDateEdit {
public:
};

void* QDateEdit_NewQDateEdit(void* parent){
	return new QDateEdit(static_cast<QWidget*>(parent));
}

void* QDateEdit_NewQDateEdit2(void* date, void* parent){
	return new QDateEdit(*static_cast<QDate*>(date), static_cast<QWidget*>(parent));
}

void QDateEdit_DestroyQDateEdit(void* ptr){
	static_cast<QDateEdit*>(ptr)->~QDateEdit();
}

#include "qplaintextdocumentlayout.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextBlock>
#include <QTextDocument>
#include <QString>
#include <QVariant>
#include <QPlainTextDocumentLayout>
#include "_cgo_export.h"

class MyQPlainTextDocumentLayout: public QPlainTextDocumentLayout {
public:
};

int QPlainTextDocumentLayout_CursorWidth(void* ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->cursorWidth();
}

void QPlainTextDocumentLayout_SetCursorWidth(void* ptr, int width){
	static_cast<QPlainTextDocumentLayout*>(ptr)->setCursorWidth(width);
}

void* QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(void* document){
	return new QPlainTextDocumentLayout(static_cast<QTextDocument*>(document));
}

void QPlainTextDocumentLayout_EnsureBlockLayout(void* ptr, void* block){
	static_cast<QPlainTextDocumentLayout*>(ptr)->ensureBlockLayout(*static_cast<QTextBlock*>(block));
}

int QPlainTextDocumentLayout_PageCount(void* ptr){
	return static_cast<QPlainTextDocumentLayout*>(ptr)->pageCount();
}

void QPlainTextDocumentLayout_RequestUpdate(void* ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->requestUpdate();
}

void QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(void* ptr){
	static_cast<QPlainTextDocumentLayout*>(ptr)->~QPlainTextDocumentLayout();
}

#include "qgraphicstransform.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsTransform>
#include "_cgo_export.h"

#include "qradiobutton.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRadioButton>
#include "_cgo_export.h"

class MyQRadioButton: public QRadioButton {
public:
};

void* QRadioButton_NewQRadioButton(void* parent){
	return new QRadioButton(static_cast<QWidget*>(parent));
}

void* QRadioButton_NewQRadioButton2(char* text, void* parent){
	return new QRadioButton(QString(text), static_cast<QWidget*>(parent));
}

void QRadioButton_DestroyQRadioButton(void* ptr){
	static_cast<QRadioButton*>(ptr)->~QRadioButton();
}

#include "qstyle.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QPoint>
#include <QSize>
#include <QString>
#include <QStyleOption>
#include <QPixmap>
#include <QApplication>
#include <QRect>
#include <QStyleHintReturn>
#include <QPainter>
#include <QVariant>
#include <QPalette>
#include <QSizePolicy>
#include <QStyleOptionComplex>
#include <QStyle>
#include "_cgo_export.h"

class MyQStyle: public QStyle {
public:
};

int QStyle_PM_MdiSubWindowMinimizedWidth_Type(){
	return QStyle::PM_MdiSubWindowMinimizedWidth;
}

int QStyle_PM_HeaderMargin_Type(){
	return QStyle::PM_HeaderMargin;
}

int QStyle_PM_HeaderMarkSize_Type(){
	return QStyle::PM_HeaderMarkSize;
}

int QStyle_PM_HeaderGripMargin_Type(){
	return QStyle::PM_HeaderGripMargin;
}

int QStyle_PM_TabBarTabShiftHorizontal_Type(){
	return QStyle::PM_TabBarTabShiftHorizontal;
}

int QStyle_PM_TabBarTabShiftVertical_Type(){
	return QStyle::PM_TabBarTabShiftVertical;
}

int QStyle_PM_TabBarScrollButtonWidth_Type(){
	return QStyle::PM_TabBarScrollButtonWidth;
}

int QStyle_PM_ToolBarFrameWidth_Type(){
	return QStyle::PM_ToolBarFrameWidth;
}

int QStyle_PM_ToolBarHandleExtent_Type(){
	return QStyle::PM_ToolBarHandleExtent;
}

int QStyle_PM_ToolBarItemSpacing_Type(){
	return QStyle::PM_ToolBarItemSpacing;
}

int QStyle_PM_ToolBarItemMargin_Type(){
	return QStyle::PM_ToolBarItemMargin;
}

int QStyle_PM_ToolBarSeparatorExtent_Type(){
	return QStyle::PM_ToolBarSeparatorExtent;
}

int QStyle_PM_ToolBarExtensionExtent_Type(){
	return QStyle::PM_ToolBarExtensionExtent;
}

int QStyle_PM_SpinBoxSliderHeight_Type(){
	return QStyle::PM_SpinBoxSliderHeight;
}

int QStyle_PM_DefaultTopLevelMargin_Type(){
	return QStyle::PM_DefaultTopLevelMargin;
}

int QStyle_PM_DefaultChildMargin_Type(){
	return QStyle::PM_DefaultChildMargin;
}

int QStyle_PM_DefaultLayoutSpacing_Type(){
	return QStyle::PM_DefaultLayoutSpacing;
}

int QStyle_PM_ToolBarIconSize_Type(){
	return QStyle::PM_ToolBarIconSize;
}

int QStyle_PM_ListViewIconSize_Type(){
	return QStyle::PM_ListViewIconSize;
}

int QStyle_PM_IconViewIconSize_Type(){
	return QStyle::PM_IconViewIconSize;
}

int QStyle_PM_SmallIconSize_Type(){
	return QStyle::PM_SmallIconSize;
}

int QStyle_PM_LargeIconSize_Type(){
	return QStyle::PM_LargeIconSize;
}

int QStyle_PM_FocusFrameVMargin_Type(){
	return QStyle::PM_FocusFrameVMargin;
}

int QStyle_PM_FocusFrameHMargin_Type(){
	return QStyle::PM_FocusFrameHMargin;
}

int QStyle_PM_ToolTipLabelFrameWidth_Type(){
	return QStyle::PM_ToolTipLabelFrameWidth;
}

int QStyle_PM_CheckBoxLabelSpacing_Type(){
	return QStyle::PM_CheckBoxLabelSpacing;
}

int QStyle_PM_TabBarIconSize_Type(){
	return QStyle::PM_TabBarIconSize;
}

int QStyle_PM_SizeGripSize_Type(){
	return QStyle::PM_SizeGripSize;
}

int QStyle_PM_DockWidgetTitleMargin_Type(){
	return QStyle::PM_DockWidgetTitleMargin;
}

int QStyle_PM_MessageBoxIconSize_Type(){
	return QStyle::PM_MessageBoxIconSize;
}

int QStyle_PM_ButtonIconSize_Type(){
	return QStyle::PM_ButtonIconSize;
}

int QStyle_PM_DockWidgetTitleBarButtonMargin_Type(){
	return QStyle::PM_DockWidgetTitleBarButtonMargin;
}

int QStyle_PM_RadioButtonLabelSpacing_Type(){
	return QStyle::PM_RadioButtonLabelSpacing;
}

int QStyle_PM_LayoutLeftMargin_Type(){
	return QStyle::PM_LayoutLeftMargin;
}

int QStyle_PM_LayoutTopMargin_Type(){
	return QStyle::PM_LayoutTopMargin;
}

int QStyle_PM_LayoutRightMargin_Type(){
	return QStyle::PM_LayoutRightMargin;
}

int QStyle_PM_LayoutBottomMargin_Type(){
	return QStyle::PM_LayoutBottomMargin;
}

int QStyle_PM_LayoutHorizontalSpacing_Type(){
	return QStyle::PM_LayoutHorizontalSpacing;
}

int QStyle_PM_LayoutVerticalSpacing_Type(){
	return QStyle::PM_LayoutVerticalSpacing;
}

int QStyle_PM_TabBar_ScrollButtonOverlap_Type(){
	return QStyle::PM_TabBar_ScrollButtonOverlap;
}

int QStyle_PM_TextCursorWidth_Type(){
	return QStyle::PM_TextCursorWidth;
}

int QStyle_PM_TabCloseIndicatorWidth_Type(){
	return QStyle::PM_TabCloseIndicatorWidth;
}

int QStyle_PM_TabCloseIndicatorHeight_Type(){
	return QStyle::PM_TabCloseIndicatorHeight;
}

int QStyle_PM_ScrollView_ScrollBarSpacing_Type(){
	return QStyle::PM_ScrollView_ScrollBarSpacing;
}

int QStyle_PM_ScrollView_ScrollBarOverlap_Type(){
	return QStyle::PM_ScrollView_ScrollBarOverlap;
}

int QStyle_PM_SubMenuOverlap_Type(){
	return QStyle::PM_SubMenuOverlap;
}

int QStyle_PM_TreeViewIndentation_Type(){
	return QStyle::PM_TreeViewIndentation;
}

int QStyle_PM_HeaderDefaultSectionSizeHorizontal_Type(){
	return QStyle::PM_HeaderDefaultSectionSizeHorizontal;
}

int QStyle_PM_HeaderDefaultSectionSizeVertical_Type(){
	return QStyle::PM_HeaderDefaultSectionSizeVertical;
}

int QStyle_PE_FrameTabWidget_Type(){
	return QStyle::PE_FrameTabWidget;
}

int QStyle_PE_FrameWindow_Type(){
	return QStyle::PE_FrameWindow;
}

int QStyle_PE_FrameButtonBevel_Type(){
	return QStyle::PE_FrameButtonBevel;
}

int QStyle_PE_FrameButtonTool_Type(){
	return QStyle::PE_FrameButtonTool;
}

int QStyle_PE_FrameTabBarBase_Type(){
	return QStyle::PE_FrameTabBarBase;
}

int QStyle_PE_PanelButtonCommand_Type(){
	return QStyle::PE_PanelButtonCommand;
}

int QStyle_PE_PanelButtonBevel_Type(){
	return QStyle::PE_PanelButtonBevel;
}

int QStyle_PE_PanelButtonTool_Type(){
	return QStyle::PE_PanelButtonTool;
}

int QStyle_PE_PanelMenuBar_Type(){
	return QStyle::PE_PanelMenuBar;
}

int QStyle_PE_PanelToolBar_Type(){
	return QStyle::PE_PanelToolBar;
}

int QStyle_PE_PanelLineEdit_Type(){
	return QStyle::PE_PanelLineEdit;
}

int QStyle_PE_IndicatorArrowDown_Type(){
	return QStyle::PE_IndicatorArrowDown;
}

int QStyle_PE_IndicatorArrowLeft_Type(){
	return QStyle::PE_IndicatorArrowLeft;
}

int QStyle_PE_IndicatorArrowRight_Type(){
	return QStyle::PE_IndicatorArrowRight;
}

int QStyle_PE_IndicatorArrowUp_Type(){
	return QStyle::PE_IndicatorArrowUp;
}

int QStyle_PE_IndicatorBranch_Type(){
	return QStyle::PE_IndicatorBranch;
}

int QStyle_PE_IndicatorButtonDropDown_Type(){
	return QStyle::PE_IndicatorButtonDropDown;
}

int QStyle_PE_IndicatorViewItemCheck_Type(){
	return QStyle::PE_IndicatorViewItemCheck;
}

int QStyle_PE_IndicatorCheckBox_Type(){
	return QStyle::PE_IndicatorCheckBox;
}

int QStyle_PE_IndicatorDockWidgetResizeHandle_Type(){
	return QStyle::PE_IndicatorDockWidgetResizeHandle;
}

int QStyle_PE_IndicatorHeaderArrow_Type(){
	return QStyle::PE_IndicatorHeaderArrow;
}

int QStyle_PE_IndicatorMenuCheckMark_Type(){
	return QStyle::PE_IndicatorMenuCheckMark;
}

int QStyle_PE_IndicatorProgressChunk_Type(){
	return QStyle::PE_IndicatorProgressChunk;
}

int QStyle_PE_IndicatorRadioButton_Type(){
	return QStyle::PE_IndicatorRadioButton;
}

int QStyle_PE_IndicatorSpinDown_Type(){
	return QStyle::PE_IndicatorSpinDown;
}

int QStyle_PE_IndicatorSpinMinus_Type(){
	return QStyle::PE_IndicatorSpinMinus;
}

int QStyle_PE_IndicatorSpinPlus_Type(){
	return QStyle::PE_IndicatorSpinPlus;
}

int QStyle_PE_IndicatorSpinUp_Type(){
	return QStyle::PE_IndicatorSpinUp;
}

int QStyle_PE_IndicatorToolBarHandle_Type(){
	return QStyle::PE_IndicatorToolBarHandle;
}

int QStyle_PE_IndicatorToolBarSeparator_Type(){
	return QStyle::PE_IndicatorToolBarSeparator;
}

int QStyle_PE_PanelTipLabel_Type(){
	return QStyle::PE_PanelTipLabel;
}

int QStyle_PE_IndicatorTabTear_Type(){
	return QStyle::PE_IndicatorTabTear;
}

int QStyle_PE_PanelScrollAreaCorner_Type(){
	return QStyle::PE_PanelScrollAreaCorner;
}

int QStyle_PE_Widget_Type(){
	return QStyle::PE_Widget;
}

int QStyle_PE_IndicatorColumnViewArrow_Type(){
	return QStyle::PE_IndicatorColumnViewArrow;
}

int QStyle_PE_IndicatorItemViewItemDrop_Type(){
	return QStyle::PE_IndicatorItemViewItemDrop;
}

int QStyle_PE_PanelItemViewItem_Type(){
	return QStyle::PE_PanelItemViewItem;
}

int QStyle_PE_PanelItemViewRow_Type(){
	return QStyle::PE_PanelItemViewRow;
}

int QStyle_PE_PanelStatusBar_Type(){
	return QStyle::PE_PanelStatusBar;
}

int QStyle_PE_IndicatorTabClose_Type(){
	return QStyle::PE_IndicatorTabClose;
}

int QStyle_PE_PanelMenu_Type(){
	return QStyle::PE_PanelMenu;
}

int QStyle_SH_BlinkCursorWhenTextSelected_Type(){
	return QStyle::SH_BlinkCursorWhenTextSelected;
}

int QStyle_SH_RichText_FullWidthSelection_Type(){
	return QStyle::SH_RichText_FullWidthSelection;
}

int QStyle_SH_Menu_Scrollable_Type(){
	return QStyle::SH_Menu_Scrollable;
}

int QStyle_SH_GroupBox_TextLabelVerticalAlignment_Type(){
	return QStyle::SH_GroupBox_TextLabelVerticalAlignment;
}

int QStyle_SH_GroupBox_TextLabelColor_Type(){
	return QStyle::SH_GroupBox_TextLabelColor;
}

int QStyle_SH_Menu_SloppySubMenus_Type(){
	return QStyle::SH_Menu_SloppySubMenus;
}

int QStyle_SH_Table_GridLineColor_Type(){
	return QStyle::SH_Table_GridLineColor;
}

int QStyle_SH_LineEdit_PasswordCharacter_Type(){
	return QStyle::SH_LineEdit_PasswordCharacter;
}

int QStyle_SH_DialogButtons_DefaultButton_Type(){
	return QStyle::SH_DialogButtons_DefaultButton;
}

int QStyle_SH_ToolBox_SelectedPageTitleBold_Type(){
	return QStyle::SH_ToolBox_SelectedPageTitleBold;
}

int QStyle_SH_TabBar_PreferNoArrows_Type(){
	return QStyle::SH_TabBar_PreferNoArrows;
}

int QStyle_SH_ScrollBar_LeftClickAbsolutePosition_Type(){
	return QStyle::SH_ScrollBar_LeftClickAbsolutePosition;
}

int QStyle_SH_ListViewExpand_SelectMouseType_Type(){
	return QStyle::SH_ListViewExpand_SelectMouseType;
}

int QStyle_SH_UnderlineShortcut_Type(){
	return QStyle::SH_UnderlineShortcut;
}

int QStyle_SH_SpinBox_AnimateButton_Type(){
	return QStyle::SH_SpinBox_AnimateButton;
}

int QStyle_SH_SpinBox_KeyPressAutoRepeatRate_Type(){
	return QStyle::SH_SpinBox_KeyPressAutoRepeatRate;
}

int QStyle_SH_SpinBox_ClickAutoRepeatRate_Type(){
	return QStyle::SH_SpinBox_ClickAutoRepeatRate;
}

int QStyle_SH_Menu_FillScreenWithScroll_Type(){
	return QStyle::SH_Menu_FillScreenWithScroll;
}

int QStyle_SH_ToolTipLabel_Opacity_Type(){
	return QStyle::SH_ToolTipLabel_Opacity;
}

int QStyle_SH_DrawMenuBarSeparator_Type(){
	return QStyle::SH_DrawMenuBarSeparator;
}

int QStyle_SH_TitleBar_ModifyNotification_Type(){
	return QStyle::SH_TitleBar_ModifyNotification;
}

int QStyle_SH_Button_FocusPolicy_Type(){
	return QStyle::SH_Button_FocusPolicy;
}

int QStyle_SH_MessageBox_UseBorderForButtonSpacing_Type(){
	return QStyle::SH_MessageBox_UseBorderForButtonSpacing;
}

int QStyle_SH_TitleBar_AutoRaise_Type(){
	return QStyle::SH_TitleBar_AutoRaise;
}

int QStyle_SH_ToolButton_PopupDelay_Type(){
	return QStyle::SH_ToolButton_PopupDelay;
}

int QStyle_SH_FocusFrame_Mask_Type(){
	return QStyle::SH_FocusFrame_Mask;
}

int QStyle_SH_RubberBand_Mask_Type(){
	return QStyle::SH_RubberBand_Mask;
}

int QStyle_SH_WindowFrame_Mask_Type(){
	return QStyle::SH_WindowFrame_Mask;
}

int QStyle_SH_SpinControls_DisableOnBounds_Type(){
	return QStyle::SH_SpinControls_DisableOnBounds;
}

int QStyle_SH_Dial_BackgroundRole_Type(){
	return QStyle::SH_Dial_BackgroundRole;
}

int QStyle_SH_ComboBox_LayoutDirection_Type(){
	return QStyle::SH_ComboBox_LayoutDirection;
}

int QStyle_SH_ItemView_EllipsisLocation_Type(){
	return QStyle::SH_ItemView_EllipsisLocation;
}

int QStyle_SH_ItemView_ShowDecorationSelected_Type(){
	return QStyle::SH_ItemView_ShowDecorationSelected;
}

int QStyle_SH_ItemView_ActivateItemOnSingleClick_Type(){
	return QStyle::SH_ItemView_ActivateItemOnSingleClick;
}

int QStyle_SH_ScrollBar_ContextMenu_Type(){
	return QStyle::SH_ScrollBar_ContextMenu;
}

int QStyle_SH_ScrollBar_RollBetweenButtons_Type(){
	return QStyle::SH_ScrollBar_RollBetweenButtons;
}

int QStyle_SH_Slider_AbsoluteSetButtons_Type(){
	return QStyle::SH_Slider_AbsoluteSetButtons;
}

int QStyle_SH_Slider_PageSetButtons_Type(){
	return QStyle::SH_Slider_PageSetButtons;
}

int QStyle_SH_Menu_KeyboardSearch_Type(){
	return QStyle::SH_Menu_KeyboardSearch;
}

int QStyle_SH_TabBar_ElideMode_Type(){
	return QStyle::SH_TabBar_ElideMode;
}

int QStyle_SH_DialogButtonLayout_Type(){
	return QStyle::SH_DialogButtonLayout;
}

int QStyle_SH_ComboBox_PopupFrameStyle_Type(){
	return QStyle::SH_ComboBox_PopupFrameStyle;
}

int QStyle_SH_MessageBox_TextInteractionFlags_Type(){
	return QStyle::SH_MessageBox_TextInteractionFlags;
}

int QStyle_SH_DialogButtonBox_ButtonsHaveIcons_Type(){
	return QStyle::SH_DialogButtonBox_ButtonsHaveIcons;
}

int QStyle_SH_SpellCheckUnderlineStyle_Type(){
	return QStyle::SH_SpellCheckUnderlineStyle;
}

int QStyle_SH_MessageBox_CenterButtons_Type(){
	return QStyle::SH_MessageBox_CenterButtons;
}

int QStyle_SH_Menu_SelectionWrap_Type(){
	return QStyle::SH_Menu_SelectionWrap;
}

int QStyle_SH_ItemView_MovementWithoutUpdatingSelection_Type(){
	return QStyle::SH_ItemView_MovementWithoutUpdatingSelection;
}

int QStyle_SH_ToolTip_Mask_Type(){
	return QStyle::SH_ToolTip_Mask;
}

int QStyle_SH_FocusFrame_AboveWidget_Type(){
	return QStyle::SH_FocusFrame_AboveWidget;
}

int QStyle_SH_TextControl_FocusIndicatorTextCharFormat_Type(){
	return QStyle::SH_TextControl_FocusIndicatorTextCharFormat;
}

int QStyle_SH_WizardStyle_Type(){
	return QStyle::SH_WizardStyle;
}

int QStyle_SH_ItemView_ArrowKeysNavigateIntoChildren_Type(){
	return QStyle::SH_ItemView_ArrowKeysNavigateIntoChildren;
}

int QStyle_SH_Menu_Mask_Type(){
	return QStyle::SH_Menu_Mask;
}

int QStyle_SH_Menu_FlashTriggeredItem_Type(){
	return QStyle::SH_Menu_FlashTriggeredItem;
}

int QStyle_SH_Menu_FadeOutOnHide_Type(){
	return QStyle::SH_Menu_FadeOutOnHide;
}

int QStyle_SH_SpinBox_ClickAutoRepeatThreshold_Type(){
	return QStyle::SH_SpinBox_ClickAutoRepeatThreshold;
}

int QStyle_SH_ItemView_PaintAlternatingRowColorsForEmptyArea_Type(){
	return QStyle::SH_ItemView_PaintAlternatingRowColorsForEmptyArea;
}

int QStyle_SH_FormLayoutWrapPolicy_Type(){
	return QStyle::SH_FormLayoutWrapPolicy;
}

int QStyle_SH_TabWidget_DefaultTabPosition_Type(){
	return QStyle::SH_TabWidget_DefaultTabPosition;
}

int QStyle_SH_ToolBar_Movable_Type(){
	return QStyle::SH_ToolBar_Movable;
}

int QStyle_SH_FormLayoutFieldGrowthPolicy_Type(){
	return QStyle::SH_FormLayoutFieldGrowthPolicy;
}

int QStyle_SH_FormLayoutFormAlignment_Type(){
	return QStyle::SH_FormLayoutFormAlignment;
}

int QStyle_SH_FormLayoutLabelAlignment_Type(){
	return QStyle::SH_FormLayoutLabelAlignment;
}

int QStyle_SH_ItemView_DrawDelegateFrame_Type(){
	return QStyle::SH_ItemView_DrawDelegateFrame;
}

int QStyle_SH_TabBar_CloseButtonPosition_Type(){
	return QStyle::SH_TabBar_CloseButtonPosition;
}

int QStyle_SH_DockWidget_ButtonsHaveFrame_Type(){
	return QStyle::SH_DockWidget_ButtonsHaveFrame;
}

int QStyle_SH_ToolButtonStyle_Type(){
	return QStyle::SH_ToolButtonStyle;
}

int QStyle_SH_RequestSoftwareInputPanel_Type(){
	return QStyle::SH_RequestSoftwareInputPanel;
}

int QStyle_SH_ScrollBar_Transient_Type(){
	return QStyle::SH_ScrollBar_Transient;
}

int QStyle_SH_Menu_SupportsSections_Type(){
	return QStyle::SH_Menu_SupportsSections;
}

int QStyle_SH_ToolTip_WakeUpDelay_Type(){
	return QStyle::SH_ToolTip_WakeUpDelay;
}

int QStyle_SH_ToolTip_FallAsleepDelay_Type(){
	return QStyle::SH_ToolTip_FallAsleepDelay;
}

int QStyle_SH_Widget_Animate_Type(){
	return QStyle::SH_Widget_Animate;
}

int QStyle_SH_Splitter_OpaqueResize_Type(){
	return QStyle::SH_Splitter_OpaqueResize;
}

int QStyle_SH_ComboBox_UseNativePopup_Type(){
	return QStyle::SH_ComboBox_UseNativePopup;
}

int QStyle_SH_LineEdit_PasswordMaskDelay_Type(){
	return QStyle::SH_LineEdit_PasswordMaskDelay;
}

int QStyle_SH_TabBar_ChangeCurrentDelay_Type(){
	return QStyle::SH_TabBar_ChangeCurrentDelay;
}

int QStyle_SH_Menu_SubMenuUniDirection_Type(){
	return QStyle::SH_Menu_SubMenuUniDirection;
}

int QStyle_SH_Menu_SubMenuUniDirectionFailCount_Type(){
	return QStyle::SH_Menu_SubMenuUniDirectionFailCount;
}

int QStyle_SH_Menu_SubMenuSloppySelectOtherActions_Type(){
	return QStyle::SH_Menu_SubMenuSloppySelectOtherActions;
}

int QStyle_SH_Menu_SubMenuSloppyCloseTimeout_Type(){
	return QStyle::SH_Menu_SubMenuSloppyCloseTimeout;
}

int QStyle_SH_Menu_SubMenuResetWhenReenteringParent_Type(){
	return QStyle::SH_Menu_SubMenuResetWhenReenteringParent;
}

int QStyle_SH_Menu_SubMenuDontStartSloppyOnLeave_Type(){
	return QStyle::SH_Menu_SubMenuDontStartSloppyOnLeave;
}

int QStyle_SE_TabBarTearIndicator_Type(){
	return QStyle::SE_TabBarTearIndicator;
}

int QStyle_SE_TreeViewDisclosureItem_Type(){
	return QStyle::SE_TreeViewDisclosureItem;
}

int QStyle_SE_LineEditContents_Type(){
	return QStyle::SE_LineEditContents;
}

int QStyle_SE_FrameContents_Type(){
	return QStyle::SE_FrameContents;
}

int QStyle_SE_DockWidgetCloseButton_Type(){
	return QStyle::SE_DockWidgetCloseButton;
}

int QStyle_SE_DockWidgetFloatButton_Type(){
	return QStyle::SE_DockWidgetFloatButton;
}

int QStyle_SE_DockWidgetTitleBarText_Type(){
	return QStyle::SE_DockWidgetTitleBarText;
}

int QStyle_SE_DockWidgetIcon_Type(){
	return QStyle::SE_DockWidgetIcon;
}

int QStyle_SE_CheckBoxLayoutItem_Type(){
	return QStyle::SE_CheckBoxLayoutItem;
}

int QStyle_SE_ComboBoxLayoutItem_Type(){
	return QStyle::SE_ComboBoxLayoutItem;
}

int QStyle_SE_DateTimeEditLayoutItem_Type(){
	return QStyle::SE_DateTimeEditLayoutItem;
}

int QStyle_SE_DialogButtonBoxLayoutItem_Type(){
	return QStyle::SE_DialogButtonBoxLayoutItem;
}

int QStyle_SE_LabelLayoutItem_Type(){
	return QStyle::SE_LabelLayoutItem;
}

int QStyle_SE_ProgressBarLayoutItem_Type(){
	return QStyle::SE_ProgressBarLayoutItem;
}

int QStyle_SE_PushButtonLayoutItem_Type(){
	return QStyle::SE_PushButtonLayoutItem;
}

int QStyle_SE_RadioButtonLayoutItem_Type(){
	return QStyle::SE_RadioButtonLayoutItem;
}

int QStyle_SE_SliderLayoutItem_Type(){
	return QStyle::SE_SliderLayoutItem;
}

int QStyle_SE_SpinBoxLayoutItem_Type(){
	return QStyle::SE_SpinBoxLayoutItem;
}

int QStyle_SE_ToolButtonLayoutItem_Type(){
	return QStyle::SE_ToolButtonLayoutItem;
}

int QStyle_SE_FrameLayoutItem_Type(){
	return QStyle::SE_FrameLayoutItem;
}

int QStyle_SE_GroupBoxLayoutItem_Type(){
	return QStyle::SE_GroupBoxLayoutItem;
}

int QStyle_SE_TabWidgetLayoutItem_Type(){
	return QStyle::SE_TabWidgetLayoutItem;
}

int QStyle_SE_ItemViewItemDecoration_Type(){
	return QStyle::SE_ItemViewItemDecoration;
}

int QStyle_SE_ItemViewItemText_Type(){
	return QStyle::SE_ItemViewItemText;
}

int QStyle_SE_ItemViewItemFocusRect_Type(){
	return QStyle::SE_ItemViewItemFocusRect;
}

int QStyle_SE_TabBarTabLeftButton_Type(){
	return QStyle::SE_TabBarTabLeftButton;
}

int QStyle_SE_TabBarTabRightButton_Type(){
	return QStyle::SE_TabBarTabRightButton;
}

int QStyle_SE_TabBarTabText_Type(){
	return QStyle::SE_TabBarTabText;
}

int QStyle_SE_ShapedFrameContents_Type(){
	return QStyle::SE_ShapedFrameContents;
}

int QStyle_SE_ToolBarHandle_Type(){
	return QStyle::SE_ToolBarHandle;
}

void QStyle_DrawItemPixmap(void* ptr, void* painter, void* rectangle, int alignment, void* pixmap){
	static_cast<QStyle*>(ptr)->drawItemPixmap(static_cast<QPainter*>(painter), *static_cast<QRect*>(rectangle), alignment, *static_cast<QPixmap*>(pixmap));
}

void QStyle_DrawItemText(void* ptr, void* painter, void* rectangle, int alignment, void* palette, int enabled, char* text, int textRole){
	static_cast<QStyle*>(ptr)->drawItemText(static_cast<QPainter*>(painter), *static_cast<QRect*>(rectangle), alignment, *static_cast<QPalette*>(palette), enabled != 0, QString(text), static_cast<QPalette::ColorRole>(textRole));
}

void QStyle_Polish2(void* ptr, void* application){
	static_cast<QStyle*>(ptr)->polish(static_cast<QApplication*>(application));
}

void QStyle_Polish3(void* ptr, void* palette){
	static_cast<QStyle*>(ptr)->polish(*static_cast<QPalette*>(palette));
}

void* QStyle_Proxy(void* ptr){
	return const_cast<QStyle*>(static_cast<QStyle*>(ptr)->proxy());
}

int QStyle_QStyle_SliderValueFromPosition(int min, int max, int position, int span, int upsideDown){
	return QStyle::sliderValueFromPosition(min, max, position, span, upsideDown != 0);
}

void QStyle_Unpolish2(void* ptr, void* application){
	static_cast<QStyle*>(ptr)->unpolish(static_cast<QApplication*>(application));
}

int QStyle_CombinedLayoutSpacing(void* ptr, int controls1, int controls2, int orientation, void* option, void* widget){
	return static_cast<QStyle*>(ptr)->combinedLayoutSpacing(static_cast<QSizePolicy::ControlType>(controls1), static_cast<QSizePolicy::ControlType>(controls2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void QStyle_DrawComplexControl(void* ptr, int control, void* option, void* painter, void* widget){
	static_cast<QStyle*>(ptr)->drawComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QStyle_DrawControl(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QStyle*>(ptr)->drawControl(static_cast<QStyle::ControlElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

void QStyle_DrawPrimitive(void* ptr, int element, void* option, void* painter, void* widget){
	static_cast<QStyle*>(ptr)->drawPrimitive(static_cast<QStyle::PrimitiveElement>(element), static_cast<QStyleOption*>(option), static_cast<QPainter*>(painter), static_cast<QWidget*>(widget));
}

int QStyle_HitTestComplexControl(void* ptr, int control, void* option, void* position, void* widget){
	return static_cast<QStyle*>(ptr)->hitTestComplexControl(static_cast<QStyle::ComplexControl>(control), static_cast<QStyleOptionComplex*>(option), *static_cast<QPoint*>(position), static_cast<QWidget*>(widget));
}

int QStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget){
	return static_cast<QStyle*>(ptr)->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

int QStyle_PixelMetric(void* ptr, int metric, void* option, void* widget){
	return static_cast<QStyle*>(ptr)->pixelMetric(static_cast<QStyle::PixelMetric>(metric), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget));
}

void QStyle_Polish(void* ptr, void* widget){
	static_cast<QStyle*>(ptr)->polish(static_cast<QWidget*>(widget));
}

int QStyle_QStyle_SliderPositionFromValue(int min, int max, int logicalValue, int span, int upsideDown){
	return QStyle::sliderPositionFromValue(min, max, logicalValue, span, upsideDown != 0);
}

int QStyle_StyleHint(void* ptr, int hint, void* option, void* widget, void* returnData){
	return static_cast<QStyle*>(ptr)->styleHint(static_cast<QStyle::StyleHint>(hint), static_cast<QStyleOption*>(option), static_cast<QWidget*>(widget), static_cast<QStyleHintReturn*>(returnData));
}

void QStyle_Unpolish(void* ptr, void* widget){
	static_cast<QStyle*>(ptr)->unpolish(static_cast<QWidget*>(widget));
}

int QStyle_QStyle_VisualAlignment(int direction, int alignment){
	return QStyle::visualAlignment(static_cast<Qt::LayoutDirection>(direction), static_cast<Qt::AlignmentFlag>(alignment));
}

void QStyle_DestroyQStyle(void* ptr){
	static_cast<QStyle*>(ptr)->~QStyle();
}

#include "qtreewidgetitemiterator.h"
#include <QTreeWidgetItem>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTreeWidget>
#include <QTreeWidgetItemIterator>
#include "_cgo_export.h"

#include "qstyleoptiondockwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionDockWidget>
#include "_cgo_export.h"

class MyQStyleOptionDockWidget: public QStyleOptionDockWidget {
public:
};

void* QStyleOptionDockWidget_NewQStyleOptionDockWidget(){
	return new QStyleOptionDockWidget();
}

void* QStyleOptionDockWidget_NewQStyleOptionDockWidget2(void* other){
	return new QStyleOptionDockWidget(*static_cast<QStyleOptionDockWidget*>(other));
}

#include "qgraphicsscene.h"
#include <QBrush>
#include <QWidget>
#include <QGraphicsItem>
#include <QUrl>
#include <QRect>
#include <QEvent>
#include <QPainter>
#include <QPolygonF>
#include <QPixmap>
#include <QString>
#include <QPalette>
#include <QFont>
#include <QMetaObject>
#include <QTransform>
#include <QStyle>
#include <QPen>
#include <QVariant>
#include <QObject>
#include <QLineF>
#include <QPolygon>
#include <QModelIndex>
#include <QRectF>
#include <QPoint>
#include <QGraphicsItemGroup>
#include <QPointF>
#include <QLine>
#include <QGraphicsWidget>
#include <QPainterPath>
#include <QGraphicsScene>
#include "_cgo_export.h"

class MyQGraphicsScene: public QGraphicsScene {
public:
void Signal_FocusItemChanged(QGraphicsItem * newFocusItem, QGraphicsItem * oldFocusItem, Qt::FocusReason reason){callbackQGraphicsSceneFocusItemChanged(this->objectName().toUtf8().data(), newFocusItem, oldFocusItem, reason);};
void Signal_SelectionChanged(){callbackQGraphicsSceneSelectionChanged(this->objectName().toUtf8().data());};
};

void* QGraphicsScene_BackgroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsScene*>(ptr)->backgroundBrush());
}

int QGraphicsScene_BspTreeDepth(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->bspTreeDepth();
}

void* QGraphicsScene_ForegroundBrush(void* ptr){
	return new QBrush(static_cast<QGraphicsScene*>(ptr)->foregroundBrush());
}

int QGraphicsScene_IsSortCacheEnabled(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->isSortCacheEnabled();
}

int QGraphicsScene_ItemIndexMethod(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->itemIndexMethod();
}

double QGraphicsScene_MinimumRenderSize(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->minimumRenderSize());
}

void QGraphicsScene_SetBackgroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsScene*>(ptr)->setBackgroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetBspTreeDepth(void* ptr, int depth){
	static_cast<QGraphicsScene*>(ptr)->setBspTreeDepth(depth);
}

void QGraphicsScene_SetFont(void* ptr, void* font){
	static_cast<QGraphicsScene*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsScene_SetForegroundBrush(void* ptr, void* brush){
	static_cast<QGraphicsScene*>(ptr)->setForegroundBrush(*static_cast<QBrush*>(brush));
}

void QGraphicsScene_SetItemIndexMethod(void* ptr, int method){
	static_cast<QGraphicsScene*>(ptr)->setItemIndexMethod(static_cast<QGraphicsScene::ItemIndexMethod>(method));
}

void QGraphicsScene_SetMinimumRenderSize(void* ptr, double minSize){
	static_cast<QGraphicsScene*>(ptr)->setMinimumRenderSize(static_cast<qreal>(minSize));
}

void QGraphicsScene_SetPalette(void* ptr, void* palette){
	static_cast<QGraphicsScene*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsScene_SetSceneRect(void* ptr, void* rect){
	static_cast<QGraphicsScene*>(ptr)->setSceneRect(*static_cast<QRectF*>(rect));
}

void QGraphicsScene_SetSortCacheEnabled(void* ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setSortCacheEnabled(enabled != 0);
}

void QGraphicsScene_SetStickyFocus(void* ptr, int enabled){
	static_cast<QGraphicsScene*>(ptr)->setStickyFocus(enabled != 0);
}

int QGraphicsScene_StickyFocus(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->stickyFocus();
}

void QGraphicsScene_Update(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "update", Q_ARG(QRectF, *static_cast<QRectF*>(rect)));
}

void* QGraphicsScene_NewQGraphicsScene(void* parent){
	return new QGraphicsScene(static_cast<QObject*>(parent));
}

void* QGraphicsScene_NewQGraphicsScene2(void* sceneRect, void* parent){
	return new QGraphicsScene(*static_cast<QRectF*>(sceneRect), static_cast<QObject*>(parent));
}

void* QGraphicsScene_NewQGraphicsScene3(double x, double y, double width, double height, void* parent){
	return new QGraphicsScene(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height), static_cast<QObject*>(parent));
}

void* QGraphicsScene_ActivePanel(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->activePanel();
}

void* QGraphicsScene_ActiveWindow(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->activeWindow();
}

void* QGraphicsScene_AddEllipse(void* ptr, void* rect, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addEllipse(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddEllipse2(void* ptr, double x, double y, double w, double h, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addEllipse(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void QGraphicsScene_AddItem(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->addItem(static_cast<QGraphicsItem*>(item));
}

void* QGraphicsScene_AddLine(void* ptr, void* line, void* pen){
	return static_cast<QGraphicsScene*>(ptr)->addLine(*static_cast<QLineF*>(line), *static_cast<QPen*>(pen));
}

void* QGraphicsScene_AddLine2(void* ptr, double x1, double y1, double x2, double y2, void* pen){
	return static_cast<QGraphicsScene*>(ptr)->addLine(static_cast<qreal>(x1), static_cast<qreal>(y1), static_cast<qreal>(x2), static_cast<qreal>(y2), *static_cast<QPen*>(pen));
}

void* QGraphicsScene_AddPath(void* ptr, void* path, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addPath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddPixmap(void* ptr, void* pixmap){
	return static_cast<QGraphicsScene*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QGraphicsScene_AddPolygon(void* ptr, void* polygon, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddRect(void* ptr, void* rect, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addRect(*static_cast<QRectF*>(rect), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddRect2(void* ptr, double x, double y, double w, double h, void* pen, void* brush){
	return static_cast<QGraphicsScene*>(ptr)->addRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), *static_cast<QPen*>(pen), *static_cast<QBrush*>(brush));
}

void* QGraphicsScene_AddSimpleText(void* ptr, char* text, void* font){
	return static_cast<QGraphicsScene*>(ptr)->addSimpleText(QString(text), *static_cast<QFont*>(font));
}

void* QGraphicsScene_AddText(void* ptr, char* text, void* font){
	return static_cast<QGraphicsScene*>(ptr)->addText(QString(text), *static_cast<QFont*>(font));
}

void* QGraphicsScene_AddWidget(void* ptr, void* widget, int wFlags){
	return static_cast<QGraphicsScene*>(ptr)->addWidget(static_cast<QWidget*>(widget), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsScene_Advance(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "advance");
}

void QGraphicsScene_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clear");
}

void QGraphicsScene_ClearFocus(void* ptr){
	static_cast<QGraphicsScene*>(ptr)->clearFocus();
}

void QGraphicsScene_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "clearSelection");
}

void QGraphicsScene_DestroyItemGroup(void* ptr, void* group){
	static_cast<QGraphicsScene*>(ptr)->destroyItemGroup(static_cast<QGraphicsItemGroup*>(group));
}

void* QGraphicsScene_FocusItem(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->focusItem();
}

void QGraphicsScene_ConnectFocusItemChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

void QGraphicsScene_DisconnectFocusItemChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&QGraphicsScene::focusItemChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)(QGraphicsItem *, QGraphicsItem *, Qt::FocusReason)>(&MyQGraphicsScene::Signal_FocusItemChanged));;
}

int QGraphicsScene_HasFocus(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->hasFocus();
}

double QGraphicsScene_Height(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->height());
}

void* QGraphicsScene_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QGraphicsScene*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QGraphicsScene_Invalidate(void* ptr, void* rect, int layers){
	QMetaObject::invokeMethod(static_cast<QGraphicsScene*>(ptr), "invalidate", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(QGraphicsScene::SceneLayer, static_cast<QGraphicsScene::SceneLayer>(layers)));
}

void QGraphicsScene_Invalidate2(void* ptr, double x, double y, double w, double h, int layers){
	static_cast<QGraphicsScene*>(ptr)->invalidate(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), static_cast<QGraphicsScene::SceneLayer>(layers));
}

int QGraphicsScene_IsActive(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->isActive();
}

void* QGraphicsScene_ItemAt(void* ptr, void* position, void* deviceTransform){
	return static_cast<QGraphicsScene*>(ptr)->itemAt(*static_cast<QPointF*>(position), *static_cast<QTransform*>(deviceTransform));
}

void* QGraphicsScene_ItemAt3(void* ptr, double x, double y, void* deviceTransform){
	return static_cast<QGraphicsScene*>(ptr)->itemAt(static_cast<qreal>(x), static_cast<qreal>(y), *static_cast<QTransform*>(deviceTransform));
}

void* QGraphicsScene_MouseGrabberItem(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->mouseGrabberItem();
}

void QGraphicsScene_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->removeItem(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_Render(void* ptr, void* painter, void* target, void* source, int aspectRatioMode){
	static_cast<QGraphicsScene*>(ptr)->render(static_cast<QPainter*>(painter), *static_cast<QRectF*>(target), *static_cast<QRectF*>(source), static_cast<Qt::AspectRatioMode>(aspectRatioMode));
}

void QGraphicsScene_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

void QGraphicsScene_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsScene*>(ptr), static_cast<void (QGraphicsScene::*)()>(&QGraphicsScene::selectionChanged), static_cast<MyQGraphicsScene*>(ptr), static_cast<void (MyQGraphicsScene::*)()>(&MyQGraphicsScene::Signal_SelectionChanged));;
}

int QGraphicsScene_SendEvent(void* ptr, void* item, void* event){
	return static_cast<QGraphicsScene*>(ptr)->sendEvent(static_cast<QGraphicsItem*>(item), static_cast<QEvent*>(event));
}

void QGraphicsScene_SetActivePanel(void* ptr, void* item){
	static_cast<QGraphicsScene*>(ptr)->setActivePanel(static_cast<QGraphicsItem*>(item));
}

void QGraphicsScene_SetActiveWindow(void* ptr, void* widget){
	static_cast<QGraphicsScene*>(ptr)->setActiveWindow(static_cast<QGraphicsWidget*>(widget));
}

void QGraphicsScene_SetFocus(void* ptr, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocus(static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetFocusItem(void* ptr, void* item, int focusReason){
	static_cast<QGraphicsScene*>(ptr)->setFocusItem(static_cast<QGraphicsItem*>(item), static_cast<Qt::FocusReason>(focusReason));
}

void QGraphicsScene_SetSceneRect2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsScene*>(ptr)->setSceneRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsScene_SetSelectionArea2(void* ptr, void* path, int mode, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea3(void* ptr, void* path, int selectionOperation, int mode, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), static_cast<Qt::ItemSelectionOperation>(selectionOperation), static_cast<Qt::ItemSelectionMode>(mode), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetSelectionArea(void* ptr, void* path, void* deviceTransform){
	static_cast<QGraphicsScene*>(ptr)->setSelectionArea(*static_cast<QPainterPath*>(path), *static_cast<QTransform*>(deviceTransform));
}

void QGraphicsScene_SetStyle(void* ptr, void* style){
	static_cast<QGraphicsScene*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void* QGraphicsScene_Style(void* ptr){
	return static_cast<QGraphicsScene*>(ptr)->style();
}

void QGraphicsScene_Update2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsScene*>(ptr)->update(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

double QGraphicsScene_Width(void* ptr){
	return static_cast<double>(static_cast<QGraphicsScene*>(ptr)->width());
}

void QGraphicsScene_DestroyQGraphicsScene(void* ptr){
	static_cast<QGraphicsScene*>(ptr)->~QGraphicsScene();
}

#include "qtoolbox.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QIcon>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QToolBox>
#include "_cgo_export.h"

class MyQToolBox: public QToolBox {
public:
void Signal_CurrentChanged(int index){callbackQToolBoxCurrentChanged(this->objectName().toUtf8().data(), index);};
};

int QToolBox_Count(void* ptr){
	return static_cast<QToolBox*>(ptr)->count();
}

int QToolBox_CurrentIndex(void* ptr){
	return static_cast<QToolBox*>(ptr)->currentIndex();
}

void QToolBox_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void* QToolBox_NewQToolBox(void* parent, int f){
	return new QToolBox(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QToolBox_AddItem2(void* ptr, void* w, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(w), QString(text));
}

int QToolBox_AddItem(void* ptr, void* widget, void* iconSet, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(widget), *static_cast<QIcon*>(iconSet), QString(text));
}

void QToolBox_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

void QToolBox_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

void* QToolBox_CurrentWidget(void* ptr){
	return static_cast<QToolBox*>(ptr)->currentWidget();
}

int QToolBox_IndexOf(void* ptr, void* widget){
	return static_cast<QToolBox*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QToolBox_InsertItem(void* ptr, int index, void* widget, void* icon, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), *static_cast<QIcon*>(icon), QString(text));
}

int QToolBox_InsertItem2(void* ptr, int index, void* widget, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), QString(text));
}

int QToolBox_IsItemEnabled(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->isItemEnabled(index);
}

char* QToolBox_ItemText(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemText(index).toUtf8().data();
}

char* QToolBox_ItemToolTip(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemToolTip(index).toUtf8().data();
}

void QToolBox_RemoveItem(void* ptr, int index){
	static_cast<QToolBox*>(ptr)->removeItem(index);
}

void QToolBox_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QToolBox_SetItemEnabled(void* ptr, int index, int enabled){
	static_cast<QToolBox*>(ptr)->setItemEnabled(index, enabled != 0);
}

void QToolBox_SetItemIcon(void* ptr, int index, void* icon){
	static_cast<QToolBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QToolBox_SetItemText(void* ptr, int index, char* text){
	static_cast<QToolBox*>(ptr)->setItemText(index, QString(text));
}

void QToolBox_SetItemToolTip(void* ptr, int index, char* toolTip){
	static_cast<QToolBox*>(ptr)->setItemToolTip(index, QString(toolTip));
}

void* QToolBox_Widget(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->widget(index);
}

void QToolBox_DestroyQToolBox(void* ptr){
	static_cast<QToolBox*>(ptr)->~QToolBox();
}

#include "qgraphicswidget.h"
#include <QVariant>
#include <QModelIndex>
#include <QPalette>
#include <QPainter>
#include <QStyle>
#include <QSizeF>
#include <QKeySequence>
#include <QRectF>
#include <QString>
#include <QObject>
#include <QAction>
#include <QGraphicsItem>
#include <QSize>
#include <QFont>
#include <QMetaObject>
#include <QUrl>
#include <QWidget>
#include <QGraphicsLayout>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QRect>
#include <QGraphicsWidget>
#include "_cgo_export.h"

class MyQGraphicsWidget: public QGraphicsWidget {
public:
void Signal_GeometryChanged(){callbackQGraphicsWidgetGeometryChanged(this->objectName().toUtf8().data());};
};

int QGraphicsWidget_AutoFillBackground(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->autoFillBackground();
}

int QGraphicsWidget_FocusPolicy(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusPolicy();
}

int QGraphicsWidget_LayoutDirection(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layoutDirection();
}

void QGraphicsWidget_Resize(void* ptr, void* size){
	static_cast<QGraphicsWidget*>(ptr)->resize(*static_cast<QSizeF*>(size));
}

void QGraphicsWidget_SetAutoFillBackground(void* ptr, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setAutoFillBackground(enabled != 0);
}

void QGraphicsWidget_SetFocusPolicy(void* ptr, int policy){
	static_cast<QGraphicsWidget*>(ptr)->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

void QGraphicsWidget_SetFont(void* ptr, void* font){
	static_cast<QGraphicsWidget*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsWidget_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsWidget*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsWidget_SetLayout(void* ptr, void* layout){
	static_cast<QGraphicsWidget*>(ptr)->setLayout(static_cast<QGraphicsLayout*>(layout));
}

void QGraphicsWidget_SetLayoutDirection(void* ptr, int direction){
	static_cast<QGraphicsWidget*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGraphicsWidget_SetPalette(void* ptr, void* palette){
	static_cast<QGraphicsWidget*>(ptr)->setPalette(*static_cast<QPalette*>(palette));
}

void QGraphicsWidget_SetWindowFlags(void* ptr, int wFlags){
	static_cast<QGraphicsWidget*>(ptr)->setWindowFlags(static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_SetWindowTitle(void* ptr, char* title){
	static_cast<QGraphicsWidget*>(ptr)->setWindowTitle(QString(title));
}

void QGraphicsWidget_UnsetLayoutDirection(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetLayoutDirection();
}

int QGraphicsWidget_WindowFlags(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowFlags();
}

char* QGraphicsWidget_WindowTitle(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowTitle().toUtf8().data();
}

void* QGraphicsWidget_NewQGraphicsWidget(void* parent, int wFlags){
	return new QGraphicsWidget(static_cast<QGraphicsItem*>(parent), static_cast<Qt::WindowType>(wFlags));
}

void QGraphicsWidget_AddAction(void* ptr, void* action){
	static_cast<QGraphicsWidget*>(ptr)->addAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_AdjustSize(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->adjustSize();
}

int QGraphicsWidget_Close(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QGraphicsWidget*>(ptr), "close");
}

void* QGraphicsWidget_FocusWidget(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->focusWidget();
}

void QGraphicsWidget_ConnectGeometryChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

void QGraphicsWidget_DisconnectGeometryChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsWidget*>(ptr), static_cast<void (QGraphicsWidget::*)()>(&QGraphicsWidget::geometryChanged), static_cast<MyQGraphicsWidget*>(ptr), static_cast<void (MyQGraphicsWidget::*)()>(&MyQGraphicsWidget::Signal_GeometryChanged));;
}

int QGraphicsWidget_GrabShortcut(void* ptr, void* sequence, int context){
	return static_cast<QGraphicsWidget*>(ptr)->grabShortcut(*static_cast<QKeySequence*>(sequence), static_cast<Qt::ShortcutContext>(context));
}

void QGraphicsWidget_InsertAction(void* ptr, void* before, void* action){
	static_cast<QGraphicsWidget*>(ptr)->insertAction(static_cast<QAction*>(before), static_cast<QAction*>(action));
}

int QGraphicsWidget_IsActiveWindow(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->isActiveWindow();
}

void* QGraphicsWidget_Layout(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->layout();
}

void QGraphicsWidget_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsWidget*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_PaintWindowFrame(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsWidget*>(ptr)->paintWindowFrame(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsWidget_ReleaseShortcut(void* ptr, int id){
	static_cast<QGraphicsWidget*>(ptr)->releaseShortcut(id);
}

void QGraphicsWidget_RemoveAction(void* ptr, void* action){
	static_cast<QGraphicsWidget*>(ptr)->removeAction(static_cast<QAction*>(action));
}

void QGraphicsWidget_Resize2(void* ptr, double w, double h){
	static_cast<QGraphicsWidget*>(ptr)->resize(static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsWidget_SetAttribute(void* ptr, int attribute, int on){
	static_cast<QGraphicsWidget*>(ptr)->setAttribute(static_cast<Qt::WidgetAttribute>(attribute), on != 0);
}

void QGraphicsWidget_SetContentsMargins(void* ptr, double left, double top, double right, double bottom){
	static_cast<QGraphicsWidget*>(ptr)->setContentsMargins(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(right), static_cast<qreal>(bottom));
}

void QGraphicsWidget_SetGeometry2(void* ptr, double x, double y, double w, double h){
	static_cast<QGraphicsWidget*>(ptr)->setGeometry(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h));
}

void QGraphicsWidget_SetShortcutAutoRepeat(void* ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutAutoRepeat(id, enabled != 0);
}

void QGraphicsWidget_SetShortcutEnabled(void* ptr, int id, int enabled){
	static_cast<QGraphicsWidget*>(ptr)->setShortcutEnabled(id, enabled != 0);
}

void QGraphicsWidget_SetStyle(void* ptr, void* style){
	static_cast<QGraphicsWidget*>(ptr)->setStyle(static_cast<QStyle*>(style));
}

void QGraphicsWidget_QGraphicsWidget_SetTabOrder(void* first, void* second){
	QGraphicsWidget::setTabOrder(static_cast<QGraphicsWidget*>(first), static_cast<QGraphicsWidget*>(second));
}

void QGraphicsWidget_SetWindowFrameMargins(void* ptr, double left, double top, double right, double bottom){
	static_cast<QGraphicsWidget*>(ptr)->setWindowFrameMargins(static_cast<qreal>(left), static_cast<qreal>(top), static_cast<qreal>(right), static_cast<qreal>(bottom));
}

void* QGraphicsWidget_Style(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->style();
}

int QGraphicsWidget_TestAttribute(void* ptr, int attribute){
	return static_cast<QGraphicsWidget*>(ptr)->testAttribute(static_cast<Qt::WidgetAttribute>(attribute));
}

int QGraphicsWidget_Type(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->type();
}

void QGraphicsWidget_UnsetWindowFrameMargins(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->unsetWindowFrameMargins();
}

int QGraphicsWidget_WindowType(void* ptr){
	return static_cast<QGraphicsWidget*>(ptr)->windowType();
}

void QGraphicsWidget_DestroyQGraphicsWidget(void* ptr){
	static_cast<QGraphicsWidget*>(ptr)->~QGraphicsWidget();
}

#include "qwizardpage.h"
#include <QWizard>
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QWizardPage>
#include "_cgo_export.h"

class MyQWizardPage: public QWizardPage {
public:
void Signal_CompleteChanged(){callbackQWizardPageCompleteChanged(this->objectName().toUtf8().data());};
};

void QWizardPage_SetSubTitle(void* ptr, char* subTitle){
	static_cast<QWizardPage*>(ptr)->setSubTitle(QString(subTitle));
}

void QWizardPage_SetTitle(void* ptr, char* title){
	static_cast<QWizardPage*>(ptr)->setTitle(QString(title));
}

char* QWizardPage_SubTitle(void* ptr){
	return static_cast<QWizardPage*>(ptr)->subTitle().toUtf8().data();
}

char* QWizardPage_Title(void* ptr){
	return static_cast<QWizardPage*>(ptr)->title().toUtf8().data();
}

void* QWizardPage_NewQWizardPage(void* parent){
	return new QWizardPage(static_cast<QWidget*>(parent));
}

char* QWizardPage_ButtonText(void* ptr, int which){
	return static_cast<QWizardPage*>(ptr)->buttonText(static_cast<QWizard::WizardButton>(which)).toUtf8().data();
}

void QWizardPage_CleanupPage(void* ptr){
	static_cast<QWizardPage*>(ptr)->cleanupPage();
}

void QWizardPage_ConnectCompleteChanged(void* ptr){
	QObject::connect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_DisconnectCompleteChanged(void* ptr){
	QObject::disconnect(static_cast<QWizardPage*>(ptr), static_cast<void (QWizardPage::*)()>(&QWizardPage::completeChanged), static_cast<MyQWizardPage*>(ptr), static_cast<void (MyQWizardPage::*)()>(&MyQWizardPage::Signal_CompleteChanged));;
}

void QWizardPage_InitializePage(void* ptr){
	static_cast<QWizardPage*>(ptr)->initializePage();
}

int QWizardPage_IsCommitPage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isCommitPage();
}

int QWizardPage_IsComplete(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isComplete();
}

int QWizardPage_IsFinalPage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->isFinalPage();
}

int QWizardPage_NextId(void* ptr){
	return static_cast<QWizardPage*>(ptr)->nextId();
}

void QWizardPage_SetButtonText(void* ptr, int which, char* text){
	static_cast<QWizardPage*>(ptr)->setButtonText(static_cast<QWizard::WizardButton>(which), QString(text));
}

void QWizardPage_SetCommitPage(void* ptr, int commitPage){
	static_cast<QWizardPage*>(ptr)->setCommitPage(commitPage != 0);
}

void QWizardPage_SetFinalPage(void* ptr, int finalPage){
	static_cast<QWizardPage*>(ptr)->setFinalPage(finalPage != 0);
}

void QWizardPage_SetPixmap(void* ptr, int which, void* pixmap){
	static_cast<QWizardPage*>(ptr)->setPixmap(static_cast<QWizard::WizardPixmap>(which), *static_cast<QPixmap*>(pixmap));
}

int QWizardPage_ValidatePage(void* ptr){
	return static_cast<QWizardPage*>(ptr)->validatePage();
}

void QWizardPage_DestroyQWizardPage(void* ptr){
	static_cast<QWizardPage*>(ptr)->~QWizardPage();
}

#include "qtablewidgetselectionrange.h"
#include <QUrl>
#include <QModelIndex>
#include <QTableWidget>
#include <QString>
#include <QVariant>
#include <QTableWidgetSelectionRange>
#include "_cgo_export.h"

class MyQTableWidgetSelectionRange: public QTableWidgetSelectionRange {
public:
};

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange(){
	return new QTableWidgetSelectionRange();
}

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange3(void* other){
	return new QTableWidgetSelectionRange(*static_cast<QTableWidgetSelectionRange*>(other));
}

void* QTableWidgetSelectionRange_NewQTableWidgetSelectionRange2(int top, int left, int bottom, int right){
	return new QTableWidgetSelectionRange(top, left, bottom, right);
}

int QTableWidgetSelectionRange_BottomRow(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->bottomRow();
}

int QTableWidgetSelectionRange_ColumnCount(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->columnCount();
}

int QTableWidgetSelectionRange_LeftColumn(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->leftColumn();
}

int QTableWidgetSelectionRange_RightColumn(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rightColumn();
}

int QTableWidgetSelectionRange_RowCount(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->rowCount();
}

int QTableWidgetSelectionRange_TopRow(void* ptr){
	return static_cast<QTableWidgetSelectionRange*>(ptr)->topRow();
}

void QTableWidgetSelectionRange_DestroyQTableWidgetSelectionRange(void* ptr){
	static_cast<QTableWidgetSelectionRange*>(ptr)->~QTableWidgetSelectionRange();
}

#include "qlistwidget.h"
#include <QList>
#include <QPoint>
#include <QItemSelection>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QListWidgetItem>
#include <QWidget>
#include <QDropEvent>
#include <QVariant>
#include <QAbstractItemView>
#include <QList>
#include <QString>
#include <QItemSelectionModel>
#include <QObject>
#include <QListWidget>
#include "_cgo_export.h"

class MyQListWidget: public QListWidget {
public:
void Signal_CurrentItemChanged(QListWidgetItem * current, QListWidgetItem * previous){callbackQListWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_CurrentRowChanged(int currentRow){callbackQListWidgetCurrentRowChanged(this->objectName().toUtf8().data(), currentRow);};
void Signal_CurrentTextChanged(const QString & currentText){callbackQListWidgetCurrentTextChanged(this->objectName().toUtf8().data(), currentText.toUtf8().data());};
void Signal_ItemActivated(QListWidgetItem * item){callbackQListWidgetItemActivated(this->objectName().toUtf8().data(), item);};
void Signal_ItemChanged(QListWidgetItem * item){callbackQListWidgetItemChanged(this->objectName().toUtf8().data(), item);};
void Signal_ItemClicked(QListWidgetItem * item){callbackQListWidgetItemClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QListWidgetItem * item){callbackQListWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemEntered(QListWidgetItem * item){callbackQListWidgetItemEntered(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QListWidgetItem * item){callbackQListWidgetItemPressed(this->objectName().toUtf8().data(), item);};
void Signal_ItemSelectionChanged(){callbackQListWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QListWidget_Count(void* ptr){
	return static_cast<QListWidget*>(ptr)->count();
}

int QListWidget_CurrentRow(void* ptr){
	return static_cast<QListWidget*>(ptr)->currentRow();
}

int QListWidget_IsSortingEnabled(void* ptr){
	return static_cast<QListWidget*>(ptr)->isSortingEnabled();
}

void QListWidget_SetCurrentRow(void* ptr, int row){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row);
}

void QListWidget_SetSortingEnabled(void* ptr, int enable){
	static_cast<QListWidget*>(ptr)->setSortingEnabled(enable != 0);
}

void* QListWidget_NewQListWidget(void* parent){
	return new QListWidget(static_cast<QWidget*>(parent));
}

void QListWidget_AddItem2(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->addItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_AddItem(void* ptr, char* label){
	static_cast<QListWidget*>(ptr)->addItem(QString(label));
}

void QListWidget_AddItems(void* ptr, char* labels){
	static_cast<QListWidget*>(ptr)->addItems(QString(labels).split("|", QString::SkipEmptyParts));
}

void QListWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "clear");
}

void QListWidget_ClosePersistentEditor(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->closePersistentEditor(static_cast<QListWidgetItem*>(item));
}

void* QListWidget_CurrentItem(void* ptr){
	return static_cast<QListWidget*>(ptr)->currentItem();
}

void QListWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_ConnectCurrentRowChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_DisconnectCurrentRowChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_ConnectCurrentTextChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DisconnectCurrentTextChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DropEvent(void* ptr, void* event){
	static_cast<QListWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QListWidget_EditItem(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->editItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem(void* ptr, int row, void* item){
	static_cast<QListWidget*>(ptr)->insertItem(row, static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem2(void* ptr, int row, char* label){
	static_cast<QListWidget*>(ptr)->insertItem(row, QString(label));
}

void QListWidget_InsertItems(void* ptr, int row, char* labels){
	static_cast<QListWidget*>(ptr)->insertItems(row, QString(labels).split("|", QString::SkipEmptyParts));
}

void* QListWidget_Item(void* ptr, int row){
	return static_cast<QListWidget*>(ptr)->item(row);
}

void QListWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

void QListWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

void* QListWidget_ItemAt(void* ptr, void* p){
	return static_cast<QListWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

void* QListWidget_ItemAt2(void* ptr, int x, int y){
	return static_cast<QListWidget*>(ptr)->itemAt(x, y);
}

void QListWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

void QListWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

void* QListWidget_ItemWidget(void* ptr, void* item){
	return static_cast<QListWidget*>(ptr)->itemWidget(static_cast<QListWidgetItem*>(item));
}

void QListWidget_OpenPersistentEditor(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->openPersistentEditor(static_cast<QListWidgetItem*>(item));
}

void QListWidget_RemoveItemWidget(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->removeItemWidget(static_cast<QListWidgetItem*>(item));
}

int QListWidget_Row(void* ptr, void* item){
	return static_cast<QListWidget*>(ptr)->row(static_cast<QListWidgetItem*>(item));
}

void QListWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "scrollToItem", Q_ARG(QListWidgetItem*, static_cast<QListWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QListWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_SetCurrentItem2(void* ptr, void* item, int command){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetCurrentRow2(void* ptr, int row, int command){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetItemWidget(void* ptr, void* item, void* widget){
	static_cast<QListWidget*>(ptr)->setItemWidget(static_cast<QListWidgetItem*>(item), static_cast<QWidget*>(widget));
}

void QListWidget_SortItems(void* ptr, int order){
	static_cast<QListWidget*>(ptr)->sortItems(static_cast<Qt::SortOrder>(order));
}

void* QListWidget_TakeItem(void* ptr, int row){
	return static_cast<QListWidget*>(ptr)->takeItem(row);
}

void QListWidget_DestroyQListWidget(void* ptr){
	static_cast<QListWidget*>(ptr)->~QListWidget();
}

#include "qgraphicsscenewheelevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QGraphicsSceneWheelEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneWheelEvent: public QGraphicsSceneWheelEvent {
public:
};

int QGraphicsSceneWheelEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->buttons();
}

int QGraphicsSceneWheelEvent_Delta(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->delta();
}

int QGraphicsSceneWheelEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->modifiers();
}

int QGraphicsSceneWheelEvent_Orientation(void* ptr){
	return static_cast<QGraphicsSceneWheelEvent*>(ptr)->orientation();
}

void QGraphicsSceneWheelEvent_DestroyQGraphicsSceneWheelEvent(void* ptr){
	static_cast<QGraphicsSceneWheelEvent*>(ptr)->~QGraphicsSceneWheelEvent();
}

#include "qtilerules.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTileRules>
#include "_cgo_export.h"

class MyQTileRules: public QTileRules {
public:
};

void* QTileRules_NewQTileRules(int horizontalRule, int verticalRule){
	return new QTileRules(static_cast<Qt::TileRule>(horizontalRule), static_cast<Qt::TileRule>(verticalRule));
}

void* QTileRules_NewQTileRules2(int rule){
	return new QTileRules(static_cast<Qt::TileRule>(rule));
}

#include "qdialog.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDial>
#include <QObject>
#include <QWidget>
#include <QMetaObject>
#include <QDialog>
#include "_cgo_export.h"

class MyQDialog: public QDialog {
public:
void Signal_Accepted(){callbackQDialogAccepted(this->objectName().toUtf8().data());};
void Signal_Finished(int result){callbackQDialogFinished(this->objectName().toUtf8().data(), result);};
void Signal_Rejected(){callbackQDialogRejected(this->objectName().toUtf8().data());};
};

int QDialog_IsSizeGripEnabled(void* ptr){
	return static_cast<QDialog*>(ptr)->isSizeGripEnabled();
}

void QDialog_SetModal(void* ptr, int modal){
	static_cast<QDialog*>(ptr)->setModal(modal != 0);
}

void QDialog_SetResult(void* ptr, int i){
	static_cast<QDialog*>(ptr)->setResult(i);
}

void QDialog_SetSizeGripEnabled(void* ptr, int v){
	static_cast<QDialog*>(ptr)->setSizeGripEnabled(v != 0);
}

void* QDialog_NewQDialog(void* parent, int f){
	return new QDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QDialog_Accept(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "accept");
}

void QDialog_ConnectAccepted(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_DisconnectAccepted(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_Done(void* ptr, int r){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "done", Q_ARG(int, r));
}

int QDialog_Exec(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "exec");
}

void QDialog_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_Open(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "open");
}

void QDialog_Reject(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "reject");
}

void QDialog_ConnectRejected(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

void QDialog_DisconnectRejected(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

int QDialog_Result(void* ptr){
	return static_cast<QDialog*>(ptr)->result();
}

void QDialog_SetVisible(void* ptr, int visible){
	static_cast<QDialog*>(ptr)->setVisible(visible != 0);
}

void QDialog_DestroyQDialog(void* ptr){
	static_cast<QDialog*>(ptr)->~QDialog();
}

#include "qtapandholdgesture.h"
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTapAndHoldGesture>
#include "_cgo_export.h"

class MyQTapAndHoldGesture: public QTapAndHoldGesture {
public:
};

void QTapAndHoldGesture_SetPosition(void* ptr, void* pos){
	static_cast<QTapAndHoldGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(int msecs){
	QTapAndHoldGesture::setTimeout(msecs);
}

int QTapAndHoldGesture_QTapAndHoldGesture_Timeout(){
	return QTapAndHoldGesture::timeout();
}

void QTapAndHoldGesture_DestroyQTapAndHoldGesture(void* ptr){
	static_cast<QTapAndHoldGesture*>(ptr)->~QTapAndHoldGesture();
}

#include "qstyleoptioncombobox.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionComboBox>
#include "_cgo_export.h"

class MyQStyleOptionComboBox: public QStyleOptionComboBox {
public:
};

void* QStyleOptionComboBox_NewQStyleOptionComboBox(){
	return new QStyleOptionComboBox();
}

void* QStyleOptionComboBox_NewQStyleOptionComboBox2(void* other){
	return new QStyleOptionComboBox(*static_cast<QStyleOptionComboBox*>(other));
}

#include "qsplitter.h"
#include <QByteArray>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSplitter>
#include "_cgo_export.h"

class MyQSplitter: public QSplitter {
public:
void Signal_SplitterMoved(int pos, int index){callbackQSplitterSplitterMoved(this->objectName().toUtf8().data(), pos, index);};
};

int QSplitter_ChildrenCollapsible(void* ptr){
	return static_cast<QSplitter*>(ptr)->childrenCollapsible();
}

int QSplitter_HandleWidth(void* ptr){
	return static_cast<QSplitter*>(ptr)->handleWidth();
}

int QSplitter_IndexOf(void* ptr, void* widget){
	return static_cast<QSplitter*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QSplitter_OpaqueResize(void* ptr){
	return static_cast<QSplitter*>(ptr)->opaqueResize();
}

int QSplitter_Orientation(void* ptr){
	return static_cast<QSplitter*>(ptr)->orientation();
}

void QSplitter_SetChildrenCollapsible(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setChildrenCollapsible(v != 0);
}

void QSplitter_SetHandleWidth(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setHandleWidth(v);
}

void QSplitter_SetOpaqueResize(void* ptr, int opaque){
	static_cast<QSplitter*>(ptr)->setOpaqueResize(opaque != 0);
}

void QSplitter_SetOrientation(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setOrientation(static_cast<Qt::Orientation>(v));
}

void* QSplitter_NewQSplitter(void* parent){
	return new QSplitter(static_cast<QWidget*>(parent));
}

void* QSplitter_NewQSplitter2(int orientation, void* parent){
	return new QSplitter(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QSplitter_AddWidget(void* ptr, void* widget){
	static_cast<QSplitter*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

int QSplitter_Count(void* ptr){
	return static_cast<QSplitter*>(ptr)->count();
}

void QSplitter_GetRange(void* ptr, int index, int min, int max){
	static_cast<QSplitter*>(ptr)->getRange(index, &min, &max);
}

void* QSplitter_Handle(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->handle(index);
}

void QSplitter_InsertWidget(void* ptr, int index, void* widget){
	static_cast<QSplitter*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

int QSplitter_IsCollapsible(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->isCollapsible(index);
}

void QSplitter_Refresh(void* ptr){
	static_cast<QSplitter*>(ptr)->refresh();
}

int QSplitter_RestoreState(void* ptr, void* state){
	return static_cast<QSplitter*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void* QSplitter_SaveState(void* ptr){
	return new QByteArray(static_cast<QSplitter*>(ptr)->saveState());
}

void QSplitter_SetCollapsible(void* ptr, int index, int collapse){
	static_cast<QSplitter*>(ptr)->setCollapsible(index, collapse != 0);
}

void QSplitter_SetStretchFactor(void* ptr, int index, int stretch){
	static_cast<QSplitter*>(ptr)->setStretchFactor(index, stretch);
}

void QSplitter_ConnectSplitterMoved(void* ptr){
	QObject::connect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

void QSplitter_DisconnectSplitterMoved(void* ptr){
	QObject::disconnect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

void* QSplitter_Widget(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->widget(index);
}

void QSplitter_DestroyQSplitter(void* ptr){
	static_cast<QSplitter*>(ptr)->~QSplitter();
}

#include "qgraphicsscenecontextmenuevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsScene>
#include <QGraphicsSceneContextMenuEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneContextMenuEvent: public QGraphicsSceneContextMenuEvent {
public:
};

int QGraphicsSceneContextMenuEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->modifiers();
}

int QGraphicsSceneContextMenuEvent_Reason(void* ptr){
	return static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->reason();
}

void QGraphicsSceneContextMenuEvent_DestroyQGraphicsSceneContextMenuEvent(void* ptr){
	static_cast<QGraphicsSceneContextMenuEvent*>(ptr)->~QGraphicsSceneContextMenuEvent();
}

#include "qitemeditorcreatorbase.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QWidget>
#include <QItemEditorCreator>
#include <QString>
#include <QVariant>
#include <QItemEditorCreatorBase>
#include "_cgo_export.h"

class MyQItemEditorCreatorBase: public QItemEditorCreatorBase {
public:
};

void QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(void* ptr){
	static_cast<QItemEditorCreatorBase*>(ptr)->~QItemEditorCreatorBase();
}

void* QItemEditorCreatorBase_CreateWidget(void* ptr, void* parent){
	return static_cast<QItemEditorCreatorBase*>(ptr)->createWidget(static_cast<QWidget*>(parent));
}

void* QItemEditorCreatorBase_ValuePropertyName(void* ptr){
	return new QByteArray(static_cast<QItemEditorCreatorBase*>(ptr)->valuePropertyName());
}

#include "qgesturerecognizer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGesture>
#include <QGestureRecognizer>
#include "_cgo_export.h"

#include "qmouseeventtransition.h"
#include <QVariant>
#include <QUrl>
#include <QState>
#include <QMouseEvent>
#include <QEvent>
#include <QPainter>
#include <QString>
#include <QPainterPath>
#include <QObject>
#include <QModelIndex>
#include <QMouseEventTransition>
#include "_cgo_export.h"

class MyQMouseEventTransition: public QMouseEventTransition {
public:
};

void* QMouseEventTransition_NewQMouseEventTransition2(void* object, int ty, int button, void* sourceState){
	return new QMouseEventTransition(static_cast<QObject*>(object), static_cast<QEvent::Type>(ty), static_cast<Qt::MouseButton>(button), static_cast<QState*>(sourceState));
}

void* QMouseEventTransition_NewQMouseEventTransition(void* sourceState){
	return new QMouseEventTransition(static_cast<QState*>(sourceState));
}

int QMouseEventTransition_Button(void* ptr){
	return static_cast<QMouseEventTransition*>(ptr)->button();
}

int QMouseEventTransition_ModifierMask(void* ptr){
	return static_cast<QMouseEventTransition*>(ptr)->modifierMask();
}

void QMouseEventTransition_SetButton(void* ptr, int button){
	static_cast<QMouseEventTransition*>(ptr)->setButton(static_cast<Qt::MouseButton>(button));
}

void QMouseEventTransition_SetHitTestPath(void* ptr, void* path){
	static_cast<QMouseEventTransition*>(ptr)->setHitTestPath(*static_cast<QPainterPath*>(path));
}

void QMouseEventTransition_SetModifierMask(void* ptr, int modifierMask){
	static_cast<QMouseEventTransition*>(ptr)->setModifierMask(static_cast<Qt::KeyboardModifier>(modifierMask));
}

void QMouseEventTransition_DestroyQMouseEventTransition(void* ptr){
	static_cast<QMouseEventTransition*>(ptr)->~QMouseEventTransition();
}

#include "qgraphicstextitem.h"
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QWidget>
#include <QVariant>
#include <QPainter>
#include <QTextCursor>
#include <QColor>
#include <QStyleOption>
#include <QString>
#include <QStyle>
#include <QGraphicsItem>
#include <QFont>
#include <QTextDocument>
#include <QObject>
#include <QStyleOptionGraphicsItem>
#include <QPointF>
#include <QGraphicsTextItem>
#include "_cgo_export.h"

class MyQGraphicsTextItem: public QGraphicsTextItem {
public:
void Signal_LinkActivated(const QString & link){callbackQGraphicsTextItemLinkActivated(this->objectName().toUtf8().data(), link.toUtf8().data());};
void Signal_LinkHovered(const QString & link){callbackQGraphicsTextItemLinkHovered(this->objectName().toUtf8().data(), link.toUtf8().data());};
};

int QGraphicsTextItem_OpenExternalLinks(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->openExternalLinks();
}

void QGraphicsTextItem_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QGraphicsTextItem*>(ptr)->setOpenExternalLinks(open != 0);
}

void QGraphicsTextItem_SetTextCursor(void* ptr, void* cursor){
	static_cast<QGraphicsTextItem*>(ptr)->setTextCursor(*static_cast<QTextCursor*>(cursor));
}

void QGraphicsTextItem_AdjustSize(void* ptr){
	static_cast<QGraphicsTextItem*>(ptr)->adjustSize();
}

int QGraphicsTextItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

void* QGraphicsTextItem_DefaultTextColor(void* ptr){
	return new QColor(static_cast<QGraphicsTextItem*>(ptr)->defaultTextColor());
}

void* QGraphicsTextItem_Document(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->document();
}

int QGraphicsTextItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsTextItem_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkActivated), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkActivated));;
}

void QGraphicsTextItem_ConnectLinkHovered(void* ptr){
	QObject::connect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_DisconnectLinkHovered(void* ptr){
	QObject::disconnect(static_cast<QGraphicsTextItem*>(ptr), static_cast<void (QGraphicsTextItem::*)(const QString &)>(&QGraphicsTextItem::linkHovered), static_cast<MyQGraphicsTextItem*>(ptr), static_cast<void (MyQGraphicsTextItem::*)(const QString &)>(&MyQGraphicsTextItem::Signal_LinkHovered));;
}

void QGraphicsTextItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsTextItem_SetDefaultTextColor(void* ptr, void* col){
	static_cast<QGraphicsTextItem*>(ptr)->setDefaultTextColor(*static_cast<QColor*>(col));
}

void QGraphicsTextItem_SetDocument(void* ptr, void* document){
	static_cast<QGraphicsTextItem*>(ptr)->setDocument(static_cast<QTextDocument*>(document));
}

void QGraphicsTextItem_SetFont(void* ptr, void* font){
	static_cast<QGraphicsTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsTextItem_SetHtml(void* ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setHtml(QString(text));
}

void QGraphicsTextItem_SetPlainText(void* ptr, char* text){
	static_cast<QGraphicsTextItem*>(ptr)->setPlainText(QString(text));
}

void QGraphicsTextItem_SetTabChangesFocus(void* ptr, int b){
	static_cast<QGraphicsTextItem*>(ptr)->setTabChangesFocus(b != 0);
}

void QGraphicsTextItem_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QGraphicsTextItem*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

void QGraphicsTextItem_SetTextWidth(void* ptr, double width){
	static_cast<QGraphicsTextItem*>(ptr)->setTextWidth(static_cast<qreal>(width));
}

int QGraphicsTextItem_TabChangesFocus(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->tabChangesFocus();
}

int QGraphicsTextItem_TextInteractionFlags(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->textInteractionFlags();
}

double QGraphicsTextItem_TextWidth(void* ptr){
	return static_cast<double>(static_cast<QGraphicsTextItem*>(ptr)->textWidth());
}

char* QGraphicsTextItem_ToHtml(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toHtml().toUtf8().data();
}

char* QGraphicsTextItem_ToPlainText(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->toPlainText().toUtf8().data();
}

int QGraphicsTextItem_Type(void* ptr){
	return static_cast<QGraphicsTextItem*>(ptr)->type();
}

void QGraphicsTextItem_DestroyQGraphicsTextItem(void* ptr){
	static_cast<QGraphicsTextItem*>(ptr)->~QGraphicsTextItem();
}

#include "qdoublespinbox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QDoubleSpinBox>
#include "_cgo_export.h"

class MyQDoubleSpinBox: public QDoubleSpinBox {
public:
};

char* QDoubleSpinBox_CleanText(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->cleanText().toUtf8().data();
}

int QDoubleSpinBox_Decimals(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->decimals();
}

char* QDoubleSpinBox_Prefix(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->prefix().toUtf8().data();
}

void QDoubleSpinBox_SetDecimals(void* ptr, int prec){
	static_cast<QDoubleSpinBox*>(ptr)->setDecimals(prec);
}

void QDoubleSpinBox_SetPrefix(void* ptr, char* prefix){
	static_cast<QDoubleSpinBox*>(ptr)->setPrefix(QString(prefix));
}

void QDoubleSpinBox_SetSuffix(void* ptr, char* suffix){
	static_cast<QDoubleSpinBox*>(ptr)->setSuffix(QString(suffix));
}

char* QDoubleSpinBox_Suffix(void* ptr){
	return static_cast<QDoubleSpinBox*>(ptr)->suffix().toUtf8().data();
}

void* QDoubleSpinBox_NewQDoubleSpinBox(void* parent){
	return new QDoubleSpinBox(static_cast<QWidget*>(parent));
}

void QDoubleSpinBox_DestroyQDoubleSpinBox(void* ptr){
	static_cast<QDoubleSpinBox*>(ptr)->~QDoubleSpinBox();
}

#include "qsystemtrayicon.h"
#include <QIcon>
#include <QMenu>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QSystemTrayIcon>
#include "_cgo_export.h"

class MyQSystemTrayIcon: public QSystemTrayIcon {
public:
void Signal_Activated(QSystemTrayIcon::ActivationReason reason){callbackQSystemTrayIconActivated(this->objectName().toUtf8().data(), reason);};
void Signal_MessageClicked(){callbackQSystemTrayIconMessageClicked(this->objectName().toUtf8().data());};
};

int QSystemTrayIcon_IsVisible(void* ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->isVisible();
}

void QSystemTrayIcon_SetIcon(void* ptr, void* icon){
	static_cast<QSystemTrayIcon*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QSystemTrayIcon_SetToolTip(void* ptr, char* tip){
	static_cast<QSystemTrayIcon*>(ptr)->setToolTip(QString(tip));
}

void QSystemTrayIcon_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QSystemTrayIcon_ShowMessage(void* ptr, char* title, char* message, int icon, int millisecondsTimeoutHint){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "showMessage", Q_ARG(QString, QString(title)), Q_ARG(QString, QString(message)), Q_ARG(QSystemTrayIcon::MessageIcon, static_cast<QSystemTrayIcon::MessageIcon>(icon)), Q_ARG(int, millisecondsTimeoutHint));
}

char* QSystemTrayIcon_ToolTip(void* ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->toolTip().toUtf8().data();
}

void* QSystemTrayIcon_NewQSystemTrayIcon(void* parent){
	return new QSystemTrayIcon(static_cast<QObject*>(parent));
}

void* QSystemTrayIcon_NewQSystemTrayIcon2(void* icon, void* parent){
	return new QSystemTrayIcon(*static_cast<QIcon*>(icon), static_cast<QObject*>(parent));
}

void QSystemTrayIcon_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&QSystemTrayIcon::activated), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&MyQSystemTrayIcon::Signal_Activated));;
}

void QSystemTrayIcon_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&QSystemTrayIcon::activated), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)(QSystemTrayIcon::ActivationReason)>(&MyQSystemTrayIcon::Signal_Activated));;
}

void* QSystemTrayIcon_ContextMenu(void* ptr){
	return static_cast<QSystemTrayIcon*>(ptr)->contextMenu();
}

void QSystemTrayIcon_Hide(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "hide");
}

int QSystemTrayIcon_QSystemTrayIcon_IsSystemTrayAvailable(){
	return QSystemTrayIcon::isSystemTrayAvailable();
}

void QSystemTrayIcon_ConnectMessageClicked(void* ptr){
	QObject::connect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)()>(&QSystemTrayIcon::messageClicked), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)()>(&MyQSystemTrayIcon::Signal_MessageClicked));;
}

void QSystemTrayIcon_DisconnectMessageClicked(void* ptr){
	QObject::disconnect(static_cast<QSystemTrayIcon*>(ptr), static_cast<void (QSystemTrayIcon::*)()>(&QSystemTrayIcon::messageClicked), static_cast<MyQSystemTrayIcon*>(ptr), static_cast<void (MyQSystemTrayIcon::*)()>(&MyQSystemTrayIcon::Signal_MessageClicked));;
}

void QSystemTrayIcon_SetContextMenu(void* ptr, void* menu){
	static_cast<QSystemTrayIcon*>(ptr)->setContextMenu(static_cast<QMenu*>(menu));
}

void QSystemTrayIcon_Show(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSystemTrayIcon*>(ptr), "show");
}

int QSystemTrayIcon_QSystemTrayIcon_SupportsMessages(){
	return QSystemTrayIcon::supportsMessages();
}

void QSystemTrayIcon_DestroyQSystemTrayIcon(void* ptr){
	static_cast<QSystemTrayIcon*>(ptr)->~QSystemTrayIcon();
}

#include "qstyleoptionspinbox.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionSpinBox>
#include "_cgo_export.h"

class MyQStyleOptionSpinBox: public QStyleOptionSpinBox {
public:
};

void* QStyleOptionSpinBox_NewQStyleOptionSpinBox(){
	return new QStyleOptionSpinBox();
}

void* QStyleOptionSpinBox_NewQStyleOptionSpinBox2(void* other){
	return new QStyleOptionSpinBox(*static_cast<QStyleOptionSpinBox*>(other));
}

#include "qcolormap.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QColormap>
#include "_cgo_export.h"

class MyQColormap: public QColormap {
public:
};

void* QColormap_NewQColormap(void* colormap){
	return new QColormap(*static_cast<QColormap*>(colormap));
}

int QColormap_Depth(void* ptr){
	return static_cast<QColormap*>(ptr)->depth();
}

int QColormap_Mode(void* ptr){
	return static_cast<QColormap*>(ptr)->mode();
}

int QColormap_Size(void* ptr){
	return static_cast<QColormap*>(ptr)->size();
}

void QColormap_DestroyQColormap(void* ptr){
	static_cast<QColormap*>(ptr)->~QColormap();
}

#include "qslider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QWidget>
#include <QSlider>
#include "_cgo_export.h"

class MyQSlider: public QSlider {
public:
};

void QSlider_SetTickInterval(void* ptr, int ti){
	static_cast<QSlider*>(ptr)->setTickInterval(ti);
}

void QSlider_SetTickPosition(void* ptr, int position){
	static_cast<QSlider*>(ptr)->setTickPosition(static_cast<QSlider::TickPosition>(position));
}

int QSlider_TickInterval(void* ptr){
	return static_cast<QSlider*>(ptr)->tickInterval();
}

int QSlider_TickPosition(void* ptr){
	return static_cast<QSlider*>(ptr)->tickPosition();
}

void* QSlider_NewQSlider(void* parent){
	return new QSlider(static_cast<QWidget*>(parent));
}

void* QSlider_NewQSlider2(int orientation, void* parent){
	return new QSlider(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QSlider_Event(void* ptr, void* event){
	return static_cast<QSlider*>(ptr)->event(static_cast<QEvent*>(event));
}

void QSlider_DestroyQSlider(void* ptr){
	static_cast<QSlider*>(ptr)->~QSlider();
}

#include "qfocusframe.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QFocusFrame>
#include "_cgo_export.h"

class MyQFocusFrame: public QFocusFrame {
public:
};

void* QFocusFrame_NewQFocusFrame(void* parent){
	return new QFocusFrame(static_cast<QWidget*>(parent));
}

void QFocusFrame_SetWidget(void* ptr, void* widget){
	static_cast<QFocusFrame*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

void* QFocusFrame_Widget(void* ptr){
	return static_cast<QFocusFrame*>(ptr)->widget();
}

void QFocusFrame_DestroyQFocusFrame(void* ptr){
	static_cast<QFocusFrame*>(ptr)->~QFocusFrame();
}

#include "qgraphicsellipseitem.h"
#include <QRect>
#include <QRectF>
#include <QPainter>
#include <QStyleOption>
#include <QGraphicsItem>
#include <QString>
#include <QModelIndex>
#include <QStyle>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QPointF>
#include <QGraphicsEllipseItem>
#include "_cgo_export.h"

class MyQGraphicsEllipseItem: public QGraphicsEllipseItem {
public:
};

int QGraphicsEllipseItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsEllipseItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsEllipseItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsEllipseItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsEllipseItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsEllipseItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsEllipseItem_SetRect(void* ptr, void* rect){
	static_cast<QGraphicsEllipseItem*>(ptr)->setRect(*static_cast<QRectF*>(rect));
}

void QGraphicsEllipseItem_SetRect2(void* ptr, double x, double y, double width, double height){
	static_cast<QGraphicsEllipseItem*>(ptr)->setRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QGraphicsEllipseItem_SetSpanAngle(void* ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setSpanAngle(angle);
}

void QGraphicsEllipseItem_SetStartAngle(void* ptr, int angle){
	static_cast<QGraphicsEllipseItem*>(ptr)->setStartAngle(angle);
}

int QGraphicsEllipseItem_SpanAngle(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->spanAngle();
}

int QGraphicsEllipseItem_StartAngle(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->startAngle();
}

int QGraphicsEllipseItem_Type(void* ptr){
	return static_cast<QGraphicsEllipseItem*>(ptr)->type();
}

void QGraphicsEllipseItem_DestroyQGraphicsEllipseItem(void* ptr){
	static_cast<QGraphicsEllipseItem*>(ptr)->~QGraphicsEllipseItem();
}

#include "qstackedwidget.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QStack>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QStackedWidget>
#include "_cgo_export.h"

class MyQStackedWidget: public QStackedWidget {
public:
void Signal_CurrentChanged(int index){callbackQStackedWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_WidgetRemoved(int index){callbackQStackedWidgetWidgetRemoved(this->objectName().toUtf8().data(), index);};
};

int QStackedWidget_Count(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->count();
}

int QStackedWidget_CurrentIndex(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->currentIndex();
}

void QStackedWidget_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QStackedWidget_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QStackedWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void* QStackedWidget_NewQStackedWidget(void* parent){
	return new QStackedWidget(static_cast<QWidget*>(parent));
}

int QStackedWidget_AddWidget(void* ptr, void* widget){
	return static_cast<QStackedWidget*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QStackedWidget_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

void QStackedWidget_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::currentChanged), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_CurrentChanged));;
}

void* QStackedWidget_CurrentWidget(void* ptr){
	return static_cast<QStackedWidget*>(ptr)->currentWidget();
}

int QStackedWidget_IndexOf(void* ptr, void* widget){
	return static_cast<QStackedWidget*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QStackedWidget_InsertWidget(void* ptr, int index, void* widget){
	return static_cast<QStackedWidget*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

void QStackedWidget_RemoveWidget(void* ptr, void* widget){
	static_cast<QStackedWidget*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void* QStackedWidget_Widget(void* ptr, int index){
	return static_cast<QStackedWidget*>(ptr)->widget(index);
}

void QStackedWidget_ConnectWidgetRemoved(void* ptr){
	QObject::connect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DisconnectWidgetRemoved(void* ptr){
	QObject::disconnect(static_cast<QStackedWidget*>(ptr), static_cast<void (QStackedWidget::*)(int)>(&QStackedWidget::widgetRemoved), static_cast<MyQStackedWidget*>(ptr), static_cast<void (MyQStackedWidget::*)(int)>(&MyQStackedWidget::Signal_WidgetRemoved));;
}

void QStackedWidget_DestroyQStackedWidget(void* ptr){
	static_cast<QStackedWidget*>(ptr)->~QStackedWidget();
}

#include "qspaceritem.h"
#include <QRect>
#include <QSize>
#include <QSizePolicy>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSpacerItem>
#include "_cgo_export.h"

class MyQSpacerItem: public QSpacerItem {
public:
};

void* QSpacerItem_NewQSpacerItem(int w, int h, int hPolicy, int vPolicy){
	return new QSpacerItem(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

void QSpacerItem_ChangeSize(void* ptr, int w, int h, int hPolicy, int vPolicy){
	static_cast<QSpacerItem*>(ptr)->changeSize(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

int QSpacerItem_ExpandingDirections(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->expandingDirections();
}

int QSpacerItem_IsEmpty(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->isEmpty();
}

void QSpacerItem_SetGeometry(void* ptr, void* r){
	static_cast<QSpacerItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void* QSpacerItem_SpacerItem(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->spacerItem();
}

void QSpacerItem_DestroyQSpacerItem(void* ptr){
	static_cast<QSpacerItem*>(ptr)->~QSpacerItem();
}

#include "qgraphicspathitem.h"
#include <QVariant>
#include <QModelIndex>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QGraphicsItem>
#include <QString>
#include <QUrl>
#include <QPainter>
#include <QStyle>
#include <QPoint>
#include <QPainterPath>
#include <QPointF>
#include <QGraphicsPathItem>
#include "_cgo_export.h"

class MyQGraphicsPathItem: public QGraphicsPathItem {
public:
};

int QGraphicsPathItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPathItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPathItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPathItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPathItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPathItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPathItem_SetPath(void* ptr, void* path){
	static_cast<QGraphicsPathItem*>(ptr)->setPath(*static_cast<QPainterPath*>(path));
}

int QGraphicsPathItem_Type(void* ptr){
	return static_cast<QGraphicsPathItem*>(ptr)->type();
}

void QGraphicsPathItem_DestroyQGraphicsPathItem(void* ptr){
	static_cast<QGraphicsPathItem*>(ptr)->~QGraphicsPathItem();
}

#include "qgridlayout.h"
#include <QLayoutItem>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLayout>
#include <QWidget>
#include <QGridLayout>
#include "_cgo_export.h"

class MyQGridLayout: public QGridLayout {
public:
};

int QGridLayout_HorizontalSpacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->horizontalSpacing();
}

void QGridLayout_SetHorizontalSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setHorizontalSpacing(spacing);
}

void QGridLayout_SetVerticalSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setVerticalSpacing(spacing);
}

int QGridLayout_VerticalSpacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->verticalSpacing();
}

void* QGridLayout_NewQGridLayout2(){
	return new QGridLayout();
}

void* QGridLayout_NewQGridLayout(void* parent){
	return new QGridLayout(static_cast<QWidget*>(parent));
}

void QGridLayout_AddItem(void* ptr, void* item, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout(void* ptr, void* layout, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddLayout2(void* ptr, void* layout, int row, int column, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), row, column, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget2(void* ptr, void* widget, int fromRow, int fromColumn, int rowSpan, int columnSpan, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), fromRow, fromColumn, rowSpan, columnSpan, static_cast<Qt::AlignmentFlag>(alignment));
}

void QGridLayout_AddWidget(void* ptr, void* widget, int row, int column, int alignment){
	static_cast<QGridLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), row, column, static_cast<Qt::AlignmentFlag>(alignment));
}

int QGridLayout_ColumnCount(void* ptr){
	return static_cast<QGridLayout*>(ptr)->columnCount();
}

int QGridLayout_ColumnMinimumWidth(void* ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnMinimumWidth(column);
}

int QGridLayout_ColumnStretch(void* ptr, int column){
	return static_cast<QGridLayout*>(ptr)->columnStretch(column);
}

int QGridLayout_Count(void* ptr){
	return static_cast<QGridLayout*>(ptr)->count();
}

int QGridLayout_ExpandingDirections(void* ptr){
	return static_cast<QGridLayout*>(ptr)->expandingDirections();
}

void QGridLayout_GetItemPosition(void* ptr, int index, int row, int column, int rowSpan, int columnSpan){
	static_cast<QGridLayout*>(ptr)->getItemPosition(index, &row, &column, &rowSpan, &columnSpan);
}

int QGridLayout_HasHeightForWidth(void* ptr){
	return static_cast<QGridLayout*>(ptr)->hasHeightForWidth();
}

int QGridLayout_HeightForWidth(void* ptr, int w){
	return static_cast<QGridLayout*>(ptr)->heightForWidth(w);
}

void QGridLayout_Invalidate(void* ptr){
	static_cast<QGridLayout*>(ptr)->invalidate();
}

void* QGridLayout_ItemAt(void* ptr, int index){
	return static_cast<QGridLayout*>(ptr)->itemAt(index);
}

void* QGridLayout_ItemAtPosition(void* ptr, int row, int column){
	return static_cast<QGridLayout*>(ptr)->itemAtPosition(row, column);
}

int QGridLayout_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QGridLayout*>(ptr)->minimumHeightForWidth(w);
}

int QGridLayout_OriginCorner(void* ptr){
	return static_cast<QGridLayout*>(ptr)->originCorner();
}

int QGridLayout_RowCount(void* ptr){
	return static_cast<QGridLayout*>(ptr)->rowCount();
}

int QGridLayout_RowMinimumHeight(void* ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowMinimumHeight(row);
}

int QGridLayout_RowStretch(void* ptr, int row){
	return static_cast<QGridLayout*>(ptr)->rowStretch(row);
}

void QGridLayout_SetColumnMinimumWidth(void* ptr, int column, int minSize){
	static_cast<QGridLayout*>(ptr)->setColumnMinimumWidth(column, minSize);
}

void QGridLayout_SetColumnStretch(void* ptr, int column, int stretch){
	static_cast<QGridLayout*>(ptr)->setColumnStretch(column, stretch);
}

void QGridLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGridLayout*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QGridLayout_SetOriginCorner(void* ptr, int corner){
	static_cast<QGridLayout*>(ptr)->setOriginCorner(static_cast<Qt::Corner>(corner));
}

void QGridLayout_SetRowMinimumHeight(void* ptr, int row, int minSize){
	static_cast<QGridLayout*>(ptr)->setRowMinimumHeight(row, minSize);
}

void QGridLayout_SetRowStretch(void* ptr, int row, int stretch){
	static_cast<QGridLayout*>(ptr)->setRowStretch(row, stretch);
}

void QGridLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QGridLayout*>(ptr)->setSpacing(spacing);
}

int QGridLayout_Spacing(void* ptr){
	return static_cast<QGridLayout*>(ptr)->spacing();
}

void* QGridLayout_TakeAt(void* ptr, int index){
	return static_cast<QGridLayout*>(ptr)->takeAt(index);
}

void QGridLayout_DestroyQGridLayout(void* ptr){
	static_cast<QGridLayout*>(ptr)->~QGridLayout();
}

#include "qgroupbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QGroupBox>
#include "_cgo_export.h"

class MyQGroupBox: public QGroupBox {
public:
void Signal_Clicked(bool checked){callbackQGroupBoxClicked(this->objectName().toUtf8().data(), checked);};
void Signal_Toggled(bool on){callbackQGroupBoxToggled(this->objectName().toUtf8().data(), on);};
};

int QGroupBox_Alignment(void* ptr){
	return static_cast<QGroupBox*>(ptr)->alignment();
}

int QGroupBox_IsCheckable(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isCheckable();
}

int QGroupBox_IsChecked(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isChecked();
}

int QGroupBox_IsFlat(void* ptr){
	return static_cast<QGroupBox*>(ptr)->isFlat();
}

void QGroupBox_SetAlignment(void* ptr, int alignment){
	static_cast<QGroupBox*>(ptr)->setAlignment(alignment);
}

void QGroupBox_SetCheckable(void* ptr, int checkable){
	static_cast<QGroupBox*>(ptr)->setCheckable(checkable != 0);
}

void QGroupBox_SetChecked(void* ptr, int checked){
	QMetaObject::invokeMethod(static_cast<QGroupBox*>(ptr), "setChecked", Q_ARG(bool, checked != 0));
}

void QGroupBox_SetFlat(void* ptr, int flat){
	static_cast<QGroupBox*>(ptr)->setFlat(flat != 0);
}

void QGroupBox_SetTitle(void* ptr, char* title){
	static_cast<QGroupBox*>(ptr)->setTitle(QString(title));
}

char* QGroupBox_Title(void* ptr){
	return static_cast<QGroupBox*>(ptr)->title().toUtf8().data();
}

void* QGroupBox_NewQGroupBox(void* parent){
	return new QGroupBox(static_cast<QWidget*>(parent));
}

void* QGroupBox_NewQGroupBox2(char* title, void* parent){
	return new QGroupBox(QString(title), static_cast<QWidget*>(parent));
}

void QGroupBox_ConnectClicked(void* ptr){
	QObject::connect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked));;
}

void QGroupBox_DisconnectClicked(void* ptr){
	QObject::disconnect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::clicked), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Clicked));;
}

void QGroupBox_ConnectToggled(void* ptr){
	QObject::connect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::toggled), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Toggled));;
}

void QGroupBox_DisconnectToggled(void* ptr){
	QObject::disconnect(static_cast<QGroupBox*>(ptr), static_cast<void (QGroupBox::*)(bool)>(&QGroupBox::toggled), static_cast<MyQGroupBox*>(ptr), static_cast<void (MyQGroupBox::*)(bool)>(&MyQGroupBox::Signal_Toggled));;
}

void QGroupBox_DestroyQGroupBox(void* ptr){
	static_cast<QGroupBox*>(ptr)->~QGroupBox();
}

#include "qtooltip.h"
#include <QRect>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QWidget>
#include <QPalette>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QToolTip>
#include "_cgo_export.h"

class MyQToolTip: public QToolTip {
public:
};

void QToolTip_QToolTip_HideText(){
	QToolTip::hideText();
}

int QToolTip_QToolTip_IsVisible(){
	return QToolTip::isVisible();
}

void QToolTip_QToolTip_SetFont(void* font){
	QToolTip::setFont(*static_cast<QFont*>(font));
}

void QToolTip_QToolTip_SetPalette(void* palette){
	QToolTip::setPalette(*static_cast<QPalette*>(palette));
}

void QToolTip_QToolTip_ShowText3(void* pos, char* text, void* w){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w));
}

void QToolTip_QToolTip_ShowText(void* pos, char* text, void* w, void* rect){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect));
}

void QToolTip_QToolTip_ShowText2(void* pos, char* text, void* w, void* rect, int msecDisplayTime){
	QToolTip::showText(*static_cast<QPoint*>(pos), QString(text), static_cast<QWidget*>(w), *static_cast<QRect*>(rect), msecDisplayTime);
}

char* QToolTip_QToolTip_Text(){
	return QToolTip::text().toUtf8().data();
}

#include "qsplitterhandle.h"
#include <QModelIndex>
#include <QSplitter>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSplitterHandle>
#include "_cgo_export.h"

class MyQSplitterHandle: public QSplitterHandle {
public:
};

void* QSplitterHandle_NewQSplitterHandle(int orientation, void* parent){
	return new QSplitterHandle(static_cast<Qt::Orientation>(orientation), static_cast<QSplitter*>(parent));
}

int QSplitterHandle_OpaqueResize(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->opaqueResize();
}

int QSplitterHandle_Orientation(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->orientation();
}

void QSplitterHandle_SetOrientation(void* ptr, int orientation){
	static_cast<QSplitterHandle*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void* QSplitterHandle_Splitter(void* ptr){
	return static_cast<QSplitterHandle*>(ptr)->splitter();
}

void QSplitterHandle_DestroyQSplitterHandle(void* ptr){
	static_cast<QSplitterHandle*>(ptr)->~QSplitterHandle();
}

#include "qtabbar.h"
#include <QObject>
#include <QColor>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QIcon>
#include <QMetaObject>
#include <QPoint>
#include <QSize>
#include <QString>
#include <QModelIndex>
#include <QTabBar>
#include "_cgo_export.h"

class MyQTabBar: public QTabBar {
public:
void Signal_CurrentChanged(int index){callbackQTabBarCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_TabBarClicked(int index){callbackQTabBarTabBarClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabBarDoubleClicked(int index){callbackQTabBarTabBarDoubleClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabCloseRequested(int index){callbackQTabBarTabCloseRequested(this->objectName().toUtf8().data(), index);};
void Signal_TabMoved(int from, int to){callbackQTabBarTabMoved(this->objectName().toUtf8().data(), from, to);};
};

int QTabBar_AutoHide(void* ptr){
	return static_cast<QTabBar*>(ptr)->autoHide();
}

int QTabBar_ChangeCurrentOnDrag(void* ptr){
	return static_cast<QTabBar*>(ptr)->changeCurrentOnDrag();
}

int QTabBar_Count(void* ptr){
	return static_cast<QTabBar*>(ptr)->count();
}

int QTabBar_CurrentIndex(void* ptr){
	return static_cast<QTabBar*>(ptr)->currentIndex();
}

int QTabBar_DocumentMode(void* ptr){
	return static_cast<QTabBar*>(ptr)->documentMode();
}

int QTabBar_DrawBase(void* ptr){
	return static_cast<QTabBar*>(ptr)->drawBase();
}

int QTabBar_ElideMode(void* ptr){
	return static_cast<QTabBar*>(ptr)->elideMode();
}

int QTabBar_Expanding(void* ptr){
	return static_cast<QTabBar*>(ptr)->expanding();
}

int QTabBar_IsMovable(void* ptr){
	return static_cast<QTabBar*>(ptr)->isMovable();
}

int QTabBar_SelectionBehaviorOnRemove(void* ptr){
	return static_cast<QTabBar*>(ptr)->selectionBehaviorOnRemove();
}

void QTabBar_SetAutoHide(void* ptr, int hide){
	static_cast<QTabBar*>(ptr)->setAutoHide(hide != 0);
}

void QTabBar_SetChangeCurrentOnDrag(void* ptr, int change){
	static_cast<QTabBar*>(ptr)->setChangeCurrentOnDrag(change != 0);
}

void QTabBar_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabBar*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabBar_SetDocumentMode(void* ptr, int set){
	static_cast<QTabBar*>(ptr)->setDocumentMode(set != 0);
}

void QTabBar_SetDrawBase(void* ptr, int drawTheBase){
	static_cast<QTabBar*>(ptr)->setDrawBase(drawTheBase != 0);
}

void QTabBar_SetElideMode(void* ptr, int v){
	static_cast<QTabBar*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabBar_SetExpanding(void* ptr, int enabled){
	static_cast<QTabBar*>(ptr)->setExpanding(enabled != 0);
}

void QTabBar_SetIconSize(void* ptr, void* size){
	static_cast<QTabBar*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabBar_SetMovable(void* ptr, int movable){
	static_cast<QTabBar*>(ptr)->setMovable(movable != 0);
}

void QTabBar_SetSelectionBehaviorOnRemove(void* ptr, int behavior){
	static_cast<QTabBar*>(ptr)->setSelectionBehaviorOnRemove(static_cast<QTabBar::SelectionBehavior>(behavior));
}

void QTabBar_SetShape(void* ptr, int shape){
	static_cast<QTabBar*>(ptr)->setShape(static_cast<QTabBar::Shape>(shape));
}

void QTabBar_SetTabsClosable(void* ptr, int closable){
	static_cast<QTabBar*>(ptr)->setTabsClosable(closable != 0);
}

void QTabBar_SetUsesScrollButtons(void* ptr, int useButtons){
	static_cast<QTabBar*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabBar_Shape(void* ptr){
	return static_cast<QTabBar*>(ptr)->shape();
}

int QTabBar_TabsClosable(void* ptr){
	return static_cast<QTabBar*>(ptr)->tabsClosable();
}

int QTabBar_UsesScrollButtons(void* ptr){
	return static_cast<QTabBar*>(ptr)->usesScrollButtons();
}

void* QTabBar_NewQTabBar(void* parent){
	return new QTabBar(static_cast<QWidget*>(parent));
}

int QTabBar_AddTab2(void* ptr, void* icon, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(*static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_AddTab(void* ptr, char* text){
	return static_cast<QTabBar*>(ptr)->addTab(QString(text));
}

void QTabBar_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

void QTabBar_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::currentChanged), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_CurrentChanged));;
}

int QTabBar_InsertTab2(void* ptr, int index, void* icon, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, *static_cast<QIcon*>(icon), QString(text));
}

int QTabBar_InsertTab(void* ptr, int index, char* text){
	return static_cast<QTabBar*>(ptr)->insertTab(index, QString(text));
}

int QTabBar_IsTabEnabled(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->isTabEnabled(index);
}

void QTabBar_MoveTab(void* ptr, int from, int to){
	static_cast<QTabBar*>(ptr)->moveTab(from, to);
}

void QTabBar_RemoveTab(void* ptr, int index){
	static_cast<QTabBar*>(ptr)->removeTab(index);
}

void QTabBar_SetTabButton(void* ptr, int index, int position, void* widget){
	static_cast<QTabBar*>(ptr)->setTabButton(index, static_cast<QTabBar::ButtonPosition>(position), static_cast<QWidget*>(widget));
}

void QTabBar_SetTabData(void* ptr, int index, void* data){
	static_cast<QTabBar*>(ptr)->setTabData(index, *static_cast<QVariant*>(data));
}

void QTabBar_SetTabEnabled(void* ptr, int index, int enabled){
	static_cast<QTabBar*>(ptr)->setTabEnabled(index, enabled != 0);
}

void QTabBar_SetTabIcon(void* ptr, int index, void* icon){
	static_cast<QTabBar*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabBar_SetTabText(void* ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabText(index, QString(text));
}

void QTabBar_SetTabTextColor(void* ptr, int index, void* color){
	static_cast<QTabBar*>(ptr)->setTabTextColor(index, *static_cast<QColor*>(color));
}

void QTabBar_SetTabToolTip(void* ptr, int index, char* tip){
	static_cast<QTabBar*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabBar_SetTabWhatsThis(void* ptr, int index, char* text){
	static_cast<QTabBar*>(ptr)->setTabWhatsThis(index, QString(text));
}

int QTabBar_TabAt(void* ptr, void* position){
	return static_cast<QTabBar*>(ptr)->tabAt(*static_cast<QPoint*>(position));
}

void QTabBar_ConnectTabBarClicked(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_DisconnectTabBarClicked(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarClicked));;
}

void QTabBar_ConnectTabBarDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

void QTabBar_DisconnectTabBarDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabBarDoubleClicked), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabBarDoubleClicked));;
}

void* QTabBar_TabButton(void* ptr, int index, int position){
	return static_cast<QTabBar*>(ptr)->tabButton(index, static_cast<QTabBar::ButtonPosition>(position));
}

void QTabBar_ConnectTabCloseRequested(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

void QTabBar_DisconnectTabCloseRequested(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int)>(&QTabBar::tabCloseRequested), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int)>(&MyQTabBar::Signal_TabCloseRequested));;
}

void* QTabBar_TabData(void* ptr, int index){
	return new QVariant(static_cast<QTabBar*>(ptr)->tabData(index));
}

void QTabBar_ConnectTabMoved(void* ptr){
	QObject::connect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

void QTabBar_DisconnectTabMoved(void* ptr){
	QObject::disconnect(static_cast<QTabBar*>(ptr), static_cast<void (QTabBar::*)(int, int)>(&QTabBar::tabMoved), static_cast<MyQTabBar*>(ptr), static_cast<void (MyQTabBar::*)(int, int)>(&MyQTabBar::Signal_TabMoved));;
}

char* QTabBar_TabText(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabText(index).toUtf8().data();
}

void* QTabBar_TabTextColor(void* ptr, int index){
	return new QColor(static_cast<QTabBar*>(ptr)->tabTextColor(index));
}

char* QTabBar_TabToolTip(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabBar_TabWhatsThis(void* ptr, int index){
	return static_cast<QTabBar*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

void QTabBar_DestroyQTabBar(void* ptr){
	static_cast<QTabBar*>(ptr)->~QTabBar();
}

#include "qgraphicspixmapitem.h"
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QStyleOption>
#include <QPixmap>
#include <QPainter>
#include <QStyle>
#include <QString>
#include <QModelIndex>
#include <QPointF>
#include <QGraphicsItem>
#include <QPoint>
#include <QVariant>
#include <QUrl>
#include <QGraphicsPixmapItem>
#include "_cgo_export.h"

class MyQGraphicsPixmapItem: public QGraphicsPixmapItem {
public:
};

void* QGraphicsPixmapItem_NewQGraphicsPixmapItem(void* parent){
	return new QGraphicsPixmapItem(static_cast<QGraphicsItem*>(parent));
}

void* QGraphicsPixmapItem_NewQGraphicsPixmapItem2(void* pixmap, void* parent){
	return new QGraphicsPixmapItem(*static_cast<QPixmap*>(pixmap), static_cast<QGraphicsItem*>(parent));
}

int QGraphicsPixmapItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPixmapItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPixmapItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPixmapItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPixmapItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPixmapItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPixmapItem_SetOffset(void* ptr, void* offset){
	static_cast<QGraphicsPixmapItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsPixmapItem_SetOffset2(void* ptr, double x, double y){
	static_cast<QGraphicsPixmapItem*>(ptr)->setOffset(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QGraphicsPixmapItem_SetPixmap(void* ptr, void* pixmap){
	static_cast<QGraphicsPixmapItem*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QGraphicsPixmapItem_SetShapeMode(void* ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setShapeMode(static_cast<QGraphicsPixmapItem::ShapeMode>(mode));
}

void QGraphicsPixmapItem_SetTransformationMode(void* ptr, int mode){
	static_cast<QGraphicsPixmapItem*>(ptr)->setTransformationMode(static_cast<Qt::TransformationMode>(mode));
}

int QGraphicsPixmapItem_ShapeMode(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->shapeMode();
}

int QGraphicsPixmapItem_TransformationMode(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->transformationMode();
}

int QGraphicsPixmapItem_Type(void* ptr){
	return static_cast<QGraphicsPixmapItem*>(ptr)->type();
}

void QGraphicsPixmapItem_DestroyQGraphicsPixmapItem(void* ptr){
	static_cast<QGraphicsPixmapItem*>(ptr)->~QGraphicsPixmapItem();
}

#include "qkeysequenceedit.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QWidget>
#include <QKeySequence>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QKeySequenceEdit>
#include "_cgo_export.h"

class MyQKeySequenceEdit: public QKeySequenceEdit {
public:
void Signal_EditingFinished(){callbackQKeySequenceEditEditingFinished(this->objectName().toUtf8().data());};
};

void QKeySequenceEdit_SetKeySequence(void* ptr, void* keySequence){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "setKeySequence", Q_ARG(QKeySequence, *static_cast<QKeySequence*>(keySequence)));
}

void* QKeySequenceEdit_NewQKeySequenceEdit(void* parent){
	return new QKeySequenceEdit(static_cast<QWidget*>(parent));
}

void* QKeySequenceEdit_NewQKeySequenceEdit2(void* keySequence, void* parent){
	return new QKeySequenceEdit(*static_cast<QKeySequence*>(keySequence), static_cast<QWidget*>(parent));
}

void QKeySequenceEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "clear");
}

void QKeySequenceEdit_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DestroyQKeySequenceEdit(void* ptr){
	static_cast<QKeySequenceEdit*>(ptr)->~QKeySequenceEdit();
}

#include "qmaccocoaviewcontainer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMacCocoaViewContainer>
#include "_cgo_export.h"

class MyQMacCocoaViewContainer: public QMacCocoaViewContainer {
public:
};

void QMacCocoaViewContainer_DestroyQMacCocoaViewContainer(void* ptr){
	static_cast<QMacCocoaViewContainer*>(ptr)->~QMacCocoaViewContainer();
}

#include "qgraphicssimpletextitem.h"
#include <QModelIndex>
#include <QPoint>
#include <QStyleOptionGraphicsItem>
#include <QFont>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWidget>
#include <QPointF>
#include <QStyleOption>
#include <QGraphicsItem>
#include <QStyle>
#include <QGraphicsSimpleTextItem>
#include "_cgo_export.h"

class MyQGraphicsSimpleTextItem: public QGraphicsSimpleTextItem {
public:
};

int QGraphicsSimpleTextItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsSimpleTextItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsSimpleTextItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsSimpleTextItem_SetFont(void* ptr, void* font){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QGraphicsSimpleTextItem_SetText(void* ptr, char* text){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->setText(QString(text));
}

char* QGraphicsSimpleTextItem_Text(void* ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->text().toUtf8().data();
}

int QGraphicsSimpleTextItem_Type(void* ptr){
	return static_cast<QGraphicsSimpleTextItem*>(ptr)->type();
}

void QGraphicsSimpleTextItem_DestroyQGraphicsSimpleTextItem(void* ptr){
	static_cast<QGraphicsSimpleTextItem*>(ptr)->~QGraphicsSimpleTextItem();
}

#include "qstyleplugin.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QString>
#include <QStylePlugin>
#include "_cgo_export.h"

class MyQStylePlugin: public QStylePlugin {
public:
};

void* QStylePlugin_Create(void* ptr, char* key){
	return static_cast<QStylePlugin*>(ptr)->create(QString(key));
}

void QStylePlugin_DestroyQStylePlugin(void* ptr){
	static_cast<QStylePlugin*>(ptr)->~QStylePlugin();
}

#include "qprogressbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
#include <QProgressBar>
#include "_cgo_export.h"

class MyQProgressBar: public QProgressBar {
public:
void Signal_ValueChanged(int value){callbackQProgressBarValueChanged(this->objectName().toUtf8().data(), value);};
};

int QProgressBar_Alignment(void* ptr){
	return static_cast<QProgressBar*>(ptr)->alignment();
}

char* QProgressBar_Format(void* ptr){
	return static_cast<QProgressBar*>(ptr)->format().toUtf8().data();
}

int QProgressBar_InvertedAppearance(void* ptr){
	return static_cast<QProgressBar*>(ptr)->invertedAppearance();
}

int QProgressBar_IsTextVisible(void* ptr){
	return static_cast<QProgressBar*>(ptr)->isTextVisible();
}

int QProgressBar_Maximum(void* ptr){
	return static_cast<QProgressBar*>(ptr)->maximum();
}

int QProgressBar_Minimum(void* ptr){
	return static_cast<QProgressBar*>(ptr)->minimum();
}

int QProgressBar_Orientation(void* ptr){
	return static_cast<QProgressBar*>(ptr)->orientation();
}

void QProgressBar_ResetFormat(void* ptr){
	static_cast<QProgressBar*>(ptr)->resetFormat();
}

void QProgressBar_SetAlignment(void* ptr, int alignment){
	static_cast<QProgressBar*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QProgressBar_SetFormat(void* ptr, char* format){
	static_cast<QProgressBar*>(ptr)->setFormat(QString(format));
}

void QProgressBar_SetInvertedAppearance(void* ptr, int invert){
	static_cast<QProgressBar*>(ptr)->setInvertedAppearance(invert != 0);
}

void QProgressBar_SetMaximum(void* ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressBar_SetMinimum(void* ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressBar_SetOrientation(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(v)));
}

void QProgressBar_SetTextDirection(void* ptr, int textDirection){
	static_cast<QProgressBar*>(ptr)->setTextDirection(static_cast<QProgressBar::Direction>(textDirection));
}

void QProgressBar_SetTextVisible(void* ptr, int visible){
	static_cast<QProgressBar*>(ptr)->setTextVisible(visible != 0);
}

void QProgressBar_SetValue(void* ptr, int value){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setValue", Q_ARG(int, value));
}

char* QProgressBar_Text(void* ptr){
	return static_cast<QProgressBar*>(ptr)->text().toUtf8().data();
}

int QProgressBar_TextDirection(void* ptr){
	return static_cast<QProgressBar*>(ptr)->textDirection();
}

int QProgressBar_Value(void* ptr){
	return static_cast<QProgressBar*>(ptr)->value();
}

void* QProgressBar_NewQProgressBar(void* parent){
	return new QProgressBar(static_cast<QWidget*>(parent));
}

void QProgressBar_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "reset");
}

void QProgressBar_SetRange(void* ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressBar_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DestroyQProgressBar(void* ptr){
	static_cast<QProgressBar*>(ptr)->~QProgressBar();
}

#include "qtabwidget.h"
#include <QSize>
#include <QObject>
#include <QIcon>
#include <QMetaObject>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QString>
#include <QTabWidget>
#include "_cgo_export.h"

class MyQTabWidget: public QTabWidget {
public:
void Signal_CurrentChanged(int index){callbackQTabWidgetCurrentChanged(this->objectName().toUtf8().data(), index);};
void Signal_TabBarClicked(int index){callbackQTabWidgetTabBarClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabBarDoubleClicked(int index){callbackQTabWidgetTabBarDoubleClicked(this->objectName().toUtf8().data(), index);};
void Signal_TabCloseRequested(int index){callbackQTabWidgetTabCloseRequested(this->objectName().toUtf8().data(), index);};
};

int QTabWidget_AddTab2(void* ptr, void* page, void* icon, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_AddTab(void* ptr, void* page, char* label){
	return static_cast<QTabWidget*>(ptr)->addTab(static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_Count(void* ptr){
	return static_cast<QTabWidget*>(ptr)->count();
}

int QTabWidget_CurrentIndex(void* ptr){
	return static_cast<QTabWidget*>(ptr)->currentIndex();
}

int QTabWidget_DocumentMode(void* ptr){
	return static_cast<QTabWidget*>(ptr)->documentMode();
}

int QTabWidget_ElideMode(void* ptr){
	return static_cast<QTabWidget*>(ptr)->elideMode();
}

int QTabWidget_InsertTab2(void* ptr, int index, void* page, void* icon, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), *static_cast<QIcon*>(icon), QString(label));
}

int QTabWidget_InsertTab(void* ptr, int index, void* page, char* label){
	return static_cast<QTabWidget*>(ptr)->insertTab(index, static_cast<QWidget*>(page), QString(label));
}

int QTabWidget_IsMovable(void* ptr){
	return static_cast<QTabWidget*>(ptr)->isMovable();
}

void QTabWidget_SetCornerWidget(void* ptr, void* widget, int corner){
	static_cast<QTabWidget*>(ptr)->setCornerWidget(static_cast<QWidget*>(widget), static_cast<Qt::Corner>(corner));
}

void QTabWidget_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QTabWidget_SetDocumentMode(void* ptr, int set){
	static_cast<QTabWidget*>(ptr)->setDocumentMode(set != 0);
}

void QTabWidget_SetElideMode(void* ptr, int v){
	static_cast<QTabWidget*>(ptr)->setElideMode(static_cast<Qt::TextElideMode>(v));
}

void QTabWidget_SetIconSize(void* ptr, void* size){
	static_cast<QTabWidget*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QTabWidget_SetMovable(void* ptr, int movable){
	static_cast<QTabWidget*>(ptr)->setMovable(movable != 0);
}

void QTabWidget_SetTabBarAutoHide(void* ptr, int enabled){
	static_cast<QTabWidget*>(ptr)->setTabBarAutoHide(enabled != 0);
}

void QTabWidget_SetTabPosition(void* ptr, int v){
	static_cast<QTabWidget*>(ptr)->setTabPosition(static_cast<QTabWidget::TabPosition>(v));
}

void QTabWidget_SetTabShape(void* ptr, int s){
	static_cast<QTabWidget*>(ptr)->setTabShape(static_cast<QTabWidget::TabShape>(s));
}

void QTabWidget_SetTabsClosable(void* ptr, int closeable){
	static_cast<QTabWidget*>(ptr)->setTabsClosable(closeable != 0);
}

void QTabWidget_SetUsesScrollButtons(void* ptr, int useButtons){
	static_cast<QTabWidget*>(ptr)->setUsesScrollButtons(useButtons != 0);
}

int QTabWidget_TabBarAutoHide(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabBarAutoHide();
}

int QTabWidget_TabPosition(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabPosition();
}

int QTabWidget_TabShape(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabShape();
}

int QTabWidget_TabsClosable(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabsClosable();
}

int QTabWidget_UsesScrollButtons(void* ptr){
	return static_cast<QTabWidget*>(ptr)->usesScrollButtons();
}

void* QTabWidget_NewQTabWidget(void* parent){
	return new QTabWidget(static_cast<QWidget*>(parent));
}

void QTabWidget_Clear(void* ptr){
	static_cast<QTabWidget*>(ptr)->clear();
}

void* QTabWidget_CornerWidget(void* ptr, int corner){
	return static_cast<QTabWidget*>(ptr)->cornerWidget(static_cast<Qt::Corner>(corner));
}

void QTabWidget_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

void QTabWidget_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::currentChanged), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_CurrentChanged));;
}

void* QTabWidget_CurrentWidget(void* ptr){
	return static_cast<QTabWidget*>(ptr)->currentWidget();
}

int QTabWidget_HasHeightForWidth(void* ptr){
	return static_cast<QTabWidget*>(ptr)->hasHeightForWidth();
}

int QTabWidget_HeightForWidth(void* ptr, int width){
	return static_cast<QTabWidget*>(ptr)->heightForWidth(width);
}

int QTabWidget_IndexOf(void* ptr, void* w){
	return static_cast<QTabWidget*>(ptr)->indexOf(static_cast<QWidget*>(w));
}

int QTabWidget_IsTabEnabled(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->isTabEnabled(index);
}

void QTabWidget_RemoveTab(void* ptr, int index){
	static_cast<QTabWidget*>(ptr)->removeTab(index);
}

void QTabWidget_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QTabWidget*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QTabWidget_SetTabEnabled(void* ptr, int index, int enable){
	static_cast<QTabWidget*>(ptr)->setTabEnabled(index, enable != 0);
}

void QTabWidget_SetTabIcon(void* ptr, int index, void* icon){
	static_cast<QTabWidget*>(ptr)->setTabIcon(index, *static_cast<QIcon*>(icon));
}

void QTabWidget_SetTabText(void* ptr, int index, char* label){
	static_cast<QTabWidget*>(ptr)->setTabText(index, QString(label));
}

void QTabWidget_SetTabToolTip(void* ptr, int index, char* tip){
	static_cast<QTabWidget*>(ptr)->setTabToolTip(index, QString(tip));
}

void QTabWidget_SetTabWhatsThis(void* ptr, int index, char* text){
	static_cast<QTabWidget*>(ptr)->setTabWhatsThis(index, QString(text));
}

void* QTabWidget_TabBar(void* ptr){
	return static_cast<QTabWidget*>(ptr)->tabBar();
}

void QTabWidget_ConnectTabBarClicked(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_DisconnectTabBarClicked(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarClicked));;
}

void QTabWidget_ConnectTabBarDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_DisconnectTabBarDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabBarDoubleClicked), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabBarDoubleClicked));;
}

void QTabWidget_ConnectTabCloseRequested(void* ptr){
	QObject::connect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

void QTabWidget_DisconnectTabCloseRequested(void* ptr){
	QObject::disconnect(static_cast<QTabWidget*>(ptr), static_cast<void (QTabWidget::*)(int)>(&QTabWidget::tabCloseRequested), static_cast<MyQTabWidget*>(ptr), static_cast<void (MyQTabWidget::*)(int)>(&MyQTabWidget::Signal_TabCloseRequested));;
}

char* QTabWidget_TabText(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabText(index).toUtf8().data();
}

char* QTabWidget_TabToolTip(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabToolTip(index).toUtf8().data();
}

char* QTabWidget_TabWhatsThis(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->tabWhatsThis(index).toUtf8().data();
}

void* QTabWidget_Widget(void* ptr, int index){
	return static_cast<QTabWidget*>(ptr)->widget(index);
}

void QTabWidget_DestroyQTabWidget(void* ptr){
	static_cast<QTabWidget*>(ptr)->~QTabWidget();
}

#include "qgraphicspolygonitem.h"
#include <QStyleOption>
#include <QGraphicsItem>
#include <QUrl>
#include <QPolygon>
#include <QPolygonF>
#include <QPointF>
#include <QStyle>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QWidget>
#include <QStyleOptionGraphicsItem>
#include <QModelIndex>
#include <QPainter>
#include <QGraphicsPolygonItem>
#include "_cgo_export.h"

class MyQGraphicsPolygonItem: public QGraphicsPolygonItem {
public:
};

int QGraphicsPolygonItem_Contains(void* ptr, void* point){
	return static_cast<QGraphicsPolygonItem*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QGraphicsPolygonItem_FillRule(void* ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->fillRule();
}

int QGraphicsPolygonItem_IsObscuredBy(void* ptr, void* item){
	return static_cast<QGraphicsPolygonItem*>(ptr)->isObscuredBy(static_cast<QGraphicsItem*>(item));
}

void QGraphicsPolygonItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsPolygonItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsPolygonItem_SetFillRule(void* ptr, int rule){
	static_cast<QGraphicsPolygonItem*>(ptr)->setFillRule(static_cast<Qt::FillRule>(rule));
}

void QGraphicsPolygonItem_SetPolygon(void* ptr, void* polygon){
	static_cast<QGraphicsPolygonItem*>(ptr)->setPolygon(*static_cast<QPolygonF*>(polygon));
}

int QGraphicsPolygonItem_Type(void* ptr){
	return static_cast<QGraphicsPolygonItem*>(ptr)->type();
}

void QGraphicsPolygonItem_DestroyQGraphicsPolygonItem(void* ptr){
	static_cast<QGraphicsPolygonItem*>(ptr)->~QGraphicsPolygonItem();
}

#include "qtoolbar.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIcon>
#include <QAction>
#include <QString>
#include <QSize>
#include <QObject>
#include <QWidget>
#include <QMetaObject>
#include <QPoint>
#include <QToolBar>
#include "_cgo_export.h"

class MyQToolBar: public QToolBar {
public:
void Signal_ActionTriggered(QAction * action){callbackQToolBarActionTriggered(this->objectName().toUtf8().data(), action);};
void Signal_AllowedAreasChanged(Qt::ToolBarAreas allowedAreas){callbackQToolBarAllowedAreasChanged(this->objectName().toUtf8().data(), allowedAreas);};
void Signal_MovableChanged(bool movable){callbackQToolBarMovableChanged(this->objectName().toUtf8().data(), movable);};
void Signal_OrientationChanged(Qt::Orientation orientation){callbackQToolBarOrientationChanged(this->objectName().toUtf8().data(), orientation);};
void Signal_ToolButtonStyleChanged(Qt::ToolButtonStyle toolButtonStyle){callbackQToolBarToolButtonStyleChanged(this->objectName().toUtf8().data(), toolButtonStyle);};
void Signal_TopLevelChanged(bool topLevel){callbackQToolBarTopLevelChanged(this->objectName().toUtf8().data(), topLevel);};
void Signal_VisibilityChanged(bool visible){callbackQToolBarVisibilityChanged(this->objectName().toUtf8().data(), visible);};
};

int QToolBar_AllowedAreas(void* ptr){
	return static_cast<QToolBar*>(ptr)->allowedAreas();
}

int QToolBar_IsFloatable(void* ptr){
	return static_cast<QToolBar*>(ptr)->isFloatable();
}

int QToolBar_IsFloating(void* ptr){
	return static_cast<QToolBar*>(ptr)->isFloating();
}

int QToolBar_IsMovable(void* ptr){
	return static_cast<QToolBar*>(ptr)->isMovable();
}

int QToolBar_Orientation(void* ptr){
	return static_cast<QToolBar*>(ptr)->orientation();
}

void QToolBar_SetAllowedAreas(void* ptr, int areas){
	static_cast<QToolBar*>(ptr)->setAllowedAreas(static_cast<Qt::ToolBarArea>(areas));
}

void QToolBar_SetFloatable(void* ptr, int floatable){
	static_cast<QToolBar*>(ptr)->setFloatable(floatable != 0);
}

void QToolBar_SetIconSize(void* ptr, void* iconSize){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setIconSize", Q_ARG(QSize, *static_cast<QSize*>(iconSize)));
}

void QToolBar_SetMovable(void* ptr, int movable){
	static_cast<QToolBar*>(ptr)->setMovable(movable != 0);
}

void QToolBar_SetOrientation(void* ptr, int orientation){
	static_cast<QToolBar*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QToolBar_SetToolButtonStyle(void* ptr, int toolButtonStyle){
	QMetaObject::invokeMethod(static_cast<QToolBar*>(ptr), "setToolButtonStyle", Q_ARG(Qt::ToolButtonStyle, static_cast<Qt::ToolButtonStyle>(toolButtonStyle)));
}

int QToolBar_ToolButtonStyle(void* ptr){
	return static_cast<QToolBar*>(ptr)->toolButtonStyle();
}

void* QToolBar_NewQToolBar2(void* parent){
	return new QToolBar(static_cast<QWidget*>(parent));
}

void* QToolBar_NewQToolBar(char* title, void* parent){
	return new QToolBar(QString(title), static_cast<QWidget*>(parent));
}

void* QToolBar_ActionAt(void* ptr, void* p){
	return static_cast<QToolBar*>(ptr)->actionAt(*static_cast<QPoint*>(p));
}

void* QToolBar_ActionAt2(void* ptr, int x, int y){
	return static_cast<QToolBar*>(ptr)->actionAt(x, y);
}

void QToolBar_ConnectActionTriggered(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

void QToolBar_DisconnectActionTriggered(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(QAction *)>(&QToolBar::actionTriggered), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(QAction *)>(&MyQToolBar::Signal_ActionTriggered));;
}

void* QToolBar_AddAction2(void* ptr, void* icon, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text));
}

void* QToolBar_AddAction4(void* ptr, void* icon, char* text, void* receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(*static_cast<QIcon*>(icon), QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QToolBar_AddAction(void* ptr, char* text){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text));
}

void* QToolBar_AddAction3(void* ptr, char* text, void* receiver, char* member){
	return static_cast<QToolBar*>(ptr)->addAction(QString(text), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QToolBar_AddSeparator(void* ptr){
	return static_cast<QToolBar*>(ptr)->addSeparator();
}

void* QToolBar_AddWidget(void* ptr, void* widget){
	return static_cast<QToolBar*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

void QToolBar_ConnectAllowedAreasChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_DisconnectAllowedAreasChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolBarAreas)>(&QToolBar::allowedAreasChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolBarAreas)>(&MyQToolBar::Signal_AllowedAreasChanged));;
}

void QToolBar_Clear(void* ptr){
	static_cast<QToolBar*>(ptr)->clear();
}

void* QToolBar_InsertSeparator(void* ptr, void* before){
	return static_cast<QToolBar*>(ptr)->insertSeparator(static_cast<QAction*>(before));
}

void* QToolBar_InsertWidget(void* ptr, void* before, void* widget){
	return static_cast<QToolBar*>(ptr)->insertWidget(static_cast<QAction*>(before), static_cast<QWidget*>(widget));
}

int QToolBar_IsAreaAllowed(void* ptr, int area){
	return static_cast<QToolBar*>(ptr)->isAreaAllowed(static_cast<Qt::ToolBarArea>(area));
}

void QToolBar_ConnectMovableChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_DisconnectMovableChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::movableChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_MovableChanged));;
}

void QToolBar_ConnectOrientationChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

void QToolBar_DisconnectOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::Orientation)>(&QToolBar::orientationChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::Orientation)>(&MyQToolBar::Signal_OrientationChanged));;
}

void* QToolBar_ToggleViewAction(void* ptr){
	return static_cast<QToolBar*>(ptr)->toggleViewAction();
}

void QToolBar_ConnectToolButtonStyleChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_DisconnectToolButtonStyleChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(Qt::ToolButtonStyle)>(&QToolBar::toolButtonStyleChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(Qt::ToolButtonStyle)>(&MyQToolBar::Signal_ToolButtonStyleChanged));;
}

void QToolBar_ConnectTopLevelChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_DisconnectTopLevelChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::topLevelChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_TopLevelChanged));;
}

void QToolBar_ConnectVisibilityChanged(void* ptr){
	QObject::connect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

void QToolBar_DisconnectVisibilityChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBar*>(ptr), static_cast<void (QToolBar::*)(bool)>(&QToolBar::visibilityChanged), static_cast<MyQToolBar*>(ptr), static_cast<void (MyQToolBar::*)(bool)>(&MyQToolBar::Signal_VisibilityChanged));;
}

void* QToolBar_WidgetForAction(void* ptr, void* action){
	return static_cast<QToolBar*>(ptr)->widgetForAction(static_cast<QAction*>(action));
}

void QToolBar_DestroyQToolBar(void* ptr){
	static_cast<QToolBar*>(ptr)->~QToolBar();
}

#include "qgraphicsscenemouseevent.h"
#include <QModelIndex>
#include <QGraphicsScene>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsSceneMouseEvent>
#include "_cgo_export.h"

class MyQGraphicsSceneMouseEvent: public QGraphicsSceneMouseEvent {
public:
};

int QGraphicsSceneMouseEvent_Button(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->button();
}

int QGraphicsSceneMouseEvent_Buttons(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->buttons();
}

int QGraphicsSceneMouseEvent_Flags(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->flags();
}

int QGraphicsSceneMouseEvent_Modifiers(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->modifiers();
}

int QGraphicsSceneMouseEvent_Source(void* ptr){
	return static_cast<QGraphicsSceneMouseEvent*>(ptr)->source();
}

void QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(void* ptr){
	static_cast<QGraphicsSceneMouseEvent*>(ptr)->~QGraphicsSceneMouseEvent();
}

#include "qscroller.h"
#include <QVariant>
#include <QUrl>
#include <QMetaObject>
#include <QPointF>
#include <QString>
#include <QRect>
#include <QRectF>
#include <QPoint>
#include <QObject>
#include <QScrollerProperties>
#include <QModelIndex>
#include <QScroller>
#include "_cgo_export.h"

class MyQScroller: public QScroller {
public:
void Signal_StateChanged(QScroller::State newState){callbackQScrollerStateChanged(this->objectName().toUtf8().data(), newState);};
};

void QScroller_SetScrollerProperties(void* ptr, void* prop){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "setScrollerProperties", Q_ARG(QScrollerProperties, *static_cast<QScrollerProperties*>(prop)));
}

void QScroller_EnsureVisible(void* ptr, void* rect, double xmargin, double ymargin){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "ensureVisible", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(qreal, static_cast<qreal>(xmargin)), Q_ARG(qreal, static_cast<qreal>(ymargin)));
}

void QScroller_EnsureVisible2(void* ptr, void* rect, double xmargin, double ymargin, int scrollTime){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "ensureVisible", Q_ARG(QRectF, *static_cast<QRectF*>(rect)), Q_ARG(qreal, static_cast<qreal>(xmargin)), Q_ARG(qreal, static_cast<qreal>(ymargin)), Q_ARG(int, scrollTime));
}

int QScroller_QScroller_GrabGesture(void* target, int scrollGestureType){
	return QScroller::grabGesture(static_cast<QObject*>(target), static_cast<QScroller::ScrollerGestureType>(scrollGestureType));
}

int QScroller_QScroller_GrabbedGesture(void* target){
	return QScroller::grabbedGesture(static_cast<QObject*>(target));
}

int QScroller_QScroller_HasScroller(void* target){
	return QScroller::hasScroller(static_cast<QObject*>(target));
}

void QScroller_ResendPrepareEvent(void* ptr){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "resendPrepareEvent");
}

void QScroller_ScrollTo(void* ptr, void* pos){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)));
}

void QScroller_ScrollTo2(void* ptr, void* pos, int scrollTime){
	QMetaObject::invokeMethod(static_cast<QScroller*>(ptr), "scrollTo", Q_ARG(QPointF, *static_cast<QPointF*>(pos)), Q_ARG(int, scrollTime));
}

void* QScroller_QScroller_Scroller(void* target){
	return QScroller::scroller(static_cast<QObject*>(target));
}

void* QScroller_QScroller_Scroller2(void* target){
	return const_cast<QScroller*>(QScroller::scroller(static_cast<QObject*>(target)));
}

void QScroller_SetSnapPositionsX2(void* ptr, double first, double interval){
	static_cast<QScroller*>(ptr)->setSnapPositionsX(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_SetSnapPositionsY2(void* ptr, double first, double interval){
	static_cast<QScroller*>(ptr)->setSnapPositionsY(static_cast<qreal>(first), static_cast<qreal>(interval));
}

void QScroller_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QScroller*>(ptr), static_cast<void (QScroller::*)(QScroller::State)>(&QScroller::stateChanged), static_cast<MyQScroller*>(ptr), static_cast<void (MyQScroller::*)(QScroller::State)>(&MyQScroller::Signal_StateChanged));;
}

void QScroller_Stop(void* ptr){
	static_cast<QScroller*>(ptr)->stop();
}

void* QScroller_Target(void* ptr){
	return static_cast<QScroller*>(ptr)->target();
}

void QScroller_QScroller_UngrabGesture(void* target){
	QScroller::ungrabGesture(static_cast<QObject*>(target));
}

#include "qstyleoptiongroupbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleOptionGroupBox>
#include "_cgo_export.h"

class MyQStyleOptionGroupBox: public QStyleOptionGroupBox {
public:
};

void* QStyleOptionGroupBox_NewQStyleOptionGroupBox(){
	return new QStyleOptionGroupBox();
}

void* QStyleOptionGroupBox_NewQStyleOptionGroupBox2(void* other){
	return new QStyleOptionGroupBox(*static_cast<QStyleOptionGroupBox*>(other));
}

#include "qgraphicscolorizeeffect.h"
#include <QModelIndex>
#include <QObject>
#include <QColor>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGraphicsColorizeEffect>
#include "_cgo_export.h"

class MyQGraphicsColorizeEffect: public QGraphicsColorizeEffect {
public:
void Signal_ColorChanged(const QColor & color){callbackQGraphicsColorizeEffectColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

void* QGraphicsColorizeEffect_Color(void* ptr){
	return new QColor(static_cast<QGraphicsColorizeEffect*>(ptr)->color());
}

void QGraphicsColorizeEffect_SetColor(void* ptr, void* c){
	QMetaObject::invokeMethod(static_cast<QGraphicsColorizeEffect*>(ptr), "setColor", Q_ARG(QColor, *static_cast<QColor*>(c)));
}

void QGraphicsColorizeEffect_SetStrength(void* ptr, double strength){
	QMetaObject::invokeMethod(static_cast<QGraphicsColorizeEffect*>(ptr), "setStrength", Q_ARG(qreal, static_cast<qreal>(strength)));
}

double QGraphicsColorizeEffect_Strength(void* ptr){
	return static_cast<double>(static_cast<QGraphicsColorizeEffect*>(ptr)->strength());
}

void* QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(void* parent){
	return new QGraphicsColorizeEffect(static_cast<QObject*>(parent));
}

void QGraphicsColorizeEffect_ConnectColorChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsColorizeEffect*>(ptr), static_cast<void (QGraphicsColorizeEffect::*)(const QColor &)>(&QGraphicsColorizeEffect::colorChanged), static_cast<MyQGraphicsColorizeEffect*>(ptr), static_cast<void (MyQGraphicsColorizeEffect::*)(const QColor &)>(&MyQGraphicsColorizeEffect::Signal_ColorChanged));;
}

void QGraphicsColorizeEffect_DisconnectColorChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsColorizeEffect*>(ptr), static_cast<void (QGraphicsColorizeEffect::*)(const QColor &)>(&QGraphicsColorizeEffect::colorChanged), static_cast<MyQGraphicsColorizeEffect*>(ptr), static_cast<void (MyQGraphicsColorizeEffect::*)(const QColor &)>(&MyQGraphicsColorizeEffect::Signal_ColorChanged));;
}

void QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(void* ptr){
	static_cast<QGraphicsColorizeEffect*>(ptr)->~QGraphicsColorizeEffect();
}

#include "qstyleoptionsizegrip.h"
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionSizeGrip>
#include "_cgo_export.h"

class MyQStyleOptionSizeGrip: public QStyleOptionSizeGrip {
public:
};

void* QStyleOptionSizeGrip_NewQStyleOptionSizeGrip(){
	return new QStyleOptionSizeGrip();
}

void* QStyleOptionSizeGrip_NewQStyleOptionSizeGrip2(void* other){
	return new QStyleOptionSizeGrip(*static_cast<QStyleOptionSizeGrip*>(other));
}


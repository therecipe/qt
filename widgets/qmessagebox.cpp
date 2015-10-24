#include "qmessagebox.h"
#include <QUrl>
#include <QCheckBox>
#include <QPixmap>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QPushButton>
#include <QObject>
#include <QAbstractButton>
#include <QMessageBox>
#include "_cgo_export.h"

class MyQMessageBox: public QMessageBox {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQMessageBoxButtonClicked(this->objectName().toUtf8().data(), button);};
};

int QMessageBox_ButtonMask_Type(){
	return QMessageBox::ButtonMask;
}

char* QMessageBox_DetailedText(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->detailedText().toUtf8().data();
}

int QMessageBox_Icon(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->icon();
}

char* QMessageBox_InformativeText(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->informativeText().toUtf8().data();
}

void QMessageBox_SetDetailedText(QtObjectPtr ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setDetailedText(QString(text));
}

void QMessageBox_SetIcon(QtObjectPtr ptr, int v){
	static_cast<QMessageBox*>(ptr)->setIcon(static_cast<QMessageBox::Icon>(v));
}

void QMessageBox_SetIconPixmap(QtObjectPtr ptr, QtObjectPtr pixmap){
	static_cast<QMessageBox*>(ptr)->setIconPixmap(*static_cast<QPixmap*>(pixmap));
}

void QMessageBox_SetInformativeText(QtObjectPtr ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setInformativeText(QString(text));
}

void QMessageBox_SetStandardButtons(QtObjectPtr ptr, int buttons){
	static_cast<QMessageBox*>(ptr)->setStandardButtons(static_cast<QMessageBox::StandardButton>(buttons));
}

void QMessageBox_SetText(QtObjectPtr ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setText(QString(text));
}

void QMessageBox_SetTextFormat(QtObjectPtr ptr, int format){
	static_cast<QMessageBox*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(format));
}

void QMessageBox_SetTextInteractionFlags(QtObjectPtr ptr, int flags){
	static_cast<QMessageBox*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

int QMessageBox_StandardButtons(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->standardButtons();
}

char* QMessageBox_Text(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->text().toUtf8().data();
}

int QMessageBox_TextFormat(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->textFormat();
}

int QMessageBox_TextInteractionFlags(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->textInteractionFlags();
}

QtObjectPtr QMessageBox_NewQMessageBox2(int icon, char* title, char* text, int buttons, QtObjectPtr parent, int f){
	return new QMessageBox(static_cast<QMessageBox::Icon>(icon), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

QtObjectPtr QMessageBox_NewQMessageBox(QtObjectPtr parent){
	return new QMessageBox(static_cast<QWidget*>(parent));
}

void QMessageBox_QMessageBox_About(QtObjectPtr parent, char* title, char* text){
	QMessageBox::about(static_cast<QWidget*>(parent), QString(title), QString(text));
}

void QMessageBox_QMessageBox_AboutQt(QtObjectPtr parent, char* title){
	QMessageBox::aboutQt(static_cast<QWidget*>(parent), QString(title));
}

QtObjectPtr QMessageBox_AddButton3(QtObjectPtr ptr, int button){
	return static_cast<QMessageBox*>(ptr)->addButton(static_cast<QMessageBox::StandardButton>(button));
}

QtObjectPtr QMessageBox_AddButton2(QtObjectPtr ptr, char* text, int role){
	return static_cast<QMessageBox*>(ptr)->addButton(QString(text), static_cast<QMessageBox::ButtonRole>(role));
}

void QMessageBox_AddButton(QtObjectPtr ptr, QtObjectPtr button, int role){
	static_cast<QMessageBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QMessageBox::ButtonRole>(role));
}

QtObjectPtr QMessageBox_Button(QtObjectPtr ptr, int which){
	return static_cast<QMessageBox*>(ptr)->button(static_cast<QMessageBox::StandardButton>(which));
}

void QMessageBox_ConnectButtonClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

void QMessageBox_DisconnectButtonClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

int QMessageBox_ButtonRole(QtObjectPtr ptr, QtObjectPtr button){
	return static_cast<QMessageBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

QtObjectPtr QMessageBox_CheckBox(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->checkBox();
}

QtObjectPtr QMessageBox_ClickedButton(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->clickedButton();
}

int QMessageBox_QMessageBox_Critical(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::critical(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

QtObjectPtr QMessageBox_DefaultButton(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->defaultButton();
}

QtObjectPtr QMessageBox_EscapeButton(QtObjectPtr ptr){
	return static_cast<QMessageBox*>(ptr)->escapeButton();
}

int QMessageBox_Exec(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QMessageBox*>(ptr), "exec");
}

int QMessageBox_QMessageBox_Information(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::information(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QMessageBox*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QMessageBox_QMessageBox_Question(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::question(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_RemoveButton(QtObjectPtr ptr, QtObjectPtr button){
	static_cast<QMessageBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetCheckBox(QtObjectPtr ptr, QtObjectPtr cb){
	static_cast<QMessageBox*>(ptr)->setCheckBox(static_cast<QCheckBox*>(cb));
}

void QMessageBox_SetDefaultButton(QtObjectPtr ptr, QtObjectPtr button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QPushButton*>(button));
}

void QMessageBox_SetDefaultButton2(QtObjectPtr ptr, int button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetEscapeButton(QtObjectPtr ptr, QtObjectPtr button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetEscapeButton2(QtObjectPtr ptr, int button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QMessageBox*>(ptr)->setVisible(visible != 0);
}

void QMessageBox_SetWindowModality(QtObjectPtr ptr, int windowModality){
	static_cast<QMessageBox*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QMessageBox_SetWindowTitle(QtObjectPtr ptr, char* title){
	static_cast<QMessageBox*>(ptr)->setWindowTitle(QString(title));
}

int QMessageBox_StandardButton(QtObjectPtr ptr, QtObjectPtr button){
	return static_cast<QMessageBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

int QMessageBox_QMessageBox_Warning(QtObjectPtr parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::warning(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_DestroyQMessageBox(QtObjectPtr ptr){
	static_cast<QMessageBox*>(ptr)->~QMessageBox();
}


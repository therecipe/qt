#include "qsysinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSysInfo>
#include "_cgo_export.h"

class MyQSysInfo: public QSysInfo {
public:
};

int QSysInfo_ByteOrder_Type(){
	return QSysInfo::ByteOrder;
}

int QSysInfo_MV_IOS_Type(){
	return QSysInfo::MV_IOS;
}

int QSysInfo_MV_IOS_4_3_Type(){
	return QSysInfo::MV_IOS_4_3;
}

int QSysInfo_MV_IOS_5_0_Type(){
	return QSysInfo::MV_IOS_5_0;
}

int QSysInfo_MV_IOS_5_1_Type(){
	return QSysInfo::MV_IOS_5_1;
}

int QSysInfo_MV_IOS_6_0_Type(){
	return QSysInfo::MV_IOS_6_0;
}

int QSysInfo_MV_IOS_6_1_Type(){
	return QSysInfo::MV_IOS_6_1;
}

int QSysInfo_MV_IOS_7_0_Type(){
	return QSysInfo::MV_IOS_7_0;
}

int QSysInfo_MV_IOS_7_1_Type(){
	return QSysInfo::MV_IOS_7_1;
}

int QSysInfo_MV_IOS_8_0_Type(){
	return QSysInfo::MV_IOS_8_0;
}

int QSysInfo_MV_IOS_8_1_Type(){
	return QSysInfo::MV_IOS_8_1;
}

int QSysInfo_MV_IOS_8_2_Type(){
	return QSysInfo::MV_IOS_8_2;
}

int QSysInfo_MV_IOS_8_3_Type(){
	return QSysInfo::MV_IOS_8_3;
}

int QSysInfo_MV_IOS_8_4_Type(){
	return QSysInfo::MV_IOS_8_4;
}

int QSysInfo_MV_IOS_9_0_Type(){
	return QSysInfo::MV_IOS_9_0;
}

int QSysInfo_WordSize_Type(){
	return QSysInfo::WordSize;
}

int QSysInfo_QSysInfo_MacVersion(){
	return QSysInfo::macVersion();
}

char* QSysInfo_QSysInfo_BuildAbi(){
	return QSysInfo::buildAbi().toUtf8().data();
}

char* QSysInfo_QSysInfo_BuildCpuArchitecture(){
	return QSysInfo::buildCpuArchitecture().toUtf8().data();
}

char* QSysInfo_QSysInfo_CurrentCpuArchitecture(){
	return QSysInfo::currentCpuArchitecture().toUtf8().data();
}

char* QSysInfo_QSysInfo_KernelType(){
	return QSysInfo::kernelType().toUtf8().data();
}

char* QSysInfo_QSysInfo_KernelVersion(){
	return QSysInfo::kernelVersion().toUtf8().data();
}

char* QSysInfo_QSysInfo_PrettyProductName(){
	return QSysInfo::prettyProductName().toUtf8().data();
}

char* QSysInfo_QSysInfo_ProductType(){
	return QSysInfo::productType().toUtf8().data();
}

char* QSysInfo_QSysInfo_ProductVersion(){
	return QSysInfo::productVersion().toUtf8().data();
}

int QSysInfo_QSysInfo_WindowsVersion(){
	return QSysInfo::windowsVersion();
}


﻿#include "Student.h"
#include <QDebug>

Student::Student(QObject* parent)
	: QObject(parent)
{
}

void Student::treat() {
	qDebug() << "请客吃饭";
}

void Student::treat(QString foodName) {
	qDebug() << "请老师吃" << foodName.toUtf8().data();
}
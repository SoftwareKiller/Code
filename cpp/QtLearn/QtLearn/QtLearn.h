#pragma once

#include <QtWidgets/QWidget>
#include <QtWidgets/qlabel.h>
#include "ui_QtLearn.h"
#include "Student.h"
#include "Teacher.h"


class QtLearn : public QWidget
{
	Q_OBJECT

public:
	QtLearn(QWidget* parent = nullptr);
	~QtLearn();

private:
	Ui::QtLearnClass ui;
	QLabel* m_lableInfo;

	Teacher* t;
	Student* st;

private:
	void afterClass();
};

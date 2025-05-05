#pragma once
#pragma execution_character_set("utf-8")

#include <QtWidgets/QMainWindow>
#include "ui_QtDialog.h"

class QtDialog : public QMainWindow
{
	Q_OBJECT

public:
	QtDialog(QWidget* parent = nullptr);
	~QtDialog();

private:
	Ui::QtDialogClass ui;
};

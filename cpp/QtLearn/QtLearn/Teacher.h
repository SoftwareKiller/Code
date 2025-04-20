#pragma once

#include <QObject>

class Teacher : public QObject
{
	Q_OBJECT;
public:
	explicit Teacher(QObject* parent = 0);

signals:
	void hungry();

	void hungry(QString foodName);
};

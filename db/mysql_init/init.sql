CREATE DATABASE IF NOT EXISTS earlGrey;
use earlGrey;

--何限目の情報を格納している
CREATE TABLE timer(
	timeNo char(3) not null,
	sTime char(5) not null,
	eTime char(5)not null,
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(timeNo)
);

--data
INSERT INTO timer(timeNo, sTime, eTime)
  VALUES("1限目", "09:15", "10:45"),
        ("2限目", "11:00", "12:30"),
        ("3限目", "13:30", "15:00"),
        ("4限目", "15:15", "16:45"),
        ("5限目", "17:00", "18:30");

--先生たちの権限情報
CREATE TABLE permission(
	perNo int auto_increment,
	permission char(2),
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(perNo)
);

--data
INSERT INTO permission(permission)
  VALUES("予約"),
        ("申請");

--先生の情報
CREATE TABLE teachers(
	teacherNo int auto_increment,
	name varchar(20) not null,
	perNo int not null,
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(teacherNo),
	foreign key(perNo) references permission(perNo)
);

--sampel data
INSERT INTO teachers(name, perNo)
  VALUES("内山豊彦", 1),
        ("武次順平", 1),
        ("小戎冴茄", 2);

--予約申請の状態を格納している（承認など）
CREATE TABLE state(
	stateNo int auto_increment,
	stateName char(4) not null,
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(stateNo)
);

--data
INSERT INTO state(stateName)
  VALUES("承認済み"),
        ("承認待ち"),
        ("否認");

--教室の情報
CREATE TABLE rooms(
	roomNo char(4),
	memo varchar(255),
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(roomNo)
);

--data
INSERT INTO rooms(roomNo, memo)
  VALUES("1204", "コンセントプラグ：床"),
        ("1205", "コンセントプラグ：床"),
        ("2031", "コンセントプラグ：机の上, ネットワーク機器あり");

--時間割り
CREATE TABLE timetable(
	No int auto_increment,
	class char(5) not null,
	roomNo char(4) not null,
	name varchar(40) not null,
	youbi char(3) not null,
	teacherNo int not null,
	timeNo char(3) not null,
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(No),
	foreign key(roomNo) references rooms(roomNo),
	foreign key(teacherNo) references teachers(teacherNo),
	foreign key(timeNo) references timer(timeNo)
);

--sample data
INSERT INTO timetable(class, roomNo, name, youbi, teacherNo, timeNo)
  VALUES("IE4A", "1205", "システム開発演習５", "Fri", 1, "3限目"),
        ("IE4A", "1205", "システム開発演習５", "Fri", 1, "4限目");

--予約
CREATE TABLE reservation(
	reseNo int auto_increment,
	teacherNo int not null,
	roomNo char(4) not null,
	reseDate date not null,
	sTime time not null,
	eTime time not null,
	purpose varchar(150) not null,
	requestDate date not null,
	stateNo int not null,
	update_at datetime,
	create_at datetime,
	delete_at datetime,

	primary key(reseNo),
	foreign key(teacherNo) references teachers(teacherNo),
	foreign key(roomNo) references rooms(roomNo),
	foreign key(stateNo) references state(stateNo)
);

--sample data
INSERT INTO reservation(teacherNo, roomNo, reseDate, sTime, eTime, purpose, requestDate, stateNo)
  VALUES(1, "1204", "2022-06-01", "12:00", "13:00", "面談", "2022-05-27", 2);

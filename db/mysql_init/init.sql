CREATE DATABASE IF NOT EXISTS earlGrey;
use earlGrey;

-- 何限目の情報を格納している
CREATE TABLE timers(
	time_no char(3) not null,
	s_time char(5) not null,
	e_time char(5)not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(time_no)
);

-- data
INSERT INTO timers(time_no, s_time, e_time)
  VALUES("1限目", "09:15", "10:45"),
        ("2限目", "11:00", "12:30"),
        ("3限目", "13:30", "15:00"),
        ("4限目", "15:15", "16:45"),
        ("5限目", "17:00", "18:30");

-- 先生たちの権限情報
CREATE TABLE permissions(
	per_no int auto_increment,
	permission char(2) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(per_no)
);

-- data
INSERT INTO permissions(permission)
  VALUES("予約"),
        ("申請");

-- 先生の情報
CREATE TABLE teachers(
	teacher_no int auto_increment,
	teacher_name varchar(20) not null,
	per_no int not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(teacher_no),
	foreign key(per_no) references permissions(per_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sampel data
INSERT INTO teachers(teacher_name, per_no)
  VALUES("内山豊彦", 1),
        ("武次順平", 1),
		("杉原宏", 1),
		("上村香代子", 1),
	("小戎冴茄", 2);

INSERT INTO teachers(teacher_name, per_no)
	VALUES("山本太", 1),
				("加藤昌", 1),
				("石田雄太", 1),
				("小出操", 1),
				("曽根国雄", 2);

-- 予約申請の状態を格納している（承認など）
CREATE TABLE states(
	state_no int auto_increment,
	state_name char(4) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(state_no)
);

-- data
INSERT INTO states(state_name)
  VALUES("承認済み"),
        ("承認待ち"),
        ("否認");

-- 教室の情報
CREATE TABLE rooms(
	room_no char(4),
	memo varchar(255),
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(room_no)
);

-- data
INSERT INTO rooms(room_no, memo)
  VALUES("1204", "コンセントプラグ：床"),
        ("1205", "コンセントプラグ：床"),
				("2301", "コンセントプラグ：机の上, ネットワーク機器あり"),
				("4301", "コンセントプラグ：床"),
				("3301", "コンセントプラグ：机の横"),
				("2302", "コンセントプラグ：机の上"),
				("4203", "コンセントプラグ：床"),
		("2031", "コンセントプラグ：机の上, ネットワーク機器あり");

INSERT INTO rooms(room_no, memo)
	VALUES("2405", "コンセントプラグ：机の横"),
				("2402", "コンセントプラグ：床"),
				("3502", "コンセントプラグ：机の上"),
				("1403", "コンセントプラグ：机の上"),
				("4202", "コンセントプラグ：床");


-- 時間割り
CREATE TABLE timetables(
	No int auto_increment,
	room_no char(4) not null,
	subject_name varchar(40) not null,
	youbi char(3) not null,
	teacher_no int not null,
	time_no char(3) not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(No),
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(teacher_no) references teachers(teacher_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(time_no) references timers(time_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sample data
INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
  VALUES("4301", "セキュリティ演習_A", "Mon", 1, "1限目"),
				("4301", "セキュリティ演習_A", "Mon", 1, "2限目"),
				("1205", "ITシステム開発演習Ⅴ", "Tue", 1, "1限目"),
				("1205", "ITシステム開発演習Ⅴ", "Tue", 1, "2限目"),
				("2301", "ITゼミ演習Ⅲ", "Tue", 1, "3限目"),
				("2301", "ITゼミ演習Ⅲ", "Tue", 1, "4限目"),
				("3301", "就職対策Ⅱ", "Wed", 5, "3限目"),
				("2302", "システム設計演習", "Wed", 4, "3限目"),
				("2302", "システム設計演習", "Wed", 4, "4限目"),
				("2301", "ITゼミ演習", "Thu", 2, "1限目"),
				("2301", "ITゼミ演習", "Thu", 2, "2限目"),
				("4203", "AIシステム開発演習Ⅰ", "Thu", 2, "3限目"),
				("4203", "AIシステム開発演習Ⅰ", "Thu", 2, "4限目"),
				("1205", "ITシステム開発演習Ⅴ", "Fri", 1, "3限目"),
        		("1205", "ITシステム開発演習Ⅴ", "Fri", 1, "4限目"),
				("3301", "ハイプロフェッショナルゼミ_IT", "Fri", 1, "5限目");

INSERT INTO timetables(room_no, subject_name, youbi, teacher_no, time_no)
	VALUES("2405", "就職対策Ⅱ", "Mon", 6, "1限目"),
				("4202", "セキュリティ演習_A", "Mon", 7, "2限目"),
				("4202", "セキュリティ演習_A", "Mon", 7, "3限目"),
				("1204", "ITシステム開発演習Ⅴ", "Tue", 1, "1限目"),
				("1204", "ITシステム開発演習Ⅴ", "Tue", 1, "2限目"),
				("3502", "ITゼミ演習Ⅲ", "Tue", 8, "3限目"),
				("3502", "ITゼミ演習Ⅲ", "Tue", 8, "4限目"),
				("2302", "システム設計演習", "Wed", 4, "1限目"),
				("2302", "システム設計演習", "Wed", 4, "2限目"),
				("3502", "ITゼミ演習Ⅲ", "Thu", 9, "3限目"),
				("3502", "ITゼミ演習Ⅲ", "Thu", 9, "4限目"),
				("1403", "AIシステム開発演習Ⅰ", "Fri", 2, "1限目"),
				("1403", "AIシステム開発演習Ⅰ", "Fri", 2, "2限目"),
				("1204", "ITシステム開発演習Ⅴ", "Tue", 1, "3限目"),
				("1204", "ITシステム開発演習Ⅴ", "Tue", 1, "4限目");

-- 予約
CREATE TABLE reservations(
	rese_no int auto_increment,
	teacher_no int not null,
	room_no char(4) not null,
	rese_date date not null,
	s_time time not null,
	e_time time not null,
	purpose varchar(150) not null,
	request_date date not null,
	state_no int not null,
	updated_at datetime,
	created_at datetime,
	deleted_at datetime,

	primary key(rese_no),
	foreign key(teacher_no) references teachers(teacher_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(room_no) references rooms(room_no) ON DELETE CASCADE ON UPDATE CASCADE,
	foreign key(state_no) references states(state_no) ON DELETE CASCADE ON UPDATE CASCADE
);

-- sample data
INSERT INTO reservations(teacher_no, room_no, rese_date, s_time, e_time, purpose, request_date, state_no)
  VALUES(1, "1204", "2022-06-01", "12:00", "13:00", "面談", "2022-05-27", 2);

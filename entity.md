# DBのEntity

## データベース名

earlGray

## テーブル

updated_at, created_at, deleted_atはすべてのテーブルに存在する。

データ型：datetime

### timetable(時間割)
- No
  - primary key
  - int
  - auto_increment
  - 時間割を特定するための値

- room_no
  - char(4)
  - not null
  - foreign key：rooms表のroomNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 教室の番号

- subject_name
  - varchar(40)
  - not null
  - 授業名

- youbi
  - char(3)
  - not null
  - 曜日
  - 書き方：英語表記の先頭3文字(例->Mon)

- teacher_no
  - int
  - not null
  - foreign key：teachers表のteacherNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 先生の番号

- time_no
  - char(3)
  - not null
  - foreign key：timer表のtimeNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - ○限目の表示

### reservation(予約)

- rese_no
  - primary key
  - int
  - auto_increment
  - 予約を特定するための値

- teacher_no
  - int
  - not null
  - foreign key：teachers表のteacherNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 予約申請した先生
- room_no
  - char(4)
  - not null
  - foreign key：rooms表のroomNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 予約したい教室番号
- rese_date
  - date
  - not null
  - 予約したい日程
- s_time
  - time
  - not null
  - 開始時間
- e_time
  - time
  - not null
  - 終了時間
- purpose
  - varchar(150)
  - not null
  - 予約理由
- request_date
    - date
    - not null
    - 予約申請をした日時
- state_no
  - int
  - not null
  - foreign key：state表のstateNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 予約の状態を表す（予約権限を持った先生が承認したかどうか）

### timers(時間)
※何限目と決まっている時間のみ(時間割作成で使用する)
- time_no
  - char(3)
  - not null
  - ○限目
- s_time
  - char(5)
  - not null
  - 開始時間
- e_time
  - char(5)
  - not null
  - 終了時間

### permission(権限)
※先生が予約する時の権限
- per_no
  - primary key
  - int
  - auto_increment
  - 権限を特定するための値
- permission
  - char(2)
  - not null
  - 権限

### teachers(先生)
- teacher_no
  - primary key
  - int
  - auto_increment
  - 先生を特定するための値
- teacher_name
  - varchar(20)
  - not null
  - 先生の名前
- per_no
  - int
  - not null
  - foreign key：parmission表のperNo
    - DELETE：CASCADE, UPDATE：CASCADE
  - 先生の権限を表す

### state(予約申請の状態)
- state_no
  - primary key
  - int
  - auto_increment
  - 予約申請の状態を特定するための値
- state_name
  - char(4)
  - not null
  - 申請の状態

### rooms(教室)
- room_no
  - primary key
  - char(4)
  - 教室番号
- memo
  - varchar(255)
  - 教室設備などを書き留めるメモ欄
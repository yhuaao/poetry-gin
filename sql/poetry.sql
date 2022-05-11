create table admins(
id int(11) not null primary key auto_increment,
pid int(11) not null default 0,
phone varchar(20) not null default '',
nick_name varchar(20) not null default '',
`password` varchar(255) not null default '',
head_img varchar(255) not null default '',
`status` tinyint(1) not null default 0,
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;

create table roles(
id int(11) not null primary key auto_increment,
role_name varchar(255) not null default '',
`status` tinyint(1) not null default 0,
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;

create table admins_roles(
admin_id int(11) not null default 0,
role_id int(11) not null default 0    
)engine=innodb charset=utf8mb4;


create table `permissions`(
id int(11) not null primary key auto_increment,
permission_name varchar(255) not null default '',
`status` tinyint(1) not null default 0,
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;

create table roles_permissions(
role_id int(11) not null default 0,
permission_id int(11) not null default 0    
)engine=innodb charset=utf8mb4;


create table actions(
id int(11) not null primary key auto_increment,
action_name varchar(255) not null default '',
`url` varchar(255) not null default '',
`status` tinyint(1) not null default 0,
`desc` varchar(255) not null default '',
method enum('POST','GET','PUT','DELETE') not null default 'GET',
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;


create table permissions_actions(
permission_id int(11) not null default 0,
action_id int(11) not null default 0    
)engine=innodb charset=utf8mb4;

create table menus(
id int(11) not null primary key auto_increment,
pid int(11) not null default 0 ,
menu_name varchar(255) not null default '',
`type` tinyint(1) not null default 0 comment '0-menu 1-dropdown',
`url` varchar(255) not null default '',
`status` tinyint(1) not null default 0,
icon  varchar(255) not null default '',
`desc` varchar(255) not null default '',
sort smallint(6) not null default 0,
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;

create table permissions_menus(
permission_id int(11) not null default 0,
menu_id int(11) not null default 0    
)engine=innodb charset=utf8mb4;



create table websites(
id int(11) not null primary key auto_increment,
website_name varchar(255) not null default '',
domain varchar(255) not null default '',
`status` tinyint(1) not null default 0,
`title` varchar(255) not null default '',
`desc` varchar(255) not null default '',
`keyword` varchar(255) not null default '',
`remark`  varchar(255) not null default '',
created_at int(11) not null default 0 ,
updated_at int(11) not null default 0
)engine=innodb charset=utf8mb4;

create table admins_websites(
admin_id int(11) not null default 0,
web_id int(11) not null default 0
)engine=innodb charset=utf8mb4;


create table categorys(
id int(11) not null primary key auto_increment,
web_id int(11) not null default 0,
pid int(11) not null default 0 ,
category_name varchar(255) not null default '',
`type` tinyint(1) not null default 0 comment '0-menu 1-dropdown',
`icon` varchar(255) not null default '',
`status` tinyint(1) not null default 0 ,
`sort` int(11) not null default 0,
`remark` varchar(255) not null default '',
`url` varchar(255) not null default '',
`title` varchar(255) not null default '',
`desc` varchar(255) not null default '',
`keyword` varchar(255) not null default '',
created_at int(11) not null default 0,
updated_at int(11) not null default 0
)engine=innodb charset=utf8mb4;


create table banner(
id int(11) not null primary key auto_increment,
category_id int(11) not null default 0,
banner_name varchar(255) not null default '',
`url` varchar(255) not null default '',
`status` tinyint(1) not null default 0 ,
`sort`  int(11) not null default 0,
created_at int(11) not null default 0,
updated_at int(11) not null default 0
)engine=innodb charset=utf8mb4;


create table articles(
id int(11) not null primary key auto_increment,
category_id int(11) not null default 0,
article_name varchar(255) not null default '',
author varchar(255) not null default '',
source varchar(255) not null default '',
pub_time int(11) not null default 0,
`type` tinyint(1) not null default 0,
`url` varchar(255) not null default '',
`status` tinyint(1) not null default 0 ,
`sort` int(11) not null default 0,
`hot` tinyint(1) not null default 0 ,
`collect_url` varchar(255) not null default '',
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;

create table article_contexts(
id int(11) not null primary key auto_increment,
article_id int(11) not null default 0,
context text default null,
created_at int(11) not null default 0,
updated_at int(11) not null default 0 
)engine=innodb charset=utf8mb4;


create table users(
id int(11) not null primary key auto_increment,
phone varchar(255) not null default '',
`password` varchar(255) not null default ''
)engine=innodb charset=utf8mb4;


-- 答题  稀疏数组

-- teacher 相关  题库  

-- 在线编辑，使用在线office  题库

-- 直播服务 ffmpeg 视频处理服务


















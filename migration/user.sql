CREATE TABLE `users` (
	`ID` bigint(20) NOT NULL AUTO_INCREMENT,
	`UserName` varchar(32) NOT NULL COMMENT '用户姓名',
	`Password` varchar(128) NOT NULL COMMENT '用户密码',
	`PasswordSalt` varchar(128) DEFAULT NULL COMMENT '密码Salt',
	`Mobile` varchar(16) DEFAULT NULL COMMENT '手机',
	`Email` varchar(256) NOT NULL COMMENT 'Email',
	`IsApproved` tinyint(1) NOT NULL COMMENT '是否审核过',
	`IsLocked` tinyint(1) NOT NULL COMMENT '是否锁定状态',
	`LastLoginDate` datetime DEFAULT NULL COMMENT '最后登录时间',
	`LastPasswordChangedDate` datetime DEFAULT NULL COMMENT '最后密码修改时间',
	`LastLockoutDate` datetime DEFAULT NULL COMMENT '最后锁定时间',
	`Remark` varchar(256) DEFAULT NULL COMMENT '备注',
	`CreatedOn` datetime NOT NULL COMMENT '创建时间',
	`IsDeleted` tinyint(1) NOT NULL DEFAULT '0',
	`UpdatedOn` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`ID`),
	KEY `idx_email` (`Email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=989187 DEFAULT CHARSET=utf8 COMMENT='用户主表';


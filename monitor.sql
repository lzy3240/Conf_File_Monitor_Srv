/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 80016
Source Host           : 127.0.0.1:3306
Source Database       : m_monitor

Target Server Type    : MYSQL
Target Server Version : 80016
File Encoding         : 65001

Date: 2020-07-27 13:18:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for file_monitor
-- ----------------------------
DROP TABLE IF EXISTS `file_monitor`;
CREATE TABLE `file_monitor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '服务所在机器IP',
  `filename` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '监听文件夹，“/”拼接',
  `stage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '监听的文件名后缀，可多个',
  `flag` varchar(255) DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for file_monitor_res
-- ----------------------------
DROP TABLE IF EXISTS `file_monitor_res`;
CREATE TABLE `file_monitor_res` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '服务所在机器IP',
  `filename` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '监听文件夹，“/”拼接',
  `flag` varchar(255) DEFAULT NULL,
  `variation` varchar(255) DEFAULT NULL,
  `varytime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;


-- ----------------------------
-- Table structure for florder_monitor
-- ----------------------------
DROP TABLE IF EXISTS `florder_monitor`;
CREATE TABLE `florder_monitor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `host` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '服务所在机器IP',
  `florder` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '监听文件夹，“/”拼接',
  `flag` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '标记，florder：监听文件夹本身；dateflorder：监听文件夹下当天日期命名的子文件夹',
  `stage` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '监听的文件名后缀，可多个',
  `num` int(11) unsigned zerofill DEFAULT NULL COMMENT '监听的结果',
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8;

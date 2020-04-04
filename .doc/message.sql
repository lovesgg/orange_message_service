CREATE TABLE `message` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `to_user` varchar(45) NOT NULL,
  `body` text,
  `c_t` int(11) DEFAULT NULL,
  `u_t` int(11) DEFAULT NULL,
  `send_status` int(1) NOT NULL,
  `source_id` int(4) NOT NULL,
  PRIMARY KEY (`id`,`to_user`,`send_status`,`source_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
SELECT * FROM testmysql.message;
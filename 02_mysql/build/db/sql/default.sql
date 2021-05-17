CREATE TABLE IF NOT EXISTS `users` (
  `id` integer NOT NULL,
  `name` varchar(256) NOT NULL,
  `age` integer(2),
  `address` varchar(256),
  `update_at` DATETIME,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

-- usersテーブルに２つレコードを追加
INSERT INTO users (id, name, age, address) VALUES ("00001", "ENDO", 30, "高円寺");
INSERT INTO users (id, name, age, address) VALUES ("00002", "ISHIYAMA", 26, "下北沢");

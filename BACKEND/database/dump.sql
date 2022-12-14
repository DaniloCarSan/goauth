CREATE TABLE IF NOT EXISTS `sigppci`.`users` (
  `CODE` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `NAME` VARCHAR(60) NOT NULL,
  `EMAIL` VARCHAR(80) NOT NULL,
  `PASSWORD` VARCHAR(100) NOT NULL,
  PRIMARY KEY (`CODE`),
  UNIQUE INDEX `CODE_UNIQUE` (`CODE` ASC) VISIBLE,
  UNIQUE INDEX `EMAIL_UNIQUE` (`EMAIL` ASC) VISIBLE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8

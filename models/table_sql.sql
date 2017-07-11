
    -- --------------------------------------------------
    --  Table Structure for `blog/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `user` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL DEFAULT '' UNIQUE ,
        `pass_word` varchar(255) NOT NULL DEFAULT '' ,
        `avatar` varchar(255) NOT NULL DEFAULT '' ,
        `signature` varchar(255) NOT NULL DEFAULT '' ,
        `email` varchar(255) NOT NULL DEFAULT '' ,
        `web_site` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';

    -- --------------------------------------------------
    --  Table Structure for `blog/models.TmpUser`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `tmp_user` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `email` varchar(255) NOT NULL DEFAULT '' ,
        `web_site` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='匿名用户表';

    -- --------------------------------------------------
    --  Table Structure for `blog/models.ArticleInfo`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article_info` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `sub_title` varchar(255) NOT NULL DEFAULT '' ,
        `author` varchar(255) NOT NULL DEFAULT '' ,
        `tags` varchar(255) NOT NULL DEFAULT '' ,
        `categories` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.ArticleContent`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article_content` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `article_id` integer NOT NULL DEFAULT 0 ,
        `content` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.ArticleStat`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article_stat` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `article_id` integer NOT NULL DEFAULT 0 ,
        `comment_count` integer NOT NULL DEFAULT 0 ,
        `favour_count` integer NOT NULL DEFAULT 0 ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Comment`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `comment` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `user_type` integer NOT NULL DEFAULT 0 ,
        `user_id` integer NOT NULL DEFAULT 0 ,
        `article_id` integer NOT NULL DEFAULT 0 ,
        `content` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.CommentStat`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `comment_stat` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `comment_id` integer NOT NULL DEFAULT 0 ,
        `favour_count` integer NOT NULL DEFAULT 0 ,
        `low_count` integer NOT NULL DEFAULT 0 ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Category`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `category` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `sort_id` integer NOT NULL DEFAULT 0 ,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `count` integer NOT NULL DEFAULT 0 ,
        `description` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `update` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.TagInfo`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `tag_info` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `count` integer NOT NULL DEFAULT 0 ,
        `description` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Social`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `social` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `sort_id` integer NOT NULL DEFAULT 0 ,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `extra_url` varchar(255) NOT NULL DEFAULT '' ,
        `icon` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Topic`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `topic` (
        `id` INTEGER AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` VARCHAR(255) NOT NULL DEFAULT '',
        `slug` VARCHAR(255) NOT NULL DEFAULT '',
        `desc` VARCHAR(255) NOT NULL DEFAULT '',
        `article_count` INTEGER NOT NULL  DEFAULT 0,
        `create` BIGINT NOT NULL DEFAULT 0,
        `updated` BIGINT NOT NULL DEFAULT 0,
        `status` INTEGER NOT NULL DEFAULT 0
    )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='专题表';
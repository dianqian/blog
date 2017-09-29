
    -- --------------------------------------------------
    --  Table Structure for `blog/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `user` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL DEFAULT '' UNIQUE ,
        `pass_word` varchar(255) NOT NULL DEFAULT '' ,
        `nick_name` varchar(255) NOT NULL  DEFAULT '',
        `avatar` varchar(255) NOT NULL DEFAULT '' ,
        `signature` varchar(255) NOT NULL DEFAULT '' ,
        `email` varchar(255) NOT NULL DEFAULT '' ,
        `web_site` varchar(255) NOT NULL DEFAULT '' ,
        `wechat` varchar(255) NOT NULL  DEFAULT '',
        `login_time` bigint NOT NULL  DEFAULT 0,
        `last_login_time` bigint NOT NULL  DEFAULT 0,
        `last_logout_time` bigint NOT NULL  DEFAULT 0,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户表';

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Blogger`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `blogger` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `sub_title` varchar(255) NOT NULL DEFAULT '' ,
        `icon` varchar(255) NOT NULL  DEFAULT '',
        `bei_an` varchar(255) NOT NULL DEFAULT '' ,
        `copyright` varchar(255) NOT NULL DEFAULT '' ,
        `series_say` varchar(255) NOT NULL DEFAULT '' ,
        `archives_say` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='blogger的配置信息';

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
    --  Table Structure for `blog/models.Article`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `title` varchar(255) NOT NULL DEFAULT '' ,
        `url` varchar(255) NOT NULL DEFAULT '' ,
        `author` integer NOT NULL,
        `publish_time` bigint NOT NULL DEFAULT 0 ,
        `content` LONGTEXT,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='文章article';

    -- --------------------------------------------------
    --  Table Structure for `blog/models.ArticleTopic`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article_topic` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `article_id` integer NOT NULL DEFAULT 0 ,
        `topic_id` integer NOT NULL DEFAULT 0 ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='关联article topic表';

    -- --------------------------------------------------
    --  Table Structure for `blog/models.ArticleTag`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `article_tag` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `article_id` integer NOT NULL DEFAULT 0 ,
        `tag_id` integer NOT NULL DEFAULT 0 ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='关联article tag表';

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

    -- --------------------------------------------------
    --  Table Structure for `blog/models.Tag`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `tag` (
        `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL DEFAULT '' ,
        `create` bigint NOT NULL DEFAULT 0 ,
        `updated` bigint NOT NULL DEFAULT 0 ,
        `status` integer NOT NULL DEFAULT 0
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='专题表';
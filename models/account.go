package models

import "time"

type Account struct {
    UserName        string
    Password        string
    Avatar          string
    Signature       string
    Email           string
    PhoneNumber     string
    Address         string

    CreateTime      time.Time
    LoginTime       time.Time
    LogoutTime      time.Time
    LoginIp         string

    Blogger
}

type Blogger struct {
    BlogName        string
    SubTitle        string

    BeiAn           string
    BTitle          string
    Copyright       string

    SeriesSay       string
    ArchivesSay     string
}

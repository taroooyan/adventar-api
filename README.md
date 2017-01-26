# adventar-api

## Description
API of http://www.adventar.org

## Usage
Return data is JSON.
```
Is_error: bool,
Title: string,
Url: string,
Creator: string,
Description: string,
Entry_count: int,
Posted_count: int,
Calendars: [
    {
        Date: int,
        User: string,
        Icon: string,
        Comment: string,
        Title: string,
        Url: string,
        Is_entry: bool,
        Is_posted: bool
    },
    ...
]
```

Access URL is /adventar/ARTICLE-NUMBER
ARTICLE-NUMBER is *** of `adventaradventar.org/calendars/***`

## LICENSE
MIT

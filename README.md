# adventar-api
![CircleCI](https://circleci.com/gh/taroooyan/adventar-api.svg?style=shield&circle-token=4f414c66211bee0d7e41206a1db98fa157422729)

## Description
API of http://www.adventar.org
This API can get each advent calendars info.

It is run https://adventar-api.appspot.com

Access URL is /ARTICLE-NUMBER. ARTICLE-NUMBER is *** of adventaradventar.org/calendars/***

Example  
To get info of [Aizu Advent Calendar 2015](http://www.adventar.org/calendars/888)  
Access URL [https://adventar-api.appspot.com/888](https://adventar-api.appspot.com/888).

## Data
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

## LICENSE
MIT

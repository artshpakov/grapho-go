# Grapho API Description

### Users

```
/graph?q={users{id,email,name,profile{firstname,lastname,location,bio},posts{id,title,text,created_at,slug}}}
/graph?q={users(id:3){id,email,name,slug}}
/graph?q={users(id:[3,4]){id,email,name}}
```


```json
{
  "data": {
    "users": [{
      "email": "vombat@local.host",
      "id": 3,
      "name": "Vombat Local",
      "posts": [{
        "created_at": "2017-04-07 07:08:15.666531 +0300 EEST",
        "id": 16,
        "text": "The quick brown fox jumps over the lazy dog",
        "title": "My First Blog Post",
        "slug": "my-first-blog-post",
      }],
      "profile": {
        "bio": "Live Today!",
        "firstname": "Vombat",
        "lastname": "Local",
        "location": "Vilnius, Lithuania"
      },
      "slug": "vombat",
    }]
  }
}
```

### Posts

```
/graph?q={posts{id,title,text,user_id,created_at,slug}}
/graph?q={posts(limit:10){id,title,text,user_id}}
/graph?q={posts(user_id:3){id,title,text,user_id}}
/graph?q={posts(id:16){id,title,text}}
/graph?q={posts(id:[12,16]){id,title,text,user_id}}
```

```json
{
  "data": {
    "posts": [{
      "created_at": "2017-04-07 07:08:15.666531 +0300 EEST",
      "id": 16,
      "slug": "my-first-blog-post",
      "text": "The quick brown fox jumps over the lazy dog",
      "title": "My First Blog Post",
      "user_id": 3
    }]
  }
}
```

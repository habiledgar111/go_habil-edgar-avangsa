# JAWABAN UNTUK SOAL FUNDAMENTAL 2

## Gambarkan Select * From Users; dan insert into users; pada beberapa database nosql dibawah


## Redis
**Select**  

HGETALL users  

**Insert**  

MULTI
HMSET users:(id number) name (name user) banknumber (bank number) numberphone (numberphone)
EXEC

## Neo4j
**Select**  

MATCH (ee:Users)
RETURN ee  

**Insert**  

CREATE(ee:User{id : "(id number)", name: "(name user)", banknumber : "(bank number)", numberphone : "(numberphone)"})

## Cassandra
**Select**  

SELECT * FROM Users  

**Insert**  

INSERT into Users (id,name,banknumber,numberphone) 
values ((id number), (name user), (bank number), (numberphone))

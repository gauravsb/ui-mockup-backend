Mongo DB useful commands


Show Databases
--------------
show dbs

Select myDb
------------
use myDb

Show Available Collections
--------------------------
db.getCollectionNames()


Show All user entries
---------------------
db.user.find()

Show All standard entries
---------------------
db.std.find()


Show specifc control under any standard
---------------------------------------
db.std.find({"controls.controlname": "IA-2 (1)"}).pretty();



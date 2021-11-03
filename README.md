# YDB 
## A simple golang database.
***
## Trait:
### 1. Purely programming in golang;
### 2. Natively combine with golang code;
### 3. Natively support with JSON data structure;
### 4. Fully fuctions in CRUD (create, read, update, delete);
### 5. Light weight;
### 6. Simple API Function, easy using；
### 7. Fast and stable (Physical Test);
### 8. Easy be expanded.

***
***
## Some Fuctions Instances:
##### // SetDBItem to use the db line unit key name and it's corresponding value to set a line unit item, the key name of the DB item can not be the same
### func SetDBItem(dbName string, tableName string, lineName string, keyName string, value string) {
***
##### // GetDBItem to use the db line unit key name to get it's corresponding value, the key name of the DB item can not be the same
### func GetDBItem(dbName string, tableName string, lineName string, keyName string) string {
***
##### // DelectDB to delect the whole specific
### func DelectDB(dbName string) {
*** 
##### // DelectTable to delect the whole specific DB Table
### func DelectTable(dbName string, tableName string) {



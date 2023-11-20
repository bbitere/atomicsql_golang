------------------------------------------
HOME
------------------------------------------

<br/>**Atomicsql_golang** is a ORM library for Golang having implemented, beside clasical implementation of any ORM, a special usage for Where() method and Select() method, using literal function aka lambda expression.
<br/>
<br/>All of these tricks are done to have a robust/flexible implementation in your code.
<br/>We have implemented: DataBase First, or Models First, and these are explained here, in a bottom section.

------------------------------------------

<br/> **Simple query interogation** . 
<br/> Let's see this Example:
<br/> var models = ctx.Users.Qry("label1").Where( func(x *m.User) bool{
<br/> &emsp;&emsp;&emsp;   return x.Name == userName}).GetModels();
<br/> 
<br/> In this example, the Where() contains a literal function aka lambda expression. This help the developer to have a robust development and the check of types between data

------------------------------------------
<br/>Also, we have 2 utilities running under .net framework 7.0 (also for linux).
<br/>
<br/>**goscanner.exe**: 
<br/>-compile the code and collects all Models marked to be in Database and generate 1 json file with these definitions, using flag:-e.
<br/>&emsp;   (Step 1 in **Model First**)
<br/>
<br/>-compile the code and generate lambda expressions, using flag:-q
<br/>&emsp;   (Step 5 in **Model First**, but also used in **DataBase First**)
<br/>
<br/>
<br/>**DBTool.exe**:
<br/>
<br/>- extract from jsons definitions of models and generate the sql scripts, using flag: -asql_migration. (Step 2 in **Model First**)
<br/>&emsp;    the sql scripts reflect the incremental updates for Database using the diferences of json files
<br/>
<br/>- apply all sql scripts to update Database, using flag: -migration_db. The directory is located at:
<br/>&emsp;	.\library\tests\test1\_db_migration
<br/>&emsp;	(Step 3 in **Model First**, but also used in **DataBase First**)
<br/><br/>
<br/>- extract directly from DataBase all tables and generate golang models using flag: -export_db
<br/>&emsp;	(Step 4 in **Model First**, but also used in **DataBase First**)<br/>
<br/>
<br/>
<br/>Both these utilities are used in batch files in a example. Its location is directory:
<br/>&emsp;    .\library\tests\test1\build\win32
<br/>
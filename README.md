------------------------------------------

Atomicsql_golang is an ORM library for Golang having implemented, beside classical features of any ORM, a special usage for Where() method and Select() method, using literal function aka lambda expression.
------------------------------------------
Description
------------------------------------------

All of these techniques are employed to ensure a robust/flexible implementation in your code.
We have implemented both: DataBase First, or Models First approaches, and these are explained in a later section.

<br/>**Atomicsql_golang** is a **ORM library** for **Golang** having implemented, beside clasical implementation of any ORM, a special usage for **Where()** method and **Select()** method, using literal function aka lambda expression.
<br/>
<br/>All of these tricks are done to have a **robust/flexible** implementation in your code.
<br/>We have implemented: **DataBase First**, or **Models First**, and these are explained here, in a bottom section.

------------------------------------------

<br/> **Simple query interogation** . 
<br/> Let's see this Example:
<br/>
<br/> var userName = "john1234";
<br/> var models = ctx.Users.Qry("label1").Where( func(x *m.User) bool{
<br/> &emsp;&emsp;&emsp;   return x.Name == userName}).GetModels();
<br/> 
<br/> Compare this kind of instruction with clasical ORM:
<br/> var models = ctx.Users.Where( "Name", "=", userName).GetModels();
<br/> 
<br/> Even the instruction is a little bit shorter, it doesn't check the types between value and field and the field is hardcoded. You don't have any guarantee that this will change appropriately in a refactor.
<br/> 
<br/> The engine translate this instruction in a sql query, using precompile lambda instructions placed inside WHERE() as following: 
<br/> &emsp;SELECT \* FROM users usr 
<br/> &emsp;&emsp;  WHERE usr.Name = @1 
<br/> 
<br/> In this example, the Where() contains a literal function aka lambda expression. This help the developer to have a robust development and the check of types between data

------------------------------------------

<br/> **Using foreign key in Where** 
<br/> Let's see next Example: 
<br/>
<br/> var roleName = "admin";
<br/> var models = ctx.Users.Qry("label2").Where( func(x *m.User) bool{
<br/> &emsp;&emsp;&emsp;   return x.RoleNameID.RoleName == roleName}).
<br/> &emsp;&emsp;     GetModels();
<br/> 
<br/> The engine translate this instruction in a sql query, using precompile lambda instructions placed inside WHERE() as following: 
<br/> &emsp; SELECT \* FROM users usr 
<br/> &emsp;&emsp;  WHERE role.RoleName = @1 
<br/> &emsp;&emsp;  LEFT JOIN user_role role on role.ID = usr.UserRole_ID
<br/> 
<br/> In this example, the Where() make a compare using the FK relation (implicit inner join) and also return the relation as a pointer. Note: ctx._Users.UserRole is the definition of FK table relation.

------------------------------------------

<br/> **Using foreign key in Where** + **get Models having pointer to relation** . 
<br/> Let's see next Example: 
<br/>
<br/> var roleName = "admin";
<br/> var models = ctx.Users.Qry("label2").Where( func(x *m.User) bool{
<br/> &emsp;&emsp;&emsp;   return x.RoleNameID.RoleName == roleName}).
<br/> &emsp;&emsp;     GetModelsRel( ctx._Users.UserRole );
<br/> 
<br/> if( len(models) > 0 && models[0].UserRoleID != nil && models[0].UserRoleID.RoleName == roleName){
<br/> &emsp;// here is showing the usage of fk pointer in code.
<br/> }
<br/> 
<br/> In this example, the Where() make a compare using the FK relation (implicit inner join) and also return the relation as a pointer. Note: ctx._Users.UserRole is the definition of FK table relation.

------------------------------------------

<br/> **Using Select()**. 
<br/>Let's see the an example with Select()
<br/>
<br/>usersAsView, _ := atmsql.Select( ctx.User.Qry("label3").
<br/>&emsp;&emsp;&emsp;                      Where(func(x *m.User) bool {
<br/>&emsp;&emsp;&emsp;&emsp;                      return x.UserRoleID.IsActive == true }),
<br/>&emsp;&emsp;&emsp;                      func(x *m.User) *vUser1 {
<br/>&emsp;&emsp;&emsp;&emsp;                            return &vUser1{ User: *x, UserRole: x.UserRoleID.RoleName, } }).
<br/>&emsp;&emsp;&emsp;                      GetModels()
<br/>
<br/> and struct vUser1 is defined as:
<br/>&emsp;type vUser1 struct {
<br/>&emsp;&emsp;		m.User   \`atomicsql:"copy-model"\`
<br/>&emsp;&emsp;		UserRole string
<br/>&emsp;	}
<br/>
<br/> The engine translate this instruction in a sql query, using precompile lambda instructions placed inside WHERE() + Select() as following: 
<br/> &emsp; SELECT usr.\*, role.RoleName AS User FROM users usr 
<br/> &emsp;&emsp;  WHERE role.IsActive = 1 
<br/> &emsp;&emsp;  LEFT JOIN user_role role on role.ID = usr.UserRole_ID
<br/>
<br/>this code does: return a list with users, but having also the UserRole as the name of RoleName from FK relation of UserRoleID. Here is no need to use GetModelsRel() method as in previous example, to obtain the FK relation.

------------------------------------------

<br/>Also, we have 2 utilities running under .net framework 7.0 (also for linux).
<br/>
<br/>**goscanner.exe**: - compile the golang source code and generate depending by case: a json with model defs or sql query for lambda expresions
<br/>**DBTool.exe**: - do a lot of things related with database or generate golang models files
<br/>
<br/>
<br/>in directory library\tests\test1\build\win32, 
<br/>you can find some scripts and you can generate the models for your project as follows:
<br/>
<br/>Paradigma **DataBase First**: generate models using **1.updateDb_DataBaseFirst.cmd**  in next steps:
<br/> - 1. extract directly from DataBase all tables and generate golang models using **DBTool.exe** with flag: -export_db
<br/> - 2. compile the code and generate sql queries for lambda expressions, using **goscanner.exe** with flag:-q
<br/>
<br/>
<br/>Paradigma: **Model First**: generate models using **1.updateDb_ModelFirst.cmd** in next steps:
<br/>- 1. compile the code and collects all Models marked to be in Database and generate 1 json file with these definitions, using **goscanner.exe** with flag:-e.
<br/>&emsp;&emsp;the marked model should be defined as: type Model struct /\*atomicsql-table:"sqlmodel"\*/{ ... }
<br/>
<br/>- 2. extract from all jsons definitions of models from dir .\library\tests\test1\_db_jsons  and generate the incremental sql scripts, using **DBTool.exe** with flag: -asql_migration. 
<br/>&emsp;    the sql scripts reflect the incremental updates for Database using the diferences of json files from json file to json file.
<br/>
<br/>- 3. apply all sql scripts to update Database, using **DBTool.exe** with flag: -migration_db. The directory is located at:
<br/>&emsp;	.\library\tests\test1\_db_migration
<br/>
<br/>- 4. extract directly from DataBase all tables and generate golang models using **DBTool.exe** with flag: -export_db
<br/>
<br/>- 5. compile the code and generate sql queries for lambda expressions, using **goscanner.exe** with flag:-q
<br/>
<br/>
<br/>Both these utilities are used in batch files in a example. Its location is in directory:
<br/>&emsp;    .\library\tests\test1\build\win32
<br/>
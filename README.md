 Atomicsql_golang is a ORM library for Golang having implemented, beside clasical implementation of any ORM, a special usage for Where() method and Select() method, using literal function aka lambda expression.

All of these tricks are done for having a robust/flexible implementation in your code.
We have implemented: DataBase First, or Models First.

------------------------------------------
Let's see this Example:<br/> 
var models = ctx.Users.Qry("label1").Where( func(x *m.User) bool{<br/> 
&emsp;&emsp;&emsp;   return x.Name == userName}).GetModels();<br/>

In this example, the Where() contains a literal function aka lambda expression. This help the developer to have a robust development and the check of types between data

------------------------------------------
Let's see next Example:<br/> 
var models = ctx.Users.Qry("label2").Where( func(x *m.User) bool{<br/> 
&emsp;&emsp;&emsp;   return x.RoleNameID.RolName == roleName}).<br/> 
&emsp;&emsp;     GetModelsRel( ctx._Users.UserRole );

if( len(models) > 0 && models[0].UserRoleID != nil && models[0].UserRoleID.RoleName == roleName){<br/>
&emsp;// here is showing the usage of fk pointer in code.<br/> 
}

In this example, the Where() make a compare using the FK relation (implicit inner join) and also return the relation as a pointer. Note: ctx._Users.UserRole is the definition of FK table relation.

------------------------------------------
Let's see the an example with Select()<br/>
usersAsView, _ := atmsql.Select( ctx.User.Qry("label3").<br/> 
&emsp;&emsp;&emsp;                      Where(func(x *m.User) bool {<br/> 
&emsp;&emsp;&emsp;&emsp;                      return x.UserRoleID.IsActive == true }),<br/> 
&emsp;&emsp;&emsp;                      func(x *m.User) *vUser1 {<br/> 
&emsp;&emsp;&emsp;&emsp;                            return &vUser1{ User: *x, UserRole: x.UserRoleID.RoleName, } }).<br/>
&emsp;&emsp;&emsp;                      GetModels()

return a list with users, but having also the UserRole as the name of RoleName from FK relation of UserRoleID. Here is no need to use GetModelsRel() method as in previous example, to obtain the FK relation.

------------------------------------------

Also, we have 2 utilities running with .net framework 7.0.<br/>
<br/>
goscanner.exe: <br/>
- compile the code and collects all Models marked to be in Database and generate 1 json file with these definitions, using flag:-e. <br/>
<br/>
- compile the code and generate lambda expressions. (Step 1 in Model First), using flag:-q <br/>
<br/>
<br/>
DBTool.exe:<br/>
<br/>
- extract from jsons definitions of models and generate the sql scripts, using flag: -asql_migration. (Step 2 in Model First)<br/>
    the sql scripts reflect the incremental update of Database using the diferences of json files,<br/>
<br/>
- apply all sql scripts to update Database, using flag: migration_db. The directory is located at:<br/>
	.\library\tests\test1\_db_migration<br/>
	(Step 3 in Model First, but also used in DataBase First)<br/>
<br/>
- extract directly from DataBase all tables and generate golang models using flag: -export_db<br/>
	(Step 4 in Model First, but also used in DataBase First)<br/>
<br/>
<br/>
Both these utilities are used in batch files in a example. Its location is directory:<br/>
    .\library\tests\test1\build\win32<br/>
<br/>
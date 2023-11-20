 Atomicsql_golang is a ORM library for Golang having implemented, beside clasical implementation of any ORM, a special usage for Where() method and Select() method, using literal function aka lambda expression.

All of these tricks are done for having a robust/flexible implementation in your code.
We have implemented: DataBase First, or Models First.

------------------------------------------
Let's see this Example:<br/> 
var models = ctx.Users.Qry("label1").Where( func(x *m.User) bool{<br/> 
&nbsp;&nbsp;&nbsp;   return x.Name == userName}).GetModels();<br/>

In this example, the Where() contains a literal function aka lambda expression. This help the developer to have a robust development and the check of types between data

------------------------------------------
Let's see next Example:<br/> 
var models = ctx.Users.Qry("label2").Where( func(x *m.User) bool{<br/> 
&nbsp;&nbsp;&nbsp;   return x.RoleNameID.RolName == roleName}).<br/> 
&nbsp;&nbsp;     GetModelsRel( ctx._Users.UserRole );

if( len(models) > 0 && models[0].UserRoleID != nil && models[0].UserRoleID.RoleName == roleName){<br/>
&nbsp;// here is showing the usage of fk pointer in code.<br/> 
}

In this example, the Where() make a compare using the FK relation (implicit inner join) and also return the relation as a pointer. Note: ctx._Users.UserRole is the definition of FK table relation.

------------------------------------------
Let's see the an example with Select()<br/>
usersAsView, _ := atmsql.Select( ctx.User.Qry("label3").<br/> 
&nbsp;&nbsp;&nbsp;                      Where(func(x *m.User) bool {<br/> 
&nbsp;&nbsp;&nbsp;&nbsp;                      return x.UserRoleID.IsActive == true }),<br/> 
&nbsp;&nbsp;&nbsp;                      func(x *m.User) *vUser1 {<br/> 
&nbsp;&nbsp;&nbsp;&nbsp;                            return &vUser1{ User: *x, UserRole: x.UserRoleID.RoleName, } }).<br/>
&nbsp;&nbsp;&nbsp;                      GetModels()

return a list with users, but having also the UserRole as the name of RoleName from FK relation of UserRoleID. Here is no need to use GetModelsRel() method as in previous example, to obtain the FK relation.
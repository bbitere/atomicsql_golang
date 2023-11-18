Atomicsql_golang is a ORM library for Golang having implemented, beside clasical implementation of any ORM, a special usage for Where() method and Select() method, using literal function aka lambda expression.

All of these tricks are done for having a robust/flexible implementation in your code.
We have implemented: DataBase First, or Models First.
------------------------------------------
Let's see this Example: var models = ctx.Users.Qry("label1").Where( func(x *m.User) bool{ return x.Name == userName}).GetModels();

in this example, the Where() contain a literal function aka lambda expression. This help the developer to have a robust development and having the check of types between data
------------------------------------------
Let's see next Example: var models = ctx.Users.Qry("label2").Where( func(x *m.User) bool{ return x.RoleNameID.RolName == roleName}). GetModelsRel( ctx._Users.UserRole );

if( len(models) > 0 && models[0].UserRoleID != nil && models[0].UserRoleID.RoleName == roleName){ // here is showing the usage of fk pointer in code. }

in this example, the Where() make a compare using the FK relation (implicit inner join) and also return the relation as a pointer. Note: ctx._Users.UserRole is the definition of FK table relation.
------------------------------------------
Let's see the an example with Select()

usersAsView, _ := atmsql.Select( ctx.User.Qry("label3"). Where(func(x *m.User) bool { return x.UserRoleID.IsActive == true }), func(x *m.User) *vUser1 { return &vUser1{ User: *x, UserRole: x.UserRoleID.RoleName, } }).GetModels()

return a list with users, but having also the UserRole as the name of RoleName from FK relation of UserRoleID. Here is no need to use GetModelsRel() method as in previous example.
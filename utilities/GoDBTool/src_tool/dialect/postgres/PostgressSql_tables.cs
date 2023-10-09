using System;
using System.Collections.Generic;
using System.Linq;
using Npgsql;

namespace src_tool
{
    partial class PostgresqlDialect
    {
        public override Dictionary<string, DbTable> readTables( string SqlLang )
        {
            var tables = new Dictionary<string, DbTable>();
            var cnn = this.connection;
            var packageImports = new Dictionary<string,string>();

            var requestText = $@"
            select
                t.table_name as table1, 
                t.table_schema  as schema1,
	            c.column_name as column1,
				format('%s',udt_name::regtype) as colType2,
				c.data_type  as colType1,	
	            c.is_nullable  as isNullable1,
				c.ordinal_position  as ordinal_position1,
	            /*c.is_identity  as is_identity1,*/
				( 
				SELECT 'YES'
						FROM information_schema.table_constraints tc1 
						JOIN information_schema.constraint_column_usage AS ccu USING (constraint_schema, constraint_name) 
						JOIN information_schema.columns AS c1 ON c1.table_schema = tc1.constraint_schema
						  AND tc1.table_name = c1.table_name AND ccu.column_name = c1.column_name
						WHERE constraint_type = 'PRIMARY KEY' 
					    and tc1.table_name = t.table_name
						and c1.column_name = c.column_name
						limit 1 
				
				) as is_identity1
				
	 
            from
                information_schema.columns c
            inner join information_schema.tables t on
                    t.table_name = c.table_name

            where
                t.table_schema = 'public'
                and t.table_type= 'BASE TABLE'
                and c.table_schema = 'public'
            group by t.table_name, t.table_schema, 
            c.column_name, udt_name::regtype, c.data_type, 
            c.is_nullable, c.is_identity,
			c.ordinal_position
			order by c.ordinal_position 
			;
            ";

            //var cmd = cnn.CreateCommand();
            using( var cmd = new NpgsqlCommand( requestText, cnn))
            {
                //cmd.CommandText = requestText;
                //cmd.ExecuteReader();requestText);
                var reader = cmd.ExecuteReader();

                if (reader.HasRows)
                {
                    DbTable table = null;
                    while (reader.Read())
                    {
                        var tableName   = getString(reader, 0);
                        var tableSchema = getString(reader, 1);

                        table = tables.ContainsKey(tableName)
                              ? tables[tableName]:new DbTable();

                        table.LangTableNameModel = GoModelTemplate.ConvertToIdent_GoLang( tableName );
                        table.Schema = tableSchema;
                        table.SqlTableNameModel = tableName;
                        tables[ table.SqlTableNameModel ] = table;
                        
                        var columnName      = getString(reader, 2);
                        if( columnName.Contains( "Opened" ))
                            Utils.Nop();
                        if( columnName == "\"OpenedItems\"" )
                            Utils.Nop();
                        var colType         = getString(reader, 3);
                        var colType2        = getString(reader, 4);
                        var colIsNullable   = getString(reader, 5);
                        //var colIsOrder      = getInteger(reader, 6);
                        var colIsIdentity   = getString(reader, 7);

                        var column  = new DbColumn();
                        column.langName  = GoModelTemplate.ConvertToIdent_GoLang( columnName, true );
                        column.langName2 = GoModelTemplate.ConvertToIdent_GoLang( columnName, false );
                                            
                        column.sqlName = columnName;
                        column.sqlType = colType;
                        column.langType =  this.getGoLangType( column, ref packageImports );
                        column.bIsIdentity = isYes(colIsIdentity);
                        column.bIsNullable = isYes(colIsNullable);
                        table.columns.Add( column );

                        if( column.bIsIdentity )
                        {
                            if( table.PrimaryColumn != null ) 
                            {
                                Console.WriteLine($"{table} has already a primary key");
                                return null;
                            }
                            table.PrimaryColumn = column;
                        }

                        //var column          = reader.GetString(6);
                        //Console.WriteLine("{0}\t{1}", reader.GetInt32(0),reader.GetString(1));
                    }
                    if( table == null )
                    {
                        printError($"table no loaded");
                        return null;
                    }
                    if( table.PrimaryColumn == null )
                    {
                        printError($"table {table.LangTableNameModel} doesnt have a primary key");
                        return null;
                    }
                    reader.Close();
                
                    return tables;
                }else
                {
                    Console.WriteLine("No rows found. Maybe, The user has no rights to read DB defs");
                }
            }
            return null;
        }

        public override bool readConstraintors( Dictionary<string, DbTable> tables, string dirJsons  )
        {
            var cnn = this.connection;
            var requestText = $@"
            select *FROM (

                    select
                        /*pgc.contype as constraint_type,*/
                        ccu.table_schema as table_schema,
                        kcu.table_name as table_name,
                        case when (pgc.contype = 'f') then kcu.column_name else ccu.column_name end as column_name, 
                        case when (pgc.contype = 'f') then ccu.table_name else (null) end as reference_table,
                        case when (pgc.contype = 'f') then ccu.column_name else (null) end as reference_col,
                        case when (pgc.contype = 'p') then 'yes' else 'no' end as auto_inc,
                        case when (pgc.contype = 'p') then 'no' else 'yes' end as is_nullable,
                        'integer' as data_type,
                        '0' as numeric_scale,
                        '32' as numeric_precision
                    from
                        pg_constraint as pgc
                        join pg_namespace nsp on nsp.oid = pgc.connamespace
                        join pg_class cls on pgc.conrelid = cls.oid
                        join information_schema.key_column_usage kcu on kcu.constraint_name = pgc.conname
                        left join information_schema.constraint_column_usage ccu on pgc.conname = ccu.constraint_name 
                        and nsp.nspname = ccu.constraint_schema
	                and ccu.constraint_schema = 'public'
                     union
                        select 
                            /*null as constraint_type ,*/
                            table_schema,
                            table_name,
                            column_name, 
                            null as refrence_table, 
                            null as refrence_col, 
                            'no' as auto_inc,
                            is_nullable,
                            data_type,
                            numeric_scale,
                            numeric_precision
                        from information_schema.columns cols 
                        where 
                            table_schema = 'public'
                            and concat(table_name, column_name) not in(
                                select concat(kcu.table_name, kcu.column_name)
                                from
                                pg_constraint as pgc
                                join pg_namespace nsp on nsp.oid = pgc.connamespace
                                join pg_class cls on pgc.conrelid = cls.oid
                                join information_schema.key_column_usage kcu on kcu.constraint_name = pgc.conname
                                left join information_schema.constraint_column_usage ccu on pgc.conname = ccu.constraint_name 
                                and nsp.nspname = ccu.constraint_schema
                            )
                    ) as foo
	                where foo.table_schema is not null
                order by table_name asc, column_name

            ";

            //var cmd = cnn.CreateCommand();
            using( var cmd = new NpgsqlCommand( requestText, cnn))
            {
                //cmd.CommandText = requestText;
                //cmd.ExecuteReader();requestText);
                var reader = cmd.ExecuteReader();

                if (reader.HasRows)
                {
                    while (reader.Read())
                    {
                        var schemaName         = getString(reader, 0);
                        var tableName          = getString(reader, 1);

                        var columnName      = getString(reader, 2);
                        var colRefTable     = getString(reader, 3);
                        var colRefColumn    = getString(reader, 4);
                        //var colIsNullable   = reader.GetString(4);
                        //var colIsIdentity   = reader.GetString(5);

                        if( colRefTable != null && colRefColumn != null )
                        if( tables.ContainsKey(tableName) )
                        {
                            var table = tables[tableName];
                            var column = table.columns.Where(x=>x.sqlName == columnName).FirstOrDefault();
                            if(column != null )
                            {
                                if( tables.ContainsKey(colRefTable) )
                                {
                                    var primary = tables[colRefTable].PrimaryColumn;
                                    if( primary != null && primary.langName == colRefColumn)
                                    {
                                        column.ForeignKey = tables[colRefTable];
                                    }else
                                    {
                                        printError($"Foreign key from table {tableName} to {colRefTable} is not pointed to primary key");
                                        return false;
                                    }
                                }else
                                {
                                    printError($"not found table {colRefTable}");
                                    return false;
                                }
                            }
                        }else
                        {
                            printError($"not found table {tableName}");
                            return false;
                        }

                        //var column          = reader.GetString(6);
                        //Console.WriteLine("{0}\t{1}", reader.GetInt32(0),reader.GetString(1));
                    }
                    reader.Close();
                    return true;
                }else
                {
                    printError("No rows found. Maybe, The user has no rights to read DB defs");
                }
            }
            return false;
        }

        private string getString(NpgsqlDataReader reader, int index)
        {
            if( !reader.IsDBNull(index) )
                return reader.GetString( index );
            return null;
        }
        

        bool isYes(string s)
        {
            return s != null && s == "YES";
        }
    }
}

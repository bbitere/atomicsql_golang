using System;
using System.Collections.Generic;
using System.Linq;
using Npgsql;

namespace src_tool
{
    partial class MongodbDialect
    {
        public override Dictionary<string, DbTable> readTables( string SqlLang )
        {
            return new Dictionary<string, DbTable>();
        }

        public override bool readConstraintors( Dictionary<string, DbTable> tables  )
        {
            return true;
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

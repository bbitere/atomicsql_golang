using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace src_tool
{
    public class GenericDialectArg
    {
        
        public string connection_string;
        public GenericDialectArg(string connection_string)
        {
            this.connection_string = connection_string;
        }
    }
}

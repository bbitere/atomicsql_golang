using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Antlr4.Runtime;

public partial class GoParser{
    public partial class PrimaryExprContext {

        public string m_funcMethodName=null;
        public string m_classMethod;
        public bool m_bIsNoSql = false;

        public void setDataPrimaryExpression( 
            string className, string methodName, bool? bIsNoSql)
        {
            if( false
                || this.UID == 96918
                || this.UID == 97770 )
                Debugger.Break();

            this.m_classMethod      = className;
            this.m_funcMethodName   = methodName;
            this.m_bIsNoSql         = bIsNoSql??false;
        }
    }
}

namespace Antlr4.Runtime
{
    

}


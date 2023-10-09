using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using goscanner.Metadata;

namespace goscanner;

public partial class ScannerBase
{
    protected TypeInfo new_TypeInfoByName(string typeName)
    {
        var v = new TypeInfo{ 
                    Name = typeName,
                    FullTypeName = typeName,
                    TypeClass = TypeClass.Simple,
                };
        return v;
    }

    protected ParameterInfo new_ParameterInfo(string nameVar, string typeName)
    {
        var v = new ParameterInfo{ 
                    Name = nameVar, 
                    Type = new_TypeInfoByName( typeName )
                };
        return v;
    }
    protected ParameterInfo new_ParameterInfo(string nameVar, TypeInfo typeinfo)
    {
        var v = new ParameterInfo{ 
                    Name = nameVar, 
                    Type = typeinfo 
                };
        return v;
    }
}

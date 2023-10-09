using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace goscanner.Metadata
{
    public class Identif
    {
        private string _name;
        public string Name{ 
            get{ return _name;}
            set{
                _name = value;
                //if( value == "User_Model" )
                //    Debugger.Break();
                } 
            }
        private string package;
        public string getPackage(){ return package;}
        public void setPackage(string n){ package = n;}
    }
}

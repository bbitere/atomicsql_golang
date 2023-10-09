

    /* this class is generated automatically by GoServerTool exporter*/

    package newton_models
        
        import (
	    uuid "github.com/google/uuid"
        )

    type User struct {
	        Generic_MODEL
	        ID                  int32                         `json:"id,omitempty"`
            UUID                uuid.uuid                     `json:"uuid"`
            UserName            string                        `json:"username"`
            UserPsw             string                        `json:"userpsw"`
            UserRole_ID         *UserRole                     `json:"userrole_id"`
            misc1               []string                      `json:"misc1"`
    }

    func (model  User) GetID() string {
	    return model.ID
    }

            
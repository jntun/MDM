package market

import (
    "fmt"

    "github.com/Nastyyy/mdm-back/config"
)

type AdminAction struct {
    uuid string 
}

func (act AdminAction) DoAction(sess *Session) error {
    config.VerboseLog(fmt.Sprintf("[ADMIN] Attempting match with %s - %s", sess.Admin.UUID.String(), act.uuid))
    if sess.Admin.UUID.String() == act.uuid {
        config.DebugLog(fmt.Sprintf("Admin found: %s", act.uuid))
        return nil
    }

    return fmt.Errorf("[ADMIN] Unauthorized request from %s", act.uuid)
}


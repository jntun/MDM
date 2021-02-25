package market

import (
    "fmt"

    "github.com/Nastyyy/mdm-back/config"
)

type AdminAction struct {
}

func (act AdminAction) DoAction(sess *Session, usr *User) error {
    config.VerboseLog(fmt.Sprintf("[ADMIN] Attempting match with %s - %s", sess.Admin.UUID.String(), usr.UUID))
    if sess.Admin.UUID.String() == usr.UUID.String() {
        config.DebugLog(fmt.Sprintf("Admin found: %s", usr.UUID))
        return nil
    }

    return fmt.Errorf("[ADMIN] Unauthorized request from %s", usr.UUID)
}

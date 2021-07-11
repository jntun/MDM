package market

import (
    "errors"
    "fmt"
    "github.com/Nastyyy/mdm-back/config"
    "reflect"
)

type AdminAction struct {
    Body interface{} `json:"body"`
}

func (act AdminAction) DoAction(sess *Session, usr *User) error {
    config.VerboseLog(fmt.Sprintf("[ADMIN] Attempting match with %s - %s", sess.Admin.UUID.String(), usr.UUID))

    if sess.Admin.Name == "default-admin" {
    	config.DebugLog(fmt.Sprintf("No admin found, setting to %s..", usr.Name))
        sess.Admin = usr
    }

    if sess.Admin.UUID.String() == usr.UUID.String() {
        //config.DebugLog(fmt.Sprintf("Admin found: %s |\nBody: %s", usr.UUID, act.Body))
        val, found := act.Body.(map[string]interface{})["admin-action"]
        if !found {
            return errors.New("No 'admin-action' field found.")
        }

        if reflect.TypeOf(val).Kind() != reflect.String {
            return errors.New("Invalid 'admin-action' command.")
        }

        switch act.Body.(map[string]interface{})["admin-action"].(string) {
        case "pause":
            sess.Game.Stop()
        case "resume":
            sess.Game.Start()
        }
        return nil
    }

    return fmt.Errorf("[ADMIN] Unauthorized request from %s", usr.UUID)
}
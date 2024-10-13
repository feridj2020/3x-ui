
package service

import (
    "encoding/json"
    "fmt"
    "strconv"
    "strings"
    "time"

    "x-ui/database"
    "x-ui/database/model"
    "x-ui/logger"
    "x-ui/util/common"
    "x-ui/xray"

    "gorm.io/gorm"
)

type InboundService struct {
    xrayApi xray.XrayAPI
}

// Updated GetInbounds function with sorting feature
func (s *InboundService) GetInbounds(userId int, sortField string) ([]*model.Inbound, error) {
    db := database.GetDB()
    var inbounds []*model.Inbound
    query := db.Model(model.Inbound{}).Preload("ClientStats").Where("user_id = ?", userId)

    // Adding sorting based on the specified field
    switch sortField {
    case "expiryTime":
        query = query.Order("expiry_time")
    case "addedDate":
        query = query.Order("created_at")
    default:
        query = query.Order("id") // Default sorting by ID
    }

    err := query.Find(&inbounds).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    return inbounds, nil
}

func (s *InboundService) GetAllInbounds() ([]*model.Inbound, error) {
    db := database.GetDB()
    var inbounds []*model.Inbound
    err := db.Model(model.Inbound{}).Preload("ClientStats").Find(&inbounds).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil, err
    }
    return inbounds, nil
}

// Rest of the code remains unchanged


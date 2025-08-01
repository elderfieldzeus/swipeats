package repositories

import (
	"github.com/SwipEats/SwipEats/server/internal/db"
	"github.com/SwipEats/SwipEats/server/internal/models"
	"gorm.io/gorm"
)

func AddUserToGroup(userID uint, groupID uint, isOwner bool) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	membership := &models.GroupMembership{
		UserID: userID,
		GroupID: groupID,
		IsOwner: isOwner,
	}

	result := db.Conn.Create(membership)
	if result.Error != nil {
		return result.Error // Error creating group membership
	}
	return nil // Group membership created successfully
}

func GetGroupMembershipByUserIDAndGroupID(userID uint, groupID uint) (*models.GroupMembership, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var membership models.GroupMembership
	result := db.Conn.Where("user_id = ? AND group_id = ? AND deleted_at IS NULL", userID, groupID).First(&membership)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Membership not found
		}
		return nil, result.Error // Other error
	}
	return &membership, nil
}

func GetGroupMembershipByUserIDAndGroupIDWithDeleted(userID uint, groupID uint) (*models.GroupMembership, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var membership models.GroupMembership
	result := db.Conn.Unscoped().Where("user_id = ? AND group_id = ?", userID, groupID).First(&membership)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Membership not found
		}
		return nil, result.Error // Other error
	}
	return &membership, nil
}

func GetGroupMembershipsByGroupID(groupID uint) ([]models.GroupMembership, error) {
	if db.Conn == nil {
		return nil, gorm.ErrInvalidDB // Database connection is not established
	}

	var memberships []models.GroupMembership
	result := db.Conn.Where("group_id = ? AND deleted_at IS NULL", groupID).Find(&memberships)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No memberships found
		}
		return nil, result.Error // Other error
	}
	return memberships, nil
}

func GetMemberCountByGroupID(groupID uint) (int64, error) {
	if db.Conn == nil {
		return 0, gorm.ErrInvalidDB // Database connection is not established
	}

	var count int64
	result := db.Conn.Model(&models.GroupMembership{}).Where("group_id = ? AND deleted_at IS NULL", groupID).Count(&count)

	if result.Error != nil {
		return 0, result.Error // Error counting memberships
	}
	return count, nil // Return the count of members
}

func UpdateGroupMembership(membership *models.GroupMembership) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Save(membership)
	if result.Error != nil {
		return result.Error // Error updating group membership
	}
	return nil // Group membership updated successfully
}

func RemoveUserFromGroup(membership *models.GroupMembership) error {
	if db.Conn == nil {
		return gorm.ErrInvalidDB // Database connection is not established
	}

	result := db.Conn.Delete(membership)
	if result.Error != nil {
		return result.Error // Error deleting group membership
	}
	return nil // Group membership deleted successfully
}
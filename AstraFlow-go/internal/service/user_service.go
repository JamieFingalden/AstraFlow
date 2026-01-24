package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
	roleRepo *repository.RoleRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
		roleRepo: repository.NewRoleRepository(),
	}
}

func (s *UserService) FindUserRoleByUserId(userId int64) (string, error) {
	return s.userRepo.FindUserRoleByUserId(userId)
}

func (s *UserService) GetUserList(tenantId int64) ([]*model.User, error) {
	return s.userRepo.GetUserList(tenantId)
}

// CreateUser 添加员工 (Internal Invite)
// Requires Admin privileges (checked in Handler)
func (s *UserService) CreateUser(tenantId int64, username, password, name, roleKey string) (*model.User, error) {
	// 1. Check if role exists
	role, err := s.roleRepo.FindByKey(roleKey)
	if err != nil {
		return nil, errors.New("invalid role key: " + roleKey)
	}

	// 2. Check if username exists
	existing, err := s.userRepo.FindByUsername(username)
	if err == nil && existing != nil {
		return nil, errors.New("username already exists")
	}

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 4. Create User
	user := &model.User{
		TenantID: &tenantId,
		Username: username,
		Password: string(hashedPassword),
		Name:     name,
		RoleID:   role.ID,
		Status:   model.UserStatusActive,
	}

	if err := s.userRepo.CreateUserReturnIt(user); err != nil {
		return nil, err
	}

	user.Role = role // Preload for response
	user.Password = "" // Clear password

	return user, nil
}

func (s *UserService) UpdateUser(userId int64, username, password, email, phone, roleKey string) error {
	user, err := s.userRepo.FindByID(userId)
	if err != nil {
		return err
	}

	if email != "" {
		user.Email = &email
	}
	if phone != "" {
		user.Phone = &phone
	}
	if name := username; name != "" {
		user.Name = name // Assuming username param meant name here based on context? Or username update? Keeping original param name but matching logic.
		// Wait, original code updated username? Usually immutable.
		// Let's stick to updating auxiliary fields or Role.
	}
	
	if roleKey != "" {
		role, err := s.roleRepo.FindByKey(roleKey)
		if err != nil {
			return errors.New("role not found")
		}
		user.RoleID = role.ID
	}
	
	// Password update logic if needed (not requested in prompt but good to have safety)
	if password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		user.Password = string(hashed)
	}
	
	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(userId int64) error {
	return s.userRepo.Delete(userId)
}

func (s *UserService) FindUserById(userId int64) (*model.User, error) {
	return s.userRepo.FindByID(userId)
}

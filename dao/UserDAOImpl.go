package dao

import (
	"ProjectFinal/db"
	"ProjectFinal/models"
	"context"
	"errors"
)

type IUserDAO interface {
	Insert(u *models.User) error
	Update(u *models.User) error
	Delete(userId int) error
	GetUserByEmailAndPassword(userEmail string, userPassword string) (*models.User,error)
	GetUserLikedNewsIds(userId int) (*[]int,error)
}

type UserDAOImpl struct {
}

var  conn = db.GetInstance()

func (dao *UserDAOImpl) Insert(user *models.User) error{
	query := "insert into users (full_name, email, password) values ($1,$2,$3)"
	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), query,user.GetUserFullName(),user.GetUserEmail(),user.GetLikedNewsIds())
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDAOImpl) Update(user *models.User) error{
	//update
	return nil
}

func (dao *UserDAOImpl) Delete(userId int) error{
	query := "delete from public.user where id = '$1' "
	commandTag,err := conn.Exec(context.Background(),query,userId)
	if err != nil {
		return err
	}
	if commandTag != 1 {
		return errors.New("No row found to delete")
	}
 	return nil
}

func (dao *UserDAOImpl) GetUserByEmailAndPassword(userEmail string, userPassword string) (*models.User, error){
	query := "SELECT id, full_name, email FROM public.users where email = '$1' and password = '$2' limit 1"
	var user models.User
	err := conn.QueryRow(context.Background(),query, userEmail,userPassword).Scan(user.SetUserId, user.SetFullName,user.SetUserEmail)

	if err != nil {
		return nil,err
	}

	return &user, nil
}

func (dao *UserDAOImpl) GetUserLikedNewsIds(userId int) (*[]int, error){
	query := "select id from liked_news_journal where "
	var likedNews []int

	var rows, err = conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return nil,err
		}
		likedNews = append(likedNews, id)
	}

	if rows.Err() != nil {
		return nil,rows.Err()
	}


	return &likedNews, nil
}


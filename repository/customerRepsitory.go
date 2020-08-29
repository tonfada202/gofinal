package repository

import (
	"database/sql"
	"fmt"

	"github.com/tonfada202/gofinal/models"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) QueryGetAllCustomer() ([]models.CustomerModel, error) {
	customers := []models.CustomerModel{}
	stmt, err := h.DB.Prepare("SELECT * FROM customer")
	if err != nil {
		return customers, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return customers, err
	}

	for rows.Next() {
		var id int
		var name, email, status string

		err := rows.Scan(&id, &name, &email, &status)
		if err != nil {
			fmt.Println("can't scan", err)
			return customers, err
		}
		cus := models.CustomerModel{id, name, email, status}
		customers = append(customers, cus)
	}
	return customers, err
}

func (h *Handler) QueryGetCustomerById(idParam int) (models.CustomerModel, error) {
	customer := models.CustomerModel{}
	stmt, err := h.DB.Prepare("SELECT * FROM customer where id=$1")
	if err != nil {
		return customer, err
	}

	row := stmt.QueryRow(idParam)
	if err != nil {
		return customer, err
	}
	var id int
	var name, email, status string
	err = row.Scan(&id, &name, &email, &status)
	if err != nil {
		fmt.Println("can't scan", err)
		return customer, err
	}
	customer = models.CustomerModel{id, name, email, status}
	return customer, err
}

func (h *Handler) QuerySaveCustomer(cusRq models.CustomerModel) (models.CustomerModel, error) {
	customer := models.CustomerModel{}
	row := h.DB.QueryRow("INSERT INTO customer (name, email , status) values ($1, $2 , $3)  RETURNING id ,name , email , status", cusRq.Name, cusRq.Email, cusRq.Status)
	var id int
	var name, email, status string
	err := row.Scan(&id, &name, &email, &status)
	if err != nil {
		fmt.Println("can't scan", err)
		return customer, err
	}
	customer = models.CustomerModel{id, name, email, status}
	return customer, err
}

func (h *Handler) QueryUpdateCustomerById(idParam int, cusRq models.CustomerModel) (models.CustomerModel, error) {
	customer := models.CustomerModel{}
	row := h.DB.QueryRow("UPDATE customer SET name = $2 , email = $3  , status = $4  WHERE id = $1 RETURNING id ,name , email , status", idParam, cusRq.Name, cusRq.Email, cusRq.Status)
	var id int
	var name, email, status string
	err := row.Scan(&id, &name, &email, &status)
	if err != nil {
		fmt.Println("can't scan", err)
		return customer, err
	}
	customer = models.CustomerModel{id, name, email, status}
	return customer, err
}

func (h *Handler) QueryDelCustomerById(idParam int) error {
	stmt, err := h.DB.Prepare("DELETE FROM customer WHERE id=$1")
	_, err = stmt.Exec(idParam)
	if err != nil {
		return err
	}
	return err
}

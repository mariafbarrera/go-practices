package models

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// palabraReservada,
//(variable para user la estructura en el contexto, estructura a la cual quieres hacer referencia),
//nombreDelaFuncion(),
//retorno de la funcion

//func (u User) GetName() string {}

func (u User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}